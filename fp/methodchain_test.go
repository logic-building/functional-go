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

// TestMap2IntMethodChain - 
func TestMap2IntMethodChain(t *testing.T) {
	if len(MakeIntSlice().Map(nil)) > 0 {
		t.Errorf("MapInt failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapIntMethodChainPtr - 
func TestMapIntMethodChainPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v9 int = 9
	expectedSquareList := []*int{&v1, &v4, &v9}
	squareList := MakeIntSlicePtr([]*int{&v1, &v2, &v3}...).MapPtr(squareIntPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapIntPtrMthodChain failed")
	}

	if len(MakeIntSlicePtr().MapPtr(squareIntPtr)) > 0 {
		t.Errorf("MapIntPtr failed.")
	}
}

// TestMapPtr2IntMethodChain -  
func TestMapPtr2IntMethodChain(t *testing.T) {
	if len(MakeIntSlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapIntPtr failed.")
	}
}

// TestFilterIntMethodChain - 
func TestFilterIntMethodChain(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	greaterThan1IntMethodChain := func(num int) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []int{v2, v3}
	filteredList := MakeIntSlice([]int{v1, v2, v3}...).Filter(greaterThan1IntMethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterInt failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2IntMethodChain - 
func TestFilter2IntMethodChain(t *testing.T) {
	if len(MakeIntSlice().Filter(nil)) > 0 {
		t.Errorf("FilterIntPtr failed.")
	}
}

// TestFilterIntPtrMethodChain - 
func TestFilterIntPtrMethodChain(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3

	greaterThan1IntMethodChain := func(num *int) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*int{&v2, &v3}
	filteredList := MakeIntSlicePtr([]*int{&v1, &v2, &v3}...).FilterPtr(greaterThan1IntMethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterIntPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2IntPtrMethodChain - 
func TestFilter2IntPtrMethodChain(t *testing.T) {
	if len(MakeIntSlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterIntPtr failed.")
	}
}

// TestRemoveIntMethodChain - 
func TestRemoveIntMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	isGreaterThanThreeInt := func (num int) bool {
		return num > 3
	}

	expectedNewList := []int{v2, v3}
	NewList := MakeIntSlice([]int{v2, v3, v4}...).Remove(isGreaterThanThreeInt)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2IntMethodChain - 
func TestRemove2IntMethodChain(t *testing.T) {
	if len(MakeIntSlice().Remove(nil)) > 0 {
		t.Errorf("RemoveInt failed.")
	}
}

// TestRemoveIntPtrMethodChain - 
func TestRemoveIntPtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	isGreaterThanThreeInt := func (num *int) bool {
		return *num > 3
	}

	expectedNewList := []*int{&v2, &v3}
	NewList := MakeIntSlicePtr([]*int{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeInt)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveIntPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrIntMethodChain - 
func TestRemove2PtrIntMethodChain(t *testing.T) {
	if len(MakeIntSlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveIntPtr failed.")
	}
}

func TestDropWhileIntMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	isEvenInt := func(num int) bool {
		return num%2 == 0
	}

	expectedNewList := []int{v3, v4, v5}
	NewList := MakeIntSlice([]int{v4, v2, v3, v4, v5}...).DropWhile(isEvenInt)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2IntMethodChain - 
func TestDropWhile2IntMethodChain(t *testing.T) {
	if len(MakeIntSlice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileInt failed.")
	}
}

func TestDropWhileIntPtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5

	isEvenIntPtr := func(num *int) bool {
		return *num%2 == 0
	}

	expectedNewList := []*int{&v3, &v4, &v5}
	NewList := MakeIntSlicePtr([]*int{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenIntPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrIntMethodChain - 
func TestDropWhile2PtrIntMethodChain(t *testing.T) {
	if len(MakeIntSlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileIntPtr failed.")
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

// TestMap2Int64MethodChain - 
func TestMap2Int64MethodChain(t *testing.T) {
	if len(MakeInt64Slice().Map(nil)) > 0 {
		t.Errorf("MapInt64 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapInt64MethodChainPtr - 
func TestMapInt64MethodChainPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v9 int64 = 9
	expectedSquareList := []*int64{&v1, &v4, &v9}
	squareList := MakeInt64SlicePtr([]*int64{&v1, &v2, &v3}...).MapPtr(squareInt64Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapInt64PtrMthodChain failed")
	}

	if len(MakeInt64SlicePtr().MapPtr(squareInt64Ptr)) > 0 {
		t.Errorf("MapInt64Ptr failed.")
	}
}

// TestMapPtr2Int64MethodChain -  
func TestMapPtr2Int64MethodChain(t *testing.T) {
	if len(MakeInt64SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapInt64Ptr failed.")
	}
}

// TestFilterInt64MethodChain - 
func TestFilterInt64MethodChain(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	greaterThan1Int64MethodChain := func(num int64) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []int64{v2, v3}
	filteredList := MakeInt64Slice([]int64{v1, v2, v3}...).Filter(greaterThan1Int64MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterInt64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int64MethodChain - 
func TestFilter2Int64MethodChain(t *testing.T) {
	if len(MakeInt64Slice().Filter(nil)) > 0 {
		t.Errorf("FilterInt64Ptr failed.")
	}
}

// TestFilterInt64PtrMethodChain - 
func TestFilterInt64PtrMethodChain(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3

	greaterThan1Int64MethodChain := func(num *int64) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*int64{&v2, &v3}
	filteredList := MakeInt64SlicePtr([]*int64{&v1, &v2, &v3}...).FilterPtr(greaterThan1Int64MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterInt64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int64PtrMethodChain - 
func TestFilter2Int64PtrMethodChain(t *testing.T) {
	if len(MakeInt64SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterInt64Ptr failed.")
	}
}

// TestRemoveInt64MethodChain - 
func TestRemoveInt64MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	isGreaterThanThreeInt64 := func (num int64) bool {
		return num > 3
	}

	expectedNewList := []int64{v2, v3}
	NewList := MakeInt64Slice([]int64{v2, v3, v4}...).Remove(isGreaterThanThreeInt64)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Int64MethodChain - 
func TestRemove2Int64MethodChain(t *testing.T) {
	if len(MakeInt64Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveInt64 failed.")
	}
}

// TestRemoveInt64PtrMethodChain - 
func TestRemoveInt64PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4

	isGreaterThanThreeInt64 := func (num *int64) bool {
		return *num > 3
	}

	expectedNewList := []*int64{&v2, &v3}
	NewList := MakeInt64SlicePtr([]*int64{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeInt64)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrInt64MethodChain - 
func TestRemove2PtrInt64MethodChain(t *testing.T) {
	if len(MakeInt64SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveInt64Ptr failed.")
	}
}

func TestDropWhileInt64MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	isEvenInt64 := func(num int64) bool {
		return num%2 == 0
	}

	expectedNewList := []int64{v3, v4, v5}
	NewList := MakeInt64Slice([]int64{v4, v2, v3, v4, v5}...).DropWhile(isEvenInt64)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Int64MethodChain - 
func TestDropWhile2Int64MethodChain(t *testing.T) {
	if len(MakeInt64Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileInt64 failed.")
	}
}

func TestDropWhileInt64PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	isEvenInt64Ptr := func(num *int64) bool {
		return *num%2 == 0
	}

	expectedNewList := []*int64{&v3, &v4, &v5}
	NewList := MakeInt64SlicePtr([]*int64{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenInt64Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrInt64MethodChain - 
func TestDropWhile2PtrInt64MethodChain(t *testing.T) {
	if len(MakeInt64SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileInt64Ptr failed.")
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

// TestMap2Int32MethodChain - 
func TestMap2Int32MethodChain(t *testing.T) {
	if len(MakeInt32Slice().Map(nil)) > 0 {
		t.Errorf("MapInt32 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapInt32MethodChainPtr - 
func TestMapInt32MethodChainPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v9 int32 = 9
	expectedSquareList := []*int32{&v1, &v4, &v9}
	squareList := MakeInt32SlicePtr([]*int32{&v1, &v2, &v3}...).MapPtr(squareInt32Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapInt32PtrMthodChain failed")
	}

	if len(MakeInt32SlicePtr().MapPtr(squareInt32Ptr)) > 0 {
		t.Errorf("MapInt32Ptr failed.")
	}
}

// TestMapPtr2Int32MethodChain -  
func TestMapPtr2Int32MethodChain(t *testing.T) {
	if len(MakeInt32SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapInt32Ptr failed.")
	}
}

// TestFilterInt32MethodChain - 
func TestFilterInt32MethodChain(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	greaterThan1Int32MethodChain := func(num int32) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []int32{v2, v3}
	filteredList := MakeInt32Slice([]int32{v1, v2, v3}...).Filter(greaterThan1Int32MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterInt32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int32MethodChain - 
func TestFilter2Int32MethodChain(t *testing.T) {
	if len(MakeInt32Slice().Filter(nil)) > 0 {
		t.Errorf("FilterInt32Ptr failed.")
	}
}

// TestFilterInt32PtrMethodChain - 
func TestFilterInt32PtrMethodChain(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	greaterThan1Int32MethodChain := func(num *int32) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*int32{&v2, &v3}
	filteredList := MakeInt32SlicePtr([]*int32{&v1, &v2, &v3}...).FilterPtr(greaterThan1Int32MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterInt32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int32PtrMethodChain - 
func TestFilter2Int32PtrMethodChain(t *testing.T) {
	if len(MakeInt32SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterInt32Ptr failed.")
	}
}

// TestRemoveInt32MethodChain - 
func TestRemoveInt32MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	isGreaterThanThreeInt32 := func (num int32) bool {
		return num > 3
	}

	expectedNewList := []int32{v2, v3}
	NewList := MakeInt32Slice([]int32{v2, v3, v4}...).Remove(isGreaterThanThreeInt32)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Int32MethodChain - 
func TestRemove2Int32MethodChain(t *testing.T) {
	if len(MakeInt32Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveInt32 failed.")
	}
}

// TestRemoveInt32PtrMethodChain - 
func TestRemoveInt32PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4

	isGreaterThanThreeInt32 := func (num *int32) bool {
		return *num > 3
	}

	expectedNewList := []*int32{&v2, &v3}
	NewList := MakeInt32SlicePtr([]*int32{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeInt32)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrInt32MethodChain - 
func TestRemove2PtrInt32MethodChain(t *testing.T) {
	if len(MakeInt32SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveInt32Ptr failed.")
	}
}

func TestDropWhileInt32MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	isEvenInt32 := func(num int32) bool {
		return num%2 == 0
	}

	expectedNewList := []int32{v3, v4, v5}
	NewList := MakeInt32Slice([]int32{v4, v2, v3, v4, v5}...).DropWhile(isEvenInt32)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Int32MethodChain - 
func TestDropWhile2Int32MethodChain(t *testing.T) {
	if len(MakeInt32Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileInt32 failed.")
	}
}

func TestDropWhileInt32PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	isEvenInt32Ptr := func(num *int32) bool {
		return *num%2 == 0
	}

	expectedNewList := []*int32{&v3, &v4, &v5}
	NewList := MakeInt32SlicePtr([]*int32{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenInt32Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrInt32MethodChain - 
func TestDropWhile2PtrInt32MethodChain(t *testing.T) {
	if len(MakeInt32SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileInt32Ptr failed.")
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

// TestMap2Int16MethodChain - 
func TestMap2Int16MethodChain(t *testing.T) {
	if len(MakeInt16Slice().Map(nil)) > 0 {
		t.Errorf("MapInt16 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapInt16MethodChainPtr - 
func TestMapInt16MethodChainPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v9 int16 = 9
	expectedSquareList := []*int16{&v1, &v4, &v9}
	squareList := MakeInt16SlicePtr([]*int16{&v1, &v2, &v3}...).MapPtr(squareInt16Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapInt16PtrMthodChain failed")
	}

	if len(MakeInt16SlicePtr().MapPtr(squareInt16Ptr)) > 0 {
		t.Errorf("MapInt16Ptr failed.")
	}
}

// TestMapPtr2Int16MethodChain -  
func TestMapPtr2Int16MethodChain(t *testing.T) {
	if len(MakeInt16SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapInt16Ptr failed.")
	}
}

// TestFilterInt16MethodChain - 
func TestFilterInt16MethodChain(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	greaterThan1Int16MethodChain := func(num int16) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []int16{v2, v3}
	filteredList := MakeInt16Slice([]int16{v1, v2, v3}...).Filter(greaterThan1Int16MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterInt16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int16MethodChain - 
func TestFilter2Int16MethodChain(t *testing.T) {
	if len(MakeInt16Slice().Filter(nil)) > 0 {
		t.Errorf("FilterInt16Ptr failed.")
	}
}

// TestFilterInt16PtrMethodChain - 
func TestFilterInt16PtrMethodChain(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3

	greaterThan1Int16MethodChain := func(num *int16) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*int16{&v2, &v3}
	filteredList := MakeInt16SlicePtr([]*int16{&v1, &v2, &v3}...).FilterPtr(greaterThan1Int16MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterInt16Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int16PtrMethodChain - 
func TestFilter2Int16PtrMethodChain(t *testing.T) {
	if len(MakeInt16SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterInt16Ptr failed.")
	}
}

// TestRemoveInt16MethodChain - 
func TestRemoveInt16MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	isGreaterThanThreeInt16 := func (num int16) bool {
		return num > 3
	}

	expectedNewList := []int16{v2, v3}
	NewList := MakeInt16Slice([]int16{v2, v3, v4}...).Remove(isGreaterThanThreeInt16)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Int16MethodChain - 
func TestRemove2Int16MethodChain(t *testing.T) {
	if len(MakeInt16Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveInt16 failed.")
	}
}

// TestRemoveInt16PtrMethodChain - 
func TestRemoveInt16PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4

	isGreaterThanThreeInt16 := func (num *int16) bool {
		return *num > 3
	}

	expectedNewList := []*int16{&v2, &v3}
	NewList := MakeInt16SlicePtr([]*int16{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeInt16)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrInt16MethodChain - 
func TestRemove2PtrInt16MethodChain(t *testing.T) {
	if len(MakeInt16SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveInt16Ptr failed.")
	}
}

func TestDropWhileInt16MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	isEvenInt16 := func(num int16) bool {
		return num%2 == 0
	}

	expectedNewList := []int16{v3, v4, v5}
	NewList := MakeInt16Slice([]int16{v4, v2, v3, v4, v5}...).DropWhile(isEvenInt16)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Int16MethodChain - 
func TestDropWhile2Int16MethodChain(t *testing.T) {
	if len(MakeInt16Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileInt16 failed.")
	}
}

func TestDropWhileInt16PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	isEvenInt16Ptr := func(num *int16) bool {
		return *num%2 == 0
	}

	expectedNewList := []*int16{&v3, &v4, &v5}
	NewList := MakeInt16SlicePtr([]*int16{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenInt16Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrInt16MethodChain - 
func TestDropWhile2PtrInt16MethodChain(t *testing.T) {
	if len(MakeInt16SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileInt16Ptr failed.")
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

// TestMap2Int8MethodChain - 
func TestMap2Int8MethodChain(t *testing.T) {
	if len(MakeInt8Slice().Map(nil)) > 0 {
		t.Errorf("MapInt8 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapInt8MethodChainPtr - 
func TestMapInt8MethodChainPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v9 int8 = 9
	expectedSquareList := []*int8{&v1, &v4, &v9}
	squareList := MakeInt8SlicePtr([]*int8{&v1, &v2, &v3}...).MapPtr(squareInt8Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapInt8PtrMthodChain failed")
	}

	if len(MakeInt8SlicePtr().MapPtr(squareInt8Ptr)) > 0 {
		t.Errorf("MapInt8Ptr failed.")
	}
}

// TestMapPtr2Int8MethodChain -  
func TestMapPtr2Int8MethodChain(t *testing.T) {
	if len(MakeInt8SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapInt8Ptr failed.")
	}
}

// TestFilterInt8MethodChain - 
func TestFilterInt8MethodChain(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	greaterThan1Int8MethodChain := func(num int8) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []int8{v2, v3}
	filteredList := MakeInt8Slice([]int8{v1, v2, v3}...).Filter(greaterThan1Int8MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterInt8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int8MethodChain - 
func TestFilter2Int8MethodChain(t *testing.T) {
	if len(MakeInt8Slice().Filter(nil)) > 0 {
		t.Errorf("FilterInt8Ptr failed.")
	}
}

// TestFilterInt8PtrMethodChain - 
func TestFilterInt8PtrMethodChain(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3

	greaterThan1Int8MethodChain := func(num *int8) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*int8{&v2, &v3}
	filteredList := MakeInt8SlicePtr([]*int8{&v1, &v2, &v3}...).FilterPtr(greaterThan1Int8MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterInt8Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Int8PtrMethodChain - 
func TestFilter2Int8PtrMethodChain(t *testing.T) {
	if len(MakeInt8SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterInt8Ptr failed.")
	}
}

// TestRemoveInt8MethodChain - 
func TestRemoveInt8MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	isGreaterThanThreeInt8 := func (num int8) bool {
		return num > 3
	}

	expectedNewList := []int8{v2, v3}
	NewList := MakeInt8Slice([]int8{v2, v3, v4}...).Remove(isGreaterThanThreeInt8)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Int8MethodChain - 
func TestRemove2Int8MethodChain(t *testing.T) {
	if len(MakeInt8Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveInt8 failed.")
	}
}

// TestRemoveInt8PtrMethodChain - 
func TestRemoveInt8PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4

	isGreaterThanThreeInt8 := func (num *int8) bool {
		return *num > 3
	}

	expectedNewList := []*int8{&v2, &v3}
	NewList := MakeInt8SlicePtr([]*int8{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeInt8)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveInt8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrInt8MethodChain - 
func TestRemove2PtrInt8MethodChain(t *testing.T) {
	if len(MakeInt8SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveInt8Ptr failed.")
	}
}

func TestDropWhileInt8MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	isEvenInt8 := func(num int8) bool {
		return num%2 == 0
	}

	expectedNewList := []int8{v3, v4, v5}
	NewList := MakeInt8Slice([]int8{v4, v2, v3, v4, v5}...).DropWhile(isEvenInt8)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Int8MethodChain - 
func TestDropWhile2Int8MethodChain(t *testing.T) {
	if len(MakeInt8Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileInt8 failed.")
	}
}

func TestDropWhileInt8PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	isEvenInt8Ptr := func(num *int8) bool {
		return *num%2 == 0
	}

	expectedNewList := []*int8{&v3, &v4, &v5}
	NewList := MakeInt8SlicePtr([]*int8{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenInt8Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrInt8MethodChain - 
func TestDropWhile2PtrInt8MethodChain(t *testing.T) {
	if len(MakeInt8SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileInt8Ptr failed.")
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

// TestMap2UintMethodChain - 
func TestMap2UintMethodChain(t *testing.T) {
	if len(MakeUintSlice().Map(nil)) > 0 {
		t.Errorf("MapUint failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapUintMethodChainPtr - 
func TestMapUintMethodChainPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v9 uint = 9
	expectedSquareList := []*uint{&v1, &v4, &v9}
	squareList := MakeUintSlicePtr([]*uint{&v1, &v2, &v3}...).MapPtr(squareUintPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapUintPtrMthodChain failed")
	}

	if len(MakeUintSlicePtr().MapPtr(squareUintPtr)) > 0 {
		t.Errorf("MapUintPtr failed.")
	}
}

// TestMapPtr2UintMethodChain -  
func TestMapPtr2UintMethodChain(t *testing.T) {
	if len(MakeUintSlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapUintPtr failed.")
	}
}

// TestFilterUintMethodChain - 
func TestFilterUintMethodChain(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	greaterThan1UintMethodChain := func(num uint) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []uint{v2, v3}
	filteredList := MakeUintSlice([]uint{v1, v2, v3}...).Filter(greaterThan1UintMethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterUint failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2UintMethodChain - 
func TestFilter2UintMethodChain(t *testing.T) {
	if len(MakeUintSlice().Filter(nil)) > 0 {
		t.Errorf("FilterUintPtr failed.")
	}
}

// TestFilterUintPtrMethodChain - 
func TestFilterUintPtrMethodChain(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3

	greaterThan1UintMethodChain := func(num *uint) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*uint{&v2, &v3}
	filteredList := MakeUintSlicePtr([]*uint{&v1, &v2, &v3}...).FilterPtr(greaterThan1UintMethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterUintPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2UintPtrMethodChain - 
func TestFilter2UintPtrMethodChain(t *testing.T) {
	if len(MakeUintSlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterUintPtr failed.")
	}
}

// TestRemoveUintMethodChain - 
func TestRemoveUintMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	isGreaterThanThreeUint := func (num uint) bool {
		return num > 3
	}

	expectedNewList := []uint{v2, v3}
	NewList := MakeUintSlice([]uint{v2, v3, v4}...).Remove(isGreaterThanThreeUint)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2UintMethodChain - 
func TestRemove2UintMethodChain(t *testing.T) {
	if len(MakeUintSlice().Remove(nil)) > 0 {
		t.Errorf("RemoveUint failed.")
	}
}

// TestRemoveUintPtrMethodChain - 
func TestRemoveUintPtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	isGreaterThanThreeUint := func (num *uint) bool {
		return *num > 3
	}

	expectedNewList := []*uint{&v2, &v3}
	NewList := MakeUintSlicePtr([]*uint{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeUint)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUintPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrUintMethodChain - 
func TestRemove2PtrUintMethodChain(t *testing.T) {
	if len(MakeUintSlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveUintPtr failed.")
	}
}

func TestDropWhileUintMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	isEvenUint := func(num uint) bool {
		return num%2 == 0
	}

	expectedNewList := []uint{v3, v4, v5}
	NewList := MakeUintSlice([]uint{v4, v2, v3, v4, v5}...).DropWhile(isEvenUint)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2UintMethodChain - 
func TestDropWhile2UintMethodChain(t *testing.T) {
	if len(MakeUintSlice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileUint failed.")
	}
}

func TestDropWhileUintPtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	isEvenUintPtr := func(num *uint) bool {
		return *num%2 == 0
	}

	expectedNewList := []*uint{&v3, &v4, &v5}
	NewList := MakeUintSlicePtr([]*uint{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenUintPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrUintMethodChain - 
func TestDropWhile2PtrUintMethodChain(t *testing.T) {
	if len(MakeUintSlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileUintPtr failed.")
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

// TestMap2Uint64MethodChain - 
func TestMap2Uint64MethodChain(t *testing.T) {
	if len(MakeUint64Slice().Map(nil)) > 0 {
		t.Errorf("MapUint64 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapUint64MethodChainPtr - 
func TestMapUint64MethodChainPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v9 uint64 = 9
	expectedSquareList := []*uint64{&v1, &v4, &v9}
	squareList := MakeUint64SlicePtr([]*uint64{&v1, &v2, &v3}...).MapPtr(squareUint64Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapUint64PtrMthodChain failed")
	}

	if len(MakeUint64SlicePtr().MapPtr(squareUint64Ptr)) > 0 {
		t.Errorf("MapUint64Ptr failed.")
	}
}

// TestMapPtr2Uint64MethodChain -  
func TestMapPtr2Uint64MethodChain(t *testing.T) {
	if len(MakeUint64SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapUint64Ptr failed.")
	}
}

// TestFilterUint64MethodChain - 
func TestFilterUint64MethodChain(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	greaterThan1Uint64MethodChain := func(num uint64) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []uint64{v2, v3}
	filteredList := MakeUint64Slice([]uint64{v1, v2, v3}...).Filter(greaterThan1Uint64MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterUint64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint64MethodChain - 
func TestFilter2Uint64MethodChain(t *testing.T) {
	if len(MakeUint64Slice().Filter(nil)) > 0 {
		t.Errorf("FilterUint64Ptr failed.")
	}
}

// TestFilterUint64PtrMethodChain - 
func TestFilterUint64PtrMethodChain(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3

	greaterThan1Uint64MethodChain := func(num *uint64) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*uint64{&v2, &v3}
	filteredList := MakeUint64SlicePtr([]*uint64{&v1, &v2, &v3}...).FilterPtr(greaterThan1Uint64MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterUint64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint64PtrMethodChain - 
func TestFilter2Uint64PtrMethodChain(t *testing.T) {
	if len(MakeUint64SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterUint64Ptr failed.")
	}
}

// TestRemoveUint64MethodChain - 
func TestRemoveUint64MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	isGreaterThanThreeUint64 := func (num uint64) bool {
		return num > 3
	}

	expectedNewList := []uint64{v2, v3}
	NewList := MakeUint64Slice([]uint64{v2, v3, v4}...).Remove(isGreaterThanThreeUint64)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Uint64MethodChain - 
func TestRemove2Uint64MethodChain(t *testing.T) {
	if len(MakeUint64Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveUint64 failed.")
	}
}

// TestRemoveUint64PtrMethodChain - 
func TestRemoveUint64PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4

	isGreaterThanThreeUint64 := func (num *uint64) bool {
		return *num > 3
	}

	expectedNewList := []*uint64{&v2, &v3}
	NewList := MakeUint64SlicePtr([]*uint64{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeUint64)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrUint64MethodChain - 
func TestRemove2PtrUint64MethodChain(t *testing.T) {
	if len(MakeUint64SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveUint64Ptr failed.")
	}
}

func TestDropWhileUint64MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	isEvenUint64 := func(num uint64) bool {
		return num%2 == 0
	}

	expectedNewList := []uint64{v3, v4, v5}
	NewList := MakeUint64Slice([]uint64{v4, v2, v3, v4, v5}...).DropWhile(isEvenUint64)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Uint64MethodChain - 
func TestDropWhile2Uint64MethodChain(t *testing.T) {
	if len(MakeUint64Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileUint64 failed.")
	}
}

func TestDropWhileUint64PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	isEvenUint64Ptr := func(num *uint64) bool {
		return *num%2 == 0
	}

	expectedNewList := []*uint64{&v3, &v4, &v5}
	NewList := MakeUint64SlicePtr([]*uint64{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenUint64Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrUint64MethodChain - 
func TestDropWhile2PtrUint64MethodChain(t *testing.T) {
	if len(MakeUint64SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileUint64Ptr failed.")
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

// TestMap2Uint32MethodChain - 
func TestMap2Uint32MethodChain(t *testing.T) {
	if len(MakeUint32Slice().Map(nil)) > 0 {
		t.Errorf("MapUint32 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapUint32MethodChainPtr - 
func TestMapUint32MethodChainPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v9 uint32 = 9
	expectedSquareList := []*uint32{&v1, &v4, &v9}
	squareList := MakeUint32SlicePtr([]*uint32{&v1, &v2, &v3}...).MapPtr(squareUint32Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapUint32PtrMthodChain failed")
	}

	if len(MakeUint32SlicePtr().MapPtr(squareUint32Ptr)) > 0 {
		t.Errorf("MapUint32Ptr failed.")
	}
}

// TestMapPtr2Uint32MethodChain -  
func TestMapPtr2Uint32MethodChain(t *testing.T) {
	if len(MakeUint32SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapUint32Ptr failed.")
	}
}

// TestFilterUint32MethodChain - 
func TestFilterUint32MethodChain(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	greaterThan1Uint32MethodChain := func(num uint32) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []uint32{v2, v3}
	filteredList := MakeUint32Slice([]uint32{v1, v2, v3}...).Filter(greaterThan1Uint32MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterUint32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint32MethodChain - 
func TestFilter2Uint32MethodChain(t *testing.T) {
	if len(MakeUint32Slice().Filter(nil)) > 0 {
		t.Errorf("FilterUint32Ptr failed.")
	}
}

// TestFilterUint32PtrMethodChain - 
func TestFilterUint32PtrMethodChain(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3

	greaterThan1Uint32MethodChain := func(num *uint32) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*uint32{&v2, &v3}
	filteredList := MakeUint32SlicePtr([]*uint32{&v1, &v2, &v3}...).FilterPtr(greaterThan1Uint32MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterUint32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint32PtrMethodChain - 
func TestFilter2Uint32PtrMethodChain(t *testing.T) {
	if len(MakeUint32SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterUint32Ptr failed.")
	}
}

// TestRemoveUint32MethodChain - 
func TestRemoveUint32MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	isGreaterThanThreeUint32 := func (num uint32) bool {
		return num > 3
	}

	expectedNewList := []uint32{v2, v3}
	NewList := MakeUint32Slice([]uint32{v2, v3, v4}...).Remove(isGreaterThanThreeUint32)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Uint32MethodChain - 
func TestRemove2Uint32MethodChain(t *testing.T) {
	if len(MakeUint32Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveUint32 failed.")
	}
}

// TestRemoveUint32PtrMethodChain - 
func TestRemoveUint32PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4

	isGreaterThanThreeUint32 := func (num *uint32) bool {
		return *num > 3
	}

	expectedNewList := []*uint32{&v2, &v3}
	NewList := MakeUint32SlicePtr([]*uint32{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeUint32)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrUint32MethodChain - 
func TestRemove2PtrUint32MethodChain(t *testing.T) {
	if len(MakeUint32SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveUint32Ptr failed.")
	}
}

func TestDropWhileUint32MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	isEvenUint32 := func(num uint32) bool {
		return num%2 == 0
	}

	expectedNewList := []uint32{v3, v4, v5}
	NewList := MakeUint32Slice([]uint32{v4, v2, v3, v4, v5}...).DropWhile(isEvenUint32)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Uint32MethodChain - 
func TestDropWhile2Uint32MethodChain(t *testing.T) {
	if len(MakeUint32Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileUint32 failed.")
	}
}

func TestDropWhileUint32PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	isEvenUint32Ptr := func(num *uint32) bool {
		return *num%2 == 0
	}

	expectedNewList := []*uint32{&v3, &v4, &v5}
	NewList := MakeUint32SlicePtr([]*uint32{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenUint32Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrUint32MethodChain - 
func TestDropWhile2PtrUint32MethodChain(t *testing.T) {
	if len(MakeUint32SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileUint32Ptr failed.")
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

// TestMap2Uint16MethodChain - 
func TestMap2Uint16MethodChain(t *testing.T) {
	if len(MakeUint16Slice().Map(nil)) > 0 {
		t.Errorf("MapUint16 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapUint16MethodChainPtr - 
func TestMapUint16MethodChainPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v9 uint16 = 9
	expectedSquareList := []*uint16{&v1, &v4, &v9}
	squareList := MakeUint16SlicePtr([]*uint16{&v1, &v2, &v3}...).MapPtr(squareUint16Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapUint16PtrMthodChain failed")
	}

	if len(MakeUint16SlicePtr().MapPtr(squareUint16Ptr)) > 0 {
		t.Errorf("MapUint16Ptr failed.")
	}
}

// TestMapPtr2Uint16MethodChain -  
func TestMapPtr2Uint16MethodChain(t *testing.T) {
	if len(MakeUint16SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapUint16Ptr failed.")
	}
}

// TestFilterUint16MethodChain - 
func TestFilterUint16MethodChain(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	greaterThan1Uint16MethodChain := func(num uint16) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []uint16{v2, v3}
	filteredList := MakeUint16Slice([]uint16{v1, v2, v3}...).Filter(greaterThan1Uint16MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterUint16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint16MethodChain - 
func TestFilter2Uint16MethodChain(t *testing.T) {
	if len(MakeUint16Slice().Filter(nil)) > 0 {
		t.Errorf("FilterUint16Ptr failed.")
	}
}

// TestFilterUint16PtrMethodChain - 
func TestFilterUint16PtrMethodChain(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3

	greaterThan1Uint16MethodChain := func(num *uint16) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*uint16{&v2, &v3}
	filteredList := MakeUint16SlicePtr([]*uint16{&v1, &v2, &v3}...).FilterPtr(greaterThan1Uint16MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterUint16Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint16PtrMethodChain - 
func TestFilter2Uint16PtrMethodChain(t *testing.T) {
	if len(MakeUint16SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterUint16Ptr failed.")
	}
}

// TestRemoveUint16MethodChain - 
func TestRemoveUint16MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	isGreaterThanThreeUint16 := func (num uint16) bool {
		return num > 3
	}

	expectedNewList := []uint16{v2, v3}
	NewList := MakeUint16Slice([]uint16{v2, v3, v4}...).Remove(isGreaterThanThreeUint16)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Uint16MethodChain - 
func TestRemove2Uint16MethodChain(t *testing.T) {
	if len(MakeUint16Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveUint16 failed.")
	}
}

// TestRemoveUint16PtrMethodChain - 
func TestRemoveUint16PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4

	isGreaterThanThreeUint16 := func (num *uint16) bool {
		return *num > 3
	}

	expectedNewList := []*uint16{&v2, &v3}
	NewList := MakeUint16SlicePtr([]*uint16{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeUint16)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint16Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrUint16MethodChain - 
func TestRemove2PtrUint16MethodChain(t *testing.T) {
	if len(MakeUint16SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveUint16Ptr failed.")
	}
}

func TestDropWhileUint16MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	isEvenUint16 := func(num uint16) bool {
		return num%2 == 0
	}

	expectedNewList := []uint16{v3, v4, v5}
	NewList := MakeUint16Slice([]uint16{v4, v2, v3, v4, v5}...).DropWhile(isEvenUint16)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Uint16MethodChain - 
func TestDropWhile2Uint16MethodChain(t *testing.T) {
	if len(MakeUint16Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileUint16 failed.")
	}
}

func TestDropWhileUint16PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	isEvenUint16Ptr := func(num *uint16) bool {
		return *num%2 == 0
	}

	expectedNewList := []*uint16{&v3, &v4, &v5}
	NewList := MakeUint16SlicePtr([]*uint16{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenUint16Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrUint16MethodChain - 
func TestDropWhile2PtrUint16MethodChain(t *testing.T) {
	if len(MakeUint16SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileUint16Ptr failed.")
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

// TestMap2Uint8MethodChain - 
func TestMap2Uint8MethodChain(t *testing.T) {
	if len(MakeUint8Slice().Map(nil)) > 0 {
		t.Errorf("MapUint8 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapUint8MethodChainPtr - 
func TestMapUint8MethodChainPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v9 uint8 = 9
	expectedSquareList := []*uint8{&v1, &v4, &v9}
	squareList := MakeUint8SlicePtr([]*uint8{&v1, &v2, &v3}...).MapPtr(squareUint8Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapUint8PtrMthodChain failed")
	}

	if len(MakeUint8SlicePtr().MapPtr(squareUint8Ptr)) > 0 {
		t.Errorf("MapUint8Ptr failed.")
	}
}

// TestMapPtr2Uint8MethodChain -  
func TestMapPtr2Uint8MethodChain(t *testing.T) {
	if len(MakeUint8SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapUint8Ptr failed.")
	}
}

// TestFilterUint8MethodChain - 
func TestFilterUint8MethodChain(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	greaterThan1Uint8MethodChain := func(num uint8) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []uint8{v2, v3}
	filteredList := MakeUint8Slice([]uint8{v1, v2, v3}...).Filter(greaterThan1Uint8MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterUint8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint8MethodChain - 
func TestFilter2Uint8MethodChain(t *testing.T) {
	if len(MakeUint8Slice().Filter(nil)) > 0 {
		t.Errorf("FilterUint8Ptr failed.")
	}
}

// TestFilterUint8PtrMethodChain - 
func TestFilterUint8PtrMethodChain(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3

	greaterThan1Uint8MethodChain := func(num *uint8) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*uint8{&v2, &v3}
	filteredList := MakeUint8SlicePtr([]*uint8{&v1, &v2, &v3}...).FilterPtr(greaterThan1Uint8MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterUint8Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Uint8PtrMethodChain - 
func TestFilter2Uint8PtrMethodChain(t *testing.T) {
	if len(MakeUint8SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterUint8Ptr failed.")
	}
}

// TestRemoveUint8MethodChain - 
func TestRemoveUint8MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	isGreaterThanThreeUint8 := func (num uint8) bool {
		return num > 3
	}

	expectedNewList := []uint8{v2, v3}
	NewList := MakeUint8Slice([]uint8{v2, v3, v4}...).Remove(isGreaterThanThreeUint8)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Uint8MethodChain - 
func TestRemove2Uint8MethodChain(t *testing.T) {
	if len(MakeUint8Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveUint8 failed.")
	}
}

// TestRemoveUint8PtrMethodChain - 
func TestRemoveUint8PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4

	isGreaterThanThreeUint8 := func (num *uint8) bool {
		return *num > 3
	}

	expectedNewList := []*uint8{&v2, &v3}
	NewList := MakeUint8SlicePtr([]*uint8{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeUint8)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveUint8Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrUint8MethodChain - 
func TestRemove2PtrUint8MethodChain(t *testing.T) {
	if len(MakeUint8SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveUint8Ptr failed.")
	}
}

func TestDropWhileUint8MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	isEvenUint8 := func(num uint8) bool {
		return num%2 == 0
	}

	expectedNewList := []uint8{v3, v4, v5}
	NewList := MakeUint8Slice([]uint8{v4, v2, v3, v4, v5}...).DropWhile(isEvenUint8)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Uint8MethodChain - 
func TestDropWhile2Uint8MethodChain(t *testing.T) {
	if len(MakeUint8Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileUint8 failed.")
	}
}

func TestDropWhileUint8PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	isEvenUint8Ptr := func(num *uint8) bool {
		return *num%2 == 0
	}

	expectedNewList := []*uint8{&v3, &v4, &v5}
	NewList := MakeUint8SlicePtr([]*uint8{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenUint8Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrUint8MethodChain - 
func TestDropWhile2PtrUint8MethodChain(t *testing.T) {
	if len(MakeUint8SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileUint8Ptr failed.")
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

// TestMap2StrMethodChain - 
func TestMap2StrMethodChain(t *testing.T) {
	if len(MakeStrSlice().Map(nil)) > 0 {
		t.Errorf("MapStr failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapStrMethodChainPtr - 
func TestMapStrMethodChainPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "11"
	var v22 string = "22"
	var v33 string = "33"
	expectedSquareList := []*string{&v11, &v22, &v33}
	squareList := MakeStrSlicePtr([]*string{&v1, &v2, &v3}...).MapPtr(squareStrPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapStrPtrMthodChain failed")
	}

	if len(MakeStrSlicePtr().MapPtr(squareStrPtr)) > 0 {
		t.Errorf("MapStrPtr failed.")
	}
}

// TestMapPtr2StrMethodChain -  
func TestMapPtr2StrMethodChain(t *testing.T) {
	if len(MakeStrSlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapStrPtr failed.")
	}
}

// TestFilterStrMethodChain - 
func TestFilterStrMethodChain(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	greaterThan1StrMethodChain := func(num string) bool {
		return num > "1"
	}

	// Test : even number in the list
	expectedFilteredList := []string{v2, v3}
	filteredList := MakeStrSlice([]string{v1, v2, v3}...).Filter(greaterThan1StrMethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterStr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2StrMethodChain - 
func TestFilter2StrMethodChain(t *testing.T) {
	if len(MakeStrSlice().Filter(nil)) > 0 {
		t.Errorf("FilterStrPtr failed.")
	}
}

// TestFilterStrPtrMethodChain - 
func TestFilterStrPtrMethodChain(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	greaterThan1StrMethodChain := func(num *string) bool {
		return *num > "1"
	}

	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v3}
	filteredList := MakeStrSlicePtr([]*string{&v1, &v2, &v3}...).FilterPtr(greaterThan1StrMethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2StrPtrMethodChain - 
func TestFilter2StrPtrMethodChain(t *testing.T) {
	if len(MakeStrSlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterStrPtr failed.")
	}
}

// TestRemoveStrMethodChain - 
func TestRemoveStrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	isGreaterThanThreeStr := func (num string) bool {
		return num > "3"
	}

	expectedNewList := []string{v2, v3}
	NewList := MakeStrSlice([]string{v2, v3, v4}...).Remove(isGreaterThanThreeStr)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2StrMethodChain - 
func TestRemove2StrMethodChain(t *testing.T) {
	if len(MakeStrSlice().Remove(nil)) > 0 {
		t.Errorf("RemoveStr failed.")
	}
}

// TestRemoveStrPtrMethodChain - 
func TestRemoveStrPtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	isGreaterThanThreeStr := func (num *string) bool {
		return *num > "3"
	}

	expectedNewList := []*string{&v2, &v3}
	NewList := MakeStrSlicePtr([]*string{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeStr)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrStrMethodChain - 
func TestRemove2PtrStrMethodChain(t *testing.T) {
	if len(MakeStrSlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveStrPtr failed.")
	}
}

func TestDropWhileStrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	isEvenStr := func(num string) bool {
		return num == "2" || num == "4"
	}

	expectedNewList := []string{v3, v4, v5}
	NewList := MakeStrSlice([]string{v4, v2, v3, v4, v5}...).DropWhile(isEvenStr)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2StrMethodChain - 
func TestDropWhile2StrMethodChain(t *testing.T) {
	if len(MakeStrSlice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileStr failed.")
	}
}

func TestDropWhileStrPtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	isEvenStrPtr := func(num *string) bool {
		return *num == "2" || *num == "4"
	}

	expectedNewList := []*string{&v3, &v4, &v5}
	NewList := MakeStrSlicePtr([]*string{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenStrPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrStrMethodChain - 
func TestDropWhile2PtrStrMethodChain(t *testing.T) {
	if len(MakeStrSlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}
}

// TestMapBoolMethodChain - 
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

	if len(MakeBoolSlice().Map(nil)) > 0 {
		t.Errorf("MapBool failed.")
	}
}

func inverseBool(v bool) bool {
	if v == true {
		return false
	} 
	return true
}

// TestMapPtrMethodChainBool - 
func TestMapPtrMethodChainBoolBoolMethodChain(t *testing.T) {
	tr := true
	f := false
	expectedSquareList := []*bool{&f, &tr, &f}
	squareList := MakeBoolSlicePtr([]*bool{&tr, &f, &tr}...).MapPtr(inverseBoolPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapBoolPtrMthodChain failed")
	}

	if len(MakeBoolSlicePtr().MapPtr(inverseBoolPtr)) > 0 {
		t.Errorf("MapBoolPtrFilterChain failed.")
	}

	if len(MakeBoolSlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapBoolPtrFilterChain failed.")
	}
}

// TestFilterBoolMethodChain - 
func TestFilterBoolMethodChain(t *testing.T) {
	var vt bool = true

	expectedSumList := []bool{vt}

	newList := MakeBoolSlice([]bool{vt}...).Filter(trueBool)
	if newList[0] != expectedSumList[0] {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(MakeBoolSlice().Filter(nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

// TestFilterBoolPtrMethodChain - 
func TestFilterBoolPtrMethodChain(t *testing.T) {
	var vt bool = true

	expectedSumList := []*bool{&vt}

	newList := MakeBoolSlicePtr([]*bool{&vt}...).FilterPtr(trueBoolPtr)
	if *newList[0] != *expectedSumList[0] {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(MakeBoolSlicePtr(&vt).FilterPtr(nil)) == 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

// TestRemoveBoolMethodChain - 
func TestRemoveBoolMethodChain(t *testing.T) {
	var vt bool = true
	r := func(vt bool) bool {
		if vt == true {
			return true
		}
		return false
	}
	if len(MakeBoolSlice(vt).Remove(r)) > 0 {
		t.Errorf("RemoveBool failed.")
	}
	if len(MakeBoolSlice(vt).Remove(nil)) == 0 {
		t.Errorf("RemoveBool failed.")
	}
}

// TestRemoveBoolPtrMethodChain - 
func TestRemoveBoolPtrMethodChain(t *testing.T) {
	var vt bool = true
	r := func(vt *bool) bool {
		if *vt == true {
			return true
		}
		return false
	}
	if len(MakeBoolSlicePtr(&vt).RemovePtr(r)) > 0 {
		t.Errorf("RemoveBool failed.")
	}
	if len(MakeBoolSlicePtr(&vt).RemovePtr(nil)) == 0 {
		t.Errorf("RemoveBoolPtr failed.")
	}
}

// TestDropWhileBoolMethodChain - 
func TestDropWhileBoolMethodChain(t *testing.T) {
	var vt bool = true
	var vf bool = false

	isTrueBool := func(num bool) bool {
		return num == true
	}

	expectedNewList := []bool{vf, vt}
	NewList := MakeBoolSlice([]bool{vt, vf, vt}...).DropWhile(isTrueBool)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("DropWhileMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeBoolSlice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhile failed.")
	}
}

// TestDropWhileBoolPtrMethodChain - 
func TestDropWhileBoolPtrMethodChain(t *testing.T) {
	var vt bool = true
	var vf bool = false

	isTrueBoolPtr := func(num *bool) bool {
		return *num == true
	}

	expectedNewList := []*bool{&vf, &vt}
	NewList := MakeBoolSlicePtr([]*bool{&vt, &vf, &vt}...).DropWhilePtr(isTrueBoolPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhilePtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeBoolSlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhilePtr failed.")
	}
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

// TestMap2Float32MethodChain - 
func TestMap2Float32MethodChain(t *testing.T) {
	if len(MakeFloat32Slice().Map(nil)) > 0 {
		t.Errorf("MapFloat32 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapFloat32MethodChainPtr - 
func TestMapFloat32MethodChainPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v9 float32 = 9
	expectedSquareList := []*float32{&v1, &v4, &v9}
	squareList := MakeFloat32SlicePtr([]*float32{&v1, &v2, &v3}...).MapPtr(squareFloat32Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapFloat32PtrMthodChain failed")
	}

	if len(MakeFloat32SlicePtr().MapPtr(squareFloat32Ptr)) > 0 {
		t.Errorf("MapFloat32Ptr failed.")
	}
}

// TestMapPtr2Float32MethodChain -  
func TestMapPtr2Float32MethodChain(t *testing.T) {
	if len(MakeFloat32SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapFloat32Ptr failed.")
	}
}

// TestFilterFloat32MethodChain - 
func TestFilterFloat32MethodChain(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	greaterThan1Float32MethodChain := func(num float32) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []float32{v2, v3}
	filteredList := MakeFloat32Slice([]float32{v1, v2, v3}...).Filter(greaterThan1Float32MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterFloat32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Float32MethodChain - 
func TestFilter2Float32MethodChain(t *testing.T) {
	if len(MakeFloat32Slice().Filter(nil)) > 0 {
		t.Errorf("FilterFloat32Ptr failed.")
	}
}

// TestFilterFloat32PtrMethodChain - 
func TestFilterFloat32PtrMethodChain(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3

	greaterThan1Float32MethodChain := func(num *float32) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*float32{&v2, &v3}
	filteredList := MakeFloat32SlicePtr([]*float32{&v1, &v2, &v3}...).FilterPtr(greaterThan1Float32MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterFloat32Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Float32PtrMethodChain - 
func TestFilter2Float32PtrMethodChain(t *testing.T) {
	if len(MakeFloat32SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterFloat32Ptr failed.")
	}
}

// TestRemoveFloat32MethodChain - 
func TestRemoveFloat32MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	isGreaterThanThreeFloat32 := func (num float32) bool {
		return num > 3
	}

	expectedNewList := []float32{v2, v3}
	NewList := MakeFloat32Slice([]float32{v2, v3, v4}...).Remove(isGreaterThanThreeFloat32)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Float32MethodChain - 
func TestRemove2Float32MethodChain(t *testing.T) {
	if len(MakeFloat32Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveFloat32 failed.")
	}
}

// TestRemoveFloat32PtrMethodChain - 
func TestRemoveFloat32PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4

	isGreaterThanThreeFloat32 := func (num *float32) bool {
		return *num > 3
	}

	expectedNewList := []*float32{&v2, &v3}
	NewList := MakeFloat32SlicePtr([]*float32{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeFloat32)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrFloat32MethodChain - 
func TestRemove2PtrFloat32MethodChain(t *testing.T) {
	if len(MakeFloat32SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveFloat32Ptr failed.")
	}
}

func TestDropWhileFloat32MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	isEvenFloat32 := func(num float32) bool {
		return int(num)%2 == 0
	}

	expectedNewList := []float32{v3, v4, v5}
	NewList := MakeFloat32Slice([]float32{v4, v2, v3, v4, v5}...).DropWhile(isEvenFloat32)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Float32MethodChain - 
func TestDropWhile2Float32MethodChain(t *testing.T) {
	if len(MakeFloat32Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileFloat32 failed.")
	}
}

func TestDropWhileFloat32PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	isEvenFloat32Ptr := func(num *float32) bool {
		return int(*num)%2 == 0
	}

	expectedNewList := []*float32{&v3, &v4, &v5}
	NewList := MakeFloat32SlicePtr([]*float32{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenFloat32Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrFloat32MethodChain - 
func TestDropWhile2PtrFloat32MethodChain(t *testing.T) {
	if len(MakeFloat32SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileFloat32Ptr failed.")
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

// TestMap2Float64MethodChain - 
func TestMap2Float64MethodChain(t *testing.T) {
	if len(MakeFloat64Slice().Map(nil)) > 0 {
		t.Errorf("MapFloat64 failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMapFloat64MethodChainPtr - 
func TestMapFloat64MethodChainPtr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v9 float64 = 9
	expectedSquareList := []*float64{&v1, &v4, &v9}
	squareList := MakeFloat64SlicePtr([]*float64{&v1, &v2, &v3}...).MapPtr(squareFloat64Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapFloat64PtrMthodChain failed")
	}

	if len(MakeFloat64SlicePtr().MapPtr(squareFloat64Ptr)) > 0 {
		t.Errorf("MapFloat64Ptr failed.")
	}
}

// TestMapPtr2Float64MethodChain -  
func TestMapPtr2Float64MethodChain(t *testing.T) {
	if len(MakeFloat64SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("MapFloat64Ptr failed.")
	}
}

// TestFilterFloat64MethodChain - 
func TestFilterFloat64MethodChain(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	greaterThan1Float64MethodChain := func(num float64) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []float64{v2, v3}
	filteredList := MakeFloat64Slice([]float64{v1, v2, v3}...).Filter(greaterThan1Float64MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterFloat64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Float64MethodChain - 
func TestFilter2Float64MethodChain(t *testing.T) {
	if len(MakeFloat64Slice().Filter(nil)) > 0 {
		t.Errorf("FilterFloat64Ptr failed.")
	}
}

// TestFilterFloat64PtrMethodChain - 
func TestFilterFloat64PtrMethodChain(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3

	greaterThan1Float64MethodChain := func(num *float64) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*float64{&v2, &v3}
	filteredList := MakeFloat64SlicePtr([]*float64{&v1, &v2, &v3}...).FilterPtr(greaterThan1Float64MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterFloat64Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2Float64PtrMethodChain - 
func TestFilter2Float64PtrMethodChain(t *testing.T) {
	if len(MakeFloat64SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("FilterFloat64Ptr failed.")
	}
}

// TestRemoveFloat64MethodChain - 
func TestRemoveFloat64MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	isGreaterThanThreeFloat64 := func (num float64) bool {
		return num > 3
	}

	expectedNewList := []float64{v2, v3}
	NewList := MakeFloat64Slice([]float64{v2, v3, v4}...).Remove(isGreaterThanThreeFloat64)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Float64MethodChain - 
func TestRemove2Float64MethodChain(t *testing.T) {
	if len(MakeFloat64Slice().Remove(nil)) > 0 {
		t.Errorf("RemoveFloat64 failed.")
	}
}

// TestRemoveFloat64PtrMethodChain - 
func TestRemoveFloat64PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4

	isGreaterThanThreeFloat64 := func (num *float64) bool {
		return *num > 3
	}

	expectedNewList := []*float64{&v2, &v3}
	NewList := MakeFloat64SlicePtr([]*float64{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeFloat64)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2PtrFloat64MethodChain - 
func TestRemove2PtrFloat64MethodChain(t *testing.T) {
	if len(MakeFloat64SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("RemoveFloat64Ptr failed.")
	}
}

func TestDropWhileFloat64MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	isEvenFloat64 := func(num float64) bool {
		return int(num)%2 == 0
	}

	expectedNewList := []float64{v3, v4, v5}
	NewList := MakeFloat64Slice([]float64{v4, v2, v3, v4, v5}...).DropWhile(isEvenFloat64)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Float64MethodChain - 
func TestDropWhile2Float64MethodChain(t *testing.T) {
	if len(MakeFloat64Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhileFloat64 failed.")
	}
}

func TestDropWhileFloat64PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5

	isEvenFloat64Ptr := func(num *float64) bool {
		return int(*num)%2 == 0
	}

	expectedNewList := []*float64{&v3, &v4, &v5}
	NewList := MakeFloat64SlicePtr([]*float64{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenFloat64Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2PtrFloat64MethodChain - 
func TestDropWhile2PtrFloat64MethodChain(t *testing.T) {
	if len(MakeFloat64SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhileFloat64Ptr failed.")
	}
}
