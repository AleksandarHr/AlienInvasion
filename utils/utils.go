package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/phuslu/log"
)

func ParseInputFile(fname string) (map[string][]string, error) {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("%v\n", err)
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
		return r, &GenerateRandomNumberError{bound: n}
	}
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(n)
	return r, nil
}

func InitializeLogger() (log.Logger, log.Logger) {
	defaultLogger := log.Logger{
		TimeFormat: "15:04:05",
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}

	debugLogger := log.Logger{
		TimeFormat: "15:04:05",
		Level:      log.DebugLevel,
		Writer: &log.FileWriter{
			Filename:     "logs/main.DEBUG",
			FileMode:     0600,
			EnsureFolder: true,
			MaxSize:      100 * 1024 * 1024,
		},
	}

	return defaultLogger, debugLogger
}

// error triggered when trying to generate a random number with an invalid upper bound
type GenerateRandomNumberError struct {
	bound int
}

func (err *GenerateRandomNumberError) Error() string {
	return fmt.Sprintf("Invalid random number bound: %d.\n", err.bound)
}
