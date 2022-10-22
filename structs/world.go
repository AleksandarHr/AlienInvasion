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
}

func CreateWorld() *World {
	return &World{
		cities: make(map[string]*City),
		aliens: make(map[string]*Alien),
	}
}

func (w *World) AddNewCity(name string) (*City, error) {
	// if city has already been added to the world
	if _, exists := w.cities[name]; exists {
		// todo: handle error
		return nil, fmt.Errorf("City already exists in the world.")
	}

	newCity := CreateCity(name)
	w.cities[name] = newCity
	return newCity, nil
}

func (w *World) InitializeWorld(mapInfo map[string][]string) {
	for cityName, neighbourNames := range mapInfo {
		city, exists := w.cities[cityName]
		if !exists {
			city = CreateCity(cityName)
			w.cities[cityName] = city
		} else {
			city = w.cities[cityName]
		}

		for _, neighbourInfo := range neighbourNames {
			temp := strings.Split(neighbourInfo, "=")
			neighbourDirection, neighbourName := temp[0], temp[1]

			neighbourCity, exists := w.cities[neighbourName]
			if !exists {
				neighbourCity = CreateCity(neighbourName)
				w.cities[neighbourName] = neighbourCity
			} else {
				neighbourCity = w.cities[neighbourName]
			}

			// add neighbour info to current city
			city.Neighbours[StringToDirection(neighbourDirection)] = neighbourCity

			// add reverse neighbour info to neighbour city
			neighbourCity.Neighbours[StringToDirection(neighbourDirection).OppositeDirection()] = city
		}
	}
}

func (w *World) PrintCitiesTopology() {
	for _, city := range w.cities {
		fmt.Printf("%s : ", city.Name)
		if len(city.Neighbours) != 0 {
			for dir, neighbour := range city.Neighbours {
				fmt.Printf("%s=%s, ", dir.String(), neighbour.Name)
			}
		}
		fmt.Println()
	}
}
