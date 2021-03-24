package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetInt16SyncAdd(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	expected := []int16{10, 20, 30, 40}
	mySet := NewInt16Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16SyncAdd failed")
		}
	}
}

func TestSetInt16SyncRemove(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	expected := []int16{10, 20}
	mySet := NewInt16Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16SyncRemove failed")
		}
	}
}

func TestSetInt16SyncClear(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt16SyncClear failed.")
	}
}

func TestSetInt16SyncContains(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt16Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt16SyncContains failed.")
	}
}

func TestSetInt16SyncSize(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet := NewInt16Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt16Size failed.")
	}
}

func TestSetInt16SyncJoin(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16Sync(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16Sync(list)

	expected := []int16{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16SyncJoin failed")
		}
	}
}

func TestSetInt16SyncIntersection(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16Sync(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16Sync(list)

	expected := []int16{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt16SyncMinus(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16Sync(list)

	list = []int16{30, 40, 50}
	mySet2 := NewInt16Sync(list)

	expected := []int16{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt16(num, expected) {
			t.Errorf("TestSetInt16SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt16SyncSubset(t *testing.T) {
	list := []int16{10, 20, 30, 20}
	mySet1 := NewInt16Sync(list)

	list = []int16{10, 20}
	mySet2 := NewInt16Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt16SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt16SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetInt16SyncSuperset(t *testing.T) {
	mySet1 := NewInt16Sync([]int16{10, 20, 30, 20})
	mySet2 := NewInt16Sync([]int16{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt16SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt16Sync([]int16{10, 20, 30, 20})
	mySet2 = NewInt16Sync([]int16{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt16Sync([]int16{10, 20, 30, 20, 40})
	mySet2 = NewInt16Sync([]int16{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt16SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetInt16SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewInt16Sync([]int16{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Int16Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int16(0); i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Int16Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int16(0); i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 1000 {
		t.Errorf("TestSetInt16SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
