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
	alien.MoveToCity(alienCity)

	assert.Equal(alien.Location.Name, randomString, "The two city names should be the same.")
	assert.Equal(alien.Location, alienCity, "The two cities should be the same.")
}

func TestPickRandomNeighbourOfIsolatedCity(t *testing.T) {
	assert := assert.New(t)

	alien := CreateAlien(1)
	alienCity := CreateCity("Random City Name")
	alien.MoveToCity(alienCity)

	assert.Equal(len(alien.Location.Neighbours), 0, "The city should have 0 neighbours.")
	_, err := alien.PickRandomNeighbourCity()
	assert.EqualError(err, "Alien is trapped")
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
