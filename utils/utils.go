package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func GenerateRandomNumber(n int) (int, error) {
	r := 0
	if n <= 0 {
		return r, fmt.Errorf("Invalid bounds for RNG.")
	}
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(n)
	return r, nil
}
