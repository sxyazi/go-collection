package collect

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
	"reflect"
	"time"
)

/**
 * Any slice
 */

func Len(v any) int {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return reflect.ValueOf(v).Len()
	default:
		return -1
	}
}

func Empty(v any) bool {
	return Len(v) == 0
}

func Each[T ~[]E, E any](items T, callback func(value E, index int)) {
	for index, value := range items {
		callback(value, index)
	}
}

func Same[T ~[]E, E any](items, target T) bool {
	if len(items) != len(target) {
		return false
	} else if len(items) == 0 {
		return true
	}

	kind := reflect.TypeOf(items).Elem().Kind()
	if kind == reflect.Slice {
		return reflect.DeepEqual(items, target)
	}

	for index, item := range items {
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

func First[T ~[]E, E any](items T) (E, bool) {
	var value E
	if len(items) == 0 {
		return value, false
	}

	value = items[0]
	return value, true
}

func Last[T ~[]E, E any](items T) (E, bool) {
	var value E
	if len(items) == 0 {
		return value, false
	}

	value = items[len(items)-1]
	return value, true
}

func Index[T ~[]E, E any](items T, target E) int {
	r1 := reflect.ValueOf(target)
	kind := reflect.TypeOf(items).Elem().Kind()

	for index, item := range items {
		if kind == reflect.Float64 {
			if math.Abs(any(target).(float64)-any(item).(float64)) <= 1e-9 {
				return index
			}
			continue
		} else if kind == reflect.Float32 {
			if math.Abs(float64(any(target).(float32))-float64(any(item).(float32))) <= 1e-9 {
				return index
			}
			continue
		} else if kind != reflect.Slice {
			if any(target) == any(item) {
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

func Contains[T ~[]E, E any](items T, item E) bool {
	return Index(items, item) != -1
}

func Diff[T ~[]E, E any](items, target T) T {
	var different T
	for _, item := range items {
		if Index(target, item) == -1 {
			different = append(different, item)
		}
	}

	return different
}

func Filter[T ~[]E, E any](items T, callback func(value E, index int) bool) T {
	var filtered T
	for index, item := range items {
		if callback(item, index) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func Map[T ~[]E, E any](items T, callback func(value E, index int) E) T {
	for index, item := range items {
		items[index] = callback(item, index)
	}

	return items
}

func Unique[T ~[]E, E any](items T) T {
	set := make(map[any]struct{})
	kind := reflect.TypeOf(items).Elem().Kind()

	return Filter(items, func(value E, index int) bool {
		switch kind {
		case reflect.Slice:
			p := reflect.ValueOf(value).UnsafePointer()
			if _, ok := set[p]; !ok {
				set[p] = struct{}{}
				return true
			}

		default:
			if _, ok := set[value]; !ok {
				set[value] = struct{}{}
				return true
			}
		}

		return false
	})
}

func Merge[T ~[]E, E any](items T, targets ...T) T {
	for _, target := range targets {
		items = append(items, target...)
	}
	return items
}

func Random[T ~[]E, E any](items T) (E, bool) {
	if len(items) == 0 {
		var zero E
		return zero, false
	}

	rand.Seed(time.Now().UnixNano())
	return items[rand.Intn(len(items))], true
}

func Reverse[T ~[]E, E any](items T) T {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	return items
}

func Shuffle[T ~[]E, E any](items T) T {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	return items
}

func Slice[T ~[]E, E any](items T, offset int, args ...int) T {
	start, end := OffsetToIndex(len(items), offset, args...)
	return items[start:end]
}

func Split[T ~[]E, E any](items T, amount int) []T {
	split := make([]T, int(math.Ceil(float64(len(items))/float64(amount))))
	for i, item := range items {
		split[i/amount] = append(split[i/amount], item)
	}

	return split
}

func Splice[T ~[]E, E any](items T, offset int, args ...any) T {
	length := len(items)
	if len(args) >= 1 {
		length = args[0].(int)
	}

	start, end := OffsetToIndex(len(items), offset, length)
	replica := make(T, start)
	copy(replica, items[:start])

	if len(args) < 2 {
		return append(replica, items[end:]...)
	}

	for _, rep := range args[1:] {
		switch v := rep.(type) {
		case E:
			replica = append(replica, v)
		case T:
			replica = append(replica, v...)
		default:
			panic("replacement type error")
		}
	}
	return append(replica, items[end:]...)
}

func Count[T ~[]E, E comparable](items T) map[E]int {
	times := make(map[E]int)
	for _, item := range items {
		times[item]++
	}

	return times
}

/**
 * Number slice
 */

func Sum[T ~[]E, E constraints.Integer | constraints.Float](items T) (total E) {
	for _, value := range items {
		total += value
	}
	return
}

func Avg[T ~[]E, E constraints.Integer | constraints.Float](items T) E {
	if len(items) == 0 {
		return 0
	}

	return Sum[T, E](items) / E(len(items))
}

func Min[T ~[]E, E constraints.Integer | constraints.Float](items T) E {
	if len(items) == 0 {
		return 0
	}

	min := items[0]
	for _, value := range items {
		if min > value {
			min = value
		}
	}

	return min
}

func Max[T ~[]E, E constraints.Integer | constraints.Float](items T) E {
	if len(items) == 0 {
		return 0
	}

	max := items[0]
	for _, value := range items {
		if max < value {
			max = value
		}
	}

	return max
}

/**
 * Map
 */

func Only[T ~map[K]V, K comparable, V any](items T, keys ...K) T {
	m := make(T)
	for _, key := range keys {
		m[key] = items[key]
	}

	return m
}

func Except[T ~map[K]V, K comparable, V any](items T, keys ...K) T {
	keysMap := map[K]struct{}{}
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}

	m := make(T)
	for key, value := range items {
		if _, ok := keysMap[key]; !ok {
			m[key] = value
		}
	}
	return m
}

func Keys[T ~map[K]V, K comparable, V any](items T) (keys []K) {
	for key := range items {
		keys = append(keys, key)
	}
	return
}

func DiffKeys[T ~map[K]V, K comparable, V any](items T, target T) T {
	m := make(T)
	for key := range items {
		if _, ok := target[key]; !ok {
			m[key] = items[key]
		}
	}

	return m
}

func Has[T ~map[K]V, K comparable, V any](items T, key K) bool {
	if _, ok := items[key]; ok {
		return true
	} else {
		return false
	}
}

func Set[T ~map[K]V, K comparable, V any](items T, key K, value V) T {
	items[key] = value
	return items
}

func Get[T ~map[K]V, K comparable, V any](items T, key K) (value V, _ bool) {
	if !Has[T, K, V](items, key) {
		return
	}

	return items[key], true
}

func MapSame[T ~map[K]V, K comparable, V any](items, target T) bool {
	if len(items) != len(target) {
		return false
	} else if len(items) == 0 {
		return true
	}

	kind := reflect.TypeOf(items).Elem().Kind()
	if kind == reflect.Slice {
		return reflect.DeepEqual(items, target)
	}

	for index, item := range items {
		tv, ok := target[index]
		if !ok {
			return false
		}

		switch kind {
		case reflect.Float32:
			if math.Abs(float64(any(item).(float32))-float64(any(tv).(float32))) > 1e-9 {
				return false
			}
		case reflect.Float64:
			if math.Abs(any(item).(float64)-any(tv).(float64)) > 1e-9 {
				return false
			}
		default:
			if any(item) != any(tv) {
				return false
			}
		}
	}
	return true
}

func MapMerge[T ~map[K]V, K comparable, V any](items T, targets ...T) T {
	for _, target := range targets {
		for key, value := range target {
			items[key] = value
		}
	}
	return items
}

func Union[T ~map[K]V, K comparable, V any](items T, target T) T {
	for key, value := range target {
		if _, ok := items[key]; !ok {
			items[key] = value
		}
	}
	return items
}
