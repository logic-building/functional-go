package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntInt64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapIntInt64(notOneIntInt64, plusOneIntInt64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt64 failed")
	}

	if len(FilterMapIntInt64(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt64(num int) bool {
	return num != 1
}

func TestFilterMapIntInt32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapIntInt32(notOneIntInt32, plusOneIntInt32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt32 failed")
	}

	if len(FilterMapIntInt32(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt32(num int) bool {
	return num != 1
}

func TestFilterMapIntInt16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapIntInt16(notOneIntInt16, plusOneIntInt16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt16 failed")
	}

	if len(FilterMapIntInt16(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt16(num int) bool {
	return num != 1
}

func TestFilterMapIntInt8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapIntInt8(notOneIntInt8, plusOneIntInt8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt8 failed")
	}

	if len(FilterMapIntInt8(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt8(num int) bool {
	return num != 1
}

func TestFilterMapIntUint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapIntUint(notOneIntUint, plusOneIntUint, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint failed")
	}

	if len(FilterMapIntUint(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint(num int) bool {
	return num != 1
}

func TestFilterMapIntUint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapIntUint64(notOneIntUint64, plusOneIntUint64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint64 failed")
	}

	if len(FilterMapIntUint64(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint64(num int) bool {
	return num != 1
}

func TestFilterMapIntUint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapIntUint32(notOneIntUint32, plusOneIntUint32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint32 failed")
	}

	if len(FilterMapIntUint32(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint32(num int) bool {
	return num != 1
}

func TestFilterMapIntUint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapIntUint16(notOneIntUint16, plusOneIntUint16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint16 failed")
	}

	if len(FilterMapIntUint16(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint16(num int) bool {
	return num != 1
}

func TestFilterMapIntUint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapIntUint8(notOneIntUint8, plusOneIntUint8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint8 failed")
	}

	if len(FilterMapIntUint8(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint8(num int) bool {
	return num != 1
}

func TestFilterMapIntStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapIntStr(notOneIntStr, someLogicIntStr, []int{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapIntStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntStr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntStr failed")
	}

	if len(FilterMapIntStr(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntStr(num int) bool {
	return num != 1
}

func TestFilterMapIntBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapIntBool(notOneIntBool, someLogicIntBool, []int{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntBool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntBool failed")
	}

	if len(FilterMapIntBool(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntBool(num int) bool {
	return num != 1
}

func TestFilterMapIntFloat32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapIntFloat32(notOneIntFloat32, plusOneIntFloat32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntFloat32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntFloat32 failed")
	}

	if len(FilterMapIntFloat32(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntFloat32(num int) bool {
	return num != 1
}

func TestFilterMapIntFloat64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapIntFloat64(notOneIntFloat64, plusOneIntFloat64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntFloat64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntFloat64 failed")
	}

	if len(FilterMapIntFloat64(nil, nil, []int{})) > 0 {
		t.Errorf("FilterMapIntFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntFloat64(num int) bool {
	return num != 1
}

func TestFilterMapInt64Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapInt64Int(notOneInt64Int, plusOneInt64Int, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int failed")
	}

	if len(FilterMapInt64Int(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapInt64Int32(notOneInt64Int32, plusOneInt64Int32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int32 failed")
	}

	if len(FilterMapInt64Int32(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int32(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapInt64Int16(notOneInt64Int16, plusOneInt64Int16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int16 failed")
	}

	if len(FilterMapInt64Int16(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int16(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapInt64Int8(notOneInt64Int8, plusOneInt64Int8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int8 failed")
	}

	if len(FilterMapInt64Int8(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int8(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapInt64Uint(notOneInt64Uint, plusOneInt64Uint, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint failed")
	}

	if len(FilterMapInt64Uint(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapInt64Uint64(notOneInt64Uint64, plusOneInt64Uint64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint64 failed")
	}

	if len(FilterMapInt64Uint64(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint64(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapInt64Uint32(notOneInt64Uint32, plusOneInt64Uint32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint32 failed")
	}

	if len(FilterMapInt64Uint32(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint32(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapInt64Uint16(notOneInt64Uint16, plusOneInt64Uint16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint16 failed")
	}

	if len(FilterMapInt64Uint16(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint16(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapInt64Uint8(notOneInt64Uint8, plusOneInt64Uint8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint8 failed")
	}

	if len(FilterMapInt64Uint8(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint8(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapInt64Str(notOneInt64Str, someLogicInt64Str, []int64{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Str failed")
	}

	if len(FilterMapInt64Str(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Str(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapInt64Bool(notOneInt64Bool, someLogicInt64Bool, []int64{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Bool failed")
	}

	if len(FilterMapInt64Bool(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Bool(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapInt64Float32(notOneInt64Float32, plusOneInt64Float32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Float32 failed")
	}

	if len(FilterMapInt64Float32(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Float32(num int64) bool {
	return num != 1
}

func TestFilterMapInt64Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapInt64Float64(notOneInt64Float64, plusOneInt64Float64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Float64 failed")
	}

	if len(FilterMapInt64Float64(nil, nil, []int64{})) > 0 {
		t.Errorf("FilterMapInt64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Float64(num int64) bool {
	return num != 1
}

func TestFilterMapInt32Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapInt32Int(notOneInt32Int, plusOneInt32Int, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int failed")
	}

	if len(FilterMapInt32Int(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapInt32Int64(notOneInt32Int64, plusOneInt32Int64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int64 failed")
	}

	if len(FilterMapInt32Int64(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int64(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapInt32Int16(notOneInt32Int16, plusOneInt32Int16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int16 failed")
	}

	if len(FilterMapInt32Int16(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int16(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapInt32Int8(notOneInt32Int8, plusOneInt32Int8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int8 failed")
	}

	if len(FilterMapInt32Int8(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int8(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapInt32Uint(notOneInt32Uint, plusOneInt32Uint, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint failed")
	}

	if len(FilterMapInt32Uint(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapInt32Uint64(notOneInt32Uint64, plusOneInt32Uint64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint64 failed")
	}

	if len(FilterMapInt32Uint64(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint64(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapInt32Uint32(notOneInt32Uint32, plusOneInt32Uint32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint32 failed")
	}

	if len(FilterMapInt32Uint32(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint32(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapInt32Uint16(notOneInt32Uint16, plusOneInt32Uint16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint16 failed")
	}

	if len(FilterMapInt32Uint16(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint16(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapInt32Uint8(notOneInt32Uint8, plusOneInt32Uint8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint8 failed")
	}

	if len(FilterMapInt32Uint8(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint8(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapInt32Str(notOneInt32Str, someLogicInt32Str, []int32{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Str failed")
	}

	if len(FilterMapInt32Str(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Str(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapInt32Bool(notOneInt32Bool, someLogicInt32Bool, []int32{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Bool failed")
	}

	if len(FilterMapInt32Bool(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Bool(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapInt32Float32(notOneInt32Float32, plusOneInt32Float32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Float32 failed")
	}

	if len(FilterMapInt32Float32(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Float32(num int32) bool {
	return num != 1
}

func TestFilterMapInt32Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapInt32Float64(notOneInt32Float64, plusOneInt32Float64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Float64 failed")
	}

	if len(FilterMapInt32Float64(nil, nil, []int32{})) > 0 {
		t.Errorf("FilterMapInt32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Float64(num int32) bool {
	return num != 1
}

func TestFilterMapInt16Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapInt16Int(notOneInt16Int, plusOneInt16Int, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int failed")
	}

	if len(FilterMapInt16Int(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapInt16Int64(notOneInt16Int64, plusOneInt16Int64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int64 failed")
	}

	if len(FilterMapInt16Int64(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int64(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapInt16Int32(notOneInt16Int32, plusOneInt16Int32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int32 failed")
	}

	if len(FilterMapInt16Int32(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int32(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapInt16Int8(notOneInt16Int8, plusOneInt16Int8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int8 failed")
	}

	if len(FilterMapInt16Int8(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int8(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapInt16Uint(notOneInt16Uint, plusOneInt16Uint, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint failed")
	}

	if len(FilterMapInt16Uint(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapInt16Uint64(notOneInt16Uint64, plusOneInt16Uint64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint64 failed")
	}

	if len(FilterMapInt16Uint64(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint64(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapInt16Uint32(notOneInt16Uint32, plusOneInt16Uint32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint32 failed")
	}

	if len(FilterMapInt16Uint32(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint32(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapInt16Uint16(notOneInt16Uint16, plusOneInt16Uint16, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint16 failed")
	}

	if len(FilterMapInt16Uint16(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint16(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapInt16Uint8(notOneInt16Uint8, plusOneInt16Uint8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint8 failed")
	}

	if len(FilterMapInt16Uint8(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint8(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapInt16Str(notOneInt16Str, someLogicInt16Str, []int16{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Str failed")
	}

	if len(FilterMapInt16Str(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Str(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapInt16Bool(notOneInt16Bool, someLogicInt16Bool, []int16{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Bool failed")
	}

	if len(FilterMapInt16Bool(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Bool(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapInt16Float32(notOneInt16Float32, plusOneInt16Float32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Float32 failed")
	}

	if len(FilterMapInt16Float32(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Float32(num int16) bool {
	return num != 1
}

func TestFilterMapInt16Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapInt16Float64(notOneInt16Float64, plusOneInt16Float64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Float64 failed")
	}

	if len(FilterMapInt16Float64(nil, nil, []int16{})) > 0 {
		t.Errorf("FilterMapInt16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Float64(num int16) bool {
	return num != 1
}

func TestFilterMapInt8Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapInt8Int(notOneInt8Int, plusOneInt8Int, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int failed")
	}

	if len(FilterMapInt8Int(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapInt8Int64(notOneInt8Int64, plusOneInt8Int64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int64 failed")
	}

	if len(FilterMapInt8Int64(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int64(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapInt8Int32(notOneInt8Int32, plusOneInt8Int32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int32 failed")
	}

	if len(FilterMapInt8Int32(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int32(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapInt8Int16(notOneInt8Int16, plusOneInt8Int16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int16 failed")
	}

	if len(FilterMapInt8Int16(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int16(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapInt8Uint(notOneInt8Uint, plusOneInt8Uint, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint failed")
	}

	if len(FilterMapInt8Uint(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapInt8Uint64(notOneInt8Uint64, plusOneInt8Uint64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint64 failed")
	}

	if len(FilterMapInt8Uint64(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint64(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapInt8Uint32(notOneInt8Uint32, plusOneInt8Uint32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint32 failed")
	}

	if len(FilterMapInt8Uint32(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint32(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapInt8Uint16(notOneInt8Uint16, plusOneInt8Uint16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint16 failed")
	}

	if len(FilterMapInt8Uint16(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint16(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapInt8Uint8(notOneInt8Uint8, plusOneInt8Uint8, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint8 failed")
	}

	if len(FilterMapInt8Uint8(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint8(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapInt8Str(notOneInt8Str, someLogicInt8Str, []int8{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Str failed")
	}

	if len(FilterMapInt8Str(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Str(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapInt8Bool(notOneInt8Bool, someLogicInt8Bool, []int8{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Bool failed")
	}

	if len(FilterMapInt8Bool(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Bool(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapInt8Float32(notOneInt8Float32, plusOneInt8Float32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Float32 failed")
	}

	if len(FilterMapInt8Float32(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Float32(num int8) bool {
	return num != 1
}

func TestFilterMapInt8Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapInt8Float64(notOneInt8Float64, plusOneInt8Float64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Float64 failed")
	}

	if len(FilterMapInt8Float64(nil, nil, []int8{})) > 0 {
		t.Errorf("FilterMapInt8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Float64(num int8) bool {
	return num != 1
}

func TestFilterMapUintInt(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapUintInt(notOneUintInt, plusOneUintInt, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt failed")
	}

	if len(FilterMapUintInt(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt(num uint) bool {
	return num != 1
}

func TestFilterMapUintInt64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapUintInt64(notOneUintInt64, plusOneUintInt64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt64 failed")
	}

	if len(FilterMapUintInt64(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt64(num uint) bool {
	return num != 1
}

func TestFilterMapUintInt32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapUintInt32(notOneUintInt32, plusOneUintInt32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt32 failed")
	}

	if len(FilterMapUintInt32(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt32(num uint) bool {
	return num != 1
}

func TestFilterMapUintInt16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapUintInt16(notOneUintInt16, plusOneUintInt16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt16 failed")
	}

	if len(FilterMapUintInt16(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt16(num uint) bool {
	return num != 1
}

func TestFilterMapUintInt8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapUintInt8(notOneUintInt8, plusOneUintInt8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt8 failed")
	}

	if len(FilterMapUintInt8(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt8(num uint) bool {
	return num != 1
}

func TestFilterMapUintUint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapUintUint64(notOneUintUint64, plusOneUintUint64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint64 failed")
	}

	if len(FilterMapUintUint64(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint64(num uint) bool {
	return num != 1
}

func TestFilterMapUintUint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapUintUint32(notOneUintUint32, plusOneUintUint32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint32 failed")
	}

	if len(FilterMapUintUint32(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint32(num uint) bool {
	return num != 1
}

func TestFilterMapUintUint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapUintUint16(notOneUintUint16, plusOneUintUint16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint16 failed")
	}

	if len(FilterMapUintUint16(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint16(num uint) bool {
	return num != 1
}

func TestFilterMapUintUint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapUintUint8(notOneUintUint8, plusOneUintUint8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint8 failed")
	}

	if len(FilterMapUintUint8(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint8(num uint) bool {
	return num != 1
}

func TestFilterMapUintStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapUintStr(notOneUintStr, someLogicUintStr, []uint{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUintStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintStr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintStr failed")
	}

	if len(FilterMapUintStr(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintStr(num uint) bool {
	return num != 1
}

func TestFilterMapUintBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapUintBool(notOneUintBool, someLogicUintBool, []uint{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintBool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintBool failed")
	}

	if len(FilterMapUintBool(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintBool(num uint) bool {
	return num != 1
}

func TestFilterMapUintFloat32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapUintFloat32(notOneUintFloat32, plusOneUintFloat32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintFloat32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintFloat32 failed")
	}

	if len(FilterMapUintFloat32(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintFloat32(num uint) bool {
	return num != 1
}

func TestFilterMapUintFloat64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapUintFloat64(notOneUintFloat64, plusOneUintFloat64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintFloat64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintFloat64 failed")
	}

	if len(FilterMapUintFloat64(nil, nil, []uint{})) > 0 {
		t.Errorf("FilterMapUintFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintFloat64(num uint) bool {
	return num != 1
}

func TestFilterMapUint64Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapUint64Int(notOneUint64Int, plusOneUint64Int, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int failed")
	}

	if len(FilterMapUint64Int(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapUint64Int64(notOneUint64Int64, plusOneUint64Int64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int64 failed")
	}

	if len(FilterMapUint64Int64(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int64(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapUint64Int32(notOneUint64Int32, plusOneUint64Int32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int32 failed")
	}

	if len(FilterMapUint64Int32(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int32(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapUint64Int16(notOneUint64Int16, plusOneUint64Int16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int16 failed")
	}

	if len(FilterMapUint64Int16(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int16(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapUint64Int8(notOneUint64Int8, plusOneUint64Int8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int8 failed")
	}

	if len(FilterMapUint64Int8(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int8(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapUint64Uint(notOneUint64Uint, plusOneUint64Uint, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint failed")
	}

	if len(FilterMapUint64Uint(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapUint64Uint32(notOneUint64Uint32, plusOneUint64Uint32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint32 failed")
	}

	if len(FilterMapUint64Uint32(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint32(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapUint64Uint16(notOneUint64Uint16, plusOneUint64Uint16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint16 failed")
	}

	if len(FilterMapUint64Uint16(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint16(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapUint64Uint8(notOneUint64Uint8, plusOneUint64Uint8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint8 failed")
	}

	if len(FilterMapUint64Uint8(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint8(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapUint64Str(notOneUint64Str, someLogicUint64Str, []uint64{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Str failed")
	}

	if len(FilterMapUint64Str(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Str(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapUint64Bool(notOneUint64Bool, someLogicUint64Bool, []uint64{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Bool failed")
	}

	if len(FilterMapUint64Bool(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Bool(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapUint64Float32(notOneUint64Float32, plusOneUint64Float32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Float32 failed")
	}

	if len(FilterMapUint64Float32(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Float32(num uint64) bool {
	return num != 1
}

func TestFilterMapUint64Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapUint64Float64(notOneUint64Float64, plusOneUint64Float64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Float64 failed")
	}

	if len(FilterMapUint64Float64(nil, nil, []uint64{})) > 0 {
		t.Errorf("FilterMapUint64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Float64(num uint64) bool {
	return num != 1
}

func TestFilterMapUint32Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapUint32Int(notOneUint32Int, plusOneUint32Int, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int failed")
	}

	if len(FilterMapUint32Int(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapUint32Int64(notOneUint32Int64, plusOneUint32Int64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int64 failed")
	}

	if len(FilterMapUint32Int64(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int64(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapUint32Int32(notOneUint32Int32, plusOneUint32Int32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int32 failed")
	}

	if len(FilterMapUint32Int32(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int32(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapUint32Int16(notOneUint32Int16, plusOneUint32Int16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int16 failed")
	}

	if len(FilterMapUint32Int16(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int16(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapUint32Int8(notOneUint32Int8, plusOneUint32Int8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int8 failed")
	}

	if len(FilterMapUint32Int8(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int8(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapUint32Uint(notOneUint32Uint, plusOneUint32Uint, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint failed")
	}

	if len(FilterMapUint32Uint(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapUint32Uint64(notOneUint32Uint64, plusOneUint32Uint64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint64 failed")
	}

	if len(FilterMapUint32Uint64(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint64(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapUint32Uint16(notOneUint32Uint16, plusOneUint32Uint16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint16 failed")
	}

	if len(FilterMapUint32Uint16(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint16(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapUint32Uint8(notOneUint32Uint8, plusOneUint32Uint8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint8 failed")
	}

	if len(FilterMapUint32Uint8(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint8(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapUint32Str(notOneUint32Str, someLogicUint32Str, []uint32{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Str failed")
	}

	if len(FilterMapUint32Str(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Str(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapUint32Bool(notOneUint32Bool, someLogicUint32Bool, []uint32{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Bool failed")
	}

	if len(FilterMapUint32Bool(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Bool(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapUint32Float32(notOneUint32Float32, plusOneUint32Float32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Float32 failed")
	}

	if len(FilterMapUint32Float32(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Float32(num uint32) bool {
	return num != 1
}

func TestFilterMapUint32Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapUint32Float64(notOneUint32Float64, plusOneUint32Float64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Float64 failed")
	}

	if len(FilterMapUint32Float64(nil, nil, []uint32{})) > 0 {
		t.Errorf("FilterMapUint32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Float64(num uint32) bool {
	return num != 1
}

func TestFilterMapUint16Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapUint16Int(notOneUint16Int, plusOneUint16Int, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int failed")
	}

	if len(FilterMapUint16Int(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapUint16Int64(notOneUint16Int64, plusOneUint16Int64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int64 failed")
	}

	if len(FilterMapUint16Int64(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int64(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapUint16Int32(notOneUint16Int32, plusOneUint16Int32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int32 failed")
	}

	if len(FilterMapUint16Int32(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int32(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapUint16Int16(notOneUint16Int16, plusOneUint16Int16, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int16 failed")
	}

	if len(FilterMapUint16Int16(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int16(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapUint16Int8(notOneUint16Int8, plusOneUint16Int8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int8 failed")
	}

	if len(FilterMapUint16Int8(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int8(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapUint16Uint(notOneUint16Uint, plusOneUint16Uint, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint failed")
	}

	if len(FilterMapUint16Uint(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapUint16Uint64(notOneUint16Uint64, plusOneUint16Uint64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint64 failed")
	}

	if len(FilterMapUint16Uint64(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint64(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapUint16Uint32(notOneUint16Uint32, plusOneUint16Uint32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint32 failed")
	}

	if len(FilterMapUint16Uint32(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint32(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapUint16Uint8(notOneUint16Uint8, plusOneUint16Uint8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint8 failed")
	}

	if len(FilterMapUint16Uint8(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint8(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapUint16Str(notOneUint16Str, someLogicUint16Str, []uint16{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Str failed")
	}

	if len(FilterMapUint16Str(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Str(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapUint16Bool(notOneUint16Bool, someLogicUint16Bool, []uint16{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Bool failed")
	}

	if len(FilterMapUint16Bool(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Bool(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapUint16Float32(notOneUint16Float32, plusOneUint16Float32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Float32 failed")
	}

	if len(FilterMapUint16Float32(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Float32(num uint16) bool {
	return num != 1
}

func TestFilterMapUint16Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapUint16Float64(notOneUint16Float64, plusOneUint16Float64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Float64 failed")
	}

	if len(FilterMapUint16Float64(nil, nil, []uint16{})) > 0 {
		t.Errorf("FilterMapUint16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Float64(num uint16) bool {
	return num != 1
}

func TestFilterMapUint8Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapUint8Int(notOneUint8Int, plusOneUint8Int, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int failed")
	}

	if len(FilterMapUint8Int(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapUint8Int64(notOneUint8Int64, plusOneUint8Int64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int64 failed")
	}

	if len(FilterMapUint8Int64(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int64(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapUint8Int32(notOneUint8Int32, plusOneUint8Int32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int32 failed")
	}

	if len(FilterMapUint8Int32(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int32(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapUint8Int16(notOneUint8Int16, plusOneUint8Int16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int16 failed")
	}

	if len(FilterMapUint8Int16(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int16(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapUint8Int8(notOneUint8Int8, plusOneUint8Int8, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int8 failed")
	}

	if len(FilterMapUint8Int8(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int8(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapUint8Uint(notOneUint8Uint, plusOneUint8Uint, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint failed")
	}

	if len(FilterMapUint8Uint(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapUint8Uint64(notOneUint8Uint64, plusOneUint8Uint64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint64 failed")
	}

	if len(FilterMapUint8Uint64(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint64(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapUint8Uint32(notOneUint8Uint32, plusOneUint8Uint32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint32 failed")
	}

	if len(FilterMapUint8Uint32(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint32(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapUint8Uint16(notOneUint8Uint16, plusOneUint8Uint16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint16 failed")
	}

	if len(FilterMapUint8Uint16(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint16(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapUint8Str(notOneUint8Str, someLogicUint8Str, []uint8{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Str failed")
	}

	if len(FilterMapUint8Str(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Str(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapUint8Bool(notOneUint8Bool, someLogicUint8Bool, []uint8{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Bool failed")
	}

	if len(FilterMapUint8Bool(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Bool(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapUint8Float32(notOneUint8Float32, plusOneUint8Float32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Float32 failed")
	}

	if len(FilterMapUint8Float32(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Float32(num uint8) bool {
	return num != 1
}

func TestFilterMapUint8Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapUint8Float64(notOneUint8Float64, plusOneUint8Float64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Float64 failed")
	}

	if len(FilterMapUint8Float64(nil, nil, []uint8{})) > 0 {
		t.Errorf("FilterMapUint8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Float64(num uint8) bool {
	return num != 1
}

func TestFilterMapStrInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10}
	newList := FilterMapStrInt(notOneStrInt, someLogicStrInt, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt failed")
	}

	if len(FilterMapStrInt(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt(num string) bool {
	return num != "one"
}

func TestFilterMapStrInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10}
	newList := FilterMapStrInt64(notOneStrInt64, someLogicStrInt64, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt64 failed")
	}

	if len(FilterMapStrInt64(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt64(num string) bool {
	return num != "one"
}

func TestFilterMapStrInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10}
	newList := FilterMapStrInt32(notOneStrInt32, someLogicStrInt32, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt32 failed")
	}

	if len(FilterMapStrInt32(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt32(num string) bool {
	return num != "one"
}

func TestFilterMapStrInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10}
	newList := FilterMapStrInt16(notOneStrInt16, someLogicStrInt16, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt16 failed")
	}

	if len(FilterMapStrInt16(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt16(num string) bool {
	return num != "one"
}

func TestFilterMapStrInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10}
	newList := FilterMapStrInt8(notOneStrInt8, someLogicStrInt8, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt8 failed")
	}

	if len(FilterMapStrInt8(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt8(num string) bool {
	return num != "one"
}

func TestFilterMapStrUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10}
	newList := FilterMapStrUint(notOneStrUint, someLogicStrUint, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint failed")
	}

	if len(FilterMapStrUint(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint(num string) bool {
	return num != "one"
}

func TestFilterMapStrUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10}
	newList := FilterMapStrUint64(notOneStrUint64, someLogicStrUint64, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint64 failed")
	}

	if len(FilterMapStrUint64(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint64(num string) bool {
	return num != "one"
}

func TestFilterMapStrUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10}
	newList := FilterMapStrUint32(notOneStrUint32, someLogicStrUint32, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint32 failed")
	}

	if len(FilterMapStrUint32(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint32(num string) bool {
	return num != "one"
}

func TestFilterMapStrUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10}
	newList := FilterMapStrUint16(notOneStrUint16, someLogicStrUint16, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint16 failed")
	}

	if len(FilterMapStrUint16(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint16(num string) bool {
	return num != "one"
}

func TestFilterMapStrUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10}
	newList := FilterMapStrUint8(notOneStrUint8, someLogicStrUint8, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint8 failed")
	}

	if len(FilterMapStrUint8(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint8(num string) bool {
	return num != "one"
}

func TestFilterMapStrBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapStrBool(notOneStrBool, someLogicStrBool, []string{"1", "10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapStrBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrBool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrBool failed")
	}

	if len(FilterMapStrBool(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrBool(num string) bool {
	return num != "1"
}

func TestFilterMapStrFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10}
	newList := FilterMapStrFloat32(notOneStrFloat32, someLogicStrFloat32, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrFloat32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrFloat32 failed")
	}

	if len(FilterMapStrFloat32(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrFloat32(num string) bool {
	return num != "one"
}

func TestFilterMapStrFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10}
	newList := FilterMapStrFloat64(notOneStrFloat64, someLogicStrFloat64, []string{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrFloat64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrFloat64 failed")
	}

	if len(FilterMapStrFloat64(nil, nil, []string{})) > 0 {
		t.Errorf("FilterMapStrFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrFloat64(num string) bool {
	return num != "one"
}

func TestFilterMapBoolInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10, 10}
	newList := FilterMapBoolInt(notOneBoolInt, someLogicBoolInt, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt failed")
	}

	if len(FilterMapBoolInt(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt(num bool) bool {
	return num == true
}

func TestFilterMapBoolInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10, 10}
	newList := FilterMapBoolInt64(notOneBoolInt64, someLogicBoolInt64, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt64 failed")
	}

	if len(FilterMapBoolInt64(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt64(num bool) bool {
	return num == true
}

func TestFilterMapBoolInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10, 10}
	newList := FilterMapBoolInt32(notOneBoolInt32, someLogicBoolInt32, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt32 failed")
	}

	if len(FilterMapBoolInt32(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt32(num bool) bool {
	return num == true
}

func TestFilterMapBoolInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10, 10}
	newList := FilterMapBoolInt16(notOneBoolInt16, someLogicBoolInt16, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt16 failed")
	}

	if len(FilterMapBoolInt16(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt16(num bool) bool {
	return num == true
}

func TestFilterMapBoolInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10, 10}
	newList := FilterMapBoolInt8(notOneBoolInt8, someLogicBoolInt8, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt8 failed")
	}

	if len(FilterMapBoolInt8(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt8(num bool) bool {
	return num == true
}

func TestFilterMapBoolUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10, 10}
	newList := FilterMapBoolUint(notOneBoolUint, someLogicBoolUint, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint failed")
	}

	if len(FilterMapBoolUint(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint(num bool) bool {
	return num == true
}

func TestFilterMapBoolUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10, 10}
	newList := FilterMapBoolUint64(notOneBoolUint64, someLogicBoolUint64, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint64 failed")
	}

	if len(FilterMapBoolUint64(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint64(num bool) bool {
	return num == true
}

func TestFilterMapBoolUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10, 10}
	newList := FilterMapBoolUint32(notOneBoolUint32, someLogicBoolUint32, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint32 failed")
	}

	if len(FilterMapBoolUint32(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint32(num bool) bool {
	return num == true
}

func TestFilterMapBoolUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10, 10}
	newList := FilterMapBoolUint16(notOneBoolUint16, someLogicBoolUint16, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint16 failed")
	}

	if len(FilterMapBoolUint16(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint16(num bool) bool {
	return num == true
}

func TestFilterMapBoolUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10, 10}
	newList := FilterMapBoolUint8(notOneBoolUint8, someLogicBoolUint8, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint8 failed")
	}

	if len(FilterMapBoolUint8(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint8(num bool) bool {
	return num == true
}

func TestFilterMapBoolStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10", "10"}
	newList := FilterMapBoolStr(notOneBoolStr, someLogicBoolStr, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolStr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolStr failed")
	}

	if len(FilterMapBoolStr(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolStr(num bool) bool {
	return num == true
}

func TestFilterMapBoolFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10, 10}
	newList := FilterMapBoolFloat32(notOneBoolFloat32, someLogicBoolFloat32, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolFloat32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolFloat32 failed")
	}

	if len(FilterMapBoolFloat32(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolFloat32(num bool) bool {
	return num == true
}

func TestFilterMapBoolFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10, 10}
	newList := FilterMapBoolFloat64(notOneBoolFloat64, someLogicBoolFloat64, []bool{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolFloat64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolFloat64 failed")
	}

	if len(FilterMapBoolFloat64(nil, nil, []bool{})) > 0 {
		t.Errorf("FilterMapBoolFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolFloat64(num bool) bool {
	return num == true
}

func TestFilterMapFloat32Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapFloat32Int(notOneFloat32Int, plusOneFloat32Int, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int failed")
	}

	if len(FilterMapFloat32Int(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapFloat32Int64(notOneFloat32Int64, plusOneFloat32Int64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int64 failed")
	}

	if len(FilterMapFloat32Int64(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int64(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapFloat32Int32(notOneFloat32Int32, plusOneFloat32Int32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int32 failed")
	}

	if len(FilterMapFloat32Int32(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int32(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapFloat32Int16(notOneFloat32Int16, plusOneFloat32Int16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int16 failed")
	}

	if len(FilterMapFloat32Int16(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int16(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapFloat32Int8(notOneFloat32Int8, plusOneFloat32Int8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int8 failed")
	}

	if len(FilterMapFloat32Int8(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int8(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapFloat32Uint(notOneFloat32Uint, plusOneFloat32Uint, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint failed")
	}

	if len(FilterMapFloat32Uint(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapFloat32Uint64(notOneFloat32Uint64, plusOneFloat32Uint64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint64 failed")
	}

	if len(FilterMapFloat32Uint64(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint64(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapFloat32Uint32(notOneFloat32Uint32, plusOneFloat32Uint32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint32 failed")
	}

	if len(FilterMapFloat32Uint32(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint32(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapFloat32Uint16(notOneFloat32Uint16, plusOneFloat32Uint16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint16 failed")
	}

	if len(FilterMapFloat32Uint16(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint16(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapFloat32Uint8(notOneFloat32Uint8, plusOneFloat32Uint8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint8 failed")
	}

	if len(FilterMapFloat32Uint8(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint8(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapFloat32Str(notOneFloat32Str, someLogicFloat32Str, []float32{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapFloat32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Str failed")
	}

	if len(FilterMapFloat32Str(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Str(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapFloat32Bool(notOneFloat32Bool, someLogicFloat32Bool, []float32{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Bool failed")
	}

	if len(FilterMapFloat32Bool(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Bool(num float32) bool {
	return num != 1
}

func TestFilterMapFloat32Float64(t *testing.T) {
	// Test : some logic
	expectedList := []float64{3, 4}
	newList := FilterMapFloat32Float64(notOneFloat32Float64, plusOneFloat32Float64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Float64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Float64 failed")
	}

	if len(FilterMapFloat32Float64(nil, nil, []float32{})) > 0 {
		t.Errorf("FilterMapFloat32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Float64(num float32) bool {
	return num != 1
}

func TestFilterMapFloat64Int(t *testing.T) {
	// Test : some logic
	expectedList := []int{3, 4}
	newList := FilterMapFloat64Int(notOneFloat64Int, plusOneFloat64Int, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int failed")
	}

	if len(FilterMapFloat64Int(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Int64(t *testing.T) {
	// Test : some logic
	expectedList := []int64{3, 4}
	newList := FilterMapFloat64Int64(notOneFloat64Int64, plusOneFloat64Int64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int64 failed")
	}

	if len(FilterMapFloat64Int64(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int64(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Int32(t *testing.T) {
	// Test : some logic
	expectedList := []int32{3, 4}
	newList := FilterMapFloat64Int32(notOneFloat64Int32, plusOneFloat64Int32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int32 failed")
	}

	if len(FilterMapFloat64Int32(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int32(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Int16(t *testing.T) {
	// Test : some logic
	expectedList := []int16{3, 4}
	newList := FilterMapFloat64Int16(notOneFloat64Int16, plusOneFloat64Int16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int16 failed")
	}

	if len(FilterMapFloat64Int16(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int16(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Int8(t *testing.T) {
	// Test : some logic
	expectedList := []int8{3, 4}
	newList := FilterMapFloat64Int8(notOneFloat64Int8, plusOneFloat64Int8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int8 failed")
	}

	if len(FilterMapFloat64Int8(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int8(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Uint(t *testing.T) {
	// Test : some logic
	expectedList := []uint{3, 4}
	newList := FilterMapFloat64Uint(notOneFloat64Uint, plusOneFloat64Uint, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint failed")
	}

	if len(FilterMapFloat64Uint(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Uint64(t *testing.T) {
	// Test : some logic
	expectedList := []uint64{3, 4}
	newList := FilterMapFloat64Uint64(notOneFloat64Uint64, plusOneFloat64Uint64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint64(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint64 failed")
	}

	if len(FilterMapFloat64Uint64(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint64(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Uint32(t *testing.T) {
	// Test : some logic
	expectedList := []uint32{3, 4}
	newList := FilterMapFloat64Uint32(notOneFloat64Uint32, plusOneFloat64Uint32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint32 failed")
	}

	if len(FilterMapFloat64Uint32(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint32(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Uint16(t *testing.T) {
	// Test : some logic
	expectedList := []uint16{3, 4}
	newList := FilterMapFloat64Uint16(notOneFloat64Uint16, plusOneFloat64Uint16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint16(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint16 failed")
	}

	if len(FilterMapFloat64Uint16(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint16(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Uint8(t *testing.T) {
	// Test : some logic
	expectedList := []uint8{3, 4}
	newList := FilterMapFloat64Uint8(notOneFloat64Uint8, plusOneFloat64Uint8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint8(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint8 failed")
	}

	if len(FilterMapFloat64Uint8(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint8(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := FilterMapFloat64Str(notOneFloat64Str, someLogicFloat64Str, []float64{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapFloat64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Str(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Str failed")
	}

	if len(FilterMapFloat64Str(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Str(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := FilterMapFloat64Bool(notOneFloat64Bool, someLogicFloat64Bool, []float64{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Bool(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Bool failed")
	}

	if len(FilterMapFloat64Bool(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Bool(num float64) bool {
	return num != 1
}

func TestFilterMapFloat64Float32(t *testing.T) {
	// Test : some logic
	expectedList := []float32{3, 4}
	newList := FilterMapFloat64Float32(notOneFloat64Float32, plusOneFloat64Float32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Float32(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Float32 failed")
	}

	if len(FilterMapFloat64Float32(nil, nil, []float64{})) > 0 {
		t.Errorf("FilterMapFloat64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Float32(num float64) bool {
	return num != 1
}
