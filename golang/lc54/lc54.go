package lc54

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	direction := Right
	output := []int{}
	rLimit, dLimit, lLimit, uLimit := m-1, n-1, 0, 1
	for i, j, k := 0, 0, 0; k < len(matrix)*len(matrix[0]); k++ {
		output = append(output, matrix[i][j])
		switch direction {
		case Right:
			if j == rLimit {
				direction = Down
				rLimit--
				i++
			} else {
				j++
			}
		case Down:
			if i == dLimit {
				direction = Left
				dLimit--
				j--
			} else {
				i++
			}
		case Left:
			if j == lLimit {
				direction = Up
				lLimit++
				i--
			} else {
				j--
			}
		case Up:
			if i == uLimit {
				direction = Right
				uLimit++
				j++
			} else {
				i--
			}
		}
	}
	return output
}
