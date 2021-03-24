package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEvenInt(t *testing.T) {
	r := EvenIntP(10)
	if !r {
		t.Errorf("EvenIntP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenIntP(1)
	if r {
		t.Errorf("EvenIntP failed. Expected=false, actual=true")
	}

	var two int = 2
	rPtr := EvenIntPPtr(&two)
	if !rPtr {
		t.Errorf("EvenIntPPtr failed. Expected=true, actual=false")
	}
}

func TestEvenInt64(t *testing.T) {
	r := EvenInt64P(10)
	if !r {
		t.Errorf("EvenInt64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt64P(1)
	if r {
		t.Errorf("EvenInt64P failed. Expected=false, actual=true")
	}

	var two int64 = 2
	rPtr := EvenInt64PPtr(&two)
	if !rPtr {
		t.Errorf("EvenInt64PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenInt32(t *testing.T) {
	r := EvenInt32P(10)
	if !r {
		t.Errorf("EvenInt32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt32P(1)
	if r {
		t.Errorf("EvenInt32P failed. Expected=false, actual=true")
	}

	var two int32 = 2
	rPtr := EvenInt32PPtr(&two)
	if !rPtr {
		t.Errorf("EvenInt32PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenInt16(t *testing.T) {
	r := EvenInt16P(10)
	if !r {
		t.Errorf("EvenInt16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt16P(1)
	if r {
		t.Errorf("EvenInt16P failed. Expected=false, actual=true")
	}

	var two int16 = 2
	rPtr := EvenInt16PPtr(&two)
	if !rPtr {
		t.Errorf("EvenInt16PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenInt8(t *testing.T) {
	r := EvenInt8P(10)
	if !r {
		t.Errorf("EvenInt8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt8P(1)
	if r {
		t.Errorf("EvenInt8P failed. Expected=false, actual=true")
	}

	var two int8 = 2
	rPtr := EvenInt8PPtr(&two)
	if !rPtr {
		t.Errorf("EvenInt8PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint(t *testing.T) {
	r := EvenUintP(10)
	if !r {
		t.Errorf("EvenUintP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUintP(1)
	if r {
		t.Errorf("EvenUintP failed. Expected=false, actual=true")
	}

	var two uint = 2
	rPtr := EvenUintPPtr(&two)
	if !rPtr {
		t.Errorf("EvenUintPPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint64(t *testing.T) {
	r := EvenUint64P(10)
	if !r {
		t.Errorf("EvenUint64P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint64P(1)
	if r {
		t.Errorf("EvenUint64P failed. Expected=false, actual=true")
	}

	var two uint64 = 2
	rPtr := EvenUint64PPtr(&two)
	if !rPtr {
		t.Errorf("EvenUint64PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint32(t *testing.T) {
	r := EvenUint32P(10)
	if !r {
		t.Errorf("EvenUint32P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint32P(1)
	if r {
		t.Errorf("EvenUint32P failed. Expected=false, actual=true")
	}

	var two uint32 = 2
	rPtr := EvenUint32PPtr(&two)
	if !rPtr {
		t.Errorf("EvenUint32PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint16(t *testing.T) {
	r := EvenUint16P(10)
	if !r {
		t.Errorf("EvenUint16P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint16P(1)
	if r {
		t.Errorf("EvenUint16P failed. Expected=false, actual=true")
	}

	var two uint16 = 2
	rPtr := EvenUint16PPtr(&two)
	if !rPtr {
		t.Errorf("EvenUint16PPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint8(t *testing.T) {
	r := EvenUint8P(10)
	if !r {
		t.Errorf("EvenUint8P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint8P(1)
	if r {
		t.Errorf("EvenUint8P failed. Expected=false, actual=true")
	}

	var two uint8 = 2
	rPtr := EvenUint8PPtr(&two)
	if !rPtr {
		t.Errorf("EvenUint8PPtr failed. Expected=true, actual=false")
	}
}
