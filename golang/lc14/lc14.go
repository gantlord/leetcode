package lc14

func longestCommonPrefix(strs []string) string {

	longestCommonPrefix := ""
	finished := false
	s1 := strs[0]

	for i := 0; i < len(s1); i++ {
		possiblePrefix := s1[:i+1]
		for _, s2 := range strs[1:] {
			if (len(s2) < i+1) || (s2[:i+1] != possiblePrefix) {
				finished = true
				break
			}
		}
		if finished {
			break
		} else {
			longestCommonPrefix = possiblePrefix
		}
	}
	return longestCommonPrefix
}
