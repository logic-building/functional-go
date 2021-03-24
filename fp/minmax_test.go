package fp

import "testing"

func TestMinMaxInt(t *testing.T) {
	// Test : get min, max number in the list
	list := []int{8, 2, 10, 4}
	min, max := MinMaxInt(list)

	if min != 2 || max != 10 {
		t.Errorf("MinMaxInt failed. Expected=2,10 actual=%v,%v", min, max)
	}

	min, max = MinMaxInt([]int{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxInt64(t *testing.T) {
	// Test : get min, max number in the list
	list := []int64{8, 2, 10, 4}
	min, max := MinMaxInt64(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxInt64 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt64([]int64{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt64 failed. Expected=0, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt64(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt64 failed. Expected=0,0 actual=%v,%v", min, max)
	}
}

func TestMinMaxInt32(t *testing.T) {
	// Test : get min, max number in the list
	list := []int32{8, 2, 10, 4}
	min, max := MinMaxInt32(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxInt32 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt32([]int32{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt32 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt32(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt32 failed. Expected=0,0 actual=%v,%v", min, max)
	}
}

func TestMinMaxInt16(t *testing.T) {
	// Test : get min, max number in the list
	list := []int16{8, 2, 10, 4}
	min, max := MinMaxInt16(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxInt16 failed. Expected=2,10 actual=%v,%v", min, max)
	}

	min, max = MinMaxInt16([]int16{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt16 failed. Expected=0,0 actual=%v,%v", min, max)
	}

	min, max = MinMaxInt16(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt16 failed. Expected=0,0 actual=%v,%v", min, max)
	}
}

func TestMinMaxInt8(t *testing.T) {
	// Test : get max number in the list
	list := []int8{8, 2, 10, 4}
	min, max := MinMaxInt8(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxInt8 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt8([]int8{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt8 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxInt8(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxInt8 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxUint(t *testing.T) {
	// Test : get min, max number in the list
	list := []uint{8, 2, 10, 4}
	min, max := MinMaxUint(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxUint failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint([]uint{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxUint64(t *testing.T) {
	// Test : get min, max number in the list
	list := []uint64{8, 2, 10, 4}
	min, max := MinMaxUint64(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxUint64 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint64([]uint64{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint64 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint64(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint64 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxUint32(t *testing.T) {
	// Test : get min, max number in the list
	list := []uint32{8, 2, 10, 4}
	min, max := MinMaxUint32(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxUint32 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint32([]uint32{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint32 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint32(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint32 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxUint16(t *testing.T) {
	// Test : get min, max number in the list
	list := []uint16{8, 2, 10, 4}
	min, max := MinMaxUint16(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxUint16 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint16([]uint16{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint16 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint16(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint16 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxUint8(t *testing.T) {
	// Test : get min, max number in the list
	list := []uint8{8, 2, 10, 4}
	min, max := MinMaxUint8(list)
	if min != 2 || max != 10 {
		t.Errorf("MinMaxUint8 failed. Expected=2,10, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint8([]uint8{})
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint8 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxUint8(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMaxUint8 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxFloat64(t *testing.T) {
	// Test : get min, max number in the list
	list := []float64{8.2, 2.2, 10.1, 4.1}
	min, max := MinMaxFloat64(list)
	if min != 2.2 || max != 10.1 {
		t.Errorf("MinFloat64 failed. Expected=2.2,10.1, actual=%f,%f", min, max)
	}

	min, max = MinMaxFloat64([]float64{})
	if min != 0 || max != 0 {
		t.Errorf("MinFloat64 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxFloat64(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinFloat64 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}

func TestMinMaxFloat32(t *testing.T) {
	// Test : get min, max number in the list
	list := []float32{8.2, 2.2, 10.1, 4.1}
	min, max := MinMaxFloat32(list)
	if min != 2.2 || max != 10.1 {
		t.Errorf("MinFloat64 failed. Expected=2.2, 10.1, actual=%f,%f", min, max)
	}

	min, max = MinMaxFloat32([]float32{})
	if min != 0 || max != 0 {
		t.Errorf("MinFloat32 failed. Expected=0,0, actual=%v,%v", min, max)
	}

	min, max = MinMaxFloat32(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinFloat32 failed. Expected=0,0, actual=%v,%v", min, max)
	}
}
