package set

import (
	"github.com/logic-building/functional-go/fp"
	"testing"
)

func TestSetStrAdd(t *testing.T) {
	list := []string{"10.2", "20", "30", "20"}
	expected := []string{"10.2", "20", "30", "40"}
	mySet := NewStr(list)
	mySet.Add("40")
	mySet.Add("40")

	for _, num := range mySet.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrAdd failed")
		}
	}
}

func TestSetStrRemove(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	expected := []string{"10", "20"}
	mySet := NewStr(list)
	mySet.Add("40")
	mySet.Add("40")
	mySet.Remove("30")
	mySet.Remove("40")

	for _, num := range mySet.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrRemove failed")
		}
	}
}

func TestSetStrClear(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStr(list)
	mySet.Add("40")
	mySet.Add("40")
	mySet.Remove("30")
	mySet.Remove("40")
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetStrClear failed.")
	}
}

func TestSetStrContains(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStr(list)

	if !mySet.Contains("20") {
		t.Errorf("TestSetStrContains failed.")
	}

	if mySet.Contains("200") {
		t.Errorf("TestSetStrContains failed.")
	}
}

func TestSetStrSize(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStr(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetStrSize failed.")
	}
}

func TestSetStrJoin(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStr(list)

	list = []string{"30", "40", "50"}
	mySet2 := NewStr(list)

	expected := []string{"10", "20", "30", "40", "50"}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrJoin failed")
		}
	}
}

func TestSetStrIntersection(t *testing.T) {
	list := []string{"10", "20", "30.3", "20"}
	mySet1 := NewStr(list)

	list = []string{"30.3", "40", "50"}
	mySet2 := NewStr(list)

	expected := []string{"30.3"}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetStrMinus(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStr(list)

	list = []string{"30", "40", "50"}
	mySet2 := NewStr(list)

	expected := []string{"10", "20"}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetStrSubset(t *testing.T) {
	list := []string{"10.2", "20", "30", "20"}
	mySet1 := NewStr(list)

	list = []string{"10.2", "20"}
	mySet2 := NewStr(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetStrSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetStrSubset failed. Expected=false, Actual=true")
	}
}

func TestSetStrSuperset(t *testing.T) {
	mySet1 := NewStr([]string{"10", "20", "30", "20"})
	mySet2 := NewStr([]string{"10", "20"})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetStrSubset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSubset failed. Expected=false, Actual=true")
	}

	mySet1 = NewStr([]string{"10", "20", "30", "20"})
	mySet2 = NewStr([]string{"10", "20", "30", "20"})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSubset failed. Expected=true, Actual=false")
	}

	mySet1 = NewStr([]string{"10", "20", "30", "20", "40"})
	mySet2 = NewStr([]string{"10", "20", "30", "20"})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSubset failed. Expected=false, Actual=true")
	}
}
