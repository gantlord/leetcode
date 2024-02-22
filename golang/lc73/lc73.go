package lc73

func setZeroes(matrix [][]int) {
	zeroRows := make([]bool, len(matrix))
	zeroColumns := make([]bool, len(matrix[0]))
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				zeroRows[i] = true
				zeroColumns[j] = true
			}
		}
	}
	for i, row := range matrix {
		for j := range row {
			if zeroRows[i] || zeroColumns[j] {
				matrix[i][j] = 0
			}
		}
	}
}
