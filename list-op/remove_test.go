package list_op

import "testing"

func TestRemoveInt(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int{2, 3}
	newList := RemoveInt(1, []int{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInt failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInt(1, []int{})

	if len(newList) != 0 {
		t.Errorf("RemoveInt failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInt(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInt failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveInts(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int{2, 3}
	newList := RemoveInts([]int{1, 4}, []int{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInts([]int{1, 4}, []int{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts([]int{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts(nil, []int{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveInt64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int64{2, 3}
	newList := RemoveInt64(1, []int64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInt64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInt64(1, []int64{})

	if len(newList) != 0 {
		t.Errorf("RemoveInt64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInt64(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInt64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveInts64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int64{2, 3}
	newList := RemoveInts64([]int64{1, 4}, []int64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInts64([]int64{1, 4}, []int64{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts64([]int64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts64(nil, []int64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveInt32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int32{2, 3}
	newList := RemoveInt32(1, []int32{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInt32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInt32(1, []int32{})

	if len(newList) != 0 {
		t.Errorf("RemoveInt32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInt32(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInt32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveInts32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int32{2, 3}
	newList := RemoveInts32([]int32{1, 4}, []int32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInts32([]int32{1, 4}, []int32{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts32([]int32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts32(nil, []int32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveInt16(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int16{2, 3}
	newList := RemoveInt16(1, []int16{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInt16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInt16(1, []int16{})

	if len(newList) != 0 {
		t.Errorf("RemoveInt16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInt16(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInt16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveInts16(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int16{2, 3}
	newList := RemoveInts16([]int16{1, 4}, []int16{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInts16([]int16{1, 4}, []int16{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts16([]int16{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts16(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts16(nil, []int16{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveInt8(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int8{2, 3}
	newList := RemoveInt8(1, []int8{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInt8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInt8(1, []int8{})

	if len(newList) != 0 {
		t.Errorf("RemoveInt8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInt8(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInt8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveInts8(t *testing.T) {
	// Test : remove number from the list
	expectedList := []int8{2, 3}
	newList := RemoveInts8([]int8{1, 4}, []int8{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveInts8([]int8{1, 4}, []int8{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts8([]int8{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts8(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveInts8(nil, []int8{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveUInt(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint{2, 3}
	newList := RemoveUint(1, []uint{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUint failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUint(1, []uint{})

	if len(newList) != 0 {
		t.Errorf("RemoveUint failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUint(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUint failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveUints(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint{2, 3}
	newList := RemoveUints([]uint{1, 4}, []uint{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUints([]uint{1, 4}, []uint{})

	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints([]uint{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints(nil, []uint{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveUInt64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint64{2, 3}
	newList := RemoveUint64(1, []uint64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUint64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUint64(1, []uint64{})

	if len(newList) != 0 {
		t.Errorf("RemoveUint64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUint64(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUint64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveUints64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint64{2, 3}
	newList := RemoveUints64([]uint64{1, 4}, []uint64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUints64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUints64([]uint64{1, 4}, []uint64{})

	if len(newList) != 0 {
		t.Errorf("RemoveUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints64([]uint64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints64(nil, []uint64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveUints64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveUInt32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint32{2, 3}
	newList := RemoveUint32(1, []uint32{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUint32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUint32(1, []uint32{})

	if len(newList) != 0 {
		t.Errorf("RemoveUint32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUint32(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUint32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveUints32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint32{2, 3}
	newList := RemoveUints32([]uint32{1, 4}, []uint32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUints32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUints32([]uint32{1, 4}, []uint32{})

	if len(newList) != 0 {
		t.Errorf("RemoveUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints32([]uint32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints32(nil, []uint32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveUints32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveUInt16(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint16{2, 3}
	newList := RemoveUint16(1, []uint16{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUint16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUint16(1, []uint16{})

	if len(newList) != 0 {
		t.Errorf("RemoveUint16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUint16(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUint16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveUints16(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint16{2, 3}
	newList := RemoveUints16([]uint16{1, 4}, []uint16{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUints16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUints16([]uint16{1, 4}, []uint16{})

	if len(newList) != 0 {
		t.Errorf("RemoveUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints16([]uint16{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints16(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints16(nil, []uint16{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveUints16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveUInt8(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint8{2, 3}
	newList := RemoveUint8(1, []uint8{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUint8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUint8(1, []uint8{})

	if len(newList) != 0 {
		t.Errorf("RemoveUint8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUint8(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUint8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveUints8(t *testing.T) {
	// Test : remove number from the list
	expectedList := []uint8{2, 3}
	newList := RemoveUints8([]uint8{1, 4}, []uint8{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveUints8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveUints8([]uint8{1, 4}, []uint8{})

	if len(newList) != 0 {
		t.Errorf("RemoveUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints8([]uint8{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints8(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveUints8(nil, []uint8{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveUints8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveFloat64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []float64{2, 3}
	newList := RemoveFloat64(1, []float64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveFloat64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveFloat64(1, []float64{})

	if len(newList) != 0 {
		t.Errorf("RemoveFloat64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloat64(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloat64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveFloats64(t *testing.T) {
	// Test : remove number from the list
	expectedList := []float64{2, 3}
	newList := RemoveFloats64([]float64{1, 4}, []float64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveFloats64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveFloats64([]float64{1, 4}, []float64{})

	if len(newList) != 0 {
		t.Errorf("RemoveFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats64([]float64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats64(nil, []float64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveFloats64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveFloat32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []float32{2.2, 3.3}
	newList := RemoveFloat32(1, []float32{1, 2.2, 3.3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveFloat32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveFloat32(1, []float32{})

	if len(newList) != 0 {
		t.Errorf("RemoveFloat32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloat32(1, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloat32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveFloats32(t *testing.T) {
	// Test : remove number from the list
	expectedList := []float32{2, 3}
	newList := RemoveFloats32([]float32{1, 4}, []float32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveFloats32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveFloats32([]float32{1, 4}, []float32{})

	if len(newList) != 0 {
		t.Errorf("RemoveFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats32([]float32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveFloats32(nil, []float32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("RemoveFloats32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveStr(t *testing.T) {
	// Test : remove string from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := RemoveStr("Nks", []string{"Nks", "Ram", "Shyam", "Nks"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveStr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveStr("1", []string{})

	if len(newList) != 0 {
		t.Errorf("RemoveStr failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStr("1", nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStr failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveStrIgnoreCase(t *testing.T) {
	// Test : remove string from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := RemoveStrIgnoreCase("nks", []string{"Nks", "Ram", "Shyam", "Nks"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveStrIgonreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveStrIgnoreCase("1", []string{})

	if len(newList) != 0 {
		t.Errorf("RemoveStrIgonreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrIgnoreCase("1", nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStrIgonreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestRemoveStrs(t *testing.T) {
	// Test : remove number from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := RemoveStrs([]string{"nks", "bharat"}, []string{"nks", "Ram", "Shyam", "Nks", "bharat"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveStrs([]string{"1", "4"}, []string{})

	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrs([]string{"1", "4"}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrs(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrs(nil, []string{"ram", "Shyam"})
	if newList[0] != "ram" || newList[1] != "Shyam" {
		t.Errorf("RemoveStrIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestRemoveStrsIgonreCase(t *testing.T) {
	// Test : remove number from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := RemoveStrsIgnoreCase([]string{"nks", "bharat"}, []string{"Nks", "Ram", "Shyam", "Nks", "bharat"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = RemoveStrsIgnoreCase([]string{"1", "4"}, []string{})

	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrsIgnoreCase([]string{"1", "4"}, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrsIgnoreCase(nil, nil)
	if len(newList) != 0 {
		t.Errorf("RemoveStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = RemoveStrsIgnoreCase(nil, []string{"ram", "Shyam"})
	if newList[0] != "ram" || newList[1] != "Shyam" {
		t.Errorf("RemoveStrIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}
