package parallelcoursesiii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinTime(t *testing.T) {
	cases := []struct {
		name      string
		n, result int
		relations [][]int
		time      []int
	}{
		{
			name:      "light",
			n:         3,
			relations: [][]int{{1, 3}, {2, 3}},
			time:      []int{3, 2, 5},
			result:    8,
		},
		{
			name:      "medium",
			n:         5,
			relations: [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}},
			time:      []int{1, 2, 3, 4, 5},
			result:    12,
		},
		{
			name: "failed test hard",
			n:    9,
			relations: [][]int{{2, 7}, {2, 6}, {3, 6}, {4, 6}, {7, 6}, {2, 1}, {3, 1}, {4, 1}, {6, 1}, {7, 1}, {3, 8}, {5, 8},
				{7, 8}, {1, 9}, {2, 9}, {6, 9}, {7, 9}},
			time:   []int{9, 5, 9, 5, 8, 7, 7, 8, 4},
			result: 32,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			require.Equal(tt, c.result, minimumTime(c.n, c.relations, c.time))
		})
	}
}
