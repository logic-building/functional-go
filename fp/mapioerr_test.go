package fp

import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntInt64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapIntInt64Err(plusOneIntInt64Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt64Err failed")
	}

	r, _ = MapIntInt64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt64Err failed")
	}

	_, err := MapIntInt64Err(plusOneIntInt64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapIntInt32Err(plusOneIntInt32Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt32Err failed")
	}

	r, _ = MapIntInt32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt32Err failed")
	}

	_, err := MapIntInt32Err(plusOneIntInt32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapIntInt16Err(plusOneIntInt16Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt16Err failed")
	}

	r, _ = MapIntInt16Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt16Err failed")
	}

	_, err := MapIntInt16Err(plusOneIntInt16Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapIntInt8Err(plusOneIntInt8Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt8Err failed")
	}

	r, _ = MapIntInt8Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt8Err failed")
	}

	_, err := MapIntInt8Err(plusOneIntInt8Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapIntUintErr(plusOneIntUintErr, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUintErr failed")
	}

	r, _ = MapIntUintErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntUintErr failed")
	}

	_, err := MapIntUintErr(plusOneIntUintErr, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntUintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapIntUint64Err(plusOneIntUint64Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint64Err failed")
	}

	r, _ = MapIntUint64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint64Err failed")
	}

	_, err := MapIntUint64Err(plusOneIntUint64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapIntUint32Err(plusOneIntUint32Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint32Err failed")
	}

	r, _ = MapIntUint32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint32Err failed")
	}

	_, err := MapIntUint32Err(plusOneIntUint32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapIntUint16Err(plusOneIntUint16Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint16Err failed")
	}

	r, _ = MapIntUint16Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint16Err failed")
	}

	_, err := MapIntUint16Err(plusOneIntUint16Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapIntUint8Err(plusOneIntUint8Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint8Err failed")
	}

	r, _ = MapIntUint8Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint8Err failed")
	}

	_, err := MapIntUint8Err(plusOneIntUint8Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapIntStrErr(someLogicIntStrErr, []int{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapIntStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntStr failed")
	}

	r, _ = MapIntStrErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntStrErr failed")
	}

	_, err := MapIntStrErr(someLogicIntStrErr, []int{0})
	if err == nil {
		t.Errorf("MapIntStrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicIntStrErr(num int) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapIntBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapIntBoolErr(someLogicIntBoolErr, []int{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapIntBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntBoolErr failed")
	}

	r, _ = MapIntBoolErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntBoolErr failed")
	}

	_, err := MapIntBoolErr(someLogicIntBoolErr, []int{10, 3})
	if err == nil {
		t.Errorf("MapIntBoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapIntFloat32Err(plusOneIntFloat32Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntFloat32Err failed")
	}

	r, _ = MapIntFloat32Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntFloat32Err failed")
	}

	_, err := MapIntFloat32Err(plusOneIntFloat32Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapIntFloat64Err(plusOneIntFloat64Err, []int{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntFloat64Err failed")
	}

	r, _ = MapIntFloat64Err(nil, []int{})
	if len(r) > 0 {
		t.Errorf("MapIntFloat64Err failed")
	}

	_, err := MapIntFloat64Err(plusOneIntFloat64Err, []int{1, 2, 3})
	if err == nil {
		t.Errorf("MapIntFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapInt64IntErr(plusOneInt64IntErr, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64IntErr failed")
	}

	r, _ = MapInt64IntErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64IntErr failed")
	}

	_, err := MapInt64IntErr(plusOneInt64IntErr, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapInt64Int32Err(plusOneInt64Int32Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int32Err failed")
	}

	r, _ = MapInt64Int32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int32Err failed")
	}

	_, err := MapInt64Int32Err(plusOneInt64Int32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapInt64Int16Err(plusOneInt64Int16Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int16Err failed")
	}

	r, _ = MapInt64Int16Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int16Err failed")
	}

	_, err := MapInt64Int16Err(plusOneInt64Int16Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapInt64Int8Err(plusOneInt64Int8Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int8Err failed")
	}

	r, _ = MapInt64Int8Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int8Err failed")
	}

	_, err := MapInt64Int8Err(plusOneInt64Int8Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapInt64UintErr(plusOneInt64UintErr, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64UintErr failed")
	}

	r, _ = MapInt64UintErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64UintErr failed")
	}

	_, err := MapInt64UintErr(plusOneInt64UintErr, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapInt64Uint64Err(plusOneInt64Uint64Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint64Err failed")
	}

	r, _ = MapInt64Uint64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint64Err failed")
	}

	_, err := MapInt64Uint64Err(plusOneInt64Uint64Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapInt64Uint32Err(plusOneInt64Uint32Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint32Err failed")
	}

	r, _ = MapInt64Uint32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint32Err failed")
	}

	_, err := MapInt64Uint32Err(plusOneInt64Uint32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapInt64Uint16Err(plusOneInt64Uint16Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint16Err failed")
	}

	r, _ = MapInt64Uint16Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint16Err failed")
	}

	_, err := MapInt64Uint16Err(plusOneInt64Uint16Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapInt64Uint8Err(plusOneInt64Uint8Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint8Err failed")
	}

	r, _ = MapInt64Uint8Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint8Err failed")
	}

	_, err := MapInt64Uint8Err(plusOneInt64Uint8Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapInt64StrErr(someLogicInt64StrErr, []int64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Str failed")
	}

	r, _ = MapInt64StrErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64StrErr failed")
	}

	_, err := MapInt64StrErr(someLogicInt64StrErr, []int64{0})
	if err == nil {
		t.Errorf("MapInt64StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt64StrErr(num int64) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapInt64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapInt64BoolErr(someLogicInt64BoolErr, []int64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64BoolErr failed")
	}

	r, _ = MapInt64BoolErr(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64BoolErr failed")
	}

	_, err := MapInt64BoolErr(someLogicInt64BoolErr, []int64{10, 3})
	if err == nil {
		t.Errorf("MapInt64BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapInt64Float32Err(plusOneInt64Float32Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Float32Err failed")
	}

	r, _ = MapInt64Float32Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Float32Err failed")
	}

	_, err := MapInt64Float32Err(plusOneInt64Float32Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapInt64Float64Err(plusOneInt64Float64Err, []int64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Float64Err failed")
	}

	r, _ = MapInt64Float64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Float64Err failed")
	}

	_, err := MapInt64Float64Err(plusOneInt64Float64Err, []int64{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt64Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapInt32IntErr(plusOneInt32IntErr, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32IntErr failed")
	}

	r, _ = MapInt32IntErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32IntErr failed")
	}

	_, err := MapInt32IntErr(plusOneInt32IntErr, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapInt32Int64Err(plusOneInt32Int64Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int64Err failed")
	}

	r, _ = MapInt32Int64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int64Err failed")
	}

	_, err := MapInt32Int64Err(plusOneInt32Int64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapInt32Int16Err(plusOneInt32Int16Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int16Err failed")
	}

	r, _ = MapInt32Int16Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int16Err failed")
	}

	_, err := MapInt32Int16Err(plusOneInt32Int16Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapInt32Int8Err(plusOneInt32Int8Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int8Err failed")
	}

	r, _ = MapInt32Int8Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int8Err failed")
	}

	_, err := MapInt32Int8Err(plusOneInt32Int8Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapInt32UintErr(plusOneInt32UintErr, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32UintErr failed")
	}

	r, _ = MapInt32UintErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32UintErr failed")
	}

	_, err := MapInt32UintErr(plusOneInt32UintErr, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapInt32Uint64Err(plusOneInt32Uint64Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint64Err failed")
	}

	r, _ = MapInt32Uint64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint64Err failed")
	}

	_, err := MapInt32Uint64Err(plusOneInt32Uint64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapInt32Uint32Err(plusOneInt32Uint32Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint32Err failed")
	}

	r, _ = MapInt32Uint32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint32Err failed")
	}

	_, err := MapInt32Uint32Err(plusOneInt32Uint32Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapInt32Uint16Err(plusOneInt32Uint16Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint16Err failed")
	}

	r, _ = MapInt32Uint16Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint16Err failed")
	}

	_, err := MapInt32Uint16Err(plusOneInt32Uint16Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapInt32Uint8Err(plusOneInt32Uint8Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint8Err failed")
	}

	r, _ = MapInt32Uint8Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint8Err failed")
	}

	_, err := MapInt32Uint8Err(plusOneInt32Uint8Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapInt32StrErr(someLogicInt32StrErr, []int32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Str failed")
	}

	r, _ = MapInt32StrErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32StrErr failed")
	}

	_, err := MapInt32StrErr(someLogicInt32StrErr, []int32{0})
	if err == nil {
		t.Errorf("MapInt32StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt32StrErr(num int32) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapInt32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapInt32BoolErr(someLogicInt32BoolErr, []int32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32BoolErr failed")
	}

	r, _ = MapInt32BoolErr(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32BoolErr failed")
	}

	_, err := MapInt32BoolErr(someLogicInt32BoolErr, []int32{10, 3})
	if err == nil {
		t.Errorf("MapInt32BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapInt32Float32Err(plusOneInt32Float32Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Float32Err failed")
	}

	r, _ = MapInt32Float32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Float32Err failed")
	}

	_, err := MapInt32Float32Err(plusOneInt32Float32Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapInt32Float64Err(plusOneInt32Float64Err, []int32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Float64Err failed")
	}

	r, _ = MapInt32Float64Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Float64Err failed")
	}

	_, err := MapInt32Float64Err(plusOneInt32Float64Err, []int32{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapInt16IntErr(plusOneInt16IntErr, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16IntErr failed")
	}

	r, _ = MapInt16IntErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16IntErr failed")
	}

	_, err := MapInt16IntErr(plusOneInt16IntErr, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapInt16Int64Err(plusOneInt16Int64Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int64Err failed")
	}

	r, _ = MapInt16Int64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int64Err failed")
	}

	_, err := MapInt16Int64Err(plusOneInt16Int64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapInt16Int32Err(plusOneInt16Int32Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int32Err failed")
	}

	r, _ = MapInt16Int32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int32Err failed")
	}

	_, err := MapInt16Int32Err(plusOneInt16Int32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapInt16Int8Err(plusOneInt16Int8Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int8Err failed")
	}

	r, _ = MapInt16Int8Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int8Err failed")
	}

	_, err := MapInt16Int8Err(plusOneInt16Int8Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapInt16UintErr(plusOneInt16UintErr, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16UintErr failed")
	}

	r, _ = MapInt16UintErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16UintErr failed")
	}

	_, err := MapInt16UintErr(plusOneInt16UintErr, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapInt16Uint64Err(plusOneInt16Uint64Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint64Err failed")
	}

	r, _ = MapInt16Uint64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint64Err failed")
	}

	_, err := MapInt16Uint64Err(plusOneInt16Uint64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapInt16Uint32Err(plusOneInt16Uint32Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint32Err failed")
	}

	r, _ = MapInt16Uint32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint32Err failed")
	}

	_, err := MapInt16Uint32Err(plusOneInt16Uint32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapInt16Uint16Err(plusOneInt16Uint16Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint16Err failed")
	}

	r, _ = MapInt16Uint16Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint16Err failed")
	}

	_, err := MapInt16Uint16Err(plusOneInt16Uint16Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapInt16Uint8Err(plusOneInt16Uint8Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint8Err failed")
	}

	r, _ = MapInt16Uint8Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint8Err failed")
	}

	_, err := MapInt16Uint8Err(plusOneInt16Uint8Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapInt16StrErr(someLogicInt16StrErr, []int16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Str failed")
	}

	r, _ = MapInt16StrErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16StrErr failed")
	}

	_, err := MapInt16StrErr(someLogicInt16StrErr, []int16{0})
	if err == nil {
		t.Errorf("MapInt16StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt16StrErr(num int16) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapInt16BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapInt16BoolErr(someLogicInt16BoolErr, []int16{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16BoolErr failed")
	}

	r, _ = MapInt16BoolErr(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16BoolErr failed")
	}

	_, err := MapInt16BoolErr(someLogicInt16BoolErr, []int16{10, 3})
	if err == nil {
		t.Errorf("MapInt16BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapInt16Float32Err(plusOneInt16Float32Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Float32Err failed")
	}

	r, _ = MapInt16Float32Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Float32Err failed")
	}

	_, err := MapInt16Float32Err(plusOneInt16Float32Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapInt16Float64Err(plusOneInt16Float64Err, []int16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Float64Err failed")
	}

	r, _ = MapInt16Float64Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Float64Err failed")
	}

	_, err := MapInt16Float64Err(plusOneInt16Float64Err, []int16{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt16Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapInt8IntErr(plusOneInt8IntErr, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8IntErr failed")
	}

	r, _ = MapInt8IntErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8IntErr failed")
	}

	_, err := MapInt8IntErr(plusOneInt8IntErr, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapInt8Int64Err(plusOneInt8Int64Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int64Err failed")
	}

	r, _ = MapInt8Int64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int64Err failed")
	}

	_, err := MapInt8Int64Err(plusOneInt8Int64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapInt8Int32Err(plusOneInt8Int32Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int32Err failed")
	}

	r, _ = MapInt8Int32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int32Err failed")
	}

	_, err := MapInt8Int32Err(plusOneInt8Int32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapInt8Int16Err(plusOneInt8Int16Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int16Err failed")
	}

	r, _ = MapInt8Int16Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int16Err failed")
	}

	_, err := MapInt8Int16Err(plusOneInt8Int16Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapInt8UintErr(plusOneInt8UintErr, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8UintErr failed")
	}

	r, _ = MapInt8UintErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8UintErr failed")
	}

	_, err := MapInt8UintErr(plusOneInt8UintErr, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapInt8Uint64Err(plusOneInt8Uint64Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint64Err failed")
	}

	r, _ = MapInt8Uint64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint64Err failed")
	}

	_, err := MapInt8Uint64Err(plusOneInt8Uint64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapInt8Uint32Err(plusOneInt8Uint32Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint32Err failed")
	}

	r, _ = MapInt8Uint32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint32Err failed")
	}

	_, err := MapInt8Uint32Err(plusOneInt8Uint32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapInt8Uint16Err(plusOneInt8Uint16Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint16Err failed")
	}

	r, _ = MapInt8Uint16Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint16Err failed")
	}

	_, err := MapInt8Uint16Err(plusOneInt8Uint16Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapInt8Uint8Err(plusOneInt8Uint8Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint8Err failed")
	}

	r, _ = MapInt8Uint8Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint8Err failed")
	}

	_, err := MapInt8Uint8Err(plusOneInt8Uint8Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapInt8StrErr(someLogicInt8StrErr, []int8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Str failed")
	}

	r, _ = MapInt8StrErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8StrErr failed")
	}

	_, err := MapInt8StrErr(someLogicInt8StrErr, []int8{0})
	if err == nil {
		t.Errorf("MapInt8StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt8StrErr(num int8) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapInt8BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapInt8BoolErr(someLogicInt8BoolErr, []int8{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8BoolErr failed")
	}

	r, _ = MapInt8BoolErr(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8BoolErr failed")
	}

	_, err := MapInt8BoolErr(someLogicInt8BoolErr, []int8{10, 3})
	if err == nil {
		t.Errorf("MapInt8BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapInt8Float32Err(plusOneInt8Float32Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Float32Err failed")
	}

	r, _ = MapInt8Float32Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Float32Err failed")
	}

	_, err := MapInt8Float32Err(plusOneInt8Float32Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapInt8Float64Err(plusOneInt8Float64Err, []int8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Float64Err failed")
	}

	r, _ = MapInt8Float64Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Float64Err failed")
	}

	_, err := MapInt8Float64Err(plusOneInt8Float64Err, []int8{1, 2, 3})
	if err == nil {
		t.Errorf("MapInt8Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintIntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapUintIntErr(plusOneUintIntErr, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintIntErr failed")
	}

	r, _ = MapUintIntErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintIntErr failed")
	}

	_, err := MapUintIntErr(plusOneUintIntErr, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintIntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapUintInt64Err(plusOneUintInt64Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt64Err failed")
	}

	r, _ = MapUintInt64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt64Err failed")
	}

	_, err := MapUintInt64Err(plusOneUintInt64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapUintInt32Err(plusOneUintInt32Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt32Err failed")
	}

	r, _ = MapUintInt32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt32Err failed")
	}

	_, err := MapUintInt32Err(plusOneUintInt32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapUintInt16Err(plusOneUintInt16Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt16Err failed")
	}

	r, _ = MapUintInt16Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt16Err failed")
	}

	_, err := MapUintInt16Err(plusOneUintInt16Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapUintInt8Err(plusOneUintInt8Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt8Err failed")
	}

	r, _ = MapUintInt8Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt8Err failed")
	}

	_, err := MapUintInt8Err(plusOneUintInt8Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapUintUint64Err(plusOneUintUint64Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint64Err failed")
	}

	r, _ = MapUintUint64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint64Err failed")
	}

	_, err := MapUintUint64Err(plusOneUintUint64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapUintUint32Err(plusOneUintUint32Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint32Err failed")
	}

	r, _ = MapUintUint32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint32Err failed")
	}

	_, err := MapUintUint32Err(plusOneUintUint32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapUintUint16Err(plusOneUintUint16Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint16Err failed")
	}

	r, _ = MapUintUint16Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint16Err failed")
	}

	_, err := MapUintUint16Err(plusOneUintUint16Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapUintUint8Err(plusOneUintUint8Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint8Err failed")
	}

	r, _ = MapUintUint8Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint8Err failed")
	}

	_, err := MapUintUint8Err(plusOneUintUint8Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapUintStrErr(someLogicUintStrErr, []uint{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUintStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintStr failed")
	}

	r, _ = MapUintStrErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintStrErr failed")
	}

	_, err := MapUintStrErr(someLogicUintStrErr, []uint{0})
	if err == nil {
		t.Errorf("MapUintStrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUintStrErr(num uint) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapUintBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapUintBoolErr(someLogicUintBoolErr, []uint{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUintBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintBoolErr failed")
	}

	r, _ = MapUintBoolErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintBoolErr failed")
	}

	_, err := MapUintBoolErr(someLogicUintBoolErr, []uint{10, 3})
	if err == nil {
		t.Errorf("MapUintBoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapUintFloat32Err(plusOneUintFloat32Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintFloat32Err failed")
	}

	r, _ = MapUintFloat32Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintFloat32Err failed")
	}

	_, err := MapUintFloat32Err(plusOneUintFloat32Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapUintFloat64Err(plusOneUintFloat64Err, []uint{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintFloat64Err failed")
	}

	r, _ = MapUintFloat64Err(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("MapUintFloat64Err failed")
	}

	_, err := MapUintFloat64Err(plusOneUintFloat64Err, []uint{1, 2, 3})
	if err == nil {
		t.Errorf("MapUintFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapUint64IntErr(plusOneUint64IntErr, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64IntErr failed")
	}

	r, _ = MapUint64IntErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64IntErr failed")
	}

	_, err := MapUint64IntErr(plusOneUint64IntErr, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapUint64Int64Err(plusOneUint64Int64Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int64Err failed")
	}

	r, _ = MapUint64Int64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int64Err failed")
	}

	_, err := MapUint64Int64Err(plusOneUint64Int64Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapUint64Int32Err(plusOneUint64Int32Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int32Err failed")
	}

	r, _ = MapUint64Int32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int32Err failed")
	}

	_, err := MapUint64Int32Err(plusOneUint64Int32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapUint64Int16Err(plusOneUint64Int16Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int16Err failed")
	}

	r, _ = MapUint64Int16Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int16Err failed")
	}

	_, err := MapUint64Int16Err(plusOneUint64Int16Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapUint64Int8Err(plusOneUint64Int8Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int8Err failed")
	}

	r, _ = MapUint64Int8Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int8Err failed")
	}

	_, err := MapUint64Int8Err(plusOneUint64Int8Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapUint64UintErr(plusOneUint64UintErr, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64UintErr failed")
	}

	r, _ = MapUint64UintErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64UintErr failed")
	}

	_, err := MapUint64UintErr(plusOneUint64UintErr, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapUint64Uint32Err(plusOneUint64Uint32Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint32Err failed")
	}

	r, _ = MapUint64Uint32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint32Err failed")
	}

	_, err := MapUint64Uint32Err(plusOneUint64Uint32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapUint64Uint16Err(plusOneUint64Uint16Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint16Err failed")
	}

	r, _ = MapUint64Uint16Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint16Err failed")
	}

	_, err := MapUint64Uint16Err(plusOneUint64Uint16Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapUint64Uint8Err(plusOneUint64Uint8Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint8Err failed")
	}

	r, _ = MapUint64Uint8Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint8Err failed")
	}

	_, err := MapUint64Uint8Err(plusOneUint64Uint8Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapUint64StrErr(someLogicUint64StrErr, []uint64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Str failed")
	}

	r, _ = MapUint64StrErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64StrErr failed")
	}

	_, err := MapUint64StrErr(someLogicUint64StrErr, []uint64{0})
	if err == nil {
		t.Errorf("MapUint64StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint64StrErr(num uint64) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapUint64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapUint64BoolErr(someLogicUint64BoolErr, []uint64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64BoolErr failed")
	}

	r, _ = MapUint64BoolErr(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64BoolErr failed")
	}

	_, err := MapUint64BoolErr(someLogicUint64BoolErr, []uint64{10, 3})
	if err == nil {
		t.Errorf("MapUint64BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapUint64Float32Err(plusOneUint64Float32Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Float32Err failed")
	}

	r, _ = MapUint64Float32Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Float32Err failed")
	}

	_, err := MapUint64Float32Err(plusOneUint64Float32Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapUint64Float64Err(plusOneUint64Float64Err, []uint64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Float64Err failed")
	}

	r, _ = MapUint64Float64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Float64Err failed")
	}

	_, err := MapUint64Float64Err(plusOneUint64Float64Err, []uint64{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint64Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapUint32IntErr(plusOneUint32IntErr, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32IntErr failed")
	}

	r, _ = MapUint32IntErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32IntErr failed")
	}

	_, err := MapUint32IntErr(plusOneUint32IntErr, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapUint32Int64Err(plusOneUint32Int64Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int64Err failed")
	}

	r, _ = MapUint32Int64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int64Err failed")
	}

	_, err := MapUint32Int64Err(plusOneUint32Int64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapUint32Int32Err(plusOneUint32Int32Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int32Err failed")
	}

	r, _ = MapUint32Int32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int32Err failed")
	}

	_, err := MapUint32Int32Err(plusOneUint32Int32Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapUint32Int16Err(plusOneUint32Int16Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int16Err failed")
	}

	r, _ = MapUint32Int16Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int16Err failed")
	}

	_, err := MapUint32Int16Err(plusOneUint32Int16Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapUint32Int8Err(plusOneUint32Int8Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int8Err failed")
	}

	r, _ = MapUint32Int8Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int8Err failed")
	}

	_, err := MapUint32Int8Err(plusOneUint32Int8Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapUint32UintErr(plusOneUint32UintErr, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32UintErr failed")
	}

	r, _ = MapUint32UintErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32UintErr failed")
	}

	_, err := MapUint32UintErr(plusOneUint32UintErr, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapUint32Uint64Err(plusOneUint32Uint64Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint64Err failed")
	}

	r, _ = MapUint32Uint64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint64Err failed")
	}

	_, err := MapUint32Uint64Err(plusOneUint32Uint64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapUint32Uint16Err(plusOneUint32Uint16Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint16Err failed")
	}

	r, _ = MapUint32Uint16Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint16Err failed")
	}

	_, err := MapUint32Uint16Err(plusOneUint32Uint16Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapUint32Uint8Err(plusOneUint32Uint8Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint8Err failed")
	}

	r, _ = MapUint32Uint8Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint8Err failed")
	}

	_, err := MapUint32Uint8Err(plusOneUint32Uint8Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapUint32StrErr(someLogicUint32StrErr, []uint32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Str failed")
	}

	r, _ = MapUint32StrErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32StrErr failed")
	}

	_, err := MapUint32StrErr(someLogicUint32StrErr, []uint32{0})
	if err == nil {
		t.Errorf("MapUint32StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint32StrErr(num uint32) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapUint32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapUint32BoolErr(someLogicUint32BoolErr, []uint32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32BoolErr failed")
	}

	r, _ = MapUint32BoolErr(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32BoolErr failed")
	}

	_, err := MapUint32BoolErr(someLogicUint32BoolErr, []uint32{10, 3})
	if err == nil {
		t.Errorf("MapUint32BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapUint32Float32Err(plusOneUint32Float32Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Float32Err failed")
	}

	r, _ = MapUint32Float32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Float32Err failed")
	}

	_, err := MapUint32Float32Err(plusOneUint32Float32Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapUint32Float64Err(plusOneUint32Float64Err, []uint32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Float64Err failed")
	}

	r, _ = MapUint32Float64Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Float64Err failed")
	}

	_, err := MapUint32Float64Err(plusOneUint32Float64Err, []uint32{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapUint16IntErr(plusOneUint16IntErr, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16IntErr failed")
	}

	r, _ = MapUint16IntErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16IntErr failed")
	}

	_, err := MapUint16IntErr(plusOneUint16IntErr, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapUint16Int64Err(plusOneUint16Int64Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int64Err failed")
	}

	r, _ = MapUint16Int64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int64Err failed")
	}

	_, err := MapUint16Int64Err(plusOneUint16Int64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapUint16Int32Err(plusOneUint16Int32Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int32Err failed")
	}

	r, _ = MapUint16Int32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int32Err failed")
	}

	_, err := MapUint16Int32Err(plusOneUint16Int32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapUint16Int16Err(plusOneUint16Int16Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int16Err failed")
	}

	r, _ = MapUint16Int16Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int16Err failed")
	}

	_, err := MapUint16Int16Err(plusOneUint16Int16Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapUint16Int8Err(plusOneUint16Int8Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int8Err failed")
	}

	r, _ = MapUint16Int8Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int8Err failed")
	}

	_, err := MapUint16Int8Err(plusOneUint16Int8Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapUint16UintErr(plusOneUint16UintErr, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16UintErr failed")
	}

	r, _ = MapUint16UintErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16UintErr failed")
	}

	_, err := MapUint16UintErr(plusOneUint16UintErr, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapUint16Uint64Err(plusOneUint16Uint64Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint64Err failed")
	}

	r, _ = MapUint16Uint64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint64Err failed")
	}

	_, err := MapUint16Uint64Err(plusOneUint16Uint64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapUint16Uint32Err(plusOneUint16Uint32Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint32Err failed")
	}

	r, _ = MapUint16Uint32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint32Err failed")
	}

	_, err := MapUint16Uint32Err(plusOneUint16Uint32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapUint16Uint8Err(plusOneUint16Uint8Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint8Err failed")
	}

	r, _ = MapUint16Uint8Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint8Err failed")
	}

	_, err := MapUint16Uint8Err(plusOneUint16Uint8Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapUint16StrErr(someLogicUint16StrErr, []uint16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Str failed")
	}

	r, _ = MapUint16StrErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16StrErr failed")
	}

	_, err := MapUint16StrErr(someLogicUint16StrErr, []uint16{0})
	if err == nil {
		t.Errorf("MapUint16StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint16StrErr(num uint16) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapUint16BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapUint16BoolErr(someLogicUint16BoolErr, []uint16{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16BoolErr failed")
	}

	r, _ = MapUint16BoolErr(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16BoolErr failed")
	}

	_, err := MapUint16BoolErr(someLogicUint16BoolErr, []uint16{10, 3})
	if err == nil {
		t.Errorf("MapUint16BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapUint16Float32Err(plusOneUint16Float32Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Float32Err failed")
	}

	r, _ = MapUint16Float32Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Float32Err failed")
	}

	_, err := MapUint16Float32Err(plusOneUint16Float32Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapUint16Float64Err(plusOneUint16Float64Err, []uint16{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Float64Err failed")
	}

	r, _ = MapUint16Float64Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Float64Err failed")
	}

	_, err := MapUint16Float64Err(plusOneUint16Float64Err, []uint16{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint16Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapUint8IntErr(plusOneUint8IntErr, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8IntErr failed")
	}

	r, _ = MapUint8IntErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8IntErr failed")
	}

	_, err := MapUint8IntErr(plusOneUint8IntErr, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapUint8Int64Err(plusOneUint8Int64Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int64Err failed")
	}

	r, _ = MapUint8Int64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int64Err failed")
	}

	_, err := MapUint8Int64Err(plusOneUint8Int64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapUint8Int32Err(plusOneUint8Int32Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int32Err failed")
	}

	r, _ = MapUint8Int32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int32Err failed")
	}

	_, err := MapUint8Int32Err(plusOneUint8Int32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapUint8Int16Err(plusOneUint8Int16Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int16Err failed")
	}

	r, _ = MapUint8Int16Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int16Err failed")
	}

	_, err := MapUint8Int16Err(plusOneUint8Int16Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapUint8Int8Err(plusOneUint8Int8Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int8Err failed")
	}

	r, _ = MapUint8Int8Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int8Err failed")
	}

	_, err := MapUint8Int8Err(plusOneUint8Int8Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapUint8UintErr(plusOneUint8UintErr, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8UintErr failed")
	}

	r, _ = MapUint8UintErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8UintErr failed")
	}

	_, err := MapUint8UintErr(plusOneUint8UintErr, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapUint8Uint64Err(plusOneUint8Uint64Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint64Err failed")
	}

	r, _ = MapUint8Uint64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint64Err failed")
	}

	_, err := MapUint8Uint64Err(plusOneUint8Uint64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapUint8Uint32Err(plusOneUint8Uint32Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint32Err failed")
	}

	r, _ = MapUint8Uint32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint32Err failed")
	}

	_, err := MapUint8Uint32Err(plusOneUint8Uint32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapUint8Uint16Err(plusOneUint8Uint16Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint16Err failed")
	}

	r, _ = MapUint8Uint16Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint16Err failed")
	}

	_, err := MapUint8Uint16Err(plusOneUint8Uint16Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapUint8StrErr(someLogicUint8StrErr, []uint8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Str failed")
	}

	r, _ = MapUint8StrErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8StrErr failed")
	}

	_, err := MapUint8StrErr(someLogicUint8StrErr, []uint8{0})
	if err == nil {
		t.Errorf("MapUint8StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint8StrErr(num uint8) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapUint8BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapUint8BoolErr(someLogicUint8BoolErr, []uint8{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8BoolErr failed")
	}

	r, _ = MapUint8BoolErr(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8BoolErr failed")
	}

	_, err := MapUint8BoolErr(someLogicUint8BoolErr, []uint8{10, 3})
	if err == nil {
		t.Errorf("MapUint8BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapUint8Float32Err(plusOneUint8Float32Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Float32Err failed")
	}

	r, _ = MapUint8Float32Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Float32Err failed")
	}

	_, err := MapUint8Float32Err(plusOneUint8Float32Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapUint8Float64Err(plusOneUint8Float64Err, []uint8{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Float64Err failed")
	}

	r, _ = MapUint8Float64Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Float64Err failed")
	}

	_, err := MapUint8Float64Err(plusOneUint8Float64Err, []uint8{1, 2, 3})
	if err == nil {
		t.Errorf("MapUint8Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrIntErr(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10}
	newList, _ := MapStrIntErr(someLogicStrIntErr, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrIntErr failed")
	}

	r, _ = MapStrIntErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrIntErr failed")
	}

	_, err := MapStrIntErr(someLogicStrIntErr, []string{"0"})
	if err == nil {
		t.Errorf("MapStrIntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrIntErr(num string) (int, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return int(10), nil
	}
	return 0, nil
}

func TestMapStrInt64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10}
	newList, _ := MapStrInt64Err(someLogicStrInt64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt64Err failed")
	}

	r, _ = MapStrInt64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt64Err failed")
	}

	_, err := MapStrInt64Err(someLogicStrInt64Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrInt64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt64Err(num string) (int64, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return int64(10), nil
	}
	return 0, nil
}

func TestMapStrInt32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10}
	newList, _ := MapStrInt32Err(someLogicStrInt32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt32Err failed")
	}

	r, _ = MapStrInt32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt32Err failed")
	}

	_, err := MapStrInt32Err(someLogicStrInt32Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrInt32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt32Err(num string) (int32, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return int32(10), nil
	}
	return 0, nil
}

func TestMapStrInt16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10}
	newList, _ := MapStrInt16Err(someLogicStrInt16Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt16Err failed")
	}

	r, _ = MapStrInt16Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt16Err failed")
	}

	_, err := MapStrInt16Err(someLogicStrInt16Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrInt16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt16Err(num string) (int16, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return int16(10), nil
	}
	return 0, nil
}

func TestMapStrInt8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10}
	newList, _ := MapStrInt8Err(someLogicStrInt8Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt8Err failed")
	}

	r, _ = MapStrInt8Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt8Err failed")
	}

	_, err := MapStrInt8Err(someLogicStrInt8Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrInt8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt8Err(num string) (int8, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return int8(10), nil
	}
	return 0, nil
}

func TestMapStrUintErr(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10}
	newList, _ := MapStrUintErr(someLogicStrUintErr, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUintErr failed")
	}

	r, _ = MapStrUintErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrUintErr failed")
	}

	_, err := MapStrUintErr(someLogicStrUintErr, []string{"0"})
	if err == nil {
		t.Errorf("MapStrUintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUintErr(num string) (uint, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return uint(10), nil
	}
	return 0, nil
}

func TestMapStrUint64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10}
	newList, _ := MapStrUint64Err(someLogicStrUint64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint64Err failed")
	}

	r, _ = MapStrUint64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint64Err failed")
	}

	_, err := MapStrUint64Err(someLogicStrUint64Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrUint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint64Err(num string) (uint64, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return uint64(10), nil
	}
	return 0, nil
}

func TestMapStrUint32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10}
	newList, _ := MapStrUint32Err(someLogicStrUint32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint32Err failed")
	}

	r, _ = MapStrUint32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint32Err failed")
	}

	_, err := MapStrUint32Err(someLogicStrUint32Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrUint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint32Err(num string) (uint32, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return uint32(10), nil
	}
	return 0, nil
}

func TestMapStrUint16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10}
	newList, _ := MapStrUint16Err(someLogicStrUint16Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint16Err failed")
	}

	r, _ = MapStrUint16Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint16Err failed")
	}

	_, err := MapStrUint16Err(someLogicStrUint16Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrUint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint16Err(num string) (uint16, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return uint16(10), nil
	}
	return 0, nil
}

func TestMapStrUint8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10}
	newList, _ := MapStrUint8Err(someLogicStrUint8Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint8Err failed")
	}

	r, _ = MapStrUint8Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint8Err failed")
	}

	_, err := MapStrUint8Err(someLogicStrUint8Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrUint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint8Err(num string) (uint8, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return uint8(10), nil
	}
	return 0, nil
}

func TestMapStrBoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapStrBoolErr(someLogicStrBoolErr, []string{"10", "0"})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapStrBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrBool failed")
	}

	r, _ = MapStrBoolErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrBoolErr failed")
	}

	_, err := MapStrBoolErr(someLogicStrBoolErr, []string{"10", "3"})
	if err == nil {
		t.Errorf("MapStrBoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrFloat32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10}
	newList, _ := MapStrFloat32Err(someLogicStrFloat32Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrFloat32Err failed")
	}

	r, _ = MapStrFloat32Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrFloat32Err failed")
	}

	_, err := MapStrFloat32Err(someLogicStrFloat32Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrFloat32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat32Err(num string) (float32, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return float32(10), nil
	}
	return 0, nil
}

func TestMapStrFloat64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10}
	newList, _ := MapStrFloat64Err(someLogicStrFloat64Err, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrFloat64Err failed")
	}

	r, _ = MapStrFloat64Err(nil, []string{})
	if len(r) > 0 {
		t.Errorf("MapStrFloat64Err failed")
	}

	_, err := MapStrFloat64Err(someLogicStrFloat64Err, []string{"0"})
	if err == nil {
		t.Errorf("MapStrFloat64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat64Err(num string) (float64, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return float64(10), nil
	}
	return 0, nil
}

func TestMapBoolIntErr(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10, 0}
	newList, _ := MapBoolIntErr(someLogicBoolIntErr, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt failed")
	}

	r, _ = MapBoolIntErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolIntErr failed")
	}
	_, err := MapBoolIntErr(someLogicBoolIntErr, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolIntErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10, 0}
	newList, _ := MapBoolInt64Err(someLogicBoolInt64Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt64 failed")
	}

	r, _ = MapBoolInt64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt64Err failed")
	}
	_, err := MapBoolInt64Err(someLogicBoolInt64Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolInt64Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10, 0}
	newList, _ := MapBoolInt32Err(someLogicBoolInt32Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt32 failed")
	}

	r, _ = MapBoolInt32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt32Err failed")
	}
	_, err := MapBoolInt32Err(someLogicBoolInt32Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolInt32Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10, 0}
	newList, _ := MapBoolInt16Err(someLogicBoolInt16Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt16 failed")
	}

	r, _ = MapBoolInt16Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt16Err failed")
	}
	_, err := MapBoolInt16Err(someLogicBoolInt16Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolInt16Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10, 0}
	newList, _ := MapBoolInt8Err(someLogicBoolInt8Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt8 failed")
	}

	r, _ = MapBoolInt8Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt8Err failed")
	}
	_, err := MapBoolInt8Err(someLogicBoolInt8Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolInt8Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUintErr(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10, 0}
	newList, _ := MapBoolUintErr(someLogicBoolUintErr, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint failed")
	}

	r, _ = MapBoolUintErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUintErr failed")
	}
	_, err := MapBoolUintErr(someLogicBoolUintErr, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolUintErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10, 0}
	newList, _ := MapBoolUint64Err(someLogicBoolUint64Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint64 failed")
	}

	r, _ = MapBoolUint64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint64Err failed")
	}
	_, err := MapBoolUint64Err(someLogicBoolUint64Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolUint64Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10, 0}
	newList, _ := MapBoolUint32Err(someLogicBoolUint32Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint32 failed")
	}

	r, _ = MapBoolUint32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint32Err failed")
	}
	_, err := MapBoolUint32Err(someLogicBoolUint32Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolUint32Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint16Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10, 0}
	newList, _ := MapBoolUint16Err(someLogicBoolUint16Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint16 failed")
	}

	r, _ = MapBoolUint16Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint16Err failed")
	}
	_, err := MapBoolUint16Err(someLogicBoolUint16Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolUint16Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint8Err(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10, 0}
	newList, _ := MapBoolUint8Err(someLogicBoolUint8Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint8 failed")
	}

	r, _ = MapBoolUint8Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint8Err failed")
	}
	_, err := MapBoolUint8Err(someLogicBoolUint8Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolUint8Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolStrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10", "0"}
	newList, _ := MapBoolStrErr(someLogicBoolStrErr, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolStrErr failed")
	}

	r, _ = MapBoolStrErr(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolStrErr failed")
	}
	_, err := MapBoolStrErr(someLogicBoolStrErr, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolStrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat32Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10, 0}
	newList, _ := MapBoolFloat32Err(someLogicBoolFloat32Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolFloat32 failed")
	}

	r, _ = MapBoolFloat32Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolFloat32Err failed")
	}
	_, err := MapBoolFloat32Err(someLogicBoolFloat32Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolFloat32Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat64Err(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10, 0}
	newList, _ := MapBoolFloat64Err(someLogicBoolFloat64Err, []bool{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolFloat64 failed")
	}

	r, _ = MapBoolFloat64Err(nil, []bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolFloat64Err failed")
	}
	_, err := MapBoolFloat64Err(someLogicBoolFloat64Err, []bool{true, false})
	if err == nil {
		t.Errorf("MapBoolFloat64Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapFloat32IntErr(plusOneFloat32IntErr, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32IntErr failed")
	}

	r, _ = MapFloat32IntErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32IntErr failed")
	}

	_, err := MapFloat32IntErr(plusOneFloat32IntErr, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapFloat32Int64Err(plusOneFloat32Int64Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int64Err failed")
	}

	r, _ = MapFloat32Int64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int64Err failed")
	}

	_, err := MapFloat32Int64Err(plusOneFloat32Int64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapFloat32Int32Err(plusOneFloat32Int32Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int32Err failed")
	}

	r, _ = MapFloat32Int32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int32Err failed")
	}

	_, err := MapFloat32Int32Err(plusOneFloat32Int32Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapFloat32Int16Err(plusOneFloat32Int16Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int16Err failed")
	}

	r, _ = MapFloat32Int16Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int16Err failed")
	}

	_, err := MapFloat32Int16Err(plusOneFloat32Int16Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapFloat32Int8Err(plusOneFloat32Int8Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int8Err failed")
	}

	r, _ = MapFloat32Int8Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int8Err failed")
	}

	_, err := MapFloat32Int8Err(plusOneFloat32Int8Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapFloat32UintErr(plusOneFloat32UintErr, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32UintErr failed")
	}

	r, _ = MapFloat32UintErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32UintErr failed")
	}

	_, err := MapFloat32UintErr(plusOneFloat32UintErr, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapFloat32Uint64Err(plusOneFloat32Uint64Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint64Err failed")
	}

	r, _ = MapFloat32Uint64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint64Err failed")
	}

	_, err := MapFloat32Uint64Err(plusOneFloat32Uint64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapFloat32Uint32Err(plusOneFloat32Uint32Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint32Err failed")
	}

	r, _ = MapFloat32Uint32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint32Err failed")
	}

	_, err := MapFloat32Uint32Err(plusOneFloat32Uint32Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapFloat32Uint16Err(plusOneFloat32Uint16Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint16Err failed")
	}

	r, _ = MapFloat32Uint16Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint16Err failed")
	}

	_, err := MapFloat32Uint16Err(plusOneFloat32Uint16Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapFloat32Uint8Err(plusOneFloat32Uint8Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint8Err failed")
	}

	r, _ = MapFloat32Uint8Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint8Err failed")
	}

	_, err := MapFloat32Uint8Err(plusOneFloat32Uint8Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapFloat32StrErr(someLogicFloat32StrErr, []float32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapFloat32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Str failed")
	}

	r, _ = MapFloat32StrErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32StrErr failed")
	}

	_, err := MapFloat32StrErr(someLogicFloat32StrErr, []float32{0})
	if err == nil {
		t.Errorf("MapFloat32StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat32StrErr(num float32) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapFloat32BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapFloat32BoolErr(someLogicFloat32BoolErr, []float32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapFloat32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32BoolErr failed")
	}

	r, _ = MapFloat32BoolErr(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32BoolErr failed")
	}

	_, err := MapFloat32BoolErr(someLogicFloat32BoolErr, []float32{10, 3})
	if err == nil {
		t.Errorf("MapFloat32BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Float64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 6}
	newList, _ := MapFloat32Float64Err(plusOneFloat32Float64Err, []float32{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Float64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Float64Err failed")
	}

	r, _ = MapFloat32Float64Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Float64Err failed")
	}

	_, err := MapFloat32Float64Err(plusOneFloat32Float64Err, []float32{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat32Float64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64IntErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 6}
	newList, _ := MapFloat64IntErr(plusOneFloat64IntErr, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64IntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64IntErr failed")
	}

	r, _ = MapFloat64IntErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64IntErr failed")
	}

	_, err := MapFloat64IntErr(plusOneFloat64IntErr, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64IntErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 6}
	newList, _ := MapFloat64Int64Err(plusOneFloat64Int64Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int64Err failed")
	}

	r, _ = MapFloat64Int64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int64Err failed")
	}

	_, err := MapFloat64Int64Err(plusOneFloat64Int64Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Int64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 6}
	newList, _ := MapFloat64Int32Err(plusOneFloat64Int32Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int32Err failed")
	}

	r, _ = MapFloat64Int32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int32Err failed")
	}

	_, err := MapFloat64Int32Err(plusOneFloat64Int32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Int32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 6}
	newList, _ := MapFloat64Int16Err(plusOneFloat64Int16Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int16Err failed")
	}

	r, _ = MapFloat64Int16Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int16Err failed")
	}

	_, err := MapFloat64Int16Err(plusOneFloat64Int16Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Int16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 6}
	newList, _ := MapFloat64Int8Err(plusOneFloat64Int8Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int8Err failed")
	}

	r, _ = MapFloat64Int8Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int8Err failed")
	}

	_, err := MapFloat64Int8Err(plusOneFloat64Int8Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Int8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64UintErr(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 6}
	newList, _ := MapFloat64UintErr(plusOneFloat64UintErr, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64UintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64UintErr failed")
	}

	r, _ = MapFloat64UintErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64UintErr failed")
	}

	_, err := MapFloat64UintErr(plusOneFloat64UintErr, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64UintErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint64Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 6}
	newList, _ := MapFloat64Uint64Err(plusOneFloat64Uint64Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint64Err failed")
	}

	r, _ = MapFloat64Uint64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint64Err failed")
	}

	_, err := MapFloat64Uint64Err(plusOneFloat64Uint64Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Uint64Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 6}
	newList, _ := MapFloat64Uint32Err(plusOneFloat64Uint32Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint32Err failed")
	}

	r, _ = MapFloat64Uint32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint32Err failed")
	}

	_, err := MapFloat64Uint32Err(plusOneFloat64Uint32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Uint32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint16Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 6}
	newList, _ := MapFloat64Uint16Err(plusOneFloat64Uint16Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint16Err failed")
	}

	r, _ = MapFloat64Uint16Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint16Err failed")
	}

	_, err := MapFloat64Uint16Err(plusOneFloat64Uint16Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Uint16Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint8Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 6}
	newList, _ := MapFloat64Uint8Err(plusOneFloat64Uint8Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint8Err failed")
	}

	r, _ = MapFloat64Uint8Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint8Err failed")
	}

	_, err := MapFloat64Uint8Err(plusOneFloat64Uint8Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Uint8Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64StrErr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList, _ := MapFloat64StrErr(someLogicFloat64StrErr, []float64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapFloat64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64StrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Str failed")
	}

	r, _ = MapFloat64StrErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64StrErr failed")
	}

	_, err := MapFloat64StrErr(someLogicFloat64StrErr, []float64{0})
	if err == nil {
		t.Errorf("MapFloat64StrErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat64StrErr(num float64) (string, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}

func TestMapFloat64BoolErr(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList, _ := MapFloat64BoolErr(someLogicFloat64BoolErr, []float64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapFloat64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64BoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64BoolErr failed")
	}

	r, _ = MapFloat64BoolErr(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64BoolErr failed")
	}

	_, err := MapFloat64BoolErr(someLogicFloat64BoolErr, []float64{10, 3})
	if err == nil {
		t.Errorf("MapFloat64BoolErr failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Float32Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 6}
	newList, _ := MapFloat64Float32Err(plusOneFloat64Float32Err, []float64{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Float32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Float32Err failed")
	}

	r, _ = MapFloat64Float32Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Float32Err failed")
	}

	_, err := MapFloat64Float32Err(plusOneFloat64Float32Err, []float64{1, 2, 3})
	if err == nil {
		t.Errorf("MapFloat64Float32Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
