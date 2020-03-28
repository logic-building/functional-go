package basic

// RemovePtrTest<FTYPE>Ptr
func SomePtrTest() string {
	return `
func TestSome<FTYPE>Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	if !Some<FTYPE>Ptr(isEven<FTYPE>Ptr, list1) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v1, &v3, &v5, &v7}
	if Some<FTYPE>Ptr(isEven<FTYPE>Ptr, list2) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(nil, nil) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{}) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}

`
}

func SomePtrTestBool() string {
	return `
func TestSome<FTYPE>Ptr(t *testing.T) {
	// Test : value exist in the list

	var vt <TYPE> = true
	var vf <TYPE> = false

	list1 := []*<TYPE>{&vt, &vf}
	if !Some<FTYPE>Ptr(func(v *bool) bool { return *v == true }, list1) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	
	if Some<FTYPE>Ptr(nil, nil) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(func(v *bool) bool { return *v == true }, []*<TYPE>{}) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}

`
}
