package collect

import "fmt"

type mapCollection[T ~map[K]V, K comparable, V any] struct {
	z T
}

func UseMap[T ~map[K]V, K, V comparable](items T) *mapCollection[T, K, V] {
	return &mapCollection[T, K, V]{items}
}

func (m *mapCollection[T, K, V]) All() T {
	return m.z
}

func (m *mapCollection[T, K, V]) New(items T) *mapCollection[T, K, V] {
	return &mapCollection[T, K, V]{items}
}

func (m *mapCollection[T, K, V]) Len() int {
	return len(m.z)
}

func (m *mapCollection[T, K, V]) Empty() bool {
	return len(m.z) == 0
}

func (m *mapCollection[T, K, V]) Print() *mapCollection[T, K, V] {
	fmt.Println(m.z)
	return m
}

func (m *mapCollection[T, K, V]) Only(keys ...K) *mapCollection[T, K, V] {
	m.z = Only[T, K, V](m.All(), keys...)
	return m
}

func (m *mapCollection[T, K, V]) Except(keys ...K) *mapCollection[T, K, V] {
	m.z = Except[T, K, V](m.All(), keys...)
	return m
}

func (m *mapCollection[T, K, V]) Keys() []K {
	return Keys[T, K, V](m.All())
}

func (m *mapCollection[T, K, V]) DiffKeys(target T) *mapCollection[T, K, V] {
	m.z = DiffKeys[T, K, V](m.All(), target)
	return m
}

func (m *mapCollection[T, K, V]) Has(key K) bool {
	return Has[T, K, V](m.All(), key)
}

func (m *mapCollection[T, K, V]) Set(key K, value V) *mapCollection[T, K, V] {
	m.z = Set(m.All(), key, value)
	return m
}

func (m *mapCollection[T, K, V]) Get(key K) (value V, _ bool) {
	return Get[T, K, V](m.All(), key)
}

func (m *mapCollection[T, K, V]) Same(target T) bool {
	return MapSame[T, K, V](m.All(), target)
}

func (m *mapCollection[T, K, V]) Merge(targets ...T) *mapCollection[T, K, V] {
	m.z = MapMerge[T, K, V](m.All(), targets...)
	return m
}

func (m *mapCollection[T, K, V]) Union(target T) *mapCollection[T, K, V] {
	return m.New(Union[T, K, V](m.All(), target))
}
