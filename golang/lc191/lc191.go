package lc191

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += int((num >> i) & 1)
	}
	return count
}
