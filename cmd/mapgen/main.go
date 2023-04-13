/*
mapgen builds random World Map for alien-invasion application

USAGE:

	mapgen [OPTIONS]

	OPTIONS:
		-n <INT>
			Number of cities
		-o <PATH>
			Optional. Output file path. Default output: stdout
		-h
			Print help information
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/zippunov/alien-invasion/internal/domain"
	"github.com/zippunov/alien-invasion/internal/encoding"
	"math/rand"
	"os"

	"html/template"
)

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var helpStr = `{{.Reset}}mapgen
Builds random map for alien-invasion.

`
var usageStr = `{{.Yellow}}USAGE:
	{{.Reset}}mapgen [OPTIONS]

{{.Yellow}}OPTIONS:
	{{.Green}}-n <INT>
		{{.Reset}}Number of cities
	{{.Green}}-o <PATH>
		{{.Reset}}Optional. Output file path. Default output: stdout
	{{.Green}}-h
		{{.Reset}}Print help information
`
var usageTemplate = template.Must(template.New("").Parse(usageStr))
var helpTemplate = template.Must(template.New("").Parse(helpStr + usageStr))
var colors = struct {
	Reset  string
	Yellow string
	Green  string
}{
	Reset:  "\033[0m",
	Yellow: "\033[0;33m",
	Green:  "\033[0;32m",
}
var log = func(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}

func main() {
	var (
		citiesCount int
		outFilePath string
		help        bool
		outFile     *os.File
		err         error
	)
	flag.Usage = func() {
		_ = usageTemplate.Execute(os.Stderr, colors)
	}
	flag.IntVar(&citiesCount, "n", 0, "")
	flag.StringVar(&outFilePath, "o", "", "")
	flag.BoolVar(&help, "h", false, "")
	flag.Parse()

	if help {
		_ = helpTemplate.Execute(os.Stderr, colors)
		os.Exit(0)
	}

	if citiesCount == 0 {
		handleError(errors.New("missing number of cities"))
	}
	if citiesCount > len(letters) {
		handleError(fmt.Errorf("max number of cities is %d", len(letters)))
	}

	if len(outFilePath) != 0 {
		if outFile, err = os.Create(outFilePath); err != nil {
			handleError(err)
		}
	}
	if outFile == nil {
		outFile = os.Stdout
	}
	defer outFile.Close()

	m := generateMap(citiesCount)
	encoding.MarshalTxt(outFile, m)
}

func generateMap(count int) domain.Map {
	m := domain.Map{}
	b := []byte(letters)
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	b = b[:count]
	for i := 0; i < len(b); i++ {
		name := string(b[i])
		otherCities := otherCitiesRandom(b, b[i])
		dirs := randomDirections()
		for i := 0; i < len(dirs) && i < len(otherCities); i++ {
			linkedName := string(otherCities[i])
			m.LinkCities(name, linkedName, dirs[i])
		}
	}
	return m
}

func otherCitiesRandom(b []byte, city byte) []byte {
	result := make([]byte, 0, len(b)-1)
	for _, c := range b {
		if c != city {
			result = append(result, c)
		}
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

func randomDirections() []domain.Direction {
	dirs := []domain.Direction{
		domain.North,
		domain.East,
		domain.South,
		domain.West,
	}
	rand.Shuffle(4, func(i, j int) {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	})
	return dirs[:rand.Intn(4)+1]
}

func handleError(err error) {
	log("%v\n\n", err)
	flag.Usage()
	os.Exit(1)
}
