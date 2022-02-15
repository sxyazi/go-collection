package collect

import (
	"fmt"
)

type SliceCollection[T ~[]E, E any] struct {
	z T
}

func UseSlice[T ~[]E, E any](items T) *SliceCollection[T, E] {
	return &SliceCollection[T, E]{items}
}

func (s *SliceCollection[T, E]) All() T {
	return s.z
}

func (s *SliceCollection[T, E]) New(items T) *SliceCollection[T, E] {
	return &SliceCollection[T, E]{items}
}

func (s *SliceCollection[T, E]) Len() int {
	return len(s.z)
}

func (s *SliceCollection[T, E]) Empty() bool {
	return len(s.z) == 0
}

func (s *SliceCollection[T, E]) Print() *SliceCollection[T, E] {
	fmt.Println(s.z)
	return s
}

func (s *SliceCollection[T, E]) Each(callback func(value E, index int)) *SliceCollection[T, E] {
	Each[T, E](s.z, callback)
	return s
}

func (s *SliceCollection[T, E]) Same(target T) bool {
	return Same[T, E](s.z, target)
}

func (s *SliceCollection[T, E]) First() (E, bool) {
	return First[T, E](s.z)
}

func (s *SliceCollection[T, E]) Last() (E, bool) {
	return Last[T, E](s.z)
}

func (s *SliceCollection[T, E]) Index(value E) int {
	return Index(s.z, value)
}

func (s *SliceCollection[T, E]) Contains(value E) bool {
	return Contains(s.z, value)
}

func (s *SliceCollection[T, E]) Diff(target T) *SliceCollection[T, E] {
	s.z = Diff[T, E](s.z, target)
	return s
}

func (s *SliceCollection[T, E]) Filter(callback func(value E, index int) bool) *SliceCollection[T, E] {
	s.z = Filter(s.z, callback)
	return s
}

func (s *SliceCollection[T, E]) Map(callback func(value E, index int) E) *SliceCollection[T, E] {
	s.z = Map(s.z, callback)
	return s
}

func (s *SliceCollection[T, E]) Unique() *SliceCollection[T, E] {
	s.z = Unique[T, E](s.z)
	return s
}

func (s *SliceCollection[T, E]) Duplicates() *MapCollection[map[int]E, int, E] {
	return UseMap[map[int]E, int, E](Duplicates[T, E](s.z))
}

func (s *SliceCollection[T, E]) Merge(targets ...T) *SliceCollection[T, E] {
	s.z = Merge[T, E](s.z, targets...)
	return s
}

func (s *SliceCollection[T, E]) Random() (E, bool) {
	return Random[T, E](s.z)
}

func (s *SliceCollection[T, E]) Reverse() *SliceCollection[T, E] {
	s.z = Reverse[T, E](s.z)
	return s
}

func (s *SliceCollection[T, E]) Shuffle() *SliceCollection[T, E] {
	s.z = Shuffle[T, E](s.z)
	return s
}

func (s *SliceCollection[T, E]) Slice(offset int, length ...int) *SliceCollection[T, E] {
	s.z = Slice[T, E](s.z, offset, length...)
	return s
}

func (s *SliceCollection[T, E]) Split(amount int) []T {
	return Split[T, E](s.z, amount)
}

func (s *SliceCollection[T, E]) Splice(offset int, args ...any) *SliceCollection[T, E] {
	return s.New(Splice[T, E](&s.z, offset, args...))
}

func (s *SliceCollection[T, E]) Reduce(initial E, callback func(carry E, value E, key int) E) E {
	return Reduce[T, E](s.z, initial, callback)
}

func (s *SliceCollection[T, E]) Pop() (E, bool) {
	return Pop[T, E](&s.z)
}

func (s *SliceCollection[T, E]) Push(item E) *SliceCollection[T, E] {
	Push[T, E](&s.z, item)
	return s
}

func (s *SliceCollection[T, E]) Where(args ...any) *SliceCollection[T, E] {
	s.z = Where[T, E](s.z, args...)
	return s
}

func (s *SliceCollection[T, E]) WhereIn(args ...any) *SliceCollection[T, E] {
	s.z = WhereIn[T, E](s.z, args...)
	return s
}

func (s *SliceCollection[T, E]) WhereNotIn(args ...any) *SliceCollection[T, E] {
	s.z = WhereNotIn[T, E](s.z, args...)
	return s
}
