package basic

// DifferenceTest is template.
func DifferenceTest() string {
	return `
func TestDifference<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	actual := Difference<FTYPE>()
	if len(actual) != 0 {
		t.Errorf("Difference<FTYPE> failed.")
	}

	expected := []<TYPE>{v8, v2, v1}
	actual = Difference<FTYPE>(expected)

	if len(actual) != 3 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Difference<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8}
	actual = Difference<FTYPE>(list1, list2)
	if len(actual) != 2 || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Difference<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := Difference<FTYPE>Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("Difference<FTYPE>Ptr failed.")
	}

	expectedPtr := []*<TYPE>{&v8, &v2, &v1}
	actualPtr = Difference<FTYPE>Ptr(expectedPtr)

	if len(actualPtr) != 3 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Difference<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8}
	actualPtr = Difference<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !Exists<FTYPE>Ptr(&v2, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Difference<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
`
}

// DifferenceBoolTest is template.
func DifferenceBoolTest() string {
	return `
func TestDifference<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = true
	var v1 <TYPE> = false

	actual := Difference<FTYPE>()
	if len(actual) != 0 {
		t.Errorf("Difference<FTYPE>ram failed.")
	}

	expected := []<TYPE>{v8, v2, v1}
	actual = Difference<FTYPE>(expected)

	if len(actual) != 2 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Difference<FTYPE> failed nks0. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v1}
	actual = Difference<FTYPE>(list1, list2)
	if len(actual) != 1 || !Exists<FTYPE>(v8, actual) {
		t.Errorf("Difference<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := Difference<FTYPE>Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("Difference<FTYPE>Ptr failed.")
	}

	expectedPtr := []*<TYPE>{&v8, &v2, &v1}
	actualPtr = Difference<FTYPE>Ptr(expectedPtr)
	
	if len(actualPtr) != 2 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Difference<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v1}
	actualPtr = Difference<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !Exists<FTYPE>Ptr(&v8, actualPtr) {
		t.Errorf("Difference<FTYPE>PtrNks failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
`
}
