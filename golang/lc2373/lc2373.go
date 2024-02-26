package lc2373

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := 1; i < len(grid)-1; i++ {
		result[i-1] = make([]int, len(grid)-2)
		for j := 1; j < len(grid)-1; j++ {
			maxVal := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					maxVal = max(maxVal, grid[i+k][j+l])
				}
			}
			result[i-1][j-1] = maxVal
		}
	}
	return result
}
