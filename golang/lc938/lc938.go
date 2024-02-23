package lc938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	retVal := root.Val
	if retVal < low || retVal > high {
		retVal = 0
	}
	return retVal + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
