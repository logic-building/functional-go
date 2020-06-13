package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestReverseIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expected := []int{v3, v2, v1}
	reversed := ReverseInts([]int{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*int{&v3, &v2, &v1}
	reversedPtrList := ReverseIntsPtr([]*int{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expected := []int64{v3, v2, v1}
	reversed := ReverseInts64([]int64{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*int64{&v3, &v2, &v1}
	reversedPtrList := ReverseInts64Ptr([]*int64{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expected := []int32{v3, v2, v1}
	reversed := ReverseInts32([]int32{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*int32{&v3, &v2, &v1}
	reversedPtrList := ReverseInts32Ptr([]*int32{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expected := []int16{v3, v2, v1}
	reversed := ReverseInts16([]int16{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*int16{&v3, &v2, &v1}
	reversedPtrList := ReverseInts16Ptr([]*int16{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expected := []int8{v3, v2, v1}
	reversed := ReverseInts8([]int8{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*int8{&v3, &v2, &v1}
	reversedPtrList := ReverseInts8Ptr([]*int8{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expected := []uint{v3, v2, v1}
	reversed := ReverseUints([]uint{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*uint{&v3, &v2, &v1}
	reversedPtrList := ReverseUintsPtr([]*uint{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expected := []uint64{v3, v2, v1}
	reversed := ReverseUint64s([]uint64{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*uint64{&v3, &v2, &v1}
	reversedPtrList := ReverseUint64sPtr([]*uint64{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expected := []uint32{v3, v2, v1}
	reversed := ReverseUints32([]uint32{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*uint32{&v3, &v2, &v1}
	reversedPtrList := ReverseUints32Ptr([]*uint32{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expected := []uint16{v3, v2, v1}
	reversed := ReverseUints16([]uint16{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*uint16{&v3, &v2, &v1}
	reversedPtrList := ReverseUints16Ptr([]*uint16{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expected := []uint8{v3, v2, v1}
	reversed := ReverseUints8([]uint8{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*uint8{&v3, &v2, &v1}
	reversedPtrList := ReverseUints8Ptr([]*uint8{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expected := []string{v3, v2, v1}
	reversed := ReverseStrs([]string{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*string{&v3, &v2, &v1}
	reversedPtrList := ReverseStrsPtr([]*string{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expected := []bool{vt, vt, vf}
	reversed := ReverseBools([]bool{vf, vt, vt})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*bool{&vt, &vt, &vf}
	reversedPtrList := ReverseBoolsPtr([]*bool{&vf, &vt, &vt})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expected := []float32{v3, v2, v1}
	reversed := ReverseFloat32s([]float32{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*float32{&v3, &v2, &v1}
	reversedPtrList := ReverseFloat32sPtr([]*float32{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}
func TestReverseFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expected := []float64{v3, v2, v1}
	reversed := ReverseFloat64s([]float64{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*float64{&v3, &v2, &v1}
	reversedPtrList := ReverseFloat64sPtr([]*float64{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}