package kthmissingpositivenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthPositive(t *testing.T) {
	cases := []struct {
		arr    []int
		k, res int
	}{
		{[]int{2, 3, 4, 7, 11}, 2, 5},
		{[]int{2, 3, 4, 7, 11}, 1, 1},
		{[]int{2, 3, 4, 7, 11}, 4, 8},
		{[]int{2, 3, 4, 7, 11}, 6, 10},
		{[]int{1, 2, 3, 4}, 2, 6},
		{[]int{5}, 5, 6},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, findKthPositive(c.arr, c.k))
		})
	}
}
