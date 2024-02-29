package lc728

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		working := i
		for {
			digit := working % 10
			if digit == 0 || i%digit != 0 {
				break
			}
			working /= 10
			if working == 0 {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
