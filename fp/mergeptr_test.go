package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMergeIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int{}
	map2 = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*int{}

	expected = map[*int]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int64{}
	map2 = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*int64{}

	expected = map[*int]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int32{}
	map2 = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*int32{}

	expected = map[*int]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt16Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int16{}
	map2 = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*int16{}

	expected = map[*int]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt8Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int8{}
	map2 = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*int8{}

	expected = map[*int]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUintPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint{}
	map2 = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*uint{}

	expected = map[*int]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint64{}
	map2 = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*uint64{}

	expected = map[*int]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint32{}
	map2 = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*uint32{}

	expected = map[*int]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint16Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint16{}
	map2 = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*uint16{}

	expected = map[*int]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint8Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint8{}
	map2 = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*uint8{}

	expected = map[*int]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntStrPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*string{}
	map2 = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*string{}

	expected = map[*int]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntStrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntStrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntBoolPtr(t *testing.T) {
	var v0 int = 0
	var v1 int = 1
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*int]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*int]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*int]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeIntBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*bool{}
	map2 = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeIntBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeIntBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeIntBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*int]*bool{}

	expected = map[*int]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeIntBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntBoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntBoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntFloat32Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float32{}
	map2 = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*float32{}

	expected = map[*int]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntFloat32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntFloat32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntFloat64Ptr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeIntFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float64{}
	map2 = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int]*float64{}

	expected = map[*int]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeIntFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntFloat64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntFloat64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64IntPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int{}
	map2 = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*int{}

	expected = map[*int64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int64{}
	map2 = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*int64{}

	expected = map[*int64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int32{}
	map2 = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*int32{}

	expected = map[*int64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int16Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int16{}
	map2 = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*int16{}

	expected = map[*int64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int8Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int8{}
	map2 = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*int8{}

	expected = map[*int64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64UintPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint{}
	map2 = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*uint{}

	expected = map[*int64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint64{}
	map2 = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*uint64{}

	expected = map[*int64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint32{}
	map2 = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*uint32{}

	expected = map[*int64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint16Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint16{}
	map2 = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*uint16{}

	expected = map[*int64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint8Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint8{}
	map2 = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*uint8{}

	expected = map[*int64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64StrPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*string{}
	map2 = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*string{}

	expected = map[*int64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64BoolPtr(t *testing.T) {
	var v0 int64 = 0
	var v1 int64 = 1
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*int64]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*int64]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*int64]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeInt64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*bool{}
	map2 = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*int64]*bool{}

	expected = map[*int64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Float32Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float32{}
	map2 = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*float32{}

	expected = map[*int64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Float64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int64]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float64{}
	map2 = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int64]*float64{}

	expected = map[*int64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32IntPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int{}
	map2 = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*int{}

	expected = map[*int32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int64{}
	map2 = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*int64{}

	expected = map[*int32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int32{}
	map2 = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*int32{}

	expected = map[*int32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int16Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int16{}
	map2 = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*int16{}

	expected = map[*int32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int8Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int8{}
	map2 = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*int8{}

	expected = map[*int32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32UintPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint{}
	map2 = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*uint{}

	expected = map[*int32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint64{}
	map2 = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*uint64{}

	expected = map[*int32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint32{}
	map2 = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*uint32{}

	expected = map[*int32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint16Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint16{}
	map2 = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*uint16{}

	expected = map[*int32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint8Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint8{}
	map2 = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*uint8{}

	expected = map[*int32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32StrPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*string{}
	map2 = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*string{}

	expected = map[*int32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32BoolPtr(t *testing.T) {
	var v0 int32 = 0
	var v1 int32 = 1
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*int32]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*int32]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*int32]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeInt32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*bool{}
	map2 = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*int32]*bool{}

	expected = map[*int32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Float32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float32{}
	map2 = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*float32{}

	expected = map[*int32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Float64Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int32]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float64{}
	map2 = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int32]*float64{}

	expected = map[*int32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16IntPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int{}
	map2 = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*int{}

	expected = map[*int16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int64{}
	map2 = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*int64{}

	expected = map[*int16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int32{}
	map2 = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*int32{}

	expected = map[*int16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int16{}
	map2 = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*int16{}

	expected = map[*int16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int8Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int8{}
	map2 = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*int8{}

	expected = map[*int16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16UintPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint{}
	map2 = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*uint{}

	expected = map[*int16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint64{}
	map2 = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*uint64{}

	expected = map[*int16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint32{}
	map2 = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*uint32{}

	expected = map[*int16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint16{}
	map2 = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*uint16{}

	expected = map[*int16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint8Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint8{}
	map2 = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*uint8{}

	expected = map[*int16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16StrPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*string{}
	map2 = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*string{}

	expected = map[*int16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16BoolPtr(t *testing.T) {
	var v0 int16 = 0
	var v1 int16 = 1
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*int16]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*int16]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*int16]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeInt16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*bool{}
	map2 = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*int16]*bool{}

	expected = map[*int16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Float32Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float32{}
	map2 = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*float32{}

	expected = map[*int16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Float64Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int16]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float64{}
	map2 = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int16]*float64{}

	expected = map[*int16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8IntPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int{}
	map2 = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*int{}

	expected = map[*int8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int64{}
	map2 = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*int64{}

	expected = map[*int8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int32{}
	map2 = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*int32{}

	expected = map[*int8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int16Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int16{}
	map2 = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*int16{}

	expected = map[*int8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int8{}
	map2 = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*int8{}

	expected = map[*int8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8UintPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint{}
	map2 = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*uint{}

	expected = map[*int8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint64{}
	map2 = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*uint64{}

	expected = map[*int8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint32{}
	map2 = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*uint32{}

	expected = map[*int8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint16Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint16{}
	map2 = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*uint16{}

	expected = map[*int8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint8{}
	map2 = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*uint8{}

	expected = map[*int8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8StrPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*string{}
	map2 = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*string{}

	expected = map[*int8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8BoolPtr(t *testing.T) {
	var v0 int8 = 0
	var v1 int8 = 1
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*int8]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*int8]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*int8]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeInt8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*bool{}
	map2 = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*int8]*bool{}

	expected = map[*int8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeInt8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Float32Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float32{}
	map2 = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*float32{}

	expected = map[*int8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Float64Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*int8]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeInt8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float64{}
	map2 = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*int8]*float64{}

	expected = map[*int8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeInt8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintIntPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int{}
	map2 = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*int{}

	expected = map[*uint]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintIntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintIntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int64{}
	map2 = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*int64{}

	expected = map[*uint]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int32{}
	map2 = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*int32{}

	expected = map[*uint]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt16Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int16{}
	map2 = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*int16{}

	expected = map[*uint]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt8Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int8{}
	map2 = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*int8{}

	expected = map[*uint]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint{}
	map2 = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*uint{}

	expected = map[*uint]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint64{}
	map2 = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*uint64{}

	expected = map[*uint]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint32{}
	map2 = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*uint32{}

	expected = map[*uint]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint16Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint16{}
	map2 = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*uint16{}

	expected = map[*uint]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint8Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint8{}
	map2 = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*uint8{}

	expected = map[*uint]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintStrPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*string{}
	map2 = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*string{}

	expected = map[*uint]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintStrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintStrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintBoolPtr(t *testing.T) {
	var v0 uint = 0
	var v1 uint = 1
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*uint]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*uint]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*uint]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeUintBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*bool{}
	map2 = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUintBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUintBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUintBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*uint]*bool{}

	expected = map[*uint]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUintBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintBoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintBoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintFloat32Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float32{}
	map2 = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*float32{}

	expected = map[*uint]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintFloat32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintFloat32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintFloat64Ptr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUintFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float64{}
	map2 = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint]*float64{}

	expected = map[*uint]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUintFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintFloat64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintFloat64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64IntPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int{}
	map2 = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*int{}

	expected = map[*uint64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int64{}
	map2 = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*int64{}

	expected = map[*uint64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int32{}
	map2 = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*int32{}

	expected = map[*uint64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int16Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int16{}
	map2 = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*int16{}

	expected = map[*uint64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int8Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int8{}
	map2 = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*int8{}

	expected = map[*uint64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64UintPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint{}
	map2 = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*uint{}

	expected = map[*uint64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint64{}
	map2 = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*uint64{}

	expected = map[*uint64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint32{}
	map2 = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*uint32{}

	expected = map[*uint64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint16Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint16{}
	map2 = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*uint16{}

	expected = map[*uint64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint8Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint8{}
	map2 = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*uint8{}

	expected = map[*uint64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64StrPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*string{}
	map2 = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*string{}

	expected = map[*uint64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64BoolPtr(t *testing.T) {
	var v0 uint64 = 0
	var v1 uint64 = 1
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*uint64]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*uint64]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*uint64]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeUint64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*bool{}
	map2 = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*uint64]*bool{}

	expected = map[*uint64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Float32Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float32{}
	map2 = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*float32{}

	expected = map[*uint64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Float64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint64]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float64{}
	map2 = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint64]*float64{}

	expected = map[*uint64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint64Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32IntPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int{}
	map2 = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*int{}

	expected = map[*uint32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int64{}
	map2 = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*int64{}

	expected = map[*uint32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int32{}
	map2 = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*int32{}

	expected = map[*uint32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int16Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int16{}
	map2 = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*int16{}

	expected = map[*uint32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int8Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int8{}
	map2 = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*int8{}

	expected = map[*uint32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32UintPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint{}
	map2 = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*uint{}

	expected = map[*uint32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint64{}
	map2 = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*uint64{}

	expected = map[*uint32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint32{}
	map2 = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*uint32{}

	expected = map[*uint32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint16Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint16{}
	map2 = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*uint16{}

	expected = map[*uint32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint8Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint8{}
	map2 = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*uint8{}

	expected = map[*uint32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32StrPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*string{}
	map2 = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*string{}

	expected = map[*uint32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32BoolPtr(t *testing.T) {
	var v0 uint32 = 0
	var v1 uint32 = 1
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*uint32]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*uint32]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*uint32]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeUint32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*bool{}
	map2 = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*uint32]*bool{}

	expected = map[*uint32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Float32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float32{}
	map2 = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*float32{}

	expected = map[*uint32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Float64Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint32]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float64{}
	map2 = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint32]*float64{}

	expected = map[*uint32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16IntPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int{}
	map2 = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*int{}

	expected = map[*uint16]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int64{}
	map2 = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*int64{}

	expected = map[*uint16]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int32{}
	map2 = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*int32{}

	expected = map[*uint16]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int16{}
	map2 = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*int16{}

	expected = map[*uint16]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int8Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int8{}
	map2 = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*int8{}

	expected = map[*uint16]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16UintPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint{}
	map2 = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*uint{}

	expected = map[*uint16]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint64{}
	map2 = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*uint64{}

	expected = map[*uint16]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint32{}
	map2 = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*uint32{}

	expected = map[*uint16]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint16{}
	map2 = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*uint16{}

	expected = map[*uint16]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint8Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint8{}
	map2 = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*uint8{}

	expected = map[*uint16]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16StrPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*string{}
	map2 = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*string{}

	expected = map[*uint16]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16BoolPtr(t *testing.T) {
	var v0 uint16 = 0
	var v1 uint16 = 1
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*uint16]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*uint16]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*uint16]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeUint16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*bool{}
	map2 = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*uint16]*bool{}

	expected = map[*uint16]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint16BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Float32Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float32{}
	map2 = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*float32{}

	expected = map[*uint16]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Float64Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint16]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float64{}
	map2 = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint16]*float64{}

	expected = map[*uint16]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint16Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8IntPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int{}
	map2 = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*int{}

	expected = map[*uint8]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int64{}
	map2 = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*int64{}

	expected = map[*uint8]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int32{}
	map2 = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*int32{}

	expected = map[*uint8]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int16Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int16{}
	map2 = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*int16{}

	expected = map[*uint8]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int8{}
	map2 = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*int8{}

	expected = map[*uint8]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8UintPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint{}
	map2 = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*uint{}

	expected = map[*uint8]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint64{}
	map2 = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*uint64{}

	expected = map[*uint8]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint32{}
	map2 = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*uint32{}

	expected = map[*uint8]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint16Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint16{}
	map2 = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*uint16{}

	expected = map[*uint8]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint8{}
	map2 = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*uint8{}

	expected = map[*uint8]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8StrPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*string{}
	map2 = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*string{}

	expected = map[*uint8]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8BoolPtr(t *testing.T) {
	var v0 uint8 = 0
	var v1 uint8 = 1
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*uint8]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*uint8]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*uint8]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeUint8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*bool{}
	map2 = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*uint8]*bool{}

	expected = map[*uint8]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeUint8BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Float32Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float32{}
	map2 = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*float32{}

	expected = map[*uint8]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Float64Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*uint8]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeUint8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float64{}
	map2 = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*uint8]*float64{}

	expected = map[*uint8]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeUint8Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrIntPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int{}
	map2 = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*int{}

	expected = map[*string]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrIntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrIntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int64{}
	map2 = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*int64{}

	expected = map[*string]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int32{}
	map2 = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*int32{}

	expected = map[*string]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt16Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int16{}
	map2 = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*int16{}

	expected = map[*string]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt8Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int8{}
	map2 = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*int8{}

	expected = map[*string]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUintPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint{}
	map2 = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*uint{}

	expected = map[*string]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint64{}
	map2 = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*uint64{}

	expected = map[*string]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint32{}
	map2 = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*uint32{}

	expected = map[*string]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint16Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint16{}
	map2 = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*uint16{}

	expected = map[*string]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint8Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint8{}
	map2 = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*uint8{}

	expected = map[*string]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"

	var vOne string = "One"
	var vTwo string = "Two"

	map1 := map[*string]*string{&v1: &vOne, &v2: &vTwo}
	map2 := map[*string]*string{&v2: &vTwo}

	expected := map[*string]*string{&v1: &vOne, &v2: &vTwo}
	actual := MergeStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*string{}
	map2 = map[*string]*string{&v1: &vOne, &v2: &vTwo}

	expected = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	actual = MergeStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*string{&v1: &vOne, &v2: &vTwo}

	expected = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	actual = MergeStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	map2 = nil

	expected = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	actual = MergeStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	map2 = map[*string]*string{}

	expected = map[*string]*string{&v1: &vOne, &v2: &vTwo}
	actual = MergeStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrBoolPtr(t *testing.T) {
	var v0 string = "0"
	var v1 string = "1"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var vt bool = true
	var vf bool = false

	map1 := map[*string]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected := map[*string]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vt, &v3: &vt}
	actual := MergeStrBoolPtr(map1, map2)

	if *expected[&v0] != *actual[&v0] {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", *expected[&v0], *actual[&v0])
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*bool{}
	map2 = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeStrBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeStrBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeStrBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*string]*bool{}

	expected = map[*string]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeStrBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrBoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrBoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrFloat32Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float32{}
	map2 = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*float32{}

	expected = map[*string]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrFloat32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrFloat32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrFloat64Ptr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*string]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeStrFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrPtrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float64{}
	map2 = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*string]*float64{}

	expected = map[*string]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeStrFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrFloat64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrFloat64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolIntPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 int = 1
	var v0 int = 2

	map1 := map[*bool]*int{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*int{&vt: &v1}

	expected := map[*bool]*int{&vt: &v1, &vf: &v0}
	actual := MergeBoolIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int{}
	map2 = map[*bool]*int{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int{&vt: &v1, &vf: &v0}
	actual = MergeBoolIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*int{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int{&vt: &v1, &vf: &v0}
	actual = MergeBoolIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*int{&vt: &v1, &vf: &v0}
	actual = MergeBoolIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*int{}

	expected = map[*bool]*int{&vt: &v1, &vf: &v0}
	actual = MergeBoolIntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolIntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolIntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 int64 = 1
	var v0 int64 = 2

	map1 := map[*bool]*int64{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*int64{&vt: &v1}

	expected := map[*bool]*int64{&vt: &v1, &vf: &v0}
	actual := MergeBoolInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int64{}
	map2 = map[*bool]*int64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int64{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*int64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int64{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int64{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*int64{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int64{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*int64{}

	expected = map[*bool]*int64{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 int32 = 1
	var v0 int32 = 2

	map1 := map[*bool]*int32{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*int32{&vt: &v1}

	expected := map[*bool]*int32{&vt: &v1, &vf: &v0}
	actual := MergeBoolInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int32{}
	map2 = map[*bool]*int32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int32{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*int32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int32{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int32{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*int32{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int32{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*int32{}

	expected = map[*bool]*int32{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt16Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 int16 = 1
	var v0 int16 = 2

	map1 := map[*bool]*int16{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*int16{&vt: &v1}

	expected := map[*bool]*int16{&vt: &v1, &vf: &v0}
	actual := MergeBoolInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int16{}
	map2 = map[*bool]*int16{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int16{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*int16{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int16{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int16{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*int16{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int16{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*int16{}

	expected = map[*bool]*int16{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt8Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 int8 = 1
	var v0 int8 = 2

	map1 := map[*bool]*int8{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*int8{&vt: &v1}

	expected := map[*bool]*int8{&vt: &v1, &vf: &v0}
	actual := MergeBoolInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int8{}
	map2 = map[*bool]*int8{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int8{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*int8{&vt: &v1, &vf: &v0}

	expected = map[*bool]*int8{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int8{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*int8{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*int8{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*int8{}

	expected = map[*bool]*int8{&vt: &v1, &vf: &v0}
	actual = MergeBoolInt8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUintPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 uint = 1
	var v0 uint = 2

	map1 := map[*bool]*uint{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*uint{&vt: &v1}

	expected := map[*bool]*uint{&vt: &v1, &vf: &v0}
	actual := MergeBoolUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint{}
	map2 = map[*bool]*uint{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint{&vt: &v1, &vf: &v0}
	actual = MergeBoolUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*uint{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint{&vt: &v1, &vf: &v0}
	actual = MergeBoolUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*uint{&vt: &v1, &vf: &v0}
	actual = MergeBoolUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*uint{}

	expected = map[*bool]*uint{&vt: &v1, &vf: &v0}
	actual = MergeBoolUintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 uint64 = 1
	var v0 uint64 = 2

	map1 := map[*bool]*uint64{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*uint64{&vt: &v1}

	expected := map[*bool]*uint64{&vt: &v1, &vf: &v0}
	actual := MergeBoolUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint64{}
	map2 = map[*bool]*uint64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*uint64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*uint64{}

	expected = map[*bool]*uint64{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 uint32 = 1
	var v0 uint32 = 2

	map1 := map[*bool]*uint32{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*uint32{&vt: &v1}

	expected := map[*bool]*uint32{&vt: &v1, &vf: &v0}
	actual := MergeBoolUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint32{}
	map2 = map[*bool]*uint32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*uint32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*uint32{}

	expected = map[*bool]*uint32{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint16Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 uint16 = 1
	var v0 uint16 = 2

	map1 := map[*bool]*uint16{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*uint16{&vt: &v1}

	expected := map[*bool]*uint16{&vt: &v1, &vf: &v0}
	actual := MergeBoolUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint16{}
	map2 = map[*bool]*uint16{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*uint16{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*uint16{}

	expected = map[*bool]*uint16{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint8Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 uint8 = 1
	var v0 uint8 = 2

	map1 := map[*bool]*uint8{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*uint8{&vt: &v1}

	expected := map[*bool]*uint8{&vt: &v1, &vf: &v0}
	actual := MergeBoolUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint8{}
	map2 = map[*bool]*uint8{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*uint8{&vt: &v1, &vf: &v0}

	expected = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*uint8{}

	expected = map[*bool]*uint8{&vt: &v1, &vf: &v0}
	actual = MergeBoolUint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolStrPtr(t *testing.T) {

	var v0 string = "0"
	var v1 string = "1"
	var v2 string = "2"

	var vt bool = true
	var vf bool = false

	map1 := map[*bool]*string{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*string{&vt: &v2}

	expected := map[*bool]*string{&vt: &v2, &vf: &v0}
	actual := MergeBoolStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*string{}
	map2 = map[*bool]*string{&vt: &v1, &vf: &v0}

	expected = map[*bool]*string{&vt: &v1, &vf: &v0}
	actual = MergeBoolStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*string{&vt: &v1, &vf: &v0}

	expected = map[*bool]*string{&vt: &v1, &vf: &v0}
	actual = MergeBoolStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*string{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*string{&vt: &v1, &vf: &v0}
	actual = MergeBoolStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*string{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*string{}

	expected = map[*bool]*string{&vt: &v1, &vf: &v0}
	actual = MergeBoolStrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolStrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolStrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolPtr(t *testing.T) {

	var vt bool = true
	var vf bool = false

	map1 := map[*bool]*bool{&vt: &vt, &vf: &vf}
	map2 := map[*bool]*bool{&vt: &vt}

	expected := map[*bool]*bool{&vt: &vt, &vf: &vf}
	actual := MergeBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*bool{}
	map2 = map[*bool]*bool{&vt: &vt, &vf: &vf}

	expected = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actual = MergeBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*bool{&vt: &vt, &vf: &vf}

	expected = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actual = MergeBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*bool{&vt: &vt, &vf: &vf}
	map2 = nil

	expected = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actual = MergeBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*bool{&vt: &vt, &vf: &vf}
	map2 = map[*bool]*bool{}

	expected = map[*bool]*bool{&vt: &vt, &vf: &vf}
	actual = MergeBoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolFloat32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 float32 = 1
	var v0 float32 = 2

	map1 := map[*bool]*float32{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*float32{&vt: &v1}

	expected := map[*bool]*float32{&vt: &v1, &vf: &v0}
	actual := MergeBoolFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float32{}
	map2 = map[*bool]*float32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*float32{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*float32{&vt: &v1, &vf: &v0}

	expected = map[*bool]*float32{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float32{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*float32{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float32{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*float32{}

	expected = map[*bool]*float32{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolFloat32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolFloat32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolFloat64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v1 float64 = 1
	var v0 float64 = 2

	map1 := map[*bool]*float64{&vt: &v1, &vf: &v0}
	map2 := map[*bool]*float64{&vt: &v1}

	expected := map[*bool]*float64{&vt: &v1, &vf: &v0}
	actual := MergeBoolFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float64{}
	map2 = map[*bool]*float64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*float64{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*bool]*float64{&vt: &v1, &vf: &v0}

	expected = map[*bool]*float64{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float64{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*bool]*float64{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*bool]*float64{&vt: &v1, &vf: &v0}
	map2 = map[*bool]*float64{}

	expected = map[*bool]*float64{&vt: &v1, &vf: &v0}
	actual = MergeBoolFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolFloat64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolFloat64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32IntPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int{}
	map2 = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*int{}

	expected = map[*float32]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int64{}
	map2 = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*int64{}

	expected = map[*float32]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int32{}
	map2 = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*int32{}

	expected = map[*float32]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int16Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int16{}
	map2 = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*int16{}

	expected = map[*float32]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int8Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int8{}
	map2 = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*int8{}

	expected = map[*float32]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32UintPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint{}
	map2 = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*uint{}

	expected = map[*float32]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint64{}
	map2 = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*uint64{}

	expected = map[*float32]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint32{}
	map2 = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*uint32{}

	expected = map[*float32]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint16Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint16{}
	map2 = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*uint16{}

	expected = map[*float32]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint8Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint8{}
	map2 = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*uint8{}

	expected = map[*float32]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32StrPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*string{}
	map2 = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*string{}

	expected = map[*float32]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32BoolPtr(t *testing.T) {
	var v0 float32 = 0
	var v1 float32 = 1
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*float32]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*float32]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*float32]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeFloat32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*bool{}
	map2 = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*float32]*bool{}

	expected = map[*float32]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat32BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float32{}
	map2 = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*float32{}

	expected = map[*float32]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Float64Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float32]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float64{}
	map2 = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float32]*float64{}

	expected = map[*float32]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat32Float64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Float64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Float64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64IntPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30
	var v40 int = 40
	var v50 int = 50

	map1 := map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*int{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int{}
	map2 = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*int{}

	expected = map[*float64]*int{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64IntPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64IntPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64IntPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64IntPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30
	var v40 int64 = 40
	var v50 int64 = 50

	map1 := map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*int64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int64{}
	map2 = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*int64{}

	expected = map[*float64]*int64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30
	var v40 int32 = 40
	var v50 int32 = 50

	map1 := map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*int32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int32{}
	map2 = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*int32{}

	expected = map[*float64]*int32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int16Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30
	var v40 int16 = 40
	var v50 int16 = 50

	map1 := map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*int16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int16{}
	map2 = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*int16{}

	expected = map[*float64]*int16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int8Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30
	var v40 int8 = 40
	var v50 int8 = 50

	map1 := map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*int8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int8{}
	map2 = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*int8{}

	expected = map[*float64]*int8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Int8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64UintPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30
	var v40 uint = 40
	var v50 uint = 50

	map1 := map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*uint{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint{}
	map2 = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*uint{}

	expected = map[*float64]*uint{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64UintPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64UintPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64UintPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64UintPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30
	var v40 uint64 = 40
	var v50 uint64 = 50

	map1 := map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*uint64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint64{}
	map2 = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*uint64{}

	expected = map[*float64]*uint64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30
	var v40 uint32 = 40
	var v50 uint32 = 50

	map1 := map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*uint32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint32{}
	map2 = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*uint32{}

	expected = map[*float64]*uint32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint16Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30
	var v40 uint16 = 40
	var v50 uint16 = 50

	map1 := map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*uint16{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint16{}
	map2 = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*uint16{}

	expected = map[*float64]*uint16{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint16Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint16Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint16Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint8Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30
	var v40 uint8 = 40
	var v50 uint8 = 50

	map1 := map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*uint8{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint8{}
	map2 = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*uint8{}

	expected = map[*float64]*uint8{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Uint8Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint8Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint8Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64StrPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"
	var v40 string = "40"
	var v50 string = "50"
	map1 := map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*string{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*string{}
	map2 = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*string{}

	expected = map[*float64]*string{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64StrPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64StrPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64StrPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64BoolPtr(t *testing.T) {
	var v0 float64 = 0
	var v1 float64 = 1
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vt bool = true
	var vf bool = false
	

	map1 := map[*float64]*bool{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*float64]*bool{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*float64]*bool{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := MergeFloat64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*bool{}
	map2 = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*float64]*bool{}

	expected = map[*float64]*bool{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = MergeFloat64BoolPtr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64BoolPtr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64BoolPtr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Float32Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30
	var v40 float32 = 40
	var v50 float32 = 50

	map1 := map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*float32{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float32{}
	map2 = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*float32{}

	expected = map[*float64]*float32{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Float32Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Float32Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Float32Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30
	var v40 float64 = 40
	var v50 float64 = 50

	map1 := map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*float64]*float64{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := MergeFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float64{}
	map2 = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*float64]*float64{}

	expected = map[*float64]*float64{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = MergeFloat64Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
