package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDedupeInt(t *testing.T) {
	var v0 int
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	expectedList := []int{v0, v1, v2, v3, v4, v3}
	givenList := []int{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeInt(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeInt failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*int{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*int{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeIntPtr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeInt failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeIntPtr failed.")
	}
}

func TestDedupeInt64(t *testing.T) {
	var v0 int64
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	expectedList := []int64{v0, v1, v2, v3, v4, v3}
	givenList := []int64{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeInt64(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeInt64 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*int64{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*int64{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeInt64Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeInt64 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeInt64Ptr failed.")
	}
}

func TestDedupeInt32(t *testing.T) {
	var v0 int32
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	expectedList := []int32{v0, v1, v2, v3, v4, v3}
	givenList := []int32{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeInt32(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeInt32 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*int32{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*int32{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeInt32Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeInt32 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeInt32Ptr failed.")
	}
}

func TestDedupeInt16(t *testing.T) {
	var v0 int16
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	expectedList := []int16{v0, v1, v2, v3, v4, v3}
	givenList := []int16{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeInt16(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeInt16 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*int16{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*int16{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeInt16Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeInt16 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeInt16Ptr failed.")
	}
}

func TestDedupeInt8(t *testing.T) {
	var v0 int8
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	expectedList := []int8{v0, v1, v2, v3, v4, v3}
	givenList := []int8{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeInt8(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeInt8 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*int8{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*int8{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeInt8Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeInt8 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeInt8Ptr failed.")
	}
}

func TestDedupeUint(t *testing.T) {
	var v0 uint
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	expectedList := []uint{v0, v1, v2, v3, v4, v3}
	givenList := []uint{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeUint(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeUint failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*uint{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*uint{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeUintPtr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeUint failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeUintPtr failed.")
	}
}

func TestDedupeUint64(t *testing.T) {
	var v0 uint64
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	expectedList := []uint64{v0, v1, v2, v3, v4, v3}
	givenList := []uint64{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeUint64(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeUint64 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*uint64{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*uint64{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeUint64Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeUint64 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeUint64Ptr failed.")
	}
}

func TestDedupeUint32(t *testing.T) {
	var v0 uint32
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	expectedList := []uint32{v0, v1, v2, v3, v4, v3}
	givenList := []uint32{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeUint32(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeUint32 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*uint32{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*uint32{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeUint32Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeUint32 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeUint32Ptr failed.")
	}
}

func TestDedupeUint16(t *testing.T) {
	var v0 uint16
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	expectedList := []uint16{v0, v1, v2, v3, v4, v3}
	givenList := []uint16{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeUint16(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeUint16 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*uint16{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*uint16{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeUint16Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeUint16 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeUint16Ptr failed.")
	}
}

func TestDedupeUint8(t *testing.T) {
	var v0 uint8
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	expectedList := []uint8{v0, v1, v2, v3, v4, v3}
	givenList := []uint8{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeUint8(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeUint8 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*uint8{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*uint8{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeUint8Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeUint8 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeUint8Ptr failed.")
	}
}

func TestDedupeStr(t *testing.T) {
	var v0 string
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	expectedList := []string{v0, v1, v2, v3, v4, v3}
	givenList := []string{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeStr(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeStr failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*string{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*string{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeStrPtr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeStr failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeStrPtr failed.")
	}
}

func TestDedupeBool(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedList := []bool{vt, vf, vt}
	givenList := []bool{vt, vt, vf, vf, vf, vt, vt, vt, vt, vt, vt}
	r := DedupeBool(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeBool failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*bool{&vt, &vf, &vt, &vf}
	givenListPtr := []*bool{&vt, &vt, &vf, &vf, &vf, &vt, &vf, &vf, &vf, &vf, &vf}
	rPtr := DedupeBoolPtr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeBoolPtr failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeBoolPtr failed.")
	}
}

func TestDedupeFloat32(t *testing.T) {
	var v0 float32
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	expectedList := []float32{v0, v1, v2, v3, v4, v3}
	givenList := []float32{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeFloat32(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeFloat32 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*float32{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*float32{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeFloat32Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeFloat32 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeFloat32Ptr failed.")
	}
}

func TestDedupeFloat64(t *testing.T) {
	var v0 float64
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	expectedList := []float64{v0, v1, v2, v3, v4, v3}
	givenList := []float64{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := DedupeFloat64(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupeFloat64 failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*float64{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*float64{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := DedupeFloat64Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupeFloat64 failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupeFloat64Ptr failed.")
	}
}
