package worldmap_test

import (
	"strings"
	"testing"

	"github.com/harry-hov/alien-invasion/worldmap"
	"github.com/stretchr/testify/assert"
)

const (
	worldMapInput string = `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee`
	validWorldMapInput     string = `Foo north=Bar west=Baz`
	invalidWorldMapInput   string = `Foo`
	invalidDirectionInput1 string = `Foo northBar west=Baz`
	invalidDirectionInput2 string = `Foo north=Bar wset=Baz`
)

func TestIsValid(t *testing.T) {
	assert.False(t, worldmap.Direction("invalid").IsValid())
	assert.True(t, worldmap.Direction("north").IsValid())
}

func TestGetOpposite(t *testing.T) {
	validDirection := worldmap.North
	oppositeValidDirection, err := validDirection.GetOpposite()
	assert.Nil(t, err)
	assert.Equal(t, worldmap.South, oppositeValidDirection)
	invalidDirection := worldmap.Direction("invalid")
	_, err = invalidDirection.GetOpposite()
	assert.NotNil(t, err)
}

func TestInitWorldMap(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(validWorldMapInput))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(worldMap.GetCities()))
	_, err = worldmap.InitWorldMap(strings.NewReader(invalidWorldMapInput))
	assert.NotNil(t, err)
	_, err = worldmap.InitWorldMap(strings.NewReader(invalidDirectionInput1))
	assert.NotNil(t, err)
	_, err = worldmap.InitWorldMap(strings.NewReader(invalidDirectionInput2))
	assert.NotNil(t, err)
}

func TestAddCity(t *testing.T) {
	worldMap := worldmap.New()
	assert.Equal(t, 0, len(worldMap.GetCities()))
	worldMap.AddCity("Foo")
	assert.Equal(t, 1, len(worldMap.GetCities()))
	assert.Equal(t, worldmap.City("Foo"), worldMap.GetCities()[0])
	worldMap.AddCity("Foo")
	assert.Equal(t, 1, len(worldMap.GetCities()))
}

func TestAddCityE(t *testing.T) {
	worldMap := worldmap.New()
	assert.Equal(t, 0, len(worldMap.GetCities()))
	assert.Nil(t, worldMap.AddCityE("Foo"))
	assert.Equal(t, 1, len(worldMap.GetCities()))
	assert.Equal(t, worldmap.City("Foo"), worldMap.GetCities()[0])
	assert.NotNil(t, worldMap.AddCityE("Foo"))
}

func TestAppendCityDirection(t *testing.T) {
	worldMap := worldmap.New()
	worldMap.AddCity("Foo")
	assert.NotNil(t, worldMap.AppendCityDirection("Foo", "Foo", worldmap.North))
	assert.Nil(t, worldMap.AppendCityDirection("Foo", "Bar", worldmap.North))
}

func TestGetCities(t *testing.T) {
	worldMap := worldmap.New()
	worldMap.AddCity("Foo")
	assert.Equal(t, []worldmap.City{"Foo"}, worldMap.GetCities())
	assert.Equal(t, 1, len(worldMap.GetCities()))
}

func TestGetCityDirections(t *testing.T) {
	worldMap := worldmap.New()
	worldMap.AddCity("Foo")
	assert.Equal(t, []worldmap.City{"Foo"}, worldMap.GetCities())
	assert.Equal(t, 1, len(worldMap.GetCities()))
}

func TestGetConnectedCities(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(worldMap.GetConnectedCities(worldmap.City("Foo"))))
}

func TestUnleaseNAliens(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	assert.Nil(t, err)
	assert.NotPanics(t, func() { worldMap.UnleaseNAliens(8) })
}

func TestDestroyCity(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	assert.Nil(t, err)
	assert.Equal(t, 5, len(worldMap.GetCities()))
	assert.NotPanics(t, func() { worldMap.DestroyCity("Foo") })
	assert.Equal(t, 4, len(worldMap.GetCities()))
}

func TestKillAliens(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	assert.Nil(t, err)
	assert.NotPanics(t, func() { worldMap.UnleaseNAliens(8) })
	assert.Equal(t, 8, len(worldMap.GetAlienList()))
	assert.NotPanics(t, func() { worldMap.KillAliens([]worldmap.Alien{"alien-0", "alien-1"}) })
	assert.Equal(t, 6, len(worldMap.GetAlienList()))
}
