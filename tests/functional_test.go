package tests

import (
	. "github.com/sxyazi/go-collection"
	"testing"
)

func TestFunctional_Len(t *testing.T) {
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
