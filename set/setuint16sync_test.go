package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetUint16SyncAdd(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	expected := []uint16{10, 20, 30, 40}
	mySet := NewUint16Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16SyncAdd failed")
		}
	}
}

func TestSetUint16SyncRemove(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	expected := []uint16{10, 20}
	mySet := NewUint16Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16SyncRemove failed")
		}
	}
}

func TestSetUint16SyncClear(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetUint16SyncClear failed.")
	}
}

func TestSetUint16SyncContains(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetUint16Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetUint16SyncContains failed.")
	}
}

func TestSetUint16SyncSize(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet := NewUint16Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetUint16Size failed.")
	}
}

func TestSetUint16SyncJoin(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16Sync(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16Sync(list)

	expected := []uint16{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16SyncJoin failed")
		}
	}
}

func TestSetUint16SyncIntersection(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16Sync(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16Sync(list)

	expected := []uint16{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint16SyncMinus(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16Sync(list)

	list = []uint16{30, 40, 50}
	mySet2 := NewUint16Sync(list)

	expected := []uint16{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsUint16(num, expected) {
			t.Errorf("TestSetUint16SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetUint16SyncSubset(t *testing.T) {
	list := []uint16{10, 20, 30, 20}
	mySet1 := NewUint16Sync(list)

	list = []uint16{10, 20}
	mySet2 := NewUint16Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetUint16SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetUint16SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetUint16SyncSuperset(t *testing.T) {
	mySet1 := NewUint16Sync([]uint16{10, 20, 30, 20})
	mySet2 := NewUint16Sync([]uint16{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetUint16SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewUint16Sync([]uint16{10, 20, 30, 20})
	mySet2 = NewUint16Sync([]uint16{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewUint16Sync([]uint16{10, 20, 30, 20, 40})
	mySet2 = NewUint16Sync([]uint16{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetUint16SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetUint16SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewUint16Sync([]uint16{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Uint16Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint16(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Uint16Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := uint16(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetUint16SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
