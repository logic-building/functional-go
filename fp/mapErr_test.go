package fp

import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntErr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v5 int = 5
	var v6 int = 6
	var v7 int = 7
	var v8 int = 8

	// Test: add 5 to each item in the list
	expectedSumList := []int{v6, v7, v8}
	partialAddInt := func(num int) (int, error) {
		r, err := addIntErr(v5, num)
		if err != nil {
			return int(0), err
		}
		return r, nil
	}
	sumList, _ := MapIntErr(partialAddInt, []int{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapIntErr failed")
	}

	r, _ := MapIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int = 0
	partialAddInt2 := func(num int) (int, error) {
		r, err := addIntErr(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapIntErr(partialAddInt2, []int{v0})
	if err == nil {
		t.Errorf("MapIntErr failed")
	}
}

func addIntErr(num1, num2 int) (int, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapInt64Err(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v5 int64 = 5
	var v6 int64 = 6
	var v7 int64 = 7
	var v8 int64 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []int64{v6, v7, v8}
	partialAddInt64 := func(num int64) (int64, error) {
		r, err := addInt64Err(v5, num)
		if err != nil {
			return int64(0), err
		}
		return r, nil
	}
	sumList, _ := MapInt64Err(partialAddInt64, []int64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt64Err failed")
	}

	r, _ := MapInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int64 = 0
	partialAddInt642 := func(num int64) (int64, error) {
		r, err := addInt64Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapInt64Err(partialAddInt642, []int64{v0})
	if err == nil {
		t.Errorf("MapInt64Err failed")
	}
}

func addInt64Err(num1, num2 int64) (int64, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapInt32Err(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v5 int32 = 5
	var v6 int32 = 6
	var v7 int32 = 7
	var v8 int32 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []int32{v6, v7, v8}
	partialAddInt32 := func(num int32) (int32, error) {
		r, err := addInt32Err(v5, num)
		if err != nil {
			return int32(0), err
		}
		return r, nil
	}
	sumList, _ := MapInt32Err(partialAddInt32, []int32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt32Err failed")
	}

	r, _ := MapInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int32 = 0
	partialAddInt322 := func(num int32) (int32, error) {
		r, err := addInt32Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapInt32Err(partialAddInt322, []int32{v0})
	if err == nil {
		t.Errorf("MapInt32Err failed")
	}
}

func addInt32Err(num1, num2 int32) (int32, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapInt16Err(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v5 int16 = 5
	var v6 int16 = 6
	var v7 int16 = 7
	var v8 int16 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []int16{v6, v7, v8}
	partialAddInt16 := func(num int16) (int16, error) {
		r, err := addInt16Err(v5, num)
		if err != nil {
			return int16(0), err
		}
		return r, nil
	}
	sumList, _ := MapInt16Err(partialAddInt16, []int16{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt16Err failed")
	}

	r, _ := MapInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int16 = 0
	partialAddInt162 := func(num int16) (int16, error) {
		r, err := addInt16Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapInt16Err(partialAddInt162, []int16{v0})
	if err == nil {
		t.Errorf("MapInt16Err failed")
	}
}

func addInt16Err(num1, num2 int16) (int16, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapInt8Err(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v5 int8 = 5
	var v6 int8 = 6
	var v7 int8 = 7
	var v8 int8 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []int8{v6, v7, v8}
	partialAddInt8 := func(num int8) (int8, error) {
		r, err := addInt8Err(v5, num)
		if err != nil {
			return int8(0), err
		}
		return r, nil
	}
	sumList, _ := MapInt8Err(partialAddInt8, []int8{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt8Err failed")
	}

	r, _ := MapInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int8 = 0
	partialAddInt82 := func(num int8) (int8, error) {
		r, err := addInt8Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapInt8Err(partialAddInt82, []int8{v0})
	if err == nil {
		t.Errorf("MapInt8Err failed")
	}
}

func addInt8Err(num1, num2 int8) (int8, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapUintErr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v5 uint = 5
	var v6 uint = 6
	var v7 uint = 7
	var v8 uint = 8

	// Test: add 5 to each item in the list
	expectedSumList := []uint{v6, v7, v8}
	partialAddUint := func(num uint) (uint, error) {
		r, err := addUintErr(v5, num)
		if err != nil {
			return uint(0), err
		}
		return r, nil
	}
	sumList, _ := MapUintErr(partialAddUint, []uint{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUintErr failed")
	}

	r, _ := MapUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint = 0
	partialAddUint2 := func(num uint) (uint, error) {
		r, err := addUintErr(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapUintErr(partialAddUint2, []uint{v0})
	if err == nil {
		t.Errorf("MapUintErr failed")
	}
}

func addUintErr(num1, num2 uint) (uint, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapUint64Err(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v5 uint64 = 5
	var v6 uint64 = 6
	var v7 uint64 = 7
	var v8 uint64 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []uint64{v6, v7, v8}
	partialAddUint64 := func(num uint64) (uint64, error) {
		r, err := addUint64Err(v5, num)
		if err != nil {
			return uint64(0), err
		}
		return r, nil
	}
	sumList, _ := MapUint64Err(partialAddUint64, []uint64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUint64Err failed")
	}

	r, _ := MapUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint64 = 0
	partialAddUint642 := func(num uint64) (uint64, error) {
		r, err := addUint64Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapUint64Err(partialAddUint642, []uint64{v0})
	if err == nil {
		t.Errorf("MapUint64Err failed")
	}
}

func addUint64Err(num1, num2 uint64) (uint64, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapUint32Err(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v5 uint32 = 5
	var v6 uint32 = 6
	var v7 uint32 = 7
	var v8 uint32 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []uint32{v6, v7, v8}
	partialAddUint32 := func(num uint32) (uint32, error) {
		r, err := addUint32Err(v5, num)
		if err != nil {
			return uint32(0), err
		}
		return r, nil
	}
	sumList, _ := MapUint32Err(partialAddUint32, []uint32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUint32Err failed")
	}

	r, _ := MapUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint32 = 0
	partialAddUint322 := func(num uint32) (uint32, error) {
		r, err := addUint32Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapUint32Err(partialAddUint322, []uint32{v0})
	if err == nil {
		t.Errorf("MapUint32Err failed")
	}
}

func addUint32Err(num1, num2 uint32) (uint32, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapUint16Err(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v5 uint16 = 5
	var v6 uint16 = 6
	var v7 uint16 = 7
	var v8 uint16 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []uint16{v6, v7, v8}
	partialAddUint16 := func(num uint16) (uint16, error) {
		r, err := addUint16Err(v5, num)
		if err != nil {
			return uint16(0), err
		}
		return r, nil
	}
	sumList, _ := MapUint16Err(partialAddUint16, []uint16{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUint16Err failed")
	}

	r, _ := MapUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint16 = 0
	partialAddUint162 := func(num uint16) (uint16, error) {
		r, err := addUint16Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapUint16Err(partialAddUint162, []uint16{v0})
	if err == nil {
		t.Errorf("MapUint16Err failed")
	}
}

func addUint16Err(num1, num2 uint16) (uint16, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapUint8Err(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v5 uint8 = 5
	var v6 uint8 = 6
	var v7 uint8 = 7
	var v8 uint8 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []uint8{v6, v7, v8}
	partialAddUint8 := func(num uint8) (uint8, error) {
		r, err := addUint8Err(v5, num)
		if err != nil {
			return uint8(0), err
		}
		return r, nil
	}
	sumList, _ := MapUint8Err(partialAddUint8, []uint8{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUint8Err failed")
	}

	r, _ := MapUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint8 = 0
	partialAddUint82 := func(num uint8) (uint8, error) {
		r, err := addUint8Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapUint8Err(partialAddUint82, []uint8{v0})
	if err == nil {
		t.Errorf("MapUint8Err failed")
	}
}

func addUint8Err(num1, num2 uint8) (uint8, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"

	// Test: add 5 to each item in the list
	expectedSumList := []string{v6, v7, v8}
	partialAddStr := func(num string) (string, error) {
		r, err := addStrErr(v5, num)
		if err != nil {
			return "0", err
		}
		return r, nil
	}
	sumList, _ := MapStrErr(partialAddStr, []string{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapStrErr failed")
	}

	r, _ := MapStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 string = "0"
	partialAddStr2 := func(num string) (string, error) {
		r, err := addStrErr(v0, num)
		if err != nil {
			return "0", err
		}
		return r, nil
	}
	_, err := MapStrErr(partialAddStr2, []string{v0})
	if err == nil {
		t.Errorf("MapStrErr failed")
	}
}

func addStrErr(num1, num2 string) (string, error) {
	if num1 == "0" {
		return "", errors.New("0 is not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapBoolErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []bool{vf}

	newList, _ := MapBoolErr(inverseBoolErr, []bool{vt})
	if newList[0] != expectedSumList[0] {
		t.Errorf("MapBoolErr failed")
	}

	_, err1 := MapBoolErr(inverseBoolErr, []bool{vf})
	if err1 == nil {
		t.Errorf("MapBoolErr failed")
	}

	r, _ := MapBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolErr failed.")
	}
}

func inverseBoolErr(num1 bool) (bool, error) {
	if num1 == false {
		return false, errors.New("False is not allowed")
	}
	vt := true
	if num1 == true {
		v := false
		return v, nil
	}
	return vt, nil
}

func TestMapFloat32Err(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v5 float32 = 5
	var v6 float32 = 6
	var v7 float32 = 7
	var v8 float32 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []float32{v6, v7, v8}
	partialAddFloat32 := func(num float32) (float32, error) {
		r, err := addFloat32Err(v5, num)
		if err != nil {
			return float32(0), err
		}
		return r, nil
	}
	sumList, _ := MapFloat32Err(partialAddFloat32, []float32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat32Err failed")
	}

	r, _ := MapFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 float32 = 0
	partialAddFloat322 := func(num float32) (float32, error) {
		r, err := addFloat32Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapFloat32Err(partialAddFloat322, []float32{v0})
	if err == nil {
		t.Errorf("MapFloat32Err failed")
	}
}

func addFloat32Err(num1, num2 float32) (float32, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}

func TestMapFloat64Err(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v5 float64 = 5
	var v6 float64 = 6
	var v7 float64 = 7
	var v8 float64 = 8

	// Test: add 5 to each item in the list
	expectedSumList := []float64{v6, v7, v8}
	partialAddFloat64 := func(num float64) (float64, error) {
		r, err := addFloat64Err(v5, num)
		if err != nil {
			return float64(0), err
		}
		return r, nil
	}
	sumList, _ := MapFloat64Err(partialAddFloat64, []float64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat64Err failed")
	}

	r, _ := MapFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 float64 = 0
	partialAddFloat642 := func(num float64) (float64, error) {
		r, err := addFloat64Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := MapFloat64Err(partialAddFloat642, []float64{v0})
	if err == nil {
		t.Errorf("MapFloat64Err failed")
	}
}

func addFloat64Err(num1, num2 float64) (float64, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}
