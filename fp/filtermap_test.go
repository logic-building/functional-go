package fp

import (
	"strings"
	"testing"
)

func TestFilterMapInt(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []int{4, 8}
	filteredList := FilterMapInt(isPositive, multiplyBy2, []int{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMap failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt failed.")
	}
}

func isPositive(num int) bool {
	return num > 0
}
func multiplyBy2(num int) int {
	return num * 2
}

func TestFilterMapInt64(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []int64{4, 8}
	filteredList := FilterMapInt64(isPositiveInt64, multiplyBy2Int64, []int64{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64 failed.")
	}
}

func BenchmarkFilterMapInt64(b *testing.B) {
	b.N = iterations
	list := make([]int64, size)
	var i int64
	for i = 0; i < size; i++ {
		list[i] = i
	}
	for i := 0; i < b.N; i++ {
		FilterMapInt64(isEvenInt64, multiplyBy2Int64, list)
	}
}

func isPositiveInt64(num int64) bool {
	return num > 0
}
func multiplyBy2Int64(num int64) int64 {
	return num * 2
}

func TestFilterMapInt32(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []int32{4, 8}
	filteredList := FilterMapInt32(isPositiveInt32, multiplyBy2Int32, []int32{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32 failed.")
	}
}

func isPositiveInt32(num int32) bool {
	return num > 0
}
func multiplyBy2Int32(num int32) int32 {
	return num * 2
}

func TestFilterMapInt16(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []int16{4, 8}
	filteredList := FilterMapInt16(isPositiveInt16, multiplyBy2Int16, []int16{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16 failed.")
	}
}

func isPositiveInt16(num int16) bool {
	return num > 0
}
func multiplyBy2Int16(num int16) int16 {
	return num * 2
}

func TestFilterMapInt8(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []int8{4, 8}
	filteredList := FilterMapInt8(isPositiveInt8, multiplyBy2Int8, []int8{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapInt8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapInt8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8 failed.")
	}
}

func isPositiveInt8(num int8) bool {
	return num > 0
}
func multiplyBy2Int8(num int8) int8 {
	return num * 2
}

func TestFilterMapUint64(t *testing.T) {
	// Test : Multiply all even numbers in the list by 2
	expectedFilteredList := []uint64{4, 8}
	filteredList := FilterMapUint64(isEvenUint64, multiplyBy2Uint64, []uint64{1, 3, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64 failed.")
	}
}

func multiplyBy2Uint64(num uint64) uint64 {
	return num * 2
}

func TestFilterMapUint32(t *testing.T) {
	// Test : Multiply all even numbers in the list by 2
	expectedFilteredList := []uint32{4, 8}
	filteredList := FilterMapUint32(isEvenUint32, multiplyBy2Uint32, []uint32{1, 3, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32 failed.")
	}
}

func multiplyBy2Uint32(num uint32) uint32 {
	return num * 2
}

func TestFilterMapUint16(t *testing.T) {
	// Test : Multiply all even numbers in the list by 2
	expectedFilteredList := []uint16{4, 8}
	filteredList := FilterMapUint16(isEvenUint16, multiplyBy2Uint16, []uint16{1, 3, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16 failed.")
	}
}

func multiplyBy2Uint16(num uint16) uint16 {
	return num * 2
}

func TestFilterMapUint8(t *testing.T) {
	// Test : Multiply all even numbers in the list by 2
	expectedFilteredList := []uint8{4, 8}
	filteredList := FilterMapUint8(isEvenUint8, multiplyBy2Uint8, []uint8{1, 3, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8 failed.")
	}
}

func multiplyBy2Uint8(num uint8) uint8 {
	return num * 2
}

func TestFilterMapUint(t *testing.T) {
	// Test : Multiply all even numbers in the list by 2
	expectedFilteredList := []uint{4, 8}
	filteredList := FilterMapUint(isEvenUint, multiplyBy2Uint, []uint{1, 3, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapUint failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapUint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint failed.")
	}
}

func multiplyBy2Uint(num uint) uint {
	return num * 2
}

func TestFilterMapFloat64(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []float64{4, 8}
	filteredList := FilterMapFloat64(isPositiveFloat64, multiplyBy2Float64, []float64{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapFloat64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapFloat64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64 failed.")
	}
}

func multiplyBy2Float64(num float64) float64 {
	return num * 2
}

func TestFilterMapFloat32(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2
	expectedFilteredList := []float32{4, 8}
	filteredList := FilterMapFloat32(isPositiveFloat32, multiplyBy2Float32, []float32{-1, 0, 2, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapFloat32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapFloat32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32 failed.")
	}
}

func multiplyBy2Float32(num float32) float32 {
	return num * 2
}

func TestFilterMapStr(t *testing.T) {
	// Test : filter all the names in the list which length is <10 and change them to upper case.
	expectedFilteredList := []string{"GOPAL", "GOVINDA"}
	filteredList := FilterMapStr(isStringLengthLessThan10, strings.ToUpper, []string{"gopal", "govinda", "nandeshwar", "Nandeshwar Sah"})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMapStr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStr failed.")
	}
}
