package tests

import (
	. "go-collection"
	"testing"
)

func TestSlice_Items(t *testing.T) {
	d := []int{1, 2, 3}

	items := Slice(d).Items()
	if !Slice(items).Same(d) {
		t.Fail()
	}
}

func TestSlice_Len(t *testing.T) {
	if Slice([]int{}).Len() != 0 {
		t.Fail()
	}
	if Slice([]int{1, 2, 3}).Len() != 3 {
		t.Fail()
	}
}

func TestSlice_Each(t *testing.T) {
	data := []float64{0, 2.71, 3.14}
	result := []float64{0, 0, 0}

	if !Slice(data).Each(func(value float64, index int) {
		result[index] = value
	}).Same(result) {
		t.Fail()
	}
}

func TestSlice_Empty(t *testing.T) {
	if !Slice([]int{}).Empty() {
		t.Fail()
	}

	if Slice([]float64{0, 2.71, 3.14}).Empty() {
		t.Fail()
	}
}

func TestSlice_Same(t *testing.T) {
	if !Slice([]int{1, 2, 3}).Same([]int{1, 2, 3}) {
		t.Fail()
	}
	if Slice([]int{1, 2, 3}).Same([]int{1, 3}) {
		t.Fail()
	}
	if !Slice([]int{}).Same([]int{}) {
		t.Fail()
	}

	f1 := Foo{}
	f2 := Foo{}
	if !Slice([]Foo{f1, f2}).Same([]Foo{f2, f1}) {
		t.Fail()
	}
	if Slice([]*Foo{&f1, &f2}).Same([]*Foo{&f2, &f1}) {
		t.Fail()
	}

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{3, 2, 1}
	s4 := [][]int{s1, s2}
	if !Slice(s4).Same(s4) {
		t.Fail()
	}
	if !Slice(s4).Same([][]int{s2, s1}) {
		t.Fail()
	}
	if Slice(s4).Same([][]int{s1, s3}) {
		t.Fail()
	}
}

func TestSlice_First(t *testing.T) {
	data := []float64{32, 2.71, 3.14}

	if v, ok := Slice(data).First(); !ok || v != 32 {
		t.Fail()
	}
}

func TestSlice_Last(t *testing.T) {
	data := []float64{32, 2.71, 3.14}

	if v, ok := Slice(data).Last(); !ok || v != 3.14 {
		t.Fail()
	}
}

func TestSlice_Index(t *testing.T) {
	// Integer
	d1 := []int{1, 2, 3}
	if v := Slice(d1).Index(10); v != -1 {
		t.Fail()
	}

	// Float
	d2 := []float64{32, 2.71, 3.14}
	if v := Slice(d2).Index(2.71); v != 1 {
		t.Fail()
	}

	// String
	d3 := []string{"a", "b", "c"}
	if v := Slice(d3).Index("d"); v != -1 {
		t.Fail()
	}

	// Struct
	f1 := Foo{}
	f2 := Foo{Bar: "b"}
	d4 := []Foo{{Bar: "xx"}, f1, f2}
	if v := Slice(d4).Index(f2); v != 2 {
		t.Fail()
	}
	if v := Slice(d4).Index(Foo{}); v != 1 {
		t.Fail()
	}

	// Nested slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	d5 := [][]int{s1, s2}
	if v := Slice(d5).Index(s2); v != 1 {
		t.Fail()
	}
	if v := Slice(d5).Index([]int{4, 5, 6}); v != -1 {
		t.Fail()
	}
}

func TestSlice_Contains(t *testing.T) {
	// Integer
	d1 := []int{1, 2, 3}
	if !Slice(d1).Contains(1) {
		t.Fail()
	}

	// Float
	d2 := []float64{32, 2.71, 3.14}
	if !Slice(d2).Contains(2.71) {
		t.Fail()
	}

	// String
	d3 := []string{"a", "b", "c"}
	if !Slice(d3).Contains("a") {
		t.Fail()
	}

	// Struct
	d4 := []Foo{{Bar: "xx"}, {Bar: "b"}, {Bar: "c"}}
	if !Slice(d4).Contains(Foo{Bar: "b"}) {
		t.Fail()
	}

	// Nested slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	d5 := [][]int{s1, s2}
	if !Slice(d5).Contains(s1) {
		t.Fail()
	}
	if Slice(d5).Contains([]int{1, 2, 3}) {
		t.Fail()
	}
}

func TestSlice_Diff(t *testing.T) {
	d1 := []int{1, 2, 3, 4, 5}
	d2 := []int{2, 4, 6, 8}
	if !Slice(d1).Diff(d2).Same([]int{1, 3, 5}) {
		t.Fail()
	}
}

func TestSlice_Filter(t *testing.T) {
	d1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !Slice(d1).Filter(func(value int, index int) bool {
		return value%2 == 0
	}).Same([]int{2, 4, 6, 8}) {
		t.Fail()
	}
}

func TestSlice_Map(t *testing.T) {
	d1 := []float64{0, 2, 3, 4, 5}
	if !Slice(d1).Map(func(value float64, index int) float64 {
		return value * 3.14
	}).Same([]float64{0, 6.28, 9.42, 12.56, 15.7}) {
		t.Fail()
	}
}

// TODO
func TestSlice_Unique(t *testing.T) {

}

// TODO
func TestSlice_Merge(t *testing.T) {

}

// TODO
func TestSlice_Random(t *testing.T) {

}

// TODO
func TestSlice_Reverse(t *testing.T) {

}

// TODO
func TestSlice_Shuffle(t *testing.T) {

}

// TODO
func TestSlice_Slice(t *testing.T) {

}

// TODO
func TestSlice_Split(t *testing.T) {

}

// TODO
func TestSlice_Splice(t *testing.T) {

}

// TODO
func TestSlice_Count(t *testing.T) {

}
