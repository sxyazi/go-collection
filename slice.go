package collect

import (
	"fmt"
)

type sliceCollection[T []E, E any] struct {
	z any
}

func UseSlice[T []E, E any](items T) *sliceCollection[T, E] {
	return &sliceCollection[T, E]{z: items}
}

func (c *sliceCollection[T, E]) All() T {
	return c.z.(T)
}

func (c *sliceCollection[T, E]) Len() int {
	return len(c.z.(T))
}

func (c *sliceCollection[T, E]) Empty() bool {
	return len(c.z.(T)) == 0
}

func (c *sliceCollection[T, E]) Print() *sliceCollection[T, E] {
	fmt.Println(c.z)
	return c
}

func (c *sliceCollection[T, E]) Each(callback func(value E, index int)) *sliceCollection[T, E] {
	Each(c.All(), callback)
	return c
}

func (c *sliceCollection[T, E]) Same(target T) bool {
	return Same[T, E](c.All(), target)
}

func (c *sliceCollection[T, E]) First() (E, bool) {
	return First[T, E](c.All())
}

func (c *sliceCollection[T, E]) Last() (E, bool) {
	return Last[T, E](c.All())
}

func (c *sliceCollection[T, E]) Index(value E) int {
	return Index(c.All(), value)
}

func (c *sliceCollection[T, E]) Contains(value E) bool {
	return Contains(c.All(), value)
}

func (c *sliceCollection[T, E]) Diff(target T) *sliceCollection[T, E] {
	c.z = Diff[T, E](c.All(), target)
	return c
}

func (c *sliceCollection[T, E]) Filter(callback func(value E, index int) bool) *sliceCollection[T, E] {
	c.z = Filter(c.All(), callback)
	return c
}

func (c *sliceCollection[T, E]) Map(callback func(value E, index int) E) *sliceCollection[T, E] {
	c.z = Map(c.All(), callback)
	return c
}

func (c *sliceCollection[T, E]) Unique() *sliceCollection[T, E] {
	c.z = Unique[T, E](c.All())
	return c
}

func (c *sliceCollection[T, E]) Merge(targets ...T) *sliceCollection[T, E] {
	c.z = Merge[T, E](c.All(), targets...)
	return c
}

func (c *sliceCollection[T, E]) Random() E {
	return Random[T, E](c.All())
}

func (c *sliceCollection[T, E]) Reverse() *sliceCollection[T, E] {
	c.z = Reverse[T, E](c.All())
	return c
}

func (c *sliceCollection[T, E]) Shuffle() *sliceCollection[T, E] {
	c.z = Shuffle[T, E](c.All())
	return c
}

func (c *sliceCollection[T, E]) Slice(offset, length int) *sliceCollection[T, E] {
	c.z = Slice[T, E](c.All(), offset, length)
	return c
}

func (c *sliceCollection[T, E]) Split(number int) *bareCollection[[]T, T] {
	return &bareCollection[[]T, T]{z: Split[T, E](c.All(), number)}
}

func (c *sliceCollection[T, E]) Splice(offset, length int) *sliceCollection[T, E] {
	c.z = Splice[T, E](c.All(), offset, length)
	return c
}
