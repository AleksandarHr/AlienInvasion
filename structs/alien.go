package structs

import (
	"fmt"
	"strconv"
	"strings"
)

type Alien struct {
	ID int
	// Name     string
	Location City
}

func CreateAlien(newAlienID int) *Alien {
	return &Alien{ID: newAlienID}
}

func (alien *Alien) MoveToCity(newLocation City) {
	alien.Location = newLocation
	fmt.Printf("Alien %d moved to %s\n", alien.ID, strings.ToUpper(newLocation.Name))
}

func (alien *Alien) String() string {
	return "Alien " + strconv.Itoa(alien.ID) + " is currently located in " + strings.ToUpper(alien.Location.Name)
}
