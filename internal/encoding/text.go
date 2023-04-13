/*
Package encoding is responsible for marshalling and unmarshalling domain.Map instance.

# Map Text Format

Map format is defined with one city per line. The city name is first, followed by 1-4 directions
(north, south, east, or west). Each one represents a road to another city that lies in that direction.

For example:

	Foo north=Bar west=Baz south=Qu-ux
	Bar south=Foo west=Bee

The city and each of the pairs are separated by a single space, and the directions are separated from their
respective cities with an equals (=) sign.
*/
package encoding

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/zippunov/alien-invasion/internal/domain"
	"io"
	"strings"
)

var newLine = []byte("\n")

// MarshalTxt writes Map into io.Writer instance according to Map text format
func MarshalTxt(w io.Writer, m domain.Map) error {
	for _, city := range m.ListCities() {
		if _, err := w.Write([]byte(city.Name)); err != nil {
			return err
		}
		for dir, neighbor := range city.OutRoad {
			s := fmt.Sprintf(" %v=%v", dir, neighbor)
			if _, err := w.Write([]byte(s)); err != nil {
				return err
			}
		}
		if _, err := w.Write(newLine); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalTxt reads stream formatted according to Map Text Format
// and fills Map with parsed Cities
func UnmarshalTxt(r io.Reader, m domain.Map) error {
	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		line++
		t := scanner.Text()
		if err := parseCityLine(t, m); err != nil {
			return fmt.Errorf("invalid line %d \"%s\": %v", line, t, err)
		}
	}
	return nil
}

// parseCityLine validates single map file line, parses it
// and adds calls Map instance to create and link Map Cities
func parseCityLine(s string, m domain.Map) error {
	tokens := trimAndFilter(strings.Split(s, " "))
	if len(tokens) < 1 {
		return errors.New("empty line")
	}
	if len(tokens) < 2 {
		return errors.New("city must have at least one outgoing road")
	}
	if len(tokens) > 5 {
		return errors.New("city must have at most four outgoing road")
	}
	cityFrom := tokens[0]
	m.InitCity(cityFrom)
	for i := 1; i < len(tokens); i++ {
		roadString := tokens[i]
		roadTokens := trimAndFilter(strings.Split(roadString, "="))
		if len(roadTokens) < 2 {
			return errors.New("invalid neighbor encoding")
		}
		direction, ok := domain.DirectionByName(roadTokens[0])
		if !ok {
			return errors.New("invalid direction name")
		}
		destination := roadTokens[1]
		if err := m.LinkCities(cityFrom, destination, direction); err != nil {
			return err
		}
	}
	return nil
}

// trimAndFilter takes slice of strings, trims each one and returns list of Non-empty string
// the order of the outgoing strings is preserved in the result.
func trimAndFilter(tokens []string) []string {
	validTokens := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		t := strings.TrimSpace(tokens[i])
		if len(t) > 0 {
			validTokens = append(validTokens, t)
		}
	}
	return validTokens
}
