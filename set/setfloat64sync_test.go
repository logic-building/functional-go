package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetFloat64SyncAdd(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	expected := []float64{10, 20, 30, 40}
	mySet := NewFloat64Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64SyncAdd failed")
		}
	}
}

func TestSetFloat64SyncRemove(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	expected := []float64{10, 20}
	mySet := NewFloat64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64SyncRemove failed")
		}
	}
}

func TestSetFloat64SyncClear(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetFloat64SyncClear failed.")
	}
}

func TestSetFloat64SyncContains(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetFloat64Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetFloat64SyncContains failed.")
	}
}

func TestSetFloat64SyncSize(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet := NewFloat64Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetFloat64Size failed.")
	}
}

func TestSetFloat64SyncJoin(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64Sync(list)

	list = []float64{30, 40, 50}
	mySet2 := NewFloat64Sync(list)

	expected := []float64{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64SyncJoin failed")
		}
	}
}

func TestSetFloat64SyncIntersection(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64Sync(list)

	list = []float64{30, 40, 50}
	mySet2 := NewFloat64Sync(list)

	expected := []float64{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat64SyncMinus(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64Sync(list)

	list = []float64{30, 40, 50}
	mySet2 := NewFloat64Sync(list)

	expected := []float64{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat64(num, expected) {
			t.Errorf("TestSetFloat64SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat64SyncSubset(t *testing.T) {
	list := []float64{10, 20, 30, 20}
	mySet1 := NewFloat64Sync(list)

	list = []float64{10, 20}
	mySet2 := NewFloat64Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetFloat64SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetFloat64SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat64SyncSuperset(t *testing.T) {
	mySet1 := NewFloat64Sync([]float64{10, 20, 30, 20})
	mySet2 := NewFloat64Sync([]float64{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetFloat64SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewFloat64Sync([]float64{10, 20, 30, 20})
	mySet2 = NewFloat64Sync([]float64{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewFloat64Sync([]float64{10, 20, 30, 20, 40})
	mySet2 = NewFloat64Sync([]float64{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat64SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat64SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewFloat64Sync([]float64{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Float64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := float64(0); i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Float64Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := float64(0); i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 1000 {
		t.Errorf("TestSetFloat64SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
