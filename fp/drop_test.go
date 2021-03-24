package fp

import "testing"

func TestDropInt(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int{2, 3}
	newList := DropInt(1, []int{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInt failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt(1, []int{})

	if len(newList) != 0 {
		t.Errorf("DropInt failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInt(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropInts(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int{2, 3}
	newList := DropInts([]int{1, 4}, []int{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts([]int{1, 4}, []int{})

	if len(newList) != 0 {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts([]int{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts(nil, []int{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int64{2, 3}
	newList := DropInt64(1, []int64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInt64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt64(1, []int64{})

	if len(newList) != 0 {
		t.Errorf("DropInt64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInt64(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropInts64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int64{2, 3}
	newList := DropInts64([]int64{1, 4}, []int64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts64([]int64{1, 4}, []int64{})

	if len(newList) != 0 {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts64([]int64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts64(nil, []int64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int32{2, 3}
	newList := DropInt32(1, []int32{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInt32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt32(1, []int32{})

	if len(newList) != 0 {
		t.Errorf("DropInt32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInt32(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropInts32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int32{2, 3}
	newList := DropInts32([]int32{1, 4}, []int32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts32([]int32{1, 4}, []int32{})

	if len(newList) != 0 {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts32([]int32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts32(nil, []int32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt16(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int16{2, 3}
	newList := DropInt16(1, []int16{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInt16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt16(1, []int16{})

	if len(newList) != 0 {
		t.Errorf("DropInt16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInt16(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropInts16(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int16{2, 3}
	newList := DropInts16([]int16{1, 4}, []int16{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts16([]int16{1, 4}, []int16{})

	if len(newList) != 0 {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts16([]int16{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts16(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts16(nil, []int16{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropInt8(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int8{2, 3}
	newList := DropInt8(1, []int8{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInt8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInt8(1, []int8{})

	if len(newList) != 0 {
		t.Errorf("DropInt8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInt8(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropInt8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropInts8(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []int8{2, 3}
	newList := DropInts8([]int8{1, 4}, []int8{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropInts8([]int8{1, 4}, []int8{})

	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts8([]int8{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts8(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropInts8(nil, []int8{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUInt(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint{2, 3}
	newList := DropUint(1, []uint{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUint failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint(1, []uint{})

	if len(newList) != 0 {
		t.Errorf("DropUint failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUint(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropUints(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint{2, 3}
	newList := DropUints([]uint{1, 4}, []uint{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints([]uint{1, 4}, []uint{})

	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints([]uint{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints(nil, []uint{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropInts8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUInt64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint64{2, 3}
	newList := DropUint64(1, []uint64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUint64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint64(1, []uint64{})

	if len(newList) != 0 {
		t.Errorf("DropUint64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUint64(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropUints64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint64{2, 3}
	newList := DropUints64([]uint64{1, 4}, []uint64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUints64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints64([]uint64{1, 4}, []uint64{})

	if len(newList) != 0 {
		t.Errorf("DropUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints64([]uint64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints64(nil, []uint64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropUints64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUInt32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint32{2, 3}
	newList := DropUint32(1, []uint32{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUint32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint32(1, []uint32{})

	if len(newList) != 0 {
		t.Errorf("DropUint32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUint32(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropUints32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint32{2, 3}
	newList := DropUints32([]uint32{1, 4}, []uint32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints32([]uint32{1, 4}, []uint32{})

	if len(newList) != 0 {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints32([]uint32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints32(nil, []uint32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropUints32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUInt16(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint16{2, 3}
	newList := DropUint16(1, []uint16{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUint16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint16(1, []uint16{})

	if len(newList) != 0 {
		t.Errorf("DropUint16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUint16(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropUints16(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint16{2, 3}
	newList := DropUints16([]uint16{1, 4}, []uint16{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints16([]uint16{1, 4}, []uint16{})

	if len(newList) != 0 {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints16([]uint16{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints16(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints16(nil, []uint16{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropUints16 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropUInt8(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint8{2, 3}
	newList := DropUint8(1, []uint8{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUint8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUint8(1, []uint8{})

	if len(newList) != 0 {
		t.Errorf("DropUint8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUint8(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropUint8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropUints8(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []uint8{2, 3}
	newList := DropUints8([]uint8{1, 4}, []uint8{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropUints8([]uint8{1, 4}, []uint8{})

	if len(newList) != 0 {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints8([]uint8{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints8(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropUints8(nil, []uint8{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropUints8 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropFloat64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []float64{2, 3}
	newList := DropFloat64(1, []float64{1, 2, 3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropFloat64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat64(1, []float64{})

	if len(newList) != 0 {
		t.Errorf("DropFloat64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloat64(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropFloats64(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []float64{2, 3}
	newList := DropFloats64([]float64{1, 4}, []float64{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropFloats64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloats64([]float64{1, 4}, []float64{})

	if len(newList) != 0 {
		t.Errorf("DropFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats64([]float64{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats64(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloats64 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats64(nil, []float64{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropFloats64 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropFloat32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []float32{2.2, 3.3}
	newList := DropFloat32(1, []float32{1, 2.2, 3.3, 1})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropFloat32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloat32(1, []float32{})

	if len(newList) != 0 {
		t.Errorf("DropFloat32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloat32(1, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloat32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropFloats32(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []float32{2, 3}
	newList := DropFloats32([]float32{1, 4}, []float32{1, 2, 3, 1, 4})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropFloats32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropFloats32([]float32{1, 4}, []float32{})

	if len(newList) != 0 {
		t.Errorf("DropFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats32([]float32{1, 4}, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats32(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropFloats32 failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropFloats32(nil, []float32{1, 4})
	if newList[0] != 1 || newList[1] != 4 {
		t.Errorf("DropFloats32 failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropStr(t *testing.T) {
	// Test : Drop string from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := DropStr("Nks", []string{"Nks", "Ram", "Shyam", "Nks"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropStr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStr("1", []string{})

	if len(newList) != 0 {
		t.Errorf("DropStr failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStr("1", nil)
	if len(newList) != 0 {
		t.Errorf("DropStr failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropStrIgnoreCase(t *testing.T) {
	// Test : Drop string from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := DropStrIgnoreCase("nks", []string{"Nks", "Ram", "Shyam", "Nks"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropStrIgonreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStrIgnoreCase("1", []string{})

	if len(newList) != 0 {
		t.Errorf("DropStrIgonreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrIgnoreCase("1", nil)
	if len(newList) != 0 {
		t.Errorf("DropStrIgonreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}
}

func TestDropStrs(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := DropStrs([]string{"nks", "bharat"}, []string{"nks", "Ram", "Shyam", "Nks", "bharat"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStrs([]string{"1", "4"}, []string{})

	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrs([]string{"1", "4"}, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrs(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrs(nil, []string{"ram", "Shyam"})
	if newList[0] != "ram" || newList[1] != "Shyam" {
		t.Errorf("DropStrIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDropStrsIgonreCase(t *testing.T) {
	// Test : Drop number from the list
	expectedList := []string{"Ram", "Shyam"}
	newList := DropStrsIgnoreCase([]string{"nks", "bharat"}, []string{"Nks", "Ram", "Shyam", "Nks", "bharat"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = DropStrsIgnoreCase([]string{"1", "4"}, []string{})

	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrsIgnoreCase([]string{"1", "4"}, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrsIgnoreCase(nil, nil)
	if len(newList) != 0 {
		t.Errorf("DropStrsIgnoreCase failed. Expected list=%v, actual list=%v", []int{}, newList)
	}

	newList = DropStrsIgnoreCase(nil, []string{"ram", "Shyam"})
	if newList[0] != "ram" || newList[1] != "Shyam" {
		t.Errorf("DropStrIgnoreCase failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}
