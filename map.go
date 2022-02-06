package collect

import "fmt"

type mapCollection[T map[K]V, K comparable, V any] struct {
	z T
}

func UseMap[T map[K]V, K, V comparable](items T) *mapCollection[T, K, V] {
	return &mapCollection[T, K, V]{z: items}
}

func (c *mapCollection[T, K, V]) All() T {
	return c.z
}

func (c *mapCollection[T, K, V]) Len() int {
	return len(c.z)
}

func (c *mapCollection[T, K, V]) Empty() bool {
	return len(c.z) == 0
}

func (c *mapCollection[T, K, V]) Print() *mapCollection[T, K, V] {
	fmt.Println(c.z)
	return c
}

func (c *mapCollection[T, K, V]) Only(keys ...K) *mapCollection[T, K, V] {
	c.z = Only[T, K, V](c.All(), keys...)
	return c
}

func (c *mapCollection[T, K, V]) Except(keys ...K) *mapCollection[T, K, V] {
	c.z = Except[T, K, V](c.All(), keys...)
	return c
}

func (c *mapCollection[T, K, V]) Keys() []K {
	return Keys[T, K, V](c.All())
}

func (c *mapCollection[T, K, V]) DiffKeys(target T) *mapCollection[T, K, V] {
	c.z = DiffKeys[T, K, V](c.All(), target)
	return c
}

func (c *mapCollection[T, K, V]) Has(key K) bool {
	return Has[T, K, V](c.All(), key)
}

func (c *mapCollection[T, K, V]) Set(key K, value V) *mapCollection[T, K, V] {
	c.z = Set(c.All(), key, value)
	return c
}

func (c *mapCollection[T, K, V]) Get(key K) (value V, _ bool) {
	return Get[T, K, V](c.All(), key)
}

func (c *mapCollection[T, K, V]) Merge(targets ...T) *mapCollection[T, K, V] {
	c.z = MapMerge[T, K, V](c.All(), targets...)
	return c
}

// TODO
func (c *mapCollection[T, K, V]) Union() *mapCollection[T, K, V] {
	return c
}
