package lc206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	ptr := head.Next
	prev.Next = nil
	for ptr != nil {
		ptr.Next, ptr, prev = prev, ptr.Next, ptr
	}
	return prev
}
