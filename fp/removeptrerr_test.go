package fp

import (
	"errors"
	"testing"
)

func TestRemoveIntPtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v20 int = 20
	var v40 int = 40

	expectedNewList := []*int{&v1, &v3}
	NewList, _ := RemoveIntPtrErr(isEvenIntPtrErr, []*int{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveIntPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int{&v1, &v3}
	partialIsEven := func(num *int) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveIntPtrErr(partialIsEven, []*int{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveIntPtr failed.")
	}

	_, err := RemoveIntPtrErr(func(v *int) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*int{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveIntPtrErr failed.")
	}
}

func TestRemoveInt64PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v20 int64 = 20
	var v40 int64 = 40

	expectedNewList := []*int64{&v1, &v3}
	NewList, _ := RemoveInt64PtrErr(isEvenInt64PtrErr, []*int64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int64{&v1, &v3}
	partialIsEven := func(num *int64) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveInt64PtrErr(partialIsEven, []*int64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt64Ptr failed.")
	}

	_, err := RemoveInt64PtrErr(func(v *int64) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*int64{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveInt64PtrErr failed.")
	}
}

func TestRemoveInt32PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v20 int32 = 20
	var v40 int32 = 40

	expectedNewList := []*int32{&v1, &v3}
	NewList, _ := RemoveInt32PtrErr(isEvenInt32PtrErr, []*int32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int32{&v1, &v3}
	partialIsEven := func(num *int32) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveInt32PtrErr(partialIsEven, []*int32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt32Ptr failed.")
	}

	_, err := RemoveInt32PtrErr(func(v *int32) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*int32{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveInt32PtrErr failed.")
	}
}

func TestRemoveInt16PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v20 int16 = 20
	var v40 int16 = 40

	expectedNewList := []*int16{&v1, &v3}
	NewList, _ := RemoveInt16PtrErr(isEvenInt16PtrErr, []*int16{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int16{&v1, &v3}
	partialIsEven := func(num *int16) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveInt16PtrErr(partialIsEven, []*int16{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt16Ptr failed.")
	}

	_, err := RemoveInt16PtrErr(func(v *int16) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*int16{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveInt16PtrErr failed.")
	}
}

func TestRemoveInt8PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v20 int8 = 20
	var v40 int8 = 40

	expectedNewList := []*int8{&v1, &v3}
	NewList, _ := RemoveInt8PtrErr(isEvenInt8PtrErr, []*int8{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*int8{&v1, &v3}
	partialIsEven := func(num *int8) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveInt8PtrErr(partialIsEven, []*int8{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt8Ptr failed.")
	}

	_, err := RemoveInt8PtrErr(func(v *int8) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*int8{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveInt8PtrErr failed.")
	}
}

func TestRemoveUintPtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v20 uint = 20
	var v40 uint = 40

	expectedNewList := []*uint{&v1, &v3}
	NewList, _ := RemoveUintPtrErr(isEvenUintPtrErr, []*uint{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUintPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint{&v1, &v3}
	partialIsEven := func(num *uint) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveUintPtrErr(partialIsEven, []*uint{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUintPtr failed.")
	}

	_, err := RemoveUintPtrErr(func(v *uint) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*uint{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveUintPtrErr failed.")
	}
}

func TestRemoveUint64PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v20 uint64 = 20
	var v40 uint64 = 40

	expectedNewList := []*uint64{&v1, &v3}
	NewList, _ := RemoveUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint64{&v1, &v3}
	partialIsEven := func(num *uint64) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveUint64PtrErr(partialIsEven, []*uint64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint64Ptr failed.")
	}

	_, err := RemoveUint64PtrErr(func(v *uint64) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*uint64{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveUint64PtrErr failed.")
	}
}

func TestRemoveUint32PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v20 uint32 = 20
	var v40 uint32 = 40

	expectedNewList := []*uint32{&v1, &v3}
	NewList, _ := RemoveUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint32{&v1, &v3}
	partialIsEven := func(num *uint32) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveUint32PtrErr(partialIsEven, []*uint32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint32Ptr failed.")
	}

	_, err := RemoveUint32PtrErr(func(v *uint32) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*uint32{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveUint32PtrErr failed.")
	}
}

func TestRemoveUint16PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v20 uint16 = 20
	var v40 uint16 = 40

	expectedNewList := []*uint16{&v1, &v3}
	NewList, _ := RemoveUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint16PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint16{&v1, &v3}
	partialIsEven := func(num *uint16) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveUint16PtrErr(partialIsEven, []*uint16{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint16Ptr failed.")
	}

	_, err := RemoveUint16PtrErr(func(v *uint16) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*uint16{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveUint16PtrErr failed.")
	}
}

func TestRemoveUint8PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v20 uint8 = 20
	var v40 uint8 = 40

	expectedNewList := []*uint8{&v1, &v3}
	NewList, _ := RemoveUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint8PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*uint8{&v1, &v3}
	partialIsEven := func(num *uint8) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveUint8PtrErr(partialIsEven, []*uint8{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint8Ptr failed.")
	}

	_, err := RemoveUint8PtrErr(func(v *uint8) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*uint8{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveUint8PtrErr failed.")
	}
}

func TestRemoveStrPtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v20 string = "20"
	var v40 string = "40"

	expectedNewList := []*string{&v1, &v3}
	NewList, _ := RemoveStrPtrErr(isEvenStrPtrErr, []*string{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*string{&v1, &v3}
	partialIsEven := func(num *string) (bool, error) {
		if *num == "0" {
			return false, errors.New("0 in invalid number for this test")
		}
		if *num == "20" || *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = RemoveStrPtrErr(partialIsEven, []*string{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveStrPtr failed.")
	}

	_, err := RemoveStrPtrErr(func(v *string) (bool, error) {
		if *v == "20" {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*string{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveStrPtrErr failed.")
	}
}

func TestRemoveBoolPtrErr(t *testing.T) {
	// Test : even number in the list
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vt}
	NewList, _ := RemoveBoolPtrErr(func(v *bool) (bool, error) { return *v == false, nil }, []*bool{&vt, &vf, &vf})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("RemoveBoolPtrErr failed. Expected New list=%v, actual list=%v", *expectedNewList[0], *NewList[0])
	}

	_, err := RemoveBoolPtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid in this test")
		}
		return true, nil
	}, []*bool{&vt, &vf, &vf})
	if err == nil {
		t.Errorf("RemoveBoolPtrErr failed.")
	}

	r, _ := RemoveBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveBoolPtrErr failed.")
	}
}

func TestRemoveFloat32PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v20 float32 = 20
	var v40 float32 = 40

	expectedNewList := []*float32{&v1, &v3}
	NewList, _ := RemoveFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float32{&v1, &v3}
	partialIsEven := func(num *float32) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(*num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat32PtrErr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveFloat32Ptr failed.")
	}

	_, err := RemoveFloat32PtrErr(func(v *float32) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*float32{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveFloat32PtrErr failed.")
	}
}

func TestRemoveFloat64PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v20 float64 = 20
	var v40 float64 = 40

	expectedNewList := []*float64{&v1, &v3}
	NewList, _ := RemoveFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float64{&v1, &v3}
	partialIsEven := func(num *float64) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(*num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat64PtrErr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveFloat64Ptr failed.")
	}

	_, err := RemoveFloat64PtrErr(func(v *float64) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*float64{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("RemoveFloat64PtrErr failed.")
	}
}
