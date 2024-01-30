package lc202

func isHappy(n int) bool {
	seenNumbers := make(map[int]bool)
	lastNumber := n
	for {
		if seenNumbers[lastNumber] {
			return false
		} else {
			seenNumbers[lastNumber] = true
		}
		nextNumber := 0
		for {
			digit := lastNumber % 10
			nextNumber += digit * digit
			lastNumber /= 10
			if lastNumber == 0 {
				break
			}
		}
		if nextNumber == 1 {
			return true
		}
		lastNumber = nextNumber
	}
}
