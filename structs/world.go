package structs

import (
	"fmt"
	"strings"
)

type World struct {
	// cities within the world
	cities map[string]*City

	// aliens within the world
	aliens map[string]*Alien

	// map a city name to the names of all the cities it is linked to
	cityConnections map[string]map[string]bool
}

func CreateWorld() *World {
	return &World{
		cities:          make(map[string]*City),
		aliens:          make(map[string]*Alien),
		cityConnections: make(map[string]map[string]bool),
	}
}

func (w *World) AddNewCity(name string) *City {
	// if city has already been added to the world
	if city, exists := w.cities[name]; exists {
		// todo: handle error
		return city
	}

	newCity := CreateCity(name)
	w.cities[name] = newCity
	return newCity
}

func (w *World) InitializeWorld(mapInfo map[string][]string) {
	for cityName, neighbourNames := range mapInfo {
		city := w.AddNewCity(cityName)

		for _, neighbourInfo := range neighbourNames {
			temp := strings.Split(neighbourInfo, "=")
			neighbourDirection, neighbourName := temp[0], temp[1]
			neighbourCity := w.AddNewCity(neighbourName)

			// add neighbour info to current city
			city.Neighbours[StringToDirection(neighbourDirection)] = neighbourCity

			// add reverse neighbour info to neighbour city
			neighbourCity.Neighbours[StringToDirection(neighbourDirection).OppositeDirection()] = city

			// add relevant links information
			w.addLinks(cityName, neighbourName)
		}
	}
}

func (w *World) addLinks(cityOneName, cityTwoName string) {
	if _, exists := w.cityConnections[cityOneName]; !exists {
		w.cityConnections[cityOneName] = make(map[string]bool)
	}
	w.cityConnections[cityOneName][cityTwoName] = true

	if _, exists := w.cityConnections[cityTwoName]; !exists {
		w.cityConnections[cityTwoName] = make(map[string]bool)
	}
	w.cityConnections[cityTwoName][cityOneName] = true
}

func (w *World) RemoveCity(cityNameToRemove string) error {
	if _, exists := w.cities[cityNameToRemove]; !exists {
		return fmt.Errorf("City with name %s does not exist", cityNameToRemove)
	}

	// delete city from world map
	delete(w.cities, cityNameToRemove)

	// delete relevant connections to this city
	connectedCities := w.cityConnections[cityNameToRemove]
	for connection, _ := range connectedCities {
		w.removeConnection(connection, cityNameToRemove)
	}
	delete(w.cityConnections, cityNameToRemove)

	return nil
}

func (w *World) removeConnection(connection, cityNameToRemove string) {
	connectionCity := w.cities[connection]
	if connectionCity.Neighbours[North] != nil && connectionCity.Neighbours[North].Name == cityNameToRemove {
		connectionCity.Neighbours[North] = nil
	} else if connectionCity.Neighbours[East] != nil && connectionCity.Neighbours[East].Name == cityNameToRemove {
		connectionCity.Neighbours[East] = nil
	} else if connectionCity.Neighbours[South] != nil && connectionCity.Neighbours[South].Name == cityNameToRemove {
		connectionCity.Neighbours[South] = nil
	} else if connectionCity.Neighbours[West] != nil && connectionCity.Neighbours[West].Name == cityNameToRemove {
		connectionCity.Neighbours[West] = nil
	}

	delete(w.cityConnections[connection], cityNameToRemove)
}

func (w *World) PrintCitiesTopology() {
	for _, city := range w.cities {
		fmt.Printf("%s : ", city.Name)
		if len(city.Neighbours) != 0 {
			for dir, neighbour := range city.Neighbours {
				if neighbour != nil {
					fmt.Printf("%s=%s, ", dir.String(), neighbour.Name)
				}
			}
		}
		fmt.Println()
	}
}

func (w *World) PrintCitiesConnections() {
	for cityName, connectionCities := range w.cityConnections {
		fmt.Printf("%s is connected to ", cityName)
		for connection, _ := range connectionCities {
			fmt.Printf("%s, ", connection)
		}
		fmt.Println()
	}
}
