package lc1374

func generateTheString(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		if i == n-1 && n%2 == 0 {
			result[i] = 'b'
		} else {
			result[i] = 'a'
		}
	}
	return string(result)
}
