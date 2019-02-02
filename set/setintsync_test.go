package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetIntSyncAdd(t *testing.T) {
	list := []int{10, 20, 30, 20}
	expected := []int{10, 20, 30, 40}
	mySet := NewIntSync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.getList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntSyncAdd failed")
		}
	}
}

func TestSetIntSyncRemove(t *testing.T) {
	list := []int{10, 20, 30, 20}
	expected := []int{10, 20}
	mySet := NewIntSync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.getList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntSyncRemove failed")
		}
	}
}

func TestSetIntSyncClear(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewIntSync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.getList()) != 0 {
		t.Errorf("TestSetIntSyncClear failed.")
	}
}

func TestSetIntSyncContains(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewIntSync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetIntContains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetIntSyncContains failed.")
	}
}

func TestSetIntSyncSize(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet := NewIntSync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetIntSize failed.")
	}
}

func TestSetIntSyncJoin(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewIntSync(list)

	list = []int{30, 40, 50}
	mySet2 := NewIntSync(list)

	expected := []int{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.getList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntSyncJoin failed")
		}
	}
}

func TestSetIntSyncIntersection(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewIntSync(list)

	list = []int{30, 40, 50}
	mySet2 := NewIntSync(list)

	expected := []int{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.getList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntSyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.getList())
		}
	}
}

func TestSetIntSyncMinus(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewIntSync(list)

	list = []int{30, 40, 50}
	mySet2 := NewIntSync(list)

	expected := []int{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.getList() {
		if !fp.ExistsInt(num, expected) {
			t.Errorf("TestSetIntSyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.getList())
		}
	}
}

func TestSetIntSyncSubset(t *testing.T) {
	list := []int{10, 20, 30, 20}
	mySet1 := NewIntSync(list)

	list = []int{10, 20}
	mySet2 := NewIntSync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetIntSyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetIntSyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetIntSyncSuperset(t *testing.T) {
	mySet1 := NewIntSync([]int{10, 20, 30, 20})
	mySet2 := NewIntSync([]int{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetIntSyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewIntSync([]int{10, 20, 30, 20})
	mySet2 = NewIntSync([]int{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewIntSync([]int{10, 20, 30, 20, 40})
	mySet2 = NewIntSync([]int{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetIntSyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetIntSyncMultipleGoRoutine(t *testing.T) {
	mySet := NewIntSync([]int{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *IntSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *IntSync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.getList()) != 1 || mySet.getList()[0] != 1000 {
		t.Errorf("TestSetIntSyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.getList())
	}

}
