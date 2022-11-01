package structs

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCity(t *testing.T) {
	assert := assert.New(t)

	str := "encodethisstring"
	randomString := base64.StdEncoding.EncodeToString([]byte(str))
	city := CreateCity(randomString)

	assert.Equal(randomString, city.Name, "The two city names should be the same.")
	assert.Equal(len(city.Neighbours), 0, "The city should have 0 neighbours.")
}

func TestAddNeighbour(t *testing.T) {
	assert := assert.New(t)

	city := CreateCity("RandomCityName")

	var invalidNeighbour *City = nil
	err := city.AddNeighbour(North, invalidNeighbour)
	assert.NotNil(t, err, "City was nil, error was expected")

	randomCity := CreateCity("RandomCity")
	err = city.AddNeighbour(Invalid, randomCity)
	assert.NotNil(t, err, "Invalid direction, error was expected")

	northNeighbour := CreateCity("NorthNeighbour")
	err = city.AddNeighbour(North, northNeighbour)
	if err != nil {
		t.Fatalf("Adding neighbour city failed: %v", err)
	}
	assert.Equal(len(city.Neighbours), 1, "The city should have 1 neighbour.")
	assert.NotNil(city.Neighbours[North], "North neighbour should be nil.")
	assert.Equal(city.Neighbours[North], northNeighbour, "The two city names should be the same.")

	eastNeighbour := CreateCity("EastNeighbour")
	err = city.AddNeighbour(East, eastNeighbour)
	if err != nil {
		t.Fatalf("Adding neighbour city failed: %v", err)
	}
	assert.Equal(len(city.Neighbours), 2, "The city should have 2 neighbours.")
	assert.NotNil(city.Neighbours[East], "East neighbour should be nil.")
	assert.Equal(city.Neighbours[East], eastNeighbour, "The two city names should be the same.")

	southNeighbour := CreateCity("SouthNeighbour")
	err = city.AddNeighbour(South, southNeighbour)
	if err != nil {
		t.Fatalf("Adding neighbour city failed: %v", err)
	}
	assert.Equal(len(city.Neighbours), 3, "The city should have 3 neighbours.")
	assert.NotNil(city.Neighbours[South], "South neighbour should be nil.")
	assert.Equal(city.Neighbours[South], southNeighbour, "The two city names should be the same.")

	westNeighbour := CreateCity("WestNeighbour")
	err = city.AddNeighbour(West, westNeighbour)
	if err != nil {
		t.Fatalf("Adding neighbour city failed: %v", err)
	}
	assert.Equal(len(city.Neighbours), 4, "The city should have 4 neighbours.")
	assert.NotNil(city.Neighbours[West], "West neighbour should be nil.")
	assert.Equal(city.Neighbours[West], westNeighbour, "The two city names should be the same.")
}

func TestHasNeighbours(t *testing.T) {
	assert := assert.New(t)

	city := CreateCity("RandomCityName")
	assert.False(city.HasNeighbours())

	neighbour := CreateCity("RandomNeighbourCityName")
	err := city.AddNeighbour(North, neighbour)
	if err != nil {
		t.Fatalf("Adding neighbour city failed: %v", err)
	}
	assert.True(city.HasNeighbours())
}
