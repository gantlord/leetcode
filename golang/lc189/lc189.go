package lc189

func rotate(nums []int, k int) {
	l := len(nums)
	kWrap := k % l
	newNums := append(nums[l-kWrap:], nums[:l-kWrap]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = newNums[i]
	}
}
