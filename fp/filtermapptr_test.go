package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntPtr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int = 1
	var v4 int = 4
	var v8 int = 8
	var v0 int
	var v2 int = 2

	expectedFilteredList := []*int{&v2, &v4, &v8}
	filteredList := FilterMapIntPtr(isPositiveIntPtr, multiplyBy2IntPtr, []*int{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapIntPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapIntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveIntPtr(num *int) bool {
	return *num > 0
}
func multiplyBy2IntPtr(num *int) *int {
	result := *num * 2
	return &result
}

func TestFilterMapInt64Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int64 = 1
	var v4 int64 = 4
	var v8 int64 = 8
	var v0 int64
	var v2 int64 = 2

	expectedFilteredList := []*int64{&v2, &v4, &v8}
	filteredList := FilterMapInt64Ptr(isPositiveInt64Ptr, multiplyBy2Int64Ptr, []*int64{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapInt64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveInt64Ptr(num *int64) bool {
	return *num > 0
}
func multiplyBy2Int64Ptr(num *int64) *int64 {
	result := *num * 2
	return &result
}

func TestFilterMapInt32Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int32 = 1
	var v4 int32 = 4
	var v8 int32 = 8
	var v0 int32
	var v2 int32 = 2

	expectedFilteredList := []*int32{&v2, &v4, &v8}
	filteredList := FilterMapInt32Ptr(isPositiveInt32Ptr, multiplyBy2Int32Ptr, []*int32{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapInt32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveInt32Ptr(num *int32) bool {
	return *num > 0
}
func multiplyBy2Int32Ptr(num *int32) *int32 {
	result := *num * 2
	return &result
}

func TestFilterMapInt16Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int16 = 1
	var v4 int16 = 4
	var v8 int16 = 8
	var v0 int16
	var v2 int16 = 2

	expectedFilteredList := []*int16{&v2, &v4, &v8}
	filteredList := FilterMapInt16Ptr(isPositiveInt16Ptr, multiplyBy2Int16Ptr, []*int16{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapInt16Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveInt16Ptr(num *int16) bool {
	return *num > 0
}
func multiplyBy2Int16Ptr(num *int16) *int16 {
	result := *num * 2
	return &result
}

func TestFilterMapInt8Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 int8 = 1
	var v4 int8 = 4
	var v8 int8 = 8
	var v0 int8
	var v2 int8 = 2

	expectedFilteredList := []*int8{&v2, &v4, &v8}
	filteredList := FilterMapInt8Ptr(isPositiveInt8Ptr, multiplyBy2Int8Ptr, []*int8{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapInt8Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveInt8Ptr(num *int8) bool {
	return *num > 0
}
func multiplyBy2Int8Ptr(num *int8) *int8 {
	result := *num * 2
	return &result
}

func TestFilterMapUintPtr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint = 1
	var v4 uint = 4
	var v8 uint = 8
	var v0 uint
	var v2 uint = 2

	expectedFilteredList := []*uint{&v2, &v4, &v8}
	filteredList := FilterMapUintPtr(isPositiveUintPtr, multiplyBy2UintPtr, []*uint{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapUintPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveUintPtr(num *uint) bool {
	return *num > 0
}
func multiplyBy2UintPtr(num *uint) *uint {
	result := *num * 2
	return &result
}

func TestFilterMapUint64Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint64 = 1
	var v4 uint64 = 4
	var v8 uint64 = 8
	var v0 uint64
	var v2 uint64 = 2

	expectedFilteredList := []*uint64{&v2, &v4, &v8}
	filteredList := FilterMapUint64Ptr(isPositiveUint64Ptr, multiplyBy2Uint64Ptr, []*uint64{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapUint64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveUint64Ptr(num *uint64) bool {
	return *num > 0
}
func multiplyBy2Uint64Ptr(num *uint64) *uint64 {
	result := *num * 2
	return &result
}

func TestFilterMapUint32Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint32 = 1
	var v4 uint32 = 4
	var v8 uint32 = 8
	var v0 uint32
	var v2 uint32 = 2

	expectedFilteredList := []*uint32{&v2, &v4, &v8}
	filteredList := FilterMapUint32Ptr(isPositiveUint32Ptr, multiplyBy2Uint32Ptr, []*uint32{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapUint32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveUint32Ptr(num *uint32) bool {
	return *num > 0
}
func multiplyBy2Uint32Ptr(num *uint32) *uint32 {
	result := *num * 2
	return &result
}

func TestFilterMapUint16Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint16 = 1
	var v4 uint16 = 4
	var v8 uint16 = 8
	var v0 uint16
	var v2 uint16 = 2

	expectedFilteredList := []*uint16{&v2, &v4, &v8}
	filteredList := FilterMapUint16Ptr(isPositiveUint16Ptr, multiplyBy2Uint16Ptr, []*uint16{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapUint16Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveUint16Ptr(num *uint16) bool {
	return *num > 0
}
func multiplyBy2Uint16Ptr(num *uint16) *uint16 {
	result := *num * 2
	return &result
}

func TestFilterMapUint8Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 uint8 = 1
	var v4 uint8 = 4
	var v8 uint8 = 8
	var v0 uint8
	var v2 uint8 = 2

	expectedFilteredList := []*uint8{&v2, &v4, &v8}
	filteredList := FilterMapUint8Ptr(isPositiveUint8Ptr, multiplyBy2Uint8Ptr, []*uint8{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapUint8Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveUint8Ptr(num *uint8) bool {
	return *num > 0
}
func multiplyBy2Uint8Ptr(num *uint8) *uint8 {
	result := *num * 2
	return &result
}

func TestFilterMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v4 string = "4"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v1, &v0, &v2, &v4}
	filteredList := FilterMapStrPtr(noFilter, concatA, []*string{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0]+"A" || *filteredList[1] != *expectedFilteredList[1]+"A" {
		t.Errorf("FilterMapStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}
}

func noFilter(num *string) bool {
	return true
}
func concatA(num *string) *string {
	result := *num + "A"
	return &result
}

func TestFilterMapBoolPtr(t *testing.T) {
	var vt1 bool = true
	var vt2 bool = true
	var vf1 bool = false
	var vf2 bool = false

	expectedFilteredList := []*bool{&vt1, &vt2}
	filteredList := FilterMapBoolPtr(isTrueBoolPtr, returnSameBoolPtr, []*bool{&vt1, &vt2, &vf1, &vf2})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapBoolPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapBoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isTrueBoolPtr(num *bool) bool {
	return *num == true
}
func returnSameBoolPtr(num *bool) *bool {
	return num
}

func TestFilterMapFloat32Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float32 = 1
	var v4 float32 = 4
	var v8 float32 = 8
	var v0 float32
	var v2 float32 = 2

	expectedFilteredList := []*float32{&v2, &v4, &v8}
	filteredList := FilterMapFloat32Ptr(isPositiveFloat32Ptr, multiplyBy2Float32Ptr, []*float32{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapFloat32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapFloat32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveFloat32Ptr(num *float32) bool {
	return *num > 0
}
func multiplyBy2Float32Ptr(num *float32) *float32 {
	result := *num * 2
	return &result
}

func TestFilterMapFloat64Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 float64 = 1
	var v4 float64 = 4
	var v8 float64 = 8
	var v0 float64
	var v2 float64 = 2

	expectedFilteredList := []*float64{&v2, &v4, &v8}
	filteredList := FilterMapFloat64Ptr(isPositiveFloat64Ptr, multiplyBy2Float64Ptr, []*float64{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapFloat64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapFloat64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveFloat64Ptr(num *float64) bool {
	return *num > 0
}
func multiplyBy2Float64Ptr(num *float64) *float64 {
	result := *num * 2
	return &result
}
