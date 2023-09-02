package minimumreplacementstosortthearray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinReplacement(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int64
	}{
		{"sorted", []int{1, 2, 3, 4, 5}, 0},
		{"two repl", []int{3, 9, 3}, 2},
		{"custom", []int{5, 2, 4, 7}, 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, minimumReplacement(c.nums))
		})
	}
}
