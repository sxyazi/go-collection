package tests

import (
	collect "go-collection"
	"testing"
)

func TestCollect_Each(t *testing.T) {
	data := []float64{0, 2.71, 3.14}
	result := []float64{0, 0, 0}

	if !collect.Slice(data).Each(func(value float64, key int) {
		result[key] = value
	}).Same(result) {
		t.Fail()
	}
}

func TestCollect_Empty(t *testing.T) {
	if !collect.Slice([]int{}).Empty() {
		t.Fail()
	}

	if collect.Slice([]float64{0, 2.71, 3.14}).Empty() {
		t.Fail()
	}
}

func TestCollect_Same(t *testing.T) {
	if !collect.Slice([]int{1, 2, 3}).Same([]int{1, 2, 3}) {
		t.Fail()
	}

	if !collect.Slice([]int{}).Same([]int{}) {
		t.Fail()
	}

	f1 := Foo{}
	f2 := Foo{}

	if !collect.Slice([]Foo{f1, f2}).Same([]Foo{f2, f1}) {
		t.Fail()
	}

	if collect.Slice([]*Foo{&f1, &f2}).Same([]*Foo{&f2, &f1}) {
		t.Fail()
	}
}
