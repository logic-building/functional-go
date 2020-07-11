package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPosInt(t *testing.T) {
	r := PosWhtInt(1)
	if !r {
		t.Errorf("PosWhtInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtInt(-1)
	if r {
		t.Errorf("PosWhtInt failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := PosWhtIntPtr(&one)
	if !rPtr {
		t.Errorf("PosWhtIntPtr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtIntPtr(&zero)
	if rPtr {
		t.Errorf("PosWhtIntPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt64(t *testing.T) {
	r := PosWhtInt64(1)
	if !r {
		t.Errorf("PosWhtInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtInt64(-1)
	if r {
		t.Errorf("PosWhtInt64 failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := PosWhtInt64Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtInt64Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtInt64Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt32(t *testing.T) {
	r := PosWhtInt32(1)
	if !r {
		t.Errorf("PosWhtInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtInt32(-1)
	if r {
		t.Errorf("PosWhtInt32 failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := PosWhtInt32Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtInt32Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtInt32Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt16(t *testing.T) {
	r := PosWhtInt16(1)
	if !r {
		t.Errorf("PosWhtInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtInt16(-1)
	if r {
		t.Errorf("PosWhtInt16 failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := PosWhtInt16Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtInt16Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtInt16Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestPosInt8(t *testing.T) {
	r := PosWhtInt8(1)
	if !r {
		t.Errorf("PosWhtInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtInt8(-1)
	if r {
		t.Errorf("PosWhtInt8 failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := PosWhtInt8Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtInt8Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtInt8Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestPosFloat32(t *testing.T) {
	r := PosWhtFloat32(1)
	if !r {
		t.Errorf("PosWhtFloat32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtFloat32(-1)
	if r {
		t.Errorf("PosWhtFloat32 failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := PosWhtFloat32Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtFloat32Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtFloat32Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestPosFloat64(t *testing.T) {
	r := PosWhtFloat64(1)
	if !r {
		t.Errorf("PosWhtFloat64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWhtFloat64(-1)
	if r {
		t.Errorf("PosWhtFloat64 failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := PosWhtFloat64Ptr(&one)
	if !rPtr {
		t.Errorf("PosWhtFloat64Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWhtFloat64Ptr(&zero)
	if rPtr {
		t.Errorf("PosWhtFloat64Ptr failed. Expected=false, actual=true")
	}
}
