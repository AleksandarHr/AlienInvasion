package main

import (
	"errors"
	"flag"
	"os"

	"github.com/AleksandarHr/AlienInvasion/structs"
	"github.com/AleksandarHr/AlienInvasion/utils"
)

func main() {
	defaultLogger, debugLogger := utils.InitializeLogger()
	var aliensCount int
	var mapFileName string

	flag.IntVar(&aliensCount, "N", 4, "Specify number of aliens. Default is 10")
	flag.StringVar(&mapFileName, "map", "map.txt", "Specify map file name. Default is map.txt")
	flag.Parse()

	mapInfo, err := utils.ParseInputFile(mapFileName)
	world := structs.CreateWorld()
	if err != nil {
		// todo: handle err
		// return nil, fmt.Errorf("City already exists in the world.")
	}

	world.InitializeWorld(mapInfo)

	world.PrintCitiesTopology(debugLogger)
	world.PrintCitiesConnections(debugLogger)

	// Spawn aliens
	for i := 0; i < aliensCount; i++ {
		alien := structs.CreateAlien(i)

		allCities, err := world.GetAllCities()
		if len(allCities) == 0 {
			defaultLogger.Info().Msg("All cities got destroyed while spawning aliens. Exitting simulation.")
			os.Exit(0)
		}
		if err != nil {
			defaultLogger.Err(errors.New("Error retrieving available cities. Exitting simulation."))
			os.Exit(0)
		}

		randCityIdx, err := utils.GenerateRandomNumber(len(allCities))
		if err != nil {
			defaultLogger.Err(errors.New("Error choosing a random spawn city. Exitting simulation."))
			os.Exit(0)
		}

		originCity := allCities[randCityIdx]
		// NOTE: It is possible to spanw an alien at a city where there already is an alien!
		added, err := world.AddAlienToCity(alien, originCity)
		if added {
			defaultLogger.Info().Msgf("%s spawned in %s.", alien.Name, originCity.Name)
		}

		world.PrintExistingCities(debugLogger)
		world.PrintAliensInfo(debugLogger)
	}

	// Simulate iterations
	maxIterations := 10
	for {
		// SHOULD THE SIMULATION STOP?
		// If the simulation has ran for 10,000 iterations, exit
		if maxIterations == 0 {
			defaultLogger.Info().Msg("Reached maximum number of iterations. Exitting simulation.")
			os.Exit(0)
		}

		// If all the aliens have died, exit
		if world.AllAliensDead() {
			defaultLogger.Info().Msg("All aliens have died. Exitting simulation.")
			os.Exit(0)
		}

		// If all remaining aliens are trapped (e.g. cannot move), exit
		if world.AllAliensTrapped() {
			defaultLogger.Info().Msg("All aliens are trapped in isolated cities. Exitting simulation.")
			os.Exit(0)
		}

		// SIMULATION CAN CONTINUE
		currentFreeAliens, _ := world.GetFreeAliens()
		for _, alien := range currentFreeAliens {

			// if the alien got trapped while executing current iterations, continue to next alien
			if !world.IsAlienStillAliveAndFree(alien) {
				debugLogger.Info().Msgf("Alien %d got trapped.", alien.ID)
				continue
			}

			// If alien is free, move to a random neighbouring city
			newAlienCity, _ := alien.PickRandomNeighbourCity()
			// update world information
			// TODO: Change the function name???
			if newAlienCity != nil {
				_, err := world.AddAlienToCity(alien, newAlienCity)
				if err != nil {
					// TODO: handle error
				}
			}
		}

		maxIterations--
	}
}
