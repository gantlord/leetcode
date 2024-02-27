package lc2520

func countDigits(num int) int {
	seen := map[int]bool{}
	count := 0
	working := num
	for working != 0 {
		digit := working % 10
		if _, exists := seen[digit]; !exists && num%(working%10) == 0 {
			count++
		}
		seen[digit] = true
		working /= 10
	}
	return count
}
