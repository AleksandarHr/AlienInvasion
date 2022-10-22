package main

import (
	"flag"

	"github.com/AleksandarHr/AlienInvasion/structs"
	"github.com/AleksandarHr/AlienInvasion/utils"
)

const mapFileName = "map.txt"

func main() {
	var aliensCount int

	flag.IntVar(&aliensCount, "N", 10, "Specify number of aliens. Default is 10")
	flag.Parse()

	mapInfo, err := utils.ParseInputFile(mapFileName)
	world := structs.CreateWorld()
	if err != nil {
		// todo: handle err
		// return nil, fmt.Errorf("City already exists in the world.")
	}

	world.InitializeWorld(mapInfo)

	// world.PrintCitiesTopology()

	// fmt.Println()
	// fmt.Println("Print city connections")
	// world.PrintCitiesConnections()

	// toRemove := "Foo"
	// fmt.Println()
	// fmt.Printf("Remove %s and print connections\n", toRemove)
	// world.RemoveCity(toRemove)
	// world.PrintCitiesConnections()

	for i := 0; i < aliensCount; i++ {
		alien := structs.CreateAlien(i)
		allCities, err := world.GetAllCities()
		if err != nil {
			// TODO
		}
		randCityIdx, err := utils.GenerateRandomNumber(len(allCities))
		if err != nil {
			// TODO
		}
		originCity := allCities[randCityIdx]
		alien.SpawnAlien(originCity)
	}
}
