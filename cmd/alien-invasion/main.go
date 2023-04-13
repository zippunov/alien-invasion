/*
alien-invasion CLI utility run scenario of the alien invasion on the given fantasy map. Prints out resulting cities map.

USAGE:

	alien-invasion [OPTIONS]

OPTIONS:

	-f <PATH>
		File path with the World Map definition
	-n <INT>
		Number of aliens invading World
	-o <PATH>
		Optional. Output file path. Default output: stdout
	-h
		{{.Reset}}Print help information
*/
package main

import (
	"flag"
	"fmt"
	infrastructure2 "github.com/zippunov/alien-invasion/internal/infrastructure"
	"github.com/zippunov/alien-invasion/internal/usecases"
	"html/template"
	"os"
)

var help = `{{.Reset}}alien-invasion
Runs scenario of the alien invasion on the given fantasy map. Prints out resulting cities map.

`

var usage = `{{.Yellow}}USAGE:
	{{.Reset}}alien-invasion [OPTIONS]

{{.Yellow}}OPTIONS:
	{{.Green}}-f <PATH>
		{{.Reset}}File path with the World Map definition
	{{.Green}}-n <INT>
		{{.Reset}}Number of aliens invading World
	{{.Green}}-o <PATH>
		{{.Reset}}Optional. Output file path. Default output: stdout
	{{.Green}}-h
		{{.Reset}}Print help information
`

// template for the Usage output of the application
var usageTemplate = template.Must(template.New("").Parse(usage))

// template for the Help output of the application after error encountered
var helpTemplate = template.Must(template.New("").Parse(help + usage))

// Predefined set of terminal colors instructions
var colors = struct {
	Reset  string
	Yellow string
	Green  string
}{
	Reset:  "\033[0m",
	Yellow: "\033[0;33m",
	Green:  "\033[0;32m",
}

// Custom log function which will be passed to the application logic
var log = func(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}

func main() {
	configFlag()
	config, err := infrastructure2.InitConfig(log)
	if err != nil {
		handleError(err)
	}
	if config.Help {
		_ = helpTemplate.Execute(os.Stderr, colors)
		os.Exit(0)
	}
	infra, err := infrastructure2.InitInfra(config)
	if err != nil {
		handleError(err)
	}
	defer infra.Shutdown()
	scenario, err := usecases.InitScenario(&infra)
	if err != nil {
		handleError(err)
	}
	if err := scenario.Run(); err != nil {
		handleError(err)
	}
	os.Exit(0)
}

// configFlag redefines Usage() behavior of the flag package
// flag.Usage will print custom preformatted usage docs
func configFlag() {
	flag.Usage = func() {
		_ = usageTemplate.Execute(os.Stderr, colors)
	}
}

// handleError default top level error handler
func handleError(err error) {
	log("%v\n\n", err)
	flag.Usage()
	os.Exit(1)
}
