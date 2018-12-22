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
}

func TestEveryInt64(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int64{8, 2, 10, 4}
	if !EveryInt64(isEvenInt64, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []int64{8, 2, 10, 5, 4}
	if EveryInt64(isEvenInt64, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryInt32(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int32{8, 2, 10, 4}
	if !EveryInt32(isEvenInt32, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []int32{8, 2, 10, 5, 4}
	if EveryInt32(isEvenInt32, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryInt16(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int16{8, 2, 10, 4}
	if !EveryInt16(isEvenInt16, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []int16{8, 2, 10, 5, 4}
	if EveryInt16(isEvenInt16, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryInt8(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []int8{8, 2, 10, 4}
	if !EveryInt8(isEvenInt8, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []int8{8, 2, 10, 5, 4}
	if EveryInt8(isEvenInt8, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryFloat64(t *testing.T) {
	// Test : every value in the list is positive number
	list1 := []float64{8.2, 2.3, 10.4, 4}
	if !EveryFloat64(isPositiveFloat64, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []float64{8, 2, 10, 5, 7, -1, 4}
	if EveryFloat64(isPositiveFloat64, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryFloat32(t *testing.T) {
	// Test : every value in the list is positive number
	list1 := []float32{8.2, 2.3, 10.4, 4}
	if !EveryFloat32(isPositiveFloat32, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []float32{8, 2, 10, 5, 7, -1, 4}
	if EveryFloat32(isPositiveFloat32, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func TestEveryStr(t *testing.T) {
	// Test : length of every string in the list is 3
	list1 := []string{"Ram", "Raj", "Sai"}
	if !EveryStr(isStrLen3, list1) {
		t.Errorf("EveryInt failed. Expected=true, actual=false")
	}

	list2 := []string{"Ram", "Krishna", "Sai"}
	if EveryStr(isStrLen3, list2) {
		t.Errorf("EveryInt failed. Expected=false, actual=true")
	}
}

func isStrLen3(str string) bool {
	return len(str) == 3
}
