package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HasApple(t *testing.T) {
	cases := []struct {
		name     string
		n        int
		edges    [][]int
		hasApple []bool

		expectedRes int
	}{
		{
			name:        "1",
			n:           7,
			edges:       [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			hasApple:    []bool{false, false, true, false, true, true, false},
			expectedRes: 8,
		},
		{
			name:        "2",
			n:           7,
			edges:       [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			hasApple:    []bool{false, false, true, false, false, true, false},
			expectedRes: 6,
		},
		{
			name:        "3",
			n:           7,
			edges:       [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			hasApple:    []bool{false, false, false, false, false, false, false},
			expectedRes: 0,
		},
		{
			name:        "4",
			n:           8,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {2, 5}, {2, 6}, {4, 7}},
			hasApple:    []bool{true, true, false, true, false, true, true, false},
			expectedRes: 10,
		},
		{
			name:        "5",
			n:           4,
			edges:       [][]int{{0, 2}, {0, 3}, {1, 2}},
			hasApple:    []bool{false, true, false, false},
			expectedRes: 4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tr *testing.T) {
			res := minTime(c.n, c.edges, c.hasApple)
			assert.Equal(tr, c.expectedRes, res)

			res = minTimeDFS(c.n, c.edges, c.hasApple)
			assert.Equal(tr, c.expectedRes, res)
		})
	}
}
