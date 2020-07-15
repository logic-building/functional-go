package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSomeIntPtr(t *testing.T) {
	// Test : value exist in the list

	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	var v7 int = 7
	var v8 int = 8
	var v10 int = 10

	list1 := []*int{&v8, &v2, &v10, &v4}
	if !SomeIntPtr(isEvenIntPtr, list1) {
		t.Errorf("SomeIntPtr failed. Expected=true, actual=false")
	}

	list2 := []*int{&v1, &v3, &v5, &v7}
	if SomeIntPtr(isEvenIntPtr, list2) {
		t.Errorf("SomeIntPtr failed. Expected=false, actual=true")
	}

	if SomeIntPtr(nil, nil) {
		t.Errorf("SomeIntPtr failed. Expected=false, actual=true")
	}

	if SomeIntPtr(isEvenIntPtr, []*int{}) {
		t.Errorf("SomeIntPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeInt64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	var v7 int64 = 7
	var v8 int64 = 8
	var v10 int64 = 10

	list1 := []*int64{&v8, &v2, &v10, &v4}
	if !SomeInt64Ptr(isEvenInt64Ptr, list1) {
		t.Errorf("SomeInt64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int64{&v1, &v3, &v5, &v7}
	if SomeInt64Ptr(isEvenInt64Ptr, list2) {
		t.Errorf("SomeInt64Ptr failed. Expected=false, actual=true")
	}

	if SomeInt64Ptr(nil, nil) {
		t.Errorf("SomeInt64Ptr failed. Expected=false, actual=true")
	}

	if SomeInt64Ptr(isEvenInt64Ptr, []*int64{}) {
		t.Errorf("SomeInt64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeInt32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	var v7 int32 = 7
	var v8 int32 = 8
	var v10 int32 = 10

	list1 := []*int32{&v8, &v2, &v10, &v4}
	if !SomeInt32Ptr(isEvenInt32Ptr, list1) {
		t.Errorf("SomeInt32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int32{&v1, &v3, &v5, &v7}
	if SomeInt32Ptr(isEvenInt32Ptr, list2) {
		t.Errorf("SomeInt32Ptr failed. Expected=false, actual=true")
	}

	if SomeInt32Ptr(nil, nil) {
		t.Errorf("SomeInt32Ptr failed. Expected=false, actual=true")
	}

	if SomeInt32Ptr(isEvenInt32Ptr, []*int32{}) {
		t.Errorf("SomeInt32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeInt16Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	var v7 int16 = 7
	var v8 int16 = 8
	var v10 int16 = 10

	list1 := []*int16{&v8, &v2, &v10, &v4}
	if !SomeInt16Ptr(isEvenInt16Ptr, list1) {
		t.Errorf("SomeInt16Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int16{&v1, &v3, &v5, &v7}
	if SomeInt16Ptr(isEvenInt16Ptr, list2) {
		t.Errorf("SomeInt16Ptr failed. Expected=false, actual=true")
	}

	if SomeInt16Ptr(nil, nil) {
		t.Errorf("SomeInt16Ptr failed. Expected=false, actual=true")
	}

	if SomeInt16Ptr(isEvenInt16Ptr, []*int16{}) {
		t.Errorf("SomeInt16Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeInt8Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	var v7 int8 = 7
	var v8 int8 = 8
	var v10 int8 = 10

	list1 := []*int8{&v8, &v2, &v10, &v4}
	if !SomeInt8Ptr(isEvenInt8Ptr, list1) {
		t.Errorf("SomeInt8Ptr failed. Expected=true, actual=false")
	}

	list2 := []*int8{&v1, &v3, &v5, &v7}
	if SomeInt8Ptr(isEvenInt8Ptr, list2) {
		t.Errorf("SomeInt8Ptr failed. Expected=false, actual=true")
	}

	if SomeInt8Ptr(nil, nil) {
		t.Errorf("SomeInt8Ptr failed. Expected=false, actual=true")
	}

	if SomeInt8Ptr(isEvenInt8Ptr, []*int8{}) {
		t.Errorf("SomeInt8Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeUintPtr(t *testing.T) {
	// Test : value exist in the list

	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	var v7 uint = 7
	var v8 uint = 8
	var v10 uint = 10

	list1 := []*uint{&v8, &v2, &v10, &v4}
	if !SomeUintPtr(isEvenUintPtr, list1) {
		t.Errorf("SomeUintPtr failed. Expected=true, actual=false")
	}

	list2 := []*uint{&v1, &v3, &v5, &v7}
	if SomeUintPtr(isEvenUintPtr, list2) {
		t.Errorf("SomeUintPtr failed. Expected=false, actual=true")
	}

	if SomeUintPtr(nil, nil) {
		t.Errorf("SomeUintPtr failed. Expected=false, actual=true")
	}

	if SomeUintPtr(isEvenUintPtr, []*uint{}) {
		t.Errorf("SomeUintPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeUint64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v7 uint64 = 7
	var v8 uint64 = 8
	var v10 uint64 = 10

	list1 := []*uint64{&v8, &v2, &v10, &v4}
	if !SomeUint64Ptr(isEvenUint64Ptr, list1) {
		t.Errorf("SomeUint64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint64{&v1, &v3, &v5, &v7}
	if SomeUint64Ptr(isEvenUint64Ptr, list2) {
		t.Errorf("SomeUint64Ptr failed. Expected=false, actual=true")
	}

	if SomeUint64Ptr(nil, nil) {
		t.Errorf("SomeUint64Ptr failed. Expected=false, actual=true")
	}

	if SomeUint64Ptr(isEvenUint64Ptr, []*uint64{}) {
		t.Errorf("SomeUint64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeUint32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v7 uint32 = 7
	var v8 uint32 = 8
	var v10 uint32 = 10

	list1 := []*uint32{&v8, &v2, &v10, &v4}
	if !SomeUint32Ptr(isEvenUint32Ptr, list1) {
		t.Errorf("SomeUint32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint32{&v1, &v3, &v5, &v7}
	if SomeUint32Ptr(isEvenUint32Ptr, list2) {
		t.Errorf("SomeUint32Ptr failed. Expected=false, actual=true")
	}

	if SomeUint32Ptr(nil, nil) {
		t.Errorf("SomeUint32Ptr failed. Expected=false, actual=true")
	}

	if SomeUint32Ptr(isEvenUint32Ptr, []*uint32{}) {
		t.Errorf("SomeUint32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeUint16Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v7 uint16 = 7
	var v8 uint16 = 8
	var v10 uint16 = 10

	list1 := []*uint16{&v8, &v2, &v10, &v4}
	if !SomeUint16Ptr(isEvenUint16Ptr, list1) {
		t.Errorf("SomeUint16Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint16{&v1, &v3, &v5, &v7}
	if SomeUint16Ptr(isEvenUint16Ptr, list2) {
		t.Errorf("SomeUint16Ptr failed. Expected=false, actual=true")
	}

	if SomeUint16Ptr(nil, nil) {
		t.Errorf("SomeUint16Ptr failed. Expected=false, actual=true")
	}

	if SomeUint16Ptr(isEvenUint16Ptr, []*uint16{}) {
		t.Errorf("SomeUint16Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeUint8Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v7 uint8 = 7
	var v8 uint8 = 8
	var v10 uint8 = 10

	list1 := []*uint8{&v8, &v2, &v10, &v4}
	if !SomeUint8Ptr(isEvenUint8Ptr, list1) {
		t.Errorf("SomeUint8Ptr failed. Expected=true, actual=false")
	}

	list2 := []*uint8{&v1, &v3, &v5, &v7}
	if SomeUint8Ptr(isEvenUint8Ptr, list2) {
		t.Errorf("SomeUint8Ptr failed. Expected=false, actual=true")
	}

	if SomeUint8Ptr(nil, nil) {
		t.Errorf("SomeUint8Ptr failed. Expected=false, actual=true")
	}

	if SomeUint8Ptr(isEvenUint8Ptr, []*uint8{}) {
		t.Errorf("SomeUint8Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeStrPtr(t *testing.T) {
	// Test : value exist in the list

	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v8 string = "8"
	var v10 string = "10"

	list1 := []*string{&v8, &v2, &v10, &v4}
	if !SomeStrPtr(isEvenStrPtr, list1) {
		t.Errorf("SomeStrPtr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v1, &v3, &v5, &v7}
	if SomeStrPtr(isEvenStrPtr, list2) {
		t.Errorf("SomeStrPtr failed. Expected=false, actual=true")
	}

	if SomeStrPtr(nil, nil) {
		t.Errorf("SomeStrPtr failed. Expected=false, actual=true")
	}

	if SomeStrPtr(isEvenStrPtr, []*string{}) {
		t.Errorf("SomeStrPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeBoolPtr(t *testing.T) {
	// Test : value exist in the list

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vf}
	if !SomeBoolPtr(func(v *bool) bool { return *v == true }, list1) {
		t.Errorf("SomeBoolPtr failed. Expected=true, actual=false")
	}

	if SomeBoolPtr(nil, nil) {
		t.Errorf("SomeBoolPtr failed. Expected=false, actual=true")
	}

	if SomeBoolPtr(func(v *bool) bool { return *v == true }, []*bool{}) {
		t.Errorf("SomeBoolPtr failed. Expected=false, actual=true")
	}
}

func TestSomeFloat32Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	var v7 float32 = 7
	var v8 float32 = 8
	var v10 float32 = 10

	list1 := []*float32{&v8, &v2, &v10, &v4}
	if !SomeFloat32Ptr(isEvenFloat32Ptr, list1) {
		t.Errorf("SomeFloat32Ptr failed. Expected=true, actual=false")
	}

	list2 := []*float32{&v1, &v3, &v5, &v7}
	if SomeFloat32Ptr(isEvenFloat32Ptr, list2) {
		t.Errorf("SomeFloat32Ptr failed. Expected=false, actual=true")
	}

	if SomeFloat32Ptr(nil, nil) {
		t.Errorf("SomeFloat32Ptr failed. Expected=false, actual=true")
	}

	if SomeFloat32Ptr(isEvenFloat32Ptr, []*float32{}) {
		t.Errorf("SomeFloat32Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestSomeFloat64Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	var v7 float64 = 7
	var v8 float64 = 8
	var v10 float64 = 10

	list1 := []*float64{&v8, &v2, &v10, &v4}
	if !SomeFloat64Ptr(isEvenFloat64Ptr, list1) {
		t.Errorf("SomeFloat64Ptr failed. Expected=true, actual=false")
	}

	list2 := []*float64{&v1, &v3, &v5, &v7}
	if SomeFloat64Ptr(isEvenFloat64Ptr, list2) {
		t.Errorf("SomeFloat64Ptr failed. Expected=false, actual=true")
	}

	if SomeFloat64Ptr(nil, nil) {
		t.Errorf("SomeFloat64Ptr failed. Expected=false, actual=true")
	}

	if SomeFloat64Ptr(isEvenFloat64Ptr, []*float64{}) {
		t.Errorf("SomeFloat64Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
