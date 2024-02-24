package lc86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	headLT, headGTE := &ListNode{}, &ListNode{}
	ptrLT, ptrGTE := headLT, headGTE
	ptr := head
	for ptr != nil {
		next := ptr.Next
		if ptr.Val < x {
			ptrLT.Next = ptr
			ptrLT = ptr
		} else {
			ptrGTE.Next = ptr
			ptrGTE = ptr
			ptrGTE.Next = nil
		}
		ptr = next
	}
	ptrLT.Next = headGTE.Next
	return headLT.Next
}
