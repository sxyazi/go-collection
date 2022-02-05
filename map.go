package collect

type mapCollection[T map[K]V, K comparable, V any] struct {
	Items T
}

func Map[T map[K]V, K, V comparable](items T) *mapCollection[T, K, V] {
	return &mapCollection[T, K, V]{Items: items}
}

func (c *mapCollection[T, K, V]) Only(keys ...K) *mapCollection[T, K, V] {
	k := Slice(keys)
	for _, key := range keys {
		if !k.Contains(key) {
			delete(c.Items, key)
		}
	}

	return c
}

func (c *mapCollection[T, K, V]) Except(keys ...K) *mapCollection[T, K, V] {
	for _, key := range keys {
		delete(c.Items, key)
	}

	return c
}

func (c *mapCollection[T, K, V]) Keys() (keys []K) {
	for key := range c.Items {
		keys = append(keys, key)
	}
	return
}

func (c *mapCollection[T, K, V]) DiffKeys(target T) *mapCollection[T, K, V] {
	items := make(T, len(target))
	for key := range c.Items {
		if _, ok := target[key]; !ok {
			items[key] = c.Items[key]
		}
	}

	return c
}

func (c *mapCollection[T, K, V]) Has(key K) bool {
	if _, ok := c.Items[key]; ok {
		return true
	} else {
		return false
	}
}

func (c *mapCollection[T, K, V]) Set(key K, value V) *mapCollection[T, K, V] {
	c.Items[key] = value
	return c
}

func (c *mapCollection[T, K, V]) Get(key K) (value V, _ bool) {
	if !c.Has(key) {
		return
	}

	return c.Items[key], true
}

func (c *mapCollection[T, K, V]) Merge(target T) *mapCollection[T, K, V] {
	for key, value := range target {
		c.Items[key] = value
	}
	return c
}

// TODO
func (c *mapCollection[T, K, V]) Union() *mapCollection[T, K, V] {
	return c
}
