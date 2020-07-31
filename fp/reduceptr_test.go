package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestReduceIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	list := []*int{&v1, &v2, &v3, &v4, &v5}
	var expected int = 15
	actual := ReduceIntPtr(plusIntPtr, list)
	if *actual != expected {
		t.Errorf("ReduceIntPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceIntPtr(plusIntPtr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceIntPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int{&v1, &v2}
	expected = 3
	actual = ReduceIntPtr(plusIntPtr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int{&v1}
	expected = 1
	actual = ReduceIntPtr(plusIntPtr, list)
	if *actual != expected {
		t.Errorf("ReduceIntPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int{}
	expected = 0
	actual = ReduceIntPtr(plusIntPtr, list)
	if *actual != expected {
		t.Errorf("ReduceIntPtr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusIntPtr(num1, num2 *int) *int {
	c := *num1 + *num2
	return &c
}

func TestReduceInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	list := []*int64{&v1, &v2, &v3, &v4, &v5}
	var expected int64 = 15
	actual := ReduceInt64Ptr(plusInt64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int64{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceInt64Ptr(plusInt64Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceInt64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int64{&v1, &v2}
	expected = 3
	actual = ReduceInt64Ptr(plusInt64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int64{&v1}
	expected = 1
	actual = ReduceInt64Ptr(plusInt64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int64{}
	expected = 0
	actual = ReduceInt64Ptr(plusInt64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt64Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusInt64Ptr(num1, num2 *int64) *int64 {
	c := *num1 + *num2
	return &c
}

func TestReduceInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	list := []*int32{&v1, &v2, &v3, &v4, &v5}
	var expected int32 = 15
	actual := ReduceInt32Ptr(plusInt32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int32{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceInt32Ptr(plusInt32Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceInt32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int32{&v1, &v2}
	expected = 3
	actual = ReduceInt32Ptr(plusInt32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int32{&v1}
	expected = 1
	actual = ReduceInt32Ptr(plusInt32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int32{}
	expected = 0
	actual = ReduceInt32Ptr(plusInt32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt32Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusInt32Ptr(num1, num2 *int32) *int32 {
	c := *num1 + *num2
	return &c
}

func TestReduceInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	list := []*int16{&v1, &v2, &v3, &v4, &v5}
	var expected int16 = 15
	actual := ReduceInt16Ptr(plusInt16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int16{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceInt16Ptr(plusInt16Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceInt16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int16{&v1, &v2}
	expected = 3
	actual = ReduceInt16Ptr(plusInt16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int16{&v1}
	expected = 1
	actual = ReduceInt16Ptr(plusInt16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int16{}
	expected = 0
	actual = ReduceInt16Ptr(plusInt16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt16Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusInt16Ptr(num1, num2 *int16) *int16 {
	c := *num1 + *num2
	return &c
}

func TestReduceInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	list := []*int8{&v1, &v2, &v3, &v4, &v5}
	var expected int8 = 15
	actual := ReduceInt8Ptr(plusInt8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int8{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceInt8Ptr(plusInt8Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceInt8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int8{&v1, &v2}
	expected = 3
	actual = ReduceInt8Ptr(plusInt8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int8{&v1}
	expected = 1
	actual = ReduceInt8Ptr(plusInt8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*int8{}
	expected = 0
	actual = ReduceInt8Ptr(plusInt8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt8Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusInt8Ptr(num1, num2 *int8) *int8 {
	c := *num1 + *num2
	return &c
}

func TestReduceUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	list := []*uint{&v1, &v2, &v3, &v4, &v5}
	var expected uint = 15
	actual := ReduceUintPtr(plusUintPtr, list)
	if *actual != expected {
		t.Errorf("ReduceUintPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceUintPtr(plusUintPtr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceUintPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint{&v1, &v2}
	expected = 3
	actual = ReduceUintPtr(plusUintPtr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint{&v1}
	expected = 1
	actual = ReduceUintPtr(plusUintPtr, list)
	if *actual != expected {
		t.Errorf("ReduceUintPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint{}
	expected = 0
	actual = ReduceUintPtr(plusUintPtr, list)
	if *actual != expected {
		t.Errorf("ReduceUintPtr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusUintPtr(num1, num2 *uint) *uint {
	c := *num1 + *num2
	return &c
}

func TestReduceUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	list := []*uint64{&v1, &v2, &v3, &v4, &v5}
	var expected uint64 = 15
	actual := ReduceUint64Ptr(plusUint64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint64{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceUint64Ptr(plusUint64Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceUint64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint64{&v1, &v2}
	expected = 3
	actual = ReduceUint64Ptr(plusUint64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint64{&v1}
	expected = 1
	actual = ReduceUint64Ptr(plusUint64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint64{}
	expected = 0
	actual = ReduceUint64Ptr(plusUint64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint64Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusUint64Ptr(num1, num2 *uint64) *uint64 {
	c := *num1 + *num2
	return &c
}

func TestReduceUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	list := []*uint32{&v1, &v2, &v3, &v4, &v5}
	var expected uint32 = 15
	actual := ReduceUint32Ptr(plusUint32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint32{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceUint32Ptr(plusUint32Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceUint32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint32{&v1, &v2}
	expected = 3
	actual = ReduceUint32Ptr(plusUint32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint32{&v1}
	expected = 1
	actual = ReduceUint32Ptr(plusUint32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint32{}
	expected = 0
	actual = ReduceUint32Ptr(plusUint32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint32Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusUint32Ptr(num1, num2 *uint32) *uint32 {
	c := *num1 + *num2
	return &c
}

func TestReduceUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	list := []*uint16{&v1, &v2, &v3, &v4, &v5}
	var expected uint16 = 15
	actual := ReduceUint16Ptr(plusUint16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint16{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceUint16Ptr(plusUint16Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceUint16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint16{&v1, &v2}
	expected = 3
	actual = ReduceUint16Ptr(plusUint16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint16{&v1}
	expected = 1
	actual = ReduceUint16Ptr(plusUint16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint16Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint16{}
	expected = 0
	actual = ReduceUint16Ptr(plusUint16Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint16Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusUint16Ptr(num1, num2 *uint16) *uint16 {
	c := *num1 + *num2
	return &c
}

func TestReduceUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	list := []*uint8{&v1, &v2, &v3, &v4, &v5}
	var expected uint8 = 15
	actual := ReduceUint8Ptr(plusUint8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint8{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceUint8Ptr(plusUint8Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceUint8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint8{&v1, &v2}
	expected = 3
	actual = ReduceUint8Ptr(plusUint8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint8{&v1}
	expected = 1
	actual = ReduceUint8Ptr(plusUint8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint8Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*uint8{}
	expected = 0
	actual = ReduceUint8Ptr(plusUint8Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceUint8Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusUint8Ptr(num1, num2 *uint8) *uint8 {
	c := *num1 + *num2
	return &c
}

func TestReduceStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	var expected string = "12345"
	actual := ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2, &v3, &v4, &v5}
	expected = "312345"
	actual = ReduceStrPtr(plusStrPtr, list, "3")
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2}
	expected = "12"
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1}
	expected = "1"
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{}
	expected = ""
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusStrPtr(num1, num2 *string) *string {
	c := *num1 + *num2
	return &c
}

func TestReduceFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	list := []*float32{&v1, &v2, &v3, &v4, &v5}
	var expected float32 = 15
	actual := ReduceFloat32Ptr(plusFloat32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float32{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceFloat32Ptr(plusFloat32Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceFloat32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float32{&v1, &v2}
	expected = 3
	actual = ReduceFloat32Ptr(plusFloat32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float32{&v1}
	expected = 1
	actual = ReduceFloat32Ptr(plusFloat32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat32Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float32{}
	expected = 0
	actual = ReduceFloat32Ptr(plusFloat32Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat32Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusFloat32Ptr(num1, num2 *float32) *float32 {
	c := *num1 + *num2
	return &c
}

func TestReduceFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	list := []*float64{&v1, &v2, &v3, &v4, &v5}
	var expected float64 = 15
	actual := ReduceFloat64Ptr(plusFloat64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float64{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceFloat64Ptr(plusFloat64Ptr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceFloat64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float64{&v1, &v2}
	expected = 3
	actual = ReduceFloat64Ptr(plusFloat64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float64{&v1}
	expected = 1
	actual = ReduceFloat64Ptr(plusFloat64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat64Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*float64{}
	expected = 0
	actual = ReduceFloat64Ptr(plusFloat64Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceFloat64Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plusFloat64Ptr(num1, num2 *float64) *float64 {
	c := *num1 + *num2
	return &c
}
