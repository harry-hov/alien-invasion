package worldmap

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"

	wmerror "github.com/harry-hov/alien-invasion/error"
)

type Alien string
type City string
type Direction string

const (
	East  = Direction("east")
	North = Direction("north")
	South = Direction("south")
	West  = Direction("west")
)

func (d Direction) IsValid() bool {
	return d == East || d == North || d == South || d == West
}

// Get opposite direction
// e.g North.Opposite() == South
func (d Direction) GetOpposite() (Direction, error) {
	switch d {
	case East:
		return West, nil
	case North:
		return South, nil
	case South:
		return North, nil
	case West:
		return East, nil
	}
	return Direction(""), wmerror.Wrap(wmerror.ErrInvalidDirection, fmt.Sprintf("(%v)", d))
}

type WorldMap struct {
	cities map[City]map[Direction]City
	aliens map[Alien]City
}

// Returns empty WorldMap
func New() *WorldMap {
	return &WorldMap{
		cities: make(map[City]map[Direction]City),
		aliens: make(map[Alien]City),
	}
}

// InitWorldMap returns WorldMap from io.Reader
func InitWorldMap(reader io.Reader) (*WorldMap, error) {
	scanner := bufio.NewScanner(reader)
	worldMap := New()

	for scanner.Scan() {
		line := scanner.Text()

		// Skip blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Tokenize line
		tokens := strings.Split(line, " ")
		city := City(tokens[0])
		if len(tokens) < 2 {
			return nil, wmerror.Wrap(wmerror.ErrInvalidCity, fmt.Sprintf("isolated city (%v)", city))
		}

		// Add city to WorldMap
		worldMap.AddCity(city)

		for _, token := range tokens[1:] {
			directionEntry := strings.Split(token, "=")
			if len(directionEntry) != 2 {
				return nil, wmerror.Wrap(wmerror.ErrInvalidDirection, "cannot parse direction entry")
			}
			direction := Direction(strings.ToLower(directionEntry[0]))
			if !direction.IsValid() {
				return nil, wmerror.Wrap(wmerror.ErrInvalidDirection, "cannot parse direction")
			}
			directionCity := City(directionEntry[1])
			if err := worldMap.AppendCityDirection(city, directionCity, direction); err != nil {
				return nil, err
			}
		}
	}

	return worldMap, nil
}

// Add a city to WorldMap
func (wm *WorldMap) AddCity(c City) {
	if _, ok := wm.cities[c]; !ok {
		wm.cities[c] = make(map[Direction]City)
	}
}

// Add a city to WorldMap with error
func (wm *WorldMap) AddCityE(c City) error {
	if _, ok := wm.cities[c]; ok {
		return wmerror.Wrap(wmerror.ErrInvalidCity, fmt.Sprintf("duplicate city (%v)", c))
	}
	wm.cities[c] = make(map[Direction]City)
	return nil
}

// Append direction to city
func (wm *WorldMap) AppendCityDirection(city, directionCity City, direction Direction) error {
	if city == directionCity {
		return wmerror.Wrap(wmerror.ErrInvalidDirection, "city cannot direct to itself")
	}

	// Add directionCity to WorldMap
	wm.AddCity(directionCity)
	if val, ok := wm.cities[city][direction]; ok && val != directionCity {
		return wmerror.Wrap(wmerror.ErrInvalidDirection, fmt.Sprintf("ambiguous direction from city (%v) to (%v)", city, directionCity))
	}

	oppositeDirection, err := direction.GetOpposite()
	if err != nil {
		return err
	}
	if val, ok := wm.cities[directionCity][oppositeDirection]; ok && val != city {
		return wmerror.Wrap(wmerror.ErrInvalidDirection, fmt.Sprintf("ambiguous direction from city (%v) to (%v)", directionCity, city))
	}

	/*
	 * Add direction to both cities (`city` and `directionCity`)
	 * i.e
	 * 		if foo's south is baz
	 * 		implies baz's north is foo
	 */
	wm.cities[city][direction] = directionCity
	wm.cities[directionCity][oppositeDirection] = city

	return nil
}

// PrintWorldMap prints the world map in the same format as the input file.
func (wm *WorldMap) Print() {
	var out string
	for city, directionEntry := range wm.cities {
		out += fmt.Sprintf("%v:", city)
		for direction, directionCity := range directionEntry {
			out += fmt.Sprintf(" %v=%v", direction, directionCity)
		}
		out += "\n"
	}
	fmt.Print(out)
}

// GetCities returns the list of cities
func (wm *WorldMap) GetCities() (cities []City) {
	for city := range wm.cities {
		cities = append(cities, city)
	}
	return
}

// UnleaseAliens unleases N aliens in the WorldMap
func (wm *WorldMap) UnleaseNAliens(aliens uint) {
	cities := wm.GetCities()
	for i := uint(0); i < aliens; i++ {
		random := rand.Intn(len(wm.cities))
		name := Alien(fmt.Sprintf("alien-%v", i))
		wm.aliens[name] = cities[random]
	}
}
