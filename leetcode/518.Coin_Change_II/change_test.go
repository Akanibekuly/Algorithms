package coinchangeii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChange(t *testing.T) {
	cases := []struct {
		name   string
		coins  []int
		amount int
		result int
	}{
		{"5", []int{1, 2, 5}, 5, 4},
		{"no result", []int{2}, 3, 0},
		{"one reult", []int{10}, 10, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.result, change(c.amount, c.coins))
		})
	}
}
