package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestNegIntWht(t *testing.T) {
	r := NegIntWht(-1)
	if !r {
		t.Errorf("NegIntWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegIntWht(1)
	if r {
		t.Errorf("NegIntWht failed. Expected=false, actual=true")
	}

	var zero int
	var one int = -1
	rPtr := NegIntWhtPtr(&one)
	if !rPtr {
		t.Errorf("NegIntWhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegIntWhtPtr(&zero)
	if rPtr {
		t.Errorf("NegIntWhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt64Wht(t *testing.T) {
	r := NegInt64Wht(-1)
	if !r {
		t.Errorf("NegInt64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt64Wht(1)
	if r {
		t.Errorf("NegInt64Wht failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = -1
	rPtr := NegInt64WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegInt64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt64WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegInt64WhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt32Wht(t *testing.T) {
	r := NegInt32Wht(-1)
	if !r {
		t.Errorf("NegInt32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt32Wht(1)
	if r {
		t.Errorf("NegInt32Wht failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = -1
	rPtr := NegInt32WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegInt32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt32WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegInt32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt16Wht(t *testing.T) {
	r := NegInt16Wht(-1)
	if !r {
		t.Errorf("NegInt16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt16Wht(1)
	if r {
		t.Errorf("NegInt16Wht failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = -1
	rPtr := NegInt16WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegInt16WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt16WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegInt16WhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegInt8Wht(t *testing.T) {
	r := NegInt8Wht(-1)
	if !r {
		t.Errorf("NegInt8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegInt8Wht(1)
	if r {
		t.Errorf("NegInt8Wht failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = -1
	rPtr := NegInt8WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegInt8WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegInt8WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegInt8WhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegFloat32Wht(t *testing.T) {
	r := NegFloat32Wht(-1)
	if !r {
		t.Errorf("NegFloat32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat32Wht(1)
	if r {
		t.Errorf("NegFloat32Wht failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = -1
	rPtr := NegFloat32WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegFloat32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat32WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegFloat32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestNegFloat64Wht(t *testing.T) {
	r := NegFloat64Wht(-1)
	if !r {
		t.Errorf("NegFloat64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = NegFloat64Wht(1)
	if r {
		t.Errorf("NegFloat64Wht failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = -1
	rPtr := NegFloat64WhtPtr(&one)
	if !rPtr {
		t.Errorf("NegFloat64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = NegFloat64WhtPtr(&zero)
	if rPtr {
		t.Errorf("NegFloat64WhtPtr failed. Expected=false, actual=true")
	}
}
