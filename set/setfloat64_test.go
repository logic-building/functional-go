package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetFloat64Add(t *testing.T) {
	list := []float64{10.2, 20, 30, 20}
	expected := []float64{10.2, 20, 30, 40}
	mySet := NewFloat64(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64Add failed")
		}
	}
}

func TestSetFloat64Remove(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	expected := []float64{10, 20}
	mySet := NewFloat64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64Remove failed")
		}
	}
}

func TestSetFloat64Clear(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetFloat64Clear failed.")
	}
}

func TestSetFloat64Contains(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetFloat64Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetFloat64Contains failed.")
	}
}

func TestSetFloat64Size(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetFloat64Size failed.")
	}
}

func TestSetFloat64Join(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64(list)

	list = []float64{30, 40, 50}
	mySet2 := NewFloat64(list)

	expected := []float64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64Join failed")
		}
	}
}

func TestSetFloat64Intersection(t *testing.T) {
	list := []float64{10, 20, 30.3, 20}
	mySet1 := NewFloat64(list)

	list = []float64{30.3, 40, 50}
	mySet2 := NewFloat64(list)

	expected := []float64{30.3}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat64Minus(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64(list)

	list = []float64{30, 40, 50}
	mySet2 := NewFloat64(list)

	expected := []float64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat64Subset(t *testing.T) {
	list := []float64{10.2, 20, 30, 20}
	mySet1 := NewFloat64(list)

	list = []float64{10.2, 20}
	mySet2 := NewFloat64(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetFloat64Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetFloat64Subset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat64Superset(t *testing.T) {
	mySet1 := NewFloat64([]float64{10, 20, 30, 20})
	mySet2 := NewFloat64([]float64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetFloat64Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewFloat64([]float64{10, 20, 30, 20})
	mySet2 = NewFloat64([]float64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewFloat64([]float64{10, 20, 30, 20, 40})
	mySet2 = NewFloat64([]float64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64Subset failed. Expected=false, Actual=true")
	}
}
