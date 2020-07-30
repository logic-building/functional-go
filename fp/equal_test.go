package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEqualInts(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	r := EqualIntsP([]int{v1, v2, v3}, []int{v1, v2, v3})
	if !r {
		t.Errorf("EqualIntsP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualIntsP([]int{v1, v3, v2}, []int{v1, v2, v3})
	if r {
		t.Errorf("EqualIntsP failed. Expected=false, actual=true")
	}

	r = EqualIntsP(nil, []int{v1, v2, v3})
	if r {
		t.Errorf("EqualIntsP failed. Expected=false, actual=true")
	}

	r = EqualIntsP([]int{v1, v2, v3}, []int{})
	if r {
		t.Errorf("EqualIntsP failed. Expected=false, actual=true")
	}

	r = EqualIntsP(nil, []int{})
	if r {
		t.Errorf("EqualIntsP failed. Expected=false, actual=true")
	}

	r = EqualIntsP([]int{v1, v2, v3}, []int{v1, v2})
	if r {
		t.Errorf("EqualIntsP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualIntsPPtr([]*int{&v1, &v2, &v3}, []*int{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualIntsPPtr([]*int{&v1, &v3, &v2}, []*int{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualIntsPPtr(nil, []*int{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualIntsPPtr([]*int{&v1, &v2, &v3}, []*int{})
	if rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualIntsPPtr(nil, []*int{})
	if rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualIntsPPtr([]*int{&v1, &v2, &v3}, []*int{&v1, &v2})
	if rPtr {
		t.Errorf("EqualIntsPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualInts64(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	r := EqualInts64P([]int64{v1, v2, v3}, []int64{v1, v2, v3})
	if !r {
		t.Errorf("EqualInts64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualInts64P([]int64{v1, v3, v2}, []int64{v1, v2, v3})
	if r {
		t.Errorf("EqualInts64P failed. Expected=false, actual=true")
	}

	r = EqualInts64P(nil, []int64{v1, v2, v3})
	if r {
		t.Errorf("EqualInts64P failed. Expected=false, actual=true")
	}

	r = EqualInts64P([]int64{v1, v2, v3}, []int64{})
	if r {
		t.Errorf("EqualInts64P failed. Expected=false, actual=true")
	}

	r = EqualInts64P(nil, []int64{})
	if r {
		t.Errorf("EqualInts64P failed. Expected=false, actual=true")
	}

	r = EqualInts64P([]int64{v1, v2, v3}, []int64{v1, v2})
	if r {
		t.Errorf("EqualInts64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualInts64PPtr([]*int64{&v1, &v2, &v3}, []*int64{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualInts64PPtr([]*int64{&v1, &v3, &v2}, []*int64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts64PPtr(nil, []*int64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts64PPtr([]*int64{&v1, &v2, &v3}, []*int64{})
	if rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts64PPtr(nil, []*int64{})
	if rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts64PPtr([]*int64{&v1, &v2, &v3}, []*int64{&v1, &v2})
	if rPtr {
		t.Errorf("EqualInts64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualInts32(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	r := EqualInts32P([]int32{v1, v2, v3}, []int32{v1, v2, v3})
	if !r {
		t.Errorf("EqualInts32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualInts32P([]int32{v1, v3, v2}, []int32{v1, v2, v3})
	if r {
		t.Errorf("EqualInts32P failed. Expected=false, actual=true")
	}

	r = EqualInts32P(nil, []int32{v1, v2, v3})
	if r {
		t.Errorf("EqualInts32P failed. Expected=false, actual=true")
	}

	r = EqualInts32P([]int32{v1, v2, v3}, []int32{})
	if r {
		t.Errorf("EqualInts32P failed. Expected=false, actual=true")
	}

	r = EqualInts32P(nil, []int32{})
	if r {
		t.Errorf("EqualInts32P failed. Expected=false, actual=true")
	}

	r = EqualInts32P([]int32{v1, v2, v3}, []int32{v1, v2})
	if r {
		t.Errorf("EqualInts32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualInts32PPtr([]*int32{&v1, &v2, &v3}, []*int32{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualInts32PPtr([]*int32{&v1, &v3, &v2}, []*int32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts32PPtr(nil, []*int32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts32PPtr([]*int32{&v1, &v2, &v3}, []*int32{})
	if rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts32PPtr(nil, []*int32{})
	if rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts32PPtr([]*int32{&v1, &v2, &v3}, []*int32{&v1, &v2})
	if rPtr {
		t.Errorf("EqualInts32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualInts16(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	r := EqualInts16P([]int16{v1, v2, v3}, []int16{v1, v2, v3})
	if !r {
		t.Errorf("EqualInts16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualInts16P([]int16{v1, v3, v2}, []int16{v1, v2, v3})
	if r {
		t.Errorf("EqualInts16P failed. Expected=false, actual=true")
	}

	r = EqualInts16P(nil, []int16{v1, v2, v3})
	if r {
		t.Errorf("EqualInts16P failed. Expected=false, actual=true")
	}

	r = EqualInts16P([]int16{v1, v2, v3}, []int16{})
	if r {
		t.Errorf("EqualInts16P failed. Expected=false, actual=true")
	}

	r = EqualInts16P(nil, []int16{})
	if r {
		t.Errorf("EqualInts16P failed. Expected=false, actual=true")
	}

	r = EqualInts16P([]int16{v1, v2, v3}, []int16{v1, v2})
	if r {
		t.Errorf("EqualInts16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualInts16PPtr([]*int16{&v1, &v2, &v3}, []*int16{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualInts16PPtr([]*int16{&v1, &v3, &v2}, []*int16{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts16PPtr(nil, []*int16{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts16PPtr([]*int16{&v1, &v2, &v3}, []*int16{})
	if rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts16PPtr(nil, []*int16{})
	if rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts16PPtr([]*int16{&v1, &v2, &v3}, []*int16{&v1, &v2})
	if rPtr {
		t.Errorf("EqualInts16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualInts8(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	r := EqualInts8P([]int8{v1, v2, v3}, []int8{v1, v2, v3})
	if !r {
		t.Errorf("EqualInts8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualInts8P([]int8{v1, v3, v2}, []int8{v1, v2, v3})
	if r {
		t.Errorf("EqualInts8P failed. Expected=false, actual=true")
	}

	r = EqualInts8P(nil, []int8{v1, v2, v3})
	if r {
		t.Errorf("EqualInts8P failed. Expected=false, actual=true")
	}

	r = EqualInts8P([]int8{v1, v2, v3}, []int8{})
	if r {
		t.Errorf("EqualInts8P failed. Expected=false, actual=true")
	}

	r = EqualInts8P(nil, []int8{})
	if r {
		t.Errorf("EqualInts8P failed. Expected=false, actual=true")
	}

	r = EqualInts8P([]int8{v1, v2, v3}, []int8{v1, v2})
	if r {
		t.Errorf("EqualInts8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualInts8PPtr([]*int8{&v1, &v2, &v3}, []*int8{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualInts8PPtr([]*int8{&v1, &v3, &v2}, []*int8{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts8PPtr(nil, []*int8{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts8PPtr([]*int8{&v1, &v2, &v3}, []*int8{})
	if rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts8PPtr(nil, []*int8{})
	if rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualInts8PPtr([]*int8{&v1, &v2, &v3}, []*int8{&v1, &v2})
	if rPtr {
		t.Errorf("EqualInts8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualUints(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	r := EqualUintsP([]uint{v1, v2, v3}, []uint{v1, v2, v3})
	if !r {
		t.Errorf("EqualUintsP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualUintsP([]uint{v1, v3, v2}, []uint{v1, v2, v3})
	if r {
		t.Errorf("EqualUintsP failed. Expected=false, actual=true")
	}

	r = EqualUintsP(nil, []uint{v1, v2, v3})
	if r {
		t.Errorf("EqualUintsP failed. Expected=false, actual=true")
	}

	r = EqualUintsP([]uint{v1, v2, v3}, []uint{})
	if r {
		t.Errorf("EqualUintsP failed. Expected=false, actual=true")
	}

	r = EqualUintsP(nil, []uint{})
	if r {
		t.Errorf("EqualUintsP failed. Expected=false, actual=true")
	}

	r = EqualUintsP([]uint{v1, v2, v3}, []uint{v1, v2})
	if r {
		t.Errorf("EqualUintsP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualUintsPPtr([]*uint{&v1, &v2, &v3}, []*uint{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualUintsPPtr([]*uint{&v1, &v3, &v2}, []*uint{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUintsPPtr(nil, []*uint{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUintsPPtr([]*uint{&v1, &v2, &v3}, []*uint{})
	if rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUintsPPtr(nil, []*uint{})
	if rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUintsPPtr([]*uint{&v1, &v2, &v3}, []*uint{&v1, &v2})
	if rPtr {
		t.Errorf("EqualUintsPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualUint64s(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	r := EqualUint64sP([]uint64{v1, v2, v3}, []uint64{v1, v2, v3})
	if !r {
		t.Errorf("EqualUint64sP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualUint64sP([]uint64{v1, v3, v2}, []uint64{v1, v2, v3})
	if r {
		t.Errorf("EqualUint64sP failed. Expected=false, actual=true")
	}

	r = EqualUint64sP(nil, []uint64{v1, v2, v3})
	if r {
		t.Errorf("EqualUint64sP failed. Expected=false, actual=true")
	}

	r = EqualUint64sP([]uint64{v1, v2, v3}, []uint64{})
	if r {
		t.Errorf("EqualUint64sP failed. Expected=false, actual=true")
	}

	r = EqualUint64sP(nil, []uint64{})
	if r {
		t.Errorf("EqualUint64sP failed. Expected=false, actual=true")
	}

	r = EqualUint64sP([]uint64{v1, v2, v3}, []uint64{v1, v2})
	if r {
		t.Errorf("EqualUint64sP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualUint64sPPtr([]*uint64{&v1, &v2, &v3}, []*uint64{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualUint64sPPtr([]*uint64{&v1, &v3, &v2}, []*uint64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUint64sPPtr(nil, []*uint64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUint64sPPtr([]*uint64{&v1, &v2, &v3}, []*uint64{})
	if rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUint64sPPtr(nil, []*uint64{})
	if rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUint64sPPtr([]*uint64{&v1, &v2, &v3}, []*uint64{&v1, &v2})
	if rPtr {
		t.Errorf("EqualUint64sPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualUints32(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	r := EqualUints32P([]uint32{v1, v2, v3}, []uint32{v1, v2, v3})
	if !r {
		t.Errorf("EqualUints32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualUints32P([]uint32{v1, v3, v2}, []uint32{v1, v2, v3})
	if r {
		t.Errorf("EqualUints32P failed. Expected=false, actual=true")
	}

	r = EqualUints32P(nil, []uint32{v1, v2, v3})
	if r {
		t.Errorf("EqualUints32P failed. Expected=false, actual=true")
	}

	r = EqualUints32P([]uint32{v1, v2, v3}, []uint32{})
	if r {
		t.Errorf("EqualUints32P failed. Expected=false, actual=true")
	}

	r = EqualUints32P(nil, []uint32{})
	if r {
		t.Errorf("EqualUints32P failed. Expected=false, actual=true")
	}

	r = EqualUints32P([]uint32{v1, v2, v3}, []uint32{v1, v2})
	if r {
		t.Errorf("EqualUints32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualUints32PPtr([]*uint32{&v1, &v2, &v3}, []*uint32{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualUints32PPtr([]*uint32{&v1, &v3, &v2}, []*uint32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints32PPtr(nil, []*uint32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints32PPtr([]*uint32{&v1, &v2, &v3}, []*uint32{})
	if rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints32PPtr(nil, []*uint32{})
	if rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints32PPtr([]*uint32{&v1, &v2, &v3}, []*uint32{&v1, &v2})
	if rPtr {
		t.Errorf("EqualUints32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualUints16(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	r := EqualUints16P([]uint16{v1, v2, v3}, []uint16{v1, v2, v3})
	if !r {
		t.Errorf("EqualUints16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualUints16P([]uint16{v1, v3, v2}, []uint16{v1, v2, v3})
	if r {
		t.Errorf("EqualUints16P failed. Expected=false, actual=true")
	}

	r = EqualUints16P(nil, []uint16{v1, v2, v3})
	if r {
		t.Errorf("EqualUints16P failed. Expected=false, actual=true")
	}

	r = EqualUints16P([]uint16{v1, v2, v3}, []uint16{})
	if r {
		t.Errorf("EqualUints16P failed. Expected=false, actual=true")
	}

	r = EqualUints16P(nil, []uint16{})
	if r {
		t.Errorf("EqualUints16P failed. Expected=false, actual=true")
	}

	r = EqualUints16P([]uint16{v1, v2, v3}, []uint16{v1, v2})
	if r {
		t.Errorf("EqualUints16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualUints16PPtr([]*uint16{&v1, &v2, &v3}, []*uint16{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualUints16PPtr([]*uint16{&v1, &v3, &v2}, []*uint16{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints16PPtr(nil, []*uint16{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints16PPtr([]*uint16{&v1, &v2, &v3}, []*uint16{})
	if rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints16PPtr(nil, []*uint16{})
	if rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints16PPtr([]*uint16{&v1, &v2, &v3}, []*uint16{&v1, &v2})
	if rPtr {
		t.Errorf("EqualUints16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualUints8(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	r := EqualUints8P([]uint8{v1, v2, v3}, []uint8{v1, v2, v3})
	if !r {
		t.Errorf("EqualUints8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualUints8P([]uint8{v1, v3, v2}, []uint8{v1, v2, v3})
	if r {
		t.Errorf("EqualUints8P failed. Expected=false, actual=true")
	}

	r = EqualUints8P(nil, []uint8{v1, v2, v3})
	if r {
		t.Errorf("EqualUints8P failed. Expected=false, actual=true")
	}

	r = EqualUints8P([]uint8{v1, v2, v3}, []uint8{})
	if r {
		t.Errorf("EqualUints8P failed. Expected=false, actual=true")
	}

	r = EqualUints8P(nil, []uint8{})
	if r {
		t.Errorf("EqualUints8P failed. Expected=false, actual=true")
	}

	r = EqualUints8P([]uint8{v1, v2, v3}, []uint8{v1, v2})
	if r {
		t.Errorf("EqualUints8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualUints8PPtr([]*uint8{&v1, &v2, &v3}, []*uint8{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualUints8PPtr([]*uint8{&v1, &v3, &v2}, []*uint8{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints8PPtr(nil, []*uint8{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints8PPtr([]*uint8{&v1, &v2, &v3}, []*uint8{})
	if rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints8PPtr(nil, []*uint8{})
	if rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualUints8PPtr([]*uint8{&v1, &v2, &v3}, []*uint8{&v1, &v2})
	if rPtr {
		t.Errorf("EqualUints8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualStrs(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	r := EqualStrsP([]string{v1, v2, v3}, []string{v1, v2, v3})
	if !r {
		t.Errorf("EqualStrsP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualStrsP([]string{v1, v3, v2}, []string{v1, v2, v3})
	if r {
		t.Errorf("EqualStrsP failed. Expected=false, actual=true")
	}

	r = EqualStrsP(nil, []string{v1, v2, v3})
	if r {
		t.Errorf("EqualStrsP failed. Expected=false, actual=true")
	}

	r = EqualStrsP([]string{v1, v2, v3}, []string{})
	if r {
		t.Errorf("EqualStrsP failed. Expected=false, actual=true")
	}

	r = EqualStrsP(nil, []string{})
	if r {
		t.Errorf("EqualStrsP failed. Expected=false, actual=true")
	}

	r = EqualStrsP([]string{v1, v2, v3}, []string{v1, v2})
	if r {
		t.Errorf("EqualStrsP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualStrsPPtr([]*string{&v1, &v2, &v3}, []*string{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualStrsPPtr([]*string{&v1, &v3, &v2}, []*string{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualStrsPPtr(nil, []*string{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualStrsPPtr([]*string{&v1, &v2, &v3}, []*string{})
	if rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualStrsPPtr(nil, []*string{})
	if rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualStrsPPtr([]*string{&v1, &v2, &v3}, []*string{&v1, &v2})
	if rPtr {
		t.Errorf("EqualStrsPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualBools(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	r := EqualBoolsP([]bool{v1, v2}, []bool{v1, v2})
	if !r {
		t.Errorf("EqualBoolsP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualBoolsP([]bool{v1, v1}, []bool{v1, v2})
	if r {
		t.Errorf("EqualBoolsP failed. Expected=false, actual=true")
	}

	r = EqualBoolsP(nil, []bool{v1, v2})
	if r {
		t.Errorf("EqualBoolsP failed. Expected=false, actual=true")
	}

	r = EqualBoolsP([]bool{v1, v2}, []bool{})
	if r {
		t.Errorf("EqualBoolsP failed. Expected=false, actual=true")
	}

	r = EqualBoolsP(nil, []bool{})
	if r {
		t.Errorf("EqualBoolsP failed. Expected=false, actual=true")
	}

	r = EqualBoolsP([]bool{v1, v2}, []bool{v1})
	if r {
		t.Errorf("EqualBoolsP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualBoolsPPtr([]*bool{&v1, &v2}, []*bool{&v1, &v2})
	if !rPtr {
		t.Errorf("EqualBoolsPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualBoolsPPtr([]*bool{&v1, &v1}, []*bool{&v1, &v2})
	if rPtr {
		t.Errorf("EqualBoolsPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualBoolsPPtr(nil, []*bool{&v1, &v2})
	if rPtr {
		t.Errorf("EqualBoolsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualBoolsPPtr([]*bool{&v1, &v2}, []*bool{})
	if rPtr {
		t.Errorf("EqualBoolsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualBoolsPPtr(nil, []*bool{})
	if rPtr {
		t.Errorf("EqualBoolsPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualBoolsPPtr([]*bool{&v1}, []*bool{&v1, &v2})
	if rPtr {
		t.Errorf("EqualBoolsPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualFloat32s(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	r := EqualFloat32sP([]float32{v1, v2, v3}, []float32{v1, v2, v3})
	if !r {
		t.Errorf("EqualFloat32sP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualFloat32sP([]float32{v1, v3, v2}, []float32{v1, v2, v3})
	if r {
		t.Errorf("EqualFloat32sP failed. Expected=false, actual=true")
	}

	r = EqualFloat32sP(nil, []float32{v1, v2, v3})
	if r {
		t.Errorf("EqualFloat32sP failed. Expected=false, actual=true")
	}

	r = EqualFloat32sP([]float32{v1, v2, v3}, []float32{})
	if r {
		t.Errorf("EqualFloat32sP failed. Expected=false, actual=true")
	}

	r = EqualFloat32sP(nil, []float32{})
	if r {
		t.Errorf("EqualFloat32sP failed. Expected=false, actual=true")
	}

	r = EqualFloat32sP([]float32{v1, v2, v3}, []float32{v1, v2})
	if r {
		t.Errorf("EqualFloat32sP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualFloat32sPPtr([]*float32{&v1, &v2, &v3}, []*float32{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualFloat32sPPtr([]*float32{&v1, &v3, &v2}, []*float32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat32sPPtr(nil, []*float32{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat32sPPtr([]*float32{&v1, &v2, &v3}, []*float32{})
	if rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat32sPPtr(nil, []*float32{})
	if rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat32sPPtr([]*float32{&v1, &v2, &v3}, []*float32{&v1, &v2})
	if rPtr {
		t.Errorf("EqualFloat32sPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualFloat64s(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	r := EqualFloat64sP([]float64{v1, v2, v3}, []float64{v1, v2, v3})
	if !r {
		t.Errorf("EqualFloat64sP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualFloat64sP([]float64{v1, v3, v2}, []float64{v1, v2, v3})
	if r {
		t.Errorf("EqualFloat64sP failed. Expected=false, actual=true")
	}

	r = EqualFloat64sP(nil, []float64{v1, v2, v3})
	if r {
		t.Errorf("EqualFloat64sP failed. Expected=false, actual=true")
	}

	r = EqualFloat64sP([]float64{v1, v2, v3}, []float64{})
	if r {
		t.Errorf("EqualFloat64sP failed. Expected=false, actual=true")
	}

	r = EqualFloat64sP(nil, []float64{})
	if r {
		t.Errorf("EqualFloat64sP failed. Expected=false, actual=true")
	}

	r = EqualFloat64sP([]float64{v1, v2, v3}, []float64{v1, v2})
	if r {
		t.Errorf("EqualFloat64sP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualFloat64sPPtr([]*float64{&v1, &v2, &v3}, []*float64{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualFloat64sPPtr([]*float64{&v1, &v3, &v2}, []*float64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat64sPPtr(nil, []*float64{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat64sPPtr([]*float64{&v1, &v2, &v3}, []*float64{})
	if rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat64sPPtr(nil, []*float64{})
	if rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualFloat64sPPtr([]*float64{&v1, &v2, &v3}, []*float64{&v1, &v2})
	if rPtr {
		t.Errorf("EqualFloat64sPPtr failed. Expected=false, actual=true")
	}
}
