package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetInt32SyncAdd(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	expected := []int32{10, 20, 30, 40}
	mySet := NewInt32Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32SyncAdd failed")
		}
	}
}

func TestSetInt32SyncRemove(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	expected := []int32{10, 20}
	mySet := NewInt32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32SyncRemove failed")
		}
	}
}

func TestSetInt32SyncClear(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt32SyncClear failed.")
	}
}

func TestSetInt32SyncContains(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt32Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt32SyncContains failed.")
	}
}

func TestSetInt32SyncSize(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet := NewInt32Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt32Size failed.")
	}
}

func TestSetInt32SyncJoin(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32Sync(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32Sync(list)

	expected := []int32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32SyncJoin failed")
		}
	}
}

func TestSetInt32SyncIntersection(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32Sync(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32Sync(list)

	expected := []int32{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt32SyncMinus(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32Sync(list)

	list = []int32{30, 40, 50}
	mySet2 := NewInt32Sync(list)

	expected := []int32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt32(num, expected) {
			t.Errorf("TestSetInt32SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt32SyncSubset(t *testing.T) {
	list := []int32{10, 20, 30, 20}
	mySet1 := NewInt32Sync(list)

	list = []int32{10, 20}
	mySet2 := NewInt32Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt32SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt32SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetInt32SyncSuperset(t *testing.T) {
	mySet1 := NewInt32Sync([]int32{10, 20, 30, 20})
	mySet2 := NewInt32Sync([]int32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt32SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt32Sync([]int32{10, 20, 30, 20})
	mySet2 = NewInt32Sync([]int32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt32Sync([]int32{10, 20, 30, 20, 40})
	mySet2 = NewInt32Sync([]int32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt32SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetInt32SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewInt32Sync([]int32{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Int32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int32(0); i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Int32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int32(0); i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 1000 {
		t.Errorf("TestSetInt32SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
