package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestOddInt(t *testing.T) {
	r := OddInt(11)
	if !r {
		t.Errorf("OddInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt(2)
	if r {
		t.Errorf("OddInt failed. Expected=false, actual=true")
	}

	var three int = 3
	rPtr := OddIntPtr(&three)
	if !rPtr {
		t.Errorf("OddIntPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt64(t *testing.T) {
	r := OddInt64(11)
	if !r {
		t.Errorf("OddInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt64(2)
	if r {
		t.Errorf("OddInt64 failed. Expected=false, actual=true")
	}

	var three int64 = 3
	rPtr := OddInt64Ptr(&three)
	if !rPtr {
		t.Errorf("OddInt64Ptr failed. Expected=true, actual=false")
	}
}

func TestOddInt32(t *testing.T) {
	r := OddInt32(11)
	if !r {
		t.Errorf("OddInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt32(2)
	if r {
		t.Errorf("OddInt32 failed. Expected=false, actual=true")
	}

	var three int32 = 3
	rPtr := OddInt32Ptr(&three)
	if !rPtr {
		t.Errorf("OddInt32Ptr failed. Expected=true, actual=false")
	}
}

func TestOddInt16(t *testing.T) {
	r := OddInt16(11)
	if !r {
		t.Errorf("OddInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt16(2)
	if r {
		t.Errorf("OddInt16 failed. Expected=false, actual=true")
	}

	var three int16 = 3
	rPtr := OddInt16Ptr(&three)
	if !rPtr {
		t.Errorf("OddInt16Ptr failed. Expected=true, actual=false")
	}
}

func TestOddInt8(t *testing.T) {
	r := OddInt8(11)
	if !r {
		t.Errorf("OddInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt8(2)
	if r {
		t.Errorf("OddInt8 failed. Expected=false, actual=true")
	}

	var three int8 = 3
	rPtr := OddInt8Ptr(&three)
	if !rPtr {
		t.Errorf("OddInt8Ptr failed. Expected=true, actual=false")
	}
}

func TestOddUint(t *testing.T) {
	r := OddUint(11)
	if !r {
		t.Errorf("OddUint failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint(2)
	if r {
		t.Errorf("OddUint failed. Expected=false, actual=true")
	}

	var three uint = 3
	rPtr := OddUintPtr(&three)
	if !rPtr {
		t.Errorf("OddUintPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint64(t *testing.T) {
	r := OddUint64(11)
	if !r {
		t.Errorf("OddUint64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint64(2)
	if r {
		t.Errorf("OddUint64 failed. Expected=false, actual=true")
	}

	var three uint64 = 3
	rPtr := OddUint64Ptr(&three)
	if !rPtr {
		t.Errorf("OddUint64Ptr failed. Expected=true, actual=false")
	}
}

func TestOddUint32(t *testing.T) {
	r := OddUint32(11)
	if !r {
		t.Errorf("OddUint32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint32(2)
	if r {
		t.Errorf("OddUint32 failed. Expected=false, actual=true")
	}

	var three uint32 = 3
	rPtr := OddUint32Ptr(&three)
	if !rPtr {
		t.Errorf("OddUint32Ptr failed. Expected=true, actual=false")
	}
}

func TestOddUint16(t *testing.T) {
	r := OddUint16(11)
	if !r {
		t.Errorf("OddUint16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint16(2)
	if r {
		t.Errorf("OddUint16 failed. Expected=false, actual=true")
	}

	var three uint16 = 3
	rPtr := OddUint16Ptr(&three)
	if !rPtr {
		t.Errorf("OddUint16Ptr failed. Expected=true, actual=false")
	}
}

func TestOddUint8(t *testing.T) {
	r := OddUint8(11)
	if !r {
		t.Errorf("OddUint8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint8(2)
	if r {
		t.Errorf("OddUint8 failed. Expected=false, actual=true")
	}

	var three uint8 = 3
	rPtr := OddUint8Ptr(&three)
	if !rPtr {
		t.Errorf("OddUint8Ptr failed. Expected=true, actual=false")
	}
}
