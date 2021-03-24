package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetUint16Add(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	expected := []uint16{10, 20, 30, 40}
	mySet := NewUint16(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16Add failed")
		}
	}
}

func TestSetUint16Remove(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	expected := []uint16{10, 20}
	mySet := NewUint16(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16Remove failed")
		}
	}
}

func TestSetUint16Clear(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint16Clear failed.")
	}
}

func TestSetUint16Contains(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint16Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint16Contains failed.")
	}
}

func TestSetUint16Size(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint16Size failed.")
	}
}

func TestSetUint16Join(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16(list)

	expected := []uint16{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16Join failed")
		}
	}
}

func TestSetUint16Intersection(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16(list)

	expected := []uint16{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint16Minus(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16(list)

	expected := []uint16{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint16Subset(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16(list)

	list = []uint16{10, 20}
	mySet2 := NewUint16(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint16Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint16Subset failed. Expected=false, Actual=true")
	}
}

func TestSetUint16Superset(t *testing.T) {
	mySet1 := NewUint16([]uint16{10, 20, 30, 20})
	mySet2 := NewUint16([]uint16{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint16Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint16([]uint16{10, 20, 30, 20})
	mySet2 = NewUint16([]uint16{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint16([]uint16{10, 20, 30, 20, 40})
	mySet2 = NewUint16([]uint16{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16Subset failed. Expected=false, Actual=true")
	}
}
