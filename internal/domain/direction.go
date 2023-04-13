package domain

import "strings"

// Direction is an integer value that identifies on of for directions (North, East, South, West).
type Direction uint

// Enumeration of all available directions
const (
	North Direction = iota
	East
	South
	West
)

// dirMap links string names to the Direction enumeration
var dirMap = map[string]Direction{
	"north": North,
	"east":  East,
	"south": South,
	"west":  West,
}

// dirNames links Direction enumeration to their names
var dirNames = map[Direction]string{
	North: "north",
	East:  "east",
	South: "south",
	West:  "west",
}

// DirectionByName fetched Direction by its name.
// Returns (0, false) if given name not found
func DirectionByName(name string) (Direction, bool) {
	result, ok := dirMap[strings.ToLower(name)]
	return result, ok
}

// String is a part of the Stringer interface implementation for the Direction
func (d Direction) String() string {
	return dirNames[d]
}
