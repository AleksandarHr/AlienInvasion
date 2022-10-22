package structs

import (
	"fmt"
	"strings"
)

type World struct {
	// cities within the world
	cities map[string]*City

	// aliens within the world
	aliens map[int]*Alien

	// aliens within the world who are untrapped (e.g. can travel to a neighboring city)
	freeAliens map[int]*Alien

	// city names mapped to alien IDs
	citiesAliens map[string]*Alien

	// map a city name to the names of all the cities it is linked to
	cityConnections map[string]map[string]bool
}

func CreateWorld() *World {
	return &World{
		cities:          make(map[string]*City),
		aliens:          make(map[int]*Alien),
		freeAliens:      make(map[int]*Alien),
		cityConnections: make(map[string]map[string]bool),
		citiesAliens:    make(map[string]*Alien),
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

// InitializeWorld creates the world based on information from the map file
func (w *World) InitializeWorld(mapInfo map[string][]string) {
	for cityName, neighbourNames := range mapInfo {
		city := w.AddNewCity(cityName)

		for _, neighbourInfo := range neighbourNames {
			temp := strings.Split(neighbourInfo, "=")
			neighbourDirection, neighbourName := temp[0], temp[1]
			var neighbourCity *City
			if nCity, exists := w.cities[neighbourName]; !exists {
				neighbourCity = w.AddNewCity(neighbourName)
			} else {
				neighbourCity = nCity
			}

			// add neighbour info to current city
			city.Neighbours[StringToDirection(neighbourDirection)] = neighbourCity

			// add reverse neighbour info to neighbour city
			neighbourCity.Neighbours[StringToDirection(neighbourDirection).OppositeDirection()] = city

			// add relevant links information
			w.addLinks(cityName, neighbourName)
		}
	}
}

// addLinks populates informatino about links between a given pair of cities
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

// RemoveCity removes the given city and any related information about it from the world
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

// removeConnection deletes information about connection between the two cities
func (w *World) removeConnection(connection, cityNameToRemove string) {
	connectionCity := w.cities[connection]
	if connectionCity.Neighbours[North] != nil && connectionCity.Neighbours[North].Name == cityNameToRemove {
		delete(connectionCity.Neighbours, North)
	} else if connectionCity.Neighbours[East] != nil && connectionCity.Neighbours[East].Name == cityNameToRemove {
		delete(connectionCity.Neighbours, East)
	} else if connectionCity.Neighbours[South] != nil && connectionCity.Neighbours[South].Name == cityNameToRemove {
		delete(connectionCity.Neighbours, South)
	} else if connectionCity.Neighbours[West] != nil && connectionCity.Neighbours[West].Name == cityNameToRemove {
		delete(connectionCity.Neighbours, West)
	}

	delete(w.cityConnections[connection], cityNameToRemove)
	// if the connectionCity is left with no connections, update freeAliens map if necessary
	if !connectionCity.HasNeighbours() {
		if alien, ok := w.citiesAliens[connectionCity.Name]; ok {
			delete(w.freeAliens, alien.ID)
		}
	}
}

// AddNewAlienToWorld attempts to spawn a new alien at origin city.
// Returns false if spawning was not successful (e.g. there already was an alien there)
func (w *World) AddAlienToCity(newAlien *Alien, from *City, to *City) (bool, error) {
	// the city already has an alien there
	if existingAlien, hasAlien := w.citiesAliens[to.Name]; hasAlien {
		// kill existing alien
		delete(w.aliens, existingAlien.ID)
		if _, ok := w.freeAliens[existingAlien.ID]; ok {
			delete(w.freeAliens, existingAlien.ID)
		}

		// kill new alien
		if _, ok := w.aliens[newAlien.ID]; ok {
			delete(w.aliens, newAlien.ID)
		}
		if _, ok := w.freeAliens[existingAlien.ID]; ok {
			delete(w.freeAliens, existingAlien.ID)
		}

		// remove mapping from city to alien in the city
		delete(w.citiesAliens, to.Name)
		if from != nil {
			if _, ok := w.citiesAliens[from.Name]; ok {
				delete(w.citiesAliens, from.Name)
			}
		}

		// destroy city
		w.RemoveCity(to.Name)

		fmt.Printf("%s has been destroyed by alien %d and alien %d\n", to.Name, existingAlien.ID, newAlien.ID)
		return false, nil
	}

	if from != nil {
		if _, ok := w.citiesAliens[from.Name]; ok {
			delete(w.citiesAliens, from.Name)
		}
	}
	// there is not alien in the origin city, spawn the new alien there
	newAlien.Location = to
	w.aliens[newAlien.ID] = newAlien
	w.citiesAliens[to.Name] = newAlien
	if to.HasNeighbours() {
		w.freeAliens[newAlien.ID] = newAlien
	}

	return true, nil
}

func (w *World) GetAllCities() ([]*City, error) {
	cities := []*City{}
	for _, city := range w.cities {
		cities = append(cities, city)
	}
	return cities, nil
}

func (w *World) GetFreeAliens() ([]*Alien, error) {
	currentFreeAliens := []*Alien{}
	for _, alien := range w.freeAliens {
		currentFreeAliens = append(currentFreeAliens, alien)
	}
	return currentFreeAliens, nil
}

func (w *World) IsAlienStillAliveAndFree(alien *Alien) bool {
	if _, exists := w.freeAliens[alien.ID]; exists {
		return true
	}
	return false
}

// =========================================================================================
// Print Helpers
// =========================================================================================
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

func (w *World) PrintAliensInfo() {
	for id, alien := range w.aliens {
		fmt.Printf("Alien %d is in %s\n", id, alien.Location.Name)
	}
	fmt.Println()
}

func (w *World) PrintExistingCities() {
	fmt.Print("Remaining cities: ")
	for cityName, _ := range w.cities {
		fmt.Printf("%s, ", cityName)
	}
	fmt.Println()
}

func (w *World) AllAliensDead() bool {
	return len(w.aliens) == 0
}

func (w *World) AllAliensTrapped() bool {
	return len(w.freeAliens) == 0
}
