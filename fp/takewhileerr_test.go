package fp

import (
	"errors"
	"testing"
)

func TestTakeWhileIntErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v7 int = 7
	var v40 int = 40
	var v0 int

	expectedNewList := []int{v4, v2, v4}
	NewList, _ := TakeWhileIntErr(isEvenIntErr, []int{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileIntErr(isEvenIntErr, []int{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int{v40}
	partialIsEvenDivisibleBy := func(num int) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileIntErr(partialIsEvenDivisibleBy, []int{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileIntErr failed.")
	}

	r, _ = TakeWhileIntErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("TakeWhileIntErr failed.")
	}
}

func TestTakeWhileInt64Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v7 int64 = 7
	var v40 int64 = 40
	var v0 int64

	expectedNewList := []int64{v4, v2, v4}
	NewList, _ := TakeWhileInt64Err(isEvenInt64Err, []int64{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt64Err(isEvenInt64Err, []int64{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int64{v40}
	partialIsEvenDivisibleBy := func(num int64) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileInt64Err(partialIsEvenDivisibleBy, []int64{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt64Err failed.")
	}

	r, _ = TakeWhileInt64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt64Err failed.")
	}
}

func TestTakeWhileInt32Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v7 int32 = 7
	var v40 int32 = 40
	var v0 int32

	expectedNewList := []int32{v4, v2, v4}
	NewList, _ := TakeWhileInt32Err(isEvenInt32Err, []int32{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt32Err(isEvenInt32Err, []int32{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int32{v40}
	partialIsEvenDivisibleBy := func(num int32) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileInt32Err(partialIsEvenDivisibleBy, []int32{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt32Err failed.")
	}

	r, _ = TakeWhileInt32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt32Err failed.")
	}
}

func TestTakeWhileInt16Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v7 int16 = 7
	var v40 int16 = 40
	var v0 int16

	expectedNewList := []int16{v4, v2, v4}
	NewList, _ := TakeWhileInt16Err(isEvenInt16Err, []int16{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt16Err(isEvenInt16Err, []int16{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int16{v40}
	partialIsEvenDivisibleBy := func(num int16) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileInt16Err(partialIsEvenDivisibleBy, []int16{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt16Err failed.")
	}

	r, _ = TakeWhileInt16Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt16Err failed.")
	}
}

func TestTakeWhileInt8Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v7 int8 = 7
	var v40 int8 = 40
	var v0 int8

	expectedNewList := []int8{v4, v2, v4}
	NewList, _ := TakeWhileInt8Err(isEvenInt8Err, []int8{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt8Err(isEvenInt8Err, []int8{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int8{v40}
	partialIsEvenDivisibleBy := func(num int8) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileInt8Err(partialIsEvenDivisibleBy, []int8{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt8Err failed.")
	}

	r, _ = TakeWhileInt8Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt8Err failed.")
	}
}

func TestTakeWhileUintErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v7 uint = 7
	var v40 uint = 40
	var v0 uint

	expectedNewList := []uint{v4, v2, v4}
	NewList, _ := TakeWhileUintErr(isEvenUintErr, []uint{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUintErr(isEvenUintErr, []uint{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint{v40}
	partialIsEvenDivisibleBy := func(num uint) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileUintErr(partialIsEvenDivisibleBy, []uint{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUintErr failed.")
	}

	r, _ = TakeWhileUintErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUintErr failed.")
	}
}

func TestTakeWhileUint64Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v7 uint64 = 7
	var v40 uint64 = 40
	var v0 uint64

	expectedNewList := []uint64{v4, v2, v4}
	NewList, _ := TakeWhileUint64Err(isEvenUint64Err, []uint64{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint64Err(isEvenUint64Err, []uint64{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint64{v40}
	partialIsEvenDivisibleBy := func(num uint64) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileUint64Err(partialIsEvenDivisibleBy, []uint64{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint64Err failed.")
	}

	r, _ = TakeWhileUint64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint64Err failed.")
	}
}

func TestTakeWhileUint32Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v7 uint32 = 7
	var v40 uint32 = 40
	var v0 uint32

	expectedNewList := []uint32{v4, v2, v4}
	NewList, _ := TakeWhileUint32Err(isEvenUint32Err, []uint32{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint32Err(isEvenUint32Err, []uint32{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint32{v40}
	partialIsEvenDivisibleBy := func(num uint32) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileUint32Err(partialIsEvenDivisibleBy, []uint32{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint32Err failed.")
	}

	r, _ = TakeWhileUint32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint32Err failed.")
	}
}

func TestTakeWhileUint16Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v7 uint16 = 7
	var v40 uint16 = 40
	var v0 uint16

	expectedNewList := []uint16{v4, v2, v4}
	NewList, _ := TakeWhileUint16Err(isEvenUint16Err, []uint16{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint16Err(isEvenUint16Err, []uint16{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint16{v40}
	partialIsEvenDivisibleBy := func(num uint16) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileUint16Err(partialIsEvenDivisibleBy, []uint16{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint16Err failed.")
	}

	r, _ = TakeWhileUint16Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint16Err failed.")
	}
}

func TestTakeWhileUint8Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v7 uint8 = 7
	var v40 uint8 = 40
	var v0 uint8

	expectedNewList := []uint8{v4, v2, v4}
	NewList, _ := TakeWhileUint8Err(isEvenUint8Err, []uint8{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint8Err(isEvenUint8Err, []uint8{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint8{v40}
	partialIsEvenDivisibleBy := func(num uint8) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileUint8Err(partialIsEvenDivisibleBy, []uint8{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint8Err failed.")
	}

	r, _ = TakeWhileUint8Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint8Err failed.")
	}
}

func TestTakeWhileStrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string = "0"

	expectedNewList := []string{v4, v2, v4}
	NewList, _ := TakeWhileStrErr(isEvenStrErr, []string{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrErr(isEvenStrErr, []string{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []string{v40}
	partialIsEvenDivisibleBy := func(num string) (bool, error) { if num == "40" { return true, nil}; return false, nil }
	NewList, _ = TakeWhileStrErr(partialIsEvenDivisibleBy, []string{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}

	r, _ = TakeWhileStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}
}

func TestTakeWhileBoolErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vt, vt, vf}
	NewList, _ := TakeWhileBoolErr(func(v bool) (bool, error) {return v == true, nil}, []bool{vt, vt, vf, vf, vf})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("TakeWhileBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []bool{vt}
	NewList, _ = TakeWhileBoolErr(func(v bool) (bool, error) {return v == true, nil}, []bool{vt})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileBoolErr(func(v bool) (bool, error) { if v == false { return false, errors.New("false is invalid for this test") }; return v == true, nil}, []bool{vf})
	
	if err == nil {
		t.Errorf("TakeWhileBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileBoolErr failed.")
	}

	r, _ = TakeWhileBoolErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("TakeWhileBoolErr failed.")
	}
}

func TestTakeWhileFloat32Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v7 float32 = 7
	var v40 float32 = 40
	var v0 float32

	expectedNewList := []float32{v4, v2, v4}
	NewList, _ := TakeWhileFloat32Err(isEvenFloat32Err, []float32{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileFloat32Err(isEvenFloat32Err, []float32{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []float32{v40}
	partialIsEvenDivisibleBy := func(num float32) (bool, error) { return int(num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat32Err(partialIsEvenDivisibleBy, []float32{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat32Err failed.")
	}

	r, _ = TakeWhileFloat32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat32Err failed.")
	}
}

func TestTakeWhileFloat64Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v7 float64 = 7
	var v40 float64 = 40
	var v0 float64

	expectedNewList := []float64{v4, v2, v4}
	NewList, _ := TakeWhileFloat64Err(isEvenFloat64Err, []float64{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileFloat64Err(isEvenFloat64Err, []float64{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []float64{v40}
	partialIsEvenDivisibleBy := func(num float64) (bool, error) { return int(num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat64Err(partialIsEvenDivisibleBy, []float64{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat64Err failed.")
	}

	r, _ = TakeWhileFloat64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat64Err failed.")
	}
}
