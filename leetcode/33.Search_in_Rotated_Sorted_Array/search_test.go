package searchinrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSEarch(t *testing.T) {
	cases := []struct {
		name        string
		nums        []int
		target, res int
	}{
		{"ok", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, search(c.nums, c.target))
		})
	}
}
