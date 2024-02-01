package lc101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	l, r := root.Left, root.Right
	if l == nil || r == nil {
		return l == r
	}
	l.Right, r.Right = r.Right, l.Right
	return l.Val == r.Val && isSymmetric(l) && isSymmetric(r)
}
