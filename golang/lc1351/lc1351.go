package lc1351

func countNegatives(grid [][]int) int {
	count := 0
	for i, row := range grid {
		allNegative := false
		for j, e := range row {
			if j == 0 && e < 0 {
				allNegative = true
				break
			} else if e < 0 {
				count += len(row) - j
				break
			}
		}
		if allNegative {
			count += (len(grid) - i) * len(grid[0])
			break
		}
	}
	return count
}
