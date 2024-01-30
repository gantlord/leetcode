package lc392

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	subsequenceIndex := 0
	for i := range t {
		if s[subsequenceIndex] == t[i] {
			subsequenceIndex++
			if subsequenceIndex == len(s) {
				return true
			}
		}
	}
	return false
}
