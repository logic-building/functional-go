package basic

// Every<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func EveryPtrTest() string {
	return `
func TestEvery<FTYPE>Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10
	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	if !Every<FTYPE>Ptr(isEven<FTYPE>Ptr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v8, &v2, &v10, &v5, &v4}
	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, list2) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, nil) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(nil, []*<TYPE>{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}

// Every<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func EveryPtrBoolTest() string {
	return `
func TestEvery<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	// Test : every value in the list is either true or false
	list1 := []*<TYPE>{&vt, &vt, &vt, &vt}
	if !Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&vt, &vt, &vt, &vf}
	if Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	
	list1 = []*bool{}
	if Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(TruePtr, nil) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(nil, []*bool{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}
	if EveryBoolPtr(TruePtr, []*bool{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}

func TruePtr(val *bool) bool {
	return *val == true
}
`
}
