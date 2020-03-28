package basic

// RemovePtrTest<FTYPE>Ptr
func RemovePtrTest() string {
	return `
func TestRemove<FTYPE>Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40

	expectedNewList := []*<TYPE>{&v1, &v3}
	NewList := Remove<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE> failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*<TYPE>{&v1, &v3}
	partialIsEven := func(num *<TYPE>) bool { 
		return *num % 10 == 0 
	}
	NewList = Remove<FTYPE>Ptr(partialIsEven, []*<TYPE>{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Remove<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

`
}

func RemovePtrTestBool() string {
	return `
func TestRemove<FTYPE>Ptr(t *testing.T) {
	// Test : even number in the list
	var vt <TYPE> = true
	var vf <TYPE> = false
	
	expectedNewList := []*<TYPE>{&vt}
	NewList := Remove<FTYPE>Ptr(func(v *bool) bool { return *v == false} , []*<TYPE>{&vt, &vf, &vf})

	if *NewList[0] != *expectedNewList[0]  {
		t.Errorf("Remove<FTYPE> failed. Expected New list=%v, actual list=%v", *expectedNewList[0], *NewList[0])
	}

	if len(Remove<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

`
}
