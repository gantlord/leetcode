package lc2485

func pivotInteger(n int) int {
	leftCount, rightCount := make([]int, n), make([]int, n)
	leftCount[0] = 1
	rightCount[n-1] = n
	for i := 1; i < n; i++ {
		leftCount[i] = leftCount[i-1] + i + 1
		rightCount[n-i-1] = rightCount[n-i] + n - i
	}
	for i := 0; i < n; i++ {
		if leftCount[i] == rightCount[i] {
			return i + 1
		}
	}
	return -1
}
