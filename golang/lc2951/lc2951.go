package lc2951

func findPeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		n := mountain[i]
		if n > mountain[i-1] && n > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}
