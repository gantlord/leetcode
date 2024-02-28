package lc1768

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result += word1[i : i+1]
		}
		if i < len(word2) {
			result += word2[i : i+1]
		}
	}
	return result
}
