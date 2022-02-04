package collect

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Get[V, K comparable](item any, key K) (V, error) {
	var zero V
	var result any
	refOfItem := reflect.ValueOf(item)

	switch refOfItem.Kind() {
	case reflect.Map:
		if r := refOfItem.MapIndex(reflect.ValueOf(key)); r.IsValid() {
			result = r.Interface()
		} else {
			return zero, errors.New("invalid map index")
		}
	case reflect.Array, reflect.Slice:
		if index, err := strconv.Atoi(fmt.Sprintf("%d", key)); err != nil {
			return zero, err
		} else {
			if index < 0 || index >= refOfItem.Len() {
				return zero, errors.New("index overflow")
			}

			result = refOfItem.Index(index).Interface()
		}
	case reflect.Struct:
		if r := refOfItem.FieldByName(fmt.Sprintf("%s", key)); r.IsValid() {
			result = r.Interface()
		} else {
			return zero, errors.New("invalid struct field")
		}
	case reflect.Interface:
		fmt.Println("reflect.Interface")
		// TODO
	case reflect.Pointer:
		return Get[V, K](refOfItem.Elem().Interface(), key)
	default:
		return zero, errors.New("failed to get")
	}

	switch result.(type) {
	case V:
		return result.(V), nil
	default:
		return zero, nil
	}
}

func Pluck[K, V comparable](items []map[K]V, key K) []V {
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

func PluckAny[V, K, I comparable](items []I, key K) []V {
	var zero V
	plucked := make([]V, len(items), cap(items))

	for i, item := range items {
		if v, err := Get[V](item, key); err == nil {
			plucked[i] = v
		} else {
			plucked[i] = zero
		}
	}

	return plucked
}
