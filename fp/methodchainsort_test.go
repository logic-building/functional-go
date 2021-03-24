package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

// TestSortIntmethodchain
func TestSortIntmethodchain(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expected := []int{v1, v2, v3}
	sorted :=  MakeIntSlice([]int{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortIntmethodchainDesc
func TestSortIntmethodchainDesc(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expected := []int{v3, v2, v1}
	sorted :=  MakeIntSlice([]int{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortIntmethodchainPtr
func TestSortIntmethodchainPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expected := []*int{&v1, &v2, &v3}
	sorted :=  MakeIntSlicePtr([]*int{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortIntmethodchainDescPtr
func TestSortIntmethodchainDescPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	expected := []*int{&v3, &v2, &v1}
	sorted :=  MakeIntSlicePtr([]*int{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt64methodchain
func TestSortInt64methodchain(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expected := []int64{v1, v2, v3}
	sorted :=  MakeInt64Slice([]int64{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt64methodchainDesc
func TestSortInt64methodchainDesc(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expected := []int64{v3, v2, v1}
	sorted :=  MakeInt64Slice([]int64{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt64methodchainPtr
func TestSortInt64methodchainPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expected := []*int64{&v1, &v2, &v3}
	sorted :=  MakeInt64SlicePtr([]*int64{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt64methodchainDescPtr
func TestSortInt64methodchainDescPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	expected := []*int64{&v3, &v2, &v1}
	sorted :=  MakeInt64SlicePtr([]*int64{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt32methodchain
func TestSortInt32methodchain(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expected := []int32{v1, v2, v3}
	sorted :=  MakeInt32Slice([]int32{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt32methodchainDesc
func TestSortInt32methodchainDesc(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expected := []int32{v3, v2, v1}
	sorted :=  MakeInt32Slice([]int32{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt32methodchainPtr
func TestSortInt32methodchainPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expected := []*int32{&v1, &v2, &v3}
	sorted :=  MakeInt32SlicePtr([]*int32{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt32methodchainDescPtr
func TestSortInt32methodchainDescPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	expected := []*int32{&v3, &v2, &v1}
	sorted :=  MakeInt32SlicePtr([]*int32{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt16methodchain
func TestSortInt16methodchain(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expected := []int16{v1, v2, v3}
	sorted :=  MakeInt16Slice([]int16{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt16methodchainDesc
func TestSortInt16methodchainDesc(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expected := []int16{v3, v2, v1}
	sorted :=  MakeInt16Slice([]int16{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt16methodchainPtr
func TestSortInt16methodchainPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expected := []*int16{&v1, &v2, &v3}
	sorted :=  MakeInt16SlicePtr([]*int16{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt16methodchainDescPtr
func TestSortInt16methodchainDescPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	expected := []*int16{&v3, &v2, &v1}
	sorted :=  MakeInt16SlicePtr([]*int16{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt8methodchain
func TestSortInt8methodchain(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expected := []int8{v1, v2, v3}
	sorted :=  MakeInt8Slice([]int8{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt8methodchainDesc
func TestSortInt8methodchainDesc(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expected := []int8{v3, v2, v1}
	sorted :=  MakeInt8Slice([]int8{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortInt8methodchainPtr
func TestSortInt8methodchainPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expected := []*int8{&v1, &v2, &v3}
	sorted :=  MakeInt8SlicePtr([]*int8{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortInt8methodchainDescPtr
func TestSortInt8methodchainDescPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	expected := []*int8{&v3, &v2, &v1}
	sorted :=  MakeInt8SlicePtr([]*int8{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUintmethodchain
func TestSortUintmethodchain(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expected := []uint{v1, v2, v3}
	sorted :=  MakeUintSlice([]uint{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUintmethodchainDesc
func TestSortUintmethodchainDesc(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expected := []uint{v3, v2, v1}
	sorted :=  MakeUintSlice([]uint{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUintmethodchainPtr
func TestSortUintmethodchainPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expected := []*uint{&v1, &v2, &v3}
	sorted :=  MakeUintSlicePtr([]*uint{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUintmethodchainDescPtr
func TestSortUintmethodchainDescPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	expected := []*uint{&v3, &v2, &v1}
	sorted :=  MakeUintSlicePtr([]*uint{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint64methodchain
func TestSortUint64methodchain(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expected := []uint64{v1, v2, v3}
	sorted :=  MakeUint64Slice([]uint64{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint64methodchainDesc
func TestSortUint64methodchainDesc(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expected := []uint64{v3, v2, v1}
	sorted :=  MakeUint64Slice([]uint64{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint64methodchainPtr
func TestSortUint64methodchainPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expected := []*uint64{&v1, &v2, &v3}
	sorted :=  MakeUint64SlicePtr([]*uint64{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint64methodchainDescPtr
func TestSortUint64methodchainDescPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	expected := []*uint64{&v3, &v2, &v1}
	sorted :=  MakeUint64SlicePtr([]*uint64{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint32methodchain
func TestSortUint32methodchain(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expected := []uint32{v1, v2, v3}
	sorted :=  MakeUint32Slice([]uint32{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint32methodchainDesc
func TestSortUint32methodchainDesc(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expected := []uint32{v3, v2, v1}
	sorted :=  MakeUint32Slice([]uint32{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint32methodchainPtr
func TestSortUint32methodchainPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expected := []*uint32{&v1, &v2, &v3}
	sorted :=  MakeUint32SlicePtr([]*uint32{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint32methodchainDescPtr
func TestSortUint32methodchainDescPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	expected := []*uint32{&v3, &v2, &v1}
	sorted :=  MakeUint32SlicePtr([]*uint32{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint16methodchain
func TestSortUint16methodchain(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expected := []uint16{v1, v2, v3}
	sorted :=  MakeUint16Slice([]uint16{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint16methodchainDesc
func TestSortUint16methodchainDesc(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expected := []uint16{v3, v2, v1}
	sorted :=  MakeUint16Slice([]uint16{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint16methodchainPtr
func TestSortUint16methodchainPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expected := []*uint16{&v1, &v2, &v3}
	sorted :=  MakeUint16SlicePtr([]*uint16{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint16methodchainDescPtr
func TestSortUint16methodchainDescPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	expected := []*uint16{&v3, &v2, &v1}
	sorted :=  MakeUint16SlicePtr([]*uint16{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint8methodchain
func TestSortUint8methodchain(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expected := []uint8{v1, v2, v3}
	sorted :=  MakeUint8Slice([]uint8{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint8methodchainDesc
func TestSortUint8methodchainDesc(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expected := []uint8{v3, v2, v1}
	sorted :=  MakeUint8Slice([]uint8{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortUint8methodchainPtr
func TestSortUint8methodchainPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expected := []*uint8{&v1, &v2, &v3}
	sorted :=  MakeUint8SlicePtr([]*uint8{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortUint8methodchainDescPtr
func TestSortUint8methodchainDescPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	expected := []*uint8{&v3, &v2, &v1}
	sorted :=  MakeUint8SlicePtr([]*uint8{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortStrmethodchain
func TestSortStrmethodchain(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expected := []string{v1, v2, v3}
	sorted :=  MakeStrSlice([]string{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortStrmethodchainDesc
func TestSortStrmethodchainDesc(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expected := []string{v3, v2, v1}
	sorted :=  MakeStrSlice([]string{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortStrmethodchainPtr
func TestSortStrmethodchainPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expected := []*string{&v1, &v2, &v3}
	sorted :=  MakeStrSlicePtr([]*string{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortStrmethodchainDescPtr
func TestSortStrmethodchainDescPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	expected := []*string{&v3, &v2, &v1}
	sorted :=  MakeStrSlicePtr([]*string{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortFloat32methodchain
func TestSortFloat32methodchain(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expected := []float32{v1, v2, v3}
	sorted :=  MakeFloat32Slice([]float32{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortFloat32methodchainDesc
func TestSortFloat32methodchainDesc(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expected := []float32{v3, v2, v1}
	sorted :=  MakeFloat32Slice([]float32{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortFloat32methodchainPtr
func TestSortFloat32methodchainPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expected := []*float32{&v1, &v2, &v3}
	sorted :=  MakeFloat32SlicePtr([]*float32{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortFloat32methodchainDescPtr
func TestSortFloat32methodchainDescPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	expected := []*float32{&v3, &v2, &v1}
	sorted :=  MakeFloat32SlicePtr([]*float32{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortFloat64methodchain
func TestSortFloat64methodchain(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expected := []float64{v1, v2, v3}
	sorted :=  MakeFloat64Slice([]float64{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortFloat64methodchainDesc
func TestSortFloat64methodchainDesc(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expected := []float64{v3, v2, v1}
	sorted :=  MakeFloat64Slice([]float64{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSortFloat64methodchainPtr
func TestSortFloat64methodchainPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expected := []*float64{&v1, &v2, &v3}
	sorted :=  MakeFloat64SlicePtr([]*float64{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSortFloat64methodchainDescPtr
func TestSortFloat64methodchainDescPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	expected := []*float64{&v3, &v2, &v1}
	sorted :=  MakeFloat64SlicePtr([]*float64{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}
