package countsubarrayswithfixedbounds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	cases := []struct {
		nums       []int
		minK, maxK int
		count      int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.count, countSubarrays(c.nums, c.minK, c.maxK))
		})
	}
}
