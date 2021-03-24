package set

import (
	"github.com/logic-building/functional-go/fp"
	"strconv"
	"sync"
	"testing"
)

func TestSetStrSyncAdd(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	expected := []string{"10", "20", "30", "40"}
	mySet := NewStrSync(list)
	mySet.Add("40")
	mySet.Add("40")

	for _, num := range mySet.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrSyncAdd failed")
		}
	}
}

func TestSetStrSyncRemove(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	expected := []string{"10", "20"}
	mySet := NewStrSync(list)
	mySet.Add("40")
	mySet.Add("40")
	mySet.Remove("30")
	mySet.Remove("40")

	for _, num := range mySet.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrSyncRemove failed")
		}
	}
}

func TestSetStrSyncClear(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStrSync(list)
	mySet.Add("40")
	mySet.Add("40")
	mySet.Remove("30")
	mySet.Remove("40")
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetStrSyncClear failed.")
	}
}

func TestSetStrSyncContains(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStrSync(list)

	if !mySet.Contains("20") {
		t.Errorf("TestSetStrContains failed.")
	}

	if mySet.Contains("200") {
		t.Errorf("TestSetStrSyncContains failed.")
	}
}

func TestSetStrSyncSize(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet := NewStrSync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetStrSize failed.")
	}
}

func TestSetStrSyncJoin(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStrSync(list)

	list = []string{"30", "40", "50"}
	mySet2 := NewStrSync(list)

	expected := []string{"10", "20", "30", "40", "50"}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrSyncJoin failed")
		}
	}
}

func TestSetStrSyncIntersection(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStrSync(list)

	list = []string{"30", "40", "50"}
	mySet2 := NewStrSync(list)

	expected := []string{"30"}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrSyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetStrSyncMinus(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStrSync(list)

	list = []string{"30", "40", "50"}
	mySet2 := NewStrSync(list)

	expected := []string{"10", "20"}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsStr(num, expected) {
			t.Errorf("TestSetStrSyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetStrSyncSubset(t *testing.T) {
	list := []string{"10", "20", "30", "20"}
	mySet1 := NewStrSync(list)

	list = []string{"10", "20"}
	mySet2 := NewStrSync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetStrSyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetStrSyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetStrSyncSuperset(t *testing.T) {
	mySet1 := NewStrSync([]string{"10", "20", "30", "20"})
	mySet2 := NewStrSync([]string{"10", "20"})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetStrSyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewStrSync([]string{"10", "20", "30", "20"})
	mySet2 = NewStrSync([]string{"10", "20", "30", "20"})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewStrSync([]string{"10", "20", "30", "20", "40"})
	mySet2 = NewStrSync([]string{"10", "20", "30", "20"})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetStrSyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetStrSyncMultipleGoRoutine(t *testing.T) {
	mySet := NewStrSync([]string{"10"})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *StrSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mySet.Add(strconv.Itoa(i))
		}
	}(mySet, &wg)

	go func(mySet *StrSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			for {
				if mySet.Remove(strconv.Itoa(i)) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != "10" {
		t.Errorf("TestSetStrSyncMultipleGoRoutine failed. Expetected [\"10\"], Actual=%v", mySet.GetList())
	}

}
