package lc2000

func reversePrefix(word string, ch byte) string {
	result := []byte(word)
	end := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			end = i
			break
		}
	}
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
