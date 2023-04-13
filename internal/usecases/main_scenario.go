/*
Package usecases contains all available application logic scenarios.
*/
package usecases

import (
	"fmt"
	domain2 "github.com/zippunov/alien-invasion/internal/domain"
	"github.com/zippunov/alien-invasion/internal/encoding"
	"io"
	"math/rand"
)

// IInfra interface specifies all required functionality from the application environment.
// Varios IInfra can be injected into Scenario in order to provide better testing.
type IInfra interface {
	In() io.Reader
	Out() io.Writer
	AliensCount() int
	Log() func(format string, a ...any)
}

// Scenario is the usecase where Alien Invasion scenario is getting executed.
//
// Aliens are represented by int number from 0 to aliensCount-1
type Scenario struct {
	out         io.Writer
	worldMap    domain2.Map                   // Cities graph
	aliensCount int                           // start Aliens count
	aliens      map[int]*domain2.City         // maps each alien to a single City
	movesLeft   []int                         // holds number of moves left for each alien by the Alien integer id.
	log         func(format string, a ...any) // logger function
}

// InitScenario scenario initialization with provided infrastructure
func InitScenario(infra IInfra) (Scenario, error) {
	m := domain2.Map{}
	if err := encoding.UnmarshalTxt(infra.In(), m); err != nil {
		return Scenario{}, err
	}
	n := infra.AliensCount()
	if len(m) < n {
		return Scenario{}, fmt.Errorf("aliens count is greater than number of  cities (%d)", len(m))
	}

	aliens := make(map[int]*domain2.City, n)
	movesLeft := make([]int, n)
	for i := 0; i < n; i++ {
		aliens[i] = nil
		movesLeft[i] = 10000
	}
	return Scenario{
		out:         infra.Out(),
		aliensCount: n,
		worldMap:    m,
		aliens:      aliens,
		movesLeft:   movesLeft,
		log:         infra.Log(),
	}, nil
}

// Run executes the Usecase
func (s *Scenario) Run() error {
	s.seedAliens()
	// queue of every alien with moves left randomized
	q := s.aliensQueue()
	// Main loop
	// - pull Alien from the queue
	// - move Alien in random available direction
	// - destroy city if conditions met
	// - mark Alien moved
	for len(q) > 0 {
		for _, alien := range q {
			if _, ok := s.aliens[alien]; !ok {
				continue
			}
			if s.movesLeft[alien] == 0 {
				continue
			}
			newCity := s.moveAlien(alien)
			if newCity == nil {
				s.movesLeft[alien] = 0
			} else {
				s.destroyCity(newCity)
			}
		}
		q = s.aliensQueue()
	}
	// Output resulting Map
	return encoding.MarshalTxt(s.out, s.worldMap)
}

// seedAliens assings single Alien to a random City
func (s *Scenario) seedAliens() {
	cities := s.worldMap.ListCities()
	rand.Shuffle(len(cities), func(i, j int) {
		cities[i], cities[j] = cities[j], cities[i]
	})
	for i := 0; i < s.aliensCount; i++ {
		s.aliens[i] = cities[i]
		s.aliens[i].Aliens = append(s.aliens[i].Aliens, i)
	}
}

// aliensQueue filters all Aliens that able to make a move and returns filtered Aliens in random order.
func (s *Scenario) aliensQueue() []int {
	queue := make([]int, 0, len(s.aliens))
	for alien := range s.aliens {
		if s.movesLeft[alien] > 0 {
			queue = append(queue, alien)
		}
	}
	rand.Shuffle(len(queue), func(i, j int) {
		queue[i], queue[j] = queue[j], queue[i]
	})
	return queue
}

// moveAlien executed single Alien move. The move Direction os randomly chosen among available out-roads in the City
func (s *Scenario) moveAlien(alien int) *domain2.City {
	city := s.aliens[alien]
	directions := city.Directions()
	dirCount := len(directions)
	if dirCount == 0 {
		return nil
	}
	d := directions[rand.Intn(dirCount)]
	nextCity := city.OutRoad[d]
	nextCity.Aliens = append(nextCity.Aliens, alien)
	city.Aliens = city.Aliens[:0]
	s.aliens[alien] = nextCity
	s.movesLeft[alien] -= 1
	return nextCity
}

// destroyCity removes City and occupying Aliens from the Map if there are 2 Aliens in the City
func (s *Scenario) destroyCity(city *domain2.City) {
	if len(city.Aliens) < 2 {
		return
	}
	s.log("%s has been destroyed by alien %d and alien %d\n", city.Name, city.Aliens[0]+1, city.Aliens[1]+1)
	for _, alien := range city.Aliens {
		delete(s.aliens, alien)
		s.movesLeft[alien] = 0
	}
	s.worldMap.DestroyCity(city)
}
