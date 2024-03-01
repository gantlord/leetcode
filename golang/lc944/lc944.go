package lc944

func minDeletionSize(strs []string) int {
	count := 0
	for _, s := range strs {
		for i := 2; i < len(s); i++ {
			valid := (s[i] >= s[i-1] && s[i-1] >= s[i-2]) || (s[i] <= s[i-1] && s[i-1] <= s[i-2])
			if !valid {
				count++
			}
		}
	}
	return count
}
