package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPosInt(t *testing.T) {
	r := PosIntP(1)
	if !r {
		t.Errorf("PosIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosIntP(-1)
	if r {
		t.Errorf("PosIntP failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := PosIntPPtr(&one)
	if !rPtr {
		t.Errorf("PosIntPPtr failed. Expected=true, actual=false")
	}

	rPtr = PosIntPPtr(&zero)
	if rPtr {
		t.Errorf("PosIntPPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt64(t *testing.T) {
	r := PosInt64P(1)
	if !r {
		t.Errorf("PosInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt64P(-1)
	if r {
		t.Errorf("PosInt64P failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := PosInt64PPtr(&one)
	if !rPtr {
		t.Errorf("PosInt64PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt64PPtr(&zero)
	if rPtr {
		t.Errorf("PosInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt32(t *testing.T) {
	r := PosInt32P(1)
	if !r {
		t.Errorf("PosInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt32P(-1)
	if r {
		t.Errorf("PosInt32P failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := PosInt32PPtr(&one)
	if !rPtr {
		t.Errorf("PosInt32PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt32PPtr(&zero)
	if rPtr {
		t.Errorf("PosInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt16(t *testing.T) {
	r := PosInt16P(1)
	if !r {
		t.Errorf("PosInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt16P(-1)
	if r {
		t.Errorf("PosInt16P failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := PosInt16PPtr(&one)
	if !rPtr {
		t.Errorf("PosInt16PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt16PPtr(&zero)
	if rPtr {
		t.Errorf("PosInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt8(t *testing.T) {
	r := PosInt8P(1)
	if !r {
		t.Errorf("PosInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt8P(-1)
	if r {
		t.Errorf("PosInt8P failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := PosInt8PPtr(&one)
	if !rPtr {
		t.Errorf("PosInt8PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt8PPtr(&zero)
	if rPtr {
		t.Errorf("PosInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestPosFloat32(t *testing.T) {
	r := PosFloat32P(1)
	if !r {
		t.Errorf("PosFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat32P(-1)
	if r {
		t.Errorf("PosFloat32P failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := PosFloat32PPtr(&one)
	if !rPtr {
		t.Errorf("PosFloat32PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat32PPtr(&zero)
	if rPtr {
		t.Errorf("PosFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestPosFloat64(t *testing.T) {
	r := PosFloat64P(1)
	if !r {
		t.Errorf("PosFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat64P(-1)
	if r {
		t.Errorf("PosFloat64P failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := PosFloat64PPtr(&one)
	if !rPtr {
		t.Errorf("PosFloat64PPtr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat64PPtr(&zero)
	if rPtr {
		t.Errorf("PosFloat64PPtr failed. Expected=false, actual=true")
	}
}
