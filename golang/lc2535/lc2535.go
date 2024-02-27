package lc2535

func differenceOfSum(nums []int) int {
	elemSum, digitSum := 0, 0
	for _, e := range nums {
		elemSum += e
		for e != 0 {
			digitSum += e % 10
			e /= 10
		}
	}
	return elemSum - digitSum
}
