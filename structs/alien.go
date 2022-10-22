package structs

import (
	"fmt"
	"strconv"
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

func (a *Alien) String() string {
	return "Alien " + strconv.Itoa(a.ID) + " is currently located in " + (a.Location.Name)
}
