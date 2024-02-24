package lc146

// generate test cases in go test format

import (
	"testing"
)

/*func TestLRUCache(t *testing.T) {
	capacity := 2
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}*/

func TestLRUCache2(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	cache.Get(2)
	cache.Put(3, 2)
	cache.Get(2)
	cache.Get(3)
}
