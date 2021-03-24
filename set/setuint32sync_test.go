package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetUint32SyncAdd(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	expected := []uint32{10, 20, 30, 40}
	mySet := NewUint32Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32SyncAdd failed")
		}
	}
}

func TestSetUint32SyncRemove(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	expected := []uint32{10, 20}
	mySet := NewUint32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32SyncRemove failed")
		}
	}
}

func TestSetUint32SyncClear(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint32SyncClear failed.")
	}
}

func TestSetUint32SyncContains(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint32Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint32SyncContains failed.")
	}
}

func TestSetUint32SyncSize(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet := NewUint32Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint32Size failed.")
	}
}

func TestSetUint32SyncJoin(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32Sync(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32Sync(list)

	expected := []uint32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32SyncJoin failed")
		}
	}
}

func TestSetUint32SyncIntersection(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32Sync(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32Sync(list)

	expected := []uint32{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint32SyncMinus(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32Sync(list)

	list = []uint32{30, 40, 50}
	mySet2 := NewUint32Sync(list)

	expected := []uint32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint32(num, expected) {
			t.Errorf("TestSetUint32SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint32SyncSubset(t *testing.T) {
	list := []uint32{10, 20, 30, 20}
	mySet1 := NewUint32Sync(list)

	list = []uint32{10, 20}
	mySet2 := NewUint32Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint32SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint32SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUint32SyncSuperset(t *testing.T) {
	mySet1 := NewUint32Sync([]uint32{10, 20, 30, 20})
	mySet2 := NewUint32Sync([]uint32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint32SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint32Sync([]uint32{10, 20, 30, 20})
	mySet2 = NewUint32Sync([]uint32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint32Sync([]uint32{10, 20, 30, 20, 40})
	mySet2 = NewUint32Sync([]uint32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint32SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetUint32SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewUint32Sync([]uint32{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Uint32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint32(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Uint32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint32(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetUint32SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
