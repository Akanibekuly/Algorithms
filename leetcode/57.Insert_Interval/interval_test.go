package insertinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expectedRes [][]int
	}{
		{
			name:        "1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expectedRes: [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "2",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{4, 5},
			expectedRes: [][]int{{1, 3}, {4, 5}, {6, 9}},
		},
		{
			name:        "3",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{0, 5},
			expectedRes: [][]int{{0, 5}, {6, 9}},
		},
		{
			name:        "4",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{3, 5},
			expectedRes: [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "5",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{4, 6},
			expectedRes: [][]int{{1, 3}, {4, 9}},
		},
		{
			name:        "6",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{0, 9},
			expectedRes: [][]int{{0, 9}},
		},
		{
			name:        "7",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expectedRes: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(
				tt,
				c.expectedRes,
				insert(c.intervals, c.newInterval),
				"case %d", i,
			)
		})
	}
}
