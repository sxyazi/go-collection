package tests

import (
	"encoding/json"
	collect "go-collection"
	"testing"
)

func TestGet(t *testing.T) {
	// Struct
	user := User{ID: 33, Name: "Lucy"}
	if v, err := collect.Get[string](user, "Name"); err != nil || v != "Lucy" {
		t.Fail()
	}
	if v, err := collect.Get[string](&user, "Name"); err != nil || v != "Lucy" {
		t.Fail()
	}

	// Slice
	users := []*User{&user}
	if _, err := collect.Get[interface{}](users, 0); err != nil {
		t.Fail()
	}
	if v, err := collect.Get[*User](users, 0); err != nil || v != &user {
		t.Fail()
	}
	if _, err := collect.Get[*User](users, 10); err == nil {
		t.Fail()
	}

	// Array
	if v, err := collect.Get[int]([]int{1, 2, 3}, 2); err != nil || v != 3 {
		t.Fail()
	}
	if v, err := collect.Get[int]([3]int{1, 2, 3}, 2); err != nil || v != 3 {
		t.Fail()
	}

	// Interface
	var i interface{}
	if _, err := collect.Get[interface{}](i, ""); err == nil {
		t.Fail()
	}

	i = make(map[int]string)
	i.(map[int]string)[0] = "Hello"
	if _, err := collect.Get[string](i, 1); err == nil {
		t.Fail()
	}
	if v, err := collect.Get[string](i, 0); err != nil || v != "Hello" {
		t.Fail()
	}

	json.Unmarshal([]byte(`["World"]`), &i)
	if _, err := collect.Get[string](i, 1); err == nil {
		t.Fail()
	}
	if v, err := collect.Get[string](i, 0); err != nil || v != "World" {
		t.Fail()
	}
}

func TestPluck(t *testing.T) {
	ids := []uint{33, 193}

	s := []map[string]uint{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}
	if !collect.Any(collect.Pluck(s, "ID")).Same(ids) {
		t.Fail()
	}

	users := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}
	if !collect.Any(collect.PluckAny[uint](users, "ID")).Same(ids) {
		t.Fail()
	}
}
