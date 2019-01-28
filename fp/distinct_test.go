package fp

import "testing"

func TestDistinctInt(t *testing.T) {
	// Test : Get distinct values
	expected := []int{8, 2, 0}
	list := []int{8, 2, 8, 0, 2, 0}
	distinct := DistinctInt(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int{8, 2, 0}
	list = []int{8, 2, 0}
	distinct = DistinctInt(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int{}
	list = []int{}
	distinct = DistinctInt(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctInt64(t *testing.T) {
	// Test : Get distinct values
	expected := []int64{8, 2, 0}
	list := []int64{8, 2, 8, 0, 2, 0}
	distinct := DistinctInt64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int64{8, 2, 0}
	list = []int64{8, 2, 0}
	distinct = DistinctInt64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int64{}
	list = []int64{}
	distinct = DistinctInt64(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt64(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt64 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctInt32(t *testing.T) {
	// Test : Get distinct values
	expected := []int32{8, 2, 0}
	list := []int32{8, 2, 8, 0, 2, 0}
	distinct := DistinctInt32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int32{8, 2, 0}
	list = []int32{8, 2, 0}
	distinct = DistinctInt32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int32{}
	list = []int32{}
	distinct = DistinctInt32(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt32(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctInt16(t *testing.T) {
	// Test : Get distinct values
	expected := []int16{8, 2, 0}
	list := []int16{8, 2, 8, 0, 2, 0}
	distinct := DistinctInt16(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int16{8, 2, 0}
	list = []int16{8, 2, 0}
	distinct = DistinctInt16(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int16{}
	list = []int16{}
	distinct = DistinctInt16(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt16(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt32 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctInt8(t *testing.T) {
	// Test : Get distinct values
	expected := []int8{8, 2, 0}
	list := []int8{8, 2, 8, 0, 2, 0}
	distinct := DistinctInt8(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int8{8, 2, 0}
	list = []int8{8, 2, 0}
	distinct = DistinctInt8(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []int8{}
	list = []int8{}
	distinct = DistinctInt8(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt8(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctUint(t *testing.T) {
	// Test : Get distinct values
	expected := []uint{8, 2, 0}
	list := []uint{8, 2, 8, 0, 2, 0}
	distinct := DistinctUint(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint{8, 2, 0}
	list = []uint{8, 2, 0}
	distinct = DistinctUint(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint{}
	list = []uint{}
	distinct = DistinctUint(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt8 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctUint64(t *testing.T) {
	// Test : Get distinct values
	expected := []uint64{8, 2, 0}
	list := []uint64{8, 2, 8, 0, 2, 0}
	distinct := DistinctUint64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint64{8, 2, 0}
	list = []uint64{8, 2, 0}
	distinct = DistinctUint64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint64{}
	list = []uint64{}
	distinct = DistinctUint64(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint64(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint64 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctUint32(t *testing.T) {
	// Test : Get distinct values
	expected := []uint32{8, 2, 0}
	list := []uint32{8, 2, 8, 0, 2, 0}
	distinct := DistinctUint32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint32{8, 2, 0}
	list = []uint32{8, 2, 0}
	distinct = DistinctUint32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint32{}
	list = []uint32{}
	distinct = DistinctUint32(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint32(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint32 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctUint16(t *testing.T) {
	// Test : Get distinct values
	expected := []uint16{8, 2, 0}
	list := []uint16{8, 2, 8, 0, 2, 0}
	distinct := DistinctUint16(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint16 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint16{8, 2, 0}
	list = []uint16{8, 2, 0}
	distinct = DistinctUint16(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint16 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint16{}
	list = []uint16{}
	distinct = DistinctUint16(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint16 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint16(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint16 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctUint8(t *testing.T) {
	// Test : Get distinct values
	expected := []uint8{8, 2, 0}
	list := []uint8{8, 2, 8, 0, 2, 0}
	distinct := DistinctUint8(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint8{8, 2, 0}
	list = []uint8{8, 2, 0}
	distinct = DistinctUint8(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctUint8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []uint8{}
	list = []uint8{}
	distinct = DistinctUint8(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint8 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint8(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint8 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctFloat64(t *testing.T) {
	// Test : Get distinct values
	expected := []float64{8, 2, 0}
	list := []float64{8, 2, 8, 0, 2, 0}
	distinct := DistinctFloat64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctFloat64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []float64{8, 2, 0}
	list = []float64{8, 2, 0}
	distinct = DistinctFloat64(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctFloat64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []float64{}
	list = []float64{}
	distinct = DistinctFloat64(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat64 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctFloat64(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat64 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctFloat32(t *testing.T) {
	// Test : Get distinct values
	expected := []float32{8, 2, 0}
	list := []float32{8, 2, 8, 0, 2, 0}
	distinct := DistinctFloat32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctFloat32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []float32{8, 2, 0}
	list = []float32{8, 2, 0}
	distinct = DistinctFloat32(list)
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctFloat32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []float32{}
	list = []float32{}
	distinct = DistinctFloat32(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat32 failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctFloat32(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat32 failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctStr(t *testing.T) {
	// Test : Get distinct values
	expected := []string{"Bharat", "Hanuman", "Sita"}
	list := []string{"Bharat", "Hanuman", "Bharat", "Sita", "Hanuman", "Sita"}
	distinct := DistinctStr(list)
	if len(distinct) != 3 || distinct[0] != expected[0] || distinct[1] != expected[1] || distinct[2] != expected[2] {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{"Bharat", "Hanuman", "Sita"}
	list = []string{"Bharat", "Hanuman", "Sita"}
	distinct = DistinctStr(list)
	if len(distinct) != 3 || distinct[0] != expected[0] || distinct[1] != expected[1] || distinct[2] != expected[2] {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{}
	list = []string{}
	distinct = DistinctStr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctStr(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinctStrIgnoreCase(t *testing.T) {
	// Test : Get distinct values
	expected := []string{"Bharat", "HanumaN", "SiTa"}
	list := []string{"Bharat", "HanumaN", "BharaT", "SiTa", "Hanuman", "Sita"}
	distinct := DistinctStrIgnoreCase(list)
	if len(distinct) != 3 || distinct[0] != expected[0] || distinct[1] != expected[1] || distinct[2] != expected[2] {
		t.Errorf("TestDistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{"Bharat", "HanumaN", "Sita"}
	list = []string{"Bharat", "HanumaN", "Sita"}
	distinct = DistinctStrIgnoreCase(list)
	if len(distinct) != 3 || distinct[0] != expected[0] || distinct[1] != expected[1] || distinct[2] != expected[2] {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{}
	list = []string{}
	distinct = DistinctStrIgnoreCase(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctStrIgnoreCase(nil)
	if len(distinct) != 0 {
		t.Errorf("DistinctStr failed. Expected=%v, actual=%v", expected, distinct)
	}
}
