package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetUint8Add(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	expected := []uint8{10, 20, 30, 40}
	mySet := NewUint8(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8Add failed")
		}
	}
}

func TestSetUint8Remove(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	expected := []uint8{10, 20}
	mySet := NewUint8(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8Remove failed")
		}
	}
}

func TestSetUint8Clear(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint8Clear failed.")
	}
}

func TestSetUint8Contains(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint8Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint8Contains failed.")
	}
}

func TestSetUint8Size(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint8Size failed.")
	}
}

func TestSetUint8Join(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8(list)

	expected := []uint8{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8Join failed")
		}
	}
}

func TestSetUint8Intersection(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8(list)

	expected := []uint8{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint8Minus(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8(list)

	expected := []uint8{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint8Subset(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8(list)

	list = []uint8{10, 20}
	mySet2 := NewUint8(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint8Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint8Subset failed. Expected=false, Actual=true")
	}
}

func TestSetUint8Superset(t *testing.T) {
	mySet1 := NewUint8([]uint8{10, 20, 30, 20})
	mySet2 := NewUint8([]uint8{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint8Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint8([]uint8{10, 20, 30, 20})
	mySet2 = NewUint8([]uint8{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint8([]uint8{10, 20, 30, 20, 40})
	mySet2 = NewUint8([]uint8{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8Subset failed. Expected=false, Actual=true")
	}
}
