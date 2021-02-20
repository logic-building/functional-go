package fp

import (
	"errors"
	"testing"
)

func TestDropWhileIntPtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	var v0 int

	expectedNewList := []*int{&v3, &v4, &v5}
	NewList, _ := DropWhileIntPtrErr(isEvenIntPtrErr, []*int{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileIntPtrErr(isEvenIntPtrErr, []*int{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileIntPtr failed.")
	}

	r, _ = DropWhileIntPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("DropWhileIntPtr failed.")
	}

	NewList, _ = DropWhileIntPtrErr(isEvenIntPtrErr, []*int{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileIntPtrErr failed")
	}
}

func TestDropWhileInt64PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	var v0 int64

	expectedNewList := []*int64{&v3, &v4, &v5}
	NewList, _ := DropWhileInt64PtrErr(isEvenInt64PtrErr, []*int64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt64PtrErr(isEvenInt64PtrErr, []*int64{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt64Ptr failed.")
	}

	r, _ = DropWhileInt64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt64Ptr failed.")
	}

	NewList, _ = DropWhileInt64PtrErr(isEvenInt64PtrErr, []*int64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt64PtrErr failed")
	}
}

func TestDropWhileInt32PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	var v0 int32

	expectedNewList := []*int32{&v3, &v4, &v5}
	NewList, _ := DropWhileInt32PtrErr(isEvenInt32PtrErr, []*int32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt32PtrErr(isEvenInt32PtrErr, []*int32{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt32Ptr failed.")
	}

	r, _ = DropWhileInt32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt32Ptr failed.")
	}

	NewList, _ = DropWhileInt32PtrErr(isEvenInt32PtrErr, []*int32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt32PtrErr failed")
	}
}

func TestDropWhileInt16PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	var v0 int16

	expectedNewList := []*int16{&v3, &v4, &v5}
	NewList, _ := DropWhileInt16PtrErr(isEvenInt16PtrErr, []*int16{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt16PtrErr(isEvenInt16PtrErr, []*int16{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt16Ptr failed.")
	}

	r, _ = DropWhileInt16PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt16Ptr failed.")
	}

	NewList, _ = DropWhileInt16PtrErr(isEvenInt16PtrErr, []*int16{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt16PtrErr failed")
	}
}

func TestDropWhileInt8PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	var v0 int8

	expectedNewList := []*int8{&v3, &v4, &v5}
	NewList, _ := DropWhileInt8PtrErr(isEvenInt8PtrErr, []*int8{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt8PtrErr(isEvenInt8PtrErr, []*int8{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt8Ptr failed.")
	}

	r, _ = DropWhileInt8PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt8Ptr failed.")
	}

	NewList, _ = DropWhileInt8PtrErr(isEvenInt8PtrErr, []*int8{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt8PtrErr failed")
	}
}

func TestDropWhileUintPtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	var v0 uint

	expectedNewList := []*uint{&v3, &v4, &v5}
	NewList, _ := DropWhileUintPtrErr(isEvenUintPtrErr, []*uint{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUintPtrErr(isEvenUintPtrErr, []*uint{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUintPtr failed.")
	}

	r, _ = DropWhileUintPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("DropWhileUintPtr failed.")
	}

	NewList, _ = DropWhileUintPtrErr(isEvenUintPtrErr, []*uint{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUintPtrErr failed")
	}
}

func TestDropWhileUint64PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v0 uint64

	expectedNewList := []*uint64{&v3, &v4, &v5}
	NewList, _ := DropWhileUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint64Ptr failed.")
	}

	r, _ = DropWhileUint64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint64Ptr failed.")
	}

	NewList, _ = DropWhileUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint64PtrErr failed")
	}
}

func TestDropWhileUint32PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v0 uint32

	expectedNewList := []*uint32{&v3, &v4, &v5}
	NewList, _ := DropWhileUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint32Ptr failed.")
	}

	r, _ = DropWhileUint32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint32Ptr failed.")
	}

	NewList, _ = DropWhileUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint32PtrErr failed")
	}
}

func TestDropWhileUint16PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v0 uint16

	expectedNewList := []*uint16{&v3, &v4, &v5}
	NewList, _ := DropWhileUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint16Ptr failed.")
	}

	r, _ = DropWhileUint16PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint16Ptr failed.")
	}

	NewList, _ = DropWhileUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint16PtrErr failed")
	}
}

func TestDropWhileUint8PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v0 uint8

	expectedNewList := []*uint8{&v3, &v4, &v5}
	NewList, _ := DropWhileUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint8Ptr failed.")
	}

	r, _ = DropWhileUint8PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint8Ptr failed.")
	}

	NewList, _ = DropWhileUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint8PtrErr failed")
	}
}

func TestDropWhileStrPtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	expectedNewList := []*string{&v3, &v4, &v5}
	NewList, _ := DropWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}

	r, _ = DropWhileStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}

	NewList, _ = DropWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileStrPtrErr failed")
	}
}

func TestDropWhileBoolPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vf, &vt}
	NewList, _ := DropWhileBoolPtrErr(isTrueBoolPtrErr, []*bool{&vt, &vf, &vt})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhileBoolPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileBoolPtrErr(isTrueBoolPtrErr2, []*bool{&vt, &vf, &vt})
	if err == nil {
		t.Errorf("DropWhileBoolPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileBoolPtr failed.")
	}

	r, _ = DropWhileBoolPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("DropWhileBoolPtr failed.")
	}

	r, _ = DropWhileBoolPtrErr(isTrueBoolPtrErr2, []*bool{})
	if len(r) > 0 {
		t.Errorf("DropWhileBoolPtr failed.")
	}
}

func isTrueBoolPtrErr(v *bool) (bool, error) {
	return *v == true, nil
}
func isTrueBoolPtrErr2(v *bool) (bool, error) {
	if *v == false {
		return false, errors.New("false is invalid for this test")
	}
	return *v == true, nil
}

func TestDropWhileFloat32PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	var v0 float32

	expectedNewList := []*float32{&v3, &v4, &v5}
	NewList, _ := DropWhileFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileFloat32Ptr failed.")
	}

	r, _ = DropWhileFloat32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("DropWhileFloat32Ptr failed.")
	}

	NewList, _ = DropWhileFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat32PtrErr failed")
	}
}

func TestDropWhileFloat64PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	var v0 float64

	expectedNewList := []*float64{&v3, &v4, &v5}
	NewList, _ := DropWhileFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhileFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileFloat64Ptr failed.")
	}

	r, _ = DropWhileFloat64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("DropWhileFloat64Ptr failed.")
	}

	NewList, _ = DropWhileFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat64PtrErr failed")
	}
}
