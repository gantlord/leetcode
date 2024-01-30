package lc13

func romanToInt(s string) int {
	index := 0
	count := 0
	for {
		switch s[index] {
		case 'M':
			count += 1000
		case 'D':
			count += 500
		case 'C':
			if index+1 < len(s) {
				switch s[index+1] {
				case 'D':
					count += 400
					index++
				case 'M':
					count += 900
					index++
				default:
					count += 100
				}
			} else {
				count += 100
			}
		case 'L':
			count += 50
		case 'X':
			if index+1 < len(s) {
				switch s[index+1] {
				case 'L':
					count += 40
					index++
				case 'C':
					count += 90
					index++
				default:
					count += 10
				}
			} else {
				count += 10
			}
		case 'V':
			count += 5
		case 'I':
			if index+1 < len(s) {
				switch s[index+1] {
				case 'V':
					count += 4
					index++
				case 'X':
					count += 9
					index++
				default:
					count += 1
				}
			} else {
				count += 1
			}
		}
		index++
		if index >= len(s) {
			break
		}
	}
	return count
}
