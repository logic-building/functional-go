package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntInt64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapIntInt64Ptr(plusOneIntInt64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt64(nil, nil)) > 0 {
		t.Errorf("MapIntInt64 failed")
	}

	if len(MapIntInt64Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapIntInt32Ptr(plusOneIntInt32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt32(nil, nil)) > 0 {
		t.Errorf("MapIntInt32 failed")
	}

	if len(MapIntInt32Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapIntInt16Ptr(plusOneIntInt16Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt16(nil, nil)) > 0 {
		t.Errorf("MapIntInt16 failed")
	}

	if len(MapIntInt16Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapIntInt8Ptr(plusOneIntInt8Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntInt8(nil, nil)) > 0 {
		t.Errorf("MapIntInt8 failed")
	}

	if len(MapIntInt8Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapIntUintPtr(plusOneIntUintPtr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint(nil, nil)) > 0 {
		t.Errorf("MapIntUint failed")
	}

	if len(MapIntUintPtr(nil, []*int{})) > 0 {
		t.Errorf("MapIntUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapIntUint64Ptr(plusOneIntUint64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint64(nil, nil)) > 0 {
		t.Errorf("MapIntUint64 failed")
	}

	if len(MapIntUint64Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapIntUint32Ptr(plusOneIntUint32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint32(nil, nil)) > 0 {
		t.Errorf("MapIntUint32 failed")
	}

	if len(MapIntUint32Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapIntUint16Ptr(plusOneIntUint16Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint16(nil, nil)) > 0 {
		t.Errorf("MapIntUint16 failed")
	}

	if len(MapIntUint16Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapIntUint8Ptr(plusOneIntUint8Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntUint8(nil, nil)) > 0 {
		t.Errorf("MapIntUint8 failed")
	}

	if len(MapIntUint8Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntStrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int = 10

	expectedList := []*string{&vo10}
	newList := MapIntStrPtr(someLogicIntStrPtr, []*int{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapIntStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntStrPtr(nil, nil)) > 0 {
		t.Errorf("MapIntStr failed")
	}

	if len(MapIntStrPtr(nil, []*int{})) > 0 {
		t.Errorf("MapIntStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicIntStrPtr(num *int) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapIntBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int

	expectedList := []*bool{&vt, &vf}
	newList := MapIntBoolPtr(someLogicIntBoolPtr, []*int{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapIntBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapIntBool failed")
	}

	if len(MapIntBoolPtr(nil, []*int{})) > 0 {
		t.Errorf("MapIntBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapIntFloat32Ptr(plusOneIntFloat32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntFloat32(nil, nil)) > 0 {
		t.Errorf("MapIntFloat32 failed")
	}

	if len(MapIntFloat32Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapIntFloat64Ptr(plusOneIntFloat64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapIntFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapIntFloat64(nil, nil)) > 0 {
		t.Errorf("MapIntFloat64 failed")
	}

	if len(MapIntFloat64Ptr(nil, []*int{})) > 0 {
		t.Errorf("MapIntFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapInt64IntPtr(plusOneInt64IntPtr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int(nil, nil)) > 0 {
		t.Errorf("MapInt64Int failed")
	}

	if len(MapInt64IntPtr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapInt64Int32Ptr(plusOneInt64Int32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int32(nil, nil)) > 0 {
		t.Errorf("MapInt64Int32 failed")
	}

	if len(MapInt64Int32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapInt64Int16Ptr(plusOneInt64Int16Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int16(nil, nil)) > 0 {
		t.Errorf("MapInt64Int16 failed")
	}

	if len(MapInt64Int16Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapInt64Int8Ptr(plusOneInt64Int8Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Int8(nil, nil)) > 0 {
		t.Errorf("MapInt64Int8 failed")
	}

	if len(MapInt64Int8Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapInt64UintPtr(plusOneInt64UintPtr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint failed")
	}

	if len(MapInt64UintPtr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapInt64Uint64Ptr(plusOneInt64Uint64Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint64 failed")
	}

	if len(MapInt64Uint64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapInt64Uint32Ptr(plusOneInt64Uint32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint32 failed")
	}

	if len(MapInt64Uint32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapInt64Uint16Ptr(plusOneInt64Uint16Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint16 failed")
	}

	if len(MapInt64Uint16Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapInt64Uint8Ptr(plusOneInt64Uint8Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt64Uint8 failed")
	}

	if len(MapInt64Uint8Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int64 = 10

	expectedList := []*string{&vo10}
	newList := MapInt64StrPtr(someLogicInt64StrPtr, []*int64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64StrPtr(nil, nil)) > 0 {
		t.Errorf("MapInt64Str failed")
	}

	if len(MapInt64StrPtr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt64StrPtr(num *int64) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapInt64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64

	expectedList := []*bool{&vt, &vf}
	newList := MapInt64BoolPtr(someLogicInt64BoolPtr, []*int64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapInt64Bool failed")
	}

	if len(MapInt64BoolPtr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapInt64Float32Ptr(plusOneInt64Float32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Float32(nil, nil)) > 0 {
		t.Errorf("MapInt64Float32 failed")
	}

	if len(MapInt64Float32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapInt64Float64Ptr(plusOneInt64Float64Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt64Float64(nil, nil)) > 0 {
		t.Errorf("MapInt64Float64 failed")
	}

	if len(MapInt64Float64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("MapInt64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapInt32IntPtr(plusOneInt32IntPtr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int(nil, nil)) > 0 {
		t.Errorf("MapInt32Int failed")
	}

	if len(MapInt32IntPtr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapInt32Int64Ptr(plusOneInt32Int64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int64(nil, nil)) > 0 {
		t.Errorf("MapInt32Int64 failed")
	}

	if len(MapInt32Int64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapInt32Int16Ptr(plusOneInt32Int16Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int16(nil, nil)) > 0 {
		t.Errorf("MapInt32Int16 failed")
	}

	if len(MapInt32Int16Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapInt32Int8Ptr(plusOneInt32Int8Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Int8(nil, nil)) > 0 {
		t.Errorf("MapInt32Int8 failed")
	}

	if len(MapInt32Int8Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapInt32UintPtr(plusOneInt32UintPtr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint failed")
	}

	if len(MapInt32UintPtr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapInt32Uint64Ptr(plusOneInt32Uint64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint64 failed")
	}

	if len(MapInt32Uint64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapInt32Uint32Ptr(plusOneInt32Uint32Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint32 failed")
	}

	if len(MapInt32Uint32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapInt32Uint16Ptr(plusOneInt32Uint16Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint16 failed")
	}

	if len(MapInt32Uint16Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapInt32Uint8Ptr(plusOneInt32Uint8Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt32Uint8 failed")
	}

	if len(MapInt32Uint8Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int32 = 10

	expectedList := []*string{&vo10}
	newList := MapInt32StrPtr(someLogicInt32StrPtr, []*int32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32StrPtr(nil, nil)) > 0 {
		t.Errorf("MapInt32Str failed")
	}

	if len(MapInt32StrPtr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt32StrPtr(num *int32) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapInt32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32

	expectedList := []*bool{&vt, &vf}
	newList := MapInt32BoolPtr(someLogicInt32BoolPtr, []*int32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapInt32Bool failed")
	}

	if len(MapInt32BoolPtr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapInt32Float32Ptr(plusOneInt32Float32Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Float32(nil, nil)) > 0 {
		t.Errorf("MapInt32Float32 failed")
	}

	if len(MapInt32Float32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapInt32Float64Ptr(plusOneInt32Float64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt32Float64(nil, nil)) > 0 {
		t.Errorf("MapInt32Float64 failed")
	}

	if len(MapInt32Float64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("MapInt32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapInt16IntPtr(plusOneInt16IntPtr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int(nil, nil)) > 0 {
		t.Errorf("MapInt16Int failed")
	}

	if len(MapInt16IntPtr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapInt16Int64Ptr(plusOneInt16Int64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int64(nil, nil)) > 0 {
		t.Errorf("MapInt16Int64 failed")
	}

	if len(MapInt16Int64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapInt16Int32Ptr(plusOneInt16Int32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int32(nil, nil)) > 0 {
		t.Errorf("MapInt16Int32 failed")
	}

	if len(MapInt16Int32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapInt16Int8Ptr(plusOneInt16Int8Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Int8(nil, nil)) > 0 {
		t.Errorf("MapInt16Int8 failed")
	}

	if len(MapInt16Int8Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapInt16UintPtr(plusOneInt16UintPtr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint failed")
	}

	if len(MapInt16UintPtr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapInt16Uint64Ptr(plusOneInt16Uint64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint64 failed")
	}

	if len(MapInt16Uint64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapInt16Uint32Ptr(plusOneInt16Uint32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint32 failed")
	}

	if len(MapInt16Uint32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapInt16Uint16Ptr(plusOneInt16Uint16Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint16 failed")
	}

	if len(MapInt16Uint16Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapInt16Uint8Ptr(plusOneInt16Uint8Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt16Uint8 failed")
	}

	if len(MapInt16Uint8Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int16 = 10

	expectedList := []*string{&vo10}
	newList := MapInt16StrPtr(someLogicInt16StrPtr, []*int16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16StrPtr(nil, nil)) > 0 {
		t.Errorf("MapInt16Str failed")
	}

	if len(MapInt16StrPtr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt16StrPtr(num *int16) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapInt16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16

	expectedList := []*bool{&vt, &vf}
	newList := MapInt16BoolPtr(someLogicInt16BoolPtr, []*int16{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapInt16Bool failed")
	}

	if len(MapInt16BoolPtr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapInt16Float32Ptr(plusOneInt16Float32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Float32(nil, nil)) > 0 {
		t.Errorf("MapInt16Float32 failed")
	}

	if len(MapInt16Float32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapInt16Float64Ptr(plusOneInt16Float64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt16Float64(nil, nil)) > 0 {
		t.Errorf("MapInt16Float64 failed")
	}

	if len(MapInt16Float64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("MapInt16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapInt8IntPtr(plusOneInt8IntPtr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int(nil, nil)) > 0 {
		t.Errorf("MapInt8Int failed")
	}

	if len(MapInt8IntPtr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapInt8Int64Ptr(plusOneInt8Int64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int64(nil, nil)) > 0 {
		t.Errorf("MapInt8Int64 failed")
	}

	if len(MapInt8Int64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapInt8Int32Ptr(plusOneInt8Int32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int32(nil, nil)) > 0 {
		t.Errorf("MapInt8Int32 failed")
	}

	if len(MapInt8Int32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapInt8Int16Ptr(plusOneInt8Int16Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Int16(nil, nil)) > 0 {
		t.Errorf("MapInt8Int16 failed")
	}

	if len(MapInt8Int16Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapInt8UintPtr(plusOneInt8UintPtr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint failed")
	}

	if len(MapInt8UintPtr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapInt8Uint64Ptr(plusOneInt8Uint64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint64(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint64 failed")
	}

	if len(MapInt8Uint64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapInt8Uint32Ptr(plusOneInt8Uint32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint32(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint32 failed")
	}

	if len(MapInt8Uint32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapInt8Uint16Ptr(plusOneInt8Uint16Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint16(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint16 failed")
	}

	if len(MapInt8Uint16Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapInt8Uint8Ptr(plusOneInt8Uint8Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Uint8(nil, nil)) > 0 {
		t.Errorf("MapInt8Uint8 failed")
	}

	if len(MapInt8Uint8Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int8 = 10

	expectedList := []*string{&vo10}
	newList := MapInt8StrPtr(someLogicInt8StrPtr, []*int8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8StrPtr(nil, nil)) > 0 {
		t.Errorf("MapInt8Str failed")
	}

	if len(MapInt8StrPtr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt8StrPtr(num *int8) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapInt8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8

	expectedList := []*bool{&vt, &vf}
	newList := MapInt8BoolPtr(someLogicInt8BoolPtr, []*int8{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapInt8Bool failed")
	}

	if len(MapInt8BoolPtr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapInt8Float32Ptr(plusOneInt8Float32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Float32(nil, nil)) > 0 {
		t.Errorf("MapInt8Float32 failed")
	}

	if len(MapInt8Float32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapInt8Float64Ptr(plusOneInt8Float64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapInt8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapInt8Float64(nil, nil)) > 0 {
		t.Errorf("MapInt8Float64 failed")
	}

	if len(MapInt8Float64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("MapInt8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintIntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapUintIntPtr(plusOneUintIntPtr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt(nil, nil)) > 0 {
		t.Errorf("MapUintInt failed")
	}

	if len(MapUintIntPtr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapUintInt64Ptr(plusOneUintInt64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt64(nil, nil)) > 0 {
		t.Errorf("MapUintInt64 failed")
	}

	if len(MapUintInt64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapUintInt32Ptr(plusOneUintInt32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt32(nil, nil)) > 0 {
		t.Errorf("MapUintInt32 failed")
	}

	if len(MapUintInt32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapUintInt16Ptr(plusOneUintInt16Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt16(nil, nil)) > 0 {
		t.Errorf("MapUintInt16 failed")
	}

	if len(MapUintInt16Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapUintInt8Ptr(plusOneUintInt8Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintInt8(nil, nil)) > 0 {
		t.Errorf("MapUintInt8 failed")
	}

	if len(MapUintInt8Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapUintUint64Ptr(plusOneUintUint64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint64(nil, nil)) > 0 {
		t.Errorf("MapUintUint64 failed")
	}

	if len(MapUintUint64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapUintUint32Ptr(plusOneUintUint32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint32(nil, nil)) > 0 {
		t.Errorf("MapUintUint32 failed")
	}

	if len(MapUintUint32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapUintUint16Ptr(plusOneUintUint16Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint16(nil, nil)) > 0 {
		t.Errorf("MapUintUint16 failed")
	}

	if len(MapUintUint16Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapUintUint8Ptr(plusOneUintUint8Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintUint8(nil, nil)) > 0 {
		t.Errorf("MapUintUint8 failed")
	}

	if len(MapUintUint8Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintStrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint = 10

	expectedList := []*string{&vo10}
	newList := MapUintStrPtr(someLogicUintStrPtr, []*uint{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUintStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintStrPtr(nil, nil)) > 0 {
		t.Errorf("MapUintStr failed")
	}

	if len(MapUintStrPtr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUintStrPtr(num *uint) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapUintBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint

	expectedList := []*bool{&vt, &vf}
	newList := MapUintBoolPtr(someLogicUintBoolPtr, []*uint{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUintBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapUintBool failed")
	}

	if len(MapUintBoolPtr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapUintFloat32Ptr(plusOneUintFloat32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintFloat32(nil, nil)) > 0 {
		t.Errorf("MapUintFloat32 failed")
	}

	if len(MapUintFloat32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapUintFloat64Ptr(plusOneUintFloat64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUintFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUintFloat64(nil, nil)) > 0 {
		t.Errorf("MapUintFloat64 failed")
	}

	if len(MapUintFloat64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("MapUintFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapUint64IntPtr(plusOneUint64IntPtr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int(nil, nil)) > 0 {
		t.Errorf("MapUint64Int failed")
	}

	if len(MapUint64IntPtr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapUint64Int64Ptr(plusOneUint64Int64Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int64(nil, nil)) > 0 {
		t.Errorf("MapUint64Int64 failed")
	}

	if len(MapUint64Int64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapUint64Int32Ptr(plusOneUint64Int32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int32(nil, nil)) > 0 {
		t.Errorf("MapUint64Int32 failed")
	}

	if len(MapUint64Int32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapUint64Int16Ptr(plusOneUint64Int16Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int16(nil, nil)) > 0 {
		t.Errorf("MapUint64Int16 failed")
	}

	if len(MapUint64Int16Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapUint64Int8Ptr(plusOneUint64Int8Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Int8(nil, nil)) > 0 {
		t.Errorf("MapUint64Int8 failed")
	}

	if len(MapUint64Int8Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapUint64UintPtr(plusOneUint64UintPtr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint failed")
	}

	if len(MapUint64UintPtr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapUint64Uint32Ptr(plusOneUint64Uint32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint32 failed")
	}

	if len(MapUint64Uint32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapUint64Uint16Ptr(plusOneUint64Uint16Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint16 failed")
	}

	if len(MapUint64Uint16Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapUint64Uint8Ptr(plusOneUint64Uint8Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint64Uint8 failed")
	}

	if len(MapUint64Uint8Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint64 = 10

	expectedList := []*string{&vo10}
	newList := MapUint64StrPtr(someLogicUint64StrPtr, []*uint64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64StrPtr(nil, nil)) > 0 {
		t.Errorf("MapUint64Str failed")
	}

	if len(MapUint64StrPtr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint64StrPtr(num *uint64) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapUint64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64

	expectedList := []*bool{&vt, &vf}
	newList := MapUint64BoolPtr(someLogicUint64BoolPtr, []*uint64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapUint64Bool failed")
	}

	if len(MapUint64BoolPtr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapUint64Float32Ptr(plusOneUint64Float32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Float32(nil, nil)) > 0 {
		t.Errorf("MapUint64Float32 failed")
	}

	if len(MapUint64Float32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapUint64Float64Ptr(plusOneUint64Float64Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint64Float64(nil, nil)) > 0 {
		t.Errorf("MapUint64Float64 failed")
	}

	if len(MapUint64Float64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("MapUint64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapUint32IntPtr(plusOneUint32IntPtr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int(nil, nil)) > 0 {
		t.Errorf("MapUint32Int failed")
	}

	if len(MapUint32IntPtr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapUint32Int64Ptr(plusOneUint32Int64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int64(nil, nil)) > 0 {
		t.Errorf("MapUint32Int64 failed")
	}

	if len(MapUint32Int64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapUint32Int32Ptr(plusOneUint32Int32Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int32(nil, nil)) > 0 {
		t.Errorf("MapUint32Int32 failed")
	}

	if len(MapUint32Int32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapUint32Int16Ptr(plusOneUint32Int16Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int16(nil, nil)) > 0 {
		t.Errorf("MapUint32Int16 failed")
	}

	if len(MapUint32Int16Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapUint32Int8Ptr(plusOneUint32Int8Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Int8(nil, nil)) > 0 {
		t.Errorf("MapUint32Int8 failed")
	}

	if len(MapUint32Int8Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapUint32UintPtr(plusOneUint32UintPtr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint failed")
	}

	if len(MapUint32UintPtr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapUint32Uint64Ptr(plusOneUint32Uint64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint64 failed")
	}

	if len(MapUint32Uint64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapUint32Uint16Ptr(plusOneUint32Uint16Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint16 failed")
	}

	if len(MapUint32Uint16Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapUint32Uint8Ptr(plusOneUint32Uint8Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint32Uint8 failed")
	}

	if len(MapUint32Uint8Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint32 = 10

	expectedList := []*string{&vo10}
	newList := MapUint32StrPtr(someLogicUint32StrPtr, []*uint32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32StrPtr(nil, nil)) > 0 {
		t.Errorf("MapUint32Str failed")
	}

	if len(MapUint32StrPtr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint32StrPtr(num *uint32) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapUint32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32

	expectedList := []*bool{&vt, &vf}
	newList := MapUint32BoolPtr(someLogicUint32BoolPtr, []*uint32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapUint32Bool failed")
	}

	if len(MapUint32BoolPtr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapUint32Float32Ptr(plusOneUint32Float32Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Float32(nil, nil)) > 0 {
		t.Errorf("MapUint32Float32 failed")
	}

	if len(MapUint32Float32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapUint32Float64Ptr(plusOneUint32Float64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint32Float64(nil, nil)) > 0 {
		t.Errorf("MapUint32Float64 failed")
	}

	if len(MapUint32Float64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("MapUint32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapUint16IntPtr(plusOneUint16IntPtr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int(nil, nil)) > 0 {
		t.Errorf("MapUint16Int failed")
	}

	if len(MapUint16IntPtr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapUint16Int64Ptr(plusOneUint16Int64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int64(nil, nil)) > 0 {
		t.Errorf("MapUint16Int64 failed")
	}

	if len(MapUint16Int64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapUint16Int32Ptr(plusOneUint16Int32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int32(nil, nil)) > 0 {
		t.Errorf("MapUint16Int32 failed")
	}

	if len(MapUint16Int32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapUint16Int16Ptr(plusOneUint16Int16Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int16(nil, nil)) > 0 {
		t.Errorf("MapUint16Int16 failed")
	}

	if len(MapUint16Int16Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapUint16Int8Ptr(plusOneUint16Int8Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Int8(nil, nil)) > 0 {
		t.Errorf("MapUint16Int8 failed")
	}

	if len(MapUint16Int8Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapUint16UintPtr(plusOneUint16UintPtr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint failed")
	}

	if len(MapUint16UintPtr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapUint16Uint64Ptr(plusOneUint16Uint64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint64 failed")
	}

	if len(MapUint16Uint64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapUint16Uint32Ptr(plusOneUint16Uint32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint32 failed")
	}

	if len(MapUint16Uint32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapUint16Uint8Ptr(plusOneUint16Uint8Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Uint8(nil, nil)) > 0 {
		t.Errorf("MapUint16Uint8 failed")
	}

	if len(MapUint16Uint8Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint16 = 10

	expectedList := []*string{&vo10}
	newList := MapUint16StrPtr(someLogicUint16StrPtr, []*uint16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16StrPtr(nil, nil)) > 0 {
		t.Errorf("MapUint16Str failed")
	}

	if len(MapUint16StrPtr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint16StrPtr(num *uint16) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapUint16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16

	expectedList := []*bool{&vt, &vf}
	newList := MapUint16BoolPtr(someLogicUint16BoolPtr, []*uint16{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapUint16Bool failed")
	}

	if len(MapUint16BoolPtr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapUint16Float32Ptr(plusOneUint16Float32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Float32(nil, nil)) > 0 {
		t.Errorf("MapUint16Float32 failed")
	}

	if len(MapUint16Float32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapUint16Float64Ptr(plusOneUint16Float64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint16Float64(nil, nil)) > 0 {
		t.Errorf("MapUint16Float64 failed")
	}

	if len(MapUint16Float64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("MapUint16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapUint8IntPtr(plusOneUint8IntPtr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int(nil, nil)) > 0 {
		t.Errorf("MapUint8Int failed")
	}

	if len(MapUint8IntPtr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapUint8Int64Ptr(plusOneUint8Int64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int64(nil, nil)) > 0 {
		t.Errorf("MapUint8Int64 failed")
	}

	if len(MapUint8Int64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapUint8Int32Ptr(plusOneUint8Int32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int32(nil, nil)) > 0 {
		t.Errorf("MapUint8Int32 failed")
	}

	if len(MapUint8Int32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapUint8Int16Ptr(plusOneUint8Int16Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int16(nil, nil)) > 0 {
		t.Errorf("MapUint8Int16 failed")
	}

	if len(MapUint8Int16Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapUint8Int8Ptr(plusOneUint8Int8Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Int8(nil, nil)) > 0 {
		t.Errorf("MapUint8Int8 failed")
	}

	if len(MapUint8Int8Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapUint8UintPtr(plusOneUint8UintPtr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint failed")
	}

	if len(MapUint8UintPtr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapUint8Uint64Ptr(plusOneUint8Uint64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint64(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint64 failed")
	}

	if len(MapUint8Uint64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapUint8Uint32Ptr(plusOneUint8Uint32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint32(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint32 failed")
	}

	if len(MapUint8Uint32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapUint8Uint16Ptr(plusOneUint8Uint16Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Uint16(nil, nil)) > 0 {
		t.Errorf("MapUint8Uint16 failed")
	}

	if len(MapUint8Uint16Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint8 = 10

	expectedList := []*string{&vo10}
	newList := MapUint8StrPtr(someLogicUint8StrPtr, []*uint8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8StrPtr(nil, nil)) > 0 {
		t.Errorf("MapUint8Str failed")
	}

	if len(MapUint8StrPtr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint8StrPtr(num *uint8) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapUint8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8

	expectedList := []*bool{&vt, &vf}
	newList := MapUint8BoolPtr(someLogicUint8BoolPtr, []*uint8{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapUint8Bool failed")
	}

	if len(MapUint8BoolPtr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapUint8Float32Ptr(plusOneUint8Float32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Float32(nil, nil)) > 0 {
		t.Errorf("MapUint8Float32 failed")
	}

	if len(MapUint8Float32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapUint8Float64Ptr(plusOneUint8Float64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapUint8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapUint8Float64(nil, nil)) > 0 {
		t.Errorf("MapUint8Float64 failed")
	}

	if len(MapUint8Float64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("MapUint8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrIntPtr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vi10 string = "10"

	expectedList := []*int{&vo10}
	newList := MapStrIntPtr(someLogicStrIntPtr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrIntPtr(nil, nil)) > 0 {
		t.Errorf("MapStrIntPtr failed")
	}

	if len(MapStrIntPtr(nil, []*string{})) > 0 {
		t.Errorf("MapStrIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrIntPtr(num *string) *int {
	var r int
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrInt64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vi10 string = "10"

	expectedList := []*int64{&vo10}
	newList := MapStrInt64Ptr(someLogicStrInt64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt64Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrInt64Ptr failed")
	}

	if len(MapStrInt64Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt64Ptr(num *string) *int64 {
	var r int64
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrInt32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vi10 string = "10"

	expectedList := []*int32{&vo10}
	newList := MapStrInt32Ptr(someLogicStrInt32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt32Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrInt32Ptr failed")
	}

	if len(MapStrInt32Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt32Ptr(num *string) *int32 {
	var r int32
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrInt16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vi10 string = "10"

	expectedList := []*int16{&vo10}
	newList := MapStrInt16Ptr(someLogicStrInt16Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt16Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrInt16Ptr failed")
	}

	if len(MapStrInt16Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt16Ptr(num *string) *int16 {
	var r int16
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrInt8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vi10 string = "10"

	expectedList := []*int8{&vo10}
	newList := MapStrInt8Ptr(someLogicStrInt8Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrInt8Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrInt8Ptr failed")
	}

	if len(MapStrInt8Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt8Ptr(num *string) *int8 {
	var r int8
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrUintPtr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vi10 string = "10"

	expectedList := []*uint{&vo10}
	newList := MapStrUintPtr(someLogicStrUintPtr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUintPtr(nil, nil)) > 0 {
		t.Errorf("MapStrUintPtr failed")
	}

	if len(MapStrUintPtr(nil, []*string{})) > 0 {
		t.Errorf("MapStrUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUintPtr(num *string) *uint {
	var r uint
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrUint64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vi10 string = "10"

	expectedList := []*uint64{&vo10}
	newList := MapStrUint64Ptr(someLogicStrUint64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint64Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrUint64Ptr failed")
	}

	if len(MapStrUint64Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint64Ptr(num *string) *uint64 {
	var r uint64
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrUint32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vi10 string = "10"

	expectedList := []*uint32{&vo10}
	newList := MapStrUint32Ptr(someLogicStrUint32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint32Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrUint32Ptr failed")
	}

	if len(MapStrUint32Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint32Ptr(num *string) *uint32 {
	var r uint32
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrUint16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vi10 string = "10"

	expectedList := []*uint16{&vo10}
	newList := MapStrUint16Ptr(someLogicStrUint16Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint16Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrUint16Ptr failed")
	}

	if len(MapStrUint16Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint16Ptr(num *string) *uint16 {
	var r uint16
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrUint8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vi10 string = "10"

	expectedList := []*uint8{&vo10}
	newList := MapStrUint8Ptr(someLogicStrUint8Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrUint8Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrUint8Ptr failed")
	}

	if len(MapStrUint8Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint8Ptr(num *string) *uint8 {
	var r uint8
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	expectedList := []*bool{&vt, &vf}
	newList := MapStrBoolPtr(someLogicStrBoolPtr, []*string{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapStrBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapStrBoolPtr failed")
	}

	if len(MapStrBoolPtr(nil, []*string{})) > 0 {
		t.Errorf("MapStrBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrFloat32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vi10 string = "10"

	expectedList := []*float32{&vo10}
	newList := MapStrFloat32Ptr(someLogicStrFloat32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrFloat32Ptr failed")
	}

	if len(MapStrFloat32Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat32Ptr(num *string) *float32 {
	var r float32
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapStrFloat64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vi10 string = "10"

	expectedList := []*float64{&vo10}
	newList := MapStrFloat64Ptr(someLogicStrFloat64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapStrFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("MapStrFloat64Ptr failed")
	}

	if len(MapStrFloat64Ptr(nil, []*string{})) > 0 {
		t.Errorf("MapStrFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat64Ptr(num *string) *float64 {
	var r float64
	if *num == "10" {
		r = 10
	}
	return &r
}

func TestMapBoolIntPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int

	// Test : someLogic
	expectedList := []*int{&v10, &v0}
	newList := MapBoolIntPtr(someLogicBoolIntPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolIntPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolIntPtr failed")
	}

	if len(MapBoolIntPtr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64

	// Test : someLogic
	expectedList := []*int64{&v10, &v0}
	newList := MapBoolInt64Ptr(someLogicBoolInt64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt64Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolInt64Ptr failed")
	}

	if len(MapBoolInt64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32

	// Test : someLogic
	expectedList := []*int32{&v10, &v0}
	newList := MapBoolInt32Ptr(someLogicBoolInt32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt32Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolInt32Ptr failed")
	}

	if len(MapBoolInt32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt16Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16

	// Test : someLogic
	expectedList := []*int16{&v10, &v0}
	newList := MapBoolInt16Ptr(someLogicBoolInt16Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt16Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolInt16Ptr failed")
	}

	if len(MapBoolInt16Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt8Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8

	// Test : someLogic
	expectedList := []*int8{&v10, &v0}
	newList := MapBoolInt8Ptr(someLogicBoolInt8Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolInt8Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolInt8Ptr failed")
	}

	if len(MapBoolInt8Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUintPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint

	// Test : someLogic
	expectedList := []*uint{&v10, &v0}
	newList := MapBoolUintPtr(someLogicBoolUintPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUintPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolUintPtr failed")
	}

	if len(MapBoolUintPtr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64

	// Test : someLogic
	expectedList := []*uint64{&v10, &v0}
	newList := MapBoolUint64Ptr(someLogicBoolUint64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint64Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolUint64Ptr failed")
	}

	if len(MapBoolUint64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32

	// Test : someLogic
	expectedList := []*uint32{&v10, &v0}
	newList := MapBoolUint32Ptr(someLogicBoolUint32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint32Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolUint32Ptr failed")
	}

	if len(MapBoolUint32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint16Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16

	// Test : someLogic
	expectedList := []*uint16{&v10, &v0}
	newList := MapBoolUint16Ptr(someLogicBoolUint16Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint16Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolUint16Ptr failed")
	}

	if len(MapBoolUint16Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint8Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8

	// Test : someLogic
	expectedList := []*uint8{&v10, &v0}
	newList := MapBoolUint8Ptr(someLogicBoolUint8Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolUint8Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolUint8Ptr failed")
	}

	if len(MapBoolUint8Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolStrPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	// Test : someLogic
	expectedList := []*string{&v10, &v0}
	newList := MapBoolStrPtr(someLogicBoolStrPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolStrPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolStr failed")
	}

	if len(MapBoolStrPtr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolStrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat32Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32

	// Test : someLogic
	expectedList := []*float32{&v10, &v0}
	newList := MapBoolFloat32Ptr(someLogicBoolFloat32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolFloat32Ptr failed")
	}

	if len(MapBoolFloat32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat64Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64

	// Test : someLogic
	expectedList := []*float64{&v10, &v0}
	newList := MapBoolFloat64Ptr(someLogicBoolFloat64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapBoolFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("MapBoolFloat64Ptr failed")
	}

	if len(MapBoolFloat64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("MapBoolFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapFloat32IntPtr(plusOneFloat32IntPtr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int failed")
	}

	if len(MapFloat32IntPtr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapFloat32Int64Ptr(plusOneFloat32Int64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int64 failed")
	}

	if len(MapFloat32Int64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapFloat32Int32Ptr(plusOneFloat32Int32Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int32(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int32 failed")
	}

	if len(MapFloat32Int32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapFloat32Int16Ptr(plusOneFloat32Int16Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int16(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int16 failed")
	}

	if len(MapFloat32Int16Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapFloat32Int8Ptr(plusOneFloat32Int8Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Int8(nil, nil)) > 0 {
		t.Errorf("MapFloat32Int8 failed")
	}

	if len(MapFloat32Int8Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapFloat32UintPtr(plusOneFloat32UintPtr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint failed")
	}

	if len(MapFloat32UintPtr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapFloat32Uint64Ptr(plusOneFloat32Uint64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint64 failed")
	}

	if len(MapFloat32Uint64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapFloat32Uint32Ptr(plusOneFloat32Uint32Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint32(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint32 failed")
	}

	if len(MapFloat32Uint32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapFloat32Uint16Ptr(plusOneFloat32Uint16Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint16(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint16 failed")
	}

	if len(MapFloat32Uint16Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapFloat32Uint8Ptr(plusOneFloat32Uint8Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Uint8(nil, nil)) > 0 {
		t.Errorf("MapFloat32Uint8 failed")
	}

	if len(MapFloat32Uint8Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float32 = 10

	expectedList := []*string{&vo10}
	newList := MapFloat32StrPtr(someLogicFloat32StrPtr, []*float32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapFloat32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32StrPtr(nil, nil)) > 0 {
		t.Errorf("MapFloat32Str failed")
	}

	if len(MapFloat32StrPtr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat32StrPtr(num *float32) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapFloat32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32

	expectedList := []*bool{&vt, &vf}
	newList := MapFloat32BoolPtr(someLogicFloat32BoolPtr, []*float32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapFloat32Bool failed")
	}

	if len(MapFloat32BoolPtr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := MapFloat32Float64Ptr(plusOneFloat32Float64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat32Float64(nil, nil)) > 0 {
		t.Errorf("MapFloat32Float64 failed")
	}

	if len(MapFloat32Float64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("MapFloat32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := MapFloat64IntPtr(plusOneFloat64IntPtr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int failed")
	}

	if len(MapFloat64IntPtr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := MapFloat64Int64Ptr(plusOneFloat64Int64Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int64(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int64 failed")
	}

	if len(MapFloat64Int64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := MapFloat64Int32Ptr(plusOneFloat64Int32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int32 failed")
	}

	if len(MapFloat64Int32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := MapFloat64Int16Ptr(plusOneFloat64Int16Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int16(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int16 failed")
	}

	if len(MapFloat64Int16Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := MapFloat64Int8Ptr(plusOneFloat64Int8Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Int8(nil, nil)) > 0 {
		t.Errorf("MapFloat64Int8 failed")
	}

	if len(MapFloat64Int8Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := MapFloat64UintPtr(plusOneFloat64UintPtr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint failed")
	}

	if len(MapFloat64UintPtr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := MapFloat64Uint64Ptr(plusOneFloat64Uint64Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint64(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint64 failed")
	}

	if len(MapFloat64Uint64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := MapFloat64Uint32Ptr(plusOneFloat64Uint32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint32 failed")
	}

	if len(MapFloat64Uint32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := MapFloat64Uint16Ptr(plusOneFloat64Uint16Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint16(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint16 failed")
	}

	if len(MapFloat64Uint16Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := MapFloat64Uint8Ptr(plusOneFloat64Uint8Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Uint8(nil, nil)) > 0 {
		t.Errorf("MapFloat64Uint8 failed")
	}

	if len(MapFloat64Uint8Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float64 = 10

	expectedList := []*string{&vo10}
	newList := MapFloat64StrPtr(someLogicFloat64StrPtr, []*float64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapFloat64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64StrPtr(nil, nil)) > 0 {
		t.Errorf("MapFloat64Str failed")
	}

	if len(MapFloat64StrPtr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat64StrPtr(num *float64) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}

func TestMapFloat64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64

	expectedList := []*bool{&vt, &vf}
	newList := MapFloat64BoolPtr(someLogicFloat64BoolPtr, []*float64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64BoolPtr(nil, nil)) > 0 {
		t.Errorf("MapFloat64Bool failed")
	}

	if len(MapFloat64BoolPtr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := MapFloat64Float32Ptr(plusOneFloat64Float32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("MapFloat64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(MapFloat64Float32(nil, nil)) > 0 {
		t.Errorf("MapFloat64Float32 failed")
	}

	if len(MapFloat64Float32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("MapFloat64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
