package lc108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	centreish := len(nums) / 2
	return &TreeNode{
		Val:   nums[centreish],
		Left:  sortedArrayToBST(nums[:centreish]),
		Right: sortedArrayToBST(nums[centreish+1:]),
	}
}
