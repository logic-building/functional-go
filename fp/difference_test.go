package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDifferenceInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	actual := DifferenceInt()
	if len(actual) != 0 {
		t.Errorf("DifferenceInt failed.")
	}

	expected := []int{v8, v2, v1}
	actual = DifferenceInt(expected)

	if len(actual) != 3 || !ExistsInt(v8, actual) || !ExistsInt(v2, actual) || !ExistsInt(v1, actual) {
		t.Errorf("DifferenceInt failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int{v8, v2, v1, v1, v2, v8}
	list2 := []int{v8}
	actual = DifferenceInt(list1, list2)
	if len(actual) != 2 || !ExistsInt(v2, actual) || !ExistsInt(v1, actual) {
		t.Errorf("DifferenceInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceIntPtr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceIntPtr failed.")
	}

	expectedPtr := []*int{&v8, &v2, &v1}
	actualPtr = DifferenceIntPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsIntPtr(&v8, actualPtr) || !ExistsIntPtr(&v2, actualPtr) || !ExistsIntPtr(&v1, actualPtr) {
		t.Errorf("DifferenceIntPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int{&v8}
	actualPtr = DifferenceIntPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsIntPtr(&v2, actualPtr) || !ExistsIntPtr(&v1, actualPtr) {
		t.Errorf("DifferenceIntPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	actual := DifferenceInt64()
	if len(actual) != 0 {
		t.Errorf("DifferenceInt64 failed.")
	}

	expected := []int64{v8, v2, v1}
	actual = DifferenceInt64(expected)

	if len(actual) != 3 || !ExistsInt64(v8, actual) || !ExistsInt64(v2, actual) || !ExistsInt64(v1, actual) {
		t.Errorf("DifferenceInt64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int64{v8, v2, v1, v1, v2, v8}
	list2 := []int64{v8}
	actual = DifferenceInt64(list1, list2)
	if len(actual) != 2 || !ExistsInt64(v2, actual) || !ExistsInt64(v1, actual) {
		t.Errorf("DifferenceInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceInt64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceInt64Ptr failed.")
	}

	expectedPtr := []*int64{&v8, &v2, &v1}
	actualPtr = DifferenceInt64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt64Ptr(&v8, actualPtr) || !ExistsInt64Ptr(&v2, actualPtr) || !ExistsInt64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int64{&v8}
	actualPtr = DifferenceInt64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsInt64Ptr(&v2, actualPtr) || !ExistsInt64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	actual := DifferenceInt32()
	if len(actual) != 0 {
		t.Errorf("DifferenceInt32 failed.")
	}

	expected := []int32{v8, v2, v1}
	actual = DifferenceInt32(expected)

	if len(actual) != 3 || !ExistsInt32(v8, actual) || !ExistsInt32(v2, actual) || !ExistsInt32(v1, actual) {
		t.Errorf("DifferenceInt32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int32{v8, v2, v1, v1, v2, v8}
	list2 := []int32{v8}
	actual = DifferenceInt32(list1, list2)
	if len(actual) != 2 || !ExistsInt32(v2, actual) || !ExistsInt32(v1, actual) {
		t.Errorf("DifferenceInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceInt32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceInt32Ptr failed.")
	}

	expectedPtr := []*int32{&v8, &v2, &v1}
	actualPtr = DifferenceInt32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt32Ptr(&v8, actualPtr) || !ExistsInt32Ptr(&v2, actualPtr) || !ExistsInt32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int32{&v8}
	actualPtr = DifferenceInt32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsInt32Ptr(&v2, actualPtr) || !ExistsInt32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	actual := DifferenceInt16()
	if len(actual) != 0 {
		t.Errorf("DifferenceInt16 failed.")
	}

	expected := []int16{v8, v2, v1}
	actual = DifferenceInt16(expected)

	if len(actual) != 3 || !ExistsInt16(v8, actual) || !ExistsInt16(v2, actual) || !ExistsInt16(v1, actual) {
		t.Errorf("DifferenceInt16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int16{v8, v2, v1, v1, v2, v8}
	list2 := []int16{v8}
	actual = DifferenceInt16(list1, list2)
	if len(actual) != 2 || !ExistsInt16(v2, actual) || !ExistsInt16(v1, actual) {
		t.Errorf("DifferenceInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceInt16Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceInt16Ptr failed.")
	}

	expectedPtr := []*int16{&v8, &v2, &v1}
	actualPtr = DifferenceInt16Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt16Ptr(&v8, actualPtr) || !ExistsInt16Ptr(&v2, actualPtr) || !ExistsInt16Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int16{&v8}
	actualPtr = DifferenceInt16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsInt16Ptr(&v2, actualPtr) || !ExistsInt16Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	actual := DifferenceInt8()
	if len(actual) != 0 {
		t.Errorf("DifferenceInt8 failed.")
	}

	expected := []int8{v8, v2, v1}
	actual = DifferenceInt8(expected)

	if len(actual) != 3 || !ExistsInt8(v8, actual) || !ExistsInt8(v2, actual) || !ExistsInt8(v1, actual) {
		t.Errorf("DifferenceInt8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int8{v8, v2, v1, v1, v2, v8}
	list2 := []int8{v8}
	actual = DifferenceInt8(list1, list2)
	if len(actual) != 2 || !ExistsInt8(v2, actual) || !ExistsInt8(v1, actual) {
		t.Errorf("DifferenceInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceInt8Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceInt8Ptr failed.")
	}

	expectedPtr := []*int8{&v8, &v2, &v1}
	actualPtr = DifferenceInt8Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt8Ptr(&v8, actualPtr) || !ExistsInt8Ptr(&v2, actualPtr) || !ExistsInt8Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int8{&v8}
	actualPtr = DifferenceInt8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsInt8Ptr(&v2, actualPtr) || !ExistsInt8Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceInt8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	actual := DifferenceUint()
	if len(actual) != 0 {
		t.Errorf("DifferenceUint failed.")
	}

	expected := []uint{v8, v2, v1}
	actual = DifferenceUint(expected)

	if len(actual) != 3 || !ExistsUint(v8, actual) || !ExistsUint(v2, actual) || !ExistsUint(v1, actual) {
		t.Errorf("DifferenceUint failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint{v8, v2, v1, v1, v2, v8}
	list2 := []uint{v8}
	actual = DifferenceUint(list1, list2)
	if len(actual) != 2 || !ExistsUint(v2, actual) || !ExistsUint(v1, actual) {
		t.Errorf("DifferenceUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceUintPtr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceUintPtr failed.")
	}

	expectedPtr := []*uint{&v8, &v2, &v1}
	actualPtr = DifferenceUintPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUintPtr(&v8, actualPtr) || !ExistsUintPtr(&v2, actualPtr) || !ExistsUintPtr(&v1, actualPtr) {
		t.Errorf("DifferenceUintPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint{&v8}
	actualPtr = DifferenceUintPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsUintPtr(&v2, actualPtr) || !ExistsUintPtr(&v1, actualPtr) {
		t.Errorf("DifferenceUintPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	actual := DifferenceUint64()
	if len(actual) != 0 {
		t.Errorf("DifferenceUint64 failed.")
	}

	expected := []uint64{v8, v2, v1}
	actual = DifferenceUint64(expected)

	if len(actual) != 3 || !ExistsUint64(v8, actual) || !ExistsUint64(v2, actual) || !ExistsUint64(v1, actual) {
		t.Errorf("DifferenceUint64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint64{v8, v2, v1, v1, v2, v8}
	list2 := []uint64{v8}
	actual = DifferenceUint64(list1, list2)
	if len(actual) != 2 || !ExistsUint64(v2, actual) || !ExistsUint64(v1, actual) {
		t.Errorf("DifferenceUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceUint64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceUint64Ptr failed.")
	}

	expectedPtr := []*uint64{&v8, &v2, &v1}
	actualPtr = DifferenceUint64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint64Ptr(&v8, actualPtr) || !ExistsUint64Ptr(&v2, actualPtr) || !ExistsUint64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint64{&v8}
	actualPtr = DifferenceUint64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsUint64Ptr(&v2, actualPtr) || !ExistsUint64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	actual := DifferenceUint32()
	if len(actual) != 0 {
		t.Errorf("DifferenceUint32 failed.")
	}

	expected := []uint32{v8, v2, v1}
	actual = DifferenceUint32(expected)

	if len(actual) != 3 || !ExistsUint32(v8, actual) || !ExistsUint32(v2, actual) || !ExistsUint32(v1, actual) {
		t.Errorf("DifferenceUint32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint32{v8, v2, v1, v1, v2, v8}
	list2 := []uint32{v8}
	actual = DifferenceUint32(list1, list2)
	if len(actual) != 2 || !ExistsUint32(v2, actual) || !ExistsUint32(v1, actual) {
		t.Errorf("DifferenceUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceUint32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceUint32Ptr failed.")
	}

	expectedPtr := []*uint32{&v8, &v2, &v1}
	actualPtr = DifferenceUint32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint32Ptr(&v8, actualPtr) || !ExistsUint32Ptr(&v2, actualPtr) || !ExistsUint32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint32{&v8}
	actualPtr = DifferenceUint32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsUint32Ptr(&v2, actualPtr) || !ExistsUint32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	actual := DifferenceUint16()
	if len(actual) != 0 {
		t.Errorf("DifferenceUint16 failed.")
	}

	expected := []uint16{v8, v2, v1}
	actual = DifferenceUint16(expected)

	if len(actual) != 3 || !ExistsUint16(v8, actual) || !ExistsUint16(v2, actual) || !ExistsUint16(v1, actual) {
		t.Errorf("DifferenceUint16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint16{v8, v2, v1, v1, v2, v8}
	list2 := []uint16{v8}
	actual = DifferenceUint16(list1, list2)
	if len(actual) != 2 || !ExistsUint16(v2, actual) || !ExistsUint16(v1, actual) {
		t.Errorf("DifferenceUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceUint16Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceUint16Ptr failed.")
	}

	expectedPtr := []*uint16{&v8, &v2, &v1}
	actualPtr = DifferenceUint16Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint16Ptr(&v8, actualPtr) || !ExistsUint16Ptr(&v2, actualPtr) || !ExistsUint16Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint16{&v8}
	actualPtr = DifferenceUint16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsUint16Ptr(&v2, actualPtr) || !ExistsUint16Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	actual := DifferenceUint8()
	if len(actual) != 0 {
		t.Errorf("DifferenceUint8 failed.")
	}

	expected := []uint8{v8, v2, v1}
	actual = DifferenceUint8(expected)

	if len(actual) != 3 || !ExistsUint8(v8, actual) || !ExistsUint8(v2, actual) || !ExistsUint8(v1, actual) {
		t.Errorf("DifferenceUint8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint8{v8, v2, v1, v1, v2, v8}
	list2 := []uint8{v8}
	actual = DifferenceUint8(list1, list2)
	if len(actual) != 2 || !ExistsUint8(v2, actual) || !ExistsUint8(v1, actual) {
		t.Errorf("DifferenceUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceUint8Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceUint8Ptr failed.")
	}

	expectedPtr := []*uint8{&v8, &v2, &v1}
	actualPtr = DifferenceUint8Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint8Ptr(&v8, actualPtr) || !ExistsUint8Ptr(&v2, actualPtr) || !ExistsUint8Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint8{&v8}
	actualPtr = DifferenceUint8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsUint8Ptr(&v2, actualPtr) || !ExistsUint8Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceUint8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	actual := DifferenceStr()
	if len(actual) != 0 {
		t.Errorf("DifferenceStr failed.")
	}

	expected := []string{v8, v2, v1}
	actual = DifferenceStr(expected)

	if len(actual) != 3 || !ExistsStr(v8, actual) || !ExistsStr(v2, actual) || !ExistsStr(v1, actual) {
		t.Errorf("DifferenceStr failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []string{v8, v2, v1, v1, v2, v8}
	list2 := []string{v8}
	actual = DifferenceStr(list1, list2)
	if len(actual) != 2 || !ExistsStr(v2, actual) || !ExistsStr(v1, actual) {
		t.Errorf("DifferenceStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceStrPtr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceStrPtr failed.")
	}

	expectedPtr := []*string{&v8, &v2, &v1}
	actualPtr = DifferenceStrPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsStrPtr(&v8, actualPtr) || !ExistsStrPtr(&v2, actualPtr) || !ExistsStrPtr(&v1, actualPtr) {
		t.Errorf("DifferenceStrPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*string{&v8}
	actualPtr = DifferenceStrPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsStrPtr(&v2, actualPtr) || !ExistsStrPtr(&v1, actualPtr) {
		t.Errorf("DifferenceStrPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = true
	var v1 bool = false

	actual := DifferenceBool()
	if len(actual) != 0 {
		t.Errorf("DifferenceBoolram failed.")
	}

	expected := []bool{v8, v2, v1}
	actual = DifferenceBool(expected)

	if len(actual) != 2 || !ExistsBool(v8, actual) || !ExistsBool(v1, actual) {
		t.Errorf("DifferenceBool failed nks0. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []bool{v8, v2, v1, v1, v2, v8}
	list2 := []bool{v1}
	actual = DifferenceBool(list1, list2)
	if len(actual) != 1 || !ExistsBool(v8, actual) {
		t.Errorf("DifferenceBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceBoolPtr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceBoolPtr failed.")
	}

	expectedPtr := []*bool{&v8, &v2, &v1}
	actualPtr = DifferenceBoolPtr(expectedPtr)
	
	if len(actualPtr) != 2 || !ExistsBoolPtr(&v8, actualPtr) || !ExistsBoolPtr(&v1, actualPtr) {
		t.Errorf("DifferenceBoolPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*bool{&v1}
	actualPtr = DifferenceBoolPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsBoolPtr(&v8, actualPtr) {
		t.Errorf("DifferenceBoolPtrNks failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	actual := DifferenceFloat32()
	if len(actual) != 0 {
		t.Errorf("DifferenceFloat32 failed.")
	}

	expected := []float32{v8, v2, v1}
	actual = DifferenceFloat32(expected)

	if len(actual) != 3 || !ExistsFloat32(v8, actual) || !ExistsFloat32(v2, actual) || !ExistsFloat32(v1, actual) {
		t.Errorf("DifferenceFloat32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []float32{v8, v2, v1, v1, v2, v8}
	list2 := []float32{v8}
	actual = DifferenceFloat32(list1, list2)
	if len(actual) != 2 || !ExistsFloat32(v2, actual) || !ExistsFloat32(v1, actual) {
		t.Errorf("DifferenceFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceFloat32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceFloat32Ptr failed.")
	}

	expectedPtr := []*float32{&v8, &v2, &v1}
	actualPtr = DifferenceFloat32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsFloat32Ptr(&v8, actualPtr) || !ExistsFloat32Ptr(&v2, actualPtr) || !ExistsFloat32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceFloat32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float32{&v8}
	actualPtr = DifferenceFloat32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsFloat32Ptr(&v2, actualPtr) || !ExistsFloat32Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceFloat32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestDifferenceFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	actual := DifferenceFloat64()
	if len(actual) != 0 {
		t.Errorf("DifferenceFloat64 failed.")
	}

	expected := []float64{v8, v2, v1}
	actual = DifferenceFloat64(expected)

	if len(actual) != 3 || !ExistsFloat64(v8, actual) || !ExistsFloat64(v2, actual) || !ExistsFloat64(v1, actual) {
		t.Errorf("DifferenceFloat64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []float64{v8, v2, v1, v1, v2, v8}
	list2 := []float64{v8}
	actual = DifferenceFloat64(list1, list2)
	if len(actual) != 2 || !ExistsFloat64(v2, actual) || !ExistsFloat64(v1, actual) {
		t.Errorf("DifferenceFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := DifferenceFloat64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("DifferenceFloat64Ptr failed.")
	}

	expectedPtr := []*float64{&v8, &v2, &v1}
	actualPtr = DifferenceFloat64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsFloat64Ptr(&v8, actualPtr) || !ExistsFloat64Ptr(&v2, actualPtr) || !ExistsFloat64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceFloat64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float64{&v8}
	actualPtr = DifferenceFloat64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsFloat64Ptr(&v2, actualPtr) || !ExistsFloat64Ptr(&v1, actualPtr) {
		t.Errorf("DifferenceFloat64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
