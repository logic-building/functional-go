package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetUintAdd(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	expected := []uint{10, 20, 30, 40}
	mySet := NewUint(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintAdd failed")
		}
	}
}

func TestSetUintRemove(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	expected := []uint{10, 20}
	mySet := NewUint(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintRemove failed")
		}
	}
}

func TestSetUintClear(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUint(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUintClear failed.")
	}
}

func TestSetUintContains(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUint(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUintContains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUintContains failed.")
	}
}

func TestSetUintSize(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUint(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUintSize failed.")
	}
}

func TestSetUintJoin(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUint(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUint(list)

	expected := []uint{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintJoin failed")
		}
	}
}

func TestSetUintIntersection(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUint(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUint(list)

	expected := []uint{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUintMinus(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUint(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUint(list)

	expected := []uint{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUintSubset(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUint(list)

	list = []uint{10, 20}
	mySet2 := NewUint(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUintSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUintSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUintSuperset(t *testing.T) {
	mySet1 := NewUint([]uint{10, 20, 30, 20})
	mySet2 := NewUint([]uint{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUintSubset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSubset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint([]uint{10, 20, 30, 20})
	mySet2 = NewUint([]uint{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSubset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint([]uint{10, 20, 30, 20, 40})
	mySet2 = NewUint([]uint{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSubset failed. Expected=false, Actual=true")
	}
}
