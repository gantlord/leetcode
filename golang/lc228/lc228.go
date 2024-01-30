package lc228

import "fmt"

func prepareOutput(s, e int, output *[]string) {
	outString := fmt.Sprint(s)
	if e != s {
		outString += "->" + fmt.Sprint(e)
	}
	*output = append(*output, outString)
}

func summaryRanges(nums []int) []string {
	output := []string{}
	if len(nums) == 0 {
		return output
	}
	prev := nums[0]
	startIndex := prev
	endIndex := prev
	for _, val := range nums[1:] {
		if val == prev+1 {
			endIndex++
		} else {
			prepareOutput(startIndex, endIndex, &output)
			startIndex = val
			endIndex = val
		}
		prev = val
	}
	prepareOutput(startIndex, endIndex, &output)
	return output
}
