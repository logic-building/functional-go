package list_op

import "testing"

func TestEveryInt(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int{8, 2, 10, 4}
	if !EveryInt(isEven, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []int{8, 2, 10, 5, 4}
	if EveryInt(isEven, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}

	if EveryInt(isEven, nil) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}

	if EveryInt(isEven, []int{}) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}

	if EveryInt(nil, []int{}) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryInt64(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int64{8, 2, 10, 4}
	if !EveryInt64(isEvenInt64, list1) {
		t.Errorf("EveryInt64 failed. Expected=true, actual=false")
	}

	list2 := []int64{8, 2, 10, 5, 4}
	if EveryInt64(isEvenInt64, list2) {
		t.Errorf("EveryInt64 failed. Expected=false, actual=true")
	}

	if EveryInt64(isEvenInt64, nil) {
		t.Errorf("EveryInt64 failed. Expected=false, actual=true")
	}

	if EveryInt64(isEvenInt64, []int64{}) {
		t.Errorf("EveryInt64 failed. Expected=false, actual=true")
	}

	if EveryInt64(nil, nil) {
		t.Errorf("EveryInt64 failed. Expected=false, actual=true")
	}
}

func TestEveryInt32(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int32{8, 2, 10, 4}
	if !EveryInt32(isEvenInt32, list1) {
		t.Errorf("EveryInt32 failed. Expected=true, actual=false")
	}

	list2 := []int32{8, 2, 10, 5, 4}
	if EveryInt32(isEvenInt32, list2) {
		t.Errorf("EveryInt32 failed. Expected=false, actual=true")
	}

	if EveryInt32(isEvenInt32, nil) {
		t.Errorf("EveryInt32 failed. Expected=false, actual=true")
	}

	if EveryInt32(isEvenInt32, []int32{}) {
		t.Errorf("EveryInt32 failed. Expected=false, actual=true")
	}

	if EveryInt32(nil, nil) {
		t.Errorf("EveryInt32 failed. Expected=false, actual=true")
	}
}

func TestEveryInt16(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int16{8, 2, 10, 4}
	if !EveryInt16(isEvenInt16, list1) {
		t.Errorf("EveryInt16 failed. Expected=true, actual=false")
	}

	list2 := []int16{8, 2, 10, 5, 4}
	if EveryInt16(isEvenInt16, list2) {
		t.Errorf("EveryInt16 failed. Expected=false, actual=true")
	}

	if EveryInt16(isEvenInt16, nil) {
		t.Errorf("EveryInt16 failed. Expected=false, actual=true")
	}

	if EveryInt16(isEvenInt16, []int16{}) {
		t.Errorf("EveryInt16 failed. Expected=false, actual=true")
	}

	if EveryInt16(nil, nil) {
		t.Errorf("EveryInt16 failed. Expected=false, actual=true")
	}
}

func TestEveryInt8(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int8{8, 2, 10, 4}
	if !EveryInt8(isEvenInt8, list1) {
		t.Errorf("EveryInt8 failed. Expected=true, actual=false")
	}

	list2 := []int8{8, 2, 10, 5, 4}
	if EveryInt8(isEvenInt8, list2) {
		t.Errorf("EveryInt8 failed. Expected=false, actual=true")
	}

	if EveryInt8(isEvenInt8, nil) {
		t.Errorf("EveryInt8 failed. Expected=false, actual=true")
	}

	if EveryInt8(isEvenInt8, []int8{}) {
		t.Errorf("EveryInt8 failed. Expected=false, actual=true")
	}

	if EveryInt8(nil, nil) {
		t.Errorf("EveryInt8 failed. Expected=false, actual=true")
	}
}

func TestEveryFloat64(t *testing.T) {
	// Test : every value in the list is positive number
	list1 := []float64{8.2, 2.3, 10.4, 4}
	if !EveryFloat64(isPositiveFloat64, list1) {
		t.Errorf("EveryFloat64 failed. Expected=true, actual=false")
	}

	list2 := []float64{8, 2, 10, 5, 7, -1, 4}
	if EveryFloat64(isPositiveFloat64, list2) {
		t.Errorf("EveryFloat64 failed. Expected=false, actual=true")
	}

	if EveryFloat64(isPositiveFloat64, nil) {
		t.Errorf("EveryFloat64 failed. Expected=false, actual=true")
	}

	if EveryFloat64(isPositiveFloat64, []float64{}) {
		t.Errorf("EveryFloat64 failed. Expected=false, actual=true")
	}

	if EveryFloat64(nil, nil) {
		t.Errorf("EveryFloat64 failed. Expected=false, actual=true")
	}
}

func TestEveryFloat32(t *testing.T) {
	// Test : every value in the list is positive number
	list1 := []float32{8.2, 2.3, 10.4, 4}
	if !EveryFloat32(isPositiveFloat32, list1) {
		t.Errorf("EveryFloat32 failed. Expected=true, actual=false")
	}

	list2 := []float32{8, 2, 10, 5, 7, -1, 4}
	if EveryFloat32(isPositiveFloat32, list2) {
		t.Errorf("EveryFloat32 failed. Expected=false, actual=true")
	}

	if EveryFloat32(isPositiveFloat32, nil) {
		t.Errorf("EveryFloat32 failed. Expected=false, actual=true")
	}

	if EveryFloat32(isPositiveFloat32, []float32{}) {
		t.Errorf("EveryFloat32 failed. Expected=false, actual=true")
	}

	if EveryFloat32(nil, nil) {
		t.Errorf("EveryFloat32 failed. Expected=false, actual=true")
	}
}

func TestEveryStr(t *testing.T) {
	// Test : length of every string in the list is 3
	list1 := []string{"Ram", "Raj", "Sai"}
	if !EveryStr(isStrLen3, list1) {
		t.Errorf("EveryStr failed. Expected=true, actual=false")
	}

	list2 := []string{"Ram", "Krishna", "Sai"}
	if EveryStr(isStrLen3, list2) {
		t.Errorf("EveryStr failed. Expected=false, actual=true")
	}

	if EveryStr(isStrLen3, nil) {
		t.Errorf("EveryStr failed. Expected=false, actual=true")
	}

	if EveryStr(isStrLen3, []string{}) {
		t.Errorf("EveryStr failed. Expected=false, actual=true")
	}

	if EveryStr(nil, nil) {
		t.Errorf("EveryStr failed. Expected=false, actual=true")
	}
}

func isStrLen3(str string) bool {
	return len(str) == 3
}
