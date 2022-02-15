package collect

import (
	"golang.org/x/exp/constraints"
	"math"
	"reflect"
)

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
	case "=", "!=":
		var eq bool
		switch av := any(a).(type) {
		case float64:
			if math.IsNaN(av) && math.IsNaN(any(b).(float64)) {
				eq = true
			} else {
				eq = math.Abs(av-any(b).(float64)) <= 1e-9
			}
		case float32:
			eq = math.Abs(float64(av)-float64(any(b).(float32))) <= 1e-9
		default:
			eq = a == b
		}

		if operator == "=" {
			return eq
		} else {
			return !eq
		}
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

	ar, br := reflect.ValueOf(a), reflect.ValueOf(b)
	switch true {
	case ar.CanInt():
		if !br.CanInt() {
			return operator != "="
		}
		return NumberCompare(ar.Int(), operator, br.Int())
	case ar.CanUint():
		if !br.CanUint() {
			return operator != "="
		}
		return NumberCompare(ar.Uint(), operator, br.Uint())
	case ar.CanFloat():
		if !br.CanFloat() {
			return operator != "="
		}
		return NumberCompare(ar.Float(), operator, br.Float())
	}

	return false
}

func Compare(a any, operator string, b any) bool {
	if a == nil && b == nil {
		return operator == "="
	} else if a == nil || b == nil {
		return operator != "="
	}

	if IsNumber(a) || IsNumber(b) {
		return AnyNumberCompare(a, operator, b)
	} else if operator != "=" && operator != "!=" {
		return false
	}

	ar, br := reflect.TypeOf(a), reflect.TypeOf(b)
	ak, bk := ar.Kind(), br.Kind()
	if ak != bk {
		return operator == "!="
	}

	if ak != reflect.Slice && ak != reflect.Map && ak != reflect.Func {
		switch operator {
		case "=":
			return a == b
		case "!=":
			return a != b
		}
	}

	p := reflect.ValueOf(a).UnsafePointer()
	switch operator {
	case "=":
		return p == reflect.ValueOf(b).UnsafePointer()
	case "!=":
		return p != reflect.ValueOf(b).UnsafePointer()
	}

	return false
}

type ComparisonSet struct {
	LooseNumber bool
	z           map[any]map[reflect.Kind]struct{}
}

func (c *ComparisonSet) Normalize(v reflect.Value) (reflect.Kind, any) {
	kind := v.Kind()
	if kind == reflect.Slice || kind == reflect.Func || kind == reflect.Map {
		return kind, v.UnsafePointer()
	}

	if c.LooseNumber {
		switch true {
		case v.CanInt():
			return reflect.Int64, v.Int()
		case v.CanUint():
			return reflect.Uint64, v.Uint()
		case v.CanFloat():
			return reflect.Float64, v.Float()
		}
	}

	return kind, v.Interface()
}

func (c *ComparisonSet) Add(v any) {
	kind, value := c.Normalize(reflect.ValueOf(v))
	if _, ok := c.z[value]; !ok {
		c.z[value] = make(map[reflect.Kind]struct{})
	}

	c.z[value][kind] = struct{}{}
}

func (c *ComparisonSet) Has(v any) bool {
	kind, value := c.Normalize(reflect.ValueOf(v))
	if m, ok := c.z[value]; ok {
		_, ok := m[kind]
		return ok
	}

	return false
}

func NewComparisonSet(looseNumber bool) *ComparisonSet {
	return &ComparisonSet{looseNumber, make(map[any]map[reflect.Kind]struct{})}
}
