package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPosInt(t *testing.T) {
	r := PosIntWht(1)
	if !r {
		t.Errorf("PosIntWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosIntWht(-1)
	if r {
		t.Errorf("PosIntWht failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := PosIntWhtPtr(&one)
	if !rPtr {
		t.Errorf("PosIntWhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosIntWhtPtr(&zero)
	if rPtr {
		t.Errorf("PosIntWhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt64(t *testing.T) {
	r := PosInt64Wht(1)
	if !r {
		t.Errorf("PosInt64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt64Wht(-1)
	if r {
		t.Errorf("PosInt64Wht failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := PosInt64WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosInt64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt64WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosInt64WhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt32(t *testing.T) {
	r := PosInt32Wht(1)
	if !r {
		t.Errorf("PosInt32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt32Wht(-1)
	if r {
		t.Errorf("PosInt32Wht failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := PosInt32WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosInt32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt32WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosInt32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt16(t *testing.T) {
	r := PosInt16Wht(1)
	if !r {
		t.Errorf("PosInt16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt16Wht(-1)
	if r {
		t.Errorf("PosInt16Wht failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := PosInt16WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosInt16WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt16WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosInt16WhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosInt8(t *testing.T) {
	r := PosInt8Wht(1)
	if !r {
		t.Errorf("PosInt8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosInt8Wht(-1)
	if r {
		t.Errorf("PosInt8Wht failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := PosInt8WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosInt8WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosInt8WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosInt8WhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosFloat32(t *testing.T) {
	r := PosFloat32Wht(1)
	if !r {
		t.Errorf("PosFloat32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat32Wht(-1)
	if r {
		t.Errorf("PosFloat32Wht failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := PosFloat32WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosFloat32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat32WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosFloat32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestPosFloat64(t *testing.T) {
	r := PosFloat64Wht(1)
	if !r {
		t.Errorf("PosFloat64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosFloat64Wht(-1)
	if r {
		t.Errorf("PosFloat64Wht failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := PosFloat64WhtPtr(&one)
	if !rPtr {
		t.Errorf("PosFloat64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = PosFloat64WhtPtr(&zero)
	if rPtr {
		t.Errorf("PosFloat64WhtPtr failed. Expected=false, actual=true")
	}
}
