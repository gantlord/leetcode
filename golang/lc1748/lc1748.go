package lc1748

func sumOfUnique(nums []int) int {
	counts := map[int]int{}
	sum := 0
	for _, n := range nums {
		if count, exists := counts[n]; exists {
			if count == 1 {
				sum -= n
			}
			counts[n] = count + 1
		} else {
			counts[n] = 1
			sum += n
		}
	}
	return sum
}
