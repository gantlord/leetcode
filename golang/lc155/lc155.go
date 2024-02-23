package lc155

type StackElem struct {
	val, min int
}

type MinStack struct {
	stackArray []StackElem
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	minSeen := val
	if len(m.stackArray) > 0 {
		minSeen = min(m.stackArray[len(m.stackArray)-1].min, val)
	}
	m.stackArray = append(m.stackArray, StackElem{val, minSeen})
}

func (m *MinStack) Pop() {
	m.stackArray = m.stackArray[:len(m.stackArray)-1]
}

func (m *MinStack) Top() int {
	return m.stackArray[len(m.stackArray)-1].val
}

func (m *MinStack) GetMin() int {
	return m.stackArray[len(m.stackArray)-1].min
}
