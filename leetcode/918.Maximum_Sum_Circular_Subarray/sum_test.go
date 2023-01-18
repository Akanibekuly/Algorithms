package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name        string
		nums        []int
		expectedRes int
	}{
		{
			name:        "1",
			nums:        []int{1, -2, 3, -2},
			expectedRes: 3,
		},
		{
			name:        "2",
			nums:        []int{5, -3, 5},
			expectedRes: 10,
		},
		{
			name:        "3",
			nums:        []int{-3, -2, -3},
			expectedRes: -2,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(tt, c.expectedRes, maxSubarraySumCircular(c.nums), "case %d", i)
		})
	}
}
