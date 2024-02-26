package lc1572

func diagonalSum(mat [][]int) int {
	sum := 0
	for i, row := range mat {
		sum += row[i]
		j := len(row) - 1 - i
		if j != i {
			sum += row[j]
		}
	}
	return sum
}
