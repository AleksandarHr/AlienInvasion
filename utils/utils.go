package utils

import (
	"bufio"
	"os"
	"strings"
)

func ParseInputFile(fname string) (map[string][]string, error) {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// store the lines from the map text file into a map
	mapInfo := make(map[string][]string)

	for fileScanner.Scan() {
		cityInfo := strings.Split(fileScanner.Text(), " ")
		mapInfo[cityInfo[0]] = cityInfo[1:]
	}

	return mapInfo, nil
}
