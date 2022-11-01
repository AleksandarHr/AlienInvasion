package cmd

import (
	"fmt"
	"os"

	"github.com/AleksandarHr/AlienInvasion/simulation"
	"github.com/spf13/cobra"
)

var (

	// Used for flags.
	initialAliensCount int
	mapFileName        string
	maxIterations      int

	rootCmd = &cobra.Command{
		Use:   "AlienInvasion",
		Short: "Simulate an alien invasion of a made-up world.",
		Long:  `Create a made-up world and spawn aliens around it. Simulate alien movement around the world.`,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().IntVarP(&initialAliensCount, "alienCount", "N", 5, "Specify number of aliens.")
	rootCmd.Flags().StringVarP(&mapFileName, "mapFileName", "m", "map.txt", "Specify map file name.")
	rootCmd.Flags().IntVarP(&maxIterations, "iterations", "i", 10000, "Specify number of maximum iterations.")
}

func run() {
	simulation, err := simulation.CreateSimulation(initialAliensCount, mapFileName, maxIterations)
	if err != nil {
		fmt.Printf("Error creating a simulation: %v", err)
		os.Exit(1)
	}
	simulation.InitializeSimulation()
	simulation.Run()
}
