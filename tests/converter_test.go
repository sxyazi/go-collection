package tests

import (
	. "github.com/sxyazi/go-collection"
	"testing"
)

// TODO
func TestStringToNumber(t *testing.T) {

}

func TestNumberFrom(t *testing.T) {
	c1 := UseSlice([]string{"1", "2", "Hello", "3"})
	if NumberFrom[float64](c1).Avg() != 1.5 {
		t.Fail()
	}

	c2 := UseSlice([]int32{392, 68, 27, 0})
	if NumberFrom[uint](c2).Avg() != 121.75 {
		t.Fail()
	}

	c3 := UseSlice([]Foo{{}})
	if NumberFrom[uint](c3).Sum() != 0 {
		t.Fail()
	}
}
