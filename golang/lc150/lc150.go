package lc150

import "strconv"

type stackType struct {
	array []int
}

func NewStack() *stackType {
	s := stackType{[]int{}}
	return &s
}

func (s *stackType) Pop() int {
	lastIndex := len(s.array) - 1
	r := s.array[lastIndex]
	s.array = s.array[:lastIndex]
	return r
}

func (s *stackType) Push(r int) {
	s.array = append(s.array, r)
}

func evalRPN(tokens []string) int {
	s := NewStack()

	for _, token := range tokens {
		switch token {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "-":
			s.Push(-s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		case "/":
			b := s.Pop()
			a := s.Pop()
			s.Push(a / b)
		default:
			v, _ := strconv.Atoi(token)
			s.Push(v)
		}
	}

	return s.Pop()

}
