package simulation

import (
	"errors"
	"os"

	"github.com/AleksandarHr/AlienInvasion/structs"
	"github.com/AleksandarHr/AlienInvasion/utils"
	"github.com/phuslu/log"
)

type Simulation struct {
	initialAliensCount int

	maxIterations int

	mapInfo map[string][]string

	world *structs.World

	defaultLogger log.Logger

	debugLogger log.Logger

	stage structs.SimulationStage
}

func CreateSimulation(aliens int, mapFileName string, maxIterations int) (*Simulation, error) {
	defaultLog, debugLog := utils.InitializeLogger()
	worldMap, err := utils.ParseInputFile(mapFileName)
	if err != nil {
		return nil, err
	}

	world := structs.CreateWorld()

	return &Simulation{
		initialAliensCount: aliens,
		maxIterations:      maxIterations,
		mapInfo:            worldMap,
		world:              world,
		defaultLogger:      defaultLog,
		debugLogger:        debugLog,
		stage:              structs.SimulationStart,
	}, nil
}

func (s *Simulation) InitializeSimulation() {
	s.stage = structs.InitializingWorld

	s.world.InitializeWorld(s.mapInfo)

	s.stage = structs.SpawningAliens
	// Spawn aliens
	for i := 0; i < s.initialAliensCount; i++ {
		alien := structs.CreateAlien(i)

		allCities, err := s.world.GetAllCities()
		if err != nil {
			s.defaultLogger.Err(errors.New("Error retrieving available cities. Exitting simulation."))
			s.debugLogger.Err(errors.New("Error retrieving available cities. Exitting simulation."))
			os.Exit(1)
		}
		if len(allCities) == 0 {
			s.defaultLogger.Info().Msg("No cities left in the world. Exitting simulation.")
			s.debugLogger.Info().Msg("No cities left in the world. Exitting simulation.")
			os.Exit(0)
		}

		randCityIdx, err := utils.GenerateRandomNumber(len(allCities))
		if err != nil {
			s.defaultLogger.Err(errors.New("Error choosing a random spawn city. Exitting simulation."))
			s.debugLogger.Info().Msg("No cities left in the world. Exitting simulation.")
			os.Exit(1)
		}

		originCity := allCities[randCityIdx]
		// NOTE: It is possible to spanw an alien at a city where there already is an alien!
		added, err := s.world.AddAlienToCity(alien, originCity, s.stage)
		if err != nil {
			s.defaultLogger.Err(err).Msgf("Cannot add an alien to an invalid city %v.", originCity)
			s.debugLogger.Err(err).Msgf("Cannot add an alien to an invalid city %v.", originCity)
			continue
		}
		if added {
			s.defaultLogger.Info().Msgf("Alien %d spawned in %s.", alien.ID, originCity.Name)
			s.debugLogger.Info().Msgf("Alien %d spawned in %s.", alien.ID, originCity.Name)
		} else {
			s.defaultLogger.Info().Msgf("Alien %d tried to spawn in %s where an alien already exists. %s was destroyed.", alien.ID, originCity.Name, originCity.Name)
			s.debugLogger.Info().Msgf("Alien %d tried to spawn in %s where an alien already exists. %s was destroyed.", alien.ID, originCity.Name, originCity.Name)
		}

		s.world.LogWorldState(s.debugLogger)
	}
}

func (s *Simulation) Run() {
	// Simulate iterations
	iteration := 0
	s.stage = structs.MovingAliens
	for {
		// If the simulation has ran for maxIterations number of iterations, exit
		if iteration == s.maxIterations {
			s.defaultLogger.Info().Msg("Reached maximum number of iterations. Exitting simulation.")
			s.debugLogger.Info().Msg("Reached maximum number of iterations. Exitting simulation.")
			os.Exit(0)
		}

		// If all the aliens have died, exit
		if s.world.AllAliensDead() {
			s.defaultLogger.Info().Msg("All aliens have died. Exitting simulation.")
			s.debugLogger.Info().Msg("All aliens have died. Exitting simulation.")
			os.Exit(0)
		}

		// If all remaining aliens are trapped (e.g. cannot move), exit
		if s.world.AllAliensTrapped() {
			s.defaultLogger.Info().Msg("All aliens are trapped in isolated cities. Exitting simulation.")
			s.debugLogger.Info().Msg("All aliens are trapped in isolated cities. Exitting simulation.")
			os.Exit(0)
		}

		// Otherwise, continue the simulation
		currentFreeAliens, _ := s.world.GetFreeAliens()
		for _, alien := range currentFreeAliens {

			// if the alien died while executing current iterations, continue to next alien
			alive, err := s.world.IsAlienAlive(alien)
			if err != nil {
				// invalid alien, simply move to the next one
				continue
			}
			if !alive {
				s.debugLogger.Debug().Msgf("Alien %d died during current iteration.", alien.ID)
				continue
			}

			// if the alien got trapped while executing current iterations, continue to next alien
			free, err := s.world.IsAlienFree(alien)
			if err != nil {
				// invalid alien, simply move to the next one
				continue
			}
			if !free {
				s.debugLogger.Debug().Msgf("Alien %d got trapped in %s during current iteration. No valid move.", alien.ID, alien.Location.Name)
				continue
			}

			// If alien is free, move to a random neighbouring city
			newAlienCity, err := alien.PickRandomNeighbourCity()
			if err != nil {
				// unable to pick a random neighbour city
				s.debugLogger.Debug().Msgf("Error trying to move alien %d to a random neighbour: %v", err)
				continue
			}
			if newAlienCity == nil {
				// alien is trapped
				s.debugLogger.Debug().Msgf("Alien %d is trapped in %s.", alien.ID, alien.Location.Name)
				continue
			}

			// move alien to neighbour and update world information
			added, err := s.world.AddAlienToCity(alien, newAlienCity, s.stage)
			if err != nil {
				// error trying to move alien to city, simply continue with next alien
				s.debugLogger.Debug().Msgf("Unable to move alien %d to a random neighbour: %v", err)
				continue
			}
			if added {
				s.defaultLogger.Info().Msgf("Alien %d moved to %s.", alien.ID, newAlienCity.Name)
				s.debugLogger.Info().Msgf("Alien %d moved to %s.", alien.ID, newAlienCity.Name)
			} else {
				s.defaultLogger.Info().Msgf("Alien %d tried to move where an alien already exists. %s was destroyed.", alien.ID, newAlienCity.Name)
				s.debugLogger.Info().Msgf("Alien %d tried to move where an alien already exists. %s was destroyed.", alien.ID, newAlienCity.Name)
			}
		}

		iteration++
	}
}
