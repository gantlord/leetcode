package lc128

func longestConsecutive(nums []int) int {
	headsToLength := map[int]int{}
	tailsToHeads := map[int]int{}
	seen := map[int]bool{}

	for _, n := range nums {
		if seen[n] {
			continue
		} else {
			seen[n] = true
		}
		len, headExists := headsToLength[n+1]
		head, tailExists := tailsToHeads[n-1]
		if headExists && tailExists {
			delete(tailsToHeads, n-1)
			delete(headsToLength, n+1)
			headsToLength[head] += 1 + len
			tailsToHeads[head+headsToLength[head]-1] = head
		} else if headExists {
			delete(headsToLength, n+1)
			headsToLength[n] = len + 1
			tailsToHeads[n+len] = n
		} else if tailExists {
			delete(tailsToHeads, n-1)
			tailsToHeads[n] = head
			headsToLength[head] += 1
		} else {
			headsToLength[n] = 1
			tailsToHeads[n] = n
		}
	}

	maxLen := 0
	for _, v := range headsToLength {
		maxLen = max(maxLen, v)
	}

	return maxLen

}
