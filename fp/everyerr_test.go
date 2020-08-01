package fp

import (
	"errors"
	"testing"
)

func TestEveryIntPErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v8 int = 8
	var v10 int = 10
	var v0 int

	list1 := []int{v2, v4}

	r, _ := EveryIntErr(isEvenIntErr, list1)
	if !r {
		t.Errorf("EveryIntErr failed. Expected=true, actual=false")
	}

	list1 = []int{v0, v4}
	_, err := EveryIntErr(isEvenIntErr, list1)
	if err == nil {
		t.Errorf("EveryIntErr failed. Expected=true, actual=false")
	}

	list2 := []int{v8, v2, v10, v5, v4}
	r, _ = EveryIntErr(isEvenIntErr, list2)
	if r {
		t.Errorf("EveryIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntErr(isEvenIntErr, nil)
	if r {
		t.Errorf("EveryIntErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntErr(isEvenIntErr, []int{})
	if r {
		t.Errorf("EveryIntErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntErr(nil, []int{})
	if r {
		t.Errorf("EveryIntErr failed. Expected=false, actual=true")
	}
}

func TestEveryInt64Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v8 int64 = 8
	var v10 int64 = 10
	var v0 int64

	list1 := []int64{v2, v4}

	r, _ := EveryInt64Err(isEvenInt64Err, list1)
	if !r {
		t.Errorf("EveryInt64Err failed. Expected=true, actual=false")
	}

	list1 = []int64{v0, v4}
	_, err := EveryInt64Err(isEvenInt64Err, list1)
	if err == nil {
		t.Errorf("EveryInt64Err failed. Expected=true, actual=false")
	}

	list2 := []int64{v8, v2, v10, v5, v4}
	r, _ = EveryInt64Err(isEvenInt64Err, list2)
	if r {
		t.Errorf("EveryInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64Err(isEvenInt64Err, nil)
	if r {
		t.Errorf("EveryInt64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64Err(isEvenInt64Err, []int64{})
	if r {
		t.Errorf("EveryInt64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64Err(nil, []int64{})
	if r {
		t.Errorf("EveryInt64Err failed. Expected=false, actual=true")
	}
}

func TestEveryInt32Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v8 int32 = 8
	var v10 int32 = 10
	var v0 int32

	list1 := []int32{v2, v4}

	r, _ := EveryInt32Err(isEvenInt32Err, list1)
	if !r {
		t.Errorf("EveryInt32Err failed. Expected=true, actual=false")
	}

	list1 = []int32{v0, v4}
	_, err := EveryInt32Err(isEvenInt32Err, list1)
	if err == nil {
		t.Errorf("EveryInt32Err failed. Expected=true, actual=false")
	}

	list2 := []int32{v8, v2, v10, v5, v4}
	r, _ = EveryInt32Err(isEvenInt32Err, list2)
	if r {
		t.Errorf("EveryInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32Err(isEvenInt32Err, nil)
	if r {
		t.Errorf("EveryInt32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32Err(isEvenInt32Err, []int32{})
	if r {
		t.Errorf("EveryInt32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32Err(nil, []int32{})
	if r {
		t.Errorf("EveryInt32Err failed. Expected=false, actual=true")
	}
}

func TestEveryInt16Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v8 int16 = 8
	var v10 int16 = 10
	var v0 int16

	list1 := []int16{v2, v4}

	r, _ := EveryInt16Err(isEvenInt16Err, list1)
	if !r {
		t.Errorf("EveryInt16Err failed. Expected=true, actual=false")
	}

	list1 = []int16{v0, v4}
	_, err := EveryInt16Err(isEvenInt16Err, list1)
	if err == nil {
		t.Errorf("EveryInt16Err failed. Expected=true, actual=false")
	}

	list2 := []int16{v8, v2, v10, v5, v4}
	r, _ = EveryInt16Err(isEvenInt16Err, list2)
	if r {
		t.Errorf("EveryInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16Err(isEvenInt16Err, nil)
	if r {
		t.Errorf("EveryInt16Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16Err(isEvenInt16Err, []int16{})
	if r {
		t.Errorf("EveryInt16Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16Err(nil, []int16{})
	if r {
		t.Errorf("EveryInt16Err failed. Expected=false, actual=true")
	}
}

func TestEveryInt8Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v8 int8 = 8
	var v10 int8 = 10
	var v0 int8

	list1 := []int8{v2, v4}

	r, _ := EveryInt8Err(isEvenInt8Err, list1)
	if !r {
		t.Errorf("EveryInt8Err failed. Expected=true, actual=false")
	}

	list1 = []int8{v0, v4}
	_, err := EveryInt8Err(isEvenInt8Err, list1)
	if err == nil {
		t.Errorf("EveryInt8Err failed. Expected=true, actual=false")
	}

	list2 := []int8{v8, v2, v10, v5, v4}
	r, _ = EveryInt8Err(isEvenInt8Err, list2)
	if r {
		t.Errorf("EveryInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8Err(isEvenInt8Err, nil)
	if r {
		t.Errorf("EveryInt8Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8Err(isEvenInt8Err, []int8{})
	if r {
		t.Errorf("EveryInt8Err failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8Err(nil, []int8{})
	if r {
		t.Errorf("EveryInt8Err failed. Expected=false, actual=true")
	}
}

func TestEveryUintErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v8 uint = 8
	var v10 uint = 10
	var v0 uint

	list1 := []uint{v2, v4}

	r, _ := EveryUintErr(isEvenUintErr, list1)
	if !r {
		t.Errorf("EveryUintErr failed. Expected=true, actual=false")
	}

	list1 = []uint{v0, v4}
	_, err := EveryUintErr(isEvenUintErr, list1)
	if err == nil {
		t.Errorf("EveryUintErr failed. Expected=true, actual=false")
	}

	list2 := []uint{v8, v2, v10, v5, v4}
	r, _ = EveryUintErr(isEvenUintErr, list2)
	if r {
		t.Errorf("EveryUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintErr(isEvenUintErr, nil)
	if r {
		t.Errorf("EveryUintErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintErr(isEvenUintErr, []uint{})
	if r {
		t.Errorf("EveryUintErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintErr(nil, []uint{})
	if r {
		t.Errorf("EveryUintErr failed. Expected=false, actual=true")
	}
}

func TestEveryUint64Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v8 uint64 = 8
	var v10 uint64 = 10
	var v0 uint64

	list1 := []uint64{v2, v4}

	r, _ := EveryUint64Err(isEvenUint64Err, list1)
	if !r {
		t.Errorf("EveryUint64Err failed. Expected=true, actual=false")
	}

	list1 = []uint64{v0, v4}
	_, err := EveryUint64Err(isEvenUint64Err, list1)
	if err == nil {
		t.Errorf("EveryUint64Err failed. Expected=true, actual=false")
	}

	list2 := []uint64{v8, v2, v10, v5, v4}
	r, _ = EveryUint64Err(isEvenUint64Err, list2)
	if r {
		t.Errorf("EveryUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64Err(isEvenUint64Err, nil)
	if r {
		t.Errorf("EveryUint64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64Err(isEvenUint64Err, []uint64{})
	if r {
		t.Errorf("EveryUint64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64Err(nil, []uint64{})
	if r {
		t.Errorf("EveryUint64Err failed. Expected=false, actual=true")
	}
}

func TestEveryUint32Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v8 uint32 = 8
	var v10 uint32 = 10
	var v0 uint32

	list1 := []uint32{v2, v4}

	r, _ := EveryUint32Err(isEvenUint32Err, list1)
	if !r {
		t.Errorf("EveryUint32Err failed. Expected=true, actual=false")
	}

	list1 = []uint32{v0, v4}
	_, err := EveryUint32Err(isEvenUint32Err, list1)
	if err == nil {
		t.Errorf("EveryUint32Err failed. Expected=true, actual=false")
	}

	list2 := []uint32{v8, v2, v10, v5, v4}
	r, _ = EveryUint32Err(isEvenUint32Err, list2)
	if r {
		t.Errorf("EveryUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32Err(isEvenUint32Err, nil)
	if r {
		t.Errorf("EveryUint32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32Err(isEvenUint32Err, []uint32{})
	if r {
		t.Errorf("EveryUint32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32Err(nil, []uint32{})
	if r {
		t.Errorf("EveryUint32Err failed. Expected=false, actual=true")
	}
}

func TestEveryUint16Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v8 uint16 = 8
	var v10 uint16 = 10
	var v0 uint16

	list1 := []uint16{v2, v4}

	r, _ := EveryUint16Err(isEvenUint16Err, list1)
	if !r {
		t.Errorf("EveryUint16Err failed. Expected=true, actual=false")
	}

	list1 = []uint16{v0, v4}
	_, err := EveryUint16Err(isEvenUint16Err, list1)
	if err == nil {
		t.Errorf("EveryUint16Err failed. Expected=true, actual=false")
	}

	list2 := []uint16{v8, v2, v10, v5, v4}
	r, _ = EveryUint16Err(isEvenUint16Err, list2)
	if r {
		t.Errorf("EveryUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16Err(isEvenUint16Err, nil)
	if r {
		t.Errorf("EveryUint16Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16Err(isEvenUint16Err, []uint16{})
	if r {
		t.Errorf("EveryUint16Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16Err(nil, []uint16{})
	if r {
		t.Errorf("EveryUint16Err failed. Expected=false, actual=true")
	}
}

func TestEveryUint8Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v8 uint8 = 8
	var v10 uint8 = 10
	var v0 uint8

	list1 := []uint8{v2, v4}

	r, _ := EveryUint8Err(isEvenUint8Err, list1)
	if !r {
		t.Errorf("EveryUint8Err failed. Expected=true, actual=false")
	}

	list1 = []uint8{v0, v4}
	_, err := EveryUint8Err(isEvenUint8Err, list1)
	if err == nil {
		t.Errorf("EveryUint8Err failed. Expected=true, actual=false")
	}

	list2 := []uint8{v8, v2, v10, v5, v4}
	r, _ = EveryUint8Err(isEvenUint8Err, list2)
	if r {
		t.Errorf("EveryUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8Err(isEvenUint8Err, nil)
	if r {
		t.Errorf("EveryUint8Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8Err(isEvenUint8Err, []uint8{})
	if r {
		t.Errorf("EveryUint8Err failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8Err(nil, []uint8{})
	if r {
		t.Errorf("EveryUint8Err failed. Expected=false, actual=true")
	}
}

func TestEveryStrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string = "0"

	list1 := []string{v2, v4}

	r, _ := EveryStrErr(isEvenStrErr, list1)
	if !r {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list1 = []string{v0, v4}
	_, err := EveryStrErr(isEvenStrErr, list1)
	if err == nil {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list2 := []string{v8, v2, v10, v5, v4}
	r, _ = EveryStrErr(isEvenStrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, nil)
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(nil, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}
}

func TestEveryBoolErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	// Test : every value in the list is either true or false
	list1 := []bool{vt, vt, vt, vt}
	r, _ := EveryBoolErr(TrueErr, list1)
	if !r {
		t.Errorf("EveryBoolErr failed. Expected=true, actual=false")
	}

	list1 = []bool{vt, vt, vt}
	r, _ = EveryBoolErr(TrueErr, list1)
	if !r {
		t.Errorf("EveryBoolErr failed. Expected=true, actual=false")
	}

	list1 = []bool{vt, vt, vt, vf}
	_, err := EveryBoolErr(TrueErr, list1)
	if err == nil {
		t.Errorf("EveryBoolErr failed. Expected=true, actual=false")
	}

	list1 = []bool{vt, vt, vf}
	r, _ = EveryBoolErr(TrueErr2, list1)
	if r {
		t.Errorf("EveryBoolErr failed. Expected=true, actual=false")
	}

	list1 = []bool{}
	r, _ = EveryBoolErr(TrueErr, list1)
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(TrueErr, nil)
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(nil, []bool{})
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(TrueErr, []bool{})
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}
}

func TrueErr(val bool) (bool, error) {
	if val == false {
		return false, errors.New("false is invalid in this test")
	}
	return val == true, nil
}

func TrueErr2(val bool) (bool, error) {
	return val == true, nil
}

func TestEveryFloat32Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v8 float32 = 8
	var v10 float32 = 10
	var v0 float32

	list1 := []float32{v2, v4}

	r, _ := EveryFloat32Err(isEvenFloat32Err, list1)
	if !r {
		t.Errorf("EveryFloat32Err failed. Expected=true, actual=false")
	}

	list1 = []float32{v0, v4}
	_, err := EveryFloat32Err(isEvenFloat32Err, list1)
	if err == nil {
		t.Errorf("EveryFloat32Err failed. Expected=true, actual=false")
	}

	list2 := []float32{v8, v2, v10, v5, v4}
	r, _ = EveryFloat32Err(isEvenFloat32Err, list2)
	if r {
		t.Errorf("EveryFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32Err(isEvenFloat32Err, nil)
	if r {
		t.Errorf("EveryFloat32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32Err(isEvenFloat32Err, []float32{})
	if r {
		t.Errorf("EveryFloat32Err failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32Err(nil, []float32{})
	if r {
		t.Errorf("EveryFloat32Err failed. Expected=false, actual=true")
	}
}

func TestEveryFloat64Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v8 float64 = 8
	var v10 float64 = 10
	var v0 float64

	list1 := []float64{v2, v4}

	r, _ := EveryFloat64Err(isEvenFloat64Err, list1)
	if !r {
		t.Errorf("EveryFloat64Err failed. Expected=true, actual=false")
	}

	list1 = []float64{v0, v4}
	_, err := EveryFloat64Err(isEvenFloat64Err, list1)
	if err == nil {
		t.Errorf("EveryFloat64Err failed. Expected=true, actual=false")
	}

	list2 := []float64{v8, v2, v10, v5, v4}
	r, _ = EveryFloat64Err(isEvenFloat64Err, list2)
	if r {
		t.Errorf("EveryFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64Err(isEvenFloat64Err, nil)
	if r {
		t.Errorf("EveryFloat64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64Err(isEvenFloat64Err, []float64{})
	if r {
		t.Errorf("EveryFloat64Err failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64Err(nil, []float64{})
	if r {
		t.Errorf("EveryFloat64Err failed. Expected=false, actual=true")
	}
}
