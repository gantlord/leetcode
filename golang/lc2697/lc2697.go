package lc2697

func makeSmallestPalindrome(s string) string {
	result := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		val := min(result[j], result[i])
		result[i], result[j] = val, val
	}
	return string(result)
}
