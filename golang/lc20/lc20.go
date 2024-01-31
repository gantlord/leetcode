package lc20

type stackType struct {
	array []rune
}

func NewStack() *stackType {
	s := stackType{[]rune{}}
	return &s
}

func (s *stackType) Pop() rune {
	lastIndex := len(s.array) - 1
	r := s.array[lastIndex]
	s.array = s.array[:lastIndex]
	return r
}

func (s *stackType) Push(r rune) {
	s.array = append(s.array, r)
}

func (s *stackType) Len() int {
	return len(s.array)
}

func isValid(s string) bool {
	expect := map[rune]rune{']': '[', '}': '{', ')': '('}
	st := NewStack()
	for _, paren := range s {
		switch paren {
		case '(', '{', '[':
			st.Push(paren)
		case ')', '}', ']':
			if st.Len() == 0 || st.Pop() != expect[paren] {
				return false
			}
		}
	}
	return st.Len() == 0
}
