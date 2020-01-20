package fp

import (
	"reflect"
	"testing"
)

func TestDropLastPtrInt(t *testing.T) {
    var v1 int = 1
	var v2 int = 2
    var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	list := []*int{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int{&v1, &v2}
	expectedList = []*int{&v1}
	actualList = DropLastPtrInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int{&v1}
	expectedList = []*int{}
	actualList = DropLastPtrInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int{}
	expectedList = []*int{}
	actualList = DropLastPtrInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*int{}
	actualList = DropLastPtrInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrInt64(t *testing.T) {
    var v1 int64 = 1
	var v2 int64 = 2
    var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	list := []*int64{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int64{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int64{&v1, &v2}
	expectedList = []*int64{&v1}
	actualList = DropLastPtrInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int64{&v1}
	expectedList = []*int64{}
	actualList = DropLastPtrInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int64{}
	expectedList = []*int64{}
	actualList = DropLastPtrInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*int64{}
	actualList = DropLastPtrInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrInt32(t *testing.T) {
    var v1 int32 = 1
	var v2 int32 = 2
    var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	list := []*int32{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int32{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int32{&v1, &v2}
	expectedList = []*int32{&v1}
	actualList = DropLastPtrInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int32{&v1}
	expectedList = []*int32{}
	actualList = DropLastPtrInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int32{}
	expectedList = []*int32{}
	actualList = DropLastPtrInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*int32{}
	actualList = DropLastPtrInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrInt16(t *testing.T) {
    var v1 int16 = 1
	var v2 int16 = 2
    var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	list := []*int16{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int16{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int16{&v1, &v2}
	expectedList = []*int16{&v1}
	actualList = DropLastPtrInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int16{&v1}
	expectedList = []*int16{}
	actualList = DropLastPtrInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int16{}
	expectedList = []*int16{}
	actualList = DropLastPtrInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*int16{}
	actualList = DropLastPtrInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrInt8(t *testing.T) {
    var v1 int8 = 1
	var v2 int8 = 2
    var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	list := []*int8{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int8{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int8{&v1, &v2}
	expectedList = []*int8{&v1}
	actualList = DropLastPtrInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int8{&v1}
	expectedList = []*int8{}
	actualList = DropLastPtrInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*int8{}
	expectedList = []*int8{}
	actualList = DropLastPtrInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*int8{}
	actualList = DropLastPtrInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrUint(t *testing.T) {
    var v1 uint = 1
	var v2 uint = 2
    var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	list := []*uint{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint{&v1, &v2}
	expectedList = []*uint{&v1}
	actualList = DropLastPtrUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint{&v1}
	expectedList = []*uint{}
	actualList = DropLastPtrUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint{}
	expectedList = []*uint{}
	actualList = DropLastPtrUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*uint{}
	actualList = DropLastPtrUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrUint64(t *testing.T) {
    var v1 uint64 = 1
	var v2 uint64 = 2
    var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	list := []*uint64{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint64{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint64{&v1, &v2}
	expectedList = []*uint64{&v1}
	actualList = DropLastPtrUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint64{&v1}
	expectedList = []*uint64{}
	actualList = DropLastPtrUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint64{}
	expectedList = []*uint64{}
	actualList = DropLastPtrUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*uint64{}
	actualList = DropLastPtrUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrUint32(t *testing.T) {
    var v1 uint32 = 1
	var v2 uint32 = 2
    var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	list := []*uint32{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint32{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint32{&v1, &v2}
	expectedList = []*uint32{&v1}
	actualList = DropLastPtrUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint32{&v1}
	expectedList = []*uint32{}
	actualList = DropLastPtrUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint32{}
	expectedList = []*uint32{}
	actualList = DropLastPtrUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*uint32{}
	actualList = DropLastPtrUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrUint16(t *testing.T) {
    var v1 uint16 = 1
	var v2 uint16 = 2
    var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	list := []*uint16{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint16{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint16{&v1, &v2}
	expectedList = []*uint16{&v1}
	actualList = DropLastPtrUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint16{&v1}
	expectedList = []*uint16{}
	actualList = DropLastPtrUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint16{}
	expectedList = []*uint16{}
	actualList = DropLastPtrUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*uint16{}
	actualList = DropLastPtrUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrUint8(t *testing.T) {
    var v1 uint8 = 1
	var v2 uint8 = 2
    var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	list := []*uint8{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint8{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint8{&v1, &v2}
	expectedList = []*uint8{&v1}
	actualList = DropLastPtrUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint8{&v1}
	expectedList = []*uint8{}
	actualList = DropLastPtrUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*uint8{}
	expectedList = []*uint8{}
	actualList = DropLastPtrUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*uint8{}
	actualList = DropLastPtrUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrStr(t *testing.T) {
    var v1 string = "1"
	var v2 string = "2"
    var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*string{&v1, &v2, &v3, &v4}
	actualList := DropLastPtrStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*string{&v1, &v2}
	expectedList = []*string{&v1}
	actualList = DropLastPtrStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*string{&v1}
	expectedList = []*string{}
	actualList = DropLastPtrStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*string{}
	expectedList = []*string{}
	actualList = DropLastPtrStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*string{}
	actualList = DropLastPtrStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastPtrBool(t *testing.T) {
	var true bool = true
    var false bool = false
	list := []*bool{&true, &true, &true, &true, &false}
	expectedList := []*bool{&true, &true, &true, &true}
	actualList := DropLastPtrBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*bool{&true, &true}
	expectedList = []*bool{&true}
	actualList = DropLastPtrBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*bool{&true}
	expectedList = []*bool{}
	actualList = DropLastPtrBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*bool{}
	expectedList = []*bool{}
	actualList = DropLastPtrBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*bool{}
	actualList = DropLastPtrBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastPtrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
