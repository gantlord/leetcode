package lc338

func numBits(n int) int {
	count := 0
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = numBits(i)
	}
	return result
}
