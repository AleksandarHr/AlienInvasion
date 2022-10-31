package structs

import "fmt"

// error triggered when trying to add a city when one with the same name already exists
type AddNewCityError struct {
	city *City
}

func (err *AddNewCityError) Error() string {
	return fmt.Sprintf("Cannot add city %s: city with such name already exists.\n", err.city.Name)
}

// error triggered when trying to add a link to an invalid direction
type InvalidDirectionError struct {
	direction Direction
}

func (err *InvalidDirectionError) Error() string {
	return fmt.Sprintf("Invalid direction: %s.\n", err.direction.String())
}

// error triggered when working with an invalid city
type InvalidCityError struct {
	city *City
}

func (err *InvalidCityError) Error() string {
	return fmt.Sprintf("Nil or invalid city: %#v.\n", err.city)
}

// error triggered when working with an invalid alien
type InvalidAlienError struct {
	alien *Alien
}

func (err *InvalidAlienError) Error() string {
	return fmt.Sprintf("Nil or invalid alien: %#v.\n", err.alien)
}

// error triggered when attempting to remove links to a non-existent city
type NonExistentCityError struct {
	cityName string
}

func (err *NonExistentCityError) Error() string {
	return fmt.Sprintf("City %s does not exist.\n", err.cityName)
}

// error triggered when unable to change alien location
type AddAlienToCityError struct {
	alien *Alien
	city  *City
}

func (err *AddAlienToCityError) Error() string {
	return fmt.Sprintf("Cannot move Alien %d to %s.\n", err.alien.ID, err.city.Name)
}
