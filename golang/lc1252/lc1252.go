package lc1252

func oddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for _, p := range indices {
		for j := 0; j < n; j++ {
			matrix[p[0]][j] = 1 - matrix[p[0]][j]
		}
		for i := 0; i < m; i++ {
			matrix[i][p[1]] = 1 - matrix[i][p[1]]
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count += matrix[i][j]
		}
	}
	return count
}
