package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntInt64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapIntInt64Ptr(notOneIntInt64Ptr, plusOneIntInt64Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt64Ptr failed")
	}

	if len(FilterMapIntInt64Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt64Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntInt64Ptr(num *int) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapIntInt32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapIntInt32Ptr(notOneIntInt32Ptr, plusOneIntInt32Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt32Ptr failed")
	}

	if len(FilterMapIntInt32Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt32Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntInt32Ptr(num *int) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapIntInt16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapIntInt16Ptr(notOneIntInt16Ptr, plusOneIntInt16Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt16Ptr failed")
	}

	if len(FilterMapIntInt16Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt16Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntInt16Ptr(num *int) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapIntInt8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapIntInt8Ptr(notOneIntInt8Ptr, plusOneIntInt8Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntInt8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntInt8Ptr failed")
	}

	if len(FilterMapIntInt8Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntInt8Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntInt8Ptr(num *int) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapIntUintPtr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapIntUintPtr(notOneIntUintPtr, plusOneIntUintPtr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUintPtr failed")
	}

	if len(FilterMapIntUintPtr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUintPtr(num *int) bool {
	return *num != 1
}
func plusOneIntUintPtr(num *int) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapIntUint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapIntUint64Ptr(notOneIntUint64Ptr, plusOneIntUint64Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint64Ptr failed")
	}

	if len(FilterMapIntUint64Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint64Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntUint64Ptr(num *int) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapIntUint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapIntUint32Ptr(notOneIntUint32Ptr, plusOneIntUint32Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint32Ptr failed")
	}

	if len(FilterMapIntUint32Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint32Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntUint32Ptr(num *int) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapIntUint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapIntUint16Ptr(notOneIntUint16Ptr, plusOneIntUint16Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint16Ptr failed")
	}

	if len(FilterMapIntUint16Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint16Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntUint16Ptr(num *int) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapIntUint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapIntUint8Ptr(notOneIntUint8Ptr, plusOneIntUint8Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntUint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntUint8Ptr failed")
	}

	if len(FilterMapIntUint8Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntUint8Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntUint8Ptr(num *int) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapIntStrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int = 1
	var iv10 int = 10
	expectedList := []*string{&ov10}
	newList := FilterMapIntStrPtr(notOneIntStrNumPtr, someLogicIntStrNumPtr, []*int{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapIntStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntStrPtr failed")
	}

	if len(FilterMapIntStrPtr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntStrNumPtr(num *int) bool {
	return *num != 1
}

func someLogicIntStrNumPtr(num *int) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapIntBoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int = 1
	var vi10 int = 10
	var vi0 int

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapIntBoolPtr(notOneIntBoolPtr, someLogicIntBoolPtr, []*int{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntBoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntBoolPtr failed")
	}

	if len(FilterMapIntBoolPtr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntBoolPtr(num *int) bool {
	return *num != 1
}

func someLogicIntBoolPtr(num *int) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapIntFloat32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapIntFloat32Ptr(notOneIntFloat32Ptr, plusOneIntFloat32Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntFloat32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntFloat32Ptr failed")
	}

	if len(FilterMapIntFloat32Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntFloat32Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntFloat32Ptr(num *int) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapIntFloat64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapIntFloat64Ptr(notOneIntFloat64Ptr, plusOneIntFloat64Ptr, []*int{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapIntFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapIntFloat64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapIntFloat64Ptr failed")
	}

	if len(FilterMapIntFloat64Ptr(nil, nil, []*int{})) > 0 {
		t.Errorf("FilterMapIntFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneIntFloat64Ptr(num *int) bool {
	return *num != 1
}
func plusOneIntFloat64Ptr(num *int) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapInt64IntPtr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapInt64IntPtr(notOneInt64IntPtr, plusOneInt64IntPtr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64IntPtr failed")
	}

	if len(FilterMapInt64IntPtr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64IntPtr(num *int64) bool {
	return *num != 1
}
func plusOneInt64IntPtr(num *int64) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapInt64Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapInt64Int32Ptr(notOneInt64Int32Ptr, plusOneInt64Int32Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int32Ptr failed")
	}

	if len(FilterMapInt64Int32Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int32Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Int32Ptr(num *int64) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapInt64Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapInt64Int16Ptr(notOneInt64Int16Ptr, plusOneInt64Int16Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int16Ptr failed")
	}

	if len(FilterMapInt64Int16Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int16Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Int16Ptr(num *int64) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapInt64Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapInt64Int8Ptr(notOneInt64Int8Ptr, plusOneInt64Int8Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Int8Ptr failed")
	}

	if len(FilterMapInt64Int8Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Int8Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Int8Ptr(num *int64) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapInt64UintPtr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapInt64UintPtr(notOneInt64UintPtr, plusOneInt64UintPtr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64UintPtr failed")
	}

	if len(FilterMapInt64UintPtr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64UintPtr(num *int64) bool {
	return *num != 1
}
func plusOneInt64UintPtr(num *int64) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapInt64Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapInt64Uint64Ptr(notOneInt64Uint64Ptr, plusOneInt64Uint64Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint64Ptr failed")
	}

	if len(FilterMapInt64Uint64Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint64Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Uint64Ptr(num *int64) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapInt64Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapInt64Uint32Ptr(notOneInt64Uint32Ptr, plusOneInt64Uint32Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint32Ptr failed")
	}

	if len(FilterMapInt64Uint32Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint32Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Uint32Ptr(num *int64) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapInt64Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapInt64Uint16Ptr(notOneInt64Uint16Ptr, plusOneInt64Uint16Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint16Ptr failed")
	}

	if len(FilterMapInt64Uint16Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint16Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Uint16Ptr(num *int64) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapInt64Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapInt64Uint8Ptr(notOneInt64Uint8Ptr, plusOneInt64Uint8Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Uint8Ptr failed")
	}

	if len(FilterMapInt64Uint8Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Uint8Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Uint8Ptr(num *int64) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapInt64StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int64 = 1
	var iv10 int64 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapInt64StrPtr(notOneInt64StrNumPtr, someLogicInt64StrNumPtr, []*int64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64StrPtr failed")
	}

	if len(FilterMapInt64StrPtr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64StrNumPtr(num *int64) bool {
	return *num != 1
}

func someLogicInt64StrNumPtr(num *int64) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapInt64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int64 = 1
	var vi10 int64 = 10
	var vi0 int64

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapInt64BoolPtr(notOneInt64BoolPtr, someLogicInt64BoolPtr, []*int64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64BoolPtr failed")
	}

	if len(FilterMapInt64BoolPtr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64BoolPtr(num *int64) bool {
	return *num != 1
}

func someLogicInt64BoolPtr(num *int64) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapInt64Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapInt64Float32Ptr(notOneInt64Float32Ptr, plusOneInt64Float32Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Float32Ptr failed")
	}

	if len(FilterMapInt64Float32Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Float32Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Float32Ptr(num *int64) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapInt64Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapInt64Float64Ptr(notOneInt64Float64Ptr, plusOneInt64Float64Ptr, []*int64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt64Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt64Float64Ptr failed")
	}

	if len(FilterMapInt64Float64Ptr(nil, nil, []*int64{})) > 0 {
		t.Errorf("FilterMapInt64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt64Float64Ptr(num *int64) bool {
	return *num != 1
}
func plusOneInt64Float64Ptr(num *int64) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapInt32IntPtr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapInt32IntPtr(notOneInt32IntPtr, plusOneInt32IntPtr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32IntPtr failed")
	}

	if len(FilterMapInt32IntPtr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32IntPtr(num *int32) bool {
	return *num != 1
}
func plusOneInt32IntPtr(num *int32) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapInt32Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapInt32Int64Ptr(notOneInt32Int64Ptr, plusOneInt32Int64Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int64Ptr failed")
	}

	if len(FilterMapInt32Int64Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int64Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Int64Ptr(num *int32) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapInt32Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapInt32Int16Ptr(notOneInt32Int16Ptr, plusOneInt32Int16Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int16Ptr failed")
	}

	if len(FilterMapInt32Int16Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int16Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Int16Ptr(num *int32) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapInt32Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapInt32Int8Ptr(notOneInt32Int8Ptr, plusOneInt32Int8Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Int8Ptr failed")
	}

	if len(FilterMapInt32Int8Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Int8Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Int8Ptr(num *int32) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapInt32UintPtr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapInt32UintPtr(notOneInt32UintPtr, plusOneInt32UintPtr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32UintPtr failed")
	}

	if len(FilterMapInt32UintPtr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32UintPtr(num *int32) bool {
	return *num != 1
}
func plusOneInt32UintPtr(num *int32) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapInt32Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapInt32Uint64Ptr(notOneInt32Uint64Ptr, plusOneInt32Uint64Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint64Ptr failed")
	}

	if len(FilterMapInt32Uint64Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint64Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Uint64Ptr(num *int32) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapInt32Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapInt32Uint32Ptr(notOneInt32Uint32Ptr, plusOneInt32Uint32Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint32Ptr failed")
	}

	if len(FilterMapInt32Uint32Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint32Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Uint32Ptr(num *int32) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapInt32Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapInt32Uint16Ptr(notOneInt32Uint16Ptr, plusOneInt32Uint16Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint16Ptr failed")
	}

	if len(FilterMapInt32Uint16Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint16Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Uint16Ptr(num *int32) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapInt32Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapInt32Uint8Ptr(notOneInt32Uint8Ptr, plusOneInt32Uint8Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Uint8Ptr failed")
	}

	if len(FilterMapInt32Uint8Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Uint8Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Uint8Ptr(num *int32) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapInt32StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int32 = 1
	var iv10 int32 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapInt32StrPtr(notOneInt32StrNumPtr, someLogicInt32StrNumPtr, []*int32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32StrPtr failed")
	}

	if len(FilterMapInt32StrPtr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32StrNumPtr(num *int32) bool {
	return *num != 1
}

func someLogicInt32StrNumPtr(num *int32) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapInt32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int32 = 1
	var vi10 int32 = 10
	var vi0 int32

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapInt32BoolPtr(notOneInt32BoolPtr, someLogicInt32BoolPtr, []*int32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32BoolPtr failed")
	}

	if len(FilterMapInt32BoolPtr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32BoolPtr(num *int32) bool {
	return *num != 1
}

func someLogicInt32BoolPtr(num *int32) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapInt32Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapInt32Float32Ptr(notOneInt32Float32Ptr, plusOneInt32Float32Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Float32Ptr failed")
	}

	if len(FilterMapInt32Float32Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Float32Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Float32Ptr(num *int32) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapInt32Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapInt32Float64Ptr(notOneInt32Float64Ptr, plusOneInt32Float64Ptr, []*int32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt32Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt32Float64Ptr failed")
	}

	if len(FilterMapInt32Float64Ptr(nil, nil, []*int32{})) > 0 {
		t.Errorf("FilterMapInt32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt32Float64Ptr(num *int32) bool {
	return *num != 1
}
func plusOneInt32Float64Ptr(num *int32) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapInt16IntPtr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapInt16IntPtr(notOneInt16IntPtr, plusOneInt16IntPtr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16IntPtr failed")
	}

	if len(FilterMapInt16IntPtr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16IntPtr(num *int16) bool {
	return *num != 1
}
func plusOneInt16IntPtr(num *int16) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapInt16Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapInt16Int64Ptr(notOneInt16Int64Ptr, plusOneInt16Int64Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int64Ptr failed")
	}

	if len(FilterMapInt16Int64Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int64Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Int64Ptr(num *int16) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapInt16Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapInt16Int32Ptr(notOneInt16Int32Ptr, plusOneInt16Int32Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int32Ptr failed")
	}

	if len(FilterMapInt16Int32Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int32Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Int32Ptr(num *int16) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapInt16Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapInt16Int8Ptr(notOneInt16Int8Ptr, plusOneInt16Int8Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Int8Ptr failed")
	}

	if len(FilterMapInt16Int8Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Int8Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Int8Ptr(num *int16) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapInt16UintPtr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapInt16UintPtr(notOneInt16UintPtr, plusOneInt16UintPtr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16UintPtr failed")
	}

	if len(FilterMapInt16UintPtr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16UintPtr(num *int16) bool {
	return *num != 1
}
func plusOneInt16UintPtr(num *int16) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapInt16Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapInt16Uint64Ptr(notOneInt16Uint64Ptr, plusOneInt16Uint64Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint64Ptr failed")
	}

	if len(FilterMapInt16Uint64Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint64Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Uint64Ptr(num *int16) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapInt16Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapInt16Uint32Ptr(notOneInt16Uint32Ptr, plusOneInt16Uint32Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint32Ptr failed")
	}

	if len(FilterMapInt16Uint32Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint32Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Uint32Ptr(num *int16) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapInt16Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapInt16Uint16Ptr(notOneInt16Uint16Ptr, plusOneInt16Uint16Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint16Ptr failed")
	}

	if len(FilterMapInt16Uint16Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint16Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Uint16Ptr(num *int16) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapInt16Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapInt16Uint8Ptr(notOneInt16Uint8Ptr, plusOneInt16Uint8Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Uint8Ptr failed")
	}

	if len(FilterMapInt16Uint8Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Uint8Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Uint8Ptr(num *int16) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapInt16StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int16 = 1
	var iv10 int16 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapInt16StrPtr(notOneInt16StrNumPtr, someLogicInt16StrNumPtr, []*int16{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16StrPtr failed")
	}

	if len(FilterMapInt16StrPtr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16StrNumPtr(num *int16) bool {
	return *num != 1
}

func someLogicInt16StrNumPtr(num *int16) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapInt16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int16 = 1
	var vi10 int16 = 10
	var vi0 int16

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapInt16BoolPtr(notOneInt16BoolPtr, someLogicInt16BoolPtr, []*int16{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16BoolPtr failed")
	}

	if len(FilterMapInt16BoolPtr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16BoolPtr(num *int16) bool {
	return *num != 1
}

func someLogicInt16BoolPtr(num *int16) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapInt16Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapInt16Float32Ptr(notOneInt16Float32Ptr, plusOneInt16Float32Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Float32Ptr failed")
	}

	if len(FilterMapInt16Float32Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Float32Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Float32Ptr(num *int16) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapInt16Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapInt16Float64Ptr(notOneInt16Float64Ptr, plusOneInt16Float64Ptr, []*int16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt16Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt16Float64Ptr failed")
	}

	if len(FilterMapInt16Float64Ptr(nil, nil, []*int16{})) > 0 {
		t.Errorf("FilterMapInt16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt16Float64Ptr(num *int16) bool {
	return *num != 1
}
func plusOneInt16Float64Ptr(num *int16) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapInt8IntPtr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapInt8IntPtr(notOneInt8IntPtr, plusOneInt8IntPtr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8IntPtr failed")
	}

	if len(FilterMapInt8IntPtr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8IntPtr(num *int8) bool {
	return *num != 1
}
func plusOneInt8IntPtr(num *int8) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapInt8Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapInt8Int64Ptr(notOneInt8Int64Ptr, plusOneInt8Int64Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int64Ptr failed")
	}

	if len(FilterMapInt8Int64Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int64Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Int64Ptr(num *int8) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapInt8Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapInt8Int32Ptr(notOneInt8Int32Ptr, plusOneInt8Int32Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int32Ptr failed")
	}

	if len(FilterMapInt8Int32Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int32Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Int32Ptr(num *int8) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapInt8Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapInt8Int16Ptr(notOneInt8Int16Ptr, plusOneInt8Int16Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Int16Ptr failed")
	}

	if len(FilterMapInt8Int16Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Int16Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Int16Ptr(num *int8) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapInt8UintPtr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapInt8UintPtr(notOneInt8UintPtr, plusOneInt8UintPtr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8UintPtr failed")
	}

	if len(FilterMapInt8UintPtr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8UintPtr(num *int8) bool {
	return *num != 1
}
func plusOneInt8UintPtr(num *int8) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapInt8Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapInt8Uint64Ptr(notOneInt8Uint64Ptr, plusOneInt8Uint64Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint64Ptr failed")
	}

	if len(FilterMapInt8Uint64Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint64Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Uint64Ptr(num *int8) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapInt8Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapInt8Uint32Ptr(notOneInt8Uint32Ptr, plusOneInt8Uint32Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint32Ptr failed")
	}

	if len(FilterMapInt8Uint32Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint32Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Uint32Ptr(num *int8) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapInt8Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapInt8Uint16Ptr(notOneInt8Uint16Ptr, plusOneInt8Uint16Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint16Ptr failed")
	}

	if len(FilterMapInt8Uint16Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint16Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Uint16Ptr(num *int8) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapInt8Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapInt8Uint8Ptr(notOneInt8Uint8Ptr, plusOneInt8Uint8Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Uint8Ptr failed")
	}

	if len(FilterMapInt8Uint8Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Uint8Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Uint8Ptr(num *int8) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapInt8StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 int8 = 1
	var iv10 int8 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapInt8StrPtr(notOneInt8StrNumPtr, someLogicInt8StrNumPtr, []*int8{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapInt8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8StrPtr failed")
	}

	if len(FilterMapInt8StrPtr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8StrNumPtr(num *int8) bool {
	return *num != 1
}

func someLogicInt8StrNumPtr(num *int8) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapInt8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 int8 = 1
	var vi10 int8 = 10
	var vi0 int8

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapInt8BoolPtr(notOneInt8BoolPtr, someLogicInt8BoolPtr, []*int8{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8BoolPtr failed")
	}

	if len(FilterMapInt8BoolPtr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8BoolPtr(num *int8) bool {
	return *num != 1
}

func someLogicInt8BoolPtr(num *int8) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapInt8Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapInt8Float32Ptr(notOneInt8Float32Ptr, plusOneInt8Float32Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Float32Ptr failed")
	}

	if len(FilterMapInt8Float32Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Float32Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Float32Ptr(num *int8) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapInt8Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapInt8Float64Ptr(notOneInt8Float64Ptr, plusOneInt8Float64Ptr, []*int8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapInt8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapInt8Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapInt8Float64Ptr failed")
	}

	if len(FilterMapInt8Float64Ptr(nil, nil, []*int8{})) > 0 {
		t.Errorf("FilterMapInt8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneInt8Float64Ptr(num *int8) bool {
	return *num != 1
}
func plusOneInt8Float64Ptr(num *int8) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapUintIntPtr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapUintIntPtr(notOneUintIntPtr, plusOneUintIntPtr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintIntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintIntPtr failed")
	}

	if len(FilterMapUintIntPtr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintIntPtr(num *uint) bool {
	return *num != 1
}
func plusOneUintIntPtr(num *uint) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapUintInt64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapUintInt64Ptr(notOneUintInt64Ptr, plusOneUintInt64Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt64Ptr failed")
	}

	if len(FilterMapUintInt64Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt64Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintInt64Ptr(num *uint) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapUintInt32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapUintInt32Ptr(notOneUintInt32Ptr, plusOneUintInt32Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt32Ptr failed")
	}

	if len(FilterMapUintInt32Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt32Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintInt32Ptr(num *uint) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapUintInt16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapUintInt16Ptr(notOneUintInt16Ptr, plusOneUintInt16Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt16Ptr failed")
	}

	if len(FilterMapUintInt16Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt16Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintInt16Ptr(num *uint) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapUintInt8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapUintInt8Ptr(notOneUintInt8Ptr, plusOneUintInt8Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintInt8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintInt8Ptr failed")
	}

	if len(FilterMapUintInt8Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintInt8Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintInt8Ptr(num *uint) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapUintUint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapUintUint64Ptr(notOneUintUint64Ptr, plusOneUintUint64Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint64Ptr failed")
	}

	if len(FilterMapUintUint64Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint64Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintUint64Ptr(num *uint) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapUintUint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapUintUint32Ptr(notOneUintUint32Ptr, plusOneUintUint32Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint32Ptr failed")
	}

	if len(FilterMapUintUint32Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint32Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintUint32Ptr(num *uint) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapUintUint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapUintUint16Ptr(notOneUintUint16Ptr, plusOneUintUint16Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint16Ptr failed")
	}

	if len(FilterMapUintUint16Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint16Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintUint16Ptr(num *uint) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapUintUint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapUintUint8Ptr(notOneUintUint8Ptr, plusOneUintUint8Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintUint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintUint8Ptr failed")
	}

	if len(FilterMapUintUint8Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintUint8Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintUint8Ptr(num *uint) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapUintStrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint = 1
	var iv10 uint = 10
	expectedList := []*string{&ov10}
	newList := FilterMapUintStrPtr(notOneUintStrNumPtr, someLogicUintStrNumPtr, []*uint{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUintStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintStrPtr failed")
	}

	if len(FilterMapUintStrPtr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintStrNumPtr(num *uint) bool {
	return *num != 1
}

func someLogicUintStrNumPtr(num *uint) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapUintBoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint = 1
	var vi10 uint = 10
	var vi0 uint

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapUintBoolPtr(notOneUintBoolPtr, someLogicUintBoolPtr, []*uint{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintBoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintBoolPtr failed")
	}

	if len(FilterMapUintBoolPtr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintBoolPtr(num *uint) bool {
	return *num != 1
}

func someLogicUintBoolPtr(num *uint) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapUintFloat32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapUintFloat32Ptr(notOneUintFloat32Ptr, plusOneUintFloat32Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintFloat32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintFloat32Ptr failed")
	}

	if len(FilterMapUintFloat32Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintFloat32Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintFloat32Ptr(num *uint) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapUintFloat64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapUintFloat64Ptr(notOneUintFloat64Ptr, plusOneUintFloat64Ptr, []*uint{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUintFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUintFloat64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUintFloat64Ptr failed")
	}

	if len(FilterMapUintFloat64Ptr(nil, nil, []*uint{})) > 0 {
		t.Errorf("FilterMapUintFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUintFloat64Ptr(num *uint) bool {
	return *num != 1
}
func plusOneUintFloat64Ptr(num *uint) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapUint64IntPtr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapUint64IntPtr(notOneUint64IntPtr, plusOneUint64IntPtr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64IntPtr failed")
	}

	if len(FilterMapUint64IntPtr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64IntPtr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64IntPtr(num *uint64) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapUint64Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapUint64Int64Ptr(notOneUint64Int64Ptr, plusOneUint64Int64Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int64Ptr failed")
	}

	if len(FilterMapUint64Int64Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int64Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Int64Ptr(num *uint64) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapUint64Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapUint64Int32Ptr(notOneUint64Int32Ptr, plusOneUint64Int32Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int32Ptr failed")
	}

	if len(FilterMapUint64Int32Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int32Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Int32Ptr(num *uint64) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapUint64Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapUint64Int16Ptr(notOneUint64Int16Ptr, plusOneUint64Int16Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int16Ptr failed")
	}

	if len(FilterMapUint64Int16Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int16Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Int16Ptr(num *uint64) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapUint64Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapUint64Int8Ptr(notOneUint64Int8Ptr, plusOneUint64Int8Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Int8Ptr failed")
	}

	if len(FilterMapUint64Int8Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Int8Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Int8Ptr(num *uint64) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapUint64UintPtr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapUint64UintPtr(notOneUint64UintPtr, plusOneUint64UintPtr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64UintPtr failed")
	}

	if len(FilterMapUint64UintPtr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64UintPtr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64UintPtr(num *uint64) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapUint64Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapUint64Uint32Ptr(notOneUint64Uint32Ptr, plusOneUint64Uint32Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint32Ptr failed")
	}

	if len(FilterMapUint64Uint32Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint32Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Uint32Ptr(num *uint64) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapUint64Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapUint64Uint16Ptr(notOneUint64Uint16Ptr, plusOneUint64Uint16Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint16Ptr failed")
	}

	if len(FilterMapUint64Uint16Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint16Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Uint16Ptr(num *uint64) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapUint64Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapUint64Uint8Ptr(notOneUint64Uint8Ptr, plusOneUint64Uint8Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Uint8Ptr failed")
	}

	if len(FilterMapUint64Uint8Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Uint8Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Uint8Ptr(num *uint64) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapUint64StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint64 = 1
	var iv10 uint64 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapUint64StrPtr(notOneUint64StrNumPtr, someLogicUint64StrNumPtr, []*uint64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64StrPtr failed")
	}

	if len(FilterMapUint64StrPtr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64StrNumPtr(num *uint64) bool {
	return *num != 1
}

func someLogicUint64StrNumPtr(num *uint64) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapUint64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint64 = 1
	var vi10 uint64 = 10
	var vi0 uint64

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapUint64BoolPtr(notOneUint64BoolPtr, someLogicUint64BoolPtr, []*uint64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64BoolPtr failed")
	}

	if len(FilterMapUint64BoolPtr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64BoolPtr(num *uint64) bool {
	return *num != 1
}

func someLogicUint64BoolPtr(num *uint64) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapUint64Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapUint64Float32Ptr(notOneUint64Float32Ptr, plusOneUint64Float32Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Float32Ptr failed")
	}

	if len(FilterMapUint64Float32Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Float32Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Float32Ptr(num *uint64) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapUint64Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapUint64Float64Ptr(notOneUint64Float64Ptr, plusOneUint64Float64Ptr, []*uint64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint64Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint64Float64Ptr failed")
	}

	if len(FilterMapUint64Float64Ptr(nil, nil, []*uint64{})) > 0 {
		t.Errorf("FilterMapUint64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint64Float64Ptr(num *uint64) bool {
	return *num != 1
}
func plusOneUint64Float64Ptr(num *uint64) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapUint32IntPtr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapUint32IntPtr(notOneUint32IntPtr, plusOneUint32IntPtr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32IntPtr failed")
	}

	if len(FilterMapUint32IntPtr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32IntPtr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32IntPtr(num *uint32) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapUint32Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapUint32Int64Ptr(notOneUint32Int64Ptr, plusOneUint32Int64Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int64Ptr failed")
	}

	if len(FilterMapUint32Int64Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int64Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Int64Ptr(num *uint32) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapUint32Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapUint32Int32Ptr(notOneUint32Int32Ptr, plusOneUint32Int32Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int32Ptr failed")
	}

	if len(FilterMapUint32Int32Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int32Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Int32Ptr(num *uint32) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapUint32Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapUint32Int16Ptr(notOneUint32Int16Ptr, plusOneUint32Int16Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int16Ptr failed")
	}

	if len(FilterMapUint32Int16Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int16Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Int16Ptr(num *uint32) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapUint32Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapUint32Int8Ptr(notOneUint32Int8Ptr, plusOneUint32Int8Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Int8Ptr failed")
	}

	if len(FilterMapUint32Int8Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Int8Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Int8Ptr(num *uint32) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapUint32UintPtr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapUint32UintPtr(notOneUint32UintPtr, plusOneUint32UintPtr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32UintPtr failed")
	}

	if len(FilterMapUint32UintPtr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32UintPtr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32UintPtr(num *uint32) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapUint32Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapUint32Uint64Ptr(notOneUint32Uint64Ptr, plusOneUint32Uint64Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint64Ptr failed")
	}

	if len(FilterMapUint32Uint64Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint64Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Uint64Ptr(num *uint32) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapUint32Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapUint32Uint16Ptr(notOneUint32Uint16Ptr, plusOneUint32Uint16Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint16Ptr failed")
	}

	if len(FilterMapUint32Uint16Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint16Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Uint16Ptr(num *uint32) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapUint32Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapUint32Uint8Ptr(notOneUint32Uint8Ptr, plusOneUint32Uint8Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Uint8Ptr failed")
	}

	if len(FilterMapUint32Uint8Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Uint8Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Uint8Ptr(num *uint32) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapUint32StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint32 = 1
	var iv10 uint32 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapUint32StrPtr(notOneUint32StrNumPtr, someLogicUint32StrNumPtr, []*uint32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32StrPtr failed")
	}

	if len(FilterMapUint32StrPtr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32StrNumPtr(num *uint32) bool {
	return *num != 1
}

func someLogicUint32StrNumPtr(num *uint32) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapUint32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint32 = 1
	var vi10 uint32 = 10
	var vi0 uint32

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapUint32BoolPtr(notOneUint32BoolPtr, someLogicUint32BoolPtr, []*uint32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32BoolPtr failed")
	}

	if len(FilterMapUint32BoolPtr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32BoolPtr(num *uint32) bool {
	return *num != 1
}

func someLogicUint32BoolPtr(num *uint32) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapUint32Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapUint32Float32Ptr(notOneUint32Float32Ptr, plusOneUint32Float32Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Float32Ptr failed")
	}

	if len(FilterMapUint32Float32Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Float32Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Float32Ptr(num *uint32) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapUint32Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapUint32Float64Ptr(notOneUint32Float64Ptr, plusOneUint32Float64Ptr, []*uint32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint32Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint32Float64Ptr failed")
	}

	if len(FilterMapUint32Float64Ptr(nil, nil, []*uint32{})) > 0 {
		t.Errorf("FilterMapUint32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint32Float64Ptr(num *uint32) bool {
	return *num != 1
}
func plusOneUint32Float64Ptr(num *uint32) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapUint16IntPtr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapUint16IntPtr(notOneUint16IntPtr, plusOneUint16IntPtr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16IntPtr failed")
	}

	if len(FilterMapUint16IntPtr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16IntPtr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16IntPtr(num *uint16) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapUint16Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapUint16Int64Ptr(notOneUint16Int64Ptr, plusOneUint16Int64Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int64Ptr failed")
	}

	if len(FilterMapUint16Int64Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int64Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Int64Ptr(num *uint16) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapUint16Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapUint16Int32Ptr(notOneUint16Int32Ptr, plusOneUint16Int32Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int32Ptr failed")
	}

	if len(FilterMapUint16Int32Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int32Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Int32Ptr(num *uint16) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapUint16Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapUint16Int16Ptr(notOneUint16Int16Ptr, plusOneUint16Int16Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int16Ptr failed")
	}

	if len(FilterMapUint16Int16Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int16Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Int16Ptr(num *uint16) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapUint16Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapUint16Int8Ptr(notOneUint16Int8Ptr, plusOneUint16Int8Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Int8Ptr failed")
	}

	if len(FilterMapUint16Int8Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Int8Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Int8Ptr(num *uint16) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapUint16UintPtr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapUint16UintPtr(notOneUint16UintPtr, plusOneUint16UintPtr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16UintPtr failed")
	}

	if len(FilterMapUint16UintPtr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16UintPtr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16UintPtr(num *uint16) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapUint16Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapUint16Uint64Ptr(notOneUint16Uint64Ptr, plusOneUint16Uint64Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint64Ptr failed")
	}

	if len(FilterMapUint16Uint64Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint64Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Uint64Ptr(num *uint16) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapUint16Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapUint16Uint32Ptr(notOneUint16Uint32Ptr, plusOneUint16Uint32Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint32Ptr failed")
	}

	if len(FilterMapUint16Uint32Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint32Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Uint32Ptr(num *uint16) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapUint16Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapUint16Uint8Ptr(notOneUint16Uint8Ptr, plusOneUint16Uint8Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Uint8Ptr failed")
	}

	if len(FilterMapUint16Uint8Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Uint8Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Uint8Ptr(num *uint16) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapUint16StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint16 = 1
	var iv10 uint16 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapUint16StrPtr(notOneUint16StrNumPtr, someLogicUint16StrNumPtr, []*uint16{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16StrPtr failed")
	}

	if len(FilterMapUint16StrPtr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16StrNumPtr(num *uint16) bool {
	return *num != 1
}

func someLogicUint16StrNumPtr(num *uint16) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapUint16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint16 = 1
	var vi10 uint16 = 10
	var vi0 uint16

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapUint16BoolPtr(notOneUint16BoolPtr, someLogicUint16BoolPtr, []*uint16{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16BoolPtr failed")
	}

	if len(FilterMapUint16BoolPtr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16BoolPtr(num *uint16) bool {
	return *num != 1
}

func someLogicUint16BoolPtr(num *uint16) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapUint16Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapUint16Float32Ptr(notOneUint16Float32Ptr, plusOneUint16Float32Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Float32Ptr failed")
	}

	if len(FilterMapUint16Float32Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Float32Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Float32Ptr(num *uint16) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapUint16Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapUint16Float64Ptr(notOneUint16Float64Ptr, plusOneUint16Float64Ptr, []*uint16{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint16Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint16Float64Ptr failed")
	}

	if len(FilterMapUint16Float64Ptr(nil, nil, []*uint16{})) > 0 {
		t.Errorf("FilterMapUint16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint16Float64Ptr(num *uint16) bool {
	return *num != 1
}
func plusOneUint16Float64Ptr(num *uint16) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapUint8IntPtr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapUint8IntPtr(notOneUint8IntPtr, plusOneUint8IntPtr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8IntPtr failed")
	}

	if len(FilterMapUint8IntPtr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8IntPtr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8IntPtr(num *uint8) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapUint8Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapUint8Int64Ptr(notOneUint8Int64Ptr, plusOneUint8Int64Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int64Ptr failed")
	}

	if len(FilterMapUint8Int64Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int64Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Int64Ptr(num *uint8) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapUint8Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapUint8Int32Ptr(notOneUint8Int32Ptr, plusOneUint8Int32Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int32Ptr failed")
	}

	if len(FilterMapUint8Int32Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int32Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Int32Ptr(num *uint8) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapUint8Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapUint8Int16Ptr(notOneUint8Int16Ptr, plusOneUint8Int16Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int16Ptr failed")
	}

	if len(FilterMapUint8Int16Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int16Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Int16Ptr(num *uint8) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapUint8Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapUint8Int8Ptr(notOneUint8Int8Ptr, plusOneUint8Int8Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Int8Ptr failed")
	}

	if len(FilterMapUint8Int8Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Int8Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Int8Ptr(num *uint8) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapUint8UintPtr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapUint8UintPtr(notOneUint8UintPtr, plusOneUint8UintPtr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8UintPtr failed")
	}

	if len(FilterMapUint8UintPtr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8UintPtr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8UintPtr(num *uint8) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapUint8Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapUint8Uint64Ptr(notOneUint8Uint64Ptr, plusOneUint8Uint64Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint64Ptr failed")
	}

	if len(FilterMapUint8Uint64Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint64Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Uint64Ptr(num *uint8) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapUint8Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapUint8Uint32Ptr(notOneUint8Uint32Ptr, plusOneUint8Uint32Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint32Ptr failed")
	}

	if len(FilterMapUint8Uint32Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint32Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Uint32Ptr(num *uint8) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapUint8Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapUint8Uint16Ptr(notOneUint8Uint16Ptr, plusOneUint8Uint16Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Uint16Ptr failed")
	}

	if len(FilterMapUint8Uint16Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Uint16Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Uint16Ptr(num *uint8) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapUint8StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 uint8 = 1
	var iv10 uint8 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapUint8StrPtr(notOneUint8StrNumPtr, someLogicUint8StrNumPtr, []*uint8{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapUint8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8StrPtr failed")
	}

	if len(FilterMapUint8StrPtr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8StrNumPtr(num *uint8) bool {
	return *num != 1
}

func someLogicUint8StrNumPtr(num *uint8) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapUint8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 uint8 = 1
	var vi10 uint8 = 10
	var vi0 uint8

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapUint8BoolPtr(notOneUint8BoolPtr, someLogicUint8BoolPtr, []*uint8{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8BoolPtr failed")
	}

	if len(FilterMapUint8BoolPtr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8BoolPtr(num *uint8) bool {
	return *num != 1
}

func someLogicUint8BoolPtr(num *uint8) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapUint8Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapUint8Float32Ptr(notOneUint8Float32Ptr, plusOneUint8Float32Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Float32Ptr failed")
	}

	if len(FilterMapUint8Float32Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Float32Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Float32Ptr(num *uint8) *float32 {
	c := float32(*num + 1)
	return &c
}

func TestFilterMapUint8Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapUint8Float64Ptr(notOneUint8Float64Ptr, plusOneUint8Float64Ptr, []*uint8{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapUint8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapUint8Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapUint8Float64Ptr failed")
	}

	if len(FilterMapUint8Float64Ptr(nil, nil, []*uint8{})) > 0 {
		t.Errorf("FilterMapUint8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneUint8Float64Ptr(num *uint8) bool {
	return *num != 1
}
func plusOneUint8Float64Ptr(num *uint8) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapStrIntPtr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*int{&vo10}
	newList := FilterMapStrIntPtr(notOneStrIntStrPtr, someLogicStrIntStrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrIntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrIntPtr failed")
	}

	if len(FilterMapStrIntPtr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrIntStrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrIntStrPtr(num *string) *int {
	var r int = int(0)
	if *num == "ten" {
		r = int(10)
		return &r
	}
	return &r
}

func TestFilterMapStrInt64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*int64{&vo10}
	newList := FilterMapStrInt64Ptr(notOneStrInt64StrPtr, someLogicStrInt64StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt64Ptr failed")
	}

	if len(FilterMapStrInt64Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt64StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrInt64StrPtr(num *string) *int64 {
	var r int64 = int64(0)
	if *num == "ten" {
		r = int64(10)
		return &r
	}
	return &r
}

func TestFilterMapStrInt32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*int32{&vo10}
	newList := FilterMapStrInt32Ptr(notOneStrInt32StrPtr, someLogicStrInt32StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt32Ptr failed")
	}

	if len(FilterMapStrInt32Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt32StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrInt32StrPtr(num *string) *int32 {
	var r int32 = int32(0)
	if *num == "ten" {
		r = int32(10)
		return &r
	}
	return &r
}

func TestFilterMapStrInt16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*int16{&vo10}
	newList := FilterMapStrInt16Ptr(notOneStrInt16StrPtr, someLogicStrInt16StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt16Ptr failed")
	}

	if len(FilterMapStrInt16Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt16StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrInt16StrPtr(num *string) *int16 {
	var r int16 = int16(0)
	if *num == "ten" {
		r = int16(10)
		return &r
	}
	return &r
}

func TestFilterMapStrInt8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*int8{&vo10}
	newList := FilterMapStrInt8Ptr(notOneStrInt8StrPtr, someLogicStrInt8StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrInt8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrInt8Ptr failed")
	}

	if len(FilterMapStrInt8Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrInt8StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrInt8StrPtr(num *string) *int8 {
	var r int8 = int8(0)
	if *num == "ten" {
		r = int8(10)
		return &r
	}
	return &r
}

func TestFilterMapStrUintPtr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*uint{&vo10}
	newList := FilterMapStrUintPtr(notOneStrUintStrPtr, someLogicStrUintStrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUintPtr failed")
	}

	if len(FilterMapStrUintPtr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUintStrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrUintStrPtr(num *string) *uint {
	var r uint = uint(0)
	if *num == "ten" {
		r = uint(10)
		return &r
	}
	return &r
}

func TestFilterMapStrUint64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*uint64{&vo10}
	newList := FilterMapStrUint64Ptr(notOneStrUint64StrPtr, someLogicStrUint64StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint64Ptr failed")
	}

	if len(FilterMapStrUint64Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint64StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrUint64StrPtr(num *string) *uint64 {
	var r uint64 = uint64(0)
	if *num == "ten" {
		r = uint64(10)
		return &r
	}
	return &r
}

func TestFilterMapStrUint32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*uint32{&vo10}
	newList := FilterMapStrUint32Ptr(notOneStrUint32StrPtr, someLogicStrUint32StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint32Ptr failed")
	}

	if len(FilterMapStrUint32Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint32StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrUint32StrPtr(num *string) *uint32 {
	var r uint32 = uint32(0)
	if *num == "ten" {
		r = uint32(10)
		return &r
	}
	return &r
}

func TestFilterMapStrUint16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*uint16{&vo10}
	newList := FilterMapStrUint16Ptr(notOneStrUint16StrPtr, someLogicStrUint16StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint16Ptr failed")
	}

	if len(FilterMapStrUint16Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint16StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrUint16StrPtr(num *string) *uint16 {
	var r uint16 = uint16(0)
	if *num == "ten" {
		r = uint16(10)
		return &r
	}
	return &r
}

func TestFilterMapStrUint8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*uint8{&vo10}
	newList := FilterMapStrUint8Ptr(notOneStrUint8StrPtr, someLogicStrUint8StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrUint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrUint8Ptr failed")
	}

	if len(FilterMapStrUint8Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrUint8StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrUint8StrPtr(num *string) *uint8 {
	var r uint8 = uint8(0)
	if *num == "ten" {
		r = uint8(10)
		return &r
	}
	return &r
}

func TestFilterMapStrBoolPtr(t *testing.T) {
	// Test : someLogic

	var vto bool = true
	var vfo bool = false

	var vi1 string = "1"
	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapStrBoolPtr(notOneStrBoolPtr, someLogicStrBoolPtr, []*string{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapStrBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrBoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrBoolPtr failed")
	}

	if len(FilterMapStrBoolPtr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrBoolPtr(num *string) bool {
	return *num != "1"
}

func someLogicStrBoolPtr(num *string) *bool {
	var t bool = true
	var f bool = false

	if *num == "10" {
		return &t
	}
	return &f
}

func TestFilterMapStrFloat32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*float32{&vo10}
	newList := FilterMapStrFloat32Ptr(notOneStrFloat32StrPtr, someLogicStrFloat32StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrFloat32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrFloat32Ptr failed")
	}

	if len(FilterMapStrFloat32Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrFloat32StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrFloat32StrPtr(num *string) *float32 {
	var r float32 = float32(0)
	if *num == "ten" {
		r = float32(10)
		return &r
	}
	return &r
}

func TestFilterMapStrFloat64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vOne string = "one"
	var vTen string = "ten"

	expectedList := []*float64{&vo10}
	newList := FilterMapStrFloat64Ptr(notOneStrFloat64StrPtr, someLogicStrFloat64StrPtr, []*string{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapStrFloat64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrFloat64Ptr failed")
	}

	if len(FilterMapStrFloat64Ptr(nil, nil, []*string{})) > 0 {
		t.Errorf("FilterMapStrFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneStrFloat64StrPtr(num *string) bool {
	return *num != "one"
}

func someLogicStrFloat64StrPtr(num *string) *float64 {
	var r float64 = float64(0)
	if *num == "ten" {
		r = float64(10)
		return &r
	}
	return &r
}

func TestFilterMapBoolIntPtr(t *testing.T) {
	// Test : someLogic

	var vo10 int = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*int{&vo10, &vo10}
	newList := FilterMapBoolIntPtr(notOneBoolIntPtr, someLogicBoolIntPtr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolIntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt failed")
	}

	if len(FilterMapBoolIntPtr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolIntPtr(num *bool) bool {
	return *num == true
}

func someLogicBoolIntPtr(num *bool) *int {
	var v10 int = 10
	var v0 int = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolInt64Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 int64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*int64{&vo10, &vo10}
	newList := FilterMapBoolInt64Ptr(notOneBoolInt64Ptr, someLogicBoolInt64Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt64 failed")
	}

	if len(FilterMapBoolInt64Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt64Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolInt64Ptr(num *bool) *int64 {
	var v10 int64 = 10
	var v0 int64 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolInt32Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 int32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*int32{&vo10, &vo10}
	newList := FilterMapBoolInt32Ptr(notOneBoolInt32Ptr, someLogicBoolInt32Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt32 failed")
	}

	if len(FilterMapBoolInt32Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt32Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolInt32Ptr(num *bool) *int32 {
	var v10 int32 = 10
	var v0 int32 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolInt16Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 int16 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*int16{&vo10, &vo10}
	newList := FilterMapBoolInt16Ptr(notOneBoolInt16Ptr, someLogicBoolInt16Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt16 failed")
	}

	if len(FilterMapBoolInt16Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt16Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolInt16Ptr(num *bool) *int16 {
	var v10 int16 = 10
	var v0 int16 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolInt8Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 int8 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*int8{&vo10, &vo10}
	newList := FilterMapBoolInt8Ptr(notOneBoolInt8Ptr, someLogicBoolInt8Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolInt8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolInt8 failed")
	}

	if len(FilterMapBoolInt8Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolInt8Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolInt8Ptr(num *bool) *int8 {
	var v10 int8 = 10
	var v0 int8 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolUintPtr(t *testing.T) {
	// Test : someLogic

	var vo10 uint = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*uint{&vo10, &vo10}
	newList := FilterMapBoolUintPtr(notOneBoolUintPtr, someLogicBoolUintPtr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint failed")
	}

	if len(FilterMapBoolUintPtr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUintPtr(num *bool) bool {
	return *num == true
}

func someLogicBoolUintPtr(num *bool) *uint {
	var v10 uint = 10
	var v0 uint = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolUint64Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 uint64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*uint64{&vo10, &vo10}
	newList := FilterMapBoolUint64Ptr(notOneBoolUint64Ptr, someLogicBoolUint64Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint64 failed")
	}

	if len(FilterMapBoolUint64Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint64Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolUint64Ptr(num *bool) *uint64 {
	var v10 uint64 = 10
	var v0 uint64 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolUint32Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 uint32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*uint32{&vo10, &vo10}
	newList := FilterMapBoolUint32Ptr(notOneBoolUint32Ptr, someLogicBoolUint32Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint32 failed")
	}

	if len(FilterMapBoolUint32Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint32Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolUint32Ptr(num *bool) *uint32 {
	var v10 uint32 = 10
	var v0 uint32 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolUint16Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 uint16 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*uint16{&vo10, &vo10}
	newList := FilterMapBoolUint16Ptr(notOneBoolUint16Ptr, someLogicBoolUint16Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint16 failed")
	}

	if len(FilterMapBoolUint16Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint16Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolUint16Ptr(num *bool) *uint16 {
	var v10 uint16 = 10
	var v0 uint16 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolUint8Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 uint8 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*uint8{&vo10, &vo10}
	newList := FilterMapBoolUint8Ptr(notOneBoolUint8Ptr, someLogicBoolUint8Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolUint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolUint8 failed")
	}

	if len(FilterMapBoolUint8Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolUint8Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolUint8Ptr(num *bool) *uint8 {
	var v10 uint8 = 10
	var v0 uint8 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolStrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"

	var vit bool = true
	var vif bool = false

	expectedList := []*string{&vo10, &vo10}
	newList := FilterMapBoolStrPtr(notOneBoolStrPtr, someLogicBoolStrPtr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolStrPtr failed")
	}

	if len(FilterMapBoolStrPtr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolStrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolStrPtr(num *bool) bool {
	return *num == true
}

func someLogicBoolStrPtr(num *bool) *string {
	var v10 string = "10"
	var v0 string = "0"

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolFloat32Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 float32 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*float32{&vo10, &vo10}
	newList := FilterMapBoolFloat32Ptr(notOneBoolFloat32Ptr, someLogicBoolFloat32Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolFloat32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolFloat32 failed")
	}

	if len(FilterMapBoolFloat32Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolFloat32Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolFloat32Ptr(num *bool) *float32 {
	var v10 float32 = 10
	var v0 float32 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapBoolFloat64Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 float64 = 10

	var vit bool = true
	var vif bool = false

	expectedList := []*float64{&vo10, &vo10}
	newList := FilterMapBoolFloat64Ptr(notOneBoolFloat64Ptr, someLogicBoolFloat64Ptr, []*bool{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapBoolFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapBoolFloat64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapBoolFloat64 failed")
	}

	if len(FilterMapBoolFloat64Ptr(nil, nil, []*bool{})) > 0 {
		t.Errorf("FilterMapBoolFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneBoolFloat64Ptr(num *bool) bool {
	return *num == true
}

func someLogicBoolFloat64Ptr(num *bool) *float64 {
	var v10 float64 = 10
	var v0 float64 = 0

	if *num == true {
		return &v10
	}
	return &v0
}

func TestFilterMapFloat32IntPtr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapFloat32IntPtr(notOneFloat32IntPtr, plusOneFloat32IntPtr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32IntPtr failed")
	}

	if len(FilterMapFloat32IntPtr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32IntPtr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32IntPtr(num *float32) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapFloat32Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapFloat32Int64Ptr(notOneFloat32Int64Ptr, plusOneFloat32Int64Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int64Ptr failed")
	}

	if len(FilterMapFloat32Int64Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int64Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Int64Ptr(num *float32) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapFloat32Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapFloat32Int32Ptr(notOneFloat32Int32Ptr, plusOneFloat32Int32Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int32Ptr failed")
	}

	if len(FilterMapFloat32Int32Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int32Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Int32Ptr(num *float32) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapFloat32Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapFloat32Int16Ptr(notOneFloat32Int16Ptr, plusOneFloat32Int16Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int16Ptr failed")
	}

	if len(FilterMapFloat32Int16Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int16Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Int16Ptr(num *float32) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapFloat32Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapFloat32Int8Ptr(notOneFloat32Int8Ptr, plusOneFloat32Int8Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Int8Ptr failed")
	}

	if len(FilterMapFloat32Int8Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Int8Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Int8Ptr(num *float32) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapFloat32UintPtr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapFloat32UintPtr(notOneFloat32UintPtr, plusOneFloat32UintPtr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32UintPtr failed")
	}

	if len(FilterMapFloat32UintPtr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32UintPtr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32UintPtr(num *float32) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapFloat32Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapFloat32Uint64Ptr(notOneFloat32Uint64Ptr, plusOneFloat32Uint64Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint64Ptr failed")
	}

	if len(FilterMapFloat32Uint64Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint64Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Uint64Ptr(num *float32) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapFloat32Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapFloat32Uint32Ptr(notOneFloat32Uint32Ptr, plusOneFloat32Uint32Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint32Ptr failed")
	}

	if len(FilterMapFloat32Uint32Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint32Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Uint32Ptr(num *float32) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapFloat32Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapFloat32Uint16Ptr(notOneFloat32Uint16Ptr, plusOneFloat32Uint16Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint16Ptr failed")
	}

	if len(FilterMapFloat32Uint16Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint16Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Uint16Ptr(num *float32) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapFloat32Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapFloat32Uint8Ptr(notOneFloat32Uint8Ptr, plusOneFloat32Uint8Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Uint8Ptr failed")
	}

	if len(FilterMapFloat32Uint8Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Uint8Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Uint8Ptr(num *float32) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapFloat32StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float32 = 1
	var iv10 float32 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapFloat32StrPtr(notOneFloat32StrNumPtr, someLogicFloat32StrNumPtr, []*float32{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapFloat32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32StrPtr failed")
	}

	if len(FilterMapFloat32StrPtr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32StrNumPtr(num *float32) bool {
	return *num != 1
}

func someLogicFloat32StrNumPtr(num *float32) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapFloat32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float32 = 1
	var vi10 float32 = 10
	var vi0 float32

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapFloat32BoolPtr(notOneFloat32BoolPtr, someLogicFloat32BoolPtr, []*float32{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32BoolPtr failed")
	}

	if len(FilterMapFloat32BoolPtr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32BoolPtr(num *float32) bool {
	return *num != 1
}

func someLogicFloat32BoolPtr(num *float32) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapFloat32Float64Ptr(t *testing.T) {
	// Test : some logic
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var vo3 float64 = 3
	var vo4 float64 = 4

	expectedList := []*float64{&vo3, &vo4}
	newList := FilterMapFloat32Float64Ptr(notOneFloat32Float64Ptr, plusOneFloat32Float64Ptr, []*float32{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat32Float64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat32Float64Ptr failed")
	}

	if len(FilterMapFloat32Float64Ptr(nil, nil, []*float32{})) > 0 {
		t.Errorf("FilterMapFloat32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat32Float64Ptr(num *float32) bool {
	return *num != 1
}
func plusOneFloat32Float64Ptr(num *float32) *float64 {
	c := float64(*num + 1)
	return &c
}

func TestFilterMapFloat64IntPtr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 int = 3
	var vo4 int = 4

	expectedList := []*int{&vo3, &vo4}
	newList := FilterMapFloat64IntPtr(notOneFloat64IntPtr, plusOneFloat64IntPtr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64IntPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64IntPtr failed")
	}

	if len(FilterMapFloat64IntPtr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64IntPtr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64IntPtr(num *float64) *int {
	c := int(*num + 1)
	return &c
}

func TestFilterMapFloat64Int64Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 int64 = 3
	var vo4 int64 = 4

	expectedList := []*int64{&vo3, &vo4}
	newList := FilterMapFloat64Int64Ptr(notOneFloat64Int64Ptr, plusOneFloat64Int64Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int64Ptr failed")
	}

	if len(FilterMapFloat64Int64Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int64Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Int64Ptr(num *float64) *int64 {
	c := int64(*num + 1)
	return &c
}

func TestFilterMapFloat64Int32Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 int32 = 3
	var vo4 int32 = 4

	expectedList := []*int32{&vo3, &vo4}
	newList := FilterMapFloat64Int32Ptr(notOneFloat64Int32Ptr, plusOneFloat64Int32Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int32Ptr failed")
	}

	if len(FilterMapFloat64Int32Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int32Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Int32Ptr(num *float64) *int32 {
	c := int32(*num + 1)
	return &c
}

func TestFilterMapFloat64Int16Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 int16 = 3
	var vo4 int16 = 4

	expectedList := []*int16{&vo3, &vo4}
	newList := FilterMapFloat64Int16Ptr(notOneFloat64Int16Ptr, plusOneFloat64Int16Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int16Ptr failed")
	}

	if len(FilterMapFloat64Int16Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int16Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Int16Ptr(num *float64) *int16 {
	c := int16(*num + 1)
	return &c
}

func TestFilterMapFloat64Int8Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 int8 = 3
	var vo4 int8 = 4

	expectedList := []*int8{&vo3, &vo4}
	newList := FilterMapFloat64Int8Ptr(notOneFloat64Int8Ptr, plusOneFloat64Int8Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Int8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Int8Ptr failed")
	}

	if len(FilterMapFloat64Int8Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Int8Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Int8Ptr(num *float64) *int8 {
	c := int8(*num + 1)
	return &c
}

func TestFilterMapFloat64UintPtr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 uint = 3
	var vo4 uint = 4

	expectedList := []*uint{&vo3, &vo4}
	newList := FilterMapFloat64UintPtr(notOneFloat64UintPtr, plusOneFloat64UintPtr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64UintPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64UintPtr failed")
	}

	if len(FilterMapFloat64UintPtr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64UintPtr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64UintPtr(num *float64) *uint {
	c := uint(*num + 1)
	return &c
}

func TestFilterMapFloat64Uint64Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 uint64 = 3
	var vo4 uint64 = 4

	expectedList := []*uint64{&vo3, &vo4}
	newList := FilterMapFloat64Uint64Ptr(notOneFloat64Uint64Ptr, plusOneFloat64Uint64Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint64Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint64Ptr failed")
	}

	if len(FilterMapFloat64Uint64Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint64Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Uint64Ptr(num *float64) *uint64 {
	c := uint64(*num + 1)
	return &c
}

func TestFilterMapFloat64Uint32Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 uint32 = 3
	var vo4 uint32 = 4

	expectedList := []*uint32{&vo3, &vo4}
	newList := FilterMapFloat64Uint32Ptr(notOneFloat64Uint32Ptr, plusOneFloat64Uint32Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint32Ptr failed")
	}

	if len(FilterMapFloat64Uint32Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint32Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Uint32Ptr(num *float64) *uint32 {
	c := uint32(*num + 1)
	return &c
}

func TestFilterMapFloat64Uint16Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 uint16 = 3
	var vo4 uint16 = 4

	expectedList := []*uint16{&vo3, &vo4}
	newList := FilterMapFloat64Uint16Ptr(notOneFloat64Uint16Ptr, plusOneFloat64Uint16Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint16Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint16Ptr failed")
	}

	if len(FilterMapFloat64Uint16Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint16Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Uint16Ptr(num *float64) *uint16 {
	c := uint16(*num + 1)
	return &c
}

func TestFilterMapFloat64Uint8Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 uint8 = 3
	var vo4 uint8 = 4

	expectedList := []*uint8{&vo3, &vo4}
	newList := FilterMapFloat64Uint8Ptr(notOneFloat64Uint8Ptr, plusOneFloat64Uint8Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Uint8Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Uint8Ptr failed")
	}

	if len(FilterMapFloat64Uint8Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Uint8Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Uint8Ptr(num *float64) *uint8 {
	c := uint8(*num + 1)
	return &c
}

func TestFilterMapFloat64StrPtr(t *testing.T) {
	// Test : someLogic
	var ov10 string = "10"
	var iv1 float64 = 1
	var iv10 float64 = 10
	expectedList := []*string{&ov10}
	newList := FilterMapFloat64StrPtr(notOneFloat64StrNumPtr, someLogicFloat64StrNumPtr, []*float64{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMapFloat64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64StrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64StrPtr failed")
	}

	if len(FilterMapFloat64StrPtr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64StrNumPtr(num *float64) bool {
	return *num != 1
}

func someLogicFloat64StrNumPtr(num *float64) *string {
	var r string = "0"
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}

func TestFilterMapFloat64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vto bool = true
	var vfo bool = false

	var vi1 float64 = 1
	var vi10 float64 = 10
	var vi0 float64

	expectedList := []*bool{&vto, &vfo}
	newList := FilterMapFloat64BoolPtr(notOneFloat64BoolPtr, someLogicFloat64BoolPtr, []*float64{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64BoolPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64BoolPtr failed")
	}

	if len(FilterMapFloat64BoolPtr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64BoolPtr(num *float64) bool {
	return *num != 1
}

func someLogicFloat64BoolPtr(num *float64) *bool {
	r := *num > 0
	return &r
}

func TestFilterMapFloat64Float32Ptr(t *testing.T) {
	// Test : some logic
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var vo3 float32 = 3
	var vo4 float32 = 4

	expectedList := []*float32{&vo3, &vo4}
	newList := FilterMapFloat64Float32Ptr(notOneFloat64Float32Ptr, plusOneFloat64Float32Ptr, []*float64{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMapFloat64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMapFloat64Float32Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapFloat64Float32Ptr failed")
	}

	if len(FilterMapFloat64Float32Ptr(nil, nil, []*float64{})) > 0 {
		t.Errorf("FilterMapFloat64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOneFloat64Float32Ptr(num *float64) bool {
	return *num != 1
}
func plusOneFloat64Float32Ptr(num *float64) *float32 {
	c := float32(*num + 1)
	return &c
}
