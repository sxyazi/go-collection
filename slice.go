package collect

type sliceCollection[T []E, E comparable] struct {
	Items T
}

func Slice[T []E, E comparable](items T) *sliceCollection[T, E] {
	return &sliceCollection[T, E]{Items: items}
}

func (c *sliceCollection[T, E]) Len() int {
	return len(c.Items)
}

func (c *sliceCollection[T, E]) Each(callback func(value E, key int)) *sliceCollection[T, E] {
	for key, value := range c.Items {
		callback(value, key)
	}
	return c
}

func (c *sliceCollection[T, E]) Empty() bool {
	return c.Len() == 0
}

func (c *sliceCollection[T, E]) Same(target T) bool {
	if c.Len() != Slice[T, E](target).Len() {
		return false
	}

	for key, value := range c.Items {
		if value != target[key] {
			return false
		}
	}

	return true
}

func (c *sliceCollection[T, E]) First() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items[0]
	return value, true
}

func (c *sliceCollection[T, E]) Last() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items[c.Len()-1]
	return value, true
}

func (c *sliceCollection[T, E]) Index(value E) int {
	for i, v := range c.Items {
		if v == value {
			return i
		}
	}
	return -1
}

func (c *sliceCollection[T, E]) Contains(value E) bool {
	return c.Index(value) != -1
}

func (c *sliceCollection[T, E]) Diff(target T) *sliceCollection[T, E] {
	t := Slice[T, E](target)
	return c.Filter(func(value E, key int) bool {
		return t.Contains(value)
	})
}

func (c *sliceCollection[T, E]) Filter(callback func(value E, key int) bool) *sliceCollection[T, E] {
	var items T
	for key, item := range c.Items {
		if callback(item, key) {
			items = append(items, item)
		}
	}

	c.Items = items
	return c
}

func (c *sliceCollection[T, E]) Map(callback func(value E, key int) E) *sliceCollection[T, E] {
	for key, item := range c.Items {
		c.Items[key] = callback(item, key)
	}

	return c
}

func (c *sliceCollection[T, E]) Unique() *sliceCollection[T, E] {
	set := make(map[E]struct{})
	var items T

	for _, item := range c.Items {
		if _, ok := set[item]; !ok {
			set[item] = struct{}{}
			items = append(items, item)
		}
	}

	c.Items = items
	return c
}
