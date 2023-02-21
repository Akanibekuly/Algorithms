package duplicate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleDuplicate(t *testing.T) {
	cases := []struct {
		nums []int
		res  int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, singleNonDuplicate(c.nums))
		})
	}
}
