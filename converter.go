package collect

import (
	"constraints"
	"errors"
	"fmt"
	"strconv"
)

func StringToNumber[T constraints.Integer | constraints.Float](s string) (result T, _ error) {
	switch interface{}(result).(type) {
	// Signed
	case int:
		n, err := strconv.Atoi(s)
		return T(n), err
	case int8:
		n, err := strconv.ParseInt(s, 10, 8)
		return T(n), err
	case int16:
		n, err := strconv.ParseInt(s, 10, 16)
		return T(n), err
	case int32:
		n, err := strconv.ParseInt(s, 10, 32)
		return T(n), err
	case int64:
		n, err := strconv.ParseInt(s, 10, 64)
		return T(n), err

	// Unsigned
	case uint:
		n, err := strconv.ParseUint(s, 10, 0)
		return T(n), err
	case uint8:
		n, err := strconv.ParseUint(s, 10, 8)
		return T(n), err
	case uint16:
		n, err := strconv.ParseUint(s, 10, 16)
		return T(n), err
	case uint32:
		n, err := strconv.ParseUint(s, 10, 32)
		return T(n), err
	case uint64:
		n, err := strconv.ParseUint(s, 10, 64)
		return T(n), err
	case uintptr:
		return result, errors.New("conversion failed")

	// Float
	case float32:
		n, err := strconv.ParseFloat(s, 32)
		return T(n), err
	case float64:
		n, err := strconv.ParseFloat(s, 64)
		return T(n), err
	}

	return
}

func NumberFrom[N constraints.Integer | constraints.Float, T []E, E comparable](c *sliceCollection[T, E]) *numberCollection[[]N, N] {
	if c.Empty() {
		return &numberCollection[[]N, N]{}
	}

	z := c.Items()
	items := make([]N, len(z), cap(z))
	for key, item := range z {
		switch v := (interface{})(item).(type) {
		case string:
			items[key], _ = StringToNumber[N](v)
		default:
			items[key], _ = StringToNumber[N](fmt.Sprintf("%d", v))
		}
	}

	return &numberCollection[[]N, N]{sliceCollection[[]N, N]{z: items}}
}
