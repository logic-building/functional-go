package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMaxIntPtr(t *testing.T) {
	// Test : get max number in the list
	var v2 int = 2
	var v4 int = 4
	var v8 int = 8
	var v10 int = 10

	list := []*int{&v8, &v2, &v10, &v4}
	max := MaxIntPtr(list)
	if *max != 10 {
		t.Errorf("MaxIntPtr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxInt64Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 int64 = 2
	var v4 int64 = 4
	var v8 int64 = 8
	var v10 int64 = 10

	list := []*int64{&v8, &v2, &v10, &v4}
	max := MaxInt64Ptr(list)
	if *max != 10 {
		t.Errorf("MaxInt64Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxInt32Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 int32 = 2
	var v4 int32 = 4
	var v8 int32 = 8
	var v10 int32 = 10

	list := []*int32{&v8, &v2, &v10, &v4}
	max := MaxInt32Ptr(list)
	if *max != 10 {
		t.Errorf("MaxInt32Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxInt16Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 int16 = 2
	var v4 int16 = 4
	var v8 int16 = 8
	var v10 int16 = 10

	list := []*int16{&v8, &v2, &v10, &v4}
	max := MaxInt16Ptr(list)
	if *max != 10 {
		t.Errorf("MaxInt16Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxInt8Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 int8 = 2
	var v4 int8 = 4
	var v8 int8 = 8
	var v10 int8 = 10

	list := []*int8{&v8, &v2, &v10, &v4}
	max := MaxInt8Ptr(list)
	if *max != 10 {
		t.Errorf("MaxInt8Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxUintPtr(t *testing.T) {
	// Test : get max number in the list
	var v2 uint = 2
	var v4 uint = 4
	var v8 uint = 8
	var v10 uint = 10

	list := []*uint{&v8, &v2, &v10, &v4}
	max := MaxUintPtr(list)
	if *max != 10 {
		t.Errorf("MaxUintPtr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxUint64Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v10 uint64 = 10

	list := []*uint64{&v8, &v2, &v10, &v4}
	max := MaxUint64Ptr(list)
	if *max != 10 {
		t.Errorf("MaxUint64Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxUint32Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v10 uint32 = 10

	list := []*uint32{&v8, &v2, &v10, &v4}
	max := MaxUint32Ptr(list)
	if *max != 10 {
		t.Errorf("MaxUint32Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxUint16Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v10 uint16 = 10

	list := []*uint16{&v8, &v2, &v10, &v4}
	max := MaxUint16Ptr(list)
	if *max != 10 {
		t.Errorf("MaxUint16Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxUint8Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v10 uint8 = 10

	list := []*uint8{&v8, &v2, &v10, &v4}
	max := MaxUint8Ptr(list)
	if *max != 10 {
		t.Errorf("MaxUint8Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxFloat32Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 float32 = 2
	var v4 float32 = 4
	var v8 float32 = 8
	var v10 float32 = 10

	list := []*float32{&v8, &v2, &v10, &v4}
	max := MaxFloat32Ptr(list)
	if *max != 10 {
		t.Errorf("MaxFloat32Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}

func TestMaxFloat64Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 float64 = 2
	var v4 float64 = 4
	var v8 float64 = 8
	var v10 float64 = 10

	list := []*float64{&v8, &v2, &v10, &v4}
	max := MaxFloat64Ptr(list)
	if *max != 10 {
		t.Errorf("MaxFloat64Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}
