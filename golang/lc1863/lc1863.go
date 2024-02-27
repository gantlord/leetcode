package lc1863

func generateSubsetsBitwise(arr []int) [][]int {
	n := len(arr)
	subsetCount := 1 << n // 2^n subsets
	subsets := make([][]int, 0, subsetCount)

	for i := 0; i < subsetCount; i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, arr[j])
			}
		}
		subsets = append(subsets, subset)
	}

	return subsets
}

func subsetXORSum(nums []int) int {
	sum := 0
	for _, subset := range generateSubsetsBitwise(nums) {
		xored := 0
		for _, v := range subset {
			xored ^= v
		}
		sum += xored

	}
	return sum
}
