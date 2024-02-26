package lc557

func reverse(s string) string {
	result := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func reverseWords(s string) string {
	result := ""
	start := 0
	for i, r := range s {
		if r == ' ' {
			result += reverse(s[start:i]) + " "
			start = i + 1
		}
		if i == len(s)-1 {
			result += reverse(s[start:])
		}
	}
	return result
}
