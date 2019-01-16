package fp

import (
	"testing"
)

func TestEveryBool(t *testing.T) {
	// Test : every value in the list is either true or false
	list1 := []bool{true, true, true, true}
	if !EveryBool(True, list1) {
		t.Errorf("EveryBool failed. Expected=true, actual=false")
	}

	list1 = []bool{true, false, true, true}
	if EveryBool(True, list1) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	list1 = []bool{true, false, true, true}
	if EveryBool(False, list1) {
		t.Errorf("EveryBool failed. Expected=true, actual=false")
	}

	list1 = []bool{false, false, false, false}
	if !EveryBool(False, list1) {
		t.Errorf("EveryBool failed. Expected=true, actual=false")
	}

	list1 = []bool{}
	if EveryBool(False, list1) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBool(False, nil) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}
}

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

func TestEveryUInt64(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []uint64{8, 2, 10, 4}
	if !EveryUint64(isEvenUint64, list1) {
		t.Errorf("EveryUInt32 failed. Expected=true, actual=false")
	}

	list2 := []uint64{8, 2, 10, 5, 4}
	if EveryUint64(isEvenUint64, list2) {
		t.Errorf("EveryUint64 failed. Expected=false, actual=true")
	}

	if EveryUint64(isEvenUint64, nil) {
		t.Errorf("EveryUint64 failed. Expected=false, actual=true")
	}

	if EveryUint64(isEvenUint64, []uint64{}) {
		t.Errorf("EveryUint64 failed. Expected=false, actual=true")
	}

	if EveryUint64(nil, nil) {
		t.Errorf("EveryUint64 failed. Expected=false, actual=true")
	}
}

func TestEveryUInt32(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []uint32{8, 2, 10, 4}
	if !EveryUint32(isEvenUint32, list1) {
		t.Errorf("EveryUint32 failed. Expected=true, actual=false")
	}

	list2 := []uint32{8, 2, 10, 5, 4}
	if EveryUint32(isEvenUint32, list2) {
		t.Errorf("EveryUint32 failed. Expected=false, actual=true")
	}

	if EveryUint32(isEvenUint32, nil) {
		t.Errorf("EveryUint32 failed. Expected=false, actual=true")
	}

	if EveryUint32(isEvenUint32, []uint32{}) {
		t.Errorf("EveryUint32 failed. Expected=false, actual=true")
	}

	if EveryUint32(nil, nil) {
		t.Errorf("EveryUint32 failed. Expected=false, actual=true")
	}
}

func TestEveryUInt16(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []uint16{8, 2, 10, 4}
	if !EveryUint16(isEvenUint16, list1) {
		t.Errorf("EveryUint16 failed. Expected=true, actual=false")
	}

	list2 := []uint16{8, 2, 10, 5, 4}
	if EveryUint16(isEvenUint16, list2) {
		t.Errorf("EveryUint16 failed. Expected=false, actual=true")
	}

	if EveryUint16(isEvenUint16, nil) {
		t.Errorf("EveryUint16 failed. Expected=false, actual=true")
	}

	if EveryUint16(isEvenUint16, []uint16{}) {
		t.Errorf("EveryUint16 failed. Expected=false, actual=true")
	}

	if EveryUint16(nil, nil) {
		t.Errorf("EveryUint16 failed. Expected=false, actual=true")
	}
}

func TestEveryUInt8(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []uint8{8, 2, 10, 4}
	if !EveryUint8(isEvenUint8, list1) {
		t.Errorf("EveryUint8 failed. Expected=true, actual=false")
	}

	list2 := []uint8{8, 2, 10, 5, 4}
	if EveryUint8(isEvenUint8, list2) {
		t.Errorf("EveryUint8 failed. Expected=false, actual=true")
	}

	if EveryUint8(isEvenUint8, nil) {
		t.Errorf("EveryUint8 failed. Expected=false, actual=true")
	}

	if EveryUint8(isEvenUint8, []uint8{}) {
		t.Errorf("EveryUint8 failed. Expected=false, actual=true")
	}

	if EveryUint8(nil, nil) {
		t.Errorf("EveryUint8 failed. Expected=false, actual=true")
	}
}

func TestEveryUInt(t *testing.T) {
	// Test : every value in the list is even number
	list1 := []uint{8, 2, 10, 4}
	if !EveryUint(isEvenUint, list1) {
		t.Errorf("EveryUint failed. Expected=true, actual=false")
	}

	list2 := []uint{8, 2, 10, 5, 4}
	if EveryUint(isEvenUint, list2) {
		t.Errorf("EveryUint failed. Expected=false, actual=true")
	}

	if EveryUint(isEvenUint, nil) {
		t.Errorf("EveryUint failed. Expected=false, actual=true")
	}

	if EveryUint(isEvenUint, []uint{}) {
		t.Errorf("EveryUint failed. Expected=false, actual=true")
	}

	if EveryUint(nil, nil) {
		t.Errorf("EveryUint failed. Expected=false, actual=true")
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
