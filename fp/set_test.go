package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSetInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	expected := []int{v8, v2, v1}
	list := []int{v8, v2, v1, v1, v2, v8}
	actual := SetInt(list)
	if len(actual) != 3 || !ExistsInt(v8, actual) || !ExistsInt(v2, actual) || !ExistsInt(v1, actual) {
		t.Errorf("UnionInt failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetInt(nil)
	if len(actual) != 0 {
		t.Errorf("SetInt failed")
	}

	actualPtr2 := SetInt([]int{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetInt failed")
	}

	actualPtr3 := SetIntPtr([]*int{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetIntPtr failed")
	}

	listPtr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetIntPtr(listPtr)
	if len(actualPtr) != 3 || !ExistsIntPtr(&v8, actualPtr) || !ExistsIntPtr(&v2, actualPtr) || !ExistsIntPtr(&v1, actualPtr) {
		t.Errorf("SetIntPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	expected := []int64{v8, v2, v1}
	list := []int64{v8, v2, v1, v1, v2, v8}
	actual := SetInt64(list)
	if len(actual) != 3 || !ExistsInt64(v8, actual) || !ExistsInt64(v2, actual) || !ExistsInt64(v1, actual) {
		t.Errorf("UnionInt64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetInt64(nil)
	if len(actual) != 0 {
		t.Errorf("SetInt64 failed")
	}

	actualPtr2 := SetInt64([]int64{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetInt64 failed")
	}

	actualPtr3 := SetInt64Ptr([]*int64{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetInt64Ptr failed")
	}

	listPtr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetInt64Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsInt64Ptr(&v8, actualPtr) || !ExistsInt64Ptr(&v2, actualPtr) || !ExistsInt64Ptr(&v1, actualPtr) {
		t.Errorf("SetInt64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	expected := []int32{v8, v2, v1}
	list := []int32{v8, v2, v1, v1, v2, v8}
	actual := SetInt32(list)
	if len(actual) != 3 || !ExistsInt32(v8, actual) || !ExistsInt32(v2, actual) || !ExistsInt32(v1, actual) {
		t.Errorf("UnionInt32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetInt32(nil)
	if len(actual) != 0 {
		t.Errorf("SetInt32 failed")
	}

	actualPtr2 := SetInt32([]int32{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetInt32 failed")
	}

	actualPtr3 := SetInt32Ptr([]*int32{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetInt32Ptr failed")
	}

	listPtr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetInt32Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsInt32Ptr(&v8, actualPtr) || !ExistsInt32Ptr(&v2, actualPtr) || !ExistsInt32Ptr(&v1, actualPtr) {
		t.Errorf("SetInt32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	expected := []int16{v8, v2, v1}
	list := []int16{v8, v2, v1, v1, v2, v8}
	actual := SetInt16(list)
	if len(actual) != 3 || !ExistsInt16(v8, actual) || !ExistsInt16(v2, actual) || !ExistsInt16(v1, actual) {
		t.Errorf("UnionInt16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetInt16(nil)
	if len(actual) != 0 {
		t.Errorf("SetInt16 failed")
	}

	actualPtr2 := SetInt16([]int16{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetInt16 failed")
	}

	actualPtr3 := SetInt16Ptr([]*int16{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetInt16Ptr failed")
	}

	listPtr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetInt16Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsInt16Ptr(&v8, actualPtr) || !ExistsInt16Ptr(&v2, actualPtr) || !ExistsInt16Ptr(&v1, actualPtr) {
		t.Errorf("SetInt16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	expected := []int8{v8, v2, v1}
	list := []int8{v8, v2, v1, v1, v2, v8}
	actual := SetInt8(list)
	if len(actual) != 3 || !ExistsInt8(v8, actual) || !ExistsInt8(v2, actual) || !ExistsInt8(v1, actual) {
		t.Errorf("UnionInt8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetInt8(nil)
	if len(actual) != 0 {
		t.Errorf("SetInt8 failed")
	}

	actualPtr2 := SetInt8([]int8{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetInt8 failed")
	}

	actualPtr3 := SetInt8Ptr([]*int8{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetInt8Ptr failed")
	}

	listPtr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetInt8Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsInt8Ptr(&v8, actualPtr) || !ExistsInt8Ptr(&v2, actualPtr) || !ExistsInt8Ptr(&v1, actualPtr) {
		t.Errorf("SetInt8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	expected := []uint{v8, v2, v1}
	list := []uint{v8, v2, v1, v1, v2, v8}
	actual := SetUint(list)
	if len(actual) != 3 || !ExistsUint(v8, actual) || !ExistsUint(v2, actual) || !ExistsUint(v1, actual) {
		t.Errorf("UnionUint failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetUint(nil)
	if len(actual) != 0 {
		t.Errorf("SetUint failed")
	}

	actualPtr2 := SetUint([]uint{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetUint failed")
	}

	actualPtr3 := SetUintPtr([]*uint{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetUintPtr failed")
	}

	listPtr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetUintPtr(listPtr)
	if len(actualPtr) != 3 || !ExistsUintPtr(&v8, actualPtr) || !ExistsUintPtr(&v2, actualPtr) || !ExistsUintPtr(&v1, actualPtr) {
		t.Errorf("SetUintPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	expected := []uint64{v8, v2, v1}
	list := []uint64{v8, v2, v1, v1, v2, v8}
	actual := SetUint64(list)
	if len(actual) != 3 || !ExistsUint64(v8, actual) || !ExistsUint64(v2, actual) || !ExistsUint64(v1, actual) {
		t.Errorf("UnionUint64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetUint64(nil)
	if len(actual) != 0 {
		t.Errorf("SetUint64 failed")
	}

	actualPtr2 := SetUint64([]uint64{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetUint64 failed")
	}

	actualPtr3 := SetUint64Ptr([]*uint64{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetUint64Ptr failed")
	}

	listPtr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetUint64Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsUint64Ptr(&v8, actualPtr) || !ExistsUint64Ptr(&v2, actualPtr) || !ExistsUint64Ptr(&v1, actualPtr) {
		t.Errorf("SetUint64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	expected := []uint32{v8, v2, v1}
	list := []uint32{v8, v2, v1, v1, v2, v8}
	actual := SetUint32(list)
	if len(actual) != 3 || !ExistsUint32(v8, actual) || !ExistsUint32(v2, actual) || !ExistsUint32(v1, actual) {
		t.Errorf("UnionUint32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetUint32(nil)
	if len(actual) != 0 {
		t.Errorf("SetUint32 failed")
	}

	actualPtr2 := SetUint32([]uint32{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetUint32 failed")
	}

	actualPtr3 := SetUint32Ptr([]*uint32{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetUint32Ptr failed")
	}

	listPtr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetUint32Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsUint32Ptr(&v8, actualPtr) || !ExistsUint32Ptr(&v2, actualPtr) || !ExistsUint32Ptr(&v1, actualPtr) {
		t.Errorf("SetUint32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	expected := []uint16{v8, v2, v1}
	list := []uint16{v8, v2, v1, v1, v2, v8}
	actual := SetUint16(list)
	if len(actual) != 3 || !ExistsUint16(v8, actual) || !ExistsUint16(v2, actual) || !ExistsUint16(v1, actual) {
		t.Errorf("UnionUint16 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetUint16(nil)
	if len(actual) != 0 {
		t.Errorf("SetUint16 failed")
	}

	actualPtr2 := SetUint16([]uint16{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetUint16 failed")
	}

	actualPtr3 := SetUint16Ptr([]*uint16{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetUint16Ptr failed")
	}

	listPtr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetUint16Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsUint16Ptr(&v8, actualPtr) || !ExistsUint16Ptr(&v2, actualPtr) || !ExistsUint16Ptr(&v1, actualPtr) {
		t.Errorf("SetUint16Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	expected := []uint8{v8, v2, v1}
	list := []uint8{v8, v2, v1, v1, v2, v8}
	actual := SetUint8(list)
	if len(actual) != 3 || !ExistsUint8(v8, actual) || !ExistsUint8(v2, actual) || !ExistsUint8(v1, actual) {
		t.Errorf("UnionUint8 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetUint8(nil)
	if len(actual) != 0 {
		t.Errorf("SetUint8 failed")
	}

	actualPtr2 := SetUint8([]uint8{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetUint8 failed")
	}

	actualPtr3 := SetUint8Ptr([]*uint8{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetUint8Ptr failed")
	}

	listPtr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetUint8Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsUint8Ptr(&v8, actualPtr) || !ExistsUint8Ptr(&v2, actualPtr) || !ExistsUint8Ptr(&v1, actualPtr) {
		t.Errorf("SetUint8Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	expected := []string{v8, v2, v1}
	list := []string{v8, v2, v1, v1, v2, v8}
	actual := SetStr(list)
	if len(actual) != 3 || !ExistsStr(v8, actual) || !ExistsStr(v2, actual) || !ExistsStr(v1, actual) {
		t.Errorf("UnionStr failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetStr(nil)
	if len(actual) != 0 {
		t.Errorf("SetStr failed")
	}

	actualPtr2 := SetStr([]string{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetStr failed")
	}

	actualPtr3 := SetStrPtr([]*string{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetStrPtr failed")
	}

	listPtr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetStrPtr(listPtr)
	if len(actualPtr) != 3 || !ExistsStrPtr(&v8, actualPtr) || !ExistsStrPtr(&v2, actualPtr) || !ExistsStrPtr(&v1, actualPtr) {
		t.Errorf("SetStrPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = false
	var v1 bool = true

	expected := []bool{v8, v2}
	list := []bool{v8, v2, v1, v1, v2, v8}
	actual := SetBool(list)
	if len(actual) != 2 || !ExistsBool(v8, actual) || !ExistsBool(v2, actual) || !ExistsBool(v1, actual) {
		t.Errorf("SetBool failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetBool(nil)
	if len(actual) != 0 {
		t.Errorf("SetBool failed")
	}

	actualPtr2 := SetBool([]bool{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetBool failed")
	}

	actualPtr3 := SetBoolPtr([]*bool{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetBool failed")
	}

	listPtr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetBoolPtr(listPtr)
	if len(actualPtr) != 2 || !ExistsBoolPtr(&v8, actualPtr) || !ExistsBoolPtr(&v2, actualPtr) {
		t.Errorf("SetBoolPtr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	expected := []float32{v8, v2, v1}
	list := []float32{v8, v2, v1, v1, v2, v8}
	actual := SetFloat32(list)
	if len(actual) != 3 || !ExistsFloat32(v8, actual) || !ExistsFloat32(v2, actual) || !ExistsFloat32(v1, actual) {
		t.Errorf("UnionFloat32 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetFloat32(nil)
	if len(actual) != 0 {
		t.Errorf("SetFloat32 failed")
	}

	actualPtr2 := SetFloat32([]float32{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetFloat32 failed")
	}

	actualPtr3 := SetFloat32Ptr([]*float32{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetFloat32Ptr failed")
	}

	listPtr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetFloat32Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsFloat32Ptr(&v8, actualPtr) || !ExistsFloat32Ptr(&v2, actualPtr) || !ExistsFloat32Ptr(&v1, actualPtr) {
		t.Errorf("SetFloat32Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}

func TestSetFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	expected := []float64{v8, v2, v1}
	list := []float64{v8, v2, v1, v1, v2, v8}
	actual := SetFloat64(list)
	if len(actual) != 3 || !ExistsFloat64(v8, actual) || !ExistsFloat64(v2, actual) || !ExistsFloat64(v1, actual) {
		t.Errorf("UnionFloat64 failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = SetFloat64(nil)
	if len(actual) != 0 {
		t.Errorf("SetFloat64 failed")
	}

	actualPtr2 := SetFloat64([]float64{})
	if len(actualPtr2) != 0 {
		t.Errorf("SetFloat64 failed")
	}

	actualPtr3 := SetFloat64Ptr([]*float64{})
	if len(actualPtr3) != 0 {
		t.Errorf("SetFloat64Ptr failed")
	}

	listPtr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := SetFloat64Ptr(listPtr)
	if len(actualPtr) != 3 || !ExistsFloat64Ptr(&v8, actualPtr) || !ExistsFloat64Ptr(&v2, actualPtr) || !ExistsFloat64Ptr(&v1, actualPtr) {
		t.Errorf("SetFloat64Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
