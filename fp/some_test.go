package fp

import (
	"testing"
)

func TestSomeInt(t *testing.T) {
	// Test : value exist in the list

	list1 := []int{8, 2, 10, 4}
	if !SomeInt(isEven, list1) {
		t.Errorf("SomeInt failed. Expected=true, actual=false")
	}

	list2 := []int{1, 3, 5, 7}
	if SomeInt(isEven, list2) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}

	if SomeInt(nil, nil) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}

	if SomeInt(isEven, []int{}) {
		t.Errorf("SomeInt failed. Expected=false, actual=true")
	}
}

func TestSomeInt64(t *testing.T) {
	// Test : value exist in the list

	list1 := []int64{8, 2, 10, 4}
	if !SomeInt64(isEvenInt64, list1) {
		t.Errorf("SomeInt64 failed. Expected=true, actual=false")
	}

	list2 := []int64{1, 3, 5, 7}
	if SomeInt64(isEvenInt64, list2) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}

	if SomeInt64(nil, nil) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}

	if SomeInt64(isEvenInt64, []int64{}) {
		t.Errorf("SomeInt64 failed. Expected=false, actual=true")
	}
}

func TestSomeInt32(t *testing.T) {
	// Test : value exist in the list

	list1 := []int32{8, 2, 10, 4}
	if !SomeInt32(isEvenInt32, list1) {
		t.Errorf("SomeInt32 failed. Expected=true, actual=false")
	}

	list2 := []int32{1, 3, 5, 7}
	if SomeInt32(isEvenInt32, list2) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}

	if SomeInt32(nil, nil) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}

	if SomeInt32(isEvenInt32, []int32{}) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}
}

func TestSomeInt16(t *testing.T) {
	// Test : value exist in the list

	list1 := []int16{8, 2, 10, 4}
	if !SomeInt16(isEvenInt16, list1) {
		t.Errorf("SomeInt16 failed. Expected=true, actual=false")
	}

	list2 := []int16{1, 3, 5, 7}
	if SomeInt16(isEvenInt16, list2) {
		t.Errorf("SomeInt16 failed. Expected=false, actual=true")
	}

	if SomeInt16(nil, nil) {
		t.Errorf("SomeInt16 failed. Expected=false, actual=true")
	}

	if SomeInt16(isEvenInt16, []int16{}) {
		t.Errorf("SomeInt32 failed. Expected=false, actual=true")
	}
}

func TestSomeInt8(t *testing.T) {
	// Test : value exist in the list

	list1 := []int8{8, 2, 10, 4}
	if !SomeInt8(isEvenInt8, list1) {
		t.Errorf("SomeInt8 failed. Expected=true, actual=false")
	}

	list2 := []int8{1, 3, 5, 7}
	if SomeInt8(isEvenInt8, list2) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}

	if SomeInt8(nil, nil) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}

	if SomeInt8(isEvenInt8, []int8{}) {
		t.Errorf("SomeInt8 failed. Expected=false, actual=true")
	}
}

func TestSomeUint64(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint64{8, 2, 10, 4}
	if !SomeUint64(isEvenUint64, list1) {
		t.Errorf("SomeUint64 failed. Expected=true, actual=false")
	}

	list2 := []uint64{1, 3, 5, 7}
	if SomeUint64(isEvenUint64, list2) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}

	if SomeUint64(nil, nil) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}

	if SomeUint64(isEvenUint64, []uint64{}) {
		t.Errorf("SomeUint64 failed. Expected=false, actual=true")
	}
}

func TestSomeUint32(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint32{8, 2, 10, 4}
	if !SomeUint32(isEvenUint32, list1) {
		t.Errorf("SomeUint32 failed. Expected=true, actual=false")
	}

	list2 := []uint32{1, 3, 5, 7}
	if SomeUint32(isEvenUint32, list2) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}

	if SomeUint32(nil, nil) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}

	if SomeUint32(isEvenUint32, []uint32{}) {
		t.Errorf("SomeUint32 failed. Expected=false, actual=true")
	}
}

func TestSomeUint16(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint16{8, 2, 10, 4}
	if !SomeUint16(isEvenUint16, list1) {
		t.Errorf("SomeUint16 failed. Expected=true, actual=false")
	}

	list2 := []uint16{1, 3, 5, 7}
	if SomeUint16(isEvenUint16, list2) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}

	if SomeUint16(isEvenUint16, nil) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}

	if SomeUint16(nil, []uint16{}) {
		t.Errorf("SomeUint16 failed. Expected=false, actual=true")
	}
}

func TestSomeUint8(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint8{8, 2, 10, 4}
	if !SomeUint8(isEvenUint8, list1) {
		t.Errorf("SomeUint8 failed. Expected=true, actual=false")
	}

	list2 := []uint8{1, 3, 5, 7}
	if SomeUint8(isEvenUint8, list2) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}

	if SomeUint8(nil, nil) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}

	if SomeUint8(isEvenUint8, []uint8{}) {
		t.Errorf("SomeUint8 failed. Expected=false, actual=true")
	}
}

func TestSomeUint(t *testing.T) {
	// Test : value exist in the list

	list1 := []uint{8, 2, 10, 4}
	if !SomeUint(isEvenUint, list1) {
		t.Errorf("SomeUint failed. Expected=true, actual=false")
	}

	list2 := []uint{1, 3, 5, 7}
	if SomeUint(isEvenUint, list2) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}

	if SomeUint(nil, nil) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}

	if SomeUint(isEvenUint, []uint{}) {
		t.Errorf("SomeUint failed. Expected=false, actual=true")
	}
}

func TestSomeFloat64(t *testing.T) {
	// Test : value exist in the list

	list1 := []float64{8, 2, 10, 4}
	if !SomeFloat64(isPositiveFloat64, list1) {
		t.Errorf("SomeFloat64 failed. Expected=true, actual=false")
	}

	list2 := []float64{-1, -3, -5, -7}
	if SomeFloat64(isPositiveFloat64, list2) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}

	if SomeFloat64(nil, nil) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}

	if SomeFloat64(isPositiveFloat64, []float64{}) {
		t.Errorf("SomeFloat64 failed. Expected=false, actual=true")
	}
}

func TestSomeFloat32(t *testing.T) {
	// Test : value exist in the list

	list1 := []float32{8, 2, 10, 4}
	if !SomeFloat32(isPositiveFloat32, list1) {
		t.Errorf("SomeFloat32 failed. Expected=true, actual=false")
	}

	list2 := []float32{-1, -3, -5, -7}
	if SomeFloat32(isPositiveFloat32, list2) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}

	if SomeFloat32(nil, nil) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}

	if SomeFloat32(isPositiveFloat32, []float32{}) {
		t.Errorf("SomeFloat32 failed. Expected=false, actual=true")
	}
}

func TestSomeStr(t *testing.T) {
	// Test : value exist in the list

	list1 := []string{"shyam", "ram", "Hanuman"}
	if !SomeStr(isStrLen3, list1) {
		t.Errorf("SomeStr failed. Expected=true, actual=false")
	}

	list2 := []string{"shyam", "rama", "Hanuman"}
	if SomeStr(isStrLen3, list2) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}

	if SomeStr(nil, nil) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}

	if SomeStr(isStrLen3, []string{}) {
		t.Errorf("SomeStr failed. Expected=false, actual=true")
	}
}
