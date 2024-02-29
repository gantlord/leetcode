package lc1295

func findNumbers(nums []int) int {
	numEvenDigits := 0
	for _, n := range nums {
		numDigits := 0
		for ; n != 0; n /= 10 {
			numDigits++
		}
		if numDigits%2 == 0 {
			numEvenDigits++
		}
	}
	return numEvenDigits
}
