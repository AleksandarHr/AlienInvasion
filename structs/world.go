package structs

import "fmt"

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

func (w *World) InitializeWorld(mapInfo map[string][]string) {

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
