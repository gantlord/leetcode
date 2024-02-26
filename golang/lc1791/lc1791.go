package lc1791

func findCenter(edges [][]int) int {
	nodes := map[int]bool{}
	for _, edge := range edges {
		if _, exists := nodes[edge[0]]; exists {
			return edge[0]
		}
		if _, exists := nodes[edge[1]]; exists {
			return edge[1]
		}
		nodes[edge[0]] = true
		nodes[edge[1]] = true
	}
	return -1
}
