package fp

import (
	"errors"
	"testing"
)

func TestReduceIntErr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	var v0 int = 0

	list := []int{v1, v2, v3, v4, v5}
	var expected int = 15
	actual, _ := ReduceIntErr(plusIntErr, list)
	if actual != expected {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []int{v1, v0, v3, v4, v5}
	_, err := ReduceIntErr(plusIntErr, list2)
	if err == nil {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceIntErr(plusIntErr, list, 3)
	if actual != expected {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{v1, v2}
	expected = 3
	actual, _ = ReduceIntErr(plusIntErr, list)
	if actual != expected {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{v1}
	expected = 1
	actual, _ = ReduceIntErr(plusIntErr, list)
	if actual != expected {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{}
	expected = 0
	actual, _ = ReduceIntErr(plusIntErr, list)
	if actual != expected {
		t.Errorf("ReduceIntErr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusIntErr(num1, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return int(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceInt64Err(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	var v0 int64 = 0

	list := []int64{v1, v2, v3, v4, v5}
	var expected int64 = 15
	actual, _ := ReduceInt64Err(plusInt64Err, list)
	if actual != expected {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []int64{v1, v0, v3, v4, v5}
	_, err := ReduceInt64Err(plusInt64Err, list2)
	if err == nil {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceInt64Err(plusInt64Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{v1, v2}
	expected = 3
	actual, _ = ReduceInt64Err(plusInt64Err, list)
	if actual != expected {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{v1}
	expected = 1
	actual, _ = ReduceInt64Err(plusInt64Err, list)
	if actual != expected {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{}
	expected = 0
	actual, _ = ReduceInt64Err(plusInt64Err, list)
	if actual != expected {
		t.Errorf("ReduceInt64Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt64Err(num1, num2 int64) (int64, error) {
	if num1 == 0 || num2 == 0 {
		return int64(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceInt32Err(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	var v0 int32 = 0

	list := []int32{v1, v2, v3, v4, v5}
	var expected int32 = 15
	actual, _ := ReduceInt32Err(plusInt32Err, list)
	if actual != expected {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []int32{v1, v0, v3, v4, v5}
	_, err := ReduceInt32Err(plusInt32Err, list2)
	if err == nil {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceInt32Err(plusInt32Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{v1, v2}
	expected = 3
	actual, _ = ReduceInt32Err(plusInt32Err, list)
	if actual != expected {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{v1}
	expected = 1
	actual, _ = ReduceInt32Err(plusInt32Err, list)
	if actual != expected {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{}
	expected = 0
	actual, _ = ReduceInt32Err(plusInt32Err, list)
	if actual != expected {
		t.Errorf("ReduceInt32Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt32Err(num1, num2 int32) (int32, error) {
	if num1 == 0 || num2 == 0 {
		return int32(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceInt16Err(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	var v0 int16 = 0

	list := []int16{v1, v2, v3, v4, v5}
	var expected int16 = 15
	actual, _ := ReduceInt16Err(plusInt16Err, list)
	if actual != expected {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []int16{v1, v0, v3, v4, v5}
	_, err := ReduceInt16Err(plusInt16Err, list2)
	if err == nil {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceInt16Err(plusInt16Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{v1, v2}
	expected = 3
	actual, _ = ReduceInt16Err(plusInt16Err, list)
	if actual != expected {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{v1}
	expected = 1
	actual, _ = ReduceInt16Err(plusInt16Err, list)
	if actual != expected {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{}
	expected = 0
	actual, _ = ReduceInt16Err(plusInt16Err, list)
	if actual != expected {
		t.Errorf("ReduceInt16Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt16Err(num1, num2 int16) (int16, error) {
	if num1 == 0 || num2 == 0 {
		return int16(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceInt8Err(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	var v0 int8 = 0

	list := []int8{v1, v2, v3, v4, v5}
	var expected int8 = 15
	actual, _ := ReduceInt8Err(plusInt8Err, list)
	if actual != expected {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []int8{v1, v0, v3, v4, v5}
	_, err := ReduceInt8Err(plusInt8Err, list2)
	if err == nil {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceInt8Err(plusInt8Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{v1, v2}
	expected = 3
	actual, _ = ReduceInt8Err(plusInt8Err, list)
	if actual != expected {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{v1}
	expected = 1
	actual, _ = ReduceInt8Err(plusInt8Err, list)
	if actual != expected {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{}
	expected = 0
	actual, _ = ReduceInt8Err(plusInt8Err, list)
	if actual != expected {
		t.Errorf("ReduceInt8Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt8Err(num1, num2 int8) (int8, error) {
	if num1 == 0 || num2 == 0 {
		return int8(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceUintErr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	var v0 uint = 0

	list := []uint{v1, v2, v3, v4, v5}
	var expected uint = 15
	actual, _ := ReduceUintErr(plusUintErr, list)
	if actual != expected {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []uint{v1, v0, v3, v4, v5}
	_, err := ReduceUintErr(plusUintErr, list2)
	if err == nil {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceUintErr(plusUintErr, list, 3)
	if actual != expected {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{v1, v2}
	expected = 3
	actual, _ = ReduceUintErr(plusUintErr, list)
	if actual != expected {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{v1}
	expected = 1
	actual, _ = ReduceUintErr(plusUintErr, list)
	if actual != expected {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{}
	expected = 0
	actual, _ = ReduceUintErr(plusUintErr, list)
	if actual != expected {
		t.Errorf("ReduceUintErr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUintErr(num1, num2 uint) (uint, error) {
	if num1 == 0 || num2 == 0 {
		return uint(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceUint64Err(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v0 uint64 = 0

	list := []uint64{v1, v2, v3, v4, v5}
	var expected uint64 = 15
	actual, _ := ReduceUint64Err(plusUint64Err, list)
	if actual != expected {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []uint64{v1, v0, v3, v4, v5}
	_, err := ReduceUint64Err(plusUint64Err, list2)
	if err == nil {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceUint64Err(plusUint64Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{v1, v2}
	expected = 3
	actual, _ = ReduceUint64Err(plusUint64Err, list)
	if actual != expected {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{v1}
	expected = 1
	actual, _ = ReduceUint64Err(plusUint64Err, list)
	if actual != expected {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{}
	expected = 0
	actual, _ = ReduceUint64Err(plusUint64Err, list)
	if actual != expected {
		t.Errorf("ReduceUint64Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint64Err(num1, num2 uint64) (uint64, error) {
	if num1 == 0 || num2 == 0 {
		return uint64(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceUint32Err(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v0 uint32 = 0

	list := []uint32{v1, v2, v3, v4, v5}
	var expected uint32 = 15
	actual, _ := ReduceUint32Err(plusUint32Err, list)
	if actual != expected {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []uint32{v1, v0, v3, v4, v5}
	_, err := ReduceUint32Err(plusUint32Err, list2)
	if err == nil {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceUint32Err(plusUint32Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{v1, v2}
	expected = 3
	actual, _ = ReduceUint32Err(plusUint32Err, list)
	if actual != expected {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{v1}
	expected = 1
	actual, _ = ReduceUint32Err(plusUint32Err, list)
	if actual != expected {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{}
	expected = 0
	actual, _ = ReduceUint32Err(plusUint32Err, list)
	if actual != expected {
		t.Errorf("ReduceUint32Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint32Err(num1, num2 uint32) (uint32, error) {
	if num1 == 0 || num2 == 0 {
		return uint32(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceUint16Err(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v0 uint16 = 0

	list := []uint16{v1, v2, v3, v4, v5}
	var expected uint16 = 15
	actual, _ := ReduceUint16Err(plusUint16Err, list)
	if actual != expected {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []uint16{v1, v0, v3, v4, v5}
	_, err := ReduceUint16Err(plusUint16Err, list2)
	if err == nil {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceUint16Err(plusUint16Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{v1, v2}
	expected = 3
	actual, _ = ReduceUint16Err(plusUint16Err, list)
	if actual != expected {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{v1}
	expected = 1
	actual, _ = ReduceUint16Err(plusUint16Err, list)
	if actual != expected {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{}
	expected = 0
	actual, _ = ReduceUint16Err(plusUint16Err, list)
	if actual != expected {
		t.Errorf("ReduceUint16Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint16Err(num1, num2 uint16) (uint16, error) {
	if num1 == 0 || num2 == 0 {
		return uint16(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceUint8Err(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v0 uint8 = 0

	list := []uint8{v1, v2, v3, v4, v5}
	var expected uint8 = 15
	actual, _ := ReduceUint8Err(plusUint8Err, list)
	if actual != expected {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []uint8{v1, v0, v3, v4, v5}
	_, err := ReduceUint8Err(plusUint8Err, list2)
	if err == nil {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceUint8Err(plusUint8Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{v1, v2}
	expected = 3
	actual, _ = ReduceUint8Err(plusUint8Err, list)
	if actual != expected {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{v1}
	expected = 1
	actual, _ = ReduceUint8Err(plusUint8Err, list)
	if actual != expected {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{}
	expected = 0
	actual, _ = ReduceUint8Err(plusUint8Err, list)
	if actual != expected {
		t.Errorf("ReduceUint8Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint8Err(num1, num2 uint8) (uint8, error) {
	if num1 == 0 || num2 == 0 {
		return uint8(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	list := []string{v1, v2, v3, v4, v5}
	var expected string = "12345"
	actual, _ := ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []string{v1, v0, v3, v4, v5}
	_, err := ReduceStrErr(plusStrErr, list2)
	if err == nil {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2, v3, v4, v5}
	expected = "312345"
	actual, _ = ReduceStrErr(plusStrErr, list, "3")
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2}
	expected = "12"
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1}
	expected = "1"
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{}
	expected = ""
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusStrErr(num1, num2 string) (string, error) {
	if num1 == "0" || num2 == "0" {
		return "", errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceFloat32Err(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	var v0 float32 = 0

	list := []float32{v1, v2, v3, v4, v5}
	var expected float32 = 15
	actual, _ := ReduceFloat32Err(plusFloat32Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []float32{v1, v0, v3, v4, v5}
	_, err := ReduceFloat32Err(plusFloat32Err, list2)
	if err == nil {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceFloat32Err(plusFloat32Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{v1, v2}
	expected = 3
	actual, _ = ReduceFloat32Err(plusFloat32Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{v1}
	expected = 1
	actual, _ = ReduceFloat32Err(plusFloat32Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{}
	expected = 0
	actual, _ = ReduceFloat32Err(plusFloat32Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat32Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusFloat32Err(num1, num2 float32) (float32, error) {
	if num1 == 0 || num2 == 0 {
		return float32(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}

func TestReduceFloat64Err(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	var v0 float64 = 0

	list := []float64{v1, v2, v3, v4, v5}
	var expected float64 = 15
	actual, _ := ReduceFloat64Err(plusFloat64Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []float64{v1, v0, v3, v4, v5}
	_, err := ReduceFloat64Err(plusFloat64Err, list2)
	if err == nil {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceFloat64Err(plusFloat64Err, list, 3)
	if actual != expected {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{v1, v2}
	expected = 3
	actual, _ = ReduceFloat64Err(plusFloat64Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{v1}
	expected = 1
	actual, _ = ReduceFloat64Err(plusFloat64Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{}
	expected = 0
	actual, _ = ReduceFloat64Err(plusFloat64Err, list)
	if actual != expected {
		t.Errorf("ReduceFloat64Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusFloat64Err(num1, num2 float64) (float64, error) {
	if num1 == 0 || num2 == 0 {
		return float64(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}
