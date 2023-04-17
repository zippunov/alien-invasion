package domain

// City is node in the Map graph. It keeps information about all incoming and outcoming roads
// as well as list of Aliens occupying this City
type City struct {
	Name      string
	OutRoad   map[Direction]*City
	inInroads RoadSet
	Aliens    []Alien
}

// String is a part of a Stringer interface implementation.
// Prints name of the City
func (c *City) String() string {
	return c.Name
}

// Directions returns list of all directions where roads out of the City available
func (c *City) Directions() []Direction {
	result := make([]Direction, 0, len(c.OutRoad))
	for d := range c.OutRoad {
		result = append(result, d)
	}
	return result
}
