package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntMethodChain(t *testing.T) {
	expectedSquareList := []int{1, 4, 9}
	squareList := MakeIntSlice([]int{1, 2, 3}...).Map(squareInt)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapIntMthodChain failed")
	}

	if len(MakeIntSlice().Map(squareInt)) > 0 {
		t.Errorf("MapInt failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapInt64MethodChain(t *testing.T) {
	expectedSquareList := []int64{1, 4, 9}
	squareList := MakeInt64Slice([]int64{1, 2, 3}...).Map(squareInt64)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt64MthodChain failed")
	}

	if len(MakeInt64Slice().Map(squareInt64)) > 0 {
		t.Errorf("MapInt64 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapInt32MethodChain(t *testing.T) {
	expectedSquareList := []int32{1, 4, 9}
	squareList := MakeInt32Slice([]int32{1, 2, 3}...).Map(squareInt32)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt32MthodChain failed")
	}

	if len(MakeInt32Slice().Map(squareInt32)) > 0 {
		t.Errorf("MapInt32 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapInt16MethodChain(t *testing.T) {
	expectedSquareList := []int16{1, 4, 9}
	squareList := MakeInt16Slice([]int16{1, 2, 3}...).Map(squareInt16)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt16MthodChain failed")
	}

	if len(MakeInt16Slice().Map(squareInt16)) > 0 {
		t.Errorf("MapInt16 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapInt8MethodChain(t *testing.T) {
	expectedSquareList := []int8{1, 4, 9}
	squareList := MakeInt8Slice([]int8{1, 2, 3}...).Map(squareInt8)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt8MthodChain failed")
	}

	if len(MakeInt8Slice().Map(squareInt8)) > 0 {
		t.Errorf("MapInt8 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapUintMethodChain(t *testing.T) {
	expectedSquareList := []uint{1, 4, 9}
	squareList := MakeUintSlice([]uint{1, 2, 3}...).Map(squareUint)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUintMthodChain failed")
	}

	if len(MakeUintSlice().Map(squareUint)) > 0 {
		t.Errorf("MapUint failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapUint64MethodChain(t *testing.T) {
	expectedSquareList := []uint64{1, 4, 9}
	squareList := MakeUint64Slice([]uint64{1, 2, 3}...).Map(squareUint64)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint64MthodChain failed")
	}

	if len(MakeUint64Slice().Map(squareUint64)) > 0 {
		t.Errorf("MapUint64 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapUint32MethodChain(t *testing.T) {
	expectedSquareList := []uint32{1, 4, 9}
	squareList := MakeUint32Slice([]uint32{1, 2, 3}...).Map(squareUint32)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint32MthodChain failed")
	}

	if len(MakeUint32Slice().Map(squareUint32)) > 0 {
		t.Errorf("MapUint32 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapUint16MethodChain(t *testing.T) {
	expectedSquareList := []uint16{1, 4, 9}
	squareList := MakeUint16Slice([]uint16{1, 2, 3}...).Map(squareUint16)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint16MthodChain failed")
	}

	if len(MakeUint16Slice().Map(squareUint16)) > 0 {
		t.Errorf("MapUint16 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapUint8MethodChain(t *testing.T) {
	expectedSquareList := []uint8{1, 4, 9}
	squareList := MakeUint8Slice([]uint8{1, 2, 3}...).Map(squareUint8)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint8MthodChain failed")
	}

	if len(MakeUint8Slice().Map(squareUint8)) > 0 {
		t.Errorf("MapUint8 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapStrMethodChain(t *testing.T) {
	expectedSquareList := []string{"11", "22", "33"}
	squareList := MakeStrSlice([]string{"1", "2", "3"}...).Map(squareStr)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapStrMthodChain failed")
	}

	if len(MakeStrSlice().Map(squareStr)) > 0 {
		t.Errorf("MapStr failed.")
		t.Errorf(reflect.String.String())
	}
}

func squareStr(s string) string {
	return s+s
}

func TestMapBoolMethodChain(t *testing.T) {
	expectedSquareList := []bool{false, true, false}
	squareList := MakeBoolSlice([]bool{true, false, true}...).Map(inverseBool)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapBoolMthodChain failed")
	}

	if len(MakeBoolSlice().Map(inverseBool)) > 0 {
		t.Errorf("MapBool failed.")
		t.Errorf(reflect.String.String())
	}
}

func inverseBool(v bool) bool {
	if v == true {
		return false
	} 
	return true
}

func TestMapFloat32MethodChain(t *testing.T) {
	expectedSquareList := []float32{1, 4, 9}
	squareList := MakeFloat32Slice([]float32{1, 2, 3}...).Map(squareFloat32)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat32MthodChain failed")
	}

	if len(MakeFloat32Slice().Map(squareFloat32)) > 0 {
		t.Errorf("MapFloat32 failed.")
		t.Errorf(reflect.String.String())
	}
}

func TestMapFloat64MethodChain(t *testing.T) {
	expectedSquareList := []float64{1, 4, 9}
	squareList := MakeFloat64Slice([]float64{1, 2, 3}...).Map(squareFloat64)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat64MthodChain failed")
	}

	if len(MakeFloat64Slice().Map(squareFloat64)) > 0 {
		t.Errorf("MapFloat64 failed.")
		t.Errorf(reflect.String.String())
	}
}
