package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestNegInt(t *testing.T) {
	r := NegInt(-1)
	if !r {
		t.Errorf("NegInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt(1)
	if r {
		t.Errorf("NegInt failed. Expected=false, actual=true")
	}

	var zero int
	var one int = -1
	rPtr := NegIntPtr(&one)
	if !rPtr {
		t.Errorf("NegIntPtr failed. Expected=true, actual=false")
	}

	rPtr = NegIntPtr(&zero)
	if rPtr {
		t.Errorf("NegIntPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt64(t *testing.T) {
	r := NegInt64(-1)
	if !r {
		t.Errorf("NegInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt64(1)
	if r {
		t.Errorf("NegInt64 failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = -1
	rPtr := NegInt64Ptr(&one)
	if !rPtr {
		t.Errorf("NegInt64Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegInt64Ptr(&zero)
	if rPtr {
		t.Errorf("NegInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestNegInt32(t *testing.T) {
	r := NegInt32(-1)
	if !r {
		t.Errorf("NegInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt32(1)
	if r {
		t.Errorf("NegInt32 failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = -1
	rPtr := NegInt32Ptr(&one)
	if !rPtr {
		t.Errorf("NegInt32Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegInt32Ptr(&zero)
	if rPtr {
		t.Errorf("NegInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestNegInt16(t *testing.T) {
	r := NegInt16(-1)
	if !r {
		t.Errorf("NegInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt16(1)
	if r {
		t.Errorf("NegInt16 failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = -1
	rPtr := NegInt16Ptr(&one)
	if !rPtr {
		t.Errorf("NegInt16Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegInt16Ptr(&zero)
	if rPtr {
		t.Errorf("NegInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestNegInt8(t *testing.T) {
	r := NegInt8(-1)
	if !r {
		t.Errorf("NegInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt8(1)
	if r {
		t.Errorf("NegInt8 failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = -1
	rPtr := NegInt8Ptr(&one)
	if !rPtr {
		t.Errorf("NegInt8Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegInt8Ptr(&zero)
	if rPtr {
		t.Errorf("NegInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestNegFloat32(t *testing.T) {
	r := NegFloat32(-1)
	if !r {
		t.Errorf("NegFloat32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat32(1)
	if r {
		t.Errorf("NegFloat32 failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = -1
	rPtr := NegFloat32Ptr(&one)
	if !rPtr {
		t.Errorf("NegFloat32Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat32Ptr(&zero)
	if rPtr {
		t.Errorf("NegFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestNegFloat64(t *testing.T) {
	r := NegFloat64(-1)
	if !r {
		t.Errorf("NegFloat64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat64(1)
	if r {
		t.Errorf("NegFloat64 failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = -1
	rPtr := NegFloat64Ptr(&one)
	if !rPtr {
		t.Errorf("NegFloat64Ptr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat64Ptr(&zero)
	if rPtr {
		t.Errorf("NegFloat64Ptr failed. Expected=false, actual=true")
	}
}
