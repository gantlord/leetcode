package lc2520

func countDigits(num int) int {
	count := 0
	working := num
	for working != 0 {
		if num%(working%10) == 0 {
			count++
		}
		working /= 10
	}
	return count
}
