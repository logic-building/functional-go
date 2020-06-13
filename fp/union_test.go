package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestUnionInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	expected := []int{v8, v2, v1}
	list1 := []int{v8, v2, v1, v1, v2, v8}
	list2 := []int{v8, v2, v1, v1, v2, v8}
	actual := UnionInt(list1, list2)
	if len(actual) != 3 || !ExistsInt(v8, actual) || !ExistsInt(v2, actual) || !ExistsInt(v1, actual) {
		t.Errorf("UnionInt failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionIntPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsIntPtr(&v8, actualPtr) || !ExistsIntPtr(&v2, actualPtr) || !ExistsIntPtr(&v1, actualPtr) {
		t.Errorf("UnionIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	expected := []int64{v8, v2, v1}
	list1 := []int64{v8, v2, v1, v1, v2, v8}
	list2 := []int64{v8, v2, v1, v1, v2, v8}
	actual := UnionInt64(list1, list2)
	if len(actual) != 3 || !ExistsInt64(v8, actual) || !ExistsInt64(v2, actual) || !ExistsInt64(v1, actual) {
		t.Errorf("UnionInt64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionInt64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsInt64Ptr(&v8, actualPtr) || !ExistsInt64Ptr(&v2, actualPtr) || !ExistsInt64Ptr(&v1, actualPtr) {
		t.Errorf("UnionInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	expected := []int32{v8, v2, v1}
	list1 := []int32{v8, v2, v1, v1, v2, v8}
	list2 := []int32{v8, v2, v1, v1, v2, v8}
	actual := UnionInt32(list1, list2)
	if len(actual) != 3 || !ExistsInt32(v8, actual) || !ExistsInt32(v2, actual) || !ExistsInt32(v1, actual) {
		t.Errorf("UnionInt32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionInt32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsInt32Ptr(&v8, actualPtr) || !ExistsInt32Ptr(&v2, actualPtr) || !ExistsInt32Ptr(&v1, actualPtr) {
		t.Errorf("UnionInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	expected := []int16{v8, v2, v1}
	list1 := []int16{v8, v2, v1, v1, v2, v8}
	list2 := []int16{v8, v2, v1, v1, v2, v8}
	actual := UnionInt16(list1, list2)
	if len(actual) != 3 || !ExistsInt16(v8, actual) || !ExistsInt16(v2, actual) || !ExistsInt16(v1, actual) {
		t.Errorf("UnionInt16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionInt16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsInt16Ptr(&v8, actualPtr) || !ExistsInt16Ptr(&v2, actualPtr) || !ExistsInt16Ptr(&v1, actualPtr) {
		t.Errorf("UnionInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	expected := []int8{v8, v2, v1}
	list1 := []int8{v8, v2, v1, v1, v2, v8}
	list2 := []int8{v8, v2, v1, v1, v2, v8}
	actual := UnionInt8(list1, list2)
	if len(actual) != 3 || !ExistsInt8(v8, actual) || !ExistsInt8(v2, actual) || !ExistsInt8(v1, actual) {
		t.Errorf("UnionInt8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionInt8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsInt8Ptr(&v8, actualPtr) || !ExistsInt8Ptr(&v2, actualPtr) || !ExistsInt8Ptr(&v1, actualPtr) {
		t.Errorf("UnionInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	expected := []uint{v8, v2, v1}
	list1 := []uint{v8, v2, v1, v1, v2, v8}
	list2 := []uint{v8, v2, v1, v1, v2, v8}
	actual := UnionUint(list1, list2)
	if len(actual) != 3 || !ExistsUint(v8, actual) || !ExistsUint(v2, actual) || !ExistsUint(v1, actual) {
		t.Errorf("UnionUint failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionUintPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsUintPtr(&v8, actualPtr) || !ExistsUintPtr(&v2, actualPtr) || !ExistsUintPtr(&v1, actualPtr) {
		t.Errorf("UnionUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	expected := []uint64{v8, v2, v1}
	list1 := []uint64{v8, v2, v1, v1, v2, v8}
	list2 := []uint64{v8, v2, v1, v1, v2, v8}
	actual := UnionUint64(list1, list2)
	if len(actual) != 3 || !ExistsUint64(v8, actual) || !ExistsUint64(v2, actual) || !ExistsUint64(v1, actual) {
		t.Errorf("UnionUint64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionUint64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsUint64Ptr(&v8, actualPtr) || !ExistsUint64Ptr(&v2, actualPtr) || !ExistsUint64Ptr(&v1, actualPtr) {
		t.Errorf("UnionUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	expected := []uint32{v8, v2, v1}
	list1 := []uint32{v8, v2, v1, v1, v2, v8}
	list2 := []uint32{v8, v2, v1, v1, v2, v8}
	actual := UnionUint32(list1, list2)
	if len(actual) != 3 || !ExistsUint32(v8, actual) || !ExistsUint32(v2, actual) || !ExistsUint32(v1, actual) {
		t.Errorf("UnionUint32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionUint32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsUint32Ptr(&v8, actualPtr) || !ExistsUint32Ptr(&v2, actualPtr) || !ExistsUint32Ptr(&v1, actualPtr) {
		t.Errorf("UnionUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	expected := []uint16{v8, v2, v1}
	list1 := []uint16{v8, v2, v1, v1, v2, v8}
	list2 := []uint16{v8, v2, v1, v1, v2, v8}
	actual := UnionUint16(list1, list2)
	if len(actual) != 3 || !ExistsUint16(v8, actual) || !ExistsUint16(v2, actual) || !ExistsUint16(v1, actual) {
		t.Errorf("UnionUint16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionUint16Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsUint16Ptr(&v8, actualPtr) || !ExistsUint16Ptr(&v2, actualPtr) || !ExistsUint16Ptr(&v1, actualPtr) {
		t.Errorf("UnionUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	expected := []uint8{v8, v2, v1}
	list1 := []uint8{v8, v2, v1, v1, v2, v8}
	list2 := []uint8{v8, v2, v1, v1, v2, v8}
	actual := UnionUint8(list1, list2)
	if len(actual) != 3 || !ExistsUint8(v8, actual) || !ExistsUint8(v2, actual) || !ExistsUint8(v1, actual) {
		t.Errorf("UnionUint8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionUint8Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsUint8Ptr(&v8, actualPtr) || !ExistsUint8Ptr(&v2, actualPtr) || !ExistsUint8Ptr(&v1, actualPtr) {
		t.Errorf("UnionUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	expected := []string{v8, v2, v1}
	list1 := []string{v8, v2, v1, v1, v2, v8}
	list2 := []string{v8, v2, v1, v1, v2, v8}
	actual := UnionStr(list1, list2)
	if len(actual) != 3 || !ExistsStr(v8, actual) || !ExistsStr(v2, actual) || !ExistsStr(v1, actual) {
		t.Errorf("UnionStr failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionStrPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsStrPtr(&v8, actualPtr) || !ExistsStrPtr(&v2, actualPtr) || !ExistsStrPtr(&v1, actualPtr) {
		t.Errorf("UnionStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = false
	var v0 bool = true

	expected := []bool{v8, v2}
	list1 := []bool{v8, v2, v0, v0, v2, v8}
	list2 := []bool{v8, v2, v0, v0, v2, v8}
	actual := UnionBool(list1, list2)
	if len(actual) != 2 || !ExistsBool(v8, actual) || !ExistsBool(v2, actual) {
		t.Errorf("UnionBool failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*bool{&v8, &v2, &v0, &v0, &v2, &v8}
	list2Ptr := []*bool{&v8, &v2, &v0, &v0, &v2, &v8}
	actualPtr := UnionBoolPtr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !ExistsBoolPtr(&v8, actualPtr) || !ExistsBoolPtr(&v2, actualPtr) {
		t.Errorf("UnionBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	expected := []float32{v8, v2, v1}
	list1 := []float32{v8, v2, v1, v1, v2, v8}
	list2 := []float32{v8, v2, v1, v1, v2, v8}
	actual := UnionFloat32(list1, list2)
	if len(actual) != 3 || !ExistsFloat32(v8, actual) || !ExistsFloat32(v2, actual) || !ExistsFloat32(v1, actual) {
		t.Errorf("UnionFloat32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionFloat32Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsFloat32Ptr(&v8, actualPtr) || !ExistsFloat32Ptr(&v2, actualPtr) || !ExistsFloat32Ptr(&v1, actualPtr) {
		t.Errorf("UnionFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestUnionFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	expected := []float64{v8, v2, v1}
	list1 := []float64{v8, v2, v1, v1, v2, v8}
	list2 := []float64{v8, v2, v1, v1, v2, v8}
	actual := UnionFloat64(list1, list2)
	if len(actual) != 3 || !ExistsFloat64(v8, actual) || !ExistsFloat64(v2, actual) || !ExistsFloat64(v1, actual) {
		t.Errorf("UnionFloat64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := UnionFloat64Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !ExistsFloat64Ptr(&v8, actualPtr) || !ExistsFloat64Ptr(&v2, actualPtr) || !ExistsFloat64Ptr(&v1, actualPtr) {
		t.Errorf("UnionFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
