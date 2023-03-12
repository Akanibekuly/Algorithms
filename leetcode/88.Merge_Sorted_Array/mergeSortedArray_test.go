package mergesortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArray(t *testing.T) {
	cases := []struct {
		nums1, nums2, res []int
		m, n              int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, []int{1, 2, 2, 3, 5, 6}, 3, 3},
		{[]int{1}, []int{}, []int{1}, 1, 0},
		{[]int{0}, []int{1}, []int{1}, 0, 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			merge(c.nums1, c.m, c.nums2, c.n)
			assert.Equal(t, c.res, c.nums1)
		})
	}
}
