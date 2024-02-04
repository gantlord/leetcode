package lc190

func reverseBits(num uint32) uint32 {
	var output uint32
	for i := 0; i < 32; i++ {
		output += ((num >> i) & 1) << (31 - i)
	}
	return output
}
