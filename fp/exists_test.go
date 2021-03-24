package fp

import "testing"

func TestExistsInt(t *testing.T) {
	// Test : value exist in the list

	list1 := []int{8, 2, 10, 4}
	if !ExistsInt(8, list1) {
		t.Errorf("ExistsInt failed. Expected=true, actual=false")
	}

	list2 := []int{8, 2, 10, 5, 4}
	if ExistsInt(80, list2) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
	}

	if ExistsInt(80, nil) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
	}

	if ExistsInt(80, []int{}) {
		t.Errorf("ExistsInt failed. Expected=false, actual=true")
	}
}

func TestExistsInt64(t *testing.T) {
	// Test : value exist in the list

	list1 := []int64{8, 2, 10, 4}
	if !ExistsInt64(8, list1) {
		t.Errorf("ExistsInt64 failed. Expected=true, actual=false")
	}

	list2 := []int64{8, 2, 10, 5, 4}
	if ExistsInt64(80, list2) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
	}

	if ExistsInt64(80, nil) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
	}

	if ExistsInt64(80, []int64{}) {
		t.Errorf("ExistsInt64 failed. Expected=false, actual=true")
	}
}

func TestExistsInt32(t *testing.T) {
	// Test : value exist in the list

	list1 := []int32{8, 2, 10, 4}
	if !ExistsInt32(8, list1) {
		t.Errorf("ExistsInt32 failed. Expected=true, actual=false")
	}

	list2 := []int32{8, 2, 10, 5, 4}
	if ExistsInt32(80, list2) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}

	if ExistsInt32(80, nil) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}

	if ExistsInt32(80, []int32{}) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}
}

func TestExistsInt16(t *testing.T) {
	// Test : value exist in the list

	list1 := []int16{8, 2, 10, 4}
	if !ExistsInt16(8, list1) {
		t.Errorf("ExistsInt16 failed. Expected=true, actual=false")
	}

	list2 := []int16{8, 2, 10, 5, 4}
	if ExistsInt16(80, list2) {
		t.Errorf("ExistsInt16 failed. Expected=false, actual=true")
	}

	if ExistsInt16(80, nil) {
		t.Errorf("ExistsInt16 failed. Expected=false, actual=true")
	}

	if ExistsInt16(80, []int16{}) {
		t.Errorf("ExistsInt32 failed. Expected=false, actual=true")
	}
}

func TestExistsInt8(t *testing.T) {
	// Test : value exist in the list

	list1 := []int8{8, 2, 10, 4}
	if !ExistsInt8(8, list1) {
		t.Errorf("ExistsInt8 failed. Expected=true, actual=false")
	}

	list2 := []int8{8, 2, 10, 5, 4}
	if ExistsInt8(80, list2) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
	}

	if ExistsInt8(80, nil) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
	}

	if ExistsInt8(80, []int8{}) {
		t.Errorf("ExistsInt8 failed. Expected=false, actual=true")
	}
}

func TestExistsUint64(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint64{8, 2, 10, 4}
	if !ExistsUint64(8, list1) {
		t.Errorf("ExistsUint64 failed. Expected=true, actual=false")
	}

	list2 := []uint64{8, 2, 10, 5, 4}
	if ExistsUint64(80, list2) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
	}

	if ExistsUint64(80, nil) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
	}

	if ExistsUint64(80, []uint64{}) {
		t.Errorf("ExistsUint64 failed. Expected=false, actual=true")
	}
}

func TestExistsUint32(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint32{8, 2, 10, 4}
	if !ExistsUint32(8, list1) {
		t.Errorf("ExistsUint32 failed. Expected=true, actual=false")
	}

	list2 := []uint32{8, 2, 10, 5, 4}
	if ExistsUint32(80, list2) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
	}

	if ExistsUint32(80, nil) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
	}

	if ExistsUint32(80, []uint32{}) {
		t.Errorf("ExistsUint32 failed. Expected=false, actual=true")
	}
}

func TestExistsUint16(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint16{8, 2, 10, 4}
	if !ExistsUint16(8, list1) {
		t.Errorf("ExistsUint16 failed. Expected=true, actual=false")
	}

	list2 := []uint16{8, 2, 10, 5, 4}
	if ExistsUint16(80, list2) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
	}

	if ExistsUint16(80, nil) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
	}

	if ExistsUint16(80, []uint16{}) {
		t.Errorf("ExistsUint16 failed. Expected=false, actual=true")
	}
}

func TestExistsUint8(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint8{8, 2, 10, 4}
	if !ExistsUint8(8, list1) {
		t.Errorf("ExistsUint8 failed. Expected=true, actual=false")
	}

	list2 := []uint8{8, 2, 10, 5, 4}
	if ExistsUint8(80, list2) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
	}

	if ExistsUint8(80, nil) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
	}

	if ExistsUint8(80, []uint8{}) {
		t.Errorf("ExistsUint8 failed. Expected=false, actual=true")
	}
}

func TestExistsUint(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint{8, 2, 10, 4}
	if !ExistsUint(8, list1) {
		t.Errorf("ExistsUint failed. Expected=true, actual=false")
	}

	list2 := []uint{8, 2, 10, 5, 4}
	if ExistsUint(80, list2) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
	}

	if ExistsUint(80, nil) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
	}

	if ExistsUint(80, []uint{}) {
		t.Errorf("ExistsUint failed. Expected=false, actual=true")
	}
}

func TestExistsFloat64(t *testing.T) {
	// Test : value exist in the list

	list1 := []float64{8, 2, 10, 4}
	if !ExistsFloat64(8, list1) {
		t.Errorf("ExistsFloat64 failed. Expected=true, actual=false")
	}

	list2 := []float64{8, 2, 10, 5, 4}
	if ExistsFloat64(80, list2) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
	}

	if ExistsFloat64(80, nil) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
	}

	if ExistsFloat64(80, []float64{}) {
		t.Errorf("ExistsFloat64 failed. Expected=false, actual=true")
	}
}

func TestExistsFloat32(t *testing.T) {
	// Test : value exist in the list

	list1 := []float32{8, 2, 10, 4}
	if !ExistsFloat32(8, list1) {
		t.Errorf("ExistsFloat32 failed. Expected=true, actual=false")
	}

	list2 := []float32{8, 2, 10, 5, 4}
	if ExistsFloat32(80, list2) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
	}

	if ExistsFloat32(80, nil) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
	}

	if ExistsFloat32(80, []float32{}) {
		t.Errorf("ExistsFloat32 failed. Expected=false, actual=true")
	}
}

func TestExistsStr(t *testing.T) {
	// Test : value exist in the list

	list1 := []string{"shyam", "ram", "Hanuman"}
	if !ExistsStr("ram", list1) {
		t.Errorf("ExistsStr failed. Expected=true, actual=false")
	}

	list2 := []string{"shyam", "ram", "Hanuman"}
	if ExistsStr("radha", list2) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
	}

	if ExistsStr("80", nil) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
	}

	if ExistsStr("80", []string{}) {
		t.Errorf("ExistsStr failed. Expected=false, actual=true")
	}
}

func TestExistsStrIgnoreCase(t *testing.T) {
	// Test : value exist in the list

	list1 := []string{"shyam", "ram", "Hanuman"}
	if !ExistsStrIgnoreCase("rAm", list1) {
		t.Errorf("ExistsStrIgnoreCase failed. Expected=true, actual=false")
	}

	list2 := []string{"shyam", "ram", "Hanuman"}
	if ExistsStrIgnoreCase("radha", list2) {
		t.Errorf("ExistsStrIgnoreCase failed. Expected=false, actual=true")
	}

	if ExistsStrIgnoreCase("80", nil) {
		t.Errorf("ExistsStrIgnoreCase failed. Expected=false, actual=true")
	}

	if ExistsStr("80", []string{}) {
		t.Errorf("ExistsStrIgnoreCase failed. Expected=false, actual=true")
	}
}
