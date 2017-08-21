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
