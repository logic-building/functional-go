package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDropIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expectedList := []*int{&v2, &v3}
	newList := DropIntPtr(&v1, []*int{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropIntPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropIntPtr(&v1, []*int{})

	if len(newList) != 0 {
		t.Errorf("DropIntPtr failed. Expected list=%v, actual list=%v", []*int{}, newList)
	}

	newList = DropIntPtr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropIntPtr failed. Expected list=%v, actual list=%v", []*int{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropIntsPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	// Test : Drop number from the list
	expectedList := []*int{&v2, &v3}
	newList := DropIntsPtr([]*int{&v1, &v4}, []*int{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropIntsPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropIntsPtr([]*int{&v1, &v4}, []*int{})

	if len(newList) != 0 {
		t.Errorf("DropIntsPtr failed. Expected list=%v, actual list=%v", []*int{}, newList)
	}

	newList = DropIntsPtr([]*int{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropIntsPtr failed. Expected list=%v, actual list=%v", []*int{}, newList)
	}

	newList = DropIntsPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropIntsPtr failed. Expected list=%v, actual list=%v", []*int{}, newList)
	}

	newList = DropIntsPtr(nil, []*int{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expectedList := []*int64{&v2, &v3}
	newList := DropInt64Ptr(&v1, []*int64{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInt64Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt64Ptr(&v1, []*int64{})

	if len(newList) != 0 {
		t.Errorf("DropInt64Ptr failed. Expected list=%v, actual list=%v", []*int64{}, newList)
	}

	newList = DropInt64Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt64Ptr failed. Expected list=%v, actual list=%v", []*int64{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropInts64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	// Test : Drop number from the list
	expectedList := []*int64{&v2, &v3}
	newList := DropInts64Ptr([]*int64{&v1, &v4}, []*int64{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInts64Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts64Ptr([]*int64{&v1, &v4}, []*int64{})

	if len(newList) != 0 {
		t.Errorf("DropInts64Ptr failed. Expected list=%v, actual list=%v", []*int64{}, newList)
	}

	newList = DropInts64Ptr([]*int64{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts64Ptr failed. Expected list=%v, actual list=%v", []*int64{}, newList)
	}

	newList = DropInts64Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts64Ptr failed. Expected list=%v, actual list=%v", []*int64{}, newList)
	}

	newList = DropInts64Ptr(nil, []*int64{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expectedList := []*int32{&v2, &v3}
	newList := DropInt32Ptr(&v1, []*int32{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInt32Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt32Ptr(&v1, []*int32{})

	if len(newList) != 0 {
		t.Errorf("DropInt32Ptr failed. Expected list=%v, actual list=%v", []*int32{}, newList)
	}

	newList = DropInt32Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt32Ptr failed. Expected list=%v, actual list=%v", []*int32{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropInts32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	// Test : Drop number from the list
	expectedList := []*int32{&v2, &v3}
	newList := DropInts32Ptr([]*int32{&v1, &v4}, []*int32{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInts32Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts32Ptr([]*int32{&v1, &v4}, []*int32{})

	if len(newList) != 0 {
		t.Errorf("DropInts32Ptr failed. Expected list=%v, actual list=%v", []*int32{}, newList)
	}

	newList = DropInts32Ptr([]*int32{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts32Ptr failed. Expected list=%v, actual list=%v", []*int32{}, newList)
	}

	newList = DropInts32Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts32Ptr failed. Expected list=%v, actual list=%v", []*int32{}, newList)
	}

	newList = DropInts32Ptr(nil, []*int32{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expectedList := []*int16{&v2, &v3}
	newList := DropInt16Ptr(&v1, []*int16{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInt16Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt16Ptr(&v1, []*int16{})

	if len(newList) != 0 {
		t.Errorf("DropInt16Ptr failed. Expected list=%v, actual list=%v", []*int16{}, newList)
	}

	newList = DropInt16Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt16Ptr failed. Expected list=%v, actual list=%v", []*int16{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropInts16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	// Test : Drop number from the list
	expectedList := []*int16{&v2, &v3}
	newList := DropInts16Ptr([]*int16{&v1, &v4}, []*int16{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInts16Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts16Ptr([]*int16{&v1, &v4}, []*int16{})

	if len(newList) != 0 {
		t.Errorf("DropInts16Ptr failed. Expected list=%v, actual list=%v", []*int16{}, newList)
	}

	newList = DropInts16Ptr([]*int16{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts16Ptr failed. Expected list=%v, actual list=%v", []*int16{}, newList)
	}

	newList = DropInts16Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts16Ptr failed. Expected list=%v, actual list=%v", []*int16{}, newList)
	}

	newList = DropInts16Ptr(nil, []*int16{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expectedList := []*int8{&v2, &v3}
	newList := DropInt8Ptr(&v1, []*int8{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInt8Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt8Ptr(&v1, []*int8{})

	if len(newList) != 0 {
		t.Errorf("DropInt8Ptr failed. Expected list=%v, actual list=%v", []*int8{}, newList)
	}

	newList = DropInt8Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt8Ptr failed. Expected list=%v, actual list=%v", []*int8{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropInts8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	// Test : Drop number from the list
	expectedList := []*int8{&v2, &v3}
	newList := DropInts8Ptr([]*int8{&v1, &v4}, []*int8{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropInts8Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts8Ptr([]*int8{&v1, &v4}, []*int8{})

	if len(newList) != 0 {
		t.Errorf("DropInts8Ptr failed. Expected list=%v, actual list=%v", []*int8{}, newList)
	}

	newList = DropInts8Ptr([]*int8{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8Ptr failed. Expected list=%v, actual list=%v", []*int8{}, newList)
	}

	newList = DropInts8Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8Ptr failed. Expected list=%v, actual list=%v", []*int8{}, newList)
	}

	newList = DropInts8Ptr(nil, []*int8{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expectedList := []*uint{&v2, &v3}
	newList := DropUintPtr(&v1, []*uint{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUintPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUintPtr(&v1, []*uint{})

	if len(newList) != 0 {
		t.Errorf("DropUintPtr failed. Expected list=%v, actual list=%v", []*uint{}, newList)
	}

	newList = DropUintPtr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUintPtr failed. Expected list=%v, actual list=%v", []*uint{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropUintsPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	// Test : Drop number from the list
	expectedList := []*uint{&v2, &v3}
	newList := DropUintsPtr([]*uint{&v1, &v4}, []*uint{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUintsPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUintsPtr([]*uint{&v1, &v4}, []*uint{})

	if len(newList) != 0 {
		t.Errorf("DropUintsPtr failed. Expected list=%v, actual list=%v", []*uint{}, newList)
	}

	newList = DropUintsPtr([]*uint{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUintsPtr failed. Expected list=%v, actual list=%v", []*uint{}, newList)
	}

	newList = DropUintsPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUintsPtr failed. Expected list=%v, actual list=%v", []*uint{}, newList)
	}

	newList = DropUintsPtr(nil, []*uint{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropUints failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expectedList := []*uint64{&v2, &v3}
	newList := DropUint64Ptr(&v1, []*uint64{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUint64Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint64Ptr(&v1, []*uint64{})

	if len(newList) != 0 {
		t.Errorf("DropUint64Ptr failed. Expected list=%v, actual list=%v", []*uint64{}, newList)
	}

	newList = DropUint64Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint64Ptr failed. Expected list=%v, actual list=%v", []*uint64{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropUint64sPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	// Test : Drop number from the list
	expectedList := []*uint64{&v2, &v3}
	newList := DropUint64sPtr([]*uint64{&v1, &v4}, []*uint64{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUint64sPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint64sPtr([]*uint64{&v1, &v4}, []*uint64{})

	if len(newList) != 0 {
		t.Errorf("DropUint64sPtr failed. Expected list=%v, actual list=%v", []*uint64{}, newList)
	}

	newList = DropUint64sPtr([]*uint64{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint64sPtr failed. Expected list=%v, actual list=%v", []*uint64{}, newList)
	}

	newList = DropUint64sPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint64sPtr failed. Expected list=%v, actual list=%v", []*uint64{}, newList)
	}

	newList = DropUint64sPtr(nil, []*uint64{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropUint64s failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expectedList := []*uint32{&v2, &v3}
	newList := DropUint32Ptr(&v1, []*uint32{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUint32Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint32Ptr(&v1, []*uint32{})

	if len(newList) != 0 {
		t.Errorf("DropUint32Ptr failed. Expected list=%v, actual list=%v", []*uint32{}, newList)
	}

	newList = DropUint32Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint32Ptr failed. Expected list=%v, actual list=%v", []*uint32{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropUints32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	// Test : Drop number from the list
	expectedList := []*uint32{&v2, &v3}
	newList := DropUints32Ptr([]*uint32{&v1, &v4}, []*uint32{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUints32Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints32Ptr([]*uint32{&v1, &v4}, []*uint32{})

	if len(newList) != 0 {
		t.Errorf("DropUints32Ptr failed. Expected list=%v, actual list=%v", []*uint32{}, newList)
	}

	newList = DropUints32Ptr([]*uint32{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints32Ptr failed. Expected list=%v, actual list=%v", []*uint32{}, newList)
	}

	newList = DropUints32Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints32Ptr failed. Expected list=%v, actual list=%v", []*uint32{}, newList)
	}

	newList = DropUints32Ptr(nil, []*uint32{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expectedList := []*uint16{&v2, &v3}
	newList := DropUint16Ptr(&v1, []*uint16{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUint16Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint16Ptr(&v1, []*uint16{})

	if len(newList) != 0 {
		t.Errorf("DropUint16Ptr failed. Expected list=%v, actual list=%v", []*uint16{}, newList)
	}

	newList = DropUint16Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint16Ptr failed. Expected list=%v, actual list=%v", []*uint16{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropUints16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	// Test : Drop number from the list
	expectedList := []*uint16{&v2, &v3}
	newList := DropUints16Ptr([]*uint16{&v1, &v4}, []*uint16{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUints16Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints16Ptr([]*uint16{&v1, &v4}, []*uint16{})

	if len(newList) != 0 {
		t.Errorf("DropUints16Ptr failed. Expected list=%v, actual list=%v", []*uint16{}, newList)
	}

	newList = DropUints16Ptr([]*uint16{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints16Ptr failed. Expected list=%v, actual list=%v", []*uint16{}, newList)
	}

	newList = DropUints16Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints16Ptr failed. Expected list=%v, actual list=%v", []*uint16{}, newList)
	}

	newList = DropUints16Ptr(nil, []*uint16{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expectedList := []*uint8{&v2, &v3}
	newList := DropUint8Ptr(&v1, []*uint8{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUint8Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint8Ptr(&v1, []*uint8{})

	if len(newList) != 0 {
		t.Errorf("DropUint8Ptr failed. Expected list=%v, actual list=%v", []*uint8{}, newList)
	}

	newList = DropUint8Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint8Ptr failed. Expected list=%v, actual list=%v", []*uint8{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropUints8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	// Test : Drop number from the list
	expectedList := []*uint8{&v2, &v3}
	newList := DropUints8Ptr([]*uint8{&v1, &v4}, []*uint8{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropUints8Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints8Ptr([]*uint8{&v1, &v4}, []*uint8{})

	if len(newList) != 0 {
		t.Errorf("DropUints8Ptr failed. Expected list=%v, actual list=%v", []*uint8{}, newList)
	}

	newList = DropUints8Ptr([]*uint8{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints8Ptr failed. Expected list=%v, actual list=%v", []*uint8{}, newList)
	}

	newList = DropUints8Ptr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints8Ptr failed. Expected list=%v, actual list=%v", []*uint8{}, newList)
	}

	newList = DropUints8Ptr(nil, []*uint8{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expectedList := []*string{&v2, &v3}
	newList := DropStrPtr(&v1, []*string{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropStrPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStrPtr(&v1, []*string{})

	if len(newList) != 0 {
		t.Errorf("DropStrPtr failed. Expected list=%v, actual list=%v", []*string{}, newList)
	}

	newList = DropStrPtr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrPtr failed. Expected list=%v, actual list=%v", []*string{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropStrsPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	// Test : Drop number from the list
	expectedList := []*string{&v2, &v3}
	newList := DropStrsPtr([]*string{&v1, &v4}, []*string{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropStrsPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStrsPtr([]*string{&v1, &v4}, []*string{})

	if len(newList) != 0 {
		t.Errorf("DropStrsPtr failed. Expected list=%v, actual list=%v", []*string{}, newList)
	}

	newList = DropStrsPtr([]*string{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsPtr failed. Expected list=%v, actual list=%v", []*string{}, newList)
	}

	newList = DropStrsPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsPtr failed. Expected list=%v, actual list=%v", []*string{}, newList)
	}

	newList = DropStrsPtr(nil, []*string{&v1, &v4})
	if *newList[0] != "1" || *newList[1] != "4" {
		t.Errorf("DropStrs failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropBoolPtr(t *testing.T) {
	var v1 bool = false
	var v2 bool = true
	var v3 bool = true

	expectedList := []*bool{&v2, &v3}
	newList := DropBoolPtr(&v1, []*bool{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropBoolPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropBoolsPtr(t *testing.T) {
	var v1 bool = false
	var v2 bool = true
	var v3 bool = true

	expectedList := []*bool{&v2, &v3}
	newList := DropBoolsPtr([]*bool{&v1}, []*bool{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropBoolsPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}


func TestDropFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expectedList := []*float32{&v2, &v3}
	newList := DropFloat32Ptr(&v1, []*float32{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropFloat32Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat32Ptr(&v1, []*float32{})

	if len(newList) != 0 {
		t.Errorf("DropFloat32Ptr failed. Expected list=%v, actual list=%v", []*float32{}, newList)
	}

	newList = DropFloat32Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat32Ptr failed. Expected list=%v, actual list=%v", []*float32{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropFloat32sPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	// Test : Drop number from the list
	expectedList := []*float32{&v2, &v3}
	newList := DropFloat32sPtr([]*float32{&v1, &v4}, []*float32{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropFloat32sPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat32sPtr([]*float32{&v1, &v4}, []*float32{})

	if len(newList) != 0 {
		t.Errorf("DropFloat32sPtr failed. Expected list=%v, actual list=%v", []*float32{}, newList)
	}

	newList = DropFloat32sPtr([]*float32{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat32sPtr failed. Expected list=%v, actual list=%v", []*float32{}, newList)
	}

	newList = DropFloat32sPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat32sPtr failed. Expected list=%v, actual list=%v", []*float32{}, newList)
	}

	newList = DropFloat32sPtr(nil, []*float32{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropFloat32s failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expectedList := []*float64{&v2, &v3}
	newList := DropFloat64Ptr(&v1, []*float64{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropFloat64Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat64Ptr(&v1, []*float64{})

	if len(newList) != 0 {
		t.Errorf("DropFloat64Ptr failed. Expected list=%v, actual list=%v", []*float64{}, newList)
	}

	newList = DropFloat64Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat64Ptr failed. Expected list=%v, actual list=%v", []*float64{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDropFloat64sPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	// Test : Drop number from the list
	expectedList := []*float64{&v2, &v3}
	newList := DropFloat64sPtr([]*float64{&v1, &v4}, []*float64{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("DropFloat64sPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat64sPtr([]*float64{&v1, &v4}, []*float64{})

	if len(newList) != 0 {
		t.Errorf("DropFloat64sPtr failed. Expected list=%v, actual list=%v", []*float64{}, newList)
	}

	newList = DropFloat64sPtr([]*float64{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat64sPtr failed. Expected list=%v, actual list=%v", []*float64{}, newList)
	}

	newList = DropFloat64sPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat64sPtr failed. Expected list=%v, actual list=%v", []*float64{}, newList)
	}

	newList = DropFloat64sPtr(nil, []*float64{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropFloat64s failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}
