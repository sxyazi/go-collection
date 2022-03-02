package collect

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"strconv"
)

func StringToNumber[T constraints.Integer | constraints.Float](s string) (result T, _ error) {
	switch any(result).(type) {
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

func OffsetToIndex(actual, offset int, args ...int) (int, int) {
	length := actual
	if len(args) >= 1 {
		length = args[0]
	}
	if actual == 0 || (offset == 0 && length < 0) {
		return 0, 0
	}

	// negative offset
	if offset < 0 {
		offset += actual
	}
	if offset >= actual || offset < 0 {
		return 0, 0
	} else if length == 0 {
		return offset, offset
	}

	// negative length
	if length < 0 {
		if offset+length < 0 {
			offset, length = 0, offset+1
		} else {
			offset, length = offset+length+1, -length
		}
	}

	length = int(math.Min(float64(length), float64(actual-offset)))
	return offset, offset + length
}

func NumberFrom[N constraints.Integer | constraints.Float, T ~[]E, E any](c *SliceCollection[T, E]) *NumberCollection[[]N, N] {
	if c.Empty() {
		return &NumberCollection[[]N, N]{}
	}

	z := c.All()
	items := make([]N, len(z), cap(z))
	for key, item := range z {
		switch v := (interface{})(item).(type) {
		case string:
			items[key], _ = StringToNumber[N](v)
		default:
			items[key], _ = StringToNumber[N](fmt.Sprintf("%d", v))
		}
	}

	return UseNumber[[]N, N](items)
}
