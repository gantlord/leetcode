package lc1588

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for oddLength := 1; oddLength <= len(arr); oddLength += 2 {
		for j := 0; j <= len(arr)-oddLength; j++ {
			for _, e := range arr[j : j+oddLength] {
				sum += e
			}
		}
	}
	return sum
}
