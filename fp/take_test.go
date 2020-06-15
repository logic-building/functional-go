package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestTakeInt(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v0 int 

	expected := []int{v8, v2, v8}
	list := []int{v8, v2, v8, v0, v2, v0}
	actual := TakeInt(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInts failed")
	}

	expected = []int{}
	actual = TakeInt(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Int failed")
	}
}

func TestTakeIntPtr(t *testing.T) {
	var v8 int = 8
	var v2 int = 2
	var v0 int

	expected := []*int{&v8, &v2, &v8}
	list := []*int{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeIntPtr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeIntPtr failed")
	}

	expected = []*int{}
	actual = TakeIntPtr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt failed")
	}
}
func TestTakeInt64(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v0 int64 

	expected := []int64{v8, v2, v8}
	list := []int64{v8, v2, v8, v0, v2, v0}
	actual := TakeInt64(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInts64 failed")
	}

	expected = []int64{}
	actual = TakeInt64(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Int64 failed")
	}
}

func TestTakeInt64Ptr(t *testing.T) {
	var v8 int64 = 8
	var v2 int64 = 2
	var v0 int64

	expected := []*int64{&v8, &v2, &v8}
	list := []*int64{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeInt64Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt64Ptr failed")
	}

	expected = []*int64{}
	actual = TakeInt64Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt64 failed")
	}
}
func TestTakeInt32(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v0 int32 

	expected := []int32{v8, v2, v8}
	list := []int32{v8, v2, v8, v0, v2, v0}
	actual := TakeInt32(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInts32 failed")
	}

	expected = []int32{}
	actual = TakeInt32(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Int32 failed")
	}
}

func TestTakeInt32Ptr(t *testing.T) {
	var v8 int32 = 8
	var v2 int32 = 2
	var v0 int32

	expected := []*int32{&v8, &v2, &v8}
	list := []*int32{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeInt32Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt32Ptr failed")
	}

	expected = []*int32{}
	actual = TakeInt32Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt32 failed")
	}
}
func TestTakeInt16(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v0 int16 

	expected := []int16{v8, v2, v8}
	list := []int16{v8, v2, v8, v0, v2, v0}
	actual := TakeInt16(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInts16 failed")
	}

	expected = []int16{}
	actual = TakeInt16(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Int16 failed")
	}
}

func TestTakeInt16Ptr(t *testing.T) {
	var v8 int16 = 8
	var v2 int16 = 2
	var v0 int16

	expected := []*int16{&v8, &v2, &v8}
	list := []*int16{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeInt16Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt16Ptr failed")
	}

	expected = []*int16{}
	actual = TakeInt16Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt16 failed")
	}
}
func TestTakeInt8(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v0 int8 

	expected := []int8{v8, v2, v8}
	list := []int8{v8, v2, v8, v0, v2, v0}
	actual := TakeInt8(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInts8 failed")
	}

	expected = []int8{}
	actual = TakeInt8(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Int8 failed")
	}
}

func TestTakeInt8Ptr(t *testing.T) {
	var v8 int8 = 8
	var v2 int8 = 2
	var v0 int8

	expected := []*int8{&v8, &v2, &v8}
	list := []*int8{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeInt8Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt8Ptr failed")
	}

	expected = []*int8{}
	actual = TakeInt8Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeInt8 failed")
	}
}
func TestTakeUint(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v0 uint 

	expected := []uint{v8, v2, v8}
	list := []uint{v8, v2, v8, v0, v2, v0}
	actual := TakeUint(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUints failed")
	}

	expected = []uint{}
	actual = TakeUint(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Uint failed")
	}
}

func TestTakeUintPtr(t *testing.T) {
	var v8 uint = 8
	var v2 uint = 2
	var v0 uint

	expected := []*uint{&v8, &v2, &v8}
	list := []*uint{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeUintPtr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUintPtr failed")
	}

	expected = []*uint{}
	actual = TakeUintPtr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint failed")
	}
}
func TestTakeUint64(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v0 uint64 

	expected := []uint64{v8, v2, v8}
	list := []uint64{v8, v2, v8, v0, v2, v0}
	actual := TakeUint64(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint64s failed")
	}

	expected = []uint64{}
	actual = TakeUint64(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Uint64 failed")
	}
}

func TestTakeUint64Ptr(t *testing.T) {
	var v8 uint64 = 8
	var v2 uint64 = 2
	var v0 uint64

	expected := []*uint64{&v8, &v2, &v8}
	list := []*uint64{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeUint64Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint64Ptr failed")
	}

	expected = []*uint64{}
	actual = TakeUint64Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint64 failed")
	}
}
func TestTakeUint32(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v0 uint32 

	expected := []uint32{v8, v2, v8}
	list := []uint32{v8, v2, v8, v0, v2, v0}
	actual := TakeUint32(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUints32 failed")
	}

	expected = []uint32{}
	actual = TakeUint32(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Uint32 failed")
	}
}

func TestTakeUint32Ptr(t *testing.T) {
	var v8 uint32 = 8
	var v2 uint32 = 2
	var v0 uint32

	expected := []*uint32{&v8, &v2, &v8}
	list := []*uint32{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeUint32Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint32Ptr failed")
	}

	expected = []*uint32{}
	actual = TakeUint32Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint32 failed")
	}
}
func TestTakeUint16(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v0 uint16 

	expected := []uint16{v8, v2, v8}
	list := []uint16{v8, v2, v8, v0, v2, v0}
	actual := TakeUint16(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUints16 failed")
	}

	expected = []uint16{}
	actual = TakeUint16(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Uint16 failed")
	}
}

func TestTakeUint16Ptr(t *testing.T) {
	var v8 uint16 = 8
	var v2 uint16 = 2
	var v0 uint16

	expected := []*uint16{&v8, &v2, &v8}
	list := []*uint16{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeUint16Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint16Ptr failed")
	}

	expected = []*uint16{}
	actual = TakeUint16Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint16 failed")
	}
}
func TestTakeUint8(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v0 uint8 

	expected := []uint8{v8, v2, v8}
	list := []uint8{v8, v2, v8, v0, v2, v0}
	actual := TakeUint8(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUints8 failed")
	}

	expected = []uint8{}
	actual = TakeUint8(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Uint8 failed")
	}
}

func TestTakeUint8Ptr(t *testing.T) {
	var v8 uint8 = 8
	var v2 uint8 = 2
	var v0 uint8

	expected := []*uint8{&v8, &v2, &v8}
	list := []*uint8{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeUint8Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint8Ptr failed")
	}

	expected = []*uint8{}
	actual = TakeUint8Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeUint8 failed")
	}
}
func TestTakeStr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v0 string 

	expected := []string{v8, v2, v8}
	list := []string{v8, v2, v8, v0, v2, v0}
	actual := TakeStr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeStrs failed")
	}

	expected = []string{}
	actual = TakeStr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Str failed")
	}
}

func TestTakeStrPtr(t *testing.T) {
	var v8 string = "8"
	var v2 string = "2"
	var v0 string

	expected := []*string{&v8, &v2, &v8}
	list := []*string{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeStrPtr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeStrPtr failed")
	}

	expected = []*string{}
	actual = TakeStrPtr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeStr failed")
	}
}
func TestTakeBool(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expected := []bool{vt, vf, vt}
	list := []bool{vt, vf, vt, vt, vt, vt}
	actual := TakeBool(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeBools failed")
	}

	expected = []bool{}
	actual = TakeBool(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Bool failed")
	}
}

func TestTakeBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expected := []*bool{&vt, &vf, &vt}
	list := []*bool{&vt, &vf, &vt, &vt, &vt, &vt}
	actual := TakeBoolPtr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeBoolPtr failed")
	}

	expected = []*bool{}
	actual = TakeBoolPtr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeBool failed")
	}
}
func TestTakeFloat32(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v0 float32 

	expected := []float32{v8, v2, v8}
	list := []float32{v8, v2, v8, v0, v2, v0}
	actual := TakeFloat32(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat32s failed")
	}

	expected = []float32{}
	actual = TakeFloat32(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Float32 failed")
	}
}

func TestTakeFloat32Ptr(t *testing.T) {
	var v8 float32 = 8
	var v2 float32 = 2
	var v0 float32

	expected := []*float32{&v8, &v2, &v8}
	list := []*float32{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeFloat32Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat32Ptr failed")
	}

	expected = []*float32{}
	actual = TakeFloat32Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat32 failed")
	}
}
func TestTakeFloat64(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v0 float64 

	expected := []float64{v8, v2, v8}
	list := []float64{v8, v2, v8, v0, v2, v0}
	actual := TakeFloat64(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat64s failed")
	}

	expected = []float64{}
	actual = TakeFloat64(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take Float64 failed")
	}
}

func TestTakeFloat64Ptr(t *testing.T) {
	var v8 float64 = 8
	var v2 float64 = 2
	var v0 float64

	expected := []*float64{&v8, &v2, &v8}
	list := []*float64{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := TakeFloat64Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat64Ptr failed")
	}

	expected = []*float64{}
	actual = TakeFloat64Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test TakeFloat64 failed")
	}
}