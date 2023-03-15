package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode
		res  bool
	}{
		{
			name: "nil tree",
			root: nil,
			res:  true,
		},
		{
			name: "complete tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			res: true,
		},
		{
			name: "wrong case",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 7},
				},
			},
			res: false,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isCompleteTree(c.root), fmt.Sprintf("case %d", i))
		})
	}
}
