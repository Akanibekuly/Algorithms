package maximumsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		nums []int
		res  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, maxSubArray(c.nums))
		})
	}
}
