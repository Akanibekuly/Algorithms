package subarraysumsdivisiblebyk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysDivByK(t *testing.T) {
	cases := []struct {
		name        string
		nums        []int
		k           int
		expectedRes int
	}{
		{
			name:        "1",
			nums:        []int{4, 5, 0, -2, -3, 1},
			k:           5,
			expectedRes: 7,
		},
		{
			name:        "2",
			nums:        []int{5},
			k:           9,
			expectedRes: 0,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(
				tt,
				c.expectedRes,
				subarraysDivByK(c.nums, c.k),
				"case %d", i,
			)
		})
	}
}
