package encoding

import (
	"bytes"
	domain2 "github.com/zippunov/alien-invasion/internal/domain"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestMarshal(t *testing.T) {
	type args struct {
		m domain2.Map
	}
	m1 := domain2.Map{}
	m2 := domain2.Map{}
	_ = m2.LinkCities("aaa", "ddd", domain2.South)
	_ = m2.LinkCities("aaa", "eee", domain2.East)
	_ = m2.LinkCities("eee", "aaa", domain2.West)
	_ = m2.LinkCities("ddd", "aaa", domain2.West)
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "Empty map",
			args: args{
				m: m1,
			},
			wantW: "",
		},
		{
			name: "Non empty map",
			args: args{
				m: m2,
			},
			wantW: `aaa south=ddd east=eee
ddd west=aaa
eee west=aaa
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			MarshalTxt(w, tt.args.m)

			if gotW := w.String(); len(gotW) != len(tt.wantW) {
				t.Errorf("MarshalTxt() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		r io.Reader
		m domain2.Map
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mapLen  int
	}{
		{
			name: "Empty file",
			args: args{
				r: strings.NewReader(""),
				m: domain2.Map{},
			},
			mapLen: 0,
		},
		{
			name: "File with invalid record",
			args: args{
				r: strings.NewReader(`aaa west=bbb
bbb south=bbb`),
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "File with 4 cities",
			args: args{
				r: strings.NewReader(`aaa west=bbb north=ddd
bbb south=ccc west=aaa`),
				m: domain2.Map{},
			},
			mapLen: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalTxt(tt.args.r, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalTxt() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && tt.mapLen != len(tt.args.m) {
				t.Errorf("parseCityLine() invalid map length want %d, actual %d", tt.mapLen, len(tt.args.m))
			}
		})
	}
}

func Test_parseCityLine(t *testing.T) {
	type args struct {
		s string
		m domain2.Map
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mapLen  int
	}{
		{
			name: "Empty line",
			args: args{
				s: "",
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "Invalid direction name",
			args: args{
				s: "aaa top=bbb",
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "Missing direction name",
			args: args{
				s: "aaa =bbb",
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "Missing neighbor name",
			args: args{
				s: "aaa west=",
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "One neighbor",
			args: args{
				s: "aaa west=bbb",
				m: domain2.Map{},
			},
			mapLen: 2,
		},
		{
			name: "Two neighbors",
			args: args{
				s: "aaa west=bbb south=ccc",
				m: domain2.Map{},
			},
			mapLen: 3,
		},
		{
			name: "Three neighbors",
			args: args{
				s: "aaa west=bbb south=ccc north=ddd ",
				m: domain2.Map{},
			},
			mapLen: 4,
		},
		{
			name: "Neighbor with two directions",
			args: args{
				s: "aaa west=bbb south=bbb north=ddd ",
				m: domain2.Map{},
			},
			mapLen: 3,
		},
		{
			name: "Two neighbors with same direction",
			args: args{
				s: "aaa west=bbb north=ccc north=ddd ",
				m: domain2.Map{},
			},
			wantErr: true,
		},
		{
			name: "City link to itself",
			args: args{
				s: "aaa west=bbb south=aaa north=ddd ",
				m: domain2.Map{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseCityLine(tt.args.s, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("parseCityLine() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && tt.mapLen != len(tt.args.m) {
				t.Errorf("parseCityLine() invalid map length want %d, actual %d", tt.mapLen, len(tt.args.m))
			}
		})
	}
}

func Test_trimAndFilter(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "Empty slice",
			args: args{
				tokens: []string{},
			},
			want: []string{},
		},
		{
			name: "Slice with empty elements",
			args: args{
				tokens: []string{"", "", ""},
			},
			want: []string{},
		},
		{
			name: "Slice with non empty elements",
			args: args{
				tokens: []string{"1", "", "2"},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimAndFilter(tt.args.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimAndFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
