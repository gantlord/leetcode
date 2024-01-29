package lc380

import "math/rand"

type RandomizedSet struct {
	set map[int]bool
}

func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.set = make(map[int]bool)
	return r
}

func (s *RandomizedSet) Insert(val int) bool {
	if s.set[val] {
		return false
	}
	s.set[val] = true
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if !s.set[val] {
		return false
	}
	delete(s.set, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	x := rand.Intn(len(s.set))
	index := 0
	for k := range s.set {
		if index == x {
			return k
		}
		index++
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
