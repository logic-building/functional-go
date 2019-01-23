package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetInt16Add(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	expected := []int16{10, 20, 30, 40}
	mySet := NewInt16(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16Add failed")
		}
	}
}

func TestSetInt16Remove(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	expected := []int16{10, 20}
	mySet := NewInt16(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16Remove failed")
		}
	}
}

func TestSetInt16Clear(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt16Clear failed.")
	}
}

func TestSetInt16Contains(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt16Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt16Contains failed.")
	}
}

func TestSetInt16Size(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt16Size failed.")
	}
}

func TestSetInt16Join(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16(list)

	expected := []int16{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16Join failed")
		}
	}
}

func TestSetInt16Intersection(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16(list)

	expected := []int16{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt16Minus(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16(list)

	expected := []int16{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt16Subset(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16(list)

	list = []int16{10, 20}
	mySet2 := NewInt16(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt16Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt16Subset failed. Expected=false, Actual=true")
	}
}

func TestSetInt16Superset(t *testing.T) {
	mySet1 := NewInt16([]int16{10, 20, 30, 20})
	mySet2 := NewInt16([]int16{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt16Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt16([]int16{10, 20, 30, 20})
	mySet2 = NewInt16([]int16{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt16([]int16{10, 20, 30, 20, 40})
	mySet2 = NewInt16([]int16{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16Subset failed. Expected=false, Actual=true")
	}
}
