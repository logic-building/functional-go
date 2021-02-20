package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetUint8SyncAdd(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	expected := []uint8{10, 20, 30, 40}
	mySet := NewUint8Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8SyncAdd failed")
		}
	}
}

func TestSetUint8SyncRemove(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	expected := []uint8{10, 20}
	mySet := NewUint8Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8SyncRemove failed")
		}
	}
}

func TestSetUint8SyncClear(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint8SyncClear failed.")
	}
}

func TestSetUint8SyncContains(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint8Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint8SyncContains failed.")
	}
}

func TestSetUint8SyncSize(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet := NewUint8Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint8Size failed.")
	}
}

func TestSetUint8SyncJoin(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8Sync(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8Sync(list)

	expected := []uint8{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8SyncJoin failed")
		}
	}
}

func TestSetUint8SyncIntersection(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8Sync(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8Sync(list)

	expected := []uint8{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint8SyncMinus(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8Sync(list)

	list = []uint8{30, 40, 50}
	mySet2 := NewUint8Sync(list)

	expected := []uint8{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint8(num, expected) {
			t.Errorf("TestSetUint8SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint8SyncSubset(t *testing.T) {
	list := []uint8{10, 20, 30, 20}
	mySet1 := NewUint8Sync(list)

	list = []uint8{10, 20}
	mySet2 := NewUint8Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint8SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint8SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUint8SyncSuperset(t *testing.T) {
	mySet1 := NewUint8Sync([]uint8{10, 20, 30, 20})
	mySet2 := NewUint8Sync([]uint8{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint8SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint8Sync([]uint8{10, 20, 30, 20})
	mySet2 = NewUint8Sync([]uint8{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint8Sync([]uint8{10, 20, 30, 20, 40})
	mySet2 = NewUint8Sync([]uint8{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint8SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetUint8SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewUint8Sync([]uint8{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Uint8Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint8(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Uint8Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint8(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetUint8SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
