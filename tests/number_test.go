package tests

import (
	. "github.com/sxyazi/go-collection"
	"testing"
)

func TestNumber_Sum(t *testing.T) {
	d := []float64{0, 1.1, 2.2, 3.3, 4.4, 5.5}
	if UseNumber(d).Sum() != 16.5 {
		t.Fail()
	}
}

func TestNumber_Min(t *testing.T) {
	d := []float64{392, 17, 65, 0, 59, 33, -4}
	if UseNumber(d).Min() != -4 {
		t.Fail()
	}
}

func TestNumber_Max(t *testing.T) {
	d := []float64{392, 17, 65, 0, 59, 33, -4}
	if UseNumber(d).Max() != 392 {
		t.Fail()
	}
}

func TestNumber_Sort(t *testing.T) {
	d1 := []float64{0, 17.5, -4.01, 0.2, 59, 33, -4}
	if !UseNumber(d1).Sort().Same([]float64{-4.01, -4, 0, 0.2, 17.5, 33, 59}) {
		t.Fail()
	}

	d2 := []int{392, 17, 65, 0, 59, 33, -4}
	if !UseNumber(d2).Sort().Same([]int{-4, 0, 17, 33, 59, 65, 392}) {
		t.Fail()
	}
}

func TestNumber_Avg(t *testing.T) {
	d := []float64{0, 1.1, 2.2, 3.3, 4.4, 5.5}
	if UseNumber(d).Avg() != 2.75 {
		t.Fail()
	}
}

func TestNumber_Median(t *testing.T) {
	if UseNumber([]int{1, 2, 3}).Median() != 2 {
		t.Fail()
	}
	if UseNumber([]int{1, 2, 3, 4}).Median() != 2.5 {
		t.Fail()
	}

	d1 := []float64{392, 17, 65.2, 0, 59, 33.33, -4}
	if UseNumber(d1).Median() != 33.33 {
		t.Fail()
	}

	d2 := []float64{392, 17, 65.2, 0, 33.33, -4}
	if UseNumber(d2).Median() != 25.165 {
		t.Fail()
	}
}
