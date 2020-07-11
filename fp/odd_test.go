package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestOddInt(t *testing.T) {
	r := OddIntP(11)
	if !r {
		t.Errorf("OddIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddIntP(2)
	if r {
		t.Errorf("OddIntP failed. Expected=false, actual=true")
	}

	var three int = 3
	rPtr := OddIntPPtr(&three)
	if !rPtr {
		t.Errorf("OddIntPPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt64(t *testing.T) {
	r := OddInt64P(11)
	if !r {
		t.Errorf("OddInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt64P(2)
	if r {
		t.Errorf("OddInt64P failed. Expected=false, actual=true")
	}

	var three int64 = 3
	rPtr := OddInt64PPtr(&three)
	if !rPtr {
		t.Errorf("OddInt64PPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt32(t *testing.T) {
	r := OddInt32P(11)
	if !r {
		t.Errorf("OddInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt32P(2)
	if r {
		t.Errorf("OddInt32P failed. Expected=false, actual=true")
	}

	var three int32 = 3
	rPtr := OddInt32PPtr(&three)
	if !rPtr {
		t.Errorf("OddInt32PPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt16(t *testing.T) {
	r := OddInt16P(11)
	if !r {
		t.Errorf("OddInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt16P(2)
	if r {
		t.Errorf("OddInt16P failed. Expected=false, actual=true")
	}

	var three int16 = 3
	rPtr := OddInt16PPtr(&three)
	if !rPtr {
		t.Errorf("OddInt16PPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt8(t *testing.T) {
	r := OddInt8P(11)
	if !r {
		t.Errorf("OddInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt8P(2)
	if r {
		t.Errorf("OddInt8P failed. Expected=false, actual=true")
	}

	var three int8 = 3
	rPtr := OddInt8PPtr(&three)
	if !rPtr {
		t.Errorf("OddInt8PPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint(t *testing.T) {
	r := OddUintP(11)
	if !r {
		t.Errorf("OddUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUintP(2)
	if r {
		t.Errorf("OddUintP failed. Expected=false, actual=true")
	}

	var three uint = 3
	rPtr := OddUintPPtr(&three)
	if !rPtr {
		t.Errorf("OddUintPPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint64(t *testing.T) {
	r := OddUint64P(11)
	if !r {
		t.Errorf("OddUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint64P(2)
	if r {
		t.Errorf("OddUint64P failed. Expected=false, actual=true")
	}

	var three uint64 = 3
	rPtr := OddUint64PPtr(&three)
	if !rPtr {
		t.Errorf("OddUint64PPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint32(t *testing.T) {
	r := OddUint32P(11)
	if !r {
		t.Errorf("OddUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint32P(2)
	if r {
		t.Errorf("OddUint32P failed. Expected=false, actual=true")
	}

	var three uint32 = 3
	rPtr := OddUint32PPtr(&three)
	if !rPtr {
		t.Errorf("OddUint32PPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint16(t *testing.T) {
	r := OddUint16P(11)
	if !r {
		t.Errorf("OddUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint16P(2)
	if r {
		t.Errorf("OddUint16P failed. Expected=false, actual=true")
	}

	var three uint16 = 3
	rPtr := OddUint16PPtr(&three)
	if !rPtr {
		t.Errorf("OddUint16PPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint8(t *testing.T) {
	r := OddUint8P(11)
	if !r {
		t.Errorf("OddUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint8P(2)
	if r {
		t.Errorf("OddUint8P failed. Expected=false, actual=true")
	}

	var three uint8 = 3
	rPtr := OddUint8PPtr(&three)
	if !rPtr {
		t.Errorf("OddUint8PPtr failed. Expected=true, actual=false")
	}
}
