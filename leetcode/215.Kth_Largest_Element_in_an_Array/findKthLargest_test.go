package kthlargestelementinanarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindKthLargest(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		k, res int
	}{
		{"1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, findKthLargest(c.nums, c.k))
		})
	}
}
