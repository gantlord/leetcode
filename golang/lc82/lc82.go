package lc82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var lastGood *ListNode = nil
	carry := head
	skipVal := 101
	for ptr := head.Next; ptr != nil; ptr = ptr.Next {
		if carry.Val == ptr.Val {
			skipVal = carry.Val
		} else if carry.Val != skipVal {
			if lastGood == nil {
				head = carry
			} else {
				lastGood.Next = carry
			}
			lastGood = carry
		}
		carry = ptr
	}
	if carry.Val != skipVal {
		if lastGood == nil {
			head = carry
		} else {
			lastGood.Next = carry
		}
	} else if lastGood != nil {
		lastGood.Next = nil
	} else {
		head = nil
	}

	return head
}
