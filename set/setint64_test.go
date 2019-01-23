package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetInt64Add(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	expected := []int64{10, 20, 30, 40}
	mySet := NewInt64(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64Add failed")
		}
	}
}

func TestSetInt64Remove(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	expected := []int64{10, 20}
	mySet := NewInt64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64Remove failed")
		}
	}
}

func TestSetInt64Clear(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt64Clear failed.")
	}
}

func TestSetInt64Contains(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt64Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt64Contains failed.")
	}
}

func TestSetInt64Size(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt64Size failed.")
	}
}

func TestSetInt64Join(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64(list)

	expected := []int64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64Join failed")
		}
	}
}

func TestSetInt64Intersection(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64(list)

	expected := []int64{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64Intersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt64Minus(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64(list)

	expected := []int64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64Minus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt64Subset(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64(list)

	list = []int64{10, 20}
	mySet2 := NewInt64(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt64Subset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt64Subset failed. Expected=false, Actual=true")
	}
}

func TestSetInt64Superset(t *testing.T) {
	mySet1 := NewInt64([]int64{10, 20, 30, 20})
	mySet2 := NewInt64([]int64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt64Subset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64Subset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt64([]int64{10, 20, 30, 20})
	mySet2 = NewInt64([]int64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64Subset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt64([]int64{10, 20, 30, 20, 40})
	mySet2 = NewInt64([]int64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64Subset failed. Expected=false, Actual=true")
	}
}
