package lc2325

func decodeMessage(key string, message string) string {
	table := map[rune]int{}
	result := []rune(message)
	index := 0
	for _, r := range key {
		if _, exists := table[r]; !exists && r != ' ' {
			table[r] = index
			index++
		}
	}

	for i, r := range message {
		if r != ' ' {
			result[i] = rune(table[r] + 'a')
		}
	}
	return string(result)
}
