package fp

import (
	"reflect"
	"testing"
)

func TestMinMaxIntPtr(t *testing.T) {
	// Test : get min number in the list
	var v2 int = 2
	var v4 int = 4
	var v8 int = 8
	var v10 int = 10

	list := []*int{&v8, &v2, &v10, &v4}
	min, max := MinMaxIntPtr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxIntPtr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxIntPtr([]*int{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxIntPtr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxIntPtr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxIntPtr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxInt64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int64 = 2
	var v4 int64 = 4
	var v8 int64 = 8
	var v10 int64 = 10

	list := []*int64{&v8, &v2, &v10, &v4}
	min, max := MinMaxInt64Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxInt64Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxInt64Ptr([]*int64{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxInt64Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxInt32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int32 = 2
	var v4 int32 = 4
	var v8 int32 = 8
	var v10 int32 = 10

	list := []*int32{&v8, &v2, &v10, &v4}
	min, max := MinMaxInt32Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxInt32Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxInt32Ptr([]*int32{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxInt32Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxInt16Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int16 = 2
	var v4 int16 = 4
	var v8 int16 = 8
	var v10 int16 = 10

	list := []*int16{&v8, &v2, &v10, &v4}
	min, max := MinMaxInt16Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxInt16Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxInt16Ptr([]*int16{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt16Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxInt16Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt16Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxInt8Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int8 = 2
	var v4 int8 = 4
	var v8 int8 = 8
	var v10 int8 = 10

	list := []*int8{&v8, &v2, &v10, &v4}
	min, max := MinMaxInt8Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxInt8Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxInt8Ptr([]*int8{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt8Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxInt8Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxInt8Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxUintPtr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint = 2
	var v4 uint = 4
	var v8 uint = 8
	var v10 uint = 10

	list := []*uint{&v8, &v2, &v10, &v4}
	min, max := MinMaxUintPtr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxUintPtr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxUintPtr([]*uint{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUintPtr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxUintPtr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUintPtr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxUint64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v10 uint64 = 10

	list := []*uint64{&v8, &v2, &v10, &v4}
	min, max := MinMaxUint64Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxUint64Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxUint64Ptr([]*uint64{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxUint64Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxUint32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v10 uint32 = 10

	list := []*uint32{&v8, &v2, &v10, &v4}
	min, max := MinMaxUint32Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxUint32Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxUint32Ptr([]*uint32{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxUint32Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxUint16Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v10 uint16 = 10

	list := []*uint16{&v8, &v2, &v10, &v4}
	min, max := MinMaxUint16Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxUint16Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxUint16Ptr([]*uint16{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint16Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxUint16Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint16Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxUint8Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v10 uint8 = 10

	list := []*uint8{&v8, &v2, &v10, &v4}
	min, max := MinMaxUint8Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxUint8Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxUint8Ptr([]*uint8{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint8Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxUint8Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxUint8Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxFloat32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 float32 = 2
	var v4 float32 = 4
	var v8 float32 = 8
	var v10 float32 = 10

	list := []*float32{&v8, &v2, &v10, &v4}
	min, max := MinMaxFloat32Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxFloat32Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxFloat32Ptr([]*float32{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxFloat32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxFloat32Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxFloat32Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}

func TestMinMaxFloat64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 float64 = 2
	var v4 float64 = 4
	var v8 float64 = 8
	var v10 float64 = 10

	list := []*float64{&v8, &v2, &v10, &v4}
	min, max := MinMaxFloat64Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMaxFloat64Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMaxFloat64Ptr([]*float64{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxFloat64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMaxFloat64Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMaxFloat64Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}
