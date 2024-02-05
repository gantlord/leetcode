package lc66

func add(a, b, carryIn int) (c, carryOut int) {
	r := a + b + carryIn
	return r % 10, r / 10
}

func plusOne(digits []int) []int {
	carry := 0
	for b, i := 1, len(digits)-1; i >= 0; i, b = i-1, 0 {
		digits[i], carry = add(digits[i], b, carry)
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
