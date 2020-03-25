package fp

import (
	"reflect"
	"testing"
)

func TestMinIntPtr(t *testing.T) {
	// Test : get min number in the list
	var v2 int = 2
	var v4 int = 4
	var v8 int = 8
	var v10 int = 10

	list := []*int{&v8, &v2, &v10, &v4}
	min := MinIntPtr(list)
	if *min != 2 {
		t.Errorf("MinIntPtr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinIntPtr([]*int{})
	if *min != 0 {
		t.Errorf("MinIntPtr failed. Expected=0, actual=%v", *min)
	}

	min = MinIntPtr(nil)
	if *min != 0 {
		t.Errorf("MinIntPtr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinInt64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int64 = 2
	var v4 int64 = 4
	var v8 int64 = 8
	var v10 int64 = 10

	list := []*int64{&v8, &v2, &v10, &v4}
	min := MinInt64Ptr(list)
	if *min != 2 {
		t.Errorf("MinInt64Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinInt64Ptr([]*int64{})
	if *min != 0 {
		t.Errorf("MinInt64Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinInt64Ptr(nil)
	if *min != 0 {
		t.Errorf("MinInt64Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinInt32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int32 = 2
	var v4 int32 = 4
	var v8 int32 = 8
	var v10 int32 = 10

	list := []*int32{&v8, &v2, &v10, &v4}
	min := MinInt32Ptr(list)
	if *min != 2 {
		t.Errorf("MinInt32Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinInt32Ptr([]*int32{})
	if *min != 0 {
		t.Errorf("MinInt32Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinInt32Ptr(nil)
	if *min != 0 {
		t.Errorf("MinInt32Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinInt16Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int16 = 2
	var v4 int16 = 4
	var v8 int16 = 8
	var v10 int16 = 10

	list := []*int16{&v8, &v2, &v10, &v4}
	min := MinInt16Ptr(list)
	if *min != 2 {
		t.Errorf("MinInt16Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinInt16Ptr([]*int16{})
	if *min != 0 {
		t.Errorf("MinInt16Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinInt16Ptr(nil)
	if *min != 0 {
		t.Errorf("MinInt16Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinInt8Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 int8 = 2
	var v4 int8 = 4
	var v8 int8 = 8
	var v10 int8 = 10

	list := []*int8{&v8, &v2, &v10, &v4}
	min := MinInt8Ptr(list)
	if *min != 2 {
		t.Errorf("MinInt8Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinInt8Ptr([]*int8{})
	if *min != 0 {
		t.Errorf("MinInt8Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinInt8Ptr(nil)
	if *min != 0 {
		t.Errorf("MinInt8Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinUintPtr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint = 2
	var v4 uint = 4
	var v8 uint = 8
	var v10 uint = 10

	list := []*uint{&v8, &v2, &v10, &v4}
	min := MinUintPtr(list)
	if *min != 2 {
		t.Errorf("MinUintPtr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinUintPtr([]*uint{})
	if *min != 0 {
		t.Errorf("MinUintPtr failed. Expected=0, actual=%v", *min)
	}

	min = MinUintPtr(nil)
	if *min != 0 {
		t.Errorf("MinUintPtr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinUint64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v10 uint64 = 10

	list := []*uint64{&v8, &v2, &v10, &v4}
	min := MinUint64Ptr(list)
	if *min != 2 {
		t.Errorf("MinUint64Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinUint64Ptr([]*uint64{})
	if *min != 0 {
		t.Errorf("MinUint64Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinUint64Ptr(nil)
	if *min != 0 {
		t.Errorf("MinUint64Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinUint32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v10 uint32 = 10

	list := []*uint32{&v8, &v2, &v10, &v4}
	min := MinUint32Ptr(list)
	if *min != 2 {
		t.Errorf("MinUint32Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinUint32Ptr([]*uint32{})
	if *min != 0 {
		t.Errorf("MinUint32Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinUint32Ptr(nil)
	if *min != 0 {
		t.Errorf("MinUint32Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinUint16Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v10 uint16 = 10

	list := []*uint16{&v8, &v2, &v10, &v4}
	min := MinUint16Ptr(list)
	if *min != 2 {
		t.Errorf("MinUint16Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinUint16Ptr([]*uint16{})
	if *min != 0 {
		t.Errorf("MinUint16Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinUint16Ptr(nil)
	if *min != 0 {
		t.Errorf("MinUint16Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinUint8Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v10 uint8 = 10

	list := []*uint8{&v8, &v2, &v10, &v4}
	min := MinUint8Ptr(list)
	if *min != 2 {
		t.Errorf("MinUint8Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinUint8Ptr([]*uint8{})
	if *min != 0 {
		t.Errorf("MinUint8Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinUint8Ptr(nil)
	if *min != 0 {
		t.Errorf("MinUint8Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinFloat32Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 float32 = 2
	var v4 float32 = 4
	var v8 float32 = 8
	var v10 float32 = 10

	list := []*float32{&v8, &v2, &v10, &v4}
	min := MinFloat32Ptr(list)
	if *min != 2 {
		t.Errorf("MinFloat32Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinFloat32Ptr([]*float32{})
	if *min != 0 {
		t.Errorf("MinFloat32Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinFloat32Ptr(nil)
	if *min != 0 {
		t.Errorf("MinFloat32Ptr failed. Expected=0, actual=%v", *min)
	}
}

func TestMinFloat64Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 float64 = 2
	var v4 float64 = 4
	var v8 float64 = 8
	var v10 float64 = 10

	list := []*float64{&v8, &v2, &v10, &v4}
	min := MinFloat64Ptr(list)
	if *min != 2 {
		t.Errorf("MinFloat64Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = MinFloat64Ptr([]*float64{})
	if *min != 0 {
		t.Errorf("MinFloat64Ptr failed. Expected=0, actual=%v", *min)
	}

	min = MinFloat64Ptr(nil)
	if *min != 0 {
		t.Errorf("MinFloat64Ptr failed. Expected=0, actual=%v", *min)
	}
}
