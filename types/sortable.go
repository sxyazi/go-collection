package types

import (
	"golang.org/x/exp/constraints"
)

/**
 * SortableSlice
 */

type SortableSlice[T ~[]E, E constraints.Ordered] struct {
	Items T
}

func (s SortableSlice[T, E]) Len() int {
	return len(s.Items)
}

func (s SortableSlice[T, E]) Less(i, j int) bool {
	return s.Items[i] < s.Items[j] || (s.Items[i] != s.Items[i] && s.Items[j] == s.Items[j])
}

func (s SortableSlice[T, E]) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

/**
 * SortableStruct
 */

type SortableStruct[E constraints.Ordered] struct {
	Value    E
	Attached any
}

type SortableStructs[T ~[]E, E constraints.Ordered] struct {
	Items []*SortableStruct[E]
}

func (s SortableStructs[T, E]) Len() int {
	return len(s.Items)
}

func (s SortableStructs[T, E]) Less(i, j int) bool {
	return s.Items[i].Value < s.Items[j].Value || (s.Items[i].Value != s.Items[i].Value && s.Items[j].Value == s.Items[j].Value)
}

func (s SortableStructs[T, E]) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}
