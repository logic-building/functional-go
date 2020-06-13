package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v5 int = 5
	var v6 int = 6
	var v7 int = 7
	var v8 int = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int{&v6, &v7, &v8}
	partialAddIntPtr := func(num *int) *int {
		return addIntPtr(&v5, num)
	}
	sumList := MapIntPtr(partialAddIntPtr, []*int{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapIntPtr failed")
	}

	if len(MapIntPtr(nil, nil)) > 0 {
		t.Errorf("MapIntPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addIntPtr(num1, num2 *int) *int {
    result := *num1 + *num2
	return &result
}


func TestMapInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v5 int64 = 5
	var v6 int64 = 6
	var v7 int64 = 7
	var v8 int64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int64{&v6, &v7, &v8}
	partialAddInt64Ptr := func(num *int64) *int64 {
		return addInt64Ptr(&v5, num)
	}
	sumList := MapInt64Ptr(partialAddInt64Ptr, []*int64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt64Ptr failed")
	}

	if len(MapInt64Ptr(nil, nil)) > 0 {
		t.Errorf("MapInt64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addInt64Ptr(num1, num2 *int64) *int64 {
    result := *num1 + *num2
	return &result
}


func TestMapInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v5 int32 = 5
	var v6 int32 = 6
	var v7 int32 = 7
	var v8 int32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int32{&v6, &v7, &v8}
	partialAddInt32Ptr := func(num *int32) *int32 {
		return addInt32Ptr(&v5, num)
	}
	sumList := MapInt32Ptr(partialAddInt32Ptr, []*int32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt32Ptr failed")
	}

	if len(MapInt32Ptr(nil, nil)) > 0 {
		t.Errorf("MapInt32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addInt32Ptr(num1, num2 *int32) *int32 {
    result := *num1 + *num2
	return &result
}


func TestMapInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v5 int16 = 5
	var v6 int16 = 6
	var v7 int16 = 7
	var v8 int16 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int16{&v6, &v7, &v8}
	partialAddInt16Ptr := func(num *int16) *int16 {
		return addInt16Ptr(&v5, num)
	}
	sumList := MapInt16Ptr(partialAddInt16Ptr, []*int16{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt16Ptr failed")
	}

	if len(MapInt16Ptr(nil, nil)) > 0 {
		t.Errorf("MapInt16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addInt16Ptr(num1, num2 *int16) *int16 {
    result := *num1 + *num2
	return &result
}


func TestMapInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v5 int8 = 5
	var v6 int8 = 6
	var v7 int8 = 7
	var v8 int8 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*int8{&v6, &v7, &v8}
	partialAddInt8Ptr := func(num *int8) *int8 {
		return addInt8Ptr(&v5, num)
	}
	sumList := MapInt8Ptr(partialAddInt8Ptr, []*int8{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapInt8Ptr failed")
	}

	if len(MapInt8Ptr(nil, nil)) > 0 {
		t.Errorf("MapInt8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addInt8Ptr(num1, num2 *int8) *int8 {
    result := *num1 + *num2
	return &result
}


func TestMapUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v5 uint = 5
	var v6 uint = 6
	var v7 uint = 7
	var v8 uint = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint{&v6, &v7, &v8}
	partialAddUintPtr := func(num *uint) *uint {
		return addUintPtr(&v5, num)
	}
	sumList := MapUintPtr(partialAddUintPtr, []*uint{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUintPtr failed")
	}

	if len(MapUintPtr(nil, nil)) > 0 {
		t.Errorf("MapUintPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addUintPtr(num1, num2 *uint) *uint {
    result := *num1 + *num2
	return &result
}


func TestMapUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v5 uint64 = 5
	var v6 uint64 = 6
	var v7 uint64 = 7
	var v8 uint64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint64{&v6, &v7, &v8}
	partialAddUint64Ptr := func(num *uint64) *uint64 {
		return addUint64Ptr(&v5, num)
	}
	sumList := MapUint64Ptr(partialAddUint64Ptr, []*uint64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint64Ptr failed")
	}

	if len(MapUint64Ptr(nil, nil)) > 0 {
		t.Errorf("MapUint64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addUint64Ptr(num1, num2 *uint64) *uint64 {
    result := *num1 + *num2
	return &result
}


func TestMapUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v5 uint32 = 5
	var v6 uint32 = 6
	var v7 uint32 = 7
	var v8 uint32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint32{&v6, &v7, &v8}
	partialAddUint32Ptr := func(num *uint32) *uint32 {
		return addUint32Ptr(&v5, num)
	}
	sumList := MapUint32Ptr(partialAddUint32Ptr, []*uint32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint32Ptr failed")
	}

	if len(MapUint32Ptr(nil, nil)) > 0 {
		t.Errorf("MapUint32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addUint32Ptr(num1, num2 *uint32) *uint32 {
    result := *num1 + *num2
	return &result
}


func TestMapUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v5 uint16 = 5
	var v6 uint16 = 6
	var v7 uint16 = 7
	var v8 uint16 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint16{&v6, &v7, &v8}
	partialAddUint16Ptr := func(num *uint16) *uint16 {
		return addUint16Ptr(&v5, num)
	}
	sumList := MapUint16Ptr(partialAddUint16Ptr, []*uint16{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint16Ptr failed")
	}

	if len(MapUint16Ptr(nil, nil)) > 0 {
		t.Errorf("MapUint16Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addUint16Ptr(num1, num2 *uint16) *uint16 {
    result := *num1 + *num2
	return &result
}


func TestMapUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v5 uint8 = 5
	var v6 uint8 = 6
	var v7 uint8 = 7
	var v8 uint8 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*uint8{&v6, &v7, &v8}
	partialAddUint8Ptr := func(num *uint8) *uint8 {
		return addUint8Ptr(&v5, num)
	}
	sumList := MapUint8Ptr(partialAddUint8Ptr, []*uint8{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapUint8Ptr failed")
	}

	if len(MapUint8Ptr(nil, nil)) > 0 {
		t.Errorf("MapUint8Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addUint8Ptr(num1, num2 *uint8) *uint8 {
    result := *num1 + *num2
	return &result
}


func TestMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		return addStrPtr(&v5, num)
	}
	sumList := MapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapStrPtr failed")
	}

	if len(MapStrPtr(nil, nil)) > 0 {
		t.Errorf("MapStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addStrPtr(num1, num2 *string) *string {
    result := *num1 + *num2
	return &result
}


func TestMapBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []*bool{&vf}
	
	newList := MapBoolPtr(inverseBoolPtr, []*bool{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("MapBoolPtr failed")
	}

	if len(MapBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

func inverseBoolPtr(num1 *bool) *bool {
	vt := true
    if *num1 == true {
		v := false
		return &v
	} 
	return &vt
}


func TestMapFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v5 float32 = 5
	var v6 float32 = 6
	var v7 float32 = 7
	var v8 float32 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*float32{&v6, &v7, &v8}
	partialAddFloat32Ptr := func(num *float32) *float32 {
		return addFloat32Ptr(&v5, num)
	}
	sumList := MapFloat32Ptr(partialAddFloat32Ptr, []*float32{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapFloat32Ptr failed")
	}

	if len(MapFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("MapFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addFloat32Ptr(num1, num2 *float32) *float32 {
    result := *num1 + *num2
	return &result
}


func TestMapFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v5 float64 = 5
	var v6 float64 = 6
	var v7 float64 = 7
	var v8 float64 = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*float64{&v6, &v7, &v8}
	partialAddFloat64Ptr := func(num *float64) *float64 {
		return addFloat64Ptr(&v5, num)
	}
	sumList := MapFloat64Ptr(partialAddFloat64Ptr, []*float64{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapFloat64Ptr failed")
	}

	if len(MapFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("MapFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func addFloat64Ptr(num1, num2 *float64) *float64 {
    result := *num1 + *num2
	return &result
}

