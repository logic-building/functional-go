package fp

import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntInt64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt64PtrErr failed")
	}

	r, _ = MapIntInt64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt64PtrErr failed")
	}

	_, err := MapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt32PtrErr failed")
	}

	r, _ = MapIntInt32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt32PtrErr failed")
	}

	_, err := MapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt16PtrErr failed")
	}

	r, _ = MapIntInt16PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt16PtrErr failed")
	}

	_, err := MapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntInt8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntInt8PtrErr failed")
	}

	r, _ = MapIntInt8PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntInt8PtrErr failed")
	}

	_, err := MapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUintPtrErr failed")
	}

	r, _ = MapIntUintPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntUintPtrErr failed")
	}

	_, err := MapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntUintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint64PtrErr failed")
	}

	r, _ = MapIntUint64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint64PtrErr failed")
	}

	_, err := MapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint32PtrErr failed")
	}

	r, _ = MapIntUint32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint32PtrErr failed")
	}

	_, err := MapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint16PtrErr failed")
	}

	r, _ = MapIntUint16PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint16PtrErr failed")
	}

	_, err := MapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntUint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntUint8PtrErr failed")
	}

	r, _ = MapIntUint8PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntUint8PtrErr failed")
	}

	_, err := MapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int = 10
	var vi0 int

	expectedList := []*string{&vo10}
	newList, _ := MapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapIntStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntStr failed")
	}

	r, _ = MapIntStrPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntStrPtrErr failed")
	}
	_, err := MapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapIntStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicIntStrPtrErr(num *int) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapIntBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int
	var v3 int = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapIntBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntBool failed")
	}

	r, _ = MapIntBoolPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntBoolPtrErr failed")
	}
	_, err := MapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v3})
	if err == nil {
		t.Errorf("MapIntBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntFloat32PtrErr failed")
	}

	r, _ = MapIntFloat32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntFloat32PtrErr failed")
	}

	_, err := MapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapIntFloat64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapIntFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapIntFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapIntFloat64PtrErr failed")
	}

	r, _ = MapIntFloat64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("MapIntFloat64PtrErr failed")
	}

	_, err := MapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapIntFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64IntPtrErr failed")
	}

	r, _ = MapInt64IntPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64IntPtrErr failed")
	}

	_, err := MapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int32PtrErr failed")
	}

	r, _ = MapInt64Int32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int32PtrErr failed")
	}

	_, err := MapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int16PtrErr failed")
	}

	r, _ = MapInt64Int16PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int16PtrErr failed")
	}

	_, err := MapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Int8PtrErr failed")
	}

	r, _ = MapInt64Int8PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Int8PtrErr failed")
	}

	_, err := MapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64UintPtrErr failed")
	}

	r, _ = MapInt64UintPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64UintPtrErr failed")
	}

	_, err := MapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint64PtrErr failed")
	}

	r, _ = MapInt64Uint64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint64PtrErr failed")
	}

	_, err := MapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint32PtrErr failed")
	}

	r, _ = MapInt64Uint32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint32PtrErr failed")
	}

	_, err := MapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint16PtrErr failed")
	}

	r, _ = MapInt64Uint16PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint16PtrErr failed")
	}

	_, err := MapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Uint8PtrErr failed")
	}

	r, _ = MapInt64Uint8PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Uint8PtrErr failed")
	}

	_, err := MapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int64 = 10
	var vi0 int64

	expectedList := []*string{&vo10}
	newList, _ := MapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Str failed")
	}

	r, _ = MapInt64StrPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64StrPtrErr failed")
	}
	_, err := MapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapInt64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt64StrPtrErr(num *int64) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapInt64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64
	var v3 int64 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Bool failed")
	}

	r, _ = MapInt64BoolPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64BoolPtrErr failed")
	}
	_, err := MapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v3})
	if err == nil {
		t.Errorf("MapInt64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Float32PtrErr failed")
	}

	r, _ = MapInt64Float32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Float32PtrErr failed")
	}

	_, err := MapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt64Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt64Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt64Float64PtrErr failed")
	}

	r, _ = MapInt64Float64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("MapInt64Float64PtrErr failed")
	}

	_, err := MapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt64Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32IntPtrErr failed")
	}

	r, _ = MapInt32IntPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32IntPtrErr failed")
	}

	_, err := MapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int64PtrErr failed")
	}

	r, _ = MapInt32Int64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int64PtrErr failed")
	}

	_, err := MapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int16PtrErr failed")
	}

	r, _ = MapInt32Int16PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int16PtrErr failed")
	}

	_, err := MapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Int8PtrErr failed")
	}

	r, _ = MapInt32Int8PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Int8PtrErr failed")
	}

	_, err := MapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32UintPtrErr failed")
	}

	r, _ = MapInt32UintPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32UintPtrErr failed")
	}

	_, err := MapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint64PtrErr failed")
	}

	r, _ = MapInt32Uint64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint64PtrErr failed")
	}

	_, err := MapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint32PtrErr failed")
	}

	r, _ = MapInt32Uint32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint32PtrErr failed")
	}

	_, err := MapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint16PtrErr failed")
	}

	r, _ = MapInt32Uint16PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint16PtrErr failed")
	}

	_, err := MapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Uint8PtrErr failed")
	}

	r, _ = MapInt32Uint8PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Uint8PtrErr failed")
	}

	_, err := MapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int32 = 10
	var vi0 int32

	expectedList := []*string{&vo10}
	newList, _ := MapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Str failed")
	}

	r, _ = MapInt32StrPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32StrPtrErr failed")
	}
	_, err := MapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapInt32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt32StrPtrErr(num *int32) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapInt32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32
	var v3 int32 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Bool failed")
	}

	r, _ = MapInt32BoolPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32BoolPtrErr failed")
	}
	_, err := MapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v3})
	if err == nil {
		t.Errorf("MapInt32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Float32PtrErr failed")
	}

	r, _ = MapInt32Float32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Float32PtrErr failed")
	}

	_, err := MapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt32Float64PtrErr failed")
	}

	r, _ = MapInt32Float64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("MapInt32Float64PtrErr failed")
	}

	_, err := MapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16IntPtrErr failed")
	}

	r, _ = MapInt16IntPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16IntPtrErr failed")
	}

	_, err := MapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int64PtrErr failed")
	}

	r, _ = MapInt16Int64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int64PtrErr failed")
	}

	_, err := MapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int32PtrErr failed")
	}

	r, _ = MapInt16Int32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int32PtrErr failed")
	}

	_, err := MapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Int8PtrErr failed")
	}

	r, _ = MapInt16Int8PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Int8PtrErr failed")
	}

	_, err := MapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16UintPtrErr failed")
	}

	r, _ = MapInt16UintPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16UintPtrErr failed")
	}

	_, err := MapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint64PtrErr failed")
	}

	r, _ = MapInt16Uint64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint64PtrErr failed")
	}

	_, err := MapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint32PtrErr failed")
	}

	r, _ = MapInt16Uint32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint32PtrErr failed")
	}

	_, err := MapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint16PtrErr failed")
	}

	r, _ = MapInt16Uint16PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint16PtrErr failed")
	}

	_, err := MapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Uint8PtrErr failed")
	}

	r, _ = MapInt16Uint8PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Uint8PtrErr failed")
	}

	_, err := MapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int16 = 10
	var vi0 int16

	expectedList := []*string{&vo10}
	newList, _ := MapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Str failed")
	}

	r, _ = MapInt16StrPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16StrPtrErr failed")
	}
	_, err := MapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapInt16StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt16StrPtrErr(num *int16) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapInt16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16
	var v3 int16 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Bool failed")
	}

	r, _ = MapInt16BoolPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16BoolPtrErr failed")
	}
	_, err := MapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v3})
	if err == nil {
		t.Errorf("MapInt16BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Float32PtrErr failed")
	}

	r, _ = MapInt16Float32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Float32PtrErr failed")
	}

	_, err := MapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt16Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt16Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt16Float64PtrErr failed")
	}

	r, _ = MapInt16Float64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("MapInt16Float64PtrErr failed")
	}

	_, err := MapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt16Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8IntPtrErr failed")
	}

	r, _ = MapInt8IntPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8IntPtrErr failed")
	}

	_, err := MapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int64PtrErr failed")
	}

	r, _ = MapInt8Int64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int64PtrErr failed")
	}

	_, err := MapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int32PtrErr failed")
	}

	r, _ = MapInt8Int32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int32PtrErr failed")
	}

	_, err := MapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Int16PtrErr failed")
	}

	r, _ = MapInt8Int16PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Int16PtrErr failed")
	}

	_, err := MapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8UintPtrErr failed")
	}

	r, _ = MapInt8UintPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8UintPtrErr failed")
	}

	_, err := MapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint64PtrErr failed")
	}

	r, _ = MapInt8Uint64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint64PtrErr failed")
	}

	_, err := MapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint32PtrErr failed")
	}

	r, _ = MapInt8Uint32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint32PtrErr failed")
	}

	_, err := MapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint16PtrErr failed")
	}

	r, _ = MapInt8Uint16PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint16PtrErr failed")
	}

	_, err := MapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Uint8PtrErr failed")
	}

	r, _ = MapInt8Uint8PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Uint8PtrErr failed")
	}

	_, err := MapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int8 = 10
	var vi0 int8

	expectedList := []*string{&vo10}
	newList, _ := MapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapInt8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Str failed")
	}

	r, _ = MapInt8StrPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8StrPtrErr failed")
	}
	_, err := MapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapInt8StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicInt8StrPtrErr(num *int8) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapInt8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8
	var v3 int8 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Bool failed")
	}

	r, _ = MapInt8BoolPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8BoolPtrErr failed")
	}
	_, err := MapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v3})
	if err == nil {
		t.Errorf("MapInt8BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Float32PtrErr failed")
	}

	r, _ = MapInt8Float32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Float32PtrErr failed")
	}

	_, err := MapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapInt8Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapInt8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapInt8Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapInt8Float64PtrErr failed")
	}

	r, _ = MapInt8Float64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("MapInt8Float64PtrErr failed")
	}

	_, err := MapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapInt8Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintIntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintIntPtrErr failed")
	}

	r, _ = MapUintIntPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintIntPtrErr failed")
	}

	_, err := MapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintIntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt64PtrErr failed")
	}

	r, _ = MapUintInt64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt64PtrErr failed")
	}

	_, err := MapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt32PtrErr failed")
	}

	r, _ = MapUintInt32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt32PtrErr failed")
	}

	_, err := MapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt16PtrErr failed")
	}

	r, _ = MapUintInt16PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt16PtrErr failed")
	}

	_, err := MapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintInt8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintInt8PtrErr failed")
	}

	r, _ = MapUintInt8PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintInt8PtrErr failed")
	}

	_, err := MapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint64PtrErr failed")
	}

	r, _ = MapUintUint64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint64PtrErr failed")
	}

	_, err := MapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint32PtrErr failed")
	}

	r, _ = MapUintUint32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint32PtrErr failed")
	}

	_, err := MapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint16PtrErr failed")
	}

	r, _ = MapUintUint16PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint16PtrErr failed")
	}

	_, err := MapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintUint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintUint8PtrErr failed")
	}

	r, _ = MapUintUint8PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintUint8PtrErr failed")
	}

	_, err := MapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint = 10
	var vi0 uint

	expectedList := []*string{&vo10}
	newList, _ := MapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUintStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintStr failed")
	}

	r, _ = MapUintStrPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintStrPtrErr failed")
	}
	_, err := MapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapUintStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUintStrPtrErr(num *uint) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapUintBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint
	var v3 uint = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUintBoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintBool failed")
	}

	r, _ = MapUintBoolPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintBoolPtrErr failed")
	}
	_, err := MapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v3})
	if err == nil {
		t.Errorf("MapUintBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintFloat32PtrErr failed")
	}

	r, _ = MapUintFloat32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintFloat32PtrErr failed")
	}

	_, err := MapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUintFloat64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUintFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUintFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUintFloat64PtrErr failed")
	}

	r, _ = MapUintFloat64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("MapUintFloat64PtrErr failed")
	}

	_, err := MapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUintFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64IntPtrErr failed")
	}

	r, _ = MapUint64IntPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64IntPtrErr failed")
	}

	_, err := MapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int64PtrErr failed")
	}

	r, _ = MapUint64Int64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int64PtrErr failed")
	}

	_, err := MapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int32PtrErr failed")
	}

	r, _ = MapUint64Int32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int32PtrErr failed")
	}

	_, err := MapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int16PtrErr failed")
	}

	r, _ = MapUint64Int16PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int16PtrErr failed")
	}

	_, err := MapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Int8PtrErr failed")
	}

	r, _ = MapUint64Int8PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Int8PtrErr failed")
	}

	_, err := MapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64UintPtrErr failed")
	}

	r, _ = MapUint64UintPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64UintPtrErr failed")
	}

	_, err := MapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint32PtrErr failed")
	}

	r, _ = MapUint64Uint32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint32PtrErr failed")
	}

	_, err := MapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint16PtrErr failed")
	}

	r, _ = MapUint64Uint16PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint16PtrErr failed")
	}

	_, err := MapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Uint8PtrErr failed")
	}

	r, _ = MapUint64Uint8PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Uint8PtrErr failed")
	}

	_, err := MapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint64 = 10
	var vi0 uint64

	expectedList := []*string{&vo10}
	newList, _ := MapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Str failed")
	}

	r, _ = MapUint64StrPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64StrPtrErr failed")
	}
	_, err := MapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapUint64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint64StrPtrErr(num *uint64) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapUint64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64
	var v3 uint64 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Bool failed")
	}

	r, _ = MapUint64BoolPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64BoolPtrErr failed")
	}
	_, err := MapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v3})
	if err == nil {
		t.Errorf("MapUint64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Float32PtrErr failed")
	}

	r, _ = MapUint64Float32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Float32PtrErr failed")
	}

	_, err := MapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint64Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint64Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint64Float64PtrErr failed")
	}

	r, _ = MapUint64Float64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("MapUint64Float64PtrErr failed")
	}

	_, err := MapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint64Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32IntPtrErr failed")
	}

	r, _ = MapUint32IntPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32IntPtrErr failed")
	}

	_, err := MapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int64PtrErr failed")
	}

	r, _ = MapUint32Int64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int64PtrErr failed")
	}

	_, err := MapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int32PtrErr failed")
	}

	r, _ = MapUint32Int32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int32PtrErr failed")
	}

	_, err := MapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int16PtrErr failed")
	}

	r, _ = MapUint32Int16PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int16PtrErr failed")
	}

	_, err := MapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Int8PtrErr failed")
	}

	r, _ = MapUint32Int8PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Int8PtrErr failed")
	}

	_, err := MapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32UintPtrErr failed")
	}

	r, _ = MapUint32UintPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32UintPtrErr failed")
	}

	_, err := MapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint64PtrErr failed")
	}

	r, _ = MapUint32Uint64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint64PtrErr failed")
	}

	_, err := MapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint16PtrErr failed")
	}

	r, _ = MapUint32Uint16PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint16PtrErr failed")
	}

	_, err := MapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Uint8PtrErr failed")
	}

	r, _ = MapUint32Uint8PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Uint8PtrErr failed")
	}

	_, err := MapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint32 = 10
	var vi0 uint32

	expectedList := []*string{&vo10}
	newList, _ := MapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Str failed")
	}

	r, _ = MapUint32StrPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32StrPtrErr failed")
	}
	_, err := MapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapUint32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint32StrPtrErr(num *uint32) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapUint32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32
	var v3 uint32 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Bool failed")
	}

	r, _ = MapUint32BoolPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32BoolPtrErr failed")
	}
	_, err := MapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v3})
	if err == nil {
		t.Errorf("MapUint32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Float32PtrErr failed")
	}

	r, _ = MapUint32Float32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Float32PtrErr failed")
	}

	_, err := MapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint32Float64PtrErr failed")
	}

	r, _ = MapUint32Float64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("MapUint32Float64PtrErr failed")
	}

	_, err := MapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16IntPtrErr failed")
	}

	r, _ = MapUint16IntPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16IntPtrErr failed")
	}

	_, err := MapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int64PtrErr failed")
	}

	r, _ = MapUint16Int64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int64PtrErr failed")
	}

	_, err := MapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int32PtrErr failed")
	}

	r, _ = MapUint16Int32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int32PtrErr failed")
	}

	_, err := MapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int16PtrErr failed")
	}

	r, _ = MapUint16Int16PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int16PtrErr failed")
	}

	_, err := MapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Int8PtrErr failed")
	}

	r, _ = MapUint16Int8PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Int8PtrErr failed")
	}

	_, err := MapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16UintPtrErr failed")
	}

	r, _ = MapUint16UintPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16UintPtrErr failed")
	}

	_, err := MapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint64PtrErr failed")
	}

	r, _ = MapUint16Uint64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint64PtrErr failed")
	}

	_, err := MapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint32PtrErr failed")
	}

	r, _ = MapUint16Uint32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint32PtrErr failed")
	}

	_, err := MapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Uint8PtrErr failed")
	}

	r, _ = MapUint16Uint8PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Uint8PtrErr failed")
	}

	_, err := MapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint16 = 10
	var vi0 uint16

	expectedList := []*string{&vo10}
	newList, _ := MapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Str failed")
	}

	r, _ = MapUint16StrPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16StrPtrErr failed")
	}
	_, err := MapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapUint16StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint16StrPtrErr(num *uint16) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapUint16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16
	var v3 uint16 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Bool failed")
	}

	r, _ = MapUint16BoolPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16BoolPtrErr failed")
	}
	_, err := MapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v3})
	if err == nil {
		t.Errorf("MapUint16BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Float32PtrErr failed")
	}

	r, _ = MapUint16Float32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Float32PtrErr failed")
	}

	_, err := MapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint16Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint16Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint16Float64PtrErr failed")
	}

	r, _ = MapUint16Float64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("MapUint16Float64PtrErr failed")
	}

	_, err := MapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint16Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8IntPtrErr failed")
	}

	r, _ = MapUint8IntPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8IntPtrErr failed")
	}

	_, err := MapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int64PtrErr failed")
	}

	r, _ = MapUint8Int64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int64PtrErr failed")
	}

	_, err := MapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int32PtrErr failed")
	}

	r, _ = MapUint8Int32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int32PtrErr failed")
	}

	_, err := MapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int16PtrErr failed")
	}

	r, _ = MapUint8Int16PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int16PtrErr failed")
	}

	_, err := MapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Int8PtrErr failed")
	}

	r, _ = MapUint8Int8PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Int8PtrErr failed")
	}

	_, err := MapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8UintPtrErr failed")
	}

	r, _ = MapUint8UintPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8UintPtrErr failed")
	}

	_, err := MapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint64PtrErr failed")
	}

	r, _ = MapUint8Uint64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint64PtrErr failed")
	}

	_, err := MapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint32PtrErr failed")
	}

	r, _ = MapUint8Uint32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint32PtrErr failed")
	}

	_, err := MapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Uint16PtrErr failed")
	}

	r, _ = MapUint8Uint16PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Uint16PtrErr failed")
	}

	_, err := MapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint8 = 10
	var vi0 uint8

	expectedList := []*string{&vo10}
	newList, _ := MapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapUint8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Str failed")
	}

	r, _ = MapUint8StrPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8StrPtrErr failed")
	}
	_, err := MapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapUint8StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicUint8StrPtrErr(num *uint8) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapUint8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8
	var v3 uint8 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Bool failed")
	}

	r, _ = MapUint8BoolPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8BoolPtrErr failed")
	}
	_, err := MapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v3})
	if err == nil {
		t.Errorf("MapUint8BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Float32PtrErr failed")
	}

	r, _ = MapUint8Float32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Float32PtrErr failed")
	}

	_, err := MapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapUint8Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapUint8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapUint8Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapUint8Float64PtrErr failed")
	}

	r, _ = MapUint8Float64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("MapUint8Float64PtrErr failed")
	}

	_, err := MapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapUint8Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrIntPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int{&vo10}
	newList, _ := MapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrIntPtrErr failed")
	}

	r, _ = MapStrIntPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrIntPtrErr failed")
	}
	_, err := MapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrIntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrIntPtrErr(num *string) (*int, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r int = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrInt64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int64{&vo10}
	newList, _ := MapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt64PtrErr failed")
	}

	r, _ = MapStrInt64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt64PtrErr failed")
	}
	_, err := MapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt64PtrErr(num *string) (*int64, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r int64 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrInt32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int32{&vo10}
	newList, _ := MapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt32PtrErr failed")
	}

	r, _ = MapStrInt32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt32PtrErr failed")
	}
	_, err := MapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt32PtrErr(num *string) (*int32, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r int32 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrInt16PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int16{&vo10}
	newList, _ := MapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt16PtrErr failed")
	}

	r, _ = MapStrInt16PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt16PtrErr failed")
	}
	_, err := MapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt16PtrErr(num *string) (*int16, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r int16 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrInt8PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int8{&vo10}
	newList, _ := MapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrInt8PtrErr failed")
	}

	r, _ = MapStrInt8PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrInt8PtrErr failed")
	}
	_, err := MapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrInt8PtrErr(num *string) (*int8, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r int8 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrUintPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint{&vo10}
	newList, _ := MapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUintPtrErr failed")
	}

	r, _ = MapStrUintPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrUintPtrErr failed")
	}
	_, err := MapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrUintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUintPtrErr(num *string) (*uint, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r uint = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrUint64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint64{&vo10}
	newList, _ := MapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint64PtrErr failed")
	}

	r, _ = MapStrUint64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint64PtrErr failed")
	}
	_, err := MapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint64PtrErr(num *string) (*uint64, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r uint64 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrUint32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint32{&vo10}
	newList, _ := MapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint32PtrErr failed")
	}

	r, _ = MapStrUint32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint32PtrErr failed")
	}
	_, err := MapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint32PtrErr(num *string) (*uint32, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r uint32 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrUint16PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint16{&vo10}
	newList, _ := MapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint16PtrErr failed")
	}

	r, _ = MapStrUint16PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint16PtrErr failed")
	}
	_, err := MapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint16PtrErr(num *string) (*uint16, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r uint16 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrUint8PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint8{&vo10}
	newList, _ := MapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrUint8PtrErr failed")
	}

	r, _ = MapStrUint8PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrUint8PtrErr failed")
	}
	_, err := MapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrUint8PtrErr(num *string) (*uint8, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r uint8 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"
	var v3 string = "3"

	expectedList := []*bool{&vt, &vf}
	newList, _ := MapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapStrBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrBoolPtrErr failed")
	}

	r, _ = MapStrBoolPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrBoolPtrErr failed")
	}
	_, err := MapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v0, &v3})
	if err == nil {
		t.Errorf("MapStrBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapStrFloat32PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*float32{&vo10}
	newList, _ := MapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrFloat32PtrErr failed")
	}

	r, _ = MapStrFloat32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrFloat32PtrErr failed")
	}
	_, err := MapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat32PtrErr(num *string) (*float32, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r float32 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapStrFloat64PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*float64{&vo10}
	newList, _ := MapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapStrFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapStrFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapStrFloat64PtrErr failed")
	}

	r, _ = MapStrFloat64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("MapStrFloat64PtrErr failed")
	}
	_, err := MapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapStrFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicStrFloat64PtrErr(num *string) (*float64, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r float64 = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}

func TestMapBoolIntPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int

	// Test : someLogic
	expectedList := []*int{&v10, &v0}
	newList, _ := MapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolIntPtrErr failed")
	}

	r, _ = MapBoolIntPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolIntPtrErr failed")
	}
	_, err := MapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolIntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt64PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64

	// Test : someLogic
	expectedList := []*int64{&v10, &v0}
	newList, _ := MapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt64PtrErr failed")
	}

	r, _ = MapBoolInt64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt64PtrErr failed")
	}
	_, err := MapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt32PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32

	// Test : someLogic
	expectedList := []*int32{&v10, &v0}
	newList, _ := MapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt32PtrErr failed")
	}

	r, _ = MapBoolInt32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt32PtrErr failed")
	}
	_, err := MapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt16PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16

	// Test : someLogic
	expectedList := []*int16{&v10, &v0}
	newList, _ := MapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt16PtrErr failed")
	}

	r, _ = MapBoolInt16PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt16PtrErr failed")
	}
	_, err := MapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolInt8PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8

	// Test : someLogic
	expectedList := []*int8{&v10, &v0}
	newList, _ := MapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolInt8PtrErr failed")
	}

	r, _ = MapBoolInt8PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolInt8PtrErr failed")
	}
	_, err := MapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUintPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint

	// Test : someLogic
	expectedList := []*uint{&v10, &v0}
	newList, _ := MapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUintPtrErr failed")
	}

	r, _ = MapBoolUintPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUintPtrErr failed")
	}
	_, err := MapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolUintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint64PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64

	// Test : someLogic
	expectedList := []*uint64{&v10, &v0}
	newList, _ := MapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint64PtrErr failed")
	}

	r, _ = MapBoolUint64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint64PtrErr failed")
	}
	_, err := MapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint32PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32

	// Test : someLogic
	expectedList := []*uint32{&v10, &v0}
	newList, _ := MapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint32PtrErr failed")
	}

	r, _ = MapBoolUint32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint32PtrErr failed")
	}
	_, err := MapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint16PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16

	// Test : someLogic
	expectedList := []*uint16{&v10, &v0}
	newList, _ := MapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint16PtrErr failed")
	}

	r, _ = MapBoolUint16PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint16PtrErr failed")
	}
	_, err := MapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolUint8PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8

	// Test : someLogic
	expectedList := []*uint8{&v10, &v0}
	newList, _ := MapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolUint8PtrErr failed")
	}

	r, _ = MapBoolUint8PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolUint8PtrErr failed")
	}
	_, err := MapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolStrPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	// Test : someLogic
	expectedList := []*string{&v10, &v0}
	newList, _ := MapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolStrPtrErr failed")
	}

	r, _ = MapBoolStrPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolStrPtrErr failed")
	}
	_, err := MapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat32PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32

	// Test : someLogic
	expectedList := []*float32{&v10, &v0}
	newList, _ := MapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolFloat32PtrErr failed")
	}

	r, _ = MapBoolFloat32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolFloat32PtrErr failed")
	}
	_, err := MapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapBoolFloat64PtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64

	// Test : someLogic
	expectedList := []*float64{&v10, &v0}
	newList, _ := MapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapBoolFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapBoolFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolFloat64PtrErr failed")
	}

	r, _ = MapBoolFloat64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("MapBoolFloat64PtrErr failed")
	}
	_, err := MapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("MapBoolFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32IntPtrErr failed")
	}

	r, _ = MapFloat32IntPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32IntPtrErr failed")
	}

	_, err := MapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int64PtrErr failed")
	}

	r, _ = MapFloat32Int64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int64PtrErr failed")
	}

	_, err := MapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int32PtrErr failed")
	}

	r, _ = MapFloat32Int32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int32PtrErr failed")
	}

	_, err := MapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int16PtrErr failed")
	}

	r, _ = MapFloat32Int16PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int16PtrErr failed")
	}

	_, err := MapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Int8PtrErr failed")
	}

	r, _ = MapFloat32Int8PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Int8PtrErr failed")
	}

	_, err := MapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32UintPtrErr failed")
	}

	r, _ = MapFloat32UintPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32UintPtrErr failed")
	}

	_, err := MapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint64PtrErr failed")
	}

	r, _ = MapFloat32Uint64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint64PtrErr failed")
	}

	_, err := MapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint32PtrErr failed")
	}

	r, _ = MapFloat32Uint32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint32PtrErr failed")
	}

	_, err := MapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint16PtrErr failed")
	}

	r, _ = MapFloat32Uint16PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint16PtrErr failed")
	}

	_, err := MapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint8PtrErr failed")
	}

	r, _ = MapFloat32Uint8PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Uint8PtrErr failed")
	}

	_, err := MapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float32 = 10
	var vi0 float32

	expectedList := []*string{&vo10}
	newList, _ := MapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapFloat32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Str failed")
	}

	r, _ = MapFloat32StrPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32StrPtrErr failed")
	}
	_, err := MapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapFloat32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat32StrPtrErr(num *float32) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapFloat32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32
	var v3 float32 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Bool failed")
	}

	r, _ = MapFloat32BoolPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32BoolPtrErr failed")
	}
	_, err := MapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v3})
	if err == nil {
		t.Errorf("MapFloat32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := MapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat32Float64PtrErr failed")
	}

	r, _ = MapFloat32Float64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("MapFloat32Float64PtrErr failed")
	}

	_, err := MapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := MapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64IntPtrErr failed")
	}

	r, _ = MapFloat64IntPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64IntPtrErr failed")
	}

	_, err := MapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := MapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int64PtrErr failed")
	}

	r, _ = MapFloat64Int64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int64PtrErr failed")
	}

	_, err := MapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := MapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int32PtrErr failed")
	}

	r, _ = MapFloat64Int32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int32PtrErr failed")
	}

	_, err := MapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := MapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int16PtrErr failed")
	}

	r, _ = MapFloat64Int16PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int16PtrErr failed")
	}

	_, err := MapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := MapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Int8PtrErr failed")
	}

	r, _ = MapFloat64Int8PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Int8PtrErr failed")
	}

	_, err := MapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := MapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64UintPtrErr failed")
	}

	r, _ = MapFloat64UintPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64UintPtrErr failed")
	}

	_, err := MapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := MapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint64PtrErr failed")
	}

	r, _ = MapFloat64Uint64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint64PtrErr failed")
	}

	_, err := MapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := MapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint32PtrErr failed")
	}

	r, _ = MapFloat64Uint32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint32PtrErr failed")
	}

	_, err := MapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := MapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint16PtrErr failed")
	}

	r, _ = MapFloat64Uint16PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint16PtrErr failed")
	}

	_, err := MapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := MapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint8PtrErr failed")
	}

	r, _ = MapFloat64Uint8PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Uint8PtrErr failed")
	}

	_, err := MapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float64 = 10
	var vi0 float64

	expectedList := []*string{&vo10}
	newList, _ := MapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("MapFloat64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Str failed")
	}

	r, _ = MapFloat64StrPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64StrPtrErr failed")
	}
	_, err := MapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi0})
	if err == nil {
		t.Errorf("MapFloat64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogicFloat64StrPtrErr(num *float64) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}

func TestMapFloat64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64
	var v3 float64 = 3
	
	expectedList := []*bool{&vt, &vf}
	newList, _ := MapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64BoolErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Bool failed")
	}

	r, _ = MapFloat64BoolPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64BoolPtrErr failed")
	}
	_, err := MapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v3})
	if err == nil {
		t.Errorf("MapFloat64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func TestMapFloat64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := MapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("MapFloat64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := MapFloat64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapFloat64Float32PtrErr failed")
	}

	r, _ = MapFloat64Float32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("MapFloat64Float32PtrErr failed")
	}

	_, err := MapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("MapFloat64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
