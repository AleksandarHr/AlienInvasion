package structs

import "fmt"

type City struct {
	Name       string
	Neighbours map[Direction]*City
}

func CreateCity(cityName string) *City {
	neighboursMap := map[Direction]*City{
		North: nil,
		East:  nil,
		South: nil,
		West:  nil,
	}
	return &City{
		Name:       cityName,
		Neighbours: neighboursMap,
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
