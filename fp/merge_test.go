package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMergeInt(t *testing.T) {
	map1 := map[int]int{1: 10, 2: 20, 3: 30}
	map2 := map[int]int{4: 40, 5: 50, 3: 30}

	expected := map[int]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int{}
	map2 = map[int]int{4: 40, 5: 50, 3: 30}

	expected = map[int]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]int{4: 40, 5: 50, 3: 30}

	expected = map[int]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int{4: 40, 5: 50, 3: 30}
	map2 = map[int]int{}

	expected = map[int]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt64(t *testing.T) {
	map1 := map[int]int64{1: 10, 2: 20, 3: 30}
	map2 := map[int]int64{4: 40, 5: 50, 3: 30}

	expected := map[int]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int64{}
	map2 = map[int]int64{4: 40, 5: 50, 3: 30}

	expected = map[int]int64{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]int64{4: 40, 5: 50, 3: 30}

	expected = map[int]int64{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]int64{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int64{4: 40, 5: 50, 3: 30}
	map2 = map[int]int64{}

	expected = map[int]int64{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt32(t *testing.T) {
	map1 := map[int]int32{1: 10, 2: 20, 3: 30}
	map2 := map[int]int32{4: 40, 5: 50, 3: 30}

	expected := map[int]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int32{}
	map2 = map[int]int32{4: 40, 5: 50, 3: 30}

	expected = map[int]int32{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]int32{4: 40, 5: 50, 3: 30}

	expected = map[int]int32{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]int32{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int32{4: 40, 5: 50, 3: 30}
	map2 = map[int]int32{}

	expected = map[int]int32{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt16(t *testing.T) {
	map1 := map[int]int16{1: 10, 2: 20, 3: 30}
	map2 := map[int]int16{4: 40, 5: 50, 3: 30}

	expected := map[int]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int16{}
	map2 = map[int]int16{4: 40, 5: 50, 3: 30}

	expected = map[int]int16{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]int16{4: 40, 5: 50, 3: 30}

	expected = map[int]int16{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]int16{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int16{4: 40, 5: 50, 3: 30}
	map2 = map[int]int16{}

	expected = map[int]int16{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntInt8(t *testing.T) {
	map1 := map[int]int8{1: 10, 2: 20, 3: 30}
	map2 := map[int]int8{4: 40, 5: 50, 3: 30}

	expected := map[int]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int8{}
	map2 = map[int]int8{4: 40, 5: 50, 3: 30}

	expected = map[int]int8{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]int8{4: 40, 5: 50, 3: 30}

	expected = map[int]int8{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]int8{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]int8{4: 40, 5: 50, 3: 30}
	map2 = map[int]int8{}

	expected = map[int]int8{4: 40, 5: 50, 3: 30}
	actual = MergeIntInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntInt8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntInt8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint(t *testing.T) {
	map1 := map[int]uint{1: 10, 2: 20, 3: 30}
	map2 := map[int]uint{4: 40, 5: 50, 3: 30}

	expected := map[int]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint{}
	map2 = map[int]uint{4: 40, 5: 50, 3: 30}

	expected = map[int]uint{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]uint{4: 40, 5: 50, 3: 30}

	expected = map[int]uint{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]uint{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint{4: 40, 5: 50, 3: 30}
	map2 = map[int]uint{}

	expected = map[int]uint{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint64(t *testing.T) {
	map1 := map[int]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[int]uint64{4: 40, 5: 50, 3: 30}

	expected := map[int]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint64{}
	map2 = map[int]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[int]uint64{}

	expected = map[int]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint32(t *testing.T) {
	map1 := map[int]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[int]uint32{4: 40, 5: 50, 3: 30}

	expected := map[int]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint32{}
	map2 = map[int]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[int]uint32{}

	expected = map[int]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint16(t *testing.T) {
	map1 := map[int]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[int]uint16{4: 40, 5: 50, 3: 30}

	expected := map[int]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint16{}
	map2 = map[int]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[int]uint16{}

	expected = map[int]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntUint8(t *testing.T) {
	map1 := map[int]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[int]uint8{4: 40, 5: 50, 3: 30}

	expected := map[int]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint8{}
	map2 = map[int]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[int]uint8{}

	expected = map[int]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeIntUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntUint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntUint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntStr(t *testing.T) {
	map1 := map[int]string{1: "10", 2: "20", 3: "30"}
	map2 := map[int]string{4: "40", 5: "50", 3: "30"}

	expected := map[int]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeIntStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]string{}
	map2 = map[int]string{4: "40", 5: "50", 3: "30"}

	expected = map[int]string{4: "40", 5: "50", 3: "30"}
	actual = MergeIntStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]string{4: "40", 5: "50", 3: "30"}

	expected = map[int]string{4: "40", 5: "50", 3: "30"}
	actual = MergeIntStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[int]string{4: "40", 5: "50", 3: "30"}
	actual = MergeIntStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]string{4: "40", 5: "50", 3: "30"}
	map2 = map[int]string{}

	expected = map[int]string{4: "40", 5: "50", 3: "30"}
	actual = MergeIntStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntStr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntStr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntBool(t *testing.T) {
	map1 := map[int]bool{1: true, 0: false, 3: true}
	map2 := map[int]bool{4: true, 5: true, 3: true}

	expected := map[int]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeIntBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]bool{}
	map2 = map[int]bool{4: true, 5: true, 3: true}

	expected = map[int]bool{4: true, 5: true, 3: true}
	actual = MergeIntBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]bool{4: true, 5: true, 3: true}

	expected = map[int]bool{4: true, 5: true, 3: true}
	actual = MergeIntBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[int]bool{4: true, 5: true, 3: true}
	actual = MergeIntBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]bool{4: true, 5: true, 3: true}
	map2 = map[int]bool{}

	expected = map[int]bool{4: true, 5: true, 3: true}
	actual = MergeIntBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntBool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntBool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntFloat32(t *testing.T) {
	map1 := map[int]float32{1: 10, 2: 20, 3: 30}
	map2 := map[int]float32{4: 40, 5: 50, 3: 30}

	expected := map[int]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float32{}
	map2 = map[int]float32{4: 40, 5: 50, 3: 30}

	expected = map[int]float32{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]float32{4: 40, 5: 50, 3: 30}

	expected = map[int]float32{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]float32{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float32{4: 40, 5: 50, 3: 30}
	map2 = map[int]float32{}

	expected = map[int]float32{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntFloat32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntFloat32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeIntFloat64(t *testing.T) {
	map1 := map[int]float64{1: 10, 2: 20, 3: 30}
	map2 := map[int]float64{4: 40, 5: 50, 3: 30}

	expected := map[int]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeIntFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float64{}
	map2 = map[int]float64{4: 40, 5: 50, 3: 30}

	expected = map[int]float64{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int]float64{4: 40, 5: 50, 3: 30}

	expected = map[int]float64{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int]float64{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int]float64{4: 40, 5: 50, 3: 30}
	map2 = map[int]float64{}

	expected = map[int]float64{4: 40, 5: 50, 3: 30}
	actual = MergeIntFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeIntFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeIntFloat64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeIntFloat64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int(t *testing.T) {
	map1 := map[int64]int{1: 10, 2: 20, 3: 30}
	map2 := map[int64]int{4: 40, 5: 50, 3: 30}

	expected := map[int64]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int{}
	map2 = map[int64]int{4: 40, 5: 50, 3: 30}

	expected = map[int64]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]int{4: 40, 5: 50, 3: 30}

	expected = map[int64]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int{4: 40, 5: 50, 3: 30}
	map2 = map[int64]int{}

	expected = map[int64]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64(t *testing.T) {
	map1 := map[int64]int64{1: 10, 2: 20, 3: 30}
	map2 := map[int64]int64{4: 40, 5: 50, 3: 30}

	expected := map[int64]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int64{}
	map2 = map[int64]int64{4: 40, 5: 50, 3: 30}

	expected = map[int64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]int64{4: 40, 5: 50, 3: 30}

	expected = map[int64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int64{4: 40, 5: 50, 3: 30}
	map2 = map[int64]int64{}

	expected = map[int64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int32(t *testing.T) {
	map1 := map[int64]int32{1: 10, 2: 20, 3: 30}
	map2 := map[int64]int32{4: 40, 5: 50, 3: 30}

	expected := map[int64]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int32{}
	map2 = map[int64]int32{4: 40, 5: 50, 3: 30}

	expected = map[int64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]int32{4: 40, 5: 50, 3: 30}

	expected = map[int64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int32{4: 40, 5: 50, 3: 30}
	map2 = map[int64]int32{}

	expected = map[int64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int16(t *testing.T) {
	map1 := map[int64]int16{1: 10, 2: 20, 3: 30}
	map2 := map[int64]int16{4: 40, 5: 50, 3: 30}

	expected := map[int64]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int16{}
	map2 = map[int64]int16{4: 40, 5: 50, 3: 30}

	expected = map[int64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]int16{4: 40, 5: 50, 3: 30}

	expected = map[int64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int16{4: 40, 5: 50, 3: 30}
	map2 = map[int64]int16{}

	expected = map[int64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Int8(t *testing.T) {
	map1 := map[int64]int8{1: 10, 2: 20, 3: 30}
	map2 := map[int64]int8{4: 40, 5: 50, 3: 30}

	expected := map[int64]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int8{}
	map2 = map[int64]int8{4: 40, 5: 50, 3: 30}

	expected = map[int64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]int8{4: 40, 5: 50, 3: 30}

	expected = map[int64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]int8{4: 40, 5: 50, 3: 30}
	map2 = map[int64]int8{}

	expected = map[int64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint(t *testing.T) {
	map1 := map[int64]uint{1: 10, 2: 20, 3: 30}
	map2 := map[int64]uint{4: 40, 5: 50, 3: 30}

	expected := map[int64]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint{}
	map2 = map[int64]uint{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]uint{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint{4: 40, 5: 50, 3: 30}
	map2 = map[int64]uint{}

	expected = map[int64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint64(t *testing.T) {
	map1 := map[int64]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[int64]uint64{4: 40, 5: 50, 3: 30}

	expected := map[int64]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint64{}
	map2 = map[int64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[int64]uint64{}

	expected = map[int64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint32(t *testing.T) {
	map1 := map[int64]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[int64]uint32{4: 40, 5: 50, 3: 30}

	expected := map[int64]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint32{}
	map2 = map[int64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[int64]uint32{}

	expected = map[int64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint16(t *testing.T) {
	map1 := map[int64]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[int64]uint16{4: 40, 5: 50, 3: 30}

	expected := map[int64]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint16{}
	map2 = map[int64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[int64]uint16{}

	expected = map[int64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Uint8(t *testing.T) {
	map1 := map[int64]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[int64]uint8{4: 40, 5: 50, 3: 30}

	expected := map[int64]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint8{}
	map2 = map[int64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[int64]uint8{}

	expected = map[int64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Str(t *testing.T) {
	map1 := map[int64]string{1: "10", 2: "20", 3: "30"}
	map2 := map[int64]string{4: "40", 5: "50", 3: "30"}

	expected := map[int64]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeInt64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]string{}
	map2 = map[int64]string{4: "40", 5: "50", 3: "30"}

	expected = map[int64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]string{4: "40", 5: "50", 3: "30"}

	expected = map[int64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[int64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]string{4: "40", 5: "50", 3: "30"}
	map2 = map[int64]string{}

	expected = map[int64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Bool(t *testing.T) {
	map1 := map[int64]bool{1: true, 0: false, 3: true}
	map2 := map[int64]bool{4: true, 5: true, 3: true}

	expected := map[int64]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeInt64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]bool{}
	map2 = map[int64]bool{4: true, 5: true, 3: true}

	expected = map[int64]bool{4: true, 5: true, 3: true}
	actual = MergeInt64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]bool{4: true, 5: true, 3: true}

	expected = map[int64]bool{4: true, 5: true, 3: true}
	actual = MergeInt64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[int64]bool{4: true, 5: true, 3: true}
	actual = MergeInt64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]bool{4: true, 5: true, 3: true}
	map2 = map[int64]bool{}

	expected = map[int64]bool{4: true, 5: true, 3: true}
	actual = MergeInt64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Float32(t *testing.T) {
	map1 := map[int64]float32{1: 10, 2: 20, 3: 30}
	map2 := map[int64]float32{4: 40, 5: 50, 3: 30}

	expected := map[int64]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float32{}
	map2 = map[int64]float32{4: 40, 5: 50, 3: 30}

	expected = map[int64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]float32{4: 40, 5: 50, 3: 30}

	expected = map[int64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float32{4: 40, 5: 50, 3: 30}
	map2 = map[int64]float32{}

	expected = map[int64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt64Float64(t *testing.T) {
	map1 := map[int64]float64{1: 10, 2: 20, 3: 30}
	map2 := map[int64]float64{4: 40, 5: 50, 3: 30}

	expected := map[int64]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float64{}
	map2 = map[int64]float64{4: 40, 5: 50, 3: 30}

	expected = map[int64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int64]float64{4: 40, 5: 50, 3: 30}

	expected = map[int64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int64]float64{4: 40, 5: 50, 3: 30}
	map2 = map[int64]float64{}

	expected = map[int64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt64Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt64Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int(t *testing.T) {
	map1 := map[int32]int{1: 10, 2: 20, 3: 30}
	map2 := map[int32]int{4: 40, 5: 50, 3: 30}

	expected := map[int32]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int{}
	map2 = map[int32]int{4: 40, 5: 50, 3: 30}

	expected = map[int32]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]int{4: 40, 5: 50, 3: 30}

	expected = map[int32]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int{4: 40, 5: 50, 3: 30}
	map2 = map[int32]int{}

	expected = map[int32]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int64(t *testing.T) {
	map1 := map[int32]int64{1: 10, 2: 20, 3: 30}
	map2 := map[int32]int64{4: 40, 5: 50, 3: 30}

	expected := map[int32]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int64{}
	map2 = map[int32]int64{4: 40, 5: 50, 3: 30}

	expected = map[int32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]int64{4: 40, 5: 50, 3: 30}

	expected = map[int32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int64{4: 40, 5: 50, 3: 30}
	map2 = map[int32]int64{}

	expected = map[int32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32(t *testing.T) {
	map1 := map[int32]int32{1: 10, 2: 20, 3: 30}
	map2 := map[int32]int32{4: 40, 5: 50, 3: 30}

	expected := map[int32]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int32{}
	map2 = map[int32]int32{4: 40, 5: 50, 3: 30}

	expected = map[int32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]int32{4: 40, 5: 50, 3: 30}

	expected = map[int32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int32{4: 40, 5: 50, 3: 30}
	map2 = map[int32]int32{}

	expected = map[int32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int16(t *testing.T) {
	map1 := map[int32]int16{1: 10, 2: 20, 3: 30}
	map2 := map[int32]int16{4: 40, 5: 50, 3: 30}

	expected := map[int32]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int16{}
	map2 = map[int32]int16{4: 40, 5: 50, 3: 30}

	expected = map[int32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]int16{4: 40, 5: 50, 3: 30}

	expected = map[int32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int16{4: 40, 5: 50, 3: 30}
	map2 = map[int32]int16{}

	expected = map[int32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Int8(t *testing.T) {
	map1 := map[int32]int8{1: 10, 2: 20, 3: 30}
	map2 := map[int32]int8{4: 40, 5: 50, 3: 30}

	expected := map[int32]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int8{}
	map2 = map[int32]int8{4: 40, 5: 50, 3: 30}

	expected = map[int32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]int8{4: 40, 5: 50, 3: 30}

	expected = map[int32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]int8{4: 40, 5: 50, 3: 30}
	map2 = map[int32]int8{}

	expected = map[int32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint(t *testing.T) {
	map1 := map[int32]uint{1: 10, 2: 20, 3: 30}
	map2 := map[int32]uint{4: 40, 5: 50, 3: 30}

	expected := map[int32]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint{}
	map2 = map[int32]uint{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]uint{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint{4: 40, 5: 50, 3: 30}
	map2 = map[int32]uint{}

	expected = map[int32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint64(t *testing.T) {
	map1 := map[int32]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[int32]uint64{4: 40, 5: 50, 3: 30}

	expected := map[int32]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint64{}
	map2 = map[int32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[int32]uint64{}

	expected = map[int32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint32(t *testing.T) {
	map1 := map[int32]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[int32]uint32{4: 40, 5: 50, 3: 30}

	expected := map[int32]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint32{}
	map2 = map[int32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[int32]uint32{}

	expected = map[int32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint16(t *testing.T) {
	map1 := map[int32]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[int32]uint16{4: 40, 5: 50, 3: 30}

	expected := map[int32]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint16{}
	map2 = map[int32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[int32]uint16{}

	expected = map[int32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Uint8(t *testing.T) {
	map1 := map[int32]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[int32]uint8{4: 40, 5: 50, 3: 30}

	expected := map[int32]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint8{}
	map2 = map[int32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[int32]uint8{}

	expected = map[int32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Str(t *testing.T) {
	map1 := map[int32]string{1: "10", 2: "20", 3: "30"}
	map2 := map[int32]string{4: "40", 5: "50", 3: "30"}

	expected := map[int32]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeInt32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]string{}
	map2 = map[int32]string{4: "40", 5: "50", 3: "30"}

	expected = map[int32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]string{4: "40", 5: "50", 3: "30"}

	expected = map[int32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[int32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]string{4: "40", 5: "50", 3: "30"}
	map2 = map[int32]string{}

	expected = map[int32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Bool(t *testing.T) {
	map1 := map[int32]bool{1: true, 0: false, 3: true}
	map2 := map[int32]bool{4: true, 5: true, 3: true}

	expected := map[int32]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeInt32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]bool{}
	map2 = map[int32]bool{4: true, 5: true, 3: true}

	expected = map[int32]bool{4: true, 5: true, 3: true}
	actual = MergeInt32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]bool{4: true, 5: true, 3: true}

	expected = map[int32]bool{4: true, 5: true, 3: true}
	actual = MergeInt32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[int32]bool{4: true, 5: true, 3: true}
	actual = MergeInt32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]bool{4: true, 5: true, 3: true}
	map2 = map[int32]bool{}

	expected = map[int32]bool{4: true, 5: true, 3: true}
	actual = MergeInt32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Float32(t *testing.T) {
	map1 := map[int32]float32{1: 10, 2: 20, 3: 30}
	map2 := map[int32]float32{4: 40, 5: 50, 3: 30}

	expected := map[int32]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float32{}
	map2 = map[int32]float32{4: 40, 5: 50, 3: 30}

	expected = map[int32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]float32{4: 40, 5: 50, 3: 30}

	expected = map[int32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float32{4: 40, 5: 50, 3: 30}
	map2 = map[int32]float32{}

	expected = map[int32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt32Float64(t *testing.T) {
	map1 := map[int32]float64{1: 10, 2: 20, 3: 30}
	map2 := map[int32]float64{4: 40, 5: 50, 3: 30}

	expected := map[int32]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float64{}
	map2 = map[int32]float64{4: 40, 5: 50, 3: 30}

	expected = map[int32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int32]float64{4: 40, 5: 50, 3: 30}

	expected = map[int32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int32]float64{4: 40, 5: 50, 3: 30}
	map2 = map[int32]float64{}

	expected = map[int32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt32Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt32Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int(t *testing.T) {
	map1 := map[int16]int{1: 10, 2: 20, 3: 30}
	map2 := map[int16]int{4: 40, 5: 50, 3: 30}

	expected := map[int16]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int{}
	map2 = map[int16]int{4: 40, 5: 50, 3: 30}

	expected = map[int16]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]int{4: 40, 5: 50, 3: 30}

	expected = map[int16]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int{4: 40, 5: 50, 3: 30}
	map2 = map[int16]int{}

	expected = map[int16]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int64(t *testing.T) {
	map1 := map[int16]int64{1: 10, 2: 20, 3: 30}
	map2 := map[int16]int64{4: 40, 5: 50, 3: 30}

	expected := map[int16]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int64{}
	map2 = map[int16]int64{4: 40, 5: 50, 3: 30}

	expected = map[int16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]int64{4: 40, 5: 50, 3: 30}

	expected = map[int16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int64{4: 40, 5: 50, 3: 30}
	map2 = map[int16]int64{}

	expected = map[int16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int32(t *testing.T) {
	map1 := map[int16]int32{1: 10, 2: 20, 3: 30}
	map2 := map[int16]int32{4: 40, 5: 50, 3: 30}

	expected := map[int16]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int32{}
	map2 = map[int16]int32{4: 40, 5: 50, 3: 30}

	expected = map[int16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]int32{4: 40, 5: 50, 3: 30}

	expected = map[int16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int32{4: 40, 5: 50, 3: 30}
	map2 = map[int16]int32{}

	expected = map[int16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16(t *testing.T) {
	map1 := map[int16]int16{1: 10, 2: 20, 3: 30}
	map2 := map[int16]int16{4: 40, 5: 50, 3: 30}

	expected := map[int16]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int16{}
	map2 = map[int16]int16{4: 40, 5: 50, 3: 30}

	expected = map[int16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]int16{4: 40, 5: 50, 3: 30}

	expected = map[int16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int16{4: 40, 5: 50, 3: 30}
	map2 = map[int16]int16{}

	expected = map[int16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Int8(t *testing.T) {
	map1 := map[int16]int8{1: 10, 2: 20, 3: 30}
	map2 := map[int16]int8{4: 40, 5: 50, 3: 30}

	expected := map[int16]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int8{}
	map2 = map[int16]int8{4: 40, 5: 50, 3: 30}

	expected = map[int16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]int8{4: 40, 5: 50, 3: 30}

	expected = map[int16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]int8{4: 40, 5: 50, 3: 30}
	map2 = map[int16]int8{}

	expected = map[int16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint(t *testing.T) {
	map1 := map[int16]uint{1: 10, 2: 20, 3: 30}
	map2 := map[int16]uint{4: 40, 5: 50, 3: 30}

	expected := map[int16]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint{}
	map2 = map[int16]uint{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]uint{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint{4: 40, 5: 50, 3: 30}
	map2 = map[int16]uint{}

	expected = map[int16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint64(t *testing.T) {
	map1 := map[int16]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[int16]uint64{4: 40, 5: 50, 3: 30}

	expected := map[int16]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint64{}
	map2 = map[int16]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[int16]uint64{}

	expected = map[int16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint32(t *testing.T) {
	map1 := map[int16]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[int16]uint32{4: 40, 5: 50, 3: 30}

	expected := map[int16]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint32{}
	map2 = map[int16]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[int16]uint32{}

	expected = map[int16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint16(t *testing.T) {
	map1 := map[int16]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[int16]uint16{4: 40, 5: 50, 3: 30}

	expected := map[int16]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint16{}
	map2 = map[int16]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[int16]uint16{}

	expected = map[int16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Uint8(t *testing.T) {
	map1 := map[int16]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[int16]uint8{4: 40, 5: 50, 3: 30}

	expected := map[int16]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint8{}
	map2 = map[int16]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[int16]uint8{}

	expected = map[int16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Str(t *testing.T) {
	map1 := map[int16]string{1: "10", 2: "20", 3: "30"}
	map2 := map[int16]string{4: "40", 5: "50", 3: "30"}

	expected := map[int16]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeInt16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]string{}
	map2 = map[int16]string{4: "40", 5: "50", 3: "30"}

	expected = map[int16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]string{4: "40", 5: "50", 3: "30"}

	expected = map[int16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[int16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]string{4: "40", 5: "50", 3: "30"}
	map2 = map[int16]string{}

	expected = map[int16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Bool(t *testing.T) {
	map1 := map[int16]bool{1: true, 0: false, 3: true}
	map2 := map[int16]bool{4: true, 5: true, 3: true}

	expected := map[int16]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeInt16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]bool{}
	map2 = map[int16]bool{4: true, 5: true, 3: true}

	expected = map[int16]bool{4: true, 5: true, 3: true}
	actual = MergeInt16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]bool{4: true, 5: true, 3: true}

	expected = map[int16]bool{4: true, 5: true, 3: true}
	actual = MergeInt16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[int16]bool{4: true, 5: true, 3: true}
	actual = MergeInt16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]bool{4: true, 5: true, 3: true}
	map2 = map[int16]bool{}

	expected = map[int16]bool{4: true, 5: true, 3: true}
	actual = MergeInt16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Float32(t *testing.T) {
	map1 := map[int16]float32{1: 10, 2: 20, 3: 30}
	map2 := map[int16]float32{4: 40, 5: 50, 3: 30}

	expected := map[int16]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float32{}
	map2 = map[int16]float32{4: 40, 5: 50, 3: 30}

	expected = map[int16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]float32{4: 40, 5: 50, 3: 30}

	expected = map[int16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float32{4: 40, 5: 50, 3: 30}
	map2 = map[int16]float32{}

	expected = map[int16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt16Float64(t *testing.T) {
	map1 := map[int16]float64{1: 10, 2: 20, 3: 30}
	map2 := map[int16]float64{4: 40, 5: 50, 3: 30}

	expected := map[int16]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float64{}
	map2 = map[int16]float64{4: 40, 5: 50, 3: 30}

	expected = map[int16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int16]float64{4: 40, 5: 50, 3: 30}

	expected = map[int16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int16]float64{4: 40, 5: 50, 3: 30}
	map2 = map[int16]float64{}

	expected = map[int16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt16Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt16Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int(t *testing.T) {
	map1 := map[int8]int{1: 10, 2: 20, 3: 30}
	map2 := map[int8]int{4: 40, 5: 50, 3: 30}

	expected := map[int8]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int{}
	map2 = map[int8]int{4: 40, 5: 50, 3: 30}

	expected = map[int8]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]int{4: 40, 5: 50, 3: 30}

	expected = map[int8]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int{4: 40, 5: 50, 3: 30}
	map2 = map[int8]int{}

	expected = map[int8]int{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int64(t *testing.T) {
	map1 := map[int8]int64{1: 10, 2: 20, 3: 30}
	map2 := map[int8]int64{4: 40, 5: 50, 3: 30}

	expected := map[int8]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int64{}
	map2 = map[int8]int64{4: 40, 5: 50, 3: 30}

	expected = map[int8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]int64{4: 40, 5: 50, 3: 30}

	expected = map[int8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int64{4: 40, 5: 50, 3: 30}
	map2 = map[int8]int64{}

	expected = map[int8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int32(t *testing.T) {
	map1 := map[int8]int32{1: 10, 2: 20, 3: 30}
	map2 := map[int8]int32{4: 40, 5: 50, 3: 30}

	expected := map[int8]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int32{}
	map2 = map[int8]int32{4: 40, 5: 50, 3: 30}

	expected = map[int8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]int32{4: 40, 5: 50, 3: 30}

	expected = map[int8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int32{4: 40, 5: 50, 3: 30}
	map2 = map[int8]int32{}

	expected = map[int8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Int16(t *testing.T) {
	map1 := map[int8]int16{1: 10, 2: 20, 3: 30}
	map2 := map[int8]int16{4: 40, 5: 50, 3: 30}

	expected := map[int8]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int16{}
	map2 = map[int8]int16{4: 40, 5: 50, 3: 30}

	expected = map[int8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]int16{4: 40, 5: 50, 3: 30}

	expected = map[int8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int16{4: 40, 5: 50, 3: 30}
	map2 = map[int8]int16{}

	expected = map[int8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8(t *testing.T) {
	map1 := map[int8]int8{1: 10, 2: 20, 3: 30}
	map2 := map[int8]int8{4: 40, 5: 50, 3: 30}

	expected := map[int8]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int8{}
	map2 = map[int8]int8{4: 40, 5: 50, 3: 30}

	expected = map[int8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]int8{4: 40, 5: 50, 3: 30}

	expected = map[int8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]int8{4: 40, 5: 50, 3: 30}
	map2 = map[int8]int8{}

	expected = map[int8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint(t *testing.T) {
	map1 := map[int8]uint{1: 10, 2: 20, 3: 30}
	map2 := map[int8]uint{4: 40, 5: 50, 3: 30}

	expected := map[int8]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint{}
	map2 = map[int8]uint{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]uint{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint{4: 40, 5: 50, 3: 30}
	map2 = map[int8]uint{}

	expected = map[int8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint64(t *testing.T) {
	map1 := map[int8]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[int8]uint64{4: 40, 5: 50, 3: 30}

	expected := map[int8]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint64{}
	map2 = map[int8]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]uint64{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[int8]uint64{}

	expected = map[int8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint32(t *testing.T) {
	map1 := map[int8]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[int8]uint32{4: 40, 5: 50, 3: 30}

	expected := map[int8]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint32{}
	map2 = map[int8]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]uint32{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[int8]uint32{}

	expected = map[int8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint16(t *testing.T) {
	map1 := map[int8]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[int8]uint16{4: 40, 5: 50, 3: 30}

	expected := map[int8]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint16{}
	map2 = map[int8]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]uint16{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[int8]uint16{}

	expected = map[int8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Uint8(t *testing.T) {
	map1 := map[int8]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[int8]uint8{4: 40, 5: 50, 3: 30}

	expected := map[int8]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint8{}
	map2 = map[int8]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]uint8{4: 40, 5: 50, 3: 30}

	expected = map[int8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[int8]uint8{}

	expected = map[int8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Str(t *testing.T) {
	map1 := map[int8]string{1: "10", 2: "20", 3: "30"}
	map2 := map[int8]string{4: "40", 5: "50", 3: "30"}

	expected := map[int8]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeInt8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]string{}
	map2 = map[int8]string{4: "40", 5: "50", 3: "30"}

	expected = map[int8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]string{4: "40", 5: "50", 3: "30"}

	expected = map[int8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[int8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]string{4: "40", 5: "50", 3: "30"}
	map2 = map[int8]string{}

	expected = map[int8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeInt8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Bool(t *testing.T) {
	map1 := map[int8]bool{1: true, 0: false, 3: true}
	map2 := map[int8]bool{4: true, 5: true, 3: true}

	expected := map[int8]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeInt8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]bool{}
	map2 = map[int8]bool{4: true, 5: true, 3: true}

	expected = map[int8]bool{4: true, 5: true, 3: true}
	actual = MergeInt8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]bool{4: true, 5: true, 3: true}

	expected = map[int8]bool{4: true, 5: true, 3: true}
	actual = MergeInt8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[int8]bool{4: true, 5: true, 3: true}
	actual = MergeInt8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]bool{4: true, 5: true, 3: true}
	map2 = map[int8]bool{}

	expected = map[int8]bool{4: true, 5: true, 3: true}
	actual = MergeInt8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Float32(t *testing.T) {
	map1 := map[int8]float32{1: 10, 2: 20, 3: 30}
	map2 := map[int8]float32{4: 40, 5: 50, 3: 30}

	expected := map[int8]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float32{}
	map2 = map[int8]float32{4: 40, 5: 50, 3: 30}

	expected = map[int8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]float32{4: 40, 5: 50, 3: 30}

	expected = map[int8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float32{4: 40, 5: 50, 3: 30}
	map2 = map[int8]float32{}

	expected = map[int8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeInt8Float64(t *testing.T) {
	map1 := map[int8]float64{1: 10, 2: 20, 3: 30}
	map2 := map[int8]float64{4: 40, 5: 50, 3: 30}

	expected := map[int8]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeInt8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float64{}
	map2 = map[int8]float64{4: 40, 5: 50, 3: 30}

	expected = map[int8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[int8]float64{4: 40, 5: 50, 3: 30}

	expected = map[int8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[int8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[int8]float64{4: 40, 5: 50, 3: 30}
	map2 = map[int8]float64{}

	expected = map[int8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeInt8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeInt8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeInt8Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeInt8Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt(t *testing.T) {
	map1 := map[uint]int{1: 10, 2: 20, 3: 30}
	map2 := map[uint]int{4: 40, 5: 50, 3: 30}

	expected := map[uint]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int{}
	map2 = map[uint]int{4: 40, 5: 50, 3: 30}

	expected = map[uint]int{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]int{4: 40, 5: 50, 3: 30}

	expected = map[uint]int{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]int{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int{4: 40, 5: 50, 3: 30}
	map2 = map[uint]int{}

	expected = map[uint]int{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt64(t *testing.T) {
	map1 := map[uint]int64{1: 10, 2: 20, 3: 30}
	map2 := map[uint]int64{4: 40, 5: 50, 3: 30}

	expected := map[uint]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int64{}
	map2 = map[uint]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int64{4: 40, 5: 50, 3: 30}
	map2 = map[uint]int64{}

	expected = map[uint]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt32(t *testing.T) {
	map1 := map[uint]int32{1: 10, 2: 20, 3: 30}
	map2 := map[uint]int32{4: 40, 5: 50, 3: 30}

	expected := map[uint]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int32{}
	map2 = map[uint]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int32{4: 40, 5: 50, 3: 30}
	map2 = map[uint]int32{}

	expected = map[uint]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt16(t *testing.T) {
	map1 := map[uint]int16{1: 10, 2: 20, 3: 30}
	map2 := map[uint]int16{4: 40, 5: 50, 3: 30}

	expected := map[uint]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int16{}
	map2 = map[uint]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int16{4: 40, 5: 50, 3: 30}
	map2 = map[uint]int16{}

	expected = map[uint]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintInt8(t *testing.T) {
	map1 := map[uint]int8{1: 10, 2: 20, 3: 30}
	map2 := map[uint]int8{4: 40, 5: 50, 3: 30}

	expected := map[uint]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int8{}
	map2 = map[uint]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]int8{4: 40, 5: 50, 3: 30}
	map2 = map[uint]int8{}

	expected = map[uint]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUintInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintInt8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintInt8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint(t *testing.T) {
	map1 := map[uint]uint{1: 10, 2: 20, 3: 30}
	map2 := map[uint]uint{4: 40, 5: 50, 3: 30}

	expected := map[uint]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint{}
	map2 = map[uint]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint{4: 40, 5: 50, 3: 30}
	map2 = map[uint]uint{}

	expected = map[uint]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint64(t *testing.T) {
	map1 := map[uint]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[uint]uint64{4: 40, 5: 50, 3: 30}

	expected := map[uint]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint64{}
	map2 = map[uint]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[uint]uint64{}

	expected = map[uint]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint32(t *testing.T) {
	map1 := map[uint]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[uint]uint32{4: 40, 5: 50, 3: 30}

	expected := map[uint]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint32{}
	map2 = map[uint]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[uint]uint32{}

	expected = map[uint]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint16(t *testing.T) {
	map1 := map[uint]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[uint]uint16{4: 40, 5: 50, 3: 30}

	expected := map[uint]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint16{}
	map2 = map[uint]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[uint]uint16{}

	expected = map[uint]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintUint8(t *testing.T) {
	map1 := map[uint]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[uint]uint8{4: 40, 5: 50, 3: 30}

	expected := map[uint]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint8{}
	map2 = map[uint]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[uint]uint8{}

	expected = map[uint]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUintUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintUint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintUint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintStr(t *testing.T) {
	map1 := map[uint]string{1: "10", 2: "20", 3: "30"}
	map2 := map[uint]string{4: "40", 5: "50", 3: "30"}

	expected := map[uint]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeUintStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]string{}
	map2 = map[uint]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUintStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUintStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[uint]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUintStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]string{4: "40", 5: "50", 3: "30"}
	map2 = map[uint]string{}

	expected = map[uint]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUintStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintStr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintStr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintBool(t *testing.T) {
	map1 := map[uint]bool{1: true, 0: false, 3: true}
	map2 := map[uint]bool{4: true, 5: true, 3: true}

	expected := map[uint]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeUintBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]bool{}
	map2 = map[uint]bool{4: true, 5: true, 3: true}

	expected = map[uint]bool{4: true, 5: true, 3: true}
	actual = MergeUintBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]bool{4: true, 5: true, 3: true}

	expected = map[uint]bool{4: true, 5: true, 3: true}
	actual = MergeUintBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[uint]bool{4: true, 5: true, 3: true}
	actual = MergeUintBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]bool{4: true, 5: true, 3: true}
	map2 = map[uint]bool{}

	expected = map[uint]bool{4: true, 5: true, 3: true}
	actual = MergeUintBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintBool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintBool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintFloat32(t *testing.T) {
	map1 := map[uint]float32{1: 10, 2: 20, 3: 30}
	map2 := map[uint]float32{4: 40, 5: 50, 3: 30}

	expected := map[uint]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float32{}
	map2 = map[uint]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float32{4: 40, 5: 50, 3: 30}
	map2 = map[uint]float32{}

	expected = map[uint]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintFloat32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintFloat32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUintFloat64(t *testing.T) {
	map1 := map[uint]float64{1: 10, 2: 20, 3: 30}
	map2 := map[uint]float64{4: 40, 5: 50, 3: 30}

	expected := map[uint]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUintFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float64{}
	map2 = map[uint]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint]float64{4: 40, 5: 50, 3: 30}
	map2 = map[uint]float64{}

	expected = map[uint]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUintFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUintFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUintFloat64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUintFloat64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int(t *testing.T) {
	map1 := map[uint64]int{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]int{4: 40, 5: 50, 3: 30}

	expected := map[uint64]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int{}
	map2 = map[uint64]int{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]int{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]int{}

	expected = map[uint64]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int64(t *testing.T) {
	map1 := map[uint64]int64{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]int64{4: 40, 5: 50, 3: 30}

	expected := map[uint64]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int64{}
	map2 = map[uint64]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int64{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]int64{}

	expected = map[uint64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int32(t *testing.T) {
	map1 := map[uint64]int32{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]int32{4: 40, 5: 50, 3: 30}

	expected := map[uint64]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int32{}
	map2 = map[uint64]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int32{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]int32{}

	expected = map[uint64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int16(t *testing.T) {
	map1 := map[uint64]int16{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]int16{4: 40, 5: 50, 3: 30}

	expected := map[uint64]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int16{}
	map2 = map[uint64]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int16{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]int16{}

	expected = map[uint64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Int8(t *testing.T) {
	map1 := map[uint64]int8{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]int8{4: 40, 5: 50, 3: 30}

	expected := map[uint64]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int8{}
	map2 = map[uint64]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]int8{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]int8{}

	expected = map[uint64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint(t *testing.T) {
	map1 := map[uint64]uint{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]uint{4: 40, 5: 50, 3: 30}

	expected := map[uint64]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint{}
	map2 = map[uint64]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]uint{}

	expected = map[uint64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64(t *testing.T) {
	map1 := map[uint64]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]uint64{4: 40, 5: 50, 3: 30}

	expected := map[uint64]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint64{}
	map2 = map[uint64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]uint64{}

	expected = map[uint64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint32(t *testing.T) {
	map1 := map[uint64]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]uint32{4: 40, 5: 50, 3: 30}

	expected := map[uint64]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint32{}
	map2 = map[uint64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]uint32{}

	expected = map[uint64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint16(t *testing.T) {
	map1 := map[uint64]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]uint16{4: 40, 5: 50, 3: 30}

	expected := map[uint64]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint16{}
	map2 = map[uint64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]uint16{}

	expected = map[uint64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Uint8(t *testing.T) {
	map1 := map[uint64]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]uint8{4: 40, 5: 50, 3: 30}

	expected := map[uint64]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint8{}
	map2 = map[uint64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]uint8{}

	expected = map[uint64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Str(t *testing.T) {
	map1 := map[uint64]string{1: "10", 2: "20", 3: "30"}
	map2 := map[uint64]string{4: "40", 5: "50", 3: "30"}

	expected := map[uint64]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeUint64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]string{}
	map2 = map[uint64]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[uint64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]string{4: "40", 5: "50", 3: "30"}
	map2 = map[uint64]string{}

	expected = map[uint64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Bool(t *testing.T) {
	map1 := map[uint64]bool{1: true, 0: false, 3: true}
	map2 := map[uint64]bool{4: true, 5: true, 3: true}

	expected := map[uint64]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeUint64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]bool{}
	map2 = map[uint64]bool{4: true, 5: true, 3: true}

	expected = map[uint64]bool{4: true, 5: true, 3: true}
	actual = MergeUint64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]bool{4: true, 5: true, 3: true}

	expected = map[uint64]bool{4: true, 5: true, 3: true}
	actual = MergeUint64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[uint64]bool{4: true, 5: true, 3: true}
	actual = MergeUint64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]bool{4: true, 5: true, 3: true}
	map2 = map[uint64]bool{}

	expected = map[uint64]bool{4: true, 5: true, 3: true}
	actual = MergeUint64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Float32(t *testing.T) {
	map1 := map[uint64]float32{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]float32{4: 40, 5: 50, 3: 30}

	expected := map[uint64]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float32{}
	map2 = map[uint64]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float32{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]float32{}

	expected = map[uint64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint64Float64(t *testing.T) {
	map1 := map[uint64]float64{1: 10, 2: 20, 3: 30}
	map2 := map[uint64]float64{4: 40, 5: 50, 3: 30}

	expected := map[uint64]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float64{}
	map2 = map[uint64]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint64]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint64]float64{4: 40, 5: 50, 3: 30}
	map2 = map[uint64]float64{}

	expected = map[uint64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint64Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint64Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint64Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint64Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int(t *testing.T) {
	map1 := map[uint32]int{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]int{4: 40, 5: 50, 3: 30}

	expected := map[uint32]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int{}
	map2 = map[uint32]int{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]int{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]int{}

	expected = map[uint32]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int64(t *testing.T) {
	map1 := map[uint32]int64{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]int64{4: 40, 5: 50, 3: 30}

	expected := map[uint32]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int64{}
	map2 = map[uint32]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int64{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]int64{}

	expected = map[uint32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int32(t *testing.T) {
	map1 := map[uint32]int32{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]int32{4: 40, 5: 50, 3: 30}

	expected := map[uint32]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int32{}
	map2 = map[uint32]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int32{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]int32{}

	expected = map[uint32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int16(t *testing.T) {
	map1 := map[uint32]int16{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]int16{4: 40, 5: 50, 3: 30}

	expected := map[uint32]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int16{}
	map2 = map[uint32]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int16{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]int16{}

	expected = map[uint32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Int8(t *testing.T) {
	map1 := map[uint32]int8{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]int8{4: 40, 5: 50, 3: 30}

	expected := map[uint32]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int8{}
	map2 = map[uint32]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]int8{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]int8{}

	expected = map[uint32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint(t *testing.T) {
	map1 := map[uint32]uint{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]uint{4: 40, 5: 50, 3: 30}

	expected := map[uint32]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint{}
	map2 = map[uint32]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]uint{}

	expected = map[uint32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint64(t *testing.T) {
	map1 := map[uint32]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]uint64{4: 40, 5: 50, 3: 30}

	expected := map[uint32]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint64{}
	map2 = map[uint32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]uint64{}

	expected = map[uint32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32(t *testing.T) {
	map1 := map[uint32]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]uint32{4: 40, 5: 50, 3: 30}

	expected := map[uint32]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint32{}
	map2 = map[uint32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]uint32{}

	expected = map[uint32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint16(t *testing.T) {
	map1 := map[uint32]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]uint16{4: 40, 5: 50, 3: 30}

	expected := map[uint32]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint16{}
	map2 = map[uint32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]uint16{}

	expected = map[uint32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Uint8(t *testing.T) {
	map1 := map[uint32]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]uint8{4: 40, 5: 50, 3: 30}

	expected := map[uint32]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint8{}
	map2 = map[uint32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]uint8{}

	expected = map[uint32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Str(t *testing.T) {
	map1 := map[uint32]string{1: "10", 2: "20", 3: "30"}
	map2 := map[uint32]string{4: "40", 5: "50", 3: "30"}

	expected := map[uint32]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeUint32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]string{}
	map2 = map[uint32]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[uint32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]string{4: "40", 5: "50", 3: "30"}
	map2 = map[uint32]string{}

	expected = map[uint32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Bool(t *testing.T) {
	map1 := map[uint32]bool{1: true, 0: false, 3: true}
	map2 := map[uint32]bool{4: true, 5: true, 3: true}

	expected := map[uint32]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeUint32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]bool{}
	map2 = map[uint32]bool{4: true, 5: true, 3: true}

	expected = map[uint32]bool{4: true, 5: true, 3: true}
	actual = MergeUint32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]bool{4: true, 5: true, 3: true}

	expected = map[uint32]bool{4: true, 5: true, 3: true}
	actual = MergeUint32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[uint32]bool{4: true, 5: true, 3: true}
	actual = MergeUint32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]bool{4: true, 5: true, 3: true}
	map2 = map[uint32]bool{}

	expected = map[uint32]bool{4: true, 5: true, 3: true}
	actual = MergeUint32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Float32(t *testing.T) {
	map1 := map[uint32]float32{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]float32{4: 40, 5: 50, 3: 30}

	expected := map[uint32]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float32{}
	map2 = map[uint32]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float32{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]float32{}

	expected = map[uint32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint32Float64(t *testing.T) {
	map1 := map[uint32]float64{1: 10, 2: 20, 3: 30}
	map2 := map[uint32]float64{4: 40, 5: 50, 3: 30}

	expected := map[uint32]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float64{}
	map2 = map[uint32]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint32]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint32]float64{4: 40, 5: 50, 3: 30}
	map2 = map[uint32]float64{}

	expected = map[uint32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint32Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint32Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int(t *testing.T) {
	map1 := map[uint16]int{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]int{4: 40, 5: 50, 3: 30}

	expected := map[uint16]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int{}
	map2 = map[uint16]int{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]int{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]int{}

	expected = map[uint16]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int64(t *testing.T) {
	map1 := map[uint16]int64{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]int64{4: 40, 5: 50, 3: 30}

	expected := map[uint16]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int64{}
	map2 = map[uint16]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int64{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]int64{}

	expected = map[uint16]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int32(t *testing.T) {
	map1 := map[uint16]int32{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]int32{4: 40, 5: 50, 3: 30}

	expected := map[uint16]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int32{}
	map2 = map[uint16]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int32{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]int32{}

	expected = map[uint16]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int16(t *testing.T) {
	map1 := map[uint16]int16{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]int16{4: 40, 5: 50, 3: 30}

	expected := map[uint16]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int16{}
	map2 = map[uint16]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int16{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]int16{}

	expected = map[uint16]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Int8(t *testing.T) {
	map1 := map[uint16]int8{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]int8{4: 40, 5: 50, 3: 30}

	expected := map[uint16]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int8{}
	map2 = map[uint16]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]int8{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]int8{}

	expected = map[uint16]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint(t *testing.T) {
	map1 := map[uint16]uint{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]uint{4: 40, 5: 50, 3: 30}

	expected := map[uint16]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint{}
	map2 = map[uint16]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]uint{}

	expected = map[uint16]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint64(t *testing.T) {
	map1 := map[uint16]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]uint64{4: 40, 5: 50, 3: 30}

	expected := map[uint16]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint64{}
	map2 = map[uint16]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]uint64{}

	expected = map[uint16]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint32(t *testing.T) {
	map1 := map[uint16]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]uint32{4: 40, 5: 50, 3: 30}

	expected := map[uint16]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint32{}
	map2 = map[uint16]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]uint32{}

	expected = map[uint16]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16(t *testing.T) {
	map1 := map[uint16]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]uint16{4: 40, 5: 50, 3: 30}

	expected := map[uint16]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint16{}
	map2 = map[uint16]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]uint16{}

	expected = map[uint16]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Uint8(t *testing.T) {
	map1 := map[uint16]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]uint8{4: 40, 5: 50, 3: 30}

	expected := map[uint16]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint8{}
	map2 = map[uint16]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]uint8{}

	expected = map[uint16]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Str(t *testing.T) {
	map1 := map[uint16]string{1: "10", 2: "20", 3: "30"}
	map2 := map[uint16]string{4: "40", 5: "50", 3: "30"}

	expected := map[uint16]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeUint16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]string{}
	map2 = map[uint16]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[uint16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]string{4: "40", 5: "50", 3: "30"}
	map2 = map[uint16]string{}

	expected = map[uint16]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint16Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Bool(t *testing.T) {
	map1 := map[uint16]bool{1: true, 0: false, 3: true}
	map2 := map[uint16]bool{4: true, 5: true, 3: true}

	expected := map[uint16]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeUint16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]bool{}
	map2 = map[uint16]bool{4: true, 5: true, 3: true}

	expected = map[uint16]bool{4: true, 5: true, 3: true}
	actual = MergeUint16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]bool{4: true, 5: true, 3: true}

	expected = map[uint16]bool{4: true, 5: true, 3: true}
	actual = MergeUint16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[uint16]bool{4: true, 5: true, 3: true}
	actual = MergeUint16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]bool{4: true, 5: true, 3: true}
	map2 = map[uint16]bool{}

	expected = map[uint16]bool{4: true, 5: true, 3: true}
	actual = MergeUint16Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Float32(t *testing.T) {
	map1 := map[uint16]float32{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]float32{4: 40, 5: 50, 3: 30}

	expected := map[uint16]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float32{}
	map2 = map[uint16]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float32{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]float32{}

	expected = map[uint16]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint16Float64(t *testing.T) {
	map1 := map[uint16]float64{1: 10, 2: 20, 3: 30}
	map2 := map[uint16]float64{4: 40, 5: 50, 3: 30}

	expected := map[uint16]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float64{}
	map2 = map[uint16]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint16]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint16]float64{4: 40, 5: 50, 3: 30}
	map2 = map[uint16]float64{}

	expected = map[uint16]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint16Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint16Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint16Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint16Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int(t *testing.T) {
	map1 := map[uint8]int{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]int{4: 40, 5: 50, 3: 30}

	expected := map[uint8]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int{}
	map2 = map[uint8]int{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]int{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]int{}

	expected = map[uint8]int{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int64(t *testing.T) {
	map1 := map[uint8]int64{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]int64{4: 40, 5: 50, 3: 30}

	expected := map[uint8]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int64{}
	map2 = map[uint8]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]int64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int64{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]int64{}

	expected = map[uint8]int64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int32(t *testing.T) {
	map1 := map[uint8]int32{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]int32{4: 40, 5: 50, 3: 30}

	expected := map[uint8]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int32{}
	map2 = map[uint8]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]int32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int32{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]int32{}

	expected = map[uint8]int32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int16(t *testing.T) {
	map1 := map[uint8]int16{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]int16{4: 40, 5: 50, 3: 30}

	expected := map[uint8]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int16{}
	map2 = map[uint8]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]int16{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int16{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]int16{}

	expected = map[uint8]int16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Int8(t *testing.T) {
	map1 := map[uint8]int8{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]int8{4: 40, 5: 50, 3: 30}

	expected := map[uint8]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int8{}
	map2 = map[uint8]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]int8{4: 40, 5: 50, 3: 30}

	expected = map[uint8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]int8{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]int8{}

	expected = map[uint8]int8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint(t *testing.T) {
	map1 := map[uint8]uint{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]uint{4: 40, 5: 50, 3: 30}

	expected := map[uint8]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint{}
	map2 = map[uint8]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]uint{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]uint{}

	expected = map[uint8]uint{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint64(t *testing.T) {
	map1 := map[uint8]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]uint64{4: 40, 5: 50, 3: 30}

	expected := map[uint8]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint64{}
	map2 = map[uint8]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]uint64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]uint64{}

	expected = map[uint8]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint32(t *testing.T) {
	map1 := map[uint8]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]uint32{4: 40, 5: 50, 3: 30}

	expected := map[uint8]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint32{}
	map2 = map[uint8]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]uint32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]uint32{}

	expected = map[uint8]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Uint16(t *testing.T) {
	map1 := map[uint8]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]uint16{4: 40, 5: 50, 3: 30}

	expected := map[uint8]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint16{}
	map2 = map[uint8]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]uint16{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]uint16{}

	expected = map[uint8]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8(t *testing.T) {
	map1 := map[uint8]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]uint8{4: 40, 5: 50, 3: 30}

	expected := map[uint8]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint8{}
	map2 = map[uint8]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]uint8{4: 40, 5: 50, 3: 30}

	expected = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]uint8{}

	expected = map[uint8]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Str(t *testing.T) {
	map1 := map[uint8]string{1: "10", 2: "20", 3: "30"}
	map2 := map[uint8]string{4: "40", 5: "50", 3: "30"}

	expected := map[uint8]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeUint8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]string{}
	map2 = map[uint8]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]string{4: "40", 5: "50", 3: "30"}

	expected = map[uint8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[uint8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]string{4: "40", 5: "50", 3: "30"}
	map2 = map[uint8]string{}

	expected = map[uint8]string{4: "40", 5: "50", 3: "30"}
	actual = MergeUint8Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Bool(t *testing.T) {
	map1 := map[uint8]bool{1: true, 0: false, 3: true}
	map2 := map[uint8]bool{4: true, 5: true, 3: true}

	expected := map[uint8]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeUint8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]bool{}
	map2 = map[uint8]bool{4: true, 5: true, 3: true}

	expected = map[uint8]bool{4: true, 5: true, 3: true}
	actual = MergeUint8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]bool{4: true, 5: true, 3: true}

	expected = map[uint8]bool{4: true, 5: true, 3: true}
	actual = MergeUint8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[uint8]bool{4: true, 5: true, 3: true}
	actual = MergeUint8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]bool{4: true, 5: true, 3: true}
	map2 = map[uint8]bool{}

	expected = map[uint8]bool{4: true, 5: true, 3: true}
	actual = MergeUint8Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Float32(t *testing.T) {
	map1 := map[uint8]float32{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]float32{4: 40, 5: 50, 3: 30}

	expected := map[uint8]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float32{}
	map2 = map[uint8]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]float32{4: 40, 5: 50, 3: 30}

	expected = map[uint8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float32{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]float32{}

	expected = map[uint8]float32{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeUint8Float64(t *testing.T) {
	map1 := map[uint8]float64{1: 10, 2: 20, 3: 30}
	map2 := map[uint8]float64{4: 40, 5: 50, 3: 30}

	expected := map[uint8]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeUint8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float64{}
	map2 = map[uint8]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[uint8]float64{4: 40, 5: 50, 3: 30}

	expected = map[uint8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[uint8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[uint8]float64{4: 40, 5: 50, 3: 30}
	map2 = map[uint8]float64{}

	expected = map[uint8]float64{4: 40, 5: 50, 3: 30}
	actual = MergeUint8Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeUint8Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeUint8Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeUint8Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt(t *testing.T) {
	map1 := map[string]int{"1": 10, "2": 20, "3": 30}
	map2 := map[string]int{"4": 40, "5": 50, "3": 30}

	expected := map[string]int{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int{}
	map2 = map[string]int{"4": 40, "5": 50, "3": 30}

	expected = map[string]int{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]int{"4": 40, "5": 50, "3": 30}

	expected = map[string]int{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]int{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int{"4": 40, "5": 50, "3": 30}
	map2 = map[string]int{}

	expected = map[string]int{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt64(t *testing.T) {
	map1 := map[string]int64{"1": 10, "2": 20, "3": 30}
	map2 := map[string]int64{"4": 40, "5": 50, "3": 30}

	expected := map[string]int64{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int64{}
	map2 = map[string]int64{"4": 40, "5": 50, "3": 30}

	expected = map[string]int64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]int64{"4": 40, "5": 50, "3": 30}

	expected = map[string]int64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int64{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]int64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int64{"4": 40, "5": 50, "3": 30}
	map2 = map[string]int64{}

	expected = map[string]int64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt32(t *testing.T) {
	map1 := map[string]int32{"1": 10, "2": 20, "3": 30}
	map2 := map[string]int32{"4": 40, "5": 50, "3": 30}

	expected := map[string]int32{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int32{}
	map2 = map[string]int32{"4": 40, "5": 50, "3": 30}

	expected = map[string]int32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]int32{"4": 40, "5": 50, "3": 30}

	expected = map[string]int32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int32{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]int32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int32{"4": 40, "5": 50, "3": 30}
	map2 = map[string]int32{}

	expected = map[string]int32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt16(t *testing.T) {
	map1 := map[string]int16{"1": 10, "2": 20, "3": 30}
	map2 := map[string]int16{"4": 40, "5": 50, "3": 30}

	expected := map[string]int16{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int16{}
	map2 = map[string]int16{"4": 40, "5": 50, "3": 30}

	expected = map[string]int16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]int16{"4": 40, "5": 50, "3": 30}

	expected = map[string]int16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int16{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]int16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int16{"4": 40, "5": 50, "3": 30}
	map2 = map[string]int16{}

	expected = map[string]int16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrInt8(t *testing.T) {
	map1 := map[string]int8{"1": 10, "2": 20, "3": 30}
	map2 := map[string]int8{"4": 40, "5": 50, "3": 30}

	expected := map[string]int8{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int8{}
	map2 = map[string]int8{"4": 40, "5": 50, "3": 30}

	expected = map[string]int8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]int8{"4": 40, "5": 50, "3": 30}

	expected = map[string]int8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int8{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]int8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]int8{"4": 40, "5": 50, "3": 30}
	map2 = map[string]int8{}

	expected = map[string]int8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrInt8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrInt8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint(t *testing.T) {
	map1 := map[string]uint{"1": 10, "2": 20, "3": 30}
	map2 := map[string]uint{"4": 40, "5": 50, "3": 30}

	expected := map[string]uint{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint{}
	map2 = map[string]uint{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]uint{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]uint{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint{"4": 40, "5": 50, "3": 30}
	map2 = map[string]uint{}

	expected = map[string]uint{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint64(t *testing.T) {
	map1 := map[string]uint64{"1": 10, "2": 20, "3": 30}
	map2 := map[string]uint64{"4": 40, "5": 50, "3": 30}

	expected := map[string]uint64{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint64{}
	map2 = map[string]uint64{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]uint64{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint64{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]uint64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint64{"4": 40, "5": 50, "3": 30}
	map2 = map[string]uint64{}

	expected = map[string]uint64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint32(t *testing.T) {
	map1 := map[string]uint32{"1": 10, "2": 20, "3": 30}
	map2 := map[string]uint32{"4": 40, "5": 50, "3": 30}

	expected := map[string]uint32{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint32{}
	map2 = map[string]uint32{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]uint32{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint32{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]uint32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint32{"4": 40, "5": 50, "3": 30}
	map2 = map[string]uint32{}

	expected = map[string]uint32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint16(t *testing.T) {
	map1 := map[string]uint16{"1": 10, "2": 20, "3": 30}
	map2 := map[string]uint16{"4": 40, "5": 50, "3": 30}

	expected := map[string]uint16{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint16{}
	map2 = map[string]uint16{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]uint16{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint16{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]uint16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint16{"4": 40, "5": 50, "3": 30}
	map2 = map[string]uint16{}

	expected = map[string]uint16{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrUint8(t *testing.T) {
	map1 := map[string]uint8{"1": 10, "2": 20, "3": 30}
	map2 := map[string]uint8{"4": 40, "5": 50, "3": 30}

	expected := map[string]uint8{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint8{}
	map2 = map[string]uint8{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]uint8{"4": 40, "5": 50, "3": 30}

	expected = map[string]uint8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint8{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]uint8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]uint8{"4": 40, "5": 50, "3": 30}
	map2 = map[string]uint8{}

	expected = map[string]uint8{"4": 40, "5": 50, "3": 30}
	actual = MergeStrUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrUint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrUint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStr(t *testing.T) {
	map1 := map[string]string{"1": "One", "2": "two"}
	map2 := map[string]string{"2": "Two"}

	expected := map[string]string{"1": "One", "2": "Two"}
	actual := MergeStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]string{}
	map2 = map[string]string{"1": "One", "2": "Two"}

	expected = map[string]string{"1": "One", "2": "Two"}
	actual = MergeStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]string{"1": "One", "2": "Two"}

	expected = map[string]string{"1": "One", "2": "Two"}
	actual = MergeStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]string{"1": "One", "2": "Two"}
	map2 = nil

	expected = map[string]string{"1": "One", "2": "Two"}
	actual = MergeStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]string{"1": "One", "2": "Two"}
	map2 = map[string]string{}

	expected = map[string]string{"1": "One", "2": "Two"}
	actual = MergeStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrBool(t *testing.T) {
	map1 := map[string]bool{"1": true, "0": false, "3": true}
	map2 := map[string]bool{"4": true, "5": true, "3": true}

	expected := map[string]bool{"1": true, "0": false, "4": true, "5": true, "3": true}
	actual := MergeStrBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]bool{}
	map2 = map[string]bool{"4": true, "5": true, "3": true}

	expected = map[string]bool{"4": true, "5": true, "3": true}
	actual = MergeStrBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]bool{"4": true, "5": true, "3": true}

	expected = map[string]bool{"4": true, "5": true, "3": true}
	actual = MergeStrBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]bool{"4": true, "5": true, "3": true}
	map2 = nil

	expected = map[string]bool{"4": true, "5": true, "3": true}
	actual = MergeStrBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]bool{"4": true, "5": true, "3": true}
	map2 = map[string]bool{}

	expected = map[string]bool{"4": true, "5": true, "3": true}
	actual = MergeStrBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrBool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrBool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrFloat32(t *testing.T) {
	map1 := map[string]float32{"1": 10, "2": 20, "3": 30}
	map2 := map[string]float32{"4": 40, "5": 50, "3": 30}

	expected := map[string]float32{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float32{}
	map2 = map[string]float32{"4": 40, "5": 50, "3": 30}

	expected = map[string]float32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]float32{"4": 40, "5": 50, "3": 30}

	expected = map[string]float32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float32{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]float32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float32{"4": 40, "5": 50, "3": 30}
	map2 = map[string]float32{}

	expected = map[string]float32{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrFloat32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrFloat32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeStrFloat64(t *testing.T) {
	map1 := map[string]float64{"1": 10, "2": 20, "3": 30}
	map2 := map[string]float64{"4": 40, "5": 50, "3": 30}

	expected := map[string]float64{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := MergeStrFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float64{}
	map2 = map[string]float64{"4": 40, "5": 50, "3": 30}

	expected = map[string]float64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[string]float64{"4": 40, "5": 50, "3": 30}

	expected = map[string]float64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float64{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[string]float64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[string]float64{"4": 40, "5": 50, "3": 30}
	map2 = map[string]float64{}

	expected = map[string]float64{"4": 40, "5": 50, "3": 30}
	actual = MergeStrFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeStrFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeStrFloat64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeStrFloat64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt(t *testing.T) {
	map1 := map[bool]int{true: 1, false: 0}
	map2 := map[bool]int{true: 2}

	expected := map[bool]int{true: 2, false: 0}
	actual := MergeBoolInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int{}
	map2 = map[bool]int{true: 1, false: 0}

	expected = map[bool]int{true: 1, false: 0}
	actual = MergeBoolInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]int{true: 1, false: 0}

	expected = map[bool]int{true: 1, false: 0}
	actual = MergeBoolInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int{true: 1, false: 0}
	map2 = nil

	expected = map[bool]int{true: 1, false: 0}
	actual = MergeBoolInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int{true: 1, false: 0}
	map2 = map[bool]int{}

	expected = map[bool]int{true: 1, false: 0}
	actual = MergeBoolInt(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt64(t *testing.T) {
	map1 := map[bool]int64{true: 1, false: 0}
	map2 := map[bool]int64{true: 2}

	expected := map[bool]int64{true: 2, false: 0}
	actual := MergeBoolInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int64{}
	map2 = map[bool]int64{true: 1, false: 0}

	expected = map[bool]int64{true: 1, false: 0}
	actual = MergeBoolInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]int64{true: 1, false: 0}

	expected = map[bool]int64{true: 1, false: 0}
	actual = MergeBoolInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int64{true: 1, false: 0}
	map2 = nil

	expected = map[bool]int64{true: 1, false: 0}
	actual = MergeBoolInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int64{true: 1, false: 0}
	map2 = map[bool]int64{}

	expected = map[bool]int64{true: 1, false: 0}
	actual = MergeBoolInt64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt32(t *testing.T) {
	map1 := map[bool]int32{true: 1, false: 0}
	map2 := map[bool]int32{true: 2}

	expected := map[bool]int32{true: 2, false: 0}
	actual := MergeBoolInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int32{}
	map2 = map[bool]int32{true: 1, false: 0}

	expected = map[bool]int32{true: 1, false: 0}
	actual = MergeBoolInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]int32{true: 1, false: 0}

	expected = map[bool]int32{true: 1, false: 0}
	actual = MergeBoolInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int32{true: 1, false: 0}
	map2 = nil

	expected = map[bool]int32{true: 1, false: 0}
	actual = MergeBoolInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int32{true: 1, false: 0}
	map2 = map[bool]int32{}

	expected = map[bool]int32{true: 1, false: 0}
	actual = MergeBoolInt32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt16(t *testing.T) {
	map1 := map[bool]int16{true: 1, false: 0}
	map2 := map[bool]int16{true: 2}

	expected := map[bool]int16{true: 2, false: 0}
	actual := MergeBoolInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int16{}
	map2 = map[bool]int16{true: 1, false: 0}

	expected = map[bool]int16{true: 1, false: 0}
	actual = MergeBoolInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]int16{true: 1, false: 0}

	expected = map[bool]int16{true: 1, false: 0}
	actual = MergeBoolInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int16{true: 1, false: 0}
	map2 = nil

	expected = map[bool]int16{true: 1, false: 0}
	actual = MergeBoolInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int16{true: 1, false: 0}
	map2 = map[bool]int16{}

	expected = map[bool]int16{true: 1, false: 0}
	actual = MergeBoolInt16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolInt8(t *testing.T) {
	map1 := map[bool]int8{true: 1, false: 0}
	map2 := map[bool]int8{true: 2}

	expected := map[bool]int8{true: 2, false: 0}
	actual := MergeBoolInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int8{}
	map2 = map[bool]int8{true: 1, false: 0}

	expected = map[bool]int8{true: 1, false: 0}
	actual = MergeBoolInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]int8{true: 1, false: 0}

	expected = map[bool]int8{true: 1, false: 0}
	actual = MergeBoolInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int8{true: 1, false: 0}
	map2 = nil

	expected = map[bool]int8{true: 1, false: 0}
	actual = MergeBoolInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]int8{true: 1, false: 0}
	map2 = map[bool]int8{}

	expected = map[bool]int8{true: 1, false: 0}
	actual = MergeBoolInt8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolInt8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolInt8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint(t *testing.T) {
	map1 := map[bool]uint{true: 1, false: 0}
	map2 := map[bool]uint{true: 2}

	expected := map[bool]uint{true: 2, false: 0}
	actual := MergeBoolUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint{}
	map2 = map[bool]uint{true: 1, false: 0}

	expected = map[bool]uint{true: 1, false: 0}
	actual = MergeBoolUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]uint{true: 1, false: 0}

	expected = map[bool]uint{true: 1, false: 0}
	actual = MergeBoolUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint{true: 1, false: 0}
	map2 = nil

	expected = map[bool]uint{true: 1, false: 0}
	actual = MergeBoolUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint{true: 1, false: 0}
	map2 = map[bool]uint{}

	expected = map[bool]uint{true: 1, false: 0}
	actual = MergeBoolUint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint64(t *testing.T) {
	map1 := map[bool]uint64{true: 1, false: 0}
	map2 := map[bool]uint64{true: 2}

	expected := map[bool]uint64{true: 2, false: 0}
	actual := MergeBoolUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint64{}
	map2 = map[bool]uint64{true: 1, false: 0}

	expected = map[bool]uint64{true: 1, false: 0}
	actual = MergeBoolUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]uint64{true: 1, false: 0}

	expected = map[bool]uint64{true: 1, false: 0}
	actual = MergeBoolUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint64{true: 1, false: 0}
	map2 = nil

	expected = map[bool]uint64{true: 1, false: 0}
	actual = MergeBoolUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint64{true: 1, false: 0}
	map2 = map[bool]uint64{}

	expected = map[bool]uint64{true: 1, false: 0}
	actual = MergeBoolUint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint32(t *testing.T) {
	map1 := map[bool]uint32{true: 1, false: 0}
	map2 := map[bool]uint32{true: 2}

	expected := map[bool]uint32{true: 2, false: 0}
	actual := MergeBoolUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint32{}
	map2 = map[bool]uint32{true: 1, false: 0}

	expected = map[bool]uint32{true: 1, false: 0}
	actual = MergeBoolUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]uint32{true: 1, false: 0}

	expected = map[bool]uint32{true: 1, false: 0}
	actual = MergeBoolUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint32{true: 1, false: 0}
	map2 = nil

	expected = map[bool]uint32{true: 1, false: 0}
	actual = MergeBoolUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint32{true: 1, false: 0}
	map2 = map[bool]uint32{}

	expected = map[bool]uint32{true: 1, false: 0}
	actual = MergeBoolUint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint16(t *testing.T) {
	map1 := map[bool]uint16{true: 1, false: 0}
	map2 := map[bool]uint16{true: 2}

	expected := map[bool]uint16{true: 2, false: 0}
	actual := MergeBoolUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint16{}
	map2 = map[bool]uint16{true: 1, false: 0}

	expected = map[bool]uint16{true: 1, false: 0}
	actual = MergeBoolUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]uint16{true: 1, false: 0}

	expected = map[bool]uint16{true: 1, false: 0}
	actual = MergeBoolUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint16{true: 1, false: 0}
	map2 = nil

	expected = map[bool]uint16{true: 1, false: 0}
	actual = MergeBoolUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint16{true: 1, false: 0}
	map2 = map[bool]uint16{}

	expected = map[bool]uint16{true: 1, false: 0}
	actual = MergeBoolUint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolUint8(t *testing.T) {
	map1 := map[bool]uint8{true: 1, false: 0}
	map2 := map[bool]uint8{true: 2}

	expected := map[bool]uint8{true: 2, false: 0}
	actual := MergeBoolUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint8{}
	map2 = map[bool]uint8{true: 1, false: 0}

	expected = map[bool]uint8{true: 1, false: 0}
	actual = MergeBoolUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]uint8{true: 1, false: 0}

	expected = map[bool]uint8{true: 1, false: 0}
	actual = MergeBoolUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint8{true: 1, false: 0}
	map2 = nil

	expected = map[bool]uint8{true: 1, false: 0}
	actual = MergeBoolUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]uint8{true: 1, false: 0}
	map2 = map[bool]uint8{}

	expected = map[bool]uint8{true: 1, false: 0}
	actual = MergeBoolUint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolUint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolUint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolStr(t *testing.T) {
	map1 := map[bool]string{true: "1", false: "0"}
	map2 := map[bool]string{true: "2"}

	expected := map[bool]string{true: "2", false: "0"}
	actual := MergeBoolStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]string{}
	map2 = map[bool]string{true: "1", false: "0"}

	expected = map[bool]string{true: "1", false: "0"}
	actual = MergeBoolStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]string{true: "1", false: "0"}

	expected = map[bool]string{true: "1", false: "0"}
	actual = MergeBoolStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]string{true: "1", false: "0"}
	map2 = nil

	expected = map[bool]string{true: "1", false: "0"}
	actual = MergeBoolStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]string{true: "1", false: "0"}
	map2 = map[bool]string{}

	expected = map[bool]string{true: "1", false: "0"}
	actual = MergeBoolStr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolStr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolStr failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBool(t *testing.T) {
	map1 := map[bool]bool{true: true, false: false}
	map2 := map[bool]bool{true: true}

	expected := map[bool]bool{true: true, false: false}
	actual := MergeBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]bool{}
	map2 = map[bool]bool{true: true, false: false}

	expected = map[bool]bool{true: true, false: false}
	actual = MergeBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]bool{true: true, false: false}

	expected = map[bool]bool{true: true, false: false}
	actual = MergeBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]bool{true: true, false: false}
	map2 = nil

	expected = map[bool]bool{true: true, false: false}
	actual = MergeBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]bool{true: true, false: false}
	map2 = map[bool]bool{}

	expected = map[bool]bool{true: true, false: false}
	actual = MergeBool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolFloat32(t *testing.T) {
	map1 := map[bool]float32{true: 1, false: 0}
	map2 := map[bool]float32{true: 2}

	expected := map[bool]float32{true: 2, false: 0}
	actual := MergeBoolFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float32{}
	map2 = map[bool]float32{true: 1, false: 0}

	expected = map[bool]float32{true: 1, false: 0}
	actual = MergeBoolFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]float32{true: 1, false: 0}

	expected = map[bool]float32{true: 1, false: 0}
	actual = MergeBoolFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float32{true: 1, false: 0}
	map2 = nil

	expected = map[bool]float32{true: 1, false: 0}
	actual = MergeBoolFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float32{true: 1, false: 0}
	map2 = map[bool]float32{}

	expected = map[bool]float32{true: 1, false: 0}
	actual = MergeBoolFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolFloat32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolFloat32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeBoolFloat64(t *testing.T) {
	map1 := map[bool]float64{true: 1, false: 0}
	map2 := map[bool]float64{true: 2}

	expected := map[bool]float64{true: 2, false: 0}
	actual := MergeBoolFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float64{}
	map2 = map[bool]float64{true: 1, false: 0}

	expected = map[bool]float64{true: 1, false: 0}
	actual = MergeBoolFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[bool]float64{true: 1, false: 0}

	expected = map[bool]float64{true: 1, false: 0}
	actual = MergeBoolFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float64{true: 1, false: 0}
	map2 = nil

	expected = map[bool]float64{true: 1, false: 0}
	actual = MergeBoolFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[bool]float64{true: 1, false: 0}
	map2 = map[bool]float64{}

	expected = map[bool]float64{true: 1, false: 0}
	actual = MergeBoolFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeBoolFloat64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeBoolFloat64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int(t *testing.T) {
	map1 := map[float32]int{1: 10, 2: 20, 3: 30}
	map2 := map[float32]int{4: 40, 5: 50, 3: 30}

	expected := map[float32]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int{}
	map2 = map[float32]int{4: 40, 5: 50, 3: 30}

	expected = map[float32]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]int{4: 40, 5: 50, 3: 30}

	expected = map[float32]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int{4: 40, 5: 50, 3: 30}
	map2 = map[float32]int{}

	expected = map[float32]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int64(t *testing.T) {
	map1 := map[float32]int64{1: 10, 2: 20, 3: 30}
	map2 := map[float32]int64{4: 40, 5: 50, 3: 30}

	expected := map[float32]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int64{}
	map2 = map[float32]int64{4: 40, 5: 50, 3: 30}

	expected = map[float32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]int64{4: 40, 5: 50, 3: 30}

	expected = map[float32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int64{4: 40, 5: 50, 3: 30}
	map2 = map[float32]int64{}

	expected = map[float32]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int32(t *testing.T) {
	map1 := map[float32]int32{1: 10, 2: 20, 3: 30}
	map2 := map[float32]int32{4: 40, 5: 50, 3: 30}

	expected := map[float32]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int32{}
	map2 = map[float32]int32{4: 40, 5: 50, 3: 30}

	expected = map[float32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]int32{4: 40, 5: 50, 3: 30}

	expected = map[float32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int32{4: 40, 5: 50, 3: 30}
	map2 = map[float32]int32{}

	expected = map[float32]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int16(t *testing.T) {
	map1 := map[float32]int16{1: 10, 2: 20, 3: 30}
	map2 := map[float32]int16{4: 40, 5: 50, 3: 30}

	expected := map[float32]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int16{}
	map2 = map[float32]int16{4: 40, 5: 50, 3: 30}

	expected = map[float32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]int16{4: 40, 5: 50, 3: 30}

	expected = map[float32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int16{4: 40, 5: 50, 3: 30}
	map2 = map[float32]int16{}

	expected = map[float32]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Int8(t *testing.T) {
	map1 := map[float32]int8{1: 10, 2: 20, 3: 30}
	map2 := map[float32]int8{4: 40, 5: 50, 3: 30}

	expected := map[float32]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int8{}
	map2 = map[float32]int8{4: 40, 5: 50, 3: 30}

	expected = map[float32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]int8{4: 40, 5: 50, 3: 30}

	expected = map[float32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]int8{4: 40, 5: 50, 3: 30}
	map2 = map[float32]int8{}

	expected = map[float32]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint(t *testing.T) {
	map1 := map[float32]uint{1: 10, 2: 20, 3: 30}
	map2 := map[float32]uint{4: 40, 5: 50, 3: 30}

	expected := map[float32]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint{}
	map2 = map[float32]uint{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]uint{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint{4: 40, 5: 50, 3: 30}
	map2 = map[float32]uint{}

	expected = map[float32]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint64(t *testing.T) {
	map1 := map[float32]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[float32]uint64{4: 40, 5: 50, 3: 30}

	expected := map[float32]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint64{}
	map2 = map[float32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]uint64{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[float32]uint64{}

	expected = map[float32]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint32(t *testing.T) {
	map1 := map[float32]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[float32]uint32{4: 40, 5: 50, 3: 30}

	expected := map[float32]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint32{}
	map2 = map[float32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]uint32{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[float32]uint32{}

	expected = map[float32]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint16(t *testing.T) {
	map1 := map[float32]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[float32]uint16{4: 40, 5: 50, 3: 30}

	expected := map[float32]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint16{}
	map2 = map[float32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]uint16{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[float32]uint16{}

	expected = map[float32]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Uint8(t *testing.T) {
	map1 := map[float32]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[float32]uint8{4: 40, 5: 50, 3: 30}

	expected := map[float32]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint8{}
	map2 = map[float32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]uint8{4: 40, 5: 50, 3: 30}

	expected = map[float32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[float32]uint8{}

	expected = map[float32]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Str(t *testing.T) {
	map1 := map[float32]string{1: "10", 2: "20", 3: "30"}
	map2 := map[float32]string{4: "40", 5: "50", 3: "30"}

	expected := map[float32]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeFloat32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]string{}
	map2 = map[float32]string{4: "40", 5: "50", 3: "30"}

	expected = map[float32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]string{4: "40", 5: "50", 3: "30"}

	expected = map[float32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[float32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]string{4: "40", 5: "50", 3: "30"}
	map2 = map[float32]string{}

	expected = map[float32]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat32Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Bool(t *testing.T) {
	map1 := map[float32]bool{1: true, 0: false, 3: true}
	map2 := map[float32]bool{4: true, 5: true, 3: true}

	expected := map[float32]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeFloat32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]bool{}
	map2 = map[float32]bool{4: true, 5: true, 3: true}

	expected = map[float32]bool{4: true, 5: true, 3: true}
	actual = MergeFloat32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]bool{4: true, 5: true, 3: true}

	expected = map[float32]bool{4: true, 5: true, 3: true}
	actual = MergeFloat32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[float32]bool{4: true, 5: true, 3: true}
	actual = MergeFloat32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]bool{4: true, 5: true, 3: true}
	map2 = map[float32]bool{}

	expected = map[float32]bool{4: true, 5: true, 3: true}
	actual = MergeFloat32Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32(t *testing.T) {
	map1 := map[float32]float32{1: 10, 2: 20, 3: 30}
	map2 := map[float32]float32{4: 40, 5: 50, 3: 30}

	expected := map[float32]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float32{}
	map2 = map[float32]float32{4: 40, 5: 50, 3: 30}

	expected = map[float32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]float32{4: 40, 5: 50, 3: 30}

	expected = map[float32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float32{4: 40, 5: 50, 3: 30}
	map2 = map[float32]float32{}

	expected = map[float32]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat32Float64(t *testing.T) {
	map1 := map[float32]float64{1: 10, 2: 20, 3: 30}
	map2 := map[float32]float64{4: 40, 5: 50, 3: 30}

	expected := map[float32]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float64{}
	map2 = map[float32]float64{4: 40, 5: 50, 3: 30}

	expected = map[float32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float32]float64{4: 40, 5: 50, 3: 30}

	expected = map[float32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float32]float64{4: 40, 5: 50, 3: 30}
	map2 = map[float32]float64{}

	expected = map[float32]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat32Float64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat32Float64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat32Float64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int(t *testing.T) {
	map1 := map[float64]int{1: 10, 2: 20, 3: 30}
	map2 := map[float64]int{4: 40, 5: 50, 3: 30}

	expected := map[float64]int{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int{}
	map2 = map[float64]int{4: 40, 5: 50, 3: 30}

	expected = map[float64]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]int{4: 40, 5: 50, 3: 30}

	expected = map[float64]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int{4: 40, 5: 50, 3: 30}
	map2 = map[float64]int{}

	expected = map[float64]int{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int64(t *testing.T) {
	map1 := map[float64]int64{1: 10, 2: 20, 3: 30}
	map2 := map[float64]int64{4: 40, 5: 50, 3: 30}

	expected := map[float64]int64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int64{}
	map2 = map[float64]int64{4: 40, 5: 50, 3: 30}

	expected = map[float64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]int64{4: 40, 5: 50, 3: 30}

	expected = map[float64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int64{4: 40, 5: 50, 3: 30}
	map2 = map[float64]int64{}

	expected = map[float64]int64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int32(t *testing.T) {
	map1 := map[float64]int32{1: 10, 2: 20, 3: 30}
	map2 := map[float64]int32{4: 40, 5: 50, 3: 30}

	expected := map[float64]int32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int32{}
	map2 = map[float64]int32{4: 40, 5: 50, 3: 30}

	expected = map[float64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]int32{4: 40, 5: 50, 3: 30}

	expected = map[float64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int32{4: 40, 5: 50, 3: 30}
	map2 = map[float64]int32{}

	expected = map[float64]int32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int16(t *testing.T) {
	map1 := map[float64]int16{1: 10, 2: 20, 3: 30}
	map2 := map[float64]int16{4: 40, 5: 50, 3: 30}

	expected := map[float64]int16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int16{}
	map2 = map[float64]int16{4: 40, 5: 50, 3: 30}

	expected = map[float64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]int16{4: 40, 5: 50, 3: 30}

	expected = map[float64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int16{4: 40, 5: 50, 3: 30}
	map2 = map[float64]int16{}

	expected = map[float64]int16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Int8(t *testing.T) {
	map1 := map[float64]int8{1: 10, 2: 20, 3: 30}
	map2 := map[float64]int8{4: 40, 5: 50, 3: 30}

	expected := map[float64]int8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int8{}
	map2 = map[float64]int8{4: 40, 5: 50, 3: 30}

	expected = map[float64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]int8{4: 40, 5: 50, 3: 30}

	expected = map[float64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]int8{4: 40, 5: 50, 3: 30}
	map2 = map[float64]int8{}

	expected = map[float64]int8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Int8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Int8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Int8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint(t *testing.T) {
	map1 := map[float64]uint{1: 10, 2: 20, 3: 30}
	map2 := map[float64]uint{4: 40, 5: 50, 3: 30}

	expected := map[float64]uint{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint{}
	map2 = map[float64]uint{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]uint{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint{4: 40, 5: 50, 3: 30}
	map2 = map[float64]uint{}

	expected = map[float64]uint{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint64(t *testing.T) {
	map1 := map[float64]uint64{1: 10, 2: 20, 3: 30}
	map2 := map[float64]uint64{4: 40, 5: 50, 3: 30}

	expected := map[float64]uint64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint64{}
	map2 = map[float64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]uint64{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint64{4: 40, 5: 50, 3: 30}
	map2 = map[float64]uint64{}

	expected = map[float64]uint64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint64 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint32(t *testing.T) {
	map1 := map[float64]uint32{1: 10, 2: 20, 3: 30}
	map2 := map[float64]uint32{4: 40, 5: 50, 3: 30}

	expected := map[float64]uint32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint32{}
	map2 = map[float64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]uint32{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint32{4: 40, 5: 50, 3: 30}
	map2 = map[float64]uint32{}

	expected = map[float64]uint32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint16(t *testing.T) {
	map1 := map[float64]uint16{1: 10, 2: 20, 3: 30}
	map2 := map[float64]uint16{4: 40, 5: 50, 3: 30}

	expected := map[float64]uint16{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint16{}
	map2 = map[float64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]uint16{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint16{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint16{4: 40, 5: 50, 3: 30}
	map2 = map[float64]uint16{}

	expected = map[float64]uint16{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint16(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint16(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint16 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Uint8(t *testing.T) {
	map1 := map[float64]uint8{1: 10, 2: 20, 3: 30}
	map2 := map[float64]uint8{4: 40, 5: 50, 3: 30}

	expected := map[float64]uint8{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint8{}
	map2 = map[float64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]uint8{4: 40, 5: 50, 3: 30}

	expected = map[float64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint8{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]uint8{4: 40, 5: 50, 3: 30}
	map2 = map[float64]uint8{}

	expected = map[float64]uint8{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Uint8(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Uint8(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Uint8 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Str(t *testing.T) {
	map1 := map[float64]string{1: "10", 2: "20", 3: "30"}
	map2 := map[float64]string{4: "40", 5: "50", 3: "30"}

	expected := map[float64]string{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := MergeFloat64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]string{}
	map2 = map[float64]string{4: "40", 5: "50", 3: "30"}

	expected = map[float64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]string{4: "40", 5: "50", 3: "30"}

	expected = map[float64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]string{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[float64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]string{4: "40", 5: "50", 3: "30"}
	map2 = map[float64]string{}

	expected = map[float64]string{4: "40", 5: "50", 3: "30"}
	actual = MergeFloat64Str(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Str failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Str(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Str failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Bool(t *testing.T) {
	map1 := map[float64]bool{1: true, 0: false, 3: true}
	map2 := map[float64]bool{4: true, 5: true, 3: true}

	expected := map[float64]bool{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := MergeFloat64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]bool{}
	map2 = map[float64]bool{4: true, 5: true, 3: true}

	expected = map[float64]bool{4: true, 5: true, 3: true}
	actual = MergeFloat64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]bool{4: true, 5: true, 3: true}

	expected = map[float64]bool{4: true, 5: true, 3: true}
	actual = MergeFloat64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]bool{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[float64]bool{4: true, 5: true, 3: true}
	actual = MergeFloat64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]bool{4: true, 5: true, 3: true}
	map2 = map[float64]bool{}

	expected = map[float64]bool{4: true, 5: true, 3: true}
	actual = MergeFloat64Bool(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Bool failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Bool(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Bool failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64Float32(t *testing.T) {
	map1 := map[float64]float32{1: 10, 2: 20, 3: 30}
	map2 := map[float64]float32{4: 40, 5: 50, 3: 30}

	expected := map[float64]float32{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float32{}
	map2 = map[float64]float32{4: 40, 5: 50, 3: 30}

	expected = map[float64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]float32{4: 40, 5: 50, 3: 30}

	expected = map[float64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float32{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float32{4: 40, 5: 50, 3: 30}
	map2 = map[float64]float32{}

	expected = map[float64]float32{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64Float32(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64Float32(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64Float32 failed. Expected=empty mape, actual=%v", actual)
	}
}

func TestMergeFloat64(t *testing.T) {
	map1 := map[float64]float64{1: 10, 2: 20, 3: 30}
	map2 := map[float64]float64{4: 40, 5: 50, 3: 30}

	expected := map[float64]float64{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := MergeFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float64{}
	map2 = map[float64]float64{4: 40, 5: 50, 3: 30}

	expected = map[float64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[float64]float64{4: 40, 5: 50, 3: 30}

	expected = map[float64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float64{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[float64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[float64]float64{4: 40, 5: 50, 3: 30}
	map2 = map[float64]float64{}

	expected = map[float64]float64{4: 40, 5: 50, 3: 30}
	actual = MergeFloat64(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMergeFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = MergeFloat64(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMergeFloat64 failed. Expected=empty mape, actual=%v", actual)
	}
}
