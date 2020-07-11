package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestZeroIntWht(t *testing.T) {
	r := ZeroIntWht(0)
	if !r {
		t.Errorf("ZeroIntWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroIntWht(1)
	if r {
		t.Errorf("ZeroIntWht failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := ZeroIntWhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroIntWhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroIntWhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroIntWhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt64Wht(t *testing.T) {
	r := ZeroInt64Wht(0)
	if !r {
		t.Errorf("ZeroInt64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt64Wht(1)
	if r {
		t.Errorf("ZeroInt64Wht failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := ZeroInt64WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt64WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt64WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt32Wht(t *testing.T) {
	r := ZeroInt32Wht(0)
	if !r {
		t.Errorf("ZeroInt32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt32Wht(1)
	if r {
		t.Errorf("ZeroInt32Wht failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := ZeroInt32WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt32WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt16Wht(t *testing.T) {
	r := ZeroInt16Wht(0)
	if !r {
		t.Errorf("ZeroInt16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt16Wht(1)
	if r {
		t.Errorf("ZeroInt16Wht failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := ZeroInt16WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt16WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt16WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt16WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt8Wht(t *testing.T) {
	r := ZeroInt8Wht(0)
	if !r {
		t.Errorf("ZeroInt8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt8Wht(1)
	if r {
		t.Errorf("ZeroInt8Wht failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := ZeroInt8WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt8WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt8WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt8WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUintWht(t *testing.T) {
	r := ZeroUintWht(0)
	if !r {
		t.Errorf("ZeroUintWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUintWht(1)
	if r {
		t.Errorf("ZeroUintWht failed. Expected=false, actual=true")
	}

	var zero uint
	var one uint = 1
	rPtr := ZeroUintWhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUintWhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUintWhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroUintWhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint64Wht(t *testing.T) {
	r := ZeroUint64Wht(0)
	if !r {
		t.Errorf("ZeroUint64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint64Wht(1)
	if r {
		t.Errorf("ZeroUint64Wht failed. Expected=false, actual=true")
	}

	var zero uint64
	var one uint64 = 1
	rPtr := ZeroUint64WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint64WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint64WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint32Wht(t *testing.T) {
	r := ZeroUint32Wht(0)
	if !r {
		t.Errorf("ZeroUint32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint32Wht(1)
	if r {
		t.Errorf("ZeroUint32Wht failed. Expected=false, actual=true")
	}

	var zero uint32
	var one uint32 = 1
	rPtr := ZeroUint32WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint32WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint16Wht(t *testing.T) {
	r := ZeroUint16Wht(0)
	if !r {
		t.Errorf("ZeroUint16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint16Wht(1)
	if r {
		t.Errorf("ZeroUint16Wht failed. Expected=false, actual=true")
	}

	var zero uint16
	var one uint16 = 1
	rPtr := ZeroUint16WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint16WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint16WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint16WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint8Wht(t *testing.T) {
	r := ZeroUint8Wht(0)
	if !r {
		t.Errorf("ZeroUint8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint8Wht(1)
	if r {
		t.Errorf("ZeroUint8Wht failed. Expected=false, actual=true")
	}

	var zero uint8
	var one uint8 = 1
	rPtr := ZeroUint8WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint8WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint8WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint8WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat32Wht(t *testing.T) {
	r := ZeroFloat32Wht(0)
	if !r {
		t.Errorf("ZeroFloat32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat32Wht(1)
	if r {
		t.Errorf("ZeroFloat32Wht failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := ZeroFloat32WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat32WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat32WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroFloat32WhtPtr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat64Wht(t *testing.T) {
	r := ZeroFloat64Wht(0)
	if !r {
		t.Errorf("ZeroFloat64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat64Wht(1)
	if r {
		t.Errorf("ZeroFloat64Wht failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := ZeroFloat64WhtPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat64WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat64WhtPtr(&one)
	if rPtr {
		t.Errorf("ZeroFloat64WhtPtr failed. Expected=false, actual=true")
	}
}
