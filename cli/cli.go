package cli

import "flag"

func ParseCli() (int, string, int) {
	var initialAliensCount int
	var mapFileName string
	var maxIterations int

	flag.IntVar(&initialAliensCount, "N", 5, "Specify number of aliens.")
	flag.StringVar(&mapFileName, "map", "map.txt", "Specify map file name.")
	flag.IntVar(&maxIterations, "iter", 10000, "Specify number of maximum iterations.")
	flag.Parse()

	return initialAliensCount, mapFileName, maxIterations
}
