package permutations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrmutations(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		result [][]int
	}{
		{"two elements [0,1]", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.result, permute(c.nums))
		})
	}
}
