package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recursive(root.Left, root.Right)
}

func recursive(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && recursive(a.Left, b.Right) && recursive(a.Right, b.Left)
}
