package fp

import (
	"errors"
	"testing"
)

func TestFilterMapIntPtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int = 1
	var v4 int = 4
	var v8 int = 8
	var v16 int = 16
	var v0 int
	var v2 int = 2

	expectedFilteredList := []*int{&v8, &v16}
	filteredList, _ := FilterMapIntPtrErr(isPositiveIntPtrErr, multiplyBy2IntPtrErr, []*int{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapIntPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapIntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntPtr failed.")
	}

	_, err := FilterMapIntPtrErr(isPositiveIntPtrErr, multiplyBy2IntPtrErr, []*int{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapIntPtrErr failed.")
	}

	_, err = FilterMapIntPtrErr(isPositiveIntPtrErr, multiplyBy2IntPtrErr, []*int{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapIntPtrErr failed.")
	}
}

func isPositiveIntPtrErr(num *int) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2IntPtrErr(num *int) (*int, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapInt64PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int64 = 1
	var v4 int64 = 4
	var v8 int64 = 8
	var v16 int64 = 16
	var v0 int64
	var v2 int64 = 2

	expectedFilteredList := []*int64{&v8, &v16}
	filteredList, _ := FilterMapInt64PtrErr(isPositiveInt64PtrErr, multiplyBy2Int64PtrErr, []*int64{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapInt64PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Ptr failed.")
	}

	_, err := FilterMapInt64PtrErr(isPositiveInt64PtrErr, multiplyBy2Int64PtrErr, []*int64{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapInt64PtrErr failed.")
	}

	_, err = FilterMapInt64PtrErr(isPositiveInt64PtrErr, multiplyBy2Int64PtrErr, []*int64{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapInt64PtrErr failed.")
	}
}

func isPositiveInt64PtrErr(num *int64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Int64PtrErr(num *int64) (*int64, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapInt32PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int32 = 1
	var v4 int32 = 4
	var v8 int32 = 8
	var v16 int32 = 16
	var v0 int32
	var v2 int32 = 2

	expectedFilteredList := []*int32{&v8, &v16}
	filteredList, _ := FilterMapInt32PtrErr(isPositiveInt32PtrErr, multiplyBy2Int32PtrErr, []*int32{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapInt32PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Ptr failed.")
	}

	_, err := FilterMapInt32PtrErr(isPositiveInt32PtrErr, multiplyBy2Int32PtrErr, []*int32{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapInt32PtrErr failed.")
	}

	_, err = FilterMapInt32PtrErr(isPositiveInt32PtrErr, multiplyBy2Int32PtrErr, []*int32{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapInt32PtrErr failed.")
	}
}

func isPositiveInt32PtrErr(num *int32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Int32PtrErr(num *int32) (*int32, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapInt16PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int16 = 1
	var v4 int16 = 4
	var v8 int16 = 8
	var v16 int16 = 16
	var v0 int16
	var v2 int16 = 2

	expectedFilteredList := []*int16{&v8, &v16}
	filteredList, _ := FilterMapInt16PtrErr(isPositiveInt16PtrErr, multiplyBy2Int16PtrErr, []*int16{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapInt16PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Ptr failed.")
	}

	_, err := FilterMapInt16PtrErr(isPositiveInt16PtrErr, multiplyBy2Int16PtrErr, []*int16{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapInt16PtrErr failed.")
	}

	_, err = FilterMapInt16PtrErr(isPositiveInt16PtrErr, multiplyBy2Int16PtrErr, []*int16{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapInt16PtrErr failed.")
	}
}

func isPositiveInt16PtrErr(num *int16) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Int16PtrErr(num *int16) (*int16, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapInt8PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int8 = 1
	var v4 int8 = 4
	var v8 int8 = 8
	var v16 int8 = 16
	var v0 int8
	var v2 int8 = 2

	expectedFilteredList := []*int8{&v8, &v16}
	filteredList, _ := FilterMapInt8PtrErr(isPositiveInt8PtrErr, multiplyBy2Int8PtrErr, []*int8{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapInt8PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapInt8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Ptr failed.")
	}

	_, err := FilterMapInt8PtrErr(isPositiveInt8PtrErr, multiplyBy2Int8PtrErr, []*int8{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapInt8PtrErr failed.")
	}

	_, err = FilterMapInt8PtrErr(isPositiveInt8PtrErr, multiplyBy2Int8PtrErr, []*int8{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapInt8PtrErr failed.")
	}
}

func isPositiveInt8PtrErr(num *int8) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Int8PtrErr(num *int8) (*int8, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapUintPtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint = 1
	var v4 uint = 4
	var v8 uint = 8
	var v16 uint = 16
	var v0 uint
	var v2 uint = 2

	expectedFilteredList := []*uint{&v8, &v16}
	filteredList, _ := FilterMapUintPtrErr(isPositiveUintPtrErr, multiplyBy2UintPtrErr, []*uint{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapUintPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintPtr failed.")
	}

	_, err := FilterMapUintPtrErr(isPositiveUintPtrErr, multiplyBy2UintPtrErr, []*uint{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapUintPtrErr failed.")
	}

	_, err = FilterMapUintPtrErr(isPositiveUintPtrErr, multiplyBy2UintPtrErr, []*uint{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapUintPtrErr failed.")
	}
}

func isPositiveUintPtrErr(num *uint) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2UintPtrErr(num *uint) (*uint, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapUint64PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint64 = 1
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v16 uint64 = 16
	var v0 uint64
	var v2 uint64 = 2

	expectedFilteredList := []*uint64{&v8, &v16}
	filteredList, _ := FilterMapUint64PtrErr(isPositiveUint64PtrErr, multiplyBy2Uint64PtrErr, []*uint64{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapUint64PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Ptr failed.")
	}

	_, err := FilterMapUint64PtrErr(isPositiveUint64PtrErr, multiplyBy2Uint64PtrErr, []*uint64{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapUint64PtrErr failed.")
	}

	_, err = FilterMapUint64PtrErr(isPositiveUint64PtrErr, multiplyBy2Uint64PtrErr, []*uint64{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapUint64PtrErr failed.")
	}
}

func isPositiveUint64PtrErr(num *uint64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Uint64PtrErr(num *uint64) (*uint64, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapUint32PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint32 = 1
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v16 uint32 = 16
	var v0 uint32
	var v2 uint32 = 2

	expectedFilteredList := []*uint32{&v8, &v16}
	filteredList, _ := FilterMapUint32PtrErr(isPositiveUint32PtrErr, multiplyBy2Uint32PtrErr, []*uint32{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapUint32PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Ptr failed.")
	}

	_, err := FilterMapUint32PtrErr(isPositiveUint32PtrErr, multiplyBy2Uint32PtrErr, []*uint32{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapUint32PtrErr failed.")
	}

	_, err = FilterMapUint32PtrErr(isPositiveUint32PtrErr, multiplyBy2Uint32PtrErr, []*uint32{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapUint32PtrErr failed.")
	}
}

func isPositiveUint32PtrErr(num *uint32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Uint32PtrErr(num *uint32) (*uint32, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapUint16PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint16 = 1
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v16 uint16 = 16
	var v0 uint16
	var v2 uint16 = 2

	expectedFilteredList := []*uint16{&v8, &v16}
	filteredList, _ := FilterMapUint16PtrErr(isPositiveUint16PtrErr, multiplyBy2Uint16PtrErr, []*uint16{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapUint16PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Ptr failed.")
	}

	_, err := FilterMapUint16PtrErr(isPositiveUint16PtrErr, multiplyBy2Uint16PtrErr, []*uint16{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapUint16PtrErr failed.")
	}

	_, err = FilterMapUint16PtrErr(isPositiveUint16PtrErr, multiplyBy2Uint16PtrErr, []*uint16{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapUint16PtrErr failed.")
	}
}

func isPositiveUint16PtrErr(num *uint16) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Uint16PtrErr(num *uint16) (*uint16, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapUint8PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint8 = 1
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v16 uint8 = 16
	var v0 uint8
	var v2 uint8 = 2

	expectedFilteredList := []*uint8{&v8, &v16}
	filteredList, _ := FilterMapUint8PtrErr(isPositiveUint8PtrErr, multiplyBy2Uint8PtrErr, []*uint8{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapUint8PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapUint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Ptr failed.")
	}

	_, err := FilterMapUint8PtrErr(isPositiveUint8PtrErr, multiplyBy2Uint8PtrErr, []*uint8{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapUint8PtrErr failed.")
	}

	_, err = FilterMapUint8PtrErr(isPositiveUint8PtrErr, multiplyBy2Uint8PtrErr, []*uint8{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapUint8PtrErr failed.")
	}
}

func isPositiveUint8PtrErr(num *uint8) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Uint8PtrErr(num *uint8) (*uint8, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapStrPtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v16 string = "16"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v8, &v16}
	filteredList, _ := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapStrPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}

	_, err := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}

	_, err = FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}
}

func isPositiveStrPtrErr(num *string) (bool, error) {
	if *num == "0" {
		return false, errors.New("Zero is invalid")
	}
	if *num == "2" || *num == "3" || *num == "4" || *num == "5" || *num == "6" || *num == "7" || *num == "8" {
		return true, nil
	} else {
		return false, nil
	}
}

func multiplyBy2StrPtrErr(num *string) (*string, error) {
	result := "100"
	if *num == "2" {
		return nil, errors.New("2 is invalid number")
	}
	if *num == "3" {
		result = "6"
	}

	if *num == "4" {
		result = "8"
	}

	if *num == "5" {
		result = "10"
	}

	if *num == "6" {
		result = "12"
	}

	if *num == "7" {
		result = "14"
	}

	if *num == "8" {
		result = "16"
	}

	return &result, nil
}

func TestFilterMapFloat32PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float32 = 1
	var v4 float32 = 4
	var v8 float32 = 8
	var v16 float32 = 16
	var v0 float32
	var v2 float32 = 2

	expectedFilteredList := []*float32{&v8, &v16}
	filteredList, _ := FilterMapFloat32PtrErr(isPositiveFloat32PtrErr, multiplyBy2Float32PtrErr, []*float32{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapFloat32PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapFloat32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Ptr failed.")
	}

	_, err := FilterMapFloat32PtrErr(isPositiveFloat32PtrErr, multiplyBy2Float32PtrErr, []*float32{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapFloat32PtrErr failed.")
	}

	_, err = FilterMapFloat32PtrErr(isPositiveFloat32PtrErr, multiplyBy2Float32PtrErr, []*float32{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapFloat32PtrErr failed.")
	}
}

func isPositiveFloat32PtrErr(num *float32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Float32PtrErr(num *float32) (*float32, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}

func TestFilterMapFloat64PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float64 = 1
	var v4 float64 = 4
	var v8 float64 = 8
	var v16 float64 = 16
	var v0 float64
	var v2 float64 = 2

	expectedFilteredList := []*float64{&v8, &v16}
	filteredList, _ := FilterMapFloat64PtrErr(isPositiveFloat64PtrErr, multiplyBy2Float64PtrErr, []*float64{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapFloat64PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapFloat64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Ptr failed.")
	}

	_, err := FilterMapFloat64PtrErr(isPositiveFloat64PtrErr, multiplyBy2Float64PtrErr, []*float64{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapFloat64PtrErr failed.")
	}

	_, err = FilterMapFloat64PtrErr(isPositiveFloat64PtrErr, multiplyBy2Float64PtrErr, []*float64{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapFloat64PtrErr failed.")
	}
}

func isPositiveFloat64PtrErr(num *float64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2Float64PtrErr(num *float64) (*float64, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}
