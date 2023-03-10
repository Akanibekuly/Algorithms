package kthlargestsuminabinarytree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return 0
	}

	level := []*TreeNode{root}
	sums := []int64{int64(root.Val)}
	for len(level) > 0 {
		var next []*TreeNode
		var sum int64
		for _, r := range next {
			if r.Left != nil {
				sum += int64(r.Left.Val)
				next = append(next, r.Left)
			}
			if r.Right != nil {
				sum += int64(r.Right.Val)
				next = append(next, r.Right)
			}
		}
		if len(next) > 0 {
			sums = append(sums, sum)
			level = next
		} else {
			break
		}
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] < sums[j]
	})

	if k <= len(sums) {
		return sums[len(sums)-k]
	}

	return -1
}
