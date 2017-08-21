package repos

import "github.com/ykhrustalev/highloadcup/collections"

func removeKeyIn(m map[int]*collections.IntSet, key int, item int) {
	set, found := m[key]
	if found {
		set.Remove(item)
	}
}

func addKeyTo(m map[int]*collections.IntSet, key int, item int) {
	set, found := m[key]
	if !found {
		set = collections.NewIntSet()
		m[key] = set
	}
	set.Add(item)
}

func removeKeyIn2(m map[string]*collections.IntSet, key string, item int) {
	set, found := m[key]
	if found {
		set.Remove(item)
	}
}

func addKeyTo2(m map[string]*collections.IntSet, key string, item int) {
	set, found := m[key]
	if !found {
		set = collections.NewIntSet()
		m[key] = set
	}
	set.Add(item)
}

func nothing() {}

func swipeKeyIfRequired(
	m map[int]*collections.IntSet,
	id int,
	oldGroup int,
	newGroup *int,
) func() {
	if newGroup == nil || oldGroup == *newGroup {
		return nothing
	}

	newGroup_ := *newGroup

	return func() {
		removeKeyIn(m, oldGroup, id)
		addKeyTo(m, newGroup_, id)
	}
}

func swipeKeyIfRequired2(
	m map[string]*collections.IntSet,
	id int,
	oldGroup string,
	newGroup *string,
) func() {
	if newGroup == nil || oldGroup == *newGroup {
		return nothing
	}

	newGroup_ := *newGroup

	return func() {
		removeKeyIn2(m, oldGroup, id)
		addKeyTo2(m, newGroup_, id)
	}
}
