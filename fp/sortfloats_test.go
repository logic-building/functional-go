package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSortFloat64(t *testing.T) {
	expectedList := []float64{1, 2, 3, 4, 5}
	sortedList := SortFloats64([]float64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredfloat64s failed")
	}

	sortedList = SortFloats64([]float64{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64 failed")
	}
}

func TestSortFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	expectedList := []*float64{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortFloats64Ptr([]*float64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredfloat64s failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}

	sortedList = SortFloats64Ptr([]*float64{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64Ptr failed")
	}
}

func TestSortFloat64Desc(t *testing.T) {
	expectedList := []float64{5, 4, 3, 2, 1}
	sortedList := SortFloats64Desc([]float64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredfloat64s failed")
	}

	sortedList = SortFloats64Desc([]float64{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64Desc failed")
	}
}

func TestSortFloat64DescPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	expectedList := []*float64{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortFloats64DescPtr([]*float64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotredfloat64s failed")
		}
	}

	sortedList = SortFloats64DescPtr([]*float64{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64DescPtr failed")
	}
}
