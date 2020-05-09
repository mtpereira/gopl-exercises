package ch4

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_removeAdjacentDuplicates(t *testing.T) {
	cases := map[string]struct {
		in   []string
		want []string
	}{
		"No duplicates": {
			in:   []string{"a", "b"},
			want: []string{"a", "b"},
		},
		"One duplicate": {
			in:   []string{"a", "a"},
			want: []string{"a"},
		},
		"One duplicate in a longer string": {
			in:   []string{"a", "a", "b", "c", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		"Two duplicates in a longer string": {
			in:   []string{"a", "a", "b", "c", "d", "d", "c"},
			want: []string{"a", "b", "c", "d", "c"},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := removeAdjacentDuplicates(tt.in)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				fmt.Println(got, len(got), cap(got))
				t.Errorf(diff)
			}
		})
	}
}
