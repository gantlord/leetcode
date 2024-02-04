package lc67

func add(a, b, carryIn int) (rune, int) {
	r := a + b + carryIn
	return rune(r%2) + '0', r / 2
}

func addBinary(a string, b string) string {
	result := []rune{}
	carry := 0
	r := '0'

	maxLen := max(len(a), len(b))
	for i := 0; i < maxLen; i++ {
		aBit, bBit := 0, 0
		if i < len(a) {
			aBit = int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			bBit = int(b[len(b)-i-1] - '0')
		}
		r, carry = add(aBit, bBit, carry)
		result = append([]rune{r}, result...)
	}
	if carry == 1 {
		result = append([]rune{'1'}, result...)
	}
	return string(result)
}
