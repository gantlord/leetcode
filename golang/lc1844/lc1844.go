package lc1844

func replaceDigits(s string) string {
	result := []byte(s)
	for i := range s {
		if i%2 == 1 {
			result[i] = s[i-1] + s[i] - '0'
		}
	}
	return string(result)
}
