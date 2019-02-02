package set

import (
	"github.com/logic-building/functional-go/fp"
	"sync"
	"testing"
)

func TestSetFloat32SyncAdd(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	expected := []float32{10, 20, 30, 40}
	mySet := NewFloat32Sync(list)
	mySet.Add(40)
	mySet.Add(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32SyncAdd failed")
		}
	}
}

func TestSetFloat32SyncRemove(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	expected := []float32{10, 20}
	mySet := NewFloat32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)

	for _, num := range mySet.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32SyncRemove failed")
		}
	}
}

func TestSetFloat32SyncClear(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32Sync(list)
	mySet.Add(40)
	mySet.Add(40)
	mySet.Remove(30)
	mySet.Remove(40)
	mySet.Clear()

	if len(mySet.GetList()) != 0 {
		t.Errorf("TestSetFloat32SyncClear failed.")
	}
}

func TestSetFloat32SyncContains(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32Sync(list)

	if !mySet.Contains(20) {
		t.Errorf("TestSetFloat32Contains failed.")
	}

	if mySet.Contains(200) {
		t.Errorf("TestSetFloat32SyncContains failed.")
	}
}

func TestSetFloat32SyncSize(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet := NewFloat32Sync(list)

	if mySet.Size() != 3 {
		t.Errorf("TestSetFloat32Size failed.")
	}
}

func TestSetFloat32SyncJoin(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32Sync(list)

	list = []float32{30, 40, 50}
	mySet2 := NewFloat32Sync(list)

	expected := []float32{10, 20, 30, 40, 50}

	mySet3 := mySet1.Union(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32SyncJoin failed")
		}
	}
}

func TestSetFloat32SyncIntersection(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32Sync(list)

	list = []float32{30, 40, 50}
	mySet2 := NewFloat32Sync(list)

	expected := []float32{30}

	mySet3 := mySet1.Intersection(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32SyncIntersection failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat32SyncMinus(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32Sync(list)

	list = []float32{30, 40, 50}
	mySet2 := NewFloat32Sync(list)

	expected := []float32{10, 20}

	mySet3 := mySet1.Minus(mySet2)

	for _, num := range mySet3.GetList() {
		if !fp.ExistsFloat32(num, expected) {
			t.Errorf("TestSetFloat32SyncMinus failed. Expected=%v, Actual=%v", expected, mySet3.GetList())
		}
	}
}

func TestSetFloat32SyncSubset(t *testing.T) {
	list := []float32{10, 20, 30, 20}
	mySet1 := NewFloat32Sync(list)

	list = []float32{10, 20}
	mySet2 := NewFloat32Sync(list)

	if !mySet2.Subset(mySet1) {
		t.Errorf("TestSetFloat32SyncSubset failed. Expected=true, Actual=false")
	}

	if mySet1.Subset(mySet2) {
		t.Errorf("TestSetFloat32SyncSubset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat32SyncSuperset(t *testing.T) {
	mySet1 := NewFloat32Sync([]float32{10, 20, 30, 20})
	mySet2 := NewFloat32Sync([]float32{10, 20})

	if !mySet1.Superset(mySet2) {
		t.Errorf("TestSetFloat32SyncSuperset failed. Expected=true, Actual=false")
	}

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32SyncSuperset failed. Expected=false, Actual=true")
	}

	mySet1 = NewFloat32Sync([]float32{10, 20, 30, 20})
	mySet2 = NewFloat32Sync([]float32{10, 20, 30, 20})

	if !mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32SyncSuperset failed. Expected=true, Actual=false")
	}

	mySet1 = NewFloat32Sync([]float32{10, 20, 30, 20, 40})
	mySet2 = NewFloat32Sync([]float32{10, 20, 30, 20})

	if mySet2.Superset(mySet1) {
		t.Errorf("TestSetFloat32SyncSuperset failed. Expected=false, Actual=true")
	}
}

func TestSetFloat32SyncMultipleGoRoutine(t *testing.T) {
	mySet := NewFloat32Sync([]float32{1000})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(mySet *Float32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := float32(0); i < 1000; i++ {
			mySet.Add(i)
		}
	}(mySet, &wg)

	go func(mySet *Float32Sync, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := float32(0); i < 1000; i++ {
			for {
				if mySet.Remove(i) {
					break
				}
			}
		}
	}(mySet, &wg)

	wg.Wait()
	if len(mySet.GetList()) != 1 || mySet.GetList()[0] != 1000 {
		t.Errorf("TestSetFloat32SyncMultipleGoRoutine failed. Expetected [1000], Actual=%v", mySet.GetList())
	}

}
