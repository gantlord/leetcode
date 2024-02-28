package lc2418

import (
	"sort"
)

type Person struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	people := make([]Person, len(names))
	for i := 0; i < len(names); i++ {
		people[i] = Person{names[i], heights[i]}
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].height < people[j].height
	})
	for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
		people[i], people[j] = people[j], people[i]
	}
	for i := 0; i < len(names); i++ {
		names[i] = people[i].name
	}
	return names
}
