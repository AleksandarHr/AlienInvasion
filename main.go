// import (
// 	"fmt"
// 	"os"

// 	"github.com/AleksandarHr/AlienInvasion/cli"
// 	"github.com/AleksandarHr/AlienInvasion/simulation"
// )

//	func main() {
//		initialAliensCount, mapFileName, maxIterations := cli.ParseCli()
//		simulation, err := simulation.CreateSimulation(initialAliensCount, mapFileName, maxIterations)
//		if err != nil {
//			fmt.Printf("Error creating a simulation: %v", err)
//			os.Exit(1)
//		}
//		simulation.InitializeSimulation()
//		simulation.Run()
//	}
package main

import (
	"github.com/AleksandarHr/AlienInvasion/cmd"
)

func main() {
	cmd.Execute()
}
