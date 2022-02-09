package collect

import (
	"fmt"
)

type SliceCollection[T ~[]E, E any] struct {
	z any
}

func UseSlice[T ~[]E, E any](items T) *SliceCollection[T, E] {
	return &SliceCollection[T, E]{items}
}

func (s *SliceCollection[T, E]) All() T {
	return s.z.(T)
}

func (s *SliceCollection[T, E]) New(items T) *SliceCollection[T, E] {
	return &SliceCollection[T, E]{items}
}

func (s *SliceCollection[T, E]) Len() int {
	return len(s.z.(T))
}

func (s *SliceCollection[T, E]) Empty() bool {
	return len(s.z.(T)) == 0
}

func (s *SliceCollection[T, E]) Print() *SliceCollection[T, E] {
	fmt.Println(s.z)
	return s
}

func (s *SliceCollection[T, E]) Each(callback func(value E, index int)) *SliceCollection[T, E] {
	Each[T, E](s.All(), callback)
	return s
}

func (s *SliceCollection[T, E]) Same(target T) bool {
	return Same[T, E](s.All(), target)
}

func (s *SliceCollection[T, E]) First() (E, bool) {
	return First[T, E](s.All())
}

func (s *SliceCollection[T, E]) Last() (E, bool) {
	return Last[T, E](s.All())
}

func (s *SliceCollection[T, E]) Index(value E) int {
	return Index(s.All(), value)
}

func (s *SliceCollection[T, E]) Contains(value E) bool {
	return Contains(s.All(), value)
}

func (s *SliceCollection[T, E]) Diff(target T) *SliceCollection[T, E] {
	s.z = Diff[T, E](s.All(), target)
	return s
}

func (s *SliceCollection[T, E]) Filter(callback func(value E, index int) bool) *SliceCollection[T, E] {
	s.z = Filter(s.All(), callback)
	return s
}

func (s *SliceCollection[T, E]) Map(callback func(value E, index int) E) *SliceCollection[T, E] {
	s.z = Map(s.All(), callback)
	return s
}

func (s *SliceCollection[T, E]) Unique() *SliceCollection[T, E] {
	s.z = Unique[T, E](s.All())
	return s
}

func (s *SliceCollection[T, E]) Merge(targets ...T) *SliceCollection[T, E] {
	s.z = Merge[T, E](s.All(), targets...)
	return s
}

func (s *SliceCollection[T, E]) Random() (E, bool) {
	return Random[T, E](s.All())
}

func (s *SliceCollection[T, E]) Reverse() *SliceCollection[T, E] {
	s.z = Reverse[T, E](s.All())
	return s
}

func (s *SliceCollection[T, E]) Shuffle() *SliceCollection[T, E] {
	s.z = Shuffle[T, E](s.All())
	return s
}

func (s *SliceCollection[T, E]) Slice(offset int, length ...int) *SliceCollection[T, E] {
	s.z = Slice[T, E](s.All(), offset, length...)
	return s
}

func (s *SliceCollection[T, E]) Split(amount int) []T {
	return Split[T, E](s.All(), amount)
}

func (s *SliceCollection[T, E]) Splice(offset int, args ...any) *SliceCollection[T, E] {
	var remaining T
	if len(args) >= 1 {
		remaining = Slice[T, E](s.All(), offset, args[0].(int))
	} else {
		remaining = Slice[T, E](s.All(), offset)
	}

	s.z = Splice[T, E](s.All(), offset, args...)
	return s.New(remaining)
}

func (s *SliceCollection[T, E]) Reduce(initial E, callback func(carry E, value E, key int) E) E {
	return Reduce[T, E](s.All(), initial, callback)
}
