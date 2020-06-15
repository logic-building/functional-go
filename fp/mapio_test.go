package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntInt64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapIntInt64(plusOneIntInt64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt64(nil, nil)) > 0 {
		t.Errorf("MapIntInt64 failed")
	}

	if len(MapIntInt64(nil, []int{})) > 0 {
		t.Errorf("MapIntInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapIntInt32(plusOneIntInt32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt32(nil, nil)) > 0 {
		t.Errorf("MapIntInt32 failed")
	}

	if len(MapIntInt32(nil, []int{})) > 0 {
		t.Errorf("MapIntInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapIntInt16(plusOneIntInt16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt16(nil, nil)) > 0 {
		t.Errorf("MapIntInt16 failed")
	}

	if len(MapIntInt16(nil, []int{})) > 0 {
		t.Errorf("MapIntInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapIntInt8(plusOneIntInt8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt8(nil, nil)) > 0 {
		t.Errorf("MapIntInt8 failed")
	}

	if len(MapIntInt8(nil, []int{})) > 0 {
		t.Errorf("MapIntInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapIntUint(plusOneIntUint, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint(nil, nil)) > 0 {
		t.Errorf("MapIntUint failed")
	}

	if len(MapIntUint(nil, []int{})) > 0 {
		t.Errorf("MapIntUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapIntUint64(plusOneIntUint64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint64(nil, nil)) > 0 {
		t.Errorf("MapIntUint64 failed")
	}

	if len(MapIntUint64(nil, []int{})) > 0 {
		t.Errorf("MapIntUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapIntUint32(plusOneIntUint32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint32(nil, nil)) > 0 {
		t.Errorf("MapIntUint32 failed")
	}

	if len(MapIntUint32(nil, []int{})) > 0 {
		t.Errorf("MapIntUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapIntUint16(plusOneIntUint16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint16(nil, nil)) > 0 {
		t.Errorf("MapIntUint16 failed")
	}

	if len(MapIntUint16(nil, []int{})) > 0 {
		t.Errorf("MapIntUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapIntUint8(plusOneIntUint8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint8(nil, nil)) > 0 {
		t.Errorf("MapIntUint8 failed")
	}

	if len(MapIntUint8(nil, []int{})) > 0 {
		t.Errorf("MapIntUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapIntStr(someLogicIntStr, []int{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapIntStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntStr(nil, nil)) > 0 {
		t.Errorf("MapIntStr failed")
	}

	if len(MapIntStr(nil, []int{})) > 0 {
		t.Errorf("MapIntStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapIntBool(someLogicIntBool, []int{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapIntBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntBool(nil, nil)) > 0 {
		t.Errorf("MapIntBool failed")
	}

	if len(MapIntBool(nil, []int{})) > 0 {
		t.Errorf("MapIntBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapIntFloat32(plusOneIntFloat32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntFloat32(nil, nil)) > 0 {
		t.Errorf("MapIntFloat32 failed")
	}

	if len(MapIntFloat32(nil, []int{})) > 0 {
		t.Errorf("MapIntFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapIntFloat64(plusOneIntFloat64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapIntFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntFloat64(nil, nil)) > 0 {
		t.Errorf("MapIntFloat64 failed")
	}

	if len(MapIntFloat64(nil, []int{})) > 0 {
		t.Errorf("MapIntFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapInt64Int(plusOneInt64Int, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int(nil, nil)) > 0 {
		t.Errorf("MapInt64Int failed")
	}

	if len(MapInt64Int(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapInt64Int32(plusOneInt64Int32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int32(nil, nil)) > 0 {
		t.Errorf("MapInt64Int32 failed")
	}

	if len(MapInt64Int32(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapInt64Int16(plusOneInt64Int16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int16(nil, nil)) > 0 {
		t.Errorf("MapInt64Int16 failed")
	}

	if len(MapInt64Int16(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapInt64Int8(plusOneInt64Int8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int8(nil, nil)) > 0 {
		t.Errorf("MapInt64Int8 failed")
	}

	if len(MapInt64Int8(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapInt64Uint(plusOneInt64Uint, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint failed")
	}

	if len(MapInt64Uint(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapInt64Uint64(plusOneInt64Uint64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint64 failed")
	}

	if len(MapInt64Uint64(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapInt64Uint32(plusOneInt64Uint32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint32 failed")
	}

	if len(MapInt64Uint32(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapInt64Uint16(plusOneInt64Uint16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint16 failed")
	}

	if len(MapInt64Uint16(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapInt64Uint8(plusOneInt64Uint8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint8 failed")
	}

	if len(MapInt64Uint8(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapInt64Str(someLogicInt64Str, []int64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Str(nil, nil)) > 0 {
		t.Errorf("MapInt64Str failed")
	}

	if len(MapInt64Str(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapInt64Bool(someLogicInt64Bool, []int64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Bool(nil, nil)) > 0 {
		t.Errorf("MapInt64Bool failed")
	}

	if len(MapInt64Bool(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapInt64Float32(plusOneInt64Float32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Float32(nil, nil)) > 0 {
		t.Errorf("MapInt64Float32 failed")
	}

	if len(MapInt64Float32(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapInt64Float64(plusOneInt64Float64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Float64(nil, nil)) > 0 {
		t.Errorf("MapInt64Float64 failed")
	}

	if len(MapInt64Float64(nil, []int64{})) > 0 {
		t.Errorf("MapInt64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapInt32Int(plusOneInt32Int, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int(nil, nil)) > 0 {
		t.Errorf("MapInt32Int failed")
	}

	if len(MapInt32Int(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapInt32Int64(plusOneInt32Int64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int64(nil, nil)) > 0 {
		t.Errorf("MapInt32Int64 failed")
	}

	if len(MapInt32Int64(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapInt32Int16(plusOneInt32Int16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int16(nil, nil)) > 0 {
		t.Errorf("MapInt32Int16 failed")
	}

	if len(MapInt32Int16(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapInt32Int8(plusOneInt32Int8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int8(nil, nil)) > 0 {
		t.Errorf("MapInt32Int8 failed")
	}

	if len(MapInt32Int8(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapInt32Uint(plusOneInt32Uint, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint failed")
	}

	if len(MapInt32Uint(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapInt32Uint64(plusOneInt32Uint64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint64 failed")
	}

	if len(MapInt32Uint64(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapInt32Uint32(plusOneInt32Uint32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint32 failed")
	}

	if len(MapInt32Uint32(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapInt32Uint16(plusOneInt32Uint16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint16 failed")
	}

	if len(MapInt32Uint16(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapInt32Uint8(plusOneInt32Uint8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint8 failed")
	}

	if len(MapInt32Uint8(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapInt32Str(someLogicInt32Str, []int32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Str(nil, nil)) > 0 {
		t.Errorf("MapInt32Str failed")
	}

	if len(MapInt32Str(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapInt32Bool(someLogicInt32Bool, []int32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Bool(nil, nil)) > 0 {
		t.Errorf("MapInt32Bool failed")
	}

	if len(MapInt32Bool(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapInt32Float32(plusOneInt32Float32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Float32(nil, nil)) > 0 {
		t.Errorf("MapInt32Float32 failed")
	}

	if len(MapInt32Float32(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapInt32Float64(plusOneInt32Float64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Float64(nil, nil)) > 0 {
		t.Errorf("MapInt32Float64 failed")
	}

	if len(MapInt32Float64(nil, []int32{})) > 0 {
		t.Errorf("MapInt32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapInt16Int(plusOneInt16Int, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int(nil, nil)) > 0 {
		t.Errorf("MapInt16Int failed")
	}

	if len(MapInt16Int(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapInt16Int64(plusOneInt16Int64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int64(nil, nil)) > 0 {
		t.Errorf("MapInt16Int64 failed")
	}

	if len(MapInt16Int64(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapInt16Int32(plusOneInt16Int32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int32(nil, nil)) > 0 {
		t.Errorf("MapInt16Int32 failed")
	}

	if len(MapInt16Int32(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapInt16Int8(plusOneInt16Int8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int8(nil, nil)) > 0 {
		t.Errorf("MapInt16Int8 failed")
	}

	if len(MapInt16Int8(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapInt16Uint(plusOneInt16Uint, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint failed")
	}

	if len(MapInt16Uint(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapInt16Uint64(plusOneInt16Uint64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint64 failed")
	}

	if len(MapInt16Uint64(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapInt16Uint32(plusOneInt16Uint32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint32 failed")
	}

	if len(MapInt16Uint32(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapInt16Uint16(plusOneInt16Uint16, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint16 failed")
	}

	if len(MapInt16Uint16(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapInt16Uint8(plusOneInt16Uint8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint8 failed")
	}

	if len(MapInt16Uint8(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapInt16Str(someLogicInt16Str, []int16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Str(nil, nil)) > 0 {
		t.Errorf("MapInt16Str failed")
	}

	if len(MapInt16Str(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapInt16Bool(someLogicInt16Bool, []int16{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Bool(nil, nil)) > 0 {
		t.Errorf("MapInt16Bool failed")
	}

	if len(MapInt16Bool(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapInt16Float32(plusOneInt16Float32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Float32(nil, nil)) > 0 {
		t.Errorf("MapInt16Float32 failed")
	}

	if len(MapInt16Float32(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapInt16Float64(plusOneInt16Float64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Float64(nil, nil)) > 0 {
		t.Errorf("MapInt16Float64 failed")
	}

	if len(MapInt16Float64(nil, []int16{})) > 0 {
		t.Errorf("MapInt16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapInt8Int(plusOneInt8Int, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int(nil, nil)) > 0 {
		t.Errorf("MapInt8Int failed")
	}

	if len(MapInt8Int(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapInt8Int64(plusOneInt8Int64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int64(nil, nil)) > 0 {
		t.Errorf("MapInt8Int64 failed")
	}

	if len(MapInt8Int64(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapInt8Int32(plusOneInt8Int32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int32(nil, nil)) > 0 {
		t.Errorf("MapInt8Int32 failed")
	}

	if len(MapInt8Int32(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapInt8Int16(plusOneInt8Int16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int16(nil, nil)) > 0 {
		t.Errorf("MapInt8Int16 failed")
	}

	if len(MapInt8Int16(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapInt8Uint(plusOneInt8Uint, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint failed")
	}

	if len(MapInt8Uint(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapInt8Uint64(plusOneInt8Uint64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint64 failed")
	}

	if len(MapInt8Uint64(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapInt8Uint32(plusOneInt8Uint32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint32 failed")
	}

	if len(MapInt8Uint32(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapInt8Uint16(plusOneInt8Uint16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint16 failed")
	}

	if len(MapInt8Uint16(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapInt8Uint8(plusOneInt8Uint8, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint8 failed")
	}

	if len(MapInt8Uint8(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapInt8Str(someLogicInt8Str, []int8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapInt8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Str(nil, nil)) > 0 {
		t.Errorf("MapInt8Str failed")
	}

	if len(MapInt8Str(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapInt8Bool(someLogicInt8Bool, []int8{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapInt8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Bool(nil, nil)) > 0 {
		t.Errorf("MapInt8Bool failed")
	}

	if len(MapInt8Bool(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapInt8Float32(plusOneInt8Float32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Float32(nil, nil)) > 0 {
		t.Errorf("MapInt8Float32 failed")
	}

	if len(MapInt8Float32(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapInt8Float64(plusOneInt8Float64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapInt8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Float64(nil, nil)) > 0 {
		t.Errorf("MapInt8Float64 failed")
	}

	if len(MapInt8Float64(nil, []int8{})) > 0 {
		t.Errorf("MapInt8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapUintInt(plusOneUintInt, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt(nil, nil)) > 0 {
		t.Errorf("MapUintInt failed")
	}

	if len(MapUintInt(nil, []uint{})) > 0 {
		t.Errorf("MapUintInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapUintInt64(plusOneUintInt64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt64(nil, nil)) > 0 {
		t.Errorf("MapUintInt64 failed")
	}

	if len(MapUintInt64(nil, []uint{})) > 0 {
		t.Errorf("MapUintInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapUintInt32(plusOneUintInt32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt32(nil, nil)) > 0 {
		t.Errorf("MapUintInt32 failed")
	}

	if len(MapUintInt32(nil, []uint{})) > 0 {
		t.Errorf("MapUintInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapUintInt16(plusOneUintInt16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt16(nil, nil)) > 0 {
		t.Errorf("MapUintInt16 failed")
	}

	if len(MapUintInt16(nil, []uint{})) > 0 {
		t.Errorf("MapUintInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapUintInt8(plusOneUintInt8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt8(nil, nil)) > 0 {
		t.Errorf("MapUintInt8 failed")
	}

	if len(MapUintInt8(nil, []uint{})) > 0 {
		t.Errorf("MapUintInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapUintUint64(plusOneUintUint64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint64(nil, nil)) > 0 {
		t.Errorf("MapUintUint64 failed")
	}

	if len(MapUintUint64(nil, []uint{})) > 0 {
		t.Errorf("MapUintUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapUintUint32(plusOneUintUint32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint32(nil, nil)) > 0 {
		t.Errorf("MapUintUint32 failed")
	}

	if len(MapUintUint32(nil, []uint{})) > 0 {
		t.Errorf("MapUintUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapUintUint16(plusOneUintUint16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint16(nil, nil)) > 0 {
		t.Errorf("MapUintUint16 failed")
	}

	if len(MapUintUint16(nil, []uint{})) > 0 {
		t.Errorf("MapUintUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapUintUint8(plusOneUintUint8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint8(nil, nil)) > 0 {
		t.Errorf("MapUintUint8 failed")
	}

	if len(MapUintUint8(nil, []uint{})) > 0 {
		t.Errorf("MapUintUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapUintStr(someLogicUintStr, []uint{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUintStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintStr(nil, nil)) > 0 {
		t.Errorf("MapUintStr failed")
	}

	if len(MapUintStr(nil, []uint{})) > 0 {
		t.Errorf("MapUintStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapUintBool(someLogicUintBool, []uint{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUintBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintBool(nil, nil)) > 0 {
		t.Errorf("MapUintBool failed")
	}

	if len(MapUintBool(nil, []uint{})) > 0 {
		t.Errorf("MapUintBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapUintFloat32(plusOneUintFloat32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintFloat32(nil, nil)) > 0 {
		t.Errorf("MapUintFloat32 failed")
	}

	if len(MapUintFloat32(nil, []uint{})) > 0 {
		t.Errorf("MapUintFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapUintFloat64(plusOneUintFloat64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUintFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintFloat64(nil, nil)) > 0 {
		t.Errorf("MapUintFloat64 failed")
	}

	if len(MapUintFloat64(nil, []uint{})) > 0 {
		t.Errorf("MapUintFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapUint64Int(plusOneUint64Int, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int(nil, nil)) > 0 {
		t.Errorf("MapUint64Int failed")
	}

	if len(MapUint64Int(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapUint64Int64(plusOneUint64Int64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int64(nil, nil)) > 0 {
		t.Errorf("MapUint64Int64 failed")
	}

	if len(MapUint64Int64(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapUint64Int32(plusOneUint64Int32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int32(nil, nil)) > 0 {
		t.Errorf("MapUint64Int32 failed")
	}

	if len(MapUint64Int32(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapUint64Int16(plusOneUint64Int16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int16(nil, nil)) > 0 {
		t.Errorf("MapUint64Int16 failed")
	}

	if len(MapUint64Int16(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapUint64Int8(plusOneUint64Int8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int8(nil, nil)) > 0 {
		t.Errorf("MapUint64Int8 failed")
	}

	if len(MapUint64Int8(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapUint64Uint(plusOneUint64Uint, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint failed")
	}

	if len(MapUint64Uint(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapUint64Uint32(plusOneUint64Uint32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint32 failed")
	}

	if len(MapUint64Uint32(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapUint64Uint16(plusOneUint64Uint16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint16 failed")
	}

	if len(MapUint64Uint16(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapUint64Uint8(plusOneUint64Uint8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint8 failed")
	}

	if len(MapUint64Uint8(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapUint64Str(someLogicUint64Str, []uint64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Str(nil, nil)) > 0 {
		t.Errorf("MapUint64Str failed")
	}

	if len(MapUint64Str(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapUint64Bool(someLogicUint64Bool, []uint64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Bool(nil, nil)) > 0 {
		t.Errorf("MapUint64Bool failed")
	}

	if len(MapUint64Bool(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapUint64Float32(plusOneUint64Float32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Float32(nil, nil)) > 0 {
		t.Errorf("MapUint64Float32 failed")
	}

	if len(MapUint64Float32(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapUint64Float64(plusOneUint64Float64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Float64(nil, nil)) > 0 {
		t.Errorf("MapUint64Float64 failed")
	}

	if len(MapUint64Float64(nil, []uint64{})) > 0 {
		t.Errorf("MapUint64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapUint32Int(plusOneUint32Int, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int(nil, nil)) > 0 {
		t.Errorf("MapUint32Int failed")
	}

	if len(MapUint32Int(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapUint32Int64(plusOneUint32Int64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int64(nil, nil)) > 0 {
		t.Errorf("MapUint32Int64 failed")
	}

	if len(MapUint32Int64(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapUint32Int32(plusOneUint32Int32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int32(nil, nil)) > 0 {
		t.Errorf("MapUint32Int32 failed")
	}

	if len(MapUint32Int32(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapUint32Int16(plusOneUint32Int16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int16(nil, nil)) > 0 {
		t.Errorf("MapUint32Int16 failed")
	}

	if len(MapUint32Int16(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapUint32Int8(plusOneUint32Int8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int8(nil, nil)) > 0 {
		t.Errorf("MapUint32Int8 failed")
	}

	if len(MapUint32Int8(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapUint32Uint(plusOneUint32Uint, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint failed")
	}

	if len(MapUint32Uint(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapUint32Uint64(plusOneUint32Uint64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint64 failed")
	}

	if len(MapUint32Uint64(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapUint32Uint16(plusOneUint32Uint16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint16 failed")
	}

	if len(MapUint32Uint16(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapUint32Uint8(plusOneUint32Uint8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint8 failed")
	}

	if len(MapUint32Uint8(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapUint32Str(someLogicUint32Str, []uint32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Str(nil, nil)) > 0 {
		t.Errorf("MapUint32Str failed")
	}

	if len(MapUint32Str(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapUint32Bool(someLogicUint32Bool, []uint32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Bool(nil, nil)) > 0 {
		t.Errorf("MapUint32Bool failed")
	}

	if len(MapUint32Bool(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapUint32Float32(plusOneUint32Float32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Float32(nil, nil)) > 0 {
		t.Errorf("MapUint32Float32 failed")
	}

	if len(MapUint32Float32(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapUint32Float64(plusOneUint32Float64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Float64(nil, nil)) > 0 {
		t.Errorf("MapUint32Float64 failed")
	}

	if len(MapUint32Float64(nil, []uint32{})) > 0 {
		t.Errorf("MapUint32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapUint16Int(plusOneUint16Int, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int(nil, nil)) > 0 {
		t.Errorf("MapUint16Int failed")
	}

	if len(MapUint16Int(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapUint16Int64(plusOneUint16Int64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int64(nil, nil)) > 0 {
		t.Errorf("MapUint16Int64 failed")
	}

	if len(MapUint16Int64(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapUint16Int32(plusOneUint16Int32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int32(nil, nil)) > 0 {
		t.Errorf("MapUint16Int32 failed")
	}

	if len(MapUint16Int32(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapUint16Int16(plusOneUint16Int16, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int16(nil, nil)) > 0 {
		t.Errorf("MapUint16Int16 failed")
	}

	if len(MapUint16Int16(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapUint16Int8(plusOneUint16Int8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int8(nil, nil)) > 0 {
		t.Errorf("MapUint16Int8 failed")
	}

	if len(MapUint16Int8(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapUint16Uint(plusOneUint16Uint, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint failed")
	}

	if len(MapUint16Uint(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapUint16Uint64(plusOneUint16Uint64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint64 failed")
	}

	if len(MapUint16Uint64(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapUint16Uint32(plusOneUint16Uint32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint32 failed")
	}

	if len(MapUint16Uint32(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapUint16Uint8(plusOneUint16Uint8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint8 failed")
	}

	if len(MapUint16Uint8(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapUint16Str(someLogicUint16Str, []uint16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Str(nil, nil)) > 0 {
		t.Errorf("MapUint16Str failed")
	}

	if len(MapUint16Str(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapUint16Bool(someLogicUint16Bool, []uint16{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Bool(nil, nil)) > 0 {
		t.Errorf("MapUint16Bool failed")
	}

	if len(MapUint16Bool(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapUint16Float32(plusOneUint16Float32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Float32(nil, nil)) > 0 {
		t.Errorf("MapUint16Float32 failed")
	}

	if len(MapUint16Float32(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapUint16Float64(plusOneUint16Float64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Float64(nil, nil)) > 0 {
		t.Errorf("MapUint16Float64 failed")
	}

	if len(MapUint16Float64(nil, []uint16{})) > 0 {
		t.Errorf("MapUint16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapUint8Int(plusOneUint8Int, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int(nil, nil)) > 0 {
		t.Errorf("MapUint8Int failed")
	}

	if len(MapUint8Int(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapUint8Int64(plusOneUint8Int64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int64(nil, nil)) > 0 {
		t.Errorf("MapUint8Int64 failed")
	}

	if len(MapUint8Int64(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapUint8Int32(plusOneUint8Int32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int32(nil, nil)) > 0 {
		t.Errorf("MapUint8Int32 failed")
	}

	if len(MapUint8Int32(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapUint8Int16(plusOneUint8Int16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int16(nil, nil)) > 0 {
		t.Errorf("MapUint8Int16 failed")
	}

	if len(MapUint8Int16(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapUint8Int8(plusOneUint8Int8, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int8(nil, nil)) > 0 {
		t.Errorf("MapUint8Int8 failed")
	}

	if len(MapUint8Int8(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapUint8Uint(plusOneUint8Uint, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint failed")
	}

	if len(MapUint8Uint(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapUint8Uint64(plusOneUint8Uint64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint64 failed")
	}

	if len(MapUint8Uint64(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapUint8Uint32(plusOneUint8Uint32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint32 failed")
	}

	if len(MapUint8Uint32(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapUint8Uint16(plusOneUint8Uint16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint16 failed")
	}

	if len(MapUint8Uint16(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapUint8Str(someLogicUint8Str, []uint8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapUint8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Str(nil, nil)) > 0 {
		t.Errorf("MapUint8Str failed")
	}

	if len(MapUint8Str(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapUint8Bool(someLogicUint8Bool, []uint8{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapUint8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Bool(nil, nil)) > 0 {
		t.Errorf("MapUint8Bool failed")
	}

	if len(MapUint8Bool(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapUint8Float32(plusOneUint8Float32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Float32(nil, nil)) > 0 {
		t.Errorf("MapUint8Float32 failed")
	}

	if len(MapUint8Float32(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapUint8Float64(plusOneUint8Float64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapUint8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Float64(nil, nil)) > 0 {
		t.Errorf("MapUint8Float64 failed")
	}

	if len(MapUint8Float64(nil, []uint8{})) > 0 {
		t.Errorf("MapUint8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10}
	newList := MapStrInt(someLogicStrInt, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt(nil, nil)) > 0 {
		t.Errorf("MapStrInt failed")
	}

	if len(MapStrInt(nil, []string{})) > 0 {
		t.Errorf("MapStrInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10}
	newList := MapStrInt64(someLogicStrInt64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt64(nil, nil)) > 0 {
		t.Errorf("MapStrInt64 failed")
	}

	if len(MapStrInt64(nil, []string{})) > 0 {
		t.Errorf("MapStrInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10}
	newList := MapStrInt32(someLogicStrInt32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt32(nil, nil)) > 0 {
		t.Errorf("MapStrInt32 failed")
	}

	if len(MapStrInt32(nil, []string{})) > 0 {
		t.Errorf("MapStrInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10}
	newList := MapStrInt16(someLogicStrInt16, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt16(nil, nil)) > 0 {
		t.Errorf("MapStrInt16 failed")
	}

	if len(MapStrInt16(nil, []string{})) > 0 {
		t.Errorf("MapStrInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10}
	newList := MapStrInt8(someLogicStrInt8, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt8(nil, nil)) > 0 {
		t.Errorf("MapStrInt8 failed")
	}

	if len(MapStrInt8(nil, []string{})) > 0 {
		t.Errorf("MapStrInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10}
	newList := MapStrUint(someLogicStrUint, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint(nil, nil)) > 0 {
		t.Errorf("MapStrUint failed")
	}

	if len(MapStrUint(nil, []string{})) > 0 {
		t.Errorf("MapStrUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10}
	newList := MapStrUint64(someLogicStrUint64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint64(nil, nil)) > 0 {
		t.Errorf("MapStrUint64 failed")
	}

	if len(MapStrUint64(nil, []string{})) > 0 {
		t.Errorf("MapStrUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10}
	newList := MapStrUint32(someLogicStrUint32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint32(nil, nil)) > 0 {
		t.Errorf("MapStrUint32 failed")
	}

	if len(MapStrUint32(nil, []string{})) > 0 {
		t.Errorf("MapStrUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10}
	newList := MapStrUint16(someLogicStrUint16, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint16(nil, nil)) > 0 {
		t.Errorf("MapStrUint16 failed")
	}

	if len(MapStrUint16(nil, []string{})) > 0 {
		t.Errorf("MapStrUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10}
	newList := MapStrUint8(someLogicStrUint8, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint8(nil, nil)) > 0 {
		t.Errorf("MapStrUint8 failed")
	}

	if len(MapStrUint8(nil, []string{})) > 0 {
		t.Errorf("MapStrUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapStrBool(someLogicStrBool, []string{"10", "0"})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapStrBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrBool(nil, nil)) > 0 {
		t.Errorf("MapStrBool failed")
	}

	if len(MapStrBool(nil, []string{})) > 0 {
		t.Errorf("MapStrBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10}
	newList := MapStrFloat32(someLogicStrFloat32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrFloat32(nil, nil)) > 0 {
		t.Errorf("MapStrFloat32 failed")
	}

	if len(MapStrFloat32(nil, []string{})) > 0 {
		t.Errorf("MapStrFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10}
	newList := MapStrFloat64(someLogicStrFloat64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("MapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrFloat64(nil, nil)) > 0 {
		t.Errorf("MapStrFloat64 failed")
	}

	if len(MapStrFloat64(nil, []string{})) > 0 {
		t.Errorf("MapStrFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10, 0}
	newList := MapBoolInt(someLogicBoolInt, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt(nil, nil)) > 0 {
		t.Errorf("MapBoolInt failed")
	}

	if len(MapBoolInt(nil, []bool{})) > 0 {
		t.Errorf("MapBoolInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10, 0}
	newList := MapBoolInt64(someLogicBoolInt64, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt64(nil, nil)) > 0 {
		t.Errorf("MapBoolInt64 failed")
	}

	if len(MapBoolInt64(nil, []bool{})) > 0 {
		t.Errorf("MapBoolInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10, 0}
	newList := MapBoolInt32(someLogicBoolInt32, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt32(nil, nil)) > 0 {
		t.Errorf("MapBoolInt32 failed")
	}

	if len(MapBoolInt32(nil, []bool{})) > 0 {
		t.Errorf("MapBoolInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10, 0}
	newList := MapBoolInt16(someLogicBoolInt16, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt16(nil, nil)) > 0 {
		t.Errorf("MapBoolInt16 failed")
	}

	if len(MapBoolInt16(nil, []bool{})) > 0 {
		t.Errorf("MapBoolInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10, 0}
	newList := MapBoolInt8(someLogicBoolInt8, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt8(nil, nil)) > 0 {
		t.Errorf("MapBoolInt8 failed")
	}

	if len(MapBoolInt8(nil, []bool{})) > 0 {
		t.Errorf("MapBoolInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10, 0}
	newList := MapBoolUint(someLogicBoolUint, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint(nil, nil)) > 0 {
		t.Errorf("MapBoolUint failed")
	}

	if len(MapBoolUint(nil, []bool{})) > 0 {
		t.Errorf("MapBoolUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10, 0}
	newList := MapBoolUint64(someLogicBoolUint64, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint64(nil, nil)) > 0 {
		t.Errorf("MapBoolUint64 failed")
	}

	if len(MapBoolUint64(nil, []bool{})) > 0 {
		t.Errorf("MapBoolUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10, 0}
	newList := MapBoolUint32(someLogicBoolUint32, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint32(nil, nil)) > 0 {
		t.Errorf("MapBoolUint32 failed")
	}

	if len(MapBoolUint32(nil, []bool{})) > 0 {
		t.Errorf("MapBoolUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10, 0}
	newList := MapBoolUint16(someLogicBoolUint16, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint16(nil, nil)) > 0 {
		t.Errorf("MapBoolUint16 failed")
	}

	if len(MapBoolUint16(nil, []bool{})) > 0 {
		t.Errorf("MapBoolUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10, 0}
	newList := MapBoolUint8(someLogicBoolUint8, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint8(nil, nil)) > 0 {
		t.Errorf("MapBoolUint8 failed")
	}

	if len(MapBoolUint8(nil, []bool{})) > 0 {
		t.Errorf("MapBoolUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10", "0"}
	newList := MapBoolStr(someLogicBoolStr, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolStr(nil, nil)) > 0 {
		t.Errorf("MapBoolStr failed")
	}

	if len(MapBoolStr(nil, []bool{})) > 0 {
		t.Errorf("MapBoolStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10, 0}
	newList := MapBoolFloat32(someLogicBoolFloat32, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolFloat32(nil, nil)) > 0 {
		t.Errorf("MapBoolFloat32 failed")
	}

	if len(MapBoolFloat32(nil, []bool{})) > 0 {
		t.Errorf("MapBoolFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10, 0}
	newList := MapBoolFloat64(someLogicBoolFloat64, []bool{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapBoolFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolFloat64(nil, nil)) > 0 {
		t.Errorf("MapBoolFloat64 failed")
	}

	if len(MapBoolFloat64(nil, []bool{})) > 0 {
		t.Errorf("MapBoolFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapFloat32Int(plusOneFloat32Int, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int failed")
	}

	if len(MapFloat32Int(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapFloat32Int64(plusOneFloat32Int64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int64 failed")
	}

	if len(MapFloat32Int64(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapFloat32Int32(plusOneFloat32Int32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int32(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int32 failed")
	}

	if len(MapFloat32Int32(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapFloat32Int16(plusOneFloat32Int16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int16(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int16 failed")
	}

	if len(MapFloat32Int16(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapFloat32Int8(plusOneFloat32Int8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int8(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int8 failed")
	}

	if len(MapFloat32Int8(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapFloat32Uint(plusOneFloat32Uint, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint failed")
	}

	if len(MapFloat32Uint(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapFloat32Uint64(plusOneFloat32Uint64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint64 failed")
	}

	if len(MapFloat32Uint64(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapFloat32Uint32(plusOneFloat32Uint32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint32(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint32 failed")
	}

	if len(MapFloat32Uint32(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapFloat32Uint16(plusOneFloat32Uint16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint16(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint16 failed")
	}

	if len(MapFloat32Uint16(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapFloat32Uint8(plusOneFloat32Uint8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint8(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint8 failed")
	}

	if len(MapFloat32Uint8(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapFloat32Str(someLogicFloat32Str, []float32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapFloat32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Str(nil, nil)) > 0 {
		t.Errorf("MapFloat32Str failed")
	}

	if len(MapFloat32Str(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapFloat32Bool(someLogicFloat32Bool, []float32{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapFloat32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Bool(nil, nil)) > 0 {
		t.Errorf("MapFloat32Bool failed")
	}

	if len(MapFloat32Bool(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := MapFloat32Float64(plusOneFloat32Float64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Float64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Float64 failed")
	}

	if len(MapFloat32Float64(nil, []float32{})) > 0 {
		t.Errorf("MapFloat32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := MapFloat64Int(plusOneFloat64Int, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int failed")
	}

	if len(MapFloat64Int(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := MapFloat64Int64(plusOneFloat64Int64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int64(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int64 failed")
	}

	if len(MapFloat64Int64(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := MapFloat64Int32(plusOneFloat64Int32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int32 failed")
	}

	if len(MapFloat64Int32(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := MapFloat64Int16(plusOneFloat64Int16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int16(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int16 failed")
	}

	if len(MapFloat64Int16(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := MapFloat64Int8(plusOneFloat64Int8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int8(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int8 failed")
	}

	if len(MapFloat64Int8(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := MapFloat64Uint(plusOneFloat64Uint, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint failed")
	}

	if len(MapFloat64Uint(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := MapFloat64Uint64(plusOneFloat64Uint64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint64(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint64 failed")
	}

	if len(MapFloat64Uint64(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := MapFloat64Uint32(plusOneFloat64Uint32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint32 failed")
	}

	if len(MapFloat64Uint32(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := MapFloat64Uint16(plusOneFloat64Uint16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint16(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint16 failed")
	}

	if len(MapFloat64Uint16(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := MapFloat64Uint8(plusOneFloat64Uint8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint8(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint8 failed")
	}

	if len(MapFloat64Uint8(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := MapFloat64Str(someLogicFloat64Str, []float64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("MapFloat64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Str(nil, nil)) > 0 {
		t.Errorf("MapFloat64Str failed")
	}

	if len(MapFloat64Str(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := MapFloat64Bool(someLogicFloat64Bool, []float64{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("MapFloat64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Bool(nil, nil)) > 0 {
		t.Errorf("MapFloat64Bool failed")
	}

	if len(MapFloat64Bool(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := MapFloat64Float32(plusOneFloat64Float32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("MapFloat64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Float32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Float32 failed")
	}

	if len(MapFloat64Float32(nil, []float64{})) > 0 {
		t.Errorf("MapFloat64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
