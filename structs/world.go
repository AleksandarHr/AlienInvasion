package structs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phuslu/log"
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

// CreateWorld construct a new world
func CreateWorld() *World {
	return &World{
		cities:          make(map[string]*City),
		aliens:          make(map[int]*Alien),
		freeAliens:      make(map[int]*Alien),
		cityConnections: make(map[string]map[string]bool),
		citiesAliens:    make(map[string]*Alien),
	}
}

// InitializeWorld creates the world based on information from the map file
func (w *World) InitializeWorld(mapInfo map[string][]string) {
	for cityName := range mapInfo {
		// add a city
		newCity := CreateCity(cityName)
		err := w.AddNewCity(newCity)
		if err != nil || cityName == "" {
			_ = fmt.Errorf("Unable to add invalid city during initialization: %v. Continue to next city.\n", newCity)
			continue
		}
	}

	for cityName, neighbourNames := range mapInfo {
		// add it's neighbours and links
		for _, neighbourInfo := range neighbourNames {
			temp := strings.Split(neighbourInfo, "=")
			neighbourDirection, neighbourName := temp[0], temp[1]
			if _, exists := w.cities[neighbourName]; !exists {
				// should NOT happen based on map.txt contents assumption
				continue
			}
			neighbourCity := w.cities[neighbourName]
			err := w.AddNewCity(neighbourCity)
			if err != nil || neighbourName == "" {
				_ = fmt.Errorf("Unable to add invalid city during initialization: %v. Continue to next city.\n", neighbourCity)
				continue
			}

			// add neighbour info to current city
			newCity := w.cities[cityName]
			newCity.Neighbours[StringToDirection(neighbourDirection)] = neighbourCity

			// add relevant links information
			w.addLinks(cityName, neighbourName)
		}
	}
}

// AddNewCity creates a new city with the specified name and adds it to the world
func (w *World) AddNewCity(newCity *City) error {
	if newCity == nil {
		return &InvalidCityError{city: newCity}
	}
	// if city has already been added to the world
	if _, exists := w.cities[newCity.Name]; !exists {
		w.cities[newCity.Name] = newCity
	}
	return nil
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
func (w *World) RemoveCity(cityToRemove *City) error {
	if cityToRemove == nil {
		return &InvalidCityError{city: cityToRemove}
	}

	cityNameToRemove := cityToRemove.Name
	if _, exists := w.cities[cityNameToRemove]; !exists {
		return &NonExistentCityError{cityName: cityNameToRemove}
	}

	// delete city from world map
	if _, ok := w.cities[cityNameToRemove]; ok {
		delete(w.cities, cityNameToRemove)
	}

	// delete relevant connections to this city
	if connectedCities, exists := w.cityConnections[cityNameToRemove]; exists {
		for connection, _ := range connectedCities {
			w.removeConnection(connection, cityNameToRemove)
		}
	}

	if _, ok := w.cityConnections[cityNameToRemove]; ok {
		delete(w.cityConnections, cityNameToRemove)
	}

	// remove mapping from city to alien in the city
	if _, ok := w.citiesAliens[cityNameToRemove]; ok {
		delete(w.citiesAliens, cityNameToRemove)
	}

	return nil
}

// removeConnection deletes information about connection between the two cities
func (w *World) removeConnection(connection, cityNameToRemove string) error {
	if _, exists := w.cities[connection]; !exists {
		return &NonExistentCityError{cityName: connection}
	}
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

	if connections, exists := w.cityConnections[connection]; exists {
		if _, exists = connections[cityNameToRemove]; exists {
			delete(w.cityConnections[connection], cityNameToRemove)
		}
	}
	// if the connectionCity is left with no connections, update freeAliens map if necessary
	if !connectionCity.HasNeighbours() {
		if alien, exists := w.citiesAliens[connectionCity.Name]; exists {
			delete(w.freeAliens, alien.ID)
		}
	}
	return nil
}

// killAlien destroys the given alien by removing references from relevant variables
func (w *World) killAlien(alien *Alien) error {
	if alien == nil {
		return &InvalidAlienError{alien: alien}
	}

	if _, exists := w.aliens[alien.ID]; exists {
		delete(w.aliens, alien.ID)
	}
	if _, exists := w.freeAliens[alien.ID]; exists {
		delete(w.freeAliens, alien.ID)
	}
	return nil
}

// AddNewAlienToCity attempts to add the given alien to the specified city.
// If the city already has an alien in it, the move is unsuccessful
func (w *World) AddAlienToCity(alien *Alien, to *City, stage SimulationStage) (bool, error) {
	if alien == nil {
		return false, &InvalidAlienError{alien: alien}
	}
	if to == nil {
		return false, &InvalidCityError{city: to}
	}
	if _, ok := w.cities[to.Name]; !ok {
		return false, &NonExistentCityError{cityName: to.Name}
	}

	// If the alien is already present in a different city, update cities-to-aliens map
	if stage != SpawningAliens {
		if _, exists := w.citiesAliens[alien.Location.Name]; exists {
			delete(w.citiesAliens, alien.Location.Name)
		}
	}

	// if the city already has an alien there
	if existingAlien, hasAlien := w.citiesAliens[to.Name]; hasAlien {

		// kill existing alien
		w.killAlien(existingAlien)
		// kill new alien
		w.killAlien(alien)

		// destroy city
		w.RemoveCity(to)

		fmt.Printf("%s has been destroyed by alien %d and alien %d\n", to.Name, existingAlien.ID, alien.ID)
		return false, nil
	}

	// there is not alien in the origin city, spawn the new alien there
	_ = w.updateAlienLocation(alien, to)
	return true, nil
}

// updateAlienLocation moves an alien to a new city and updates relevant information
func (w *World) updateAlienLocation(alien *Alien, newCity *City) error {
	if alien == nil {
		return &InvalidAlienError{alien: alien}
	}
	if newCity == nil {
		return &InvalidCityError{city: newCity}
	}

	// move alien to new city
	alien.MoveToCity(newCity)

	// update relevant variables
	w.aliens[alien.ID] = alien
	w.citiesAliens[newCity.Name] = alien
	if newCity.HasNeighbours() {
		w.freeAliens[alien.ID] = alien
	}
	return nil
}

// GetAllCities returns all currentlt existing cities
func (w *World) GetAllCities() ([]*City, error) {
	cities := []*City{}
	for _, city := range w.cities {
		cities = append(cities, city)
	}
	return cities, nil
}

// GetFreeAliens returns all aliens which are not trapped
func (w *World) GetFreeAliens() ([]*Alien, error) {
	currentFreeAliens := []*Alien{}
	for _, alien := range w.freeAliens {
		currentFreeAliens = append(currentFreeAliens, alien)
	}
	return currentFreeAliens, nil
}

// IsAlienFree checks if a given alien is still free
func (w *World) IsAlienFree(alien *Alien) (bool, error) {
	if alien == nil {
		return false, &InvalidAlienError{alien: alien}
	}

	if _, exists := w.freeAliens[alien.ID]; exists {
		return true, nil
	}
	return false, nil
}

// IsAlienAlive checks if a given alien is still alive
func (w *World) IsAlienAlive(alien *Alien) (bool, error) {
	if alien == nil {
		return false, &InvalidAlienError{alien: alien}
	}

	if _, exists := w.aliens[alien.ID]; exists {
		return true, nil
	}
	return false, nil
}

// AllAliensDead checks if all aliens have died
func (w *World) AllAliensDead() bool {
	return len(w.aliens) == 0
}

// AllAliensTrapped checks if all remaining aliens are trapped
func (w *World) AllAliensTrapped() bool {
	return len(w.freeAliens) == 0
}

// =========================================================================================
// Print Helpers
// =========================================================================================
func (w *World) LogWorldState(debugLogger log.Logger) {
	w.printCitiesTopology(debugLogger)
	w.printCitiesConnections(debugLogger)
	w.printExistingCities(debugLogger)
	w.printAliensInfo(debugLogger)
}

func (w *World) printCitiesTopology(debugLogger log.Logger) {
	for _, city := range w.cities {
		var topology strings.Builder
		topology.WriteString(city.Name + ": ")
		if len(city.Neighbours) != 0 {
			for dir, neighbour := range city.Neighbours {
				if neighbour != nil {
					topology.WriteString(dir.String() + "=" + neighbour.Name)
				}
			}
		}
		debugLogger.Debug().Msgf("Topology: %s", topology.String())
	}
}

func (w *World) printCitiesConnections(debugLogger log.Logger) {
	for cityName, connectionCities := range w.cityConnections {
		var connections strings.Builder
		connections.WriteString(cityName + " is connected to: ")
		for connection, _ := range connectionCities {
			connections.WriteString(connection + ", ")
		}
		debugLogger.Debug().Msgf("Connections: %s", connections.String())
	}
}

func (w *World) printAliensInfo(debugLogger log.Logger) {
	var aliens strings.Builder
	for id, alien := range w.aliens {
		if _, isFree := w.freeAliens[alien.ID]; isFree {
			aliens.WriteString("Alien " + strconv.Itoa(id) + " is in " + alien.Location.Name + ", ")
		} else {
			aliens.WriteString("Alien " + strconv.Itoa(id) + " is TRAPPED in " + alien.Location.Name + ", ")
		}
	}
	debugLogger.Debug().Msgf("Aliens: %s", aliens.String())
}

func (w *World) printExistingCities(debugLogger log.Logger) {
	var cities strings.Builder

	for cityName, _ := range w.cities {
		cities.WriteString(", ")
		cities.WriteString(cityName)
	}
	debugLogger.Debug().Msgf("Remaining cities: %s", cities.String())
}
