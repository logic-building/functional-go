package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetIntAdd(t *testing.T) {
	list := []int{10, 20, 30, 20}
	expected := []int{10, 20, 30, 40}
	mySet := NewInt(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntAdd failed")
		}
	}
}

func TestSetIntRemove(t *testing.T) {
	list := []int{10, 20, 30, 20}
	expected := []int{10, 20}
	mySet := NewInt(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntRemove failed")
		}
	}
}

func TestSetIntClear(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewInt(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetIntClear failed.")
	}
}

func TestSetIntContains(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewInt(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetIntContains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetIntContains failed.")
	}
}

func TestSetIntSize(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewInt(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetIntSize failed.")
	}
}

func TestSetIntJoin(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewInt(list)

	list = []int{30, 40, 50}
	mySet2 := NewInt(list)

	expected := []int{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntJoin failed")
		}
	}
}

func TestSetIntIntersection(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewInt(list)

	list = []int{30, 40, 50}
	mySet2 := NewInt(list)

	expected := []int{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetIntMinus(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewInt(list)

	list = []int{30, 40, 50}
	mySet2 := NewInt(list)

	expected := []int{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetIntSubset(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewInt(list)

	list = []int{10, 20}
	mySet2 := NewInt(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetIntSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetIntSubset failed. Expected=false, Actual=true")
	}
}

func TestSetIntSuperset(t *testing.T) {
	mySet1 := NewInt([]int{10, 20, 30, 20})
	mySet2 := NewInt([]int{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetIntSubset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSubset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt([]int{10, 20, 30, 20})
	mySet2 = NewInt([]int{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSubset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt([]int{10, 20, 30, 20, 40})
	mySet2 = NewInt([]int{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSubset failed. Expected=false, Actual=true")
	}
}
