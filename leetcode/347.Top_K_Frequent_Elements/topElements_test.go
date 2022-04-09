package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		res  []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{3, 0, 1, 0}, 1, []int{3}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case [%d]", i+1), func(tr *testing.T) {
			require.Equal(tr, topKFrequent(c.nums, c.k), c.res)
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	}
}
