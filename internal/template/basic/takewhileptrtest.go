package basic

// TakeWhilePtrTest<FTYPE>Ptr
func TakeWhilePtrTest() string {
	return `
func TestTakeWhile<FTYPE>Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v40 <TYPE> = 40


	expectedNewList := []*<TYPE>{&v4, &v2, &v4}
	NewList := TakeWhile<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&v40}
	partialIsEvenDivisibleBy := func(num *<TYPE>) bool { return *num % 10 == 0 }
	NewList = TakeWhile<FTYPE>Ptr(partialIsEvenDivisibleBy, []*<TYPE>{&v40})
	
	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}

	if len(TakeWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

`
}

func TakeWhileTestBool() string {
	return `
func TestTakeWhile<FTYPE>Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt <TYPE> = true
	var vf <TYPE> = false


	expectedNewList := []*<TYPE>{&vt, &vt, &vf}
	NewList := TakeWhile<FTYPE>Ptr(func(v *bool) bool {return *v == true}, []*<TYPE>{&vt, &vt, &vf, &vf, &vf})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&vt}
	NewList = TakeWhile<FTYPE>Ptr(func(v *bool) bool {return *v == true}, []*<TYPE>{&vt})
	
	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}

	if len(TakeWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

`
}
