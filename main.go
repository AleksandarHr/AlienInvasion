package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AleksandarHr/AlienInvasion/structs"
	"github.com/AleksandarHr/AlienInvasion/utils"
)

const mapFileName = "map.txt"

func main() {
	var aliensCount int

	flag.IntVar(&aliensCount, "N", 4, "Specify number of aliens. Default is 10")
	flag.Parse()

	mapInfo, err := utils.ParseInputFile(mapFileName)
	world := structs.CreateWorld()
	if err != nil {
		// todo: handle err
		// return nil, fmt.Errorf("City already exists in the world.")
	}

	world.InitializeWorld(mapInfo)

	world.PrintCitiesTopology()
	fmt.Println()

	// fmt.Println("Print city connections")
	// world.PrintCitiesConnections()

	// toRemove := "Foo"
	// fmt.Println()
	// fmt.Printf("Remove %s and print connections\n", toRemove)
	// world.RemoveCity(toRemove)
	// world.PrintCitiesConnections()

	// Spawn aliens
	for i := 0; i < aliensCount; i++ {
		alien := structs.CreateAlien(i)
		allCities, err := world.GetAllCities()
		if len(allCities) == 0 {
			fmt.Println("All cities got destroyed while spawning aliens. Exit simulation.")
			os.Exit(0)
		}
		if err != nil {
			// TODO
		}

		randCityIdx, err := utils.GenerateRandomNumber(len(allCities))
		if err != nil {
			// TODO
		}

		originCity := allCities[randCityIdx]
		// NOTE: It is possible to spanw an alien at a city where there already is an alien!
		alien.SpawnAlien(originCity)
		world.AddAlienToCity(alien, nil, originCity)

		world.PrintExistingCities()
		world.PrintAliensInfo()
	}

	// Simulate iterations
	maxIterations := 10
	for {
		// SHOULD THE SIMULATION STOP?
		// If the simulation has ran for 10,000 iterations, exit
		if maxIterations == 0 {
			fmt.Println("Reached maximum number of iterations. Exitting simulation.")
			os.Exit(0)
		}

		// If all the aliens have died, exit
		if world.AllAliensDead() {
			fmt.Println("All aliens have diead. Exitting simulation.")
			os.Exit(0)
		}

		// If all remaining aliens are trapped (e.g. cannot move), exit
		if world.AllAliensTrapped() {
			fmt.Println("All remaining aliens are trapped. Exitting simulation.")
			os.Exit(0)
		}

		// SIMULATION CAN CONTINUE
		currentFreeAliens, _ := world.GetFreeAliens()
		for _, alien := range currentFreeAliens {
			// if the alien got trapped while executing current iterations, continue to next alien
			if !world.IsAlienStillAliveAndFree(alien) {
				continue
			}

			// If alien is free, move to a random neighbouring city
			newAlienCity, _ := alien.PickRandomNeighbourCity()
			// update world information
			// TODO: Change the function name???
			if newAlienCity != nil {
				_, err := world.AddAlienToCity(alien, alien.Location, newAlienCity)
				if err != nil {
					// TODO: handle error
				}
			}
		}
	}
}
