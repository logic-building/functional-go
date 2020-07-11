package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestZeroIntP(t *testing.T) {
	r := ZeroIntP(0)
	if !r {
		t.Errorf("ZeroIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroIntP(1)
	if r {
		t.Errorf("ZeroIntP failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := ZeroIntPPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroIntPPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroIntPPtr(&one)
	if rPtr {
		t.Errorf("ZeroIntPPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt64P(t *testing.T) {
	r := ZeroInt64P(0)
	if !r {
		t.Errorf("ZeroInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt64P(1)
	if r {
		t.Errorf("ZeroInt64P failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := ZeroInt64PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt64PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt64PPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt64PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt32P(t *testing.T) {
	r := ZeroInt32P(0)
	if !r {
		t.Errorf("ZeroInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt32P(1)
	if r {
		t.Errorf("ZeroInt32P failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := ZeroInt32PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt32PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt32PPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt32PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt16P(t *testing.T) {
	r := ZeroInt16P(0)
	if !r {
		t.Errorf("ZeroInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt16P(1)
	if r {
		t.Errorf("ZeroInt16P failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := ZeroInt16PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt16PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt16PPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt16PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt8P(t *testing.T) {
	r := ZeroInt8P(0)
	if !r {
		t.Errorf("ZeroInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt8P(1)
	if r {
		t.Errorf("ZeroInt8P failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := ZeroInt8PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt8PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt8PPtr(&one)
	if rPtr {
		t.Errorf("ZeroInt8PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUintP(t *testing.T) {
	r := ZeroUintP(0)
	if !r {
		t.Errorf("ZeroUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUintP(1)
	if r {
		t.Errorf("ZeroUintP failed. Expected=false, actual=true")
	}

	var zero uint
	var one uint = 1
	rPtr := ZeroUintPPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUintPPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUintPPtr(&one)
	if rPtr {
		t.Errorf("ZeroUintPPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint64P(t *testing.T) {
	r := ZeroUint64P(0)
	if !r {
		t.Errorf("ZeroUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint64P(1)
	if r {
		t.Errorf("ZeroUint64P failed. Expected=false, actual=true")
	}

	var zero uint64
	var one uint64 = 1
	rPtr := ZeroUint64PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint64PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint64PPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint64PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint32P(t *testing.T) {
	r := ZeroUint32P(0)
	if !r {
		t.Errorf("ZeroUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint32P(1)
	if r {
		t.Errorf("ZeroUint32P failed. Expected=false, actual=true")
	}

	var zero uint32
	var one uint32 = 1
	rPtr := ZeroUint32PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint32PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint32PPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint32PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint16P(t *testing.T) {
	r := ZeroUint16P(0)
	if !r {
		t.Errorf("ZeroUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint16P(1)
	if r {
		t.Errorf("ZeroUint16P failed. Expected=false, actual=true")
	}

	var zero uint16
	var one uint16 = 1
	rPtr := ZeroUint16PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint16PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint16PPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint16PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint8P(t *testing.T) {
	r := ZeroUint8P(0)
	if !r {
		t.Errorf("ZeroUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint8P(1)
	if r {
		t.Errorf("ZeroUint8P failed. Expected=false, actual=true")
	}

	var zero uint8
	var one uint8 = 1
	rPtr := ZeroUint8PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint8PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint8PPtr(&one)
	if rPtr {
		t.Errorf("ZeroUint8PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat32P(t *testing.T) {
	r := ZeroFloat32P(0)
	if !r {
		t.Errorf("ZeroFloat32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat32P(1)
	if r {
		t.Errorf("ZeroFloat32P failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := ZeroFloat32PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat32PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat32PPtr(&one)
	if rPtr {
		t.Errorf("ZeroFloat32PPtr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat64P(t *testing.T) {
	r := ZeroFloat64P(0)
	if !r {
		t.Errorf("ZeroFloat64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat64P(1)
	if r {
		t.Errorf("ZeroFloat64P failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := ZeroFloat64PPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat64PPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat64PPtr(&one)
	if rPtr {
		t.Errorf("ZeroFloat64PPtr failed. Expected=false, actual=true")
	}
}
