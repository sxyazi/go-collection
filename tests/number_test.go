package tests

import (
	. "go-collection"
	"testing"
)

func TestNumber_Sum(t *testing.T) {
	d := []float64{0, 1.1, 2.2, 3.3, 4.4, 5.5}
	if Number(d).Sum() != 16.5 {
		t.Fail()
	}
}

func TestNumber_Avg(t *testing.T) {
	d := []float64{0, 1.1, 2.2, 3.3, 4.4, 5.5}
	if Number(d).Avg() != 2.75 {
		t.Fail()
	}
}

func TestNumber_Min(t *testing.T) {
	d := []float64{392, 17, 65, 0, 59, 33, -4}
	if Number(d).Min() != -4 {
		t.Fail()
	}
}

func TestNumber_Max(t *testing.T) {
	d := []float64{392, 17, 65, 0, 59, 33, -4}
	if Number(d).Max() != 392 {
		t.Fail()
	}
}
