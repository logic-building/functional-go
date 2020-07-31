package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestTakeWhileIntPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int = 2
	var v4 int = 4
	var v5 int = 5
	var v7 int = 7
	var v40 int = 40

	expectedNewList := []*int{&v4, &v2, &v4}
	NewList := TakeWhileIntPtr(isEvenIntPtr, []*int{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int{&v40}
	partialIsEvenDivisibleBy := func(num *int) bool { return *num%10 == 0 }
	NewList = TakeWhileIntPtr(partialIsEvenDivisibleBy, []*int{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileIntPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileIntPtr failed.")
	}

	if len(TakeWhileIntPtr(nil, []*int{})) > 0 {
		t.Errorf("TakeWhileIntPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileInt64Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int64 = 2
	var v4 int64 = 4
	var v5 int64 = 5
	var v7 int64 = 7
	var v40 int64 = 40

	expectedNewList := []*int64{&v4, &v2, &v4}
	NewList := TakeWhileInt64Ptr(isEvenInt64Ptr, []*int64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int64{&v40}
	partialIsEvenDivisibleBy := func(num *int64) bool { return *num%10 == 0 }
	NewList = TakeWhileInt64Ptr(partialIsEvenDivisibleBy, []*int64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt64Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt64Ptr failed.")
	}

	if len(TakeWhileInt64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("TakeWhileInt64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileInt32Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int32 = 2
	var v4 int32 = 4
	var v5 int32 = 5
	var v7 int32 = 7
	var v40 int32 = 40

	expectedNewList := []*int32{&v4, &v2, &v4}
	NewList := TakeWhileInt32Ptr(isEvenInt32Ptr, []*int32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int32{&v40}
	partialIsEvenDivisibleBy := func(num *int32) bool { return *num%10 == 0 }
	NewList = TakeWhileInt32Ptr(partialIsEvenDivisibleBy, []*int32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt32Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt32Ptr failed.")
	}

	if len(TakeWhileInt32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("TakeWhileInt32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileInt16Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int16 = 2
	var v4 int16 = 4
	var v5 int16 = 5
	var v7 int16 = 7
	var v40 int16 = 40

	expectedNewList := []*int16{&v4, &v2, &v4}
	NewList := TakeWhileInt16Ptr(isEvenInt16Ptr, []*int16{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int16{&v40}
	partialIsEvenDivisibleBy := func(num *int16) bool { return *num%10 == 0 }
	NewList = TakeWhileInt16Ptr(partialIsEvenDivisibleBy, []*int16{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt16Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt16Ptr failed.")
	}

	if len(TakeWhileInt16Ptr(nil, []*int16{})) > 0 {
		t.Errorf("TakeWhileInt16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileInt8Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 int8 = 2
	var v4 int8 = 4
	var v5 int8 = 5
	var v7 int8 = 7
	var v40 int8 = 40

	expectedNewList := []*int8{&v4, &v2, &v4}
	NewList := TakeWhileInt8Ptr(isEvenInt8Ptr, []*int8{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*int8{&v40}
	partialIsEvenDivisibleBy := func(num *int8) bool { return *num%10 == 0 }
	NewList = TakeWhileInt8Ptr(partialIsEvenDivisibleBy, []*int8{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt8Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt8Ptr failed.")
	}

	if len(TakeWhileInt8Ptr(nil, []*int8{})) > 0 {
		t.Errorf("TakeWhileInt8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileUintPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint = 2
	var v4 uint = 4
	var v5 uint = 5
	var v7 uint = 7
	var v40 uint = 40

	expectedNewList := []*uint{&v4, &v2, &v4}
	NewList := TakeWhileUintPtr(isEvenUintPtr, []*uint{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint{&v40}
	partialIsEvenDivisibleBy := func(num *uint) bool { return *num%10 == 0 }
	NewList = TakeWhileUintPtr(partialIsEvenDivisibleBy, []*uint{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUintPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileUintPtr failed.")
	}

	if len(TakeWhileUintPtr(nil, []*uint{})) > 0 {
		t.Errorf("TakeWhileUintPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileUint64Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint64 = 2
	var v4 uint64 = 4
	var v5 uint64 = 5
	var v7 uint64 = 7
	var v40 uint64 = 40

	expectedNewList := []*uint64{&v4, &v2, &v4}
	NewList := TakeWhileUint64Ptr(isEvenUint64Ptr, []*uint64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint64{&v40}
	partialIsEvenDivisibleBy := func(num *uint64) bool { return *num%10 == 0 }
	NewList = TakeWhileUint64Ptr(partialIsEvenDivisibleBy, []*uint64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint64Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint64Ptr failed.")
	}

	if len(TakeWhileUint64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("TakeWhileUint64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileUint32Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint32 = 2
	var v4 uint32 = 4
	var v5 uint32 = 5
	var v7 uint32 = 7
	var v40 uint32 = 40

	expectedNewList := []*uint32{&v4, &v2, &v4}
	NewList := TakeWhileUint32Ptr(isEvenUint32Ptr, []*uint32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint32{&v40}
	partialIsEvenDivisibleBy := func(num *uint32) bool { return *num%10 == 0 }
	NewList = TakeWhileUint32Ptr(partialIsEvenDivisibleBy, []*uint32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint32Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint32Ptr failed.")
	}

	if len(TakeWhileUint32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("TakeWhileUint32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileUint16Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint16 = 2
	var v4 uint16 = 4
	var v5 uint16 = 5
	var v7 uint16 = 7
	var v40 uint16 = 40

	expectedNewList := []*uint16{&v4, &v2, &v4}
	NewList := TakeWhileUint16Ptr(isEvenUint16Ptr, []*uint16{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint16{&v40}
	partialIsEvenDivisibleBy := func(num *uint16) bool { return *num%10 == 0 }
	NewList = TakeWhileUint16Ptr(partialIsEvenDivisibleBy, []*uint16{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint16Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint16Ptr failed.")
	}

	if len(TakeWhileUint16Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("TakeWhileUint16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileUint8Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 uint8 = 2
	var v4 uint8 = 4
	var v5 uint8 = 5
	var v7 uint8 = 7
	var v40 uint8 = 40

	expectedNewList := []*uint8{&v4, &v2, &v4}
	NewList := TakeWhileUint8Ptr(isEvenUint8Ptr, []*uint8{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*uint8{&v40}
	partialIsEvenDivisibleBy := func(num *uint8) bool { return *num%10 == 0 }
	NewList = TakeWhileUint8Ptr(partialIsEvenDivisibleBy, []*uint8{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint8Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint8Ptr failed.")
	}

	if len(TakeWhileUint8Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("TakeWhileUint8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileStrPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList := TakeWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) bool {
		if *num == "40" {
			return true
		}
		return false
	}
	NewList = TakeWhileStrPtr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileStrPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}

	if len(TakeWhileStrPtr(nil, []*string{})) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileBoolPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vt, &vt, &vf}
	NewList := TakeWhileBoolPtr(func(v *bool) bool { return *v == true }, []*bool{&vt, &vt, &vf, &vf, &vf})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*bool{&vt}
	NewList = TakeWhileBoolPtr(func(v *bool) bool { return *v == true }, []*bool{&vt})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileBoolPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}

	if len(TakeWhileBoolPtr(nil, []*bool{})) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileFloat32Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float32 = 2
	var v4 float32 = 4
	var v5 float32 = 5
	var v7 float32 = 7
	var v40 float32 = 40

	expectedNewList := []*float32{&v4, &v2, &v4}
	NewList := TakeWhileFloat32Ptr(isEvenFloat32Ptr, []*float32{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*float32{&v40}
	partialIsEvenDivisibleBy := func(num *float32) bool { return int(*num)%10 == 0 }
	NewList = TakeWhileFloat32Ptr(partialIsEvenDivisibleBy, []*float32{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileFloat32Ptr failed.")
	}

	if len(TakeWhileFloat32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("TakeWhileFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestTakeWhileFloat64Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 float64 = 2
	var v4 float64 = 4
	var v5 float64 = 5
	var v7 float64 = 7
	var v40 float64 = 40

	expectedNewList := []*float64{&v4, &v2, &v4}
	NewList := TakeWhileFloat64Ptr(isEvenFloat64Ptr, []*float64{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*float64{&v40}
	partialIsEvenDivisibleBy := func(num *float64) bool { return int(*num)%10 == 0 }
	NewList = TakeWhileFloat64Ptr(partialIsEvenDivisibleBy, []*float64{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhileFloat64Ptr failed.")
	}

	if len(TakeWhileFloat64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("TakeWhileFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
