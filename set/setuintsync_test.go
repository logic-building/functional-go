package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetUintSyncAdd(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	expected := []uint{10, 20, 30, 40}
	mySet := NewUintSync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintSyncAdd failed")
		}
	}
}

func TestSetUintSyncRemove(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	expected := []uint{10, 20}
	mySet := NewUintSync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintSyncRemove failed")
		}
	}
}

func TestSetUintSyncClear(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUintSync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUintSyncClear failed.")
	}
}

func TestSetUintSyncContains(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUintSync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUintContains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUintSyncContains failed.")
	}
}

func TestSetUintSyncSize(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet := NewUintSync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUintSize failed.")
	}
}

func TestSetUintSyncJoin(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUintSync(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUintSync(list)

	expected := []uint{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintSyncJoin failed")
		}
	}
}

func TestSetUintSyncIntersection(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUintSync(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUintSync(list)

	expected := []uint{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintSyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUintSyncMinus(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUintSync(list)

	list = []uint{30, 40, 50}
	mySet2 := NewUintSync(list)

	expected := []uint{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint(num, expected) {
			t.Errorf("TestSetUintSyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUintSyncSubset(t *testing.T) {
	list := []uint{10, 20, 30, 20}
	mySet1 := NewUintSync(list)

	list = []uint{10, 20}
	mySet2 := NewUintSync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUintSyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUintSyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUintSyncSuperset(t *testing.T) {
	mySet1 := NewUintSync([]uint{10, 20, 30, 20})
	mySet2 := NewUintSync([]uint{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUintSyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUintSync([]uint{10, 20, 30, 20})
	mySet2 = NewUintSync([]uint{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUintSync([]uint{10, 20, 30, 20, 40})
	mySet2 = NewUintSync([]uint{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUintSyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetUintSyncMultipleGoRoutine(t *testing.T) {
	mySet := NewUintSync([]uint{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *UintSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *UintSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetUintSyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
