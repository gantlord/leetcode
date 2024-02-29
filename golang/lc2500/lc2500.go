package lc2500

func deleteGreatestValue(grid [][]int) int {
	sum := 0
	for len(grid[0]) > 0 {
		gridLargest := 0
		for j, row := range grid {
			index := 0
			rowLargest := 0
			for k, e := range row {
				if e > rowLargest {
					index = k
					rowLargest = e
				}
			}
			grid[j] = append(row[:index], row[index+1:]...)
			gridLargest = max(gridLargest, rowLargest)
		}
		sum += gridLargest
	}
	return sum
}
