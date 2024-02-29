package lc933

type RecentCounter struct {
	recents map[int]bool
}

func Constructor() RecentCounter {
	return RecentCounter{map[int]bool{}}
}

func (c *RecentCounter) Ping(t int) int {
	c.recents[t] = true
	for k := range c.recents {
		if k < t-3000 {
			delete(c.recents, k)
		}
	}
	return len(c.recents)
}
