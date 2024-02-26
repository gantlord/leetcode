package lc1859

import (
	"strings"
)

func sortSentence(s string) string {
	array := make([]string, 9)
	start := 0
	count := 0
	for i, c := range s {
		if c >= '1' && c <= '9' {
			array[c-'1'] = strings.TrimSpace(string(s[start:i]))
			start = i + 1
			count++
		}
	}
	result := ""
	array = array[:count]
	for i, w := range array {
		result += w
		if i != len(array)-1 {
			result += " "
		}
	}
	return result
}
