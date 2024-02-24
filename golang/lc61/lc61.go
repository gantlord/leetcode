package lc61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	maxLength := 500
	if head == nil {
		return nil
	}
	nodeWindow := make([]*ListNode, min(maxLength, k+1))
	ptr := head
	l := 0
	for ptr != nil {
		l++
		for i := 0; i < min(maxLength-1, k); i++ {
			nodeWindow[i] = nodeWindow[i+1]
		}
		nodeWindow[min(maxLength-1, k)] = ptr
		ptr = ptr.Next
	}
	k = k % l
	if k == 0 {
		return head
	}
	nodeWindow = nodeWindow[len(nodeWindow)-(k+1):]
	nodeWindow[0].Next = nil
	nodeWindow[k].Next = head
	return nodeWindow[1]
}
