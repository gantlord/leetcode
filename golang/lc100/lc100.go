package lc100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkBranchSame(p, q *TreeNode) bool {
	var same bool
	if p == nil || q == nil {
		if p != q {
			same = false
		} else {
			same = true
		}
	} else {
		same = isSameTree(p, q)
	}
	return same
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && checkBranchSame(p.Left, q.Left) && checkBranchSame(p.Right, q.Right)
}
