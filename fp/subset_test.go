package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSubsetInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	actual := SubsetInt([]int{1}, []int{})
	if actual == true {
		t.Errorf("SubsetInt failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetInt([]int{v1, v2}, []int{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetIntRam failed")
	}

	list1 := []int{v8, v2, v1, v1, v2, v8}
	list2 := []int{v8, v2, v1, v1, v2, v8}
	actual = SubsetInt(list1, list2)
	if !actual {
		t.Errorf("SubsetInt failed.")
	}

	actual2 := SubsetIntPtr([]*int{&v1}, []*int{})
	if actual2 == true {
		t.Errorf("SubsetIntPtr failed")
	}

	actual2 = SubsetIntPtr([]*int{&v1, &v2}, []*int{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetIntPtr failed")
	}

	list1Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetIntPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetIntPtr failed.")
	}
}

func TestSubsetInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	actual := SubsetInt64([]int64{1}, []int64{})
	if actual == true {
		t.Errorf("SubsetInt64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetInt64([]int64{v1, v2}, []int64{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetInt64Ram failed")
	}

	list1 := []int64{v8, v2, v1, v1, v2, v8}
	list2 := []int64{v8, v2, v1, v1, v2, v8}
	actual = SubsetInt64(list1, list2)
	if !actual {
		t.Errorf("SubsetInt64 failed.")
	}

	actual2 := SubsetInt64Ptr([]*int64{&v1}, []*int64{})
	if actual2 == true {
		t.Errorf("SubsetInt64Ptr failed")
	}

	actual2 = SubsetInt64Ptr([]*int64{&v1, &v2}, []*int64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetInt64Ptr failed")
	}

	list1Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetInt64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetInt64Ptr failed.")
	}
}

func TestSubsetInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	actual := SubsetInt32([]int32{1}, []int32{})
	if actual == true {
		t.Errorf("SubsetInt32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetInt32([]int32{v1, v2}, []int32{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetInt32Ram failed")
	}

	list1 := []int32{v8, v2, v1, v1, v2, v8}
	list2 := []int32{v8, v2, v1, v1, v2, v8}
	actual = SubsetInt32(list1, list2)
	if !actual {
		t.Errorf("SubsetInt32 failed.")
	}

	actual2 := SubsetInt32Ptr([]*int32{&v1}, []*int32{})
	if actual2 == true {
		t.Errorf("SubsetInt32Ptr failed")
	}

	actual2 = SubsetInt32Ptr([]*int32{&v1, &v2}, []*int32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetInt32Ptr failed")
	}

	list1Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetInt32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetInt32Ptr failed.")
	}
}

func TestSubsetInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	actual := SubsetInt16([]int16{1}, []int16{})
	if actual == true {
		t.Errorf("SubsetInt16 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetInt16([]int16{v1, v2}, []int16{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetInt16Ram failed")
	}

	list1 := []int16{v8, v2, v1, v1, v2, v8}
	list2 := []int16{v8, v2, v1, v1, v2, v8}
	actual = SubsetInt16(list1, list2)
	if !actual {
		t.Errorf("SubsetInt16 failed.")
	}

	actual2 := SubsetInt16Ptr([]*int16{&v1}, []*int16{})
	if actual2 == true {
		t.Errorf("SubsetInt16Ptr failed")
	}

	actual2 = SubsetInt16Ptr([]*int16{&v1, &v2}, []*int16{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetInt16Ptr failed")
	}

	list1Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetInt16Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetInt16Ptr failed.")
	}
}

func TestSubsetInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	actual := SubsetInt8([]int8{1}, []int8{})
	if actual == true {
		t.Errorf("SubsetInt8 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetInt8([]int8{v1, v2}, []int8{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetInt8Ram failed")
	}

	list1 := []int8{v8, v2, v1, v1, v2, v8}
	list2 := []int8{v8, v2, v1, v1, v2, v8}
	actual = SubsetInt8(list1, list2)
	if !actual {
		t.Errorf("SubsetInt8 failed.")
	}

	actual2 := SubsetInt8Ptr([]*int8{&v1}, []*int8{})
	if actual2 == true {
		t.Errorf("SubsetInt8Ptr failed")
	}

	actual2 = SubsetInt8Ptr([]*int8{&v1, &v2}, []*int8{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetInt8Ptr failed")
	}

	list1Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetInt8Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetInt8Ptr failed.")
	}
}

func TestSubsetUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	actual := SubsetUint([]uint{1}, []uint{})
	if actual == true {
		t.Errorf("SubsetUint failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetUint([]uint{v1, v2}, []uint{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetUintRam failed")
	}

	list1 := []uint{v8, v2, v1, v1, v2, v8}
	list2 := []uint{v8, v2, v1, v1, v2, v8}
	actual = SubsetUint(list1, list2)
	if !actual {
		t.Errorf("SubsetUint failed.")
	}

	actual2 := SubsetUintPtr([]*uint{&v1}, []*uint{})
	if actual2 == true {
		t.Errorf("SubsetUintPtr failed")
	}

	actual2 = SubsetUintPtr([]*uint{&v1, &v2}, []*uint{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetUintPtr failed")
	}

	list1Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetUintPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetUintPtr failed.")
	}
}

func TestSubsetUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	actual := SubsetUint64([]uint64{1}, []uint64{})
	if actual == true {
		t.Errorf("SubsetUint64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetUint64([]uint64{v1, v2}, []uint64{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetUint64Ram failed")
	}

	list1 := []uint64{v8, v2, v1, v1, v2, v8}
	list2 := []uint64{v8, v2, v1, v1, v2, v8}
	actual = SubsetUint64(list1, list2)
	if !actual {
		t.Errorf("SubsetUint64 failed.")
	}

	actual2 := SubsetUint64Ptr([]*uint64{&v1}, []*uint64{})
	if actual2 == true {
		t.Errorf("SubsetUint64Ptr failed")
	}

	actual2 = SubsetUint64Ptr([]*uint64{&v1, &v2}, []*uint64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetUint64Ptr failed")
	}

	list1Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetUint64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetUint64Ptr failed.")
	}
}

func TestSubsetUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	actual := SubsetUint32([]uint32{1}, []uint32{})
	if actual == true {
		t.Errorf("SubsetUint32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetUint32([]uint32{v1, v2}, []uint32{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetUint32Ram failed")
	}

	list1 := []uint32{v8, v2, v1, v1, v2, v8}
	list2 := []uint32{v8, v2, v1, v1, v2, v8}
	actual = SubsetUint32(list1, list2)
	if !actual {
		t.Errorf("SubsetUint32 failed.")
	}

	actual2 := SubsetUint32Ptr([]*uint32{&v1}, []*uint32{})
	if actual2 == true {
		t.Errorf("SubsetUint32Ptr failed")
	}

	actual2 = SubsetUint32Ptr([]*uint32{&v1, &v2}, []*uint32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetUint32Ptr failed")
	}

	list1Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetUint32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetUint32Ptr failed.")
	}
}

func TestSubsetUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	actual := SubsetUint16([]uint16{1}, []uint16{})
	if actual == true {
		t.Errorf("SubsetUint16 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetUint16([]uint16{v1, v2}, []uint16{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetUint16Ram failed")
	}

	list1 := []uint16{v8, v2, v1, v1, v2, v8}
	list2 := []uint16{v8, v2, v1, v1, v2, v8}
	actual = SubsetUint16(list1, list2)
	if !actual {
		t.Errorf("SubsetUint16 failed.")
	}

	actual2 := SubsetUint16Ptr([]*uint16{&v1}, []*uint16{})
	if actual2 == true {
		t.Errorf("SubsetUint16Ptr failed")
	}

	actual2 = SubsetUint16Ptr([]*uint16{&v1, &v2}, []*uint16{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetUint16Ptr failed")
	}

	list1Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetUint16Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetUint16Ptr failed.")
	}
}

func TestSubsetUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	actual := SubsetUint8([]uint8{1}, []uint8{})
	if actual == true {
		t.Errorf("SubsetUint8 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetUint8([]uint8{v1, v2}, []uint8{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetUint8Ram failed")
	}

	list1 := []uint8{v8, v2, v1, v1, v2, v8}
	list2 := []uint8{v8, v2, v1, v1, v2, v8}
	actual = SubsetUint8(list1, list2)
	if !actual {
		t.Errorf("SubsetUint8 failed.")
	}

	actual2 := SubsetUint8Ptr([]*uint8{&v1}, []*uint8{})
	if actual2 == true {
		t.Errorf("SubsetUint8Ptr failed")
	}

	actual2 = SubsetUint8Ptr([]*uint8{&v1, &v2}, []*uint8{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetUint8Ptr failed")
	}

	list1Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetUint8Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetUint8Ptr failed.")
	}
}

func TestSubsetStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	actual := SubsetStr([]string{"1"}, []string{})
	if actual == true {
		t.Errorf("SubsetStr failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetStr([]string{v1, v2}, []string{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetStrRam failed")
	}

	list1 := []string{v8, v2, v1, v1, v2, v8}
	list2 := []string{v8, v2, v1, v1, v2, v8}
	actual = SubsetStr(list1, list2)
	if !actual {
		t.Errorf("SubsetStr failed.")
	}

	actual2 := SubsetStrPtr([]*string{&v1}, []*string{})
	if actual2 == true {
		t.Errorf("SubsetStrPtr failed")
	}

	actual2 = SubsetStrPtr([]*string{&v1, &v2}, []*string{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetStrPtr failed")
	}

	list1Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetStrPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetStrPtr failed.")
	}
}

func TestSubsetBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = false
	var v1 bool = true

	actual := SubsetBool([]bool{true}, []bool{})
	if actual == true {
		t.Errorf("SubsetBool failed")
		t.Errorf(reflect.String.String())
	}

	list1 := []bool{v8, v2, v1, v1, v2, v8}
	list2 := []bool{v8, v2, v1, v1, v2, v8}
	actual = SubsetBool(list1, list2)
	if !actual {
		t.Errorf("SubsetBool failed.")
	}

	actual2 := SubsetBoolPtr([]*bool{&v1}, []*bool{})
	if actual2 == true {
		t.Errorf("SubsetBoolPtr failed")
	}

	list1Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetBoolPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetBoolPtr failed.")
	}
}

func TestSubsetFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	actual := SubsetFloat32([]float32{1}, []float32{})
	if actual == true {
		t.Errorf("SubsetFloat32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetFloat32([]float32{v1, v2}, []float32{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetFloat32Ram failed")
	}

	list1 := []float32{v8, v2, v1, v1, v2, v8}
	list2 := []float32{v8, v2, v1, v1, v2, v8}
	actual = SubsetFloat32(list1, list2)
	if !actual {
		t.Errorf("SubsetFloat32 failed.")
	}

	actual2 := SubsetFloat32Ptr([]*float32{&v1}, []*float32{})
	if actual2 == true {
		t.Errorf("SubsetFloat32Ptr failed")
	}

	actual2 = SubsetFloat32Ptr([]*float32{&v1, &v2}, []*float32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetFloat32Ptr failed")
	}

	list1Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetFloat32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetFloat32Ptr failed.")
	}
}

func TestSubsetFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	actual := SubsetFloat64([]float64{1}, []float64{})
	if actual == true {
		t.Errorf("SubsetFloat64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SubsetFloat64([]float64{v1, v2}, []float64{v1, v8, v8})
	if actual == true {
		t.Errorf("SubsetFloat64Ram failed")
	}

	list1 := []float64{v8, v2, v1, v1, v2, v8}
	list2 := []float64{v8, v2, v1, v1, v2, v8}
	actual = SubsetFloat64(list1, list2)
	if !actual {
		t.Errorf("SubsetFloat64 failed.")
	}

	actual2 := SubsetFloat64Ptr([]*float64{&v1}, []*float64{})
	if actual2 == true {
		t.Errorf("SubsetFloat64Ptr failed")
	}

	actual2 = SubsetFloat64Ptr([]*float64{&v1, &v2}, []*float64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SubsetFloat64Ptr failed")
	}

	list1Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SubsetFloat64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SubsetFloat64Ptr failed.")
	}
}
