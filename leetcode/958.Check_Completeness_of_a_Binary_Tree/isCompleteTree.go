package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	curr, levelLen, prevCount := []*TreeNode{root}, 1, 1
	for len(curr) > 0 {
		next, count := make([]*TreeNode, levelLen*2), 0
		for i, t := range curr {
			if t == nil {
				continue
			}

			if t.Left != nil {
				if count != i*2 {
					return false
				}
				next[i*2], count = t.Left, count+1
			}
			if t.Right != nil {
				if count != i*2+1 {
					return false
				}
				next[i*2+1], count = t.Right, count+1
			}

		}

		if count == 0 {
			break
		} else if prevCount != levelLen {
			return false
		}

		prevCount, levelLen, curr = count, levelLen*2, next
	}

	return true
}
