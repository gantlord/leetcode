package lc49

import "slices"

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, s := range strs {
		r := []rune(s)
		slices.Sort(r)
		sKey := string(r)
		if array, exists := anagrams[sKey]; exists {
			anagrams[sKey] = append(array, s)
		} else {
			anagrams[sKey] = []string{s}
		}
	}
	output := [][]string{}
	for _, v := range anagrams {
		output = append(output, v)
	}
	return output
}
