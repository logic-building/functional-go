package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetInt64SyncAdd(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	expected := []int64{10, 20, 30, 40}
	mySet := NewInt64Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64SyncAdd failed")
		}
	}
}

func TestSetInt64SyncRemove(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	expected := []int64{10, 20}
	mySet := NewInt64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64SyncRemove failed")
		}
	}
}

func TestSetInt64SyncClear(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt64SyncClear failed.")
	}
}

func TestSetInt64SyncContains(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt64Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetInt64SyncContains failed.")
	}
}

func TestSetInt64SyncSize(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet := NewInt64Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt64Size failed.")
	}
}

func TestSetInt64SyncJoin(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64Sync(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64Sync(list)

	expected := []int64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64SyncJoin failed")
		}
	}
}

func TestSetInt64SyncIntersection(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64Sync(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64Sync(list)

	expected := []int64{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt64SyncMinus(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64Sync(list)

	list = []int64{30, 40, 50}
	mySet2 := NewInt64Sync(list)

	expected := []int64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt64(num, expected) {
			t.Errorf("TestSetInt64SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt64SyncSubset(t *testing.T) {
	list := []int64{10, 20, 30, 20}
	mySet1 := NewInt64Sync(list)

	list = []int64{10, 20}
	mySet2 := NewInt64Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt64SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt64SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetInt64SyncSuperset(t *testing.T) {
	mySet1 := NewInt64Sync([]int64{10, 20, 30, 20})
	mySet2 := NewInt64Sync([]int64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt64SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt64Sync([]int64{10, 20, 30, 20})
	mySet2 = NewInt64Sync([]int64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt64Sync([]int64{10, 20, 30, 20, 40})
	mySet2 = NewInt64Sync([]int64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt64SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetInt64SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewInt64Sync([]int64{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Int64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int64(0); i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Int64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int64(0); i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 1000 {
		t.Errorf("TestSetInt64SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
