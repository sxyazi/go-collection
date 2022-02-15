package collect

import "fmt"

type MapCollection[T ~map[K]V, K comparable, V any] struct {
	z T
}

func UseMap[T ~map[K]V, K comparable, V any](items T) *MapCollection[T, K, V] {
	return &MapCollection[T, K, V]{items}
}

func (m *MapCollection[T, K, V]) All() T {
	return m.z
}

func (m *MapCollection[T, K, V]) New(items T) *MapCollection[T, K, V] {
	return &MapCollection[T, K, V]{items}
}

func (m *MapCollection[T, K, V]) Len() int {
	return len(m.z)
}

func (m *MapCollection[T, K, V]) Empty() bool {
	return len(m.z) == 0
}

func (m *MapCollection[T, K, V]) Print() *MapCollection[T, K, V] {
	fmt.Println(m.z)
	return m
}

func (m *MapCollection[T, K, V]) Only(keys ...K) *MapCollection[T, K, V] {
	m.z = Only[T, K, V](m.z, keys...)
	return m
}

func (m *MapCollection[T, K, V]) Except(keys ...K) *MapCollection[T, K, V] {
	m.z = Except[T, K, V](m.z, keys...)
	return m
}

func (m *MapCollection[T, K, V]) Keys() []K {
	return Keys[T, K, V](m.z)
}

func (m *MapCollection[T, K, V]) DiffKeys(target T) *MapCollection[T, K, V] {
	m.z = DiffKeys[T, K, V](m.z, target)
	return m
}

func (m *MapCollection[T, K, V]) Has(key K) bool {
	return Has[T, K, V](m.z, key)
}

func (m *MapCollection[T, K, V]) Get(key K) (value V, _ bool) {
	return Get[T, K, V](m.z, key)
}

func (m *MapCollection[T, K, V]) Put(key K, value V) *MapCollection[T, K, V] {
	Put[T, K, V](m.z, key, value)
	return m
}

func (m *MapCollection[T, K, V]) Pull(key K) (V, bool) {
	return Pull[T, K, V](m.z, key)
}

func (m *MapCollection[T, K, V]) Same(target T) bool {
	return MapSame[T, K, V](m.z, target)
}

func (m *MapCollection[T, K, V]) Merge(targets ...T) *MapCollection[T, K, V] {
	m.z = MapMerge[T, K, V](m.z, targets...)
	return m
}

func (m *MapCollection[T, K, V]) Union(target T) *MapCollection[T, K, V] {
	return m.New(Union[T, K, V](m.z, target))
}
