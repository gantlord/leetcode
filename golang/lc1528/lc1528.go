package lc1528

func restoreString(s string, indices []int) string {
	working := make([]byte, len(indices))
	for i, index := range indices {
		working[index] = s[i]
	}
	return string(working)
}
