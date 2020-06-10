package basic

// UnionTest is template.
func UnionTest() string {
	return `
func TestUnion<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	expected := []<TYPE>{v8, v2, v1}
	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual := Union<FTYPE>(list1, list2)
	if len(actual) != 3 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Union<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := Union<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 3 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Union<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
`
}

// UnionBoolTest is template.
func UnionBoolTest() string {
	return `
func TestUnion<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = false
	var v0 <TYPE> = true

	expected := []<TYPE>{v8, v2}
	list1 := []<TYPE>{v8, v2, v0, v0, v2, v8}
	list2 := []<TYPE>{v8, v2, v0, v0, v2, v8}
	actual := Union<FTYPE>(list1, list2)
	if len(actual) != 2 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) {
		t.Errorf("Union<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v0, &v0, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v0, &v0, &v2, &v8}
	actualPtr := Union<FTYPE>Ptr(list1Ptr, list2Ptr)
	if len(actualPtr) != 2 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) {
		t.Errorf("Union<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
`
}
