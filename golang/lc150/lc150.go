package lc150

import "strconv"

type OpType int

const (
	NIL OpType = iota
	MUL
	DIV
	ADD
	SUB
	RES
)

type StackElem struct {
	a, b      *int64
	operation OpType
}

type stackType struct {
	array []StackElem
}

func NewStack() *stackType {
	s := stackType{[]StackElem{}}
	return &s
}

func (s *stackType) Pop() StackElem {
	lastIndex := len(s.array) - 1
	r := s.array[lastIndex]
	s.array = s.array[:lastIndex]
	return r
}

func (s *stackType) Push(r StackElem) {
	s.array = append(s.array, r)
}

func (s *stackType) Len() int {
	return len(s.array)
}

func (s *stackType) GetArray() []StackElem {
	return s.array
}

func evalRPN(tokens []string) int {
	elem := StackElem{}
	s := NewStack()

	for i := len(tokens) - 1; i >= 0; i-- {
		var operation OpType = NIL
		val := new(int64)
		switch tokens[i] {
		case "*":
			operation = MUL
		case "/":
			operation = DIV
		case "+":
			operation = ADD
		case "-":
			operation = SUB
		default:
			*val, _ = strconv.ParseInt(tokens[i], 10, 64)
		}

		elem = processElem(s, elem, operation, val)

	}
	return int(*elem.a)
}

func processElem(s *stackType, elem StackElem, operation OpType, val *int64) StackElem {
	if elem.operation != NIL && operation != NIL {
		s.Push(elem)
		elem = StackElem{}
		elem.operation = operation
	} else if operation != NIL {
		elem.operation = operation
	} else {
		if elem.a == nil {
			elem.a = val
		} else if elem.b == nil {
			elem.b = val
			var result int64 = 0
			switch elem.operation {
			case MUL:
				result = *elem.b * *elem.a
			case DIV:
				result = *elem.b / *elem.a
			case ADD:
				result = *elem.b + *elem.a
			case SUB:
				result = *elem.b - *elem.a
			}
			if s.Len() == 0 {
				return StackElem{&result, nil, RES}
			}
			next := s.Pop()
			elem = processElem(s, next, NIL, &result)
		}
	}
	return elem
}
