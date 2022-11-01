package structs

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAlien(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	randomAlienID := rand.Intn(100)

	alien := CreateAlien(randomAlienID)
	assert.Equal(t, alien.ID, randomAlienID, "The two numbers should be the same.")

}

func TestMoveAlienToCity(t *testing.T) {
	assert := assert.New(t)

	str := "encode this string"
	randomString := base64.StdEncoding.EncodeToString([]byte(str))

	alien := CreateAlien(1)
	alienCity := CreateCity(randomString)
	err := alien.MoveToCity(alienCity)
	if err != nil {
		t.Fatalf("Adding alien to city failed: %v", err)
	}
	assert.Equal(alien.Location.Name, randomString, "The two city names should be the same.")
	assert.Equal(alien.Location, alienCity, "The two cities should be the same.")

	var invalidCity *City = nil
	err = alien.MoveToCity(invalidCity)
	assert.NotNil(t, err, "Invalid city, error should not be nil.")
}

func TestPickRandomNeighbourOfIsolatedCity(t *testing.T) {
	assert := assert.New(t)

	alien := CreateAlien(1)
	alienCity := CreateCity("RandomCityName")
	err := alien.MoveToCity(alienCity)
	if err != nil {
		t.Fatalf("Adding alien to city failed: %v", err)
	}

	assert.Equal(len(alien.Location.Neighbours), 0, "The city should have 0 neighbours.")
	randomNeighbour, _ := alien.PickRandomNeighbourCity()
	assert.Nil(randomNeighbour, "Alien is trapped, no city should have been chosen.")
}

func TestPickRandomNeighbour(t *testing.T) {
	assert := assert.New(t)

	alien := CreateAlien(1)
	alienCity := CreateCity("RandomCityName")
	neighbourCity := CreateCity("NeighbourCityName")
	alienCity.AddNeighbour(North, neighbourCity)
	alien.MoveToCity(alienCity)

	assert.Equal(len(alien.Location.Neighbours), 1, "The city should have 1 neighbour.")
	randomNeighbour, _ := alien.PickRandomNeighbourCity()
	assert.NotNil(randomNeighbour, "Alien is not trapped, a neighbour city should have been chosen.")
	assert.Equal(randomNeighbour, neighbourCity, "Cities should be the same.")
}

func TestGiveAlienPetName(t *testing.T) {
	assert := assert.New(t)

	rand.Seed(time.Now().UnixNano())
	randomAlienID := rand.Intn(100)
	alien := CreateAlien(randomAlienID)

	randomAlienNameWordCount := rand.Intn(3) + 1
	alienName := alien.GiveAlienPetName(randomAlienNameWordCount, "-")

	assert.Equal(alien.Name, alienName, "The two names should be the same.")
	assert.True(strings.HasSuffix(alien.Name, strconv.Itoa(alien.ID)), "The alien name should end with the alien ID.")
}
