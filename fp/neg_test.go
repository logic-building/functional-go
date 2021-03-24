package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestNegIntP(t *testing.T) {
	r := NegIntP(-1)
	if !r {
		t.Errorf("NegIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegIntP(1)
	if r {
		t.Errorf("NegIntP failed. Expected=false, actual=true")
	}

	var zero int
	var one int = -1
	rPtr := NegIntPPtr(&one)
	if !rPtr {
		t.Errorf("NegIntPPtr failed. Expected=true, actual=false")
	}

	rPtr = NegIntPPtr(&zero)
	if rPtr {
		t.Errorf("NegIntPPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt64P(t *testing.T) {
	r := NegInt64P(-1)
	if !r {
		t.Errorf("NegInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt64P(1)
	if r {
		t.Errorf("NegInt64P failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = -1
	rPtr := NegInt64PPtr(&one)
	if !rPtr {
		t.Errorf("NegInt64PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt64PPtr(&zero)
	if rPtr {
		t.Errorf("NegInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt32P(t *testing.T) {
	r := NegInt32P(-1)
	if !r {
		t.Errorf("NegInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt32P(1)
	if r {
		t.Errorf("NegInt32P failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = -1
	rPtr := NegInt32PPtr(&one)
	if !rPtr {
		t.Errorf("NegInt32PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt32PPtr(&zero)
	if rPtr {
		t.Errorf("NegInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt16P(t *testing.T) {
	r := NegInt16P(-1)
	if !r {
		t.Errorf("NegInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt16P(1)
	if r {
		t.Errorf("NegInt16P failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = -1
	rPtr := NegInt16PPtr(&one)
	if !rPtr {
		t.Errorf("NegInt16PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt16PPtr(&zero)
	if rPtr {
		t.Errorf("NegInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt8P(t *testing.T) {
	r := NegInt8P(-1)
	if !r {
		t.Errorf("NegInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt8P(1)
	if r {
		t.Errorf("NegInt8P failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = -1
	rPtr := NegInt8PPtr(&one)
	if !rPtr {
		t.Errorf("NegInt8PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt8PPtr(&zero)
	if rPtr {
		t.Errorf("NegInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestNegFloat32P(t *testing.T) {
	r := NegFloat32P(-1)
	if !r {
		t.Errorf("NegFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat32P(1)
	if r {
		t.Errorf("NegFloat32P failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = -1
	rPtr := NegFloat32PPtr(&one)
	if !rPtr {
		t.Errorf("NegFloat32PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat32PPtr(&zero)
	if rPtr {
		t.Errorf("NegFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestNegFloat64P(t *testing.T) {
	r := NegFloat64P(-1)
	if !r {
		t.Errorf("NegFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat64P(1)
	if r {
		t.Errorf("NegFloat64P failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = -1
	rPtr := NegFloat64PPtr(&one)
	if !rPtr {
		t.Errorf("NegFloat64PPtr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat64PPtr(&zero)
	if rPtr {
		t.Errorf("NegFloat64PPtr failed. Expected=false, actual=true")
	}
}
