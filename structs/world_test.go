package structs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewCity(t *testing.T) {
	world := CreateWorld()
	city := CreateCity("Foo")
	err := world.AddNewCity(city)

	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, len(world.cities), 1, "World should have 1 city.")
	assert.Contains(t, world.cities, city.Name, "World should contain city.")
	assert.Equal(t, world.cities[city.Name], city, "Foo should correspond to the correct city.")
}

// GetAllCities returns all currentlt existing cities
func TestGetAllCities(t *testing.T) {
	world := CreateWorld()
	cityFoo := CreateCity("Foo")
	cityBee := CreateCity("Bee")
	world.AddNewCity(cityFoo)
	world.AddNewCity(cityBee)

	cities, err := world.GetAllCities()

	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, len(cities), 2, "Function should return 2 cities.")

	citiesNames := []string{cities[0].Name, cities[1].Name}
	sort.Sort(sort.StringSlice(citiesNames))
	assert.Equal(t, citiesNames[0], "Bee", "City names should be the same.")
	assert.Equal(t, citiesNames[1], "Foo", "City names should be the same.")
}

func TestInitializeWorld(t *testing.T) {
	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo":   {"north=Bar", "west=Baz", "south=Qu-ux"},
		"Bar":   {"south=Foo", "west=Bee"},
		"Baz":   {"east=Foo"},
		"Qu-ux": {"north=Foo"},
		"Bee":   {"east=Bar"},
	}

	world.InitializeWorld(mapInfo)

	// cities added to world.cities
	assert.Equal(t, len(world.cities), 5, "The numbers of cities should be 5.")
	assert.Contains(t, world.cities, "Foo", "Foo should exist in the cities map.")
	assert.Contains(t, world.cities, "Bar", "Bar should exist in the cities map.")
	assert.Contains(t, world.cities, "Baz", "Baz should exist in the cities map.")
	assert.Contains(t, world.cities, "Bee", "Bee should exist in the cities map.")
	assert.Contains(t, world.cities, "Qu-ux", "Qu-ux should exist in the cities map.")
	assert.Equal(t, len(world.aliens), 0, "The number of aliens should be 0.")

	// cities added to world.cityConnections
	assert.Equal(t, len(world.cityConnections), 5, "The city connections map should have 5 entries.")
	assert.Contains(t, world.cityConnections, "Foo", "Foo should exist in the cityConnections map.")
	fooConnections := world.cityConnections["Foo"]
	assert.Equal(t, len(fooConnections), 3, "Foo shuold have 3 connections.")
	assert.Contains(t, fooConnections, "Bar", "Bar should be a connection of Foo.")
	assert.Contains(t, fooConnections, "Baz", "Baz should be a connection of Foo.")
	assert.Contains(t, fooConnections, "Qu-ux", "Qu-ux should be a connection of Foo.")

	// citiyConnections contain the correct information
	assert.Contains(t, world.cityConnections, "Bar", "Bar should exist in the cityConnections map.")
	barConnections := world.cityConnections["Bar"]
	assert.Equal(t, len(barConnections), 2, "Bar shuold have 2 connections.")
	assert.Contains(t, barConnections, "Foo", "Foo should be a connection of Bar.")
	assert.Contains(t, barConnections, "Bee", "Bee should be a connection of Bar.")

	assert.Contains(t, world.cityConnections, "Baz", "Baz should exist in the cityConnections map.")
	bazConnections := world.cityConnections["Baz"]
	assert.Equal(t, len(bazConnections), 1, "Baz shuold have 1 connection.")
	assert.Contains(t, bazConnections, "Foo", "Foo should be a connection of Baz.")

	assert.Contains(t, world.cityConnections, "Qu-ux", "Qu-ux should exist in the cityConnections map.")
	quuxConnections := world.cityConnections["Qu-ux"]
	assert.Equal(t, len(quuxConnections), 1, "Qu-ux shuold have 1 connections.")
	assert.Contains(t, quuxConnections, "Foo", "Foo should be a connection of Qu-ux.")

	assert.Contains(t, world.cityConnections, "Bee", "Bee should exist in the cityConnections map.")
	beeConnections := world.cityConnections["Bee"]
	assert.Equal(t, len(beeConnections), 1, "Bee shuold have 1 connections.")
	assert.Contains(t, beeConnections, "Bar", "Bar should be a connection of Bee.")

	// neighbour data of cities is updated properly
	foo := world.cities["Foo"]
	bar := world.cities["Bar"]
	baz := world.cities["Baz"]
	bee := world.cities["Bee"]
	quux := world.cities["Qu-ux"]

	assert.Equal(t, len(foo.Neighbours), 3, "Foo should have 3 neighbours")
	assert.Equal(t, foo.Neighbours[North].Name, "Bar", "Cities should be the same")
	assert.Equal(t, foo.Neighbours[West].Name, "Baz", "Cities should be the same")
	assert.Equal(t, foo.Neighbours[South].Name, "Qu-ux", "Cities should be the same")

	assert.Equal(t, len(bar.Neighbours), 2, "Bar should have 2 neighbours")
	assert.Equal(t, bar.Neighbours[South].Name, "Foo", "Cities should be the same")
	assert.Equal(t, bar.Neighbours[West].Name, "Bee", "Cities should be the same")

	assert.Equal(t, len(baz.Neighbours), 1, "Baz should have 1 neighbour1")
	assert.Equal(t, baz.Neighbours[East].Name, "Foo", "Cities should be the same")

	assert.Equal(t, len(bee.Neighbours), 1, "Bee should have 1 neighbour1")
	assert.Equal(t, bee.Neighbours[East].Name, "Bar", "Cities should be the same")

	assert.Equal(t, len(quux.Neighbours), 1, "Qu-ux should have 1 neighbour")
	assert.Equal(t, quux.Neighbours[North].Name, "Foo", "Cities should be the same")
}

func TestRemoveCity(t *testing.T) {
	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {"north=Bar"},
		"Bar": {"south=Foo", "west=Bee"},
		"Bee": {"east=Bar"},
	}

	world.InitializeWorld(mapInfo)
	world.RemoveCity(world.cities["Foo"])

	assert.Equal(t, len(world.cities), 2, "The numbers of cities should be 2.")
	assert.Contains(t, world.cities, "Bar", "Foo should exist in the cities map.")
	assert.Contains(t, world.cities, "Bee", "Bar should exist in the cities map.")
	assert.NotContains(t, world.cities, "Foo", "Foo should not exist in the cities map.")

	assert.Equal(t, len(world.cityConnections), 2, "The city connections map should have 2 entries.")
	assert.NotContains(t, world.cityConnections, "Foo", "Foo should not exist in the cityConnections map.")

	assert.Contains(t, world.cityConnections, "Bee", "Bee should exist in the cityConnections map.")
	beeConnections := world.cityConnections["Bee"]
	assert.Equal(t, len(beeConnections), 1, "Bee shuold have 1 connections.")
	assert.Contains(t, beeConnections, "Bar", "Bar should be a connection of Bee.")

	assert.Contains(t, world.cityConnections, "Bar", "Bar should exist in the cityConnections map.")
	barConnections := world.cityConnections["Bar"]
	assert.Equal(t, len(barConnections), 1, "Bar shuold have 1 connection.")
	assert.Contains(t, barConnections, "Bee", "Bee should be a connection of Bar.")
	assert.NotContains(t, barConnections, "Foo", "Foo should be a connection of Bar.")

	bar := world.cities["Bar"]
	assert.Equal(t, len(bar.Neighbours), 1, "Bar should have 1 neighbour")
	assert.NotContains(t, bar.Neighbours, South, "Bar should not have south connection")
	assert.Equal(t, bar.Neighbours[West].Name, "Bee", "Cities should be the same")
}
