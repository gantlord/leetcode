package lc530

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortVals(vals []int) {
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMinimumDifference(root *TreeNode) int {
	nodes := []*TreeNode{root}
	vals := []int{root.Val}
	for {
		l := len(nodes)
		for _, n := range nodes {
			if n.Left != nil {
				nodes = append(nodes, n.Left)
				vals = append(vals, n.Left.Val)
			}
			if n.Right != nil {
				nodes = append(nodes, n.Right)
				vals = append(vals, n.Right.Val)
			}
		}
		if len(nodes) == l {
			break
		}
		nodes = nodes[l:]
	}
	sortVals(vals)
	minDiff := abs(vals[1] - vals[0])
	prev := vals[1]
	for _, n := range vals[2:] {
		diff := abs(n - prev)
		if diff < minDiff {
			minDiff = diff
		}
		prev = n
	}
	return minDiff
}
