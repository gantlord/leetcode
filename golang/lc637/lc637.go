package lc637

func averageOfLevels(root *TreeNode) []float64 {
	levelOrdered := []*TreeNode{root}
	averages := []float64{}
	levelCount := 0
	for l, a := 1, 0.0; l != 0; l, a = len(levelOrdered), 0.0 {
		for _, n := range levelOrdered {
			a += float64(n.Val)
			if n.Left != nil {
				levelOrdered = append(levelOrdered, n.Left)
			}
			if n.Right != nil {
				levelOrdered = append(levelOrdered, n.Right)
			}
		}
		a /= float64(l)
		averages = append(averages, a)
		levelCount++
		levelOrdered = levelOrdered[l:]
	}
	return averages
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
