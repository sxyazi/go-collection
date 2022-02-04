package tests

import (
	collect "go-collection"
	"testing"
)

func TestNumberFrom(t *testing.T) {
	c1 := collect.Any([]string{"1", "2", "Hello", "3"})
	if collect.NumberFrom[float64](c1).Avg() != 1.5 {
		t.Fail()
	}

	c2 := collect.Any([]int32{392, 68, 27, 0})
	if collect.NumberFrom[uint](c2).Avg() != 121 {
		t.Fail()
	}

	c3 := collect.Any([]Foo{Foo{}})
	if collect.NumberFrom[uint](c3).Sum() != 0 {
		t.Fail()
	}
}
