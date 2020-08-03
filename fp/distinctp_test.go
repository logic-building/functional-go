package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDistinctIntP(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v0 int

	list0 := []int{v8, v2, v8, v0, v2, v0}
	result := DistinctIntP(list0)
	if result {
		t.Errorf("DistinctIntP failed. Expected=false, actual=true")
	}

	list0 = []int{v8, v2, v0}
	result = DistinctIntP(list0)
	if !result {
		t.Errorf("DistinctIntP failed. Expected=true, actual=false")
	}

	list0 = []int{}
	result = DistinctIntP(list0)
	if result {
		t.Errorf("DistinctIntP failed. Expected=false, actual=true")
	}

	result = DistinctIntP(nil)
	if result {
		t.Errorf("DistinctintP failed. Expected=false, actual=true")
	}

	list := []*int{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctIntPPtr(list)
	if distinct {
		t.Errorf("DistinctIntPPtr failed. Expected=false, actual=true")
	}

	list = []*int{&v8, &v2, &v0}
	distinct = DistinctIntPPtr(list)
	if !distinct {
		t.Errorf("DistinctIntPPtr failed. Expected=true, actual=false")
	}

	list = []*int{}
	distinct = DistinctIntPPtr(list)
	if distinct {
		t.Errorf("DistinctIntPPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctIntPPtr(nil)
	if distinct {
		t.Errorf("DistinctintPPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt64P(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v0 int64

	list0 := []int64{v8, v2, v8, v0, v2, v0}
	result := DistinctInt64P(list0)
	if result {
		t.Errorf("DistinctInt64P failed. Expected=false, actual=true")
	}

	list0 = []int64{v8, v2, v0}
	result = DistinctInt64P(list0)
	if !result {
		t.Errorf("DistinctInt64P failed. Expected=true, actual=false")
	}

	list0 = []int64{}
	result = DistinctInt64P(list0)
	if result {
		t.Errorf("DistinctInt64P failed. Expected=false, actual=true")
	}

	result = DistinctInt64P(nil)
	if result {
		t.Errorf("Distinctint64P failed. Expected=false, actual=true")
	}

	list := []*int64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt64PPtr(list)
	if distinct {
		t.Errorf("DistinctInt64PPtr failed. Expected=false, actual=true")
	}

	list = []*int64{&v8, &v2, &v0}
	distinct = DistinctInt64PPtr(list)
	if !distinct {
		t.Errorf("DistinctInt64PPtr failed. Expected=true, actual=false")
	}

	list = []*int64{}
	distinct = DistinctInt64PPtr(list)
	if distinct {
		t.Errorf("DistinctInt64PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctInt64PPtr(nil)
	if distinct {
		t.Errorf("Distinctint64PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt32P(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v0 int32

	list0 := []int32{v8, v2, v8, v0, v2, v0}
	result := DistinctInt32P(list0)
	if result {
		t.Errorf("DistinctInt32P failed. Expected=false, actual=true")
	}

	list0 = []int32{v8, v2, v0}
	result = DistinctInt32P(list0)
	if !result {
		t.Errorf("DistinctInt32P failed. Expected=true, actual=false")
	}

	list0 = []int32{}
	result = DistinctInt32P(list0)
	if result {
		t.Errorf("DistinctInt32P failed. Expected=false, actual=true")
	}

	result = DistinctInt32P(nil)
	if result {
		t.Errorf("Distinctint32P failed. Expected=false, actual=true")
	}

	list := []*int32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt32PPtr(list)
	if distinct {
		t.Errorf("DistinctInt32PPtr failed. Expected=false, actual=true")
	}

	list = []*int32{&v8, &v2, &v0}
	distinct = DistinctInt32PPtr(list)
	if !distinct {
		t.Errorf("DistinctInt32PPtr failed. Expected=true, actual=false")
	}

	list = []*int32{}
	distinct = DistinctInt32PPtr(list)
	if distinct {
		t.Errorf("DistinctInt32PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctInt32PPtr(nil)
	if distinct {
		t.Errorf("Distinctint32PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt16P(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v0 int16

	list0 := []int16{v8, v2, v8, v0, v2, v0}
	result := DistinctInt16P(list0)
	if result {
		t.Errorf("DistinctInt16P failed. Expected=false, actual=true")
	}

	list0 = []int16{v8, v2, v0}
	result = DistinctInt16P(list0)
	if !result {
		t.Errorf("DistinctInt16P failed. Expected=true, actual=false")
	}

	list0 = []int16{}
	result = DistinctInt16P(list0)
	if result {
		t.Errorf("DistinctInt16P failed. Expected=false, actual=true")
	}

	result = DistinctInt16P(nil)
	if result {
		t.Errorf("Distinctint16P failed. Expected=false, actual=true")
	}

	list := []*int16{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt16PPtr(list)
	if distinct {
		t.Errorf("DistinctInt16PPtr failed. Expected=false, actual=true")
	}

	list = []*int16{&v8, &v2, &v0}
	distinct = DistinctInt16PPtr(list)
	if !distinct {
		t.Errorf("DistinctInt16PPtr failed. Expected=true, actual=false")
	}

	list = []*int16{}
	distinct = DistinctInt16PPtr(list)
	if distinct {
		t.Errorf("DistinctInt16PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctInt16PPtr(nil)
	if distinct {
		t.Errorf("Distinctint16PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt8P(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v0 int8

	list0 := []int8{v8, v2, v8, v0, v2, v0}
	result := DistinctInt8P(list0)
	if result {
		t.Errorf("DistinctInt8P failed. Expected=false, actual=true")
	}

	list0 = []int8{v8, v2, v0}
	result = DistinctInt8P(list0)
	if !result {
		t.Errorf("DistinctInt8P failed. Expected=true, actual=false")
	}

	list0 = []int8{}
	result = DistinctInt8P(list0)
	if result {
		t.Errorf("DistinctInt8P failed. Expected=false, actual=true")
	}

	result = DistinctInt8P(nil)
	if result {
		t.Errorf("Distinctint8P failed. Expected=false, actual=true")
	}

	list := []*int8{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt8PPtr(list)
	if distinct {
		t.Errorf("DistinctInt8PPtr failed. Expected=false, actual=true")
	}

	list = []*int8{&v8, &v2, &v0}
	distinct = DistinctInt8PPtr(list)
	if !distinct {
		t.Errorf("DistinctInt8PPtr failed. Expected=true, actual=false")
	}

	list = []*int8{}
	distinct = DistinctInt8PPtr(list)
	if distinct {
		t.Errorf("DistinctInt8PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctInt8PPtr(nil)
	if distinct {
		t.Errorf("Distinctint8PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUintP(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v0 uint

	list0 := []uint{v8, v2, v8, v0, v2, v0}
	result := DistinctUintP(list0)
	if result {
		t.Errorf("DistinctUintP failed. Expected=false, actual=true")
	}

	list0 = []uint{v8, v2, v0}
	result = DistinctUintP(list0)
	if !result {
		t.Errorf("DistinctUintP failed. Expected=true, actual=false")
	}

	list0 = []uint{}
	result = DistinctUintP(list0)
	if result {
		t.Errorf("DistinctUintP failed. Expected=false, actual=true")
	}

	result = DistinctUintP(nil)
	if result {
		t.Errorf("DistinctuintP failed. Expected=false, actual=true")
	}

	list := []*uint{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUintPPtr(list)
	if distinct {
		t.Errorf("DistinctUintPPtr failed. Expected=false, actual=true")
	}

	list = []*uint{&v8, &v2, &v0}
	distinct = DistinctUintPPtr(list)
	if !distinct {
		t.Errorf("DistinctUintPPtr failed. Expected=true, actual=false")
	}

	list = []*uint{}
	distinct = DistinctUintPPtr(list)
	if distinct {
		t.Errorf("DistinctUintPPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctUintPPtr(nil)
	if distinct {
		t.Errorf("DistinctuintPPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint64P(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v0 uint64

	list0 := []uint64{v8, v2, v8, v0, v2, v0}
	result := DistinctUint64P(list0)
	if result {
		t.Errorf("DistinctUint64P failed. Expected=false, actual=true")
	}

	list0 = []uint64{v8, v2, v0}
	result = DistinctUint64P(list0)
	if !result {
		t.Errorf("DistinctUint64P failed. Expected=true, actual=false")
	}

	list0 = []uint64{}
	result = DistinctUint64P(list0)
	if result {
		t.Errorf("DistinctUint64P failed. Expected=false, actual=true")
	}

	result = DistinctUint64P(nil)
	if result {
		t.Errorf("Distinctuint64P failed. Expected=false, actual=true")
	}

	list := []*uint64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint64PPtr(list)
	if distinct {
		t.Errorf("DistinctUint64PPtr failed. Expected=false, actual=true")
	}

	list = []*uint64{&v8, &v2, &v0}
	distinct = DistinctUint64PPtr(list)
	if !distinct {
		t.Errorf("DistinctUint64PPtr failed. Expected=true, actual=false")
	}

	list = []*uint64{}
	distinct = DistinctUint64PPtr(list)
	if distinct {
		t.Errorf("DistinctUint64PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctUint64PPtr(nil)
	if distinct {
		t.Errorf("Distinctuint64PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint32P(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v0 uint32

	list0 := []uint32{v8, v2, v8, v0, v2, v0}
	result := DistinctUint32P(list0)
	if result {
		t.Errorf("DistinctUint32P failed. Expected=false, actual=true")
	}

	list0 = []uint32{v8, v2, v0}
	result = DistinctUint32P(list0)
	if !result {
		t.Errorf("DistinctUint32P failed. Expected=true, actual=false")
	}

	list0 = []uint32{}
	result = DistinctUint32P(list0)
	if result {
		t.Errorf("DistinctUint32P failed. Expected=false, actual=true")
	}

	result = DistinctUint32P(nil)
	if result {
		t.Errorf("Distinctuint32P failed. Expected=false, actual=true")
	}

	list := []*uint32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint32PPtr(list)
	if distinct {
		t.Errorf("DistinctUint32PPtr failed. Expected=false, actual=true")
	}

	list = []*uint32{&v8, &v2, &v0}
	distinct = DistinctUint32PPtr(list)
	if !distinct {
		t.Errorf("DistinctUint32PPtr failed. Expected=true, actual=false")
	}

	list = []*uint32{}
	distinct = DistinctUint32PPtr(list)
	if distinct {
		t.Errorf("DistinctUint32PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctUint32PPtr(nil)
	if distinct {
		t.Errorf("Distinctuint32PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint16P(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v0 uint16

	list0 := []uint16{v8, v2, v8, v0, v2, v0}
	result := DistinctUint16P(list0)
	if result {
		t.Errorf("DistinctUint16P failed. Expected=false, actual=true")
	}

	list0 = []uint16{v8, v2, v0}
	result = DistinctUint16P(list0)
	if !result {
		t.Errorf("DistinctUint16P failed. Expected=true, actual=false")
	}

	list0 = []uint16{}
	result = DistinctUint16P(list0)
	if result {
		t.Errorf("DistinctUint16P failed. Expected=false, actual=true")
	}

	result = DistinctUint16P(nil)
	if result {
		t.Errorf("Distinctuint16P failed. Expected=false, actual=true")
	}

	list := []*uint16{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint16PPtr(list)
	if distinct {
		t.Errorf("DistinctUint16PPtr failed. Expected=false, actual=true")
	}

	list = []*uint16{&v8, &v2, &v0}
	distinct = DistinctUint16PPtr(list)
	if !distinct {
		t.Errorf("DistinctUint16PPtr failed. Expected=true, actual=false")
	}

	list = []*uint16{}
	distinct = DistinctUint16PPtr(list)
	if distinct {
		t.Errorf("DistinctUint16PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctUint16PPtr(nil)
	if distinct {
		t.Errorf("Distinctuint16PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint8P(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v0 uint8

	list0 := []uint8{v8, v2, v8, v0, v2, v0}
	result := DistinctUint8P(list0)
	if result {
		t.Errorf("DistinctUint8P failed. Expected=false, actual=true")
	}

	list0 = []uint8{v8, v2, v0}
	result = DistinctUint8P(list0)
	if !result {
		t.Errorf("DistinctUint8P failed. Expected=true, actual=false")
	}

	list0 = []uint8{}
	result = DistinctUint8P(list0)
	if result {
		t.Errorf("DistinctUint8P failed. Expected=false, actual=true")
	}

	result = DistinctUint8P(nil)
	if result {
		t.Errorf("Distinctuint8P failed. Expected=false, actual=true")
	}

	list := []*uint8{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint8PPtr(list)
	if distinct {
		t.Errorf("DistinctUint8PPtr failed. Expected=false, actual=true")
	}

	list = []*uint8{&v8, &v2, &v0}
	distinct = DistinctUint8PPtr(list)
	if !distinct {
		t.Errorf("DistinctUint8PPtr failed. Expected=true, actual=false")
	}

	list = []*uint8{}
	distinct = DistinctUint8PPtr(list)
	if distinct {
		t.Errorf("DistinctUint8PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctUint8PPtr(nil)
	if distinct {
		t.Errorf("Distinctuint8PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctStrP(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v0 string

	list0 := []string{v8, v2, v8, v0, v2, v0}
	result := DistinctStrP(list0)
	if result {
		t.Errorf("DistinctStrP failed. Expected=false, actual=true")
	}

	list0 = []string{v8, v2, v0}
	result = DistinctStrP(list0)
	if !result {
		t.Errorf("DistinctStrP failed. Expected=true, actual=false")
	}

	list0 = []string{}
	result = DistinctStrP(list0)
	if result {
		t.Errorf("DistinctStrP failed. Expected=false, actual=true")
	}

	result = DistinctStrP(nil)
	if result {
		t.Errorf("DistinctstringP failed. Expected=false, actual=true")
	}

	list := []*string{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctStrPPtr(list)
	if distinct {
		t.Errorf("DistinctStrPPtr failed. Expected=false, actual=true")
	}

	list = []*string{&v8, &v2, &v0}
	distinct = DistinctStrPPtr(list)
	if !distinct {
		t.Errorf("DistinctStrPPtr failed. Expected=true, actual=false")
	}

	list = []*string{}
	distinct = DistinctStrPPtr(list)
	if distinct {
		t.Errorf("DistinctStrPPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctStrPPtr(nil)
	if distinct {
		t.Errorf("DistinctstringPPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctBoolP(t *testing.T) {
	var v8 bool = true
	var v2 bool = false

	list0 := []bool{v8, v2, v8, v2}
	result := DistinctBoolP(list0)
	if result {
		t.Errorf("DistinctBoolP failed. Expected=false, actual=true")
	}

	list0 = []bool{v8, v2}
	result = DistinctBoolP(list0)
	if !result {
		t.Errorf("DistinctBoolP failed. Expected=true, actual=false")
	}

	list0 = []bool{}
	result = DistinctBoolP(list0)
	if result {
		t.Errorf("DistinctBoolP failed. Expected=false, actual=true")
	}

	result = DistinctBoolP(nil)
	if result {
		t.Errorf("DistinctboolP failed. Expected=false, actual=true")
	}

	list := []*bool{&v8, &v2, &v8, &v2}
	distinct := DistinctBoolPPtr(list)
	if distinct {
		t.Errorf("DistinctBoolPPtr failed. Expected=false, actual=true")
	}

	list = []*bool{&v8, &v2}
	distinct = DistinctBoolPPtr(list)
	if !distinct {
		t.Errorf("DistinctBoolPPtr failed. Expected=true, actual=false")
	}

	list = []*bool{}
	distinct = DistinctBoolPPtr(list)
	if distinct {
		t.Errorf("DistinctBoolPPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctBoolPPtr(nil)
	if distinct {
		t.Errorf("DistinctboolPPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctFloat32P(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v0 float32

	list0 := []float32{v8, v2, v8, v0, v2, v0}
	result := DistinctFloat32P(list0)
	if result {
		t.Errorf("DistinctFloat32P failed. Expected=false, actual=true")
	}

	list0 = []float32{v8, v2, v0}
	result = DistinctFloat32P(list0)
	if !result {
		t.Errorf("DistinctFloat32P failed. Expected=true, actual=false")
	}

	list0 = []float32{}
	result = DistinctFloat32P(list0)
	if result {
		t.Errorf("DistinctFloat32P failed. Expected=false, actual=true")
	}

	result = DistinctFloat32P(nil)
	if result {
		t.Errorf("Distinctfloat32P failed. Expected=false, actual=true")
	}

	list := []*float32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctFloat32PPtr(list)
	if distinct {
		t.Errorf("DistinctFloat32PPtr failed. Expected=false, actual=true")
	}

	list = []*float32{&v8, &v2, &v0}
	distinct = DistinctFloat32PPtr(list)
	if !distinct {
		t.Errorf("DistinctFloat32PPtr failed. Expected=true, actual=false")
	}

	list = []*float32{}
	distinct = DistinctFloat32PPtr(list)
	if distinct {
		t.Errorf("DistinctFloat32PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctFloat32PPtr(nil)
	if distinct {
		t.Errorf("Distinctfloat32PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctFloat64P(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v0 float64

	list0 := []float64{v8, v2, v8, v0, v2, v0}
	result := DistinctFloat64P(list0)
	if result {
		t.Errorf("DistinctFloat64P failed. Expected=false, actual=true")
	}

	list0 = []float64{v8, v2, v0}
	result = DistinctFloat64P(list0)
	if !result {
		t.Errorf("DistinctFloat64P failed. Expected=true, actual=false")
	}

	list0 = []float64{}
	result = DistinctFloat64P(list0)
	if result {
		t.Errorf("DistinctFloat64P failed. Expected=false, actual=true")
	}

	result = DistinctFloat64P(nil)
	if result {
		t.Errorf("Distinctfloat64P failed. Expected=false, actual=true")
	}

	list := []*float64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctFloat64PPtr(list)
	if distinct {
		t.Errorf("DistinctFloat64PPtr failed. Expected=false, actual=true")
	}

	list = []*float64{&v8, &v2, &v0}
	distinct = DistinctFloat64PPtr(list)
	if !distinct {
		t.Errorf("DistinctFloat64PPtr failed. Expected=true, actual=false")
	}

	list = []*float64{}
	distinct = DistinctFloat64PPtr(list)
	if distinct {
		t.Errorf("DistinctFloat64PPtr failed. Expected=false, actual=true")
	}

	distinct = DistinctFloat64PPtr(nil)
	if distinct {
		t.Errorf("Distinctfloat64PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
