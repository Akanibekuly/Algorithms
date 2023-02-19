package binarytreezigzaglevelordertraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigzag(t *testing.T) {

	cases := []struct {
		root *TreeNode
		res  [][]int
	}{
		{
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			res: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root: &TreeNode{Val: 1},
			res:  [][]int{{1}},
		},
		{},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, zigzagLevelOrder(c.root))
			assert.Equal(tt, c.res, zigzagLevelOrder2(c.root))
		})
	}
}
