package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEqualMapsInt(t *testing.T) {
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

	r = EqualMapIntP(map[int]int{v1: v10, v3: v10, v2: v20}, map[int]int{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapIntPPtr(map[*int]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*int{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsIntInt64(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapIntInt64P(map[int]int64{v1: v10, v2: v20, v3: v30}, map[int]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntInt64P(map[int]int64{v1: v10, v3: v10, v2: v20}, map[int]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt64P(nil, map[int]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt64P(map[int]int64{v1: v10, v2: v20, v3: v30}, map[int]int64{})
	if r {
		t.Errorf("EqualMapIntInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt64P(nil, map[int]int64{})
	if r {
		t.Errorf("EqualMapIntInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt64P(map[int]int64{v1: v10, v2: v20, v3: v30}, map[int]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntInt64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntInt64PPtr(map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntInt64PPtr(map[*int]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt64PPtr(nil, map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt64PPtr(map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int64{})
	if rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt64PPtr(nil, map[*int]*int64{})
	if rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt64PPtr(map[*int]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntInt32(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapIntInt32P(map[int]int32{v1: v10, v2: v20, v3: v30}, map[int]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntInt32P(map[int]int32{v1: v10, v3: v10, v2: v20}, map[int]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt32P(nil, map[int]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt32P(map[int]int32{v1: v10, v2: v20, v3: v30}, map[int]int32{})
	if r {
		t.Errorf("EqualMapIntInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt32P(nil, map[int]int32{})
	if r {
		t.Errorf("EqualMapIntInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt32P(map[int]int32{v1: v10, v2: v20, v3: v30}, map[int]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntInt32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntInt32PPtr(map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntInt32PPtr(map[*int]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt32PPtr(nil, map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt32PPtr(map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int32{})
	if rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt32PPtr(nil, map[*int]*int32{})
	if rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt32PPtr(map[*int]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntInt16(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapIntInt16P(map[int]int16{v1: v10, v2: v20, v3: v30}, map[int]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntInt16P(map[int]int16{v1: v10, v3: v10, v2: v20}, map[int]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt16P(nil, map[int]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt16P(map[int]int16{v1: v10, v2: v20, v3: v30}, map[int]int16{})
	if r {
		t.Errorf("EqualMapIntInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt16P(nil, map[int]int16{})
	if r {
		t.Errorf("EqualMapIntInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt16P(map[int]int16{v1: v10, v2: v20, v3: v30}, map[int]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntInt16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntInt16PPtr(map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntInt16PPtr(map[*int]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt16PPtr(nil, map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt16PPtr(map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int16{})
	if rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt16PPtr(nil, map[*int]*int16{})
	if rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt16PPtr(map[*int]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntInt8(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapIntInt8P(map[int]int8{v1: v10, v2: v20, v3: v30}, map[int]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntInt8P(map[int]int8{v1: v10, v3: v10, v2: v20}, map[int]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt8P(nil, map[int]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt8P(map[int]int8{v1: v10, v2: v20, v3: v30}, map[int]int8{})
	if r {
		t.Errorf("EqualMapIntInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt8P(nil, map[int]int8{})
	if r {
		t.Errorf("EqualMapIntInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntInt8P(map[int]int8{v1: v10, v2: v20, v3: v30}, map[int]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntInt8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntInt8PPtr(map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntInt8PPtr(map[*int]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt8PPtr(nil, map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt8PPtr(map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int8{})
	if rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt8PPtr(nil, map[*int]*int8{})
	if rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntInt8PPtr(map[*int]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntUint(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapIntUintP(map[int]uint{v1: v10, v2: v20, v3: v30}, map[int]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntUintP(map[int]uint{v1: v10, v3: v10, v2: v20}, map[int]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUintP failed. Expected=false, actual=true")
	}

	r = EqualMapIntUintP(nil, map[int]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUintP failed. Expected=false, actual=true")
	}

	r = EqualMapIntUintP(map[int]uint{v1: v10, v2: v20, v3: v30}, map[int]uint{})
	if r {
		t.Errorf("EqualMapIntUintP failed. Expected=false, actual=true")
	}

	r = EqualMapIntUintP(nil, map[int]uint{})
	if r {
		t.Errorf("EqualMapIntUintP failed. Expected=false, actual=true")
	}

	r = EqualMapIntUintP(map[int]uint{v1: v10, v2: v20, v3: v30}, map[int]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntUintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntUintPPtr(map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntUintPPtr(map[*int]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUintPPtr(nil, map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUintPPtr(map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint{})
	if rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUintPPtr(nil, map[*int]*uint{})
	if rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUintPPtr(map[*int]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntUintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntUint64(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapIntUint64P(map[int]uint64{v1: v10, v2: v20, v3: v30}, map[int]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntUint64P(map[int]uint64{v1: v10, v3: v10, v2: v20}, map[int]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint64P(nil, map[int]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint64P(map[int]uint64{v1: v10, v2: v20, v3: v30}, map[int]uint64{})
	if r {
		t.Errorf("EqualMapIntUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint64P(nil, map[int]uint64{})
	if r {
		t.Errorf("EqualMapIntUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint64P(map[int]uint64{v1: v10, v2: v20, v3: v30}, map[int]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntUint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntUint64PPtr(map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntUint64PPtr(map[*int]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint64PPtr(nil, map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint64PPtr(map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint64{})
	if rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint64PPtr(nil, map[*int]*uint64{})
	if rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint64PPtr(map[*int]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntUint32(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapIntUint32P(map[int]uint32{v1: v10, v2: v20, v3: v30}, map[int]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntUint32P(map[int]uint32{v1: v10, v3: v10, v2: v20}, map[int]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint32P(nil, map[int]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint32P(map[int]uint32{v1: v10, v2: v20, v3: v30}, map[int]uint32{})
	if r {
		t.Errorf("EqualMapIntUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint32P(nil, map[int]uint32{})
	if r {
		t.Errorf("EqualMapIntUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint32P(map[int]uint32{v1: v10, v2: v20, v3: v30}, map[int]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntUint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntUint32PPtr(map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntUint32PPtr(map[*int]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint32PPtr(nil, map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint32PPtr(map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint32{})
	if rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint32PPtr(nil, map[*int]*uint32{})
	if rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint32PPtr(map[*int]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntUint16(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapIntUint16P(map[int]uint16{v1: v10, v2: v20, v3: v30}, map[int]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntUint16P(map[int]uint16{v1: v10, v3: v10, v2: v20}, map[int]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint16P(nil, map[int]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint16P(map[int]uint16{v1: v10, v2: v20, v3: v30}, map[int]uint16{})
	if r {
		t.Errorf("EqualMapIntUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint16P(nil, map[int]uint16{})
	if r {
		t.Errorf("EqualMapIntUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint16P(map[int]uint16{v1: v10, v2: v20, v3: v30}, map[int]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntUint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntUint16PPtr(map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntUint16PPtr(map[*int]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint16PPtr(nil, map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint16PPtr(map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint16{})
	if rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint16PPtr(nil, map[*int]*uint16{})
	if rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint16PPtr(map[*int]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntUint8(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapIntUint8P(map[int]uint8{v1: v10, v2: v20, v3: v30}, map[int]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntUint8P(map[int]uint8{v1: v10, v3: v10, v2: v20}, map[int]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint8P(nil, map[int]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint8P(map[int]uint8{v1: v10, v2: v20, v3: v30}, map[int]uint8{})
	if r {
		t.Errorf("EqualMapIntUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint8P(nil, map[int]uint8{})
	if r {
		t.Errorf("EqualMapIntUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapIntUint8P(map[int]uint8{v1: v10, v2: v20, v3: v30}, map[int]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntUint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntUint8PPtr(map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntUint8PPtr(map[*int]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint8PPtr(nil, map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint8PPtr(map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint8{})
	if rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint8PPtr(nil, map[*int]*uint8{})
	if rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntUint8PPtr(map[*int]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntStr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapIntStrP(map[int]string{v1: v10, v2: v20, v3: v30}, map[int]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntStrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntStrP(map[int]string{v1: v10, v3: v10, v2: v20}, map[int]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntStrP failed. Expected=false, actual=true")
	}

	r = EqualMapIntStrP(nil, map[int]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntStrP failed. Expected=false, actual=true")
	}

	r = EqualMapIntStrP(map[int]string{v1: v10, v2: v20, v3: v30}, map[int]string{})
	if r {
		t.Errorf("EqualMapIntStrP failed. Expected=false, actual=true")
	}

	r = EqualMapIntStrP(nil, map[int]string{})
	if r {
		t.Errorf("EqualMapIntStrP failed. Expected=false, actual=true")
	}

	r = EqualMapIntStrP(map[int]string{v1: v10, v2: v20, v3: v30}, map[int]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntStrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntStrPPtr(map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntStrPPtr(map[*int]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntStrPPtr(nil, map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntStrPPtr(map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*string{})
	if rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntStrPPtr(nil, map[*int]*string{})
	if rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntStrPPtr(map[*int]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntStrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntBool(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapIntBoolP(map[int]bool{v1: v10, v2: v20, v3: v30}, map[int]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntBoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntBoolP(map[int]bool{v1: v20, v3: v10, v2: v20}, map[int]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapIntBoolP(nil, map[int]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapIntBoolP(map[int]bool{v1: v10, v2: v20, v3: v30}, map[int]bool{})
	if r {
		t.Errorf("EqualMapIntBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapIntBoolP(nil, map[int]bool{})
	if r {
		t.Errorf("EqualMapIntBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapIntBoolP(map[int]bool{v1: v10, v2: v20, v3: v30}, map[int]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntBoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntBoolPPtr(map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntBoolPPtr(map[*int]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntBoolPPtr(nil, map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntBoolPPtr(map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*bool{})
	if rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntBoolPPtr(nil, map[*int]*bool{})
	if rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntBoolPPtr(map[*int]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntBoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntFloat32(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapIntFloat32P(map[int]float32{v1: v10, v2: v20, v3: v30}, map[int]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntFloat32P(map[int]float32{v1: v10, v3: v10, v2: v20}, map[int]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat32P(nil, map[int]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat32P(map[int]float32{v1: v10, v2: v20, v3: v30}, map[int]float32{})
	if r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat32P(nil, map[int]float32{})
	if r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat32P(map[int]float32{v1: v10, v2: v20, v3: v30}, map[int]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntFloat32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntFloat32PPtr(map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntFloat32PPtr(map[*int]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat32PPtr(nil, map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat32PPtr(map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float32{})
	if rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat32PPtr(nil, map[*int]*float32{})
	if rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat32PPtr(map[*int]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsIntFloat64(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapIntFloat64P(map[int]float64{v1: v10, v2: v20, v3: v30}, map[int]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapIntFloat64P(map[int]float64{v1: v10, v3: v10, v2: v20}, map[int]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat64P(nil, map[int]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat64P(map[int]float64{v1: v10, v2: v20, v3: v30}, map[int]float64{})
	if r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat64P(nil, map[int]float64{})
	if r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapIntFloat64P(map[int]float64{v1: v10, v2: v20, v3: v30}, map[int]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapIntFloat64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapIntFloat64PPtr(map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapIntFloat64PPtr(map[*int]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat64PPtr(nil, map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat64PPtr(map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float64{})
	if rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat64PPtr(nil, map[*int]*float64{})
	if rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapIntFloat64PPtr(map[*int]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapIntFloat64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Int(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapInt64IntP(map[int64]int{v1: v10, v2: v20, v3: v30}, map[int64]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64IntP(map[int64]int{v1: v10, v3: v10, v2: v20}, map[int64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64IntP(nil, map[int64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64IntP(map[int64]int{v1: v10, v2: v20, v3: v30}, map[int64]int{})
	if r {
		t.Errorf("EqualMapInt64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64IntP(nil, map[int64]int{})
	if r {
		t.Errorf("EqualMapInt64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64IntP(map[int64]int{v1: v10, v2: v20, v3: v30}, map[int64]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64IntPPtr(map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64IntPPtr(map[*int64]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64IntPPtr(nil, map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64IntPPtr(map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int{})
	if rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64IntPPtr(nil, map[*int64]*int{})
	if rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64IntPPtr(map[*int64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64(t *testing.T) {
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

	r = EqualMapInt64P(map[int64]int64{v1: v10, v3: v10, v2: v20}, map[int64]int64{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapInt64PPtr(map[*int64]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsInt64Int32(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapInt64Int32P(map[int64]int32{v1: v10, v2: v20, v3: v30}, map[int64]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Int32P(map[int64]int32{v1: v10, v3: v10, v2: v20}, map[int64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int32P(nil, map[int64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int32P(map[int64]int32{v1: v10, v2: v20, v3: v30}, map[int64]int32{})
	if r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int32P(nil, map[int64]int32{})
	if r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int32P(map[int64]int32{v1: v10, v2: v20, v3: v30}, map[int64]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Int32PPtr(map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Int32PPtr(map[*int64]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int32PPtr(nil, map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int32PPtr(map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int32PPtr(nil, map[*int64]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int32PPtr(map[*int64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Int16(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapInt64Int16P(map[int64]int16{v1: v10, v2: v20, v3: v30}, map[int64]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Int16P(map[int64]int16{v1: v10, v3: v10, v2: v20}, map[int64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int16P(nil, map[int64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int16P(map[int64]int16{v1: v10, v2: v20, v3: v30}, map[int64]int16{})
	if r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int16P(nil, map[int64]int16{})
	if r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int16P(map[int64]int16{v1: v10, v2: v20, v3: v30}, map[int64]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Int16PPtr(map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Int16PPtr(map[*int64]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int16PPtr(nil, map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int16PPtr(map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int16PPtr(nil, map[*int64]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int16PPtr(map[*int64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Int8(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapInt64Int8P(map[int64]int8{v1: v10, v2: v20, v3: v30}, map[int64]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Int8P(map[int64]int8{v1: v10, v3: v10, v2: v20}, map[int64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int8P(nil, map[int64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int8P(map[int64]int8{v1: v10, v2: v20, v3: v30}, map[int64]int8{})
	if r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int8P(nil, map[int64]int8{})
	if r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Int8P(map[int64]int8{v1: v10, v2: v20, v3: v30}, map[int64]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Int8PPtr(map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Int8PPtr(map[*int64]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int8PPtr(nil, map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int8PPtr(map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int8PPtr(nil, map[*int64]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Int8PPtr(map[*int64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Uint(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapInt64UintP(map[int64]uint{v1: v10, v2: v20, v3: v30}, map[int64]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64UintP(map[int64]uint{v1: v10, v3: v10, v2: v20}, map[int64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64UintP(nil, map[int64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64UintP(map[int64]uint{v1: v10, v2: v20, v3: v30}, map[int64]uint{})
	if r {
		t.Errorf("EqualMapInt64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64UintP(nil, map[int64]uint{})
	if r {
		t.Errorf("EqualMapInt64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64UintP(map[int64]uint{v1: v10, v2: v20, v3: v30}, map[int64]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64UintPPtr(map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64UintPPtr(map[*int64]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64UintPPtr(nil, map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64UintPPtr(map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64UintPPtr(nil, map[*int64]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64UintPPtr(map[*int64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Uint64(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapInt64Uint64P(map[int64]uint64{v1: v10, v2: v20, v3: v30}, map[int64]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Uint64P(map[int64]uint64{v1: v10, v3: v10, v2: v20}, map[int64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint64P(nil, map[int64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint64P(map[int64]uint64{v1: v10, v2: v20, v3: v30}, map[int64]uint64{})
	if r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint64P(nil, map[int64]uint64{})
	if r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint64P(map[int64]uint64{v1: v10, v2: v20, v3: v30}, map[int64]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Uint64PPtr(map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Uint64PPtr(map[*int64]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint64PPtr(nil, map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint64PPtr(map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint64PPtr(nil, map[*int64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint64PPtr(map[*int64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Uint32(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapInt64Uint32P(map[int64]uint32{v1: v10, v2: v20, v3: v30}, map[int64]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Uint32P(map[int64]uint32{v1: v10, v3: v10, v2: v20}, map[int64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint32P(nil, map[int64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint32P(map[int64]uint32{v1: v10, v2: v20, v3: v30}, map[int64]uint32{})
	if r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint32P(nil, map[int64]uint32{})
	if r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint32P(map[int64]uint32{v1: v10, v2: v20, v3: v30}, map[int64]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Uint32PPtr(map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Uint32PPtr(map[*int64]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint32PPtr(nil, map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint32PPtr(map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint32PPtr(nil, map[*int64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint32PPtr(map[*int64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Uint16(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapInt64Uint16P(map[int64]uint16{v1: v10, v2: v20, v3: v30}, map[int64]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Uint16P(map[int64]uint16{v1: v10, v3: v10, v2: v20}, map[int64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint16P(nil, map[int64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint16P(map[int64]uint16{v1: v10, v2: v20, v3: v30}, map[int64]uint16{})
	if r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint16P(nil, map[int64]uint16{})
	if r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint16P(map[int64]uint16{v1: v10, v2: v20, v3: v30}, map[int64]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Uint16PPtr(map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Uint16PPtr(map[*int64]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint16PPtr(nil, map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint16PPtr(map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint16PPtr(nil, map[*int64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint16PPtr(map[*int64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Uint8(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapInt64Uint8P(map[int64]uint8{v1: v10, v2: v20, v3: v30}, map[int64]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Uint8P(map[int64]uint8{v1: v10, v3: v10, v2: v20}, map[int64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint8P(nil, map[int64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint8P(map[int64]uint8{v1: v10, v2: v20, v3: v30}, map[int64]uint8{})
	if r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint8P(nil, map[int64]uint8{})
	if r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Uint8P(map[int64]uint8{v1: v10, v2: v20, v3: v30}, map[int64]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Uint8PPtr(map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Uint8PPtr(map[*int64]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint8PPtr(nil, map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint8PPtr(map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint8PPtr(nil, map[*int64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Uint8PPtr(map[*int64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Str(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapInt64StrP(map[int64]string{v1: v10, v2: v20, v3: v30}, map[int64]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64StrP(map[int64]string{v1: v10, v3: v10, v2: v20}, map[int64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64StrP(nil, map[int64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64StrP(map[int64]string{v1: v10, v2: v20, v3: v30}, map[int64]string{})
	if r {
		t.Errorf("EqualMapInt64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64StrP(nil, map[int64]string{})
	if r {
		t.Errorf("EqualMapInt64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64StrP(map[int64]string{v1: v10, v2: v20, v3: v30}, map[int64]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64StrPPtr(map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64StrPPtr(map[*int64]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64StrPPtr(nil, map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64StrPPtr(map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*string{})
	if rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64StrPPtr(nil, map[*int64]*string{})
	if rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64StrPPtr(map[*int64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Bool(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapInt64BoolP(map[int64]bool{v1: v10, v2: v20, v3: v30}, map[int64]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64BoolP(map[int64]bool{v1: v20, v3: v10, v2: v20}, map[int64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64BoolP(nil, map[int64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64BoolP(map[int64]bool{v1: v10, v2: v20, v3: v30}, map[int64]bool{})
	if r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64BoolP(nil, map[int64]bool{})
	if r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt64BoolP(map[int64]bool{v1: v10, v2: v20, v3: v30}, map[int64]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64BoolPPtr(map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64BoolPPtr(map[*int64]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64BoolPPtr(nil, map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64BoolPPtr(map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64BoolPPtr(nil, map[*int64]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64BoolPPtr(map[*int64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Float32(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapInt64Float32P(map[int64]float32{v1: v10, v2: v20, v3: v30}, map[int64]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Float32P(map[int64]float32{v1: v10, v3: v10, v2: v20}, map[int64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float32P(nil, map[int64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float32P(map[int64]float32{v1: v10, v2: v20, v3: v30}, map[int64]float32{})
	if r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float32P(nil, map[int64]float32{})
	if r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float32P(map[int64]float32{v1: v10, v2: v20, v3: v30}, map[int64]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Float32PPtr(map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Float32PPtr(map[*int64]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float32PPtr(nil, map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float32PPtr(map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float32PPtr(nil, map[*int64]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float32PPtr(map[*int64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt64Float64(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapInt64Float64P(map[int64]float64{v1: v10, v2: v20, v3: v30}, map[int64]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt64Float64P(map[int64]float64{v1: v10, v3: v10, v2: v20}, map[int64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float64P(nil, map[int64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float64P(map[int64]float64{v1: v10, v2: v20, v3: v30}, map[int64]float64{})
	if r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float64P(nil, map[int64]float64{})
	if r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt64Float64P(map[int64]float64{v1: v10, v2: v20, v3: v30}, map[int64]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt64Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt64Float64PPtr(map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt64Float64PPtr(map[*int64]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float64PPtr(nil, map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float64PPtr(map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float64PPtr(nil, map[*int64]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt64Float64PPtr(map[*int64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int64]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt64Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Int(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapInt32IntP(map[int32]int{v1: v10, v2: v20, v3: v30}, map[int32]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32IntP(map[int32]int{v1: v10, v3: v10, v2: v20}, map[int32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32IntP(nil, map[int32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32IntP(map[int32]int{v1: v10, v2: v20, v3: v30}, map[int32]int{})
	if r {
		t.Errorf("EqualMapInt32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32IntP(nil, map[int32]int{})
	if r {
		t.Errorf("EqualMapInt32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32IntP(map[int32]int{v1: v10, v2: v20, v3: v30}, map[int32]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32IntPPtr(map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32IntPPtr(map[*int32]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32IntPPtr(nil, map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32IntPPtr(map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int{})
	if rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32IntPPtr(nil, map[*int32]*int{})
	if rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32IntPPtr(map[*int32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Int64(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapInt32Int64P(map[int32]int64{v1: v10, v2: v20, v3: v30}, map[int32]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Int64P(map[int32]int64{v1: v10, v3: v10, v2: v20}, map[int32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int64P(nil, map[int32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int64P(map[int32]int64{v1: v10, v2: v20, v3: v30}, map[int32]int64{})
	if r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int64P(nil, map[int32]int64{})
	if r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int64P(map[int32]int64{v1: v10, v2: v20, v3: v30}, map[int32]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Int64PPtr(map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Int64PPtr(map[*int32]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int64PPtr(nil, map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int64PPtr(map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int64PPtr(nil, map[*int32]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int64PPtr(map[*int32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32(t *testing.T) {
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

	r = EqualMapInt32P(map[int32]int32{v1: v10, v3: v10, v2: v20}, map[int32]int32{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapInt32PPtr(map[*int32]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsInt32Int16(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapInt32Int16P(map[int32]int16{v1: v10, v2: v20, v3: v30}, map[int32]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Int16P(map[int32]int16{v1: v10, v3: v10, v2: v20}, map[int32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int16P(nil, map[int32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int16P(map[int32]int16{v1: v10, v2: v20, v3: v30}, map[int32]int16{})
	if r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int16P(nil, map[int32]int16{})
	if r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int16P(map[int32]int16{v1: v10, v2: v20, v3: v30}, map[int32]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Int16PPtr(map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Int16PPtr(map[*int32]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int16PPtr(nil, map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int16PPtr(map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int16PPtr(nil, map[*int32]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int16PPtr(map[*int32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Int8(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapInt32Int8P(map[int32]int8{v1: v10, v2: v20, v3: v30}, map[int32]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Int8P(map[int32]int8{v1: v10, v3: v10, v2: v20}, map[int32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int8P(nil, map[int32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int8P(map[int32]int8{v1: v10, v2: v20, v3: v30}, map[int32]int8{})
	if r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int8P(nil, map[int32]int8{})
	if r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Int8P(map[int32]int8{v1: v10, v2: v20, v3: v30}, map[int32]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Int8PPtr(map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Int8PPtr(map[*int32]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int8PPtr(nil, map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int8PPtr(map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int8PPtr(nil, map[*int32]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Int8PPtr(map[*int32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Uint(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapInt32UintP(map[int32]uint{v1: v10, v2: v20, v3: v30}, map[int32]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32UintP(map[int32]uint{v1: v10, v3: v10, v2: v20}, map[int32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32UintP(nil, map[int32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32UintP(map[int32]uint{v1: v10, v2: v20, v3: v30}, map[int32]uint{})
	if r {
		t.Errorf("EqualMapInt32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32UintP(nil, map[int32]uint{})
	if r {
		t.Errorf("EqualMapInt32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32UintP(map[int32]uint{v1: v10, v2: v20, v3: v30}, map[int32]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32UintPPtr(map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32UintPPtr(map[*int32]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32UintPPtr(nil, map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32UintPPtr(map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32UintPPtr(nil, map[*int32]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32UintPPtr(map[*int32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Uint64(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapInt32Uint64P(map[int32]uint64{v1: v10, v2: v20, v3: v30}, map[int32]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Uint64P(map[int32]uint64{v1: v10, v3: v10, v2: v20}, map[int32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint64P(nil, map[int32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint64P(map[int32]uint64{v1: v10, v2: v20, v3: v30}, map[int32]uint64{})
	if r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint64P(nil, map[int32]uint64{})
	if r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint64P(map[int32]uint64{v1: v10, v2: v20, v3: v30}, map[int32]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Uint64PPtr(map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Uint64PPtr(map[*int32]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint64PPtr(nil, map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint64PPtr(map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint64PPtr(nil, map[*int32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint64PPtr(map[*int32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Uint32(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapInt32Uint32P(map[int32]uint32{v1: v10, v2: v20, v3: v30}, map[int32]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Uint32P(map[int32]uint32{v1: v10, v3: v10, v2: v20}, map[int32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint32P(nil, map[int32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint32P(map[int32]uint32{v1: v10, v2: v20, v3: v30}, map[int32]uint32{})
	if r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint32P(nil, map[int32]uint32{})
	if r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint32P(map[int32]uint32{v1: v10, v2: v20, v3: v30}, map[int32]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Uint32PPtr(map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Uint32PPtr(map[*int32]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint32PPtr(nil, map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint32PPtr(map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint32PPtr(nil, map[*int32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint32PPtr(map[*int32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Uint16(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapInt32Uint16P(map[int32]uint16{v1: v10, v2: v20, v3: v30}, map[int32]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Uint16P(map[int32]uint16{v1: v10, v3: v10, v2: v20}, map[int32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint16P(nil, map[int32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint16P(map[int32]uint16{v1: v10, v2: v20, v3: v30}, map[int32]uint16{})
	if r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint16P(nil, map[int32]uint16{})
	if r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint16P(map[int32]uint16{v1: v10, v2: v20, v3: v30}, map[int32]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Uint16PPtr(map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Uint16PPtr(map[*int32]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint16PPtr(nil, map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint16PPtr(map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint16PPtr(nil, map[*int32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint16PPtr(map[*int32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Uint8(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapInt32Uint8P(map[int32]uint8{v1: v10, v2: v20, v3: v30}, map[int32]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Uint8P(map[int32]uint8{v1: v10, v3: v10, v2: v20}, map[int32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint8P(nil, map[int32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint8P(map[int32]uint8{v1: v10, v2: v20, v3: v30}, map[int32]uint8{})
	if r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint8P(nil, map[int32]uint8{})
	if r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Uint8P(map[int32]uint8{v1: v10, v2: v20, v3: v30}, map[int32]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Uint8PPtr(map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Uint8PPtr(map[*int32]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint8PPtr(nil, map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint8PPtr(map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint8PPtr(nil, map[*int32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Uint8PPtr(map[*int32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Str(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapInt32StrP(map[int32]string{v1: v10, v2: v20, v3: v30}, map[int32]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32StrP(map[int32]string{v1: v10, v3: v10, v2: v20}, map[int32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32StrP(nil, map[int32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32StrP(map[int32]string{v1: v10, v2: v20, v3: v30}, map[int32]string{})
	if r {
		t.Errorf("EqualMapInt32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32StrP(nil, map[int32]string{})
	if r {
		t.Errorf("EqualMapInt32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32StrP(map[int32]string{v1: v10, v2: v20, v3: v30}, map[int32]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32StrPPtr(map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32StrPPtr(map[*int32]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32StrPPtr(nil, map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32StrPPtr(map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*string{})
	if rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32StrPPtr(nil, map[*int32]*string{})
	if rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32StrPPtr(map[*int32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Bool(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapInt32BoolP(map[int32]bool{v1: v10, v2: v20, v3: v30}, map[int32]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32BoolP(map[int32]bool{v1: v20, v3: v10, v2: v20}, map[int32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32BoolP(nil, map[int32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32BoolP(map[int32]bool{v1: v10, v2: v20, v3: v30}, map[int32]bool{})
	if r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32BoolP(nil, map[int32]bool{})
	if r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt32BoolP(map[int32]bool{v1: v10, v2: v20, v3: v30}, map[int32]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32BoolPPtr(map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32BoolPPtr(map[*int32]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32BoolPPtr(nil, map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32BoolPPtr(map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32BoolPPtr(nil, map[*int32]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32BoolPPtr(map[*int32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Float32(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapInt32Float32P(map[int32]float32{v1: v10, v2: v20, v3: v30}, map[int32]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Float32P(map[int32]float32{v1: v10, v3: v10, v2: v20}, map[int32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float32P(nil, map[int32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float32P(map[int32]float32{v1: v10, v2: v20, v3: v30}, map[int32]float32{})
	if r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float32P(nil, map[int32]float32{})
	if r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float32P(map[int32]float32{v1: v10, v2: v20, v3: v30}, map[int32]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Float32PPtr(map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Float32PPtr(map[*int32]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float32PPtr(nil, map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float32PPtr(map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float32PPtr(nil, map[*int32]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float32PPtr(map[*int32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt32Float64(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapInt32Float64P(map[int32]float64{v1: v10, v2: v20, v3: v30}, map[int32]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt32Float64P(map[int32]float64{v1: v10, v3: v10, v2: v20}, map[int32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float64P(nil, map[int32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float64P(map[int32]float64{v1: v10, v2: v20, v3: v30}, map[int32]float64{})
	if r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float64P(nil, map[int32]float64{})
	if r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt32Float64P(map[int32]float64{v1: v10, v2: v20, v3: v30}, map[int32]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt32Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt32Float64PPtr(map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt32Float64PPtr(map[*int32]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float64PPtr(nil, map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float64PPtr(map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float64PPtr(nil, map[*int32]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt32Float64PPtr(map[*int32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int32]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt32Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Int(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapInt16IntP(map[int16]int{v1: v10, v2: v20, v3: v30}, map[int16]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16IntP(map[int16]int{v1: v10, v3: v10, v2: v20}, map[int16]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16IntP(nil, map[int16]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16IntP(map[int16]int{v1: v10, v2: v20, v3: v30}, map[int16]int{})
	if r {
		t.Errorf("EqualMapInt16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16IntP(nil, map[int16]int{})
	if r {
		t.Errorf("EqualMapInt16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16IntP(map[int16]int{v1: v10, v2: v20, v3: v30}, map[int16]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16IntPPtr(map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16IntPPtr(map[*int16]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16IntPPtr(nil, map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16IntPPtr(map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int{})
	if rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16IntPPtr(nil, map[*int16]*int{})
	if rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16IntPPtr(map[*int16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Int64(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapInt16Int64P(map[int16]int64{v1: v10, v2: v20, v3: v30}, map[int16]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Int64P(map[int16]int64{v1: v10, v3: v10, v2: v20}, map[int16]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int64P(nil, map[int16]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int64P(map[int16]int64{v1: v10, v2: v20, v3: v30}, map[int16]int64{})
	if r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int64P(nil, map[int16]int64{})
	if r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int64P(map[int16]int64{v1: v10, v2: v20, v3: v30}, map[int16]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Int64PPtr(map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Int64PPtr(map[*int16]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int64PPtr(nil, map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int64PPtr(map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int64PPtr(nil, map[*int16]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int64PPtr(map[*int16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Int32(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapInt16Int32P(map[int16]int32{v1: v10, v2: v20, v3: v30}, map[int16]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Int32P(map[int16]int32{v1: v10, v3: v10, v2: v20}, map[int16]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int32P(nil, map[int16]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int32P(map[int16]int32{v1: v10, v2: v20, v3: v30}, map[int16]int32{})
	if r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int32P(nil, map[int16]int32{})
	if r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int32P(map[int16]int32{v1: v10, v2: v20, v3: v30}, map[int16]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Int32PPtr(map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Int32PPtr(map[*int16]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int32PPtr(nil, map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int32PPtr(map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int32PPtr(nil, map[*int16]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int32PPtr(map[*int16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16(t *testing.T) {
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

	r = EqualMapInt16P(map[int16]int16{v1: v10, v3: v10, v2: v20}, map[int16]int16{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapInt16PPtr(map[*int16]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsInt16Int8(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapInt16Int8P(map[int16]int8{v1: v10, v2: v20, v3: v30}, map[int16]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Int8P(map[int16]int8{v1: v10, v3: v10, v2: v20}, map[int16]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int8P(nil, map[int16]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int8P(map[int16]int8{v1: v10, v2: v20, v3: v30}, map[int16]int8{})
	if r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int8P(nil, map[int16]int8{})
	if r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Int8P(map[int16]int8{v1: v10, v2: v20, v3: v30}, map[int16]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Int8PPtr(map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Int8PPtr(map[*int16]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int8PPtr(nil, map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int8PPtr(map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int8PPtr(nil, map[*int16]*int8{})
	if rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Int8PPtr(map[*int16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Uint(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapInt16UintP(map[int16]uint{v1: v10, v2: v20, v3: v30}, map[int16]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16UintP(map[int16]uint{v1: v10, v3: v10, v2: v20}, map[int16]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16UintP(nil, map[int16]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16UintP(map[int16]uint{v1: v10, v2: v20, v3: v30}, map[int16]uint{})
	if r {
		t.Errorf("EqualMapInt16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16UintP(nil, map[int16]uint{})
	if r {
		t.Errorf("EqualMapInt16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16UintP(map[int16]uint{v1: v10, v2: v20, v3: v30}, map[int16]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16UintPPtr(map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16UintPPtr(map[*int16]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16UintPPtr(nil, map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16UintPPtr(map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16UintPPtr(nil, map[*int16]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16UintPPtr(map[*int16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Uint64(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapInt16Uint64P(map[int16]uint64{v1: v10, v2: v20, v3: v30}, map[int16]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Uint64P(map[int16]uint64{v1: v10, v3: v10, v2: v20}, map[int16]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint64P(nil, map[int16]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint64P(map[int16]uint64{v1: v10, v2: v20, v3: v30}, map[int16]uint64{})
	if r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint64P(nil, map[int16]uint64{})
	if r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint64P(map[int16]uint64{v1: v10, v2: v20, v3: v30}, map[int16]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Uint64PPtr(map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Uint64PPtr(map[*int16]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint64PPtr(nil, map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint64PPtr(map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint64PPtr(nil, map[*int16]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint64PPtr(map[*int16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Uint32(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapInt16Uint32P(map[int16]uint32{v1: v10, v2: v20, v3: v30}, map[int16]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Uint32P(map[int16]uint32{v1: v10, v3: v10, v2: v20}, map[int16]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint32P(nil, map[int16]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint32P(map[int16]uint32{v1: v10, v2: v20, v3: v30}, map[int16]uint32{})
	if r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint32P(nil, map[int16]uint32{})
	if r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint32P(map[int16]uint32{v1: v10, v2: v20, v3: v30}, map[int16]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Uint32PPtr(map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Uint32PPtr(map[*int16]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint32PPtr(nil, map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint32PPtr(map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint32PPtr(nil, map[*int16]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint32PPtr(map[*int16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Uint16(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapInt16Uint16P(map[int16]uint16{v1: v10, v2: v20, v3: v30}, map[int16]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Uint16P(map[int16]uint16{v1: v10, v3: v10, v2: v20}, map[int16]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint16P(nil, map[int16]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint16P(map[int16]uint16{v1: v10, v2: v20, v3: v30}, map[int16]uint16{})
	if r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint16P(nil, map[int16]uint16{})
	if r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint16P(map[int16]uint16{v1: v10, v2: v20, v3: v30}, map[int16]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Uint16PPtr(map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Uint16PPtr(map[*int16]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint16PPtr(nil, map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint16PPtr(map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint16PPtr(nil, map[*int16]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint16PPtr(map[*int16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Uint8(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapInt16Uint8P(map[int16]uint8{v1: v10, v2: v20, v3: v30}, map[int16]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Uint8P(map[int16]uint8{v1: v10, v3: v10, v2: v20}, map[int16]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint8P(nil, map[int16]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint8P(map[int16]uint8{v1: v10, v2: v20, v3: v30}, map[int16]uint8{})
	if r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint8P(nil, map[int16]uint8{})
	if r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Uint8P(map[int16]uint8{v1: v10, v2: v20, v3: v30}, map[int16]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Uint8PPtr(map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Uint8PPtr(map[*int16]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint8PPtr(nil, map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint8PPtr(map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint8PPtr(nil, map[*int16]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Uint8PPtr(map[*int16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Str(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapInt16StrP(map[int16]string{v1: v10, v2: v20, v3: v30}, map[int16]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16StrP(map[int16]string{v1: v10, v3: v10, v2: v20}, map[int16]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16StrP(nil, map[int16]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16StrP(map[int16]string{v1: v10, v2: v20, v3: v30}, map[int16]string{})
	if r {
		t.Errorf("EqualMapInt16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16StrP(nil, map[int16]string{})
	if r {
		t.Errorf("EqualMapInt16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16StrP(map[int16]string{v1: v10, v2: v20, v3: v30}, map[int16]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16StrPPtr(map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16StrPPtr(map[*int16]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16StrPPtr(nil, map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16StrPPtr(map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*string{})
	if rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16StrPPtr(nil, map[*int16]*string{})
	if rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16StrPPtr(map[*int16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Bool(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapInt16BoolP(map[int16]bool{v1: v10, v2: v20, v3: v30}, map[int16]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16BoolP(map[int16]bool{v1: v20, v3: v10, v2: v20}, map[int16]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16BoolP(nil, map[int16]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16BoolP(map[int16]bool{v1: v10, v2: v20, v3: v30}, map[int16]bool{})
	if r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16BoolP(nil, map[int16]bool{})
	if r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt16BoolP(map[int16]bool{v1: v10, v2: v20, v3: v30}, map[int16]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16BoolPPtr(map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16BoolPPtr(map[*int16]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16BoolPPtr(nil, map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16BoolPPtr(map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16BoolPPtr(nil, map[*int16]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16BoolPPtr(map[*int16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Float32(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapInt16Float32P(map[int16]float32{v1: v10, v2: v20, v3: v30}, map[int16]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Float32P(map[int16]float32{v1: v10, v3: v10, v2: v20}, map[int16]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float32P(nil, map[int16]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float32P(map[int16]float32{v1: v10, v2: v20, v3: v30}, map[int16]float32{})
	if r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float32P(nil, map[int16]float32{})
	if r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float32P(map[int16]float32{v1: v10, v2: v20, v3: v30}, map[int16]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Float32PPtr(map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Float32PPtr(map[*int16]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float32PPtr(nil, map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float32PPtr(map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float32PPtr(nil, map[*int16]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float32PPtr(map[*int16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt16Float64(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapInt16Float64P(map[int16]float64{v1: v10, v2: v20, v3: v30}, map[int16]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt16Float64P(map[int16]float64{v1: v10, v3: v10, v2: v20}, map[int16]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float64P(nil, map[int16]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float64P(map[int16]float64{v1: v10, v2: v20, v3: v30}, map[int16]float64{})
	if r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float64P(nil, map[int16]float64{})
	if r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt16Float64P(map[int16]float64{v1: v10, v2: v20, v3: v30}, map[int16]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt16Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt16Float64PPtr(map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt16Float64PPtr(map[*int16]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float64PPtr(nil, map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float64PPtr(map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float64PPtr(nil, map[*int16]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt16Float64PPtr(map[*int16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int16]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt16Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Int(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapInt8IntP(map[int8]int{v1: v10, v2: v20, v3: v30}, map[int8]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8IntP(map[int8]int{v1: v10, v3: v10, v2: v20}, map[int8]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8IntP(nil, map[int8]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8IntP(map[int8]int{v1: v10, v2: v20, v3: v30}, map[int8]int{})
	if r {
		t.Errorf("EqualMapInt8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8IntP(nil, map[int8]int{})
	if r {
		t.Errorf("EqualMapInt8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8IntP(map[int8]int{v1: v10, v2: v20, v3: v30}, map[int8]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8IntPPtr(map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8IntPPtr(map[*int8]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8IntPPtr(nil, map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8IntPPtr(map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int{})
	if rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8IntPPtr(nil, map[*int8]*int{})
	if rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8IntPPtr(map[*int8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Int64(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapInt8Int64P(map[int8]int64{v1: v10, v2: v20, v3: v30}, map[int8]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Int64P(map[int8]int64{v1: v10, v3: v10, v2: v20}, map[int8]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int64P(nil, map[int8]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int64P(map[int8]int64{v1: v10, v2: v20, v3: v30}, map[int8]int64{})
	if r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int64P(nil, map[int8]int64{})
	if r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int64P(map[int8]int64{v1: v10, v2: v20, v3: v30}, map[int8]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Int64PPtr(map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Int64PPtr(map[*int8]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int64PPtr(nil, map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int64PPtr(map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int64PPtr(nil, map[*int8]*int64{})
	if rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int64PPtr(map[*int8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Int32(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapInt8Int32P(map[int8]int32{v1: v10, v2: v20, v3: v30}, map[int8]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Int32P(map[int8]int32{v1: v10, v3: v10, v2: v20}, map[int8]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int32P(nil, map[int8]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int32P(map[int8]int32{v1: v10, v2: v20, v3: v30}, map[int8]int32{})
	if r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int32P(nil, map[int8]int32{})
	if r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int32P(map[int8]int32{v1: v10, v2: v20, v3: v30}, map[int8]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Int32PPtr(map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Int32PPtr(map[*int8]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int32PPtr(nil, map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int32PPtr(map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int32PPtr(nil, map[*int8]*int32{})
	if rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int32PPtr(map[*int8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Int16(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapInt8Int16P(map[int8]int16{v1: v10, v2: v20, v3: v30}, map[int8]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Int16P(map[int8]int16{v1: v10, v3: v10, v2: v20}, map[int8]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int16P(nil, map[int8]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int16P(map[int8]int16{v1: v10, v2: v20, v3: v30}, map[int8]int16{})
	if r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int16P(nil, map[int8]int16{})
	if r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Int16P(map[int8]int16{v1: v10, v2: v20, v3: v30}, map[int8]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Int16PPtr(map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Int16PPtr(map[*int8]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int16PPtr(nil, map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int16PPtr(map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int16PPtr(nil, map[*int8]*int16{})
	if rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Int16PPtr(map[*int8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8(t *testing.T) {
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

	r = EqualMapInt8P(map[int8]int8{v1: v10, v3: v10, v2: v20}, map[int8]int8{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapInt8PPtr(map[*int8]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsInt8Uint(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapInt8UintP(map[int8]uint{v1: v10, v2: v20, v3: v30}, map[int8]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8UintP(map[int8]uint{v1: v10, v3: v10, v2: v20}, map[int8]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8UintP(nil, map[int8]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8UintP(map[int8]uint{v1: v10, v2: v20, v3: v30}, map[int8]uint{})
	if r {
		t.Errorf("EqualMapInt8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8UintP(nil, map[int8]uint{})
	if r {
		t.Errorf("EqualMapInt8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8UintP(map[int8]uint{v1: v10, v2: v20, v3: v30}, map[int8]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8UintPPtr(map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8UintPPtr(map[*int8]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8UintPPtr(nil, map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8UintPPtr(map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8UintPPtr(nil, map[*int8]*uint{})
	if rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8UintPPtr(map[*int8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Uint64(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapInt8Uint64P(map[int8]uint64{v1: v10, v2: v20, v3: v30}, map[int8]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Uint64P(map[int8]uint64{v1: v10, v3: v10, v2: v20}, map[int8]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint64P(nil, map[int8]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint64P(map[int8]uint64{v1: v10, v2: v20, v3: v30}, map[int8]uint64{})
	if r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint64P(nil, map[int8]uint64{})
	if r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint64P(map[int8]uint64{v1: v10, v2: v20, v3: v30}, map[int8]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Uint64PPtr(map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Uint64PPtr(map[*int8]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint64PPtr(nil, map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint64PPtr(map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint64PPtr(nil, map[*int8]*uint64{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint64PPtr(map[*int8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Uint32(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapInt8Uint32P(map[int8]uint32{v1: v10, v2: v20, v3: v30}, map[int8]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Uint32P(map[int8]uint32{v1: v10, v3: v10, v2: v20}, map[int8]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint32P(nil, map[int8]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint32P(map[int8]uint32{v1: v10, v2: v20, v3: v30}, map[int8]uint32{})
	if r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint32P(nil, map[int8]uint32{})
	if r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint32P(map[int8]uint32{v1: v10, v2: v20, v3: v30}, map[int8]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Uint32PPtr(map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Uint32PPtr(map[*int8]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint32PPtr(nil, map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint32PPtr(map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint32PPtr(nil, map[*int8]*uint32{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint32PPtr(map[*int8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Uint16(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapInt8Uint16P(map[int8]uint16{v1: v10, v2: v20, v3: v30}, map[int8]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Uint16P(map[int8]uint16{v1: v10, v3: v10, v2: v20}, map[int8]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint16P(nil, map[int8]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint16P(map[int8]uint16{v1: v10, v2: v20, v3: v30}, map[int8]uint16{})
	if r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint16P(nil, map[int8]uint16{})
	if r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint16P(map[int8]uint16{v1: v10, v2: v20, v3: v30}, map[int8]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Uint16PPtr(map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Uint16PPtr(map[*int8]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint16PPtr(nil, map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint16PPtr(map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint16PPtr(nil, map[*int8]*uint16{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint16PPtr(map[*int8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Uint8(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapInt8Uint8P(map[int8]uint8{v1: v10, v2: v20, v3: v30}, map[int8]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Uint8P(map[int8]uint8{v1: v10, v3: v10, v2: v20}, map[int8]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint8P(nil, map[int8]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint8P(map[int8]uint8{v1: v10, v2: v20, v3: v30}, map[int8]uint8{})
	if r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint8P(nil, map[int8]uint8{})
	if r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Uint8P(map[int8]uint8{v1: v10, v2: v20, v3: v30}, map[int8]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Uint8PPtr(map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Uint8PPtr(map[*int8]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint8PPtr(nil, map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint8PPtr(map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint8PPtr(nil, map[*int8]*uint8{})
	if rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Uint8PPtr(map[*int8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Str(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapInt8StrP(map[int8]string{v1: v10, v2: v20, v3: v30}, map[int8]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8StrP(map[int8]string{v1: v10, v3: v10, v2: v20}, map[int8]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8StrP(nil, map[int8]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8StrP(map[int8]string{v1: v10, v2: v20, v3: v30}, map[int8]string{})
	if r {
		t.Errorf("EqualMapInt8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8StrP(nil, map[int8]string{})
	if r {
		t.Errorf("EqualMapInt8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8StrP(map[int8]string{v1: v10, v2: v20, v3: v30}, map[int8]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8StrPPtr(map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8StrPPtr(map[*int8]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8StrPPtr(nil, map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8StrPPtr(map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*string{})
	if rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8StrPPtr(nil, map[*int8]*string{})
	if rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8StrPPtr(map[*int8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Bool(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapInt8BoolP(map[int8]bool{v1: v10, v2: v20, v3: v30}, map[int8]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8BoolP(map[int8]bool{v1: v20, v3: v10, v2: v20}, map[int8]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8BoolP(nil, map[int8]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8BoolP(map[int8]bool{v1: v10, v2: v20, v3: v30}, map[int8]bool{})
	if r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8BoolP(nil, map[int8]bool{})
	if r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapInt8BoolP(map[int8]bool{v1: v10, v2: v20, v3: v30}, map[int8]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8BoolPPtr(map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8BoolPPtr(map[*int8]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8BoolPPtr(nil, map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8BoolPPtr(map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8BoolPPtr(nil, map[*int8]*bool{})
	if rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8BoolPPtr(map[*int8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Float32(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapInt8Float32P(map[int8]float32{v1: v10, v2: v20, v3: v30}, map[int8]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Float32P(map[int8]float32{v1: v10, v3: v10, v2: v20}, map[int8]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float32P(nil, map[int8]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float32P(map[int8]float32{v1: v10, v2: v20, v3: v30}, map[int8]float32{})
	if r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float32P(nil, map[int8]float32{})
	if r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float32P(map[int8]float32{v1: v10, v2: v20, v3: v30}, map[int8]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Float32PPtr(map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Float32PPtr(map[*int8]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float32PPtr(nil, map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float32PPtr(map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float32PPtr(nil, map[*int8]*float32{})
	if rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float32PPtr(map[*int8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsInt8Float64(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapInt8Float64P(map[int8]float64{v1: v10, v2: v20, v3: v30}, map[int8]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapInt8Float64P(map[int8]float64{v1: v10, v3: v10, v2: v20}, map[int8]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float64P(nil, map[int8]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float64P(map[int8]float64{v1: v10, v2: v20, v3: v30}, map[int8]float64{})
	if r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float64P(nil, map[int8]float64{})
	if r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapInt8Float64P(map[int8]float64{v1: v10, v2: v20, v3: v30}, map[int8]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapInt8Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapInt8Float64PPtr(map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapInt8Float64PPtr(map[*int8]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float64PPtr(nil, map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float64PPtr(map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float64PPtr(nil, map[*int8]*float64{})
	if rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapInt8Float64PPtr(map[*int8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*int8]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapInt8Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintInt(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapUintIntP(map[uint]int{v1: v10, v2: v20, v3: v30}, map[uint]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintIntP(map[uint]int{v1: v10, v3: v10, v2: v20}, map[uint]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintIntP failed. Expected=false, actual=true")
	}

	r = EqualMapUintIntP(nil, map[uint]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintIntP failed. Expected=false, actual=true")
	}

	r = EqualMapUintIntP(map[uint]int{v1: v10, v2: v20, v3: v30}, map[uint]int{})
	if r {
		t.Errorf("EqualMapUintIntP failed. Expected=false, actual=true")
	}

	r = EqualMapUintIntP(nil, map[uint]int{})
	if r {
		t.Errorf("EqualMapUintIntP failed. Expected=false, actual=true")
	}

	r = EqualMapUintIntP(map[uint]int{v1: v10, v2: v20, v3: v30}, map[uint]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintIntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintIntPPtr(map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintIntPPtr(map[*uint]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintIntPPtr(nil, map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintIntPPtr(map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int{})
	if rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintIntPPtr(nil, map[*uint]*int{})
	if rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintIntPPtr(map[*uint]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintIntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintInt64(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapUintInt64P(map[uint]int64{v1: v10, v2: v20, v3: v30}, map[uint]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintInt64P(map[uint]int64{v1: v10, v3: v10, v2: v20}, map[uint]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt64P(nil, map[uint]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt64P(map[uint]int64{v1: v10, v2: v20, v3: v30}, map[uint]int64{})
	if r {
		t.Errorf("EqualMapUintInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt64P(nil, map[uint]int64{})
	if r {
		t.Errorf("EqualMapUintInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt64P(map[uint]int64{v1: v10, v2: v20, v3: v30}, map[uint]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintInt64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintInt64PPtr(map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintInt64PPtr(map[*uint]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt64PPtr(nil, map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt64PPtr(map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int64{})
	if rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt64PPtr(nil, map[*uint]*int64{})
	if rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt64PPtr(map[*uint]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintInt32(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapUintInt32P(map[uint]int32{v1: v10, v2: v20, v3: v30}, map[uint]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintInt32P(map[uint]int32{v1: v10, v3: v10, v2: v20}, map[uint]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt32P(nil, map[uint]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt32P(map[uint]int32{v1: v10, v2: v20, v3: v30}, map[uint]int32{})
	if r {
		t.Errorf("EqualMapUintInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt32P(nil, map[uint]int32{})
	if r {
		t.Errorf("EqualMapUintInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt32P(map[uint]int32{v1: v10, v2: v20, v3: v30}, map[uint]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintInt32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintInt32PPtr(map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintInt32PPtr(map[*uint]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt32PPtr(nil, map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt32PPtr(map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int32{})
	if rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt32PPtr(nil, map[*uint]*int32{})
	if rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt32PPtr(map[*uint]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintInt16(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapUintInt16P(map[uint]int16{v1: v10, v2: v20, v3: v30}, map[uint]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintInt16P(map[uint]int16{v1: v10, v3: v10, v2: v20}, map[uint]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt16P(nil, map[uint]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt16P(map[uint]int16{v1: v10, v2: v20, v3: v30}, map[uint]int16{})
	if r {
		t.Errorf("EqualMapUintInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt16P(nil, map[uint]int16{})
	if r {
		t.Errorf("EqualMapUintInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt16P(map[uint]int16{v1: v10, v2: v20, v3: v30}, map[uint]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintInt16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintInt16PPtr(map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintInt16PPtr(map[*uint]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt16PPtr(nil, map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt16PPtr(map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int16{})
	if rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt16PPtr(nil, map[*uint]*int16{})
	if rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt16PPtr(map[*uint]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintInt8(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapUintInt8P(map[uint]int8{v1: v10, v2: v20, v3: v30}, map[uint]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintInt8P(map[uint]int8{v1: v10, v3: v10, v2: v20}, map[uint]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt8P(nil, map[uint]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt8P(map[uint]int8{v1: v10, v2: v20, v3: v30}, map[uint]int8{})
	if r {
		t.Errorf("EqualMapUintInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt8P(nil, map[uint]int8{})
	if r {
		t.Errorf("EqualMapUintInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintInt8P(map[uint]int8{v1: v10, v2: v20, v3: v30}, map[uint]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintInt8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintInt8PPtr(map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintInt8PPtr(map[*uint]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt8PPtr(nil, map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt8PPtr(map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int8{})
	if rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt8PPtr(nil, map[*uint]*int8{})
	if rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintInt8PPtr(map[*uint]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint(t *testing.T) {
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

	r = EqualMapUintP(map[uint]uint{v1: v10, v3: v10, v2: v20}, map[uint]uint{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapUintPPtr(map[*uint]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsUintUint64(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapUintUint64P(map[uint]uint64{v1: v10, v2: v20, v3: v30}, map[uint]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintUint64P(map[uint]uint64{v1: v10, v3: v10, v2: v20}, map[uint]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint64P(nil, map[uint]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint64P(map[uint]uint64{v1: v10, v2: v20, v3: v30}, map[uint]uint64{})
	if r {
		t.Errorf("EqualMapUintUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint64P(nil, map[uint]uint64{})
	if r {
		t.Errorf("EqualMapUintUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint64P(map[uint]uint64{v1: v10, v2: v20, v3: v30}, map[uint]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintUint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintUint64PPtr(map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintUint64PPtr(map[*uint]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint64PPtr(nil, map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint64PPtr(map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint64PPtr(nil, map[*uint]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint64PPtr(map[*uint]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintUint32(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapUintUint32P(map[uint]uint32{v1: v10, v2: v20, v3: v30}, map[uint]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintUint32P(map[uint]uint32{v1: v10, v3: v10, v2: v20}, map[uint]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint32P(nil, map[uint]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint32P(map[uint]uint32{v1: v10, v2: v20, v3: v30}, map[uint]uint32{})
	if r {
		t.Errorf("EqualMapUintUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint32P(nil, map[uint]uint32{})
	if r {
		t.Errorf("EqualMapUintUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint32P(map[uint]uint32{v1: v10, v2: v20, v3: v30}, map[uint]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintUint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintUint32PPtr(map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintUint32PPtr(map[*uint]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint32PPtr(nil, map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint32PPtr(map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint32PPtr(nil, map[*uint]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint32PPtr(map[*uint]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintUint16(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapUintUint16P(map[uint]uint16{v1: v10, v2: v20, v3: v30}, map[uint]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintUint16P(map[uint]uint16{v1: v10, v3: v10, v2: v20}, map[uint]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint16P(nil, map[uint]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint16P(map[uint]uint16{v1: v10, v2: v20, v3: v30}, map[uint]uint16{})
	if r {
		t.Errorf("EqualMapUintUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint16P(nil, map[uint]uint16{})
	if r {
		t.Errorf("EqualMapUintUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint16P(map[uint]uint16{v1: v10, v2: v20, v3: v30}, map[uint]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintUint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintUint16PPtr(map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintUint16PPtr(map[*uint]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint16PPtr(nil, map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint16PPtr(map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint16PPtr(nil, map[*uint]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint16PPtr(map[*uint]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintUint8(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapUintUint8P(map[uint]uint8{v1: v10, v2: v20, v3: v30}, map[uint]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintUint8P(map[uint]uint8{v1: v10, v3: v10, v2: v20}, map[uint]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint8P(nil, map[uint]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint8P(map[uint]uint8{v1: v10, v2: v20, v3: v30}, map[uint]uint8{})
	if r {
		t.Errorf("EqualMapUintUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint8P(nil, map[uint]uint8{})
	if r {
		t.Errorf("EqualMapUintUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUintUint8P(map[uint]uint8{v1: v10, v2: v20, v3: v30}, map[uint]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintUint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintUint8PPtr(map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintUint8PPtr(map[*uint]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint8PPtr(nil, map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint8PPtr(map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint8PPtr(nil, map[*uint]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintUint8PPtr(map[*uint]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintStr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapUintStrP(map[uint]string{v1: v10, v2: v20, v3: v30}, map[uint]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintStrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintStrP(map[uint]string{v1: v10, v3: v10, v2: v20}, map[uint]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintStrP failed. Expected=false, actual=true")
	}

	r = EqualMapUintStrP(nil, map[uint]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintStrP failed. Expected=false, actual=true")
	}

	r = EqualMapUintStrP(map[uint]string{v1: v10, v2: v20, v3: v30}, map[uint]string{})
	if r {
		t.Errorf("EqualMapUintStrP failed. Expected=false, actual=true")
	}

	r = EqualMapUintStrP(nil, map[uint]string{})
	if r {
		t.Errorf("EqualMapUintStrP failed. Expected=false, actual=true")
	}

	r = EqualMapUintStrP(map[uint]string{v1: v10, v2: v20, v3: v30}, map[uint]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintStrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintStrPPtr(map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintStrPPtr(map[*uint]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintStrPPtr(nil, map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintStrPPtr(map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*string{})
	if rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintStrPPtr(nil, map[*uint]*string{})
	if rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintStrPPtr(map[*uint]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintStrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintBool(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapUintBoolP(map[uint]bool{v1: v10, v2: v20, v3: v30}, map[uint]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintBoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintBoolP(map[uint]bool{v1: v20, v3: v10, v2: v20}, map[uint]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUintBoolP(nil, map[uint]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUintBoolP(map[uint]bool{v1: v10, v2: v20, v3: v30}, map[uint]bool{})
	if r {
		t.Errorf("EqualMapUintBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUintBoolP(nil, map[uint]bool{})
	if r {
		t.Errorf("EqualMapUintBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUintBoolP(map[uint]bool{v1: v10, v2: v20, v3: v30}, map[uint]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintBoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintBoolPPtr(map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintBoolPPtr(map[*uint]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintBoolPPtr(nil, map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintBoolPPtr(map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*bool{})
	if rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintBoolPPtr(nil, map[*uint]*bool{})
	if rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintBoolPPtr(map[*uint]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintBoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintFloat32(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapUintFloat32P(map[uint]float32{v1: v10, v2: v20, v3: v30}, map[uint]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintFloat32P(map[uint]float32{v1: v10, v3: v10, v2: v20}, map[uint]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat32P(nil, map[uint]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat32P(map[uint]float32{v1: v10, v2: v20, v3: v30}, map[uint]float32{})
	if r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat32P(nil, map[uint]float32{})
	if r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat32P(map[uint]float32{v1: v10, v2: v20, v3: v30}, map[uint]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintFloat32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintFloat32PPtr(map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintFloat32PPtr(map[*uint]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat32PPtr(nil, map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat32PPtr(map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float32{})
	if rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat32PPtr(nil, map[*uint]*float32{})
	if rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat32PPtr(map[*uint]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUintFloat64(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapUintFloat64P(map[uint]float64{v1: v10, v2: v20, v3: v30}, map[uint]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUintFloat64P(map[uint]float64{v1: v10, v3: v10, v2: v20}, map[uint]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat64P(nil, map[uint]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat64P(map[uint]float64{v1: v10, v2: v20, v3: v30}, map[uint]float64{})
	if r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat64P(nil, map[uint]float64{})
	if r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapUintFloat64P(map[uint]float64{v1: v10, v2: v20, v3: v30}, map[uint]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUintFloat64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUintFloat64PPtr(map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUintFloat64PPtr(map[*uint]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat64PPtr(nil, map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat64PPtr(map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float64{})
	if rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat64PPtr(nil, map[*uint]*float64{})
	if rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUintFloat64PPtr(map[*uint]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUintFloat64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Int(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapUint64IntP(map[uint64]int{v1: v10, v2: v20, v3: v30}, map[uint64]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64IntP(map[uint64]int{v1: v10, v3: v10, v2: v20}, map[uint64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64IntP(nil, map[uint64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64IntP(map[uint64]int{v1: v10, v2: v20, v3: v30}, map[uint64]int{})
	if r {
		t.Errorf("EqualMapUint64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64IntP(nil, map[uint64]int{})
	if r {
		t.Errorf("EqualMapUint64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64IntP(map[uint64]int{v1: v10, v2: v20, v3: v30}, map[uint64]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64IntPPtr(map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64IntPPtr(map[*uint64]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64IntPPtr(nil, map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64IntPPtr(map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int{})
	if rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64IntPPtr(nil, map[*uint64]*int{})
	if rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64IntPPtr(map[*uint64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Int64(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapUint64Int64P(map[uint64]int64{v1: v10, v2: v20, v3: v30}, map[uint64]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Int64P(map[uint64]int64{v1: v10, v3: v10, v2: v20}, map[uint64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int64P(nil, map[uint64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int64P(map[uint64]int64{v1: v10, v2: v20, v3: v30}, map[uint64]int64{})
	if r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int64P(nil, map[uint64]int64{})
	if r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int64P(map[uint64]int64{v1: v10, v2: v20, v3: v30}, map[uint64]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Int64PPtr(map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Int64PPtr(map[*uint64]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int64PPtr(nil, map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int64PPtr(map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int64PPtr(nil, map[*uint64]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int64PPtr(map[*uint64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Int32(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapUint64Int32P(map[uint64]int32{v1: v10, v2: v20, v3: v30}, map[uint64]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Int32P(map[uint64]int32{v1: v10, v3: v10, v2: v20}, map[uint64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int32P(nil, map[uint64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int32P(map[uint64]int32{v1: v10, v2: v20, v3: v30}, map[uint64]int32{})
	if r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int32P(nil, map[uint64]int32{})
	if r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int32P(map[uint64]int32{v1: v10, v2: v20, v3: v30}, map[uint64]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Int32PPtr(map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Int32PPtr(map[*uint64]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int32PPtr(nil, map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int32PPtr(map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int32PPtr(nil, map[*uint64]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int32PPtr(map[*uint64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Int16(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapUint64Int16P(map[uint64]int16{v1: v10, v2: v20, v3: v30}, map[uint64]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Int16P(map[uint64]int16{v1: v10, v3: v10, v2: v20}, map[uint64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int16P(nil, map[uint64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int16P(map[uint64]int16{v1: v10, v2: v20, v3: v30}, map[uint64]int16{})
	if r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int16P(nil, map[uint64]int16{})
	if r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int16P(map[uint64]int16{v1: v10, v2: v20, v3: v30}, map[uint64]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Int16PPtr(map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Int16PPtr(map[*uint64]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int16PPtr(nil, map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int16PPtr(map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int16PPtr(nil, map[*uint64]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int16PPtr(map[*uint64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Int8(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapUint64Int8P(map[uint64]int8{v1: v10, v2: v20, v3: v30}, map[uint64]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Int8P(map[uint64]int8{v1: v10, v3: v10, v2: v20}, map[uint64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int8P(nil, map[uint64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int8P(map[uint64]int8{v1: v10, v2: v20, v3: v30}, map[uint64]int8{})
	if r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int8P(nil, map[uint64]int8{})
	if r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Int8P(map[uint64]int8{v1: v10, v2: v20, v3: v30}, map[uint64]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Int8PPtr(map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Int8PPtr(map[*uint64]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int8PPtr(nil, map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int8PPtr(map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int8PPtr(nil, map[*uint64]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Int8PPtr(map[*uint64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Uint(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapUint64UintP(map[uint64]uint{v1: v10, v2: v20, v3: v30}, map[uint64]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64UintP(map[uint64]uint{v1: v10, v3: v10, v2: v20}, map[uint64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64UintP(nil, map[uint64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64UintP(map[uint64]uint{v1: v10, v2: v20, v3: v30}, map[uint64]uint{})
	if r {
		t.Errorf("EqualMapUint64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64UintP(nil, map[uint64]uint{})
	if r {
		t.Errorf("EqualMapUint64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64UintP(map[uint64]uint{v1: v10, v2: v20, v3: v30}, map[uint64]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64UintPPtr(map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64UintPPtr(map[*uint64]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64UintPPtr(nil, map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64UintPPtr(map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64UintPPtr(nil, map[*uint64]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64UintPPtr(map[*uint64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64(t *testing.T) {
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

	r = EqualMapUint64P(map[uint64]uint64{v1: v10, v3: v10, v2: v20}, map[uint64]uint64{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapUint64PPtr(map[*uint64]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsUint64Uint32(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapUint64Uint32P(map[uint64]uint32{v1: v10, v2: v20, v3: v30}, map[uint64]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Uint32P(map[uint64]uint32{v1: v10, v3: v10, v2: v20}, map[uint64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint32P(nil, map[uint64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint32P(map[uint64]uint32{v1: v10, v2: v20, v3: v30}, map[uint64]uint32{})
	if r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint32P(nil, map[uint64]uint32{})
	if r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint32P(map[uint64]uint32{v1: v10, v2: v20, v3: v30}, map[uint64]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Uint32PPtr(map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Uint32PPtr(map[*uint64]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint32PPtr(nil, map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint32PPtr(map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint32PPtr(nil, map[*uint64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint32PPtr(map[*uint64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Uint16(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapUint64Uint16P(map[uint64]uint16{v1: v10, v2: v20, v3: v30}, map[uint64]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Uint16P(map[uint64]uint16{v1: v10, v3: v10, v2: v20}, map[uint64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint16P(nil, map[uint64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint16P(map[uint64]uint16{v1: v10, v2: v20, v3: v30}, map[uint64]uint16{})
	if r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint16P(nil, map[uint64]uint16{})
	if r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint16P(map[uint64]uint16{v1: v10, v2: v20, v3: v30}, map[uint64]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Uint16PPtr(map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Uint16PPtr(map[*uint64]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint16PPtr(nil, map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint16PPtr(map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint16PPtr(nil, map[*uint64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint16PPtr(map[*uint64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Uint8(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapUint64Uint8P(map[uint64]uint8{v1: v10, v2: v20, v3: v30}, map[uint64]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Uint8P(map[uint64]uint8{v1: v10, v3: v10, v2: v20}, map[uint64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint8P(nil, map[uint64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint8P(map[uint64]uint8{v1: v10, v2: v20, v3: v30}, map[uint64]uint8{})
	if r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint8P(nil, map[uint64]uint8{})
	if r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Uint8P(map[uint64]uint8{v1: v10, v2: v20, v3: v30}, map[uint64]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Uint8PPtr(map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Uint8PPtr(map[*uint64]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint8PPtr(nil, map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint8PPtr(map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint8PPtr(nil, map[*uint64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Uint8PPtr(map[*uint64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Str(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapUint64StrP(map[uint64]string{v1: v10, v2: v20, v3: v30}, map[uint64]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64StrP(map[uint64]string{v1: v10, v3: v10, v2: v20}, map[uint64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64StrP(nil, map[uint64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64StrP(map[uint64]string{v1: v10, v2: v20, v3: v30}, map[uint64]string{})
	if r {
		t.Errorf("EqualMapUint64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64StrP(nil, map[uint64]string{})
	if r {
		t.Errorf("EqualMapUint64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64StrP(map[uint64]string{v1: v10, v2: v20, v3: v30}, map[uint64]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64StrPPtr(map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64StrPPtr(map[*uint64]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64StrPPtr(nil, map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64StrPPtr(map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*string{})
	if rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64StrPPtr(nil, map[*uint64]*string{})
	if rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64StrPPtr(map[*uint64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Bool(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapUint64BoolP(map[uint64]bool{v1: v10, v2: v20, v3: v30}, map[uint64]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64BoolP(map[uint64]bool{v1: v20, v3: v10, v2: v20}, map[uint64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64BoolP(nil, map[uint64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64BoolP(map[uint64]bool{v1: v10, v2: v20, v3: v30}, map[uint64]bool{})
	if r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64BoolP(nil, map[uint64]bool{})
	if r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint64BoolP(map[uint64]bool{v1: v10, v2: v20, v3: v30}, map[uint64]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64BoolPPtr(map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64BoolPPtr(map[*uint64]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64BoolPPtr(nil, map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64BoolPPtr(map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64BoolPPtr(nil, map[*uint64]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64BoolPPtr(map[*uint64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Float32(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapUint64Float32P(map[uint64]float32{v1: v10, v2: v20, v3: v30}, map[uint64]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Float32P(map[uint64]float32{v1: v10, v3: v10, v2: v20}, map[uint64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float32P(nil, map[uint64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float32P(map[uint64]float32{v1: v10, v2: v20, v3: v30}, map[uint64]float32{})
	if r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float32P(nil, map[uint64]float32{})
	if r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float32P(map[uint64]float32{v1: v10, v2: v20, v3: v30}, map[uint64]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Float32PPtr(map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Float32PPtr(map[*uint64]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float32PPtr(nil, map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float32PPtr(map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float32PPtr(nil, map[*uint64]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float32PPtr(map[*uint64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint64Float64(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapUint64Float64P(map[uint64]float64{v1: v10, v2: v20, v3: v30}, map[uint64]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint64Float64P(map[uint64]float64{v1: v10, v3: v10, v2: v20}, map[uint64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float64P(nil, map[uint64]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float64P(map[uint64]float64{v1: v10, v2: v20, v3: v30}, map[uint64]float64{})
	if r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float64P(nil, map[uint64]float64{})
	if r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint64Float64P(map[uint64]float64{v1: v10, v2: v20, v3: v30}, map[uint64]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint64Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint64Float64PPtr(map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint64Float64PPtr(map[*uint64]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float64PPtr(nil, map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float64PPtr(map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float64PPtr(nil, map[*uint64]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint64Float64PPtr(map[*uint64]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint64]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint64Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Int(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapUint32IntP(map[uint32]int{v1: v10, v2: v20, v3: v30}, map[uint32]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32IntP(map[uint32]int{v1: v10, v3: v10, v2: v20}, map[uint32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32IntP(nil, map[uint32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32IntP(map[uint32]int{v1: v10, v2: v20, v3: v30}, map[uint32]int{})
	if r {
		t.Errorf("EqualMapUint32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32IntP(nil, map[uint32]int{})
	if r {
		t.Errorf("EqualMapUint32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32IntP(map[uint32]int{v1: v10, v2: v20, v3: v30}, map[uint32]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32IntPPtr(map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32IntPPtr(map[*uint32]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32IntPPtr(nil, map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32IntPPtr(map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int{})
	if rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32IntPPtr(nil, map[*uint32]*int{})
	if rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32IntPPtr(map[*uint32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Int64(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapUint32Int64P(map[uint32]int64{v1: v10, v2: v20, v3: v30}, map[uint32]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Int64P(map[uint32]int64{v1: v10, v3: v10, v2: v20}, map[uint32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int64P(nil, map[uint32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int64P(map[uint32]int64{v1: v10, v2: v20, v3: v30}, map[uint32]int64{})
	if r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int64P(nil, map[uint32]int64{})
	if r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int64P(map[uint32]int64{v1: v10, v2: v20, v3: v30}, map[uint32]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Int64PPtr(map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Int64PPtr(map[*uint32]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int64PPtr(nil, map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int64PPtr(map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int64PPtr(nil, map[*uint32]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int64PPtr(map[*uint32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Int32(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapUint32Int32P(map[uint32]int32{v1: v10, v2: v20, v3: v30}, map[uint32]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Int32P(map[uint32]int32{v1: v10, v3: v10, v2: v20}, map[uint32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int32P(nil, map[uint32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int32P(map[uint32]int32{v1: v10, v2: v20, v3: v30}, map[uint32]int32{})
	if r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int32P(nil, map[uint32]int32{})
	if r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int32P(map[uint32]int32{v1: v10, v2: v20, v3: v30}, map[uint32]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Int32PPtr(map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Int32PPtr(map[*uint32]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int32PPtr(nil, map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int32PPtr(map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int32PPtr(nil, map[*uint32]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int32PPtr(map[*uint32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Int16(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapUint32Int16P(map[uint32]int16{v1: v10, v2: v20, v3: v30}, map[uint32]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Int16P(map[uint32]int16{v1: v10, v3: v10, v2: v20}, map[uint32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int16P(nil, map[uint32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int16P(map[uint32]int16{v1: v10, v2: v20, v3: v30}, map[uint32]int16{})
	if r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int16P(nil, map[uint32]int16{})
	if r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int16P(map[uint32]int16{v1: v10, v2: v20, v3: v30}, map[uint32]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Int16PPtr(map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Int16PPtr(map[*uint32]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int16PPtr(nil, map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int16PPtr(map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int16PPtr(nil, map[*uint32]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int16PPtr(map[*uint32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Int8(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapUint32Int8P(map[uint32]int8{v1: v10, v2: v20, v3: v30}, map[uint32]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Int8P(map[uint32]int8{v1: v10, v3: v10, v2: v20}, map[uint32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int8P(nil, map[uint32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int8P(map[uint32]int8{v1: v10, v2: v20, v3: v30}, map[uint32]int8{})
	if r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int8P(nil, map[uint32]int8{})
	if r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Int8P(map[uint32]int8{v1: v10, v2: v20, v3: v30}, map[uint32]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Int8PPtr(map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Int8PPtr(map[*uint32]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int8PPtr(nil, map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int8PPtr(map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int8PPtr(nil, map[*uint32]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Int8PPtr(map[*uint32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Uint(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapUint32UintP(map[uint32]uint{v1: v10, v2: v20, v3: v30}, map[uint32]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32UintP(map[uint32]uint{v1: v10, v3: v10, v2: v20}, map[uint32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32UintP(nil, map[uint32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32UintP(map[uint32]uint{v1: v10, v2: v20, v3: v30}, map[uint32]uint{})
	if r {
		t.Errorf("EqualMapUint32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32UintP(nil, map[uint32]uint{})
	if r {
		t.Errorf("EqualMapUint32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32UintP(map[uint32]uint{v1: v10, v2: v20, v3: v30}, map[uint32]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32UintPPtr(map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32UintPPtr(map[*uint32]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32UintPPtr(nil, map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32UintPPtr(map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32UintPPtr(nil, map[*uint32]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32UintPPtr(map[*uint32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Uint64(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapUint32Uint64P(map[uint32]uint64{v1: v10, v2: v20, v3: v30}, map[uint32]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Uint64P(map[uint32]uint64{v1: v10, v3: v10, v2: v20}, map[uint32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint64P(nil, map[uint32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint64P(map[uint32]uint64{v1: v10, v2: v20, v3: v30}, map[uint32]uint64{})
	if r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint64P(nil, map[uint32]uint64{})
	if r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint64P(map[uint32]uint64{v1: v10, v2: v20, v3: v30}, map[uint32]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Uint64PPtr(map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Uint64PPtr(map[*uint32]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint64PPtr(nil, map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint64PPtr(map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint64PPtr(nil, map[*uint32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint64PPtr(map[*uint32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32(t *testing.T) {
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

	r = EqualMapUint32P(map[uint32]uint32{v1: v10, v3: v10, v2: v20}, map[uint32]uint32{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapUint32PPtr(map[*uint32]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsUint32Uint16(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapUint32Uint16P(map[uint32]uint16{v1: v10, v2: v20, v3: v30}, map[uint32]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Uint16P(map[uint32]uint16{v1: v10, v3: v10, v2: v20}, map[uint32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint16P(nil, map[uint32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint16P(map[uint32]uint16{v1: v10, v2: v20, v3: v30}, map[uint32]uint16{})
	if r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint16P(nil, map[uint32]uint16{})
	if r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint16P(map[uint32]uint16{v1: v10, v2: v20, v3: v30}, map[uint32]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Uint16PPtr(map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Uint16PPtr(map[*uint32]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint16PPtr(nil, map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint16PPtr(map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint16PPtr(nil, map[*uint32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint16PPtr(map[*uint32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Uint8(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapUint32Uint8P(map[uint32]uint8{v1: v10, v2: v20, v3: v30}, map[uint32]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Uint8P(map[uint32]uint8{v1: v10, v3: v10, v2: v20}, map[uint32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint8P(nil, map[uint32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint8P(map[uint32]uint8{v1: v10, v2: v20, v3: v30}, map[uint32]uint8{})
	if r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint8P(nil, map[uint32]uint8{})
	if r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Uint8P(map[uint32]uint8{v1: v10, v2: v20, v3: v30}, map[uint32]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Uint8PPtr(map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Uint8PPtr(map[*uint32]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint8PPtr(nil, map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint8PPtr(map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint8PPtr(nil, map[*uint32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Uint8PPtr(map[*uint32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Str(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapUint32StrP(map[uint32]string{v1: v10, v2: v20, v3: v30}, map[uint32]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32StrP(map[uint32]string{v1: v10, v3: v10, v2: v20}, map[uint32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32StrP(nil, map[uint32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32StrP(map[uint32]string{v1: v10, v2: v20, v3: v30}, map[uint32]string{})
	if r {
		t.Errorf("EqualMapUint32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32StrP(nil, map[uint32]string{})
	if r {
		t.Errorf("EqualMapUint32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32StrP(map[uint32]string{v1: v10, v2: v20, v3: v30}, map[uint32]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32StrPPtr(map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32StrPPtr(map[*uint32]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32StrPPtr(nil, map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32StrPPtr(map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*string{})
	if rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32StrPPtr(nil, map[*uint32]*string{})
	if rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32StrPPtr(map[*uint32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Bool(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapUint32BoolP(map[uint32]bool{v1: v10, v2: v20, v3: v30}, map[uint32]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32BoolP(map[uint32]bool{v1: v20, v3: v10, v2: v20}, map[uint32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32BoolP(nil, map[uint32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32BoolP(map[uint32]bool{v1: v10, v2: v20, v3: v30}, map[uint32]bool{})
	if r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32BoolP(nil, map[uint32]bool{})
	if r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint32BoolP(map[uint32]bool{v1: v10, v2: v20, v3: v30}, map[uint32]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32BoolPPtr(map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32BoolPPtr(map[*uint32]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32BoolPPtr(nil, map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32BoolPPtr(map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32BoolPPtr(nil, map[*uint32]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32BoolPPtr(map[*uint32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Float32(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapUint32Float32P(map[uint32]float32{v1: v10, v2: v20, v3: v30}, map[uint32]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Float32P(map[uint32]float32{v1: v10, v3: v10, v2: v20}, map[uint32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float32P(nil, map[uint32]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float32P(map[uint32]float32{v1: v10, v2: v20, v3: v30}, map[uint32]float32{})
	if r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float32P(nil, map[uint32]float32{})
	if r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float32P(map[uint32]float32{v1: v10, v2: v20, v3: v30}, map[uint32]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Float32PPtr(map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Float32PPtr(map[*uint32]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float32PPtr(nil, map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float32PPtr(map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float32PPtr(nil, map[*uint32]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float32PPtr(map[*uint32]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint32Float64(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapUint32Float64P(map[uint32]float64{v1: v10, v2: v20, v3: v30}, map[uint32]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint32Float64P(map[uint32]float64{v1: v10, v3: v10, v2: v20}, map[uint32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float64P(nil, map[uint32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float64P(map[uint32]float64{v1: v10, v2: v20, v3: v30}, map[uint32]float64{})
	if r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float64P(nil, map[uint32]float64{})
	if r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint32Float64P(map[uint32]float64{v1: v10, v2: v20, v3: v30}, map[uint32]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint32Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint32Float64PPtr(map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint32Float64PPtr(map[*uint32]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float64PPtr(nil, map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float64PPtr(map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float64PPtr(nil, map[*uint32]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint32Float64PPtr(map[*uint32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint32]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint32Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Int(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapUint16IntP(map[uint16]int{v1: v10, v2: v20, v3: v30}, map[uint16]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16IntP(map[uint16]int{v1: v10, v3: v10, v2: v20}, map[uint16]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16IntP(nil, map[uint16]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16IntP(map[uint16]int{v1: v10, v2: v20, v3: v30}, map[uint16]int{})
	if r {
		t.Errorf("EqualMapUint16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16IntP(nil, map[uint16]int{})
	if r {
		t.Errorf("EqualMapUint16IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16IntP(map[uint16]int{v1: v10, v2: v20, v3: v30}, map[uint16]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16IntPPtr(map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16IntPPtr(map[*uint16]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16IntPPtr(nil, map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16IntPPtr(map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int{})
	if rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16IntPPtr(nil, map[*uint16]*int{})
	if rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16IntPPtr(map[*uint16]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Int64(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapUint16Int64P(map[uint16]int64{v1: v10, v2: v20, v3: v30}, map[uint16]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Int64P(map[uint16]int64{v1: v10, v3: v10, v2: v20}, map[uint16]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int64P(nil, map[uint16]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int64P(map[uint16]int64{v1: v10, v2: v20, v3: v30}, map[uint16]int64{})
	if r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int64P(nil, map[uint16]int64{})
	if r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int64P(map[uint16]int64{v1: v10, v2: v20, v3: v30}, map[uint16]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Int64PPtr(map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Int64PPtr(map[*uint16]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int64PPtr(nil, map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int64PPtr(map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int64PPtr(nil, map[*uint16]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int64PPtr(map[*uint16]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Int32(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapUint16Int32P(map[uint16]int32{v1: v10, v2: v20, v3: v30}, map[uint16]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Int32P(map[uint16]int32{v1: v10, v3: v10, v2: v20}, map[uint16]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int32P(nil, map[uint16]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int32P(map[uint16]int32{v1: v10, v2: v20, v3: v30}, map[uint16]int32{})
	if r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int32P(nil, map[uint16]int32{})
	if r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int32P(map[uint16]int32{v1: v10, v2: v20, v3: v30}, map[uint16]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Int32PPtr(map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Int32PPtr(map[*uint16]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int32PPtr(nil, map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int32PPtr(map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int32PPtr(nil, map[*uint16]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int32PPtr(map[*uint16]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Int16(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapUint16Int16P(map[uint16]int16{v1: v10, v2: v20, v3: v30}, map[uint16]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Int16P(map[uint16]int16{v1: v10, v3: v10, v2: v20}, map[uint16]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int16P(nil, map[uint16]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int16P(map[uint16]int16{v1: v10, v2: v20, v3: v30}, map[uint16]int16{})
	if r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int16P(nil, map[uint16]int16{})
	if r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int16P(map[uint16]int16{v1: v10, v2: v20, v3: v30}, map[uint16]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Int16PPtr(map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Int16PPtr(map[*uint16]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int16PPtr(nil, map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int16PPtr(map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int16PPtr(nil, map[*uint16]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int16PPtr(map[*uint16]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Int8(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapUint16Int8P(map[uint16]int8{v1: v10, v2: v20, v3: v30}, map[uint16]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Int8P(map[uint16]int8{v1: v10, v3: v10, v2: v20}, map[uint16]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int8P(nil, map[uint16]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int8P(map[uint16]int8{v1: v10, v2: v20, v3: v30}, map[uint16]int8{})
	if r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int8P(nil, map[uint16]int8{})
	if r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Int8P(map[uint16]int8{v1: v10, v2: v20, v3: v30}, map[uint16]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Int8PPtr(map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Int8PPtr(map[*uint16]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int8PPtr(nil, map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int8PPtr(map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int8PPtr(nil, map[*uint16]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Int8PPtr(map[*uint16]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Uint(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapUint16UintP(map[uint16]uint{v1: v10, v2: v20, v3: v30}, map[uint16]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16UintP(map[uint16]uint{v1: v10, v3: v10, v2: v20}, map[uint16]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16UintP(nil, map[uint16]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16UintP(map[uint16]uint{v1: v10, v2: v20, v3: v30}, map[uint16]uint{})
	if r {
		t.Errorf("EqualMapUint16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16UintP(nil, map[uint16]uint{})
	if r {
		t.Errorf("EqualMapUint16UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16UintP(map[uint16]uint{v1: v10, v2: v20, v3: v30}, map[uint16]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16UintPPtr(map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16UintPPtr(map[*uint16]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16UintPPtr(nil, map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16UintPPtr(map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16UintPPtr(nil, map[*uint16]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16UintPPtr(map[*uint16]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Uint64(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapUint16Uint64P(map[uint16]uint64{v1: v10, v2: v20, v3: v30}, map[uint16]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Uint64P(map[uint16]uint64{v1: v10, v3: v10, v2: v20}, map[uint16]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint64P(nil, map[uint16]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint64P(map[uint16]uint64{v1: v10, v2: v20, v3: v30}, map[uint16]uint64{})
	if r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint64P(nil, map[uint16]uint64{})
	if r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint64P(map[uint16]uint64{v1: v10, v2: v20, v3: v30}, map[uint16]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Uint64PPtr(map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Uint64PPtr(map[*uint16]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint64PPtr(nil, map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint64PPtr(map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint64PPtr(nil, map[*uint16]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint64PPtr(map[*uint16]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Uint32(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapUint16Uint32P(map[uint16]uint32{v1: v10, v2: v20, v3: v30}, map[uint16]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Uint32P(map[uint16]uint32{v1: v10, v3: v10, v2: v20}, map[uint16]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint32P(nil, map[uint16]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint32P(map[uint16]uint32{v1: v10, v2: v20, v3: v30}, map[uint16]uint32{})
	if r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint32P(nil, map[uint16]uint32{})
	if r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint32P(map[uint16]uint32{v1: v10, v2: v20, v3: v30}, map[uint16]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Uint32PPtr(map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Uint32PPtr(map[*uint16]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint32PPtr(nil, map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint32PPtr(map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint32PPtr(nil, map[*uint16]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint32PPtr(map[*uint16]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16(t *testing.T) {
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

	r = EqualMapUint16P(map[uint16]uint16{v1: v10, v3: v10, v2: v20}, map[uint16]uint16{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapUint16PPtr(map[*uint16]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsUint16Uint8(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapUint16Uint8P(map[uint16]uint8{v1: v10, v2: v20, v3: v30}, map[uint16]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Uint8P(map[uint16]uint8{v1: v10, v3: v10, v2: v20}, map[uint16]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint8P(nil, map[uint16]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint8P(map[uint16]uint8{v1: v10, v2: v20, v3: v30}, map[uint16]uint8{})
	if r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint8P(nil, map[uint16]uint8{})
	if r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Uint8P(map[uint16]uint8{v1: v10, v2: v20, v3: v30}, map[uint16]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Uint8PPtr(map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Uint8PPtr(map[*uint16]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint8PPtr(nil, map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint8PPtr(map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint8PPtr(nil, map[*uint16]*uint8{})
	if rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Uint8PPtr(map[*uint16]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Str(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapUint16StrP(map[uint16]string{v1: v10, v2: v20, v3: v30}, map[uint16]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16StrP(map[uint16]string{v1: v10, v3: v10, v2: v20}, map[uint16]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16StrP(nil, map[uint16]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16StrP(map[uint16]string{v1: v10, v2: v20, v3: v30}, map[uint16]string{})
	if r {
		t.Errorf("EqualMapUint16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16StrP(nil, map[uint16]string{})
	if r {
		t.Errorf("EqualMapUint16StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16StrP(map[uint16]string{v1: v10, v2: v20, v3: v30}, map[uint16]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16StrPPtr(map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16StrPPtr(map[*uint16]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16StrPPtr(nil, map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16StrPPtr(map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*string{})
	if rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16StrPPtr(nil, map[*uint16]*string{})
	if rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16StrPPtr(map[*uint16]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Bool(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapUint16BoolP(map[uint16]bool{v1: v10, v2: v20, v3: v30}, map[uint16]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16BoolP(map[uint16]bool{v1: v20, v3: v10, v2: v20}, map[uint16]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16BoolP(nil, map[uint16]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16BoolP(map[uint16]bool{v1: v10, v2: v20, v3: v30}, map[uint16]bool{})
	if r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16BoolP(nil, map[uint16]bool{})
	if r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint16BoolP(map[uint16]bool{v1: v10, v2: v20, v3: v30}, map[uint16]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16BoolPPtr(map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16BoolPPtr(map[*uint16]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16BoolPPtr(nil, map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16BoolPPtr(map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16BoolPPtr(nil, map[*uint16]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16BoolPPtr(map[*uint16]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Float32(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapUint16Float32P(map[uint16]float32{v1: v10, v2: v20, v3: v30}, map[uint16]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Float32P(map[uint16]float32{v1: v10, v3: v10, v2: v20}, map[uint16]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float32P(nil, map[uint16]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float32P(map[uint16]float32{v1: v10, v2: v20, v3: v30}, map[uint16]float32{})
	if r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float32P(nil, map[uint16]float32{})
	if r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float32P(map[uint16]float32{v1: v10, v2: v20, v3: v30}, map[uint16]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Float32PPtr(map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Float32PPtr(map[*uint16]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float32PPtr(nil, map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float32PPtr(map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float32PPtr(nil, map[*uint16]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float32PPtr(map[*uint16]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint16Float64(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapUint16Float64P(map[uint16]float64{v1: v10, v2: v20, v3: v30}, map[uint16]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint16Float64P(map[uint16]float64{v1: v10, v3: v10, v2: v20}, map[uint16]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float64P(nil, map[uint16]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float64P(map[uint16]float64{v1: v10, v2: v20, v3: v30}, map[uint16]float64{})
	if r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float64P(nil, map[uint16]float64{})
	if r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint16Float64P(map[uint16]float64{v1: v10, v2: v20, v3: v30}, map[uint16]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint16Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint16Float64PPtr(map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint16Float64PPtr(map[*uint16]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float64PPtr(nil, map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float64PPtr(map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float64PPtr(nil, map[*uint16]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint16Float64PPtr(map[*uint16]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint16]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint16Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Int(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapUint8IntP(map[uint8]int{v1: v10, v2: v20, v3: v30}, map[uint8]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8IntP(map[uint8]int{v1: v10, v3: v10, v2: v20}, map[uint8]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8IntP(nil, map[uint8]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8IntP(map[uint8]int{v1: v10, v2: v20, v3: v30}, map[uint8]int{})
	if r {
		t.Errorf("EqualMapUint8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8IntP(nil, map[uint8]int{})
	if r {
		t.Errorf("EqualMapUint8IntP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8IntP(map[uint8]int{v1: v10, v2: v20, v3: v30}, map[uint8]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8IntPPtr(map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8IntPPtr(map[*uint8]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8IntPPtr(nil, map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8IntPPtr(map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int{})
	if rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8IntPPtr(nil, map[*uint8]*int{})
	if rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8IntPPtr(map[*uint8]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Int64(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapUint8Int64P(map[uint8]int64{v1: v10, v2: v20, v3: v30}, map[uint8]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Int64P(map[uint8]int64{v1: v10, v3: v10, v2: v20}, map[uint8]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int64P(nil, map[uint8]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int64P(map[uint8]int64{v1: v10, v2: v20, v3: v30}, map[uint8]int64{})
	if r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int64P(nil, map[uint8]int64{})
	if r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int64P(map[uint8]int64{v1: v10, v2: v20, v3: v30}, map[uint8]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Int64PPtr(map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Int64PPtr(map[*uint8]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int64PPtr(nil, map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int64PPtr(map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int64PPtr(nil, map[*uint8]*int64{})
	if rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int64PPtr(map[*uint8]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Int32(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapUint8Int32P(map[uint8]int32{v1: v10, v2: v20, v3: v30}, map[uint8]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Int32P(map[uint8]int32{v1: v10, v3: v10, v2: v20}, map[uint8]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int32P(nil, map[uint8]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int32P(map[uint8]int32{v1: v10, v2: v20, v3: v30}, map[uint8]int32{})
	if r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int32P(nil, map[uint8]int32{})
	if r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int32P(map[uint8]int32{v1: v10, v2: v20, v3: v30}, map[uint8]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Int32PPtr(map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Int32PPtr(map[*uint8]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int32PPtr(nil, map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int32PPtr(map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int32PPtr(nil, map[*uint8]*int32{})
	if rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int32PPtr(map[*uint8]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Int16(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapUint8Int16P(map[uint8]int16{v1: v10, v2: v20, v3: v30}, map[uint8]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Int16P(map[uint8]int16{v1: v10, v3: v10, v2: v20}, map[uint8]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int16P(nil, map[uint8]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int16P(map[uint8]int16{v1: v10, v2: v20, v3: v30}, map[uint8]int16{})
	if r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int16P(nil, map[uint8]int16{})
	if r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int16P(map[uint8]int16{v1: v10, v2: v20, v3: v30}, map[uint8]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Int16PPtr(map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Int16PPtr(map[*uint8]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int16PPtr(nil, map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int16PPtr(map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int16PPtr(nil, map[*uint8]*int16{})
	if rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int16PPtr(map[*uint8]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Int8(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapUint8Int8P(map[uint8]int8{v1: v10, v2: v20, v3: v30}, map[uint8]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Int8P(map[uint8]int8{v1: v10, v3: v10, v2: v20}, map[uint8]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int8P(nil, map[uint8]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int8P(map[uint8]int8{v1: v10, v2: v20, v3: v30}, map[uint8]int8{})
	if r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int8P(nil, map[uint8]int8{})
	if r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Int8P(map[uint8]int8{v1: v10, v2: v20, v3: v30}, map[uint8]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Int8PPtr(map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Int8PPtr(map[*uint8]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int8PPtr(nil, map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int8PPtr(map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int8PPtr(nil, map[*uint8]*int8{})
	if rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Int8PPtr(map[*uint8]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Uint(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapUint8UintP(map[uint8]uint{v1: v10, v2: v20, v3: v30}, map[uint8]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8UintP(map[uint8]uint{v1: v10, v3: v10, v2: v20}, map[uint8]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8UintP(nil, map[uint8]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8UintP(map[uint8]uint{v1: v10, v2: v20, v3: v30}, map[uint8]uint{})
	if r {
		t.Errorf("EqualMapUint8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8UintP(nil, map[uint8]uint{})
	if r {
		t.Errorf("EqualMapUint8UintP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8UintP(map[uint8]uint{v1: v10, v2: v20, v3: v30}, map[uint8]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8UintPPtr(map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8UintPPtr(map[*uint8]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8UintPPtr(nil, map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8UintPPtr(map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8UintPPtr(nil, map[*uint8]*uint{})
	if rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8UintPPtr(map[*uint8]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Uint64(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapUint8Uint64P(map[uint8]uint64{v1: v10, v2: v20, v3: v30}, map[uint8]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Uint64P(map[uint8]uint64{v1: v10, v3: v10, v2: v20}, map[uint8]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint64P(nil, map[uint8]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint64P(map[uint8]uint64{v1: v10, v2: v20, v3: v30}, map[uint8]uint64{})
	if r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint64P(nil, map[uint8]uint64{})
	if r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint64P(map[uint8]uint64{v1: v10, v2: v20, v3: v30}, map[uint8]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Uint64PPtr(map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Uint64PPtr(map[*uint8]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint64PPtr(nil, map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint64PPtr(map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint64PPtr(nil, map[*uint8]*uint64{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint64PPtr(map[*uint8]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Uint32(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapUint8Uint32P(map[uint8]uint32{v1: v10, v2: v20, v3: v30}, map[uint8]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Uint32P(map[uint8]uint32{v1: v10, v3: v10, v2: v20}, map[uint8]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint32P(nil, map[uint8]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint32P(map[uint8]uint32{v1: v10, v2: v20, v3: v30}, map[uint8]uint32{})
	if r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint32P(nil, map[uint8]uint32{})
	if r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint32P(map[uint8]uint32{v1: v10, v2: v20, v3: v30}, map[uint8]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Uint32PPtr(map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Uint32PPtr(map[*uint8]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint32PPtr(nil, map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint32PPtr(map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint32PPtr(nil, map[*uint8]*uint32{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint32PPtr(map[*uint8]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Uint16(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapUint8Uint16P(map[uint8]uint16{v1: v10, v2: v20, v3: v30}, map[uint8]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Uint16P(map[uint8]uint16{v1: v10, v3: v10, v2: v20}, map[uint8]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint16P(nil, map[uint8]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint16P(map[uint8]uint16{v1: v10, v2: v20, v3: v30}, map[uint8]uint16{})
	if r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint16P(nil, map[uint8]uint16{})
	if r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Uint16P(map[uint8]uint16{v1: v10, v2: v20, v3: v30}, map[uint8]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Uint16PPtr(map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Uint16PPtr(map[*uint8]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint16PPtr(nil, map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint16PPtr(map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint16PPtr(nil, map[*uint8]*uint16{})
	if rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Uint16PPtr(map[*uint8]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8(t *testing.T) {
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

	r = EqualMapUint8P(map[uint8]uint8{v1: v10, v3: v10, v2: v20}, map[uint8]uint8{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapUint8PPtr(map[*uint8]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsUint8Str(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapUint8StrP(map[uint8]string{v1: v10, v2: v20, v3: v30}, map[uint8]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8StrP(map[uint8]string{v1: v10, v3: v10, v2: v20}, map[uint8]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8StrP(nil, map[uint8]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8StrP(map[uint8]string{v1: v10, v2: v20, v3: v30}, map[uint8]string{})
	if r {
		t.Errorf("EqualMapUint8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8StrP(nil, map[uint8]string{})
	if r {
		t.Errorf("EqualMapUint8StrP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8StrP(map[uint8]string{v1: v10, v2: v20, v3: v30}, map[uint8]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8StrPPtr(map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8StrPPtr(map[*uint8]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8StrPPtr(nil, map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8StrPPtr(map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*string{})
	if rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8StrPPtr(nil, map[*uint8]*string{})
	if rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8StrPPtr(map[*uint8]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Bool(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapUint8BoolP(map[uint8]bool{v1: v10, v2: v20, v3: v30}, map[uint8]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8BoolP(map[uint8]bool{v1: v20, v3: v10, v2: v20}, map[uint8]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8BoolP(nil, map[uint8]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8BoolP(map[uint8]bool{v1: v10, v2: v20, v3: v30}, map[uint8]bool{})
	if r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8BoolP(nil, map[uint8]bool{})
	if r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapUint8BoolP(map[uint8]bool{v1: v10, v2: v20, v3: v30}, map[uint8]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8BoolPPtr(map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8BoolPPtr(map[*uint8]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8BoolPPtr(nil, map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8BoolPPtr(map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8BoolPPtr(nil, map[*uint8]*bool{})
	if rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8BoolPPtr(map[*uint8]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Float32(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapUint8Float32P(map[uint8]float32{v1: v10, v2: v20, v3: v30}, map[uint8]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Float32P(map[uint8]float32{v1: v10, v3: v10, v2: v20}, map[uint8]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float32P(nil, map[uint8]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float32P(map[uint8]float32{v1: v10, v2: v20, v3: v30}, map[uint8]float32{})
	if r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float32P(nil, map[uint8]float32{})
	if r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float32P(map[uint8]float32{v1: v10, v2: v20, v3: v30}, map[uint8]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Float32PPtr(map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Float32PPtr(map[*uint8]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float32PPtr(nil, map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float32PPtr(map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float32PPtr(nil, map[*uint8]*float32{})
	if rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float32PPtr(map[*uint8]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsUint8Float64(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapUint8Float64P(map[uint8]float64{v1: v10, v2: v20, v3: v30}, map[uint8]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapUint8Float64P(map[uint8]float64{v1: v10, v3: v10, v2: v20}, map[uint8]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float64P(nil, map[uint8]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float64P(map[uint8]float64{v1: v10, v2: v20, v3: v30}, map[uint8]float64{})
	if r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float64P(nil, map[uint8]float64{})
	if r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapUint8Float64P(map[uint8]float64{v1: v10, v2: v20, v3: v30}, map[uint8]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapUint8Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapUint8Float64PPtr(map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapUint8Float64PPtr(map[*uint8]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float64PPtr(nil, map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float64PPtr(map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float64PPtr(nil, map[*uint8]*float64{})
	if rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapUint8Float64PPtr(map[*uint8]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*uint8]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapUint8Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrInt(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapStrIntP(map[string]int{v1: v10, v2: v20, v3: v30}, map[string]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrIntP(map[string]int{v1: v10, v3: v10, v2: v20}, map[string]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrIntP failed. Expected=false, actual=true")
	}

	r = EqualMapStrIntP(nil, map[string]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrIntP failed. Expected=false, actual=true")
	}

	r = EqualMapStrIntP(map[string]int{v1: v10, v2: v20, v3: v30}, map[string]int{})
	if r {
		t.Errorf("EqualMapStrIntP failed. Expected=false, actual=true")
	}

	r = EqualMapStrIntP(nil, map[string]int{})
	if r {
		t.Errorf("EqualMapStrIntP failed. Expected=false, actual=true")
	}

	r = EqualMapStrIntP(map[string]int{v1: v10, v2: v20, v3: v30}, map[string]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrIntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrIntPPtr(map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrIntPPtr(map[*string]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrIntPPtr(nil, map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrIntPPtr(map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int{})
	if rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrIntPPtr(nil, map[*string]*int{})
	if rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrIntPPtr(map[*string]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrIntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrInt64(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapStrInt64P(map[string]int64{v1: v10, v2: v20, v3: v30}, map[string]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrInt64P(map[string]int64{v1: v10, v3: v10, v2: v20}, map[string]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt64P(nil, map[string]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt64P(map[string]int64{v1: v10, v2: v20, v3: v30}, map[string]int64{})
	if r {
		t.Errorf("EqualMapStrInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt64P(nil, map[string]int64{})
	if r {
		t.Errorf("EqualMapStrInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt64P(map[string]int64{v1: v10, v2: v20, v3: v30}, map[string]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrInt64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrInt64PPtr(map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrInt64PPtr(map[*string]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt64PPtr(nil, map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt64PPtr(map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int64{})
	if rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt64PPtr(nil, map[*string]*int64{})
	if rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt64PPtr(map[*string]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrInt32(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapStrInt32P(map[string]int32{v1: v10, v2: v20, v3: v30}, map[string]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrInt32P(map[string]int32{v1: v10, v3: v10, v2: v20}, map[string]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt32P(nil, map[string]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt32P(map[string]int32{v1: v10, v2: v20, v3: v30}, map[string]int32{})
	if r {
		t.Errorf("EqualMapStrInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt32P(nil, map[string]int32{})
	if r {
		t.Errorf("EqualMapStrInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt32P(map[string]int32{v1: v10, v2: v20, v3: v30}, map[string]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrInt32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrInt32PPtr(map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrInt32PPtr(map[*string]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt32PPtr(nil, map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt32PPtr(map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int32{})
	if rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt32PPtr(nil, map[*string]*int32{})
	if rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt32PPtr(map[*string]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrInt16(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapStrInt16P(map[string]int16{v1: v10, v2: v20, v3: v30}, map[string]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrInt16P(map[string]int16{v1: v10, v3: v10, v2: v20}, map[string]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt16P(nil, map[string]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt16P(map[string]int16{v1: v10, v2: v20, v3: v30}, map[string]int16{})
	if r {
		t.Errorf("EqualMapStrInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt16P(nil, map[string]int16{})
	if r {
		t.Errorf("EqualMapStrInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt16P(map[string]int16{v1: v10, v2: v20, v3: v30}, map[string]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrInt16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrInt16PPtr(map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrInt16PPtr(map[*string]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt16PPtr(nil, map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt16PPtr(map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int16{})
	if rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt16PPtr(nil, map[*string]*int16{})
	if rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt16PPtr(map[*string]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrInt8(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapStrInt8P(map[string]int8{v1: v10, v2: v20, v3: v30}, map[string]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrInt8P(map[string]int8{v1: v10, v3: v10, v2: v20}, map[string]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt8P(nil, map[string]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt8P(map[string]int8{v1: v10, v2: v20, v3: v30}, map[string]int8{})
	if r {
		t.Errorf("EqualMapStrInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt8P(nil, map[string]int8{})
	if r {
		t.Errorf("EqualMapStrInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrInt8P(map[string]int8{v1: v10, v2: v20, v3: v30}, map[string]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrInt8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrInt8PPtr(map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrInt8PPtr(map[*string]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt8PPtr(nil, map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt8PPtr(map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int8{})
	if rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt8PPtr(nil, map[*string]*int8{})
	if rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrInt8PPtr(map[*string]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrUint(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapStrUintP(map[string]uint{v1: v10, v2: v20, v3: v30}, map[string]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrUintP(map[string]uint{v1: v10, v3: v10, v2: v20}, map[string]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUintP failed. Expected=false, actual=true")
	}

	r = EqualMapStrUintP(nil, map[string]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUintP failed. Expected=false, actual=true")
	}

	r = EqualMapStrUintP(map[string]uint{v1: v10, v2: v20, v3: v30}, map[string]uint{})
	if r {
		t.Errorf("EqualMapStrUintP failed. Expected=false, actual=true")
	}

	r = EqualMapStrUintP(nil, map[string]uint{})
	if r {
		t.Errorf("EqualMapStrUintP failed. Expected=false, actual=true")
	}

	r = EqualMapStrUintP(map[string]uint{v1: v10, v2: v20, v3: v30}, map[string]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrUintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrUintPPtr(map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrUintPPtr(map[*string]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUintPPtr(nil, map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUintPPtr(map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint{})
	if rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUintPPtr(nil, map[*string]*uint{})
	if rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUintPPtr(map[*string]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrUintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrUint64(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapStrUint64P(map[string]uint64{v1: v10, v2: v20, v3: v30}, map[string]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrUint64P(map[string]uint64{v1: v10, v3: v10, v2: v20}, map[string]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint64P(nil, map[string]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint64P(map[string]uint64{v1: v10, v2: v20, v3: v30}, map[string]uint64{})
	if r {
		t.Errorf("EqualMapStrUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint64P(nil, map[string]uint64{})
	if r {
		t.Errorf("EqualMapStrUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint64P(map[string]uint64{v1: v10, v2: v20, v3: v30}, map[string]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrUint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrUint64PPtr(map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrUint64PPtr(map[*string]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint64PPtr(nil, map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint64PPtr(map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint64{})
	if rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint64PPtr(nil, map[*string]*uint64{})
	if rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint64PPtr(map[*string]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrUint32(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapStrUint32P(map[string]uint32{v1: v10, v2: v20, v3: v30}, map[string]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrUint32P(map[string]uint32{v1: v10, v3: v10, v2: v20}, map[string]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint32P(nil, map[string]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint32P(map[string]uint32{v1: v10, v2: v20, v3: v30}, map[string]uint32{})
	if r {
		t.Errorf("EqualMapStrUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint32P(nil, map[string]uint32{})
	if r {
		t.Errorf("EqualMapStrUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint32P(map[string]uint32{v1: v10, v2: v20, v3: v30}, map[string]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrUint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrUint32PPtr(map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrUint32PPtr(map[*string]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint32PPtr(nil, map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint32PPtr(map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint32{})
	if rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint32PPtr(nil, map[*string]*uint32{})
	if rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint32PPtr(map[*string]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrUint16(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapStrUint16P(map[string]uint16{v1: v10, v2: v20, v3: v30}, map[string]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrUint16P(map[string]uint16{v1: v10, v3: v10, v2: v20}, map[string]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint16P(nil, map[string]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint16P(map[string]uint16{v1: v10, v2: v20, v3: v30}, map[string]uint16{})
	if r {
		t.Errorf("EqualMapStrUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint16P(nil, map[string]uint16{})
	if r {
		t.Errorf("EqualMapStrUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint16P(map[string]uint16{v1: v10, v2: v20, v3: v30}, map[string]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrUint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrUint16PPtr(map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrUint16PPtr(map[*string]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint16PPtr(nil, map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint16PPtr(map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint16{})
	if rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint16PPtr(nil, map[*string]*uint16{})
	if rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint16PPtr(map[*string]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrUint8(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapStrUint8P(map[string]uint8{v1: v10, v2: v20, v3: v30}, map[string]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrUint8P(map[string]uint8{v1: v10, v3: v10, v2: v20}, map[string]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint8P(nil, map[string]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint8P(map[string]uint8{v1: v10, v2: v20, v3: v30}, map[string]uint8{})
	if r {
		t.Errorf("EqualMapStrUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint8P(nil, map[string]uint8{})
	if r {
		t.Errorf("EqualMapStrUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapStrUint8P(map[string]uint8{v1: v10, v2: v20, v3: v30}, map[string]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrUint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrUint8PPtr(map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrUint8PPtr(map[*string]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint8PPtr(nil, map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint8PPtr(map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint8{})
	if rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint8PPtr(nil, map[*string]*uint8{})
	if rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrUint8PPtr(map[*string]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStr(t *testing.T) {
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

	r = EqualMapStrP(map[string]string{v1: v10, v3: v10, v2: v20}, map[string]string{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapStrPPtr(map[*string]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*string{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsStrBool(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapStrBoolP(map[string]bool{v1: v10, v2: v20, v3: v30}, map[string]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrBoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrBoolP(map[string]bool{v1: v20, v3: v10, v2: v20}, map[string]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapStrBoolP(nil, map[string]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapStrBoolP(map[string]bool{v1: v10, v2: v20, v3: v30}, map[string]bool{})
	if r {
		t.Errorf("EqualMapStrBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapStrBoolP(nil, map[string]bool{})
	if r {
		t.Errorf("EqualMapStrBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapStrBoolP(map[string]bool{v1: v10, v2: v20, v3: v30}, map[string]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrBoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrBoolPPtr(map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrBoolPPtr(map[*string]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrBoolPPtr(nil, map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrBoolPPtr(map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*bool{})
	if rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrBoolPPtr(nil, map[*string]*bool{})
	if rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrBoolPPtr(map[*string]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrBoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrFloat32(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapStrFloat32P(map[string]float32{v1: v10, v2: v20, v3: v30}, map[string]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrFloat32P(map[string]float32{v1: v10, v3: v10, v2: v20}, map[string]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat32P(nil, map[string]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat32P(map[string]float32{v1: v10, v2: v20, v3: v30}, map[string]float32{})
	if r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat32P(nil, map[string]float32{})
	if r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat32P(map[string]float32{v1: v10, v2: v20, v3: v30}, map[string]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrFloat32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrFloat32PPtr(map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrFloat32PPtr(map[*string]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat32PPtr(nil, map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat32PPtr(map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float32{})
	if rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat32PPtr(nil, map[*string]*float32{})
	if rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat32PPtr(map[*string]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsStrFloat64(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapStrFloat64P(map[string]float64{v1: v10, v2: v20, v3: v30}, map[string]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapStrFloat64P(map[string]float64{v1: v10, v3: v10, v2: v20}, map[string]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat64P(nil, map[string]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat64P(map[string]float64{v1: v10, v2: v20, v3: v30}, map[string]float64{})
	if r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat64P(nil, map[string]float64{})
	if r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapStrFloat64P(map[string]float64{v1: v10, v2: v20, v3: v30}, map[string]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapStrFloat64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapStrFloat64PPtr(map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapStrFloat64PPtr(map[*string]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat64PPtr(nil, map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat64PPtr(map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float64{})
	if rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat64PPtr(nil, map[*string]*float64{})
	if rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapStrFloat64PPtr(map[*string]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*string]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapStrFloat64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolInt(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapBoolIntP(map[bool]int{v1: v10, v2: v20}, map[bool]int{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolIntP(map[bool]int{v1: v20, v2: v10}, map[bool]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolIntP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolIntP(nil, map[bool]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolIntP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolIntP(map[bool]int{v1: v10, v2: v20}, map[bool]int{})
	if r {
		t.Errorf("EqualMapBoolIntP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolIntP(nil, map[bool]int{})
	if r {
		t.Errorf("EqualMapBoolIntP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolIntP(map[bool]int{v1: v10, v2: v20}, map[bool]int{v1: v10})
	if r {
		t.Errorf("EqualMapBoolIntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolIntPPtr(map[*bool]*int{&v1: &v10, &v2: &v20}, map[*bool]*int{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolIntPPtr(map[*bool]*int{&v1: &v20, &v2: &v20}, map[*bool]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolIntPPtr(nil, map[*bool]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolIntPPtr(map[*bool]*int{&v1: &v10, &v2: &v20}, map[*bool]*int{})
	if rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolIntPPtr(nil, map[*bool]*int{})
	if rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolIntPPtr(map[*bool]*int{&v1: &v10, &v2: &v20}, map[*bool]*int{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolIntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolInt64(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapBoolInt64P(map[bool]int64{v1: v10, v2: v20}, map[bool]int64{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolInt64P(map[bool]int64{v1: v20, v2: v10}, map[bool]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt64P(nil, map[bool]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt64P(map[bool]int64{v1: v10, v2: v20}, map[bool]int64{})
	if r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt64P(nil, map[bool]int64{})
	if r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt64P(map[bool]int64{v1: v10, v2: v20}, map[bool]int64{v1: v10})
	if r {
		t.Errorf("EqualMapBoolInt64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolInt64PPtr(map[*bool]*int64{&v1: &v10, &v2: &v20}, map[*bool]*int64{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolInt64PPtr(map[*bool]*int64{&v1: &v20, &v2: &v20}, map[*bool]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt64PPtr(nil, map[*bool]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt64PPtr(map[*bool]*int64{&v1: &v10, &v2: &v20}, map[*bool]*int64{})
	if rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt64PPtr(nil, map[*bool]*int64{})
	if rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt64PPtr(map[*bool]*int64{&v1: &v10, &v2: &v20}, map[*bool]*int64{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolInt32(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapBoolInt32P(map[bool]int32{v1: v10, v2: v20}, map[bool]int32{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolInt32P(map[bool]int32{v1: v20, v2: v10}, map[bool]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt32P(nil, map[bool]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt32P(map[bool]int32{v1: v10, v2: v20}, map[bool]int32{})
	if r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt32P(nil, map[bool]int32{})
	if r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt32P(map[bool]int32{v1: v10, v2: v20}, map[bool]int32{v1: v10})
	if r {
		t.Errorf("EqualMapBoolInt32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolInt32PPtr(map[*bool]*int32{&v1: &v10, &v2: &v20}, map[*bool]*int32{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolInt32PPtr(map[*bool]*int32{&v1: &v20, &v2: &v20}, map[*bool]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt32PPtr(nil, map[*bool]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt32PPtr(map[*bool]*int32{&v1: &v10, &v2: &v20}, map[*bool]*int32{})
	if rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt32PPtr(nil, map[*bool]*int32{})
	if rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt32PPtr(map[*bool]*int32{&v1: &v10, &v2: &v20}, map[*bool]*int32{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolInt16(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapBoolInt16P(map[bool]int16{v1: v10, v2: v20}, map[bool]int16{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolInt16P(map[bool]int16{v1: v20, v2: v10}, map[bool]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt16P(nil, map[bool]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt16P(map[bool]int16{v1: v10, v2: v20}, map[bool]int16{})
	if r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt16P(nil, map[bool]int16{})
	if r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt16P(map[bool]int16{v1: v10, v2: v20}, map[bool]int16{v1: v10})
	if r {
		t.Errorf("EqualMapBoolInt16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolInt16PPtr(map[*bool]*int16{&v1: &v10, &v2: &v20}, map[*bool]*int16{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolInt16PPtr(map[*bool]*int16{&v1: &v20, &v2: &v20}, map[*bool]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt16PPtr(nil, map[*bool]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt16PPtr(map[*bool]*int16{&v1: &v10, &v2: &v20}, map[*bool]*int16{})
	if rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt16PPtr(nil, map[*bool]*int16{})
	if rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt16PPtr(map[*bool]*int16{&v1: &v10, &v2: &v20}, map[*bool]*int16{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolInt8(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapBoolInt8P(map[bool]int8{v1: v10, v2: v20}, map[bool]int8{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolInt8P(map[bool]int8{v1: v20, v2: v10}, map[bool]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt8P(nil, map[bool]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt8P(map[bool]int8{v1: v10, v2: v20}, map[bool]int8{})
	if r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt8P(nil, map[bool]int8{})
	if r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolInt8P(map[bool]int8{v1: v10, v2: v20}, map[bool]int8{v1: v10})
	if r {
		t.Errorf("EqualMapBoolInt8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolInt8PPtr(map[*bool]*int8{&v1: &v10, &v2: &v20}, map[*bool]*int8{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolInt8PPtr(map[*bool]*int8{&v1: &v20, &v2: &v20}, map[*bool]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt8PPtr(nil, map[*bool]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt8PPtr(map[*bool]*int8{&v1: &v10, &v2: &v20}, map[*bool]*int8{})
	if rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt8PPtr(nil, map[*bool]*int8{})
	if rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolInt8PPtr(map[*bool]*int8{&v1: &v10, &v2: &v20}, map[*bool]*int8{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolUint(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapBoolUintP(map[bool]uint{v1: v10, v2: v20}, map[bool]uint{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolUintP(map[bool]uint{v1: v20, v2: v10}, map[bool]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUintP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUintP(nil, map[bool]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUintP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUintP(map[bool]uint{v1: v10, v2: v20}, map[bool]uint{})
	if r {
		t.Errorf("EqualMapBoolUintP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUintP(nil, map[bool]uint{})
	if r {
		t.Errorf("EqualMapBoolUintP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUintP(map[bool]uint{v1: v10, v2: v20}, map[bool]uint{v1: v10})
	if r {
		t.Errorf("EqualMapBoolUintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolUintPPtr(map[*bool]*uint{&v1: &v10, &v2: &v20}, map[*bool]*uint{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolUintPPtr(map[*bool]*uint{&v1: &v20, &v2: &v20}, map[*bool]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUintPPtr(nil, map[*bool]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUintPPtr(map[*bool]*uint{&v1: &v10, &v2: &v20}, map[*bool]*uint{})
	if rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUintPPtr(nil, map[*bool]*uint{})
	if rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUintPPtr(map[*bool]*uint{&v1: &v10, &v2: &v20}, map[*bool]*uint{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolUintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolUint64(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapBoolUint64P(map[bool]uint64{v1: v10, v2: v20}, map[bool]uint64{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolUint64P(map[bool]uint64{v1: v20, v2: v10}, map[bool]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint64P(nil, map[bool]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint64P(map[bool]uint64{v1: v10, v2: v20}, map[bool]uint64{})
	if r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint64P(nil, map[bool]uint64{})
	if r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint64P(map[bool]uint64{v1: v10, v2: v20}, map[bool]uint64{v1: v10})
	if r {
		t.Errorf("EqualMapBoolUint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolUint64PPtr(map[*bool]*uint64{&v1: &v10, &v2: &v20}, map[*bool]*uint64{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolUint64PPtr(map[*bool]*uint64{&v1: &v20, &v2: &v20}, map[*bool]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint64PPtr(nil, map[*bool]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint64PPtr(map[*bool]*uint64{&v1: &v10, &v2: &v20}, map[*bool]*uint64{})
	if rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint64PPtr(nil, map[*bool]*uint64{})
	if rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint64PPtr(map[*bool]*uint64{&v1: &v10, &v2: &v20}, map[*bool]*uint64{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolUint32(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapBoolUint32P(map[bool]uint32{v1: v10, v2: v20}, map[bool]uint32{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolUint32P(map[bool]uint32{v1: v20, v2: v10}, map[bool]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint32P(nil, map[bool]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint32P(map[bool]uint32{v1: v10, v2: v20}, map[bool]uint32{})
	if r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint32P(nil, map[bool]uint32{})
	if r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint32P(map[bool]uint32{v1: v10, v2: v20}, map[bool]uint32{v1: v10})
	if r {
		t.Errorf("EqualMapBoolUint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolUint32PPtr(map[*bool]*uint32{&v1: &v10, &v2: &v20}, map[*bool]*uint32{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolUint32PPtr(map[*bool]*uint32{&v1: &v20, &v2: &v20}, map[*bool]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint32PPtr(nil, map[*bool]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint32PPtr(map[*bool]*uint32{&v1: &v10, &v2: &v20}, map[*bool]*uint32{})
	if rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint32PPtr(nil, map[*bool]*uint32{})
	if rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint32PPtr(map[*bool]*uint32{&v1: &v10, &v2: &v20}, map[*bool]*uint32{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolUint16(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapBoolUint16P(map[bool]uint16{v1: v10, v2: v20}, map[bool]uint16{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolUint16P(map[bool]uint16{v1: v20, v2: v10}, map[bool]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint16P(nil, map[bool]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint16P(map[bool]uint16{v1: v10, v2: v20}, map[bool]uint16{})
	if r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint16P(nil, map[bool]uint16{})
	if r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint16P(map[bool]uint16{v1: v10, v2: v20}, map[bool]uint16{v1: v10})
	if r {
		t.Errorf("EqualMapBoolUint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolUint16PPtr(map[*bool]*uint16{&v1: &v10, &v2: &v20}, map[*bool]*uint16{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolUint16PPtr(map[*bool]*uint16{&v1: &v20, &v2: &v20}, map[*bool]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint16PPtr(nil, map[*bool]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint16PPtr(map[*bool]*uint16{&v1: &v10, &v2: &v20}, map[*bool]*uint16{})
	if rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint16PPtr(nil, map[*bool]*uint16{})
	if rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint16PPtr(map[*bool]*uint16{&v1: &v10, &v2: &v20}, map[*bool]*uint16{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolUint8(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapBoolUint8P(map[bool]uint8{v1: v10, v2: v20}, map[bool]uint8{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolUint8P(map[bool]uint8{v1: v20, v2: v10}, map[bool]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint8P(nil, map[bool]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint8P(map[bool]uint8{v1: v10, v2: v20}, map[bool]uint8{})
	if r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint8P(nil, map[bool]uint8{})
	if r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolUint8P(map[bool]uint8{v1: v10, v2: v20}, map[bool]uint8{v1: v10})
	if r {
		t.Errorf("EqualMapBoolUint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolUint8PPtr(map[*bool]*uint8{&v1: &v10, &v2: &v20}, map[*bool]*uint8{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolUint8PPtr(map[*bool]*uint8{&v1: &v20, &v2: &v20}, map[*bool]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint8PPtr(nil, map[*bool]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint8PPtr(map[*bool]*uint8{&v1: &v10, &v2: &v20}, map[*bool]*uint8{})
	if rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint8PPtr(nil, map[*bool]*uint8{})
	if rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolUint8PPtr(map[*bool]*uint8{&v1: &v10, &v2: &v20}, map[*bool]*uint8{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolStr(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapBoolStrP(map[bool]string{v1: v10, v2: v20}, map[bool]string{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolStrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolStrP(map[bool]string{v1: v20, v2: v10}, map[bool]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolStrP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolStrP(nil, map[bool]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolStrP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolStrP(map[bool]string{v1: v10, v2: v20}, map[bool]string{})
	if r {
		t.Errorf("EqualMapBoolStrP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolStrP(nil, map[bool]string{})
	if r {
		t.Errorf("EqualMapBoolStrP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolStrP(map[bool]string{v1: v10, v2: v20}, map[bool]string{v1: v10})
	if r {
		t.Errorf("EqualMapBoolStrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolStrPPtr(map[*bool]*string{&v1: &v10, &v2: &v20}, map[*bool]*string{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolStrPPtr(map[*bool]*string{&v1: &v20, &v2: &v20}, map[*bool]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolStrPPtr(nil, map[*bool]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolStrPPtr(map[*bool]*string{&v1: &v10, &v2: &v20}, map[*bool]*string{})
	if rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolStrPPtr(nil, map[*bool]*string{})
	if rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolStrPPtr(map[*bool]*string{&v1: &v10, &v2: &v20}, map[*bool]*string{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolStrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBools(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 bool = true
	var v20 bool = false

	r := EqualMapBoolP(map[bool]bool{v1: v10, v2: v20}, map[bool]bool{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolP(map[bool]bool{v1: v20, v2: v10}, map[bool]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(nil, map[bool]bool{v1: v20, v2: v20})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v10, v2: v20}, map[bool]bool{})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(nil, map[bool]bool{})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v10, v2: v20}, map[bool]bool{v1: v10})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	r = EqualMapBoolP(map[bool]bool{v1: v20}, map[bool]bool{v1: v10})
	if r {
		t.Errorf("EqualMapBoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolPPtr(map[*bool]*bool{&v1: &v10, &v2: &v20}, map[*bool]*bool{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v10, &v2: &v20}, map[*bool]*bool{&v1: &v20, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(nil, map[*bool]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v10, &v2: &v20}, map[*bool]*bool{})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(nil, map[*bool]*bool{})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v10}, map[*bool]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolPPtr(map[*bool]*bool{&v1: &v20}, map[*bool]*bool{&v2: &v10})
	if rPtr {
		t.Errorf("EqualMapBoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolFloat32(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapBoolFloat32P(map[bool]float32{v1: v10, v2: v20}, map[bool]float32{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolFloat32P(map[bool]float32{v1: v20, v2: v10}, map[bool]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat32P(nil, map[bool]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat32P(map[bool]float32{v1: v10, v2: v20}, map[bool]float32{})
	if r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat32P(nil, map[bool]float32{})
	if r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat32P(map[bool]float32{v1: v10, v2: v20}, map[bool]float32{v1: v10})
	if r {
		t.Errorf("EqualMapBoolFloat32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolFloat32PPtr(map[*bool]*float32{&v1: &v10, &v2: &v20}, map[*bool]*float32{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolFloat32PPtr(map[*bool]*float32{&v1: &v20, &v2: &v20}, map[*bool]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat32PPtr(nil, map[*bool]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat32PPtr(map[*bool]*float32{&v1: &v10, &v2: &v20}, map[*bool]*float32{})
	if rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat32PPtr(nil, map[*bool]*float32{})
	if rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat32PPtr(map[*bool]*float32{&v1: &v10, &v2: &v20}, map[*bool]*float32{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsBoolFloat64(t *testing.T) {
	var v1 bool = true
	var v2 bool = false

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapBoolFloat64P(map[bool]float64{v1: v10, v2: v20}, map[bool]float64{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapBoolFloat64P(map[bool]float64{v1: v20, v2: v10}, map[bool]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat64P(nil, map[bool]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat64P(map[bool]float64{v1: v10, v2: v20}, map[bool]float64{})
	if r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat64P(nil, map[bool]float64{})
	if r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=false, actual=true")
	}

	r = EqualMapBoolFloat64P(map[bool]float64{v1: v10, v2: v20}, map[bool]float64{v1: v10})
	if r {
		t.Errorf("EqualMapBoolFloat64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapBoolFloat64PPtr(map[*bool]*float64{&v1: &v10, &v2: &v20}, map[*bool]*float64{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapBoolFloat64PPtr(map[*bool]*float64{&v1: &v20, &v2: &v20}, map[*bool]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat64PPtr(nil, map[*bool]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat64PPtr(map[*bool]*float64{&v1: &v10, &v2: &v20}, map[*bool]*float64{})
	if rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat64PPtr(nil, map[*bool]*float64{})
	if rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapBoolFloat64PPtr(map[*bool]*float64{&v1: &v10, &v2: &v20}, map[*bool]*float64{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMapBoolFloat64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Int(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapFloat32IntP(map[float32]int{v1: v10, v2: v20, v3: v30}, map[float32]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32IntP(map[float32]int{v1: v10, v3: v10, v2: v20}, map[float32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32IntP(nil, map[float32]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32IntP(map[float32]int{v1: v10, v2: v20, v3: v30}, map[float32]int{})
	if r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32IntP(nil, map[float32]int{})
	if r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32IntP(map[float32]int{v1: v10, v2: v20, v3: v30}, map[float32]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32IntPPtr(map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32IntPPtr(map[*float32]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32IntPPtr(nil, map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32IntPPtr(map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int{})
	if rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32IntPPtr(nil, map[*float32]*int{})
	if rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32IntPPtr(map[*float32]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Int64(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapFloat32Int64P(map[float32]int64{v1: v10, v2: v20, v3: v30}, map[float32]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Int64P(map[float32]int64{v1: v10, v3: v10, v2: v20}, map[float32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int64P(nil, map[float32]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int64P(map[float32]int64{v1: v10, v2: v20, v3: v30}, map[float32]int64{})
	if r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int64P(nil, map[float32]int64{})
	if r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int64P(map[float32]int64{v1: v10, v2: v20, v3: v30}, map[float32]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Int64PPtr(map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Int64PPtr(map[*float32]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int64PPtr(nil, map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int64PPtr(map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int64PPtr(nil, map[*float32]*int64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int64PPtr(map[*float32]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Int32(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapFloat32Int32P(map[float32]int32{v1: v10, v2: v20, v3: v30}, map[float32]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Int32P(map[float32]int32{v1: v10, v3: v10, v2: v20}, map[float32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int32P(nil, map[float32]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int32P(map[float32]int32{v1: v10, v2: v20, v3: v30}, map[float32]int32{})
	if r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int32P(nil, map[float32]int32{})
	if r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int32P(map[float32]int32{v1: v10, v2: v20, v3: v30}, map[float32]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Int32PPtr(map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Int32PPtr(map[*float32]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int32PPtr(nil, map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int32PPtr(map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int32{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int32PPtr(nil, map[*float32]*int32{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int32PPtr(map[*float32]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Int16(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapFloat32Int16P(map[float32]int16{v1: v10, v2: v20, v3: v30}, map[float32]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Int16P(map[float32]int16{v1: v10, v3: v10, v2: v20}, map[float32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int16P(nil, map[float32]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int16P(map[float32]int16{v1: v10, v2: v20, v3: v30}, map[float32]int16{})
	if r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int16P(nil, map[float32]int16{})
	if r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int16P(map[float32]int16{v1: v10, v2: v20, v3: v30}, map[float32]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Int16PPtr(map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Int16PPtr(map[*float32]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int16PPtr(nil, map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int16PPtr(map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int16{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int16PPtr(nil, map[*float32]*int16{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int16PPtr(map[*float32]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Int8(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapFloat32Int8P(map[float32]int8{v1: v10, v2: v20, v3: v30}, map[float32]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Int8P(map[float32]int8{v1: v10, v3: v10, v2: v20}, map[float32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int8P(nil, map[float32]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int8P(map[float32]int8{v1: v10, v2: v20, v3: v30}, map[float32]int8{})
	if r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int8P(nil, map[float32]int8{})
	if r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Int8P(map[float32]int8{v1: v10, v2: v20, v3: v30}, map[float32]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Int8PPtr(map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Int8PPtr(map[*float32]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int8PPtr(nil, map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int8PPtr(map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int8{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int8PPtr(nil, map[*float32]*int8{})
	if rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Int8PPtr(map[*float32]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Uint(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapFloat32UintP(map[float32]uint{v1: v10, v2: v20, v3: v30}, map[float32]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32UintP(map[float32]uint{v1: v10, v3: v10, v2: v20}, map[float32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32UintP(nil, map[float32]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32UintP(map[float32]uint{v1: v10, v2: v20, v3: v30}, map[float32]uint{})
	if r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32UintP(nil, map[float32]uint{})
	if r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32UintP(map[float32]uint{v1: v10, v2: v20, v3: v30}, map[float32]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32UintPPtr(map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32UintPPtr(map[*float32]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32UintPPtr(nil, map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32UintPPtr(map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint{})
	if rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32UintPPtr(nil, map[*float32]*uint{})
	if rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32UintPPtr(map[*float32]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Uint64(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapFloat32Uint64P(map[float32]uint64{v1: v10, v2: v20, v3: v30}, map[float32]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Uint64P(map[float32]uint64{v1: v10, v3: v10, v2: v20}, map[float32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint64P(nil, map[float32]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint64P(map[float32]uint64{v1: v10, v2: v20, v3: v30}, map[float32]uint64{})
	if r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint64P(nil, map[float32]uint64{})
	if r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint64P(map[float32]uint64{v1: v10, v2: v20, v3: v30}, map[float32]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Uint64PPtr(map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Uint64PPtr(map[*float32]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint64PPtr(nil, map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint64PPtr(map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint64PPtr(nil, map[*float32]*uint64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint64PPtr(map[*float32]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Uint32(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapFloat32Uint32P(map[float32]uint32{v1: v10, v2: v20, v3: v30}, map[float32]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Uint32P(map[float32]uint32{v1: v10, v3: v10, v2: v20}, map[float32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint32P(nil, map[float32]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint32P(map[float32]uint32{v1: v10, v2: v20, v3: v30}, map[float32]uint32{})
	if r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint32P(nil, map[float32]uint32{})
	if r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint32P(map[float32]uint32{v1: v10, v2: v20, v3: v30}, map[float32]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Uint32PPtr(map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Uint32PPtr(map[*float32]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint32PPtr(nil, map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint32PPtr(map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint32PPtr(nil, map[*float32]*uint32{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint32PPtr(map[*float32]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Uint16(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapFloat32Uint16P(map[float32]uint16{v1: v10, v2: v20, v3: v30}, map[float32]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Uint16P(map[float32]uint16{v1: v10, v3: v10, v2: v20}, map[float32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint16P(nil, map[float32]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint16P(map[float32]uint16{v1: v10, v2: v20, v3: v30}, map[float32]uint16{})
	if r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint16P(nil, map[float32]uint16{})
	if r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint16P(map[float32]uint16{v1: v10, v2: v20, v3: v30}, map[float32]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Uint16PPtr(map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Uint16PPtr(map[*float32]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint16PPtr(nil, map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint16PPtr(map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint16PPtr(nil, map[*float32]*uint16{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint16PPtr(map[*float32]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Uint8(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapFloat32Uint8P(map[float32]uint8{v1: v10, v2: v20, v3: v30}, map[float32]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Uint8P(map[float32]uint8{v1: v10, v3: v10, v2: v20}, map[float32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint8P(nil, map[float32]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint8P(map[float32]uint8{v1: v10, v2: v20, v3: v30}, map[float32]uint8{})
	if r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint8P(nil, map[float32]uint8{})
	if r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Uint8P(map[float32]uint8{v1: v10, v2: v20, v3: v30}, map[float32]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Uint8PPtr(map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Uint8PPtr(map[*float32]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint8PPtr(nil, map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint8PPtr(map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint8PPtr(nil, map[*float32]*uint8{})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Uint8PPtr(map[*float32]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Str(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapFloat32StrP(map[float32]string{v1: v10, v2: v20, v3: v30}, map[float32]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32StrP(map[float32]string{v1: v10, v3: v10, v2: v20}, map[float32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32StrP(nil, map[float32]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32StrP(map[float32]string{v1: v10, v2: v20, v3: v30}, map[float32]string{})
	if r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32StrP(nil, map[float32]string{})
	if r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32StrP(map[float32]string{v1: v10, v2: v20, v3: v30}, map[float32]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32StrPPtr(map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32StrPPtr(map[*float32]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32StrPPtr(nil, map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32StrPPtr(map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*string{})
	if rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32StrPPtr(nil, map[*float32]*string{})
	if rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32StrPPtr(map[*float32]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32Bool(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapFloat32BoolP(map[float32]bool{v1: v10, v2: v20, v3: v30}, map[float32]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32BoolP(map[float32]bool{v1: v20, v3: v10, v2: v20}, map[float32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32BoolP(nil, map[float32]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32BoolP(map[float32]bool{v1: v10, v2: v20, v3: v30}, map[float32]bool{})
	if r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32BoolP(nil, map[float32]bool{})
	if r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32BoolP(map[float32]bool{v1: v10, v2: v20, v3: v30}, map[float32]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32BoolPPtr(map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32BoolPPtr(map[*float32]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32BoolPPtr(nil, map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32BoolPPtr(map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*bool{})
	if rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32BoolPPtr(nil, map[*float32]*bool{})
	if rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32BoolPPtr(map[*float32]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat32(t *testing.T) {
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

	r = EqualMapFloat32P(map[float32]float32{v1: v10, v3: v10, v2: v20}, map[float32]float32{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapFloat32PPtr(map[*float32]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
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

func TestEqualMapsFloat32Float64(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	var v10 float64 = 10
	var v20 float64 = 20
	var v30 float64 = 30

	r := EqualMapFloat32Float64P(map[float32]float64{v1: v10, v2: v20, v3: v30}, map[float32]float64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat32Float64P(map[float32]float64{v1: v10, v3: v10, v2: v20}, map[float32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Float64P(nil, map[float32]float64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Float64P(map[float32]float64{v1: v10, v2: v20, v3: v30}, map[float32]float64{})
	if r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Float64P(nil, map[float32]float64{})
	if r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat32Float64P(map[float32]float64{v1: v10, v2: v20, v3: v30}, map[float32]float64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat32Float64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat32Float64PPtr(map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat32Float64PPtr(map[*float32]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Float64PPtr(nil, map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Float64PPtr(map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Float64PPtr(nil, map[*float32]*float64{})
	if rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat32Float64PPtr(map[*float32]*float64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float32]*float64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat32Float64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Int(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 int = 10
	var v20 int = 20
	var v30 int = 30

	r := EqualMapFloat64IntP(map[float64]int{v1: v10, v2: v20, v3: v30}, map[float64]int{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64IntP(map[float64]int{v1: v10, v3: v10, v2: v20}, map[float64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64IntP(nil, map[float64]int{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64IntP(map[float64]int{v1: v10, v2: v20, v3: v30}, map[float64]int{})
	if r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64IntP(nil, map[float64]int{})
	if r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64IntP(map[float64]int{v1: v10, v2: v20, v3: v30}, map[float64]int{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64IntP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64IntPPtr(map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64IntPPtr(map[*float64]*int{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64IntPPtr(nil, map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64IntPPtr(map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int{})
	if rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64IntPPtr(nil, map[*float64]*int{})
	if rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64IntPPtr(map[*float64]*int{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64IntPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Int64(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 int64 = 10
	var v20 int64 = 20
	var v30 int64 = 30

	r := EqualMapFloat64Int64P(map[float64]int64{v1: v10, v2: v20, v3: v30}, map[float64]int64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Int64P(map[float64]int64{v1: v10, v3: v10, v2: v20}, map[float64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int64P(nil, map[float64]int64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int64P(map[float64]int64{v1: v10, v2: v20, v3: v30}, map[float64]int64{})
	if r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int64P(nil, map[float64]int64{})
	if r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int64P(map[float64]int64{v1: v10, v2: v20, v3: v30}, map[float64]int64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Int64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Int64PPtr(map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Int64PPtr(map[*float64]*int64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int64PPtr(nil, map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int64PPtr(map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int64{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int64PPtr(nil, map[*float64]*int64{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int64PPtr(map[*float64]*int64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Int64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Int32(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 int32 = 10
	var v20 int32 = 20
	var v30 int32 = 30

	r := EqualMapFloat64Int32P(map[float64]int32{v1: v10, v2: v20, v3: v30}, map[float64]int32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Int32P(map[float64]int32{v1: v10, v3: v10, v2: v20}, map[float64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int32P(nil, map[float64]int32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int32P(map[float64]int32{v1: v10, v2: v20, v3: v30}, map[float64]int32{})
	if r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int32P(nil, map[float64]int32{})
	if r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int32P(map[float64]int32{v1: v10, v2: v20, v3: v30}, map[float64]int32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Int32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Int32PPtr(map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Int32PPtr(map[*float64]*int32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int32PPtr(nil, map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int32PPtr(map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int32PPtr(nil, map[*float64]*int32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int32PPtr(map[*float64]*int32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Int32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Int16(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 int16 = 10
	var v20 int16 = 20
	var v30 int16 = 30

	r := EqualMapFloat64Int16P(map[float64]int16{v1: v10, v2: v20, v3: v30}, map[float64]int16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Int16P(map[float64]int16{v1: v10, v3: v10, v2: v20}, map[float64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int16P(nil, map[float64]int16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int16P(map[float64]int16{v1: v10, v2: v20, v3: v30}, map[float64]int16{})
	if r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int16P(nil, map[float64]int16{})
	if r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int16P(map[float64]int16{v1: v10, v2: v20, v3: v30}, map[float64]int16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Int16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Int16PPtr(map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Int16PPtr(map[*float64]*int16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int16PPtr(nil, map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int16PPtr(map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int16{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int16PPtr(nil, map[*float64]*int16{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int16PPtr(map[*float64]*int16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Int16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Int8(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 int8 = 10
	var v20 int8 = 20
	var v30 int8 = 30

	r := EqualMapFloat64Int8P(map[float64]int8{v1: v10, v2: v20, v3: v30}, map[float64]int8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Int8P(map[float64]int8{v1: v10, v3: v10, v2: v20}, map[float64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int8P(nil, map[float64]int8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int8P(map[float64]int8{v1: v10, v2: v20, v3: v30}, map[float64]int8{})
	if r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int8P(nil, map[float64]int8{})
	if r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Int8P(map[float64]int8{v1: v10, v2: v20, v3: v30}, map[float64]int8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Int8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Int8PPtr(map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Int8PPtr(map[*float64]*int8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int8PPtr(nil, map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int8PPtr(map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int8{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int8PPtr(nil, map[*float64]*int8{})
	if rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Int8PPtr(map[*float64]*int8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*int8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Int8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Uint(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 uint = 10
	var v20 uint = 20
	var v30 uint = 30

	r := EqualMapFloat64UintP(map[float64]uint{v1: v10, v2: v20, v3: v30}, map[float64]uint{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64UintP(map[float64]uint{v1: v10, v3: v10, v2: v20}, map[float64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64UintP(nil, map[float64]uint{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64UintP(map[float64]uint{v1: v10, v2: v20, v3: v30}, map[float64]uint{})
	if r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64UintP(nil, map[float64]uint{})
	if r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64UintP(map[float64]uint{v1: v10, v2: v20, v3: v30}, map[float64]uint{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64UintP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64UintPPtr(map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64UintPPtr(map[*float64]*uint{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64UintPPtr(nil, map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64UintPPtr(map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint{})
	if rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64UintPPtr(nil, map[*float64]*uint{})
	if rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64UintPPtr(map[*float64]*uint{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64UintPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Uint64(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 uint64 = 10
	var v20 uint64 = 20
	var v30 uint64 = 30

	r := EqualMapFloat64Uint64P(map[float64]uint64{v1: v10, v2: v20, v3: v30}, map[float64]uint64{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Uint64P(map[float64]uint64{v1: v10, v3: v10, v2: v20}, map[float64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint64P(nil, map[float64]uint64{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint64P(map[float64]uint64{v1: v10, v2: v20, v3: v30}, map[float64]uint64{})
	if r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint64P(nil, map[float64]uint64{})
	if r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint64P(map[float64]uint64{v1: v10, v2: v20, v3: v30}, map[float64]uint64{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Uint64P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Uint64PPtr(map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Uint64PPtr(map[*float64]*uint64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint64PPtr(nil, map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint64PPtr(map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint64PPtr(nil, map[*float64]*uint64{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint64PPtr(map[*float64]*uint64{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint64{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint64PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Uint32(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 uint32 = 10
	var v20 uint32 = 20
	var v30 uint32 = 30

	r := EqualMapFloat64Uint32P(map[float64]uint32{v1: v10, v2: v20, v3: v30}, map[float64]uint32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Uint32P(map[float64]uint32{v1: v10, v3: v10, v2: v20}, map[float64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint32P(nil, map[float64]uint32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint32P(map[float64]uint32{v1: v10, v2: v20, v3: v30}, map[float64]uint32{})
	if r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint32P(nil, map[float64]uint32{})
	if r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint32P(map[float64]uint32{v1: v10, v2: v20, v3: v30}, map[float64]uint32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Uint32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Uint32PPtr(map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Uint32PPtr(map[*float64]*uint32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint32PPtr(nil, map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint32PPtr(map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint32PPtr(nil, map[*float64]*uint32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint32PPtr(map[*float64]*uint32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Uint16(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 uint16 = 10
	var v20 uint16 = 20
	var v30 uint16 = 30

	r := EqualMapFloat64Uint16P(map[float64]uint16{v1: v10, v2: v20, v3: v30}, map[float64]uint16{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Uint16P(map[float64]uint16{v1: v10, v3: v10, v2: v20}, map[float64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint16P(nil, map[float64]uint16{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint16P(map[float64]uint16{v1: v10, v2: v20, v3: v30}, map[float64]uint16{})
	if r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint16P(nil, map[float64]uint16{})
	if r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint16P(map[float64]uint16{v1: v10, v2: v20, v3: v30}, map[float64]uint16{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Uint16P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Uint16PPtr(map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Uint16PPtr(map[*float64]*uint16{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint16PPtr(nil, map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint16PPtr(map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint16PPtr(nil, map[*float64]*uint16{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint16PPtr(map[*float64]*uint16{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint16{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint16PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Uint8(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 uint8 = 10
	var v20 uint8 = 20
	var v30 uint8 = 30

	r := EqualMapFloat64Uint8P(map[float64]uint8{v1: v10, v2: v20, v3: v30}, map[float64]uint8{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Uint8P(map[float64]uint8{v1: v10, v3: v10, v2: v20}, map[float64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint8P(nil, map[float64]uint8{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint8P(map[float64]uint8{v1: v10, v2: v20, v3: v30}, map[float64]uint8{})
	if r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint8P(nil, map[float64]uint8{})
	if r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Uint8P(map[float64]uint8{v1: v10, v2: v20, v3: v30}, map[float64]uint8{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Uint8P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Uint8PPtr(map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Uint8PPtr(map[*float64]*uint8{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint8PPtr(nil, map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint8PPtr(map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint8PPtr(nil, map[*float64]*uint8{})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Uint8PPtr(map[*float64]*uint8{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*uint8{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Uint8PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Str(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 string = "10"
	var v20 string = "20"
	var v30 string = "30"

	r := EqualMapFloat64StrP(map[float64]string{v1: v10, v2: v20, v3: v30}, map[float64]string{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64StrP(map[float64]string{v1: v10, v3: v10, v2: v20}, map[float64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64StrP(nil, map[float64]string{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64StrP(map[float64]string{v1: v10, v2: v20, v3: v30}, map[float64]string{})
	if r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64StrP(nil, map[float64]string{})
	if r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64StrP(map[float64]string{v1: v10, v2: v20, v3: v30}, map[float64]string{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64StrP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64StrPPtr(map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64StrPPtr(map[*float64]*string{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64StrPPtr(nil, map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64StrPPtr(map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*string{})
	if rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64StrPPtr(nil, map[*float64]*string{})
	if rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64StrPPtr(map[*float64]*string{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*string{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64StrPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Bool(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 bool = true
	var v20 bool = false
	var v30 bool = true

	r := EqualMapFloat64BoolP(map[float64]bool{v1: v10, v2: v20, v3: v30}, map[float64]bool{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64BoolP(map[float64]bool{v1: v20, v3: v10, v2: v20}, map[float64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64BoolP(nil, map[float64]bool{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64BoolP(map[float64]bool{v1: v10, v2: v20, v3: v30}, map[float64]bool{})
	if r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64BoolP(nil, map[float64]bool{})
	if r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64BoolP(map[float64]bool{v1: v10, v2: v20, v3: v30}, map[float64]bool{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64BoolP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64BoolPPtr(map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64BoolPPtr(map[*float64]*bool{&v1: &v20, &v3: &v10, &v2: &v20}, map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64BoolPPtr(nil, map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64BoolPPtr(map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*bool{})
	if rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64BoolPPtr(nil, map[*float64]*bool{})
	if rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64BoolPPtr(map[*float64]*bool{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*bool{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64BoolPPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64Float32(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	var v10 float32 = 10
	var v20 float32 = 20
	var v30 float32 = 30

	r := EqualMapFloat64Float32P(map[float64]float32{v1: v10, v2: v20, v3: v30}, map[float64]float32{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMapFloat64Float32P(map[float64]float32{v1: v10, v3: v10, v2: v20}, map[float64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Float32P(nil, map[float64]float32{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Float32P(map[float64]float32{v1: v10, v2: v20, v3: v30}, map[float64]float32{})
	if r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Float32P(nil, map[float64]float32{})
	if r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=false, actual=true")
	}

	r = EqualMapFloat64Float32P(map[float64]float32{v1: v10, v2: v20, v3: v30}, map[float64]float32{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMapFloat64Float32P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMapFloat64Float32PPtr(map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMapFloat64Float32PPtr(map[*float64]*float32{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Float32PPtr(nil, map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Float32PPtr(map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Float32PPtr(nil, map[*float64]*float32{})
	if rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMapFloat64Float32PPtr(map[*float64]*float32{&v1: &v10, &v2: &v20, &v3: &v30}, map[*float64]*float32{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMapFloat64Float32PPtr failed. Expected=false, actual=true")
	}
}

func TestEqualMapsFloat64(t *testing.T) {
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

	r = EqualMapFloat64P(map[float64]float64{v1: v10, v3: v10, v2: v20}, map[float64]float64{v1: v10, v2: v20, v3: v30})
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

	rPtr = EqualMapFloat64PPtr(map[*float64]*float64{&v1: &v10, &v3: &v10, &v2: &v20}, map[*float64]*float64{&v1: &v10, &v2: &v20, &v3: &v30})
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
