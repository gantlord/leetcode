package lc6

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	index := 0
	rows := make([][]rune, min(numRows, len(s)))
	increasing := true
	for _, c := range s {
		rows[index] = append(rows[index], c)
		if (increasing && index != numRows-1) || index == 0 {
			index++
			increasing = true
		} else if (!increasing && index != 0) || index == numRows-1 {
			index--
			increasing = false
		}
	}
	for i, result := 0, ""; i <= numRows; i, result = i+1, result+string(rows[i]) {
		if i == numRows {
			return result
		}
	}
	return ""
}
