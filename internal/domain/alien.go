package domain

import "fmt"

// Alien can be represented as an integer ID. It is just enough for all required functionality.
type Alien int

// String is the implementation of the Stringer interface
func (a Alien) String() string {
	return fmt.Sprintf("%d", a)
}
