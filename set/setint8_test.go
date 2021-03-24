package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetInt8Add(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	expected := []int8{10, 20, 30, 40}
	mySet := NewInt8(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8Add failed")
		}
	}
}

func TestSetInt8Remove(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	expected := []int8{10, 20}
	mySet := NewInt8(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8Remove failed")
		}
	}
}

func TestSetInt8Clear(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt8Clear failed.")
	}
}

func TestSetInt8Contains(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt8Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetInt8Contains failed.")
	}
}

func TestSetInt8Size(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt8Size failed.")
	}
}

func TestSetInt8Join(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8(list)

	expected := []int8{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8Join failed")
		}
	}
}

func TestSetInt8Intersection(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8(list)

	expected := []int8{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt8Minus(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8(list)

	expected := []int8{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt8Subset(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8(list)

	list = []int8{10, 20}
	mySet2 := NewInt8(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt8Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt8Subset failed. Expected=false, Actual=true")
	}
}

func TestSetInt8Superset(t *testing.T) {
	mySet1 := NewInt8([]int8{10, 20, 30, 20})
	mySet2 := NewInt8([]int8{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt8Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt8([]int8{10, 20, 30, 20})
	mySet2 = NewInt8([]int8{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt8([]int8{10, 20, 30, 20, 40})
	mySet2 = NewInt8([]int8{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8Subset failed. Expected=false, Actual=true")
	}
}
