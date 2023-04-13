package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func buildMap1() Map {
	//C west=E south=J east=D north=B
	//J east=C south=D north=T
	//D south=T north=B
	//T east=J west=B north=E south=C
	//B north=E east=T
	//E west=D
	m := Map{}
	m.LinkCities("C", "E", West)
	m.LinkCities("C", "J", South)
	m.LinkCities("C", "D", East)
	m.LinkCities("C", "B", North)
	m.LinkCities("J", "C", East)
	m.LinkCities("J", "D", South)
	m.LinkCities("J", "T", North)
	m.LinkCities("D", "T", South)
	m.LinkCities("D", "B", North)
	m.LinkCities("T", "J", East)
	m.LinkCities("T", "B", West)
	m.LinkCities("T", "E", North)
	m.LinkCities("T", "C", South)
	m.LinkCities("B", "E", North)
	m.LinkCities("B", "T", East)
	m.LinkCities("E", "D", West)
	return m
}

func TestMap_DestroyCity(t *testing.T) {
	m := buildMap1()
	type args struct {
		city *City
	}
	type roadCount struct {
		name  string
		count int
	}
	tests := []struct {
		name                 string
		m                    Map
		args                 args
		expectLen            int
		expectedOutRoadCount []roadCount
	}{
		{
			name: "Empty map",
			m:    Map{},
			args: args{
				city: &City{
					Name: "A",
				},
			},
			expectLen: 0,
		},
		{
			name: "City not from the map",
			m:    m,
			args: args{
				city: &City{
					Name: "A",
				},
			},
			expectLen: 6,
		},
		{
			name: "City from the map",
			m:    m,
			args: args{
				city: m["D"],
			},
			expectLen: 5,
			expectedOutRoadCount: []roadCount{
				{
					name:  "C",
					count: 3,
				},
				{
					name:  "J",
					count: 2,
				},
				{
					name:  "E",
					count: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.DestroyCity(tt.args.city)
			if len(tt.m) != tt.expectLen {
				t.Errorf("Length mismath, want %v, got %v", tt.expectLen, len(tt.m))
			}
			for _, rc := range tt.expectedOutRoadCount {
				c := tt.m[rc.name]
				if len(c.OutRoad) != rc.count {
					t.Errorf("Out road length mismath for city %v, want %v, got %v", c.Name, rc.count, len(c.OutRoad))
				}
			}
		})
	}
}

func TestMap_InitCity(t *testing.T) {
	m1 := buildMap1()
	type args struct {
		name string
	}
	tests := []struct {
		name string
		m    Map
		args args
		want *City
	}{
		{
			name: "Existing city",
			m:    m1,
			args: args{
				name: "D",
			},
			want: m1["D"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.InitCity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_LinkCities(t *testing.T) {
	type args struct {
		from      string
		to        string
		direction Direction
	}
	tests := []struct {
		name    string
		m       Map
		args    args
		wantErr bool
	}{
		{
			name: "Link to itself",
			m:    buildMap1(),
			args: args{
				from:      "C",
				to:        "C",
				direction: North,
			},
			wantErr: true,
		},
		{
			name: "Link with already taken direction",
			m:    buildMap1(),
			args: args{
				from:      "C",
				to:        "D",
				direction: East,
			},
			wantErr: true,
		},
		{
			name: "New City",
			m:    buildMap1(),
			args: args{
				from:      "A",
				to:        "D",
				direction: East,
			},
			wantErr: false,
		},
		{
			name: "Two existing Cities",
			m:    buildMap1(),
			args: args{
				from:      "D",
				to:        "B",
				direction: East,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.LinkCities(tt.args.from, tt.args.to, tt.args.direction); (err != nil) != tt.wantErr {
				t.Errorf("LinkCities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMap_ListCities(t *testing.T) {
	m1 := buildMap1()
	tests := []struct {
		name string
		m    Map
		want []*City
	}{
		{
			name: "Empty map",
			m:    Map{},
			want: []*City{},
		},
		{
			name: "Non Empty map",
			m:    buildMap1(),
			want: []*City{
				m1["B"],
				m1["C"],
				m1["D"],
				m1["E"],
				m1["J"],
				m1["T"],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ListCities(); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("ListCities() = %v, want %v", got, tt.want)
			}
		})
	}
}
