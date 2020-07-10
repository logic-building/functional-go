package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEvenInt(t *testing.T) {
	r := EvenInt(10)
	if !r {
		t.Errorf("EvenInt failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt(1)
	if r {
		t.Errorf("EvenInt failed. Expected=false, actual=true")
	}

	var two int = 2
	rPtr := EvenIntPtr(&two)
	if !rPtr {
		t.Errorf("EvenIntPtr failed. Expected=true, actual=false")
	}
}

func TestEvenInt64(t *testing.T) {
	r := EvenInt64(10)
	if !r {
		t.Errorf("EvenInt64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt64(1)
	if r {
		t.Errorf("EvenInt64 failed. Expected=false, actual=true")
	}

	var two int64 = 2
	rPtr := EvenInt64Ptr(&two)
	if !rPtr {
		t.Errorf("EvenInt64Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenInt32(t *testing.T) {
	r := EvenInt32(10)
	if !r {
		t.Errorf("EvenInt32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt32(1)
	if r {
		t.Errorf("EvenInt32 failed. Expected=false, actual=true")
	}

	var two int32 = 2
	rPtr := EvenInt32Ptr(&two)
	if !rPtr {
		t.Errorf("EvenInt32Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenInt16(t *testing.T) {
	r := EvenInt16(10)
	if !r {
		t.Errorf("EvenInt16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt16(1)
	if r {
		t.Errorf("EvenInt16 failed. Expected=false, actual=true")
	}

	var two int16 = 2
	rPtr := EvenInt16Ptr(&two)
	if !rPtr {
		t.Errorf("EvenInt16Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenInt8(t *testing.T) {
	r := EvenInt8(10)
	if !r {
		t.Errorf("EvenInt8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenInt8(1)
	if r {
		t.Errorf("EvenInt8 failed. Expected=false, actual=true")
	}

	var two int8 = 2
	rPtr := EvenInt8Ptr(&two)
	if !rPtr {
		t.Errorf("EvenInt8Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenUint(t *testing.T) {
	r := EvenUint(10)
	if !r {
		t.Errorf("EvenUint failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint(1)
	if r {
		t.Errorf("EvenUint failed. Expected=false, actual=true")
	}

	var two uint = 2
	rPtr := EvenUintPtr(&two)
	if !rPtr {
		t.Errorf("EvenUintPtr failed. Expected=true, actual=false")
	}
}

func TestEvenUint64(t *testing.T) {
	r := EvenUint64(10)
	if !r {
		t.Errorf("EvenUint64 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint64(1)
	if r {
		t.Errorf("EvenUint64 failed. Expected=false, actual=true")
	}

	var two uint64 = 2
	rPtr := EvenUint64Ptr(&two)
	if !rPtr {
		t.Errorf("EvenUint64Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenUint32(t *testing.T) {
	r := EvenUint32(10)
	if !r {
		t.Errorf("EvenUint32 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint32(1)
	if r {
		t.Errorf("EvenUint32 failed. Expected=false, actual=true")
	}

	var two uint32 = 2
	rPtr := EvenUint32Ptr(&two)
	if !rPtr {
		t.Errorf("EvenUint32Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenUint16(t *testing.T) {
	r := EvenUint16(10)
	if !r {
		t.Errorf("EvenUint16 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint16(1)
	if r {
		t.Errorf("EvenUint16 failed. Expected=false, actual=true")
	}

	var two uint16 = 2
	rPtr := EvenUint16Ptr(&two)
	if !rPtr {
		t.Errorf("EvenUint16Ptr failed. Expected=true, actual=false")
	}
}

func TestEvenUint8(t *testing.T) {
	r := EvenUint8(10)
	if !r {
		t.Errorf("EvenUint8 failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EvenUint8(1)
	if r {
		t.Errorf("EvenUint8 failed. Expected=false, actual=true")
	}

	var two uint8 = 2
	rPtr := EvenUint8Ptr(&two)
	if !rPtr {
		t.Errorf("EvenUint8Ptr failed. Expected=true, actual=false")
	}
}
