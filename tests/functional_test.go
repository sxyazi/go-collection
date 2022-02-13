package tests

import (
	. "github.com/sxyazi/go-collection"
	"strconv"
	"testing"
)

func TestFunctional_Len(t *testing.T) {
	if Len(nil) != -1 {
		t.Fail()
	}

	if Len([...]int{}) != 0 {
		t.Fail()
	}

	c := make(chan int, 10)
	c <- 1
	if Len(c) != 1 {
		t.Fail()
	}

	if Len(map[int]bool{1: true, 2: false}) != 2 {
		t.Fail()
	}

	if Len([]int{1, 2, 3}) != 3 {
		t.Fail()
	}

	if Len("Hello") != 5 {
		t.Fail()
	}

	if Len(struct{}{}) != -1 {
		t.Fail()
	}
}

func TestFunctional_Empty(t *testing.T) {
	if !Empty([]int{}) {
		t.Fail()
	}

	if Empty([]int{1, 2, 3}) {
		t.Fail()
	}
}

func TestFunctional_Count(t *testing.T) {
	m := Count([]int{1, 2, 2, 3})
	if m[1] != 1 || m[2] != 2 || m[3] != 1 {
		t.Fail()
	}
}

func TestFunctional_Times(t *testing.T) {
	if !Times(3, func(number int) float64 {
		return float64(number) * 3.14
	}).Same([]float64{3.14, 6.28, 9.42}) {
		t.Fail()
	}
}

func TestFunctional_SortBy(t *testing.T) {
	if !SortBy([]int{2, 1, 3}, func(item, index int) string {
		return strconv.Itoa(item)
	}).Same([]int{1, 2, 3}) {
		t.Fail()
	}
}

func TestFunctional_SortByDesc(t *testing.T) {
	if !SortByDesc([]int{2, 1, 3}, func(item, index int) string {
		return strconv.Itoa(item)
	}).Same([]int{3, 2, 1}) {
		t.Fail()
	}
}
