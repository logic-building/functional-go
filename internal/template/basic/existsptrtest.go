package basic

// ExistsPtrTest is template
func ExistsPtrTest() string {
	return `
func TestExists<FTYPE>Ptr(t *testing.T) {
	// Test : value exist in the list

	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10
	var v80 <TYPE> = 80

	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	if !Exists<FTYPE>Ptr(&v8, list1) {
		t.Errorf("Exists<FTYPE> failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v8, &v2, &v10, &v5, &v4}
	if Exists<FTYPE>Ptr(&v80, list2) {
		t.Errorf("Exists<FTYPE> failed. Expected=false, actual=true")
	}

	if Exists<FTYPE>Ptr(&v80, nil) {
		t.Errorf("Exists<FTYPE> failed. Expected=false, actual=true")
	}

	if Exists<FTYPE>Ptr(&v80, []*<TYPE>{}) {
		t.Errorf("Exists<FTYPE> failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}
