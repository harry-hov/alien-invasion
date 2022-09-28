package invasion

import "github.com/harry-hov/alien-invasion/worldmap"

const maxMoves = 10000

type Invasion struct {
	worldMap *worldmap.WorldMap
	move     int
}

func (i *Invasion) GetWorldMap() *worldmap.WorldMap {
	return i.worldMap
}

func (i *Invasion) GetCurrentMove() int {
	return i.move
}

func InitInvasion(worldMap *worldmap.WorldMap, aliens uint) *Invasion {
	invasion := &Invasion{
		worldMap: worldMap,
		move:     0,
	}
	invasion.worldMap.UnleaseNAliens(aliens)
	return invasion
}
