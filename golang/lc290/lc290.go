package lc290

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	patternToWord := make(map[rune]string)
	wordMapped := make(map[string]bool)
	for i, r := range pattern {
		word := words[i]
		if storedWord, exists := patternToWord[r]; !exists {
			patternToWord[r] = word
			if wordMapped[word] {
				return false
			} else {
				wordMapped[word] = true
			}
		} else {
			if storedWord != word {
				return false
			}
		}
	}
	return true
}
