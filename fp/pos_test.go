package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPosInt(t *testing.T) {
	r := PosInt(1)
	if !r {
		t.Errorf("PosInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt(-1)
	if r {
		t.Errorf("PosInt failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := PosIntPtr(&one)
	if !rPtr {
		t.Errorf("PosIntPtr failed. Expected=true, actual=false")
	}

	rPtr = PosIntPtr(&zero)
	if rPtr {
		t.Errorf("PosIntPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt64(t *testing.T) {
	r := PosInt64(1)
	if !r {
		t.Errorf("PosInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt64(-1)
	if r {
		t.Errorf("PosInt64 failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := PosInt64Ptr(&one)
	if !rPtr {
		t.Errorf("PosInt64Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosInt64Ptr(&zero)
	if rPtr {
		t.Errorf("PosInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt32(t *testing.T) {
	r := PosInt32(1)
	if !r {
		t.Errorf("PosInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt32(-1)
	if r {
		t.Errorf("PosInt32 failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := PosInt32Ptr(&one)
	if !rPtr {
		t.Errorf("PosInt32Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosInt32Ptr(&zero)
	if rPtr {
		t.Errorf("PosInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt16(t *testing.T) {
	r := PosInt16(1)
	if !r {
		t.Errorf("PosInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt16(-1)
	if r {
		t.Errorf("PosInt16 failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := PosInt16Ptr(&one)
	if !rPtr {
		t.Errorf("PosInt16Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosInt16Ptr(&zero)
	if rPtr {
		t.Errorf("PosInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt8(t *testing.T) {
	r := PosInt8(1)
	if !r {
		t.Errorf("PosInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt8(-1)
	if r {
		t.Errorf("PosInt8 failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := PosInt8Ptr(&one)
	if !rPtr {
		t.Errorf("PosInt8Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosInt8Ptr(&zero)
	if rPtr {
		t.Errorf("PosInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestPosFloat32(t *testing.T) {
	r := PosFloat32(1)
	if !r {
		t.Errorf("PosFloat32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat32(-1)
	if r {
		t.Errorf("PosFloat32 failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := PosFloat32Ptr(&one)
	if !rPtr {
		t.Errorf("PosFloat32Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat32Ptr(&zero)
	if rPtr {
		t.Errorf("PosFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestPosFloat64(t *testing.T) {
	r := PosFloat64(1)
	if !r {
		t.Errorf("PosFloat64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat64(-1)
	if r {
		t.Errorf("PosFloat64 failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := PosFloat64Ptr(&one)
	if !rPtr {
		t.Errorf("PosFloat64Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat64Ptr(&zero)
	if rPtr {
		t.Errorf("PosFloat64Ptr failed. Expected=false, actual=true")
	}
}
