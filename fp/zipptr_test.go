package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestZipIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int{}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*int]*int{}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int{}

	expectedMap = map[*int]*int{}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*int{}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*int{}
	actualMap = ZipIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int64{}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*int]*int64{}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int64{}

	expectedMap = map[*int]*int64{}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*int64{}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*int64{}
	actualMap = ZipIntInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int32{}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*int]*int32{}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int32{}

	expectedMap = map[*int]*int32{}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*int32{}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*int32{}
	actualMap = ZipIntInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt16Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int16{}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*int]*int16{}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int16{}

	expectedMap = map[*int]*int16{}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*int16{}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*int16{}
	actualMap = ZipIntInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt8Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*int8{}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*int]*int8{}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*int8{}

	expectedMap = map[*int]*int8{}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*int8{}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*int8{}
	actualMap = ZipIntInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUintPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint{}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*int]*uint{}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint{}

	expectedMap = map[*int]*uint{}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*uint{}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*uint{}
	actualMap = ZipIntUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint64{}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*int]*uint64{}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint64{}

	expectedMap = map[*int]*uint64{}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*uint64{}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*uint64{}
	actualMap = ZipIntUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint32{}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*int]*uint32{}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint32{}

	expectedMap = map[*int]*uint32{}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*uint32{}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*uint32{}
	actualMap = ZipIntUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint16Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint16{}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*int]*uint16{}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint16{}

	expectedMap = map[*int]*uint16{}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*uint16{}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*uint16{}
	actualMap = ZipIntUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint8Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*uint8{}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*int]*uint8{}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*uint8{}

	expectedMap = map[*int]*uint8{}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*uint8{}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*uint8{}
	actualMap = ZipIntUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntStrPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*string{}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*int]*string{}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*string{}

	expectedMap = map[*int]*string{}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*string{}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*string{}
	actualMap = ZipIntStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntBoolPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var vt bool = true
	var vf bool = false

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*int]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*int]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*int]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*int]*bool{}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*int]*bool{}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*bool{}

	expectedMap = map[*int]*bool{}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*bool{}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*bool{}
	actualMap = ZipIntBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntFloat32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*float32{}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*int]*float32{}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*float32{}

	expectedMap = map[*int]*float32{}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*float32{}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*float32{}
	actualMap = ZipIntFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntFloat64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*int{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int]*float64{}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*int]*float64{}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = []*float64{}

	expectedMap = map[*int]*float64{}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int]*float64{}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int{}
	list2 = nil

	expectedMap = map[*int]*float64{}
	actualMap = ZipIntFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64IntPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int{}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*int64]*int{}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int{}

	expectedMap = map[*int64]*int{}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*int{}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*int{}
	actualMap = ZipInt64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int64{}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*int64]*int64{}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int64{}

	expectedMap = map[*int64]*int64{}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*int64{}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*int64{}
	actualMap = ZipInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int32{}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*int64]*int32{}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int32{}

	expectedMap = map[*int64]*int32{}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*int32{}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*int32{}
	actualMap = ZipInt64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int16Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int16{}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*int64]*int16{}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int16{}

	expectedMap = map[*int64]*int16{}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*int16{}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*int16{}
	actualMap = ZipInt64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int8Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*int8{}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*int64]*int8{}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*int8{}

	expectedMap = map[*int64]*int8{}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*int8{}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*int8{}
	actualMap = ZipInt64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64UintPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint{}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*int64]*uint{}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint{}

	expectedMap = map[*int64]*uint{}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*uint{}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*uint{}
	actualMap = ZipInt64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint64{}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*int64]*uint64{}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint64{}

	expectedMap = map[*int64]*uint64{}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*uint64{}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*uint64{}
	actualMap = ZipInt64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint32{}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*int64]*uint32{}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint32{}

	expectedMap = map[*int64]*uint32{}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*uint32{}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*uint32{}
	actualMap = ZipInt64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint16Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint16{}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*int64]*uint16{}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint16{}

	expectedMap = map[*int64]*uint16{}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*uint16{}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*uint16{}
	actualMap = ZipInt64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint8Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*uint8{}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*int64]*uint8{}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*uint8{}

	expectedMap = map[*int64]*uint8{}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*uint8{}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*uint8{}
	actualMap = ZipInt64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64StrPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*string{}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*int64]*string{}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*string{}

	expectedMap = map[*int64]*string{}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*string{}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*string{}
	actualMap = ZipInt64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64BoolPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*int64]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*int64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*int64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*int64]*bool{}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*int64]*bool{}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*bool{}

	expectedMap = map[*int64]*bool{}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*bool{}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*bool{}
	actualMap = ZipInt64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Float32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*float32{}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*int64]*float32{}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*float32{}

	expectedMap = map[*int64]*float32{}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*float32{}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*float32{}
	actualMap = ZipInt64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Float64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*int64{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int64]*float64{}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*int64]*float64{}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = []*float64{}

	expectedMap = map[*int64]*float64{}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int64]*float64{}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int64{}
	list2 = nil

	expectedMap = map[*int64]*float64{}
	actualMap = ZipInt64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32IntPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int{}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*int32]*int{}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int{}

	expectedMap = map[*int32]*int{}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*int{}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*int{}
	actualMap = ZipInt32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int64{}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*int32]*int64{}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int64{}

	expectedMap = map[*int32]*int64{}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*int64{}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*int64{}
	actualMap = ZipInt32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int32{}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*int32]*int32{}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int32{}

	expectedMap = map[*int32]*int32{}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*int32{}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*int32{}
	actualMap = ZipInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int16Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int16{}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*int32]*int16{}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int16{}

	expectedMap = map[*int32]*int16{}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*int16{}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*int16{}
	actualMap = ZipInt32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int8Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*int8{}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*int32]*int8{}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*int8{}

	expectedMap = map[*int32]*int8{}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*int8{}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*int8{}
	actualMap = ZipInt32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32UintPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint{}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*int32]*uint{}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint{}

	expectedMap = map[*int32]*uint{}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*uint{}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*uint{}
	actualMap = ZipInt32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint64{}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*int32]*uint64{}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint64{}

	expectedMap = map[*int32]*uint64{}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*uint64{}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*uint64{}
	actualMap = ZipInt32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint32{}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*int32]*uint32{}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint32{}

	expectedMap = map[*int32]*uint32{}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*uint32{}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*uint32{}
	actualMap = ZipInt32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint16Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint16{}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*int32]*uint16{}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint16{}

	expectedMap = map[*int32]*uint16{}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*uint16{}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*uint16{}
	actualMap = ZipInt32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint8Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*uint8{}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*int32]*uint8{}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*uint8{}

	expectedMap = map[*int32]*uint8{}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*uint8{}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*uint8{}
	actualMap = ZipInt32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32StrPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*string{}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*int32]*string{}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*string{}

	expectedMap = map[*int32]*string{}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*string{}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*string{}
	actualMap = ZipInt32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32BoolPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*int32]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*int32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*int32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*int32]*bool{}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*int32]*bool{}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*bool{}

	expectedMap = map[*int32]*bool{}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*bool{}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*bool{}
	actualMap = ZipInt32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Float32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*float32{}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*int32]*float32{}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*float32{}

	expectedMap = map[*int32]*float32{}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*float32{}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*float32{}
	actualMap = ZipInt32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Float64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*int32{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int32]*float64{}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*int32]*float64{}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = []*float64{}

	expectedMap = map[*int32]*float64{}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int32]*float64{}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int32{}
	list2 = nil

	expectedMap = map[*int32]*float64{}
	actualMap = ZipInt32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16IntPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int{}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*int16]*int{}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int{}

	expectedMap = map[*int16]*int{}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*int{}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*int{}
	actualMap = ZipInt16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int64{}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*int16]*int64{}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int64{}

	expectedMap = map[*int16]*int64{}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*int64{}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*int64{}
	actualMap = ZipInt16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int32{}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*int16]*int32{}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int32{}

	expectedMap = map[*int16]*int32{}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*int32{}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*int32{}
	actualMap = ZipInt16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int16{}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*int16]*int16{}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int16{}

	expectedMap = map[*int16]*int16{}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*int16{}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*int16{}
	actualMap = ZipInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int8Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*int8{}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*int16]*int8{}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*int8{}

	expectedMap = map[*int16]*int8{}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*int8{}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*int8{}
	actualMap = ZipInt16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16UintPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint{}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*int16]*uint{}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint{}

	expectedMap = map[*int16]*uint{}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*uint{}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*uint{}
	actualMap = ZipInt16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint64{}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*int16]*uint64{}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint64{}

	expectedMap = map[*int16]*uint64{}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*uint64{}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*uint64{}
	actualMap = ZipInt16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint32{}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*int16]*uint32{}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint32{}

	expectedMap = map[*int16]*uint32{}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*uint32{}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*uint32{}
	actualMap = ZipInt16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint16{}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*int16]*uint16{}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint16{}

	expectedMap = map[*int16]*uint16{}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*uint16{}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*uint16{}
	actualMap = ZipInt16Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint8Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*uint8{}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*int16]*uint8{}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*uint8{}

	expectedMap = map[*int16]*uint8{}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*uint8{}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*uint8{}
	actualMap = ZipInt16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16StrPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*string{}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*int16]*string{}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*string{}

	expectedMap = map[*int16]*string{}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*string{}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*string{}
	actualMap = ZipInt16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16BoolPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*int16]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*int16]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*int16]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*int16]*bool{}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*int16]*bool{}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*bool{}

	expectedMap = map[*int16]*bool{}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*bool{}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*bool{}
	actualMap = ZipInt16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Float32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*float32{}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*int16]*float32{}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*float32{}

	expectedMap = map[*int16]*float32{}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*float32{}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*float32{}
	actualMap = ZipInt16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Float64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*int16{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int16]*float64{}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*int16]*float64{}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = []*float64{}

	expectedMap = map[*int16]*float64{}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int16]*float64{}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int16{}
	list2 = nil

	expectedMap = map[*int16]*float64{}
	actualMap = ZipInt16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8IntPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int{}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*int8]*int{}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int{}

	expectedMap = map[*int8]*int{}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*int{}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*int{}
	actualMap = ZipInt8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int64{}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*int8]*int64{}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int64{}

	expectedMap = map[*int8]*int64{}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*int64{}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*int64{}
	actualMap = ZipInt8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int32{}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*int8]*int32{}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int32{}

	expectedMap = map[*int8]*int32{}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*int32{}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*int32{}
	actualMap = ZipInt8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int16Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int16{}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*int8]*int16{}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int16{}

	expectedMap = map[*int8]*int16{}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*int16{}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*int16{}
	actualMap = ZipInt8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*int8{}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*int8]*int8{}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*int8{}

	expectedMap = map[*int8]*int8{}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*int8{}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*int8{}
	actualMap = ZipInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8UintPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint{}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*int8]*uint{}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint{}

	expectedMap = map[*int8]*uint{}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*uint{}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*uint{}
	actualMap = ZipInt8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint64{}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*int8]*uint64{}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint64{}

	expectedMap = map[*int8]*uint64{}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*uint64{}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*uint64{}
	actualMap = ZipInt8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint32{}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*int8]*uint32{}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint32{}

	expectedMap = map[*int8]*uint32{}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*uint32{}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*uint32{}
	actualMap = ZipInt8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint16Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint16{}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*int8]*uint16{}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint16{}

	expectedMap = map[*int8]*uint16{}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*uint16{}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*uint16{}
	actualMap = ZipInt8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*uint8{}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*int8]*uint8{}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*uint8{}

	expectedMap = map[*int8]*uint8{}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*uint8{}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*uint8{}
	actualMap = ZipInt8Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8StrPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*string{}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*int8]*string{}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*string{}

	expectedMap = map[*int8]*string{}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*string{}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*string{}
	actualMap = ZipInt8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8BoolPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*int8]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*int8]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*int8]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*int8]*bool{}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*int8]*bool{}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*bool{}

	expectedMap = map[*int8]*bool{}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*bool{}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*bool{}
	actualMap = ZipInt8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Float32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*float32{}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*int8]*float32{}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*float32{}

	expectedMap = map[*int8]*float32{}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*float32{}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*float32{}
	actualMap = ZipInt8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Float64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*int8{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*int8]*float64{}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*int8]*float64{}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = []*float64{}

	expectedMap = map[*int8]*float64{}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*int8]*float64{}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*int8{}
	list2 = nil

	expectedMap = map[*int8]*float64{}
	actualMap = ZipInt8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintIntPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int{}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*uint]*int{}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int{}

	expectedMap = map[*uint]*int{}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*int{}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*int{}
	actualMap = ZipUintIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int64{}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*uint]*int64{}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int64{}

	expectedMap = map[*uint]*int64{}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*int64{}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*int64{}
	actualMap = ZipUintInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int32{}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*uint]*int32{}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int32{}

	expectedMap = map[*uint]*int32{}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*int32{}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*int32{}
	actualMap = ZipUintInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt16Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int16{}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*uint]*int16{}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int16{}

	expectedMap = map[*uint]*int16{}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*int16{}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*int16{}
	actualMap = ZipUintInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt8Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*int8{}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*uint]*int8{}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*int8{}

	expectedMap = map[*uint]*int8{}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*int8{}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*int8{}
	actualMap = ZipUintInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint{}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*uint]*uint{}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint{}

	expectedMap = map[*uint]*uint{}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*uint{}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*uint{}
	actualMap = ZipUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint64{}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*uint]*uint64{}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint64{}

	expectedMap = map[*uint]*uint64{}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*uint64{}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*uint64{}
	actualMap = ZipUintUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint32{}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*uint]*uint32{}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint32{}

	expectedMap = map[*uint]*uint32{}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*uint32{}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*uint32{}
	actualMap = ZipUintUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint16Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint16{}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*uint]*uint16{}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint16{}

	expectedMap = map[*uint]*uint16{}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*uint16{}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*uint16{}
	actualMap = ZipUintUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint8Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*uint8{}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*uint]*uint8{}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*uint8{}

	expectedMap = map[*uint]*uint8{}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*uint8{}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*uint8{}
	actualMap = ZipUintUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintStrPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*string{}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*uint]*string{}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*string{}

	expectedMap = map[*uint]*string{}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*string{}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*string{}
	actualMap = ZipUintStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintBoolPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var vt bool = true
	var vf bool = false

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*uint]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*uint]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*uint]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*uint]*bool{}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*uint]*bool{}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*bool{}

	expectedMap = map[*uint]*bool{}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*bool{}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*bool{}
	actualMap = ZipUintBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintFloat32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*float32{}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*uint]*float32{}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*float32{}

	expectedMap = map[*uint]*float32{}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*float32{}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*float32{}
	actualMap = ZipUintFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintFloat64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*uint{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint]*float64{}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*uint]*float64{}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = []*float64{}

	expectedMap = map[*uint]*float64{}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint]*float64{}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint{}
	list2 = nil

	expectedMap = map[*uint]*float64{}
	actualMap = ZipUintFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64IntPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int{}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*uint64]*int{}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int{}

	expectedMap = map[*uint64]*int{}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*int{}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*int{}
	actualMap = ZipUint64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int64{}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*uint64]*int64{}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int64{}

	expectedMap = map[*uint64]*int64{}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*int64{}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*int64{}
	actualMap = ZipUint64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int32{}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*uint64]*int32{}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int32{}

	expectedMap = map[*uint64]*int32{}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*int32{}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*int32{}
	actualMap = ZipUint64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int16Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int16{}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*uint64]*int16{}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int16{}

	expectedMap = map[*uint64]*int16{}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*int16{}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*int16{}
	actualMap = ZipUint64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int8Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*int8{}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*uint64]*int8{}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*int8{}

	expectedMap = map[*uint64]*int8{}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*int8{}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*int8{}
	actualMap = ZipUint64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64UintPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint{}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*uint64]*uint{}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint{}

	expectedMap = map[*uint64]*uint{}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*uint{}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*uint{}
	actualMap = ZipUint64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint64{}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*uint64]*uint64{}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint64{}

	expectedMap = map[*uint64]*uint64{}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*uint64{}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*uint64{}
	actualMap = ZipUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint32{}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*uint64]*uint32{}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint32{}

	expectedMap = map[*uint64]*uint32{}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*uint32{}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*uint32{}
	actualMap = ZipUint64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint16Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint16{}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*uint64]*uint16{}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint16{}

	expectedMap = map[*uint64]*uint16{}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*uint16{}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*uint16{}
	actualMap = ZipUint64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint8Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*uint8{}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*uint64]*uint8{}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*uint8{}

	expectedMap = map[*uint64]*uint8{}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*uint8{}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*uint8{}
	actualMap = ZipUint64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64StrPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*string{}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*uint64]*string{}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*string{}

	expectedMap = map[*uint64]*string{}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*string{}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*string{}
	actualMap = ZipUint64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64BoolPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*uint64]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*uint64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*uint64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*uint64]*bool{}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*uint64]*bool{}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*bool{}

	expectedMap = map[*uint64]*bool{}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*bool{}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*bool{}
	actualMap = ZipUint64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Float32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*float32{}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*uint64]*float32{}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*float32{}

	expectedMap = map[*uint64]*float32{}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*float32{}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*float32{}
	actualMap = ZipUint64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Float64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*uint64{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint64]*float64{}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*uint64]*float64{}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = []*float64{}

	expectedMap = map[*uint64]*float64{}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint64]*float64{}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint64{}
	list2 = nil

	expectedMap = map[*uint64]*float64{}
	actualMap = ZipUint64Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32IntPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int{}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*uint32]*int{}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int{}

	expectedMap = map[*uint32]*int{}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*int{}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*int{}
	actualMap = ZipUint32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int64{}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*uint32]*int64{}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int64{}

	expectedMap = map[*uint32]*int64{}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*int64{}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*int64{}
	actualMap = ZipUint32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int32{}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*uint32]*int32{}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int32{}

	expectedMap = map[*uint32]*int32{}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*int32{}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*int32{}
	actualMap = ZipUint32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int16Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int16{}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*uint32]*int16{}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int16{}

	expectedMap = map[*uint32]*int16{}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*int16{}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*int16{}
	actualMap = ZipUint32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int8Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*int8{}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*uint32]*int8{}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*int8{}

	expectedMap = map[*uint32]*int8{}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*int8{}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*int8{}
	actualMap = ZipUint32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32UintPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint{}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*uint32]*uint{}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint{}

	expectedMap = map[*uint32]*uint{}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*uint{}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*uint{}
	actualMap = ZipUint32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint64{}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*uint32]*uint64{}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint64{}

	expectedMap = map[*uint32]*uint64{}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*uint64{}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*uint64{}
	actualMap = ZipUint32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint32{}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*uint32]*uint32{}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint32{}

	expectedMap = map[*uint32]*uint32{}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*uint32{}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*uint32{}
	actualMap = ZipUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint16Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint16{}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*uint32]*uint16{}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint16{}

	expectedMap = map[*uint32]*uint16{}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*uint16{}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*uint16{}
	actualMap = ZipUint32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint8Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*uint8{}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*uint32]*uint8{}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*uint8{}

	expectedMap = map[*uint32]*uint8{}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*uint8{}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*uint8{}
	actualMap = ZipUint32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32StrPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*string{}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*uint32]*string{}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*string{}

	expectedMap = map[*uint32]*string{}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*string{}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*string{}
	actualMap = ZipUint32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32BoolPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*uint32]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*uint32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*uint32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*uint32]*bool{}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*uint32]*bool{}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*bool{}

	expectedMap = map[*uint32]*bool{}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*bool{}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*bool{}
	actualMap = ZipUint32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Float32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*float32{}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*uint32]*float32{}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*float32{}

	expectedMap = map[*uint32]*float32{}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*float32{}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*float32{}
	actualMap = ZipUint32Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Float64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*uint32{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint32]*float64{}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*uint32]*float64{}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = []*float64{}

	expectedMap = map[*uint32]*float64{}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint32]*float64{}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint32{}
	list2 = nil

	expectedMap = map[*uint32]*float64{}
	actualMap = ZipUint32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16IntPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int{}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*uint16]*int{}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int{}

	expectedMap = map[*uint16]*int{}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*int{}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*int{}
	actualMap = ZipUint16IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int64{}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*uint16]*int64{}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int64{}

	expectedMap = map[*uint16]*int64{}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*int64{}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*int64{}
	actualMap = ZipUint16Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int32{}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*uint16]*int32{}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int32{}

	expectedMap = map[*uint16]*int32{}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*int32{}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*int32{}
	actualMap = ZipUint16Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int16{}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*uint16]*int16{}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int16{}

	expectedMap = map[*uint16]*int16{}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*int16{}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*int16{}
	actualMap = ZipUint16Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int8Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*int8{}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*uint16]*int8{}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*int8{}

	expectedMap = map[*uint16]*int8{}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*int8{}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*int8{}
	actualMap = ZipUint16Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16UintPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint{}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*uint16]*uint{}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint{}

	expectedMap = map[*uint16]*uint{}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*uint{}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*uint{}
	actualMap = ZipUint16UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint64{}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*uint16]*uint64{}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint64{}

	expectedMap = map[*uint16]*uint64{}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*uint64{}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*uint64{}
	actualMap = ZipUint16Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint32{}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*uint16]*uint32{}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint32{}

	expectedMap = map[*uint16]*uint32{}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*uint32{}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*uint32{}
	actualMap = ZipUint16Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint16{}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*uint16]*uint16{}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint16{}

	expectedMap = map[*uint16]*uint16{}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*uint16{}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*uint16{}
	actualMap = ZipUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint8Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*uint8{}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*uint16]*uint8{}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*uint8{}

	expectedMap = map[*uint16]*uint8{}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*uint8{}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*uint8{}
	actualMap = ZipUint16Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16StrPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*string{}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*uint16]*string{}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*string{}

	expectedMap = map[*uint16]*string{}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*string{}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*string{}
	actualMap = ZipUint16StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16BoolPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*uint16]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*uint16]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*uint16]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*uint16]*bool{}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*uint16]*bool{}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*bool{}

	expectedMap = map[*uint16]*bool{}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*bool{}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*bool{}
	actualMap = ZipUint16BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Float32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*float32{}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*uint16]*float32{}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*float32{}

	expectedMap = map[*uint16]*float32{}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*float32{}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*float32{}
	actualMap = ZipUint16Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Float64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*uint16{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint16]*float64{}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*uint16]*float64{}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = []*float64{}

	expectedMap = map[*uint16]*float64{}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint16]*float64{}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint16{}
	list2 = nil

	expectedMap = map[*uint16]*float64{}
	actualMap = ZipUint16Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8IntPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int{}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*uint8]*int{}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int{}

	expectedMap = map[*uint8]*int{}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*int{}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*int{}
	actualMap = ZipUint8IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int64{}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*uint8]*int64{}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int64{}

	expectedMap = map[*uint8]*int64{}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*int64{}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*int64{}
	actualMap = ZipUint8Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int32{}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*uint8]*int32{}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int32{}

	expectedMap = map[*uint8]*int32{}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*int32{}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*int32{}
	actualMap = ZipUint8Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int16Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int16{}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*uint8]*int16{}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int16{}

	expectedMap = map[*uint8]*int16{}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*int16{}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*int16{}
	actualMap = ZipUint8Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*int8{}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*uint8]*int8{}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*int8{}

	expectedMap = map[*uint8]*int8{}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*int8{}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*int8{}
	actualMap = ZipUint8Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8UintPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint{}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*uint8]*uint{}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint{}

	expectedMap = map[*uint8]*uint{}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*uint{}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*uint{}
	actualMap = ZipUint8UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint64{}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*uint8]*uint64{}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint64{}

	expectedMap = map[*uint8]*uint64{}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*uint64{}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*uint64{}
	actualMap = ZipUint8Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint32{}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*uint8]*uint32{}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint32{}

	expectedMap = map[*uint8]*uint32{}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*uint32{}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*uint32{}
	actualMap = ZipUint8Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint16Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint16{}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*uint8]*uint16{}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint16{}

	expectedMap = map[*uint8]*uint16{}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*uint16{}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*uint16{}
	actualMap = ZipUint8Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*uint8{}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*uint8]*uint8{}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*uint8{}

	expectedMap = map[*uint8]*uint8{}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*uint8{}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*uint8{}
	actualMap = ZipUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8StrPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*string{}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*uint8]*string{}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*string{}

	expectedMap = map[*uint8]*string{}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*string{}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*string{}
	actualMap = ZipUint8StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8BoolPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*uint8]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*uint8]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*uint8]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*uint8]*bool{}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*uint8]*bool{}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*bool{}

	expectedMap = map[*uint8]*bool{}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*bool{}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*bool{}
	actualMap = ZipUint8BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Float32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*float32{}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*uint8]*float32{}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*float32{}

	expectedMap = map[*uint8]*float32{}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*float32{}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*float32{}
	actualMap = ZipUint8Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Float64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*uint8{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*uint8]*float64{}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*uint8]*float64{}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = []*float64{}

	expectedMap = map[*uint8]*float64{}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*uint8]*float64{}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*uint8{}
	list2 = nil

	expectedMap = map[*uint8]*float64{}
	actualMap = ZipUint8Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrIntPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int{}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*string]*int{}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int{}

	expectedMap = map[*string]*int{}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*int{}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*int{}
	actualMap = ZipStrIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int64{}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*string]*int64{}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int64{}

	expectedMap = map[*string]*int64{}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*int64{}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*int64{}
	actualMap = ZipStrInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int32{}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*string]*int32{}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int32{}

	expectedMap = map[*string]*int32{}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*int32{}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*int32{}
	actualMap = ZipStrInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt16Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int16{}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*string]*int16{}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int16{}

	expectedMap = map[*string]*int16{}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*int16{}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*int16{}
	actualMap = ZipStrInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt8Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*int8{}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*string]*int8{}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*int8{}

	expectedMap = map[*string]*int8{}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*int8{}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*int8{}
	actualMap = ZipStrInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUintPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint{}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*string]*uint{}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint{}

	expectedMap = map[*string]*uint{}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*uint{}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*uint{}
	actualMap = ZipStrUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint64{}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*string]*uint64{}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint64{}

	expectedMap = map[*string]*uint64{}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*uint64{}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*uint64{}
	actualMap = ZipStrUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint32{}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*string]*uint32{}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint32{}

	expectedMap = map[*string]*uint32{}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*uint32{}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*uint32{}
	actualMap = ZipStrUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint16Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint16{}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*string]*uint16{}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint16{}

	expectedMap = map[*string]*uint16{}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*uint16{}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*uint16{}
	actualMap = ZipStrUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint8Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*uint8{}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*string]*uint8{}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*uint8{}

	expectedMap = map[*string]*uint8{}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*uint8{}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*uint8{}
	actualMap = ZipStrUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*string{}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*string]*string{}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*string{}

	expectedMap = map[*string]*string{}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*string{}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*string{}
	actualMap = ZipStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrBoolPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var vt bool = true
	var vf bool = false

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*string]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*string]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*string]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*string]*bool{}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*string]*bool{}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*bool{}

	expectedMap = map[*string]*bool{}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*bool{}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*bool{}
	actualMap = ZipStrBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrFloat32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*float32{}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*string]*float32{}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*float32{}

	expectedMap = map[*string]*float32{}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*float32{}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*float32{}
	actualMap = ZipStrFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrFloat64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*string{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*string]*float64{}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*string]*float64{}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = []*float64{}

	expectedMap = map[*string]*float64{}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*string]*float64{}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*string{}
	list2 = nil

	expectedMap = map[*string]*float64{}
	actualMap = ZipStrFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*int{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*int{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int{&v1, &v2, &v3}

	expectedMap = map[*bool]*int{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*int{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int{}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int{}

	expectedMap = map[*bool]*int{}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int{}

	expectedMap = map[*bool]*int{}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*int{}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*int{}
	actualMap = ZipBoolIntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolIntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*int64{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*int64{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int64{&v1, &v2, &v3}

	expectedMap = map[*bool]*int64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*int64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int64{}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int64{}

	expectedMap = map[*bool]*int64{}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int64{}

	expectedMap = map[*bool]*int64{}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*int64{}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*int64{}
	actualMap = ZipBoolInt64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*int32{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*int32{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int32{&v1, &v2, &v3}

	expectedMap = map[*bool]*int32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*int32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int32{}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int32{}

	expectedMap = map[*bool]*int32{}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int32{}

	expectedMap = map[*bool]*int32{}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*int32{}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*int32{}
	actualMap = ZipBoolInt32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*int16{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*int16{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int16{&v1, &v2, &v3}

	expectedMap = map[*bool]*int16{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*int16{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int16{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int16{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int16{}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int16{}

	expectedMap = map[*bool]*int16{}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int16{}

	expectedMap = map[*bool]*int16{}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*int16{}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*int16{}
	actualMap = ZipBoolInt16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*int8{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*int8{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int8{&v1, &v2, &v3}

	expectedMap = map[*bool]*int8{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*int8{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int8{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int8{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*int8{}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*int8{}

	expectedMap = map[*bool]*int8{}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*int8{}

	expectedMap = map[*bool]*int8{}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*int8{}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*int8{}
	actualMap = ZipBoolInt8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*uint{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*uint{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint{&v1, &v2, &v3}

	expectedMap = map[*bool]*uint{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*uint{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint{}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint{}

	expectedMap = map[*bool]*uint{}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint{}

	expectedMap = map[*bool]*uint{}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*uint{}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*uint{}
	actualMap = ZipBoolUintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*uint64{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*uint64{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint64{&v1, &v2, &v3}

	expectedMap = map[*bool]*uint64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*uint64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint64{}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint64{}

	expectedMap = map[*bool]*uint64{}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint64{}

	expectedMap = map[*bool]*uint64{}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*uint64{}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*uint64{}
	actualMap = ZipBoolUint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*uint32{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*uint32{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint32{&v1, &v2, &v3}

	expectedMap = map[*bool]*uint32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*uint32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint32{}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint32{}

	expectedMap = map[*bool]*uint32{}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint32{}

	expectedMap = map[*bool]*uint32{}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*uint32{}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*uint32{}
	actualMap = ZipBoolUint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*uint16{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*uint16{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint16{&v1, &v2, &v3}

	expectedMap = map[*bool]*uint16{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*uint16{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint16{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint16{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint16{}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint16{}

	expectedMap = map[*bool]*uint16{}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint16{}

	expectedMap = map[*bool]*uint16{}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*uint16{}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*uint16{}
	actualMap = ZipBoolUint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*uint8{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*uint8{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint8{&v1, &v2, &v3}

	expectedMap = map[*bool]*uint8{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*uint8{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint8{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint8{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*uint8{}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*uint8{}

	expectedMap = map[*bool]*uint8{}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*uint8{}

	expectedMap = map[*bool]*uint8{}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*uint8{}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*uint8{}
	actualMap = ZipBoolUint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*string{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*string{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*string{&v1, &v2, &v3}

	expectedMap = map[*bool]*string{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*string{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*string{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*string{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*string{}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*string{}

	expectedMap = map[*bool]*string{}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*string{}

	expectedMap = map[*bool]*string{}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*string{}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*string{}
	actualMap = ZipBoolStrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*bool]*bool{&vt: &vt, &vf: &vf}
	actualMap := ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*bool]*bool{}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*bool{}

	expectedMap = map[*bool]*bool{}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*bool{}

	expectedMap = map[*bool]*bool{}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*bool{}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*bool{}
	actualMap = ZipBoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*float32{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*float32{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*float32{&v1, &v2, &v3}

	expectedMap = map[*bool]*float32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*float32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*float32{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*float32{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*float32{}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*float32{}

	expectedMap = map[*bool]*float32{}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*float32{}

	expectedMap = map[*bool]*float32{}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*float32{}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*float32{}
	actualMap = ZipBoolFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vt, &vf, &vt}
	list2 := []*float64{&v1, &v2, &v3, &v4}

	expectedMap := map[*bool]*float64{&vt: &v4, &vf: &v3}
	actualMap := ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*float64{&v1, &v2, &v3}

	expectedMap = map[*bool]*float64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf}
	list2 = []*float64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*float64{&vt: &v2, &vf: &v3}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*float64{&v1, &v2, &v3, &v4}

	expectedMap = map[*bool]*float64{}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{&vt, &vt, &vf, &vt}
	list2 = []*float64{}

	expectedMap = map[*bool]*float64{}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = []*float64{}

	expectedMap = map[*bool]*float64{}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*bool]*float64{}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*bool{}
	list2 = nil

	expectedMap = map[*bool]*float64{}
	actualMap = ZipBoolFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32IntPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int{}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*float32]*int{}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int{}

	expectedMap = map[*float32]*int{}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*int{}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*int{}
	actualMap = ZipFloat32IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int64{}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*float32]*int64{}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int64{}

	expectedMap = map[*float32]*int64{}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*int64{}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*int64{}
	actualMap = ZipFloat32Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int32{}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*float32]*int32{}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int32{}

	expectedMap = map[*float32]*int32{}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*int32{}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*int32{}
	actualMap = ZipFloat32Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int16Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int16{}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*float32]*int16{}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int16{}

	expectedMap = map[*float32]*int16{}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*int16{}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*int16{}
	actualMap = ZipFloat32Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int8Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*int8{}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*float32]*int8{}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*int8{}

	expectedMap = map[*float32]*int8{}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*int8{}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*int8{}
	actualMap = ZipFloat32Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32UintPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint{}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*float32]*uint{}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint{}

	expectedMap = map[*float32]*uint{}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*uint{}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*uint{}
	actualMap = ZipFloat32UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint64{}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*float32]*uint64{}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint64{}

	expectedMap = map[*float32]*uint64{}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*uint64{}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*uint64{}
	actualMap = ZipFloat32Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint32{}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*float32]*uint32{}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint32{}

	expectedMap = map[*float32]*uint32{}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*uint32{}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*uint32{}
	actualMap = ZipFloat32Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint16Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint16{}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*float32]*uint16{}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint16{}

	expectedMap = map[*float32]*uint16{}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*uint16{}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*uint16{}
	actualMap = ZipFloat32Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint8Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*uint8{}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*float32]*uint8{}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*uint8{}

	expectedMap = map[*float32]*uint8{}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*uint8{}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*uint8{}
	actualMap = ZipFloat32Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32StrPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*string{}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*float32]*string{}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*string{}

	expectedMap = map[*float32]*string{}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*string{}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*string{}
	actualMap = ZipFloat32StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32BoolPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*float32]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*float32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*float32]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*float32]*bool{}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*float32]*bool{}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*bool{}

	expectedMap = map[*float32]*bool{}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*bool{}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*bool{}
	actualMap = ZipFloat32BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*float32{}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*float32]*float32{}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*float32{}

	expectedMap = map[*float32]*float32{}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*float32{}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*float32{}
	actualMap = ZipFloat32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Float64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*float32{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float32]*float64{}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*float32]*float64{}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = []*float64{}

	expectedMap = map[*float32]*float64{}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float32]*float64{}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float32{}
	list2 = nil

	expectedMap = map[*float32]*float64{}
	actualMap = ZipFloat32Float64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64IntPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*int{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int{&v10, &v20, &v30}

	expectedMap = map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int{}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int{}

	expectedMap = map[*float64]*int{}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int{}

	expectedMap = map[*float64]*int{}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*int{}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*int{}
	actualMap = ZipFloat64IntPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64IntPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*int64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int64{&v10, &v20, &v30}

	expectedMap = map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int64{}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int64{}

	expectedMap = map[*float64]*int64{}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int64{}

	expectedMap = map[*float64]*int64{}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*int64{}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*int64{}
	actualMap = ZipFloat64Int64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*int32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int32{&v10, &v20, &v30}

	expectedMap = map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int32{}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int32{}

	expectedMap = map[*float64]*int32{}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int32{}

	expectedMap = map[*float64]*int32{}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*int32{}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*int32{}
	actualMap = ZipFloat64Int32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int16Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*int16{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int16{&v10, &v20, &v30}

	expectedMap = map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int16{}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int16{}

	expectedMap = map[*float64]*int16{}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int16{}

	expectedMap = map[*float64]*int16{}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*int16{}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*int16{}
	actualMap = ZipFloat64Int16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int8Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*int8{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int8{&v10, &v20, &v30}

	expectedMap = map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*int8{}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*int8{}

	expectedMap = map[*float64]*int8{}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*int8{}

	expectedMap = map[*float64]*int8{}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*int8{}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*int8{}
	actualMap = ZipFloat64Int8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64UintPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*uint{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint{&v10, &v20, &v30}

	expectedMap = map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint{}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint{}

	expectedMap = map[*float64]*uint{}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint{}

	expectedMap = map[*float64]*uint{}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*uint{}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*uint{}
	actualMap = ZipFloat64UintPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64UintPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*uint64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{&v10, &v20, &v30}

	expectedMap = map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint64{}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint64{}

	expectedMap = map[*float64]*uint64{}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint64{}

	expectedMap = map[*float64]*uint64{}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*uint64{}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*uint64{}
	actualMap = ZipFloat64Uint64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*uint32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{&v10, &v20, &v30}

	expectedMap = map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint32{}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint32{}

	expectedMap = map[*float64]*uint32{}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint32{}

	expectedMap = map[*float64]*uint32{}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*uint32{}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*uint32{}
	actualMap = ZipFloat64Uint32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint16Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*uint16{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{&v10, &v20, &v30}

	expectedMap = map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint16{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint16{}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint16{}

	expectedMap = map[*float64]*uint16{}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint16{}

	expectedMap = map[*float64]*uint16{}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*uint16{}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*uint16{}
	actualMap = ZipFloat64Uint16Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint8Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*uint8{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{&v10, &v20, &v30}

	expectedMap = map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint8{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*uint8{}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*uint8{}

	expectedMap = map[*float64]*uint8{}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*uint8{}

	expectedMap = map[*float64]*uint8{}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*uint8{}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*uint8{}
	actualMap = ZipFloat64Uint8Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64StrPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*string{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*string{&v10, &v20, &v30}

	expectedMap = map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*string{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*string{}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*string{}

	expectedMap = map[*float64]*string{}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*string{}

	expectedMap = map[*float64]*string{}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*string{}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*string{}
	actualMap = ZipFloat64StrPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64StrPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64BoolPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var vt bool = true
	var vf bool = false

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*bool{&vt, &vt, &vf, &vt}

	expectedMap := map[*float64]*bool{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*bool{&vt, &vt, &vf}

	expectedMap = map[*float64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*bool{&vt, &vt, &vf, &vt}

	expectedMap = map[*float64]*bool{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*bool{&vt, &vt, &vt, &vt}

	expectedMap = map[*float64]*bool{}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*bool{}

	expectedMap = map[*float64]*bool{}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*bool{}

	expectedMap = map[*float64]*bool{}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*bool{}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*bool{}
	actualMap = ZipFloat64BoolPtr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64BoolPtr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Float32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*float32{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*float32{&v10, &v20, &v30}

	expectedMap = map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*float32{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*float32{}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*float32{}

	expectedMap = map[*float64]*float32{}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*float32{}

	expectedMap = map[*float64]*float32{}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*float32{}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*float32{}
	actualMap = ZipFloat64Float32Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40

	list1 := []*float64{&v1, &v2, &v3, &v4}
	list2 := []*float64{&v10, &v20, &v30, &v40}

	expectedMap := map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*float64{&v10, &v20, &v30}

	expectedMap = map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*float64{&v10, &v20, &v30, &v40}

	expectedMap = map[*float64]*float64{}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{&v1, &v2, &v3, &v4}
	list2 = []*float64{}

	expectedMap = map[*float64]*float64{}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = []*float64{}

	expectedMap = map[*float64]*float64{}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*float64]*float64{}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*float64{}
	list2 = nil

	expectedMap = map[*float64]*float64{}
	actualMap = ZipFloat64Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
