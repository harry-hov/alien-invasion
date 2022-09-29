package invasion

import (
	"fmt"

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

func (i *Invasion) GetWorldMap() *worldmap.WorldMap {
	return i.worldMap
}

func (i *Invasion) GetCurrentMove() int {
	return i.move
}

func (i *Invasion) Conclusion() Conclusion {
	return i.conclusion
}

func InitInvasion(worldMap *worldmap.WorldMap, aliens uint) *Invasion {
	invasion := &Invasion{
		worldMap: worldMap,
		move:     0,
	}
	invasion.worldMap.UnleaseNAliens(aliens)
	return invasion
}

func (i *Invasion) MakeMove() {
	i.worldMap.RandWalkAlien()
	i.move++
}

func (i *Invasion) IsFinished() bool {
	if i.move >= maxMoves {
		i.finished = true
		i.conclusion = Conclusion("exceeds maximum moves")
	}
	if i.worldMap.GetCities() == nil {
		i.finished = true
		i.conclusion = Conclusion("all cities destroyed")
	}

	aliens := i.worldMap.GetAlienList()
	if aliens == nil {
		i.finished = true
		i.conclusion = Conclusion("all aliens died")
	}
	if len(aliens) == 1 {
		i.finished = true
		i.conclusion = Conclusion(fmt.Sprintf("alien (%v) won", aliens[0]))
	}

	if trappedAliens := i.worldMap.GetTrappedAlienCount(); (uint(len(aliens)) - trappedAliens) == 0 {
		i.finished = true
		i.conclusion = Conclusion("all aliens trapped")
	}

	return i.finished
}
