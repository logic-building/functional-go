package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSortStr(t *testing.T) {
	expectedList := []string{"1", "2", "3", "4", "5"}
	sortedList := SortStrs([]string{"5", "1", "4", "2", "3"})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredstrings failed")
	}

	sortedList = SortStrs([]string{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrs failed")
	}
}

func TestSortStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	expectedList := []*string{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortStrsPtr([]*string{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredstrings failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}

	sortedList = SortStrsPtr([]*string{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsPtr failed")
	}
}

func TestSortStrDesc(t *testing.T) {
	expectedList := []string{"5", "4", "3", "2", "1"}
	sortedList := SortStrsDesc([]string{"5", "1", "4", "2", "3"})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredstrings failed")
	}

	sortedList = SortStrsDesc([]string{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsDesc failed")
	}
}

func TestSortStrDescPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	expectedList := []*string{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortStrsDescPtr([]*string{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredstrings failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}
	sortedList = SortStrsDescPtr([]*string{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsDescPtr failed")
	}
}
