package tests

import (
	. "github.com/sxyazi/go-collection"
	"testing"
)

func TestMap_All(t *testing.T) {
	d := map[string]int{"foo": 1, "bar": 0}

	items := UseMap(d).All()
	if !UseMap(items).Same(d) {
		t.Fail()
	}
}

func TestMap_New(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	d2 := map[string]int{"foo": 1, "bar": 0}

	m1 := UseMap(d1)
	m2 := m1.New(d2)

	m2.Put("bar", 100)
	if !m1.Same(d1) || m2.Same(d1) {
		t.Fail()
	}
}

func TestMap_Len(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if UseMap(d1).Len() != 2 {
		t.Fail()
	}

	d2 := map[string]int{}
	if UseMap(d2).Len() != 0 {
		t.Fail()
	}
}

func TestMap_Empty(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if UseMap(d1).Empty() {
		t.Fail()
	}

	d2 := map[string]int{}
	if !UseMap(d2).Empty() {
		t.Fail()
	}
}

func TestMap_Only(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if !UseMap(d1).Only("foo", "bar").Same(d1) {
		t.Fail()
	}
	if !UseMap(d1).Only("bar").Same(map[string]int{"bar": 0}) {
		t.Fail()
	}

	d2 := map[string]int{}
	if UseMap(d2).Only("foo", "bar").Same(d2) {
		t.Fail()
	}
}

func TestMap_Except(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if !UseMap(d1).Except("foo", "bar").Empty() {
		t.Fail()
	}
	if !UseMap(d1).Except("bar").Same(map[string]int{"foo": 1}) {
		t.Fail()
	}

	d2 := map[string]int{}
	if !UseMap(d2).Except("foo", "bar").Same(d2) {
		t.Fail()
	}
}

func TestMap_Keys(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0, "baz": 2}
	c := UseSlice(UseMap(d1).Keys())
	if c.Len() != 3 || !c.Contains("foo") || !c.Contains("bar") || !c.Contains("baz") {
		t.Fail()
	}

	d2 := map[float64]int{}
	if !UseSlice(UseMap(d2).Keys()).Same([]float64{}) {
		t.Fail()
	}
}

func TestMap_DiffKeys(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0, "baz": 2}
	d2 := map[string]int{"foo": 1, "bar": 0}
	if !UseMap(d1).DiffKeys(d2).Same(map[string]int{"baz": 2}) {
		t.Fail()
	}

	d3 := map[string]int{}
	if !UseMap(d1).DiffKeys(d3).Same(d1) {
		t.Fail()
	}
}

func TestMap_Has(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if !UseMap(d1).Has("foo") {
		t.Fail()
	}
	if UseMap(d1).Has("baz") {
		t.Fail()
	}

	d2 := map[string]int{}
	if UseMap(d2).Has("foo") {
		t.Fail()
	}
}

func TestMap_Get(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 0}
	if v, ok := UseMap(d1).Get("foo"); !ok || v != 1 {
		t.Fail()
	}
	if v, ok := UseMap(d1).Get("baz"); ok || v != 0 {
		t.Fail()
	}

	d2 := map[string]int{}
	if v, ok := UseMap(d2).Get("foo"); ok || v != 0 {
		t.Fail()
	}
}

func TestMap_Put(t *testing.T) {
	d1 := map[string]int{"foo": 1}

	UseMap(d1).Put("bar", 20)
	if !UseMap(d1).Same(map[string]int{"foo": 1, "bar": 20}) {
		t.Fail()
	}

	// Functional test
	d2 := map[string]int{"foo": 1}
	Put(d2, "bar", 20)
	if !UseMap(d2).Same(map[string]int{"foo": 1, "bar": 20}) {
		t.Fail()
	}
}

func TestMap_Pull(t *testing.T) {
	d1 := map[string]int{"foo": 1, "bar": 2}
	c1 := UseMap(d1)
	if v, ok := c1.Pull("bar"); !ok || v != 2 {
		t.Fail()
	}
	if v, ok := c1.Pull("bar"); ok || v != 0 {
		t.Fail()
	}
	if !c1.Same(map[string]int{"foo": 1}) {
		t.Fail()
	}
	if !UseMap(d1).Same(map[string]int{"foo": 1}) {
		t.Fail()
	}

	// Functional test
	d2 := map[string]int{"foo": 1, "bar": 2}
	if v, ok := Pull(d2, "bar"); !ok || v != 2 {
		t.Fail()
	}
	if !UseMap(d2).Same(map[string]int{"foo": 1}) {
		t.Fail()
	}
}

func TestMap_Same(t *testing.T) {
	if !UseMap[map[int]int, int, int](nil).Same(nil) {
		t.Fail()
	}

	d1 := map[string]int{"foo": 1, "bar": 0}
	d2 := map[string]int{"foo": 1, "bar": 0}
	if !UseMap(d1).Same(d1) {
		t.Fail()
	}
	if !UseMap(d1).Same(d2) {
		t.Fail()
	}

	if !UseMap(map[bool]struct{}{}).Same(map[bool]struct{}{}) {
		t.Fail()
	}

	d3 := map[string]Foo{"foo": {Bar: "aaa"}, "bar": {Bar: "bbb"}}
	d4 := map[string]Foo{"foo": {Bar: "aaa"}, "bar": {Bar: "bbb"}}
	if !UseMap(d3).Same(d3) {
		t.Fail()
	}
	if !UseMap(d3).Same(d4) {
		t.Fail()
	}

	UseMap(d4).Put("bar", Foo{Bar: "ccc"})
	if UseMap(d3).Same(d4) {
		t.Fail()
	}

	UseMap(d4).Put("bar", Foo{Bar: "bbb"})
	if !UseMap(d3).Same(d4) {
		t.Fail()
	}
}

func TestMap_Merge(t *testing.T) {
	d1 := map[string]int{"a": 1, "b": 2, "c": 3}
	d2 := map[string]int{"c": 33}
	d3 := map[string]int{"c": 333, "d": 444}

	if !UseMap(d1).Merge().Same(d1) {
		t.Fail()
	}
	if !UseMap(d1).Merge(d2).Same(map[string]int{"a": 1, "b": 2, "c": 33}) {
		t.Fail()
	}
	if !UseMap(d1).Merge(d2, d3).Same(map[string]int{"a": 1, "b": 2, "c": 333, "d": 444}) {
		t.Fail()
	}
}

func TestMap_Union(t *testing.T) {
	d1 := map[string]int{"a": 1, "b": 2, "c": 3}
	d2 := map[string]int{"b": 22, "d": 44}

	if !UseMap(d1).Union(d2).Same(map[string]int{"a": 1, "b": 2, "c": 3, "d": 44}) {
		t.Fail()
	}

	f1 := Foo{Bar: "foo1"}
	f2 := Foo{Bar: "foo2"}
	f3 := Foo{Bar: "foo3"}

	d3 := map[Foo]int{f1: 1, f2: 2}
	d4 := map[Foo]int{f1: 11, f3: 33}
	if !UseMap(d3).Union(d4).Same(map[Foo]int{f1: 1, f2: 2, f3: 33}) {
		t.Fail()
	}

	d5 := map[*Foo]int{&f1: 1, &f2: 2}
	d6 := map[*Foo]int{&f1: 11, &f3: 33}
	if !UseMap(d5).Union(d6).Same(map[*Foo]int{&f1: 1, &f2: 2, &f3: 33}) {
		t.Fail()
	}
}
