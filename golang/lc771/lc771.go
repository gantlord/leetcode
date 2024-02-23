package lc771

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, s := range stones {
		if strings.ContainsRune(jewels, s) {
			count++
		}
	}
	return count
}
