package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestKeysInt(t *testing.T) {
	m := map[int]int{1: 1}
	expectedList := []int{1}
	actualList := KeysInt(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntInt64(t *testing.T) {
	m := map[int]int64{1: 1}
	expectedList := []int{1}
	actualList := KeysIntInt64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntInt32(t *testing.T) {
	m := map[int]int32{1: 1}
	expectedList := []int{1}
	actualList := KeysIntInt32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntInt16(t *testing.T) {
	m := map[int]int16{1: 1}
	expectedList := []int{1}
	actualList := KeysIntInt16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntInt8(t *testing.T) {
	m := map[int]int8{1: 1}
	expectedList := []int{1}
	actualList := KeysIntInt8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntUint(t *testing.T) {
	m := map[int]uint{1: 1}
	expectedList := []int{1}
	actualList := KeysIntUint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntUint64(t *testing.T) {
	m := map[int]uint64{1: 1}
	expectedList := []int{1}
	actualList := KeysIntUint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntUint32(t *testing.T) {
	m := map[int]uint32{1: 1}
	expectedList := []int{1}
	actualList := KeysIntUint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntUint16(t *testing.T) {
	m := map[int]uint16{1: 1}
	expectedList := []int{1}
	actualList := KeysIntUint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntUint8(t *testing.T) {
	m := map[int]uint8{1: 1}
	expectedList := []int{1}
	actualList := KeysIntUint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntStr(t *testing.T) {
	m := map[int]string{1: "1"}
	expectedList := []int{1}
	actualList := KeysIntStr(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntBool(t *testing.T) {
	m := map[int]bool{1: true}
	expectedList := []int{1}
	actualList := KeysIntBool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntFloat32(t *testing.T) {
	m := map[int]float32{1: 1}
	expectedList := []int{1}
	actualList := KeysIntFloat32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysIntFloat64(t *testing.T) {
	m := map[int]float64{1: 1}
	expectedList := []int{1}
	actualList := KeysIntFloat64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysIntFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Int(t *testing.T) {
	m := map[int64]int{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64(t *testing.T) {
	m := map[int64]int64{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Int32(t *testing.T) {
	m := map[int64]int32{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Int16(t *testing.T) {
	m := map[int64]int16{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Int8(t *testing.T) {
	m := map[int64]int8{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Uint(t *testing.T) {
	m := map[int64]uint{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Uint64(t *testing.T) {
	m := map[int64]uint64{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Uint32(t *testing.T) {
	m := map[int64]uint32{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Uint16(t *testing.T) {
	m := map[int64]uint16{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Uint8(t *testing.T) {
	m := map[int64]uint8{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Str(t *testing.T) {
	m := map[int64]string{1: "1"}
	expectedList := []int64{1}
	actualList := KeysInt64Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Bool(t *testing.T) {
	m := map[int64]bool{1: true}
	expectedList := []int64{1}
	actualList := KeysInt64Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Float32(t *testing.T) {
	m := map[int64]float32{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt64Float64(t *testing.T) {
	m := map[int64]float64{1: 1}
	expectedList := []int64{1}
	actualList := KeysInt64Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt64Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Int(t *testing.T) {
	m := map[int32]int{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Int64(t *testing.T) {
	m := map[int32]int64{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32(t *testing.T) {
	m := map[int32]int32{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Int16(t *testing.T) {
	m := map[int32]int16{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Int8(t *testing.T) {
	m := map[int32]int8{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Uint(t *testing.T) {
	m := map[int32]uint{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Uint64(t *testing.T) {
	m := map[int32]uint64{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Uint32(t *testing.T) {
	m := map[int32]uint32{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Uint16(t *testing.T) {
	m := map[int32]uint16{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Uint8(t *testing.T) {
	m := map[int32]uint8{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Str(t *testing.T) {
	m := map[int32]string{1: "1"}
	expectedList := []int32{1}
	actualList := KeysInt32Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Bool(t *testing.T) {
	m := map[int32]bool{1: true}
	expectedList := []int32{1}
	actualList := KeysInt32Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Float32(t *testing.T) {
	m := map[int32]float32{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt32Float64(t *testing.T) {
	m := map[int32]float64{1: 1}
	expectedList := []int32{1}
	actualList := KeysInt32Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt32Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Int(t *testing.T) {
	m := map[int16]int{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Int64(t *testing.T) {
	m := map[int16]int64{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Int32(t *testing.T) {
	m := map[int16]int32{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16(t *testing.T) {
	m := map[int16]int16{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Int8(t *testing.T) {
	m := map[int16]int8{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Uint(t *testing.T) {
	m := map[int16]uint{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Uint64(t *testing.T) {
	m := map[int16]uint64{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Uint32(t *testing.T) {
	m := map[int16]uint32{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Uint16(t *testing.T) {
	m := map[int16]uint16{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Uint8(t *testing.T) {
	m := map[int16]uint8{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Str(t *testing.T) {
	m := map[int16]string{1: "1"}
	expectedList := []int16{1}
	actualList := KeysInt16Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Bool(t *testing.T) {
	m := map[int16]bool{1: true}
	expectedList := []int16{1}
	actualList := KeysInt16Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Float32(t *testing.T) {
	m := map[int16]float32{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt16Float64(t *testing.T) {
	m := map[int16]float64{1: 1}
	expectedList := []int16{1}
	actualList := KeysInt16Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt16Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Int(t *testing.T) {
	m := map[int8]int{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Int64(t *testing.T) {
	m := map[int8]int64{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Int32(t *testing.T) {
	m := map[int8]int32{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Int16(t *testing.T) {
	m := map[int8]int16{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8(t *testing.T) {
	m := map[int8]int8{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Uint(t *testing.T) {
	m := map[int8]uint{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Uint64(t *testing.T) {
	m := map[int8]uint64{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Uint32(t *testing.T) {
	m := map[int8]uint32{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Uint16(t *testing.T) {
	m := map[int8]uint16{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Uint8(t *testing.T) {
	m := map[int8]uint8{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Str(t *testing.T) {
	m := map[int8]string{1: "1"}
	expectedList := []int8{1}
	actualList := KeysInt8Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Bool(t *testing.T) {
	m := map[int8]bool{1: true}
	expectedList := []int8{1}
	actualList := KeysInt8Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Float32(t *testing.T) {
	m := map[int8]float32{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysInt8Float64(t *testing.T) {
	m := map[int8]float64{1: 1}
	expectedList := []int8{1}
	actualList := KeysInt8Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysInt8Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintInt(t *testing.T) {
	m := map[uint]int{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintInt(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintInt64(t *testing.T) {
	m := map[uint]int64{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintInt64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintInt32(t *testing.T) {
	m := map[uint]int32{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintInt32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintInt16(t *testing.T) {
	m := map[uint]int16{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintInt16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintInt8(t *testing.T) {
	m := map[uint]int8{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintInt8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint(t *testing.T) {
	m := map[uint]uint{1: 1}
	expectedList := []uint{1}
	actualList := KeysUint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintUint64(t *testing.T) {
	m := map[uint]uint64{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintUint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintUint32(t *testing.T) {
	m := map[uint]uint32{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintUint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintUint16(t *testing.T) {
	m := map[uint]uint16{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintUint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintUint8(t *testing.T) {
	m := map[uint]uint8{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintUint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintStr(t *testing.T) {
	m := map[uint]string{1: "1"}
	expectedList := []uint{1}
	actualList := KeysUintStr(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintBool(t *testing.T) {
	m := map[uint]bool{1: true}
	expectedList := []uint{1}
	actualList := KeysUintBool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintFloat32(t *testing.T) {
	m := map[uint]float32{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintFloat32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUintFloat64(t *testing.T) {
	m := map[uint]float64{1: 1}
	expectedList := []uint{1}
	actualList := KeysUintFloat64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUintFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Int(t *testing.T) {
	m := map[uint64]int{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Int64(t *testing.T) {
	m := map[uint64]int64{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Int32(t *testing.T) {
	m := map[uint64]int32{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Int16(t *testing.T) {
	m := map[uint64]int16{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Int8(t *testing.T) {
	m := map[uint64]int8{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Uint(t *testing.T) {
	m := map[uint64]uint{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64(t *testing.T) {
	m := map[uint64]uint64{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Uint32(t *testing.T) {
	m := map[uint64]uint32{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Uint16(t *testing.T) {
	m := map[uint64]uint16{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Uint8(t *testing.T) {
	m := map[uint64]uint8{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Str(t *testing.T) {
	m := map[uint64]string{1: "1"}
	expectedList := []uint64{1}
	actualList := KeysUint64Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Bool(t *testing.T) {
	m := map[uint64]bool{1: true}
	expectedList := []uint64{1}
	actualList := KeysUint64Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Float32(t *testing.T) {
	m := map[uint64]float32{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint64Float64(t *testing.T) {
	m := map[uint64]float64{1: 1}
	expectedList := []uint64{1}
	actualList := KeysUint64Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint64Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Int(t *testing.T) {
	m := map[uint32]int{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Int64(t *testing.T) {
	m := map[uint32]int64{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Int32(t *testing.T) {
	m := map[uint32]int32{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Int16(t *testing.T) {
	m := map[uint32]int16{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Int8(t *testing.T) {
	m := map[uint32]int8{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Uint(t *testing.T) {
	m := map[uint32]uint{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Uint64(t *testing.T) {
	m := map[uint32]uint64{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32(t *testing.T) {
	m := map[uint32]uint32{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Uint16(t *testing.T) {
	m := map[uint32]uint16{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Uint8(t *testing.T) {
	m := map[uint32]uint8{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Str(t *testing.T) {
	m := map[uint32]string{1: "1"}
	expectedList := []uint32{1}
	actualList := KeysUint32Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Bool(t *testing.T) {
	m := map[uint32]bool{1: true}
	expectedList := []uint32{1}
	actualList := KeysUint32Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Float32(t *testing.T) {
	m := map[uint32]float32{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint32Float64(t *testing.T) {
	m := map[uint32]float64{1: 1}
	expectedList := []uint32{1}
	actualList := KeysUint32Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint32Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Int(t *testing.T) {
	m := map[uint16]int{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Int64(t *testing.T) {
	m := map[uint16]int64{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Int32(t *testing.T) {
	m := map[uint16]int32{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Int16(t *testing.T) {
	m := map[uint16]int16{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Int8(t *testing.T) {
	m := map[uint16]int8{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Uint(t *testing.T) {
	m := map[uint16]uint{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Uint64(t *testing.T) {
	m := map[uint16]uint64{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Uint32(t *testing.T) {
	m := map[uint16]uint32{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16(t *testing.T) {
	m := map[uint16]uint16{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Uint8(t *testing.T) {
	m := map[uint16]uint8{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Str(t *testing.T) {
	m := map[uint16]string{1: "1"}
	expectedList := []uint16{1}
	actualList := KeysUint16Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Bool(t *testing.T) {
	m := map[uint16]bool{1: true}
	expectedList := []uint16{1}
	actualList := KeysUint16Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Float32(t *testing.T) {
	m := map[uint16]float32{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint16Float64(t *testing.T) {
	m := map[uint16]float64{1: 1}
	expectedList := []uint16{1}
	actualList := KeysUint16Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint16Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Int(t *testing.T) {
	m := map[uint8]int{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Int64(t *testing.T) {
	m := map[uint8]int64{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Int32(t *testing.T) {
	m := map[uint8]int32{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Int16(t *testing.T) {
	m := map[uint8]int16{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Int8(t *testing.T) {
	m := map[uint8]int8{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Uint(t *testing.T) {
	m := map[uint8]uint{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Uint64(t *testing.T) {
	m := map[uint8]uint64{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Uint32(t *testing.T) {
	m := map[uint8]uint32{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Uint16(t *testing.T) {
	m := map[uint8]uint16{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8(t *testing.T) {
	m := map[uint8]uint8{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Str(t *testing.T) {
	m := map[uint8]string{1: "1"}
	expectedList := []uint8{1}
	actualList := KeysUint8Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Bool(t *testing.T) {
	m := map[uint8]bool{1: true}
	expectedList := []uint8{1}
	actualList := KeysUint8Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Float32(t *testing.T) {
	m := map[uint8]float32{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysUint8Float64(t *testing.T) {
	m := map[uint8]float64{1: 1}
	expectedList := []uint8{1}
	actualList := KeysUint8Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysUint8Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrInt(t *testing.T) {
	m := map[string]int{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrInt(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrInt64(t *testing.T) {
	m := map[string]int64{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrInt64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrInt32(t *testing.T) {
	m := map[string]int32{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrInt32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrInt16(t *testing.T) {
	m := map[string]int16{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrInt16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrInt8(t *testing.T) {
	m := map[string]int8{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrInt8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrUint(t *testing.T) {
	m := map[string]uint{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrUint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrUint64(t *testing.T) {
	m := map[string]uint64{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrUint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrUint32(t *testing.T) {
	m := map[string]uint32{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrUint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrUint16(t *testing.T) {
	m := map[string]uint16{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrUint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrUint8(t *testing.T) {
	m := map[string]uint8{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrUint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStr(t *testing.T) {
	m := map[string]string{"ram": "ram"}
	expectedList := []string{"ram"}
	actualList := KeysStr(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrBool(t *testing.T) {
	m := map[string]bool{"1": true}
	expectedList := []string{"1"}
	actualList := KeysStrBool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrFloat32(t *testing.T) {
	m := map[string]float32{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrFloat32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysStrFloat64(t *testing.T) {
	m := map[string]float64{"1": 1}
	expectedList := []string{"1"}
	actualList := KeysStrFloat64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysStrFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolInt(t *testing.T) {
	m := map[bool]int{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolInt(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolInt64(t *testing.T) {
	m := map[bool]int64{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolInt64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolInt32(t *testing.T) {
	m := map[bool]int32{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolInt32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolInt16(t *testing.T) {
	m := map[bool]int16{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolInt16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolInt8(t *testing.T) {
	m := map[bool]int8{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolInt8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolUint(t *testing.T) {
	m := map[bool]uint{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolUint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolUint64(t *testing.T) {
	m := map[bool]uint64{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolUint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolUint32(t *testing.T) {
	m := map[bool]uint32{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolUint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolUint16(t *testing.T) {
	m := map[bool]uint16{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolUint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolUint8(t *testing.T) {
	m := map[bool]uint8{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolUint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolStr(t *testing.T) {
	m := map[bool]string{true: "1"}
	expectedList := []bool{true}
	actualList := KeysBoolStr(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBool(t *testing.T) {
	m := map[bool]bool{true: true}
	expectedList := []bool{true}
	actualList := KeysBool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolFloat32(t *testing.T) {
	m := map[bool]float32{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolFloat32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysBoolFloat64(t *testing.T) {
	m := map[bool]float64{true: 1}
	expectedList := []bool{true}
	actualList := KeysBoolFloat64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysBoolFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Int(t *testing.T) {
	m := map[float32]int{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Int64(t *testing.T) {
	m := map[float32]int64{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Int32(t *testing.T) {
	m := map[float32]int32{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Int16(t *testing.T) {
	m := map[float32]int16{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Int8(t *testing.T) {
	m := map[float32]int8{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Uint(t *testing.T) {
	m := map[float32]uint{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Uint64(t *testing.T) {
	m := map[float32]uint64{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Uint32(t *testing.T) {
	m := map[float32]uint32{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Uint16(t *testing.T) {
	m := map[float32]uint16{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Uint8(t *testing.T) {
	m := map[float32]uint8{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Str(t *testing.T) {
	m := map[float32]string{1: "1"}
	expectedList := []float32{1}
	actualList := KeysFloat32Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Bool(t *testing.T) {
	m := map[float32]bool{1: true}
	expectedList := []float32{1}
	actualList := KeysFloat32Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32(t *testing.T) {
	m := map[float32]float32{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat32Float64(t *testing.T) {
	m := map[float32]float64{1: 1}
	expectedList := []float32{1}
	actualList := KeysFloat32Float64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat32Float64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Int(t *testing.T) {
	m := map[float64]int{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Int(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Int failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Int64(t *testing.T) {
	m := map[float64]int64{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Int64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Int64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Int32(t *testing.T) {
	m := map[float64]int32{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Int32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Int32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Int16(t *testing.T) {
	m := map[float64]int16{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Int16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Int16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Int8(t *testing.T) {
	m := map[float64]int8{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Int8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Int8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Uint(t *testing.T) {
	m := map[float64]uint{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Uint(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Uint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Uint64(t *testing.T) {
	m := map[float64]uint64{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Uint64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Uint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Uint32(t *testing.T) {
	m := map[float64]uint32{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Uint32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Uint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Uint16(t *testing.T) {
	m := map[float64]uint16{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Uint16(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Uint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Uint8(t *testing.T) {
	m := map[float64]uint8{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Uint8(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Uint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Str(t *testing.T) {
	m := map[float64]string{1: "1"}
	expectedList := []float64{1}
	actualList := KeysFloat64Str(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Str failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Bool(t *testing.T) {
	m := map[float64]bool{1: true}
	expectedList := []float64{1}
	actualList := KeysFloat64Bool(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Bool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64Float32(t *testing.T) {
	m := map[float64]float32{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64Float32(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64Float32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestKeysFloat64(t *testing.T) {
	m := map[float64]float64{1: 1}
	expectedList := []float64{1}
	actualList := KeysFloat64(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test KeysFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
