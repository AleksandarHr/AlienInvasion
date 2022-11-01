package structs

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/AleksandarHr/AlienInvasion/utils"
	petname "github.com/dustinkirkland/golang-petname"
)

// Alien structure to represent information about an alien
type Alien struct {
	ID       int
	Name     string
	Location *City
}

// CreateAlien constructs an alien with a provided integer ID
//
//	and a randomly generated pet name
func CreateAlien(newAlienID int) *Alien {
	alien := &Alien{ID: newAlienID}
	alienName := alien.GiveAlienPetName(3, "-")
	alien.Name = alienName
	return alien
}

// MovetoCity changes the location of the alien to the provided city
func (a *Alien) MoveToCity(newLocation *City) error {
	if newLocation == nil {
		return &InvalidCityError{city: newLocation}
	}
	a.Location = newLocation
	return nil
}

// PickRandomNeighbourCity randomly chooses a city neighbouring the current alien location
func (a *Alien) PickRandomNeighbourCity() (*City, error) {
	alienCity := a.Location

	if len(alienCity.Neighbours) == 0 {
		return nil, nil
	}

	availableDirections := []Direction{}
	for dir := range alienCity.Neighbours {
		availableDirections = append(availableDirections, dir)
	}

	directionIndex, err := utils.GenerateRandomNumber(len(availableDirections))
	if err != nil {
		return nil, err
	}
	randomDirection := availableDirections[directionIndex]
	randomNextCity := alienCity.Neighbours[randomDirection]

	return randomNextCity, nil
}

// GiveAlienPetName generates a petname for the alien
func (a *Alien) GiveAlienPetName(wordCount int, nameSeparator string) string {
	rand.Seed(time.Now().UnixNano())
	alienName := petname.Generate(wordCount, nameSeparator) + "_" + strconv.Itoa(a.ID)
	a.Name = alienName
	return alienName
}
