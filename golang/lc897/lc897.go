package lc897

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	return dfs(root, nil)
}

func dfs(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := dfs(root.Left, root)
	root.Left = nil
	root.Right = dfs(root.Right, tail)
	return res
}
