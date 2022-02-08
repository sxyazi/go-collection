package tests

import (
	"encoding/json"
	"fmt"
	. "github.com/sxyazi/go-collection"
	"testing"
)

func TestGet(t *testing.T) {
	// Struct
	user := User{ID: 33, Name: "Lucy"}
	if v, err := AnyGet[string](user, "Name"); err != nil || v != "Lucy" {
		t.Fail()
	}
	if v, err := AnyGet[string](&user, "Name"); err != nil || v != "Lucy" {
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

	json.Unmarshal([]byte(`["World"]`), &i)
	if _, err := AnyGet[string](i, 1); err == nil {
		t.Fail()
	}
	if v, err := AnyGet[string](i, 0); err != nil || v != "World" {
		t.Fail()
	}
}

func TestPluck(t *testing.T) {
	ids := []uint{33, 193}
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}

	if !UseSlice(Pluck[uint](users, "ID")).Same(ids) {
		t.Fail()
	}
}

func TestMapPluck(t *testing.T) {
	ids := []uint{33, 193}
	s := []map[string]uint{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}

	if !UseSlice(MapPluck(s, "ID")).Same(ids) {
		t.Fail()
	}
}

func TestKeyBy(t *testing.T) {
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
	r := KeyBy[string](users, "Name")
	fmt.Println(r)
	if len(r) != 2 {
		t.Fail()
	}
	if r["Lucy"].ID != 33 || r["Peter"].ID != 194 {
		t.Fail()
	}
}

func TestMapKeyBy(t *testing.T) {
	m := []map[string]int{{"ID": 33, "Age": 40}, {"ID": 193, "Age": 25}, {"ID": 194, "Age": 25}}
	r := MapKeyBy(m, "Age")
	if len(r) != 2 {
		t.Fail()
	}
	if r[40]["ID"] != 33 || r[25]["ID"] != 194 {
		t.Fail()
	}
}

func TestGroupBy(t *testing.T) {
	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Lacie"}}
	r2 := GroupBy[uint](users, "ID")
	if len(r2) != 3 || len(r2[33]) != 1 || len(r2[193]) != 1 || len(r2[194]) != 1 {
		t.Fail()
	}
	if r2[33][0].Name != "Lucy" || r2[193][0].Name != "Peter" || r2[194][0].Name != "Lacie" {
		t.Fail()
	}
}

func TestMapGroupBy(t *testing.T) {
	m := []map[string]int{{"ID": 33, "Age": 40}, {"ID": 193, "Age": 25}, {"ID": 194, "Age": 25}}
	r := MapGroupBy(m, "Age")
	if len(r) != 2 || len(r[40]) != 1 || len(r[25]) != 2 {
		t.Fail()
	}
	if r[40][0]["ID"] != 33 || r[25][0]["ID"] != 193 || r[25][1]["ID"] != 194 {
		t.Fail()
	}
}
