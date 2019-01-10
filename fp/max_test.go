package fp

import "testing"

func TestMaxInt(t *testing.T) {
	// Test : get max number in the list
	list := []int{8, 2, 10, 4}
	max := MaxInt(list)
	if max != 10 {
		t.Errorf("MaxInt failed. Expected=10, actual=%v", max)
	}
}

func TestMaxInt64(t *testing.T) {
	// Test : get max number in the list
	list := []int64{8, 2, 10, 4}
	max := MaxInt64(list)
	if max != 10 {
		t.Errorf("MaxInt64 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxInt32(t *testing.T) {
	// Test : get max number in the list
	list := []int32{8, 2, 10, 4}
	max := MaxInt32(list)
	if max != 10 {
		t.Errorf("MaxInt32 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxInt16(t *testing.T) {
	// Test : get max number in the list
	list := []int16{8, 2, 10, 4}
	max := MaxInt16(list)
	if max != 10 {
		t.Errorf("MaxInt16 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxInt8(t *testing.T) {
	// Test : get max number in the list
	list := []int8{8, 2, 10, 4}
	max := MaxInt8(list)
	if max != 10 {
		t.Errorf("MaxInt8 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxUint(t *testing.T) {
	// Test : get max number in the list
	list := []uint{8, 2, 10, 4}
	max := MaxUint(list)
	if max != 10 {
		t.Errorf("MaxUint failed. Expected=10, actual=%v", max)
	}
}

func TestMaxUint64(t *testing.T) {
	// Test : get max number in the list
	list := []uint64{8, 2, 10, 4}
	max := MaxUint64(list)
	if max != 10 {
		t.Errorf("MaxUint64 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxUint32(t *testing.T) {
	// Test : get max number in the list
	list := []uint32{8, 2, 10, 4}
	max := MaxUint32(list)
	if max != 10 {
		t.Errorf("MaxUint32 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxUint16(t *testing.T) {
	// Test : get max number in the list
	list := []uint16{8, 2, 10, 4}
	max := MaxUint16(list)
	if max != 10 {
		t.Errorf("MaxUint16 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxUint8(t *testing.T) {
	// Test : get max number in the list
	list := []uint8{8, 2, 10, 4}
	max := MaxUint8(list)
	if max != 10 {
		t.Errorf("MaxUint8 failed. Expected=10, actual=%v", max)
	}
}

func TestMaxFloat64(t *testing.T) {
	// Test : get max number in the list
	list := []float64{8.2, 2.2, 10.1, 4.1}
	max := MaxFloat64(list)
	if max != 10.1 {
		t.Errorf("MaxFloat64 failed. Expected=10.1, actual=%f", max)
	}
}

func TestMaxFloat32(t *testing.T) {
	// Test : get max number in the list
	list := []float32{8.2, 2.2, 10.1, 4.1}
	max := MaxFloat32(list)
	if max != 10.1 {
		t.Errorf("MaxFloat64 failed. Expected=10.1, actual=%f", max)
	}
}
