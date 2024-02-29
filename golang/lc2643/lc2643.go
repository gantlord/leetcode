package lc2643

func rowAndMaximumOnes(mat [][]int) []int {
	indexMax := 0
	maxOnes := 0
	for i, row := range mat {
		oneCount := 0
		for _, n := range row {
			if n == 1 {
				oneCount++
			}
			if oneCount > maxOnes {
				maxOnes = oneCount
				indexMax = i
			}
		}
	}
	return []int{indexMax, maxOnes}

}
