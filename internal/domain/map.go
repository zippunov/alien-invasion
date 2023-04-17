package domain

import (
	"fmt"
	"sort"
)

// Map is the representation of the Cities graph. It is responsible for indexing Cities by name,
// and provides main operations over the graph.
type Map map[string]*City

// ListCities returns list of the City entities sorted by City name.
func (m *Map) ListCities() []*City {
	result := make([]*City, 0, len(*m))
	for _, c := range *m {
		result = append(result, c)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

// LinkCities links two Cities with roads. Roads are properly registered as in-road or out-road
// in the correspondent City. If the City with given name is not found in the Map, then new City
// will be created and indexed with the name.
func (m *Map) LinkCities(from, to string, direction Direction) error {
	if err := m.validateLink(from, to, direction); err != nil {
		return err
	}
	sourceCity, targetCity := m.InitCity(from), m.InitCity(to)
	sourceCity.OutRoad[direction] = targetCity
	targetCity.inInroads.add(road{sourceCity, direction})
	return nil
}

// validateLink ensures that given parameters are valid for creation Cities link.
// It is impossible to create two out-road from the City in the same Direction
// It is impossible to create road from the City to itself.
func (m *Map) validateLink(from, to string, direction Direction) error {
	city, ok := (*m)[from]
	if !ok {
		return nil
	}
	if _, ok := city.OutRoad[direction]; ok {
		return fmt.Errorf("direction %v is taken for city %v", direction, city)
	}
	if from == to {
		return fmt.Errorf("attempt to link city %v to itself", city)
	}
	return nil
}

// DestroyCity removes City from the map. All ingoing and outgoing roads will be deleted in all linked cities.
func (m *Map) DestroyCity(city *City) {
	for dir, c := range city.OutRoad {
		c.inInroads.remove(road{c, dir})
	}
	for r := range city.inInroads {
		delete(r.from.OutRoad, r.direction)
	}
	delete(*m, city.Name)
}

// InitCity creates empty City instance with given name and indexes City in the Map
func (m *Map) InitCity(name string) *City {
	city, ok := (*m)[name]
	if ok {
		return city
	}
	(*m)[name] = &City{
		Name:      name,
		OutRoad:   map[Direction]*City{},
		inInroads: RoadSet{},
		Aliens:    []Alien{},
	}
	return (*m)[name]
}
