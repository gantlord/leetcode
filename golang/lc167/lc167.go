package lc167

func twoSum(numbers []int, target int) []int {
	j, k := 0, len(numbers)-1
	for {
		if numbers[j]+numbers[k] < target {
			j++
		} else if numbers[j]+numbers[k] > target {
			k--
		} else {
			return []int{j + 1, k + 1}
		}
	}
}
