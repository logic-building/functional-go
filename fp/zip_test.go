package fp

import (
	"reflect"
	"testing"
)

func TestZipInt(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[int]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[int]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int]int{}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[int]int{}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int{}

	expectedMap = map[int]int{}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]int{}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]int{}
	actualMap = ZipInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
