package lc2652

func sumOfMultiples(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			result += i
		}
	}
	return result
}
