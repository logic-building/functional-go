package fp

import (
	"errors"
	"testing"
)

func TestSomeIntErr(t *testing.T) {
	// Test : value exist in the list

	var v0 int = 0
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	var v7 int = 7
	var v8 int = 8
	var v10 int = 10

	list1 := []int{v8, v2, v10, v4}
	r, _ := SomeIntErr(isEvenIntErr, list1)
	if !r {
		t.Errorf("SomeIntErr failed. Expected=true, actual=false")
	}

	list3 := []int{v0, v4}
	_, err := SomeIntErr(isEvenIntErr, list3)
	if err == nil {
		t.Errorf("SomeIntErr failed. Expected=true, actual=false")
	}

	list2 := []int{v1, v3, v5, v7}
	r, _ = SomeIntErr(isEvenIntErr, list2)
	if r {
		t.Errorf("SomeIntErr failed. Expected=false, actual=true")
	}

	r, _ = SomeIntErr(nil, nil)
	if r {
		t.Errorf("SomeIntErr failed. Expected=false, actual=true")
	}

	r, _ = SomeIntErr(isEvenIntErr, []int{})
	if r {
		t.Errorf("SomeIntPtr failed. Expected=false, actual=true")
	}
}

func TestSomeInt64Err(t *testing.T) {
	// Test : value exist in the list

	var v0 int64 = 0
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	var v7 int64 = 7
	var v8 int64 = 8
	var v10 int64 = 10

	list1 := []int64{v8, v2, v10, v4}
	r, _ := SomeInt64Err(isEvenInt64Err, list1)
	if !r {
		t.Errorf("SomeInt64Err failed. Expected=true, actual=false")
	}

	list3 := []int64{v0, v4}
	_, err := SomeInt64Err(isEvenInt64Err, list3)
	if err == nil {
		t.Errorf("SomeInt64Err failed. Expected=true, actual=false")
	}

	list2 := []int64{v1, v3, v5, v7}
	r, _ = SomeInt64Err(isEvenInt64Err, list2)
	if r {
		t.Errorf("SomeInt64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt64Err(nil, nil)
	if r {
		t.Errorf("SomeInt64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt64Err(isEvenInt64Err, []int64{})
	if r {
		t.Errorf("SomeInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt32Err(t *testing.T) {
	// Test : value exist in the list

	var v0 int32 = 0
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	var v7 int32 = 7
	var v8 int32 = 8
	var v10 int32 = 10

	list1 := []int32{v8, v2, v10, v4}
	r, _ := SomeInt32Err(isEvenInt32Err, list1)
	if !r {
		t.Errorf("SomeInt32Err failed. Expected=true, actual=false")
	}

	list3 := []int32{v0, v4}
	_, err := SomeInt32Err(isEvenInt32Err, list3)
	if err == nil {
		t.Errorf("SomeInt32Err failed. Expected=true, actual=false")
	}

	list2 := []int32{v1, v3, v5, v7}
	r, _ = SomeInt32Err(isEvenInt32Err, list2)
	if r {
		t.Errorf("SomeInt32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt32Err(nil, nil)
	if r {
		t.Errorf("SomeInt32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt32Err(isEvenInt32Err, []int32{})
	if r {
		t.Errorf("SomeInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt16Err(t *testing.T) {
	// Test : value exist in the list

	var v0 int16 = 0
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	var v7 int16 = 7
	var v8 int16 = 8
	var v10 int16 = 10

	list1 := []int16{v8, v2, v10, v4}
	r, _ := SomeInt16Err(isEvenInt16Err, list1)
	if !r {
		t.Errorf("SomeInt16Err failed. Expected=true, actual=false")
	}

	list3 := []int16{v0, v4}
	_, err := SomeInt16Err(isEvenInt16Err, list3)
	if err == nil {
		t.Errorf("SomeInt16Err failed. Expected=true, actual=false")
	}

	list2 := []int16{v1, v3, v5, v7}
	r, _ = SomeInt16Err(isEvenInt16Err, list2)
	if r {
		t.Errorf("SomeInt16Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt16Err(nil, nil)
	if r {
		t.Errorf("SomeInt16Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt16Err(isEvenInt16Err, []int16{})
	if r {
		t.Errorf("SomeInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt8Err(t *testing.T) {
	// Test : value exist in the list

	var v0 int8 = 0
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	var v7 int8 = 7
	var v8 int8 = 8
	var v10 int8 = 10

	list1 := []int8{v8, v2, v10, v4}
	r, _ := SomeInt8Err(isEvenInt8Err, list1)
	if !r {
		t.Errorf("SomeInt8Err failed. Expected=true, actual=false")
	}

	list3 := []int8{v0, v4}
	_, err := SomeInt8Err(isEvenInt8Err, list3)
	if err == nil {
		t.Errorf("SomeInt8Err failed. Expected=true, actual=false")
	}

	list2 := []int8{v1, v3, v5, v7}
	r, _ = SomeInt8Err(isEvenInt8Err, list2)
	if r {
		t.Errorf("SomeInt8Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt8Err(nil, nil)
	if r {
		t.Errorf("SomeInt8Err failed. Expected=false, actual=true")
	}

	r, _ = SomeInt8Err(isEvenInt8Err, []int8{})
	if r {
		t.Errorf("SomeInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUintErr(t *testing.T) {
	// Test : value exist in the list

	var v0 uint = 0
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	var v7 uint = 7
	var v8 uint = 8
	var v10 uint = 10

	list1 := []uint{v8, v2, v10, v4}
	r, _ := SomeUintErr(isEvenUintErr, list1)
	if !r {
		t.Errorf("SomeUintErr failed. Expected=true, actual=false")
	}

	list3 := []uint{v0, v4}
	_, err := SomeUintErr(isEvenUintErr, list3)
	if err == nil {
		t.Errorf("SomeUintErr failed. Expected=true, actual=false")
	}

	list2 := []uint{v1, v3, v5, v7}
	r, _ = SomeUintErr(isEvenUintErr, list2)
	if r {
		t.Errorf("SomeUintErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUintErr(nil, nil)
	if r {
		t.Errorf("SomeUintErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUintErr(isEvenUintErr, []uint{})
	if r {
		t.Errorf("SomeUintPtr failed. Expected=false, actual=true")
	}
}

func TestSomeUint64Err(t *testing.T) {
	// Test : value exist in the list

	var v0 uint64 = 0
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v7 uint64 = 7
	var v8 uint64 = 8
	var v10 uint64 = 10

	list1 := []uint64{v8, v2, v10, v4}
	r, _ := SomeUint64Err(isEvenUint64Err, list1)
	if !r {
		t.Errorf("SomeUint64Err failed. Expected=true, actual=false")
	}

	list3 := []uint64{v0, v4}
	_, err := SomeUint64Err(isEvenUint64Err, list3)
	if err == nil {
		t.Errorf("SomeUint64Err failed. Expected=true, actual=false")
	}

	list2 := []uint64{v1, v3, v5, v7}
	r, _ = SomeUint64Err(isEvenUint64Err, list2)
	if r {
		t.Errorf("SomeUint64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint64Err(nil, nil)
	if r {
		t.Errorf("SomeUint64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint64Err(isEvenUint64Err, []uint64{})
	if r {
		t.Errorf("SomeUint64Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint32Err(t *testing.T) {
	// Test : value exist in the list

	var v0 uint32 = 0
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v7 uint32 = 7
	var v8 uint32 = 8
	var v10 uint32 = 10

	list1 := []uint32{v8, v2, v10, v4}
	r, _ := SomeUint32Err(isEvenUint32Err, list1)
	if !r {
		t.Errorf("SomeUint32Err failed. Expected=true, actual=false")
	}

	list3 := []uint32{v0, v4}
	_, err := SomeUint32Err(isEvenUint32Err, list3)
	if err == nil {
		t.Errorf("SomeUint32Err failed. Expected=true, actual=false")
	}

	list2 := []uint32{v1, v3, v5, v7}
	r, _ = SomeUint32Err(isEvenUint32Err, list2)
	if r {
		t.Errorf("SomeUint32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint32Err(nil, nil)
	if r {
		t.Errorf("SomeUint32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint32Err(isEvenUint32Err, []uint32{})
	if r {
		t.Errorf("SomeUint32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint16Err(t *testing.T) {
	// Test : value exist in the list

	var v0 uint16 = 0
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v7 uint16 = 7
	var v8 uint16 = 8
	var v10 uint16 = 10

	list1 := []uint16{v8, v2, v10, v4}
	r, _ := SomeUint16Err(isEvenUint16Err, list1)
	if !r {
		t.Errorf("SomeUint16Err failed. Expected=true, actual=false")
	}

	list3 := []uint16{v0, v4}
	_, err := SomeUint16Err(isEvenUint16Err, list3)
	if err == nil {
		t.Errorf("SomeUint16Err failed. Expected=true, actual=false")
	}

	list2 := []uint16{v1, v3, v5, v7}
	r, _ = SomeUint16Err(isEvenUint16Err, list2)
	if r {
		t.Errorf("SomeUint16Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint16Err(nil, nil)
	if r {
		t.Errorf("SomeUint16Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint16Err(isEvenUint16Err, []uint16{})
	if r {
		t.Errorf("SomeUint16Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint8Err(t *testing.T) {
	// Test : value exist in the list

	var v0 uint8 = 0
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v7 uint8 = 7
	var v8 uint8 = 8
	var v10 uint8 = 10

	list1 := []uint8{v8, v2, v10, v4}
	r, _ := SomeUint8Err(isEvenUint8Err, list1)
	if !r {
		t.Errorf("SomeUint8Err failed. Expected=true, actual=false")
	}

	list3 := []uint8{v0, v4}
	_, err := SomeUint8Err(isEvenUint8Err, list3)
	if err == nil {
		t.Errorf("SomeUint8Err failed. Expected=true, actual=false")
	}

	list2 := []uint8{v1, v3, v5, v7}
	r, _ = SomeUint8Err(isEvenUint8Err, list2)
	if r {
		t.Errorf("SomeUint8Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint8Err(nil, nil)
	if r {
		t.Errorf("SomeUint8Err failed. Expected=false, actual=true")
	}

	r, _ = SomeUint8Err(isEvenUint8Err, []uint8{})
	if r {
		t.Errorf("SomeUint8Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeStrErr(t *testing.T) {
	// Test : value exist in the list

	var v0 string = "0"
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v8 string = "8"
	var v10 string = "10"

	list1 := []string{v8, v2, v10, v4}
	r, _ := SomeStrErr(isEvenStrErr, list1)
	if !r {
		t.Errorf("SomeStrErr failed. Expected=true, actual=false")
	}

	list3 := []string{v0, v4}
	_, err := SomeStrErr(isEvenStrErr, list3)
	if err == nil {
		t.Errorf("SomeStrErr failed. Expected=true, actual=false")
	}

	list2 := []string{v1, v3, v5, v7}
	r, _ = SomeStrErr(isEvenStrErr, list2)
	if r {
		t.Errorf("SomeStrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeStrErr(nil, nil)
	if r {
		t.Errorf("SomeStrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeStrErr(isEvenStrErr, []string{})
	if r {
		t.Errorf("SomeStrPtr failed. Expected=false, actual=true")
	}
}

func TestSomeBoolErr(t *testing.T) {
	// Test : value exist in the list

	var vt bool = true
	var vf bool = false

	list1 := []bool{vt, vf}
	r, _ := SomeBoolErr(func(v bool) (bool, error) { return v == true, nil }, list1)
	if !r {
		t.Errorf("SomeBool Err failed. Expected=true, actual=false")
	}

	r, _ = SomeBoolErr(nil, nil)
	
	if r{
		t.Errorf("SomeBoolErr failed. Expected=false, actual=true")
	}

	r, _ = SomeBoolErr(func(v bool) (bool, error) { return v == true, nil }, []bool{})
	if  r {
		t.Errorf("SomeBoolErr failed. Expected=false, actual=true")
	}

	_, err := SomeBoolErr(func(v bool) (bool, error) { if v == false { return false, errors.New("false is invalid in this test") }; return v == true, nil }, []bool{vf})
	if  err == nil {
		t.Errorf("SomeBoolErr failed. Expected=false, actual=true")
	}
}

func TestSomeFloat32Err(t *testing.T) {
	// Test : value exist in the list

	var v0 float32 = 0
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	var v7 float32 = 7
	var v8 float32 = 8
	var v10 float32 = 10

	list1 := []float32{v8, v2, v10, v4}
	r, _ := SomeFloat32Err(isEvenFloat32Err, list1)
	if !r {
		t.Errorf("SomeFloat32Err failed. Expected=true, actual=false")
	}

	list3 := []float32{v0, v4}
	_, err := SomeFloat32Err(isEvenFloat32Err, list3)
	if err == nil {
		t.Errorf("SomeFloat32Err failed. Expected=true, actual=false")
	}

	list2 := []float32{v1, v3, v5, v7}
	r, _ = SomeFloat32Err(isEvenFloat32Err, list2)
	if r {
		t.Errorf("SomeFloat32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat32Err(nil, nil)
	if r {
		t.Errorf("SomeFloat32Err failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat32Err(isEvenFloat32Err, []float32{})
	if r {
		t.Errorf("SomeFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeFloat64Err(t *testing.T) {
	// Test : value exist in the list

	var v0 float64 = 0
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	var v7 float64 = 7
	var v8 float64 = 8
	var v10 float64 = 10

	list1 := []float64{v8, v2, v10, v4}
	r, _ := SomeFloat64Err(isEvenFloat64Err, list1)
	if !r {
		t.Errorf("SomeFloat64Err failed. Expected=true, actual=false")
	}

	list3 := []float64{v0, v4}
	_, err := SomeFloat64Err(isEvenFloat64Err, list3)
	if err == nil {
		t.Errorf("SomeFloat64Err failed. Expected=true, actual=false")
	}

	list2 := []float64{v1, v3, v5, v7}
	r, _ = SomeFloat64Err(isEvenFloat64Err, list2)
	if r {
		t.Errorf("SomeFloat64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat64Err(nil, nil)
	if r {
		t.Errorf("SomeFloat64Err failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat64Err(isEvenFloat64Err, []float64{})
	if r {
		t.Errorf("SomeFloat64Ptr failed. Expected=false, actual=true")
	}
}
