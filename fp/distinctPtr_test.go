package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDistinctIntPtr(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v0 int

	// Test : Get distinct values
	expected := []*int{&v8, &v2, &v0}
	list := []*int{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctIntPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctIntPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int{&v8, &v2, &v0}
	list = []*int{&v8, &v2, &v0}
	distinct = DistinctIntPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctIntPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int{}
	list = []*int{}
	distinct = DistinctIntPtr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctIntPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctIntPtr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctint failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt64Ptr(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v0 int64

	// Test : Get distinct values
	expected := []*int64{&v8, &v2, &v0}
	list := []*int64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int64{&v8, &v2, &v0}
	list = []*int64{&v8, &v2, &v0}
	distinct = DistinctInt64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int64{}
	list = []*int64{}
	distinct = DistinctInt64Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt64Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctint64 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt32Ptr(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v0 int32

	// Test : Get distinct values
	expected := []*int32{&v8, &v2, &v0}
	list := []*int32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int32{&v8, &v2, &v0}
	list = []*int32{&v8, &v2, &v0}
	distinct = DistinctInt32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int32{}
	list = []*int32{}
	distinct = DistinctInt32Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt32Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctint32 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt16Ptr(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v0 int16

	// Test : Get distinct values
	expected := []*int16{&v8, &v2, &v0}
	list := []*int16{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt16Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int16{&v8, &v2, &v0}
	list = []*int16{&v8, &v2, &v0}
	distinct = DistinctInt16Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int16{}
	list = []*int16{}
	distinct = DistinctInt16Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt16Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctint16 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctInt8Ptr(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v0 int8

	// Test : Get distinct values
	expected := []*int8{&v8, &v2, &v0}
	list := []*int8{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctInt8Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int8{&v8, &v2, &v0}
	list = []*int8{&v8, &v2, &v0}
	distinct = DistinctInt8Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctInt8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*int8{}
	list = []*int8{}
	distinct = DistinctInt8Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctInt8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctInt8Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctint8 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUintPtr(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v0 uint

	// Test : Get distinct values
	expected := []*uint{&v8, &v2, &v0}
	list := []*uint{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUintPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUintPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint{&v8, &v2, &v0}
	list = []*uint{&v8, &v2, &v0}
	distinct = DistinctUintPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUintPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint{}
	list = []*uint{}
	distinct = DistinctUintPtr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUintPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUintPtr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctuint failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint64Ptr(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v0 uint64

	// Test : Get distinct values
	expected := []*uint64{&v8, &v2, &v0}
	list := []*uint64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint64{&v8, &v2, &v0}
	list = []*uint64{&v8, &v2, &v0}
	distinct = DistinctUint64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint64{}
	list = []*uint64{}
	distinct = DistinctUint64Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint64Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctuint64 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint32Ptr(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v0 uint32

	// Test : Get distinct values
	expected := []*uint32{&v8, &v2, &v0}
	list := []*uint32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint32{&v8, &v2, &v0}
	list = []*uint32{&v8, &v2, &v0}
	distinct = DistinctUint32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint32{}
	list = []*uint32{}
	distinct = DistinctUint32Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint32Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctuint32 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint16Ptr(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v0 uint16

	// Test : Get distinct values
	expected := []*uint16{&v8, &v2, &v0}
	list := []*uint16{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint16Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint16{&v8, &v2, &v0}
	list = []*uint16{&v8, &v2, &v0}
	distinct = DistinctUint16Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint16{}
	list = []*uint16{}
	distinct = DistinctUint16Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint16Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint16Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctuint16 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctUint8Ptr(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v0 uint8

	// Test : Get distinct values
	expected := []*uint8{&v8, &v2, &v0}
	list := []*uint8{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctUint8Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint8{&v8, &v2, &v0}
	list = []*uint8{&v8, &v2, &v0}
	distinct = DistinctUint8Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctUint8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*uint8{}
	list = []*uint8{}
	distinct = DistinctUint8Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctUint8Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctUint8Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctuint8 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctStrPtr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v0 string

	// Test : Get distinct values
	expected := []*string{&v8, &v2, &v0}
	list := []*string{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctStrPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctStrPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*string{&v8, &v2, &v0}
	list = []*string{&v8, &v2, &v0}
	distinct = DistinctStrPtr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctStrPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*string{}
	list = []*string{}
	distinct = DistinctStrPtr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctStrPtr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctStrPtr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctstring failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctBoolPtr(t *testing.T) {
	var vt bool = true

	newList := DistinctBoolPtr([]*bool{&vt, &vt})
	if *newList[0] != vt {
		t.Errorf("DistinctBoolPtr failed")
	}

	if len(DistinctBoolPtr(nil)) > 0 {
		t.Errorf("DistinctBoolPtr failed.")
	}
}

func TestDistinctFloat32Ptr(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v0 float32

	// Test : Get distinct values
	expected := []*float32{&v8, &v2, &v0}
	list := []*float32{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctFloat32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctFloat32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*float32{&v8, &v2, &v0}
	list = []*float32{&v8, &v2, &v0}
	distinct = DistinctFloat32Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctFloat32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*float32{}
	list = []*float32{}
	distinct = DistinctFloat32Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat32Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctFloat32Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctfloat32 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}

func TestDistinctFloat64Ptr(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v0 float64

	// Test : Get distinct values
	expected := []*float64{&v8, &v2, &v0}
	list := []*float64{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := DistinctFloat64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctFloat64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*float64{&v8, &v2, &v0}
	list = []*float64{&v8, &v2, &v0}
	distinct = DistinctFloat64Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("DistinctFloat64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*float64{}
	list = []*float64{}
	distinct = DistinctFloat64Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("DistinctFloat64Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = DistinctFloat64Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinctfloat64 failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}
