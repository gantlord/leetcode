package lc1816

import "strings"

func truncateSentence(s string, k int) string {
	result := ""
	a := strings.SplitAfter(s, " ")
	count := 0
	for _, w := range a {
		if w != "" {
			count++
			if count == k {
				result += strings.ReplaceAll(w, " ", "")
				break
			}
			result += w
		}
	}
	return result
}
