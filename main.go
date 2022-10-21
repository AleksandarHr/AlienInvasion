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

}
