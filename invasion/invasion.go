package invasion

import (
	"fmt"

	"github.com/harry-hov/alien-invasion/utils"
	"github.com/harry-hov/alien-invasion/worldmap"
)

const maxMoves = 10000

type Conclusion string

type Invasion struct {
	worldMap   *worldmap.WorldMap
	move       int
	finished   bool
	conclusion Conclusion
}

// GetRelease returns WorldMap
func (i *Invasion) GetWorldMap() *worldmap.WorldMap {
	return i.worldMap
}

// GetRelease returns current move
func (i *Invasion) GetCurrentMove() int {
	return i.move
}

// GetRelease returns the Conclusion of invasion
func (i *Invasion) Conclusion() Conclusion {
	return i.conclusion
}

// SetReleaseCount sets the WorldMap in the invasion
func (i *Invasion) SetWorldMap(wm *worldmap.WorldMap) {
	i.worldMap = wm
}

// SetReleaseCount sets the current move of the invasion
func (i *Invasion) SetMove(move int) {
	i.move = move
}

// SetReleaseCount sets the finished in invasion
func (i *Invasion) SetFinished(val bool) {
	i.finished = val
}

// SetReleaseCount sets the conclusion on the invasion
func (i *Invasion) SetConclusion(c Conclusion) {
	i.conclusion = c
}

// InitInvasion Unleases aliens on WorldMap and returns Invasion
func InitInvasion(worldMap *worldmap.WorldMap, aliens uint) *Invasion {
	invasion := &Invasion{
		worldMap: worldMap,
		move:     0,
	}
	invasion.worldMap.UnleaseNAliens(aliens)
	return invasion
}

// MakeMove reallocate aliens to random connected city
// and increment the current move count
func (i *Invasion) MakeMove() {
	i.worldMap.RandWalkAlien()
	i.move++
}

// IsFinished checks if invasion is finished
// If finished, it sets conclusion and return
func (i *Invasion) IsFinished() bool {
	if i.move >= maxMoves {
		i.finished = true
		i.conclusion = Conclusion("exceeds maximum moves")
		return i.finished
	}
	if i.worldMap.GetCities() == nil {
		i.finished = true
		i.conclusion = Conclusion("all cities destroyed")
		return i.finished
	}

	aliens := i.worldMap.GetAlienList()
	if aliens == nil {
		i.finished = true
		i.conclusion = Conclusion("all aliens died")
		return i.finished
	}
	if len(aliens) == 1 {
		i.finished = true
		i.conclusion = Conclusion(fmt.Sprintf("alien (%v) won", aliens[0]))
		return i.finished
	}

	if trappedAliens := i.worldMap.GetTrappedAlienCount(); (uint(len(aliens)) - trappedAliens) == 0 {
		i.finished = true
		i.conclusion = Conclusion("all aliens trapped")
		return i.finished
	}

	return i.finished
}

// Fight makes the aliens fight if city has more than 2 aliens.
// In process, destroys city and Kill aliens on the destroyed city.
func (i *Invasion) Fight() {
	for city, aliens := range i.worldMap.GetAliensByCity() {
		if len(aliens) >= 2 {
			i.worldMap.DestroyCity(city)
			i.worldMap.KillAliens(aliens)

			fmt.Println(fmt.Sprintf("%v has been destroyed by %v!", city, utils.PrettyJoinAliens(aliens)))
		}
	}
}
