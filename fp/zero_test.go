package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestZeroInt(t *testing.T) {
	r := ZeroInt(0)
	if !r {
		t.Errorf("ZeroInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt(1)
	if r {
		t.Errorf("ZeroInt failed. Expected=false, actual=true")
	}

	var zero int
	var one int = 1
	rPtr := ZeroIntPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroIntPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroIntPtr(&one)
	if rPtr {
		t.Errorf("ZeroIntPtr failed. Expected=false, actual=true")
	}
}

func TestZeroInt64(t *testing.T) {
	r := ZeroInt64(0)
	if !r {
		t.Errorf("ZeroInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt64(1)
	if r {
		t.Errorf("ZeroInt64 failed. Expected=false, actual=true")
	}

	var zero int64
	var one int64 = 1
	rPtr := ZeroInt64Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt64Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt64Ptr(&one)
	if rPtr {
		t.Errorf("ZeroInt64Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroInt32(t *testing.T) {
	r := ZeroInt32(0)
	if !r {
		t.Errorf("ZeroInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt32(1)
	if r {
		t.Errorf("ZeroInt32 failed. Expected=false, actual=true")
	}

	var zero int32
	var one int32 = 1
	rPtr := ZeroInt32Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt32Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt32Ptr(&one)
	if rPtr {
		t.Errorf("ZeroInt32Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroInt16(t *testing.T) {
	r := ZeroInt16(0)
	if !r {
		t.Errorf("ZeroInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt16(1)
	if r {
		t.Errorf("ZeroInt16 failed. Expected=false, actual=true")
	}

	var zero int16
	var one int16 = 1
	rPtr := ZeroInt16Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt16Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt16Ptr(&one)
	if rPtr {
		t.Errorf("ZeroInt16Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroInt8(t *testing.T) {
	r := ZeroInt8(0)
	if !r {
		t.Errorf("ZeroInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroInt8(1)
	if r {
		t.Errorf("ZeroInt8 failed. Expected=false, actual=true")
	}

	var zero int8
	var one int8 = 1
	rPtr := ZeroInt8Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroInt8Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroInt8Ptr(&one)
	if rPtr {
		t.Errorf("ZeroInt8Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroUint(t *testing.T) {
	r := ZeroUint(0)
	if !r {
		t.Errorf("ZeroUint failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint(1)
	if r {
		t.Errorf("ZeroUint failed. Expected=false, actual=true")
	}

	var zero uint
	var one uint = 1
	rPtr := ZeroUintPtr(&zero)
	if !rPtr {
		t.Errorf("ZeroUintPtr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUintPtr(&one)
	if rPtr {
		t.Errorf("ZeroUintPtr failed. Expected=false, actual=true")
	}
}

func TestZeroUint64(t *testing.T) {
	r := ZeroUint64(0)
	if !r {
		t.Errorf("ZeroUint64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint64(1)
	if r {
		t.Errorf("ZeroUint64 failed. Expected=false, actual=true")
	}

	var zero uint64
	var one uint64 = 1
	rPtr := ZeroUint64Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint64Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint64Ptr(&one)
	if rPtr {
		t.Errorf("ZeroUint64Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroUint32(t *testing.T) {
	r := ZeroUint32(0)
	if !r {
		t.Errorf("ZeroUint32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint32(1)
	if r {
		t.Errorf("ZeroUint32 failed. Expected=false, actual=true")
	}

	var zero uint32
	var one uint32 = 1
	rPtr := ZeroUint32Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint32Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint32Ptr(&one)
	if rPtr {
		t.Errorf("ZeroUint32Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroUint16(t *testing.T) {
	r := ZeroUint16(0)
	if !r {
		t.Errorf("ZeroUint16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint16(1)
	if r {
		t.Errorf("ZeroUint16 failed. Expected=false, actual=true")
	}

	var zero uint16
	var one uint16 = 1
	rPtr := ZeroUint16Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint16Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint16Ptr(&one)
	if rPtr {
		t.Errorf("ZeroUint16Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroUint8(t *testing.T) {
	r := ZeroUint8(0)
	if !r {
		t.Errorf("ZeroUint8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroUint8(1)
	if r {
		t.Errorf("ZeroUint8 failed. Expected=false, actual=true")
	}

	var zero uint8
	var one uint8 = 1
	rPtr := ZeroUint8Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroUint8Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroUint8Ptr(&one)
	if rPtr {
		t.Errorf("ZeroUint8Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat32(t *testing.T) {
	r := ZeroFloat32(0)
	if !r {
		t.Errorf("ZeroFloat32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat32(1)
	if r {
		t.Errorf("ZeroFloat32 failed. Expected=false, actual=true")
	}

	var zero float32
	var one float32 = 1
	rPtr := ZeroFloat32Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat32Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat32Ptr(&one)
	if rPtr {
		t.Errorf("ZeroFloat32Ptr failed. Expected=false, actual=true")
	}
}

func TestZeroFloat64(t *testing.T) {
	r := ZeroFloat64(0)
	if !r {
		t.Errorf("ZeroFloat64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = ZeroFloat64(1)
	if r {
		t.Errorf("ZeroFloat64 failed. Expected=false, actual=true")
	}

	var zero float64
	var one float64 = 1
	rPtr := ZeroFloat64Ptr(&zero)
	if !rPtr {
		t.Errorf("ZeroFloat64Ptr failed. Expected=true, actual=false")
	}

	rPtr = ZeroFloat64Ptr(&one)
	if rPtr {
		t.Errorf("ZeroFloat64Ptr failed. Expected=false, actual=true")
	}
}
