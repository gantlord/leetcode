package lc2859

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, n := range nums {
		count := 0
		for i != 0 {
			if i%2 == 1 {
				count++
			}
			i >>= 1
		}
		if count == k {
			sum += n
		}
	}
	return sum
}
