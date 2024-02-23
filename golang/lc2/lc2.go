package lc2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	carry := 0
	tail := result
	firstPass := true
	for l1 != nil || l2 != nil {

		if firstPass {
			firstPass = false
		} else {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		r := a + b + carry
		tail.Val = r % 10
		carry = r / 10
	}
	if carry > 0 {
		tail.Next = &ListNode{}
		tail.Next.Val = carry
	}
	return result
}
