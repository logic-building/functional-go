package fp

import (
	_ "errors"
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

func TestZipIntInt64(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[int]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[int]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int]int64{}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[int]int64{}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int64{}

	expectedMap = map[int]int64{}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]int64{}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]int64{}
	actualMap = ZipIntInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt32(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[int]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[int]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int]int32{}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[int]int32{}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int32{}

	expectedMap = map[int]int32{}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]int32{}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]int32{}
	actualMap = ZipIntInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt16(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[int]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[int]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int]int16{}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[int]int16{}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int16{}

	expectedMap = map[int]int16{}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]int16{}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]int16{}
	actualMap = ZipIntInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntInt8(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[int]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[int]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int]int8{}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[int]int8{}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []int8{}

	expectedMap = map[int]int8{}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]int8{}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]int8{}
	actualMap = ZipIntInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[int]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[int]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int]uint{}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[int]uint{}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint{}

	expectedMap = map[int]uint{}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]uint{}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]uint{}
	actualMap = ZipIntUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint64(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[int]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[int]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int]uint64{}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[int]uint64{}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint64{}

	expectedMap = map[int]uint64{}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]uint64{}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]uint64{}
	actualMap = ZipIntUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint32(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[int]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[int]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int]uint32{}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[int]uint32{}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint32{}

	expectedMap = map[int]uint32{}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]uint32{}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]uint32{}
	actualMap = ZipIntUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint16(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[int]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[int]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int]uint16{}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[int]uint16{}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint16{}

	expectedMap = map[int]uint16{}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]uint16{}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]uint16{}
	actualMap = ZipIntUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntUint8(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[int]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[int]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int]uint8{}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[int]uint8{}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []uint8{}

	expectedMap = map[int]uint8{}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]uint8{}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]uint8{}
	actualMap = ZipIntUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntStr(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[int]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[int]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int]string{}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[int]string{}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []string{}

	expectedMap = map[int]string{}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]string{}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]string{}
	actualMap = ZipIntStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntBool(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[int]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[int]bool{1: true, 2: true, 3: false}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[int]bool{1: true, 2: true, 3: false}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[int]bool{}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[int]bool{}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []bool{}

	expectedMap = map[int]bool{}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]bool{}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]bool{}
	actualMap = ZipIntBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntFloat32(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[int]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[int]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int]float32{}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[int]float32{}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []float32{}

	expectedMap = map[int]float32{}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]float32{}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]float32{}
	actualMap = ZipIntFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipIntFloat64(t *testing.T) {

	list1 := []int{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[int]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[int]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int]float64{}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[int]float64{}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = []float64{}

	expectedMap = map[int]float64{}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int]float64{}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int{}
	list2 = nil

	expectedMap = map[int]float64{}
	actualMap = ZipIntFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipIntFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[int64]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[int64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int64]int{}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[int64]int{}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int{}

	expectedMap = map[int64]int{}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]int{}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]int{}
	actualMap = ZipInt64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[int64]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[int64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int64]int64{}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[int64]int64{}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int64{}

	expectedMap = map[int64]int64{}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]int64{}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]int64{}
	actualMap = ZipInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int32(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[int64]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[int64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int64]int32{}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[int64]int32{}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int32{}

	expectedMap = map[int64]int32{}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]int32{}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]int32{}
	actualMap = ZipInt64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int16(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[int64]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[int64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int64]int16{}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[int64]int16{}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int16{}

	expectedMap = map[int64]int16{}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]int16{}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]int16{}
	actualMap = ZipInt64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Int8(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[int64]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[int64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int64]int8{}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[int64]int8{}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []int8{}

	expectedMap = map[int64]int8{}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]int8{}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]int8{}
	actualMap = ZipInt64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[int64]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[int64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int64]uint{}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[int64]uint{}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint{}

	expectedMap = map[int64]uint{}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]uint{}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]uint{}
	actualMap = ZipInt64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint64(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[int64]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[int64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int64]uint64{}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[int64]uint64{}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint64{}

	expectedMap = map[int64]uint64{}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]uint64{}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]uint64{}
	actualMap = ZipInt64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint32(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[int64]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[int64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int64]uint32{}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[int64]uint32{}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint32{}

	expectedMap = map[int64]uint32{}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]uint32{}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]uint32{}
	actualMap = ZipInt64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint16(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[int64]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[int64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int64]uint16{}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[int64]uint16{}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint16{}

	expectedMap = map[int64]uint16{}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]uint16{}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]uint16{}
	actualMap = ZipInt64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Uint8(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[int64]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[int64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int64]uint8{}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[int64]uint8{}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []uint8{}

	expectedMap = map[int64]uint8{}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]uint8{}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]uint8{}
	actualMap = ZipInt64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Str(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[int64]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[int64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int64]string{}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[int64]string{}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []string{}

	expectedMap = map[int64]string{}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]string{}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]string{}
	actualMap = ZipInt64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Bool(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[int64]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[int64]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[int64]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[int64]bool{}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[int64]bool{}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []bool{}

	expectedMap = map[int64]bool{}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]bool{}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]bool{}
	actualMap = ZipInt64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Float32(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[int64]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[int64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int64]float32{}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[int64]float32{}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []float32{}

	expectedMap = map[int64]float32{}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]float32{}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]float32{}
	actualMap = ZipInt64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt64Float64(t *testing.T) {

	list1 := []int64{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[int64]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[int64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int64]float64{}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[int64]float64{}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = []float64{}

	expectedMap = map[int64]float64{}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int64]float64{}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int64{}
	list2 = nil

	expectedMap = map[int64]float64{}
	actualMap = ZipInt64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[int32]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[int32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int32]int{}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[int32]int{}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int{}

	expectedMap = map[int32]int{}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]int{}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]int{}
	actualMap = ZipInt32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int64(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[int32]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[int32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int32]int64{}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[int32]int64{}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int64{}

	expectedMap = map[int32]int64{}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]int64{}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]int64{}
	actualMap = ZipInt32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[int32]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[int32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int32]int32{}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[int32]int32{}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int32{}

	expectedMap = map[int32]int32{}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]int32{}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]int32{}
	actualMap = ZipInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int16(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[int32]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[int32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int32]int16{}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[int32]int16{}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int16{}

	expectedMap = map[int32]int16{}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]int16{}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]int16{}
	actualMap = ZipInt32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Int8(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[int32]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[int32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int32]int8{}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[int32]int8{}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []int8{}

	expectedMap = map[int32]int8{}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]int8{}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]int8{}
	actualMap = ZipInt32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[int32]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[int32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int32]uint{}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[int32]uint{}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint{}

	expectedMap = map[int32]uint{}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]uint{}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]uint{}
	actualMap = ZipInt32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint64(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[int32]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[int32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int32]uint64{}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[int32]uint64{}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint64{}

	expectedMap = map[int32]uint64{}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]uint64{}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]uint64{}
	actualMap = ZipInt32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint32(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[int32]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[int32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int32]uint32{}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[int32]uint32{}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint32{}

	expectedMap = map[int32]uint32{}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]uint32{}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]uint32{}
	actualMap = ZipInt32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint16(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[int32]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[int32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int32]uint16{}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[int32]uint16{}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint16{}

	expectedMap = map[int32]uint16{}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]uint16{}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]uint16{}
	actualMap = ZipInt32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Uint8(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[int32]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[int32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int32]uint8{}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[int32]uint8{}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []uint8{}

	expectedMap = map[int32]uint8{}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]uint8{}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]uint8{}
	actualMap = ZipInt32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Str(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[int32]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[int32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int32]string{}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[int32]string{}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []string{}

	expectedMap = map[int32]string{}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]string{}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]string{}
	actualMap = ZipInt32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Bool(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[int32]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[int32]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[int32]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[int32]bool{}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[int32]bool{}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []bool{}

	expectedMap = map[int32]bool{}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]bool{}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]bool{}
	actualMap = ZipInt32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Float32(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[int32]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[int32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int32]float32{}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[int32]float32{}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []float32{}

	expectedMap = map[int32]float32{}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]float32{}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]float32{}
	actualMap = ZipInt32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt32Float64(t *testing.T) {

	list1 := []int32{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[int32]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[int32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int32]float64{}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[int32]float64{}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = []float64{}

	expectedMap = map[int32]float64{}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int32]float64{}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int32{}
	list2 = nil

	expectedMap = map[int32]float64{}
	actualMap = ZipInt32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[int16]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[int16]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int16]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int16]int{}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[int16]int{}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int{}

	expectedMap = map[int16]int{}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]int{}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]int{}
	actualMap = ZipInt16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int64(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[int16]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[int16]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int16]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int16]int64{}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[int16]int64{}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int64{}

	expectedMap = map[int16]int64{}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]int64{}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]int64{}
	actualMap = ZipInt16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int32(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[int16]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[int16]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int16]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int16]int32{}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[int16]int32{}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int32{}

	expectedMap = map[int16]int32{}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]int32{}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]int32{}
	actualMap = ZipInt16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[int16]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[int16]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int16]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int16]int16{}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[int16]int16{}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int16{}

	expectedMap = map[int16]int16{}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]int16{}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]int16{}
	actualMap = ZipInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Int8(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[int16]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[int16]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int16]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int16]int8{}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[int16]int8{}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []int8{}

	expectedMap = map[int16]int8{}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]int8{}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]int8{}
	actualMap = ZipInt16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[int16]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[int16]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int16]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int16]uint{}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[int16]uint{}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint{}

	expectedMap = map[int16]uint{}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]uint{}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]uint{}
	actualMap = ZipInt16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint64(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[int16]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[int16]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int16]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int16]uint64{}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[int16]uint64{}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint64{}

	expectedMap = map[int16]uint64{}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]uint64{}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]uint64{}
	actualMap = ZipInt16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint32(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[int16]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[int16]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int16]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int16]uint32{}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[int16]uint32{}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint32{}

	expectedMap = map[int16]uint32{}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]uint32{}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]uint32{}
	actualMap = ZipInt16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint16(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[int16]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[int16]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int16]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int16]uint16{}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[int16]uint16{}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint16{}

	expectedMap = map[int16]uint16{}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]uint16{}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]uint16{}
	actualMap = ZipInt16Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Uint8(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[int16]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[int16]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int16]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int16]uint8{}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[int16]uint8{}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []uint8{}

	expectedMap = map[int16]uint8{}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]uint8{}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]uint8{}
	actualMap = ZipInt16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Str(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[int16]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[int16]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int16]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int16]string{}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[int16]string{}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []string{}

	expectedMap = map[int16]string{}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]string{}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]string{}
	actualMap = ZipInt16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Bool(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[int16]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[int16]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[int16]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[int16]bool{}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[int16]bool{}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []bool{}

	expectedMap = map[int16]bool{}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]bool{}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]bool{}
	actualMap = ZipInt16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Float32(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[int16]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[int16]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int16]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int16]float32{}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[int16]float32{}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []float32{}

	expectedMap = map[int16]float32{}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]float32{}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]float32{}
	actualMap = ZipInt16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt16Float64(t *testing.T) {

	list1 := []int16{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[int16]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[int16]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int16]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int16]float64{}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[int16]float64{}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = []float64{}

	expectedMap = map[int16]float64{}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int16]float64{}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int16{}
	list2 = nil

	expectedMap = map[int16]float64{}
	actualMap = ZipInt16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[int8]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[int8]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int8]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[int8]int{}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[int8]int{}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int{}

	expectedMap = map[int8]int{}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]int{}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]int{}
	actualMap = ZipInt8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int64(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[int8]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[int8]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int8]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[int8]int64{}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[int8]int64{}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int64{}

	expectedMap = map[int8]int64{}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]int64{}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]int64{}
	actualMap = ZipInt8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int32(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[int8]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[int8]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int8]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[int8]int32{}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[int8]int32{}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int32{}

	expectedMap = map[int8]int32{}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]int32{}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]int32{}
	actualMap = ZipInt8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Int16(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[int8]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[int8]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int8]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[int8]int16{}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[int8]int16{}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int16{}

	expectedMap = map[int8]int16{}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]int16{}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]int16{}
	actualMap = ZipInt8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[int8]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[int8]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int8]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[int8]int8{}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[int8]int8{}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []int8{}

	expectedMap = map[int8]int8{}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]int8{}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]int8{}
	actualMap = ZipInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[int8]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[int8]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int8]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[int8]uint{}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[int8]uint{}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint{}

	expectedMap = map[int8]uint{}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]uint{}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]uint{}
	actualMap = ZipInt8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint64(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[int8]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[int8]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int8]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[int8]uint64{}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[int8]uint64{}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint64{}

	expectedMap = map[int8]uint64{}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]uint64{}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]uint64{}
	actualMap = ZipInt8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint32(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[int8]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[int8]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int8]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[int8]uint32{}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[int8]uint32{}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint32{}

	expectedMap = map[int8]uint32{}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]uint32{}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]uint32{}
	actualMap = ZipInt8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint16(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[int8]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[int8]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int8]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[int8]uint16{}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[int8]uint16{}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint16{}

	expectedMap = map[int8]uint16{}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]uint16{}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]uint16{}
	actualMap = ZipInt8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Uint8(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[int8]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[int8]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int8]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[int8]uint8{}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[int8]uint8{}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []uint8{}

	expectedMap = map[int8]uint8{}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]uint8{}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]uint8{}
	actualMap = ZipInt8Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Str(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[int8]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[int8]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int8]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[int8]string{}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[int8]string{}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []string{}

	expectedMap = map[int8]string{}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]string{}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]string{}
	actualMap = ZipInt8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Bool(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[int8]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[int8]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[int8]bool{1: true, 2: true, 3: false}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[int8]bool{}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[int8]bool{}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []bool{}

	expectedMap = map[int8]bool{}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]bool{}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]bool{}
	actualMap = ZipInt8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Float32(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[int8]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[int8]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int8]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[int8]float32{}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[int8]float32{}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []float32{}

	expectedMap = map[int8]float32{}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]float32{}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]float32{}
	actualMap = ZipInt8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipInt8Float64(t *testing.T) {

	list1 := []int8{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[int8]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[int8]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int8]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[int8]float64{}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[int8]float64{}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = []float64{}

	expectedMap = map[int8]float64{}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[int8]float64{}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []int8{}
	list2 = nil

	expectedMap = map[int8]float64{}
	actualMap = ZipInt8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipInt8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[uint]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[uint]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint]int{}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[uint]int{}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int{}

	expectedMap = map[uint]int{}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]int{}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]int{}
	actualMap = ZipUintInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt64(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[uint]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[uint]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint]int64{}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[uint]int64{}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int64{}

	expectedMap = map[uint]int64{}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]int64{}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]int64{}
	actualMap = ZipUintInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt32(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[uint]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[uint]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint]int32{}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[uint]int32{}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int32{}

	expectedMap = map[uint]int32{}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]int32{}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]int32{}
	actualMap = ZipUintInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt16(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[uint]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[uint]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint]int16{}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[uint]int16{}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int16{}

	expectedMap = map[uint]int16{}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]int16{}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]int16{}
	actualMap = ZipUintInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintInt8(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[uint]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[uint]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint]int8{}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[uint]int8{}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []int8{}

	expectedMap = map[uint]int8{}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]int8{}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]int8{}
	actualMap = ZipUintInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[uint]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[uint]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint]uint{}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[uint]uint{}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint{}

	expectedMap = map[uint]uint{}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]uint{}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]uint{}
	actualMap = ZipUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint64(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[uint]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[uint]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint]uint64{}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[uint]uint64{}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint64{}

	expectedMap = map[uint]uint64{}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]uint64{}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]uint64{}
	actualMap = ZipUintUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint32(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[uint]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[uint]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint]uint32{}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[uint]uint32{}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint32{}

	expectedMap = map[uint]uint32{}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]uint32{}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]uint32{}
	actualMap = ZipUintUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint16(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[uint]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[uint]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint]uint16{}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[uint]uint16{}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint16{}

	expectedMap = map[uint]uint16{}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]uint16{}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]uint16{}
	actualMap = ZipUintUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintUint8(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[uint]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[uint]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint]uint8{}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[uint]uint8{}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []uint8{}

	expectedMap = map[uint]uint8{}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]uint8{}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]uint8{}
	actualMap = ZipUintUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintStr(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[uint]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[uint]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint]string{}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[uint]string{}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []string{}

	expectedMap = map[uint]string{}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]string{}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]string{}
	actualMap = ZipUintStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintBool(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[uint]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[uint]bool{1: true, 2: true, 3: false}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[uint]bool{1: true, 2: true, 3: false}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[uint]bool{}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[uint]bool{}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []bool{}

	expectedMap = map[uint]bool{}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]bool{}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]bool{}
	actualMap = ZipUintBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintFloat32(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[uint]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[uint]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint]float32{}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[uint]float32{}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []float32{}

	expectedMap = map[uint]float32{}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]float32{}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]float32{}
	actualMap = ZipUintFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUintFloat64(t *testing.T) {

	list1 := []uint{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[uint]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[uint]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint]float64{}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[uint]float64{}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = []float64{}

	expectedMap = map[uint]float64{}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint]float64{}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint{}
	list2 = nil

	expectedMap = map[uint]float64{}
	actualMap = ZipUintFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUintFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[uint64]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[uint64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint64]int{}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[uint64]int{}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int{}

	expectedMap = map[uint64]int{}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]int{}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]int{}
	actualMap = ZipUint64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int64(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[uint64]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[uint64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint64]int64{}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[uint64]int64{}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int64{}

	expectedMap = map[uint64]int64{}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]int64{}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]int64{}
	actualMap = ZipUint64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int32(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[uint64]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[uint64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint64]int32{}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[uint64]int32{}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int32{}

	expectedMap = map[uint64]int32{}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]int32{}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]int32{}
	actualMap = ZipUint64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int16(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[uint64]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[uint64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint64]int16{}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[uint64]int16{}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int16{}

	expectedMap = map[uint64]int16{}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]int16{}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]int16{}
	actualMap = ZipUint64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Int8(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[uint64]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[uint64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint64]int8{}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[uint64]int8{}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []int8{}

	expectedMap = map[uint64]int8{}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]int8{}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]int8{}
	actualMap = ZipUint64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[uint64]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[uint64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint64]uint{}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[uint64]uint{}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint{}

	expectedMap = map[uint64]uint{}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]uint{}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]uint{}
	actualMap = ZipUint64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[uint64]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[uint64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint64]uint64{}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[uint64]uint64{}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint64{}

	expectedMap = map[uint64]uint64{}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]uint64{}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]uint64{}
	actualMap = ZipUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint32(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[uint64]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[uint64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint64]uint32{}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[uint64]uint32{}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint32{}

	expectedMap = map[uint64]uint32{}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]uint32{}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]uint32{}
	actualMap = ZipUint64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint16(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[uint64]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[uint64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint64]uint16{}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[uint64]uint16{}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint16{}

	expectedMap = map[uint64]uint16{}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]uint16{}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]uint16{}
	actualMap = ZipUint64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Uint8(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[uint64]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[uint64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint64]uint8{}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[uint64]uint8{}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []uint8{}

	expectedMap = map[uint64]uint8{}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]uint8{}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]uint8{}
	actualMap = ZipUint64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Str(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[uint64]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[uint64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint64]string{}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[uint64]string{}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []string{}

	expectedMap = map[uint64]string{}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]string{}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]string{}
	actualMap = ZipUint64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Bool(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[uint64]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[uint64]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[uint64]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[uint64]bool{}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[uint64]bool{}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []bool{}

	expectedMap = map[uint64]bool{}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]bool{}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]bool{}
	actualMap = ZipUint64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Float32(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[uint64]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[uint64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint64]float32{}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[uint64]float32{}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []float32{}

	expectedMap = map[uint64]float32{}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]float32{}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]float32{}
	actualMap = ZipUint64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint64Float64(t *testing.T) {

	list1 := []uint64{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[uint64]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[uint64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint64]float64{}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[uint64]float64{}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = []float64{}

	expectedMap = map[uint64]float64{}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint64]float64{}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint64{}
	list2 = nil

	expectedMap = map[uint64]float64{}
	actualMap = ZipUint64Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint64Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[uint32]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[uint32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint32]int{}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[uint32]int{}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int{}

	expectedMap = map[uint32]int{}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]int{}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]int{}
	actualMap = ZipUint32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int64(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[uint32]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[uint32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint32]int64{}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[uint32]int64{}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int64{}

	expectedMap = map[uint32]int64{}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]int64{}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]int64{}
	actualMap = ZipUint32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int32(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[uint32]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[uint32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint32]int32{}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[uint32]int32{}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int32{}

	expectedMap = map[uint32]int32{}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]int32{}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]int32{}
	actualMap = ZipUint32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int16(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[uint32]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[uint32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint32]int16{}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[uint32]int16{}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int16{}

	expectedMap = map[uint32]int16{}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]int16{}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]int16{}
	actualMap = ZipUint32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Int8(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[uint32]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[uint32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint32]int8{}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[uint32]int8{}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []int8{}

	expectedMap = map[uint32]int8{}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]int8{}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]int8{}
	actualMap = ZipUint32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[uint32]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[uint32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint32]uint{}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[uint32]uint{}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint{}

	expectedMap = map[uint32]uint{}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]uint{}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]uint{}
	actualMap = ZipUint32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint64(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[uint32]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[uint32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint32]uint64{}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[uint32]uint64{}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint64{}

	expectedMap = map[uint32]uint64{}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]uint64{}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]uint64{}
	actualMap = ZipUint32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[uint32]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[uint32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint32]uint32{}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[uint32]uint32{}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint32{}

	expectedMap = map[uint32]uint32{}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]uint32{}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]uint32{}
	actualMap = ZipUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint16(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[uint32]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[uint32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint32]uint16{}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[uint32]uint16{}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint16{}

	expectedMap = map[uint32]uint16{}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]uint16{}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]uint16{}
	actualMap = ZipUint32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Uint8(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[uint32]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[uint32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint32]uint8{}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[uint32]uint8{}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []uint8{}

	expectedMap = map[uint32]uint8{}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]uint8{}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]uint8{}
	actualMap = ZipUint32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Str(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[uint32]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[uint32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint32]string{}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[uint32]string{}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []string{}

	expectedMap = map[uint32]string{}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]string{}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]string{}
	actualMap = ZipUint32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Bool(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[uint32]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[uint32]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[uint32]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[uint32]bool{}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[uint32]bool{}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []bool{}

	expectedMap = map[uint32]bool{}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]bool{}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]bool{}
	actualMap = ZipUint32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Float32(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[uint32]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[uint32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint32]float32{}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[uint32]float32{}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []float32{}

	expectedMap = map[uint32]float32{}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]float32{}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]float32{}
	actualMap = ZipUint32Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint32Float64(t *testing.T) {

	list1 := []uint32{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[uint32]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[uint32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint32]float64{}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[uint32]float64{}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = []float64{}

	expectedMap = map[uint32]float64{}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint32]float64{}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint32{}
	list2 = nil

	expectedMap = map[uint32]float64{}
	actualMap = ZipUint32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[uint16]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[uint16]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint16]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint16]int{}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[uint16]int{}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int{}

	expectedMap = map[uint16]int{}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]int{}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]int{}
	actualMap = ZipUint16Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int64(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[uint16]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[uint16]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint16]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint16]int64{}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[uint16]int64{}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int64{}

	expectedMap = map[uint16]int64{}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]int64{}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]int64{}
	actualMap = ZipUint16Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int32(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[uint16]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[uint16]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint16]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint16]int32{}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[uint16]int32{}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int32{}

	expectedMap = map[uint16]int32{}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]int32{}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]int32{}
	actualMap = ZipUint16Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int16(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[uint16]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[uint16]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint16]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint16]int16{}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[uint16]int16{}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int16{}

	expectedMap = map[uint16]int16{}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]int16{}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]int16{}
	actualMap = ZipUint16Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Int8(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[uint16]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[uint16]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint16]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint16]int8{}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[uint16]int8{}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []int8{}

	expectedMap = map[uint16]int8{}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]int8{}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]int8{}
	actualMap = ZipUint16Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[uint16]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[uint16]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint16]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint16]uint{}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[uint16]uint{}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint{}

	expectedMap = map[uint16]uint{}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]uint{}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]uint{}
	actualMap = ZipUint16Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint64(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[uint16]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[uint16]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint16]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint16]uint64{}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[uint16]uint64{}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint64{}

	expectedMap = map[uint16]uint64{}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]uint64{}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]uint64{}
	actualMap = ZipUint16Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint32(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[uint16]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[uint16]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint16]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint16]uint32{}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[uint16]uint32{}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint32{}

	expectedMap = map[uint16]uint32{}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]uint32{}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]uint32{}
	actualMap = ZipUint16Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[uint16]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[uint16]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint16]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint16]uint16{}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[uint16]uint16{}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint16{}

	expectedMap = map[uint16]uint16{}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]uint16{}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]uint16{}
	actualMap = ZipUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Uint8(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[uint16]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[uint16]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint16]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint16]uint8{}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[uint16]uint8{}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []uint8{}

	expectedMap = map[uint16]uint8{}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]uint8{}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]uint8{}
	actualMap = ZipUint16Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Str(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[uint16]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[uint16]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint16]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint16]string{}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[uint16]string{}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []string{}

	expectedMap = map[uint16]string{}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]string{}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]string{}
	actualMap = ZipUint16Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Bool(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[uint16]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[uint16]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[uint16]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[uint16]bool{}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[uint16]bool{}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []bool{}

	expectedMap = map[uint16]bool{}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]bool{}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]bool{}
	actualMap = ZipUint16Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Float32(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[uint16]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[uint16]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint16]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint16]float32{}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[uint16]float32{}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []float32{}

	expectedMap = map[uint16]float32{}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]float32{}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]float32{}
	actualMap = ZipUint16Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint16Float64(t *testing.T) {

	list1 := []uint16{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[uint16]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[uint16]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint16]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint16]float64{}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[uint16]float64{}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = []float64{}

	expectedMap = map[uint16]float64{}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint16]float64{}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint16{}
	list2 = nil

	expectedMap = map[uint16]float64{}
	actualMap = ZipUint16Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint16Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[uint8]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[uint8]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint8]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[uint8]int{}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[uint8]int{}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int{}

	expectedMap = map[uint8]int{}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]int{}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]int{}
	actualMap = ZipUint8Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int64(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[uint8]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[uint8]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint8]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[uint8]int64{}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[uint8]int64{}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int64{}

	expectedMap = map[uint8]int64{}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]int64{}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]int64{}
	actualMap = ZipUint8Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int32(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[uint8]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[uint8]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint8]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[uint8]int32{}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[uint8]int32{}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int32{}

	expectedMap = map[uint8]int32{}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]int32{}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]int32{}
	actualMap = ZipUint8Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int16(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[uint8]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[uint8]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint8]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[uint8]int16{}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[uint8]int16{}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int16{}

	expectedMap = map[uint8]int16{}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]int16{}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]int16{}
	actualMap = ZipUint8Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Int8(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[uint8]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[uint8]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint8]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[uint8]int8{}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[uint8]int8{}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []int8{}

	expectedMap = map[uint8]int8{}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]int8{}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]int8{}
	actualMap = ZipUint8Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[uint8]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[uint8]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint8]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[uint8]uint{}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[uint8]uint{}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint{}

	expectedMap = map[uint8]uint{}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]uint{}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]uint{}
	actualMap = ZipUint8Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint64(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[uint8]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[uint8]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint8]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[uint8]uint64{}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[uint8]uint64{}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint64{}

	expectedMap = map[uint8]uint64{}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]uint64{}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]uint64{}
	actualMap = ZipUint8Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint32(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[uint8]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[uint8]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint8]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[uint8]uint32{}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[uint8]uint32{}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint32{}

	expectedMap = map[uint8]uint32{}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]uint32{}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]uint32{}
	actualMap = ZipUint8Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Uint16(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[uint8]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[uint8]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint8]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[uint8]uint16{}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[uint8]uint16{}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint16{}

	expectedMap = map[uint8]uint16{}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]uint16{}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]uint16{}
	actualMap = ZipUint8Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[uint8]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[uint8]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint8]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[uint8]uint8{}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[uint8]uint8{}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []uint8{}

	expectedMap = map[uint8]uint8{}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]uint8{}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]uint8{}
	actualMap = ZipUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Str(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[uint8]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[uint8]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint8]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[uint8]string{}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[uint8]string{}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []string{}

	expectedMap = map[uint8]string{}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]string{}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]string{}
	actualMap = ZipUint8Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Bool(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[uint8]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[uint8]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[uint8]bool{1: true, 2: true, 3: false}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[uint8]bool{}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[uint8]bool{}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []bool{}

	expectedMap = map[uint8]bool{}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]bool{}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]bool{}
	actualMap = ZipUint8Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Float32(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[uint8]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[uint8]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint8]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[uint8]float32{}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[uint8]float32{}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []float32{}

	expectedMap = map[uint8]float32{}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]float32{}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]float32{}
	actualMap = ZipUint8Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipUint8Float64(t *testing.T) {

	list1 := []uint8{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[uint8]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[uint8]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint8]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[uint8]float64{}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[uint8]float64{}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = []float64{}

	expectedMap = map[uint8]float64{}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[uint8]float64{}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []uint8{}
	list2 = nil

	expectedMap = map[uint8]float64{}
	actualMap = ZipUint8Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipUint8Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[string]int{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int{10, 20, 30}

	expectedMap = map[string]int{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[string]int{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[string]int{}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int{}

	expectedMap = map[string]int{}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int{}

	expectedMap = map[string]int{}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]int{}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]int{}
	actualMap = ZipStrInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt64(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[string]int64{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int64{10, 20, 30}

	expectedMap = map[string]int64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[string]int64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[string]int64{}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int64{}

	expectedMap = map[string]int64{}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int64{}

	expectedMap = map[string]int64{}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]int64{}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]int64{}
	actualMap = ZipStrInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt32(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[string]int32{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int32{10, 20, 30}

	expectedMap = map[string]int32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[string]int32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[string]int32{}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int32{}

	expectedMap = map[string]int32{}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int32{}

	expectedMap = map[string]int32{}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]int32{}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]int32{}
	actualMap = ZipStrInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt16(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[string]int16{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int16{10, 20, 30}

	expectedMap = map[string]int16{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[string]int16{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[string]int16{}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int16{}

	expectedMap = map[string]int16{}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int16{}

	expectedMap = map[string]int16{}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]int16{}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]int16{}
	actualMap = ZipStrInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrInt8(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[string]int8{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int8{10, 20, 30}

	expectedMap = map[string]int8{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[string]int8{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[string]int8{}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []int8{}

	expectedMap = map[string]int8{}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []int8{}

	expectedMap = map[string]int8{}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]int8{}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]int8{}
	actualMap = ZipStrInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[string]uint{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint{10, 20, 30}

	expectedMap = map[string]uint{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[string]uint{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[string]uint{}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint{}

	expectedMap = map[string]uint{}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint{}

	expectedMap = map[string]uint{}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]uint{}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]uint{}
	actualMap = ZipStrUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint64(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[string]uint64{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[string]uint64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[string]uint64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[string]uint64{}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint64{}

	expectedMap = map[string]uint64{}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint64{}

	expectedMap = map[string]uint64{}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]uint64{}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]uint64{}
	actualMap = ZipStrUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint32(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[string]uint32{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[string]uint32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[string]uint32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[string]uint32{}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint32{}

	expectedMap = map[string]uint32{}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint32{}

	expectedMap = map[string]uint32{}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]uint32{}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]uint32{}
	actualMap = ZipStrUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint16(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[string]uint16{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[string]uint16{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[string]uint16{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[string]uint16{}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint16{}

	expectedMap = map[string]uint16{}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint16{}

	expectedMap = map[string]uint16{}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]uint16{}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]uint16{}
	actualMap = ZipStrUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrUint8(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[string]uint8{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[string]uint8{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[string]uint8{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[string]uint8{}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []uint8{}

	expectedMap = map[string]uint8{}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []uint8{}

	expectedMap = map[string]uint8{}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]uint8{}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]uint8{}
	actualMap = ZipStrUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStr(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[string]string{"1": "10", "2": "20", "3": "30", "4": "40"}
	actualMap := ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[string]string{"1": "10", "2": "20", "3": "30"}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[string]string{"1": "10", "2": "20", "3": "30"}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[string]string{}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []string{}

	expectedMap = map[string]string{}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []string{}

	expectedMap = map[string]string{}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]string{}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]string{}
	actualMap = ZipStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrBool(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []bool{true, true, false, true}

	expectedMap := map[string]bool{"1": true, "2": true, "3": false, "4": true}
	actualMap := ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []bool{true, true, false}

	expectedMap = map[string]bool{"1": true, "2": true, "3": false}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []bool{true, true, false, true}

	expectedMap = map[string]bool{"1": true, "2": true, "3": false}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[string]bool{}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []bool{}

	expectedMap = map[string]bool{}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []bool{}

	expectedMap = map[string]bool{}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]bool{}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]bool{}
	actualMap = ZipStrBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrFloat32(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[string]float32{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []float32{10, 20, 30}

	expectedMap = map[string]float32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[string]float32{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[string]float32{}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []float32{}

	expectedMap = map[string]float32{}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []float32{}

	expectedMap = map[string]float32{}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]float32{}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]float32{}
	actualMap = ZipStrFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipStrFloat64(t *testing.T) {

	list1 := []string{"1", "2", "3", "4"}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[string]float64{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []float64{10, 20, 30}

	expectedMap = map[string]float64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3"}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[string]float64{"1": 10, "2": 20, "3": 30}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[string]float64{}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{"1", "2", "3", "4"}
	list2 = []float64{}

	expectedMap = map[string]float64{}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = []float64{}

	expectedMap = map[string]float64{}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[string]float64{}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []string{}
	list2 = nil

	expectedMap = map[string]float64{}
	actualMap = ZipStrFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipStrFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []int{1, 2, 3, 4}

	expectedMap := map[bool]int{true: 4, false: 3}
	actualMap := ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int{1, 2, 3}

	expectedMap = map[bool]int{true: 2, false: 3}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []int{1, 2, 3, 4}

	expectedMap = map[bool]int{true: 2, false: 3}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int{1, 2, 3, 4}

	expectedMap = map[bool]int{}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int{}

	expectedMap = map[bool]int{}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int{}

	expectedMap = map[bool]int{}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]int{}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]int{}
	actualMap = ZipBoolInt(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt64(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []int64{1, 2, 3, 4}

	expectedMap := map[bool]int64{true: 4, false: 3}
	actualMap := ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int64{1, 2, 3}

	expectedMap = map[bool]int64{true: 2, false: 3}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []int64{1, 2, 3, 4}

	expectedMap = map[bool]int64{true: 2, false: 3}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int64{1, 2, 3, 4}

	expectedMap = map[bool]int64{}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int64{}

	expectedMap = map[bool]int64{}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int64{}

	expectedMap = map[bool]int64{}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]int64{}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]int64{}
	actualMap = ZipBoolInt64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt32(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []int32{1, 2, 3, 4}

	expectedMap := map[bool]int32{true: 4, false: 3}
	actualMap := ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int32{1, 2, 3}

	expectedMap = map[bool]int32{true: 2, false: 3}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []int32{1, 2, 3, 4}

	expectedMap = map[bool]int32{true: 2, false: 3}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int32{1, 2, 3, 4}

	expectedMap = map[bool]int32{}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int32{}

	expectedMap = map[bool]int32{}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int32{}

	expectedMap = map[bool]int32{}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]int32{}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]int32{}
	actualMap = ZipBoolInt32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt16(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []int16{1, 2, 3, 4}

	expectedMap := map[bool]int16{true: 4, false: 3}
	actualMap := ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int16{1, 2, 3}

	expectedMap = map[bool]int16{true: 2, false: 3}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []int16{1, 2, 3, 4}

	expectedMap = map[bool]int16{true: 2, false: 3}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int16{1, 2, 3, 4}

	expectedMap = map[bool]int16{}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int16{}

	expectedMap = map[bool]int16{}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int16{}

	expectedMap = map[bool]int16{}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]int16{}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]int16{}
	actualMap = ZipBoolInt16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolInt8(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []int8{1, 2, 3, 4}

	expectedMap := map[bool]int8{true: 4, false: 3}
	actualMap := ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int8{1, 2, 3}

	expectedMap = map[bool]int8{true: 2, false: 3}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []int8{1, 2, 3, 4}

	expectedMap = map[bool]int8{true: 2, false: 3}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int8{1, 2, 3, 4}

	expectedMap = map[bool]int8{}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []int8{}

	expectedMap = map[bool]int8{}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []int8{}

	expectedMap = map[bool]int8{}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]int8{}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]int8{}
	actualMap = ZipBoolInt8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolInt8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []uint{1, 2, 3, 4}

	expectedMap := map[bool]uint{true: 4, false: 3}
	actualMap := ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint{1, 2, 3}

	expectedMap = map[bool]uint{true: 2, false: 3}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []uint{1, 2, 3, 4}

	expectedMap = map[bool]uint{true: 2, false: 3}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint{1, 2, 3, 4}

	expectedMap = map[bool]uint{}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint{}

	expectedMap = map[bool]uint{}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint{}

	expectedMap = map[bool]uint{}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]uint{}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]uint{}
	actualMap = ZipBoolUint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint64(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []uint64{1, 2, 3, 4}

	expectedMap := map[bool]uint64{true: 4, false: 3}
	actualMap := ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint64{1, 2, 3}

	expectedMap = map[bool]uint64{true: 2, false: 3}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []uint64{1, 2, 3, 4}

	expectedMap = map[bool]uint64{true: 2, false: 3}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint64{1, 2, 3, 4}

	expectedMap = map[bool]uint64{}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint64{}

	expectedMap = map[bool]uint64{}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint64{}

	expectedMap = map[bool]uint64{}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]uint64{}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]uint64{}
	actualMap = ZipBoolUint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint32(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []uint32{1, 2, 3, 4}

	expectedMap := map[bool]uint32{true: 4, false: 3}
	actualMap := ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint32{1, 2, 3}

	expectedMap = map[bool]uint32{true: 2, false: 3}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []uint32{1, 2, 3, 4}

	expectedMap = map[bool]uint32{true: 2, false: 3}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint32{1, 2, 3, 4}

	expectedMap = map[bool]uint32{}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint32{}

	expectedMap = map[bool]uint32{}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint32{}

	expectedMap = map[bool]uint32{}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]uint32{}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]uint32{}
	actualMap = ZipBoolUint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint16(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []uint16{1, 2, 3, 4}

	expectedMap := map[bool]uint16{true: 4, false: 3}
	actualMap := ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint16{1, 2, 3}

	expectedMap = map[bool]uint16{true: 2, false: 3}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []uint16{1, 2, 3, 4}

	expectedMap = map[bool]uint16{true: 2, false: 3}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint16{1, 2, 3, 4}

	expectedMap = map[bool]uint16{}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint16{}

	expectedMap = map[bool]uint16{}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint16{}

	expectedMap = map[bool]uint16{}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]uint16{}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]uint16{}
	actualMap = ZipBoolUint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolUint8(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []uint8{1, 2, 3, 4}

	expectedMap := map[bool]uint8{true: 4, false: 3}
	actualMap := ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint8{1, 2, 3}

	expectedMap = map[bool]uint8{true: 2, false: 3}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []uint8{1, 2, 3, 4}

	expectedMap = map[bool]uint8{true: 2, false: 3}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint8{1, 2, 3, 4}

	expectedMap = map[bool]uint8{}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []uint8{}

	expectedMap = map[bool]uint8{}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []uint8{}

	expectedMap = map[bool]uint8{}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]uint8{}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]uint8{}
	actualMap = ZipBoolUint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolUint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolStr(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []string{"1", "2", "3", "4"}

	expectedMap := map[bool]string{true: "4", false: "3"}
	actualMap := ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []string{"1", "2", "3"}

	expectedMap = map[bool]string{true: "2", false: "3"}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []string{"1", "2", "3", "4"}

	expectedMap = map[bool]string{true: "2", false: "3"}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []string{"1", "2", "3", "4"}

	expectedMap = map[bool]string{}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []string{}

	expectedMap = map[bool]string{}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []string{}

	expectedMap = map[bool]string{}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]string{}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]string{}
	actualMap = ZipBoolStr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolStr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBool(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []bool{true, true, false, true}

	expectedMap := map[bool]bool{true: true, false: false}
	actualMap := ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []bool{true, true, false}

	expectedMap = map[bool]bool{true: true, false: false}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []bool{true, true, false, true}

	expectedMap = map[bool]bool{true: true, false: false}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []bool{true, true, false, true}

	expectedMap = map[bool]bool{}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []bool{}

	expectedMap = map[bool]bool{}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []bool{}

	expectedMap = map[bool]bool{}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]bool{}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]bool{}
	actualMap = ZipBool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolFloat32(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []float32{1, 2, 3, 4}

	expectedMap := map[bool]float32{true: 4, false: 3}
	actualMap := ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []float32{1, 2, 3}

	expectedMap = map[bool]float32{true: 2, false: 3}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []float32{1, 2, 3, 4}

	expectedMap = map[bool]float32{true: 2, false: 3}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []float32{1, 2, 3, 4}

	expectedMap = map[bool]float32{}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []float32{}

	expectedMap = map[bool]float32{}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []float32{}

	expectedMap = map[bool]float32{}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]float32{}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]float32{}
	actualMap = ZipBoolFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipBoolFloat64(t *testing.T) {

	list1 := []bool{true, true, false, true}
	list2 := []float64{1, 2, 3, 4}

	expectedMap := map[bool]float64{true: 4, false: 3}
	actualMap := ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []float64{1, 2, 3}

	expectedMap = map[bool]float64{true: 2, false: 3}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false}
	list2 = []float64{1, 2, 3, 4}

	expectedMap = map[bool]float64{true: 2, false: 3}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []float64{1, 2, 3, 4}

	expectedMap = map[bool]float64{}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{true, true, false, true}
	list2 = []float64{}

	expectedMap = map[bool]float64{}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = []float64{}

	expectedMap = map[bool]float64{}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[bool]float64{}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []bool{}
	list2 = nil

	expectedMap = map[bool]float64{}
	actualMap = ZipBoolFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipBoolFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[float32]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[float32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[float32]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[float32]int{}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[float32]int{}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int{}

	expectedMap = map[float32]int{}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]int{}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]int{}
	actualMap = ZipFloat32Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int64(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[float32]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[float32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[float32]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[float32]int64{}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[float32]int64{}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int64{}

	expectedMap = map[float32]int64{}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]int64{}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]int64{}
	actualMap = ZipFloat32Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int32(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[float32]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[float32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[float32]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[float32]int32{}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[float32]int32{}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int32{}

	expectedMap = map[float32]int32{}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]int32{}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]int32{}
	actualMap = ZipFloat32Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int16(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[float32]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[float32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[float32]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[float32]int16{}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[float32]int16{}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int16{}

	expectedMap = map[float32]int16{}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]int16{}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]int16{}
	actualMap = ZipFloat32Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Int8(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[float32]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[float32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[float32]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[float32]int8{}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[float32]int8{}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []int8{}

	expectedMap = map[float32]int8{}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]int8{}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]int8{}
	actualMap = ZipFloat32Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[float32]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[float32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[float32]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[float32]uint{}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[float32]uint{}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint{}

	expectedMap = map[float32]uint{}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]uint{}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]uint{}
	actualMap = ZipFloat32Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint64(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[float32]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[float32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[float32]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[float32]uint64{}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[float32]uint64{}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint64{}

	expectedMap = map[float32]uint64{}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]uint64{}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]uint64{}
	actualMap = ZipFloat32Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint32(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[float32]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[float32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[float32]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[float32]uint32{}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[float32]uint32{}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint32{}

	expectedMap = map[float32]uint32{}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]uint32{}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]uint32{}
	actualMap = ZipFloat32Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint16(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[float32]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[float32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[float32]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[float32]uint16{}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[float32]uint16{}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint16{}

	expectedMap = map[float32]uint16{}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]uint16{}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]uint16{}
	actualMap = ZipFloat32Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Uint8(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[float32]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[float32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[float32]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[float32]uint8{}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[float32]uint8{}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []uint8{}

	expectedMap = map[float32]uint8{}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]uint8{}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]uint8{}
	actualMap = ZipFloat32Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Str(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[float32]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[float32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[float32]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[float32]string{}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[float32]string{}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []string{}

	expectedMap = map[float32]string{}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]string{}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]string{}
	actualMap = ZipFloat32Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Bool(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[float32]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[float32]bool{1: true, 2: true, 3: false}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[float32]bool{1: true, 2: true, 3: false}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[float32]bool{}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[float32]bool{}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []bool{}

	expectedMap = map[float32]bool{}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]bool{}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]bool{}
	actualMap = ZipFloat32Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[float32]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[float32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[float32]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[float32]float32{}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[float32]float32{}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []float32{}

	expectedMap = map[float32]float32{}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]float32{}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]float32{}
	actualMap = ZipFloat32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat32Float64(t *testing.T) {

	list1 := []float32{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[float32]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[float32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[float32]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[float32]float64{}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[float32]float64{}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = []float64{}

	expectedMap = map[float32]float64{}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float32]float64{}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float32{}
	list2 = nil

	expectedMap = map[float32]float64{}
	actualMap = ZipFloat32Float64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat32Float64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []int{10, 20, 30, 40}

	expectedMap := map[float64]int{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int{10, 20, 30}

	expectedMap = map[float64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[float64]int{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int{10, 20, 30, 40}

	expectedMap = map[float64]int{}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int{}

	expectedMap = map[float64]int{}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int{}

	expectedMap = map[float64]int{}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]int{}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]int{}
	actualMap = ZipFloat64Int(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int64(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []int64{10, 20, 30, 40}

	expectedMap := map[float64]int64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int64{10, 20, 30}

	expectedMap = map[float64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[float64]int64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int64{10, 20, 30, 40}

	expectedMap = map[float64]int64{}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int64{}

	expectedMap = map[float64]int64{}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int64{}

	expectedMap = map[float64]int64{}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]int64{}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]int64{}
	actualMap = ZipFloat64Int64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int32(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []int32{10, 20, 30, 40}

	expectedMap := map[float64]int32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int32{10, 20, 30}

	expectedMap = map[float64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[float64]int32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int32{10, 20, 30, 40}

	expectedMap = map[float64]int32{}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int32{}

	expectedMap = map[float64]int32{}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int32{}

	expectedMap = map[float64]int32{}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]int32{}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]int32{}
	actualMap = ZipFloat64Int32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int16(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []int16{10, 20, 30, 40}

	expectedMap := map[float64]int16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int16{10, 20, 30}

	expectedMap = map[float64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[float64]int16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int16{10, 20, 30, 40}

	expectedMap = map[float64]int16{}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int16{}

	expectedMap = map[float64]int16{}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int16{}

	expectedMap = map[float64]int16{}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]int16{}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]int16{}
	actualMap = ZipFloat64Int16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Int8(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []int8{10, 20, 30, 40}

	expectedMap := map[float64]int8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int8{10, 20, 30}

	expectedMap = map[float64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[float64]int8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int8{10, 20, 30, 40}

	expectedMap = map[float64]int8{}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []int8{}

	expectedMap = map[float64]int8{}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []int8{}

	expectedMap = map[float64]int8{}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]int8{}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]int8{}
	actualMap = ZipFloat64Int8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Int8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []uint{10, 20, 30, 40}

	expectedMap := map[float64]uint{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint{10, 20, 30}

	expectedMap = map[float64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[float64]uint{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint{10, 20, 30, 40}

	expectedMap = map[float64]uint{}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint{}

	expectedMap = map[float64]uint{}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint{}

	expectedMap = map[float64]uint{}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]uint{}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]uint{}
	actualMap = ZipFloat64Uint(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint64(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []uint64{10, 20, 30, 40}

	expectedMap := map[float64]uint64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint64{10, 20, 30}

	expectedMap = map[float64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[float64]uint64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint64{10, 20, 30, 40}

	expectedMap = map[float64]uint64{}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint64{}

	expectedMap = map[float64]uint64{}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint64{}

	expectedMap = map[float64]uint64{}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]uint64{}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]uint64{}
	actualMap = ZipFloat64Uint64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint32(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []uint32{10, 20, 30, 40}

	expectedMap := map[float64]uint32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint32{10, 20, 30}

	expectedMap = map[float64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[float64]uint32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint32{10, 20, 30, 40}

	expectedMap = map[float64]uint32{}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint32{}

	expectedMap = map[float64]uint32{}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint32{}

	expectedMap = map[float64]uint32{}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]uint32{}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]uint32{}
	actualMap = ZipFloat64Uint32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint16(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []uint16{10, 20, 30, 40}

	expectedMap := map[float64]uint16{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint16{10, 20, 30}

	expectedMap = map[float64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[float64]uint16{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint16{10, 20, 30, 40}

	expectedMap = map[float64]uint16{}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint16{}

	expectedMap = map[float64]uint16{}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint16{}

	expectedMap = map[float64]uint16{}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]uint16{}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]uint16{}
	actualMap = ZipFloat64Uint16(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint16 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Uint8(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []uint8{10, 20, 30, 40}

	expectedMap := map[float64]uint8{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint8{10, 20, 30}

	expectedMap = map[float64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[float64]uint8{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint8{10, 20, 30, 40}

	expectedMap = map[float64]uint8{}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []uint8{}

	expectedMap = map[float64]uint8{}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []uint8{}

	expectedMap = map[float64]uint8{}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]uint8{}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]uint8{}
	actualMap = ZipFloat64Uint8(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Uint8 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Str(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []string{"10", "20", "30", "40"}

	expectedMap := map[float64]string{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []string{"10", "20", "30"}

	expectedMap = map[float64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[float64]string{1: "10", 2: "20", 3: "30"}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []string{"10", "20", "30", "40"}

	expectedMap = map[float64]string{}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []string{}

	expectedMap = map[float64]string{}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []string{}

	expectedMap = map[float64]string{}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]string{}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]string{}
	actualMap = ZipFloat64Str(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Str failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Bool(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []bool{true, true, false, true}

	expectedMap := map[float64]bool{1: true, 2: true, 3: false, 4: true}
	actualMap := ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []bool{true, true, false}

	expectedMap = map[float64]bool{1: true, 2: true, 3: false}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []bool{true, true, false, true}

	expectedMap = map[float64]bool{1: true, 2: true, 3: false}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []bool{true, true, true, true}

	expectedMap = map[float64]bool{}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []bool{}

	expectedMap = map[float64]bool{}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []bool{}

	expectedMap = map[float64]bool{}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]bool{}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]bool{}
	actualMap = ZipFloat64Bool(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Bool failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64Float32(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []float32{10, 20, 30, 40}

	expectedMap := map[float64]float32{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []float32{10, 20, 30}

	expectedMap = map[float64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[float64]float32{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []float32{10, 20, 30, 40}

	expectedMap = map[float64]float32{}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []float32{}

	expectedMap = map[float64]float32{}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []float32{}

	expectedMap = map[float64]float32{}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]float32{}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]float32{}
	actualMap = ZipFloat64Float32(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64Float32 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}

func TestZipFloat64(t *testing.T) {

	list1 := []float64{1, 2, 3, 4}
	list2 := []float64{10, 20, 30, 40}

	expectedMap := map[float64]float64{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []float64{10, 20, 30}

	expectedMap = map[float64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[float64]float64{1: 10, 2: 20, 3: 30}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []float64{10, 20, 30, 40}

	expectedMap = map[float64]float64{}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{1, 2, 3, 4}
	list2 = []float64{}

	expectedMap = map[float64]float64{}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = []float64{}

	expectedMap = map[float64]float64{}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[float64]float64{}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []float64{}
	list2 = nil

	expectedMap = map[float64]float64{}
	actualMap = ZipFloat64(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZipFloat64 failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
