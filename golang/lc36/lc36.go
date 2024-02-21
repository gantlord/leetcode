package lc36

func isValidSudoku(board [][]byte) bool {
	var columns, rows, boxes [9][9]bool
	for i, row := range board {
		for j, elem := range row {
			if elem == '.' {
				continue
			}
			boxIndex := j/3 + (i/3)*3
			v := elem - '1'
			if columns[i][v] || rows[j][v] || boxes[boxIndex][v] {
				return false
			}
			columns[i][v] = true
			rows[j][v] = true
			boxes[boxIndex][v] = true
		}
	}
	return true
}
