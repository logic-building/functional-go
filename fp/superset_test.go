package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSupersetInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v1 int = 1

	actual := SupersetInt([]int{1}, []int{})
	if actual == true {
		t.Errorf("SupersetInt failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetInt([]int{v1, v2}, []int{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetIntRam failed")
	}

	list1 := []int{v8, v2, v1, v1, v2, v8}
	list2 := []int{v8, v2, v1, v1, v2, v8}
	actual = SupersetInt(list1, list2)
	if !actual {
		t.Errorf("SupersetInt failed.")
	}

	actual2 := SupersetIntPtr([]*int{&v1}, []*int{})
	if actual2 == true {
		t.Errorf("SupersetIntPtr failed")
	}

	actual2 = SupersetIntPtr([]*int{&v1, &v2}, []*int{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetIntPtr failed")
	}

	list1Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetIntPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetIntPtr failed.")
	}
}

func TestSupersetInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v1 int64 = 1

	actual := SupersetInt64([]int64{1}, []int64{})
	if actual == true {
		t.Errorf("SupersetInt64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetInt64([]int64{v1, v2}, []int64{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetInt64Ram failed")
	}

	list1 := []int64{v8, v2, v1, v1, v2, v8}
	list2 := []int64{v8, v2, v1, v1, v2, v8}
	actual = SupersetInt64(list1, list2)
	if !actual {
		t.Errorf("SupersetInt64 failed.")
	}

	actual2 := SupersetInt64Ptr([]*int64{&v1}, []*int64{})
	if actual2 == true {
		t.Errorf("SupersetInt64Ptr failed")
	}

	actual2 = SupersetInt64Ptr([]*int64{&v1, &v2}, []*int64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetInt64Ptr failed")
	}

	list1Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetInt64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetInt64Ptr failed.")
	}
}

func TestSupersetInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v1 int32 = 1

	actual := SupersetInt32([]int32{1}, []int32{})
	if actual == true {
		t.Errorf("SupersetInt32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetInt32([]int32{v1, v2}, []int32{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetInt32Ram failed")
	}

	list1 := []int32{v8, v2, v1, v1, v2, v8}
	list2 := []int32{v8, v2, v1, v1, v2, v8}
	actual = SupersetInt32(list1, list2)
	if !actual {
		t.Errorf("SupersetInt32 failed.")
	}

	actual2 := SupersetInt32Ptr([]*int32{&v1}, []*int32{})
	if actual2 == true {
		t.Errorf("SupersetInt32Ptr failed")
	}

	actual2 = SupersetInt32Ptr([]*int32{&v1, &v2}, []*int32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetInt32Ptr failed")
	}

	list1Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetInt32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetInt32Ptr failed.")
	}
}

func TestSupersetInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v1 int16 = 1

	actual := SupersetInt16([]int16{1}, []int16{})
	if actual == true {
		t.Errorf("SupersetInt16 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetInt16([]int16{v1, v2}, []int16{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetInt16Ram failed")
	}

	list1 := []int16{v8, v2, v1, v1, v2, v8}
	list2 := []int16{v8, v2, v1, v1, v2, v8}
	actual = SupersetInt16(list1, list2)
	if !actual {
		t.Errorf("SupersetInt16 failed.")
	}

	actual2 := SupersetInt16Ptr([]*int16{&v1}, []*int16{})
	if actual2 == true {
		t.Errorf("SupersetInt16Ptr failed")
	}

	actual2 = SupersetInt16Ptr([]*int16{&v1, &v2}, []*int16{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetInt16Ptr failed")
	}

	list1Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int16{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetInt16Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetInt16Ptr failed.")
	}
}

func TestSupersetInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v1 int8 = 1

	actual := SupersetInt8([]int8{1}, []int8{})
	if actual == true {
		t.Errorf("SupersetInt8 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetInt8([]int8{v1, v2}, []int8{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetInt8Ram failed")
	}

	list1 := []int8{v8, v2, v1, v1, v2, v8}
	list2 := []int8{v8, v2, v1, v1, v2, v8}
	actual = SupersetInt8(list1, list2)
	if !actual {
		t.Errorf("SupersetInt8 failed.")
	}

	actual2 := SupersetInt8Ptr([]*int8{&v1}, []*int8{})
	if actual2 == true {
		t.Errorf("SupersetInt8Ptr failed")
	}

	actual2 = SupersetInt8Ptr([]*int8{&v1, &v2}, []*int8{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetInt8Ptr failed")
	}

	list1Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*int8{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetInt8Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetInt8Ptr failed.")
	}
}

func TestSupersetUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v1 uint = 1

	actual := SupersetUint([]uint{1}, []uint{})
	if actual == true {
		t.Errorf("SupersetUint failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetUint([]uint{v1, v2}, []uint{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetUintRam failed")
	}

	list1 := []uint{v8, v2, v1, v1, v2, v8}
	list2 := []uint{v8, v2, v1, v1, v2, v8}
	actual = SupersetUint(list1, list2)
	if !actual {
		t.Errorf("SupersetUint failed.")
	}

	actual2 := SupersetUintPtr([]*uint{&v1}, []*uint{})
	if actual2 == true {
		t.Errorf("SupersetUintPtr failed")
	}

	actual2 = SupersetUintPtr([]*uint{&v1, &v2}, []*uint{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetUintPtr failed")
	}

	list1Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetUintPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetUintPtr failed.")
	}
}

func TestSupersetUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v1 uint64 = 1

	actual := SupersetUint64([]uint64{1}, []uint64{})
	if actual == true {
		t.Errorf("SupersetUint64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetUint64([]uint64{v1, v2}, []uint64{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetUint64Ram failed")
	}

	list1 := []uint64{v8, v2, v1, v1, v2, v8}
	list2 := []uint64{v8, v2, v1, v1, v2, v8}
	actual = SupersetUint64(list1, list2)
	if !actual {
		t.Errorf("SupersetUint64 failed.")
	}

	actual2 := SupersetUint64Ptr([]*uint64{&v1}, []*uint64{})
	if actual2 == true {
		t.Errorf("SupersetUint64Ptr failed")
	}

	actual2 = SupersetUint64Ptr([]*uint64{&v1, &v2}, []*uint64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetUint64Ptr failed")
	}

	list1Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetUint64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetUint64Ptr failed.")
	}
}

func TestSupersetUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v1 uint32 = 1

	actual := SupersetUint32([]uint32{1}, []uint32{})
	if actual == true {
		t.Errorf("SupersetUint32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetUint32([]uint32{v1, v2}, []uint32{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetUint32Ram failed")
	}

	list1 := []uint32{v8, v2, v1, v1, v2, v8}
	list2 := []uint32{v8, v2, v1, v1, v2, v8}
	actual = SupersetUint32(list1, list2)
	if !actual {
		t.Errorf("SupersetUint32 failed.")
	}

	actual2 := SupersetUint32Ptr([]*uint32{&v1}, []*uint32{})
	if actual2 == true {
		t.Errorf("SupersetUint32Ptr failed")
	}

	actual2 = SupersetUint32Ptr([]*uint32{&v1, &v2}, []*uint32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetUint32Ptr failed")
	}

	list1Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetUint32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetUint32Ptr failed.")
	}
}

func TestSupersetUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v1 uint16 = 1

	actual := SupersetUint16([]uint16{1}, []uint16{})
	if actual == true {
		t.Errorf("SupersetUint16 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetUint16([]uint16{v1, v2}, []uint16{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetUint16Ram failed")
	}

	list1 := []uint16{v8, v2, v1, v1, v2, v8}
	list2 := []uint16{v8, v2, v1, v1, v2, v8}
	actual = SupersetUint16(list1, list2)
	if !actual {
		t.Errorf("SupersetUint16 failed.")
	}

	actual2 := SupersetUint16Ptr([]*uint16{&v1}, []*uint16{})
	if actual2 == true {
		t.Errorf("SupersetUint16Ptr failed")
	}

	actual2 = SupersetUint16Ptr([]*uint16{&v1, &v2}, []*uint16{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetUint16Ptr failed")
	}

	list1Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint16{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetUint16Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetUint16Ptr failed.")
	}
}

func TestSupersetUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v1 uint8 = 1

	actual := SupersetUint8([]uint8{1}, []uint8{})
	if actual == true {
		t.Errorf("SupersetUint8 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetUint8([]uint8{v1, v2}, []uint8{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetUint8Ram failed")
	}

	list1 := []uint8{v8, v2, v1, v1, v2, v8}
	list2 := []uint8{v8, v2, v1, v1, v2, v8}
	actual = SupersetUint8(list1, list2)
	if !actual {
		t.Errorf("SupersetUint8 failed.")
	}

	actual2 := SupersetUint8Ptr([]*uint8{&v1}, []*uint8{})
	if actual2 == true {
		t.Errorf("SupersetUint8Ptr failed")
	}

	actual2 = SupersetUint8Ptr([]*uint8{&v1, &v2}, []*uint8{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetUint8Ptr failed")
	}

	list1Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*uint8{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetUint8Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetUint8Ptr failed.")
	}
}

func TestSupersetStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v1 string = "1"

	actual := SupersetStr([]string{"1"}, []string{})
	if actual == true {
		t.Errorf("SupersetStr failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetStr([]string{v1, v2}, []string{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetStrRam failed")
	}

	list1 := []string{v8, v2, v1, v1, v2, v8}
	list2 := []string{v8, v2, v1, v1, v2, v8}
	actual = SupersetStr(list1, list2)
	if !actual {
		t.Errorf("SupersetStr failed.")
	}

	actual2 := SupersetStrPtr([]*string{&v1}, []*string{})
	if actual2 == true {
		t.Errorf("SupersetStrPtr failed")
	}

	actual2 = SupersetStrPtr([]*string{&v1, &v2}, []*string{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetStrPtr failed")
	}

	list1Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*string{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetStrPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetStrPtr failed.")
	}
}

func TestSupersetBool(t *testing.T) {
	var v8 bool = true
	var v2 bool = false
	var v1 bool = true

	actual := SupersetBool([]bool{true}, []bool{})
	if actual == true {
		t.Errorf("SupersetBool failed")
	}

	list1 := []bool{v8, v2, v1, v1, v2, v8}
	list2 := []bool{v8, v2, v1, v1, v2, v8}
	actual = SupersetBool(list1, list2)
	if !actual {
		t.Errorf("SupersetBool failed.")
	}

	actual2 := SupersetBoolPtr([]*bool{&v1}, []*bool{})
	if actual2 == true {
		t.Errorf("SupersetBoolPtr failed")
	}

	list1Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*bool{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetBoolPtr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetBoolPtr failed.")
	}
}

func TestSupersetFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v1 float32 = 1

	actual := SupersetFloat32([]float32{1}, []float32{})
	if actual == true {
		t.Errorf("SupersetFloat32 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetFloat32([]float32{v1, v2}, []float32{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetFloat32Ram failed")
	}

	list1 := []float32{v8, v2, v1, v1, v2, v8}
	list2 := []float32{v8, v2, v1, v1, v2, v8}
	actual = SupersetFloat32(list1, list2)
	if !actual {
		t.Errorf("SupersetFloat32 failed.")
	}

	actual2 := SupersetFloat32Ptr([]*float32{&v1}, []*float32{})
	if actual2 == true {
		t.Errorf("SupersetFloat32Ptr failed")
	}

	actual2 = SupersetFloat32Ptr([]*float32{&v1, &v2}, []*float32{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetFloat32Ptr failed")
	}

	list1Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float32{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetFloat32Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetFloat32Ptr failed.")
	}
}

func TestSupersetFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v1 float64 = 1

	actual := SupersetFloat64([]float64{1}, []float64{})
	if actual == true {
		t.Errorf("SupersetFloat64 failed")
		t.Errorf(reflect.String.String())
	}

	actual = SupersetFloat64([]float64{v1, v2}, []float64{v1, v8, v8})
	if actual == true {
		t.Errorf("SupersetFloat64Ram failed")
	}

	list1 := []float64{v8, v2, v1, v1, v2, v8}
	list2 := []float64{v8, v2, v1, v1, v2, v8}
	actual = SupersetFloat64(list1, list2)
	if !actual {
		t.Errorf("SupersetFloat64 failed.")
	}

	actual2 := SupersetFloat64Ptr([]*float64{&v1}, []*float64{})
	if actual2 == true {
		t.Errorf("SupersetFloat64Ptr failed")
	}

	actual2 = SupersetFloat64Ptr([]*float64{&v1, &v2}, []*float64{&v8, &v1})
	if actual2 == true {
		t.Errorf("SupersetFloat64Ptr failed")
	}

	list1Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*float64{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = SupersetFloat64Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("SupersetFloat64Ptr failed.")
	}
}
