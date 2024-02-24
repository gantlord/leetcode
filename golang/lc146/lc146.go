package lc146

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cache     map[int]*Node
	usageHead *Node
	usageTail *Node
	capacity  int
}

func Constructor(capacity int) LRUCache {
	usageHead, usageTail := &Node{}, &Node{}
	usageHead.next = usageTail
	usageTail.prev = usageHead
	return LRUCache{map[int]*Node{}, usageHead, usageTail, capacity}
}

func (c *LRUCache) insert(node *Node) {
	c.cache[node.key] = node
	node.next = c.usageTail
	node.prev = c.usageTail.prev
	node.prev.next = node
	c.usageTail.prev = node
}

func (c *LRUCache) remove(node *Node) {
	delete(c.cache, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) Get(key int) int {
	if node, exists := c.cache[key]; exists {
		c.remove(node)
		c.insert(node)
		return node.value
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		c.remove(node)
		node.value = value
		c.insert(node)
	} else {
		if len(c.cache) == c.capacity {
			eject := c.usageHead.next
			c.remove(eject)
		}
		node := &Node{key, value, nil, nil}
		c.insert(node)
	}
}
