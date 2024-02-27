package lc106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	node := &TreeNode{}
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node.Val = postorder[len(postorder)-1]
	for i, v := range inorder {
		if v == node.Val {
			node.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			node.Left = buildTree(inorder[:i], postorder[:i])
		}
	}
	return node
}
