package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetUint64SyncAdd(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	expected := []uint64{10, 20, 30, 40}
	mySet := NewUint64Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64SyncAdd failed")
		}
	}
}

func TestSetUint64SyncRemove(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	expected := []uint64{10, 20}
	mySet := NewUint64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64SyncRemove failed")
		}
	}
}

func TestSetUint64SyncClear(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint64SyncClear failed.")
	}
}

func TestSetUint64SyncContains(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint64Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint64SyncContains failed.")
	}
}

func TestSetUint64SyncSize(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet := NewUint64Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint64Size failed.")
	}
}

func TestSetUint64SyncJoin(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64Sync(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64Sync(list)

	expected := []uint64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64SyncJoin failed")
		}
	}
}

func TestSetUint64SyncIntersection(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64Sync(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64Sync(list)

	expected := []uint64{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint64SyncMinus(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64Sync(list)

	list = []uint64{30, 40, 50}
	mySet2 := NewUint64Sync(list)

	expected := []uint64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint64(num, expected) {
			t.Errorf("TestSetUint64SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint64SyncSubset(t *testing.T) {
	list := []uint64{10, 20, 30, 20}
	mySet1 := NewUint64Sync(list)

	list = []uint64{10, 20}
	mySet2 := NewUint64Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint64SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint64SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUint64SyncSuperset(t *testing.T) {
	mySet1 := NewUint64Sync([]uint64{10, 20, 30, 20})
	mySet2 := NewUint64Sync([]uint64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint64SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint64Sync([]uint64{10, 20, 30, 20})
	mySet2 = NewUint64Sync([]uint64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint64Sync([]uint64{10, 20, 30, 20, 40})
	mySet2 = NewUint64Sync([]uint64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint64SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetUint64SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewUint64Sync([]uint64{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Uint64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint64(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Uint64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint64(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetUint64SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
