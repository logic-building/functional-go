package methodchain

import "strings"

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
`
}

// MethodChainMapBoolTest is template to generate itself for different combination of data type.
func MethodChainMapBoolTest() string {
	return `
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
}

func inverseBool(v bool) bool {
	if v == true {
		return false
	} 
	return true
}

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
}

func TestFilterBoolMethodChain(t *testing.T) {
	var vt bool = true

	expectedSumList := []bool{vt}

	newList := Make<FTYPE>Slice([]bool{vt}...).Filter(trueBool)
	if newList[0] != expectedSumList[0] {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(FilterBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

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

func TestRemoveBoolPtrMethodChain(t *testing.T) {
	var vt bool = true
	if len(Make<FTYPE>Slice(vt).Remove(nil)) == 0 {
		t.Errorf("RemoveBool failed.")
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
	return code
}
