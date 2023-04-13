package domain

// road holds information from which City and in what direction it goes FROM the City
type road struct {
	from      *City
	direction Direction
}
