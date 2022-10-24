package structs

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/AleksandarHr/AlienInvasion/utils"
	petname "github.com/dustinkirkland/golang-petname"
)

type Alien struct {
	ID       int
	Name     string
	Location *City
}

func CreateAlien(newAlienID int) *Alien {
	alien := &Alien{ID: newAlienID}
	alienName := alien.GiveAlienPetName(3, "-")
	alien.Name = alienName
	return alien
}

func (a *Alien) MoveToCity(newLocation *City) {
	a.Location = newLocation
	fmt.Printf("Alien %d moved to %s\n", a.ID, (newLocation.Name))
}

func (a *Alien) PickRandomNeighbourCity() (*City, error) {
	alienCity := a.Location
	availableDirections := []Direction{}
	for dir, _ := range alienCity.Neighbours {
		availableDirections = append(availableDirections, dir)
	}

	if len(availableDirections) == 0 {
		fmt.Printf("Alien %d is trapped in %s\n", a.ID, a.Location.Name)
		return nil, fmt.Errorf("Alien is trapped")
	}

	directionIndex, _ := utils.GenerateRandomNumber(len(availableDirections))
	randomDirection := availableDirections[directionIndex]
	randomNextCity := alienCity.Neighbours[randomDirection]

	// fmt.Printf("Alien %d wants to move %s from %s to %s\n", a.ID, randomDirection.String(), a.Location.Name, randomNextCity.Name)
	return randomNextCity, nil
}

func (a *Alien) String() string {
	return "Alien " + strconv.Itoa(a.ID) + " is currently located in " + (a.Location.Name)
}

func (a *Alien) GiveAlienPetName(wordCount int, nameSeparator string) string {
	rand.Seed(time.Now().UnixNano())
	alienName := petname.Generate(wordCount, nameSeparator) + "_" + strconv.Itoa(a.ID)
	a.Name = alienName
	return alienName
}
