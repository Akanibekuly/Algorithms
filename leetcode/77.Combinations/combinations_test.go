package combinations

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		n, k int
		res  [][]int
	}{
		{1, 1, [][]int{{1}}},
		{3, 1, [][]int{{1}, {2}, {3}}},
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d -> %d", c.n, c.k), func(t *testing.T) {
			require.Equal(t, c.res, combine(c.n, c.k))
		})
	}
}

func TestCnK(t *testing.T) {
	cases := []struct {
		n, k, res int
	}{
		{4, 1, 4},
		{4, 2, 6},
		{4, 3, 4},
		{4, 4, 1},
		{5, 3, 10},
		{6, 3, 20},
		{6, 4, 15},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("C(%d, %d)", c.n, c.k), func(t *testing.T) {
			require.Equal(t, c.res, cnk(c.n, c.k))
		})
	}
}

func TestSwitchToTheNext(t *testing.T) {
	cases := []struct {
		name    string
		hasNext bool
	}{
		// get 2 from 4 Cnk = 6
		{"1100 -> 1010", true},  // 12 -> 13
		{"1010 -> 1001", true},  // 13 -> 14
		{"1001 -> 0110", true},  // 14 -> 23
		{"0110 -> 0101", true},  // 23 -> 24
		{"0101 -> 0011", true},  // 24 -> 34
		{"0011 -> 0011", false}, // 34 -> x
		// get 3 from 6 Cnk = 20
		{"111000 -> 110100", true}, // 123 -> 124
		{"110100 -> 110010", true}, // 124 -> 125
		{"110010 -> 110001", true}, // 125 -> 126
		{"110001 -> 101100", true}, // 126 -> 134
		{"101100 -> 101010", true}, // 134 -> 135
		{"101010 -> 101001", true}, // 135 -> 136
		{"101001 -> 100110", true}, // 136 -> 145
		{"100110 -> 100101", true}, // 145 -> 146
		{"100101 -> 100011", true}, // 146 -> 156
		{"100011 -> 011100", true}, // 156 -> 234
		{"011100 -> 011010", true}, // 234 -> 235
		{"011010 -> 011001", true}, // 235 -> 236
		{"011001 -> 010110", true}, // 236 -> 245
		{"010110 -> 010101", true}, // 245 -> 246
		{"010101 -> 010011", true}, // 246 -> 256
		{"010011 -> 001110", true}, // 256 -> 345
		{"001110 -> 001101", true}, // 345 -> 346
		{"001101 -> 001011", true}, // 346 -> 356
		{"001011 -> 000111", true}, // 356 -> 456
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arr := strings.Split(c.name, " -> ")
			require.Equal(t, len(arr), 2)

			in := getArrayFromStr(arr[0])
			out := getArrayFromStr(arr[1])

			require.Equal(t, c.hasNext, switchToTheNextVal(in))
			require.Equal(t, out, in)
		})
	}
}

func getArrayFromStr(s string) []int {
	arr := make([]int, len(s))
	for i, r := range s {
		if r == '1' {
			arr[i] = 1
		}
	}

	return arr
}
