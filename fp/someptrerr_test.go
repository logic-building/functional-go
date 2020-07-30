package fp

import (
	"errors"
	"testing"
)

func TestSomeIntPtrErr(t *testing.T) {
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

	list1 := []*int{&v8, &v2, &v10, &v4}
	r, _ := SomeIntPtrErr(isEvenIntPtrErr, list1)
	if !r {
		t.Errorf("SomeIntPtrErr failed. Expected=true, actual=false")
	}

	list3 := []*int{&v0, &v4}
	_, err := SomeIntPtrErr(isEvenIntPtrErr, list3)
	if err == nil {
		t.Errorf("SomeIntPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int{&v1, &v3, &v5, &v7}
	r, _ = SomeIntPtrErr(isEvenIntPtrErr, list2)
	if r {
		t.Errorf("SomeIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeIntPtrErr(nil, nil)
	if r {
		t.Errorf("SomeIntPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeIntPtrErr(isEvenIntPtrErr, []*int{})
	if r {
		t.Errorf("SomeIntPtr failed. Expected=false, actual=true")
	}
}

func TestSomeInt64PtrErr(t *testing.T) {
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

	list1 := []*int64{&v8, &v2, &v10, &v4}
	r, _ := SomeInt64PtrErr(isEvenInt64PtrErr, list1)
	if !r {
		t.Errorf("SomeInt64PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*int64{&v0, &v4}
	_, err := SomeInt64PtrErr(isEvenInt64PtrErr, list3)
	if err == nil {
		t.Errorf("SomeInt64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int64{&v1, &v3, &v5, &v7}
	r, _ = SomeInt64PtrErr(isEvenInt64PtrErr, list2)
	if r {
		t.Errorf("SomeInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt64PtrErr(nil, nil)
	if r {
		t.Errorf("SomeInt64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt64PtrErr(isEvenInt64PtrErr, []*int64{})
	if r {
		t.Errorf("SomeInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt32PtrErr(t *testing.T) {
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

	list1 := []*int32{&v8, &v2, &v10, &v4}
	r, _ := SomeInt32PtrErr(isEvenInt32PtrErr, list1)
	if !r {
		t.Errorf("SomeInt32PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*int32{&v0, &v4}
	_, err := SomeInt32PtrErr(isEvenInt32PtrErr, list3)
	if err == nil {
		t.Errorf("SomeInt32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int32{&v1, &v3, &v5, &v7}
	r, _ = SomeInt32PtrErr(isEvenInt32PtrErr, list2)
	if r {
		t.Errorf("SomeInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt32PtrErr(nil, nil)
	if r {
		t.Errorf("SomeInt32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt32PtrErr(isEvenInt32PtrErr, []*int32{})
	if r {
		t.Errorf("SomeInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt16PtrErr(t *testing.T) {
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

	list1 := []*int16{&v8, &v2, &v10, &v4}
	r, _ := SomeInt16PtrErr(isEvenInt16PtrErr, list1)
	if !r {
		t.Errorf("SomeInt16PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*int16{&v0, &v4}
	_, err := SomeInt16PtrErr(isEvenInt16PtrErr, list3)
	if err == nil {
		t.Errorf("SomeInt16PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int16{&v1, &v3, &v5, &v7}
	r, _ = SomeInt16PtrErr(isEvenInt16PtrErr, list2)
	if r {
		t.Errorf("SomeInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt16PtrErr(nil, nil)
	if r {
		t.Errorf("SomeInt16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt16PtrErr(isEvenInt16PtrErr, []*int16{})
	if r {
		t.Errorf("SomeInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeInt8PtrErr(t *testing.T) {
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

	list1 := []*int8{&v8, &v2, &v10, &v4}
	r, _ := SomeInt8PtrErr(isEvenInt8PtrErr, list1)
	if !r {
		t.Errorf("SomeInt8PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*int8{&v0, &v4}
	_, err := SomeInt8PtrErr(isEvenInt8PtrErr, list3)
	if err == nil {
		t.Errorf("SomeInt8PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*int8{&v1, &v3, &v5, &v7}
	r, _ = SomeInt8PtrErr(isEvenInt8PtrErr, list2)
	if r {
		t.Errorf("SomeInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt8PtrErr(nil, nil)
	if r {
		t.Errorf("SomeInt8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeInt8PtrErr(isEvenInt8PtrErr, []*int8{})
	if r {
		t.Errorf("SomeInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUintPtrErr(t *testing.T) {
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

	list1 := []*uint{&v8, &v2, &v10, &v4}
	r, _ := SomeUintPtrErr(isEvenUintPtrErr, list1)
	if !r {
		t.Errorf("SomeUintPtrErr failed. Expected=true, actual=false")
	}

	list3 := []*uint{&v0, &v4}
	_, err := SomeUintPtrErr(isEvenUintPtrErr, list3)
	if err == nil {
		t.Errorf("SomeUintPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint{&v1, &v3, &v5, &v7}
	r, _ = SomeUintPtrErr(isEvenUintPtrErr, list2)
	if r {
		t.Errorf("SomeUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUintPtrErr(nil, nil)
	if r {
		t.Errorf("SomeUintPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUintPtrErr(isEvenUintPtrErr, []*uint{})
	if r {
		t.Errorf("SomeUintPtr failed. Expected=false, actual=true")
	}
}

func TestSomeUint64PtrErr(t *testing.T) {
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

	list1 := []*uint64{&v8, &v2, &v10, &v4}
	r, _ := SomeUint64PtrErr(isEvenUint64PtrErr, list1)
	if !r {
		t.Errorf("SomeUint64PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*uint64{&v0, &v4}
	_, err := SomeUint64PtrErr(isEvenUint64PtrErr, list3)
	if err == nil {
		t.Errorf("SomeUint64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint64{&v1, &v3, &v5, &v7}
	r, _ = SomeUint64PtrErr(isEvenUint64PtrErr, list2)
	if r {
		t.Errorf("SomeUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint64PtrErr(nil, nil)
	if r {
		t.Errorf("SomeUint64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint64PtrErr(isEvenUint64PtrErr, []*uint64{})
	if r {
		t.Errorf("SomeUint64Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint32PtrErr(t *testing.T) {
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

	list1 := []*uint32{&v8, &v2, &v10, &v4}
	r, _ := SomeUint32PtrErr(isEvenUint32PtrErr, list1)
	if !r {
		t.Errorf("SomeUint32PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*uint32{&v0, &v4}
	_, err := SomeUint32PtrErr(isEvenUint32PtrErr, list3)
	if err == nil {
		t.Errorf("SomeUint32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint32{&v1, &v3, &v5, &v7}
	r, _ = SomeUint32PtrErr(isEvenUint32PtrErr, list2)
	if r {
		t.Errorf("SomeUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint32PtrErr(nil, nil)
	if r {
		t.Errorf("SomeUint32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint32PtrErr(isEvenUint32PtrErr, []*uint32{})
	if r {
		t.Errorf("SomeUint32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint16PtrErr(t *testing.T) {
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

	list1 := []*uint16{&v8, &v2, &v10, &v4}
	r, _ := SomeUint16PtrErr(isEvenUint16PtrErr, list1)
	if !r {
		t.Errorf("SomeUint16PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*uint16{&v0, &v4}
	_, err := SomeUint16PtrErr(isEvenUint16PtrErr, list3)
	if err == nil {
		t.Errorf("SomeUint16PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint16{&v1, &v3, &v5, &v7}
	r, _ = SomeUint16PtrErr(isEvenUint16PtrErr, list2)
	if r {
		t.Errorf("SomeUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint16PtrErr(nil, nil)
	if r {
		t.Errorf("SomeUint16PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint16PtrErr(isEvenUint16PtrErr, []*uint16{})
	if r {
		t.Errorf("SomeUint16Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeUint8PtrErr(t *testing.T) {
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

	list1 := []*uint8{&v8, &v2, &v10, &v4}
	r, _ := SomeUint8PtrErr(isEvenUint8PtrErr, list1)
	if !r {
		t.Errorf("SomeUint8PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*uint8{&v0, &v4}
	_, err := SomeUint8PtrErr(isEvenUint8PtrErr, list3)
	if err == nil {
		t.Errorf("SomeUint8PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*uint8{&v1, &v3, &v5, &v7}
	r, _ = SomeUint8PtrErr(isEvenUint8PtrErr, list2)
	if r {
		t.Errorf("SomeUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint8PtrErr(nil, nil)
	if r {
		t.Errorf("SomeUint8PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeUint8PtrErr(isEvenUint8PtrErr, []*uint8{})
	if r {
		t.Errorf("SomeUint8Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeStrPtrErr(t *testing.T) {
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

	list1 := []*string{&v8, &v2, &v10, &v4}
	r, _ := SomeStrPtrErr(isEvenStrPtrErr, list1)
	if !r {
		t.Errorf("SomeStrPtrErr failed. Expected=true, actual=false")
	}

	list3 := []*string{&v0, &v4}
	_, err := SomeStrPtrErr(isEvenStrPtrErr, list3)
	if err == nil {
		t.Errorf("SomeStrPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v1, &v3, &v5, &v7}
	r, _ = SomeStrPtrErr(isEvenStrPtrErr, list2)
	if r {
		t.Errorf("SomeStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeStrPtrErr(nil, nil)
	if r {
		t.Errorf("SomeStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeStrPtrErr(isEvenStrPtrErr, []*string{})
	if r {
		t.Errorf("SomeStrPtr failed. Expected=false, actual=true")
	}
}

func TestSomeBoolPtrErr(t *testing.T) {
	// Test : value exist in the list

	var vt bool = true
	var vf bool = false

	list1 := []*bool{&vt, &vf}
	r, _ := SomeBoolPtrErr(func(v *bool) (bool, error) { return *v == true, nil }, list1)
	if !r {
		t.Errorf("SomeBoolPtrErr failed. Expected=true, actual=false")
	}

	r, _ = SomeBoolPtrErr(nil, nil)

	if r {
		t.Errorf("SomeBoolPtr failed. Expected=false, actual=true")
	}

	r, _ = SomeBoolPtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*bool{})
	if r {
		t.Errorf("SomeBoolPtrErr failed. Expected=false, actual=true")
	}

	_, err := SomeBoolPtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid in this test")
		}
		return *v == true, nil
	}, []*bool{&vf})

	if err == nil {
		t.Errorf("SomeBoolPtrErr failed. Expected=false, actual=true")
	}
}

func TestSomeFloat32PtrErr(t *testing.T) {
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

	list1 := []*float32{&v8, &v2, &v10, &v4}
	r, _ := SomeFloat32PtrErr(isEvenFloat32PtrErr, list1)
	if !r {
		t.Errorf("SomeFloat32PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*float32{&v0, &v4}
	_, err := SomeFloat32PtrErr(isEvenFloat32PtrErr, list3)
	if err == nil {
		t.Errorf("SomeFloat32PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*float32{&v1, &v3, &v5, &v7}
	r, _ = SomeFloat32PtrErr(isEvenFloat32PtrErr, list2)
	if r {
		t.Errorf("SomeFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat32PtrErr(nil, nil)
	if r {
		t.Errorf("SomeFloat32PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat32PtrErr(isEvenFloat32PtrErr, []*float32{})
	if r {
		t.Errorf("SomeFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestSomeFloat64PtrErr(t *testing.T) {
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

	list1 := []*float64{&v8, &v2, &v10, &v4}
	r, _ := SomeFloat64PtrErr(isEvenFloat64PtrErr, list1)
	if !r {
		t.Errorf("SomeFloat64PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*float64{&v0, &v4}
	_, err := SomeFloat64PtrErr(isEvenFloat64PtrErr, list3)
	if err == nil {
		t.Errorf("SomeFloat64PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*float64{&v1, &v3, &v5, &v7}
	r, _ = SomeFloat64PtrErr(isEvenFloat64PtrErr, list2)
	if r {
		t.Errorf("SomeFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat64PtrErr(nil, nil)
	if r {
		t.Errorf("SomeFloat64PtrErr failed. Expected=false, actual=true")
	}

	r, _ = SomeFloat64PtrErr(isEvenFloat64PtrErr, []*float64{})
	if r {
		t.Errorf("SomeFloat64Ptr failed. Expected=false, actual=true")
	}
}
