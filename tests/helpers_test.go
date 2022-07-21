package tests

import (
	"encoding/json"
	. "github.com/sxyazi/go-collection"
	"testing"
)

func TestHelpers_AnyGet(t *testing.T) {
	// Nil
	if _, err := AnyGet[any](nil, ""); err == nil {
		t.Fail()
	}

	// Struct
	user := User{ID: 33, Name: "Lucy"}
	if v, err := AnyGet[string](user, "Name"); err != nil || v != "Lucy" {
		t.Fail()
	}
	if v, err := AnyGet[string](&user, "Name"); err != nil || v != "Lucy" {
		t.Fail()
	}
	if v, err := AnyGet[any](&user, "Name"); err != nil || v.(string) != "Lucy" {
		t.Fail()
	}

	// Slice
	users := []*User{&user}
	if _, err := AnyGet[any](users, 0); err != nil {
		t.Fail()
	}
	if v, err := AnyGet[*User](users, 0); err != nil || v != &user {
		t.Fail()
	}
	if v, err := AnyGet[*User](users, "0"); err != nil || v != &user {
		t.Fail()
	}
	if _, err := AnyGet[*User](users, 10); err == nil {
		t.Fail()
	}

	// Array
	if v, err := AnyGet[int]([]int{1, 2, 3}, 2); err != nil || v != 3 {
		t.Fail()
	}
	if v, err := AnyGet[int]([3]int{1, 2, 3}, 2); err != nil || v != 3 {
		t.Fail()
	}
	if v, err := AnyGet[int]([3]int{1, 2, 3}, "2"); err != nil || v != 3 {
		t.Fail()
	}
	if v, err := AnyGet[any]([3]int{1, 2, 3}, 2); err != nil || v.(int) != 3 {
		t.Fail()
	}

	// Interface
	var i any
	if _, err := AnyGet[any](i, ""); err == nil {
		t.Fail()
	}

	i = make(map[int]string)
	i.(map[int]string)[0] = "Hello"
	if _, err := AnyGet[string](i, 1); err == nil {
		t.Fail()
	}
	if v, err := AnyGet[string](i, 0); err != nil || v != "Hello" {
		t.Fail()
	}
	if v, err := AnyGet[any](i, 0); err != nil || v.(string) != "Hello" {
		t.Fail()
	}

	json.Unmarshal([]byte(`["World"]`), &i)
	if _, err := AnyGet[string](i, 1); err == nil {
		t.Fail()
	}
	if v, err := AnyGet[string](i, 0); err != nil || v != "World" {
		t.Fail()
	}
	if v, err := AnyGet[any](i, 0); err != nil || v.(string) != "World" {
		t.Fail()
	}
}

func TestHelpers_Pluck(t *testing.T) {
	ids := []uint{33, 193}
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}

	if !UseSlice(Pluck[uint](users, "ID")).Same(ids) {
		t.Fail()
	}
}

func TestHelpers_MapPluck(t *testing.T) {
	ids := []uint{33, 193}
	s := []map[string]uint{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}

	if !UseSlice(MapPluck(s, "ID")).Same(ids) {
		t.Fail()
	}
}

func TestHelpers_KeyBy(t *testing.T) {
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
	r := KeyBy[string](users, "Name")
	if len(r) != 2 {
		t.Fail()
	}
	if r["Lucy"].ID != 33 || r["Peter"].ID != 194 {
		t.Fail()
	}
}

func TestHelpers_MapKeyBy(t *testing.T) {
	m := []map[string]int{{"ID": 33, "Age": 40}, {"ID": 193, "Age": 25}, {"ID": 194, "Age": 25}}
	r := MapKeyBy(m, "Age")
	if len(r) != 2 {
		t.Fail()
	}
	if r[40]["ID"] != 33 || r[25]["ID"] != 194 {
		t.Fail()
	}
}

func TestHelpers_GroupBy(t *testing.T) {
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Lacie"}}
	r2 := GroupBy[uint](users, "ID")
	if len(r2) != 3 || len(r2[33]) != 1 || len(r2[193]) != 1 || len(r2[194]) != 1 {
		t.Fail()
	}
	if r2[33][0].Name != "Lucy" || r2[193][0].Name != "Peter" || r2[194][0].Name != "Lacie" {
		t.Fail()
	}
}

func TestHelpers_MapGroupBy(t *testing.T) {
	m := []map[string]int{{"ID": 33, "Age": 40}, {"ID": 193, "Age": 25}, {"ID": 194, "Age": 25}}
	r := MapGroupBy(m, "Age")
	if len(r) != 2 || len(r[40]) != 1 || len(r[25]) != 2 {
		t.Fail()
	}
	if r[40][0]["ID"] != 33 || r[25][0]["ID"] != 193 || r[25][1]["ID"] != 194 {
		t.Fail()
	}
}

func TestHelper_IsNumber(t *testing.T) {
	var d1 int = 123
	if !IsNumber(d1) {
		t.Fail()
	}

	var d2 int32 = 123
	if !IsNumber(d2) {
		t.Fail()
	}

	var d3 float32 = 3.14
	if !IsNumber(d3) {
		t.Fail()
	}

	var d4 float64 = 2.71
	if !IsNumber(d4) {
		t.Fail()
	}

	d5 := true
	if IsNumber(d5) {
		t.Fail()
	}

	d6 := []int{1, 2, 3}
	if IsNumber(d6) {
		t.Fail()
	}
}

func TestHelper_NumberCompare(t *testing.T) {
	if !NumberCompare(10.3, "=", 10.3) {
		t.Fail()
	}
	if NumberCompare(10.3, "!=", 10.3) {
		t.Fail()
	}
	if !NumberCompare(10.3, ">", 4.7) {
		t.Fail()
	}
	if !NumberCompare(10.3, "<", 20.5) {
		t.Fail()
	}
	if !NumberCompare(10.3, ">=", 10.3) {
		t.Fail()
	}
	if !NumberCompare(10.3, "<=", 10.3) {
		t.Fail()
	}
}

func TestHelper_AnyNumberCompare(t *testing.T) {
	if !AnyNumberCompare(10.3, "=", 10.3) {
		t.Fail()
	}
	if AnyNumberCompare(10.3, "=", 10) {
		t.Fail()
	}
	if AnyNumberCompare(10, "=", 10.3) {
		t.Fail()
	}
	if !AnyNumberCompare(10.0, "!=", 10) {
		t.Fail()
	}
	if !AnyNumberCompare(10, "!=", 10.0) {
		t.Fail()
	}
	if AnyNumberCompare(10.1, "=", 10.0) {
		t.Fail()
	}

	if !AnyNumberCompare(10.3, ">", 4.7) {
		t.Fail()
	}
	if !AnyNumberCompare(10.3, "<", 20.5) {
		t.Fail()
	}
	if !AnyNumberCompare(10.3, ">=", 10.3) {
		t.Fail()
	}
	if !AnyNumberCompare(10.3, "<=", 10.3) {
		t.Fail()
	}

}

func TestHelpers_Compare(t *testing.T) {
	// Nil
	if !Compare(nil, "=", nil) {
		t.Fail()
	}
	if Compare(nil, "!=", nil) {
		t.Fail()
	}
	if Compare(nil, ">", nil) {
		t.Fail()
	}

	// Slice
	d1 := []int{1, 2, 3}
	if Compare(d1, "!=", d1) {
		t.Fail()
	}
	if Compare(d1, "=", []int{4, 5, 6}) {
		t.Fail()
	}

	// Array
	d2 := [...]int{1, 2, 3}
	if Compare(d2, "!=", d2) {
		t.Fail()
	}
	if Compare(d2, "=", []int{4, 5, 6}) {
		t.Fail()
	}
	if Compare(d2, "=", [...]int{4, 5, 6}) {
		t.Fail()
	}

	// Channel
	d3 := make(chan int)
	if Compare(d3, "!=", d3) {
		t.Fail()
	}
	if Compare(d3, "=", false) {
		t.Fail()
	}
	if Compare(d3, "=", make(chan int)) {
		t.Fail()
	}

	// Function
	d4 := func() {}
	if Compare(d4, "!=", d4) {
		t.Fail()
	}
	if Compare(d4, "=", false) {
		t.Fail()
	}
	if Compare(d4, "=", func() {}) {
		t.Fail()
	}

	// Interface
	d5 := interface{}(3.14)
	if Compare(d5, "!=", d5) {
		t.Fail()
	}
	if Compare(d5, "=", false) {
		t.Fail()
	}
	if Compare(d5, "!=", 3.14) {
		t.Fail()
	}

	// Map
	d6 := map[string]string{"name": "Lucy"}
	if Compare(d6, "!=", d6) {
		t.Fail()
	}
	if Compare(d6, "=", nil) {
		t.Fail()
	}
	if !Compare(d6, "!=", nil) {
		t.Fail()
	}

	// Struct
	d7 := struct{ Name string }{"Lucy"}
	if Compare(d7, "!=", d7) {
		t.Fail()
	}
	if Compare(d7, "=", nil) {
		t.Fail()
	}
	if !Compare(d7, "!=", nil) {
		t.Fail()
	}

	// Pointer
	d8 := &Foo{Bar: "abc"}
	if Compare(d8, "!=", d8) {
		t.Fail()
	}
	if Compare(d8, "=", &d8) {
		t.Fail()
	}
	if !Compare(d8, "!=", true) {
		t.Fail()
	}
	if Compare(d8, "=", nil) {
		t.Fail()
	}
	if !Compare(d8, "!=", nil) {
		t.Fail()
	}
}
