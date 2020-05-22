package tempconv

import (
	"testing"
)

func TestCToK(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Kelvin
	}{
		{
			name: "0K",
			c:    Celsius(-273.15),
			want: Kelvin(0),
		},
		{
			name: "273.15K",
			c:    Celsius(0),
			want: Kelvin(273.15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.c); got != tt.want {
				t.Errorf("CToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want Celsius
	}{
		{
			name: "0 C",
			k:    Kelvin(273.15),
			want: Celsius(0),
		},
		{
			name: "-273.15 C",
			k:    Kelvin(0),
			want: Celsius(-273.15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToC(tt.k); got != tt.want {
				t.Errorf("KToC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want string
	}{
		{
			name: "zero",
			k:    Kelvin(0),
			want: "0K",
		},
		{
			name: "some other value",
			k:    Kelvin(25),
			want: "25K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
