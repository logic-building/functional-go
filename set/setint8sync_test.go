package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetInt8SyncAdd(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	expected := []int8{10, 20, 30, 40}
	mySet := NewInt8Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8SyncAdd failed")
		}
	}
}

func TestSetInt8SyncRemove(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	expected := []int8{10, 20}
	mySet := NewInt8Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8SyncRemove failed")
		}
	}
}

func TestSetInt8SyncClear(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetInt8SyncClear failed.")
	}
}

func TestSetInt8SyncContains(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetInt8Contains failed.")
	}

	if mySet.Contains(100) {
		t.Errorf("TestSetInt8SyncContains failed.")
	}
}

func TestSetInt8SyncSize(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet := NewInt8Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetInt8Size failed.")
	}
}

func TestSetInt8SyncJoin(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8Sync(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8Sync(list)

	expected := []int8{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8SyncJoin failed")
		}
	}
}

func TestSetInt8SyncIntersection(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8Sync(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8Sync(list)

	expected := []int8{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt8SyncMinus(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8Sync(list)

	list = []int8{30, 40, 50}
	mySet2 := NewInt8Sync(list)

	expected := []int8{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsInt8(num, expected) {
			t.Errorf("TestSetInt8SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetInt8SyncSubset(t *testing.T) {
	list := []int8{10, 20, 30, 20}
	mySet1 := NewInt8Sync(list)

	list = []int8{10, 20}
	mySet2 := NewInt8Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetInt8SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetInt8SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetInt8SyncSuperset(t *testing.T) {
	mySet1 := NewInt8Sync([]int8{10, 20, 30, 20})
	mySet2 := NewInt8Sync([]int8{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetInt8SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewInt8Sync([]int8{10, 20, 30, 20})
	mySet2 = NewInt8Sync([]int8{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewInt8Sync([]int8{10, 20, 30, 20, 40})
	mySet2 = NewInt8Sync([]int8{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetInt8SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetInt8SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewInt8Sync([]int8{100})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Int8Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int8(0); i < 100; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Int8Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := int8(0); i < 100; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 100 {
		t.Errorf("TestSetInt8SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
