package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapIntInt64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapIntInt64Err(plusOneIntInt64Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt64Err failed")
	}

	r, _ = PMapIntInt64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt64Err failed")
	}
	_, err := PMapIntInt64Err(plusOneIntInt64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntInt32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapIntInt32Err(plusOneIntInt32Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt32Err failed")
	}

	r, _ = PMapIntInt32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt32Err failed")
	}
	_, err := PMapIntInt32Err(plusOneIntInt32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntInt16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapIntInt16Err(plusOneIntInt16Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt16Err failed")
	}

	r, _ = PMapIntInt16Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt16Err failed")
	}
	_, err := PMapIntInt16Err(plusOneIntInt16Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntInt8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapIntInt8Err(plusOneIntInt8Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt8Err failed")
	}

	r, _ = PMapIntInt8Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt8Err failed")
	}
	_, err := PMapIntInt8Err(plusOneIntInt8Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntUintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapIntUintErr(plusOneIntUintErr, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUintErr failed")
	}

	r, _ = PMapIntUintErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUintErr failed")
	}
	_, err := PMapIntUintErr(plusOneIntUintErr, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntUintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntUint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapIntUint64Err(plusOneIntUint64Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint64Err failed")
	}

	r, _ = PMapIntUint64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint64Err failed")
	}
	_, err := PMapIntUint64Err(plusOneIntUint64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntUint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapIntUint32Err(plusOneIntUint32Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint32Err failed")
	}

	r, _ = PMapIntUint32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint32Err failed")
	}
	_, err := PMapIntUint32Err(plusOneIntUint32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntUint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapIntUint16Err(plusOneIntUint16Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint16Err failed")
	}

	r, _ = PMapIntUint16Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint16Err failed")
	}
	_, err := PMapIntUint16Err(plusOneIntUint16Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntUint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapIntUint8Err(plusOneIntUint8Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint8Err failed")
	}

	r, _ = PMapIntUint8Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint8Err failed")
	}
	_, err := PMapIntUint8Err(plusOneIntUint8Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapIntStrErr(someLogicIntStrErr, []int{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapIntStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntStrErr failed")
	}

	r, _ = PMapIntStrErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntStrErr failed")
	}
	_, err := PMapIntStrErr(someLogicIntStrErr, []int{10, 0})
	if err == nil {
		t.Errorf("PMapIntStrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapIntBoolErr(someLogicIntBoolErr, []int{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntBoolErr failed")
	}

	r, _ = PMapIntBoolErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntBoolErr failed")
	}
	_, err := PMapIntBoolErr(someLogicIntBoolErr, []int{10, 0, 3})
	if err == nil {
		t.Errorf("PMapIntBoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntFloat32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapIntFloat32Err(plusOneIntFloat32Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntFloat32Err failed")
	}

	r, _ = PMapIntFloat32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntFloat32Err failed")
	}
	_, err := PMapIntFloat32Err(plusOneIntFloat32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapIntFloat64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapIntFloat64Err(plusOneIntFloat64Err, []int{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntFloat64Err failed")
	}

	r, _ = PMapIntFloat64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntFloat64Err failed")
	}
	_, err := PMapIntFloat64Err(plusOneIntFloat64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("PMapIntFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapInt64IntErr(plusOneInt64IntErr, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64IntErr failed")
	}

	r, _ = PMapInt64IntErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64IntErr failed")
	}
	_, err := PMapInt64IntErr(plusOneInt64IntErr, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapInt64Int32Err(plusOneInt64Int32Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int32Err failed")
	}

	r, _ = PMapInt64Int32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int32Err failed")
	}
	_, err := PMapInt64Int32Err(plusOneInt64Int32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapInt64Int16Err(plusOneInt64Int16Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int16Err failed")
	}

	r, _ = PMapInt64Int16Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int16Err failed")
	}
	_, err := PMapInt64Int16Err(plusOneInt64Int16Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapInt64Int8Err(plusOneInt64Int8Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int8Err failed")
	}

	r, _ = PMapInt64Int8Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int8Err failed")
	}
	_, err := PMapInt64Int8Err(plusOneInt64Int8Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapInt64UintErr(plusOneInt64UintErr, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64UintErr failed")
	}

	r, _ = PMapInt64UintErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64UintErr failed")
	}
	_, err := PMapInt64UintErr(plusOneInt64UintErr, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapInt64Uint64Err(plusOneInt64Uint64Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint64Err failed")
	}

	r, _ = PMapInt64Uint64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint64Err failed")
	}
	_, err := PMapInt64Uint64Err(plusOneInt64Uint64Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapInt64Uint32Err(plusOneInt64Uint32Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint32Err failed")
	}

	r, _ = PMapInt64Uint32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint32Err failed")
	}
	_, err := PMapInt64Uint32Err(plusOneInt64Uint32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapInt64Uint16Err(plusOneInt64Uint16Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint16Err failed")
	}

	r, _ = PMapInt64Uint16Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint16Err failed")
	}
	_, err := PMapInt64Uint16Err(plusOneInt64Uint16Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapInt64Uint8Err(plusOneInt64Uint8Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint8Err failed")
	}

	r, _ = PMapInt64Uint8Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint8Err failed")
	}
	_, err := PMapInt64Uint8Err(plusOneInt64Uint8Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapInt64StrErr(someLogicInt64StrErr, []int64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64StrErr failed")
	}

	r, _ = PMapInt64StrErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64StrErr failed")
	}
	_, err := PMapInt64StrErr(someLogicInt64StrErr, []int64{10, 0})
	if err == nil {
		t.Errorf("PMapInt64StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapInt64BoolErr(someLogicInt64BoolErr, []int64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64BoolErr failed")
	}

	r, _ = PMapInt64BoolErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64BoolErr failed")
	}
	_, err := PMapInt64BoolErr(someLogicInt64BoolErr, []int64{10, 0, 3})
	if err == nil {
		t.Errorf("PMapInt64BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapInt64Float32Err(plusOneInt64Float32Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Float32Err failed")
	}

	r, _ = PMapInt64Float32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Float32Err failed")
	}
	_, err := PMapInt64Float32Err(plusOneInt64Float32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt64Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapInt64Float64Err(plusOneInt64Float64Err, []int64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Float64Err failed")
	}

	r, _ = PMapInt64Float64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Float64Err failed")
	}
	_, err := PMapInt64Float64Err(plusOneInt64Float64Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt64Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapInt32IntErr(plusOneInt32IntErr, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32IntErr failed")
	}

	r, _ = PMapInt32IntErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32IntErr failed")
	}
	_, err := PMapInt32IntErr(plusOneInt32IntErr, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapInt32Int64Err(plusOneInt32Int64Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int64Err failed")
	}

	r, _ = PMapInt32Int64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int64Err failed")
	}
	_, err := PMapInt32Int64Err(plusOneInt32Int64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapInt32Int16Err(plusOneInt32Int16Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int16Err failed")
	}

	r, _ = PMapInt32Int16Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int16Err failed")
	}
	_, err := PMapInt32Int16Err(plusOneInt32Int16Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapInt32Int8Err(plusOneInt32Int8Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int8Err failed")
	}

	r, _ = PMapInt32Int8Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int8Err failed")
	}
	_, err := PMapInt32Int8Err(plusOneInt32Int8Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapInt32UintErr(plusOneInt32UintErr, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32UintErr failed")
	}

	r, _ = PMapInt32UintErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32UintErr failed")
	}
	_, err := PMapInt32UintErr(plusOneInt32UintErr, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapInt32Uint64Err(plusOneInt32Uint64Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint64Err failed")
	}

	r, _ = PMapInt32Uint64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint64Err failed")
	}
	_, err := PMapInt32Uint64Err(plusOneInt32Uint64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapInt32Uint32Err(plusOneInt32Uint32Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint32Err failed")
	}

	r, _ = PMapInt32Uint32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint32Err failed")
	}
	_, err := PMapInt32Uint32Err(plusOneInt32Uint32Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapInt32Uint16Err(plusOneInt32Uint16Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint16Err failed")
	}

	r, _ = PMapInt32Uint16Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint16Err failed")
	}
	_, err := PMapInt32Uint16Err(plusOneInt32Uint16Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapInt32Uint8Err(plusOneInt32Uint8Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint8Err failed")
	}

	r, _ = PMapInt32Uint8Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint8Err failed")
	}
	_, err := PMapInt32Uint8Err(plusOneInt32Uint8Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapInt32StrErr(someLogicInt32StrErr, []int32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32StrErr failed")
	}

	r, _ = PMapInt32StrErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32StrErr failed")
	}
	_, err := PMapInt32StrErr(someLogicInt32StrErr, []int32{10, 0})
	if err == nil {
		t.Errorf("PMapInt32StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapInt32BoolErr(someLogicInt32BoolErr, []int32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32BoolErr failed")
	}

	r, _ = PMapInt32BoolErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32BoolErr failed")
	}
	_, err := PMapInt32BoolErr(someLogicInt32BoolErr, []int32{10, 0, 3})
	if err == nil {
		t.Errorf("PMapInt32BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapInt32Float32Err(plusOneInt32Float32Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Float32Err failed")
	}

	r, _ = PMapInt32Float32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Float32Err failed")
	}
	_, err := PMapInt32Float32Err(plusOneInt32Float32Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapInt32Float64Err(plusOneInt32Float64Err, []int32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Float64Err failed")
	}

	r, _ = PMapInt32Float64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Float64Err failed")
	}
	_, err := PMapInt32Float64Err(plusOneInt32Float64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapInt16IntErr(plusOneInt16IntErr, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16IntErr failed")
	}

	r, _ = PMapInt16IntErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16IntErr failed")
	}
	_, err := PMapInt16IntErr(plusOneInt16IntErr, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapInt16Int64Err(plusOneInt16Int64Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int64Err failed")
	}

	r, _ = PMapInt16Int64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int64Err failed")
	}
	_, err := PMapInt16Int64Err(plusOneInt16Int64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapInt16Int32Err(plusOneInt16Int32Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int32Err failed")
	}

	r, _ = PMapInt16Int32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int32Err failed")
	}
	_, err := PMapInt16Int32Err(plusOneInt16Int32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapInt16Int8Err(plusOneInt16Int8Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int8Err failed")
	}

	r, _ = PMapInt16Int8Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int8Err failed")
	}
	_, err := PMapInt16Int8Err(plusOneInt16Int8Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapInt16UintErr(plusOneInt16UintErr, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16UintErr failed")
	}

	r, _ = PMapInt16UintErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16UintErr failed")
	}
	_, err := PMapInt16UintErr(plusOneInt16UintErr, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapInt16Uint64Err(plusOneInt16Uint64Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint64Err failed")
	}

	r, _ = PMapInt16Uint64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint64Err failed")
	}
	_, err := PMapInt16Uint64Err(plusOneInt16Uint64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapInt16Uint32Err(plusOneInt16Uint32Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint32Err failed")
	}

	r, _ = PMapInt16Uint32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint32Err failed")
	}
	_, err := PMapInt16Uint32Err(plusOneInt16Uint32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapInt16Uint16Err(plusOneInt16Uint16Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint16Err failed")
	}

	r, _ = PMapInt16Uint16Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint16Err failed")
	}
	_, err := PMapInt16Uint16Err(plusOneInt16Uint16Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapInt16Uint8Err(plusOneInt16Uint8Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint8Err failed")
	}

	r, _ = PMapInt16Uint8Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint8Err failed")
	}
	_, err := PMapInt16Uint8Err(plusOneInt16Uint8Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapInt16StrErr(someLogicInt16StrErr, []int16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16StrErr failed")
	}

	r, _ = PMapInt16StrErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16StrErr failed")
	}
	_, err := PMapInt16StrErr(someLogicInt16StrErr, []int16{10, 0})
	if err == nil {
		t.Errorf("PMapInt16StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapInt16BoolErr(someLogicInt16BoolErr, []int16{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16BoolErr failed")
	}

	r, _ = PMapInt16BoolErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16BoolErr failed")
	}
	_, err := PMapInt16BoolErr(someLogicInt16BoolErr, []int16{10, 0, 3})
	if err == nil {
		t.Errorf("PMapInt16BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapInt16Float32Err(plusOneInt16Float32Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Float32Err failed")
	}

	r, _ = PMapInt16Float32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Float32Err failed")
	}
	_, err := PMapInt16Float32Err(plusOneInt16Float32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt16Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapInt16Float64Err(plusOneInt16Float64Err, []int16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Float64Err failed")
	}

	r, _ = PMapInt16Float64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Float64Err failed")
	}
	_, err := PMapInt16Float64Err(plusOneInt16Float64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt16Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapInt8IntErr(plusOneInt8IntErr, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8IntErr failed")
	}

	r, _ = PMapInt8IntErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8IntErr failed")
	}
	_, err := PMapInt8IntErr(plusOneInt8IntErr, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapInt8Int64Err(plusOneInt8Int64Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int64Err failed")
	}

	r, _ = PMapInt8Int64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int64Err failed")
	}
	_, err := PMapInt8Int64Err(plusOneInt8Int64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapInt8Int32Err(plusOneInt8Int32Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int32Err failed")
	}

	r, _ = PMapInt8Int32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int32Err failed")
	}
	_, err := PMapInt8Int32Err(plusOneInt8Int32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapInt8Int16Err(plusOneInt8Int16Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int16Err failed")
	}

	r, _ = PMapInt8Int16Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int16Err failed")
	}
	_, err := PMapInt8Int16Err(plusOneInt8Int16Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapInt8UintErr(plusOneInt8UintErr, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8UintErr failed")
	}

	r, _ = PMapInt8UintErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8UintErr failed")
	}
	_, err := PMapInt8UintErr(plusOneInt8UintErr, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapInt8Uint64Err(plusOneInt8Uint64Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint64Err failed")
	}

	r, _ = PMapInt8Uint64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint64Err failed")
	}
	_, err := PMapInt8Uint64Err(plusOneInt8Uint64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapInt8Uint32Err(plusOneInt8Uint32Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint32Err failed")
	}

	r, _ = PMapInt8Uint32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint32Err failed")
	}
	_, err := PMapInt8Uint32Err(plusOneInt8Uint32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapInt8Uint16Err(plusOneInt8Uint16Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint16Err failed")
	}

	r, _ = PMapInt8Uint16Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint16Err failed")
	}
	_, err := PMapInt8Uint16Err(plusOneInt8Uint16Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapInt8Uint8Err(plusOneInt8Uint8Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint8Err failed")
	}

	r, _ = PMapInt8Uint8Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint8Err failed")
	}
	_, err := PMapInt8Uint8Err(plusOneInt8Uint8Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapInt8StrErr(someLogicInt8StrErr, []int8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8StrErr failed")
	}

	r, _ = PMapInt8StrErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8StrErr failed")
	}
	_, err := PMapInt8StrErr(someLogicInt8StrErr, []int8{10, 0})
	if err == nil {
		t.Errorf("PMapInt8StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapInt8BoolErr(someLogicInt8BoolErr, []int8{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8BoolErr failed")
	}

	r, _ = PMapInt8BoolErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8BoolErr failed")
	}
	_, err := PMapInt8BoolErr(someLogicInt8BoolErr, []int8{10, 0, 3})
	if err == nil {
		t.Errorf("PMapInt8BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapInt8Float32Err(plusOneInt8Float32Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Float32Err failed")
	}

	r, _ = PMapInt8Float32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Float32Err failed")
	}
	_, err := PMapInt8Float32Err(plusOneInt8Float32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapInt8Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapInt8Float64Err(plusOneInt8Float64Err, []int8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Float64Err failed")
	}

	r, _ = PMapInt8Float64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Float64Err failed")
	}
	_, err := PMapInt8Float64Err(plusOneInt8Float64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapInt8Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintIntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapUintIntErr(plusOneUintIntErr, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintIntErr failed")
	}

	r, _ = PMapUintIntErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintIntErr failed")
	}
	_, err := PMapUintIntErr(plusOneUintIntErr, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintIntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintInt64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapUintInt64Err(plusOneUintInt64Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt64Err failed")
	}

	r, _ = PMapUintInt64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt64Err failed")
	}
	_, err := PMapUintInt64Err(plusOneUintInt64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintInt32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapUintInt32Err(plusOneUintInt32Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt32Err failed")
	}

	r, _ = PMapUintInt32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt32Err failed")
	}
	_, err := PMapUintInt32Err(plusOneUintInt32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintInt16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapUintInt16Err(plusOneUintInt16Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt16Err failed")
	}

	r, _ = PMapUintInt16Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt16Err failed")
	}
	_, err := PMapUintInt16Err(plusOneUintInt16Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintInt8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapUintInt8Err(plusOneUintInt8Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt8Err failed")
	}

	r, _ = PMapUintInt8Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt8Err failed")
	}
	_, err := PMapUintInt8Err(plusOneUintInt8Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintUint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapUintUint64Err(plusOneUintUint64Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint64Err failed")
	}

	r, _ = PMapUintUint64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint64Err failed")
	}
	_, err := PMapUintUint64Err(plusOneUintUint64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintUint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapUintUint32Err(plusOneUintUint32Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint32Err failed")
	}

	r, _ = PMapUintUint32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint32Err failed")
	}
	_, err := PMapUintUint32Err(plusOneUintUint32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintUint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapUintUint16Err(plusOneUintUint16Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint16Err failed")
	}

	r, _ = PMapUintUint16Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint16Err failed")
	}
	_, err := PMapUintUint16Err(plusOneUintUint16Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintUint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapUintUint8Err(plusOneUintUint8Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint8Err failed")
	}

	r, _ = PMapUintUint8Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint8Err failed")
	}
	_, err := PMapUintUint8Err(plusOneUintUint8Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapUintStrErr(someLogicUintStrErr, []uint{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUintStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintStrErr failed")
	}

	r, _ = PMapUintStrErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintStrErr failed")
	}
	_, err := PMapUintStrErr(someLogicUintStrErr, []uint{10, 0})
	if err == nil {
		t.Errorf("PMapUintStrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapUintBoolErr(someLogicUintBoolErr, []uint{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintBoolErr failed")
	}

	r, _ = PMapUintBoolErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintBoolErr failed")
	}
	_, err := PMapUintBoolErr(someLogicUintBoolErr, []uint{10, 0, 3})
	if err == nil {
		t.Errorf("PMapUintBoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintFloat32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapUintFloat32Err(plusOneUintFloat32Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintFloat32Err failed")
	}

	r, _ = PMapUintFloat32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintFloat32Err failed")
	}
	_, err := PMapUintFloat32Err(plusOneUintFloat32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUintFloat64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapUintFloat64Err(plusOneUintFloat64Err, []uint{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintFloat64Err failed")
	}

	r, _ = PMapUintFloat64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintFloat64Err failed")
	}
	_, err := PMapUintFloat64Err(plusOneUintFloat64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUintFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapUint64IntErr(plusOneUint64IntErr, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64IntErr failed")
	}

	r, _ = PMapUint64IntErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64IntErr failed")
	}
	_, err := PMapUint64IntErr(plusOneUint64IntErr, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapUint64Int64Err(plusOneUint64Int64Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int64Err failed")
	}

	r, _ = PMapUint64Int64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int64Err failed")
	}
	_, err := PMapUint64Int64Err(plusOneUint64Int64Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapUint64Int32Err(plusOneUint64Int32Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int32Err failed")
	}

	r, _ = PMapUint64Int32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int32Err failed")
	}
	_, err := PMapUint64Int32Err(plusOneUint64Int32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapUint64Int16Err(plusOneUint64Int16Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int16Err failed")
	}

	r, _ = PMapUint64Int16Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int16Err failed")
	}
	_, err := PMapUint64Int16Err(plusOneUint64Int16Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapUint64Int8Err(plusOneUint64Int8Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int8Err failed")
	}

	r, _ = PMapUint64Int8Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int8Err failed")
	}
	_, err := PMapUint64Int8Err(plusOneUint64Int8Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapUint64UintErr(plusOneUint64UintErr, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64UintErr failed")
	}

	r, _ = PMapUint64UintErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64UintErr failed")
	}
	_, err := PMapUint64UintErr(plusOneUint64UintErr, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapUint64Uint32Err(plusOneUint64Uint32Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint32Err failed")
	}

	r, _ = PMapUint64Uint32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint32Err failed")
	}
	_, err := PMapUint64Uint32Err(plusOneUint64Uint32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapUint64Uint16Err(plusOneUint64Uint16Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint16Err failed")
	}

	r, _ = PMapUint64Uint16Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint16Err failed")
	}
	_, err := PMapUint64Uint16Err(plusOneUint64Uint16Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapUint64Uint8Err(plusOneUint64Uint8Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint8Err failed")
	}

	r, _ = PMapUint64Uint8Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint8Err failed")
	}
	_, err := PMapUint64Uint8Err(plusOneUint64Uint8Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapUint64StrErr(someLogicUint64StrErr, []uint64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64StrErr failed")
	}

	r, _ = PMapUint64StrErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64StrErr failed")
	}
	_, err := PMapUint64StrErr(someLogicUint64StrErr, []uint64{10, 0})
	if err == nil {
		t.Errorf("PMapUint64StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapUint64BoolErr(someLogicUint64BoolErr, []uint64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64BoolErr failed")
	}

	r, _ = PMapUint64BoolErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64BoolErr failed")
	}
	_, err := PMapUint64BoolErr(someLogicUint64BoolErr, []uint64{10, 0, 3})
	if err == nil {
		t.Errorf("PMapUint64BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapUint64Float32Err(plusOneUint64Float32Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Float32Err failed")
	}

	r, _ = PMapUint64Float32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Float32Err failed")
	}
	_, err := PMapUint64Float32Err(plusOneUint64Float32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint64Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapUint64Float64Err(plusOneUint64Float64Err, []uint64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Float64Err failed")
	}

	r, _ = PMapUint64Float64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Float64Err failed")
	}
	_, err := PMapUint64Float64Err(plusOneUint64Float64Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint64Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapUint32IntErr(plusOneUint32IntErr, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32IntErr failed")
	}

	r, _ = PMapUint32IntErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32IntErr failed")
	}
	_, err := PMapUint32IntErr(plusOneUint32IntErr, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapUint32Int64Err(plusOneUint32Int64Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int64Err failed")
	}

	r, _ = PMapUint32Int64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int64Err failed")
	}
	_, err := PMapUint32Int64Err(plusOneUint32Int64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapUint32Int32Err(plusOneUint32Int32Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int32Err failed")
	}

	r, _ = PMapUint32Int32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int32Err failed")
	}
	_, err := PMapUint32Int32Err(plusOneUint32Int32Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapUint32Int16Err(plusOneUint32Int16Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int16Err failed")
	}

	r, _ = PMapUint32Int16Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int16Err failed")
	}
	_, err := PMapUint32Int16Err(plusOneUint32Int16Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapUint32Int8Err(plusOneUint32Int8Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int8Err failed")
	}

	r, _ = PMapUint32Int8Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int8Err failed")
	}
	_, err := PMapUint32Int8Err(plusOneUint32Int8Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapUint32UintErr(plusOneUint32UintErr, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32UintErr failed")
	}

	r, _ = PMapUint32UintErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32UintErr failed")
	}
	_, err := PMapUint32UintErr(plusOneUint32UintErr, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapUint32Uint64Err(plusOneUint32Uint64Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint64Err failed")
	}

	r, _ = PMapUint32Uint64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint64Err failed")
	}
	_, err := PMapUint32Uint64Err(plusOneUint32Uint64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapUint32Uint16Err(plusOneUint32Uint16Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint16Err failed")
	}

	r, _ = PMapUint32Uint16Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint16Err failed")
	}
	_, err := PMapUint32Uint16Err(plusOneUint32Uint16Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapUint32Uint8Err(plusOneUint32Uint8Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint8Err failed")
	}

	r, _ = PMapUint32Uint8Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint8Err failed")
	}
	_, err := PMapUint32Uint8Err(plusOneUint32Uint8Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapUint32StrErr(someLogicUint32StrErr, []uint32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32StrErr failed")
	}

	r, _ = PMapUint32StrErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32StrErr failed")
	}
	_, err := PMapUint32StrErr(someLogicUint32StrErr, []uint32{10, 0})
	if err == nil {
		t.Errorf("PMapUint32StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapUint32BoolErr(someLogicUint32BoolErr, []uint32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32BoolErr failed")
	}

	r, _ = PMapUint32BoolErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32BoolErr failed")
	}
	_, err := PMapUint32BoolErr(someLogicUint32BoolErr, []uint32{10, 0, 3})
	if err == nil {
		t.Errorf("PMapUint32BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapUint32Float32Err(plusOneUint32Float32Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Float32Err failed")
	}

	r, _ = PMapUint32Float32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Float32Err failed")
	}
	_, err := PMapUint32Float32Err(plusOneUint32Float32Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapUint32Float64Err(plusOneUint32Float64Err, []uint32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Float64Err failed")
	}

	r, _ = PMapUint32Float64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Float64Err failed")
	}
	_, err := PMapUint32Float64Err(plusOneUint32Float64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapUint16IntErr(plusOneUint16IntErr, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16IntErr failed")
	}

	r, _ = PMapUint16IntErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16IntErr failed")
	}
	_, err := PMapUint16IntErr(plusOneUint16IntErr, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapUint16Int64Err(plusOneUint16Int64Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int64Err failed")
	}

	r, _ = PMapUint16Int64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int64Err failed")
	}
	_, err := PMapUint16Int64Err(plusOneUint16Int64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapUint16Int32Err(plusOneUint16Int32Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int32Err failed")
	}

	r, _ = PMapUint16Int32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int32Err failed")
	}
	_, err := PMapUint16Int32Err(plusOneUint16Int32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapUint16Int16Err(plusOneUint16Int16Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int16Err failed")
	}

	r, _ = PMapUint16Int16Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int16Err failed")
	}
	_, err := PMapUint16Int16Err(plusOneUint16Int16Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapUint16Int8Err(plusOneUint16Int8Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int8Err failed")
	}

	r, _ = PMapUint16Int8Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int8Err failed")
	}
	_, err := PMapUint16Int8Err(plusOneUint16Int8Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapUint16UintErr(plusOneUint16UintErr, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16UintErr failed")
	}

	r, _ = PMapUint16UintErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16UintErr failed")
	}
	_, err := PMapUint16UintErr(plusOneUint16UintErr, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapUint16Uint64Err(plusOneUint16Uint64Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint64Err failed")
	}

	r, _ = PMapUint16Uint64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint64Err failed")
	}
	_, err := PMapUint16Uint64Err(plusOneUint16Uint64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapUint16Uint32Err(plusOneUint16Uint32Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint32Err failed")
	}

	r, _ = PMapUint16Uint32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint32Err failed")
	}
	_, err := PMapUint16Uint32Err(plusOneUint16Uint32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapUint16Uint8Err(plusOneUint16Uint8Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint8Err failed")
	}

	r, _ = PMapUint16Uint8Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint8Err failed")
	}
	_, err := PMapUint16Uint8Err(plusOneUint16Uint8Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapUint16StrErr(someLogicUint16StrErr, []uint16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16StrErr failed")
	}

	r, _ = PMapUint16StrErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16StrErr failed")
	}
	_, err := PMapUint16StrErr(someLogicUint16StrErr, []uint16{10, 0})
	if err == nil {
		t.Errorf("PMapUint16StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapUint16BoolErr(someLogicUint16BoolErr, []uint16{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16BoolErr failed")
	}

	r, _ = PMapUint16BoolErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16BoolErr failed")
	}
	_, err := PMapUint16BoolErr(someLogicUint16BoolErr, []uint16{10, 0, 3})
	if err == nil {
		t.Errorf("PMapUint16BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapUint16Float32Err(plusOneUint16Float32Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Float32Err failed")
	}

	r, _ = PMapUint16Float32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Float32Err failed")
	}
	_, err := PMapUint16Float32Err(plusOneUint16Float32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint16Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapUint16Float64Err(plusOneUint16Float64Err, []uint16{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Float64Err failed")
	}

	r, _ = PMapUint16Float64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Float64Err failed")
	}
	_, err := PMapUint16Float64Err(plusOneUint16Float64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint16Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapUint8IntErr(plusOneUint8IntErr, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8IntErr failed")
	}

	r, _ = PMapUint8IntErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8IntErr failed")
	}
	_, err := PMapUint8IntErr(plusOneUint8IntErr, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapUint8Int64Err(plusOneUint8Int64Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int64Err failed")
	}

	r, _ = PMapUint8Int64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int64Err failed")
	}
	_, err := PMapUint8Int64Err(plusOneUint8Int64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapUint8Int32Err(plusOneUint8Int32Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int32Err failed")
	}

	r, _ = PMapUint8Int32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int32Err failed")
	}
	_, err := PMapUint8Int32Err(plusOneUint8Int32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapUint8Int16Err(plusOneUint8Int16Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int16Err failed")
	}

	r, _ = PMapUint8Int16Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int16Err failed")
	}
	_, err := PMapUint8Int16Err(plusOneUint8Int16Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapUint8Int8Err(plusOneUint8Int8Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int8Err failed")
	}

	r, _ = PMapUint8Int8Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int8Err failed")
	}
	_, err := PMapUint8Int8Err(plusOneUint8Int8Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapUint8UintErr(plusOneUint8UintErr, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8UintErr failed")
	}

	r, _ = PMapUint8UintErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8UintErr failed")
	}
	_, err := PMapUint8UintErr(plusOneUint8UintErr, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapUint8Uint64Err(plusOneUint8Uint64Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint64Err failed")
	}

	r, _ = PMapUint8Uint64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint64Err failed")
	}
	_, err := PMapUint8Uint64Err(plusOneUint8Uint64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapUint8Uint32Err(plusOneUint8Uint32Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint32Err failed")
	}

	r, _ = PMapUint8Uint32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint32Err failed")
	}
	_, err := PMapUint8Uint32Err(plusOneUint8Uint32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapUint8Uint16Err(plusOneUint8Uint16Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint16Err failed")
	}

	r, _ = PMapUint8Uint16Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint16Err failed")
	}
	_, err := PMapUint8Uint16Err(plusOneUint8Uint16Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapUint8StrErr(someLogicUint8StrErr, []uint8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8StrErr failed")
	}

	r, _ = PMapUint8StrErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8StrErr failed")
	}
	_, err := PMapUint8StrErr(someLogicUint8StrErr, []uint8{10, 0})
	if err == nil {
		t.Errorf("PMapUint8StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapUint8BoolErr(someLogicUint8BoolErr, []uint8{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8BoolErr failed")
	}

	r, _ = PMapUint8BoolErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8BoolErr failed")
	}
	_, err := PMapUint8BoolErr(someLogicUint8BoolErr, []uint8{10, 0, 3})
	if err == nil {
		t.Errorf("PMapUint8BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapUint8Float32Err(plusOneUint8Float32Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Float32Err failed")
	}

	r, _ = PMapUint8Float32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Float32Err failed")
	}
	_, err := PMapUint8Float32Err(plusOneUint8Float32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapUint8Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapUint8Float64Err(plusOneUint8Float64Err, []uint8{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Float64Err failed")
	}

	r, _ = PMapUint8Float64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Float64Err failed")
	}
	_, err := PMapUint8Float64Err(plusOneUint8Float64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("PMapUint8Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrIntErr(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10}
	newList, _ := PMapStrIntErr(someLogicStrIntErr, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrIntErr failed")
	}

	r, _ = PMapStrIntErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt failed")
	}
	_, err := PMapStrIntErr(someLogicStrIntErr, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrIntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrInt64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10}
	newList, _ := PMapStrInt64Err(someLogicStrInt64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt64Err failed")
	}

	r, _ = PMapStrInt64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt64 failed")
	}
	_, err := PMapStrInt64Err(someLogicStrInt64Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrInt32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10}
	newList, _ := PMapStrInt32Err(someLogicStrInt32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt32Err failed")
	}

	r, _ = PMapStrInt32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt32 failed")
	}
	_, err := PMapStrInt32Err(someLogicStrInt32Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrInt16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10}
	newList, _ := PMapStrInt16Err(someLogicStrInt16Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt16Err failed")
	}

	r, _ = PMapStrInt16Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt16 failed")
	}
	_, err := PMapStrInt16Err(someLogicStrInt16Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrInt8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10}
	newList, _ := PMapStrInt8Err(someLogicStrInt8Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt8Err failed")
	}

	r, _ = PMapStrInt8Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt8 failed")
	}
	_, err := PMapStrInt8Err(someLogicStrInt8Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrUintErr(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10}
	newList, _ := PMapStrUintErr(someLogicStrUintErr, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUintErr failed")
	}

	r, _ = PMapStrUintErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint failed")
	}
	_, err := PMapStrUintErr(someLogicStrUintErr, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrUintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrUint64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10}
	newList, _ := PMapStrUint64Err(someLogicStrUint64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint64Err failed")
	}

	r, _ = PMapStrUint64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint64 failed")
	}
	_, err := PMapStrUint64Err(someLogicStrUint64Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrUint32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10}
	newList, _ := PMapStrUint32Err(someLogicStrUint32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint32Err failed")
	}

	r, _ = PMapStrUint32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint32 failed")
	}
	_, err := PMapStrUint32Err(someLogicStrUint32Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrUint16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10}
	newList, _ := PMapStrUint16Err(someLogicStrUint16Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint16Err failed")
	}

	r, _ = PMapStrUint16Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint16 failed")
	}
	_, err := PMapStrUint16Err(someLogicStrUint16Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrUint8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10}
	newList, _ := PMapStrUint8Err(someLogicStrUint8Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint8Err failed")
	}

	r, _ = PMapStrUint8Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint8 failed")
	}
	_, err := PMapStrUint8Err(someLogicStrUint8Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapStrBoolErr(someLogicStrBoolErr, []string{"10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapStrBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrBoolErr failed")
	}

	r, _ = PMapStrBoolErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrBoolErr failed")
	}
	_, err := PMapStrBoolErr(someLogicStrBoolErr, []string{"10", "0", "3"})
	if err == nil {
		t.Errorf("PMapStrBoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrFloat32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10}
	newList, _ := PMapStrFloat32Err(someLogicStrFloat32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrFloat32Err failed")
	}

	r, _ = PMapStrFloat32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrFloat32 failed")
	}
	_, err := PMapStrFloat32Err(someLogicStrFloat32Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapStrFloat64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10}
	newList, _ := PMapStrFloat64Err(someLogicStrFloat64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrFloat64Err failed")
	}

	r, _ = PMapStrFloat64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrFloat64 failed")
	}
	_, err := PMapStrFloat64Err(someLogicStrFloat64Err, []string{"ten", "0"})
	if err == nil {
		t.Errorf("PMapStrFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolIntErr(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10, 10}
	newList, _ := PMapBoolIntErr(someLogicBoolIntErr, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolIntErr failed")
	}

	r, _ = PMapBoolIntErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolIntErr failed")
	}
	_, err := PMapBoolIntErr(someLogicBoolIntErr, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolIntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolInt64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10, 10}
	newList, _ := PMapBoolInt64Err(someLogicBoolInt64Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt64Err failed")
	}

	r, _ = PMapBoolInt64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt64Err failed")
	}
	_, err := PMapBoolInt64Err(someLogicBoolInt64Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolInt32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10, 10}
	newList, _ := PMapBoolInt32Err(someLogicBoolInt32Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt32Err failed")
	}

	r, _ = PMapBoolInt32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt32Err failed")
	}
	_, err := PMapBoolInt32Err(someLogicBoolInt32Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolInt16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10, 10}
	newList, _ := PMapBoolInt16Err(someLogicBoolInt16Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt16Err failed")
	}

	r, _ = PMapBoolInt16Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt16Err failed")
	}
	_, err := PMapBoolInt16Err(someLogicBoolInt16Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolInt8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10, 10}
	newList, _ := PMapBoolInt8Err(someLogicBoolInt8Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt8Err failed")
	}

	r, _ = PMapBoolInt8Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt8Err failed")
	}
	_, err := PMapBoolInt8Err(someLogicBoolInt8Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolUintErr(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10, 10}
	newList, _ := PMapBoolUintErr(someLogicBoolUintErr, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUintErr failed")
	}

	r, _ = PMapBoolUintErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUintErr failed")
	}
	_, err := PMapBoolUintErr(someLogicBoolUintErr, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolUintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolUint64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10, 10}
	newList, _ := PMapBoolUint64Err(someLogicBoolUint64Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint64Err failed")
	}

	r, _ = PMapBoolUint64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint64Err failed")
	}
	_, err := PMapBoolUint64Err(someLogicBoolUint64Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolUint32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10, 10}
	newList, _ := PMapBoolUint32Err(someLogicBoolUint32Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint32Err failed")
	}

	r, _ = PMapBoolUint32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint32Err failed")
	}
	_, err := PMapBoolUint32Err(someLogicBoolUint32Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolUint16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10, 10}
	newList, _ := PMapBoolUint16Err(someLogicBoolUint16Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint16Err failed")
	}

	r, _ = PMapBoolUint16Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint16Err failed")
	}
	_, err := PMapBoolUint16Err(someLogicBoolUint16Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolUint8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10, 10}
	newList, _ := PMapBoolUint8Err(someLogicBoolUint8Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint8Err failed")
	}

	r, _ = PMapBoolUint8Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint8Err failed")
	}
	_, err := PMapBoolUint8Err(someLogicBoolUint8Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10", "10"}
	newList, _ := PMapBoolStrErr(someLogicBoolStrErr, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolStrErr failed")
	}

	r, _ = PMapBoolStrErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolStrErr failed")
	}
	_, err := PMapBoolStrErr(someLogicBoolStrErr, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolStrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolFloat32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10, 10}
	newList, _ := PMapBoolFloat32Err(someLogicBoolFloat32Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat32Err failed")
	}

	r, _ = PMapBoolFloat32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat32Err failed")
	}
	_, err := PMapBoolFloat32Err(someLogicBoolFloat32Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapBoolFloat64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10, 10}
	newList, _ := PMapBoolFloat64Err(someLogicBoolFloat64Err, []bool{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat64Err failed")
	}

	r, _ = PMapBoolFloat64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat64Err failed")
	}
	_, err := PMapBoolFloat64Err(someLogicBoolFloat64Err, []bool{true, false})
	if err == nil {
		t.Errorf("PMapBoolFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapFloat32IntErr(plusOneFloat32IntErr, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32IntErr failed")
	}

	r, _ = PMapFloat32IntErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32IntErr failed")
	}
	_, err := PMapFloat32IntErr(plusOneFloat32IntErr, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapFloat32Int64Err(plusOneFloat32Int64Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int64Err failed")
	}

	r, _ = PMapFloat32Int64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int64Err failed")
	}
	_, err := PMapFloat32Int64Err(plusOneFloat32Int64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapFloat32Int32Err(plusOneFloat32Int32Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int32Err failed")
	}

	r, _ = PMapFloat32Int32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int32Err failed")
	}
	_, err := PMapFloat32Int32Err(plusOneFloat32Int32Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapFloat32Int16Err(plusOneFloat32Int16Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int16Err failed")
	}

	r, _ = PMapFloat32Int16Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int16Err failed")
	}
	_, err := PMapFloat32Int16Err(plusOneFloat32Int16Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapFloat32Int8Err(plusOneFloat32Int8Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int8Err failed")
	}

	r, _ = PMapFloat32Int8Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int8Err failed")
	}
	_, err := PMapFloat32Int8Err(plusOneFloat32Int8Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapFloat32UintErr(plusOneFloat32UintErr, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32UintErr failed")
	}

	r, _ = PMapFloat32UintErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32UintErr failed")
	}
	_, err := PMapFloat32UintErr(plusOneFloat32UintErr, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapFloat32Uint64Err(plusOneFloat32Uint64Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint64Err failed")
	}

	r, _ = PMapFloat32Uint64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint64Err failed")
	}
	_, err := PMapFloat32Uint64Err(plusOneFloat32Uint64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapFloat32Uint32Err(plusOneFloat32Uint32Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint32Err failed")
	}

	r, _ = PMapFloat32Uint32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint32Err failed")
	}
	_, err := PMapFloat32Uint32Err(plusOneFloat32Uint32Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapFloat32Uint16Err(plusOneFloat32Uint16Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint16Err failed")
	}

	r, _ = PMapFloat32Uint16Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint16Err failed")
	}
	_, err := PMapFloat32Uint16Err(plusOneFloat32Uint16Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapFloat32Uint8Err(plusOneFloat32Uint8Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint8Err failed")
	}

	r, _ = PMapFloat32Uint8Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint8Err failed")
	}
	_, err := PMapFloat32Uint8Err(plusOneFloat32Uint8Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapFloat32StrErr(someLogicFloat32StrErr, []float32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapFloat32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32StrErr failed")
	}

	r, _ = PMapFloat32StrErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32StrErr failed")
	}
	_, err := PMapFloat32StrErr(someLogicFloat32StrErr, []float32{10, 0})
	if err == nil {
		t.Errorf("PMapFloat32StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapFloat32BoolErr(someLogicFloat32BoolErr, []float32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32BoolErr failed")
	}

	r, _ = PMapFloat32BoolErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32BoolErr failed")
	}
	_, err := PMapFloat32BoolErr(someLogicFloat32BoolErr, []float32{10, 0, 3})
	if err == nil {
		t.Errorf("PMapFloat32BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3}
	newList, _ := PMapFloat32Float64Err(plusOneFloat32Float64Err, []float32{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Float64Err failed")
	}

	r, _ = PMapFloat32Float64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Float64Err failed")
	}
	_, err := PMapFloat32Float64Err(plusOneFloat32Float64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3}
	newList, _ := PMapFloat64IntErr(plusOneFloat64IntErr, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64IntErr failed")
	}

	r, _ = PMapFloat64IntErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64IntErr failed")
	}
	_, err := PMapFloat64IntErr(plusOneFloat64IntErr, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3}
	newList, _ := PMapFloat64Int64Err(plusOneFloat64Int64Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int64Err failed")
	}

	r, _ = PMapFloat64Int64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int64Err failed")
	}
	_, err := PMapFloat64Int64Err(plusOneFloat64Int64Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3}
	newList, _ := PMapFloat64Int32Err(plusOneFloat64Int32Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int32Err failed")
	}

	r, _ = PMapFloat64Int32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int32Err failed")
	}
	_, err := PMapFloat64Int32Err(plusOneFloat64Int32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3}
	newList, _ := PMapFloat64Int16Err(plusOneFloat64Int16Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int16Err failed")
	}

	r, _ = PMapFloat64Int16Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int16Err failed")
	}
	_, err := PMapFloat64Int16Err(plusOneFloat64Int16Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3}
	newList, _ := PMapFloat64Int8Err(plusOneFloat64Int8Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int8Err failed")
	}

	r, _ = PMapFloat64Int8Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int8Err failed")
	}
	_, err := PMapFloat64Int8Err(plusOneFloat64Int8Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3}
	newList, _ := PMapFloat64UintErr(plusOneFloat64UintErr, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64UintErr failed")
	}

	r, _ = PMapFloat64UintErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64UintErr failed")
	}
	_, err := PMapFloat64UintErr(plusOneFloat64UintErr, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3}
	newList, _ := PMapFloat64Uint64Err(plusOneFloat64Uint64Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint64Err failed")
	}

	r, _ = PMapFloat64Uint64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint64Err failed")
	}
	_, err := PMapFloat64Uint64Err(plusOneFloat64Uint64Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3}
	newList, _ := PMapFloat64Uint32Err(plusOneFloat64Uint32Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint32Err failed")
	}

	r, _ = PMapFloat64Uint32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint32Err failed")
	}
	_, err := PMapFloat64Uint32Err(plusOneFloat64Uint32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3}
	newList, _ := PMapFloat64Uint16Err(plusOneFloat64Uint16Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint16Err failed")
	}

	r, _ = PMapFloat64Uint16Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint16Err failed")
	}
	_, err := PMapFloat64Uint16Err(plusOneFloat64Uint16Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3}
	newList, _ := PMapFloat64Uint8Err(plusOneFloat64Uint8Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint8Err failed")
	}

	r, _ = PMapFloat64Uint8Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint8Err failed")
	}
	_, err := PMapFloat64Uint8Err(plusOneFloat64Uint8Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := PMapFloat64StrErr(someLogicFloat64StrErr, []float64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapFloat64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64StrErr failed")
	}

	r, _ = PMapFloat64StrErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64StrErr failed")
	}
	_, err := PMapFloat64StrErr(someLogicFloat64StrErr, []float64{10, 0})
	if err == nil {
		t.Errorf("PMapFloat64StrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := PMapFloat64BoolErr(someLogicFloat64BoolErr, []float64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64BoolErr failed")
	}

	r, _ = PMapFloat64BoolErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64BoolErr failed")
	}
	_, err := PMapFloat64BoolErr(someLogicFloat64BoolErr, []float64{10, 0, 3})
	if err == nil {
		t.Errorf("PMapFloat64BoolErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestPmapFloat64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3}
	newList, _ := PMapFloat64Float32Err(plusOneFloat64Float32Err, []float64{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Float32Err failed")
	}

	r, _ = PMapFloat64Float32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Float32Err failed")
	}
	_, err := PMapFloat64Float32Err(plusOneFloat64Float32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("PMapFloat64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
