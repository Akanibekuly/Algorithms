package uniquepathsii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquePath(t *testing.T) {
	cases := []struct {
		name   string
		grid   [][]int
		result int
	}{
		{"1x1", [][]int{{1}}, 0},
		{"3x3", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"2x2", [][]int{{0, 1}, {0, 0}}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.result, uniquePathsWithObstacles(c.grid))
		})
	}
}
