package lc1313

func decompressRLElist(nums []int) []int {
	output := []int{}
	for i := 0; i < len(nums)/2; i++ {
		freq := nums[i*2]
		val := nums[i*2+1]
		for j := 0; j < freq; j++ {
			output = append(output, val)
		}
	}
	return output
}
