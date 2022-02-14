package collect

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
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
		if index, err := strconv.Atoi(fmt.Sprintf("%d", key)); err != nil {
			return zero, err
		} else {
			if index < 0 || index >= ref.Len() {
				return zero, errors.New("index overflow")
			}

			result = ref.Index(index).Interface()
		}
	case reflect.Struct:
		if r := ref.FieldByName(fmt.Sprintf("%s", key)); r.IsValid() {
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

func IsNumber(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func NumberCompare[T constraints.Integer | constraints.Float](a T, operator string, b T) bool {
	switch operator {
	case "=":
		return a == b
	case "!=":
		return a != b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	}
	return false
}

func AnyNumberCompare(a any, operator string, b any) bool {
	if a == nil || b == nil {
		return false
	} else if !IsNumber(a) {
		return false
	}

	in := operator == "in" || operator == "not in"
	positive := operator == "=" || operator == "in"
	ar, br := reflect.ValueOf(a), reflect.ValueOf(b)
	if in && ((br.Kind() != reflect.Slice && br.Kind() != reflect.Array) || br.Len() == 0 || !IsNumber(br.Index(0).Interface())) {
		return !positive
	}

	switch true {
	case ar.CanInt():
		if !in {
			if !br.CanInt() {
				return !positive
			}
			return NumberCompare(ar.Int(), operator, br.Int())
		}
		for i := 0; i < br.Len(); i++ {
			e := br.Index(i)
			if e.CanInt() && e.Int() == ar.Int() {
				return positive
			}
		}
		return !positive
	case ar.CanUint():
		if !in {
			if !br.CanUint() {
				return !positive
			}
			return NumberCompare(ar.Uint(), operator, br.Uint())
		}
		for i := 0; i < br.Len(); i++ {
			e := br.Index(i)
			if e.CanUint() && e.Uint() == ar.Uint() {
				return positive
			}
		}
		return !positive
	case ar.CanFloat():
		if !in {
			if !br.CanFloat() {
				return !positive
			}
			return NumberCompare(ar.Float(), operator, br.Float())
		}
		for i := 0; i < br.Len(); i++ {
			e := br.Index(i)
			if e.CanFloat() && e.Float() == ar.Float() {
				return positive
			}
		}
		return !positive
	}

	return false
}

func Compare(a any, operator string, b any) bool {
	in := operator == "in" || operator == "not in"
	positive := operator == "=" || operator == "in"

	if a == nil && b == nil {
		return !in && positive
	} else if a == nil || b == nil {
		return !positive
	}

	if !in && (IsNumber(a) || IsNumber(b)) {
		return AnyNumberCompare(a, operator, b)
	} else if !Contains([]string{"=", "!=", "in", "not in"}, operator) {
		return false
	}

	ar, br := reflect.TypeOf(a), reflect.TypeOf(b)
	if !in && ar.Kind() != br.Kind() {
		return !positive
	} else if in && ((br.Kind() != reflect.Slice && br.Kind() != reflect.Array) || br.Elem().Kind() != ar.Kind()) {
		return !positive
	}

	if !Contains([]reflect.Kind{reflect.Func, reflect.Map, reflect.Slice}, ar.Kind()) {
		switch operator {
		case "=":
			return a == b
		case "!=":
			return a != b
		case "in", "not in":
			for i, s := 0, reflect.ValueOf(b); i < s.Len(); i++ {
				if s.Index(i).Interface() == a {
					return positive
				}
			}
			return !positive
		}
	}

	p := reflect.ValueOf(a).UnsafePointer()
	switch operator {
	case "=":
		return p == reflect.ValueOf(b).UnsafePointer()
	case "!=":
		return p != reflect.ValueOf(b).UnsafePointer()
	case "in", "not in":
		for i, s := 0, reflect.ValueOf(b); i < s.Len(); i++ {
			if s.Index(i).UnsafePointer() == p {
				return positive
			}
		}
		return !positive
	}

	return false
}
