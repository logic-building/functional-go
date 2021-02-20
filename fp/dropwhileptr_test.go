package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDropWhileIntPtr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	expectedNewList := []*int{&v3, &v4, &v5}
	NewList := DropWhileIntPtr(isEvenIntPtr, []*int{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileIntPtr(nil, nil)) > 0 {
		t.Errorf("DropWhileIntPtr failed.")
	}

	if len(DropWhileIntPtr(nil, []*int{})) > 0 {
		t.Errorf("DropWhileIntPtr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileIntPtr(isEvenIntPtr, []*int{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileIntPtr failed")
	}
}

func TestDropWhileInt64Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	expectedNewList := []*int64{&v3, &v4, &v5}
	NewList := DropWhileInt64Ptr(isEvenInt64Ptr, []*int64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt64Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileInt64Ptr failed.")
	}

	if len(DropWhileInt64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("DropWhileInt64Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileInt64Ptr(isEvenInt64Ptr, []*int64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt64Ptr failed")
	}
}

func TestDropWhileInt32Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	expectedNewList := []*int32{&v3, &v4, &v5}
	NewList := DropWhileInt32Ptr(isEvenInt32Ptr, []*int32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt32Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileInt32Ptr failed.")
	}

	if len(DropWhileInt32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("DropWhileInt32Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileInt32Ptr(isEvenInt32Ptr, []*int32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt32Ptr failed")
	}
}

func TestDropWhileInt16Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	expectedNewList := []*int16{&v3, &v4, &v5}
	NewList := DropWhileInt16Ptr(isEvenInt16Ptr, []*int16{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt16Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileInt16Ptr failed.")
	}

	if len(DropWhileInt16Ptr(nil, []*int16{})) > 0 {
		t.Errorf("DropWhileInt16Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileInt16Ptr(isEvenInt16Ptr, []*int16{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt16Ptr failed")
	}
}

func TestDropWhileInt8Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	expectedNewList := []*int8{&v3, &v4, &v5}
	NewList := DropWhileInt8Ptr(isEvenInt8Ptr, []*int8{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt8Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileInt8Ptr failed.")
	}

	if len(DropWhileInt8Ptr(nil, []*int8{})) > 0 {
		t.Errorf("DropWhileInt8Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileInt8Ptr(isEvenInt8Ptr, []*int8{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt8Ptr failed")
	}
}

func TestDropWhileUintPtr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	expectedNewList := []*uint{&v3, &v4, &v5}
	NewList := DropWhileUintPtr(isEvenUintPtr, []*uint{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUintPtr(nil, nil)) > 0 {
		t.Errorf("DropWhileUintPtr failed.")
	}

	if len(DropWhileUintPtr(nil, []*uint{})) > 0 {
		t.Errorf("DropWhileUintPtr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileUintPtr(isEvenUintPtr, []*uint{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUintPtr failed")
	}
}

func TestDropWhileUint64Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	expectedNewList := []*uint64{&v3, &v4, &v5}
	NewList := DropWhileUint64Ptr(isEvenUint64Ptr, []*uint64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint64Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileUint64Ptr failed.")
	}

	if len(DropWhileUint64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("DropWhileUint64Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileUint64Ptr(isEvenUint64Ptr, []*uint64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint64Ptr failed")
	}
}

func TestDropWhileUint32Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	expectedNewList := []*uint32{&v3, &v4, &v5}
	NewList := DropWhileUint32Ptr(isEvenUint32Ptr, []*uint32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint32Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileUint32Ptr failed.")
	}

	if len(DropWhileUint32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("DropWhileUint32Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileUint32Ptr(isEvenUint32Ptr, []*uint32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint32Ptr failed")
	}
}

func TestDropWhileUint16Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	expectedNewList := []*uint16{&v3, &v4, &v5}
	NewList := DropWhileUint16Ptr(isEvenUint16Ptr, []*uint16{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint16Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileUint16Ptr failed.")
	}

	if len(DropWhileUint16Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("DropWhileUint16Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileUint16Ptr(isEvenUint16Ptr, []*uint16{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint16Ptr failed")
	}
}

func TestDropWhileUint8Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	expectedNewList := []*uint8{&v3, &v4, &v5}
	NewList := DropWhileUint8Ptr(isEvenUint8Ptr, []*uint8{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint8Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileUint8Ptr failed.")
	}

	if len(DropWhileUint8Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("DropWhileUint8Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileUint8Ptr(isEvenUint8Ptr, []*uint8{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint8Ptr failed")
	}
}

func TestDropWhileStrPtr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	expectedNewList := []*string{&v3, &v4, &v5}
	NewList := DropWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileStrPtr(nil, nil)) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}

	if len(DropWhileStrPtr(nil, []*string{})) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}
	NewList = DropWhileStrPtr(isEvenStrPtr, []*string{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileStrPtr failed")
	}
}

func TestDropWhileBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vf, &vt}
	NewList := DropWhileBoolPtr(isTrueBoolPtr, []*bool{&vt, &vf, &vt})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

func TestDropWhileFloat32Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	expectedNewList := []*float32{&v3, &v4, &v5}
	NewList := DropWhileFloat32Ptr(isEvenFloat32Ptr, []*float32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileFloat32Ptr failed.")
	}

	if len(DropWhileFloat32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("DropWhileFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileFloat32Ptr(isEvenFloat32Ptr, []*float32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat32Ptr failed")
	}
}

func TestDropWhileFloat64Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	expectedNewList := []*float64{&v3, &v4, &v5}
	NewList := DropWhileFloat64Ptr(isEvenFloat64Ptr, []*float64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhileFloat64Ptr failed.")
	}

	if len(DropWhileFloat64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("DropWhileFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhileFloat64Ptr(isEvenFloat64Ptr, []*float64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat64Ptr failed")
	}
}
