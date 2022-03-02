package tests

import (
	. "github.com/sxyazi/go-collection"
	"math"
	"testing"
)

func TestSlice_All(t *testing.T) {
	d := []int{1, 2, 3}

	items := UseSlice(d).All()
	if !UseSlice(items).Same(d) {
		t.Fail()
	}
}

func TestSlice_Len(t *testing.T) {
	var d1 []int
	if UseSlice(d1).Len() != 0 {
		t.Fail()
	}

	d2 := [3]int{1, 2, 3}
	if UseSlice(d2[:]).Len() != 3 {
		t.Fail()
	}
}

func TestSlice_Each(t *testing.T) {
	data := []float64{0, 2.71, 3.14}
	result := []float64{0, 0, 0}

	if !UseSlice(data).Each(func(value float64, index int) {
		result[index] = value
	}).Same(result) {
		t.Fail()
	}
}

func TestSlice_Empty(t *testing.T) {
	if !UseSlice([]int{}).Empty() {
		t.Fail()
	}

	if UseSlice([]float64{0, 2.71, 3.14}).Empty() {
		t.Fail()
	}
}

func TestSlice_Same(t *testing.T) {
	if !UseSlice[[]int, int](nil).Same(nil) {
		t.Fail()
	}

	if !UseSlice([]int{1, 2, 3}).Same([]int{1, 2, 3}) {
		t.Fail()
	}
	if UseSlice([]int{1, 2, 3}).Same([]int{1, 3}) {
		t.Fail()
	}
	if !UseSlice([]int{}).Same([]int{}) {
		t.Fail()
	}

	f1 := Foo{}
	f2 := Foo{}
	if !UseSlice([]Foo{f1, f2}).Same([]Foo{f2, f1}) {
		t.Fail()
	}
	if UseSlice([]*Foo{&f1, &f2}).Same([]*Foo{&f2, &f1}) {
		t.Fail()
	}

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{3, 2, 1}
	s4 := [][]int{s1, s2}
	if !UseSlice(s4).Same(s4) {
		t.Fail()
	}
	if !UseSlice(s4).Same([][]int{s2, s1}) {
		t.Fail()
	}
	if UseSlice(s4).Same([][]int{s1, s3}) {
		t.Fail()
	}

	if !UseSlice([]float64{math.NaN(), math.NaN()}).Same([]float64{math.NaN(), math.NaN()}) {
		t.Fail()
	}
	if UseSlice([]float64{math.NaN(), math.NaN()}).Same([]float64{0, math.NaN()}) {
		t.Fail()
	}
}

func TestSlice_First(t *testing.T) {
	data := []float64{32, 2.71, 3.14}

	if v, ok := UseSlice(data).First(); !ok || v != 32 {
		t.Fail()
	}
}

func TestSlice_Last(t *testing.T) {
	data := []float64{32, 2.71, 3.14}

	if v, ok := UseSlice(data).Last(); !ok || v != 3.14 {
		t.Fail()
	}
}

func TestSlice_Index(t *testing.T) {
	// Nil
	if v := UseSlice[[]int, int](nil).Index(1); v != -1 {
		t.Fail()
	}

	// Integer
	d1 := []int{1, 2, 3}
	if v := UseSlice(d1).Index(2); v != 1 {
		t.Fail()
	}
	if v := UseSlice(d1).Index(10); v != -1 {
		t.Fail()
	}

	// Float
	d2 := []float64{32, 2.71, 3.14}
	if v := UseSlice(d2).Index(2.71); v != 1 {
		t.Fail()
	}

	// String
	d3 := []string{"a", "b", "c"}
	if v := UseSlice(d3).Index("d"); v != -1 {
		t.Fail()
	}

	// Struct
	f1 := Foo{}
	f2 := Foo{Bar: "b"}
	d4 := []Foo{{Bar: "xx"}, f1, f2}
	if v := UseSlice(d4).Index(f2); v != 2 {
		t.Fail()
	}
	if v := UseSlice(d4).Index(Foo{}); v != 1 {
		t.Fail()
	}

	// Nested slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	d5 := [][]int{s1, s2}
	if v := UseSlice(d5).Index(s2); v != 1 {
		t.Fail()
	}
	if v := UseSlice(d5).Index([]int{4, 5, 6}); v != -1 {
		t.Fail()
	}
}

func TestSlice_Contains(t *testing.T) {
	// Integer
	d1 := []int{1, 2, 3}
	if !UseSlice(d1).Contains(1) {
		t.Fail()
	}

	// Float
	d2 := []float64{32, 2.71, 3.14}
	if !UseSlice(d2).Contains(2.71) {
		t.Fail()
	}

	// String
	d3 := []string{"a", "b", "c"}
	if !UseSlice(d3).Contains("a") {
		t.Fail()
	}

	// Struct
	d4 := []Foo{{Bar: "xx"}, {Bar: "b"}, {Bar: "c"}}
	if !UseSlice(d4).Contains(Foo{Bar: "b"}) {
		t.Fail()
	}

	// Nested slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	d5 := [][]int{s1, s2}
	if !UseSlice(d5).Contains(s1) {
		t.Fail()
	}
	if UseSlice(d5).Contains([]int{1, 2, 3}) {
		t.Fail()
	}
}

func TestSlice_Diff(t *testing.T) {
	d1 := []int{1, 2, 3, 4, 5}
	d2 := []int{2, 4, 6, 8}
	if !UseSlice(d1).Diff(d2).Same([]int{1, 3, 5}) {
		t.Fail()
	}
}

func TestSlice_Filter(t *testing.T) {
	d1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !UseSlice(d1).Filter(func(value int, index int) bool {
		return value%2 == 0
	}).Same([]int{2, 4, 6, 8}) {
		t.Fail()
	}
}

func TestSlice_Map(t *testing.T) {
	d1 := []float64{0, 2, 3, 4, 5}
	if !UseSlice(d1).Map(func(value float64, index int) float64 {
		return value * 3.14
	}).Same([]float64{0, 6.28, 9.42, 12.56, 15.7}) {
		t.Fail()
	}
}

func TestSlice_Unique(t *testing.T) {
	if !UseSlice([]int{1, 2, 2, 3}).Unique().Same([]int{1, 2, 3}) {
		t.Fail()
	}

	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6}
	if !UseSlice([][]int{s1, s2, s1}).Unique().Same([][]int{s1, s2}) {
		t.Fail()
	}
	if !UseSlice([][]int{s1, nil, s2, nil, s1}).Unique().Same([][]int{s1, nil, s2}) {
		t.Fail()
	}
	if !UseSlice([]*[]int{&s1, &s2, &s1}).Unique().Same([]*[]int{&s1, &s2}) {
		t.Fail()
	}

	s3, s4 := &[]int{1, 2, 3}, &[]int{4, 5, 6}
	s5 := (*s4)[:2]
	if !UseSlice([]*[]int{s3, s4, &s5}).Unique().Same([]*[]int{s3, s4, &s5}) {
		t.Fail()
	}

	if !UseSlice([]any{s1, s1, s2, s3, s3, s4, s5, &s5}).Unique().Same([]any{s1, s2, s3, s4, s5, &s5}) {
		t.Fail()
	}
}

func TestSlice_Duplicates(t *testing.T) {
	if !UseSlice([]int{1, 2, 2, 3}).Duplicates().Same(map[int]int{2: 2}) {
		t.Fail()
	}
	if !UseSlice([]string{"a", "b", "a", "c"}).Duplicates().Same(map[int]string{2: "a"}) {
		t.Fail()
	}

	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6}
	if !UseSlice([][]int{s1, s2, s1}).Duplicates().Same(map[int][]int{2: s1}) {
		t.Fail()
	}
	if !UseSlice([][]int{s1, nil, s2, nil, s1}).Duplicates().Same(map[int][]int{3: nil, 4: s1}) {
		t.Fail()
	}
	if !UseSlice([]*[]int{&s1, &s2, &s1}).Duplicates().Same(map[int]*[]int{2: &s1}) {
		t.Fail()
	}

	s3, s4 := &[]int{1, 2, 3}, &[]int{4, 5, 6}
	s5 := (*s4)[:2]
	if !UseSlice([]*[]int{s3, s4, &s5}).Duplicates().Same(map[int]*[]int{}) {
		t.Fail()
	}

	if !UseSlice([]any{s1, s1, s2, s3, s3, s4, s5, &s5}).Duplicates().Same(map[int]any{1: s1, 4: s3}) {
		t.Fail()
	}
}

func TestSlice_Merge(t *testing.T) {
	if !UseSlice([]int{1, 2}).Merge([]int{3, 4}).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2}).Merge([]int{3, 4}, []int{5, 6}).Same([]int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestSlice_Random(t *testing.T) {
	if v, ok := UseSlice([]int{}).Random(); ok || v != 0 {
		t.Fail()
	}
	if v, ok := UseSlice([]int{1}).Random(); !ok || v == 0 {
		t.Fail()
	}
}

func TestSlice_Reverse(t *testing.T) {
	if !UseSlice([]int{1, 2}).Reverse().Same([]int{2, 1}) {
		t.Fail()
	}
	if UseSlice([]any{}).Reverse().Len() != 0 {
		t.Fail()
	}
}

func TestSlice_Shuffle(t *testing.T) {
	if UseSlice([]any{}).Shuffle().Len() != 0 {
		t.Fail()
	}
	if v, ok := UseSlice([]int{1}).Shuffle().First(); !ok || v != 1 {
		t.Fail()
	}

	s1 := UseSlice([]int{1, 2}).Shuffle().All()
	if !UseSlice(s1).Same([]int{1, 2}) && !UseSlice(s1).Same([]int{2, 1}) {
		t.Fail()
	}
}

func TestSlice_Slice(t *testing.T) {
	d := []int{1, 2, 3, 4}

	// Normal
	if !UseSlice(d).Slice(0, 0).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(0, 2).Same([]int{1, 2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(2, 2).Same([]int{3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(0, 4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}

	// Offset out of range
	if !UseSlice(d).Slice(4, 0).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(4, 2).Same([]int{}) {
		t.Fail()
	}

	// (offset + length) out of range
	if !UseSlice(d).Slice(3, 2).Same([]int{4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(0, 5).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(0, 100).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}

	// Negative offset
	if !UseSlice(d).Slice(-2, 2).Same([]int{3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4, 2).Same([]int{1, 2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4, 4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4, 5).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-5, 2).Same([]int{}) {
		t.Fail()
	}

	// Negative length
	if !UseSlice(d).Slice(0, -2).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(0, -10).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(1, -1).Same([]int{2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(1, -10).Same([]int{1, 2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(3, -4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(3, -10).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(4, -10).Same([]int{}) {
		t.Fail()
	}

	// Negative offset and length
	if !UseSlice(d).Slice(-1, -1).Same([]int{4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-1, -2).Same([]int{3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-1, -4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-1, -10).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-3, -1).Same([]int{2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-3, -10).Same([]int{1, 2}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4, -1).Same([]int{1}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4, -10).Same([]int{1}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-5, -1).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-5, -10).Same([]int{}) {
		t.Fail()
	}

	// Pass only offset
	if !UseSlice(d).Slice(0).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(1).Same([]int{2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(3).Same([]int{4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(4).Same([]int{}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-1).Same([]int{4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-3).Same([]int{2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice(d).Slice(-5).Same([]int{}) {
		t.Fail()
	}
}

func TestSlice_Split(t *testing.T) {
	d := []int{1, 2, 3, 4, 5}
	if !UseSlice(UseSlice(d).Split(2)).Same([][]int{{1, 2}, {3, 4}, {5}}) {
		t.Fail()
	}
}

func TestSlice_Splice(t *testing.T) {
	test := func(offset int, args ...any) *SliceCollection[[]int, int] {
		s := UseSlice([]int{1, 2, 3, 4})
		chunk := s.Splice(offset, args...)

		s2 := []int{1, 2, 3, 4}
		var start, end int
		if len(args) >= 1 {
			start, end = OffsetToIndex(len(s2), offset, args[0].(int))
		} else {
			start, end = OffsetToIndex(len(s2), offset)
		}

		if !s.Same(append(s2[:start], s2[end:]...)) {
			t.Fail()
		}

		return chunk
	}

	// Normal offset
	if !test(0).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !test(1).Same([]int{2, 3, 4}) {
		t.Fail()
	}
	if !test(3).Same([]int{4}) {
		t.Fail()
	}
	if !test(4).Same([]int{}) {
		t.Fail()
	}
	if !test(-1).Same([]int{4}) {
		t.Fail()
	}
	if !test(-3).Same([]int{2, 3, 4}) {
		t.Fail()
	}
	if !test(-4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}

	// Offset out of range
	if !test(5).Same([]int{}) {
		t.Fail()
	}
	if !test(10).Same([]int{}) {
		t.Fail()
	}
	if !test(-5).Same([]int{}) {
		t.Fail()
	}

	// Normal length
	if !test(0, 1).Same([]int{1}) {
		t.Fail()
	}
	if !test(0, 3).Same([]int{1, 2, 3}) {
		t.Fail()
	}
	if !test(0, 4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !test(3, -4).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}

	// Length out of range
	if !test(0, 5).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !test(0, 10).Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}
	if !test(0, -1).Same([]int{}) {
		t.Fail()
	}
	if !test(4, -1).Same([]int{}) {
		t.Fail()
	}
	if !test(4, -4).Same([]int{}) {
		t.Fail()
	}

	// Replacement
	s := UseSlice([]int{1, 2, 3, 4})
	if !s.Splice(1, 2, []int{22, 33}).Same([]int{2, 3}) || !s.Same([]int{1, 22, 33, 4}) {
		t.Fail()
	}
	s = UseSlice([]int{1, 2, 3, 4})
	if !s.Splice(1, 0, []int{22, 33}).Same([]int{}) || !s.Same([]int{1, 22, 33, 2, 3, 4}) {
		t.Fail()
	}
	s = UseSlice([]int{1, 2, 3, 4})
	if !s.Splice(1, 2, 22, 33).Same([]int{2, 3}) || !s.Same([]int{1, 22, 33, 4}) {
		t.Fail()
	}
	s = UseSlice([]int{1, 2, 3, 4})
	if !s.Splice(-4, 4, 11, 22, 33, 44).Same([]int{1, 2, 3, 4}) || !s.Same([]int{11, 22, 33, 44}) {
		t.Fail()
	}
}

func TestSlice_Reduce(t *testing.T) {
	if UseSlice([]int{1, 2, 3}).Reduce(100, func(carry, value, key int) int {
		return carry + value
	}) != 106 {
		t.Fail()
	}
}

func TestSlice_Pop(t *testing.T) {
	c := UseSlice([]int{1, 2})
	if v, ok := c.Pop(); !ok || v != 2 {
		t.Fail()
	}
	if v, ok := c.Pop(); !ok || v != 1 {
		t.Fail()
	}
	if !c.Empty() {
		t.Fail()
	}
	if v, ok := c.Pop(); ok || v != 0 {
		t.Fail()
	}

	// Functional test
	d := []int{1, 2}
	Pop(&d)
	if !Same(d, []int{1}) {
		t.Fail()
	}
}

func TestSlice_Push(t *testing.T) {
	c := UseSlice([]int{1, 2, 3})
	c.Push(4)
	if !c.Same([]int{1, 2, 3, 4}) {
		t.Fail()
	}

	// Functional test
	d := []int{1, 2, 3}
	Push(&d, 4)
	if !Same(d, []int{1, 2, 3, 4}) {
		t.Fail()
	}
}

func TestSlice_Where(t *testing.T) {
	// Only target
	if !UseSlice([]int{1, 2, 3}).Where(2).Same([]int{2}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3}).Where(4).Same([]int{}) {
		t.Fail()
	}

	// Operator and target
	if !UseSlice([]int{1, 2, 3}).Where("!=", 4).Same([]int{1, 2, 3}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3}).Where("!=", 2).Same([]int{1, 3}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3, 4}).Where(">", 2).Same([]int{3, 4}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3, 4}).Where(">=", 2).Same([]int{2, 3, 4}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3, 4}).Where("<", 3).Same([]int{1, 2}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3, 4}).Where("<=", 3).Same([]int{1, 2, 3}) {
		t.Fail()
	}

	u1, u2, u3, u4 := User{1, "Hugo"}, User{2, "Lisa"}, User{3, "Iris"}, User{4, "Lisa"}
	if !UseSlice([]User{u1, u2, u3, u4}).Where("!=", u2).Same([]User{u1, u3, u4}) {
		t.Fail()
	}
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).Where("!=", &u2).Same([]*User{&u1, &u3, &u4}) {
		t.Fail()
	}
	if UseSlice([]*User{&u1, &u2, &u3, &u4}).Where("!=", u2).Same([]*User{&u1, &u3, &u4}) {
		t.Fail()
	}

	// Key and target
	d1 := []User{u1, u2, u3, u4}
	if !UseSlice(d1).Where("Name", "Lisa").Same([]User{{2, "Lisa"}, {4, "Lisa"}}) {
		t.Fail()
	}

	// Key, operator and target
	if !UseSlice(d1).Where("Name", "!=", "Lisa").Same([]User{{1, "Hugo"}, {3, "Iris"}}) {
		t.Fail()
	}
}

func TestSlice_WhereIn(t *testing.T) {
	// Only targets
	if !UseSlice([]int{1, 2, 3}).WhereIn([]int{1, 3}).Same([]int{1, 3}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3}).WhereIn([]int{0, 2, -1}).Same([]int{2}) {
		t.Fail()
	}

	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6}
	if !UseSlice([][]int{s1, s2}).WhereIn([][]int{s2}).Same([][]int{s2}) {
		t.Fail()
	}

	u1, u2, u3, u4 := User{1, "Hugo"}, User{2, "Lisa"}, User{3, "Iris"}, User{4, "Lisa"}
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).WhereIn([]*User{nil, &u2, &u1}).Same([]*User{&u1, &u2}) {
		t.Fail()
	}
	if !UseSlice([]User{u1, u2, u3, u4}).WhereIn([]User{{5, "Kite"}, u2, u1}).Same([]User{u1, u2}) {
		t.Fail()
	}

	// Key and targets
	d1 := []User{u1, u2, u3, u4}
	if !UseSlice(d1).WhereIn("Name", []string{"Iris", "Hugo"}).Same([]User{{1, "Hugo"}, {3, "Iris"}}) {
		t.Fail()
	}

	// Nil
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).WhereIn(nil).Same([]*User{}) {
		t.Fail()
	}
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).WhereIn([]*User{nil, &u3}).Same([]*User{&u3}) {
		t.Fail()
	}
}

func TestSlice_WhereNotIn(t *testing.T) {
	// Only targets
	if !UseSlice([]int{1, 2, 3}).WhereNotIn([]int{1, 3}).Same([]int{2}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3}).WhereNotIn([]int{0, 2, -1}).Same([]int{1, 3}) {
		t.Fail()
	}
	if !UseSlice([]int{1, 2, 3}).WhereNotIn([]float64{1, 2, 3.14}).Same([]int{1, 2, 3}) {
		t.Fail()
	}
	if !UseSlice([]float64{1, 2, 3}).WhereNotIn([]int{1, 2, 3}).Same([]float64{1, 2, 3}) {
		t.Fail()
	}
	if !UseSlice([]int64{1, 2, 3}).WhereNotIn([]int8{1, 2, 3}).Same([]int64{}) {
		t.Fail()
	}

	u1, u2, u3, u4 := User{1, "Hugo"}, User{2, "Lisa"}, User{3, "Iris"}, User{4, "Lisa"}
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).WhereNotIn([]*User{nil, &u2, &u1}).Same([]*User{&u3, &u4}) {
		t.Fail()
	}

	// Key and targets
	d1 := []User{u1, u2, u3, u4}
	if !UseSlice(d1).WhereNotIn("Name", []string{"Lisa"}).Same([]User{{1, "Hugo"}, {3, "Iris"}}) {
		t.Fail()
	}

	// Nil
	if !UseSlice([]*User{&u1, &u2, &u3, &u4}).WhereNotIn(nil).Same([]*User{&u1, &u2, &u3, &u4}) {
		t.Fail()
	}
}
