package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestExistsIntPtr(t *testing.T) {
	// Test : value exist in the list

	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v8 int = 8
	var v10 int = 10
	var v80 int = 80

	list1 := []*int{&v8, &v2, &v10, &v4}
	if !ExistsIntPtr(&v8, list1) {
		t.Errorf("ExistsInt failed. Expected=true, actual=false")
	}

	list2 := []*int{&v8, &v2, &v10, &v5, &v4}
	if ExistsIntPtr(&v80, list2) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
	}

	if ExistsIntPtr(&v80, nil) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
	}

	if ExistsIntPtr(&v80, []*int{}) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsInt64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v8 int64 = 8
	var v10 int64 = 10
	var v80 int64 = 80

	list1 := []*int64{&v8, &v2, &v10, &v4}
	if !ExistsInt64Ptr(&v8, list1) {
		t.Errorf("ExistsInt64 failed. Expected=true, actual=false")
	}

	list2 := []*int64{&v8, &v2, &v10, &v5, &v4}
	if ExistsInt64Ptr(&v80, list2) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
	}

	if ExistsInt64Ptr(&v80, nil) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
	}

	if ExistsInt64Ptr(&v80, []*int64{}) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsInt32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v8 int32 = 8
	var v10 int32 = 10
	var v80 int32 = 80

	list1 := []*int32{&v8, &v2, &v10, &v4}
	if !ExistsInt32Ptr(&v8, list1) {
		t.Errorf("ExistsInt32 failed. Expected=true, actual=false")
	}

	list2 := []*int32{&v8, &v2, &v10, &v5, &v4}
	if ExistsInt32Ptr(&v80, list2) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}

	if ExistsInt32Ptr(&v80, nil) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}

	if ExistsInt32Ptr(&v80, []*int32{}) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsInt16Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v8 int16 = 8
	var v10 int16 = 10
	var v80 int16 = 80

	list1 := []*int16{&v8, &v2, &v10, &v4}
	if !ExistsInt16Ptr(&v8, list1) {
		t.Errorf("ExistsInt16 failed. Expected=true, actual=false")
	}

	list2 := []*int16{&v8, &v2, &v10, &v5, &v4}
	if ExistsInt16Ptr(&v80, list2) {
		t.Errorf("ExistsInt16 failed. Expected=false, actual=true")
	}

	if ExistsInt16Ptr(&v80, nil) {
		t.Errorf("ExistsInt16 failed. Expected=false, actual=true")
	}

	if ExistsInt16Ptr(&v80, []*int16{}) {
		t.Errorf("ExistsInt16 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsInt8Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v8 int8 = 8
	var v10 int8 = 10
	var v80 int8 = 80

	list1 := []*int8{&v8, &v2, &v10, &v4}
	if !ExistsInt8Ptr(&v8, list1) {
		t.Errorf("ExistsInt8 failed. Expected=true, actual=false")
	}

	list2 := []*int8{&v8, &v2, &v10, &v5, &v4}
	if ExistsInt8Ptr(&v80, list2) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
	}

	if ExistsInt8Ptr(&v80, nil) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
	}

	if ExistsInt8Ptr(&v80, []*int8{}) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsUintPtr(t *testing.T) {
	// Test : value exist in the list

	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v8 uint = 8
	var v10 uint = 10
	var v80 uint = 80

	list1 := []*uint{&v8, &v2, &v10, &v4}
	if !ExistsUintPtr(&v8, list1) {
		t.Errorf("ExistsUint failed. Expected=true, actual=false")
	}

	list2 := []*uint{&v8, &v2, &v10, &v5, &v4}
	if ExistsUintPtr(&v80, list2) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
	}

	if ExistsUintPtr(&v80, nil) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
	}

	if ExistsUintPtr(&v80, []*uint{}) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsUint64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v8 uint64 = 8
	var v10 uint64 = 10
	var v80 uint64 = 80

	list1 := []*uint64{&v8, &v2, &v10, &v4}
	if !ExistsUint64Ptr(&v8, list1) {
		t.Errorf("ExistsUint64 failed. Expected=true, actual=false")
	}

	list2 := []*uint64{&v8, &v2, &v10, &v5, &v4}
	if ExistsUint64Ptr(&v80, list2) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
	}

	if ExistsUint64Ptr(&v80, nil) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
	}

	if ExistsUint64Ptr(&v80, []*uint64{}) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsUint32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v8 uint32 = 8
	var v10 uint32 = 10
	var v80 uint32 = 80

	list1 := []*uint32{&v8, &v2, &v10, &v4}
	if !ExistsUint32Ptr(&v8, list1) {
		t.Errorf("ExistsUint32 failed. Expected=true, actual=false")
	}

	list2 := []*uint32{&v8, &v2, &v10, &v5, &v4}
	if ExistsUint32Ptr(&v80, list2) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
	}

	if ExistsUint32Ptr(&v80, nil) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
	}

	if ExistsUint32Ptr(&v80, []*uint32{}) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsUint16Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v8 uint16 = 8
	var v10 uint16 = 10
	var v80 uint16 = 80

	list1 := []*uint16{&v8, &v2, &v10, &v4}
	if !ExistsUint16Ptr(&v8, list1) {
		t.Errorf("ExistsUint16 failed. Expected=true, actual=false")
	}

	list2 := []*uint16{&v8, &v2, &v10, &v5, &v4}
	if ExistsUint16Ptr(&v80, list2) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
	}

	if ExistsUint16Ptr(&v80, nil) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
	}

	if ExistsUint16Ptr(&v80, []*uint16{}) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsUint8Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v8 uint8 = 8
	var v10 uint8 = 10
	var v80 uint8 = 80

	list1 := []*uint8{&v8, &v2, &v10, &v4}
	if !ExistsUint8Ptr(&v8, list1) {
		t.Errorf("ExistsUint8 failed. Expected=true, actual=false")
	}

	list2 := []*uint8{&v8, &v2, &v10, &v5, &v4}
	if ExistsUint8Ptr(&v80, list2) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
	}

	if ExistsUint8Ptr(&v80, nil) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
	}

	if ExistsUint8Ptr(&v80, []*uint8{}) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsStrPtr(t *testing.T) {
	// Test : value exist in the list

	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v80 string = "80"

	list1 := []*string{&v8, &v2, &v10, &v4}
	if !ExistsStrPtr(&v8, list1) {
		t.Errorf("ExistsStr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v8, &v2, &v10, &v5, &v4}
	if ExistsStrPtr(&v80, list2) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
	}

	if ExistsStrPtr(&v80, nil) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
	}

	if ExistsStrPtr(&v80, []*string{}) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsFloat32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v8 float32 = 8
	var v10 float32 = 10
	var v80 float32 = 80

	list1 := []*float32{&v8, &v2, &v10, &v4}
	if !ExistsFloat32Ptr(&v8, list1) {
		t.Errorf("ExistsFloat32 failed. Expected=true, actual=false")
	}

	list2 := []*float32{&v8, &v2, &v10, &v5, &v4}
	if ExistsFloat32Ptr(&v80, list2) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
	}

	if ExistsFloat32Ptr(&v80, nil) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
	}

	if ExistsFloat32Ptr(&v80, []*float32{}) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestExistsFloat64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v8 float64 = 8
	var v10 float64 = 10
	var v80 float64 = 80

	list1 := []*float64{&v8, &v2, &v10, &v4}
	if !ExistsFloat64Ptr(&v8, list1) {
		t.Errorf("ExistsFloat64 failed. Expected=true, actual=false")
	}

	list2 := []*float64{&v8, &v2, &v10, &v5, &v4}
	if ExistsFloat64Ptr(&v80, list2) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
	}

	if ExistsFloat64Ptr(&v80, nil) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
	}

	if ExistsFloat64Ptr(&v80, []*float64{}) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
