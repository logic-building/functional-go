package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEqualMapsInts(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapIntP(map[int]int{v1: v10, v2: v20, v3: v30}, map[int]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntP(map[int]int{v1: v10, v3: v3, v2: v20}, map[int]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntP failed. Expected=false, actual=true")
	}

	r = EqualMapIntP(nil, map[int]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntP failed. Expected=false, actual=true")
	}

	r = EqualMapIntP(map[int]int{v1: v10, v2: v20, v3: v30}, map[int]int{})
	if r {
		t.Errorf("EqualMapIntP failed. Expected=false, actual=true")
	}

	r = EqualMapIntP(nil, map[int]int{})
	if r {
		t.Errorf("EqualMapIntP failed. Expected=false, actual=true")
	}

	r = EqualMapIntP(map[int]int{v1: v10, v2: v20, v3: v30}, map[int]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntPPtr(map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntPPtr(map[*int]*int{&v1: &v10, &v3: &v3, &v2: &v20}, map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntPPtr(nil, map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntPPtr(map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int{})
	if rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntPPtr(nil, map[*int]*int{})
	if rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntPPtr(map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInts64(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapInt64P(map[int64]int64{v1: v10, v2: v20, v3: v30}, map[int64]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64P(map[int64]int64{v1: v10, v3: v3, v2: v20}, map[int64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64P(nil, map[int64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64P(map[int64]int64{v1: v10, v2: v20, v3: v30}, map[int64]int64{})
	if r {
		t.Errorf("EqualMapInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64P(nil, map[int64]int64{})
	if r {
		t.Errorf("EqualMapInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64P(map[int64]int64{v1: v10, v2: v20, v3: v30}, map[int64]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64PPtr(map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64PPtr(map[*int64]*int64{&v1: &v10, &v3: &v3, &v2: &v20}, map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64PPtr(nil, map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64PPtr(map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64PPtr(nil, map[*int64]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64PPtr(map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInts32(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapInt32P(map[int32]int32{v1: v10, v2: v20, v3: v30}, map[int32]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32P(map[int32]int32{v1: v10, v3: v3, v2: v20}, map[int32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32P(nil, map[int32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32P(map[int32]int32{v1: v10, v2: v20, v3: v30}, map[int32]int32{})
	if r {
		t.Errorf("EqualMapInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32P(nil, map[int32]int32{})
	if r {
		t.Errorf("EqualMapInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32P(map[int32]int32{v1: v10, v2: v20, v3: v30}, map[int32]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32PPtr(map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32PPtr(map[*int32]*int32{&v1: &v10, &v3: &v3, &v2: &v20}, map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32PPtr(nil, map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32PPtr(map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32PPtr(nil, map[*int32]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32PPtr(map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInts16(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapInt16P(map[int16]int16{v1: v10, v2: v20, v3: v30}, map[int16]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16P(map[int16]int16{v1: v10, v3: v3, v2: v20}, map[int16]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16P(nil, map[int16]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16P(map[int16]int16{v1: v10, v2: v20, v3: v30}, map[int16]int16{})
	if r {
		t.Errorf("EqualMapInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16P(nil, map[int16]int16{})
	if r {
		t.Errorf("EqualMapInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16P(map[int16]int16{v1: v10, v2: v20, v3: v30}, map[int16]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16PPtr(map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16PPtr(map[*int16]*int16{&v1: &v10, &v3: &v3, &v2: &v20}, map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16PPtr(nil, map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16PPtr(map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16PPtr(nil, map[*int16]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16PPtr(map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInts8(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapInt8P(map[int8]int8{v1: v10, v2: v20, v3: v30}, map[int8]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8P(map[int8]int8{v1: v10, v3: v3, v2: v20}, map[int8]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8P(nil, map[int8]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8P(map[int8]int8{v1: v10, v2: v20, v3: v30}, map[int8]int8{})
	if r {
		t.Errorf("EqualMapInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8P(nil, map[int8]int8{})
	if r {
		t.Errorf("EqualMapInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8P(map[int8]int8{v1: v10, v2: v20, v3: v30}, map[int8]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8PPtr(map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8PPtr(map[*int8]*int8{&v1: &v10, &v3: &v3, &v2: &v20}, map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8PPtr(nil, map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8PPtr(map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8PPtr(nil, map[*int8]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8PPtr(map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUints(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapUintP(map[uint]uint{v1: v10, v2: v20, v3: v30}, map[uint]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintP(map[uint]uint{v1: v10, v3: v3, v2: v20}, map[uint]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintP failed. Expected=false, actual=true")
	}

	r = EqualMapUintP(nil, map[uint]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintP failed. Expected=false, actual=true")
	}

	r = EqualMapUintP(map[uint]uint{v1: v10, v2: v20, v3: v30}, map[uint]uint{})
	if r {
		t.Errorf("EqualMapUintP failed. Expected=false, actual=true")
	}

	r = EqualMapUintP(nil, map[uint]uint{})
	if r {
		t.Errorf("EqualMapUintP failed. Expected=false, actual=true")
	}

	r = EqualMapUintP(map[uint]uint{v1: v10, v2: v20, v3: v30}, map[uint]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintPPtr(map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintPPtr(map[*uint]*uint{&v1: &v10, &v3: &v3, &v2: &v20}, map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintPPtr(nil, map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintPPtr(map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint{})
	if rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintPPtr(nil, map[*uint]*uint{})
	if rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintPPtr(map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64s(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapUint64P(map[uint64]uint64{v1: v10, v2: v20, v3: v30}, map[uint64]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64P(map[uint64]uint64{v1: v10, v3: v3, v2: v20}, map[uint64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64P(nil, map[uint64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64P(map[uint64]uint64{v1: v10, v2: v20, v3: v30}, map[uint64]uint64{})
	if r {
		t.Errorf("EqualMapUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64P(nil, map[uint64]uint64{})
	if r {
		t.Errorf("EqualMapUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64P(map[uint64]uint64{v1: v10, v2: v20, v3: v30}, map[uint64]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64PPtr(map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64PPtr(map[*uint64]*uint64{&v1: &v10, &v3: &v3, &v2: &v20}, map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64PPtr(nil, map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64PPtr(map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64PPtr(nil, map[*uint64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64PPtr(map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUints32(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapUint32P(map[uint32]uint32{v1: v10, v2: v20, v3: v30}, map[uint32]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32P(map[uint32]uint32{v1: v10, v3: v3, v2: v20}, map[uint32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32P(nil, map[uint32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32P(map[uint32]uint32{v1: v10, v2: v20, v3: v30}, map[uint32]uint32{})
	if r {
		t.Errorf("EqualMapUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32P(nil, map[uint32]uint32{})
	if r {
		t.Errorf("EqualMapUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32P(map[uint32]uint32{v1: v10, v2: v20, v3: v30}, map[uint32]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32PPtr(map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32PPtr(map[*uint32]*uint32{&v1: &v10, &v3: &v3, &v2: &v20}, map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32PPtr(nil, map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32PPtr(map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32PPtr(nil, map[*uint32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32PPtr(map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUints16(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapUint16P(map[uint16]uint16{v1: v10, v2: v20, v3: v30}, map[uint16]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16P(map[uint16]uint16{v1: v10, v3: v3, v2: v20}, map[uint16]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16P(nil, map[uint16]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16P(map[uint16]uint16{v1: v10, v2: v20, v3: v30}, map[uint16]uint16{})
	if r {
		t.Errorf("EqualMapUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16P(nil, map[uint16]uint16{})
	if r {
		t.Errorf("EqualMapUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16P(map[uint16]uint16{v1: v10, v2: v20, v3: v30}, map[uint16]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16PPtr(map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16PPtr(map[*uint16]*uint16{&v1: &v10, &v3: &v3, &v2: &v20}, map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16PPtr(nil, map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16PPtr(map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16PPtr(nil, map[*uint16]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16PPtr(map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUints8(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapUint8P(map[uint8]uint8{v1: v10, v2: v20, v3: v30}, map[uint8]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8P(map[uint8]uint8{v1: v10, v3: v3, v2: v20}, map[uint8]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8P(nil, map[uint8]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8P(map[uint8]uint8{v1: v10, v2: v20, v3: v30}, map[uint8]uint8{})
	if r {
		t.Errorf("EqualMapUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8P(nil, map[uint8]uint8{})
	if r {
		t.Errorf("EqualMapUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8P(map[uint8]uint8{v1: v10, v2: v20, v3: v30}, map[uint8]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8PPtr(map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8PPtr(map[*uint8]*uint8{&v1: &v10, &v3: &v3, &v2: &v20}, map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8PPtr(nil, map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8PPtr(map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8PPtr(nil, map[*uint8]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8PPtr(map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrs(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapStrP(map[string]string{v1: v10, v2: v20, v3: v30}, map[string]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrP(map[string]string{v1: v10, v3: v3, v2: v20}, map[string]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrP failed. Expected=false, actual=true")
	}

	r = EqualMapStrP(nil, map[string]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrP failed. Expected=false, actual=true")
	}

	r = EqualMapStrP(map[string]string{v1: v10, v2: v20, v3: v30}, map[string]string{})
	if r {
		t.Errorf("EqualMapStrP failed. Expected=false, actual=true")
	}

	r = EqualMapStrP(nil, map[string]string{})
	if r {
		t.Errorf("EqualMapStrP failed. Expected=false, actual=true")
	}

	r = EqualMapStrP(map[string]string{v1: v10, v2: v20, v3: v30}, map[string]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrPPtr(map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrPPtr(map[*string]*string{&v1: &v10, &v3: &v3, &v2: &v20}, map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrPPtr(nil, map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrPPtr(map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*string{})
	if rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrPPtr(nil, map[*string]*string{})
	if rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrPPtr(map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBools(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	r := EqualMapBoolP(map[bool]bool{v1: v1, v2: v2}, map[bool]bool{v1: v1, v2: v2})
	if !r {
		t.Errorf("EqualMapBoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolP(map[bool]bool{v1: v1, v1: v1}, map[bool]bool{v1: v1, v2: v2})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(nil, map[bool]bool{v1: v1, v2: v2})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v1, v2: v2}, map[bool]bool{})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(nil, map[bool]bool{})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v1, v2: v2}, map[bool]bool{v1: v1})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v2}, map[bool]bool{v1: v1})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolPPtr(map[*bool]*bool{&v1: &v1, &v2: &v2}, map[*bool]*bool{&v1: &v1, &v2: &v2})
	if !rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v1, &v1: &v1}, map[*bool]*bool{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(nil, map[*bool]*bool{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v1, &v2: &v2}, map[*bool]*bool{})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(nil, map[*bool]*bool{})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v1}, map[*bool]*bool{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v2}, map[*bool]*bool{&v1: &v1})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32s(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapFloat32P(map[float32]float32{v1: v10, v2: v20, v3: v30}, map[float32]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32P(map[float32]float32{v1: v10, v3: v3, v2: v20}, map[float32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32P(nil, map[float32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32P(map[float32]float32{v1: v10, v2: v20, v3: v30}, map[float32]float32{})
	if r {
		t.Errorf("EqualMapFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32P(nil, map[float32]float32{})
	if r {
		t.Errorf("EqualMapFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32P(map[float32]float32{v1: v10, v2: v20, v3: v30}, map[float32]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32PPtr(map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32PPtr(map[*float32]*float32{&v1: &v10, &v3: &v3, &v2: &v20}, map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32PPtr(nil, map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32PPtr(map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float32{})
	if rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32PPtr(nil, map[*float32]*float32{})
	if rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32PPtr(map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64s(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapFloat64P(map[float64]float64{v1: v10, v2: v20, v3: v30}, map[float64]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64P(map[float64]float64{v1: v10, v3: v3, v2: v20}, map[float64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64P(nil, map[float64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64P(map[float64]float64{v1: v10, v2: v20, v3: v30}, map[float64]float64{})
	if r {
		t.Errorf("EqualMapFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64P(nil, map[float64]float64{})
	if r {
		t.Errorf("EqualMapFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64P(map[float64]float64{v1: v10, v2: v20, v3: v30}, map[float64]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64PPtr(map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64PPtr(map[*float64]*float64{&v1: &v10, &v3: &v3, &v2: &v20}, map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64PPtr(nil, map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64PPtr(map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float64{})
	if rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64PPtr(nil, map[*float64]*float64{})
	if rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64PPtr(map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64PPtr failed. Expected=false, actual=true")
	}
}
