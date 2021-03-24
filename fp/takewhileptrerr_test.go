package fp

import (
	"errors"
	"testing"
)

func TestTakeWhileIntPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v7 int = 7
	var v40 int = 40
	var v0 int

	expectedNewList := []*int{&v4, &v2, &v4}
	NewList, _ := TakeWhileIntPtrErr(isEvenIntPtrErr, []*int{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileIntPtrErr(isEvenIntPtrErr, []*int{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int{&v40}
	partialIsEvenDivisibleBy := func(num *int) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileIntPtrErr(partialIsEvenDivisibleBy, []*int{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileIntPtrErr failed.")
	}

	r, _ = TakeWhileIntPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("TakeWhileIntPtr failed.")
	}
}

func TestTakeWhileInt64PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v7 int64 = 7
	var v40 int64 = 40
	var v0 int64

	expectedNewList := []*int64{&v4, &v2, &v4}
	NewList, _ := TakeWhileInt64PtrErr(isEvenInt64PtrErr, []*int64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt64PtrErr(isEvenInt64PtrErr, []*int64{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int64{&v40}
	partialIsEvenDivisibleBy := func(num *int64) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileInt64PtrErr(partialIsEvenDivisibleBy, []*int64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt64PtrErr failed.")
	}

	r, _ = TakeWhileInt64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt64Ptr failed.")
	}
}

func TestTakeWhileInt32PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v7 int32 = 7
	var v40 int32 = 40
	var v0 int32

	expectedNewList := []*int32{&v4, &v2, &v4}
	NewList, _ := TakeWhileInt32PtrErr(isEvenInt32PtrErr, []*int32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt32PtrErr(isEvenInt32PtrErr, []*int32{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int32{&v40}
	partialIsEvenDivisibleBy := func(num *int32) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileInt32PtrErr(partialIsEvenDivisibleBy, []*int32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt32PtrErr failed.")
	}

	r, _ = TakeWhileInt32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt32Ptr failed.")
	}
}

func TestTakeWhileInt16PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v7 int16 = 7
	var v40 int16 = 40
	var v0 int16

	expectedNewList := []*int16{&v4, &v2, &v4}
	NewList, _ := TakeWhileInt16PtrErr(isEvenInt16PtrErr, []*int16{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt16PtrErr(isEvenInt16PtrErr, []*int16{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int16{&v40}
	partialIsEvenDivisibleBy := func(num *int16) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileInt16PtrErr(partialIsEvenDivisibleBy, []*int16{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt16PtrErr failed.")
	}

	r, _ = TakeWhileInt16PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt16Ptr failed.")
	}
}

func TestTakeWhileInt8PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v7 int8 = 7
	var v40 int8 = 40
	var v0 int8

	expectedNewList := []*int8{&v4, &v2, &v4}
	NewList, _ := TakeWhileInt8PtrErr(isEvenInt8PtrErr, []*int8{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileInt8PtrErr(isEvenInt8PtrErr, []*int8{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int8{&v40}
	partialIsEvenDivisibleBy := func(num *int8) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileInt8PtrErr(partialIsEvenDivisibleBy, []*int8{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileInt8PtrErr failed.")
	}

	r, _ = TakeWhileInt8PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("TakeWhileInt8Ptr failed.")
	}
}

func TestTakeWhileUintPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v7 uint = 7
	var v40 uint = 40
	var v0 uint

	expectedNewList := []*uint{&v4, &v2, &v4}
	NewList, _ := TakeWhileUintPtrErr(isEvenUintPtrErr, []*uint{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUintPtrErr(isEvenUintPtrErr, []*uint{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint{&v40}
	partialIsEvenDivisibleBy := func(num *uint) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileUintPtrErr(partialIsEvenDivisibleBy, []*uint{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUintPtrErr failed.")
	}

	r, _ = TakeWhileUintPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUintPtr failed.")
	}
}

func TestTakeWhileUint64PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v7 uint64 = 7
	var v40 uint64 = 40
	var v0 uint64

	expectedNewList := []*uint64{&v4, &v2, &v4}
	NewList, _ := TakeWhileUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint64{&v40}
	partialIsEvenDivisibleBy := func(num *uint64) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileUint64PtrErr(partialIsEvenDivisibleBy, []*uint64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint64PtrErr failed.")
	}

	r, _ = TakeWhileUint64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint64Ptr failed.")
	}
}

func TestTakeWhileUint32PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v7 uint32 = 7
	var v40 uint32 = 40
	var v0 uint32

	expectedNewList := []*uint32{&v4, &v2, &v4}
	NewList, _ := TakeWhileUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint32{&v40}
	partialIsEvenDivisibleBy := func(num *uint32) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileUint32PtrErr(partialIsEvenDivisibleBy, []*uint32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint32PtrErr failed.")
	}

	r, _ = TakeWhileUint32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint32Ptr failed.")
	}
}

func TestTakeWhileUint16PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v7 uint16 = 7
	var v40 uint16 = 40
	var v0 uint16

	expectedNewList := []*uint16{&v4, &v2, &v4}
	NewList, _ := TakeWhileUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint16{&v40}
	partialIsEvenDivisibleBy := func(num *uint16) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileUint16PtrErr(partialIsEvenDivisibleBy, []*uint16{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint16PtrErr failed.")
	}

	r, _ = TakeWhileUint16PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint16Ptr failed.")
	}
}

func TestTakeWhileUint8PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v7 uint8 = 7
	var v40 uint8 = 40
	var v0 uint8

	expectedNewList := []*uint8{&v4, &v2, &v4}
	NewList, _ := TakeWhileUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint8{&v40}
	partialIsEvenDivisibleBy := func(num *uint8) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileUint8PtrErr(partialIsEvenDivisibleBy, []*uint8{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileUint8PtrErr failed.")
	}

	r, _ = TakeWhileUint8PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("TakeWhileUint8Ptr failed.")
	}
}

func TestTakeWhileStrPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string = "0"

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList, _ := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) (bool, error) {
		if *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = TakeWhileStrPtrErr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtrErr failed.")
	}

	r, _ = TakeWhileStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}
}

func TestTakeWhileBoolPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vt, &vt, &vf}
	NewList, _ := TakeWhileBoolPtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*bool{&vt, &vt, &vf, &vf, &vf})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhileBoolPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*bool{&vt}
	NewList, _ = TakeWhileBoolPtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*bool{&vt})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileBoolPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileBoolPtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid for this test")
		}
		return *v == true, nil
	}, []*bool{&vf})

	if err == nil {
		t.Errorf("TakeWhileBoolPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileBoolPtrErr failed.")
	}

	r, _ = TakeWhileBoolPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}
}

func TestTakeWhileFloat32PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v7 float32 = 7
	var v40 float32 = 40
	var v0 float32

	expectedNewList := []*float32{&v4, &v2, &v4}
	NewList, _ := TakeWhileFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*float32{&v40}
	partialIsEvenDivisibleBy := func(num *float32) (bool, error) { return int(*num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat32PtrErr(partialIsEvenDivisibleBy, []*float32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat32PtrErr failed.")
	}

	r, _ = TakeWhileFloat32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat32Ptr failed.")
	}
}

func TestTakeWhileFloat64PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v7 float64 = 7
	var v40 float64 = 40
	var v0 float64

	expectedNewList := []*float64{&v4, &v2, &v4}
	NewList, _ := TakeWhileFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*float64{&v40}
	partialIsEvenDivisibleBy := func(num *float64) (bool, error) { return int(*num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat64PtrErr(partialIsEvenDivisibleBy, []*float64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat64PtrErr failed.")
	}

	r, _ = TakeWhileFloat64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("TakeWhileFloat64Ptr failed.")
	}
}
