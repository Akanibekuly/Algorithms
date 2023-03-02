package besttimetobuyandsellstock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfit(t *testing.T) {
	cases := []struct {
		prices []int
		res    int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, maxProfit(c.prices))
		})
	}
}
