package lc1290

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ptr := head
	result := 0
	for ptr != nil {
		result = result<<1 + ptr.Val
		ptr = ptr.Next
	}
	return result
}
