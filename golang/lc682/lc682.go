package lc682

import "strconv"

func calPoints(operations []string) int {
	record := []int{}
	for _, o := range operations {
		l := len(record)
		switch o[0] {
		case '+':
			record = append(record, record[l-1]+record[l-2])
		case 'C':
			record = record[:l-1]
		case 'D':
			record = append(record, 2*record[l-1])
		default:
			v, _ := strconv.Atoi(o)
			record = append(record, v)
		}
	}
	sum := 0
	for _, n := range record {
		sum += n
	}
	return sum
}
