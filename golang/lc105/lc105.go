package lc105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	node := &TreeNode{}
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node.Val = preorder[0]
	for i, v := range inorder {
		if v == node.Val {
			node.Left = buildTree(preorder[1:1+i], inorder[:i])
			node.Right = buildTree(preorder[1+i:], inorder[i+1:])
		}
	}
	return node
}
