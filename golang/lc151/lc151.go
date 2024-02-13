package lc151

import "strings"

func reverseWords(s string) string {
	array := strings.Split(s, " ")
	result := ""
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == "" {
			continue
		}
		if len(result) > 0 {
			result += " "
		}
		result += array[i]
	}
	return result
}
