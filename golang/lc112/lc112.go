package lc112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	targetSum -= root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
