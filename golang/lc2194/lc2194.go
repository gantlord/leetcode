package lc2194

func cellsInRange(s string) []string {
	result := []string{}
	startColumn, startRow, endColumn, endRow := s[0], s[1], s[3], s[4]
	for i := startColumn; i <= endColumn; i++ {
		for j := startRow; j <= endRow; j++ {
			result = append(result, string([]byte{i, j}))
		}
	}
	return result
}
