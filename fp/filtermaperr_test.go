package fp

import (
	"errors"
	"testing"
)

func TestFilterMapIntErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int = 1
	var v4 int = 4
	var v8 int = 8
	var v16 int = 16
	var v0 int = 0
	var v2 int = 2
	
	expectedFilteredList := []int{v8, v16}
	filteredList, _ := FilterMapIntErr(isPositiveIntErr, multiplyBy2IntErr, []int{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapIntErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapIntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntErr failed.")
	}

	_, err := FilterMapIntErr(isPositiveIntErr, multiplyBy2IntErr, []int{v0, v1})
	if err == nil {
		t.Errorf("FilterMapIntErr failed.")
	}

	_, err = FilterMapIntErr(isPositiveIntErr, multiplyBy2IntErr, []int{v1, v2})
	if err == nil {
		t.Errorf("FilterMapIntErr failed.")
	}
}

func isPositiveIntErr(num int) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2IntErr(num int) (int, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapInt64Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int64 = 1
	var v4 int64 = 4
	var v8 int64 = 8
	var v16 int64 = 16
	var v0 int64 = 0
	var v2 int64 = 2
	
	expectedFilteredList := []int64{v8, v16}
	filteredList, _ := FilterMapInt64Err(isPositiveInt64Err, multiplyBy2Int64Err, []int64{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt64Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Err failed.")
	}

	_, err := FilterMapInt64Err(isPositiveInt64Err, multiplyBy2Int64Err, []int64{v0, v1})
	if err == nil {
		t.Errorf("FilterMapInt64Err failed.")
	}

	_, err = FilterMapInt64Err(isPositiveInt64Err, multiplyBy2Int64Err, []int64{v1, v2})
	if err == nil {
		t.Errorf("FilterMapInt64Err failed.")
	}
}

func isPositiveInt64Err(num int64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Int64Err(num int64) (int64, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapInt32Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int32 = 1
	var v4 int32 = 4
	var v8 int32 = 8
	var v16 int32 = 16
	var v0 int32 = 0
	var v2 int32 = 2
	
	expectedFilteredList := []int32{v8, v16}
	filteredList, _ := FilterMapInt32Err(isPositiveInt32Err, multiplyBy2Int32Err, []int32{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt32Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Err failed.")
	}

	_, err := FilterMapInt32Err(isPositiveInt32Err, multiplyBy2Int32Err, []int32{v0, v1})
	if err == nil {
		t.Errorf("FilterMapInt32Err failed.")
	}

	_, err = FilterMapInt32Err(isPositiveInt32Err, multiplyBy2Int32Err, []int32{v1, v2})
	if err == nil {
		t.Errorf("FilterMapInt32Err failed.")
	}
}

func isPositiveInt32Err(num int32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Int32Err(num int32) (int32, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapInt16Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int16 = 1
	var v4 int16 = 4
	var v8 int16 = 8
	var v16 int16 = 16
	var v0 int16 = 0
	var v2 int16 = 2
	
	expectedFilteredList := []int16{v8, v16}
	filteredList, _ := FilterMapInt16Err(isPositiveInt16Err, multiplyBy2Int16Err, []int16{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt16Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Err failed.")
	}

	_, err := FilterMapInt16Err(isPositiveInt16Err, multiplyBy2Int16Err, []int16{v0, v1})
	if err == nil {
		t.Errorf("FilterMapInt16Err failed.")
	}

	_, err = FilterMapInt16Err(isPositiveInt16Err, multiplyBy2Int16Err, []int16{v1, v2})
	if err == nil {
		t.Errorf("FilterMapInt16Err failed.")
	}
}

func isPositiveInt16Err(num int16) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Int16Err(num int16) (int16, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapInt8Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int8 = 1
	var v4 int8 = 4
	var v8 int8 = 8
	var v16 int8 = 16
	var v0 int8 = 0
	var v2 int8 = 2
	
	expectedFilteredList := []int8{v8, v16}
	filteredList, _ := FilterMapInt8Err(isPositiveInt8Err, multiplyBy2Int8Err, []int8{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt8Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Err failed.")
	}

	_, err := FilterMapInt8Err(isPositiveInt8Err, multiplyBy2Int8Err, []int8{v0, v1})
	if err == nil {
		t.Errorf("FilterMapInt8Err failed.")
	}

	_, err = FilterMapInt8Err(isPositiveInt8Err, multiplyBy2Int8Err, []int8{v1, v2})
	if err == nil {
		t.Errorf("FilterMapInt8Err failed.")
	}
}

func isPositiveInt8Err(num int8) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Int8Err(num int8) (int8, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapUintErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint = 1
	var v4 uint = 4
	var v8 uint = 8
	var v16 uint = 16
	var v0 uint = 0
	var v2 uint = 2
	
	expectedFilteredList := []uint{v8, v16}
	filteredList, _ := FilterMapUintErr(isPositiveUintErr, multiplyBy2UintErr, []uint{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUintErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintErr failed.")
	}

	_, err := FilterMapUintErr(isPositiveUintErr, multiplyBy2UintErr, []uint{v0, v1})
	if err == nil {
		t.Errorf("FilterMapUintErr failed.")
	}

	_, err = FilterMapUintErr(isPositiveUintErr, multiplyBy2UintErr, []uint{v1, v2})
	if err == nil {
		t.Errorf("FilterMapUintErr failed.")
	}
}

func isPositiveUintErr(num uint) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2UintErr(num uint) (uint, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapUint64Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint64 = 1
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v16 uint64 = 16
	var v0 uint64 = 0
	var v2 uint64 = 2
	
	expectedFilteredList := []uint64{v8, v16}
	filteredList, _ := FilterMapUint64Err(isPositiveUint64Err, multiplyBy2Uint64Err, []uint64{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint64Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Err failed.")
	}

	_, err := FilterMapUint64Err(isPositiveUint64Err, multiplyBy2Uint64Err, []uint64{v0, v1})
	if err == nil {
		t.Errorf("FilterMapUint64Err failed.")
	}

	_, err = FilterMapUint64Err(isPositiveUint64Err, multiplyBy2Uint64Err, []uint64{v1, v2})
	if err == nil {
		t.Errorf("FilterMapUint64Err failed.")
	}
}

func isPositiveUint64Err(num uint64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Uint64Err(num uint64) (uint64, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapUint32Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint32 = 1
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v16 uint32 = 16
	var v0 uint32 = 0
	var v2 uint32 = 2
	
	expectedFilteredList := []uint32{v8, v16}
	filteredList, _ := FilterMapUint32Err(isPositiveUint32Err, multiplyBy2Uint32Err, []uint32{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint32Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Err failed.")
	}

	_, err := FilterMapUint32Err(isPositiveUint32Err, multiplyBy2Uint32Err, []uint32{v0, v1})
	if err == nil {
		t.Errorf("FilterMapUint32Err failed.")
	}

	_, err = FilterMapUint32Err(isPositiveUint32Err, multiplyBy2Uint32Err, []uint32{v1, v2})
	if err == nil {
		t.Errorf("FilterMapUint32Err failed.")
	}
}

func isPositiveUint32Err(num uint32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Uint32Err(num uint32) (uint32, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapUint16Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint16 = 1
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v16 uint16 = 16
	var v0 uint16 = 0
	var v2 uint16 = 2
	
	expectedFilteredList := []uint16{v8, v16}
	filteredList, _ := FilterMapUint16Err(isPositiveUint16Err, multiplyBy2Uint16Err, []uint16{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint16Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Err failed.")
	}

	_, err := FilterMapUint16Err(isPositiveUint16Err, multiplyBy2Uint16Err, []uint16{v0, v1})
	if err == nil {
		t.Errorf("FilterMapUint16Err failed.")
	}

	_, err = FilterMapUint16Err(isPositiveUint16Err, multiplyBy2Uint16Err, []uint16{v1, v2})
	if err == nil {
		t.Errorf("FilterMapUint16Err failed.")
	}
}

func isPositiveUint16Err(num uint16) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Uint16Err(num uint16) (uint16, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapUint8Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint8 = 1
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v16 uint8 = 16
	var v0 uint8 = 0
	var v2 uint8 = 2
	
	expectedFilteredList := []uint8{v8, v16}
	filteredList, _ := FilterMapUint8Err(isPositiveUint8Err, multiplyBy2Uint8Err, []uint8{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint8Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Err failed.")
	}

	_, err := FilterMapUint8Err(isPositiveUint8Err, multiplyBy2Uint8Err, []uint8{v0, v1})
	if err == nil {
		t.Errorf("FilterMapUint8Err failed.")
	}

	_, err = FilterMapUint8Err(isPositiveUint8Err, multiplyBy2Uint8Err, []uint8{v1, v2})
	if err == nil {
		t.Errorf("FilterMapUint8Err failed.")
	}
}

func isPositiveUint8Err(num uint8) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Uint8Err(num uint8) (uint8, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapStrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v16 string = "16"
	var v0 string = "0"
	var v2 string = "2"
	
	expectedFilteredList := []string{v8, v16}
	filteredList, _ := FilterMapStrErr(isPositiveStrErr, multiplyBy2StrErr, []string{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapStrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapStrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrErr failed.")
	}

	_, err := FilterMapStrErr(isPositiveStrErr, multiplyBy2StrErr, []string{v0, v1})
	if err == nil {
		t.Errorf("FilterMapStrErr failed.")
	}

	_, err = FilterMapStrErr(isPositiveStrErr, multiplyBy2StrErr, []string{v1, v2})
	if err == nil {
		t.Errorf("FilterMapStrErr failed.")
	}
}

func isPositiveStrErr(num string) (bool, error) {
	if num == "0" {
		return false, errors.New("Zero is invalid")
	}
	if num == "2" || num == "3" || num == "4" || num == "5" || num == "6" || num == "7" || num == "8" {
		return true, nil
	}
	return false, nil
}

func multiplyBy2StrErr(num string) (string, error) {
	result := "100"
	if num == "2" {
		return "0", errors.New("2 is invalid number")
	}
	if num == "3" {
		result = "6"
	} 

	if num == "4" {
		result = "8"
	} 

	if num == "5" {
		result = "10"
	} 

	if num == "6" {
		result = "12"
	} 

	if num == "7" {
		result = "14"
	} 

	if num == "8" {
		result = "16"
	} 
	
	return result, nil
}


func TestFilterMapFloat32Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float32 = 1
	var v4 float32 = 4
	var v8 float32 = 8
	var v16 float32 = 16
	var v0 float32 = 0
	var v2 float32 = 2
	
	expectedFilteredList := []float32{v8, v16}
	filteredList, _ := FilterMapFloat32Err(isPositiveFloat32Err, multiplyBy2Float32Err, []float32{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapFloat32Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapFloat32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Err failed.")
	}

	_, err := FilterMapFloat32Err(isPositiveFloat32Err, multiplyBy2Float32Err, []float32{v0, v1})
	if err == nil {
		t.Errorf("FilterMapFloat32Err failed.")
	}

	_, err = FilterMapFloat32Err(isPositiveFloat32Err, multiplyBy2Float32Err, []float32{v1, v2})
	if err == nil {
		t.Errorf("FilterMapFloat32Err failed.")
	}
}

func isPositiveFloat32Err(num float32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Float32Err(num float32) (float32, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}


func TestFilterMapFloat64Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float64 = 1
	var v4 float64 = 4
	var v8 float64 = 8
	var v16 float64 = 16
	var v0 float64 = 0
	var v2 float64 = 2
	
	expectedFilteredList := []float64{v8, v16}
	filteredList, _ := FilterMapFloat64Err(isPositiveFloat64Err, multiplyBy2Float64Err, []float64{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapFloat64Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapFloat64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Err failed.")
	}

	_, err := FilterMapFloat64Err(isPositiveFloat64Err, multiplyBy2Float64Err, []float64{v0, v1})
	if err == nil {
		t.Errorf("FilterMapFloat64Err failed.")
	}

	_, err = FilterMapFloat64Err(isPositiveFloat64Err, multiplyBy2Float64Err, []float64{v1, v2})
	if err == nil {
		t.Errorf("FilterMapFloat64Err failed.")
	}
}

func isPositiveFloat64Err(num float64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2Float64Err(num float64) (float64, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}

