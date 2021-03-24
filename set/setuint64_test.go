package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetUint64Add(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	expected := []uint64{10, 20, 30, 40}
	mySet := NewUint64(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64Add failed")
		}
	}
}

func TestSetUint64Remove(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	expected := []uint64{10, 20}
	mySet := NewUint64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64Remove failed")
		}
	}
}

func TestSetUint64Clear(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint64Clear failed.")
	}
}

func TestSetUint64Contains(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint64Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint64Contains failed.")
	}
}

func TestSetUint64Size(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint64Size failed.")
	}
}

func TestSetUint64Join(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64(list)

	expected := []uint64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64Join failed")
		}
	}
}

func TestSetUint64Intersection(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64(list)

	expected := []uint64{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint64Minus(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64(list)

	expected := []uint64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint64Subset(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64(list)

	list = []uint64{10, 20}
	mySet2 := NewUint64(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint64Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint64Subset failed. Expected=false, Actual=true")
	}
}

func TestSetUint64Superset(t *testing.T) {
	mySet1 := NewUint64([]uint64{10, 20, 30, 20})
	mySet2 := NewUint64([]uint64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint64Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint64([]uint64{10, 20, 30, 20})
	mySet2 = NewUint64([]uint64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint64([]uint64{10, 20, 30, 20, 40})
	mySet2 = NewUint64([]uint64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64Subset failed. Expected=false, Actual=true")
	}
}
