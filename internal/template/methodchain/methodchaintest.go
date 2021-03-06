package methodchain

import (
	"strings"
)

// MethodChainMapTest is template to generate itself for different combination of data type.
func MethodChainMapTest() string {
	return `
func TestMap<FTYPE>MethodChain(t *testing.T) {
	expectedSquareList := []<TYPE>{1, 4, 9}
	squareList := Make<FTYPE>Slice([]<TYPE>{1, 2, 3}...).Map(square<FTYPE>)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("Map<FTYPE>MthodChain failed")
	}

	if len(Make<FTYPE>Slice().Map(square<FTYPE>)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMap2<FTYPE>MethodChain - 
func TestMap2<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>Slice().Map(nil)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
		t.Errorf(reflect.String.String())
	}
}

// TestMap<FTYPE>MethodChainPtr - 
func TestMap<FTYPE>MethodChainPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v9 <TYPE> = 9
	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList := Make<FTYPE>SlicePtr([]*<TYPE>{&v1, &v2, &v3}...).MapPtr(square<FTYPE>Ptr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("Map<FTYPE>PtrMthodChain failed")
	}

	if len(Make<FTYPE>SlicePtr().MapPtr(square<FTYPE>Ptr)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

// TestMapPtr2<FTYPE>MethodChain -  
func TestMapPtr2<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

// TestFilter<FTYPE>MethodChain - 
func TestFilter<FTYPE>MethodChain(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	greaterThan1<FTYPE>MethodChain := func(num <TYPE>) bool {
		return num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []<TYPE>{v2, v3}
	filteredList := Make<FTYPE>Slice([]<TYPE>{v1, v2, v3}...).Filter(greaterThan1<FTYPE>MethodChain)

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("Filter<FTYPE> failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2<FTYPE>MethodChain - 
func TestFilter2<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>Slice().Filter(nil)) > 0 {
		t.Errorf("Filter<FTYPE>Ptr failed.")
	}
}

// TestFilter<FTYPE>PtrMethodChain - 
func TestFilter<FTYPE>PtrMethodChain(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	greaterThan1<FTYPE>MethodChain := func(num *<TYPE>) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*<TYPE>{&v2, &v3}
	filteredList := Make<FTYPE>SlicePtr([]*<TYPE>{&v1, &v2, &v3}...).FilterPtr(greaterThan1<FTYPE>MethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("Filter<FTYPE>Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

// TestFilter2<FTYPE>PtrMethodChain - 
func TestFilter2<FTYPE>PtrMethodChain(t *testing.T) {
	if len(Make<FTYPE>SlicePtr().FilterPtr(nil)) > 0 {
		t.Errorf("Filter<FTYPE>Ptr failed.")
	}
}

// TestRemove<FTYPE>MethodChain - 
func TestRemove<FTYPE>MethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4

	isGreaterThanThree<FTYPE> := func (num <TYPE>) bool {
		return num > 3
	}

	expectedNewList := []<TYPE>{v2, v3}
	NewList := Make<FTYPE>Slice([]<TYPE>{v2, v3, v4}...).Remove(isGreaterThanThree<FTYPE>)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("Remove<FTYPE> failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2<FTYPE>MethodChain - 
func TestRemove2<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>Slice().Remove(nil)) > 0 {
		t.Errorf("Remove<FTYPE> failed.")
	}
}

// TestRemove<FTYPE>PtrMethodChain - 
func TestRemove<FTYPE>PtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4

	isGreaterThanThree<FTYPE> := func (num *<TYPE>) bool {
		return *num > 3
	}

	expectedNewList := []*<TYPE>{&v2, &v3}
	NewList := Make<FTYPE>SlicePtr([]*<TYPE>{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThree<FTYPE>)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestRemove2Ptr<FTYPE>MethodChain - 
func TestRemove2Ptr<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>SlicePtr().RemovePtr(nil)) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
	}
}

func TestDropWhile<FTYPE>MethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	isEven<FTYPE> := func(num <TYPE>) bool {
		return num%2 == 0
	}

	expectedNewList := []<TYPE>{v3, v4, v5}
	NewList := Make<FTYPE>Slice([]<TYPE>{v4, v2, v3, v4, v5}...).DropWhile(isEven<FTYPE>)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2<FTYPE>MethodChain - 
func TestDropWhile2<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>Slice().DropWhile(nil)) > 0 {
		t.Errorf("DropWhile<FTYPE> failed.")
	}
}

func TestDropWhile<FTYPE>PtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	isEven<FTYPE>Ptr := func(num *<TYPE>) bool {
		return *num%2 == 0
	}

	expectedNewList := []*<TYPE>{&v3, &v4, &v5}
	NewList := Make<FTYPE>SlicePtr([]*<TYPE>{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEven<FTYPE>Ptr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

// TestDropWhile2Ptr<FTYPE>MethodChain - 
func TestDropWhile2Ptr<FTYPE>MethodChain(t *testing.T) {
	if len(Make<FTYPE>SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}
}

// TestReverse<FTYPE>methodchain
func TestReverse<FTYPE>methodchain(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []<TYPE>{v3, v2, v1}
	reversed :=  Make<FTYPE>Slice([]<TYPE>{v1, v2, v3}...).Reverse()
	if expected[0] != reversed[0] || expected[1] != reversed[1] || expected[2] != reversed[2] {
		t.Errorf("Reverse<Type>s failed")
	}
}

// TestReverse<FTYPE>Ptrmethodchain
func TestReverse<FTYPE>Ptrmethodchain(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []*<TYPE>{&v3, &v2, &v1}
	reversed :=  Make<FTYPE>SlicePtr([]*<TYPE>{&v1, &v2, &v3}...).ReversePtr()
	if *expected[0] != *reversed[0] || *expected[1] != *reversed[1] || *expected[2] != *reversed[2] {
		t.Errorf("Reverse<Type>sMethodChain failed")
	}
}

// TestDistinct<FTYPE>MethodChain - 
func TestDistinct<FTYPE>MethodChain(t *testing.T) {
	// Test : Get distinct values
	expected := []<TYPE>{8, 2, 0}
	list := []<TYPE>{8, 2, 8, 0, 2, 0}
	distinct := Make<FTYPE>Slice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []<TYPE>{8, 2, 0}
	list = []<TYPE>{8, 2, 0}
	distinct = Make<FTYPE>Slice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []<TYPE>{}
	list = []<TYPE>{}
	distinct = Make<FTYPE>Slice(list...).Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = Make<FTYPE>Slice().Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}
}

func TestDistinct<FTYPE>PtrMethodChain(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v0 <TYPE>

	// Test : Get distinct values
	expected := []*<TYPE>{&v8, &v2, &v0}
	list := []*<TYPE>{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := Make<FTYPE>SlicePtr(list...).DistinctPtr()
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("Distinct<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*<TYPE>{&v8, &v2, &v0}
	list = []*<TYPE>{&v8, &v2, &v0}
	distinct = Make<FTYPE>SlicePtr(list...).DistinctPtr()
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("Distinct<TYPE>PtrMethodChain failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*<TYPE>{}
	list = []*<TYPE>{}
	distinct = Make<FTYPE>SlicePtr(list...).DistinctPtr()
	if len(distinct) != 0 {
		t.Errorf("Distinct<FTYPE>PtrMethodChain failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = Make<FTYPE>SlicePtr().DistinctPtr()
	if len(distinct) != 0 {
		t.Errorf("Distinct<FTYPE>MethodChain failed. Expected=%v, actual=%v", expected, distinct)
	}
}

// TestTakeWhile<FTYPE>MethodChain - 
func TestTakeWhile<FTYPE>MethodChain(t *testing.T) {
	isEvenMethodChain := func(num <TYPE>) bool {
		return num%2 == 0
	}

	// Test : Take the numbers as long as condition match
	expectedNewList := []<TYPE>{4, 2, 4}
	NewList := Make<FTYPE>Slice([]<TYPE>{4, 2, 4, 7, 5}...).TakeWhile(isEvenMethodChain)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>Slice().TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}

	if len(Make<FTYPE>Slice([]<TYPE>{}...).TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}
}

// TestTakeWhile<FTYPE>MethodChainPtr - 
func TestTakeWhile<FTYPE>MethodChainPtr(t *testing.T) {
	isEvenMethodChain := func(num *<TYPE>) bool {
		return *num%2 == 0
	}

	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7

	// Test : Take the numbers as long as condition match
	expectedNewList := []*<TYPE>{&v4, &v2, &v4}
	NewList := Make<FTYPE>SlicePtr([]*<TYPE>{&v4, &v2, &v4, &v7, &v5}...).TakeWhilePtr(isEvenMethodChain)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhile<FTYPE>methodchain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>SlicePtr().TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE> failed.")
	}

	if len(Make<FTYPE>SlicePtr([]*<TYPE>{}...).TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE> failed.")
	}
}
`
}

// MethodChainSortTest is template to generate itself for different combination of data type.
func MethodChainSortTest() string {
	return `
// TestSort<FTYPE>methodchain
func TestSort<FTYPE>methodchain(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []<TYPE>{v1, v2, v3}
	sorted :=  Make<FTYPE>Slice([]<TYPE>{v3, v2, v1}...).Sort()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSort<FTYPE>methodchainDesc
func TestSort<FTYPE>methodchainDesc(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []<TYPE>{v3, v2, v1}
	sorted :=  Make<FTYPE>Slice([]<TYPE>{v3, v2, v1}...).SortDesc()
	if expected[0] != sorted[0] || expected[1] != sorted[1] || expected[2] != sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
		t.Errorf(reflect.String.String())
	}
}

// TestSort<FTYPE>methodchainPtr
func TestSort<FTYPE>methodchainPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []*<TYPE>{&v1, &v2, &v3}
	sorted :=  Make<FTYPE>SlicePtr([]*<TYPE>{&v3, &v2, &v1}...).SortPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}

// TestSort<FTYPE>methodchainDescPtr
func TestSort<FTYPE>methodchainDescPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []*<TYPE>{&v3, &v2, &v1}
	sorted :=  Make<FTYPE>SlicePtr([]*<TYPE>{&v3, &v2, &v1}...).SortDescPtr()
	if *expected[0] != *sorted[0] || *expected[1] != *sorted[1] || *expected[2] != *sorted[2] {
		t.Errorf("Sort<Type>sMethodChain failed")
	}
}
`
}

// MethodChainMapBoolTest is template to generate itself for different combination of data type.
func MethodChainMapBoolTest() string {
	return `
// TestMap<FTYPE>MethodChain - 
func TestMap<FTYPE>MethodChain(t *testing.T) {
	expectedSquareList := []<TYPE>{false, true, false}
	squareList := Make<FTYPE>Slice([]<TYPE>{true, false, true}...).Map(inverseBool)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("Map<FTYPE>MthodChain failed")
	}

	if len(Make<FTYPE>Slice().Map(inverseBool)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
		t.Errorf(reflect.String.String())
	}

	if len(Make<FTYPE>Slice().Map(nil)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
	}
}

func inverseBool(v bool) bool {
	if v == true {
		return false
	} 
	return true
}

// TestMapPtrMethodChainBool - 
func TestMapPtrMethodChainBool<FTYPE>MethodChain(t *testing.T) {
	tr := true
	f := false
	expectedSquareList := []*<TYPE>{&f, &tr, &f}
	squareList := Make<FTYPE>SlicePtr([]*<TYPE>{&tr, &f, &tr}...).MapPtr(inverseBoolPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("Map<FTYPE>PtrMthodChain failed")
	}

	if len(Make<FTYPE>SlicePtr().MapPtr(inverseBoolPtr)) > 0 {
		t.Errorf("Map<FTYPE>PtrFilterChain failed.")
	}

	if len(Make<FTYPE>SlicePtr().MapPtr(nil)) > 0 {
		t.Errorf("Map<FTYPE>PtrFilterChain failed.")
	}
}

// TestFilterBoolMethodChain - 
func TestFilterBoolMethodChain(t *testing.T) {
	var vt bool = true

	expectedSumList := []bool{vt}

	newList := Make<FTYPE>Slice([]bool{vt}...).Filter(trueBool)
	if newList[0] != expectedSumList[0] {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(Make<FTYPE>Slice().Filter(nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

// TestFilterBoolPtrMethodChain - 
func TestFilterBoolPtrMethodChain(t *testing.T) {
	var vt bool = true

	expectedSumList := []*bool{&vt}

	newList := Make<FTYPE>SlicePtr([]*bool{&vt}...).FilterPtr(trueBoolPtr)
	if *newList[0] != *expectedSumList[0] {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(Make<FTYPE>SlicePtr(&vt).FilterPtr(nil)) == 0 {
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
	if len(Make<FTYPE>Slice(vt).Remove(r)) > 0 {
		t.Errorf("RemoveBool failed.")
	}
	if len(Make<FTYPE>Slice(vt).Remove(nil)) == 0 {
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
	if len(Make<FTYPE>SlicePtr(&vt).RemovePtr(r)) > 0 {
		t.Errorf("RemoveBool failed.")
	}
	if len(Make<FTYPE>SlicePtr(&vt).RemovePtr(nil)) == 0 {
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
	NewList := Make<FTYPE>Slice([]bool{vt, vf, vt}...).DropWhile(isTrueBool)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("DropWhileMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>Slice().DropWhile(nil)) > 0 {
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
	NewList := Make<FTYPE>SlicePtr([]*bool{&vt, &vf, &vt}...).DropWhilePtr(isTrueBoolPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhilePtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>SlicePtr().DropWhilePtr(nil)) > 0 {
		t.Errorf("DropWhilePtr failed.")
	}
}

// TestReverse<FTYPE>methodchain
func TestReverse<FTYPE>methodchain(t *testing.T) {
	var v1 <TYPE> = true
	var v2 <TYPE> = false
	var v3 <TYPE> = false

	expected := []<TYPE>{v3, v2, v1}
	reversed :=  Make<FTYPE>Slice([]<TYPE>{v1, v2, v3}...).Reverse()
	if expected[0] != reversed[0] || expected[1] != reversed[1] || expected[2] != reversed[2] {
		t.Errorf("Reverse<Type>s failed")
	}
}

// TestReverse<FTYPE>Ptrmethodchain
func TestReverse<FTYPE>Ptrmethodchain(t *testing.T) {
	var v1 <TYPE> = true
	var v2 <TYPE> = false
	var v3 <TYPE> = false

	expected := []*<TYPE>{&v3, &v2, &v1}
	reversed :=  Make<FTYPE>SlicePtr([]*<TYPE>{&v1, &v2, &v3}...).ReversePtr()
	if *expected[0] != *reversed[0] || *expected[1] != *reversed[1] || *expected[2] != *reversed[2] {
		t.Errorf("Reverse<Type>sMethodChain failed")
	}
}

// TestDistinctBoolMethodChain
func TestDistinctBoolMethodChain(t *testing.T) {
	var vt bool = true

	newList := Make<FTYPE>Slice([]bool{vt, vt}...).Distinct()
	if newList[0] != vt {
		t.Errorf("DistinctBool failed")
	}

	if len(Make<FTYPE>Slice().Distinct()) > 0 {
		t.Errorf("DistinctBool failed.")
	}
}

// TestDistinctBoolPtrMethodChain
func TestDistinctBoolPtrMethodChain(t *testing.T) {
	var vt bool = true

	newList := Make<FTYPE>SlicePtr([]*bool{&vt, &vt}...).DistinctPtr()
	if *newList[0] != vt {
		t.Errorf("DistinctPtrBool failed")
	}

	if len(Make<FTYPE>SlicePtr().DistinctPtr()) > 0 {
		t.Errorf("DistinctPtrBool failed.")
	}
}

// TestTakeWhileBoolMethodChain - 
func TestTakeWhileBoolMethodChain(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vt, vt, vf}
	NewList := Make<FTYPE>Slice([]bool{vt, vt, vf, vf, vf}...).TakeWhile(func(v bool) bool { return v == true })
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("TakeWhileBool failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []bool{vt}
	NewList = Make<FTYPE>Slice([]bool{vt}...).TakeWhile(func(v bool) bool { return v == true })

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>Slice().TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}

	if len(Make<FTYPE>Slice([]bool{}...).TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}
}

// TestTakeWhileBoolPtrMethodChain - 
func TestTakeWhileBoolPtrMethodChain(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []*bool{&vt, &vt, &vf}
	NewList := Make<FTYPE>SlicePtr([]*bool{&vt, &vt, &vf, &vf, &vf}...).TakeWhilePtr(func(v *bool) bool { return *v == true } )
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*bool{&vt}
	NewList = Make<FTYPE>SlicePtr([]*bool{&vt}...).TakeWhilePtr(func(v *bool) bool { return *v == true })

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Make<FTYPE>SlicePtr().TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}

	if len(Make<FTYPE>SlicePtr([]*bool{}...).TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// ReplaceActivityMethodChainMap - Replace activity for string type
func ReplaceActivityMethodChainMap(code string) string {
	t1 := `func TestMapStrMethodChain(t *testing.T) {
	expectedSquareList := []string{1, 4, 9}
	squareList := MakeStrSlice([]string{1, 2, 3}...).Map(squareStr)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapStrMthodChain failed")
	}

	if len(MakeStrSlice().Map(squareStr)) > 0 {
		t.Errorf("MapStr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	t2 := `func TestMapStrMethodChain(t *testing.T) {
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
}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `greaterThan1StrMethodChain := func(num string) bool {
		return num > 1
	}`
	t2 = `greaterThan1StrMethodChain := func(num string) bool {
		return num > "1"
	}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestMapStrMethodChainPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v9 string = "9"
	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList := MakeStrSlicePtr([]*string{&v1, &v2, &v3}...).MapPtr(squareStrPtr)

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("MapStrPtrMthodChain failed")
	}

	if len(MakeStrSlicePtr().MapPtr(squareStrPtr)) > 0 {
		t.Errorf("MapStrPtr failed.")
	}
}`
	t2 = `func TestMapStrMethodChainPtr(t *testing.T) {
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
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestFilterStrPtrMethodChain(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"

	greaterThan1StrMethodChain := func(num *string) bool {
		return *num > 1
	}

	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v3}
	filteredList := MakeStrSlicePtr([]*string{&v1, &v2, &v3}...).FilterPtr(greaterThan1StrMethodChain)

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}`
	t2 = `func TestFilterStrPtrMethodChain(t *testing.T) {
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
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestRemoveStrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	isGreaterThanThreeStr := func (num string) bool {
		return num > 3
	}

	expectedNewList := []string{v2, v3}
	NewList := MakeStrSlice([]string{v2, v3, v4}...).Remove(isGreaterThanThreeStr)

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}`
	t2 = `func TestRemoveStrMethodChain(t *testing.T) {
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
}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestRemoveStrPtrMethodChain(t *testing.T) {
	// Test : even number in the list
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	isGreaterThanThreeStr := func (num *string) bool {
		return *num > 3
	}

	expectedNewList := []*string{&v2, &v3}
	NewList := MakeStrSlicePtr([]*string{&v2, &v3, &v4}...).RemovePtr(isGreaterThanThreeStr)

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}`

	t2 = `func TestRemoveStrPtrMethodChain(t *testing.T) {
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
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenFloat32 := func(num float32) bool {
		return num%2 == 0
	}`
	t2 = `isEvenFloat32 := func(num float32) bool {
		return int(num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenFloat64 := func(num float64) bool {
		return num%2 == 0
	}`
	t2 = `isEvenFloat64 := func(num float64) bool {
		return int(num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestDropWhileStrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	isEvenStr := func(num string) bool {
		return num%2 == 0
	}

	expectedNewList := []string{v3, v4, v5}
	NewList := MakeStrSlice([]string{v4, v2, v3, v4, v5}...).DropWhile(isEvenStr)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile>MethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}`
	t2 = `func TestDropWhileStrMethodChain(t *testing.T) {
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
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenFloat32Ptr := func(num *float32) bool {
		return *num%2 == 0
	}`
	t2 = `isEvenFloat32Ptr := func(num *float32) bool {
		return int(*num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenFloat64Ptr := func(num *float64) bool {
		return *num%2 == 0
	}`
	t2 = `isEvenFloat64Ptr := func(num *float64) bool {
		return int(*num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestDropWhileStrPtrMethodChain(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	isEvenStrPtr := func(num *string) bool {
		return *num%2 == 0
	}

	expectedNewList := []*string{&v3, &v4, &v5}
	NewList := MakeStrSlicePtr([]*string{&v4, &v2, &v3, &v4, &v5}...).DropWhilePtr(isEvenStrPtr)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile>PtrMethodChain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}`
	t2 = `func TestDropWhileStrPtrMethodChain(t *testing.T) {
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
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `// TestDistinctStrMethodChain - 
func TestDistinctStrMethodChain(t *testing.T) {
	// Test : Get distinct values
	expected := []string{8, 2, 0}
	list := []string{8, 2, 8, 0, 2, 0}
	distinct := MakeStrSlice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{8, 2, 0}
	list = []string{8, 2, 0}
	distinct = MakeStrSlice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != 8 || distinct[1] != 2 || distinct[2] != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{}
	list = []string{}
	distinct = MakeStrSlice(list...).Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = MakeStrSlice().Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}
}`

	t2 = `// TestDistinctStrMethodChain -
func TestDistinctStrMethodChain(t *testing.T) {
	// Test : Get distinct values
	expected := []string{"8", "2", "0"}
	list := []string{"8", "2", "8", "0", "2", "0"}
	distinct := MakeStrSlice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != "8" || distinct[1] != "2" || distinct[2] != "0" {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{"8", "2", "0"}
	list = []string{"8", "2", "0"}
	distinct = MakeStrSlice(list...).Distinct()
	if len(distinct) != 3 || distinct[0] != "8" || distinct[1] != "2" || distinct[2] != "0" {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []string{}
	list = []string{}
	distinct = MakeStrSlice(list...).Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = MakeStrSlice().Distinct()
	if len(distinct) != 0 {
		t.Errorf("DistinctInt failed. Expected=%v, actual=%v", expected, distinct)
	}
}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `func TestTakeWhileStrMethodChain(t *testing.T) {
	isEvenMethodChain := func(num string) bool {
		return num%2 == 0
	}

	// Test : Take the numbers as long as condition match
	expectedNewList := []string{4, 2, 4}
	NewList := MakeStrSlice([]string{4, 2, 4, 7, 5}...).TakeWhile(isEvenMethodChain)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeStrSlice().TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}

	if len(MakeStrSlice([]string{}...).TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}
}`
	t2 = `func TestTakeWhileStrMethodChain(t *testing.T) {
	isEvenMethodChain := func(num string) bool {
		return num == "2" || num == "4"
	}

	// Test : Take the numbers as long as condition match
	expectedNewList := []string{"4", "2", "4"}
	NewList := MakeStrSlice([]string{"4", "2", "4", "7", "5"}...).TakeWhile(isEvenMethodChain)
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeStrSlice().TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}

	if len(MakeStrSlice([]string{}...).TakeWhile(nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}
}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenMethodChain := func(num float32) bool {
		return num%2 == 0
	}`
	t2 = `isEvenMethodChain := func(num float32) bool {
		return int(num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenMethodChain := func(num float64) bool {
		return num%2 == 0
	}`
	t2 = `isEvenMethodChain := func(num float64) bool {
		return int(num)%2 == 0
	}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `
// TestTakeWhileStrMethodChainPtr - 
func TestTakeWhileStrMethodChainPtr(t *testing.T) {
	isEvenMethodChain := func(num *string) bool {
		return *num%2 == 0
	}

	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"

	// Test : Take the numbers as long as condition match
	expectedNewList := []*string{&v4, &v2, &v4}
	NewList := MakeStrSlicePtr([]*string{&v4, &v2, &v4, &v7, &v5}...).TakeWhilePtr(isEvenMethodChain)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrmethodchain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeStrSlicePtr().TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}

	if len(MakeStrSlicePtr([]*string{}...).TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}
}`
	t2 = `// TestTakeWhileStrMethodChainPtr - 
func TestTakeWhileStrMethodChainPtr(t *testing.T) {
	isEvenMethodChain := func(num *string) bool {
		return *num == "2" || *num == "4"
	}

	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"

	// Test : Take the numbers as long as condition match
	expectedNewList := []*string{&v4, &v2, &v4}
	NewList := MakeStrSlicePtr([]*string{&v4, &v2, &v4, &v7, &v5}...).TakeWhilePtr(isEvenMethodChain)
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrmethodchain failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(MakeStrSlicePtr().TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}

	if len(MakeStrSlicePtr([]*string{}...).TakeWhilePtr(nil)) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}
}`
	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenMethodChain := func(num *float32) bool {
		return *num%2 == 0
	}`

	t2 = `isEvenMethodChain := func(num *float32) bool {
		return int(*num)%2 == 0
	}`

	code = strings.ReplaceAll(code, t1, t2)

	t1 = `isEvenMethodChain := func(num *float64) bool {
		return *num%2 == 0
	}`

	t2 = `isEvenMethodChain := func(num *float64) bool {
		return int(*num)%2 == 0
	}`

	code = strings.ReplaceAll(code, t1, t2)

	return code
}
