package list_op

import "testing"

func TestSomeInt(t *testing.T) {
	// Test : value exist in the list

	list1 := []int{8, 2, 10, 4}
	if !SomeInt(8, list1) {
		t.Errorf("SomeInt failed. Expected=true, actual=false")
	}

	list2 := []int{8, 2, 10, 5, 4}
	if SomeInt(80, list2) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}

	if SomeInt(80, nil) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}

	if SomeInt(80, []int{}) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}
}

func TestSomeInt64(t *testing.T) {
	// Test : value exist in the list

	list1 := []int64{8, 2, 10, 4}
	if !SomeInt64(8, list1) {
		t.Errorf("SomeInt64 failed. Expected=true, actual=false")
	}

	list2 := []int64{8, 2, 10, 5, 4}
	if SomeInt64(80, list2) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}

	if SomeInt64(80, nil) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}

	if SomeInt64(80, []int64{}) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}
}

func TestSomeInt32(t *testing.T) {
	// Test : value exist in the list

	list1 := []int32{8, 2, 10, 4}
	if !SomeInt32(8, list1) {
		t.Errorf("SomeInt32 failed. Expected=true, actual=false")
	}

	list2 := []int32{8, 2, 10, 5, 4}
	if SomeInt32(80, list2) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}

	if SomeInt32(80, nil) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}

	if SomeInt32(80, []int32{}) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}
}

func TestSomeInt16(t *testing.T) {
	// Test : value exist in the list

	list1 := []int16{8, 2, 10, 4}
	if !SomeInt16(8, list1) {
		t.Errorf("SomeInt16 failed. Expected=true, actual=false")
	}

	list2 := []int16{8, 2, 10, 5, 4}
	if SomeInt16(80, list2) {
		t.Errorf("SomeInt16 failed. Expected=false, actual=true")
	}

	if SomeInt16(80, nil) {
		t.Errorf("SomeInt16 failed. Expected=false, actual=true")
	}

	if SomeInt16(80, []int16{}) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}
}

func TestSomeInt8(t *testing.T) {
	// Test : value exist in the list

	list1 := []int8{8, 2, 10, 4}
	if !SomeInt8(8, list1) {
		t.Errorf("SomeInt8 failed. Expected=true, actual=false")
	}

	list2 := []int8{8, 2, 10, 5, 4}
	if SomeInt8(80, list2) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}

	if SomeInt8(80, nil) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}

	if SomeInt8(80, []int8{}) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}
}

func TestSomeUint64(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint64{8, 2, 10, 4}
	if !SomeUint64(8, list1) {
		t.Errorf("SomeUint64 failed. Expected=true, actual=false")
	}

	list2 := []uint64{8, 2, 10, 5, 4}
	if SomeUint64(80, list2) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}

	if SomeUint64(80, nil) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}

	if SomeUint64(80, []uint64{}) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}
}

func TestSomeUint32(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint32{8, 2, 10, 4}
	if !SomeUint32(8, list1) {
		t.Errorf("SomeUint32 failed. Expected=true, actual=false")
	}

	list2 := []uint32{8, 2, 10, 5, 4}
	if SomeUint32(80, list2) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}

	if SomeUint32(80, nil) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}

	if SomeUint32(80, []uint32{}) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}
}

func TestSomeUint16(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint16{8, 2, 10, 4}
	if !SomeUint16(8, list1) {
		t.Errorf("SomeUint16 failed. Expected=true, actual=false")
	}

	list2 := []uint16{8, 2, 10, 5, 4}
	if SomeUint16(80, list2) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}

	if SomeUint16(80, nil) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}

	if SomeUint16(80, []uint16{}) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}
}

func TestSomeUint8(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint8{8, 2, 10, 4}
	if !SomeUint8(8, list1) {
		t.Errorf("SomeUint8 failed. Expected=true, actual=false")
	}

	list2 := []uint8{8, 2, 10, 5, 4}
	if SomeUint8(80, list2) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}

	if SomeUint8(80, nil) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}

	if SomeUint8(80, []uint8{}) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}
}

func TestSomeUint(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint{8, 2, 10, 4}
	if !SomeUint(8, list1) {
		t.Errorf("SomeUint failed. Expected=true, actual=false")
	}

	list2 := []uint{8, 2, 10, 5, 4}
	if SomeUint(80, list2) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}

	if SomeUint(80, nil) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}

	if SomeUint(80, []uint{}) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}
}

func TestSomeFloat64(t *testing.T) {
	// Test : value exist in the list

	list1 := []float64{8, 2, 10, 4}
	if !SomeFloat64(8, list1) {
		t.Errorf("SomeFloat64 failed. Expected=true, actual=false")
	}

	list2 := []float64{8, 2, 10, 5, 4}
	if SomeFloat64(80, list2) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}

	if SomeFloat64(80, nil) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}

	if SomeFloat64(80, []float64{}) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}
}

func TestSomeFloat32(t *testing.T) {
	// Test : value exist in the list

	list1 := []float32{8, 2, 10, 4}
	if !SomeFloat32(8, list1) {
		t.Errorf("SomeFloat32 failed. Expected=true, actual=false")
	}

	list2 := []float32{8, 2, 10, 5, 4}
	if SomeFloat32(80, list2) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}

	if SomeFloat32(80, nil) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}

	if SomeFloat32(80, []float32{}) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}
}

func TestSomeStr(t *testing.T) {
	// Test : value exist in the list

	list1 := []string{"shyam", "ram", "Hanuman"}
	if !SomeStr("ram", list1) {
		t.Errorf("SomeStr failed. Expected=true, actual=false")
	}

	list2 := []string{"shyam", "ram", "Hanuman"}
	if SomeStr("radha", list2) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}

	if SomeStr("80", nil) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}

	if SomeStr("80", []string{}) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}
}

func TestSomeStrIgnoreCase(t *testing.T) {
	// Test : value exist in the list

	list1 := []string{"shyam", "ram", "Hanuman"}
	if !SomeStrIgnoreCase("rAm", list1) {
		t.Errorf("SomeStrIgnoreCase failed. Expected=true, actual=false")
	}

	list2 := []string{"shyam", "ram", "Hanuman"}
	if SomeStrIgnoreCase("radha", list2) {
		t.Errorf("SomeStrIgnoreCase failed. Expected=false, actual=true")
	}

	if SomeStrIgnoreCase("80", nil) {
		t.Errorf("SomeStrIgnoreCase failed. Expected=false, actual=true")
	}

	if SomeStr("80", []string{}) {
		t.Errorf("SomeStrIgnoreCase failed. Expected=false, actual=true")
	}
}
