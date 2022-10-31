package structs

// City structure to represent information about a city
type City struct {
	Name       string
	Neighbours map[Direction]*City
}

// CreateCity constructs a City with provided cityName
func CreateCity(cityName string) *City {
	return &City{
		Name:       cityName,
		Neighbours: map[Direction]*City{},
	}
}

// AddNeigbhour adds a neighbour city in the provided direction
func (c *City) AddNeighbour(dir Direction, neighbour *City) error {
	if dir != North && dir != East && dir != South && dir != West {
		return &InvalidDirectionError{direction: dir}
	}

	if neighbour == nil {
		return &InvalidCityError{city: neighbour}
	}

	c.Neighbours[dir] = neighbour
	return nil
}

// HasNeighbours checks if a city has any neighbouring cities
func (c *City) HasNeighbours() bool {
	if c.Neighbours[North] != nil || c.Neighbours[East] != nil || c.Neighbours[South] != nil || c.Neighbours[West] != nil {
		return true
	}

	return false
}
