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
		if ptr.Val < x {
			ptrLT.Next = ptr
			ptrLT = ptr
		} else {
			ptrGTE.Next = ptr
			ptrGTE = ptr
		}
		ptr = ptr.Next
	}
	ptrLT.Next = headGTE.Next
	ptrGTE.Next = nil
	return headLT.Next
}
