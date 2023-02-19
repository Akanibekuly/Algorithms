package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	curr := []*TreeNode{root}

	var levels [][]*TreeNode
	levels = append(levels, curr)
	res = append(res, []int{root.Val})
	for len(curr) != 0 {
		var next []*TreeNode
		var vals []int
		for i := len(curr) - 1; i >= 0; i-- {
			if len(levels)%2 == 0 {
				if curr[i].Left != nil {
					next = append(next, curr[i].Left)
					vals = append(vals, curr[i].Left.Val)
				}
				if curr[i].Right != nil {
					next = append(next, curr[i].Right)
					vals = append(vals, curr[i].Right.Val)
				}
			} else {
				if curr[i].Right != nil {
					next = append(next, curr[i].Right)
					vals = append(vals, curr[i].Right.Val)
				}
				if curr[i].Left != nil {
					next = append(next, curr[i].Left)
					vals = append(vals, curr[i].Left.Val)
				}
			}
		}

		curr = next
		if len(next) != 0 {
			levels = append(levels, curr)
			res = append(res, vals)
		}
	}

	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count := 0
	var result [][]int

	for len(queue) > 0 {
		tmpArray := make([]int, len(queue))
		for i, v := range queue {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			if count%2 == 0 {
				tmpArray[i] = v.Val
			} else {
				tmpArray[len(tmpArray)-1-i] = v.Val
			}

			queue = queue[1:]
		}
		result = append(result, tmpArray)
		count++
	}
	return result
}
