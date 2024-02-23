package lc19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	array := []*ListNode{}
	for i := 0; ptr != nil; i++ {
		array = append(array, ptr)
		ptr = ptr.Next
	}
	if len(array) == 1 {
		return nil
	} else {
		if n == len(array) {
			head = array[1]
		} else if n == 1 {
			array[len(array)-n-1].Next = nil
		} else {
			array[len(array)-n-1].Next = array[len(array)-n+1]
		}
	}
	return head
}
