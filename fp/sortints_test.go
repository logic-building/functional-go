package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSortInt(t *testing.T) {
	expectedList := []int{1, 2, 3, 4, 5}
	sortedList := SortInts([]int{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredints failed")
	}

	sortedList = SortInts([]int{})
	if len(sortedList) > 0 {
		t.Errorf("SortInts failed")
	}
}

func TestSortIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	expectedList := []*int{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortIntsPtr([]*int{&v5, &v1, &v4, &v2, &v3})
	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredints failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}
	sortedList = SortIntsPtr([]*int{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsPtr failed")
	}
}

func TestSortIntDesc(t *testing.T) {
	expectedList := []int{5, 4, 3, 2, 1}
	sortedList := SortIntsDesc([]int{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredints failed")
	}
	sortedList = SortIntsDesc([]int{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsDesc failed")
	}
}

func TestSortIntDescPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	expectedList := []*int{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortIntsDescPtr([]*int{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredints failed")
		}
	}

	sortedList = SortIntsDescPtr([]*int{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsDescPtr failed")
	}
}
