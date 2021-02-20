package fp

import (
	"errors"
	"testing"
)

func TestFilterMapIntInt64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapIntInt64PtrErr(notOneIntInt64PtrErr, plusOneIntInt64PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt64PtrErr failed")
	}
	r, _ = FilterMapIntInt64PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt64PtrErr failed")
	}

	_, err := FilterMapIntInt64PtrErr(notOneIntInt64PtrErr, plusOneIntInt64PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntInt64PtrErr failed")
	}
	_, err = FilterMapIntInt64PtrErr(notOneIntInt64PtrErr, plusOneIntInt64PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntInt64PtrErr failed")
	}
}
func notOneIntInt64PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntInt64PtrErr(num *int) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapIntInt32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapIntInt32PtrErr(notOneIntInt32PtrErr, plusOneIntInt32PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt32PtrErr failed")
	}
	r, _ = FilterMapIntInt32PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt32PtrErr failed")
	}

	_, err := FilterMapIntInt32PtrErr(notOneIntInt32PtrErr, plusOneIntInt32PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntInt32PtrErr failed")
	}
	_, err = FilterMapIntInt32PtrErr(notOneIntInt32PtrErr, plusOneIntInt32PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntInt32PtrErr failed")
	}
}
func notOneIntInt32PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntInt32PtrErr(num *int) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapIntInt16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapIntInt16PtrErr(notOneIntInt16PtrErr, plusOneIntInt16PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt16PtrErr failed")
	}
	r, _ = FilterMapIntInt16PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt16PtrErr failed")
	}

	_, err := FilterMapIntInt16PtrErr(notOneIntInt16PtrErr, plusOneIntInt16PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntInt16PtrErr failed")
	}
	_, err = FilterMapIntInt16PtrErr(notOneIntInt16PtrErr, plusOneIntInt16PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntInt16PtrErr failed")
	}
}
func notOneIntInt16PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntInt16PtrErr(num *int) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapIntInt8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapIntInt8PtrErr(notOneIntInt8PtrErr, plusOneIntInt8PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntInt8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntInt8PtrErr failed")
	}
	r, _ = FilterMapIntInt8PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntInt8PtrErr failed")
	}

	_, err := FilterMapIntInt8PtrErr(notOneIntInt8PtrErr, plusOneIntInt8PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntInt8PtrErr failed")
	}
	_, err = FilterMapIntInt8PtrErr(notOneIntInt8PtrErr, plusOneIntInt8PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntInt8PtrErr failed")
	}
}
func notOneIntInt8PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntInt8PtrErr(num *int) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapIntUintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapIntUintPtrErr(notOneIntUintPtrErr, plusOneIntUintPtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUintPtrErr failed")
	}
	r, _ = FilterMapIntUintPtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUintPtrErr failed")
	}

	_, err := FilterMapIntUintPtrErr(notOneIntUintPtrErr, plusOneIntUintPtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntUintPtrErr failed")
	}
	_, err = FilterMapIntUintPtrErr(notOneIntUintPtrErr, plusOneIntUintPtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntUintPtrErr failed")
	}
}
func notOneIntUintPtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntUintPtrErr(num *int) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapIntUint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapIntUint64PtrErr(notOneIntUint64PtrErr, plusOneIntUint64PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint64PtrErr failed")
	}
	r, _ = FilterMapIntUint64PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint64PtrErr failed")
	}

	_, err := FilterMapIntUint64PtrErr(notOneIntUint64PtrErr, plusOneIntUint64PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntUint64PtrErr failed")
	}
	_, err = FilterMapIntUint64PtrErr(notOneIntUint64PtrErr, plusOneIntUint64PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntUint64PtrErr failed")
	}
}
func notOneIntUint64PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntUint64PtrErr(num *int) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapIntUint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapIntUint32PtrErr(notOneIntUint32PtrErr, plusOneIntUint32PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint32PtrErr failed")
	}
	r, _ = FilterMapIntUint32PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint32PtrErr failed")
	}

	_, err := FilterMapIntUint32PtrErr(notOneIntUint32PtrErr, plusOneIntUint32PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntUint32PtrErr failed")
	}
	_, err = FilterMapIntUint32PtrErr(notOneIntUint32PtrErr, plusOneIntUint32PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntUint32PtrErr failed")
	}
}
func notOneIntUint32PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntUint32PtrErr(num *int) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapIntUint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapIntUint16PtrErr(notOneIntUint16PtrErr, plusOneIntUint16PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint16PtrErr failed")
	}
	r, _ = FilterMapIntUint16PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint16PtrErr failed")
	}

	_, err := FilterMapIntUint16PtrErr(notOneIntUint16PtrErr, plusOneIntUint16PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntUint16PtrErr failed")
	}
	_, err = FilterMapIntUint16PtrErr(notOneIntUint16PtrErr, plusOneIntUint16PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntUint16PtrErr failed")
	}
}
func notOneIntUint16PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntUint16PtrErr(num *int) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapIntUint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapIntUint8PtrErr(notOneIntUint8PtrErr, plusOneIntUint8PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntUint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntUint8PtrErr failed")
	}
	r, _ = FilterMapIntUint8PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntUint8PtrErr failed")
	}

	_, err := FilterMapIntUint8PtrErr(notOneIntUint8PtrErr, plusOneIntUint8PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntUint8PtrErr failed")
	}
	_, err = FilterMapIntUint8PtrErr(notOneIntUint8PtrErr, plusOneIntUint8PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntUint8PtrErr failed")
	}
}
func notOneIntUint8PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntUint8PtrErr(num *int) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapIntStrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int = 1
	var iv2 int = 2
	var iv3 int = 3
	var iv10 int = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapIntStrPtrErr(notOneIntStrNumPtrErr, someLogicIntStrNumPtrErr, []*int{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapIntStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntStrPtrErr failed")
	}

	r, _ = FilterMapIntStrPtrErr(nil, nil, []*int{})
	if len(r) > 0 {
		t.Errorf("FilterMapIntStrPtrErr failed")
	}

	_, err := FilterMapIntStrPtrErr(notOneIntStrNumPtrErr, someLogicIntStrNumPtrErr, []*int{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapIntStrPtrErr failed")
	}
	_, err = FilterMapIntStrPtrErr(notOneIntStrNumPtrErr, someLogicIntStrNumPtrErr, []*int{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapIntStrPtrErr failed")
	}
}

func notOneIntStrNumPtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicIntStrNumPtrErr(num *int) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapIntBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	var vi10 int = 10
	var vi0 int = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapIntBoolPtrErr(notOneIntBoolPtrErr, someLogicIntBoolPtrErr, []*int{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntBoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntBoolPtr failed")
	}

	r, _ = FilterMapIntBoolPtrErr(nil, nil, []*int{})
	if len(r) > 0 {
		t.Errorf("FilterMapIntBoolPtrErr failed")
	}

	_, err := FilterMapIntBoolPtrErr(notOneIntBoolPtrErr, someLogicIntBoolPtrErr, []*int{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapIntBoolPtrErr failed")
	}

	_, err = FilterMapIntBoolPtrErr(notOneIntBoolPtrErr, someLogicIntBoolPtrErr, []*int{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapIntBoolPtrErr failed")
	}
}
func notOneIntBoolPtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicIntBoolPtrErr(num *int) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapIntFloat32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapIntFloat32PtrErr(notOneIntFloat32PtrErr, plusOneIntFloat32PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntFloat32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat32PtrErr failed")
	}
	r, _ = FilterMapIntFloat32PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat32PtrErr failed")
	}

	_, err := FilterMapIntFloat32PtrErr(notOneIntFloat32PtrErr, plusOneIntFloat32PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat32PtrErr failed")
	}
	_, err = FilterMapIntFloat32PtrErr(notOneIntFloat32PtrErr, plusOneIntFloat32PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat32PtrErr failed")
	}
}
func notOneIntFloat32PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntFloat32PtrErr(num *int) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapIntFloat64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapIntFloat64PtrErr(notOneIntFloat64PtrErr, plusOneIntFloat64PtrErr, []*int{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapIntFloat64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat64PtrErr failed")
	}
	r, _ = FilterMapIntFloat64PtrErr(nil, nil, []*int{})

	if len(r) > 0 {
		t.Errorf("FilterMapIntFloat64PtrErr failed")
	}

	_, err := FilterMapIntFloat64PtrErr(notOneIntFloat64PtrErr, plusOneIntFloat64PtrErr, []*int{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat64PtrErr failed")
	}
	_, err = FilterMapIntFloat64PtrErr(notOneIntFloat64PtrErr, plusOneIntFloat64PtrErr, []*int{&v3})
	if err == nil {
		t.Errorf("FilterMapIntFloat64PtrErr failed")
	}
}
func notOneIntFloat64PtrErr(num *int) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneIntFloat64PtrErr(num *int) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapInt64IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapInt64IntPtrErr(notOneInt64IntPtrErr, plusOneInt64IntPtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64IntPtrErr failed")
	}
	r, _ = FilterMapInt64IntPtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64IntPtrErr failed")
	}

	_, err := FilterMapInt64IntPtrErr(notOneInt64IntPtrErr, plusOneInt64IntPtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64IntPtrErr failed")
	}
	_, err = FilterMapInt64IntPtrErr(notOneInt64IntPtrErr, plusOneInt64IntPtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64IntPtrErr failed")
	}
}
func notOneInt64IntPtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64IntPtrErr(num *int64) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapInt64Int32PtrErr(notOneInt64Int32PtrErr, plusOneInt64Int32PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int32PtrErr failed")
	}
	r, _ = FilterMapInt64Int32PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int32PtrErr failed")
	}

	_, err := FilterMapInt64Int32PtrErr(notOneInt64Int32PtrErr, plusOneInt64Int32PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int32PtrErr failed")
	}
	_, err = FilterMapInt64Int32PtrErr(notOneInt64Int32PtrErr, plusOneInt64Int32PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int32PtrErr failed")
	}
}
func notOneInt64Int32PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Int32PtrErr(num *int64) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapInt64Int16PtrErr(notOneInt64Int16PtrErr, plusOneInt64Int16PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int16PtrErr failed")
	}
	r, _ = FilterMapInt64Int16PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int16PtrErr failed")
	}

	_, err := FilterMapInt64Int16PtrErr(notOneInt64Int16PtrErr, plusOneInt64Int16PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int16PtrErr failed")
	}
	_, err = FilterMapInt64Int16PtrErr(notOneInt64Int16PtrErr, plusOneInt64Int16PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int16PtrErr failed")
	}
}
func notOneInt64Int16PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Int16PtrErr(num *int64) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapInt64Int8PtrErr(notOneInt64Int8PtrErr, plusOneInt64Int8PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int8PtrErr failed")
	}
	r, _ = FilterMapInt64Int8PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Int8PtrErr failed")
	}

	_, err := FilterMapInt64Int8PtrErr(notOneInt64Int8PtrErr, plusOneInt64Int8PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int8PtrErr failed")
	}
	_, err = FilterMapInt64Int8PtrErr(notOneInt64Int8PtrErr, plusOneInt64Int8PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Int8PtrErr failed")
	}
}
func notOneInt64Int8PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Int8PtrErr(num *int64) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapInt64UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapInt64UintPtrErr(notOneInt64UintPtrErr, plusOneInt64UintPtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64UintPtrErr failed")
	}
	r, _ = FilterMapInt64UintPtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64UintPtrErr failed")
	}

	_, err := FilterMapInt64UintPtrErr(notOneInt64UintPtrErr, plusOneInt64UintPtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64UintPtrErr failed")
	}
	_, err = FilterMapInt64UintPtrErr(notOneInt64UintPtrErr, plusOneInt64UintPtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64UintPtrErr failed")
	}
}
func notOneInt64UintPtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64UintPtrErr(num *int64) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapInt64Uint64PtrErr(notOneInt64Uint64PtrErr, plusOneInt64Uint64PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint64PtrErr failed")
	}
	r, _ = FilterMapInt64Uint64PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint64PtrErr failed")
	}

	_, err := FilterMapInt64Uint64PtrErr(notOneInt64Uint64PtrErr, plusOneInt64Uint64PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint64PtrErr failed")
	}
	_, err = FilterMapInt64Uint64PtrErr(notOneInt64Uint64PtrErr, plusOneInt64Uint64PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint64PtrErr failed")
	}
}
func notOneInt64Uint64PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Uint64PtrErr(num *int64) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapInt64Uint32PtrErr(notOneInt64Uint32PtrErr, plusOneInt64Uint32PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint32PtrErr failed")
	}
	r, _ = FilterMapInt64Uint32PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint32PtrErr failed")
	}

	_, err := FilterMapInt64Uint32PtrErr(notOneInt64Uint32PtrErr, plusOneInt64Uint32PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint32PtrErr failed")
	}
	_, err = FilterMapInt64Uint32PtrErr(notOneInt64Uint32PtrErr, plusOneInt64Uint32PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint32PtrErr failed")
	}
}
func notOneInt64Uint32PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Uint32PtrErr(num *int64) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapInt64Uint16PtrErr(notOneInt64Uint16PtrErr, plusOneInt64Uint16PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint16PtrErr failed")
	}
	r, _ = FilterMapInt64Uint16PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint16PtrErr failed")
	}

	_, err := FilterMapInt64Uint16PtrErr(notOneInt64Uint16PtrErr, plusOneInt64Uint16PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint16PtrErr failed")
	}
	_, err = FilterMapInt64Uint16PtrErr(notOneInt64Uint16PtrErr, plusOneInt64Uint16PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint16PtrErr failed")
	}
}
func notOneInt64Uint16PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Uint16PtrErr(num *int64) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapInt64Uint8PtrErr(notOneInt64Uint8PtrErr, plusOneInt64Uint8PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint8PtrErr failed")
	}
	r, _ = FilterMapInt64Uint8PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Uint8PtrErr failed")
	}

	_, err := FilterMapInt64Uint8PtrErr(notOneInt64Uint8PtrErr, plusOneInt64Uint8PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint8PtrErr failed")
	}
	_, err = FilterMapInt64Uint8PtrErr(notOneInt64Uint8PtrErr, plusOneInt64Uint8PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Uint8PtrErr failed")
	}
}
func notOneInt64Uint8PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Uint8PtrErr(num *int64) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapInt64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int64 = 1
	var iv2 int64 = 2
	var iv3 int64 = 3
	var iv10 int64 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapInt64StrPtrErr(notOneInt64StrNumPtrErr, someLogicInt64StrNumPtrErr, []*int64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64StrPtrErr failed")
	}

	r, _ = FilterMapInt64StrPtrErr(nil, nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt64StrPtrErr failed")
	}

	_, err := FilterMapInt64StrPtrErr(notOneInt64StrNumPtrErr, someLogicInt64StrNumPtrErr, []*int64{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapInt64StrPtrErr failed")
	}
	_, err = FilterMapInt64StrPtrErr(notOneInt64StrNumPtrErr, someLogicInt64StrNumPtrErr, []*int64{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapInt64StrPtrErr failed")
	}
}

func notOneInt64StrNumPtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicInt64StrNumPtrErr(num *int64) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapInt64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	var vi10 int64 = 10
	var vi0 int64 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapInt64BoolPtrErr(notOneInt64BoolPtrErr, someLogicInt64BoolPtrErr, []*int64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64BoolPtr failed")
	}

	r, _ = FilterMapInt64BoolPtrErr(nil, nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt64BoolPtrErr failed")
	}

	_, err := FilterMapInt64BoolPtrErr(notOneInt64BoolPtrErr, someLogicInt64BoolPtrErr, []*int64{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt64BoolPtrErr failed")
	}

	_, err = FilterMapInt64BoolPtrErr(notOneInt64BoolPtrErr, someLogicInt64BoolPtrErr, []*int64{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt64BoolPtrErr failed")
	}
}
func notOneInt64BoolPtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicInt64BoolPtrErr(num *int64) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapInt64Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapInt64Float32PtrErr(notOneInt64Float32PtrErr, plusOneInt64Float32PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float32PtrErr failed")
	}
	r, _ = FilterMapInt64Float32PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float32PtrErr failed")
	}

	_, err := FilterMapInt64Float32PtrErr(notOneInt64Float32PtrErr, plusOneInt64Float32PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float32PtrErr failed")
	}
	_, err = FilterMapInt64Float32PtrErr(notOneInt64Float32PtrErr, plusOneInt64Float32PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float32PtrErr failed")
	}
}
func notOneInt64Float32PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Float32PtrErr(num *int64) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapInt64Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapInt64Float64PtrErr(notOneInt64Float64PtrErr, plusOneInt64Float64PtrErr, []*int64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt64Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float64PtrErr failed")
	}
	r, _ = FilterMapInt64Float64PtrErr(nil, nil, []*int64{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt64Float64PtrErr failed")
	}

	_, err := FilterMapInt64Float64PtrErr(notOneInt64Float64PtrErr, plusOneInt64Float64PtrErr, []*int64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float64PtrErr failed")
	}
	_, err = FilterMapInt64Float64PtrErr(notOneInt64Float64PtrErr, plusOneInt64Float64PtrErr, []*int64{&v3})
	if err == nil {
		t.Errorf("FilterMapInt64Float64PtrErr failed")
	}
}
func notOneInt64Float64PtrErr(num *int64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt64Float64PtrErr(num *int64) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapInt32IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapInt32IntPtrErr(notOneInt32IntPtrErr, plusOneInt32IntPtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32IntPtrErr failed")
	}
	r, _ = FilterMapInt32IntPtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32IntPtrErr failed")
	}

	_, err := FilterMapInt32IntPtrErr(notOneInt32IntPtrErr, plusOneInt32IntPtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32IntPtrErr failed")
	}
	_, err = FilterMapInt32IntPtrErr(notOneInt32IntPtrErr, plusOneInt32IntPtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32IntPtrErr failed")
	}
}
func notOneInt32IntPtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32IntPtrErr(num *int32) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapInt32Int64PtrErr(notOneInt32Int64PtrErr, plusOneInt32Int64PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int64PtrErr failed")
	}
	r, _ = FilterMapInt32Int64PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int64PtrErr failed")
	}

	_, err := FilterMapInt32Int64PtrErr(notOneInt32Int64PtrErr, plusOneInt32Int64PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int64PtrErr failed")
	}
	_, err = FilterMapInt32Int64PtrErr(notOneInt32Int64PtrErr, plusOneInt32Int64PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int64PtrErr failed")
	}
}
func notOneInt32Int64PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Int64PtrErr(num *int32) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapInt32Int16PtrErr(notOneInt32Int16PtrErr, plusOneInt32Int16PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int16PtrErr failed")
	}
	r, _ = FilterMapInt32Int16PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int16PtrErr failed")
	}

	_, err := FilterMapInt32Int16PtrErr(notOneInt32Int16PtrErr, plusOneInt32Int16PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int16PtrErr failed")
	}
	_, err = FilterMapInt32Int16PtrErr(notOneInt32Int16PtrErr, plusOneInt32Int16PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int16PtrErr failed")
	}
}
func notOneInt32Int16PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Int16PtrErr(num *int32) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapInt32Int8PtrErr(notOneInt32Int8PtrErr, plusOneInt32Int8PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int8PtrErr failed")
	}
	r, _ = FilterMapInt32Int8PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Int8PtrErr failed")
	}

	_, err := FilterMapInt32Int8PtrErr(notOneInt32Int8PtrErr, plusOneInt32Int8PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int8PtrErr failed")
	}
	_, err = FilterMapInt32Int8PtrErr(notOneInt32Int8PtrErr, plusOneInt32Int8PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Int8PtrErr failed")
	}
}
func notOneInt32Int8PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Int8PtrErr(num *int32) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapInt32UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapInt32UintPtrErr(notOneInt32UintPtrErr, plusOneInt32UintPtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32UintPtrErr failed")
	}
	r, _ = FilterMapInt32UintPtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32UintPtrErr failed")
	}

	_, err := FilterMapInt32UintPtrErr(notOneInt32UintPtrErr, plusOneInt32UintPtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32UintPtrErr failed")
	}
	_, err = FilterMapInt32UintPtrErr(notOneInt32UintPtrErr, plusOneInt32UintPtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32UintPtrErr failed")
	}
}
func notOneInt32UintPtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32UintPtrErr(num *int32) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapInt32Uint64PtrErr(notOneInt32Uint64PtrErr, plusOneInt32Uint64PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint64PtrErr failed")
	}
	r, _ = FilterMapInt32Uint64PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint64PtrErr failed")
	}

	_, err := FilterMapInt32Uint64PtrErr(notOneInt32Uint64PtrErr, plusOneInt32Uint64PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint64PtrErr failed")
	}
	_, err = FilterMapInt32Uint64PtrErr(notOneInt32Uint64PtrErr, plusOneInt32Uint64PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint64PtrErr failed")
	}
}
func notOneInt32Uint64PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Uint64PtrErr(num *int32) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapInt32Uint32PtrErr(notOneInt32Uint32PtrErr, plusOneInt32Uint32PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint32PtrErr failed")
	}
	r, _ = FilterMapInt32Uint32PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint32PtrErr failed")
	}

	_, err := FilterMapInt32Uint32PtrErr(notOneInt32Uint32PtrErr, plusOneInt32Uint32PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint32PtrErr failed")
	}
	_, err = FilterMapInt32Uint32PtrErr(notOneInt32Uint32PtrErr, plusOneInt32Uint32PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint32PtrErr failed")
	}
}
func notOneInt32Uint32PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Uint32PtrErr(num *int32) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapInt32Uint16PtrErr(notOneInt32Uint16PtrErr, plusOneInt32Uint16PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint16PtrErr failed")
	}
	r, _ = FilterMapInt32Uint16PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint16PtrErr failed")
	}

	_, err := FilterMapInt32Uint16PtrErr(notOneInt32Uint16PtrErr, plusOneInt32Uint16PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint16PtrErr failed")
	}
	_, err = FilterMapInt32Uint16PtrErr(notOneInt32Uint16PtrErr, plusOneInt32Uint16PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint16PtrErr failed")
	}
}
func notOneInt32Uint16PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Uint16PtrErr(num *int32) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapInt32Uint8PtrErr(notOneInt32Uint8PtrErr, plusOneInt32Uint8PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint8PtrErr failed")
	}
	r, _ = FilterMapInt32Uint8PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Uint8PtrErr failed")
	}

	_, err := FilterMapInt32Uint8PtrErr(notOneInt32Uint8PtrErr, plusOneInt32Uint8PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint8PtrErr failed")
	}
	_, err = FilterMapInt32Uint8PtrErr(notOneInt32Uint8PtrErr, plusOneInt32Uint8PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Uint8PtrErr failed")
	}
}
func notOneInt32Uint8PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Uint8PtrErr(num *int32) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapInt32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int32 = 1
	var iv2 int32 = 2
	var iv3 int32 = 3
	var iv10 int32 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapInt32StrPtrErr(notOneInt32StrNumPtrErr, someLogicInt32StrNumPtrErr, []*int32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32StrPtrErr failed")
	}

	r, _ = FilterMapInt32StrPtrErr(nil, nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt32StrPtrErr failed")
	}

	_, err := FilterMapInt32StrPtrErr(notOneInt32StrNumPtrErr, someLogicInt32StrNumPtrErr, []*int32{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapInt32StrPtrErr failed")
	}
	_, err = FilterMapInt32StrPtrErr(notOneInt32StrNumPtrErr, someLogicInt32StrNumPtrErr, []*int32{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapInt32StrPtrErr failed")
	}
}

func notOneInt32StrNumPtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicInt32StrNumPtrErr(num *int32) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapInt32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	var vi10 int32 = 10
	var vi0 int32 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapInt32BoolPtrErr(notOneInt32BoolPtrErr, someLogicInt32BoolPtrErr, []*int32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32BoolPtr failed")
	}

	r, _ = FilterMapInt32BoolPtrErr(nil, nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt32BoolPtrErr failed")
	}

	_, err := FilterMapInt32BoolPtrErr(notOneInt32BoolPtrErr, someLogicInt32BoolPtrErr, []*int32{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt32BoolPtrErr failed")
	}

	_, err = FilterMapInt32BoolPtrErr(notOneInt32BoolPtrErr, someLogicInt32BoolPtrErr, []*int32{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt32BoolPtrErr failed")
	}
}
func notOneInt32BoolPtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicInt32BoolPtrErr(num *int32) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapInt32Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapInt32Float32PtrErr(notOneInt32Float32PtrErr, plusOneInt32Float32PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float32PtrErr failed")
	}
	r, _ = FilterMapInt32Float32PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float32PtrErr failed")
	}

	_, err := FilterMapInt32Float32PtrErr(notOneInt32Float32PtrErr, plusOneInt32Float32PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float32PtrErr failed")
	}
	_, err = FilterMapInt32Float32PtrErr(notOneInt32Float32PtrErr, plusOneInt32Float32PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float32PtrErr failed")
	}
}
func notOneInt32Float32PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Float32PtrErr(num *int32) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapInt32Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapInt32Float64PtrErr(notOneInt32Float64PtrErr, plusOneInt32Float64PtrErr, []*int32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt32Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float64PtrErr failed")
	}
	r, _ = FilterMapInt32Float64PtrErr(nil, nil, []*int32{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt32Float64PtrErr failed")
	}

	_, err := FilterMapInt32Float64PtrErr(notOneInt32Float64PtrErr, plusOneInt32Float64PtrErr, []*int32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float64PtrErr failed")
	}
	_, err = FilterMapInt32Float64PtrErr(notOneInt32Float64PtrErr, plusOneInt32Float64PtrErr, []*int32{&v3})
	if err == nil {
		t.Errorf("FilterMapInt32Float64PtrErr failed")
	}
}
func notOneInt32Float64PtrErr(num *int32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt32Float64PtrErr(num *int32) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapInt16IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapInt16IntPtrErr(notOneInt16IntPtrErr, plusOneInt16IntPtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16IntPtrErr failed")
	}
	r, _ = FilterMapInt16IntPtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16IntPtrErr failed")
	}

	_, err := FilterMapInt16IntPtrErr(notOneInt16IntPtrErr, plusOneInt16IntPtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16IntPtrErr failed")
	}
	_, err = FilterMapInt16IntPtrErr(notOneInt16IntPtrErr, plusOneInt16IntPtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16IntPtrErr failed")
	}
}
func notOneInt16IntPtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16IntPtrErr(num *int16) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapInt16Int64PtrErr(notOneInt16Int64PtrErr, plusOneInt16Int64PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int64PtrErr failed")
	}
	r, _ = FilterMapInt16Int64PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int64PtrErr failed")
	}

	_, err := FilterMapInt16Int64PtrErr(notOneInt16Int64PtrErr, plusOneInt16Int64PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int64PtrErr failed")
	}
	_, err = FilterMapInt16Int64PtrErr(notOneInt16Int64PtrErr, plusOneInt16Int64PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int64PtrErr failed")
	}
}
func notOneInt16Int64PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Int64PtrErr(num *int16) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapInt16Int32PtrErr(notOneInt16Int32PtrErr, plusOneInt16Int32PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int32PtrErr failed")
	}
	r, _ = FilterMapInt16Int32PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int32PtrErr failed")
	}

	_, err := FilterMapInt16Int32PtrErr(notOneInt16Int32PtrErr, plusOneInt16Int32PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int32PtrErr failed")
	}
	_, err = FilterMapInt16Int32PtrErr(notOneInt16Int32PtrErr, plusOneInt16Int32PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int32PtrErr failed")
	}
}
func notOneInt16Int32PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Int32PtrErr(num *int16) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapInt16Int8PtrErr(notOneInt16Int8PtrErr, plusOneInt16Int8PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int8PtrErr failed")
	}
	r, _ = FilterMapInt16Int8PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Int8PtrErr failed")
	}

	_, err := FilterMapInt16Int8PtrErr(notOneInt16Int8PtrErr, plusOneInt16Int8PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int8PtrErr failed")
	}
	_, err = FilterMapInt16Int8PtrErr(notOneInt16Int8PtrErr, plusOneInt16Int8PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Int8PtrErr failed")
	}
}
func notOneInt16Int8PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Int8PtrErr(num *int16) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapInt16UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapInt16UintPtrErr(notOneInt16UintPtrErr, plusOneInt16UintPtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16UintPtrErr failed")
	}
	r, _ = FilterMapInt16UintPtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16UintPtrErr failed")
	}

	_, err := FilterMapInt16UintPtrErr(notOneInt16UintPtrErr, plusOneInt16UintPtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16UintPtrErr failed")
	}
	_, err = FilterMapInt16UintPtrErr(notOneInt16UintPtrErr, plusOneInt16UintPtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16UintPtrErr failed")
	}
}
func notOneInt16UintPtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16UintPtrErr(num *int16) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapInt16Uint64PtrErr(notOneInt16Uint64PtrErr, plusOneInt16Uint64PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint64PtrErr failed")
	}
	r, _ = FilterMapInt16Uint64PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint64PtrErr failed")
	}

	_, err := FilterMapInt16Uint64PtrErr(notOneInt16Uint64PtrErr, plusOneInt16Uint64PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint64PtrErr failed")
	}
	_, err = FilterMapInt16Uint64PtrErr(notOneInt16Uint64PtrErr, plusOneInt16Uint64PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint64PtrErr failed")
	}
}
func notOneInt16Uint64PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Uint64PtrErr(num *int16) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapInt16Uint32PtrErr(notOneInt16Uint32PtrErr, plusOneInt16Uint32PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint32PtrErr failed")
	}
	r, _ = FilterMapInt16Uint32PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint32PtrErr failed")
	}

	_, err := FilterMapInt16Uint32PtrErr(notOneInt16Uint32PtrErr, plusOneInt16Uint32PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint32PtrErr failed")
	}
	_, err = FilterMapInt16Uint32PtrErr(notOneInt16Uint32PtrErr, plusOneInt16Uint32PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint32PtrErr failed")
	}
}
func notOneInt16Uint32PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Uint32PtrErr(num *int16) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapInt16Uint16PtrErr(notOneInt16Uint16PtrErr, plusOneInt16Uint16PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint16PtrErr failed")
	}
	r, _ = FilterMapInt16Uint16PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint16PtrErr failed")
	}

	_, err := FilterMapInt16Uint16PtrErr(notOneInt16Uint16PtrErr, plusOneInt16Uint16PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint16PtrErr failed")
	}
	_, err = FilterMapInt16Uint16PtrErr(notOneInt16Uint16PtrErr, plusOneInt16Uint16PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint16PtrErr failed")
	}
}
func notOneInt16Uint16PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Uint16PtrErr(num *int16) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapInt16Uint8PtrErr(notOneInt16Uint8PtrErr, plusOneInt16Uint8PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint8PtrErr failed")
	}
	r, _ = FilterMapInt16Uint8PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Uint8PtrErr failed")
	}

	_, err := FilterMapInt16Uint8PtrErr(notOneInt16Uint8PtrErr, plusOneInt16Uint8PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint8PtrErr failed")
	}
	_, err = FilterMapInt16Uint8PtrErr(notOneInt16Uint8PtrErr, plusOneInt16Uint8PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Uint8PtrErr failed")
	}
}
func notOneInt16Uint8PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Uint8PtrErr(num *int16) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapInt16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int16 = 1
	var iv2 int16 = 2
	var iv3 int16 = 3
	var iv10 int16 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapInt16StrPtrErr(notOneInt16StrNumPtrErr, someLogicInt16StrNumPtrErr, []*int16{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16StrPtrErr failed")
	}

	r, _ = FilterMapInt16StrPtrErr(nil, nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt16StrPtrErr failed")
	}

	_, err := FilterMapInt16StrPtrErr(notOneInt16StrNumPtrErr, someLogicInt16StrNumPtrErr, []*int16{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapInt16StrPtrErr failed")
	}
	_, err = FilterMapInt16StrPtrErr(notOneInt16StrNumPtrErr, someLogicInt16StrNumPtrErr, []*int16{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapInt16StrPtrErr failed")
	}
}

func notOneInt16StrNumPtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicInt16StrNumPtrErr(num *int16) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapInt16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	var vi10 int16 = 10
	var vi0 int16 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapInt16BoolPtrErr(notOneInt16BoolPtrErr, someLogicInt16BoolPtrErr, []*int16{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16BoolPtr failed")
	}

	r, _ = FilterMapInt16BoolPtrErr(nil, nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt16BoolPtrErr failed")
	}

	_, err := FilterMapInt16BoolPtrErr(notOneInt16BoolPtrErr, someLogicInt16BoolPtrErr, []*int16{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt16BoolPtrErr failed")
	}

	_, err = FilterMapInt16BoolPtrErr(notOneInt16BoolPtrErr, someLogicInt16BoolPtrErr, []*int16{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt16BoolPtrErr failed")
	}
}
func notOneInt16BoolPtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicInt16BoolPtrErr(num *int16) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapInt16Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapInt16Float32PtrErr(notOneInt16Float32PtrErr, plusOneInt16Float32PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float32PtrErr failed")
	}
	r, _ = FilterMapInt16Float32PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float32PtrErr failed")
	}

	_, err := FilterMapInt16Float32PtrErr(notOneInt16Float32PtrErr, plusOneInt16Float32PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float32PtrErr failed")
	}
	_, err = FilterMapInt16Float32PtrErr(notOneInt16Float32PtrErr, plusOneInt16Float32PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float32PtrErr failed")
	}
}
func notOneInt16Float32PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Float32PtrErr(num *int16) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapInt16Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapInt16Float64PtrErr(notOneInt16Float64PtrErr, plusOneInt16Float64PtrErr, []*int16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt16Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float64PtrErr failed")
	}
	r, _ = FilterMapInt16Float64PtrErr(nil, nil, []*int16{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt16Float64PtrErr failed")
	}

	_, err := FilterMapInt16Float64PtrErr(notOneInt16Float64PtrErr, plusOneInt16Float64PtrErr, []*int16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float64PtrErr failed")
	}
	_, err = FilterMapInt16Float64PtrErr(notOneInt16Float64PtrErr, plusOneInt16Float64PtrErr, []*int16{&v3})
	if err == nil {
		t.Errorf("FilterMapInt16Float64PtrErr failed")
	}
}
func notOneInt16Float64PtrErr(num *int16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt16Float64PtrErr(num *int16) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapInt8IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapInt8IntPtrErr(notOneInt8IntPtrErr, plusOneInt8IntPtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8IntPtrErr failed")
	}
	r, _ = FilterMapInt8IntPtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8IntPtrErr failed")
	}

	_, err := FilterMapInt8IntPtrErr(notOneInt8IntPtrErr, plusOneInt8IntPtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8IntPtrErr failed")
	}
	_, err = FilterMapInt8IntPtrErr(notOneInt8IntPtrErr, plusOneInt8IntPtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8IntPtrErr failed")
	}
}
func notOneInt8IntPtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8IntPtrErr(num *int8) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapInt8Int64PtrErr(notOneInt8Int64PtrErr, plusOneInt8Int64PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int64PtrErr failed")
	}
	r, _ = FilterMapInt8Int64PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int64PtrErr failed")
	}

	_, err := FilterMapInt8Int64PtrErr(notOneInt8Int64PtrErr, plusOneInt8Int64PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int64PtrErr failed")
	}
	_, err = FilterMapInt8Int64PtrErr(notOneInt8Int64PtrErr, plusOneInt8Int64PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int64PtrErr failed")
	}
}
func notOneInt8Int64PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Int64PtrErr(num *int8) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapInt8Int32PtrErr(notOneInt8Int32PtrErr, plusOneInt8Int32PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int32PtrErr failed")
	}
	r, _ = FilterMapInt8Int32PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int32PtrErr failed")
	}

	_, err := FilterMapInt8Int32PtrErr(notOneInt8Int32PtrErr, plusOneInt8Int32PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int32PtrErr failed")
	}
	_, err = FilterMapInt8Int32PtrErr(notOneInt8Int32PtrErr, plusOneInt8Int32PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int32PtrErr failed")
	}
}
func notOneInt8Int32PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Int32PtrErr(num *int8) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapInt8Int16PtrErr(notOneInt8Int16PtrErr, plusOneInt8Int16PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int16PtrErr failed")
	}
	r, _ = FilterMapInt8Int16PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Int16PtrErr failed")
	}

	_, err := FilterMapInt8Int16PtrErr(notOneInt8Int16PtrErr, plusOneInt8Int16PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int16PtrErr failed")
	}
	_, err = FilterMapInt8Int16PtrErr(notOneInt8Int16PtrErr, plusOneInt8Int16PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Int16PtrErr failed")
	}
}
func notOneInt8Int16PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Int16PtrErr(num *int8) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapInt8UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapInt8UintPtrErr(notOneInt8UintPtrErr, plusOneInt8UintPtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8UintPtrErr failed")
	}
	r, _ = FilterMapInt8UintPtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8UintPtrErr failed")
	}

	_, err := FilterMapInt8UintPtrErr(notOneInt8UintPtrErr, plusOneInt8UintPtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8UintPtrErr failed")
	}
	_, err = FilterMapInt8UintPtrErr(notOneInt8UintPtrErr, plusOneInt8UintPtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8UintPtrErr failed")
	}
}
func notOneInt8UintPtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8UintPtrErr(num *int8) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapInt8Uint64PtrErr(notOneInt8Uint64PtrErr, plusOneInt8Uint64PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint64PtrErr failed")
	}
	r, _ = FilterMapInt8Uint64PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint64PtrErr failed")
	}

	_, err := FilterMapInt8Uint64PtrErr(notOneInt8Uint64PtrErr, plusOneInt8Uint64PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint64PtrErr failed")
	}
	_, err = FilterMapInt8Uint64PtrErr(notOneInt8Uint64PtrErr, plusOneInt8Uint64PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint64PtrErr failed")
	}
}
func notOneInt8Uint64PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Uint64PtrErr(num *int8) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapInt8Uint32PtrErr(notOneInt8Uint32PtrErr, plusOneInt8Uint32PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint32PtrErr failed")
	}
	r, _ = FilterMapInt8Uint32PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint32PtrErr failed")
	}

	_, err := FilterMapInt8Uint32PtrErr(notOneInt8Uint32PtrErr, plusOneInt8Uint32PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint32PtrErr failed")
	}
	_, err = FilterMapInt8Uint32PtrErr(notOneInt8Uint32PtrErr, plusOneInt8Uint32PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint32PtrErr failed")
	}
}
func notOneInt8Uint32PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Uint32PtrErr(num *int8) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapInt8Uint16PtrErr(notOneInt8Uint16PtrErr, plusOneInt8Uint16PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint16PtrErr failed")
	}
	r, _ = FilterMapInt8Uint16PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint16PtrErr failed")
	}

	_, err := FilterMapInt8Uint16PtrErr(notOneInt8Uint16PtrErr, plusOneInt8Uint16PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint16PtrErr failed")
	}
	_, err = FilterMapInt8Uint16PtrErr(notOneInt8Uint16PtrErr, plusOneInt8Uint16PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint16PtrErr failed")
	}
}
func notOneInt8Uint16PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Uint16PtrErr(num *int8) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapInt8Uint8PtrErr(notOneInt8Uint8PtrErr, plusOneInt8Uint8PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint8PtrErr failed")
	}
	r, _ = FilterMapInt8Uint8PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Uint8PtrErr failed")
	}

	_, err := FilterMapInt8Uint8PtrErr(notOneInt8Uint8PtrErr, plusOneInt8Uint8PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint8PtrErr failed")
	}
	_, err = FilterMapInt8Uint8PtrErr(notOneInt8Uint8PtrErr, plusOneInt8Uint8PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Uint8PtrErr failed")
	}
}
func notOneInt8Uint8PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Uint8PtrErr(num *int8) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapInt8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int8 = 1
	var iv2 int8 = 2
	var iv3 int8 = 3
	var iv10 int8 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapInt8StrPtrErr(notOneInt8StrNumPtrErr, someLogicInt8StrNumPtrErr, []*int8{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8StrPtrErr failed")
	}

	r, _ = FilterMapInt8StrPtrErr(nil, nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt8StrPtrErr failed")
	}

	_, err := FilterMapInt8StrPtrErr(notOneInt8StrNumPtrErr, someLogicInt8StrNumPtrErr, []*int8{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapInt8StrPtrErr failed")
	}
	_, err = FilterMapInt8StrPtrErr(notOneInt8StrNumPtrErr, someLogicInt8StrNumPtrErr, []*int8{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapInt8StrPtrErr failed")
	}
}

func notOneInt8StrNumPtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicInt8StrNumPtrErr(num *int8) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapInt8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	var vi10 int8 = 10
	var vi0 int8 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapInt8BoolPtrErr(notOneInt8BoolPtrErr, someLogicInt8BoolPtrErr, []*int8{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8BoolPtr failed")
	}

	r, _ = FilterMapInt8BoolPtrErr(nil, nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("FilterMapInt8BoolPtrErr failed")
	}

	_, err := FilterMapInt8BoolPtrErr(notOneInt8BoolPtrErr, someLogicInt8BoolPtrErr, []*int8{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt8BoolPtrErr failed")
	}

	_, err = FilterMapInt8BoolPtrErr(notOneInt8BoolPtrErr, someLogicInt8BoolPtrErr, []*int8{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapInt8BoolPtrErr failed")
	}
}
func notOneInt8BoolPtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicInt8BoolPtrErr(num *int8) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapInt8Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapInt8Float32PtrErr(notOneInt8Float32PtrErr, plusOneInt8Float32PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float32PtrErr failed")
	}
	r, _ = FilterMapInt8Float32PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float32PtrErr failed")
	}

	_, err := FilterMapInt8Float32PtrErr(notOneInt8Float32PtrErr, plusOneInt8Float32PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float32PtrErr failed")
	}
	_, err = FilterMapInt8Float32PtrErr(notOneInt8Float32PtrErr, plusOneInt8Float32PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float32PtrErr failed")
	}
}
func notOneInt8Float32PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Float32PtrErr(num *int8) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapInt8Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapInt8Float64PtrErr(notOneInt8Float64PtrErr, plusOneInt8Float64PtrErr, []*int8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapInt8Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float64PtrErr failed")
	}
	r, _ = FilterMapInt8Float64PtrErr(nil, nil, []*int8{})

	if len(r) > 0 {
		t.Errorf("FilterMapInt8Float64PtrErr failed")
	}

	_, err := FilterMapInt8Float64PtrErr(notOneInt8Float64PtrErr, plusOneInt8Float64PtrErr, []*int8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float64PtrErr failed")
	}
	_, err = FilterMapInt8Float64PtrErr(notOneInt8Float64PtrErr, plusOneInt8Float64PtrErr, []*int8{&v3})
	if err == nil {
		t.Errorf("FilterMapInt8Float64PtrErr failed")
	}
}
func notOneInt8Float64PtrErr(num *int8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneInt8Float64PtrErr(num *int8) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapUintIntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapUintIntPtrErr(notOneUintIntPtrErr, plusOneUintIntPtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintIntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintIntPtrErr failed")
	}
	r, _ = FilterMapUintIntPtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintIntPtrErr failed")
	}

	_, err := FilterMapUintIntPtrErr(notOneUintIntPtrErr, plusOneUintIntPtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintIntPtrErr failed")
	}
	_, err = FilterMapUintIntPtrErr(notOneUintIntPtrErr, plusOneUintIntPtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintIntPtrErr failed")
	}
}
func notOneUintIntPtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintIntPtrErr(num *uint) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapUintInt64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapUintInt64PtrErr(notOneUintInt64PtrErr, plusOneUintInt64PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt64PtrErr failed")
	}
	r, _ = FilterMapUintInt64PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt64PtrErr failed")
	}

	_, err := FilterMapUintInt64PtrErr(notOneUintInt64PtrErr, plusOneUintInt64PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintInt64PtrErr failed")
	}
	_, err = FilterMapUintInt64PtrErr(notOneUintInt64PtrErr, plusOneUintInt64PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintInt64PtrErr failed")
	}
}
func notOneUintInt64PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintInt64PtrErr(num *uint) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapUintInt32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapUintInt32PtrErr(notOneUintInt32PtrErr, plusOneUintInt32PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt32PtrErr failed")
	}
	r, _ = FilterMapUintInt32PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt32PtrErr failed")
	}

	_, err := FilterMapUintInt32PtrErr(notOneUintInt32PtrErr, plusOneUintInt32PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintInt32PtrErr failed")
	}
	_, err = FilterMapUintInt32PtrErr(notOneUintInt32PtrErr, plusOneUintInt32PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintInt32PtrErr failed")
	}
}
func notOneUintInt32PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintInt32PtrErr(num *uint) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapUintInt16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapUintInt16PtrErr(notOneUintInt16PtrErr, plusOneUintInt16PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt16PtrErr failed")
	}
	r, _ = FilterMapUintInt16PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt16PtrErr failed")
	}

	_, err := FilterMapUintInt16PtrErr(notOneUintInt16PtrErr, plusOneUintInt16PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintInt16PtrErr failed")
	}
	_, err = FilterMapUintInt16PtrErr(notOneUintInt16PtrErr, plusOneUintInt16PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintInt16PtrErr failed")
	}
}
func notOneUintInt16PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintInt16PtrErr(num *uint) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapUintInt8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapUintInt8PtrErr(notOneUintInt8PtrErr, plusOneUintInt8PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintInt8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintInt8PtrErr failed")
	}
	r, _ = FilterMapUintInt8PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintInt8PtrErr failed")
	}

	_, err := FilterMapUintInt8PtrErr(notOneUintInt8PtrErr, plusOneUintInt8PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintInt8PtrErr failed")
	}
	_, err = FilterMapUintInt8PtrErr(notOneUintInt8PtrErr, plusOneUintInt8PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintInt8PtrErr failed")
	}
}
func notOneUintInt8PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintInt8PtrErr(num *uint) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapUintUint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapUintUint64PtrErr(notOneUintUint64PtrErr, plusOneUintUint64PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint64PtrErr failed")
	}
	r, _ = FilterMapUintUint64PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint64PtrErr failed")
	}

	_, err := FilterMapUintUint64PtrErr(notOneUintUint64PtrErr, plusOneUintUint64PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintUint64PtrErr failed")
	}
	_, err = FilterMapUintUint64PtrErr(notOneUintUint64PtrErr, plusOneUintUint64PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintUint64PtrErr failed")
	}
}
func notOneUintUint64PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintUint64PtrErr(num *uint) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapUintUint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapUintUint32PtrErr(notOneUintUint32PtrErr, plusOneUintUint32PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint32PtrErr failed")
	}
	r, _ = FilterMapUintUint32PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint32PtrErr failed")
	}

	_, err := FilterMapUintUint32PtrErr(notOneUintUint32PtrErr, plusOneUintUint32PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintUint32PtrErr failed")
	}
	_, err = FilterMapUintUint32PtrErr(notOneUintUint32PtrErr, plusOneUintUint32PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintUint32PtrErr failed")
	}
}
func notOneUintUint32PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintUint32PtrErr(num *uint) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapUintUint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapUintUint16PtrErr(notOneUintUint16PtrErr, plusOneUintUint16PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint16PtrErr failed")
	}
	r, _ = FilterMapUintUint16PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint16PtrErr failed")
	}

	_, err := FilterMapUintUint16PtrErr(notOneUintUint16PtrErr, plusOneUintUint16PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintUint16PtrErr failed")
	}
	_, err = FilterMapUintUint16PtrErr(notOneUintUint16PtrErr, plusOneUintUint16PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintUint16PtrErr failed")
	}
}
func notOneUintUint16PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintUint16PtrErr(num *uint) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapUintUint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapUintUint8PtrErr(notOneUintUint8PtrErr, plusOneUintUint8PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintUint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintUint8PtrErr failed")
	}
	r, _ = FilterMapUintUint8PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintUint8PtrErr failed")
	}

	_, err := FilterMapUintUint8PtrErr(notOneUintUint8PtrErr, plusOneUintUint8PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintUint8PtrErr failed")
	}
	_, err = FilterMapUintUint8PtrErr(notOneUintUint8PtrErr, plusOneUintUint8PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintUint8PtrErr failed")
	}
}
func notOneUintUint8PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintUint8PtrErr(num *uint) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapUintStrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint = 1
	var iv2 uint = 2
	var iv3 uint = 3
	var iv10 uint = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapUintStrPtrErr(notOneUintStrNumPtrErr, someLogicUintStrNumPtrErr, []*uint{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUintStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintStrPtrErr failed")
	}

	r, _ = FilterMapUintStrPtrErr(nil, nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("FilterMapUintStrPtrErr failed")
	}

	_, err := FilterMapUintStrPtrErr(notOneUintStrNumPtrErr, someLogicUintStrNumPtrErr, []*uint{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapUintStrPtrErr failed")
	}
	_, err = FilterMapUintStrPtrErr(notOneUintStrNumPtrErr, someLogicUintStrNumPtrErr, []*uint{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapUintStrPtrErr failed")
	}
}

func notOneUintStrNumPtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicUintStrNumPtrErr(num *uint) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapUintBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	var vi10 uint = 10
	var vi0 uint = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapUintBoolPtrErr(notOneUintBoolPtrErr, someLogicUintBoolPtrErr, []*uint{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintBoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintBoolPtr failed")
	}

	r, _ = FilterMapUintBoolPtrErr(nil, nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("FilterMapUintBoolPtrErr failed")
	}

	_, err := FilterMapUintBoolPtrErr(notOneUintBoolPtrErr, someLogicUintBoolPtrErr, []*uint{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUintBoolPtrErr failed")
	}

	_, err = FilterMapUintBoolPtrErr(notOneUintBoolPtrErr, someLogicUintBoolPtrErr, []*uint{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUintBoolPtrErr failed")
	}
}
func notOneUintBoolPtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicUintBoolPtrErr(num *uint) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapUintFloat32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapUintFloat32PtrErr(notOneUintFloat32PtrErr, plusOneUintFloat32PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintFloat32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat32PtrErr failed")
	}
	r, _ = FilterMapUintFloat32PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat32PtrErr failed")
	}

	_, err := FilterMapUintFloat32PtrErr(notOneUintFloat32PtrErr, plusOneUintFloat32PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat32PtrErr failed")
	}
	_, err = FilterMapUintFloat32PtrErr(notOneUintFloat32PtrErr, plusOneUintFloat32PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat32PtrErr failed")
	}
}
func notOneUintFloat32PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintFloat32PtrErr(num *uint) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapUintFloat64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapUintFloat64PtrErr(notOneUintFloat64PtrErr, plusOneUintFloat64PtrErr, []*uint{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUintFloat64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat64PtrErr failed")
	}
	r, _ = FilterMapUintFloat64PtrErr(nil, nil, []*uint{})

	if len(r) > 0 {
		t.Errorf("FilterMapUintFloat64PtrErr failed")
	}

	_, err := FilterMapUintFloat64PtrErr(notOneUintFloat64PtrErr, plusOneUintFloat64PtrErr, []*uint{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat64PtrErr failed")
	}
	_, err = FilterMapUintFloat64PtrErr(notOneUintFloat64PtrErr, plusOneUintFloat64PtrErr, []*uint{&v3})
	if err == nil {
		t.Errorf("FilterMapUintFloat64PtrErr failed")
	}
}
func notOneUintFloat64PtrErr(num *uint) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUintFloat64PtrErr(num *uint) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapUint64IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapUint64IntPtrErr(notOneUint64IntPtrErr, plusOneUint64IntPtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64IntPtrErr failed")
	}
	r, _ = FilterMapUint64IntPtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64IntPtrErr failed")
	}

	_, err := FilterMapUint64IntPtrErr(notOneUint64IntPtrErr, plusOneUint64IntPtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64IntPtrErr failed")
	}
	_, err = FilterMapUint64IntPtrErr(notOneUint64IntPtrErr, plusOneUint64IntPtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64IntPtrErr failed")
	}
}
func notOneUint64IntPtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64IntPtrErr(num *uint64) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapUint64Int64PtrErr(notOneUint64Int64PtrErr, plusOneUint64Int64PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int64PtrErr failed")
	}
	r, _ = FilterMapUint64Int64PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int64PtrErr failed")
	}

	_, err := FilterMapUint64Int64PtrErr(notOneUint64Int64PtrErr, plusOneUint64Int64PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int64PtrErr failed")
	}
	_, err = FilterMapUint64Int64PtrErr(notOneUint64Int64PtrErr, plusOneUint64Int64PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int64PtrErr failed")
	}
}
func notOneUint64Int64PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Int64PtrErr(num *uint64) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapUint64Int32PtrErr(notOneUint64Int32PtrErr, plusOneUint64Int32PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int32PtrErr failed")
	}
	r, _ = FilterMapUint64Int32PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int32PtrErr failed")
	}

	_, err := FilterMapUint64Int32PtrErr(notOneUint64Int32PtrErr, plusOneUint64Int32PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int32PtrErr failed")
	}
	_, err = FilterMapUint64Int32PtrErr(notOneUint64Int32PtrErr, plusOneUint64Int32PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int32PtrErr failed")
	}
}
func notOneUint64Int32PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Int32PtrErr(num *uint64) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapUint64Int16PtrErr(notOneUint64Int16PtrErr, plusOneUint64Int16PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int16PtrErr failed")
	}
	r, _ = FilterMapUint64Int16PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int16PtrErr failed")
	}

	_, err := FilterMapUint64Int16PtrErr(notOneUint64Int16PtrErr, plusOneUint64Int16PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int16PtrErr failed")
	}
	_, err = FilterMapUint64Int16PtrErr(notOneUint64Int16PtrErr, plusOneUint64Int16PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int16PtrErr failed")
	}
}
func notOneUint64Int16PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Int16PtrErr(num *uint64) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapUint64Int8PtrErr(notOneUint64Int8PtrErr, plusOneUint64Int8PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int8PtrErr failed")
	}
	r, _ = FilterMapUint64Int8PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Int8PtrErr failed")
	}

	_, err := FilterMapUint64Int8PtrErr(notOneUint64Int8PtrErr, plusOneUint64Int8PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int8PtrErr failed")
	}
	_, err = FilterMapUint64Int8PtrErr(notOneUint64Int8PtrErr, plusOneUint64Int8PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Int8PtrErr failed")
	}
}
func notOneUint64Int8PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Int8PtrErr(num *uint64) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapUint64UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapUint64UintPtrErr(notOneUint64UintPtrErr, plusOneUint64UintPtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64UintPtrErr failed")
	}
	r, _ = FilterMapUint64UintPtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64UintPtrErr failed")
	}

	_, err := FilterMapUint64UintPtrErr(notOneUint64UintPtrErr, plusOneUint64UintPtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64UintPtrErr failed")
	}
	_, err = FilterMapUint64UintPtrErr(notOneUint64UintPtrErr, plusOneUint64UintPtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64UintPtrErr failed")
	}
}
func notOneUint64UintPtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64UintPtrErr(num *uint64) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapUint64Uint32PtrErr(notOneUint64Uint32PtrErr, plusOneUint64Uint32PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint32PtrErr failed")
	}
	r, _ = FilterMapUint64Uint32PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint32PtrErr failed")
	}

	_, err := FilterMapUint64Uint32PtrErr(notOneUint64Uint32PtrErr, plusOneUint64Uint32PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint32PtrErr failed")
	}
	_, err = FilterMapUint64Uint32PtrErr(notOneUint64Uint32PtrErr, plusOneUint64Uint32PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint32PtrErr failed")
	}
}
func notOneUint64Uint32PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Uint32PtrErr(num *uint64) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapUint64Uint16PtrErr(notOneUint64Uint16PtrErr, plusOneUint64Uint16PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint16PtrErr failed")
	}
	r, _ = FilterMapUint64Uint16PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint16PtrErr failed")
	}

	_, err := FilterMapUint64Uint16PtrErr(notOneUint64Uint16PtrErr, plusOneUint64Uint16PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint16PtrErr failed")
	}
	_, err = FilterMapUint64Uint16PtrErr(notOneUint64Uint16PtrErr, plusOneUint64Uint16PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint16PtrErr failed")
	}
}
func notOneUint64Uint16PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Uint16PtrErr(num *uint64) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapUint64Uint8PtrErr(notOneUint64Uint8PtrErr, plusOneUint64Uint8PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint8PtrErr failed")
	}
	r, _ = FilterMapUint64Uint8PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Uint8PtrErr failed")
	}

	_, err := FilterMapUint64Uint8PtrErr(notOneUint64Uint8PtrErr, plusOneUint64Uint8PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint8PtrErr failed")
	}
	_, err = FilterMapUint64Uint8PtrErr(notOneUint64Uint8PtrErr, plusOneUint64Uint8PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Uint8PtrErr failed")
	}
}
func notOneUint64Uint8PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Uint8PtrErr(num *uint64) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapUint64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint64 = 1
	var iv2 uint64 = 2
	var iv3 uint64 = 3
	var iv10 uint64 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapUint64StrPtrErr(notOneUint64StrNumPtrErr, someLogicUint64StrNumPtrErr, []*uint64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64StrPtrErr failed")
	}

	r, _ = FilterMapUint64StrPtrErr(nil, nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint64StrPtrErr failed")
	}

	_, err := FilterMapUint64StrPtrErr(notOneUint64StrNumPtrErr, someLogicUint64StrNumPtrErr, []*uint64{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapUint64StrPtrErr failed")
	}
	_, err = FilterMapUint64StrPtrErr(notOneUint64StrNumPtrErr, someLogicUint64StrNumPtrErr, []*uint64{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapUint64StrPtrErr failed")
	}
}

func notOneUint64StrNumPtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicUint64StrNumPtrErr(num *uint64) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapUint64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	var vi10 uint64 = 10
	var vi0 uint64 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapUint64BoolPtrErr(notOneUint64BoolPtrErr, someLogicUint64BoolPtrErr, []*uint64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64BoolPtr failed")
	}

	r, _ = FilterMapUint64BoolPtrErr(nil, nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint64BoolPtrErr failed")
	}

	_, err := FilterMapUint64BoolPtrErr(notOneUint64BoolPtrErr, someLogicUint64BoolPtrErr, []*uint64{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint64BoolPtrErr failed")
	}

	_, err = FilterMapUint64BoolPtrErr(notOneUint64BoolPtrErr, someLogicUint64BoolPtrErr, []*uint64{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint64BoolPtrErr failed")
	}
}
func notOneUint64BoolPtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicUint64BoolPtrErr(num *uint64) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapUint64Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapUint64Float32PtrErr(notOneUint64Float32PtrErr, plusOneUint64Float32PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float32PtrErr failed")
	}
	r, _ = FilterMapUint64Float32PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float32PtrErr failed")
	}

	_, err := FilterMapUint64Float32PtrErr(notOneUint64Float32PtrErr, plusOneUint64Float32PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float32PtrErr failed")
	}
	_, err = FilterMapUint64Float32PtrErr(notOneUint64Float32PtrErr, plusOneUint64Float32PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float32PtrErr failed")
	}
}
func notOneUint64Float32PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Float32PtrErr(num *uint64) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapUint64Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapUint64Float64PtrErr(notOneUint64Float64PtrErr, plusOneUint64Float64PtrErr, []*uint64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint64Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float64PtrErr failed")
	}
	r, _ = FilterMapUint64Float64PtrErr(nil, nil, []*uint64{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint64Float64PtrErr failed")
	}

	_, err := FilterMapUint64Float64PtrErr(notOneUint64Float64PtrErr, plusOneUint64Float64PtrErr, []*uint64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float64PtrErr failed")
	}
	_, err = FilterMapUint64Float64PtrErr(notOneUint64Float64PtrErr, plusOneUint64Float64PtrErr, []*uint64{&v3})
	if err == nil {
		t.Errorf("FilterMapUint64Float64PtrErr failed")
	}
}
func notOneUint64Float64PtrErr(num *uint64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint64Float64PtrErr(num *uint64) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapUint32IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapUint32IntPtrErr(notOneUint32IntPtrErr, plusOneUint32IntPtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32IntPtrErr failed")
	}
	r, _ = FilterMapUint32IntPtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32IntPtrErr failed")
	}

	_, err := FilterMapUint32IntPtrErr(notOneUint32IntPtrErr, plusOneUint32IntPtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32IntPtrErr failed")
	}
	_, err = FilterMapUint32IntPtrErr(notOneUint32IntPtrErr, plusOneUint32IntPtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32IntPtrErr failed")
	}
}
func notOneUint32IntPtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32IntPtrErr(num *uint32) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapUint32Int64PtrErr(notOneUint32Int64PtrErr, plusOneUint32Int64PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int64PtrErr failed")
	}
	r, _ = FilterMapUint32Int64PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int64PtrErr failed")
	}

	_, err := FilterMapUint32Int64PtrErr(notOneUint32Int64PtrErr, plusOneUint32Int64PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int64PtrErr failed")
	}
	_, err = FilterMapUint32Int64PtrErr(notOneUint32Int64PtrErr, plusOneUint32Int64PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int64PtrErr failed")
	}
}
func notOneUint32Int64PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Int64PtrErr(num *uint32) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapUint32Int32PtrErr(notOneUint32Int32PtrErr, plusOneUint32Int32PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int32PtrErr failed")
	}
	r, _ = FilterMapUint32Int32PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int32PtrErr failed")
	}

	_, err := FilterMapUint32Int32PtrErr(notOneUint32Int32PtrErr, plusOneUint32Int32PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int32PtrErr failed")
	}
	_, err = FilterMapUint32Int32PtrErr(notOneUint32Int32PtrErr, plusOneUint32Int32PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int32PtrErr failed")
	}
}
func notOneUint32Int32PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Int32PtrErr(num *uint32) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapUint32Int16PtrErr(notOneUint32Int16PtrErr, plusOneUint32Int16PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int16PtrErr failed")
	}
	r, _ = FilterMapUint32Int16PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int16PtrErr failed")
	}

	_, err := FilterMapUint32Int16PtrErr(notOneUint32Int16PtrErr, plusOneUint32Int16PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int16PtrErr failed")
	}
	_, err = FilterMapUint32Int16PtrErr(notOneUint32Int16PtrErr, plusOneUint32Int16PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int16PtrErr failed")
	}
}
func notOneUint32Int16PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Int16PtrErr(num *uint32) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapUint32Int8PtrErr(notOneUint32Int8PtrErr, plusOneUint32Int8PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int8PtrErr failed")
	}
	r, _ = FilterMapUint32Int8PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Int8PtrErr failed")
	}

	_, err := FilterMapUint32Int8PtrErr(notOneUint32Int8PtrErr, plusOneUint32Int8PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int8PtrErr failed")
	}
	_, err = FilterMapUint32Int8PtrErr(notOneUint32Int8PtrErr, plusOneUint32Int8PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Int8PtrErr failed")
	}
}
func notOneUint32Int8PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Int8PtrErr(num *uint32) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapUint32UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapUint32UintPtrErr(notOneUint32UintPtrErr, plusOneUint32UintPtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32UintPtrErr failed")
	}
	r, _ = FilterMapUint32UintPtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32UintPtrErr failed")
	}

	_, err := FilterMapUint32UintPtrErr(notOneUint32UintPtrErr, plusOneUint32UintPtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32UintPtrErr failed")
	}
	_, err = FilterMapUint32UintPtrErr(notOneUint32UintPtrErr, plusOneUint32UintPtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32UintPtrErr failed")
	}
}
func notOneUint32UintPtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32UintPtrErr(num *uint32) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapUint32Uint64PtrErr(notOneUint32Uint64PtrErr, plusOneUint32Uint64PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint64PtrErr failed")
	}
	r, _ = FilterMapUint32Uint64PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint64PtrErr failed")
	}

	_, err := FilterMapUint32Uint64PtrErr(notOneUint32Uint64PtrErr, plusOneUint32Uint64PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint64PtrErr failed")
	}
	_, err = FilterMapUint32Uint64PtrErr(notOneUint32Uint64PtrErr, plusOneUint32Uint64PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint64PtrErr failed")
	}
}
func notOneUint32Uint64PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Uint64PtrErr(num *uint32) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapUint32Uint16PtrErr(notOneUint32Uint16PtrErr, plusOneUint32Uint16PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint16PtrErr failed")
	}
	r, _ = FilterMapUint32Uint16PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint16PtrErr failed")
	}

	_, err := FilterMapUint32Uint16PtrErr(notOneUint32Uint16PtrErr, plusOneUint32Uint16PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint16PtrErr failed")
	}
	_, err = FilterMapUint32Uint16PtrErr(notOneUint32Uint16PtrErr, plusOneUint32Uint16PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint16PtrErr failed")
	}
}
func notOneUint32Uint16PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Uint16PtrErr(num *uint32) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapUint32Uint8PtrErr(notOneUint32Uint8PtrErr, plusOneUint32Uint8PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint8PtrErr failed")
	}
	r, _ = FilterMapUint32Uint8PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Uint8PtrErr failed")
	}

	_, err := FilterMapUint32Uint8PtrErr(notOneUint32Uint8PtrErr, plusOneUint32Uint8PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint8PtrErr failed")
	}
	_, err = FilterMapUint32Uint8PtrErr(notOneUint32Uint8PtrErr, plusOneUint32Uint8PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Uint8PtrErr failed")
	}
}
func notOneUint32Uint8PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Uint8PtrErr(num *uint32) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapUint32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint32 = 1
	var iv2 uint32 = 2
	var iv3 uint32 = 3
	var iv10 uint32 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapUint32StrPtrErr(notOneUint32StrNumPtrErr, someLogicUint32StrNumPtrErr, []*uint32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32StrPtrErr failed")
	}

	r, _ = FilterMapUint32StrPtrErr(nil, nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint32StrPtrErr failed")
	}

	_, err := FilterMapUint32StrPtrErr(notOneUint32StrNumPtrErr, someLogicUint32StrNumPtrErr, []*uint32{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapUint32StrPtrErr failed")
	}
	_, err = FilterMapUint32StrPtrErr(notOneUint32StrNumPtrErr, someLogicUint32StrNumPtrErr, []*uint32{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapUint32StrPtrErr failed")
	}
}

func notOneUint32StrNumPtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicUint32StrNumPtrErr(num *uint32) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapUint32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	var vi10 uint32 = 10
	var vi0 uint32 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapUint32BoolPtrErr(notOneUint32BoolPtrErr, someLogicUint32BoolPtrErr, []*uint32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32BoolPtr failed")
	}

	r, _ = FilterMapUint32BoolPtrErr(nil, nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint32BoolPtrErr failed")
	}

	_, err := FilterMapUint32BoolPtrErr(notOneUint32BoolPtrErr, someLogicUint32BoolPtrErr, []*uint32{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint32BoolPtrErr failed")
	}

	_, err = FilterMapUint32BoolPtrErr(notOneUint32BoolPtrErr, someLogicUint32BoolPtrErr, []*uint32{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint32BoolPtrErr failed")
	}
}
func notOneUint32BoolPtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicUint32BoolPtrErr(num *uint32) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapUint32Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapUint32Float32PtrErr(notOneUint32Float32PtrErr, plusOneUint32Float32PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float32PtrErr failed")
	}
	r, _ = FilterMapUint32Float32PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float32PtrErr failed")
	}

	_, err := FilterMapUint32Float32PtrErr(notOneUint32Float32PtrErr, plusOneUint32Float32PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float32PtrErr failed")
	}
	_, err = FilterMapUint32Float32PtrErr(notOneUint32Float32PtrErr, plusOneUint32Float32PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float32PtrErr failed")
	}
}
func notOneUint32Float32PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Float32PtrErr(num *uint32) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapUint32Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapUint32Float64PtrErr(notOneUint32Float64PtrErr, plusOneUint32Float64PtrErr, []*uint32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint32Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float64PtrErr failed")
	}
	r, _ = FilterMapUint32Float64PtrErr(nil, nil, []*uint32{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint32Float64PtrErr failed")
	}

	_, err := FilterMapUint32Float64PtrErr(notOneUint32Float64PtrErr, plusOneUint32Float64PtrErr, []*uint32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float64PtrErr failed")
	}
	_, err = FilterMapUint32Float64PtrErr(notOneUint32Float64PtrErr, plusOneUint32Float64PtrErr, []*uint32{&v3})
	if err == nil {
		t.Errorf("FilterMapUint32Float64PtrErr failed")
	}
}
func notOneUint32Float64PtrErr(num *uint32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint32Float64PtrErr(num *uint32) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapUint16IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapUint16IntPtrErr(notOneUint16IntPtrErr, plusOneUint16IntPtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16IntPtrErr failed")
	}
	r, _ = FilterMapUint16IntPtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16IntPtrErr failed")
	}

	_, err := FilterMapUint16IntPtrErr(notOneUint16IntPtrErr, plusOneUint16IntPtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16IntPtrErr failed")
	}
	_, err = FilterMapUint16IntPtrErr(notOneUint16IntPtrErr, plusOneUint16IntPtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16IntPtrErr failed")
	}
}
func notOneUint16IntPtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16IntPtrErr(num *uint16) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapUint16Int64PtrErr(notOneUint16Int64PtrErr, plusOneUint16Int64PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int64PtrErr failed")
	}
	r, _ = FilterMapUint16Int64PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int64PtrErr failed")
	}

	_, err := FilterMapUint16Int64PtrErr(notOneUint16Int64PtrErr, plusOneUint16Int64PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int64PtrErr failed")
	}
	_, err = FilterMapUint16Int64PtrErr(notOneUint16Int64PtrErr, plusOneUint16Int64PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int64PtrErr failed")
	}
}
func notOneUint16Int64PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Int64PtrErr(num *uint16) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapUint16Int32PtrErr(notOneUint16Int32PtrErr, plusOneUint16Int32PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int32PtrErr failed")
	}
	r, _ = FilterMapUint16Int32PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int32PtrErr failed")
	}

	_, err := FilterMapUint16Int32PtrErr(notOneUint16Int32PtrErr, plusOneUint16Int32PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int32PtrErr failed")
	}
	_, err = FilterMapUint16Int32PtrErr(notOneUint16Int32PtrErr, plusOneUint16Int32PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int32PtrErr failed")
	}
}
func notOneUint16Int32PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Int32PtrErr(num *uint16) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapUint16Int16PtrErr(notOneUint16Int16PtrErr, plusOneUint16Int16PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int16PtrErr failed")
	}
	r, _ = FilterMapUint16Int16PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int16PtrErr failed")
	}

	_, err := FilterMapUint16Int16PtrErr(notOneUint16Int16PtrErr, plusOneUint16Int16PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int16PtrErr failed")
	}
	_, err = FilterMapUint16Int16PtrErr(notOneUint16Int16PtrErr, plusOneUint16Int16PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int16PtrErr failed")
	}
}
func notOneUint16Int16PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Int16PtrErr(num *uint16) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapUint16Int8PtrErr(notOneUint16Int8PtrErr, plusOneUint16Int8PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int8PtrErr failed")
	}
	r, _ = FilterMapUint16Int8PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Int8PtrErr failed")
	}

	_, err := FilterMapUint16Int8PtrErr(notOneUint16Int8PtrErr, plusOneUint16Int8PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int8PtrErr failed")
	}
	_, err = FilterMapUint16Int8PtrErr(notOneUint16Int8PtrErr, plusOneUint16Int8PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Int8PtrErr failed")
	}
}
func notOneUint16Int8PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Int8PtrErr(num *uint16) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapUint16UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapUint16UintPtrErr(notOneUint16UintPtrErr, plusOneUint16UintPtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16UintPtrErr failed")
	}
	r, _ = FilterMapUint16UintPtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16UintPtrErr failed")
	}

	_, err := FilterMapUint16UintPtrErr(notOneUint16UintPtrErr, plusOneUint16UintPtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16UintPtrErr failed")
	}
	_, err = FilterMapUint16UintPtrErr(notOneUint16UintPtrErr, plusOneUint16UintPtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16UintPtrErr failed")
	}
}
func notOneUint16UintPtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16UintPtrErr(num *uint16) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapUint16Uint64PtrErr(notOneUint16Uint64PtrErr, plusOneUint16Uint64PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint64PtrErr failed")
	}
	r, _ = FilterMapUint16Uint64PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint64PtrErr failed")
	}

	_, err := FilterMapUint16Uint64PtrErr(notOneUint16Uint64PtrErr, plusOneUint16Uint64PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint64PtrErr failed")
	}
	_, err = FilterMapUint16Uint64PtrErr(notOneUint16Uint64PtrErr, plusOneUint16Uint64PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint64PtrErr failed")
	}
}
func notOneUint16Uint64PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Uint64PtrErr(num *uint16) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapUint16Uint32PtrErr(notOneUint16Uint32PtrErr, plusOneUint16Uint32PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint32PtrErr failed")
	}
	r, _ = FilterMapUint16Uint32PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint32PtrErr failed")
	}

	_, err := FilterMapUint16Uint32PtrErr(notOneUint16Uint32PtrErr, plusOneUint16Uint32PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint32PtrErr failed")
	}
	_, err = FilterMapUint16Uint32PtrErr(notOneUint16Uint32PtrErr, plusOneUint16Uint32PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint32PtrErr failed")
	}
}
func notOneUint16Uint32PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Uint32PtrErr(num *uint16) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapUint16Uint8PtrErr(notOneUint16Uint8PtrErr, plusOneUint16Uint8PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint8PtrErr failed")
	}
	r, _ = FilterMapUint16Uint8PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Uint8PtrErr failed")
	}

	_, err := FilterMapUint16Uint8PtrErr(notOneUint16Uint8PtrErr, plusOneUint16Uint8PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint8PtrErr failed")
	}
	_, err = FilterMapUint16Uint8PtrErr(notOneUint16Uint8PtrErr, plusOneUint16Uint8PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Uint8PtrErr failed")
	}
}
func notOneUint16Uint8PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Uint8PtrErr(num *uint16) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapUint16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint16 = 1
	var iv2 uint16 = 2
	var iv3 uint16 = 3
	var iv10 uint16 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapUint16StrPtrErr(notOneUint16StrNumPtrErr, someLogicUint16StrNumPtrErr, []*uint16{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16StrPtrErr failed")
	}

	r, _ = FilterMapUint16StrPtrErr(nil, nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint16StrPtrErr failed")
	}

	_, err := FilterMapUint16StrPtrErr(notOneUint16StrNumPtrErr, someLogicUint16StrNumPtrErr, []*uint16{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapUint16StrPtrErr failed")
	}
	_, err = FilterMapUint16StrPtrErr(notOneUint16StrNumPtrErr, someLogicUint16StrNumPtrErr, []*uint16{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapUint16StrPtrErr failed")
	}
}

func notOneUint16StrNumPtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicUint16StrNumPtrErr(num *uint16) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapUint16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	var vi10 uint16 = 10
	var vi0 uint16 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapUint16BoolPtrErr(notOneUint16BoolPtrErr, someLogicUint16BoolPtrErr, []*uint16{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16BoolPtr failed")
	}

	r, _ = FilterMapUint16BoolPtrErr(nil, nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint16BoolPtrErr failed")
	}

	_, err := FilterMapUint16BoolPtrErr(notOneUint16BoolPtrErr, someLogicUint16BoolPtrErr, []*uint16{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint16BoolPtrErr failed")
	}

	_, err = FilterMapUint16BoolPtrErr(notOneUint16BoolPtrErr, someLogicUint16BoolPtrErr, []*uint16{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint16BoolPtrErr failed")
	}
}
func notOneUint16BoolPtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicUint16BoolPtrErr(num *uint16) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapUint16Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapUint16Float32PtrErr(notOneUint16Float32PtrErr, plusOneUint16Float32PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float32PtrErr failed")
	}
	r, _ = FilterMapUint16Float32PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float32PtrErr failed")
	}

	_, err := FilterMapUint16Float32PtrErr(notOneUint16Float32PtrErr, plusOneUint16Float32PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float32PtrErr failed")
	}
	_, err = FilterMapUint16Float32PtrErr(notOneUint16Float32PtrErr, plusOneUint16Float32PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float32PtrErr failed")
	}
}
func notOneUint16Float32PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Float32PtrErr(num *uint16) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapUint16Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapUint16Float64PtrErr(notOneUint16Float64PtrErr, plusOneUint16Float64PtrErr, []*uint16{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint16Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float64PtrErr failed")
	}
	r, _ = FilterMapUint16Float64PtrErr(nil, nil, []*uint16{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint16Float64PtrErr failed")
	}

	_, err := FilterMapUint16Float64PtrErr(notOneUint16Float64PtrErr, plusOneUint16Float64PtrErr, []*uint16{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float64PtrErr failed")
	}
	_, err = FilterMapUint16Float64PtrErr(notOneUint16Float64PtrErr, plusOneUint16Float64PtrErr, []*uint16{&v3})
	if err == nil {
		t.Errorf("FilterMapUint16Float64PtrErr failed")
	}
}
func notOneUint16Float64PtrErr(num *uint16) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint16Float64PtrErr(num *uint16) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapUint8IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapUint8IntPtrErr(notOneUint8IntPtrErr, plusOneUint8IntPtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8IntPtrErr failed")
	}
	r, _ = FilterMapUint8IntPtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8IntPtrErr failed")
	}

	_, err := FilterMapUint8IntPtrErr(notOneUint8IntPtrErr, plusOneUint8IntPtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8IntPtrErr failed")
	}
	_, err = FilterMapUint8IntPtrErr(notOneUint8IntPtrErr, plusOneUint8IntPtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8IntPtrErr failed")
	}
}
func notOneUint8IntPtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8IntPtrErr(num *uint8) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapUint8Int64PtrErr(notOneUint8Int64PtrErr, plusOneUint8Int64PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int64PtrErr failed")
	}
	r, _ = FilterMapUint8Int64PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int64PtrErr failed")
	}

	_, err := FilterMapUint8Int64PtrErr(notOneUint8Int64PtrErr, plusOneUint8Int64PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int64PtrErr failed")
	}
	_, err = FilterMapUint8Int64PtrErr(notOneUint8Int64PtrErr, plusOneUint8Int64PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int64PtrErr failed")
	}
}
func notOneUint8Int64PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Int64PtrErr(num *uint8) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapUint8Int32PtrErr(notOneUint8Int32PtrErr, plusOneUint8Int32PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int32PtrErr failed")
	}
	r, _ = FilterMapUint8Int32PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int32PtrErr failed")
	}

	_, err := FilterMapUint8Int32PtrErr(notOneUint8Int32PtrErr, plusOneUint8Int32PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int32PtrErr failed")
	}
	_, err = FilterMapUint8Int32PtrErr(notOneUint8Int32PtrErr, plusOneUint8Int32PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int32PtrErr failed")
	}
}
func notOneUint8Int32PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Int32PtrErr(num *uint8) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapUint8Int16PtrErr(notOneUint8Int16PtrErr, plusOneUint8Int16PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int16PtrErr failed")
	}
	r, _ = FilterMapUint8Int16PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int16PtrErr failed")
	}

	_, err := FilterMapUint8Int16PtrErr(notOneUint8Int16PtrErr, plusOneUint8Int16PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int16PtrErr failed")
	}
	_, err = FilterMapUint8Int16PtrErr(notOneUint8Int16PtrErr, plusOneUint8Int16PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int16PtrErr failed")
	}
}
func notOneUint8Int16PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Int16PtrErr(num *uint8) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapUint8Int8PtrErr(notOneUint8Int8PtrErr, plusOneUint8Int8PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int8PtrErr failed")
	}
	r, _ = FilterMapUint8Int8PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Int8PtrErr failed")
	}

	_, err := FilterMapUint8Int8PtrErr(notOneUint8Int8PtrErr, plusOneUint8Int8PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int8PtrErr failed")
	}
	_, err = FilterMapUint8Int8PtrErr(notOneUint8Int8PtrErr, plusOneUint8Int8PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Int8PtrErr failed")
	}
}
func notOneUint8Int8PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Int8PtrErr(num *uint8) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapUint8UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapUint8UintPtrErr(notOneUint8UintPtrErr, plusOneUint8UintPtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8UintPtrErr failed")
	}
	r, _ = FilterMapUint8UintPtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8UintPtrErr failed")
	}

	_, err := FilterMapUint8UintPtrErr(notOneUint8UintPtrErr, plusOneUint8UintPtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8UintPtrErr failed")
	}
	_, err = FilterMapUint8UintPtrErr(notOneUint8UintPtrErr, plusOneUint8UintPtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8UintPtrErr failed")
	}
}
func notOneUint8UintPtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8UintPtrErr(num *uint8) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapUint8Uint64PtrErr(notOneUint8Uint64PtrErr, plusOneUint8Uint64PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint64PtrErr failed")
	}
	r, _ = FilterMapUint8Uint64PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint64PtrErr failed")
	}

	_, err := FilterMapUint8Uint64PtrErr(notOneUint8Uint64PtrErr, plusOneUint8Uint64PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint64PtrErr failed")
	}
	_, err = FilterMapUint8Uint64PtrErr(notOneUint8Uint64PtrErr, plusOneUint8Uint64PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint64PtrErr failed")
	}
}
func notOneUint8Uint64PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Uint64PtrErr(num *uint8) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapUint8Uint32PtrErr(notOneUint8Uint32PtrErr, plusOneUint8Uint32PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint32PtrErr failed")
	}
	r, _ = FilterMapUint8Uint32PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint32PtrErr failed")
	}

	_, err := FilterMapUint8Uint32PtrErr(notOneUint8Uint32PtrErr, plusOneUint8Uint32PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint32PtrErr failed")
	}
	_, err = FilterMapUint8Uint32PtrErr(notOneUint8Uint32PtrErr, plusOneUint8Uint32PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint32PtrErr failed")
	}
}
func notOneUint8Uint32PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Uint32PtrErr(num *uint8) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapUint8Uint16PtrErr(notOneUint8Uint16PtrErr, plusOneUint8Uint16PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint16PtrErr failed")
	}
	r, _ = FilterMapUint8Uint16PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Uint16PtrErr failed")
	}

	_, err := FilterMapUint8Uint16PtrErr(notOneUint8Uint16PtrErr, plusOneUint8Uint16PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint16PtrErr failed")
	}
	_, err = FilterMapUint8Uint16PtrErr(notOneUint8Uint16PtrErr, plusOneUint8Uint16PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Uint16PtrErr failed")
	}
}
func notOneUint8Uint16PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Uint16PtrErr(num *uint8) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapUint8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint8 = 1
	var iv2 uint8 = 2
	var iv3 uint8 = 3
	var iv10 uint8 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapUint8StrPtrErr(notOneUint8StrNumPtrErr, someLogicUint8StrNumPtrErr, []*uint8{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8StrPtrErr failed")
	}

	r, _ = FilterMapUint8StrPtrErr(nil, nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint8StrPtrErr failed")
	}

	_, err := FilterMapUint8StrPtrErr(notOneUint8StrNumPtrErr, someLogicUint8StrNumPtrErr, []*uint8{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapUint8StrPtrErr failed")
	}
	_, err = FilterMapUint8StrPtrErr(notOneUint8StrNumPtrErr, someLogicUint8StrNumPtrErr, []*uint8{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapUint8StrPtrErr failed")
	}
}

func notOneUint8StrNumPtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicUint8StrNumPtrErr(num *uint8) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapUint8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	var vi10 uint8 = 10
	var vi0 uint8 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapUint8BoolPtrErr(notOneUint8BoolPtrErr, someLogicUint8BoolPtrErr, []*uint8{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8BoolPtr failed")
	}

	r, _ = FilterMapUint8BoolPtrErr(nil, nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("FilterMapUint8BoolPtrErr failed")
	}

	_, err := FilterMapUint8BoolPtrErr(notOneUint8BoolPtrErr, someLogicUint8BoolPtrErr, []*uint8{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint8BoolPtrErr failed")
	}

	_, err = FilterMapUint8BoolPtrErr(notOneUint8BoolPtrErr, someLogicUint8BoolPtrErr, []*uint8{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapUint8BoolPtrErr failed")
	}
}
func notOneUint8BoolPtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicUint8BoolPtrErr(num *uint8) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapUint8Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapUint8Float32PtrErr(notOneUint8Float32PtrErr, plusOneUint8Float32PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float32PtrErr failed")
	}
	r, _ = FilterMapUint8Float32PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float32PtrErr failed")
	}

	_, err := FilterMapUint8Float32PtrErr(notOneUint8Float32PtrErr, plusOneUint8Float32PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float32PtrErr failed")
	}
	_, err = FilterMapUint8Float32PtrErr(notOneUint8Float32PtrErr, plusOneUint8Float32PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float32PtrErr failed")
	}
}
func notOneUint8Float32PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Float32PtrErr(num *uint8) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}

func TestFilterMapUint8Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapUint8Float64PtrErr(notOneUint8Float64PtrErr, plusOneUint8Float64PtrErr, []*uint8{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapUint8Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float64PtrErr failed")
	}
	r, _ = FilterMapUint8Float64PtrErr(nil, nil, []*uint8{})

	if len(r) > 0 {
		t.Errorf("FilterMapUint8Float64PtrErr failed")
	}

	_, err := FilterMapUint8Float64PtrErr(notOneUint8Float64PtrErr, plusOneUint8Float64PtrErr, []*uint8{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float64PtrErr failed")
	}
	_, err = FilterMapUint8Float64PtrErr(notOneUint8Float64PtrErr, plusOneUint8Float64PtrErr, []*uint8{&v3})
	if err == nil {
		t.Errorf("FilterMapUint8Float64PtrErr failed")
	}
}
func notOneUint8Float64PtrErr(num *uint8) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneUint8Float64PtrErr(num *uint8) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapStrIntPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*int{&vo10}
	newList, _ := FilterMapStrIntPtrErr(notOneStrIntStrPtrErr, someLogicStrIntStrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrIntPtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrIntPtrErr failed")
	}

	r, _ = FilterMapStrIntPtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrIntPtrErr failed")
	}

	_, err := FilterMapStrIntPtrErr(notOneStrIntStrPtrErr, someLogicStrIntStrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrIntPtrErr failed")
	}
	_, err = FilterMapStrIntPtrErr(notOneStrIntStrPtrErr, someLogicStrIntStrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrIntPtrErr failed")
	}
}
func notOneStrIntStrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrIntStrPtrErr(num *string) (*int, error) {
	var r int = int(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = int(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrInt64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*int64{&vo10}
	newList, _ := FilterMapStrInt64PtrErr(notOneStrInt64StrPtrErr, someLogicStrInt64StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt64PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt64PtrErr failed")
	}

	r, _ = FilterMapStrInt64PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt64PtrErr failed")
	}

	_, err := FilterMapStrInt64PtrErr(notOneStrInt64StrPtrErr, someLogicStrInt64StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt64PtrErr failed")
	}
	_, err = FilterMapStrInt64PtrErr(notOneStrInt64StrPtrErr, someLogicStrInt64StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt64PtrErr failed")
	}
}
func notOneStrInt64StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrInt64StrPtrErr(num *string) (*int64, error) {
	var r int64 = int64(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = int64(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrInt32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*int32{&vo10}
	newList, _ := FilterMapStrInt32PtrErr(notOneStrInt32StrPtrErr, someLogicStrInt32StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt32PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt32PtrErr failed")
	}

	r, _ = FilterMapStrInt32PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt32PtrErr failed")
	}

	_, err := FilterMapStrInt32PtrErr(notOneStrInt32StrPtrErr, someLogicStrInt32StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt32PtrErr failed")
	}
	_, err = FilterMapStrInt32PtrErr(notOneStrInt32StrPtrErr, someLogicStrInt32StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt32PtrErr failed")
	}
}
func notOneStrInt32StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrInt32StrPtrErr(num *string) (*int32, error) {
	var r int32 = int32(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = int32(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrInt16PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*int16{&vo10}
	newList, _ := FilterMapStrInt16PtrErr(notOneStrInt16StrPtrErr, someLogicStrInt16StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt16PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt16PtrErr failed")
	}

	r, _ = FilterMapStrInt16PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt16PtrErr failed")
	}

	_, err := FilterMapStrInt16PtrErr(notOneStrInt16StrPtrErr, someLogicStrInt16StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt16PtrErr failed")
	}
	_, err = FilterMapStrInt16PtrErr(notOneStrInt16StrPtrErr, someLogicStrInt16StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt16PtrErr failed")
	}
}
func notOneStrInt16StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrInt16StrPtrErr(num *string) (*int16, error) {
	var r int16 = int16(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = int16(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrInt8PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*int8{&vo10}
	newList, _ := FilterMapStrInt8PtrErr(notOneStrInt8StrPtrErr, someLogicStrInt8StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrInt8PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrInt8PtrErr failed")
	}

	r, _ = FilterMapStrInt8PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrInt8PtrErr failed")
	}

	_, err := FilterMapStrInt8PtrErr(notOneStrInt8StrPtrErr, someLogicStrInt8StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt8PtrErr failed")
	}
	_, err = FilterMapStrInt8PtrErr(notOneStrInt8StrPtrErr, someLogicStrInt8StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrInt8PtrErr failed")
	}
}
func notOneStrInt8StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrInt8StrPtrErr(num *string) (*int8, error) {
	var r int8 = int8(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = int8(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrUintPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*uint{&vo10}
	newList, _ := FilterMapStrUintPtrErr(notOneStrUintStrPtrErr, someLogicStrUintStrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUintPtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUintPtrErr failed")
	}

	r, _ = FilterMapStrUintPtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUintPtrErr failed")
	}

	_, err := FilterMapStrUintPtrErr(notOneStrUintStrPtrErr, someLogicStrUintStrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrUintPtrErr failed")
	}
	_, err = FilterMapStrUintPtrErr(notOneStrUintStrPtrErr, someLogicStrUintStrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrUintPtrErr failed")
	}
}
func notOneStrUintStrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrUintStrPtrErr(num *string) (*uint, error) {
	var r uint = uint(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = uint(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrUint64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*uint64{&vo10}
	newList, _ := FilterMapStrUint64PtrErr(notOneStrUint64StrPtrErr, someLogicStrUint64StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint64PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint64PtrErr failed")
	}

	r, _ = FilterMapStrUint64PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint64PtrErr failed")
	}

	_, err := FilterMapStrUint64PtrErr(notOneStrUint64StrPtrErr, someLogicStrUint64StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint64PtrErr failed")
	}
	_, err = FilterMapStrUint64PtrErr(notOneStrUint64StrPtrErr, someLogicStrUint64StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint64PtrErr failed")
	}
}
func notOneStrUint64StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrUint64StrPtrErr(num *string) (*uint64, error) {
	var r uint64 = uint64(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = uint64(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrUint32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*uint32{&vo10}
	newList, _ := FilterMapStrUint32PtrErr(notOneStrUint32StrPtrErr, someLogicStrUint32StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint32PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint32PtrErr failed")
	}

	r, _ = FilterMapStrUint32PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint32PtrErr failed")
	}

	_, err := FilterMapStrUint32PtrErr(notOneStrUint32StrPtrErr, someLogicStrUint32StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint32PtrErr failed")
	}
	_, err = FilterMapStrUint32PtrErr(notOneStrUint32StrPtrErr, someLogicStrUint32StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint32PtrErr failed")
	}
}
func notOneStrUint32StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrUint32StrPtrErr(num *string) (*uint32, error) {
	var r uint32 = uint32(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = uint32(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrUint16PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*uint16{&vo10}
	newList, _ := FilterMapStrUint16PtrErr(notOneStrUint16StrPtrErr, someLogicStrUint16StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint16PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint16PtrErr failed")
	}

	r, _ = FilterMapStrUint16PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint16PtrErr failed")
	}

	_, err := FilterMapStrUint16PtrErr(notOneStrUint16StrPtrErr, someLogicStrUint16StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint16PtrErr failed")
	}
	_, err = FilterMapStrUint16PtrErr(notOneStrUint16StrPtrErr, someLogicStrUint16StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint16PtrErr failed")
	}
}
func notOneStrUint16StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrUint16StrPtrErr(num *string) (*uint16, error) {
	var r uint16 = uint16(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = uint16(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrUint8PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*uint8{&vo10}
	newList, _ := FilterMapStrUint8PtrErr(notOneStrUint8StrPtrErr, someLogicStrUint8StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrUint8PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrUint8PtrErr failed")
	}

	r, _ = FilterMapStrUint8PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrUint8PtrErr failed")
	}

	_, err := FilterMapStrUint8PtrErr(notOneStrUint8StrPtrErr, someLogicStrUint8StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint8PtrErr failed")
	}
	_, err = FilterMapStrUint8PtrErr(notOneStrUint8StrPtrErr, someLogicStrUint8StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrUint8PtrErr failed")
	}
}
func notOneStrUint8StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrUint8StrPtrErr(num *string) (*uint8, error) {
	var r uint8 = uint8(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = uint8(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrBoolPtrErr(t *testing.T) {
	// Test : someLogic

	var vto bool = true
	var vfo bool = false

	var vi1 string = "1"
	var vi2 string = "2"
	var vi3 string = "3"
	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapStrBoolPtrErr(notOneStrBoolPtrErr, someLogicStrBoolPtrErr, []*string{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapStrBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrBoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrBoolPtrErr failed")
	}

	r, _ = FilterMapStrBoolPtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrBoolPtrErr failed")
	}

	_, err := FilterMapStrBoolPtrErr(notOneStrBoolPtrErr, someLogicStrBoolPtrErr, []*string{&vi1, &vi10, &vi2})
	if err == nil {
		t.Errorf("FilterMapStrBoolPtrErr failed")
	}

	_, err = FilterMapStrBoolPtrErr(notOneStrBoolPtrErr, someLogicStrBoolPtrErr, []*string{&vi1, &vi10, &vi3})
	if err == nil {
		t.Errorf("FilterMapStrBoolPtrErr failed")
	}
}
func notOneStrBoolPtrErr(num *string) (bool, error) {
	if *num == "2" {
		return false, errors.New("2 is not valid value for this test")
	}
	return *num != "1", nil
}

func someLogicStrBoolPtrErr(num *string) (*bool, error) {
	if *num == "3" {
		return nil, errors.New("3 is not valid value for this test")
	}
	var t bool = true
	var f bool = false

	if *num == "10" {
		return &t, nil
	} 
	return &f, nil
}

func TestFilterMapStrFloat32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*float32{&vo10}
	newList, _ := FilterMapStrFloat32PtrErr(notOneStrFloat32StrPtrErr, someLogicStrFloat32StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrFloat32PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat32PtrErr failed")
	}

	r, _ = FilterMapStrFloat32PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat32PtrErr failed")
	}

	_, err := FilterMapStrFloat32PtrErr(notOneStrFloat32StrPtrErr, someLogicStrFloat32StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat32PtrErr failed")
	}
	_, err = FilterMapStrFloat32PtrErr(notOneStrFloat32StrPtrErr, someLogicStrFloat32StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat32PtrErr failed")
	}
}
func notOneStrFloat32StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrFloat32StrPtrErr(num *string) (*float32, error) {
	var r float32 = float32(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = float32(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapStrFloat64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vOne string = "one"
	var vTwo string = "two"
	var vThree string = "three"
	var vTen string = "ten"

	expectedList := []*float64{&vo10}
	newList, _ := FilterMapStrFloat64PtrErr(notOneStrFloat64StrPtrErr, someLogicStrFloat64StrPtrErr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapStrFloat64PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat64PtrErr failed")
	}

	r, _ = FilterMapStrFloat64PtrErr(nil, nil, []*string{})
	if len(r) > 0 {
		t.Errorf("FilterMapStrFloat64PtrErr failed")
	}

	_, err := FilterMapStrFloat64PtrErr(notOneStrFloat64StrPtrErr, someLogicStrFloat64StrPtrErr, []*string{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat64PtrErr failed")
	}
	_, err = FilterMapStrFloat64PtrErr(notOneStrFloat64StrPtrErr, someLogicStrFloat64StrPtrErr, []*string{&vThree})
	if err == nil {
		t.Errorf("FilterMapStrFloat64PtrErr failed")
	}
}
func notOneStrFloat64StrPtrErr(num *string) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogicStrFloat64StrPtrErr(num *string) (*float64, error) {
	var r float64 = float64(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = float64(10)
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapBoolIntPtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 int = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*int{&vo10, &vo10}
	newList, _ := FilterMapBoolIntPtrErr(notOneBoolIntPtrErr, someLogicBoolIntPtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolIntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolIntErr failed")
	}

	r, _ = FilterMapBoolIntPtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolIntPtrErr failed")
	}
	_, err := FilterMapBoolIntPtrErr(notOneBoolIntPtrErr, someLogicBoolIntPtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolIntPtrErr failed")
	}

	_, err = FilterMapBoolIntPtrErr(notOneBoolIntPtrErr2, someLogicBoolIntPtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolIntPtrErr failed")
	}
}
func notOneBoolIntPtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolIntPtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolIntPtrErr(num *bool) (*int, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 int = 10
	var v0 int = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolInt64PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 int64 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*int64{&vo10, &vo10}
	newList, _ := FilterMapBoolInt64PtrErr(notOneBoolInt64PtrErr, someLogicBoolInt64PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt64Err failed")
	}

	r, _ = FilterMapBoolInt64PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt64PtrErr failed")
	}
	_, err := FilterMapBoolInt64PtrErr(notOneBoolInt64PtrErr, someLogicBoolInt64PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolInt64PtrErr failed")
	}

	_, err = FilterMapBoolInt64PtrErr(notOneBoolInt64PtrErr2, someLogicBoolInt64PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt64PtrErr failed")
	}
}
func notOneBoolInt64PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolInt64PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolInt64PtrErr(num *bool) (*int64, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 int64 = 10
	var v0 int64 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolInt32PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 int32 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*int32{&vo10, &vo10}
	newList, _ := FilterMapBoolInt32PtrErr(notOneBoolInt32PtrErr, someLogicBoolInt32PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt32Err failed")
	}

	r, _ = FilterMapBoolInt32PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt32PtrErr failed")
	}
	_, err := FilterMapBoolInt32PtrErr(notOneBoolInt32PtrErr, someLogicBoolInt32PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolInt32PtrErr failed")
	}

	_, err = FilterMapBoolInt32PtrErr(notOneBoolInt32PtrErr2, someLogicBoolInt32PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt32PtrErr failed")
	}
}
func notOneBoolInt32PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolInt32PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolInt32PtrErr(num *bool) (*int32, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 int32 = 10
	var v0 int32 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolInt16PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 int16 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*int16{&vo10, &vo10}
	newList, _ := FilterMapBoolInt16PtrErr(notOneBoolInt16PtrErr, someLogicBoolInt16PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt16Err failed")
	}

	r, _ = FilterMapBoolInt16PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt16PtrErr failed")
	}
	_, err := FilterMapBoolInt16PtrErr(notOneBoolInt16PtrErr, someLogicBoolInt16PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolInt16PtrErr failed")
	}

	_, err = FilterMapBoolInt16PtrErr(notOneBoolInt16PtrErr2, someLogicBoolInt16PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt16PtrErr failed")
	}
}
func notOneBoolInt16PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolInt16PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolInt16PtrErr(num *bool) (*int16, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 int16 = 10
	var v0 int16 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolInt8PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 int8 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*int8{&vo10, &vo10}
	newList, _ := FilterMapBoolInt8PtrErr(notOneBoolInt8PtrErr, someLogicBoolInt8PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolInt8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt8Err failed")
	}

	r, _ = FilterMapBoolInt8PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolInt8PtrErr failed")
	}
	_, err := FilterMapBoolInt8PtrErr(notOneBoolInt8PtrErr, someLogicBoolInt8PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolInt8PtrErr failed")
	}

	_, err = FilterMapBoolInt8PtrErr(notOneBoolInt8PtrErr2, someLogicBoolInt8PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolInt8PtrErr failed")
	}
}
func notOneBoolInt8PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolInt8PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolInt8PtrErr(num *bool) (*int8, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 int8 = 10
	var v0 int8 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolUintPtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*uint{&vo10, &vo10}
	newList, _ := FilterMapBoolUintPtrErr(notOneBoolUintPtrErr, someLogicBoolUintPtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUintErr failed")
	}

	r, _ = FilterMapBoolUintPtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUintPtrErr failed")
	}
	_, err := FilterMapBoolUintPtrErr(notOneBoolUintPtrErr, someLogicBoolUintPtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolUintPtrErr failed")
	}

	_, err = FilterMapBoolUintPtrErr(notOneBoolUintPtrErr2, someLogicBoolUintPtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolUintPtrErr failed")
	}
}
func notOneBoolUintPtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolUintPtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolUintPtrErr(num *bool) (*uint, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 uint = 10
	var v0 uint = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolUint64PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint64 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*uint64{&vo10, &vo10}
	newList, _ := FilterMapBoolUint64PtrErr(notOneBoolUint64PtrErr, someLogicBoolUint64PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint64Err failed")
	}

	r, _ = FilterMapBoolUint64PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint64PtrErr failed")
	}
	_, err := FilterMapBoolUint64PtrErr(notOneBoolUint64PtrErr, someLogicBoolUint64PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolUint64PtrErr failed")
	}

	_, err = FilterMapBoolUint64PtrErr(notOneBoolUint64PtrErr2, someLogicBoolUint64PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint64PtrErr failed")
	}
}
func notOneBoolUint64PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolUint64PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolUint64PtrErr(num *bool) (*uint64, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 uint64 = 10
	var v0 uint64 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolUint32PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint32 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*uint32{&vo10, &vo10}
	newList, _ := FilterMapBoolUint32PtrErr(notOneBoolUint32PtrErr, someLogicBoolUint32PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint32Err failed")
	}

	r, _ = FilterMapBoolUint32PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint32PtrErr failed")
	}
	_, err := FilterMapBoolUint32PtrErr(notOneBoolUint32PtrErr, someLogicBoolUint32PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolUint32PtrErr failed")
	}

	_, err = FilterMapBoolUint32PtrErr(notOneBoolUint32PtrErr2, someLogicBoolUint32PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint32PtrErr failed")
	}
}
func notOneBoolUint32PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolUint32PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolUint32PtrErr(num *bool) (*uint32, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 uint32 = 10
	var v0 uint32 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolUint16PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint16 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*uint16{&vo10, &vo10}
	newList, _ := FilterMapBoolUint16PtrErr(notOneBoolUint16PtrErr, someLogicBoolUint16PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint16Err failed")
	}

	r, _ = FilterMapBoolUint16PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint16PtrErr failed")
	}
	_, err := FilterMapBoolUint16PtrErr(notOneBoolUint16PtrErr, someLogicBoolUint16PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolUint16PtrErr failed")
	}

	_, err = FilterMapBoolUint16PtrErr(notOneBoolUint16PtrErr2, someLogicBoolUint16PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint16PtrErr failed")
	}
}
func notOneBoolUint16PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolUint16PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolUint16PtrErr(num *bool) (*uint16, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 uint16 = 10
	var v0 uint16 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolUint8PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 uint8 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*uint8{&vo10, &vo10}
	newList, _ := FilterMapBoolUint8PtrErr(notOneBoolUint8PtrErr, someLogicBoolUint8PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolUint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint8Err failed")
	}

	r, _ = FilterMapBoolUint8PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolUint8PtrErr failed")
	}
	_, err := FilterMapBoolUint8PtrErr(notOneBoolUint8PtrErr, someLogicBoolUint8PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolUint8PtrErr failed")
	}

	_, err = FilterMapBoolUint8PtrErr(notOneBoolUint8PtrErr2, someLogicBoolUint8PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolUint8PtrErr failed")
	}
}
func notOneBoolUint8PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolUint8PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolUint8PtrErr(num *bool) (*uint8, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 uint8 = 10
	var v0 uint8 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	
	var vit bool = true
	var vif bool = false

	expectedList := []*string{&vo10, &vo10}
	newList, _ := FilterMapBoolStrPtrErr(notOneBoolStrPtrErr, someLogicBoolStrPtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolStrPtrErr failed")
	}

	r, _ = FilterMapBoolStrPtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolStrPtrErr failed")
	}
	_, err := FilterMapBoolStrPtrErr(notOneBoolStrPtrErr, someLogicBoolStrPtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolStrPtrErr failed")
	}
	_, err = FilterMapBoolStrPtrErr(notOneBoolStrPtrErr2, someLogicBoolStrPtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolStrPtrErr failed")
	}
}
func notOneBoolStrPtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolStrPtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolStrPtrErr(num *bool) (*string, error) {
	if *num == false {
		return nil, errors.New("false is error in this test")
	}
	var v10 string = "10"
	var v0 string = "0"

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolFloat32PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 float32 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*float32{&vo10, &vo10}
	newList, _ := FilterMapBoolFloat32PtrErr(notOneBoolFloat32PtrErr, someLogicBoolFloat32PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolFloat32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat32Err failed")
	}

	r, _ = FilterMapBoolFloat32PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat32PtrErr failed")
	}
	_, err := FilterMapBoolFloat32PtrErr(notOneBoolFloat32PtrErr, someLogicBoolFloat32PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolFloat32PtrErr failed")
	}

	_, err = FilterMapBoolFloat32PtrErr(notOneBoolFloat32PtrErr2, someLogicBoolFloat32PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat32PtrErr failed")
	}
}
func notOneBoolFloat32PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolFloat32PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolFloat32PtrErr(num *bool) (*float32, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 float32 = 10
	var v0 float32 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapBoolFloat64PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 float64 = 10
	
	var vit bool = true
	var vif bool = false

	expectedList := []*float64{&vo10, &vo10}
	newList, _ := FilterMapBoolFloat64PtrErr(notOneBoolFloat64PtrErr, someLogicBoolFloat64PtrErr, []*bool{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapBoolFloat64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat64Err failed")
	}

	r, _ = FilterMapBoolFloat64PtrErr(nil, nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("FilterMapBoolFloat64PtrErr failed")
	}
	_, err := FilterMapBoolFloat64PtrErr(notOneBoolFloat64PtrErr, someLogicBoolFloat64PtrErr, []*bool{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMapBoolFloat64PtrErr failed")
	}

	_, err = FilterMapBoolFloat64PtrErr(notOneBoolFloat64PtrErr2, someLogicBoolFloat64PtrErr, []*bool{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMapBoolFloat64PtrErr failed")
	}
}
func notOneBoolFloat64PtrErr(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOneBoolFloat64PtrErr2(num *bool) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogicBoolFloat64PtrErr(num *bool) (*float64, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 float64 = 10
	var v0 float64 = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}

func TestFilterMapFloat32IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapFloat32IntPtrErr(notOneFloat32IntPtrErr, plusOneFloat32IntPtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32IntPtrErr failed")
	}
	r, _ = FilterMapFloat32IntPtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32IntPtrErr failed")
	}

	_, err := FilterMapFloat32IntPtrErr(notOneFloat32IntPtrErr, plusOneFloat32IntPtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32IntPtrErr failed")
	}
	_, err = FilterMapFloat32IntPtrErr(notOneFloat32IntPtrErr, plusOneFloat32IntPtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32IntPtrErr failed")
	}
}
func notOneFloat32IntPtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32IntPtrErr(num *float32) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapFloat32Int64PtrErr(notOneFloat32Int64PtrErr, plusOneFloat32Int64PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int64PtrErr failed")
	}
	r, _ = FilterMapFloat32Int64PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int64PtrErr failed")
	}

	_, err := FilterMapFloat32Int64PtrErr(notOneFloat32Int64PtrErr, plusOneFloat32Int64PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int64PtrErr failed")
	}
	_, err = FilterMapFloat32Int64PtrErr(notOneFloat32Int64PtrErr, plusOneFloat32Int64PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int64PtrErr failed")
	}
}
func notOneFloat32Int64PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Int64PtrErr(num *float32) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapFloat32Int32PtrErr(notOneFloat32Int32PtrErr, plusOneFloat32Int32PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int32PtrErr failed")
	}
	r, _ = FilterMapFloat32Int32PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int32PtrErr failed")
	}

	_, err := FilterMapFloat32Int32PtrErr(notOneFloat32Int32PtrErr, plusOneFloat32Int32PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int32PtrErr failed")
	}
	_, err = FilterMapFloat32Int32PtrErr(notOneFloat32Int32PtrErr, plusOneFloat32Int32PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int32PtrErr failed")
	}
}
func notOneFloat32Int32PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Int32PtrErr(num *float32) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapFloat32Int16PtrErr(notOneFloat32Int16PtrErr, plusOneFloat32Int16PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int16PtrErr failed")
	}
	r, _ = FilterMapFloat32Int16PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int16PtrErr failed")
	}

	_, err := FilterMapFloat32Int16PtrErr(notOneFloat32Int16PtrErr, plusOneFloat32Int16PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int16PtrErr failed")
	}
	_, err = FilterMapFloat32Int16PtrErr(notOneFloat32Int16PtrErr, plusOneFloat32Int16PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int16PtrErr failed")
	}
}
func notOneFloat32Int16PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Int16PtrErr(num *float32) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapFloat32Int8PtrErr(notOneFloat32Int8PtrErr, plusOneFloat32Int8PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int8PtrErr failed")
	}
	r, _ = FilterMapFloat32Int8PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Int8PtrErr failed")
	}

	_, err := FilterMapFloat32Int8PtrErr(notOneFloat32Int8PtrErr, plusOneFloat32Int8PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int8PtrErr failed")
	}
	_, err = FilterMapFloat32Int8PtrErr(notOneFloat32Int8PtrErr, plusOneFloat32Int8PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Int8PtrErr failed")
	}
}
func notOneFloat32Int8PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Int8PtrErr(num *float32) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapFloat32UintPtrErr(notOneFloat32UintPtrErr, plusOneFloat32UintPtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32UintPtrErr failed")
	}
	r, _ = FilterMapFloat32UintPtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32UintPtrErr failed")
	}

	_, err := FilterMapFloat32UintPtrErr(notOneFloat32UintPtrErr, plusOneFloat32UintPtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32UintPtrErr failed")
	}
	_, err = FilterMapFloat32UintPtrErr(notOneFloat32UintPtrErr, plusOneFloat32UintPtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32UintPtrErr failed")
	}
}
func notOneFloat32UintPtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32UintPtrErr(num *float32) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapFloat32Uint64PtrErr(notOneFloat32Uint64PtrErr, plusOneFloat32Uint64PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed")
	}
	r, _ = FilterMapFloat32Uint64PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed")
	}

	_, err := FilterMapFloat32Uint64PtrErr(notOneFloat32Uint64PtrErr, plusOneFloat32Uint64PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed")
	}
	_, err = FilterMapFloat32Uint64PtrErr(notOneFloat32Uint64PtrErr, plusOneFloat32Uint64PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint64PtrErr failed")
	}
}
func notOneFloat32Uint64PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Uint64PtrErr(num *float32) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapFloat32Uint32PtrErr(notOneFloat32Uint32PtrErr, plusOneFloat32Uint32PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed")
	}
	r, _ = FilterMapFloat32Uint32PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed")
	}

	_, err := FilterMapFloat32Uint32PtrErr(notOneFloat32Uint32PtrErr, plusOneFloat32Uint32PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed")
	}
	_, err = FilterMapFloat32Uint32PtrErr(notOneFloat32Uint32PtrErr, plusOneFloat32Uint32PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint32PtrErr failed")
	}
}
func notOneFloat32Uint32PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Uint32PtrErr(num *float32) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapFloat32Uint16PtrErr(notOneFloat32Uint16PtrErr, plusOneFloat32Uint16PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed")
	}
	r, _ = FilterMapFloat32Uint16PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed")
	}

	_, err := FilterMapFloat32Uint16PtrErr(notOneFloat32Uint16PtrErr, plusOneFloat32Uint16PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed")
	}
	_, err = FilterMapFloat32Uint16PtrErr(notOneFloat32Uint16PtrErr, plusOneFloat32Uint16PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint16PtrErr failed")
	}
}
func notOneFloat32Uint16PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Uint16PtrErr(num *float32) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapFloat32Uint8PtrErr(notOneFloat32Uint8PtrErr, plusOneFloat32Uint8PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed")
	}
	r, _ = FilterMapFloat32Uint8PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed")
	}

	_, err := FilterMapFloat32Uint8PtrErr(notOneFloat32Uint8PtrErr, plusOneFloat32Uint8PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed")
	}
	_, err = FilterMapFloat32Uint8PtrErr(notOneFloat32Uint8PtrErr, plusOneFloat32Uint8PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Uint8PtrErr failed")
	}
}
func notOneFloat32Uint8PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Uint8PtrErr(num *float32) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapFloat32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float32 = 1
	var iv2 float32 = 2
	var iv3 float32 = 3
	var iv10 float32 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapFloat32StrPtrErr(notOneFloat32StrNumPtrErr, someLogicFloat32StrNumPtrErr, []*float32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapFloat32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32StrPtrErr failed")
	}

	r, _ = FilterMapFloat32StrPtrErr(nil, nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32StrPtrErr failed")
	}

	_, err := FilterMapFloat32StrPtrErr(notOneFloat32StrNumPtrErr, someLogicFloat32StrNumPtrErr, []*float32{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapFloat32StrPtrErr failed")
	}
	_, err = FilterMapFloat32StrPtrErr(notOneFloat32StrNumPtrErr, someLogicFloat32StrNumPtrErr, []*float32{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapFloat32StrPtrErr failed")
	}
}

func notOneFloat32StrNumPtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicFloat32StrNumPtrErr(num *float32) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapFloat32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	var vi10 float32 = 10
	var vi0 float32 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapFloat32BoolPtrErr(notOneFloat32BoolPtrErr, someLogicFloat32BoolPtrErr, []*float32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32BoolPtr failed")
	}

	r, _ = FilterMapFloat32BoolPtrErr(nil, nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32BoolPtrErr failed")
	}

	_, err := FilterMapFloat32BoolPtrErr(notOneFloat32BoolPtrErr, someLogicFloat32BoolPtrErr, []*float32{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapFloat32BoolPtrErr failed")
	}

	_, err = FilterMapFloat32BoolPtrErr(notOneFloat32BoolPtrErr, someLogicFloat32BoolPtrErr, []*float32{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapFloat32BoolPtrErr failed")
	}
}
func notOneFloat32BoolPtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicFloat32BoolPtrErr(num *float32) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapFloat32Float64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	var vo5 float64 = 5
	var vo6 float64 = 6

	expectedList := []*float64{&vo5, &vo6}
	newList, _ := FilterMapFloat32Float64PtrErr(notOneFloat32Float64PtrErr, plusOneFloat32Float64PtrErr, []*float32{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat32Float64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Float64PtrErr failed")
	}
	r, _ = FilterMapFloat32Float64PtrErr(nil, nil, []*float32{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat32Float64PtrErr failed")
	}

	_, err := FilterMapFloat32Float64PtrErr(notOneFloat32Float64PtrErr, plusOneFloat32Float64PtrErr, []*float32{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Float64PtrErr failed")
	}
	_, err = FilterMapFloat32Float64PtrErr(notOneFloat32Float64PtrErr, plusOneFloat32Float64PtrErr, []*float32{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat32Float64PtrErr failed")
	}
}
func notOneFloat32Float64PtrErr(num *float32) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat32Float64PtrErr(num *float32) (*float64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float64(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64IntPtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int = 5
	var vo6 int = 6

	expectedList := []*int{&vo5, &vo6}
	newList, _ := FilterMapFloat64IntPtrErr(notOneFloat64IntPtrErr, plusOneFloat64IntPtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64IntPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64IntPtrErr failed")
	}
	r, _ = FilterMapFloat64IntPtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64IntPtrErr failed")
	}

	_, err := FilterMapFloat64IntPtrErr(notOneFloat64IntPtrErr, plusOneFloat64IntPtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64IntPtrErr failed")
	}
	_, err = FilterMapFloat64IntPtrErr(notOneFloat64IntPtrErr, plusOneFloat64IntPtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64IntPtrErr failed")
	}
}
func notOneFloat64IntPtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64IntPtrErr(num *float64) (*int, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Int64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int64 = 5
	var vo6 int64 = 6

	expectedList := []*int64{&vo5, &vo6}
	newList, _ := FilterMapFloat64Int64PtrErr(notOneFloat64Int64PtrErr, plusOneFloat64Int64PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int64PtrErr failed")
	}
	r, _ = FilterMapFloat64Int64PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int64PtrErr failed")
	}

	_, err := FilterMapFloat64Int64PtrErr(notOneFloat64Int64PtrErr, plusOneFloat64Int64PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int64PtrErr failed")
	}
	_, err = FilterMapFloat64Int64PtrErr(notOneFloat64Int64PtrErr, plusOneFloat64Int64PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int64PtrErr failed")
	}
}
func notOneFloat64Int64PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Int64PtrErr(num *float64) (*int64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int64(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Int32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int32 = 5
	var vo6 int32 = 6

	expectedList := []*int32{&vo5, &vo6}
	newList, _ := FilterMapFloat64Int32PtrErr(notOneFloat64Int32PtrErr, plusOneFloat64Int32PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int32PtrErr failed")
	}
	r, _ = FilterMapFloat64Int32PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int32PtrErr failed")
	}

	_, err := FilterMapFloat64Int32PtrErr(notOneFloat64Int32PtrErr, plusOneFloat64Int32PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int32PtrErr failed")
	}
	_, err = FilterMapFloat64Int32PtrErr(notOneFloat64Int32PtrErr, plusOneFloat64Int32PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int32PtrErr failed")
	}
}
func notOneFloat64Int32PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Int32PtrErr(num *float64) (*int32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int32(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Int16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int16 = 5
	var vo6 int16 = 6

	expectedList := []*int16{&vo5, &vo6}
	newList, _ := FilterMapFloat64Int16PtrErr(notOneFloat64Int16PtrErr, plusOneFloat64Int16PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int16PtrErr failed")
	}
	r, _ = FilterMapFloat64Int16PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int16PtrErr failed")
	}

	_, err := FilterMapFloat64Int16PtrErr(notOneFloat64Int16PtrErr, plusOneFloat64Int16PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int16PtrErr failed")
	}
	_, err = FilterMapFloat64Int16PtrErr(notOneFloat64Int16PtrErr, plusOneFloat64Int16PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int16PtrErr failed")
	}
}
func notOneFloat64Int16PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Int16PtrErr(num *float64) (*int16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int16(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Int8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 int8 = 5
	var vo6 int8 = 6

	expectedList := []*int8{&vo5, &vo6}
	newList, _ := FilterMapFloat64Int8PtrErr(notOneFloat64Int8PtrErr, plusOneFloat64Int8PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Int8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int8PtrErr failed")
	}
	r, _ = FilterMapFloat64Int8PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Int8PtrErr failed")
	}

	_, err := FilterMapFloat64Int8PtrErr(notOneFloat64Int8PtrErr, plusOneFloat64Int8PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int8PtrErr failed")
	}
	_, err = FilterMapFloat64Int8PtrErr(notOneFloat64Int8PtrErr, plusOneFloat64Int8PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Int8PtrErr failed")
	}
}
func notOneFloat64Int8PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Int8PtrErr(num *float64) (*int8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := int8(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64UintPtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint = 5
	var vo6 uint = 6

	expectedList := []*uint{&vo5, &vo6}
	newList, _ := FilterMapFloat64UintPtrErr(notOneFloat64UintPtrErr, plusOneFloat64UintPtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64UintPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64UintPtrErr failed")
	}
	r, _ = FilterMapFloat64UintPtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64UintPtrErr failed")
	}

	_, err := FilterMapFloat64UintPtrErr(notOneFloat64UintPtrErr, plusOneFloat64UintPtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64UintPtrErr failed")
	}
	_, err = FilterMapFloat64UintPtrErr(notOneFloat64UintPtrErr, plusOneFloat64UintPtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64UintPtrErr failed")
	}
}
func notOneFloat64UintPtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64UintPtrErr(num *float64) (*uint, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Uint64PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint64 = 5
	var vo6 uint64 = 6

	expectedList := []*uint64{&vo5, &vo6}
	newList, _ := FilterMapFloat64Uint64PtrErr(notOneFloat64Uint64PtrErr, plusOneFloat64Uint64PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint64PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed")
	}
	r, _ = FilterMapFloat64Uint64PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed")
	}

	_, err := FilterMapFloat64Uint64PtrErr(notOneFloat64Uint64PtrErr, plusOneFloat64Uint64PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed")
	}
	_, err = FilterMapFloat64Uint64PtrErr(notOneFloat64Uint64PtrErr, plusOneFloat64Uint64PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint64PtrErr failed")
	}
}
func notOneFloat64Uint64PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Uint64PtrErr(num *float64) (*uint64, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint64(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Uint32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint32 = 5
	var vo6 uint32 = 6

	expectedList := []*uint32{&vo5, &vo6}
	newList, _ := FilterMapFloat64Uint32PtrErr(notOneFloat64Uint32PtrErr, plusOneFloat64Uint32PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed")
	}
	r, _ = FilterMapFloat64Uint32PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed")
	}

	_, err := FilterMapFloat64Uint32PtrErr(notOneFloat64Uint32PtrErr, plusOneFloat64Uint32PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed")
	}
	_, err = FilterMapFloat64Uint32PtrErr(notOneFloat64Uint32PtrErr, plusOneFloat64Uint32PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint32PtrErr failed")
	}
}
func notOneFloat64Uint32PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Uint32PtrErr(num *float64) (*uint32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint32(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Uint16PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint16 = 5
	var vo6 uint16 = 6

	expectedList := []*uint16{&vo5, &vo6}
	newList, _ := FilterMapFloat64Uint16PtrErr(notOneFloat64Uint16PtrErr, plusOneFloat64Uint16PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint16PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed")
	}
	r, _ = FilterMapFloat64Uint16PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed")
	}

	_, err := FilterMapFloat64Uint16PtrErr(notOneFloat64Uint16PtrErr, plusOneFloat64Uint16PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed")
	}
	_, err = FilterMapFloat64Uint16PtrErr(notOneFloat64Uint16PtrErr, plusOneFloat64Uint16PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint16PtrErr failed")
	}
}
func notOneFloat64Uint16PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Uint16PtrErr(num *float64) (*uint16, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint16(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64Uint8PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 uint8 = 5
	var vo6 uint8 = 6

	expectedList := []*uint8{&vo5, &vo6}
	newList, _ := FilterMapFloat64Uint8PtrErr(notOneFloat64Uint8PtrErr, plusOneFloat64Uint8PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Uint8PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed")
	}
	r, _ = FilterMapFloat64Uint8PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed")
	}

	_, err := FilterMapFloat64Uint8PtrErr(notOneFloat64Uint8PtrErr, plusOneFloat64Uint8PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed")
	}
	_, err = FilterMapFloat64Uint8PtrErr(notOneFloat64Uint8PtrErr, plusOneFloat64Uint8PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Uint8PtrErr failed")
	}
}
func notOneFloat64Uint8PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Uint8PtrErr(num *float64) (*uint8, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := uint8(*num + 1)
	return &c, nil
}

func TestFilterMapFloat64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float64 = 1
	var iv2 float64 = 2
	var iv3 float64 = 3
	var iv10 float64 = 10
	expectedList := []*string{&ov10}
	newList, _ := FilterMapFloat64StrPtrErr(notOneFloat64StrNumPtrErr, someLogicFloat64StrNumPtrErr, []*float64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapFloat64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64StrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64StrPtrErr failed")
	}

	r, _ = FilterMapFloat64StrPtrErr(nil, nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64StrPtrErr failed")
	}

	_, err := FilterMapFloat64StrPtrErr(notOneFloat64StrNumPtrErr, someLogicFloat64StrNumPtrErr, []*float64{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMapFloat64StrPtrErr failed")
	}
	_, err = FilterMapFloat64StrPtrErr(notOneFloat64StrNumPtrErr, someLogicFloat64StrNumPtrErr, []*float64{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMapFloat64StrPtrErr failed")
	}
}

func notOneFloat64StrNumPtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogicFloat64StrNumPtrErr(num *float64) (*string, error) {
	var r string = "0"
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}

func TestFilterMapFloat64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	var vi10 float64 = 10
	var vi0 float64 = 0

	expectedList := []*bool{&vto, &vfo}
	newList, _ := FilterMapFloat64BoolPtrErr(notOneFloat64BoolPtrErr, someLogicFloat64BoolPtrErr, []*float64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64BoolPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64BoolPtr failed")
	}

	r, _ = FilterMapFloat64BoolPtrErr(nil, nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64BoolPtrErr failed")
	}

	_, err := FilterMapFloat64BoolPtrErr(notOneFloat64BoolPtrErr, someLogicFloat64BoolPtrErr, []*float64{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapFloat64BoolPtrErr failed")
	}

	_, err = FilterMapFloat64BoolPtrErr(notOneFloat64BoolPtrErr, someLogicFloat64BoolPtrErr, []*float64{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMapFloat64BoolPtrErr failed")
	}
}
func notOneFloat64BoolPtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogicFloat64BoolPtrErr(num *float64) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}

func TestFilterMapFloat64Float32PtrErr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	var vo5 float32 = 5
	var vo6 float32 = 6

	expectedList := []*float32{&vo5, &vo6}
	newList, _ := FilterMapFloat64Float32PtrErr(notOneFloat64Float32PtrErr, plusOneFloat64Float32PtrErr, []*float64{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMapFloat64Float32PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Float32PtrErr failed")
	}
	r, _ = FilterMapFloat64Float32PtrErr(nil, nil, []*float64{})

	if len(r) > 0 {
		t.Errorf("FilterMapFloat64Float32PtrErr failed")
	}

	_, err := FilterMapFloat64Float32PtrErr(notOneFloat64Float32PtrErr, plusOneFloat64Float32PtrErr, []*float64{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Float32PtrErr failed")
	}
	_, err = FilterMapFloat64Float32PtrErr(notOneFloat64Float32PtrErr, plusOneFloat64Float32PtrErr, []*float64{&v3})
	if err == nil {
		t.Errorf("FilterMapFloat64Float32PtrErr failed")
	}
}
func notOneFloat64Float32PtrErr(num *float64) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOneFloat64Float32PtrErr(num *float64) (*float32, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := float32(*num + 1)
	return &c, nil
}
