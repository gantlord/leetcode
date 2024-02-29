package lc876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	byIndex := map[int]*ListNode{}
	i := 0
	for ptr := head; ptr != nil; i, ptr = i+1, ptr.Next {
		byIndex[i] = ptr
	}
	return byIndex[i/2]
}
