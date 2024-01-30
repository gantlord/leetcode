package lc242

import "sort"

func sortString(s string) string {
	r := []rune(s)
	f := func(i, j int) bool {
		return r[i] < r[j]
	}
	sort.Slice(r, f)
	return string(r)
}

func isAnagram(s string, t string) bool {
	return sortString(s) == sortString(t)
}
