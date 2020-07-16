package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestIntersectionInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	actual := IntersectionInt()
	if len(actual) != 0 {
		t.Errorf("IntersectionInt failed.")
	}

	expected := []int{v8, v2, v1}
	actual = IntersectionInt(expected)

	if len(actual) != 3 || !ExistsInt(v8, actual) || !ExistsInt(v2, actual) || !ExistsInt(v1, actual) {
		t.Errorf("IntersectionInt failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int{v8, v2, v1, v1, v2, v8}
	list2 := []int{v8}
	actual = IntersectionInt(list1, list2)
	if len(actual) != 1 || !ExistsInt(v8, actual) {
		t.Errorf("IntersectionInt failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionIntPtr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionIntPtr failed.")
	}

	expectedPtr := []*int{&v8, &v2, &v1}
	actualPtr = IntersectionIntPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsIntPtr(&v8, actualPtr) || !ExistsIntPtr(&v2, actualPtr) || !ExistsIntPtr(&v1, actualPtr) {
		t.Errorf("IntersectionIntPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int{&v8}
	actualPtr = IntersectionIntPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsIntPtr(&v8, actualPtr) {
		t.Errorf("IntersectionIntPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	actual := IntersectionInt64()
	if len(actual) != 0 {
		t.Errorf("IntersectionInt64 failed.")
	}

	expected := []int64{v8, v2, v1}
	actual = IntersectionInt64(expected)

	if len(actual) != 3 || !ExistsInt64(v8, actual) || !ExistsInt64(v2, actual) || !ExistsInt64(v1, actual) {
		t.Errorf("IntersectionInt64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int64{v8, v2, v1, v1, v2, v8}
	list2 := []int64{v8}
	actual = IntersectionInt64(list1, list2)
	if len(actual) != 1 || !ExistsInt64(v8, actual) {
		t.Errorf("IntersectionInt64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionInt64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionInt64Ptr failed.")
	}

	expectedPtr := []*int64{&v8, &v2, &v1}
	actualPtr = IntersectionInt64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt64Ptr(&v8, actualPtr) || !ExistsInt64Ptr(&v2, actualPtr) || !ExistsInt64Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionInt64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int64{&v8}
	actualPtr = IntersectionInt64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsInt64Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionInt64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	actual := IntersectionInt32()
	if len(actual) != 0 {
		t.Errorf("IntersectionInt32 failed.")
	}

	expected := []int32{v8, v2, v1}
	actual = IntersectionInt32(expected)

	if len(actual) != 3 || !ExistsInt32(v8, actual) || !ExistsInt32(v2, actual) || !ExistsInt32(v1, actual) {
		t.Errorf("IntersectionInt32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int32{v8, v2, v1, v1, v2, v8}
	list2 := []int32{v8}
	actual = IntersectionInt32(list1, list2)
	if len(actual) != 1 || !ExistsInt32(v8, actual) {
		t.Errorf("IntersectionInt32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionInt32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionInt32Ptr failed.")
	}

	expectedPtr := []*int32{&v8, &v2, &v1}
	actualPtr = IntersectionInt32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt32Ptr(&v8, actualPtr) || !ExistsInt32Ptr(&v2, actualPtr) || !ExistsInt32Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionInt32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int32{&v8}
	actualPtr = IntersectionInt32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsInt32Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionInt32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	actual := IntersectionInt16()
	if len(actual) != 0 {
		t.Errorf("IntersectionInt16 failed.")
	}

	expected := []int16{v8, v2, v1}
	actual = IntersectionInt16(expected)

	if len(actual) != 3 || !ExistsInt16(v8, actual) || !ExistsInt16(v2, actual) || !ExistsInt16(v1, actual) {
		t.Errorf("IntersectionInt16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int16{v8, v2, v1, v1, v2, v8}
	list2 := []int16{v8}
	actual = IntersectionInt16(list1, list2)
	if len(actual) != 1 || !ExistsInt16(v8, actual) {
		t.Errorf("IntersectionInt16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionInt16Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionInt16Ptr failed.")
	}

	expectedPtr := []*int16{&v8, &v2, &v1}
	actualPtr = IntersectionInt16Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt16Ptr(&v8, actualPtr) || !ExistsInt16Ptr(&v2, actualPtr) || !ExistsInt16Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionInt16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int16{&v8}
	actualPtr = IntersectionInt16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsInt16Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionInt16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	actual := IntersectionInt8()
	if len(actual) != 0 {
		t.Errorf("IntersectionInt8 failed.")
	}

	expected := []int8{v8, v2, v1}
	actual = IntersectionInt8(expected)

	if len(actual) != 3 || !ExistsInt8(v8, actual) || !ExistsInt8(v2, actual) || !ExistsInt8(v1, actual) {
		t.Errorf("IntersectionInt8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []int8{v8, v2, v1, v1, v2, v8}
	list2 := []int8{v8}
	actual = IntersectionInt8(list1, list2)
	if len(actual) != 1 || !ExistsInt8(v8, actual) {
		t.Errorf("IntersectionInt8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionInt8Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionInt8Ptr failed.")
	}

	expectedPtr := []*int8{&v8, &v2, &v1}
	actualPtr = IntersectionInt8Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsInt8Ptr(&v8, actualPtr) || !ExistsInt8Ptr(&v2, actualPtr) || !ExistsInt8Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionInt8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int8{&v8}
	actualPtr = IntersectionInt8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsInt8Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionInt8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	actual := IntersectionUint()
	if len(actual) != 0 {
		t.Errorf("IntersectionUint failed.")
	}

	expected := []uint{v8, v2, v1}
	actual = IntersectionUint(expected)

	if len(actual) != 3 || !ExistsUint(v8, actual) || !ExistsUint(v2, actual) || !ExistsUint(v1, actual) {
		t.Errorf("IntersectionUint failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint{v8, v2, v1, v1, v2, v8}
	list2 := []uint{v8}
	actual = IntersectionUint(list1, list2)
	if len(actual) != 1 || !ExistsUint(v8, actual) {
		t.Errorf("IntersectionUint failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionUintPtr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionUintPtr failed.")
	}

	expectedPtr := []*uint{&v8, &v2, &v1}
	actualPtr = IntersectionUintPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUintPtr(&v8, actualPtr) || !ExistsUintPtr(&v2, actualPtr) || !ExistsUintPtr(&v1, actualPtr) {
		t.Errorf("IntersectionUintPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint{&v8}
	actualPtr = IntersectionUintPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsUintPtr(&v8, actualPtr) {
		t.Errorf("IntersectionUintPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	actual := IntersectionUint64()
	if len(actual) != 0 {
		t.Errorf("IntersectionUint64 failed.")
	}

	expected := []uint64{v8, v2, v1}
	actual = IntersectionUint64(expected)

	if len(actual) != 3 || !ExistsUint64(v8, actual) || !ExistsUint64(v2, actual) || !ExistsUint64(v1, actual) {
		t.Errorf("IntersectionUint64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint64{v8, v2, v1, v1, v2, v8}
	list2 := []uint64{v8}
	actual = IntersectionUint64(list1, list2)
	if len(actual) != 1 || !ExistsUint64(v8, actual) {
		t.Errorf("IntersectionUint64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionUint64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionUint64Ptr failed.")
	}

	expectedPtr := []*uint64{&v8, &v2, &v1}
	actualPtr = IntersectionUint64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint64Ptr(&v8, actualPtr) || !ExistsUint64Ptr(&v2, actualPtr) || !ExistsUint64Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionUint64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint64{&v8}
	actualPtr = IntersectionUint64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsUint64Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionUint64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	actual := IntersectionUint32()
	if len(actual) != 0 {
		t.Errorf("IntersectionUint32 failed.")
	}

	expected := []uint32{v8, v2, v1}
	actual = IntersectionUint32(expected)

	if len(actual) != 3 || !ExistsUint32(v8, actual) || !ExistsUint32(v2, actual) || !ExistsUint32(v1, actual) {
		t.Errorf("IntersectionUint32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint32{v8, v2, v1, v1, v2, v8}
	list2 := []uint32{v8}
	actual = IntersectionUint32(list1, list2)
	if len(actual) != 1 || !ExistsUint32(v8, actual) {
		t.Errorf("IntersectionUint32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionUint32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionUint32Ptr failed.")
	}

	expectedPtr := []*uint32{&v8, &v2, &v1}
	actualPtr = IntersectionUint32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint32Ptr(&v8, actualPtr) || !ExistsUint32Ptr(&v2, actualPtr) || !ExistsUint32Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionUint32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint32{&v8}
	actualPtr = IntersectionUint32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsUint32Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionUint32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	actual := IntersectionUint16()
	if len(actual) != 0 {
		t.Errorf("IntersectionUint16 failed.")
	}

	expected := []uint16{v8, v2, v1}
	actual = IntersectionUint16(expected)

	if len(actual) != 3 || !ExistsUint16(v8, actual) || !ExistsUint16(v2, actual) || !ExistsUint16(v1, actual) {
		t.Errorf("IntersectionUint16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint16{v8, v2, v1, v1, v2, v8}
	list2 := []uint16{v8}
	actual = IntersectionUint16(list1, list2)
	if len(actual) != 1 || !ExistsUint16(v8, actual) {
		t.Errorf("IntersectionUint16 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionUint16Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionUint16Ptr failed.")
	}

	expectedPtr := []*uint16{&v8, &v2, &v1}
	actualPtr = IntersectionUint16Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint16Ptr(&v8, actualPtr) || !ExistsUint16Ptr(&v2, actualPtr) || !ExistsUint16Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionUint16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint16{&v8}
	actualPtr = IntersectionUint16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsUint16Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionUint16Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	actual := IntersectionUint8()
	if len(actual) != 0 {
		t.Errorf("IntersectionUint8 failed.")
	}

	expected := []uint8{v8, v2, v1}
	actual = IntersectionUint8(expected)

	if len(actual) != 3 || !ExistsUint8(v8, actual) || !ExistsUint8(v2, actual) || !ExistsUint8(v1, actual) {
		t.Errorf("IntersectionUint8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []uint8{v8, v2, v1, v1, v2, v8}
	list2 := []uint8{v8}
	actual = IntersectionUint8(list1, list2)
	if len(actual) != 1 || !ExistsUint8(v8, actual) {
		t.Errorf("IntersectionUint8 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionUint8Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionUint8Ptr failed.")
	}

	expectedPtr := []*uint8{&v8, &v2, &v1}
	actualPtr = IntersectionUint8Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsUint8Ptr(&v8, actualPtr) || !ExistsUint8Ptr(&v2, actualPtr) || !ExistsUint8Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionUint8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint8{&v8}
	actualPtr = IntersectionUint8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsUint8Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionUint8Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	actual := IntersectionStr()
	if len(actual) != 0 {
		t.Errorf("IntersectionStr failed.")
	}

	expected := []string{v8, v2, v1}
	actual = IntersectionStr(expected)

	if len(actual) != 3 || !ExistsStr(v8, actual) || !ExistsStr(v2, actual) || !ExistsStr(v1, actual) {
		t.Errorf("IntersectionStr failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []string{v8, v2, v1, v1, v2, v8}
	list2 := []string{v8}
	actual = IntersectionStr(list1, list2)
	if len(actual) != 1 || !ExistsStr(v8, actual) {
		t.Errorf("IntersectionStr failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionStrPtr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionStrPtr failed.")
	}

	expectedPtr := []*string{&v8, &v2, &v1}
	actualPtr = IntersectionStrPtr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsStrPtr(&v8, actualPtr) || !ExistsStrPtr(&v2, actualPtr) || !ExistsStrPtr(&v1, actualPtr) {
		t.Errorf("IntersectionStrPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*string{&v8}
	actualPtr = IntersectionStrPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsStrPtr(&v8, actualPtr) {
		t.Errorf("IntersectionStrPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = true
	var v1 bool = false

	actual := IntersectionBool()
	if len(actual) != 0 {
		t.Errorf("IntersectionBool failed.")
	}

	expected := []bool{v8, v2, v1}
	actual = IntersectionBool(expected)

	if len(actual) != 2 || !ExistsBool(v8, actual) || !ExistsBool(v1, actual) {
		t.Errorf("IntersectionBool failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []bool{v8, v2, v1, v1, v2, v8}
	list2 := []bool{v8}
	actual = IntersectionBool(list1, list2)
	if len(actual) != 1 || !ExistsBool(v8, actual) {
		t.Errorf("IntersectionBool failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionBoolPtr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionBoolPtr failed.")
	}

	expectedPtr := []*bool{&v8, &v2, &v1}
	actualPtr = IntersectionBoolPtr(expectedPtr)

	if len(actualPtr) != 2 || !ExistsBoolPtr(&v8, actualPtr) || !ExistsBoolPtr(&v1, actualPtr) {
		t.Errorf("IntersectionBoolPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*bool{&v8}
	actualPtr = IntersectionBoolPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsBoolPtr(&v8, actualPtr) {
		t.Errorf("IntersectionBoolPtr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	actual := IntersectionFloat32()
	if len(actual) != 0 {
		t.Errorf("IntersectionFloat32 failed.")
	}

	expected := []float32{v8, v2, v1}
	actual = IntersectionFloat32(expected)

	if len(actual) != 3 || !ExistsFloat32(v8, actual) || !ExistsFloat32(v2, actual) || !ExistsFloat32(v1, actual) {
		t.Errorf("IntersectionFloat32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []float32{v8, v2, v1, v1, v2, v8}
	list2 := []float32{v8}
	actual = IntersectionFloat32(list1, list2)
	if len(actual) != 1 || !ExistsFloat32(v8, actual) {
		t.Errorf("IntersectionFloat32 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionFloat32Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionFloat32Ptr failed.")
	}

	expectedPtr := []*float32{&v8, &v2, &v1}
	actualPtr = IntersectionFloat32Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsFloat32Ptr(&v8, actualPtr) || !ExistsFloat32Ptr(&v2, actualPtr) || !ExistsFloat32Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionFloat32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float32{&v8}
	actualPtr = IntersectionFloat32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsFloat32Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionFloat32Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}

func TestIntersectionFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	actual := IntersectionFloat64()
	if len(actual) != 0 {
		t.Errorf("IntersectionFloat64 failed.")
	}

	expected := []float64{v8, v2, v1}
	actual = IntersectionFloat64(expected)

	if len(actual) != 3 || !ExistsFloat64(v8, actual) || !ExistsFloat64(v2, actual) || !ExistsFloat64(v1, actual) {
		t.Errorf("IntersectionFloat64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []float64{v8, v2, v1, v1, v2, v8}
	list2 := []float64{v8}
	actual = IntersectionFloat64(list1, list2)
	if len(actual) != 1 || !ExistsFloat64(v8, actual) {
		t.Errorf("IntersectionFloat64 failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := IntersectionFloat64Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("IntersectionFloat64Ptr failed.")
	}

	expectedPtr := []*float64{&v8, &v2, &v1}
	actualPtr = IntersectionFloat64Ptr(expectedPtr)

	if len(actualPtr) != 3 || !ExistsFloat64Ptr(&v8, actualPtr) || !ExistsFloat64Ptr(&v2, actualPtr) || !ExistsFloat64Ptr(&v1, actualPtr) {
		t.Errorf("IntersectionFloat64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float64{&v8}
	actualPtr = IntersectionFloat64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !ExistsFloat64Ptr(&v8, actualPtr) {
		t.Errorf("IntersectionFloat64Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
