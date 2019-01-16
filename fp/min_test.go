package fp

import "testing"

func TestMinInt(t *testing.T) {
	// Test : get min number in the list
	list := []int{8, 2, 10, 4}
	min := MinInt(list)
	if min != 2 {
		t.Errorf("MinInt failed. Expected=2, actual=%v", min)
	}

	min = MinInt([]int{})
	if min != 0 {
		t.Errorf("MinInt failed. Expected=0, actual=%v", min)
	}

	min = MinInt(nil)
	if min != 0 {
		t.Errorf("MinInt failed. Expected=0, actual=%v", min)
	}
}

func TestMinInt64(t *testing.T) {
	// Test : get min number in the list
	list := []int64{8, 2, 10, 4}
	min := MinInt64(list)
	if min != 2 {
		t.Errorf("MinInt64 failed. Expected=2, actual=%v", min)
	}

	min = MinInt64([]int64{})
	if min != 0 {
		t.Errorf("MinInt64 failed. Expected=0, actual=%v", min)
	}

	min = MinInt64(nil)
	if min != 0 {
		t.Errorf("MinInt64 failed. Expected=0, actual=%v", min)
	}
}

func TestMinInt32(t *testing.T) {
	// Test : get min number in the list
	list := []int32{8, 2, 10, 4}
	min := MinInt32(list)
	if min != 2 {
		t.Errorf("MinInt32 failed. Expected=2, actual=%v", min)
	}

	min = MinInt32([]int32{})
	if min != 0 {
		t.Errorf("MinInt32 failed. Expected=0, actual=%v", min)
	}

	min = MinInt32(nil)
	if min != 0 {
		t.Errorf("MinInt32 failed. Expected=0, actual=%v", min)
	}
}

func TestMinInt16(t *testing.T) {
	// Test : get min number in the list
	list := []int16{8, 2, 10, 4}
	min := MinInt16(list)
	if min != 2 {
		t.Errorf("MinInt16 failed. Expected=2, actual=%v", min)
	}

	min = MinInt16([]int16{})
	if min != 0 {
		t.Errorf("MinInt16 failed. Expected=0, actual=%v", min)
	}

	min = MinInt16(nil)
	if min != 0 {
		t.Errorf("MinInt16 failed. Expected=0, actual=%v", min)
	}
}

func TestMinInt8(t *testing.T) {
	// Test : get max number in the list
	list := []int8{8, 2, 10, 4}
	min := MinInt8(list)
	if min != 2 {
		t.Errorf("MinInt8 failed. Expected=2, actual=%v", min)
	}

	min = MinInt8([]int8{})
	if min != 0 {
		t.Errorf("MinInt8 failed. Expected=0, actual=%v", min)
	}

	min = MinInt8(nil)
	if min != 0 {
		t.Errorf("MinInt8 failed. Expected=0, actual=%v", min)
	}
}

func TestMinUint(t *testing.T) {
	// Test : get min number in the list
	list := []uint{8, 2, 10, 4}
	min := MinUint(list)
	if min != 2 {
		t.Errorf("MinUint failed. Expected=2, actual=%v", min)
	}

	min = MinUint([]uint{})
	if min != 0 {
		t.Errorf("MinUint failed. Expected=0, actual=%v", min)
	}

	min = MinUint(nil)
	if min != 0 {
		t.Errorf("MinUint failed. Expected=0, actual=%v", min)
	}
}

func TestMinUint64(t *testing.T) {
	// Test : get min number in the list
	list := []uint64{8, 2, 10, 4}
	min := MinUint64(list)
	if min != 2 {
		t.Errorf("MinUint64 failed. Expected=2, actual=%v", min)
	}

	min = MinUint64([]uint64{})
	if min != 0 {
		t.Errorf("MinUint64 failed. Expected=0, actual=%v", min)
	}

	min = MinUint64(nil)
	if min != 0 {
		t.Errorf("MinUint64 failed. Expected=0, actual=%v", min)
	}
}

func TestMinUint32(t *testing.T) {
	// Test : get min number in the list
	list := []uint32{8, 2, 10, 4}
	min := MinUint32(list)
	if min != 2 {
		t.Errorf("MinUint32 failed. Expected=2, actual=%v", min)
	}

	min = MinUint32([]uint32{})
	if min != 0 {
		t.Errorf("MinUint32 failed. Expected=0, actual=%v", min)
	}

	min = MinUint32(nil)
	if min != 0 {
		t.Errorf("MinUint32 failed. Expected=0, actual=%v", min)
	}
}

func TestMinUint16(t *testing.T) {
	// Test : get min number in the list
	list := []uint16{8, 2, 10, 4}
	min := MinUint16(list)
	if min != 2 {
		t.Errorf("MinUint16 failed. Expected=10, actual=%v", min)
	}

	min = MinUint16([]uint16{})
	if min != 0 {
		t.Errorf("MinUint16 failed. Expected=0, actual=%v", min)
	}

	min = MinUint16(nil)
	if min != 0 {
		t.Errorf("MinUint16 failed. Expected=0, actual=%v", min)
	}
}

func TestMinUint8(t *testing.T) {
	// Test : get min number in the list
	list := []uint8{8, 2, 10, 4}
	min := MinUint8(list)
	if min != 2 {
		t.Errorf("MinUint8 failed. Expected=2, actual=%v", min)
	}

	min = MinUint8([]uint8{})
	if min != 0 {
		t.Errorf("MinUint8 failed. Expected=0, actual=%v", min)
	}

	min = MinUint8(nil)
	if min != 0 {
		t.Errorf("MinUint8 failed. Expected=0, actual=%v", min)
	}
}

func TestMinFloat64(t *testing.T) {
	// Test : get min number in the list
	list := []float64{8.2, 2.2, 10.1, 4.1}
	min := MinFloat64(list)
	if min != 2.2 {
		t.Errorf("MinFloat64 failed. Expected=2.2, actual=%f", min)
	}

	min = MinFloat64([]float64{})
	if min != 0 {
		t.Errorf("MinFloat64 failed. Expected=0, actual=%v", min)
	}

	min = MinFloat64(nil)
	if min != 0 {
		t.Errorf("MinFloat64 failed. Expected=0, actual=%v", min)
	}
}

func TestMinFloat32(t *testing.T) {
	// Test : get min number in the list
	list := []float32{8.2, 2.2, 10.1, 4.1}
	min := MinFloat32(list)
	if min != 2.2 {
		t.Errorf("MinFloat64 failed. Expected=2.2, actual=%f", min)
	}

	min = MinFloat32([]float32{})
	if min != 0 {
		t.Errorf("MinFloat32 failed. Expected=0, actual=%v", min)
	}

	min = MinFloat32(nil)
	if min != 0 {
		t.Errorf("MinFloat32 failed. Expected=0, actual=%v", min)
	}
}
