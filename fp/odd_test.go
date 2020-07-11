package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestOddInt(t *testing.T) {
	r := OddIntWht(11)
	if !r {
		t.Errorf("OddIntWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddIntWht(2)
	if r {
		t.Errorf("OddIntWht failed. Expected=false, actual=true")
	}

	var three int = 3
	rPtr := OddIntWhtPtr(&three)
	if !rPtr {
		t.Errorf("OddIntWhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt64(t *testing.T) {
	r := OddInt64Wht(11)
	if !r {
		t.Errorf("OddInt64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt64Wht(2)
	if r {
		t.Errorf("OddInt64Wht failed. Expected=false, actual=true")
	}

	var three int64 = 3
	rPtr := OddInt64WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddInt64WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt32(t *testing.T) {
	r := OddInt32Wht(11)
	if !r {
		t.Errorf("OddInt32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt32Wht(2)
	if r {
		t.Errorf("OddInt32Wht failed. Expected=false, actual=true")
	}

	var three int32 = 3
	rPtr := OddInt32WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddInt32WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt16(t *testing.T) {
	r := OddInt16Wht(11)
	if !r {
		t.Errorf("OddInt16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt16Wht(2)
	if r {
		t.Errorf("OddInt16Wht failed. Expected=false, actual=true")
	}

	var three int16 = 3
	rPtr := OddInt16WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddInt16WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddInt8(t *testing.T) {
	r := OddInt8Wht(11)
	if !r {
		t.Errorf("OddInt8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddInt8Wht(2)
	if r {
		t.Errorf("OddInt8Wht failed. Expected=false, actual=true")
	}

	var three int8 = 3
	rPtr := OddInt8WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddInt8WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint(t *testing.T) {
	r := OddUintWht(11)
	if !r {
		t.Errorf("OddUintWht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUintWht(2)
	if r {
		t.Errorf("OddUintWht failed. Expected=false, actual=true")
	}

	var three uint = 3
	rPtr := OddUintWhtPtr(&three)
	if !rPtr {
		t.Errorf("OddUintWhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint64(t *testing.T) {
	r := OddUint64Wht(11)
	if !r {
		t.Errorf("OddUint64Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint64Wht(2)
	if r {
		t.Errorf("OddUint64Wht failed. Expected=false, actual=true")
	}

	var three uint64 = 3
	rPtr := OddUint64WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddUint64WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint32(t *testing.T) {
	r := OddUint32Wht(11)
	if !r {
		t.Errorf("OddUint32Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint32Wht(2)
	if r {
		t.Errorf("OddUint32Wht failed. Expected=false, actual=true")
	}

	var three uint32 = 3
	rPtr := OddUint32WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddUint32WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint16(t *testing.T) {
	r := OddUint16Wht(11)
	if !r {
		t.Errorf("OddUint16Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint16Wht(2)
	if r {
		t.Errorf("OddUint16Wht failed. Expected=false, actual=true")
	}

	var three uint16 = 3
	rPtr := OddUint16WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddUint16WhtPtr failed. Expected=true, actual=false")
	}
}

func TestOddUint8(t *testing.T) {
	r := OddUint8Wht(11)
	if !r {
		t.Errorf("OddUint8Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = OddUint8Wht(2)
	if r {
		t.Errorf("OddUint8Wht failed. Expected=false, actual=true")
	}

	var three uint8 = 3
	rPtr := OddUint8WhtPtr(&three)
	if !rPtr {
		t.Errorf("OddUint8WhtPtr failed. Expected=true, actual=false")
	}
}
