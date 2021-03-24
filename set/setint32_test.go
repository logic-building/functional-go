package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetInt32Add(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	expected := []int32{10, 20, 30, 40}
	mySet := NewInt32(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32Add failed")
		}
	}
}

func TestSetInt32Remove(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	expected := []int32{10, 20}
	mySet := NewInt32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32Remove failed")
		}
	}
}

func TestSetInt32Clear(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt32Clear failed.")
	}
}

func TestSetInt32Contains(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt32Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt32Contains failed.")
	}
}

func TestSetInt32Size(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt32Size failed.")
	}
}

func TestSetInt32Join(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32(list)

	expected := []int32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32Join failed")
		}
	}
}

func TestSetInt32Intersection(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32(list)

	expected := []int32{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt32Minus(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32(list)

	expected := []int32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt32Subset(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32(list)

	list = []int32{10, 20}
	mySet2 := NewInt32(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt32Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt32Subset failed. Expected=false, Actual=true")
	}
}

func TestSetInt32Superset(t *testing.T) {
	mySet1 := NewInt32([]int32{10, 20, 30, 20})
	mySet2 := NewInt32([]int32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt32Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt32([]int32{10, 20, 30, 20})
	mySet2 = NewInt32([]int32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt32([]int32{10, 20, 30, 20, 40})
	mySet2 = NewInt32([]int32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32Subset failed. Expected=false, Actual=true")
	}
}
