package usecases

import (
	"github.com/zippunov/alien-invasion/internal/domain"
	"io"
	"reflect"
	"testing"
)

func TestInitScenario(t *testing.T) {
	type args struct {
		infra IInfra
	}
	tests := []struct {
		name    string
		args    args
		want    Scenario
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitScenario(tt.args.infra)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitScenario() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitScenario() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScenario_Run(t *testing.T) {
	type fields struct {
		out         io.Writer
		worldMap    domain.Map
		aliensCount int
		aliens      map[domain.Alien]*domain.City
		movesLeft   []int
		log         func(format string, a ...any)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scenario{
				out:         tt.fields.out,
				worldMap:    tt.fields.worldMap,
				aliensCount: tt.fields.aliensCount,
				aliens:      tt.fields.aliens,
				movesLeft:   tt.fields.movesLeft,
				log:         tt.fields.log,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScenario_aliensQueue(t *testing.T) {
	type fields struct {
		out         io.Writer
		worldMap    domain.Map
		aliensCount int
		aliens      map[domain.Alien]*domain.City
		movesLeft   []int
		log         func(format string, a ...any)
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scenario{
				out:         tt.fields.out,
				worldMap:    tt.fields.worldMap,
				aliensCount: tt.fields.aliensCount,
				aliens:      tt.fields.aliens,
				movesLeft:   tt.fields.movesLeft,
				log:         tt.fields.log,
			}
			if got := s.aliensQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aliensQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScenario_destroyCity(t *testing.T) {
	type fields struct {
		out         io.Writer
		worldMap    domain.Map
		aliensCount int
		aliens      map[domain.Alien]*domain.City
		movesLeft   []int
		log         func(format string, a ...any)
	}
	type args struct {
		city *domain.City
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scenario{
				out:         tt.fields.out,
				worldMap:    tt.fields.worldMap,
				aliensCount: tt.fields.aliensCount,
				aliens:      tt.fields.aliens,
				movesLeft:   tt.fields.movesLeft,
				log:         tt.fields.log,
			}
			s.destroyCity(tt.args.city)
		})
	}
}

func TestScenario_moveAlien(t *testing.T) {
	type fields struct {
		out         io.Writer
		worldMap    domain.Map
		aliensCount int
		aliens      map[domain.Alien]*domain.City
		movesLeft   []int
		log         func(format string, a ...any)
	}
	type args struct {
		alien domain.Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.City
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scenario{
				out:         tt.fields.out,
				worldMap:    tt.fields.worldMap,
				aliensCount: tt.fields.aliensCount,
				aliens:      tt.fields.aliens,
				movesLeft:   tt.fields.movesLeft,
				log:         tt.fields.log,
			}
			if got := s.moveAlien(tt.args.alien); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveAlien() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScenario_seedAliens(t *testing.T) {
	type fields struct {
		out         io.Writer
		worldMap    domain.Map
		aliensCount int
		aliens      map[domain.Alien]*domain.City
		movesLeft   []int
		log         func(format string, a ...any)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scenario{
				out:         tt.fields.out,
				worldMap:    tt.fields.worldMap,
				aliensCount: tt.fields.aliensCount,
				aliens:      tt.fields.aliens,
				movesLeft:   tt.fields.movesLeft,
				log:         tt.fields.log,
			}
			s.seedAliens()
		})
	}
}
