package basic

// FilterMap<FTYPE>Ptr
func FilterMapPtrTest() string {
	return `
func TestFilterMap<FTYPE>Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 <TYPE> = 1
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v0 <TYPE> = 0
	var v2 <TYPE> = 2

	expectedFilteredList := []*<TYPE>{&v2, &v4, &v8}
	filteredList := FilterMap<FTYPE>Ptr(isPositive<FTYPE>Ptr, multiplyBy2<FTYPE>Ptr, []*<TYPE>{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2]{
		t.Errorf("FilterMap<FTYPE>Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMap<FTYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositive<FTYPE>Ptr(num *<TYPE>) bool {
	return *num > 0
}
func multiplyBy2<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	result := *num * 2
	return &result
}

`
}

// FilterMapPtrBool
func FilterMapPtrBoolTest() string {
	return `
func TestFilterMap<FTYPE>Ptr(t *testing.T) {
	var vt1 <TYPE> = true
	var vt2 <TYPE> = true
	var vf1 <TYPE> = false
	var vf2 <TYPE> = false

	expectedFilteredList := []*<TYPE>{&vt1, &vt2}
	filteredList := FilterMap<FTYPE>Ptr(isTrue<FTYPE>Ptr, returnSame<FTYPE>Ptr, []*<TYPE>{&vt1, &vt2, &vf1, &vf2})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMap<FTYPE>Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMap<FTYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isTrue<FTYPE>Ptr(num *<TYPE>) bool {
	return *num == true
}
func returnSame<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	return num
}
`
}
