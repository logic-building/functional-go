package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetFloat32Add(t *testing.T) {
	list := []float32{10.2, 20, 30, 20}
	expected := []float32{10.2, 20, 30, 40}
	mySet := NewFloat32(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32Add failed")
		}
	}
}

func TestSetFloat32Remove(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	expected := []float32{10, 20}
	mySet := NewFloat32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32Remove failed")
		}
	}
}

func TestSetFloat32Clear(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetFloat32Clear failed.")
	}
}

func TestSetFloat32Contains(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetFloat32Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetFloat32Contains failed.")
	}
}

func TestSetFloat32Size(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetFloat32Size failed.")
	}
}

func TestSetFloat32Join(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32(list)

	list = []float32{30, 40, 50}
	mySet2 := NewFloat32(list)

	expected := []float32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32Join failed")
		}
	}
}

func TestSetFloat32Intersection(t *testing.T) {
	list := []float32{10, 20, 30.3, 20}
	mySet1 := NewFloat32(list)

	list = []float32{30.3, 40, 50}
	mySet2 := NewFloat32(list)

	expected := []float32{30.3}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat32Minus(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32(list)

	list = []float32{30, 40, 50}
	mySet2 := NewFloat32(list)

	expected := []float32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat32Subset(t *testing.T) {
	list := []float32{10.2, 20, 30, 20}
	mySet1 := NewFloat32(list)

	list = []float32{10.2, 20}
	mySet2 := NewFloat32(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetFloat32Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetFloat32Subset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat32Superset(t *testing.T) {
	mySet1 := NewFloat32([]float32{10, 20, 30, 20})
	mySet2 := NewFloat32([]float32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetFloat32Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewFloat32([]float32{10, 20, 30, 20})
	mySet2 = NewFloat32([]float32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewFloat32([]float32{10, 20, 30, 20, 40})
	mySet2 = NewFloat32([]float32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32Subset failed. Expected=false, Actual=true")
	}
}
