package lc138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	oldToNew := map[*Node]*Node{}
	ptr := head
	newHead := &Node{}
	newPtr := newHead

	for {
		newPtr.Val = ptr.Val
		oldToNew[ptr] = newPtr
		if ptr.Next == nil {
			break
		}
		ptr = ptr.Next
		newPtr.Next = &Node{}
		newPtr = newPtr.Next
	}
	ptr = head
	newPtr = newHead
	for ptr != nil {
		if ptr.Random != nil {
			newPtr.Random = oldToNew[ptr.Random]
		}
		ptr = ptr.Next
		newPtr = newPtr.Next
	}
	return newHead
}
