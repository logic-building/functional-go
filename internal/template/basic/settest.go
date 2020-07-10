package basic

// SetTest is template.
func SetTest() string {
	return `
func TestSet<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	expected := []<TYPE>{v8, v2, v1}
	list := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual := Set<FTYPE>(list)
	if len(actual) != 3 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Union<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = Set<FTYPE>(nil)
	if len(actual) != 0 {
		t.Errorf("Set<FTYPE> failed")
	}

	actualPtr2 := Set<FTYPE>([]<TYPE>{})
	if len(actualPtr2) != 0 {
		t.Errorf("Set<FTYPE> failed")
	}

	actualPtr3 := Set<FTYPE>Ptr([]*<TYPE>{})
	if len(actualPtr3) != 0 {
		t.Errorf("Set<FTYPE>Ptr failed")
	}

	listPtr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := Set<FTYPE>Ptr(listPtr)
	if len(actualPtr) != 3 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) || !Exists<FTYPE>Ptr(&v1, actualPtr) {
		t.Errorf("Set<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
`
}

// SetBoolTest is template.
func SetBoolTest() string {
	return `
func TestSet<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = false
	var v1 <TYPE> = true

	expected := []<TYPE>{v8, v2}
	list := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual := Set<FTYPE>(list)
	if len(actual) != 2 || !Exists<FTYPE>(v8, actual) || !Exists<FTYPE>(v2, actual) || !Exists<FTYPE>(v1, actual) {
		t.Errorf("Set<FTYPE> failed. Expected=%v, actual=%v", expected, actual)
		t.Errorf(reflect.String.String())
	}

	actual = Set<FTYPE>(nil)
	if len(actual) != 0 {
		t.Errorf("Set<FTYPE> failed")
	}

	actualPtr2 := Set<FTYPE>([]<TYPE>{})
	if len(actualPtr2) != 0 {
		t.Errorf("Set<FTYPE> failed")
	}

	actualPtr3 := Set<FTYPE>Ptr([]*<TYPE>{})
	if len(actualPtr3) != 0 {
		t.Errorf("Set<FTYPE> failed")
	}

	listPtr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actualPtr := Set<FTYPE>Ptr(listPtr)
	if len(actualPtr) != 2 || !Exists<FTYPE>Ptr(&v8, actualPtr) || !Exists<FTYPE>Ptr(&v2, actualPtr) {
		t.Errorf("Set<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}
}
`
}
