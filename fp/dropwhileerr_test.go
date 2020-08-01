package fp

import (
	"errors"
	"testing"
)

func TestDropWhileIntErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	var v0 int

	expectedNewList := []int{v3, v4, v5}
	NewList, _ := DropWhileIntErr(isEvenIntErr, []int{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileIntErr(isEvenIntErr, []int{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileIntErr failed.")
	}

	r, _ = DropWhileIntErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("DropWhileIntErr failed.")
	}

	NewList, _ = DropWhileIntErr(isEvenIntErr, []int{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileIntErr failed")
	}
}

func TestDropWhileInt64Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	var v0 int64

	expectedNewList := []int64{v3, v4, v5}
	NewList, _ := DropWhileInt64Err(isEvenInt64Err, []int64{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt64Err(isEvenInt64Err, []int64{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt64Err failed.")
	}

	r, _ = DropWhileInt64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt64Err failed.")
	}

	NewList, _ = DropWhileInt64Err(isEvenInt64Err, []int64{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt64Err failed")
	}
}

func TestDropWhileInt32Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	var v0 int32

	expectedNewList := []int32{v3, v4, v5}
	NewList, _ := DropWhileInt32Err(isEvenInt32Err, []int32{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt32Err(isEvenInt32Err, []int32{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt32Err failed.")
	}

	r, _ = DropWhileInt32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt32Err failed.")
	}

	NewList, _ = DropWhileInt32Err(isEvenInt32Err, []int32{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt32Err failed")
	}
}

func TestDropWhileInt16Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	var v0 int16

	expectedNewList := []int16{v3, v4, v5}
	NewList, _ := DropWhileInt16Err(isEvenInt16Err, []int16{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt16Err(isEvenInt16Err, []int16{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt16Err failed.")
	}

	r, _ = DropWhileInt16Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt16Err failed.")
	}

	NewList, _ = DropWhileInt16Err(isEvenInt16Err, []int16{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt16Err failed")
	}
}

func TestDropWhileInt8Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	var v0 int8

	expectedNewList := []int8{v3, v4, v5}
	NewList, _ := DropWhileInt8Err(isEvenInt8Err, []int8{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileInt8Err(isEvenInt8Err, []int8{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileInt8Err failed.")
	}

	r, _ = DropWhileInt8Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("DropWhileInt8Err failed.")
	}

	NewList, _ = DropWhileInt8Err(isEvenInt8Err, []int8{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt8Err failed")
	}
}

func TestDropWhileUintErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	var v0 uint

	expectedNewList := []uint{v3, v4, v5}
	NewList, _ := DropWhileUintErr(isEvenUintErr, []uint{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUintErr(isEvenUintErr, []uint{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUintErr failed.")
	}

	r, _ = DropWhileUintErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("DropWhileUintErr failed.")
	}

	NewList, _ = DropWhileUintErr(isEvenUintErr, []uint{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUintErr failed")
	}
}

func TestDropWhileUint64Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v0 uint64

	expectedNewList := []uint64{v3, v4, v5}
	NewList, _ := DropWhileUint64Err(isEvenUint64Err, []uint64{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint64Err(isEvenUint64Err, []uint64{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint64Err failed.")
	}

	r, _ = DropWhileUint64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint64Err failed.")
	}

	NewList, _ = DropWhileUint64Err(isEvenUint64Err, []uint64{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint64Err failed")
	}
}

func TestDropWhileUint32Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v0 uint32

	expectedNewList := []uint32{v3, v4, v5}
	NewList, _ := DropWhileUint32Err(isEvenUint32Err, []uint32{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint32Err(isEvenUint32Err, []uint32{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint32Err failed.")
	}

	r, _ = DropWhileUint32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint32Err failed.")
	}

	NewList, _ = DropWhileUint32Err(isEvenUint32Err, []uint32{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint32Err failed")
	}
}

func TestDropWhileUint16Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v0 uint16

	expectedNewList := []uint16{v3, v4, v5}
	NewList, _ := DropWhileUint16Err(isEvenUint16Err, []uint16{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint16Err(isEvenUint16Err, []uint16{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint16Err failed.")
	}

	r, _ = DropWhileUint16Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint16Err failed.")
	}

	NewList, _ = DropWhileUint16Err(isEvenUint16Err, []uint16{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint16Err failed")
	}
}

func TestDropWhileUint8Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v0 uint8

	expectedNewList := []uint8{v3, v4, v5}
	NewList, _ := DropWhileUint8Err(isEvenUint8Err, []uint8{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileUint8Err(isEvenUint8Err, []uint8{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileUint8Err failed.")
	}

	r, _ = DropWhileUint8Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("DropWhileUint8Err failed.")
	}

	NewList, _ = DropWhileUint8Err(isEvenUint8Err, []uint8{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint8Err failed")
	}
}

func TestDropWhileStrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	expectedNewList := []string{v3, v4, v5}
	NewList, _ := DropWhileStrErr(isEvenStrErr, []string{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileStrErr(isEvenStrErr, []string{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileStrErr failed.")
	}

	r, _ = DropWhileStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("DropWhileStrErr failed.")
	}

	NewList, _ = DropWhileStrErr(isEvenStrErr, []string{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileStrErr failed")
	}
}

func TestDropWhileBoolErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vf, vt}
	NewList, _ := DropWhileBoolErr(isTrueBoolErr, []bool{vt, vf, vt})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("DropWhileBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileBoolErr(isTrueBoolErr2, []bool{vt, vf, vt})
	if err == nil {
		t.Errorf("DropWhileBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileBoolErr failed.")
	}

	r, _ = DropWhileBoolErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("DropWhileBoolErr failed.")
	}

	r, _ = DropWhileBoolErr(isTrueBoolErr2, []bool{})
	if len(r) > 0 {
		t.Errorf("DropWhileBoolErr failed.")
	}
}

func isTrueBoolErr(v bool) (bool, error) {
	return v == true, nil
}
func isTrueBoolErr2(v bool) (bool, error) {
	if v == false {
		return false, errors.New("false is invalid for this test")
	}
	return v == true, nil
}

func TestDropWhileFloat32Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	var v0 float32

	expectedNewList := []float32{v3, v4, v5}
	NewList, _ := DropWhileFloat32Err(isEvenFloat32Err, []float32{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileFloat32Err(isEvenFloat32Err, []float32{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileFloat32Err failed.")
	}

	r, _ = DropWhileFloat32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("DropWhileFloat32Err failed.")
	}

	NewList, _ = DropWhileFloat32Err(isEvenFloat32Err, []float32{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat32Err failed")
	}
}

func TestDropWhileFloat64Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	var v0 float64

	expectedNewList := []float64{v3, v4, v5}
	NewList, _ := DropWhileFloat64Err(isEvenFloat64Err, []float64{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhileFloat64Err(isEvenFloat64Err, []float64{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhileFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhileFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhileFloat64Err failed.")
	}

	r, _ = DropWhileFloat64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("DropWhileFloat64Err failed.")
	}

	NewList, _ = DropWhileFloat64Err(isEvenFloat64Err, []float64{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat64Err failed")
	}
}
