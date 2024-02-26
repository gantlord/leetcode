package lc1656

type OrderedStream struct {
	storage []string
	index   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	result := []string{}
	s.storage[idKey-1] = value
	for i := s.index; i < len(s.storage); i++ {
		if s.storage[i] != "" {
			s.index++
			result = append(result, s.storage[i])
		} else {
			break
		}
	}
	return result
}
