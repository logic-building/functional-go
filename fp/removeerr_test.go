package fp

import (
	"errors"
	"testing"
)

func TestRemoveIntErr(t *testing.T) {
	// Test : even number in the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v20 int = 20
	var v40 int = 40

	expectedNewList := []int{v1, v3}
	NewList, _ := RemoveIntErr(isEvenIntErr, []int{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int{v1, v3}
	partialIsEven := func(num int) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveIntErr(partialIsEven, []int{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveIntErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveIntErr failed.")
	}

	_, err := RemoveIntErr(func(v int) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []int{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveIntErr failed.")
	}
}

func TestRemoveInt64Err(t *testing.T) {
	// Test : even number in the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v20 int64 = 20
	var v40 int64 = 40

	expectedNewList := []int64{v1, v3}
	NewList, _ := RemoveInt64Err(isEvenInt64Err, []int64{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int64{v1, v3}
	partialIsEven := func(num int64) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveInt64Err(partialIsEven, []int64{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt64Err failed.")
	}

	_, err := RemoveInt64Err(func(v int64) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []int64{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveInt64Err failed.")
	}
}

func TestRemoveInt32Err(t *testing.T) {
	// Test : even number in the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v20 int32 = 20
	var v40 int32 = 40

	expectedNewList := []int32{v1, v3}
	NewList, _ := RemoveInt32Err(isEvenInt32Err, []int32{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int32{v1, v3}
	partialIsEven := func(num int32) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveInt32Err(partialIsEven, []int32{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt32Err failed.")
	}

	_, err := RemoveInt32Err(func(v int32) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []int32{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveInt32Err failed.")
	}
}

func TestRemoveInt16Err(t *testing.T) {
	// Test : even number in the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v20 int16 = 20
	var v40 int16 = 40

	expectedNewList := []int16{v1, v3}
	NewList, _ := RemoveInt16Err(isEvenInt16Err, []int16{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int16{v1, v3}
	partialIsEven := func(num int16) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveInt16Err(partialIsEven, []int16{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt16Err failed.")
	}

	_, err := RemoveInt16Err(func(v int16) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []int16{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveInt16Err failed.")
	}
}

func TestRemoveInt8Err(t *testing.T) {
	// Test : even number in the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v20 int8 = 20
	var v40 int8 = 40

	expectedNewList := []int8{v1, v3}
	NewList, _ := RemoveInt8Err(isEvenInt8Err, []int8{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int8{v1, v3}
	partialIsEven := func(num int8) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveInt8Err(partialIsEven, []int8{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveInt8Err failed.")
	}

	_, err := RemoveInt8Err(func(v int8) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []int8{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveInt8Err failed.")
	}
}

func TestRemoveUintErr(t *testing.T) {
	// Test : even number in the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v20 uint = 20
	var v40 uint = 40

	expectedNewList := []uint{v1, v3}
	NewList, _ := RemoveUintErr(isEvenUintErr, []uint{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint{v1, v3}
	partialIsEven := func(num uint) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveUintErr(partialIsEven, []uint{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUintErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUintErr failed.")
	}

	_, err := RemoveUintErr(func(v uint) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []uint{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveUintErr failed.")
	}
}

func TestRemoveUint64Err(t *testing.T) {
	// Test : even number in the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v20 uint64 = 20
	var v40 uint64 = 40

	expectedNewList := []uint64{v1, v3}
	NewList, _ := RemoveUint64Err(isEvenUint64Err, []uint64{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint64{v1, v3}
	partialIsEven := func(num uint64) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveUint64Err(partialIsEven, []uint64{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint64Err failed.")
	}

	_, err := RemoveUint64Err(func(v uint64) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []uint64{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveUint64Err failed.")
	}
}

func TestRemoveUint32Err(t *testing.T) {
	// Test : even number in the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v20 uint32 = 20
	var v40 uint32 = 40

	expectedNewList := []uint32{v1, v3}
	NewList, _ := RemoveUint32Err(isEvenUint32Err, []uint32{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint32{v1, v3}
	partialIsEven := func(num uint32) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveUint32Err(partialIsEven, []uint32{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint32Err failed.")
	}

	_, err := RemoveUint32Err(func(v uint32) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []uint32{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveUint32Err failed.")
	}
}

func TestRemoveUint16Err(t *testing.T) {
	// Test : even number in the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v20 uint16 = 20
	var v40 uint16 = 40

	expectedNewList := []uint16{v1, v3}
	NewList, _ := RemoveUint16Err(isEvenUint16Err, []uint16{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint16{v1, v3}
	partialIsEven := func(num uint16) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveUint16Err(partialIsEven, []uint16{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint16Err failed.")
	}

	_, err := RemoveUint16Err(func(v uint16) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []uint16{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveUint16Err failed.")
	}
}

func TestRemoveUint8Err(t *testing.T) {
	// Test : even number in the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v20 uint8 = 20
	var v40 uint8 = 40

	expectedNewList := []uint8{v1, v3}
	NewList, _ := RemoveUint8Err(isEvenUint8Err, []uint8{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint8{v1, v3}
	partialIsEven := func(num uint8) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveUint8Err(partialIsEven, []uint8{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint8Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveUint8Err failed.")
	}

	_, err := RemoveUint8Err(func(v uint8) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []uint8{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveUint8Err failed.")
	}
}

func TestRemoveStrErr(t *testing.T) {
	// Test : even number in the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v20 string = "20"
	var v40 string = "40"

	expectedNewList := []string{v1, v3}
	NewList, _ := RemoveStrErr(isEvenStrErr, []string{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []string{v1, v3}
	partialIsEven := func(num string) (bool, error) {
		if num == "0" {
			return false, errors.New("0 in invalid number for this test")
		}
		if num == "20" || num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = RemoveStrErr(partialIsEven, []string{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveStrErr failed.")
	}

	_, err := RemoveStrErr(func(v string) (bool, error) {
		if v == "20" {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []string{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveStrErr failed.")
	}
}

func TestRemoveBoolErr(t *testing.T) {
	// Test : even number in the list
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vt}
	NewList, _ := RemoveBoolErr(func(v bool) (bool, error) { return v == false, nil }, []bool{vt, vf, vf})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("RemoveBoolErr failed. Expected New list=%v, actual list=%v", expectedNewList[0], NewList[0])
	}

	_, err := RemoveBoolErr(func(v bool) (bool, error) {
		if v == false {
			return false, errors.New("false is invalid in this test")
		}
		return true, nil
	}, []bool{vt, vf, vf})

	if err == nil {
		t.Errorf("RemoveBoolErr failed.")
	}

	r, _ := RemoveBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveBoolErr failed.")
	}
}

func TestRemoveFloat32Err(t *testing.T) {
	// Test : even number in the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v20 float32 = 20
	var v40 float32 = 40

	expectedNewList := []float32{v1, v3}
	NewList, _ := RemoveFloat32Err(isEvenFloat32Err, []float32{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []float32{v1, v3}
	partialIsEven := func(num float32) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat32Err(partialIsEven, []float32{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat32Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveFloat32Err failed.")
	}

	_, err := RemoveFloat32Err(func(v float32) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []float32{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveFloat32Err failed.")
	}
}

func TestRemoveFloat64Err(t *testing.T) {
	// Test : even number in the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v20 float64 = 20
	var v40 float64 = 40

	expectedNewList := []float64{v1, v3}
	NewList, _ := RemoveFloat64Err(isEvenFloat64Err, []float64{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []float64{v1, v3}
	partialIsEven := func(num float64) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat64Err(partialIsEven, []float64{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat64Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := RemoveFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("RemoveFloat64Err failed.")
	}

	_, err := RemoveFloat64Err(func(v float64) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []float64{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("RemoveFloat64Err failed.")
	}
}
