package lc86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var headLT, headGTE *ListNode
	ptrLT, ptrGTE := headLT, headGTE
	ptr := head
	for ptr != nil {
		next := ptr.Next
		if ptr.Val < x {
			if headLT == nil {
				ptrLT = ptr
				headLT = ptr
			} else {
				ptrLT.Next = ptr
				ptrLT = ptr
			}
		} else {
			if headGTE == nil {
				ptrGTE = ptr
				headGTE = ptr
			} else {
				ptrGTE.Next = ptr
				ptrGTE = ptr
			}
			ptrGTE.Next = nil
		}
		ptr = next
	}
	if headLT == nil {
		return headGTE
	} else {
		ptrLT.Next = headGTE
	}
	return headLT
}
