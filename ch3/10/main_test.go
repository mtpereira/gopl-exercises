package main

import (
	"testing"
)

func Test_comma(t *testing.T) {
	cases := map[string]struct {
		in   string
		want string
	}{
		"4 digits": {
			in:   "1234",
			want: "1,234",
		},
		"many digits": {
			in:   "1234567891230",
			want: "1,234,567,891,230",
		},
		"2 digits": {
			in:   "12",
			want: "12",
		},
		"0 digits": {
			in:   "",
			want: "",
		},
		"1 digits": {
			in:   "1",
			want: "1",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {

			got := comma(tt.in)

			if got != tt.want {
				t.Errorf("comma(%v) = %v; want %v", tt.in, got, tt.want)
			}

		})
	}
}
