package invasion_test

import (
	"strings"
	"testing"

	"github.com/harry-hov/alien-invasion/invasion"
	"github.com/harry-hov/alien-invasion/worldmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const worldMapInput string = `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
`

func ResetInvasion(i *invasion.Invasion) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	if err != nil {
		panic(err)
	}
	worldMap.UnleaseNAliens(8)
	i.SetWorldMap(worldMap)
	i.SetMove(0)
	i.SetFinished(false)
	i.SetConclusion("")
}

func TestInitInvasion(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	require.Nil(t, err)
	assert.NotPanics(t, func() { invasion.InitInvasion(worldMap, 8) })
}

func TestSetAndGetGetWorldMap(t *testing.T) {
	wm := worldmap.New()
	var in invasion.Invasion

	in.SetWorldMap(wm)
	assert.Equal(t, wm, in.GetWorldMap())
}

func TestSetAndGetCurrentMove(t *testing.T) {
	var in invasion.Invasion

	in.SetMove(10)
	assert.Equal(t, 10, in.GetCurrentMove())
}

func TestSetAndGetConclusion(t *testing.T) {
	var in invasion.Invasion

	in.SetConclusion("concluded")
	assert.Equal(t, invasion.Conclusion("concluded"), in.Conclusion())
}

func TestMakeMove(t *testing.T) {
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	require.Nil(t, err)
	i := invasion.InitInvasion(worldMap, 8)
	assert.Equal(t, 0, i.GetCurrentMove())
	i.MakeMove()
	assert.Equal(t, 1, i.GetCurrentMove())
}

func TestIsFinished(t *testing.T) {
	var in *invasion.Invasion
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	require.Nil(t, err)
	assert.NotPanics(t, func() { in = invasion.InitInvasion(worldMap, 8) })

	// Test moves exceeds limit
	in.SetMove(10000)
	res := in.IsFinished()
	assert.Equal(t, true, res)
	assert.Equal(t, invasion.Conclusion("exceeds maximum moves"), in.Conclusion())

	ResetInvasion(in)

	// Test all cities destroyed
	for _, city := range []worldmap.City{"Foo", "Bar", "Baz", "Bee", "Qu-ux"} {
		in.GetWorldMap().DestroyCity(city)
	}
	assert.Equal(t, true, in.IsFinished())
	assert.Equal(t, invasion.Conclusion("all cities destroyed"), in.Conclusion())

	ResetInvasion(in)

	// Test all aliens died
	aliens := in.GetWorldMap().GetAlienList()
	in.GetWorldMap().KillAliens(aliens)
	assert.Equal(t, true, in.IsFinished())
	assert.Equal(t, invasion.Conclusion("all aliens died"), in.Conclusion())

	ResetInvasion(in)

	// Test aliens won
	aliens = in.GetWorldMap().GetAlienList()
	in.GetWorldMap().KillAliens(aliens[1:])
	assert.Equal(t, true, in.IsFinished())
	assert.True(t, strings.HasSuffix(string(in.Conclusion()), "won"))

	ResetInvasion(in)

	// All aliens trapped
	for _, city := range []worldmap.City{"Foo", "Bar", "Baz", "Bee"} {
		in.GetWorldMap().DestroyCity(city)
	}
	assert.Equal(t, true, in.IsFinished())
	assert.Equal(t, invasion.Conclusion("all aliens trapped"), in.Conclusion())
}

func TestFight(t *testing.T) {
	const worldMapInput string = `Foo north=Bar`
	var in *invasion.Invasion
	worldMap, err := worldmap.InitWorldMap(strings.NewReader(worldMapInput))
	require.Nil(t, err)
	assert.NotPanics(t, func() { in = invasion.InitInvasion(worldMap, 1000) })
	in.Fight()
	assert.Equal(t, 0, len(in.GetWorldMap().GetAlienList()))
	assert.Equal(t, 0, len(in.GetWorldMap().GetCities()))
}
