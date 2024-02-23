package lc2114

import "strings"

func mostWordsFound(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		count := 0
		for _, w := range strings.Split(s, " ") {
			if w != "" {
				count++
			}
		}
		maxWords = max(maxWords, count)
	}
	return maxWords
}
