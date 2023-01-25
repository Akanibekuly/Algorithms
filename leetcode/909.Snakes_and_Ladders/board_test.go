package snakesandladders

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeAndLadders(t *testing.T) {
	cases := []struct {
		board [][]int
		res   int
	}{
		// {
		// 	board: [][]int{{-1, -1}, {-1, 3}},
		// 	res:   1,
		// },
		// {
		// 	board: [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}},
		// 	res:   4,
		// },
		{
			board: [][]int{
				{-1, -1, -1, 46, 47, -1, -1, -1},
				{51, -1, -1, 63, -1, 31, 21, -1},
				{-1, -1, 26, -1, -1, 38, -1, -1},
				{-1, -1, 11, -1, 14, 23, 56, 57},
				{11, -1, -1, -1, 49, 36, -1, 48},
				{-1, -1, -1, 33, 56, -1, 57, 21},
				{-1, -1, -1, -1, -1, -1, 2, -1},
				{-1, -1, -1, 8, 3, -1, 6, 56},
			},
			res: 4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, snakesAndLadders(c.board))
		})
	}
}

// func TestGetCoordiantes(t *testing.T) {
// 	cases := []struct {
// 		n, curr, x, y int
// 	}{
// 		{3, 1, 2, 0},
// 		{3, 2, 2, 1},
// 		{3, 3, 2, 2},
// 		{3, 4, 1, 2},
// 		{3, 5, 1, 1},
// 		{3, 6, 1, 0},
// 		{3, 7, 0, 0},
// 		{3, 8, 0, 1},
// 		{3, 9, 0, 2},
// 		{6, 17, 3, 4},
// 		{6, 23, 2, 1},
// 	}

// 	for i, c := range cases {
// 		t.Run(fmt.Sprintf("case[%d] board %dx%d, curr [%d]", i, c.n, c.n, c.curr), func(tt *testing.T) {
// 			x, y := getCoordinates(c.n, c.curr)
// 			assert.Equal(tt, []int{x, y}, []int{c.x, c.y})
// 		})
// 	}
// }
