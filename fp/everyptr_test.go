package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEveryIntPtr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v8 int = 8
	var v10 int = 10
	list1 := []*int{&v8, &v2, &v10, &v4}
	if !EveryIntPtr(isEvenIntPtr, list1) {
		t.Errorf("EveryIntPtr failed. Expected=true, actual=false")
	}

	list2 := []*int{&v8, &v2, &v10, &v5, &v4}
	if EveryIntPtr(isEvenIntPtr, list2) {
		t.Errorf("EveryIntPtr failed. Expected=false, actual=true")
	}

	if EveryIntPtr(isEvenIntPtr, nil) {
		t.Errorf("EveryIntPtr failed. Expected=false, actual=true")
	}

	if EveryIntPtr(isEvenIntPtr, []*int{}) {
		t.Errorf("EveryIntPtr failed. Expected=false, actual=true")
	}

	if EveryIntPtr(nil, []*int{}) {
		t.Errorf("EveryIntPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryInt64Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v8 int64 = 8
	var v10 int64 = 10
	list1 := []*int64{&v8, &v2, &v10, &v4}
	if !EveryInt64Ptr(isEvenInt64Ptr, list1) {
		t.Errorf("EveryInt64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int64{&v8, &v2, &v10, &v5, &v4}
	if EveryInt64Ptr(isEvenInt64Ptr, list2) {
		t.Errorf("EveryInt64Ptr failed. Expected=false, actual=true")
	}

	if EveryInt64Ptr(isEvenInt64Ptr, nil) {
		t.Errorf("EveryInt64Ptr failed. Expected=false, actual=true")
	}

	if EveryInt64Ptr(isEvenInt64Ptr, []*int64{}) {
		t.Errorf("EveryInt64Ptr failed. Expected=false, actual=true")
	}

	if EveryInt64Ptr(nil, []*int64{}) {
		t.Errorf("EveryInt64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryInt32Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v8 int32 = 8
	var v10 int32 = 10
	list1 := []*int32{&v8, &v2, &v10, &v4}
	if !EveryInt32Ptr(isEvenInt32Ptr, list1) {
		t.Errorf("EveryInt32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int32{&v8, &v2, &v10, &v5, &v4}
	if EveryInt32Ptr(isEvenInt32Ptr, list2) {
		t.Errorf("EveryInt32Ptr failed. Expected=false, actual=true")
	}

	if EveryInt32Ptr(isEvenInt32Ptr, nil) {
		t.Errorf("EveryInt32Ptr failed. Expected=false, actual=true")
	}

	if EveryInt32Ptr(isEvenInt32Ptr, []*int32{}) {
		t.Errorf("EveryInt32Ptr failed. Expected=false, actual=true")
	}

	if EveryInt32Ptr(nil, []*int32{}) {
		t.Errorf("EveryInt32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryInt16Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v8 int16 = 8
	var v10 int16 = 10
	list1 := []*int16{&v8, &v2, &v10, &v4}
	if !EveryInt16Ptr(isEvenInt16Ptr, list1) {
		t.Errorf("EveryInt16Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int16{&v8, &v2, &v10, &v5, &v4}
	if EveryInt16Ptr(isEvenInt16Ptr, list2) {
		t.Errorf("EveryInt16Ptr failed. Expected=false, actual=true")
	}

	if EveryInt16Ptr(isEvenInt16Ptr, nil) {
		t.Errorf("EveryInt16Ptr failed. Expected=false, actual=true")
	}

	if EveryInt16Ptr(isEvenInt16Ptr, []*int16{}) {
		t.Errorf("EveryInt16Ptr failed. Expected=false, actual=true")
	}

	if EveryInt16Ptr(nil, []*int16{}) {
		t.Errorf("EveryInt16Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryInt8Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v8 int8 = 8
	var v10 int8 = 10
	list1 := []*int8{&v8, &v2, &v10, &v4}
	if !EveryInt8Ptr(isEvenInt8Ptr, list1) {
		t.Errorf("EveryInt8Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int8{&v8, &v2, &v10, &v5, &v4}
	if EveryInt8Ptr(isEvenInt8Ptr, list2) {
		t.Errorf("EveryInt8Ptr failed. Expected=false, actual=true")
	}

	if EveryInt8Ptr(isEvenInt8Ptr, nil) {
		t.Errorf("EveryInt8Ptr failed. Expected=false, actual=true")
	}

	if EveryInt8Ptr(isEvenInt8Ptr, []*int8{}) {
		t.Errorf("EveryInt8Ptr failed. Expected=false, actual=true")
	}

	if EveryInt8Ptr(nil, []*int8{}) {
		t.Errorf("EveryInt8Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryUintPtr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v8 uint = 8
	var v10 uint = 10
	list1 := []*uint{&v8, &v2, &v10, &v4}
	if !EveryUintPtr(isEvenUintPtr, list1) {
		t.Errorf("EveryUintPtr failed. Expected=true, actual=false")
	}

	list2 := []*uint{&v8, &v2, &v10, &v5, &v4}
	if EveryUintPtr(isEvenUintPtr, list2) {
		t.Errorf("EveryUintPtr failed. Expected=false, actual=true")
	}

	if EveryUintPtr(isEvenUintPtr, nil) {
		t.Errorf("EveryUintPtr failed. Expected=false, actual=true")
	}

	if EveryUintPtr(isEvenUintPtr, []*uint{}) {
		t.Errorf("EveryUintPtr failed. Expected=false, actual=true")
	}

	if EveryUintPtr(nil, []*uint{}) {
		t.Errorf("EveryUintPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryUint64Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v8 uint64 = 8
	var v10 uint64 = 10
	list1 := []*uint64{&v8, &v2, &v10, &v4}
	if !EveryUint64Ptr(isEvenUint64Ptr, list1) {
		t.Errorf("EveryUint64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint64{&v8, &v2, &v10, &v5, &v4}
	if EveryUint64Ptr(isEvenUint64Ptr, list2) {
		t.Errorf("EveryUint64Ptr failed. Expected=false, actual=true")
	}

	if EveryUint64Ptr(isEvenUint64Ptr, nil) {
		t.Errorf("EveryUint64Ptr failed. Expected=false, actual=true")
	}

	if EveryUint64Ptr(isEvenUint64Ptr, []*uint64{}) {
		t.Errorf("EveryUint64Ptr failed. Expected=false, actual=true")
	}

	if EveryUint64Ptr(nil, []*uint64{}) {
		t.Errorf("EveryUint64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryUint32Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v8 uint32 = 8
	var v10 uint32 = 10
	list1 := []*uint32{&v8, &v2, &v10, &v4}
	if !EveryUint32Ptr(isEvenUint32Ptr, list1) {
		t.Errorf("EveryUint32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint32{&v8, &v2, &v10, &v5, &v4}
	if EveryUint32Ptr(isEvenUint32Ptr, list2) {
		t.Errorf("EveryUint32Ptr failed. Expected=false, actual=true")
	}

	if EveryUint32Ptr(isEvenUint32Ptr, nil) {
		t.Errorf("EveryUint32Ptr failed. Expected=false, actual=true")
	}

	if EveryUint32Ptr(isEvenUint32Ptr, []*uint32{}) {
		t.Errorf("EveryUint32Ptr failed. Expected=false, actual=true")
	}

	if EveryUint32Ptr(nil, []*uint32{}) {
		t.Errorf("EveryUint32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryUint16Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v8 uint16 = 8
	var v10 uint16 = 10
	list1 := []*uint16{&v8, &v2, &v10, &v4}
	if !EveryUint16Ptr(isEvenUint16Ptr, list1) {
		t.Errorf("EveryUint16Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint16{&v8, &v2, &v10, &v5, &v4}
	if EveryUint16Ptr(isEvenUint16Ptr, list2) {
		t.Errorf("EveryUint16Ptr failed. Expected=false, actual=true")
	}

	if EveryUint16Ptr(isEvenUint16Ptr, nil) {
		t.Errorf("EveryUint16Ptr failed. Expected=false, actual=true")
	}

	if EveryUint16Ptr(isEvenUint16Ptr, []*uint16{}) {
		t.Errorf("EveryUint16Ptr failed. Expected=false, actual=true")
	}

	if EveryUint16Ptr(nil, []*uint16{}) {
		t.Errorf("EveryUint16Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryUint8Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v8 uint8 = 8
	var v10 uint8 = 10
	list1 := []*uint8{&v8, &v2, &v10, &v4}
	if !EveryUint8Ptr(isEvenUint8Ptr, list1) {
		t.Errorf("EveryUint8Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint8{&v8, &v2, &v10, &v5, &v4}
	if EveryUint8Ptr(isEvenUint8Ptr, list2) {
		t.Errorf("EveryUint8Ptr failed. Expected=false, actual=true")
	}

	if EveryUint8Ptr(isEvenUint8Ptr, nil) {
		t.Errorf("EveryUint8Ptr failed. Expected=false, actual=true")
	}

	if EveryUint8Ptr(isEvenUint8Ptr, []*uint8{}) {
		t.Errorf("EveryUint8Ptr failed. Expected=false, actual=true")
	}

	if EveryUint8Ptr(nil, []*uint8{}) {
		t.Errorf("EveryUint8Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryStrPtr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	list1 := []*string{&v2, &v4}
	if !EveryStrPtr(isEvenStrPtr, list1) {
		t.Errorf("EveryStrPtr failed. Expected=true, actual=false")
	}

	list1 = []*string{&v2, &v4, &v3}
	if EveryStrPtr(isEvenStrPtr, list1) {
		t.Errorf("EveryStrPtr failed. Expected=true, actual=false")
	}

	if EveryStrPtr(isEvenStrPtr, nil) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(isEvenStrPtr, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(nil, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	// Test : every value in the list is either true or false
	list1 := []*bool{&vt, &vt, &vt, &vt}
	if !EveryBoolPtr(TruePtr, list1) {
		t.Errorf("EveryBoolPtr failed. Expected=true, actual=false")
	}

	list1 = []*bool{&vt, &vt, &vt, &vf}
	if EveryBoolPtr(TruePtr, list1) {
		t.Errorf("EveryBoolPtr failed. Expected=true, actual=false")
	}

	list1 = []*bool{}
	if EveryBoolPtr(TruePtr, list1) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(TruePtr, nil) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(nil, []*bool{}) {
		t.Errorf("EveryBoolPtr failed. Expected=false, actual=true")
	}
	if EveryBoolPtr(TruePtr, []*bool{}) {
		t.Errorf("EveryBoolPtr failed. Expected=false, actual=true")
	}
}

func TruePtr(val *bool) bool {
	return *val == true
}

func TestEveryFloat32Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v8 float32 = 8
	var v10 float32 = 10
	list1 := []*float32{&v8, &v2, &v10, &v4}
	if !EveryFloat32Ptr(isEvenFloat32Ptr, list1) {
		t.Errorf("EveryFloat32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*float32{&v8, &v2, &v10, &v5, &v4}
	if EveryFloat32Ptr(isEvenFloat32Ptr, list2) {
		t.Errorf("EveryFloat32Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat32Ptr(isEvenFloat32Ptr, nil) {
		t.Errorf("EveryFloat32Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat32Ptr(isEvenFloat32Ptr, []*float32{}) {
		t.Errorf("EveryFloat32Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat32Ptr(nil, []*float32{}) {
		t.Errorf("EveryFloat32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestEveryFloat64Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v8 float64 = 8
	var v10 float64 = 10
	list1 := []*float64{&v8, &v2, &v10, &v4}
	if !EveryFloat64Ptr(isEvenFloat64Ptr, list1) {
		t.Errorf("EveryFloat64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*float64{&v8, &v2, &v10, &v5, &v4}
	if EveryFloat64Ptr(isEvenFloat64Ptr, list2) {
		t.Errorf("EveryFloat64Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat64Ptr(isEvenFloat64Ptr, nil) {
		t.Errorf("EveryFloat64Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat64Ptr(isEvenFloat64Ptr, []*float64{}) {
		t.Errorf("EveryFloat64Ptr failed. Expected=false, actual=true")
	}

	if EveryFloat64Ptr(nil, []*float64{}) {
		t.Errorf("EveryFloat64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
