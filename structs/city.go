package structs

import "fmt"

type City struct {
	Name       string
	Neighbours map[Direction]*City
}

func CreateCity(cityName string) *City {
	return &City{
		Name:       cityName,
		Neighbours: map[Direction]*City{},
	}
}

func (c *City) AddNeighbour(dir Direction, neighbour *City) error {
	if dir != North && dir != East && dir != South && dir != West {
		// todo: handle error
		return fmt.Errorf("Invalid direction.")
	}
	c.Neighbours[dir] = neighbour
	return nil
}

func (c *City) HasNeighbours() bool {
	if c.Neighbours[North] != nil || c.Neighbours[East] != nil || c.Neighbours[South] != nil || c.Neighbours[West] != nil {
		return true
	}

	return false
}
