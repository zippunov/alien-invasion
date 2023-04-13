![workflow passed](https://github.com/zippunov/alien-invasion/actions/workflows/go.yml/badge.svg)

# Aliens Invasion

## Table of contents
1. [Assignment](#assignment)
2. [Design and Architecture](#design)
3. [The Project](#project)
	1. [Requirements](#requirements)
	2. [Project Structure](#structure)
	3. [Tests](#tests)
	4. [Build](#build)
	5. [Usage](#usage)

<a name="assignment"></a>
## Assignment

A Group of Space Aliens invades the fantasy World X. Map of World X is represented by the city's directional graph.
Given the number of Aliens wandering through the graph. When two Aliens met with each other they destroy the City they met.
City destruction causes occupying Aliens to get destroyed, as well as all in and out roads connecting the destroyed City.

The complete and more or less formal definition is provided in the [Assignment Document](docs/assignment.md)

<a name="design"></a>
## Design and Architecture

The functionality described in the assignment is stated in general terms leaving space for multiple interpretations.

During the design phase of the project number of decisions were made in order to define every functionality aspect.
Please refer to the [Design and Architecture Document](docs/design.md) for further details about the design process and application
architecture.

<a name="project"></a>
## The Project

<a name="requirements"></a>
### Requirements
The application implementation assumes no use of any 3rd party library staying within the bounds of the Go 1.20 standard library.

- Go language tooling and library [All releases. The Go](https://go.dev/dl/)

<a name="structure"></a>
### Structure

```txt
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── alien-invasion                  // alien-invasion - main app
│   │   └── main.go
│   └── mapgen                          // alien-mapgen utility to generate test map files
│       └── main.go
├── dist                                // Holder for compilation results
├── docs                                // Project documentation
├── go.mod
├── go.sum
├── internal
│   ├── domain                          // package for domain entities
│   │   ├── city.go                     // City entity definition
│   │   ├── city_test.go                // Unit tests
│   │   ├── direction.go                // Direction enum definition
│   │   ├── doc.go                      // package docs
│   │   ├── map.go                      // Map entity definition
│   │   ├── map_test.go                 // Unit tests
│   │   ├── road.go                     // Road entity structure
│   │   └── roadset.go                  // Set of Roads datastructure
│   ├── encoding                        // Package encoding
│   │   ├── text.go                     // Marshalling and Unmarshalling of the map files
│   │   └── text_test.go                // Unit tests
│   ├── infrastructure                  // Package infrastructure
│   │   ├── config.go                   // Infrastructure configuration
│   │   ├── doc.go                      // Package documentation
│   │   └── infra.go                    // Infra struct definitions
│   └── usecases                        // Package usecases
│       ├── main_scenario.go            // Main Scenario Usecase
│       └── main_scenario_test.go       // Unit tests
└── test                                // Generated test maps
```

<a name="tests"></a>
### Tests

```
$ make test
```

<a name="build"></a>
### Build

```
$ make build-all
```

<a name="usage"></a>
### Usage

```
$ ./dist/alien-invasion -h
alien-invasion
Runs the scenario of the alien invasion on the given fantasy map. Prints out resulting cities map.

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
		Print help information
```
