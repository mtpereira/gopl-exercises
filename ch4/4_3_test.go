package ch4

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	cases := map[string]struct {
		in   *[4]int
		want *[4]int
	}{
		"Populated array": {
			in:   &[4]int{1, 2, 3, 4},
			want: &[4]int{4, 3, 2, 1},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			originalIn := *tt.in
			if reverse(tt.in); *tt.in != *tt.want {
				t.Errorf("reverse(%v) = %v; want %v", originalIn, tt.in, tt.want)
			}
		})
	}
}
