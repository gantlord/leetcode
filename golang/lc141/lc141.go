package lc141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	visited := make(map[*ListNode]bool)
	for n := head.Next; n != nil; n = n.Next {
		if visited[n] {
			return true
		}
		visited[n] = true
	}
	return false
}
