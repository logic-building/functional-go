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
