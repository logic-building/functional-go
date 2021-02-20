package fp

import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntPtrErr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v5 int = 5
	var v6 int = 6
	var v7 int = 7
	var v8 int = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int{&v6, &v7, &v8}
	partialAddIntPtr := func(num *int) (*int, error) {
		r, err := addIntPtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapIntPtrErr(partialAddIntPtr, []*int{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapIntPtrErr failed")
	}

	r, _ := MapIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntPtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int = 0
	partialAddIntPtr2 := func(num *int) (*int, error) {
		r, err := addIntPtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapIntPtrErr(partialAddIntPtr2, []*int{&v0})
	if err == nil {
		t.Errorf("MapIntPtrErr failed")
	}

	
}

func addIntPtrErr(num1, num2 *int) (*int, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapInt64PtrErr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v5 int64 = 5
	var v6 int64 = 6
	var v7 int64 = 7
	var v8 int64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int64{&v6, &v7, &v8}
	partialAddInt64Ptr := func(num *int64) (*int64, error) {
		r, err := addInt64PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapInt64PtrErr(partialAddInt64Ptr, []*int64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt64PtrErr failed")
	}

	r, _ := MapInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int64 = 0
	partialAddInt64Ptr2 := func(num *int64) (*int64, error) {
		r, err := addInt64PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapInt64PtrErr(partialAddInt64Ptr2, []*int64{&v0})
	if err == nil {
		t.Errorf("MapInt64PtrErr failed")
	}

	
}

func addInt64PtrErr(num1, num2 *int64) (*int64, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapInt32PtrErr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v5 int32 = 5
	var v6 int32 = 6
	var v7 int32 = 7
	var v8 int32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int32{&v6, &v7, &v8}
	partialAddInt32Ptr := func(num *int32) (*int32, error) {
		r, err := addInt32PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapInt32PtrErr(partialAddInt32Ptr, []*int32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt32PtrErr failed")
	}

	r, _ := MapInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int32 = 0
	partialAddInt32Ptr2 := func(num *int32) (*int32, error) {
		r, err := addInt32PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapInt32PtrErr(partialAddInt32Ptr2, []*int32{&v0})
	if err == nil {
		t.Errorf("MapInt32PtrErr failed")
	}

	
}

func addInt32PtrErr(num1, num2 *int32) (*int32, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapInt16PtrErr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v5 int16 = 5
	var v6 int16 = 6
	var v7 int16 = 7
	var v8 int16 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int16{&v6, &v7, &v8}
	partialAddInt16Ptr := func(num *int16) (*int16, error) {
		r, err := addInt16PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapInt16PtrErr(partialAddInt16Ptr, []*int16{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt16PtrErr failed")
	}

	r, _ := MapInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int16 = 0
	partialAddInt16Ptr2 := func(num *int16) (*int16, error) {
		r, err := addInt16PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapInt16PtrErr(partialAddInt16Ptr2, []*int16{&v0})
	if err == nil {
		t.Errorf("MapInt16PtrErr failed")
	}

	
}

func addInt16PtrErr(num1, num2 *int16) (*int16, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapInt8PtrErr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v5 int8 = 5
	var v6 int8 = 6
	var v7 int8 = 7
	var v8 int8 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int8{&v6, &v7, &v8}
	partialAddInt8Ptr := func(num *int8) (*int8, error) {
		r, err := addInt8PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapInt8PtrErr(partialAddInt8Ptr, []*int8{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt8PtrErr failed")
	}

	r, _ := MapInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 int8 = 0
	partialAddInt8Ptr2 := func(num *int8) (*int8, error) {
		r, err := addInt8PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapInt8PtrErr(partialAddInt8Ptr2, []*int8{&v0})
	if err == nil {
		t.Errorf("MapInt8PtrErr failed")
	}

	
}

func addInt8PtrErr(num1, num2 *int8) (*int8, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapUintPtrErr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v5 uint = 5
	var v6 uint = 6
	var v7 uint = 7
	var v8 uint = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint{&v6, &v7, &v8}
	partialAddUintPtr := func(num *uint) (*uint, error) {
		r, err := addUintPtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapUintPtrErr(partialAddUintPtr, []*uint{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUintPtrErr failed")
	}

	r, _ := MapUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintPtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint = 0
	partialAddUintPtr2 := func(num *uint) (*uint, error) {
		r, err := addUintPtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapUintPtrErr(partialAddUintPtr2, []*uint{&v0})
	if err == nil {
		t.Errorf("MapUintPtrErr failed")
	}

	
}

func addUintPtrErr(num1, num2 *uint) (*uint, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapUint64PtrErr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v5 uint64 = 5
	var v6 uint64 = 6
	var v7 uint64 = 7
	var v8 uint64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint64{&v6, &v7, &v8}
	partialAddUint64Ptr := func(num *uint64) (*uint64, error) {
		r, err := addUint64PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapUint64PtrErr(partialAddUint64Ptr, []*uint64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint64PtrErr failed")
	}

	r, _ := MapUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint64 = 0
	partialAddUint64Ptr2 := func(num *uint64) (*uint64, error) {
		r, err := addUint64PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapUint64PtrErr(partialAddUint64Ptr2, []*uint64{&v0})
	if err == nil {
		t.Errorf("MapUint64PtrErr failed")
	}

	
}

func addUint64PtrErr(num1, num2 *uint64) (*uint64, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapUint32PtrErr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v5 uint32 = 5
	var v6 uint32 = 6
	var v7 uint32 = 7
	var v8 uint32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint32{&v6, &v7, &v8}
	partialAddUint32Ptr := func(num *uint32) (*uint32, error) {
		r, err := addUint32PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapUint32PtrErr(partialAddUint32Ptr, []*uint32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint32PtrErr failed")
	}

	r, _ := MapUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint32 = 0
	partialAddUint32Ptr2 := func(num *uint32) (*uint32, error) {
		r, err := addUint32PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapUint32PtrErr(partialAddUint32Ptr2, []*uint32{&v0})
	if err == nil {
		t.Errorf("MapUint32PtrErr failed")
	}

	
}

func addUint32PtrErr(num1, num2 *uint32) (*uint32, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapUint16PtrErr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v5 uint16 = 5
	var v6 uint16 = 6
	var v7 uint16 = 7
	var v8 uint16 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint16{&v6, &v7, &v8}
	partialAddUint16Ptr := func(num *uint16) (*uint16, error) {
		r, err := addUint16PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapUint16PtrErr(partialAddUint16Ptr, []*uint16{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint16PtrErr failed")
	}

	r, _ := MapUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint16 = 0
	partialAddUint16Ptr2 := func(num *uint16) (*uint16, error) {
		r, err := addUint16PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapUint16PtrErr(partialAddUint16Ptr2, []*uint16{&v0})
	if err == nil {
		t.Errorf("MapUint16PtrErr failed")
	}

	
}

func addUint16PtrErr(num1, num2 *uint16) (*uint16, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapUint8PtrErr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v5 uint8 = 5
	var v6 uint8 = 6
	var v7 uint8 = 7
	var v8 uint8 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint8{&v6, &v7, &v8}
	partialAddUint8Ptr := func(num *uint8) (*uint8, error) {
		r, err := addUint8PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapUint8PtrErr(partialAddUint8Ptr, []*uint8{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint8PtrErr failed")
	}

	r, _ := MapUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 uint8 = 0
	partialAddUint8Ptr2 := func(num *uint8) (*uint8, error) {
		r, err := addUint8PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapUint8PtrErr(partialAddUint8Ptr2, []*uint8{&v0})
	if err == nil {
		t.Errorf("MapUint8PtrErr failed")
	}

	
}

func addUint8PtrErr(num1, num2 *uint8) (*uint8, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) (*string, error) {
		r, err := addStrPtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapStrPtrErr(partialAddStrPtr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapStrPtrErr failed")
	}

	r, _ := MapStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrPtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 string = "0"
	partialAddStrPtr2 := func(num *string) (*string, error) {
		r, err := addStrPtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapStrPtrErr(partialAddStrPtr2, []*string{&v0})
	if err == nil {
		t.Errorf("MapStrPtrErr failed")
	}

	
}

func addStrPtrErr(num1, num2 *string) (*string, error) {
	if *num1 == "0" {
		return nil, errors.New("0 is not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapBoolPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []*bool{&vf}
	
	newList, _ := MapBoolPtrErr(inverseBoolPtrErr, []*bool{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("MapBoolPtrErr failed")
	}

	_, err1 := MapBoolPtrErr(inverseBoolPtrErr, []*bool{&vf})
	if err1 == nil  {
		t.Errorf("MapBoolPtrErr failed")
	}

	r, _ := MapBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolPtrErr failed.")
	}
}

func inverseBoolPtrErr(num1 *bool) (*bool, error) {
	if *num1 == false {
		return nil, errors.New("False is not allowed")
	}
	vt := true
    if *num1 == true {
		v := false
		return &v, nil
	} 
	return &vt, nil
}


func TestMapFloat32PtrErr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v5 float32 = 5
	var v6 float32 = 6
	var v7 float32 = 7
	var v8 float32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*float32{&v6, &v7, &v8}
	partialAddFloat32Ptr := func(num *float32) (*float32, error) {
		r, err := addFloat32PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapFloat32PtrErr(partialAddFloat32Ptr, []*float32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapFloat32PtrErr failed")
	}

	r, _ := MapFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 float32 = 0
	partialAddFloat32Ptr2 := func(num *float32) (*float32, error) {
		r, err := addFloat32PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapFloat32PtrErr(partialAddFloat32Ptr2, []*float32{&v0})
	if err == nil {
		t.Errorf("MapFloat32PtrErr failed")
	}

	
}

func addFloat32PtrErr(num1, num2 *float32) (*float32, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}


func TestMapFloat64PtrErr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v5 float64 = 5
	var v6 float64 = 6
	var v7 float64 = 7
	var v8 float64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*float64{&v6, &v7, &v8}
	partialAddFloat64Ptr := func(num *float64) (*float64, error) {
		r, err := addFloat64PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := MapFloat64PtrErr(partialAddFloat64Ptr, []*float64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapFloat64PtrErr failed")
	}

	r, _ := MapFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 float64 = 0
	partialAddFloat64Ptr2 := func(num *float64) (*float64, error) {
		r, err := addFloat64PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := MapFloat64PtrErr(partialAddFloat64Ptr2, []*float64{&v0})
	if err == nil {
		t.Errorf("MapFloat64PtrErr failed")
	}

	
}

func addFloat64PtrErr(num1, num2 *float64) (*float64, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}

