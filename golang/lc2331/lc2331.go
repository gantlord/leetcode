package lc2331

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		if root.Val == 1 {
			return true
		} else {
			return false
		}
	}
	l, r := evaluateTree(root.Left), evaluateTree(root.Right)
	if root.Val == 2 {
		return l || r
	} else {
		return l && r
	}
}
