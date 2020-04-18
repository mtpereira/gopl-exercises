package ch4

import (
	"testing"
)

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Test_rotate(t *testing.T) {
	cases := map[string]struct {
		s    []int
		n    int
		want []int
	}{
		"Empty slice": {
			s:    nil,
			n:    1,
			want: nil,
		},
		"Rotate once": {
			s:    []int{1, 2, 3},
			n:    1,
			want: []int{3, 1, 2},
		},
		"Rotate twice": {
			s:    []int{1, 2, 3},
			n:    2,
			want: []int{2, 3, 1},
		},
		"Rotate thrice": {
			s:    []int{1, 2, 3},
			n:    3,
			want: []int{1, 2, 3},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			if got := rotate(tt.s, tt.n); !equal(got, tt.want) {
				t.Errorf("rotate(%v) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
