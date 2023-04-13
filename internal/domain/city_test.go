package domain

import (
	"testing"
)

func TestCity_Directions(t *testing.T) {
	type fields struct {
		Name      string
		OutRoad   map[Direction]*City
		inInroads RoadSet
		Aliens    []int
	}
	tests := []struct {
		name                string
		fields              fields
		wantDirectionsCount int
	}{
		{
			name: "Empty out roads",
			fields: fields{
				OutRoad: map[Direction]*City{},
			},
			wantDirectionsCount: 0,
		},
		{
			name: "Two out roads",
			fields: fields{
				OutRoad: map[Direction]*City{
					West:  &City{},
					North: &City{},
				},
			},
			wantDirectionsCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				Name:      tt.fields.Name,
				OutRoad:   tt.fields.OutRoad,
				inInroads: tt.fields.inInroads,
				Aliens:    tt.fields.Aliens,
			}
			if dirs := c.Directions(); len(dirs) != tt.wantDirectionsCount {
				t.Errorf("Directions() = %v, count: %v, want %v", dirs, len(dirs), tt.wantDirectionsCount)
			}
		})
	}
}

func TestCity_String(t *testing.T) {
	type fields struct {
		Name      string
		OutRoad   map[Direction]*City
		inInroads RoadSet
		Aliens    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Empty name",
			fields: fields{},
			want:   "",
		},
		{
			name: "Non empty name",
			fields: fields{
				Name: "Some",
			},
			want: "Some",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				Name:      tt.fields.Name,
				OutRoad:   tt.fields.OutRoad,
				inInroads: tt.fields.inInroads,
				Aliens:    tt.fields.Aliens,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
