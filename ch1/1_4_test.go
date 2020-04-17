package ch4

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_dup4(t *testing.T) {
	cases := map[string]struct {
		in   []string
		want []string
	}{
		"file1": {
			in:   []string{"file1"},
			want: []string{"file1"},
		},
		"file2": {
			in:   []string{"file2"},
			want: nil,
		},
		"file{1,2}": {
			in:   []string{"file1", "file2"},
			want: []string{"file1"},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(dup4(tt.in), tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
