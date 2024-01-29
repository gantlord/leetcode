package lc45

func jump(nums []int) int {
	finalIndex := len(nums) - 1
	minJumps := make([]int, len(nums))
	minJumps[finalIndex] = 0
	for i := finalIndex - 1; i >= 0; i-- {
		elem := nums[i]
		lowestJumps := 10_000
		for j := i + 1; j <= i+elem && j < len(nums); j++ {
			lowestJumps = min(lowestJumps, minJumps[j])
		}
		minJumps[i] = lowestJumps + 1
	}
	return minJumps[0]
}
