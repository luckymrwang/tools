package utils

import (
	"sort"
)

type Map struct {
	Key   string
	Value int
}

type MapSlice []Map

func (m MapSlice) Len() int           { return len(m) }
func (m MapSlice) Less(i, j int) bool { return m[i].Value < m[j].Value }
func (m MapSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func Sort(m map[string]int) MapSlice {
	ms := make(MapSlice, 0)
	for k, v := range m {
		ms = append(ms, Map{k, v})
	}
	sort.Sort(ms)

	return ms
}

func SortReverse(m map[string]int) MapSlice {
	ms := make(MapSlice, 0)
	for k, v := range m {
		ms = append(ms, Map{k, v})
	}
	sort.Sort(sort.Reverse(ms))

	return ms
}
