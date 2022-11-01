package structs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewCity(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()

	var nilCity *City = nil
	err := world.AddNewCity(nilCity)
	assert.NotNil(err, "Should not be able to add nil city")

	city := CreateCity("Foo")
	err = world.AddNewCity(city)
	assert.Nil(err, "Should be able to add city.")
	assert.Equal(len(world.cities), 1, "World should have 1 city.")
	assert.Contains(world.cities, city.Name, "World should contain city.")
	assert.Equal(world.cities[city.Name], city, "Foo should correspond to the correct city.")
}

// GetAllCities returns all currentlt existing cities
func TestGetAllCities(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	cityFoo := CreateCity("Foo")
	cityBee := CreateCity("Bee")
	world.AddNewCity(cityFoo)
	world.AddNewCity(cityBee)

	cities, err := world.GetAllCities()

	assert.Nil(err, "Error should be nil.")
	assert.Equal(len(cities), 2, "Function should return 2 cities.")

	citiesNames := []string{cities[0].Name, cities[1].Name}
	sort.Sort(sort.StringSlice(citiesNames))
	assert.Equal(citiesNames[0], "Bee", "City names should be the same.")
	assert.Equal(citiesNames[1], "Foo", "City names should be the same.")
}

func TestInitializeWorld(t *testing.T) {
	assert := assert.New(t)

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
	assert.Equal(len(world.cities), 5, "The numbers of cities should be 5.")
	assert.Contains(world.cities, "Foo", "Foo should exist in the cities map.")
	assert.Contains(world.cities, "Bar", "Bar should exist in the cities map.")
	assert.Contains(world.cities, "Baz", "Baz should exist in the cities map.")
	assert.Contains(world.cities, "Bee", "Bee should exist in the cities map.")
	assert.Contains(world.cities, "Qu-ux", "Qu-ux should exist in the cities map.")
	assert.Equal(len(world.aliens), 0, "The number of aliens should be 0.")

	// cities added to world.cityConnections
	assert.Equal(len(world.cityConnections), 5, "The city connections map should have 5 entries.")
	assert.Contains(world.cityConnections, "Foo", "Foo should exist in the cityConnections map.")
	fooConnections := world.cityConnections["Foo"]
	assert.Equal(len(fooConnections), 3, "Foo shuold have 3 connections.")
	assert.Contains(fooConnections, "Bar", "Bar should be a connection of Foo.")
	assert.Contains(fooConnections, "Baz", "Baz should be a connection of Foo.")
	assert.Contains(fooConnections, "Qu-ux", "Qu-ux should be a connection of Foo.")

	// citiyConnections contain the correct information
	assert.Contains(world.cityConnections, "Bar", "Bar should exist in the cityConnections map.")
	barConnections := world.cityConnections["Bar"]
	assert.Equal(len(barConnections), 2, "Bar shuold have 2 connections.")
	assert.Contains(barConnections, "Foo", "Foo should be a connection of Bar.")
	assert.Contains(barConnections, "Bee", "Bee should be a connection of Bar.")

	assert.Contains(world.cityConnections, "Baz", "Baz should exist in the cityConnections map.")
	bazConnections := world.cityConnections["Baz"]
	assert.Equal(len(bazConnections), 1, "Baz shuold have 1 connection.")
	assert.Contains(bazConnections, "Foo", "Foo should be a connection of Baz.")

	assert.Contains(world.cityConnections, "Qu-ux", "Qu-ux should exist in the cityConnections map.")
	quuxConnections := world.cityConnections["Qu-ux"]
	assert.Equal(len(quuxConnections), 1, "Qu-ux shuold have 1 connections.")
	assert.Contains(quuxConnections, "Foo", "Foo should be a connection of Qu-ux.")

	assert.Contains(world.cityConnections, "Bee", "Bee should exist in the cityConnections map.")
	beeConnections := world.cityConnections["Bee"]
	assert.Equal(len(beeConnections), 1, "Bee shuold have 1 connections.")
	assert.Contains(beeConnections, "Bar", "Bar should be a connection of Bee.")

	// neighbour data of cities is updated properly
	foo := world.cities["Foo"]
	bar := world.cities["Bar"]
	baz := world.cities["Baz"]
	bee := world.cities["Bee"]
	quux := world.cities["Qu-ux"]

	assert.Equal(len(foo.Neighbours), 3, "Foo should have 3 neighbours")
	assert.Equal(foo.Neighbours[North].Name, "Bar", "Cities should be the same")
	assert.Equal(foo.Neighbours[West].Name, "Baz", "Cities should be the same")
	assert.Equal(foo.Neighbours[South].Name, "Qu-ux", "Cities should be the same")

	assert.Equal(len(bar.Neighbours), 2, "Bar should have 2 neighbours")
	assert.Equal(bar.Neighbours[South].Name, "Foo", "Cities should be the same")
	assert.Equal(bar.Neighbours[West].Name, "Bee", "Cities should be the same")

	assert.Equal(len(baz.Neighbours), 1, "Baz should have 1 neighbour1")
	assert.Equal(baz.Neighbours[East].Name, "Foo", "Cities should be the same")

	assert.Equal(len(bee.Neighbours), 1, "Bee should have 1 neighbour1")
	assert.Equal(bee.Neighbours[East].Name, "Bar", "Cities should be the same")

	assert.Equal(len(quux.Neighbours), 1, "Qu-ux should have 1 neighbour")
	assert.Equal(quux.Neighbours[North].Name, "Foo", "Cities should be the same")
}

func TestRemoveCity(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {"north=Bar"},
		"Bar": {"south=Foo", "west=Bee"},
		"Bee": {"east=Bar"},
	}

	world.InitializeWorld(mapInfo)

	var nilCity *City = nil
	err := world.RemoveCity(nilCity)
	assert.NotNil(err, "Should not be able to remove nil city")

	cityToRemove := world.cities["Foo"]
	err = world.RemoveCity(cityToRemove)
	assert.Nil(err, "Should be able to remove existing city.")

	err = world.RemoveCity(cityToRemove)
	assert.NotNil(err, "Should not be able to remove non-existing city.")

	assert.Equal(len(world.cities), 2, "The numbers of cities should be 2.")
	assert.Contains(world.cities, "Bar", "Foo should exist in the cities map.")
	assert.Contains(world.cities, "Bee", "Bar should exist in the cities map.")
	assert.NotContains(world.cities, "Foo", "Foo should not exist in the cities map.")

	assert.Equal(len(world.cityConnections), 2, "The city connections map should have 2 entries.")
	assert.NotContains(world.cityConnections, "Foo", "Foo should not exist in the cityConnections map.")

	assert.Contains(world.cityConnections, "Bee", "Bee should exist in the cityConnections map.")
	beeConnections := world.cityConnections["Bee"]
	assert.Equal(len(beeConnections), 1, "Bee shuold have 1 connections.")
	assert.Contains(beeConnections, "Bar", "Bar should be a connection of Bee.")

	assert.Contains(world.cityConnections, "Bar", "Bar should exist in the cityConnections map.")
	barConnections := world.cityConnections["Bar"]
	assert.Equal(len(barConnections), 1, "Bar shuold have 1 connection.")
	assert.Contains(barConnections, "Bee", "Bee should be a connection of Bar.")
	assert.NotContains(barConnections, "Foo", "Foo should be a connection of Bar.")

	bar := world.cities["Bar"]
	assert.Equal(len(bar.Neighbours), 1, "Bar should have 1 neighbour")
	assert.NotContains(bar.Neighbours, South, "Bar should not have south connection")
	assert.Equal(bar.Neighbours[West].Name, "Bee", "Cities should be the same")
}

func TestFreeAliens(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {"north=Bar"},
		"Bar": {"south=Foo"},
		"Baz": {},
	}

	world.InitializeWorld(mapInfo)
	alien := CreateAlien(0)
	world.AddAlienToCity(alien, world.cities["Foo"], SpawningAliens)
	trappedAlien := CreateAlien(1)
	world.AddAlienToCity(trappedAlien, world.cities["Baz"], SpawningAliens)

	freeAliens, err := world.GetFreeAliens()
	assert.Nil(err, "Should be returning all free aliens without an error.")
	assert.Equal(len(freeAliens), 1, "There is only one free alien")
	assert.Equal(freeAliens[0].ID, 0, "The free alien should be with ID 1.")

	free, err := world.IsAlienFree(alien)
	assert.Nil(err, "Should be returning true for a free alien.")
	assert.True(free, "Should be returning true for a free alien.")

	world.RemoveCity(world.cities["Bar"])
	freeAliens, err = world.GetFreeAliens()
	assert.Nil(err, "Should be returning all free aliens without an error.")
	assert.Equal(len(freeAliens), 0, "There should be no free aliens.")

	free, err = world.IsAlienFree(trappedAlien)
	assert.Nil(err, "Should be returning false for a trapped alien.")
	assert.False(free, "Should be returning false for a trapped alien.")

	var nilAlien *Alien = nil
	_, err = world.IsAlienFree(nilAlien)
	assert.NotNil(err, "Should not be handling nil aliens.")
}

func TestIsAlienAlive(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {},
	}

	world.InitializeWorld(mapInfo)

	var nilAlien *Alien = nil
	_, err := world.IsAlienAlive(nilAlien)
	assert.NotNil(err, "Should not be able to handle nil alien")

	alien := CreateAlien(0)
	world.AddAlienToCity(alien, world.cities["Foo"], SpawningAliens)

	deadAlien := CreateAlien(1)
	alive, err := world.IsAlienAlive(deadAlien)
	assert.Nil(err, "Should return false for a dead alien")
	assert.False(alive, "Should return false for a dead alien")

	alive, err = world.IsAlienAlive(alien)
	assert.Nil(err, "Should return true for an alive alien")
	assert.True(alive, "Should return true for an alive alien")

}

func TestAllAliens(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {"north=Bar"},
		"Bar": {"south=Foo"},
	}

	world.InitializeWorld(mapInfo)
	foo := world.cities["Foo"]
	alien := CreateAlien(0)
	world.AddAlienToCity(alien, foo, SpawningAliens)

	dead := world.AllAliensDead()
	assert.False(dead, "There is one alien alive.")
	trapped := world.AllAliensTrapped()
	assert.False(trapped, "There is one alien not trapped.")

	world.RemoveCity(world.cities["Bar"])
	trapped = world.AllAliensTrapped()
	assert.True(trapped, "All aliens are trapped.")
	world.killAlien(alien)
	dead = world.AllAliensDead()
	assert.True(dead, "All aliens are dead.")
}

func TestAddAlienToCity(t *testing.T) {
	assert := assert.New(t)

	world := CreateWorld()
	mapInfo := map[string][]string{
		"Foo": {},
		"Bar": {"south=Baz"},
		"Baz": {"north=Bar"},
	}

	world.InitializeWorld(mapInfo)
	foo := world.cities["Foo"]
	var nilAlien *Alien = nil

	added, err := world.AddAlienToCity(nilAlien, foo, SpawningAliens)
	assert.False(added, "Alien should not have been added.")
	assert.NotNil(err, "Nil alien should raise an error.")

	alien := CreateAlien(0)
	var nilCity *City = nil
	added, err = world.AddAlienToCity(alien, nilCity, SpawningAliens)
	assert.False(added, "Alien should not have been added.")
	assert.NotNil(err, "Nil city should raise an error.")

	nonExistentCity := CreateCity("NonExistent")
	added, err = world.AddAlienToCity(alien, nonExistentCity, SpawningAliens)
	assert.False(added, "Alien should not have been added.")
	assert.NotNil(err, "Nonexistent city should raise an error.")

	alienTwo := CreateAlien(1)
	added, err = world.AddAlienToCity(alien, foo, SpawningAliens)
	assert.True(added, "Alien should have been added successfully.")
	assert.Nil(err, "Alien should have been added without an error.")

	added, err = world.AddAlienToCity(alienTwo, foo, SpawningAliens)
	assert.False(added, "Alien should not have been but rather destroyed.")
	assert.Nil(err, "No error should have been raised for destroying the alien.")

	alienThree := CreateAlien(1)
	bar := world.cities["Bar"]
	added, err = world.AddAlienToCity(alienThree, bar, SpawningAliens)
	assert.True(added, "Alien should have been added successfully.")
	assert.Nil(err, "Alien should have been added without an error.")
	baz := world.cities["Baz"]
	added, err = world.AddAlienToCity(alienThree, baz, MovingAliens)
	assert.True(added, "Alien should have been added successfully.")
	assert.Nil(err, "Alien should have been added without an error.")
}
