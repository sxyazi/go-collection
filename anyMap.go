package collect

type anyMap[T map[K]V, K, V comparable] struct {
	Items T
}

func AnyMap[T map[K]V, K, V comparable](items T) *anyMap[T, K, V] {
	return &anyMap[T, K, V]{Items: items}
}

func (c *anyMap[T, K, V]) Except(keys ...K) *anyMap[T, K, V] {
	for _, key := range keys {
		delete(c.Items, key)
	}

	return c
}

func (c *anyMap[T, K, V]) Keys() (keys []K) {
	for key := range c.Items {
		keys = append(keys, key)
	}
	return
}

func (c *anyMap[T, K, V]) diffKeys(target T) *anyMap[T, K, V] {
	items := make(T, len(target))
	for key := range c.Items {
		if _, ok := target[key]; !ok {
			items[key] = c.Items[key]
		}
	}

	return c
}
