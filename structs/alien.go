package structs

import (
	"fmt"
	"strconv"

	"github.com/AleksandarHr/AlienInvasion/utils"
)

type Alien struct {
	ID int
	// Name     string
	Location *City
}

func CreateAlien(newAlienID int) *Alien {
	return &Alien{ID: newAlienID}
}

func (a *Alien) SpawnAlien(originCity *City) {
	a.Location = originCity
	fmt.Printf("Alien %d spawned at %s\n", a.ID, (originCity.Name))
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

	fmt.Printf("Alien %d wants to move %s from %s to %s\n", a.ID, randomDirection.String(), a.Location.Name, randomNextCity.Name)
	return randomNextCity, nil
}

func (a *Alien) String() string {
	return "Alien " + strconv.Itoa(a.ID) + " is currently located in " + (a.Location.Name)
}
