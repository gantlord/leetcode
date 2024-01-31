package lc104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	done := make(map[*TreeNode]bool)
	maxDepth := 1
	cDepth := 1
	c := root
	if root == nil {
		return 0
	}
	for {
		if c.Left != nil && !done[c.Left] {
			c = c.Left
			cDepth++
		} else if c.Right != nil && !done[c.Right] {
			c = c.Right
			cDepth++
		} else {
			if cDepth == 1 {
				return maxDepth
			}
			if c.Left == nil && c.Right == nil {
				maxDepth = max(cDepth, maxDepth)
			}
			done[c] = true
			c = root
			cDepth = 1

		}
	}
}
