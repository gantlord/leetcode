package lc2810

func reverse(s string) string {
	result := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func finalString(s string) string {
	start := 0
	result := ""
	for i, c := range s {
		if c == 'i' {
			result = reverse(result + s[start:i])
			start = i + 1
		}
		if i == len(s)-1 {
			result += s[start : i+1]
		}
	}
	return result
}
