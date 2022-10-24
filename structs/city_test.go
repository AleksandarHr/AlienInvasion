package structs

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCity(t *testing.T) {
	assert := assert.New(t)

	str := "encode this string"
	randomString := base64.StdEncoding.EncodeToString([]byte(str))
	city := CreateCity(randomString)

	assert.Equal(randomString, city.Name, "The two city names should be the same.")
	assert.Equal(len(city.Neighbours), 0, "The city should have 0 neighbours.")
}

func TestAddNeighbour(t *testing.T) {
	assert := assert.New(t)

	city := CreateCity("Random City Name")

	northNeighbour := CreateCity("North Neighbour")
	city.AddNeighbour(North, northNeighbour)
	assert.Equal(len(city.Neighbours), 1, "The city should have 1 neighbour.")
	assert.NotNil(city.Neighbours[North], "North neighbour should be nil.")
	assert.Equal(city.Neighbours[North], northNeighbour, "The two city names should be the same.")

	eastNeighbour := CreateCity("East Neighbour")
	city.AddNeighbour(East, eastNeighbour)
	assert.Equal(len(city.Neighbours), 2, "The city should have 2 neighbours.")
	assert.NotNil(city.Neighbours[East], "East neighbour should be nil.")
	assert.Equal(city.Neighbours[East], eastNeighbour, "The two city names should be the same.")

	southNeighbour := CreateCity("South Neighbour")
	city.AddNeighbour(South, southNeighbour)
	assert.Equal(len(city.Neighbours), 3, "The city should have 3 neighbours.")
	assert.NotNil(city.Neighbours[South], "South neighbour should be nil.")
	assert.Equal(city.Neighbours[South], southNeighbour, "The two city names should be the same.")

	westNeighbour := CreateCity("West Neighbour")
	city.AddNeighbour(West, westNeighbour)
	assert.Equal(len(city.Neighbours), 4, "The city should have 4 neighbours.")
	assert.NotNil(city.Neighbours[West], "West neighbour should be nil.")
	assert.Equal(city.Neighbours[West], westNeighbour, "The two city names should be the same.")
}

func TestHasNeighbours(t *testing.T) {
	assert := assert.New(t)

	city := CreateCity("Random City Name")
	assert.False(city.HasNeighbours())

	neighbour := CreateCity("Random Neighbour City Name")
	city.AddNeighbour(North, neighbour)
	assert.True(city.HasNeighbours())
}
