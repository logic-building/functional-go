package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetUint32Add(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	expected := []uint32{10, 20, 30, 40}
	mySet := NewUint32(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32Add failed")
		}
	}
}

func TestSetUint32Remove(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	expected := []uint32{10, 20}
	mySet := NewUint32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32Remove failed")
		}
	}
}

func TestSetUint32Clear(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint32Clear failed.")
	}
}

func TestSetUint32Contains(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint32Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint32Contains failed.")
	}
}

func TestSetUint32Size(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint32Size failed.")
	}
}

func TestSetUint32Join(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32(list)

	expected := []uint32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32Join failed")
		}
	}
}

func TestSetUint32Intersection(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32(list)

	expected := []uint32{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint32Minus(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32(list)

	expected := []uint32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint32Subset(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32(list)

	list = []uint32{10, 20}
	mySet2 := NewUint32(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint32Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint32Subset failed. Expected=false, Actual=true")
	}
}

func TestSetUint32Superset(t *testing.T) {
	mySet1 := NewUint32([]uint32{10, 20, 30, 20})
	mySet2 := NewUint32([]uint32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint32Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint32([]uint32{10, 20, 30, 20})
	mySet2 = NewUint32([]uint32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint32([]uint32{10, 20, 30, 20, 40})
	mySet2 = NewUint32([]uint32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32Subset failed. Expected=false, Actual=true")
	}
}
