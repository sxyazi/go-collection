package collect

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func AnyGet[V, K any](item any, key K) (zero V, _ error) {
	var result any
	ref := reflect.ValueOf(item)

	switch ref.Kind() {
	case reflect.Map:
		if r := ref.MapIndex(reflect.ValueOf(key)); r.IsValid() {
			result = r.Interface()
		} else {
			return zero, errors.New("invalid map index")
		}
	case reflect.Array, reflect.Slice:
		if index, err := strconv.Atoi(fmt.Sprintf("%v", key)); err != nil {
			return zero, err
		} else {
			if index < 0 || index >= ref.Len() {
				return zero, errors.New("index overflow")
			}

			result = ref.Index(index).Interface()
		}
	case reflect.Struct:
		if r := ref.FieldByName(fmt.Sprintf("%v", key)); r.IsValid() {
			result = r.Interface()
		} else {
			return zero, errors.New("invalid struct field")
		}
	case reflect.Pointer:
		return AnyGet[V, K](ref.Elem().Interface(), key)
	default:
		return zero, errors.New("failed to get")
	}

	switch result.(type) {
	case V:
		return result.(V), nil
	default:
		return zero, errors.New("type mismatch")
	}
}

func Pluck[V, K, I comparable](items []I, key K) []V {
	var zero V
	plucked := make([]V, len(items), cap(items))

	for i, item := range items {
		if v, err := AnyGet[V](item, key); err == nil {
			plucked[i] = v
		} else {
			plucked[i] = zero
		}
	}

	return plucked
}

func MapPluck[K, V comparable](items []map[K]V, key K) []V {
	var zero V
	plucked := make([]V, len(items), cap(items))

	for i, item := range items {
		if v, ok := item[key]; ok {
			plucked[i] = v
		} else {
			plucked[i] = zero
		}
	}

	return plucked
}

func KeyBy[V, K, I comparable](items []I, key K) map[V]I {
	result := make(map[V]I)
	for _, item := range items {
		if v, err := AnyGet[V](item, key); err == nil {
			result[v] = item
		}
	}
	return result
}

func MapKeyBy[K, V comparable](items []map[K]V, key K) map[V]map[K]V {
	result := make(map[V]map[K]V)
	for _, item := range items {
		result[item[key]] = item
	}
	return result
}

func GroupBy[V, K, I comparable](items []I, key K) map[V][]I {
	result := make(map[V][]I)
	for _, item := range items {
		if v, err := AnyGet[V](item, key); err == nil {
			result[v] = append(result[v], item)
		}
	}
	return result
}

func MapGroupBy[K, V comparable](items []map[K]V, key K) map[V][]map[K]V {
	result := make(map[V][]map[K]V)
	for _, item := range items {
		result[item[key]] = append(result[item[key]], item)
	}
	return result
}
