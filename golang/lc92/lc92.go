package lc92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseFor(head *ListNode, length int) *ListNode {
	tail := head
	var carry *ListNode = nil
	for i := 0; i <= length; i++ {
		carry, head.Next, head = head, carry, head.Next
	}
	tail.Next = head
	return carry
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var carry *ListNode
	ptr := head
	for i := 1; ptr != nil; i++ {
		if i == left {
			ptr = reverseFor(ptr, right-left)
			if carry != nil {
				carry.Next = ptr
			}
			if i == 1 {
				head = ptr
			}
			break
		}
		carry = ptr
		ptr = ptr.Next
	}
	return head
}
