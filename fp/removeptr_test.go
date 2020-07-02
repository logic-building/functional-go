package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestRemoveIntPtr(t *testing.T) {
	// Test : even number in the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v20 int = 20
	var v40 int = 40

	expectedNewList := []*int{&v1, &v3}
	NewList := RemoveIntPtr(isEvenIntPtr, []*int{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int{&v1, &v3}
	partialIsEven := func(num *int) bool {
		return *num%10 == 0
	}
	NewList = RemoveIntPtr(partialIsEven, []*int{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveIntPtr(nil, nil)) > 0 {
		t.Errorf("RemoveIntPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveInt64Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v20 int64 = 20
	var v40 int64 = 40

	expectedNewList := []*int64{&v1, &v3}
	NewList := RemoveInt64Ptr(isEvenInt64Ptr, []*int64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int64{&v1, &v3}
	partialIsEven := func(num *int64) bool {
		return *num%10 == 0
	}
	NewList = RemoveInt64Ptr(partialIsEven, []*int64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveInt64Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveInt64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveInt32Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v20 int32 = 20
	var v40 int32 = 40

	expectedNewList := []*int32{&v1, &v3}
	NewList := RemoveInt32Ptr(isEvenInt32Ptr, []*int32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int32{&v1, &v3}
	partialIsEven := func(num *int32) bool {
		return *num%10 == 0
	}
	NewList = RemoveInt32Ptr(partialIsEven, []*int32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveInt32Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveInt32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveInt16Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v20 int16 = 20
	var v40 int16 = 40

	expectedNewList := []*int16{&v1, &v3}
	NewList := RemoveInt16Ptr(isEvenInt16Ptr, []*int16{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int16{&v1, &v3}
	partialIsEven := func(num *int16) bool {
		return *num%10 == 0
	}
	NewList = RemoveInt16Ptr(partialIsEven, []*int16{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveInt16Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveInt16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveInt8Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v20 int8 = 20
	var v40 int8 = 40

	expectedNewList := []*int8{&v1, &v3}
	NewList := RemoveInt8Ptr(isEvenInt8Ptr, []*int8{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int8{&v1, &v3}
	partialIsEven := func(num *int8) bool {
		return *num%10 == 0
	}
	NewList = RemoveInt8Ptr(partialIsEven, []*int8{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveInt8Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveInt8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveUintPtr(t *testing.T) {
	// Test : even number in the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v20 uint = 20
	var v40 uint = 40

	expectedNewList := []*uint{&v1, &v3}
	NewList := RemoveUintPtr(isEvenUintPtr, []*uint{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint{&v1, &v3}
	partialIsEven := func(num *uint) bool {
		return *num%10 == 0
	}
	NewList = RemoveUintPtr(partialIsEven, []*uint{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveUintPtr(nil, nil)) > 0 {
		t.Errorf("RemoveUintPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveUint64Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v20 uint64 = 20
	var v40 uint64 = 40

	expectedNewList := []*uint64{&v1, &v3}
	NewList := RemoveUint64Ptr(isEvenUint64Ptr, []*uint64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint64{&v1, &v3}
	partialIsEven := func(num *uint64) bool {
		return *num%10 == 0
	}
	NewList = RemoveUint64Ptr(partialIsEven, []*uint64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveUint64Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveUint64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveUint32Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v20 uint32 = 20
	var v40 uint32 = 40

	expectedNewList := []*uint32{&v1, &v3}
	NewList := RemoveUint32Ptr(isEvenUint32Ptr, []*uint32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint32{&v1, &v3}
	partialIsEven := func(num *uint32) bool {
		return *num%10 == 0
	}
	NewList = RemoveUint32Ptr(partialIsEven, []*uint32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveUint32Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveUint32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveUint16Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v20 uint16 = 20
	var v40 uint16 = 40

	expectedNewList := []*uint16{&v1, &v3}
	NewList := RemoveUint16Ptr(isEvenUint16Ptr, []*uint16{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint16{&v1, &v3}
	partialIsEven := func(num *uint16) bool {
		return *num%10 == 0
	}
	NewList = RemoveUint16Ptr(partialIsEven, []*uint16{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveUint16Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveUint16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveUint8Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v20 uint8 = 20
	var v40 uint8 = 40

	expectedNewList := []*uint8{&v1, &v3}
	NewList := RemoveUint8Ptr(isEvenUint8Ptr, []*uint8{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint8{&v1, &v3}
	partialIsEven := func(num *uint8) bool {
		return *num%10 == 0
	}
	NewList = RemoveUint8Ptr(partialIsEven, []*uint8{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveUint8Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveUint8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveStrPtr(t *testing.T) {
	// Test : even number in the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v20 string = "20"
	var v40 string = "40"

	expectedNewList := []*string{&v1, &v3}
	NewList := RemoveStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*string{&v1, &v3}
	partialIsEven := func(num *string) bool {
		if *num == "10" || *num == "20" {
			return true
		}
		return false
	}
	NewList = RemoveStrPtr(partialIsEven, []*string{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveStrPtr(nil, nil)) > 0 {
		t.Errorf("RemoveStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveBoolPtr(t *testing.T) {
	// Test : even number in the list
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vt}
	NewList := RemoveBoolPtr(func(v *bool) bool { return *v == false }, []*bool{&vt, &vf, &vf})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("RemoveBool failed. Expected New list=%v, actual list=%v", *expectedNewList[0], *NewList[0])
	}

	if len(RemoveBoolPtr(nil, nil)) > 0 {
		t.Errorf("RemoveBoolPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveFloat32Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v20 float32 = 20
	var v40 float32 = 40

	expectedNewList := []*float32{&v1, &v3}
	NewList := RemoveFloat32Ptr(isEvenFloat32Ptr, []*float32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float32{&v1, &v3}
	partialIsEven := func(num *float32) bool {
		return int(*num)%10 == 0
	}
	NewList = RemoveFloat32Ptr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestRemoveFloat64Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v20 float64 = 20
	var v40 float64 = 40

	expectedNewList := []*float64{&v1, &v3}
	NewList := RemoveFloat64Ptr(isEvenFloat64Ptr, []*float64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float64{&v1, &v3}
	partialIsEven := func(num *float64) bool {
		return int(*num)%10 == 0
	}
	NewList = RemoveFloat64Ptr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
