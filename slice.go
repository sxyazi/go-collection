package collect

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

type baseCollection[T []E, E any] struct {
	z any
}

type sliceCollection[T []E, E any] struct {
	z any
}

func Slice[T []E, E any](items T) *sliceCollection[T, E] {
	return &sliceCollection[T, E]{z: items}
}

func (c *sliceCollection[T, E]) Items() T {
	return c.z.(T)
}

func (c *sliceCollection[T, E]) Len() int {
	return len(c.z.(T))
}

func (c *sliceCollection[T, E]) Each(callback func(value E, index int)) *sliceCollection[T, E] {
	for index, value := range c.Items() {
		callback(value, index)
	}
	return c
}

func (c *sliceCollection[T, E]) Empty() bool {
	return c.Len() == 0
}

func (c *sliceCollection[T, E]) Same(target T) bool {
	if c.Len() != len(target) {
		return false
	} else if c.Empty() {
		return true
	}

	kind := reflect.TypeOf(c.z).Elem().Kind()
	if kind == reflect.Slice {
		return reflect.DeepEqual(c.z, target)
	}

	for index, item := range c.Items() {
		switch kind {
		case reflect.Float32:
			if math.Abs(float64(any(item).(float32))-float64(any(target[index]).(float32))) > 1e-9 {
				return false
			}
		case reflect.Float64:
			if math.Abs(any(item).(float64)-any(target[index]).(float64)) > 1e-9 {
				return false
			}
		default:
			if any(item) != any(target[index]) {
				return false
			}
		}
	}
	return true
}

func (c *sliceCollection[T, E]) First() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items()[0]
	return value, true
}

func (c *sliceCollection[T, E]) Last() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items()[c.Len()-1]
	return value, true
}

func (c *sliceCollection[T, E]) Index(value E) int {
	r1 := reflect.ValueOf(value)
	kind := reflect.TypeOf(c.z).Elem().Kind()

	for index, item := range c.Items() {
		if kind == reflect.Float64 {
			if math.Abs(any(item).(float64)-any(value).(float64)) <= 1e-9 {
				return index
			}
			continue
		} else if kind == reflect.Float32 {
			if math.Abs(float64(any(item).(float32))-float64(any(value).(float32))) <= 1e-9 {
				return index
			}
			continue
		} else if kind != reflect.Slice {
			if any(item) == any(value) {
				return index
			}
			continue
		}

		r2 := reflect.ValueOf(item)
		if r1.IsNil() != r2.IsNil() {
			continue
		} else if r1.Len() != r2.Len() {
			continue
		} else if r1.UnsafePointer() == r2.UnsafePointer() {
			return index
		}
	}

	return -1
}

func (c *sliceCollection[T, E]) Contains(value E) bool {
	return c.Index(value) != -1
}

func (c *sliceCollection[T, E]) Diff(target T) *sliceCollection[T, E] {
	t := Slice[T, E](target)
	return c.Filter(func(value E, index int) bool {
		return !t.Contains(value)
	})
}

func (c *sliceCollection[T, E]) Filter(callback func(value E, index int) bool) *sliceCollection[T, E] {
	var items T
	for index, item := range c.Items() {
		if callback(item, index) {
			items = append(items, item)
		}
	}

	c.z = items
	return c
}

func (c *sliceCollection[T, E]) Map(callback func(value E, index int) E) *sliceCollection[T, E] {
	for index, item := range c.Items() {
		c.z.(T)[index] = callback(item, index)
	}

	return c
}

func (c *sliceCollection[T, E]) Unique() *sliceCollection[T, E] {
	set := make(map[any]struct{})
	return c.Filter(func(value E, index int) bool {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
			return true
		}
		return false
	})
}

func (c *sliceCollection[T, E]) Merge(target T) *sliceCollection[T, E] {
	c.z = append(c.z.(T), target...)
	return c
}

func (c *sliceCollection[T, E]) Random() E {
	if c.Empty() {
		// TODO: returns false
		var zero E
		return zero
	}

	rand.Seed(time.Now().UnixNano())
	return c.Items()[rand.Intn(c.Len())]
}

func (c *sliceCollection[T, E]) Reverse() *sliceCollection[T, E] {
	z := c.Items()
	for i, j := 0, c.Len()-1; i < j; i, j = i+1, j-1 {
		z[i], z[j] = z[j], z[i]
	}
	return c
}

func (c *sliceCollection[T, E]) Shuffle() *sliceCollection[T, E] {
	z := c.Items()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(c.Len(), func(i, j int) { z[i], z[j] = z[j], z[i] })
	return c
}

func (c *sliceCollection[T, E]) Slice(offset, length int) *sliceCollection[T, E] {
	l := c.Len()
	if l == 0 || offset >= l || length == 0 {
		return c
	}

	// TODO: negative offset and length
	if offset+length > l {
		length = l - offset
	}

	c.z = c.Items()[offset : offset+length]
	return c
}

func (c *sliceCollection[T, E]) Split(number int) *baseCollection[[]T, T] {
	items := make([][]any, int(math.Ceil(float64(c.Len())/float64(number))))
	for i, item := range c.Items() {
		items[i/number] = append(items[i/number], item)
	}

	return &baseCollection[[]T, T]{z: items}
}

func (c *sliceCollection[T, E]) Splice(offset, length int) *sliceCollection[T, E] {
	l := c.Len()
	if l == 0 || offset >= l || length == 0 {
		return c
	}

	// TODO: negative offset and length
	if offset+length > l {
		length = l - offset
	}

	z := c.Items()
	items := z[offset : offset+length]
	c.z = append(z[:offset], z[offset+length:]...)
	return Slice[T, E](items)
}

func (c *sliceCollection[T, E]) Count() *mapCollection[map[any]int, any, int] {
	times := map[any]int{}
	for _, item := range c.Items() {
		if _, ok := times[item]; ok {
			times[item]++
		} else {
			times[item] = 0
		}
	}

	return Map[map[any]int, any, int](times)
}
