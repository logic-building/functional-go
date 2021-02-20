package fp

import (
	"errors"
	"testing"
)

func TestEveryIntPtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v8 int = 8
	var v10 int = 10
	var v0 int

	list1 := []*int{&v2, &v4}

	r, _ := EveryIntPtrErr(isEvenIntPtrErr, list1)
	if !r {
		t.Errorf("EveryIntPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*int{&v0, &v4}
	_, err := EveryIntPtrErr(isEvenIntPtrErr, list1)
	if err == nil {
		t.Errorf("EveryIntPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryIntPtrErr(isEvenIntPtrErr, list2)
	if r {
		t.Errorf("EveryIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntPtrErr(isEvenIntPtrErr, nil)
	if r {
		t.Errorf("EveryIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntPtrErr(isEvenIntPtrErr, []*int{})
	if r {
		t.Errorf("EveryIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryIntPtrErr(nil, []*int{})
	if r {
		t.Errorf("EveryIntPtr failed. Expected=false, actual=true")
	}
}

func TestEveryInt64PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v8 int64 = 8
	var v10 int64 = 10
	var v0 int64

	list1 := []*int64{&v2, &v4}

	r, _ := EveryInt64PtrErr(isEvenInt64PtrErr, list1)
	if !r {
		t.Errorf("EveryInt64PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*int64{&v0, &v4}
	_, err := EveryInt64PtrErr(isEvenInt64PtrErr, list1)
	if err == nil {
		t.Errorf("EveryInt64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int64{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryInt64PtrErr(isEvenInt64PtrErr, list2)
	if r {
		t.Errorf("EveryInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64PtrErr(isEvenInt64PtrErr, nil)
	if r {
		t.Errorf("EveryInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64PtrErr(isEvenInt64PtrErr, []*int64{})
	if r {
		t.Errorf("EveryInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt64PtrErr(nil, []*int64{})
	if r {
		t.Errorf("EveryInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryInt32PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v8 int32 = 8
	var v10 int32 = 10
	var v0 int32

	list1 := []*int32{&v2, &v4}

	r, _ := EveryInt32PtrErr(isEvenInt32PtrErr, list1)
	if !r {
		t.Errorf("EveryInt32PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*int32{&v0, &v4}
	_, err := EveryInt32PtrErr(isEvenInt32PtrErr, list1)
	if err == nil {
		t.Errorf("EveryInt32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int32{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryInt32PtrErr(isEvenInt32PtrErr, list2)
	if r {
		t.Errorf("EveryInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32PtrErr(isEvenInt32PtrErr, nil)
	if r {
		t.Errorf("EveryInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32PtrErr(isEvenInt32PtrErr, []*int32{})
	if r {
		t.Errorf("EveryInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt32PtrErr(nil, []*int32{})
	if r {
		t.Errorf("EveryInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryInt16PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v8 int16 = 8
	var v10 int16 = 10
	var v0 int16

	list1 := []*int16{&v2, &v4}

	r, _ := EveryInt16PtrErr(isEvenInt16PtrErr, list1)
	if !r {
		t.Errorf("EveryInt16PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*int16{&v0, &v4}
	_, err := EveryInt16PtrErr(isEvenInt16PtrErr, list1)
	if err == nil {
		t.Errorf("EveryInt16PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int16{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryInt16PtrErr(isEvenInt16PtrErr, list2)
	if r {
		t.Errorf("EveryInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16PtrErr(isEvenInt16PtrErr, nil)
	if r {
		t.Errorf("EveryInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16PtrErr(isEvenInt16PtrErr, []*int16{})
	if r {
		t.Errorf("EveryInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt16PtrErr(nil, []*int16{})
	if r {
		t.Errorf("EveryInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryInt8PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v8 int8 = 8
	var v10 int8 = 10
	var v0 int8

	list1 := []*int8{&v2, &v4}

	r, _ := EveryInt8PtrErr(isEvenInt8PtrErr, list1)
	if !r {
		t.Errorf("EveryInt8PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*int8{&v0, &v4}
	_, err := EveryInt8PtrErr(isEvenInt8PtrErr, list1)
	if err == nil {
		t.Errorf("EveryInt8PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int8{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryInt8PtrErr(isEvenInt8PtrErr, list2)
	if r {
		t.Errorf("EveryInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8PtrErr(isEvenInt8PtrErr, nil)
	if r {
		t.Errorf("EveryInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8PtrErr(isEvenInt8PtrErr, []*int8{})
	if r {
		t.Errorf("EveryInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryInt8PtrErr(nil, []*int8{})
	if r {
		t.Errorf("EveryInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryUintPtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v8 uint = 8
	var v10 uint = 10
	var v0 uint

	list1 := []*uint{&v2, &v4}

	r, _ := EveryUintPtrErr(isEvenUintPtrErr, list1)
	if !r {
		t.Errorf("EveryUintPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*uint{&v0, &v4}
	_, err := EveryUintPtrErr(isEvenUintPtrErr, list1)
	if err == nil {
		t.Errorf("EveryUintPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryUintPtrErr(isEvenUintPtrErr, list2)
	if r {
		t.Errorf("EveryUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintPtrErr(isEvenUintPtrErr, nil)
	if r {
		t.Errorf("EveryUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintPtrErr(isEvenUintPtrErr, []*uint{})
	if r {
		t.Errorf("EveryUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUintPtrErr(nil, []*uint{})
	if r {
		t.Errorf("EveryUintPtr failed. Expected=false, actual=true")
	}
}

func TestEveryUint64PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v8 uint64 = 8
	var v10 uint64 = 10
	var v0 uint64

	list1 := []*uint64{&v2, &v4}

	r, _ := EveryUint64PtrErr(isEvenUint64PtrErr, list1)
	if !r {
		t.Errorf("EveryUint64PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*uint64{&v0, &v4}
	_, err := EveryUint64PtrErr(isEvenUint64PtrErr, list1)
	if err == nil {
		t.Errorf("EveryUint64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint64{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryUint64PtrErr(isEvenUint64PtrErr, list2)
	if r {
		t.Errorf("EveryUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64PtrErr(isEvenUint64PtrErr, nil)
	if r {
		t.Errorf("EveryUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64PtrErr(isEvenUint64PtrErr, []*uint64{})
	if r {
		t.Errorf("EveryUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint64PtrErr(nil, []*uint64{})
	if r {
		t.Errorf("EveryUint64Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryUint32PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v8 uint32 = 8
	var v10 uint32 = 10
	var v0 uint32

	list1 := []*uint32{&v2, &v4}

	r, _ := EveryUint32PtrErr(isEvenUint32PtrErr, list1)
	if !r {
		t.Errorf("EveryUint32PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*uint32{&v0, &v4}
	_, err := EveryUint32PtrErr(isEvenUint32PtrErr, list1)
	if err == nil {
		t.Errorf("EveryUint32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint32{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryUint32PtrErr(isEvenUint32PtrErr, list2)
	if r {
		t.Errorf("EveryUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32PtrErr(isEvenUint32PtrErr, nil)
	if r {
		t.Errorf("EveryUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32PtrErr(isEvenUint32PtrErr, []*uint32{})
	if r {
		t.Errorf("EveryUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint32PtrErr(nil, []*uint32{})
	if r {
		t.Errorf("EveryUint32Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryUint16PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v8 uint16 = 8
	var v10 uint16 = 10
	var v0 uint16

	list1 := []*uint16{&v2, &v4}

	r, _ := EveryUint16PtrErr(isEvenUint16PtrErr, list1)
	if !r {
		t.Errorf("EveryUint16PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*uint16{&v0, &v4}
	_, err := EveryUint16PtrErr(isEvenUint16PtrErr, list1)
	if err == nil {
		t.Errorf("EveryUint16PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint16{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryUint16PtrErr(isEvenUint16PtrErr, list2)
	if r {
		t.Errorf("EveryUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16PtrErr(isEvenUint16PtrErr, nil)
	if r {
		t.Errorf("EveryUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16PtrErr(isEvenUint16PtrErr, []*uint16{})
	if r {
		t.Errorf("EveryUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint16PtrErr(nil, []*uint16{})
	if r {
		t.Errorf("EveryUint16Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryUint8PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v8 uint8 = 8
	var v10 uint8 = 10
	var v0 uint8

	list1 := []*uint8{&v2, &v4}

	r, _ := EveryUint8PtrErr(isEvenUint8PtrErr, list1)
	if !r {
		t.Errorf("EveryUint8PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*uint8{&v0, &v4}
	_, err := EveryUint8PtrErr(isEvenUint8PtrErr, list1)
	if err == nil {
		t.Errorf("EveryUint8PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint8{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryUint8PtrErr(isEvenUint8PtrErr, list2)
	if r {
		t.Errorf("EveryUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8PtrErr(isEvenUint8PtrErr, nil)
	if r {
		t.Errorf("EveryUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8PtrErr(isEvenUint8PtrErr, []*uint8{})
	if r {
		t.Errorf("EveryUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryUint8PtrErr(nil, []*uint8{})
	if r {
		t.Errorf("EveryUint8Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryStrPtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string = "0"

	list1 := []*string{&v2, &v4}

	r, _ := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if !r {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*string{&v0, &v4}
	_, err := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if err == nil {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryStrPtrErr(isEvenStrPtrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, nil)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, []*string{})
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(nil, []*string{})
	if r {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}
}

func TestEveryBoolPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	// Test : every value in the list is either true or false
	list1 := []*bool{&vt, &vt, &vt, &vt}
	r, _ := EveryBoolPtrErr(TruePtrErr, list1)
	if !r {
		t.Errorf("EveryBoolPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*bool{&vt, &vt, &vt}
	r, _ = EveryBoolPtrErr(TruePtrErr, list1)
	if !r {
		t.Errorf("EveryBoolPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*bool{&vt, &vt, &vt, &vf}
	_, err := EveryBoolPtrErr(TruePtrErr, list1)
	if err == nil {
		t.Errorf("EveryBoolPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*bool{&vt, &vt, &vf}
	r, _ = EveryBoolPtrErr(TruePtrErr2, list1)
	if r {
		t.Errorf("EveryBoolPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*bool{}
	r, _ = EveryBoolPtrErr(TruePtrErr, list1)
	if r {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(TruePtrErr, nil)
	if r {
		t.Errorf("EveryBoolPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(nil, []*bool{})
	if r {
		t.Errorf("EveryBoolPtr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(TruePtrErr, []*bool{})
	if r {
		t.Errorf("EveryBoolPtrErr failed. Expected=false, actual=true")
	}
}

func TruePtrErr(val *bool) (bool, error) {
	if *val == false {
		return false, errors.New("false is invalid in this test")
	}
	return *val == true, nil
}

func TruePtrErr2(val *bool) (bool, error) {
	return *val == true, nil
}

func TestEveryFloat32PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v8 float32 = 8
	var v10 float32 = 10
	var v0 float32

	list1 := []*float32{&v2, &v4}

	r, _ := EveryFloat32PtrErr(isEvenFloat32PtrErr, list1)
	if !r {
		t.Errorf("EveryFloat32PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*float32{&v0, &v4}
	_, err := EveryFloat32PtrErr(isEvenFloat32PtrErr, list1)
	if err == nil {
		t.Errorf("EveryFloat32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*float32{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryFloat32PtrErr(isEvenFloat32PtrErr, list2)
	if r {
		t.Errorf("EveryFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32PtrErr(isEvenFloat32PtrErr, nil)
	if r {
		t.Errorf("EveryFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32PtrErr(isEvenFloat32PtrErr, []*float32{})
	if r {
		t.Errorf("EveryFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat32PtrErr(nil, []*float32{})
	if r {
		t.Errorf("EveryFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestEveryFloat64PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v8 float64 = 8
	var v10 float64 = 10
	var v0 float64

	list1 := []*float64{&v2, &v4}

	r, _ := EveryFloat64PtrErr(isEvenFloat64PtrErr, list1)
	if !r {
		t.Errorf("EveryFloat64PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*float64{&v0, &v4}
	_, err := EveryFloat64PtrErr(isEvenFloat64PtrErr, list1)
	if err == nil {
		t.Errorf("EveryFloat64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*float64{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryFloat64PtrErr(isEvenFloat64PtrErr, list2)
	if r {
		t.Errorf("EveryFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64PtrErr(isEvenFloat64PtrErr, nil)
	if r {
		t.Errorf("EveryFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64PtrErr(isEvenFloat64PtrErr, []*float64{})
	if r {
		t.Errorf("EveryFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryFloat64PtrErr(nil, []*float64{})
	if r {
		t.Errorf("EveryFloat64Ptr failed. Expected=false, actual=true")
	}
}
