package lc289

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	next := make([][]int, m)
	for i, row := range board {
		next[i] = make([]int, n)
		for j, v := range row {
			neighbours := 0
			if i > 0 {
				if j > 0 {
					neighbours += board[i-1][j-1]
				}
				neighbours += board[i-1][j]
				if j < n-1 {
					neighbours += board[i-1][j+1]
				}
			}
			if j > 0 {
				neighbours += board[i][j-1]
			}
			if j < n-1 {
				neighbours += board[i][j+1]
			}
			if i < m-1 {
				if j > 0 {
					neighbours += board[i+1][j-1]
				}
				neighbours += board[i+1][j]
				if j < n-1 {
					neighbours += board[i+1][j+1]
				}
			}
			if neighbours == 3 {
				next[i][j] = 1
			} else if neighbours == 2 {
				next[i][j] = v
			} else {
				next[i][j] = 0
			}
		}
	}
	for i := range board {
		board[i] = next[i]
	}
}
