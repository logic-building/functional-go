package fp

import (
	"errors"
	"testing"
)

func TestFilterMapIntInt64Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapIntInt64Err(notOneIntInt64Err, plusOneIntInt64Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt64Err failed")
	}
	r, _ = FilterMapIntInt64Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt64PtrErr failed")
	}

	_, err := FilterMapIntInt64Err(notOneIntInt64Err, plusOneIntInt64Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntInt64Err failed")
	}
	_, err = FilterMapIntInt64Err(notOneIntInt64Err, plusOneIntInt64Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntInt64Err failed")
	}
}
func notOneIntInt64Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntInt64Err(num int) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapIntInt32Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapIntInt32Err(notOneIntInt32Err, plusOneIntInt32Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt32Err failed")
	}
	r, _ = FilterMapIntInt32Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt32PtrErr failed")
	}

	_, err := FilterMapIntInt32Err(notOneIntInt32Err, plusOneIntInt32Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntInt32Err failed")
	}
	_, err = FilterMapIntInt32Err(notOneIntInt32Err, plusOneIntInt32Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntInt32Err failed")
	}
}
func notOneIntInt32Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntInt32Err(num int) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapIntInt16Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapIntInt16Err(notOneIntInt16Err, plusOneIntInt16Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt16Err failed")
	}
	r, _ = FilterMapIntInt16Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt16PtrErr failed")
	}

	_, err := FilterMapIntInt16Err(notOneIntInt16Err, plusOneIntInt16Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntInt16Err failed")
	}
	_, err = FilterMapIntInt16Err(notOneIntInt16Err, plusOneIntInt16Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntInt16Err failed")
	}
}
func notOneIntInt16Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntInt16Err(num int) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapIntInt8Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapIntInt8Err(notOneIntInt8Err, plusOneIntInt8Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt8Err failed")
	}
	r, _ = FilterMapIntInt8Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt8PtrErr failed")
	}

	_, err := FilterMapIntInt8Err(notOneIntInt8Err, plusOneIntInt8Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntInt8Err failed")
	}
	_, err = FilterMapIntInt8Err(notOneIntInt8Err, plusOneIntInt8Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntInt8Err failed")
	}
}
func notOneIntInt8Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntInt8Err(num int) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapIntUintErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapIntUintErr(notOneIntUintErr, plusOneIntUintErr, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUintErr failed")
	}
	r, _ = FilterMapIntUintErr(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUintPtrErr failed")
	}

	_, err := FilterMapIntUintErr(notOneIntUintErr, plusOneIntUintErr, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntUintErr failed")
	}
	_, err = FilterMapIntUintErr(notOneIntUintErr, plusOneIntUintErr, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntUintErr failed")
	}
}
func notOneIntUintErr(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntUintErr(num int) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapIntUint64Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapIntUint64Err(notOneIntUint64Err, plusOneIntUint64Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint64Err failed")
	}
	r, _ = FilterMapIntUint64Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint64PtrErr failed")
	}

	_, err := FilterMapIntUint64Err(notOneIntUint64Err, plusOneIntUint64Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntUint64Err failed")
	}
	_, err = FilterMapIntUint64Err(notOneIntUint64Err, plusOneIntUint64Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntUint64Err failed")
	}
}
func notOneIntUint64Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntUint64Err(num int) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapIntUint32Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapIntUint32Err(notOneIntUint32Err, plusOneIntUint32Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint32Err failed")
	}
	r, _ = FilterMapIntUint32Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint32PtrErr failed")
	}

	_, err := FilterMapIntUint32Err(notOneIntUint32Err, plusOneIntUint32Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntUint32Err failed")
	}
	_, err = FilterMapIntUint32Err(notOneIntUint32Err, plusOneIntUint32Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntUint32Err failed")
	}
}
func notOneIntUint32Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntUint32Err(num int) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapIntUint16Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapIntUint16Err(notOneIntUint16Err, plusOneIntUint16Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint16Err failed")
	}
	r, _ = FilterMapIntUint16Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint16PtrErr failed")
	}

	_, err := FilterMapIntUint16Err(notOneIntUint16Err, plusOneIntUint16Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntUint16Err failed")
	}
	_, err = FilterMapIntUint16Err(notOneIntUint16Err, plusOneIntUint16Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntUint16Err failed")
	}
}
func notOneIntUint16Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntUint16Err(num int) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapIntUint8Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapIntUint8Err(notOneIntUint8Err, plusOneIntUint8Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint8Err failed")
	}
	r, _ = FilterMapIntUint8Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint8PtrErr failed")
	}

	_, err := FilterMapIntUint8Err(notOneIntUint8Err, plusOneIntUint8Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntUint8Err failed")
	}
	_, err = FilterMapIntUint8Err(notOneIntUint8Err, plusOneIntUint8Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntUint8Err failed")
	}
}
func notOneIntUint8Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntUint8Err(num int) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapIntStrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int = 1
	var iv2 int = 2
	var iv3 int = 3
	var iv10 int = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapIntStrErr(notOneIntStrNumErr, someLogicIntStrNumErr, []int{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapIntStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntStrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntStrErr failed")
	}

	r, _ = FilterMapIntStrErr(nil, nil, []int{})
	if len(r) > 0 {
		t.Errorf("FilterMapIntStrErr failed")
	}

	_, err := FilterMapIntStrErr(notOneIntStrNumErr, someLogicIntStrNumErr, []int{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapIntStrErr failed")
	}
	_, err = FilterMapIntStrErr(notOneIntStrNumErr, someLogicIntStrNumErr, []int{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapIntStrErr failed")
	}
}

func notOneIntStrNumErr(num int) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicIntStrNumErr(num int) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapIntBoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	var vi10 int = 10
	var vi0 int

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapIntBoolErr(notOneIntBoolErr, someLogicIntBoolErr, []int{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntBoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntBoolErr failed")
	}

	r, _ = FilterMapIntBoolErr(nil, nil, []int{})
	if len(r) > 0 {
		t.Errorf("FilterMapIntBoolPtrErr failed")
	}

	_, err := FilterMapIntBoolErr(notOneIntBoolErr, someLogicIntBoolErr, []int{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapIntBoolErr failed")
	}

	_, err = FilterMapIntBoolErr(notOneIntBoolErr, someLogicIntBoolErr, []int{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapIntBoolErr failed")
	}
}
func notOneIntBoolErr(num int) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicIntBoolErr(num int) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapIntFloat32Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapIntFloat32Err(notOneIntFloat32Err, plusOneIntFloat32Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntFloat32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat32Err failed")
	}
	r, _ = FilterMapIntFloat32Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat32PtrErr failed")
	}

	_, err := FilterMapIntFloat32Err(notOneIntFloat32Err, plusOneIntFloat32Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat32Err failed")
	}
	_, err = FilterMapIntFloat32Err(notOneIntFloat32Err, plusOneIntFloat32Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat32Err failed")
	}
}
func notOneIntFloat32Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntFloat32Err(num int) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapIntFloat64Err(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapIntFloat64Err(notOneIntFloat64Err, plusOneIntFloat64Err, []int{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapIntFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntFloat64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat64Err failed")
	}
	r, _ = FilterMapIntFloat64Err(nil, nil, []int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat64PtrErr failed")
	}

	_, err := FilterMapIntFloat64Err(notOneIntFloat64Err, plusOneIntFloat64Err, []int{v2, v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat64Err failed")
	}
	_, err = FilterMapIntFloat64Err(notOneIntFloat64Err, plusOneIntFloat64Err, []int{v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat64Err failed")
	}
}
func notOneIntFloat64Err(num int) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneIntFloat64Err(num int) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapInt64IntErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapInt64IntErr(notOneInt64IntErr, plusOneInt64IntErr, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64IntErr failed")
	}
	r, _ = FilterMapInt64IntErr(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64IntPtrErr failed")
	}

	_, err := FilterMapInt64IntErr(notOneInt64IntErr, plusOneInt64IntErr, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64IntErr failed")
	}
	_, err = FilterMapInt64IntErr(notOneInt64IntErr, plusOneInt64IntErr, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64IntErr failed")
	}
}
func notOneInt64IntErr(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64IntErr(num int64) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapInt64Int32Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapInt64Int32Err(notOneInt64Int32Err, plusOneInt64Int32Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int32Err failed")
	}
	r, _ = FilterMapInt64Int32Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int32PtrErr failed")
	}

	_, err := FilterMapInt64Int32Err(notOneInt64Int32Err, plusOneInt64Int32Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int32Err failed")
	}
	_, err = FilterMapInt64Int32Err(notOneInt64Int32Err, plusOneInt64Int32Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int32Err failed")
	}
}
func notOneInt64Int32Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Int32Err(num int64) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapInt64Int16Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapInt64Int16Err(notOneInt64Int16Err, plusOneInt64Int16Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int16Err failed")
	}
	r, _ = FilterMapInt64Int16Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int16PtrErr failed")
	}

	_, err := FilterMapInt64Int16Err(notOneInt64Int16Err, plusOneInt64Int16Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int16Err failed")
	}
	_, err = FilterMapInt64Int16Err(notOneInt64Int16Err, plusOneInt64Int16Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int16Err failed")
	}
}
func notOneInt64Int16Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Int16Err(num int64) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapInt64Int8Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapInt64Int8Err(notOneInt64Int8Err, plusOneInt64Int8Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int8Err failed")
	}
	r, _ = FilterMapInt64Int8Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int8PtrErr failed")
	}

	_, err := FilterMapInt64Int8Err(notOneInt64Int8Err, plusOneInt64Int8Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int8Err failed")
	}
	_, err = FilterMapInt64Int8Err(notOneInt64Int8Err, plusOneInt64Int8Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int8Err failed")
	}
}
func notOneInt64Int8Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Int8Err(num int64) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapInt64UintErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapInt64UintErr(notOneInt64UintErr, plusOneInt64UintErr, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64UintErr failed")
	}
	r, _ = FilterMapInt64UintErr(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64UintPtrErr failed")
	}

	_, err := FilterMapInt64UintErr(notOneInt64UintErr, plusOneInt64UintErr, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64UintErr failed")
	}
	_, err = FilterMapInt64UintErr(notOneInt64UintErr, plusOneInt64UintErr, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64UintErr failed")
	}
}
func notOneInt64UintErr(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64UintErr(num int64) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapInt64Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapInt64Uint64Err(notOneInt64Uint64Err, plusOneInt64Uint64Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint64Err failed")
	}
	r, _ = FilterMapInt64Uint64Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint64PtrErr failed")
	}

	_, err := FilterMapInt64Uint64Err(notOneInt64Uint64Err, plusOneInt64Uint64Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint64Err failed")
	}
	_, err = FilterMapInt64Uint64Err(notOneInt64Uint64Err, plusOneInt64Uint64Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint64Err failed")
	}
}
func notOneInt64Uint64Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Uint64Err(num int64) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapInt64Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapInt64Uint32Err(notOneInt64Uint32Err, plusOneInt64Uint32Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint32Err failed")
	}
	r, _ = FilterMapInt64Uint32Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint32PtrErr failed")
	}

	_, err := FilterMapInt64Uint32Err(notOneInt64Uint32Err, plusOneInt64Uint32Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint32Err failed")
	}
	_, err = FilterMapInt64Uint32Err(notOneInt64Uint32Err, plusOneInt64Uint32Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint32Err failed")
	}
}
func notOneInt64Uint32Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Uint32Err(num int64) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapInt64Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapInt64Uint16Err(notOneInt64Uint16Err, plusOneInt64Uint16Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint16Err failed")
	}
	r, _ = FilterMapInt64Uint16Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint16PtrErr failed")
	}

	_, err := FilterMapInt64Uint16Err(notOneInt64Uint16Err, plusOneInt64Uint16Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint16Err failed")
	}
	_, err = FilterMapInt64Uint16Err(notOneInt64Uint16Err, plusOneInt64Uint16Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint16Err failed")
	}
}
func notOneInt64Uint16Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Uint16Err(num int64) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapInt64Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapInt64Uint8Err(notOneInt64Uint8Err, plusOneInt64Uint8Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint8Err failed")
	}
	r, _ = FilterMapInt64Uint8Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint8PtrErr failed")
	}

	_, err := FilterMapInt64Uint8Err(notOneInt64Uint8Err, plusOneInt64Uint8Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint8Err failed")
	}
	_, err = FilterMapInt64Uint8Err(notOneInt64Uint8Err, plusOneInt64Uint8Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint8Err failed")
	}
}
func notOneInt64Uint8Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Uint8Err(num int64) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapInt64StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int64 = 1
	var iv2 int64 = 2
	var iv3 int64 = 3
	var iv10 int64 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapInt64StrErr(notOneInt64StrNumErr, someLogicInt64StrNumErr, []int64{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64StrErr failed")
	}

	r, _ = FilterMapInt64StrErr(nil, nil, []int64{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt64StrErr failed")
	}

	_, err := FilterMapInt64StrErr(notOneInt64StrNumErr, someLogicInt64StrNumErr, []int64{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapInt64StrErr failed")
	}
	_, err = FilterMapInt64StrErr(notOneInt64StrNumErr, someLogicInt64StrNumErr, []int64{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapInt64StrErr failed")
	}
}

func notOneInt64StrNumErr(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt64StrNumErr(num int64) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapInt64BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	var vi10 int64 = 10
	var vi0 int64

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapInt64BoolErr(notOneInt64BoolErr, someLogicInt64BoolErr, []int64{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64BoolErr failed")
	}

	r, _ = FilterMapInt64BoolErr(nil, nil, []int64{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt64BoolPtrErr failed")
	}

	_, err := FilterMapInt64BoolErr(notOneInt64BoolErr, someLogicInt64BoolErr, []int64{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt64BoolErr failed")
	}

	_, err = FilterMapInt64BoolErr(notOneInt64BoolErr, someLogicInt64BoolErr, []int64{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt64BoolErr failed")
	}
}
func notOneInt64BoolErr(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt64BoolErr(num int64) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapInt64Float32Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapInt64Float32Err(notOneInt64Float32Err, plusOneInt64Float32Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float32Err failed")
	}
	r, _ = FilterMapInt64Float32Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float32PtrErr failed")
	}

	_, err := FilterMapInt64Float32Err(notOneInt64Float32Err, plusOneInt64Float32Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float32Err failed")
	}
	_, err = FilterMapInt64Float32Err(notOneInt64Float32Err, plusOneInt64Float32Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float32Err failed")
	}
}
func notOneInt64Float32Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Float32Err(num int64) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapInt64Float64Err(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapInt64Float64Err(notOneInt64Float64Err, plusOneInt64Float64Err, []int64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt64Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float64Err failed")
	}
	r, _ = FilterMapInt64Float64Err(nil, nil, []int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float64PtrErr failed")
	}

	_, err := FilterMapInt64Float64Err(notOneInt64Float64Err, plusOneInt64Float64Err, []int64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float64Err failed")
	}
	_, err = FilterMapInt64Float64Err(notOneInt64Float64Err, plusOneInt64Float64Err, []int64{v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float64Err failed")
	}
}
func notOneInt64Float64Err(num int64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt64Float64Err(num int64) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapInt32IntErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapInt32IntErr(notOneInt32IntErr, plusOneInt32IntErr, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32IntErr failed")
	}
	r, _ = FilterMapInt32IntErr(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32IntPtrErr failed")
	}

	_, err := FilterMapInt32IntErr(notOneInt32IntErr, plusOneInt32IntErr, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32IntErr failed")
	}
	_, err = FilterMapInt32IntErr(notOneInt32IntErr, plusOneInt32IntErr, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32IntErr failed")
	}
}
func notOneInt32IntErr(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32IntErr(num int32) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapInt32Int64Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapInt32Int64Err(notOneInt32Int64Err, plusOneInt32Int64Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int64Err failed")
	}
	r, _ = FilterMapInt32Int64Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int64PtrErr failed")
	}

	_, err := FilterMapInt32Int64Err(notOneInt32Int64Err, plusOneInt32Int64Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int64Err failed")
	}
	_, err = FilterMapInt32Int64Err(notOneInt32Int64Err, plusOneInt32Int64Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int64Err failed")
	}
}
func notOneInt32Int64Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Int64Err(num int32) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapInt32Int16Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapInt32Int16Err(notOneInt32Int16Err, plusOneInt32Int16Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int16Err failed")
	}
	r, _ = FilterMapInt32Int16Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int16PtrErr failed")
	}

	_, err := FilterMapInt32Int16Err(notOneInt32Int16Err, plusOneInt32Int16Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int16Err failed")
	}
	_, err = FilterMapInt32Int16Err(notOneInt32Int16Err, plusOneInt32Int16Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int16Err failed")
	}
}
func notOneInt32Int16Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Int16Err(num int32) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapInt32Int8Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapInt32Int8Err(notOneInt32Int8Err, plusOneInt32Int8Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int8Err failed")
	}
	r, _ = FilterMapInt32Int8Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int8PtrErr failed")
	}

	_, err := FilterMapInt32Int8Err(notOneInt32Int8Err, plusOneInt32Int8Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int8Err failed")
	}
	_, err = FilterMapInt32Int8Err(notOneInt32Int8Err, plusOneInt32Int8Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int8Err failed")
	}
}
func notOneInt32Int8Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Int8Err(num int32) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapInt32UintErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapInt32UintErr(notOneInt32UintErr, plusOneInt32UintErr, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32UintErr failed")
	}
	r, _ = FilterMapInt32UintErr(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32UintPtrErr failed")
	}

	_, err := FilterMapInt32UintErr(notOneInt32UintErr, plusOneInt32UintErr, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32UintErr failed")
	}
	_, err = FilterMapInt32UintErr(notOneInt32UintErr, plusOneInt32UintErr, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32UintErr failed")
	}
}
func notOneInt32UintErr(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32UintErr(num int32) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapInt32Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapInt32Uint64Err(notOneInt32Uint64Err, plusOneInt32Uint64Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint64Err failed")
	}
	r, _ = FilterMapInt32Uint64Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint64PtrErr failed")
	}

	_, err := FilterMapInt32Uint64Err(notOneInt32Uint64Err, plusOneInt32Uint64Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint64Err failed")
	}
	_, err = FilterMapInt32Uint64Err(notOneInt32Uint64Err, plusOneInt32Uint64Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint64Err failed")
	}
}
func notOneInt32Uint64Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Uint64Err(num int32) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapInt32Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapInt32Uint32Err(notOneInt32Uint32Err, plusOneInt32Uint32Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint32Err failed")
	}
	r, _ = FilterMapInt32Uint32Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint32PtrErr failed")
	}

	_, err := FilterMapInt32Uint32Err(notOneInt32Uint32Err, plusOneInt32Uint32Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint32Err failed")
	}
	_, err = FilterMapInt32Uint32Err(notOneInt32Uint32Err, plusOneInt32Uint32Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint32Err failed")
	}
}
func notOneInt32Uint32Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Uint32Err(num int32) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapInt32Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapInt32Uint16Err(notOneInt32Uint16Err, plusOneInt32Uint16Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint16Err failed")
	}
	r, _ = FilterMapInt32Uint16Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint16PtrErr failed")
	}

	_, err := FilterMapInt32Uint16Err(notOneInt32Uint16Err, plusOneInt32Uint16Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint16Err failed")
	}
	_, err = FilterMapInt32Uint16Err(notOneInt32Uint16Err, plusOneInt32Uint16Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint16Err failed")
	}
}
func notOneInt32Uint16Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Uint16Err(num int32) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapInt32Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapInt32Uint8Err(notOneInt32Uint8Err, plusOneInt32Uint8Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint8Err failed")
	}
	r, _ = FilterMapInt32Uint8Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint8PtrErr failed")
	}

	_, err := FilterMapInt32Uint8Err(notOneInt32Uint8Err, plusOneInt32Uint8Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint8Err failed")
	}
	_, err = FilterMapInt32Uint8Err(notOneInt32Uint8Err, plusOneInt32Uint8Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint8Err failed")
	}
}
func notOneInt32Uint8Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Uint8Err(num int32) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapInt32StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int32 = 1
	var iv2 int32 = 2
	var iv3 int32 = 3
	var iv10 int32 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapInt32StrErr(notOneInt32StrNumErr, someLogicInt32StrNumErr, []int32{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32StrErr failed")
	}

	r, _ = FilterMapInt32StrErr(nil, nil, []int32{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt32StrErr failed")
	}

	_, err := FilterMapInt32StrErr(notOneInt32StrNumErr, someLogicInt32StrNumErr, []int32{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapInt32StrErr failed")
	}
	_, err = FilterMapInt32StrErr(notOneInt32StrNumErr, someLogicInt32StrNumErr, []int32{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapInt32StrErr failed")
	}
}

func notOneInt32StrNumErr(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt32StrNumErr(num int32) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapInt32BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	var vi10 int32 = 10
	var vi0 int32

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapInt32BoolErr(notOneInt32BoolErr, someLogicInt32BoolErr, []int32{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32BoolErr failed")
	}

	r, _ = FilterMapInt32BoolErr(nil, nil, []int32{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt32BoolPtrErr failed")
	}

	_, err := FilterMapInt32BoolErr(notOneInt32BoolErr, someLogicInt32BoolErr, []int32{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt32BoolErr failed")
	}

	_, err = FilterMapInt32BoolErr(notOneInt32BoolErr, someLogicInt32BoolErr, []int32{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt32BoolErr failed")
	}
}
func notOneInt32BoolErr(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt32BoolErr(num int32) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapInt32Float32Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapInt32Float32Err(notOneInt32Float32Err, plusOneInt32Float32Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float32Err failed")
	}
	r, _ = FilterMapInt32Float32Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float32PtrErr failed")
	}

	_, err := FilterMapInt32Float32Err(notOneInt32Float32Err, plusOneInt32Float32Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float32Err failed")
	}
	_, err = FilterMapInt32Float32Err(notOneInt32Float32Err, plusOneInt32Float32Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float32Err failed")
	}
}
func notOneInt32Float32Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Float32Err(num int32) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapInt32Float64Err(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapInt32Float64Err(notOneInt32Float64Err, plusOneInt32Float64Err, []int32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float64Err failed")
	}
	r, _ = FilterMapInt32Float64Err(nil, nil, []int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float64PtrErr failed")
	}

	_, err := FilterMapInt32Float64Err(notOneInt32Float64Err, plusOneInt32Float64Err, []int32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float64Err failed")
	}
	_, err = FilterMapInt32Float64Err(notOneInt32Float64Err, plusOneInt32Float64Err, []int32{v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float64Err failed")
	}
}
func notOneInt32Float64Err(num int32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt32Float64Err(num int32) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapInt16IntErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapInt16IntErr(notOneInt16IntErr, plusOneInt16IntErr, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16IntErr failed")
	}
	r, _ = FilterMapInt16IntErr(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16IntPtrErr failed")
	}

	_, err := FilterMapInt16IntErr(notOneInt16IntErr, plusOneInt16IntErr, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16IntErr failed")
	}
	_, err = FilterMapInt16IntErr(notOneInt16IntErr, plusOneInt16IntErr, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16IntErr failed")
	}
}
func notOneInt16IntErr(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16IntErr(num int16) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapInt16Int64Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapInt16Int64Err(notOneInt16Int64Err, plusOneInt16Int64Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int64Err failed")
	}
	r, _ = FilterMapInt16Int64Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int64PtrErr failed")
	}

	_, err := FilterMapInt16Int64Err(notOneInt16Int64Err, plusOneInt16Int64Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int64Err failed")
	}
	_, err = FilterMapInt16Int64Err(notOneInt16Int64Err, plusOneInt16Int64Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int64Err failed")
	}
}
func notOneInt16Int64Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Int64Err(num int16) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapInt16Int32Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapInt16Int32Err(notOneInt16Int32Err, plusOneInt16Int32Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int32Err failed")
	}
	r, _ = FilterMapInt16Int32Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int32PtrErr failed")
	}

	_, err := FilterMapInt16Int32Err(notOneInt16Int32Err, plusOneInt16Int32Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int32Err failed")
	}
	_, err = FilterMapInt16Int32Err(notOneInt16Int32Err, plusOneInt16Int32Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int32Err failed")
	}
}
func notOneInt16Int32Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Int32Err(num int16) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapInt16Int8Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapInt16Int8Err(notOneInt16Int8Err, plusOneInt16Int8Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int8Err failed")
	}
	r, _ = FilterMapInt16Int8Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int8PtrErr failed")
	}

	_, err := FilterMapInt16Int8Err(notOneInt16Int8Err, plusOneInt16Int8Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int8Err failed")
	}
	_, err = FilterMapInt16Int8Err(notOneInt16Int8Err, plusOneInt16Int8Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int8Err failed")
	}
}
func notOneInt16Int8Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Int8Err(num int16) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapInt16UintErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapInt16UintErr(notOneInt16UintErr, plusOneInt16UintErr, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16UintErr failed")
	}
	r, _ = FilterMapInt16UintErr(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16UintPtrErr failed")
	}

	_, err := FilterMapInt16UintErr(notOneInt16UintErr, plusOneInt16UintErr, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16UintErr failed")
	}
	_, err = FilterMapInt16UintErr(notOneInt16UintErr, plusOneInt16UintErr, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16UintErr failed")
	}
}
func notOneInt16UintErr(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16UintErr(num int16) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapInt16Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapInt16Uint64Err(notOneInt16Uint64Err, plusOneInt16Uint64Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint64Err failed")
	}
	r, _ = FilterMapInt16Uint64Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint64PtrErr failed")
	}

	_, err := FilterMapInt16Uint64Err(notOneInt16Uint64Err, plusOneInt16Uint64Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint64Err failed")
	}
	_, err = FilterMapInt16Uint64Err(notOneInt16Uint64Err, plusOneInt16Uint64Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint64Err failed")
	}
}
func notOneInt16Uint64Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Uint64Err(num int16) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapInt16Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapInt16Uint32Err(notOneInt16Uint32Err, plusOneInt16Uint32Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint32Err failed")
	}
	r, _ = FilterMapInt16Uint32Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint32PtrErr failed")
	}

	_, err := FilterMapInt16Uint32Err(notOneInt16Uint32Err, plusOneInt16Uint32Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint32Err failed")
	}
	_, err = FilterMapInt16Uint32Err(notOneInt16Uint32Err, plusOneInt16Uint32Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint32Err failed")
	}
}
func notOneInt16Uint32Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Uint32Err(num int16) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapInt16Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapInt16Uint16Err(notOneInt16Uint16Err, plusOneInt16Uint16Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint16Err failed")
	}
	r, _ = FilterMapInt16Uint16Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint16PtrErr failed")
	}

	_, err := FilterMapInt16Uint16Err(notOneInt16Uint16Err, plusOneInt16Uint16Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint16Err failed")
	}
	_, err = FilterMapInt16Uint16Err(notOneInt16Uint16Err, plusOneInt16Uint16Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint16Err failed")
	}
}
func notOneInt16Uint16Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Uint16Err(num int16) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapInt16Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapInt16Uint8Err(notOneInt16Uint8Err, plusOneInt16Uint8Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint8Err failed")
	}
	r, _ = FilterMapInt16Uint8Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint8PtrErr failed")
	}

	_, err := FilterMapInt16Uint8Err(notOneInt16Uint8Err, plusOneInt16Uint8Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint8Err failed")
	}
	_, err = FilterMapInt16Uint8Err(notOneInt16Uint8Err, plusOneInt16Uint8Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint8Err failed")
	}
}
func notOneInt16Uint8Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Uint8Err(num int16) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapInt16StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int16 = 1
	var iv2 int16 = 2
	var iv3 int16 = 3
	var iv10 int16 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapInt16StrErr(notOneInt16StrNumErr, someLogicInt16StrNumErr, []int16{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16StrErr failed")
	}

	r, _ = FilterMapInt16StrErr(nil, nil, []int16{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt16StrErr failed")
	}

	_, err := FilterMapInt16StrErr(notOneInt16StrNumErr, someLogicInt16StrNumErr, []int16{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapInt16StrErr failed")
	}
	_, err = FilterMapInt16StrErr(notOneInt16StrNumErr, someLogicInt16StrNumErr, []int16{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapInt16StrErr failed")
	}
}

func notOneInt16StrNumErr(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt16StrNumErr(num int16) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapInt16BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	var vi10 int16 = 10
	var vi0 int16

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapInt16BoolErr(notOneInt16BoolErr, someLogicInt16BoolErr, []int16{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16BoolErr failed")
	}

	r, _ = FilterMapInt16BoolErr(nil, nil, []int16{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt16BoolPtrErr failed")
	}

	_, err := FilterMapInt16BoolErr(notOneInt16BoolErr, someLogicInt16BoolErr, []int16{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt16BoolErr failed")
	}

	_, err = FilterMapInt16BoolErr(notOneInt16BoolErr, someLogicInt16BoolErr, []int16{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt16BoolErr failed")
	}
}
func notOneInt16BoolErr(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt16BoolErr(num int16) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapInt16Float32Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapInt16Float32Err(notOneInt16Float32Err, plusOneInt16Float32Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float32Err failed")
	}
	r, _ = FilterMapInt16Float32Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float32PtrErr failed")
	}

	_, err := FilterMapInt16Float32Err(notOneInt16Float32Err, plusOneInt16Float32Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float32Err failed")
	}
	_, err = FilterMapInt16Float32Err(notOneInt16Float32Err, plusOneInt16Float32Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float32Err failed")
	}
}
func notOneInt16Float32Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Float32Err(num int16) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapInt16Float64Err(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapInt16Float64Err(notOneInt16Float64Err, plusOneInt16Float64Err, []int16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt16Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float64Err failed")
	}
	r, _ = FilterMapInt16Float64Err(nil, nil, []int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float64PtrErr failed")
	}

	_, err := FilterMapInt16Float64Err(notOneInt16Float64Err, plusOneInt16Float64Err, []int16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float64Err failed")
	}
	_, err = FilterMapInt16Float64Err(notOneInt16Float64Err, plusOneInt16Float64Err, []int16{v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float64Err failed")
	}
}
func notOneInt16Float64Err(num int16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt16Float64Err(num int16) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapInt8IntErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapInt8IntErr(notOneInt8IntErr, plusOneInt8IntErr, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8IntErr failed")
	}
	r, _ = FilterMapInt8IntErr(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8IntPtrErr failed")
	}

	_, err := FilterMapInt8IntErr(notOneInt8IntErr, plusOneInt8IntErr, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8IntErr failed")
	}
	_, err = FilterMapInt8IntErr(notOneInt8IntErr, plusOneInt8IntErr, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8IntErr failed")
	}
}
func notOneInt8IntErr(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8IntErr(num int8) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapInt8Int64Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapInt8Int64Err(notOneInt8Int64Err, plusOneInt8Int64Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int64Err failed")
	}
	r, _ = FilterMapInt8Int64Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int64PtrErr failed")
	}

	_, err := FilterMapInt8Int64Err(notOneInt8Int64Err, plusOneInt8Int64Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int64Err failed")
	}
	_, err = FilterMapInt8Int64Err(notOneInt8Int64Err, plusOneInt8Int64Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int64Err failed")
	}
}
func notOneInt8Int64Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Int64Err(num int8) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapInt8Int32Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapInt8Int32Err(notOneInt8Int32Err, plusOneInt8Int32Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int32Err failed")
	}
	r, _ = FilterMapInt8Int32Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int32PtrErr failed")
	}

	_, err := FilterMapInt8Int32Err(notOneInt8Int32Err, plusOneInt8Int32Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int32Err failed")
	}
	_, err = FilterMapInt8Int32Err(notOneInt8Int32Err, plusOneInt8Int32Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int32Err failed")
	}
}
func notOneInt8Int32Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Int32Err(num int8) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapInt8Int16Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapInt8Int16Err(notOneInt8Int16Err, plusOneInt8Int16Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int16Err failed")
	}
	r, _ = FilterMapInt8Int16Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int16PtrErr failed")
	}

	_, err := FilterMapInt8Int16Err(notOneInt8Int16Err, plusOneInt8Int16Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int16Err failed")
	}
	_, err = FilterMapInt8Int16Err(notOneInt8Int16Err, plusOneInt8Int16Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int16Err failed")
	}
}
func notOneInt8Int16Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Int16Err(num int8) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapInt8UintErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapInt8UintErr(notOneInt8UintErr, plusOneInt8UintErr, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8UintErr failed")
	}
	r, _ = FilterMapInt8UintErr(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8UintPtrErr failed")
	}

	_, err := FilterMapInt8UintErr(notOneInt8UintErr, plusOneInt8UintErr, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8UintErr failed")
	}
	_, err = FilterMapInt8UintErr(notOneInt8UintErr, plusOneInt8UintErr, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8UintErr failed")
	}
}
func notOneInt8UintErr(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8UintErr(num int8) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapInt8Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapInt8Uint64Err(notOneInt8Uint64Err, plusOneInt8Uint64Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint64Err failed")
	}
	r, _ = FilterMapInt8Uint64Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint64PtrErr failed")
	}

	_, err := FilterMapInt8Uint64Err(notOneInt8Uint64Err, plusOneInt8Uint64Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint64Err failed")
	}
	_, err = FilterMapInt8Uint64Err(notOneInt8Uint64Err, plusOneInt8Uint64Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint64Err failed")
	}
}
func notOneInt8Uint64Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Uint64Err(num int8) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapInt8Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapInt8Uint32Err(notOneInt8Uint32Err, plusOneInt8Uint32Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint32Err failed")
	}
	r, _ = FilterMapInt8Uint32Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint32PtrErr failed")
	}

	_, err := FilterMapInt8Uint32Err(notOneInt8Uint32Err, plusOneInt8Uint32Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint32Err failed")
	}
	_, err = FilterMapInt8Uint32Err(notOneInt8Uint32Err, plusOneInt8Uint32Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint32Err failed")
	}
}
func notOneInt8Uint32Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Uint32Err(num int8) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapInt8Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapInt8Uint16Err(notOneInt8Uint16Err, plusOneInt8Uint16Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint16Err failed")
	}
	r, _ = FilterMapInt8Uint16Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint16PtrErr failed")
	}

	_, err := FilterMapInt8Uint16Err(notOneInt8Uint16Err, plusOneInt8Uint16Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint16Err failed")
	}
	_, err = FilterMapInt8Uint16Err(notOneInt8Uint16Err, plusOneInt8Uint16Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint16Err failed")
	}
}
func notOneInt8Uint16Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Uint16Err(num int8) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapInt8Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapInt8Uint8Err(notOneInt8Uint8Err, plusOneInt8Uint8Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint8Err failed")
	}
	r, _ = FilterMapInt8Uint8Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint8PtrErr failed")
	}

	_, err := FilterMapInt8Uint8Err(notOneInt8Uint8Err, plusOneInt8Uint8Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint8Err failed")
	}
	_, err = FilterMapInt8Uint8Err(notOneInt8Uint8Err, plusOneInt8Uint8Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint8Err failed")
	}
}
func notOneInt8Uint8Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Uint8Err(num int8) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapInt8StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int8 = 1
	var iv2 int8 = 2
	var iv3 int8 = 3
	var iv10 int8 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapInt8StrErr(notOneInt8StrNumErr, someLogicInt8StrNumErr, []int8{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapInt8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8StrErr failed")
	}

	r, _ = FilterMapInt8StrErr(nil, nil, []int8{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt8StrErr failed")
	}

	_, err := FilterMapInt8StrErr(notOneInt8StrNumErr, someLogicInt8StrNumErr, []int8{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapInt8StrErr failed")
	}
	_, err = FilterMapInt8StrErr(notOneInt8StrNumErr, someLogicInt8StrNumErr, []int8{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapInt8StrErr failed")
	}
}

func notOneInt8StrNumErr(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt8StrNumErr(num int8) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapInt8BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	var vi10 int8 = 10
	var vi0 int8

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapInt8BoolErr(notOneInt8BoolErr, someLogicInt8BoolErr, []int8{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8BoolErr failed")
	}

	r, _ = FilterMapInt8BoolErr(nil, nil, []int8{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt8BoolPtrErr failed")
	}

	_, err := FilterMapInt8BoolErr(notOneInt8BoolErr, someLogicInt8BoolErr, []int8{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt8BoolErr failed")
	}

	_, err = FilterMapInt8BoolErr(notOneInt8BoolErr, someLogicInt8BoolErr, []int8{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapInt8BoolErr failed")
	}
}
func notOneInt8BoolErr(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicInt8BoolErr(num int8) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapInt8Float32Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapInt8Float32Err(notOneInt8Float32Err, plusOneInt8Float32Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float32Err failed")
	}
	r, _ = FilterMapInt8Float32Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float32PtrErr failed")
	}

	_, err := FilterMapInt8Float32Err(notOneInt8Float32Err, plusOneInt8Float32Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float32Err failed")
	}
	_, err = FilterMapInt8Float32Err(notOneInt8Float32Err, plusOneInt8Float32Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float32Err failed")
	}
}
func notOneInt8Float32Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Float32Err(num int8) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapInt8Float64Err(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapInt8Float64Err(notOneInt8Float64Err, plusOneInt8Float64Err, []int8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapInt8Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float64Err failed")
	}
	r, _ = FilterMapInt8Float64Err(nil, nil, []int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float64PtrErr failed")
	}

	_, err := FilterMapInt8Float64Err(notOneInt8Float64Err, plusOneInt8Float64Err, []int8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float64Err failed")
	}
	_, err = FilterMapInt8Float64Err(notOneInt8Float64Err, plusOneInt8Float64Err, []int8{v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float64Err failed")
	}
}
func notOneInt8Float64Err(num int8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneInt8Float64Err(num int8) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapUintIntErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapUintIntErr(notOneUintIntErr, plusOneUintIntErr, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintIntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintIntErr failed")
	}
	r, _ = FilterMapUintIntErr(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintIntPtrErr failed")
	}

	_, err := FilterMapUintIntErr(notOneUintIntErr, plusOneUintIntErr, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintIntErr failed")
	}
	_, err = FilterMapUintIntErr(notOneUintIntErr, plusOneUintIntErr, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintIntErr failed")
	}
}
func notOneUintIntErr(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintIntErr(num uint) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapUintInt64Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapUintInt64Err(notOneUintInt64Err, plusOneUintInt64Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt64Err failed")
	}
	r, _ = FilterMapUintInt64Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt64PtrErr failed")
	}

	_, err := FilterMapUintInt64Err(notOneUintInt64Err, plusOneUintInt64Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintInt64Err failed")
	}
	_, err = FilterMapUintInt64Err(notOneUintInt64Err, plusOneUintInt64Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintInt64Err failed")
	}
}
func notOneUintInt64Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintInt64Err(num uint) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapUintInt32Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapUintInt32Err(notOneUintInt32Err, plusOneUintInt32Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt32Err failed")
	}
	r, _ = FilterMapUintInt32Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt32PtrErr failed")
	}

	_, err := FilterMapUintInt32Err(notOneUintInt32Err, plusOneUintInt32Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintInt32Err failed")
	}
	_, err = FilterMapUintInt32Err(notOneUintInt32Err, plusOneUintInt32Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintInt32Err failed")
	}
}
func notOneUintInt32Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintInt32Err(num uint) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapUintInt16Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapUintInt16Err(notOneUintInt16Err, plusOneUintInt16Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt16Err failed")
	}
	r, _ = FilterMapUintInt16Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt16PtrErr failed")
	}

	_, err := FilterMapUintInt16Err(notOneUintInt16Err, plusOneUintInt16Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintInt16Err failed")
	}
	_, err = FilterMapUintInt16Err(notOneUintInt16Err, plusOneUintInt16Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintInt16Err failed")
	}
}
func notOneUintInt16Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintInt16Err(num uint) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapUintInt8Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapUintInt8Err(notOneUintInt8Err, plusOneUintInt8Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt8Err failed")
	}
	r, _ = FilterMapUintInt8Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt8PtrErr failed")
	}

	_, err := FilterMapUintInt8Err(notOneUintInt8Err, plusOneUintInt8Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintInt8Err failed")
	}
	_, err = FilterMapUintInt8Err(notOneUintInt8Err, plusOneUintInt8Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintInt8Err failed")
	}
}
func notOneUintInt8Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintInt8Err(num uint) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapUintUint64Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapUintUint64Err(notOneUintUint64Err, plusOneUintUint64Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint64Err failed")
	}
	r, _ = FilterMapUintUint64Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint64PtrErr failed")
	}

	_, err := FilterMapUintUint64Err(notOneUintUint64Err, plusOneUintUint64Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintUint64Err failed")
	}
	_, err = FilterMapUintUint64Err(notOneUintUint64Err, plusOneUintUint64Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintUint64Err failed")
	}
}
func notOneUintUint64Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintUint64Err(num uint) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapUintUint32Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapUintUint32Err(notOneUintUint32Err, plusOneUintUint32Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint32Err failed")
	}
	r, _ = FilterMapUintUint32Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint32PtrErr failed")
	}

	_, err := FilterMapUintUint32Err(notOneUintUint32Err, plusOneUintUint32Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintUint32Err failed")
	}
	_, err = FilterMapUintUint32Err(notOneUintUint32Err, plusOneUintUint32Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintUint32Err failed")
	}
}
func notOneUintUint32Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintUint32Err(num uint) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapUintUint16Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapUintUint16Err(notOneUintUint16Err, plusOneUintUint16Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint16Err failed")
	}
	r, _ = FilterMapUintUint16Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint16PtrErr failed")
	}

	_, err := FilterMapUintUint16Err(notOneUintUint16Err, plusOneUintUint16Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintUint16Err failed")
	}
	_, err = FilterMapUintUint16Err(notOneUintUint16Err, plusOneUintUint16Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintUint16Err failed")
	}
}
func notOneUintUint16Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintUint16Err(num uint) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapUintUint8Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapUintUint8Err(notOneUintUint8Err, plusOneUintUint8Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint8Err failed")
	}
	r, _ = FilterMapUintUint8Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint8PtrErr failed")
	}

	_, err := FilterMapUintUint8Err(notOneUintUint8Err, plusOneUintUint8Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintUint8Err failed")
	}
	_, err = FilterMapUintUint8Err(notOneUintUint8Err, plusOneUintUint8Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintUint8Err failed")
	}
}
func notOneUintUint8Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintUint8Err(num uint) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapUintStrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint = 1
	var iv2 uint = 2
	var iv3 uint = 3
	var iv10 uint = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapUintStrErr(notOneUintStrNumErr, someLogicUintStrNumErr, []uint{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUintStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintStrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintStrErr failed")
	}

	r, _ = FilterMapUintStrErr(nil, nil, []uint{})
	if len(r) > 0 {
		t.Errorf("FilterMapUintStrErr failed")
	}

	_, err := FilterMapUintStrErr(notOneUintStrNumErr, someLogicUintStrNumErr, []uint{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapUintStrErr failed")
	}
	_, err = FilterMapUintStrErr(notOneUintStrNumErr, someLogicUintStrNumErr, []uint{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapUintStrErr failed")
	}
}

func notOneUintStrNumErr(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUintStrNumErr(num uint) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapUintBoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	var vi10 uint = 10
	var vi0 uint

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapUintBoolErr(notOneUintBoolErr, someLogicUintBoolErr, []uint{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintBoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintBoolErr failed")
	}

	r, _ = FilterMapUintBoolErr(nil, nil, []uint{})
	if len(r) > 0 {
		t.Errorf("FilterMapUintBoolPtrErr failed")
	}

	_, err := FilterMapUintBoolErr(notOneUintBoolErr, someLogicUintBoolErr, []uint{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUintBoolErr failed")
	}

	_, err = FilterMapUintBoolErr(notOneUintBoolErr, someLogicUintBoolErr, []uint{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUintBoolErr failed")
	}
}
func notOneUintBoolErr(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUintBoolErr(num uint) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapUintFloat32Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapUintFloat32Err(notOneUintFloat32Err, plusOneUintFloat32Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintFloat32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat32Err failed")
	}
	r, _ = FilterMapUintFloat32Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat32PtrErr failed")
	}

	_, err := FilterMapUintFloat32Err(notOneUintFloat32Err, plusOneUintFloat32Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat32Err failed")
	}
	_, err = FilterMapUintFloat32Err(notOneUintFloat32Err, plusOneUintFloat32Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat32Err failed")
	}
}
func notOneUintFloat32Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintFloat32Err(num uint) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapUintFloat64Err(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapUintFloat64Err(notOneUintFloat64Err, plusOneUintFloat64Err, []uint{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUintFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintFloat64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat64Err failed")
	}
	r, _ = FilterMapUintFloat64Err(nil, nil, []uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat64PtrErr failed")
	}

	_, err := FilterMapUintFloat64Err(notOneUintFloat64Err, plusOneUintFloat64Err, []uint{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat64Err failed")
	}
	_, err = FilterMapUintFloat64Err(notOneUintFloat64Err, plusOneUintFloat64Err, []uint{v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat64Err failed")
	}
}
func notOneUintFloat64Err(num uint) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUintFloat64Err(num uint) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapUint64IntErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapUint64IntErr(notOneUint64IntErr, plusOneUint64IntErr, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64IntErr failed")
	}
	r, _ = FilterMapUint64IntErr(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64IntPtrErr failed")
	}

	_, err := FilterMapUint64IntErr(notOneUint64IntErr, plusOneUint64IntErr, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64IntErr failed")
	}
	_, err = FilterMapUint64IntErr(notOneUint64IntErr, plusOneUint64IntErr, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64IntErr failed")
	}
}
func notOneUint64IntErr(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64IntErr(num uint64) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapUint64Int64Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapUint64Int64Err(notOneUint64Int64Err, plusOneUint64Int64Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int64Err failed")
	}
	r, _ = FilterMapUint64Int64Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int64PtrErr failed")
	}

	_, err := FilterMapUint64Int64Err(notOneUint64Int64Err, plusOneUint64Int64Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int64Err failed")
	}
	_, err = FilterMapUint64Int64Err(notOneUint64Int64Err, plusOneUint64Int64Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int64Err failed")
	}
}
func notOneUint64Int64Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Int64Err(num uint64) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapUint64Int32Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapUint64Int32Err(notOneUint64Int32Err, plusOneUint64Int32Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int32Err failed")
	}
	r, _ = FilterMapUint64Int32Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int32PtrErr failed")
	}

	_, err := FilterMapUint64Int32Err(notOneUint64Int32Err, plusOneUint64Int32Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int32Err failed")
	}
	_, err = FilterMapUint64Int32Err(notOneUint64Int32Err, plusOneUint64Int32Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int32Err failed")
	}
}
func notOneUint64Int32Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Int32Err(num uint64) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapUint64Int16Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapUint64Int16Err(notOneUint64Int16Err, plusOneUint64Int16Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int16Err failed")
	}
	r, _ = FilterMapUint64Int16Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int16PtrErr failed")
	}

	_, err := FilterMapUint64Int16Err(notOneUint64Int16Err, plusOneUint64Int16Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int16Err failed")
	}
	_, err = FilterMapUint64Int16Err(notOneUint64Int16Err, plusOneUint64Int16Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int16Err failed")
	}
}
func notOneUint64Int16Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Int16Err(num uint64) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapUint64Int8Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapUint64Int8Err(notOneUint64Int8Err, plusOneUint64Int8Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int8Err failed")
	}
	r, _ = FilterMapUint64Int8Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int8PtrErr failed")
	}

	_, err := FilterMapUint64Int8Err(notOneUint64Int8Err, plusOneUint64Int8Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int8Err failed")
	}
	_, err = FilterMapUint64Int8Err(notOneUint64Int8Err, plusOneUint64Int8Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int8Err failed")
	}
}
func notOneUint64Int8Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Int8Err(num uint64) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapUint64UintErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapUint64UintErr(notOneUint64UintErr, plusOneUint64UintErr, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64UintErr failed")
	}
	r, _ = FilterMapUint64UintErr(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64UintPtrErr failed")
	}

	_, err := FilterMapUint64UintErr(notOneUint64UintErr, plusOneUint64UintErr, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64UintErr failed")
	}
	_, err = FilterMapUint64UintErr(notOneUint64UintErr, plusOneUint64UintErr, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64UintErr failed")
	}
}
func notOneUint64UintErr(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64UintErr(num uint64) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapUint64Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapUint64Uint32Err(notOneUint64Uint32Err, plusOneUint64Uint32Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint32Err failed")
	}
	r, _ = FilterMapUint64Uint32Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint32PtrErr failed")
	}

	_, err := FilterMapUint64Uint32Err(notOneUint64Uint32Err, plusOneUint64Uint32Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint32Err failed")
	}
	_, err = FilterMapUint64Uint32Err(notOneUint64Uint32Err, plusOneUint64Uint32Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint32Err failed")
	}
}
func notOneUint64Uint32Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Uint32Err(num uint64) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapUint64Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapUint64Uint16Err(notOneUint64Uint16Err, plusOneUint64Uint16Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint16Err failed")
	}
	r, _ = FilterMapUint64Uint16Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint16PtrErr failed")
	}

	_, err := FilterMapUint64Uint16Err(notOneUint64Uint16Err, plusOneUint64Uint16Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint16Err failed")
	}
	_, err = FilterMapUint64Uint16Err(notOneUint64Uint16Err, plusOneUint64Uint16Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint16Err failed")
	}
}
func notOneUint64Uint16Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Uint16Err(num uint64) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapUint64Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapUint64Uint8Err(notOneUint64Uint8Err, plusOneUint64Uint8Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint8Err failed")
	}
	r, _ = FilterMapUint64Uint8Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint8PtrErr failed")
	}

	_, err := FilterMapUint64Uint8Err(notOneUint64Uint8Err, plusOneUint64Uint8Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint8Err failed")
	}
	_, err = FilterMapUint64Uint8Err(notOneUint64Uint8Err, plusOneUint64Uint8Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint8Err failed")
	}
}
func notOneUint64Uint8Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Uint8Err(num uint64) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapUint64StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint64 = 1
	var iv2 uint64 = 2
	var iv3 uint64 = 3
	var iv10 uint64 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapUint64StrErr(notOneUint64StrNumErr, someLogicUint64StrNumErr, []uint64{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64StrErr failed")
	}

	r, _ = FilterMapUint64StrErr(nil, nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint64StrErr failed")
	}

	_, err := FilterMapUint64StrErr(notOneUint64StrNumErr, someLogicUint64StrNumErr, []uint64{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapUint64StrErr failed")
	}
	_, err = FilterMapUint64StrErr(notOneUint64StrNumErr, someLogicUint64StrNumErr, []uint64{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapUint64StrErr failed")
	}
}

func notOneUint64StrNumErr(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint64StrNumErr(num uint64) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapUint64BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	var vi10 uint64 = 10
	var vi0 uint64

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapUint64BoolErr(notOneUint64BoolErr, someLogicUint64BoolErr, []uint64{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64BoolErr failed")
	}

	r, _ = FilterMapUint64BoolErr(nil, nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint64BoolPtrErr failed")
	}

	_, err := FilterMapUint64BoolErr(notOneUint64BoolErr, someLogicUint64BoolErr, []uint64{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint64BoolErr failed")
	}

	_, err = FilterMapUint64BoolErr(notOneUint64BoolErr, someLogicUint64BoolErr, []uint64{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint64BoolErr failed")
	}
}
func notOneUint64BoolErr(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint64BoolErr(num uint64) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapUint64Float32Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapUint64Float32Err(notOneUint64Float32Err, plusOneUint64Float32Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float32Err failed")
	}
	r, _ = FilterMapUint64Float32Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float32PtrErr failed")
	}

	_, err := FilterMapUint64Float32Err(notOneUint64Float32Err, plusOneUint64Float32Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float32Err failed")
	}
	_, err = FilterMapUint64Float32Err(notOneUint64Float32Err, plusOneUint64Float32Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float32Err failed")
	}
}
func notOneUint64Float32Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Float32Err(num uint64) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapUint64Float64Err(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapUint64Float64Err(notOneUint64Float64Err, plusOneUint64Float64Err, []uint64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint64Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float64Err failed")
	}
	r, _ = FilterMapUint64Float64Err(nil, nil, []uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float64PtrErr failed")
	}

	_, err := FilterMapUint64Float64Err(notOneUint64Float64Err, plusOneUint64Float64Err, []uint64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float64Err failed")
	}
	_, err = FilterMapUint64Float64Err(notOneUint64Float64Err, plusOneUint64Float64Err, []uint64{v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float64Err failed")
	}
}
func notOneUint64Float64Err(num uint64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint64Float64Err(num uint64) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapUint32IntErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapUint32IntErr(notOneUint32IntErr, plusOneUint32IntErr, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32IntErr failed")
	}
	r, _ = FilterMapUint32IntErr(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32IntPtrErr failed")
	}

	_, err := FilterMapUint32IntErr(notOneUint32IntErr, plusOneUint32IntErr, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32IntErr failed")
	}
	_, err = FilterMapUint32IntErr(notOneUint32IntErr, plusOneUint32IntErr, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32IntErr failed")
	}
}
func notOneUint32IntErr(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32IntErr(num uint32) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapUint32Int64Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapUint32Int64Err(notOneUint32Int64Err, plusOneUint32Int64Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int64Err failed")
	}
	r, _ = FilterMapUint32Int64Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int64PtrErr failed")
	}

	_, err := FilterMapUint32Int64Err(notOneUint32Int64Err, plusOneUint32Int64Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int64Err failed")
	}
	_, err = FilterMapUint32Int64Err(notOneUint32Int64Err, plusOneUint32Int64Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int64Err failed")
	}
}
func notOneUint32Int64Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Int64Err(num uint32) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapUint32Int32Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapUint32Int32Err(notOneUint32Int32Err, plusOneUint32Int32Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int32Err failed")
	}
	r, _ = FilterMapUint32Int32Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int32PtrErr failed")
	}

	_, err := FilterMapUint32Int32Err(notOneUint32Int32Err, plusOneUint32Int32Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int32Err failed")
	}
	_, err = FilterMapUint32Int32Err(notOneUint32Int32Err, plusOneUint32Int32Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int32Err failed")
	}
}
func notOneUint32Int32Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Int32Err(num uint32) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapUint32Int16Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapUint32Int16Err(notOneUint32Int16Err, plusOneUint32Int16Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int16Err failed")
	}
	r, _ = FilterMapUint32Int16Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int16PtrErr failed")
	}

	_, err := FilterMapUint32Int16Err(notOneUint32Int16Err, plusOneUint32Int16Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int16Err failed")
	}
	_, err = FilterMapUint32Int16Err(notOneUint32Int16Err, plusOneUint32Int16Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int16Err failed")
	}
}
func notOneUint32Int16Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Int16Err(num uint32) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapUint32Int8Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapUint32Int8Err(notOneUint32Int8Err, plusOneUint32Int8Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int8Err failed")
	}
	r, _ = FilterMapUint32Int8Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int8PtrErr failed")
	}

	_, err := FilterMapUint32Int8Err(notOneUint32Int8Err, plusOneUint32Int8Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int8Err failed")
	}
	_, err = FilterMapUint32Int8Err(notOneUint32Int8Err, plusOneUint32Int8Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int8Err failed")
	}
}
func notOneUint32Int8Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Int8Err(num uint32) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapUint32UintErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapUint32UintErr(notOneUint32UintErr, plusOneUint32UintErr, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32UintErr failed")
	}
	r, _ = FilterMapUint32UintErr(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32UintPtrErr failed")
	}

	_, err := FilterMapUint32UintErr(notOneUint32UintErr, plusOneUint32UintErr, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32UintErr failed")
	}
	_, err = FilterMapUint32UintErr(notOneUint32UintErr, plusOneUint32UintErr, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32UintErr failed")
	}
}
func notOneUint32UintErr(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32UintErr(num uint32) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapUint32Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapUint32Uint64Err(notOneUint32Uint64Err, plusOneUint32Uint64Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint64Err failed")
	}
	r, _ = FilterMapUint32Uint64Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint64PtrErr failed")
	}

	_, err := FilterMapUint32Uint64Err(notOneUint32Uint64Err, plusOneUint32Uint64Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint64Err failed")
	}
	_, err = FilterMapUint32Uint64Err(notOneUint32Uint64Err, plusOneUint32Uint64Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint64Err failed")
	}
}
func notOneUint32Uint64Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Uint64Err(num uint32) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapUint32Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapUint32Uint16Err(notOneUint32Uint16Err, plusOneUint32Uint16Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint16Err failed")
	}
	r, _ = FilterMapUint32Uint16Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint16PtrErr failed")
	}

	_, err := FilterMapUint32Uint16Err(notOneUint32Uint16Err, plusOneUint32Uint16Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint16Err failed")
	}
	_, err = FilterMapUint32Uint16Err(notOneUint32Uint16Err, plusOneUint32Uint16Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint16Err failed")
	}
}
func notOneUint32Uint16Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Uint16Err(num uint32) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapUint32Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapUint32Uint8Err(notOneUint32Uint8Err, plusOneUint32Uint8Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint8Err failed")
	}
	r, _ = FilterMapUint32Uint8Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint8PtrErr failed")
	}

	_, err := FilterMapUint32Uint8Err(notOneUint32Uint8Err, plusOneUint32Uint8Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint8Err failed")
	}
	_, err = FilterMapUint32Uint8Err(notOneUint32Uint8Err, plusOneUint32Uint8Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint8Err failed")
	}
}
func notOneUint32Uint8Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Uint8Err(num uint32) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapUint32StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint32 = 1
	var iv2 uint32 = 2
	var iv3 uint32 = 3
	var iv10 uint32 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapUint32StrErr(notOneUint32StrNumErr, someLogicUint32StrNumErr, []uint32{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32StrErr failed")
	}

	r, _ = FilterMapUint32StrErr(nil, nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint32StrErr failed")
	}

	_, err := FilterMapUint32StrErr(notOneUint32StrNumErr, someLogicUint32StrNumErr, []uint32{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapUint32StrErr failed")
	}
	_, err = FilterMapUint32StrErr(notOneUint32StrNumErr, someLogicUint32StrNumErr, []uint32{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapUint32StrErr failed")
	}
}

func notOneUint32StrNumErr(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint32StrNumErr(num uint32) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapUint32BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	var vi10 uint32 = 10
	var vi0 uint32

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapUint32BoolErr(notOneUint32BoolErr, someLogicUint32BoolErr, []uint32{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32BoolErr failed")
	}

	r, _ = FilterMapUint32BoolErr(nil, nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint32BoolPtrErr failed")
	}

	_, err := FilterMapUint32BoolErr(notOneUint32BoolErr, someLogicUint32BoolErr, []uint32{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint32BoolErr failed")
	}

	_, err = FilterMapUint32BoolErr(notOneUint32BoolErr, someLogicUint32BoolErr, []uint32{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint32BoolErr failed")
	}
}
func notOneUint32BoolErr(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint32BoolErr(num uint32) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapUint32Float32Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapUint32Float32Err(notOneUint32Float32Err, plusOneUint32Float32Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float32Err failed")
	}
	r, _ = FilterMapUint32Float32Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float32PtrErr failed")
	}

	_, err := FilterMapUint32Float32Err(notOneUint32Float32Err, plusOneUint32Float32Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float32Err failed")
	}
	_, err = FilterMapUint32Float32Err(notOneUint32Float32Err, plusOneUint32Float32Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float32Err failed")
	}
}
func notOneUint32Float32Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Float32Err(num uint32) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapUint32Float64Err(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapUint32Float64Err(notOneUint32Float64Err, plusOneUint32Float64Err, []uint32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float64Err failed")
	}
	r, _ = FilterMapUint32Float64Err(nil, nil, []uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float64PtrErr failed")
	}

	_, err := FilterMapUint32Float64Err(notOneUint32Float64Err, plusOneUint32Float64Err, []uint32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float64Err failed")
	}
	_, err = FilterMapUint32Float64Err(notOneUint32Float64Err, plusOneUint32Float64Err, []uint32{v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float64Err failed")
	}
}
func notOneUint32Float64Err(num uint32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint32Float64Err(num uint32) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapUint16IntErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapUint16IntErr(notOneUint16IntErr, plusOneUint16IntErr, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16IntErr failed")
	}
	r, _ = FilterMapUint16IntErr(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16IntPtrErr failed")
	}

	_, err := FilterMapUint16IntErr(notOneUint16IntErr, plusOneUint16IntErr, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16IntErr failed")
	}
	_, err = FilterMapUint16IntErr(notOneUint16IntErr, plusOneUint16IntErr, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16IntErr failed")
	}
}
func notOneUint16IntErr(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16IntErr(num uint16) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapUint16Int64Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapUint16Int64Err(notOneUint16Int64Err, plusOneUint16Int64Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int64Err failed")
	}
	r, _ = FilterMapUint16Int64Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int64PtrErr failed")
	}

	_, err := FilterMapUint16Int64Err(notOneUint16Int64Err, plusOneUint16Int64Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int64Err failed")
	}
	_, err = FilterMapUint16Int64Err(notOneUint16Int64Err, plusOneUint16Int64Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int64Err failed")
	}
}
func notOneUint16Int64Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Int64Err(num uint16) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapUint16Int32Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapUint16Int32Err(notOneUint16Int32Err, plusOneUint16Int32Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int32Err failed")
	}
	r, _ = FilterMapUint16Int32Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int32PtrErr failed")
	}

	_, err := FilterMapUint16Int32Err(notOneUint16Int32Err, plusOneUint16Int32Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int32Err failed")
	}
	_, err = FilterMapUint16Int32Err(notOneUint16Int32Err, plusOneUint16Int32Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int32Err failed")
	}
}
func notOneUint16Int32Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Int32Err(num uint16) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapUint16Int16Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapUint16Int16Err(notOneUint16Int16Err, plusOneUint16Int16Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int16Err failed")
	}
	r, _ = FilterMapUint16Int16Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int16PtrErr failed")
	}

	_, err := FilterMapUint16Int16Err(notOneUint16Int16Err, plusOneUint16Int16Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int16Err failed")
	}
	_, err = FilterMapUint16Int16Err(notOneUint16Int16Err, plusOneUint16Int16Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int16Err failed")
	}
}
func notOneUint16Int16Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Int16Err(num uint16) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapUint16Int8Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapUint16Int8Err(notOneUint16Int8Err, plusOneUint16Int8Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int8Err failed")
	}
	r, _ = FilterMapUint16Int8Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int8PtrErr failed")
	}

	_, err := FilterMapUint16Int8Err(notOneUint16Int8Err, plusOneUint16Int8Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int8Err failed")
	}
	_, err = FilterMapUint16Int8Err(notOneUint16Int8Err, plusOneUint16Int8Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int8Err failed")
	}
}
func notOneUint16Int8Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Int8Err(num uint16) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapUint16UintErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapUint16UintErr(notOneUint16UintErr, plusOneUint16UintErr, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16UintErr failed")
	}
	r, _ = FilterMapUint16UintErr(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16UintPtrErr failed")
	}

	_, err := FilterMapUint16UintErr(notOneUint16UintErr, plusOneUint16UintErr, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16UintErr failed")
	}
	_, err = FilterMapUint16UintErr(notOneUint16UintErr, plusOneUint16UintErr, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16UintErr failed")
	}
}
func notOneUint16UintErr(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16UintErr(num uint16) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapUint16Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapUint16Uint64Err(notOneUint16Uint64Err, plusOneUint16Uint64Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint64Err failed")
	}
	r, _ = FilterMapUint16Uint64Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint64PtrErr failed")
	}

	_, err := FilterMapUint16Uint64Err(notOneUint16Uint64Err, plusOneUint16Uint64Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint64Err failed")
	}
	_, err = FilterMapUint16Uint64Err(notOneUint16Uint64Err, plusOneUint16Uint64Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint64Err failed")
	}
}
func notOneUint16Uint64Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Uint64Err(num uint16) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapUint16Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapUint16Uint32Err(notOneUint16Uint32Err, plusOneUint16Uint32Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint32Err failed")
	}
	r, _ = FilterMapUint16Uint32Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint32PtrErr failed")
	}

	_, err := FilterMapUint16Uint32Err(notOneUint16Uint32Err, plusOneUint16Uint32Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint32Err failed")
	}
	_, err = FilterMapUint16Uint32Err(notOneUint16Uint32Err, plusOneUint16Uint32Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint32Err failed")
	}
}
func notOneUint16Uint32Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Uint32Err(num uint16) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapUint16Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapUint16Uint8Err(notOneUint16Uint8Err, plusOneUint16Uint8Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint8Err failed")
	}
	r, _ = FilterMapUint16Uint8Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint8PtrErr failed")
	}

	_, err := FilterMapUint16Uint8Err(notOneUint16Uint8Err, plusOneUint16Uint8Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint8Err failed")
	}
	_, err = FilterMapUint16Uint8Err(notOneUint16Uint8Err, plusOneUint16Uint8Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint8Err failed")
	}
}
func notOneUint16Uint8Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Uint8Err(num uint16) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapUint16StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint16 = 1
	var iv2 uint16 = 2
	var iv3 uint16 = 3
	var iv10 uint16 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapUint16StrErr(notOneUint16StrNumErr, someLogicUint16StrNumErr, []uint16{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint16StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16StrErr failed")
	}

	r, _ = FilterMapUint16StrErr(nil, nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint16StrErr failed")
	}

	_, err := FilterMapUint16StrErr(notOneUint16StrNumErr, someLogicUint16StrNumErr, []uint16{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapUint16StrErr failed")
	}
	_, err = FilterMapUint16StrErr(notOneUint16StrNumErr, someLogicUint16StrNumErr, []uint16{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapUint16StrErr failed")
	}
}

func notOneUint16StrNumErr(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint16StrNumErr(num uint16) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapUint16BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	var vi10 uint16 = 10
	var vi0 uint16

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapUint16BoolErr(notOneUint16BoolErr, someLogicUint16BoolErr, []uint16{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16BoolErr failed")
	}

	r, _ = FilterMapUint16BoolErr(nil, nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint16BoolPtrErr failed")
	}

	_, err := FilterMapUint16BoolErr(notOneUint16BoolErr, someLogicUint16BoolErr, []uint16{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint16BoolErr failed")
	}

	_, err = FilterMapUint16BoolErr(notOneUint16BoolErr, someLogicUint16BoolErr, []uint16{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint16BoolErr failed")
	}
}
func notOneUint16BoolErr(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint16BoolErr(num uint16) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapUint16Float32Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapUint16Float32Err(notOneUint16Float32Err, plusOneUint16Float32Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float32Err failed")
	}
	r, _ = FilterMapUint16Float32Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float32PtrErr failed")
	}

	_, err := FilterMapUint16Float32Err(notOneUint16Float32Err, plusOneUint16Float32Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float32Err failed")
	}
	_, err = FilterMapUint16Float32Err(notOneUint16Float32Err, plusOneUint16Float32Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float32Err failed")
	}
}
func notOneUint16Float32Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Float32Err(num uint16) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapUint16Float64Err(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapUint16Float64Err(notOneUint16Float64Err, plusOneUint16Float64Err, []uint16{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint16Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float64Err failed")
	}
	r, _ = FilterMapUint16Float64Err(nil, nil, []uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float64PtrErr failed")
	}

	_, err := FilterMapUint16Float64Err(notOneUint16Float64Err, plusOneUint16Float64Err, []uint16{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float64Err failed")
	}
	_, err = FilterMapUint16Float64Err(notOneUint16Float64Err, plusOneUint16Float64Err, []uint16{v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float64Err failed")
	}
}
func notOneUint16Float64Err(num uint16) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint16Float64Err(num uint16) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapUint8IntErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapUint8IntErr(notOneUint8IntErr, plusOneUint8IntErr, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8IntErr failed")
	}
	r, _ = FilterMapUint8IntErr(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8IntPtrErr failed")
	}

	_, err := FilterMapUint8IntErr(notOneUint8IntErr, plusOneUint8IntErr, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8IntErr failed")
	}
	_, err = FilterMapUint8IntErr(notOneUint8IntErr, plusOneUint8IntErr, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8IntErr failed")
	}
}
func notOneUint8IntErr(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8IntErr(num uint8) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapUint8Int64Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapUint8Int64Err(notOneUint8Int64Err, plusOneUint8Int64Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int64Err failed")
	}
	r, _ = FilterMapUint8Int64Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int64PtrErr failed")
	}

	_, err := FilterMapUint8Int64Err(notOneUint8Int64Err, plusOneUint8Int64Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int64Err failed")
	}
	_, err = FilterMapUint8Int64Err(notOneUint8Int64Err, plusOneUint8Int64Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int64Err failed")
	}
}
func notOneUint8Int64Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Int64Err(num uint8) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapUint8Int32Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapUint8Int32Err(notOneUint8Int32Err, plusOneUint8Int32Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int32Err failed")
	}
	r, _ = FilterMapUint8Int32Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int32PtrErr failed")
	}

	_, err := FilterMapUint8Int32Err(notOneUint8Int32Err, plusOneUint8Int32Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int32Err failed")
	}
	_, err = FilterMapUint8Int32Err(notOneUint8Int32Err, plusOneUint8Int32Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int32Err failed")
	}
}
func notOneUint8Int32Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Int32Err(num uint8) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapUint8Int16Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapUint8Int16Err(notOneUint8Int16Err, plusOneUint8Int16Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int16Err failed")
	}
	r, _ = FilterMapUint8Int16Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int16PtrErr failed")
	}

	_, err := FilterMapUint8Int16Err(notOneUint8Int16Err, plusOneUint8Int16Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int16Err failed")
	}
	_, err = FilterMapUint8Int16Err(notOneUint8Int16Err, plusOneUint8Int16Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int16Err failed")
	}
}
func notOneUint8Int16Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Int16Err(num uint8) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapUint8Int8Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapUint8Int8Err(notOneUint8Int8Err, plusOneUint8Int8Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int8Err failed")
	}
	r, _ = FilterMapUint8Int8Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int8PtrErr failed")
	}

	_, err := FilterMapUint8Int8Err(notOneUint8Int8Err, plusOneUint8Int8Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int8Err failed")
	}
	_, err = FilterMapUint8Int8Err(notOneUint8Int8Err, plusOneUint8Int8Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int8Err failed")
	}
}
func notOneUint8Int8Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Int8Err(num uint8) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapUint8UintErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapUint8UintErr(notOneUint8UintErr, plusOneUint8UintErr, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8UintErr failed")
	}
	r, _ = FilterMapUint8UintErr(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8UintPtrErr failed")
	}

	_, err := FilterMapUint8UintErr(notOneUint8UintErr, plusOneUint8UintErr, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8UintErr failed")
	}
	_, err = FilterMapUint8UintErr(notOneUint8UintErr, plusOneUint8UintErr, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8UintErr failed")
	}
}
func notOneUint8UintErr(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8UintErr(num uint8) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapUint8Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapUint8Uint64Err(notOneUint8Uint64Err, plusOneUint8Uint64Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint64Err failed")
	}
	r, _ = FilterMapUint8Uint64Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint64PtrErr failed")
	}

	_, err := FilterMapUint8Uint64Err(notOneUint8Uint64Err, plusOneUint8Uint64Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint64Err failed")
	}
	_, err = FilterMapUint8Uint64Err(notOneUint8Uint64Err, plusOneUint8Uint64Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint64Err failed")
	}
}
func notOneUint8Uint64Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Uint64Err(num uint8) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapUint8Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapUint8Uint32Err(notOneUint8Uint32Err, plusOneUint8Uint32Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint32Err failed")
	}
	r, _ = FilterMapUint8Uint32Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint32PtrErr failed")
	}

	_, err := FilterMapUint8Uint32Err(notOneUint8Uint32Err, plusOneUint8Uint32Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint32Err failed")
	}
	_, err = FilterMapUint8Uint32Err(notOneUint8Uint32Err, plusOneUint8Uint32Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint32Err failed")
	}
}
func notOneUint8Uint32Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Uint32Err(num uint8) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapUint8Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapUint8Uint16Err(notOneUint8Uint16Err, plusOneUint8Uint16Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint16Err failed")
	}
	r, _ = FilterMapUint8Uint16Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint16PtrErr failed")
	}

	_, err := FilterMapUint8Uint16Err(notOneUint8Uint16Err, plusOneUint8Uint16Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint16Err failed")
	}
	_, err = FilterMapUint8Uint16Err(notOneUint8Uint16Err, plusOneUint8Uint16Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint16Err failed")
	}
}
func notOneUint8Uint16Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Uint16Err(num uint8) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapUint8StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint8 = 1
	var iv2 uint8 = 2
	var iv3 uint8 = 3
	var iv10 uint8 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapUint8StrErr(notOneUint8StrNumErr, someLogicUint8StrNumErr, []uint8{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapUint8StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8StrErr failed")
	}

	r, _ = FilterMapUint8StrErr(nil, nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint8StrErr failed")
	}

	_, err := FilterMapUint8StrErr(notOneUint8StrNumErr, someLogicUint8StrNumErr, []uint8{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapUint8StrErr failed")
	}
	_, err = FilterMapUint8StrErr(notOneUint8StrNumErr, someLogicUint8StrNumErr, []uint8{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapUint8StrErr failed")
	}
}

func notOneUint8StrNumErr(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint8StrNumErr(num uint8) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapUint8BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	var vi10 uint8 = 10
	var vi0 uint8

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapUint8BoolErr(notOneUint8BoolErr, someLogicUint8BoolErr, []uint8{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8BoolErr failed")
	}

	r, _ = FilterMapUint8BoolErr(nil, nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint8BoolPtrErr failed")
	}

	_, err := FilterMapUint8BoolErr(notOneUint8BoolErr, someLogicUint8BoolErr, []uint8{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint8BoolErr failed")
	}

	_, err = FilterMapUint8BoolErr(notOneUint8BoolErr, someLogicUint8BoolErr, []uint8{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapUint8BoolErr failed")
	}
}
func notOneUint8BoolErr(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicUint8BoolErr(num uint8) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapUint8Float32Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapUint8Float32Err(notOneUint8Float32Err, plusOneUint8Float32Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float32Err failed")
	}
	r, _ = FilterMapUint8Float32Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float32PtrErr failed")
	}

	_, err := FilterMapUint8Float32Err(notOneUint8Float32Err, plusOneUint8Float32Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float32Err failed")
	}
	_, err = FilterMapUint8Float32Err(notOneUint8Float32Err, plusOneUint8Float32Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float32Err failed")
	}
}
func notOneUint8Float32Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Float32Err(num uint8) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}

func TestFilterMapUint8Float64Err(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapUint8Float64Err(notOneUint8Float64Err, plusOneUint8Float64Err, []uint8{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapUint8Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float64Err failed")
	}
	r, _ = FilterMapUint8Float64Err(nil, nil, []uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float64PtrErr failed")
	}

	_, err := FilterMapUint8Float64Err(notOneUint8Float64Err, plusOneUint8Float64Err, []uint8{v2, v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float64Err failed")
	}
	_, err = FilterMapUint8Float64Err(notOneUint8Float64Err, plusOneUint8Float64Err, []uint8{v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float64Err failed")
	}
}
func notOneUint8Float64Err(num uint8) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneUint8Float64Err(num uint8) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapStrIntErr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []int{vo10}
	newList, _ := FilterMapStrIntErr(notOneStrIntStrErr, someLogicStrIntStrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrIntErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrIntErr failed")
	}

	r, _ = FilterMapStrIntErr(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrIntErr failed")
	}

	_, err := FilterMapStrIntErr(notOneStrIntStrErr, someLogicStrIntStrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrIntErr failed")
	}
	_, err = FilterMapStrIntErr(notOneStrIntStrErr, someLogicStrIntStrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrIntErr failed")
	}
}
func notOneStrIntStrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrIntStrErr(num string) (int, error) {
	var r int = int(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = int(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrInt64Err(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []int64{vo10}
	newList, _ := FilterMapStrInt64Err(notOneStrInt64StrErr, someLogicStrInt64StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt64Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt64Err failed")
	}

	r, _ = FilterMapStrInt64Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt64Err failed")
	}

	_, err := FilterMapStrInt64Err(notOneStrInt64StrErr, someLogicStrInt64StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt64Err failed")
	}
	_, err = FilterMapStrInt64Err(notOneStrInt64StrErr, someLogicStrInt64StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt64Err failed")
	}
}
func notOneStrInt64StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrInt64StrErr(num string) (int64, error) {
	var r int64 = int64(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = int64(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrInt32Err(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []int32{vo10}
	newList, _ := FilterMapStrInt32Err(notOneStrInt32StrErr, someLogicStrInt32StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt32Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt32Err failed")
	}

	r, _ = FilterMapStrInt32Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt32Err failed")
	}

	_, err := FilterMapStrInt32Err(notOneStrInt32StrErr, someLogicStrInt32StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt32Err failed")
	}
	_, err = FilterMapStrInt32Err(notOneStrInt32StrErr, someLogicStrInt32StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt32Err failed")
	}
}
func notOneStrInt32StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrInt32StrErr(num string) (int32, error) {
	var r int32 = int32(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = int32(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrInt16Err(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []int16{vo10}
	newList, _ := FilterMapStrInt16Err(notOneStrInt16StrErr, someLogicStrInt16StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt16Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt16Err failed")
	}

	r, _ = FilterMapStrInt16Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt16Err failed")
	}

	_, err := FilterMapStrInt16Err(notOneStrInt16StrErr, someLogicStrInt16StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt16Err failed")
	}
	_, err = FilterMapStrInt16Err(notOneStrInt16StrErr, someLogicStrInt16StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt16Err failed")
	}
}
func notOneStrInt16StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrInt16StrErr(num string) (int16, error) {
	var r int16 = int16(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = int16(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrInt8Err(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []int8{vo10}
	newList, _ := FilterMapStrInt8Err(notOneStrInt8StrErr, someLogicStrInt8StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt8Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt8Err failed")
	}

	r, _ = FilterMapStrInt8Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt8Err failed")
	}

	_, err := FilterMapStrInt8Err(notOneStrInt8StrErr, someLogicStrInt8StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt8Err failed")
	}
	_, err = FilterMapStrInt8Err(notOneStrInt8StrErr, someLogicStrInt8StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt8Err failed")
	}
}
func notOneStrInt8StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrInt8StrErr(num string) (int8, error) {
	var r int8 = int8(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = int8(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrUintErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []uint{vo10}
	newList, _ := FilterMapStrUintErr(notOneStrUintStrErr, someLogicStrUintStrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUintErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUintErr failed")
	}

	r, _ = FilterMapStrUintErr(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUintErr failed")
	}

	_, err := FilterMapStrUintErr(notOneStrUintStrErr, someLogicStrUintStrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrUintErr failed")
	}
	_, err = FilterMapStrUintErr(notOneStrUintStrErr, someLogicStrUintStrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrUintErr failed")
	}
}
func notOneStrUintStrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrUintStrErr(num string) (uint, error) {
	var r uint = uint(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = uint(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrUint64Err(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []uint64{vo10}
	newList, _ := FilterMapStrUint64Err(notOneStrUint64StrErr, someLogicStrUint64StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint64Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint64Err failed")
	}

	r, _ = FilterMapStrUint64Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint64Err failed")
	}

	_, err := FilterMapStrUint64Err(notOneStrUint64StrErr, someLogicStrUint64StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint64Err failed")
	}
	_, err = FilterMapStrUint64Err(notOneStrUint64StrErr, someLogicStrUint64StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint64Err failed")
	}
}
func notOneStrUint64StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrUint64StrErr(num string) (uint64, error) {
	var r uint64 = uint64(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = uint64(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrUint32Err(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []uint32{vo10}
	newList, _ := FilterMapStrUint32Err(notOneStrUint32StrErr, someLogicStrUint32StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint32Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint32Err failed")
	}

	r, _ = FilterMapStrUint32Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint32Err failed")
	}

	_, err := FilterMapStrUint32Err(notOneStrUint32StrErr, someLogicStrUint32StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint32Err failed")
	}
	_, err = FilterMapStrUint32Err(notOneStrUint32StrErr, someLogicStrUint32StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint32Err failed")
	}
}
func notOneStrUint32StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrUint32StrErr(num string) (uint32, error) {
	var r uint32 = uint32(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = uint32(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrUint16Err(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []uint16{vo10}
	newList, _ := FilterMapStrUint16Err(notOneStrUint16StrErr, someLogicStrUint16StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint16Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint16Err failed")
	}

	r, _ = FilterMapStrUint16Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint16Err failed")
	}

	_, err := FilterMapStrUint16Err(notOneStrUint16StrErr, someLogicStrUint16StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint16Err failed")
	}
	_, err = FilterMapStrUint16Err(notOneStrUint16StrErr, someLogicStrUint16StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint16Err failed")
	}
}
func notOneStrUint16StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrUint16StrErr(num string) (uint16, error) {
	var r uint16 = uint16(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = uint16(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrUint8Err(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []uint8{vo10}
	newList, _ := FilterMapStrUint8Err(notOneStrUint8StrErr, someLogicStrUint8StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint8Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint8Err failed")
	}

	r, _ = FilterMapStrUint8Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint8Err failed")
	}

	_, err := FilterMapStrUint8Err(notOneStrUint8StrErr, someLogicStrUint8StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint8Err failed")
	}
	_, err = FilterMapStrUint8Err(notOneStrUint8StrErr, someLogicStrUint8StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint8Err failed")
	}
}
func notOneStrUint8StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrUint8StrErr(num string) (uint8, error) {
	var r uint8 = uint8(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = uint8(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrBoolErr(t *testing.T) {
	// Test : someLogic

	var vto bool = true
	var vfo bool = false

	var vi1 string = "1"
	var vi2 string = "2"
	var vi3 string = "3"
	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapStrBoolErr(notOneStrBoolErr, someLogicStrBoolErr, []string{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapStrBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrBoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrBoolErr failed")
	}

	r, _ = FilterMapStrBoolErr(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrBoolErr failed")
	}

	_, err := FilterMapStrBoolErr(notOneStrBoolErr, someLogicStrBoolErr, []string{vi1, vi10, vi2})
	if err == nil {
		t.Errorf("FilterMapStrBoolErr failed")
	}

	_, err = FilterMapStrBoolErr(notOneStrBoolErr, someLogicStrBoolErr, []string{vi1, vi10, vi3})
	if err == nil {
		t.Errorf("FilterMapStrBoolErr failed")
	}
}
func notOneStrBoolErr(num string) (bool, error) {
	if num == "2" {
		return false, errors.New("2 is not valid value for this test")
	}
	return num != "1", nil
}

func someLogicStrBoolErr(num string) (bool, error) {
	if num == "3" {
		return false, errors.New("3 is not valid value for this test")
	}
	var t bool = true
	var f bool = false

	if num == "10" {
		return t, nil
	}
	return f, nil
}

func TestFilterMapStrFloat32Err(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []float32{vo10}
	newList, _ := FilterMapStrFloat32Err(notOneStrFloat32StrErr, someLogicStrFloat32StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrFloat32Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat32Err failed")
	}

	r, _ = FilterMapStrFloat32Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat32Err failed")
	}

	_, err := FilterMapStrFloat32Err(notOneStrFloat32StrErr, someLogicStrFloat32StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat32Err failed")
	}
	_, err = FilterMapStrFloat32Err(notOneStrFloat32StrErr, someLogicStrFloat32StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat32Err failed")
	}
}
func notOneStrFloat32StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrFloat32StrErr(num string) (float32, error) {
	var r float32 = float32(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = float32(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapStrFloat64Err(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []float64{vo10}
	newList, _ := FilterMapStrFloat64Err(notOneStrFloat64StrErr, someLogicStrFloat64StrErr, []string{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapStrFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrFloat64Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat64Err failed")
	}

	r, _ = FilterMapStrFloat64Err(nil, nil, []string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat64Err failed")
	}

	_, err := FilterMapStrFloat64Err(notOneStrFloat64StrErr, someLogicStrFloat64StrErr, []string{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat64Err failed")
	}
	_, err = FilterMapStrFloat64Err(notOneStrFloat64StrErr, someLogicStrFloat64StrErr, []string{vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat64Err failed")
	}
}
func notOneStrFloat64StrErr(num string) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogicStrFloat64StrErr(num string) (float64, error) {
	var r float64 = float64(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = float64(10)
		return r, nil
	}
	return r, nil
}

func TestFilterMapBoolIntErr(t *testing.T) {
	// Test : someLogic

	var vo10 int = 10

	var vit bool = true
	var vif bool = false

	expectedList := []int{vo10, vo10}
	newList, _ := FilterMapBoolIntErr(notOneBoolIntErr, someLogicBoolIntErr, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolIntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolIntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolIntErr failed")
	}

	r, _ = FilterMapBoolIntErr(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolIntErr failed")
	}
	_, err := FilterMapBoolIntErr(notOneBoolIntErr, someLogicBoolIntErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolIntErr failed")
	}

	_, err = FilterMapBoolIntErr(notOneBoolIntErr2, someLogicBoolIntErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolIntErr failed")
	}
}
func notOneBoolIntErr(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolIntErr2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolIntErr(num bool) (int, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 int = 10
	var v0 int

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolInt64Err(t *testing.T) {
	// Test : someLogic

	var vo10 int64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []int64{vo10, vo10}
	newList, _ := FilterMapBoolInt64Err(notOneBoolInt64Err, someLogicBoolInt64Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt64Err failed")
	}

	r, _ = FilterMapBoolInt64Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt64Err failed")
	}
	_, err := FilterMapBoolInt64Err(notOneBoolInt64Err, someLogicBoolInt64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt64Err failed")
	}

	_, err = FilterMapBoolInt64Err(notOneBoolInt64Err2, someLogicBoolInt64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt64Err failed")
	}
}
func notOneBoolInt64Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolInt64Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolInt64Err(num bool) (int64, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 int64 = 10
	var v0 int64

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolInt32Err(t *testing.T) {
	// Test : someLogic

	var vo10 int32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []int32{vo10, vo10}
	newList, _ := FilterMapBoolInt32Err(notOneBoolInt32Err, someLogicBoolInt32Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt32Err failed")
	}

	r, _ = FilterMapBoolInt32Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt32Err failed")
	}
	_, err := FilterMapBoolInt32Err(notOneBoolInt32Err, someLogicBoolInt32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt32Err failed")
	}

	_, err = FilterMapBoolInt32Err(notOneBoolInt32Err2, someLogicBoolInt32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt32Err failed")
	}
}
func notOneBoolInt32Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolInt32Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolInt32Err(num bool) (int32, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 int32 = 10
	var v0 int32

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolInt16Err(t *testing.T) {
	// Test : someLogic

	var vo10 int16 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []int16{vo10, vo10}
	newList, _ := FilterMapBoolInt16Err(notOneBoolInt16Err, someLogicBoolInt16Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt16Err failed")
	}

	r, _ = FilterMapBoolInt16Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt16Err failed")
	}
	_, err := FilterMapBoolInt16Err(notOneBoolInt16Err, someLogicBoolInt16Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt16Err failed")
	}

	_, err = FilterMapBoolInt16Err(notOneBoolInt16Err2, someLogicBoolInt16Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt16Err failed")
	}
}
func notOneBoolInt16Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolInt16Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolInt16Err(num bool) (int16, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 int16 = 10
	var v0 int16

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolInt8Err(t *testing.T) {
	// Test : someLogic

	var vo10 int8 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []int8{vo10, vo10}
	newList, _ := FilterMapBoolInt8Err(notOneBoolInt8Err, someLogicBoolInt8Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolInt8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt8Err failed")
	}

	r, _ = FilterMapBoolInt8Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt8Err failed")
	}
	_, err := FilterMapBoolInt8Err(notOneBoolInt8Err, someLogicBoolInt8Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt8Err failed")
	}

	_, err = FilterMapBoolInt8Err(notOneBoolInt8Err2, someLogicBoolInt8Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt8Err failed")
	}
}
func notOneBoolInt8Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolInt8Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolInt8Err(num bool) (int8, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 int8 = 10
	var v0 int8

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolUintErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint = 10

	var vit bool = true
	var vif bool = false

	expectedList := []uint{vo10, vo10}
	newList, _ := FilterMapBoolUintErr(notOneBoolUintErr, someLogicBoolUintErr, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUintErr failed")
	}

	r, _ = FilterMapBoolUintErr(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUintErr failed")
	}
	_, err := FilterMapBoolUintErr(notOneBoolUintErr, someLogicBoolUintErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUintErr failed")
	}

	_, err = FilterMapBoolUintErr(notOneBoolUintErr2, someLogicBoolUintErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUintErr failed")
	}
}
func notOneBoolUintErr(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolUintErr2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolUintErr(num bool) (uint, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 uint = 10
	var v0 uint

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolUint64Err(t *testing.T) {
	// Test : someLogic

	var vo10 uint64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []uint64{vo10, vo10}
	newList, _ := FilterMapBoolUint64Err(notOneBoolUint64Err, someLogicBoolUint64Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint64Err failed")
	}

	r, _ = FilterMapBoolUint64Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint64Err failed")
	}
	_, err := FilterMapBoolUint64Err(notOneBoolUint64Err, someLogicBoolUint64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint64Err failed")
	}

	_, err = FilterMapBoolUint64Err(notOneBoolUint64Err2, someLogicBoolUint64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint64Err failed")
	}
}
func notOneBoolUint64Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolUint64Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolUint64Err(num bool) (uint64, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 uint64 = 10
	var v0 uint64

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolUint32Err(t *testing.T) {
	// Test : someLogic

	var vo10 uint32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []uint32{vo10, vo10}
	newList, _ := FilterMapBoolUint32Err(notOneBoolUint32Err, someLogicBoolUint32Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint32Err failed")
	}

	r, _ = FilterMapBoolUint32Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint32Err failed")
	}
	_, err := FilterMapBoolUint32Err(notOneBoolUint32Err, someLogicBoolUint32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint32Err failed")
	}

	_, err = FilterMapBoolUint32Err(notOneBoolUint32Err2, someLogicBoolUint32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint32Err failed")
	}
}
func notOneBoolUint32Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolUint32Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolUint32Err(num bool) (uint32, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 uint32 = 10
	var v0 uint32

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolUint16Err(t *testing.T) {
	// Test : someLogic

	var vo10 uint16 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []uint16{vo10, vo10}
	newList, _ := FilterMapBoolUint16Err(notOneBoolUint16Err, someLogicBoolUint16Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint16Err failed")
	}

	r, _ = FilterMapBoolUint16Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint16Err failed")
	}
	_, err := FilterMapBoolUint16Err(notOneBoolUint16Err, someLogicBoolUint16Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint16Err failed")
	}

	_, err = FilterMapBoolUint16Err(notOneBoolUint16Err2, someLogicBoolUint16Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint16Err failed")
	}
}
func notOneBoolUint16Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolUint16Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolUint16Err(num bool) (uint16, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 uint16 = 10
	var v0 uint16

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolUint8Err(t *testing.T) {
	// Test : someLogic

	var vo10 uint8 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []uint8{vo10, vo10}
	newList, _ := FilterMapBoolUint8Err(notOneBoolUint8Err, someLogicBoolUint8Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolUint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint8Err failed")
	}

	r, _ = FilterMapBoolUint8Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint8Err failed")
	}
	_, err := FilterMapBoolUint8Err(notOneBoolUint8Err, someLogicBoolUint8Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint8Err failed")
	}

	_, err = FilterMapBoolUint8Err(notOneBoolUint8Err2, someLogicBoolUint8Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint8Err failed")
	}
}
func notOneBoolUint8Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolUint8Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolUint8Err(num bool) (uint8, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 uint8 = 10
	var v0 uint8

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolStrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"

	var vit bool = true
	var vif bool = false

	expectedList := []string{vo10, vo10}
	newList, _ := FilterMapBoolStrErr(notOneBoolStrErr, someLogicBoolStrErr, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolStrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolStrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolStrErr failed")
	}

	r, _ = FilterMapBoolStrErr(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolStrErr failed")
	}
	_, err := FilterMapBoolStrErr(notOneBoolStrErr, someLogicBoolStrErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolStrErr failed")
	}
	_, err = FilterMapBoolStrErr(notOneBoolStrErr2, someLogicBoolStrErr, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolStrErr failed")
	}
}
func notOneBoolStrErr(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolStrErr2(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolStrErr(num bool) (string, error) {
	if num == false {
		return "", errors.New("false is error in this test")
	}
	var v10 string = "10"
	var v0 string = "0"

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolFloat32Err(t *testing.T) {
	// Test : someLogic

	var vo10 float32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []float32{vo10, vo10}
	newList, _ := FilterMapBoolFloat32Err(notOneBoolFloat32Err, someLogicBoolFloat32Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolFloat32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolFloat32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat32Err failed")
	}

	r, _ = FilterMapBoolFloat32Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat32Err failed")
	}
	_, err := FilterMapBoolFloat32Err(notOneBoolFloat32Err, someLogicBoolFloat32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat32Err failed")
	}

	_, err = FilterMapBoolFloat32Err(notOneBoolFloat32Err2, someLogicBoolFloat32Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat32Err failed")
	}
}
func notOneBoolFloat32Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolFloat32Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolFloat32Err(num bool) (float32, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 float32 = 10
	var v0 float32

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapBoolFloat64Err(t *testing.T) {
	// Test : someLogic

	var vo10 float64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []float64{vo10, vo10}
	newList, _ := FilterMapBoolFloat64Err(notOneBoolFloat64Err, someLogicBoolFloat64Err, []bool{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapBoolFloat64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolFloat64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat64Err failed")
	}

	r, _ = FilterMapBoolFloat64Err(nil, nil, []bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat64Err failed")
	}
	_, err := FilterMapBoolFloat64Err(notOneBoolFloat64Err, someLogicBoolFloat64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat64Err failed")
	}

	_, err = FilterMapBoolFloat64Err(notOneBoolFloat64Err2, someLogicBoolFloat64Err, []bool{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat64Err failed")
	}
}
func notOneBoolFloat64Err(num bool) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOneBoolFloat64Err2(num bool) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogicBoolFloat64Err(num bool) (float64, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 float64 = 10
	var v0 float64

	if num == true {
		return v10, nil
	}
	return v0, nil
}

func TestFilterMapFloat32IntErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapFloat32IntErr(notOneFloat32IntErr, plusOneFloat32IntErr, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32IntErr failed")
	}
	r, _ = FilterMapFloat32IntErr(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32IntPtrErr failed")
	}

	_, err := FilterMapFloat32IntErr(notOneFloat32IntErr, plusOneFloat32IntErr, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32IntErr failed")
	}
	_, err = FilterMapFloat32IntErr(notOneFloat32IntErr, plusOneFloat32IntErr, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32IntErr failed")
	}
}
func notOneFloat32IntErr(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32IntErr(num float32) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapFloat32Int64Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapFloat32Int64Err(notOneFloat32Int64Err, plusOneFloat32Int64Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int64Err failed")
	}
	r, _ = FilterMapFloat32Int64Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int64PtrErr failed")
	}

	_, err := FilterMapFloat32Int64Err(notOneFloat32Int64Err, plusOneFloat32Int64Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int64Err failed")
	}
	_, err = FilterMapFloat32Int64Err(notOneFloat32Int64Err, plusOneFloat32Int64Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int64Err failed")
	}
}
func notOneFloat32Int64Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Int64Err(num float32) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapFloat32Int32Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapFloat32Int32Err(notOneFloat32Int32Err, plusOneFloat32Int32Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int32Err failed")
	}
	r, _ = FilterMapFloat32Int32Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int32PtrErr failed")
	}

	_, err := FilterMapFloat32Int32Err(notOneFloat32Int32Err, plusOneFloat32Int32Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int32Err failed")
	}
	_, err = FilterMapFloat32Int32Err(notOneFloat32Int32Err, plusOneFloat32Int32Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int32Err failed")
	}
}
func notOneFloat32Int32Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Int32Err(num float32) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapFloat32Int16Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapFloat32Int16Err(notOneFloat32Int16Err, plusOneFloat32Int16Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int16Err failed")
	}
	r, _ = FilterMapFloat32Int16Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int16PtrErr failed")
	}

	_, err := FilterMapFloat32Int16Err(notOneFloat32Int16Err, plusOneFloat32Int16Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int16Err failed")
	}
	_, err = FilterMapFloat32Int16Err(notOneFloat32Int16Err, plusOneFloat32Int16Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int16Err failed")
	}
}
func notOneFloat32Int16Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Int16Err(num float32) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapFloat32Int8Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapFloat32Int8Err(notOneFloat32Int8Err, plusOneFloat32Int8Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int8Err failed")
	}
	r, _ = FilterMapFloat32Int8Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int8PtrErr failed")
	}

	_, err := FilterMapFloat32Int8Err(notOneFloat32Int8Err, plusOneFloat32Int8Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int8Err failed")
	}
	_, err = FilterMapFloat32Int8Err(notOneFloat32Int8Err, plusOneFloat32Int8Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int8Err failed")
	}
}
func notOneFloat32Int8Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Int8Err(num float32) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapFloat32UintErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapFloat32UintErr(notOneFloat32UintErr, plusOneFloat32UintErr, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32UintErr failed")
	}
	r, _ = FilterMapFloat32UintErr(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32UintPtrErr failed")
	}

	_, err := FilterMapFloat32UintErr(notOneFloat32UintErr, plusOneFloat32UintErr, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32UintErr failed")
	}
	_, err = FilterMapFloat32UintErr(notOneFloat32UintErr, plusOneFloat32UintErr, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32UintErr failed")
	}
}
func notOneFloat32UintErr(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32UintErr(num float32) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapFloat32Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapFloat32Uint64Err(notOneFloat32Uint64Err, plusOneFloat32Uint64Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint64Err failed")
	}
	r, _ = FilterMapFloat32Uint64Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed")
	}

	_, err := FilterMapFloat32Uint64Err(notOneFloat32Uint64Err, plusOneFloat32Uint64Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint64Err failed")
	}
	_, err = FilterMapFloat32Uint64Err(notOneFloat32Uint64Err, plusOneFloat32Uint64Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint64Err failed")
	}
}
func notOneFloat32Uint64Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Uint64Err(num float32) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapFloat32Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapFloat32Uint32Err(notOneFloat32Uint32Err, plusOneFloat32Uint32Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint32Err failed")
	}
	r, _ = FilterMapFloat32Uint32Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed")
	}

	_, err := FilterMapFloat32Uint32Err(notOneFloat32Uint32Err, plusOneFloat32Uint32Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint32Err failed")
	}
	_, err = FilterMapFloat32Uint32Err(notOneFloat32Uint32Err, plusOneFloat32Uint32Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint32Err failed")
	}
}
func notOneFloat32Uint32Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Uint32Err(num float32) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapFloat32Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapFloat32Uint16Err(notOneFloat32Uint16Err, plusOneFloat32Uint16Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint16Err failed")
	}
	r, _ = FilterMapFloat32Uint16Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed")
	}

	_, err := FilterMapFloat32Uint16Err(notOneFloat32Uint16Err, plusOneFloat32Uint16Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint16Err failed")
	}
	_, err = FilterMapFloat32Uint16Err(notOneFloat32Uint16Err, plusOneFloat32Uint16Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint16Err failed")
	}
}
func notOneFloat32Uint16Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Uint16Err(num float32) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapFloat32Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapFloat32Uint8Err(notOneFloat32Uint8Err, plusOneFloat32Uint8Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint8Err failed")
	}
	r, _ = FilterMapFloat32Uint8Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed")
	}

	_, err := FilterMapFloat32Uint8Err(notOneFloat32Uint8Err, plusOneFloat32Uint8Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint8Err failed")
	}
	_, err = FilterMapFloat32Uint8Err(notOneFloat32Uint8Err, plusOneFloat32Uint8Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint8Err failed")
	}
}
func notOneFloat32Uint8Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Uint8Err(num float32) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapFloat32StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float32 = 1
	var iv2 float32 = 2
	var iv3 float32 = 3
	var iv10 float32 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapFloat32StrErr(notOneFloat32StrNumErr, someLogicFloat32StrNumErr, []float32{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapFloat32StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32StrErr failed")
	}

	r, _ = FilterMapFloat32StrErr(nil, nil, []float32{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32StrErr failed")
	}

	_, err := FilterMapFloat32StrErr(notOneFloat32StrNumErr, someLogicFloat32StrNumErr, []float32{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapFloat32StrErr failed")
	}
	_, err = FilterMapFloat32StrErr(notOneFloat32StrNumErr, someLogicFloat32StrNumErr, []float32{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapFloat32StrErr failed")
	}
}

func notOneFloat32StrNumErr(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicFloat32StrNumErr(num float32) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapFloat32BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	var vi10 float32 = 10
	var vi0 float32

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapFloat32BoolErr(notOneFloat32BoolErr, someLogicFloat32BoolErr, []float32{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32BoolErr failed")
	}

	r, _ = FilterMapFloat32BoolErr(nil, nil, []float32{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32BoolPtrErr failed")
	}

	_, err := FilterMapFloat32BoolErr(notOneFloat32BoolErr, someLogicFloat32BoolErr, []float32{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapFloat32BoolErr failed")
	}

	_, err = FilterMapFloat32BoolErr(notOneFloat32BoolErr, someLogicFloat32BoolErr, []float32{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapFloat32BoolErr failed")
	}
}
func notOneFloat32BoolErr(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicFloat32BoolErr(num float32) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapFloat32Float64Err(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []float64{vo5, vo6}
	newList, _ := FilterMapFloat32Float64Err(notOneFloat32Float64Err, plusOneFloat32Float64Err, []float32{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat32Float64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Float64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Float64Err failed")
	}
	r, _ = FilterMapFloat32Float64Err(nil, nil, []float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Float64PtrErr failed")
	}

	_, err := FilterMapFloat32Float64Err(notOneFloat32Float64Err, plusOneFloat32Float64Err, []float32{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Float64Err failed")
	}
	_, err = FilterMapFloat32Float64Err(notOneFloat32Float64Err, plusOneFloat32Float64Err, []float32{v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Float64Err failed")
	}
}
func notOneFloat32Float64Err(num float32) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat32Float64Err(num float32) (float64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float64(num + 1)
	return c, nil
}

func TestFilterMapFloat64IntErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []int{vo5, vo6}
	newList, _ := FilterMapFloat64IntErr(notOneFloat64IntErr, plusOneFloat64IntErr, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64IntErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64IntErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64IntErr failed")
	}
	r, _ = FilterMapFloat64IntErr(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64IntPtrErr failed")
	}

	_, err := FilterMapFloat64IntErr(notOneFloat64IntErr, plusOneFloat64IntErr, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64IntErr failed")
	}
	_, err = FilterMapFloat64IntErr(notOneFloat64IntErr, plusOneFloat64IntErr, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64IntErr failed")
	}
}
func notOneFloat64IntErr(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64IntErr(num float64) (int, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int(num + 1)
	return c, nil
}

func TestFilterMapFloat64Int64Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []int64{vo5, vo6}
	newList, _ := FilterMapFloat64Int64Err(notOneFloat64Int64Err, plusOneFloat64Int64Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int64Err failed")
	}
	r, _ = FilterMapFloat64Int64Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int64PtrErr failed")
	}

	_, err := FilterMapFloat64Int64Err(notOneFloat64Int64Err, plusOneFloat64Int64Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int64Err failed")
	}
	_, err = FilterMapFloat64Int64Err(notOneFloat64Int64Err, plusOneFloat64Int64Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int64Err failed")
	}
}
func notOneFloat64Int64Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Int64Err(num float64) (int64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int64(num + 1)
	return c, nil
}

func TestFilterMapFloat64Int32Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []int32{vo5, vo6}
	newList, _ := FilterMapFloat64Int32Err(notOneFloat64Int32Err, plusOneFloat64Int32Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int32Err failed")
	}
	r, _ = FilterMapFloat64Int32Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int32PtrErr failed")
	}

	_, err := FilterMapFloat64Int32Err(notOneFloat64Int32Err, plusOneFloat64Int32Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int32Err failed")
	}
	_, err = FilterMapFloat64Int32Err(notOneFloat64Int32Err, plusOneFloat64Int32Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int32Err failed")
	}
}
func notOneFloat64Int32Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Int32Err(num float64) (int32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int32(num + 1)
	return c, nil
}

func TestFilterMapFloat64Int16Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []int16{vo5, vo6}
	newList, _ := FilterMapFloat64Int16Err(notOneFloat64Int16Err, plusOneFloat64Int16Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int16Err failed")
	}
	r, _ = FilterMapFloat64Int16Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int16PtrErr failed")
	}

	_, err := FilterMapFloat64Int16Err(notOneFloat64Int16Err, plusOneFloat64Int16Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int16Err failed")
	}
	_, err = FilterMapFloat64Int16Err(notOneFloat64Int16Err, plusOneFloat64Int16Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int16Err failed")
	}
}
func notOneFloat64Int16Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Int16Err(num float64) (int16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int16(num + 1)
	return c, nil
}

func TestFilterMapFloat64Int8Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []int8{vo5, vo6}
	newList, _ := FilterMapFloat64Int8Err(notOneFloat64Int8Err, plusOneFloat64Int8Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Int8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int8Err failed")
	}
	r, _ = FilterMapFloat64Int8Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int8PtrErr failed")
	}

	_, err := FilterMapFloat64Int8Err(notOneFloat64Int8Err, plusOneFloat64Int8Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int8Err failed")
	}
	_, err = FilterMapFloat64Int8Err(notOneFloat64Int8Err, plusOneFloat64Int8Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int8Err failed")
	}
}
func notOneFloat64Int8Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Int8Err(num float64) (int8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := int8(num + 1)
	return c, nil
}

func TestFilterMapFloat64UintErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []uint{vo5, vo6}
	newList, _ := FilterMapFloat64UintErr(notOneFloat64UintErr, plusOneFloat64UintErr, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64UintErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64UintErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64UintErr failed")
	}
	r, _ = FilterMapFloat64UintErr(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64UintPtrErr failed")
	}

	_, err := FilterMapFloat64UintErr(notOneFloat64UintErr, plusOneFloat64UintErr, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64UintErr failed")
	}
	_, err = FilterMapFloat64UintErr(notOneFloat64UintErr, plusOneFloat64UintErr, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64UintErr failed")
	}
}
func notOneFloat64UintErr(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64UintErr(num float64) (uint, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint(num + 1)
	return c, nil
}

func TestFilterMapFloat64Uint64Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []uint64{vo5, vo6}
	newList, _ := FilterMapFloat64Uint64Err(notOneFloat64Uint64Err, plusOneFloat64Uint64Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint64Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint64Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint64Err failed")
	}
	r, _ = FilterMapFloat64Uint64Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed")
	}

	_, err := FilterMapFloat64Uint64Err(notOneFloat64Uint64Err, plusOneFloat64Uint64Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint64Err failed")
	}
	_, err = FilterMapFloat64Uint64Err(notOneFloat64Uint64Err, plusOneFloat64Uint64Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint64Err failed")
	}
}
func notOneFloat64Uint64Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Uint64Err(num float64) (uint64, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint64(num + 1)
	return c, nil
}

func TestFilterMapFloat64Uint32Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []uint32{vo5, vo6}
	newList, _ := FilterMapFloat64Uint32Err(notOneFloat64Uint32Err, plusOneFloat64Uint32Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint32Err failed")
	}
	r, _ = FilterMapFloat64Uint32Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed")
	}

	_, err := FilterMapFloat64Uint32Err(notOneFloat64Uint32Err, plusOneFloat64Uint32Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint32Err failed")
	}
	_, err = FilterMapFloat64Uint32Err(notOneFloat64Uint32Err, plusOneFloat64Uint32Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint32Err failed")
	}
}
func notOneFloat64Uint32Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Uint32Err(num float64) (uint32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint32(num + 1)
	return c, nil
}

func TestFilterMapFloat64Uint16Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []uint16{vo5, vo6}
	newList, _ := FilterMapFloat64Uint16Err(notOneFloat64Uint16Err, plusOneFloat64Uint16Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint16Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint16Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint16Err failed")
	}
	r, _ = FilterMapFloat64Uint16Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed")
	}

	_, err := FilterMapFloat64Uint16Err(notOneFloat64Uint16Err, plusOneFloat64Uint16Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint16Err failed")
	}
	_, err = FilterMapFloat64Uint16Err(notOneFloat64Uint16Err, plusOneFloat64Uint16Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint16Err failed")
	}
}
func notOneFloat64Uint16Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Uint16Err(num float64) (uint16, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint16(num + 1)
	return c, nil
}

func TestFilterMapFloat64Uint8Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []uint8{vo5, vo6}
	newList, _ := FilterMapFloat64Uint8Err(notOneFloat64Uint8Err, plusOneFloat64Uint8Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Uint8Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint8Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint8Err failed")
	}
	r, _ = FilterMapFloat64Uint8Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed")
	}

	_, err := FilterMapFloat64Uint8Err(notOneFloat64Uint8Err, plusOneFloat64Uint8Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint8Err failed")
	}
	_, err = FilterMapFloat64Uint8Err(notOneFloat64Uint8Err, plusOneFloat64Uint8Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint8Err failed")
	}
}
func notOneFloat64Uint8Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Uint8Err(num float64) (uint8, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := uint8(num + 1)
	return c, nil
}

func TestFilterMapFloat64StrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float64 = 1
	var iv2 float64 = 2
	var iv3 float64 = 3
	var iv10 float64 = 10
	expectedList := []string{ov10}
	newList, _ := FilterMapFloat64StrErr(notOneFloat64StrNumErr, someLogicFloat64StrNumErr, []float64{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMapFloat64StrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64StrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64StrErr failed")
	}

	r, _ = FilterMapFloat64StrErr(nil, nil, []float64{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64StrErr failed")
	}

	_, err := FilterMapFloat64StrErr(notOneFloat64StrNumErr, someLogicFloat64StrNumErr, []float64{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMapFloat64StrErr failed")
	}
	_, err = FilterMapFloat64StrErr(notOneFloat64StrNumErr, someLogicFloat64StrNumErr, []float64{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMapFloat64StrErr failed")
	}
}

func notOneFloat64StrNumErr(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicFloat64StrNumErr(num float64) (string, error) {
	var r string = "0"
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}

func TestFilterMapFloat64BoolErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	var vi10 float64 = 10
	var vi0 float64

	expectedList := []bool{vto, vfo}
	newList, _ := FilterMapFloat64BoolErr(notOneFloat64BoolErr, someLogicFloat64BoolErr, []float64{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64BoolErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64BoolErr failed")
	}

	r, _ = FilterMapFloat64BoolErr(nil, nil, []float64{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64BoolPtrErr failed")
	}

	_, err := FilterMapFloat64BoolErr(notOneFloat64BoolErr, someLogicFloat64BoolErr, []float64{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapFloat64BoolErr failed")
	}

	_, err = FilterMapFloat64BoolErr(notOneFloat64BoolErr, someLogicFloat64BoolErr, []float64{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMapFloat64BoolErr failed")
	}
}
func notOneFloat64BoolErr(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogicFloat64BoolErr(num float64) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}

func TestFilterMapFloat64Float32Err(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []float32{vo5, vo6}
	newList, _ := FilterMapFloat64Float32Err(notOneFloat64Float32Err, plusOneFloat64Float32Err, []float64{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMapFloat64Float32Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Float32Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Float32Err failed")
	}
	r, _ = FilterMapFloat64Float32Err(nil, nil, []float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Float32PtrErr failed")
	}

	_, err := FilterMapFloat64Float32Err(notOneFloat64Float32Err, plusOneFloat64Float32Err, []float64{v2, v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Float32Err failed")
	}
	_, err = FilterMapFloat64Float32Err(notOneFloat64Float32Err, plusOneFloat64Float32Err, []float64{v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Float32Err failed")
	}
}
func notOneFloat64Float32Err(num float64) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOneFloat64Float32Err(num float64) (float32, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := float32(num + 1)
	return c, nil
}
