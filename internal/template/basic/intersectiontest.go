package basic

// IntersectionTest is template.
func IntersectionTest() string {
	return `
func TestIntersection<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	actual := Intersection<FTYPE>()
	if len(actual) != 0 {
		t.Errorf("Intersection<FTYPE> failed.")
	}

	expected := []<TYPE>{v8, v2, v1}
	actual = Intersection<FTYPE>(expected)

	if len(actual) != 3 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Intersection<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8}
	actual = Intersection<FTYPE>(list1, list2)
	if len(actual) != 1 || !Exists<FTYPE>(v8, actual) {
		t.Errorf("Intersection<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := Intersection<FTYPE>Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("Intersection<FTYPE>Ptr failed.")
	}

	expectedPtr := []*<TYPE>{&v8, &v2, &v1}
	actualPtr = Intersection<FTYPE>Ptr(expectedPtr)

	if len(actualPtr) != 3 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Intersection<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8}
	actualPtr = Intersection<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !Exists<FTYPE>Ptr(&v8, actualPtr) {
		t.Errorf("Intersection<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
`
}

// IntersectionBoolTest is template.
func IntersectionBoolTest() string {
	return `
func TestIntersection<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = true
	var v1 <TYPE> = false

	actual := Intersection<FTYPE>()
	if len(actual) != 0 {
		t.Errorf("Intersection<FTYPE> failed.")
	}

	expected := []<TYPE>{v8, v2, v1}
	actual = Intersection<FTYPE>(expected)

	if len(actual) != 2 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Intersection<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8}
	actual = Intersection<FTYPE>(list1, list2)
	if len(actual) != 1 || !Exists<FTYPE>(v8, actual) {
		t.Errorf("Intersection<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
	}

	actualPtr := Intersection<FTYPE>Ptr()
	if len(actualPtr) != 0 {
		t.Errorf("Intersection<FTYPE>Ptr failed.")
	}

	expectedPtr := []*<TYPE>{&v8, &v2, &v1}
	actualPtr = Intersection<FTYPE>Ptr(expectedPtr)

	if len(actualPtr) != 2 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Intersection<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8}
	actualPtr = Intersection<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 1 || !Exists<FTYPE>Ptr(&v8, actualPtr) {
		t.Errorf("Intersection<FTYPE>Ptr failed. Expected=%v, actual=%v", expectedPtr, actualPtr)
	}
}
`
}
