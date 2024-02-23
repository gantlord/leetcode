package lc71

import "strings"

type stackType struct {
	array []string
}

func NewStack() *stackType {
	s := stackType{[]string{}}
	return &s
}

func (s *stackType) Pop() string {
	if len(s.array) == 0 {
		return ""
	}
	lastIndex := len(s.array) - 1
	r := s.array[lastIndex]
	s.array = s.array[:lastIndex]
	return r
}

func (s *stackType) Push(r string) {
	s.array = append(s.array, r)
}

func (s *stackType) Len() int {
	return len(s.array)
}

func (s *stackType) GetArray() []string {
	return s.array
}

func simplifyPath(path string) string {
	stack := NewStack()
	split := strings.Split(path, "/")
	for _, s := range split {
		if s == "" || s == "." {
			continue
		} else if s == ".." {
			stack.Pop()
		} else {
			stack.Push(s)
		}
	}
	result := ""
	for _, s := range stack.GetArray() {
		result += "/" + s
	}
	if result == "" {
		return "/"
	}
	return result
}
