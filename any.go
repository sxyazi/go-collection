package collect

type anySlice[T []E, E comparable] struct {
	Items T
}

func Any[T []E, E comparable](items T) *anySlice[T, E] {
	return &anySlice[T, E]{Items: items}
}

func (c *anySlice[T, E]) Len() int {
	return len(c.Items)
}

func (c *anySlice[T, E]) Each(callback func(value E, key int)) *anySlice[T, E] {
	for key, value := range c.Items {
		callback(value, key)
	}
	return c
}

func (c *anySlice[T, E]) Empty() bool {
	return c.Len() == 0
}

func (c *anySlice[T, E]) Same(target T) bool {
	if c.Len() != Any[T, E](target).Len() {
		return false
	}

	for key, value := range c.Items {
		if value != target[key] {
			return false
		}
	}

	return true
}

func (c *anySlice[T, E]) First() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items[0]
	return value, true
}

func (c *anySlice[T, E]) Last() (E, bool) {
	var value E
	if c.Empty() {
		return value, false
	}

	value = c.Items[c.Len()-1]
	return value, true
}

func (c *anySlice[T, E]) Index(value E) int {
	for i, v := range c.Items {
		if v == value {
			return i
		}
	}
	return -1
}

func (c *anySlice[T, E]) Contains(value E) bool {
	return c.Index(value) != -1
}

func (c *anySlice[T, E]) Diff(target T) *anySlice[T, E] {
	t := Any[T, E](target)
	return c.Filter(func(value E, key int) bool {
		return t.Contains(value)
	})
}

func (c *anySlice[T, E]) Filter(callback func(value E, key int) bool) *anySlice[T, E] {
	var items T
	for key, item := range c.Items {
		if callback(item, key) {
			items = append(items, item)
		}
	}

	c.Items = items
	return c
}
