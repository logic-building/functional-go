package basic

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func MinPtrTest() string {
	return `
func TestMin<FTYPE>Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list := []*<TYPE>{&v8, &v2, &v10, &v4}
	min := Min<FTYPE>Ptr(list)
	if *min != 2 {
		t.Errorf("Min<FTYPE>Ptr failed. Expected=10, actual=%v", min)
		t.Errorf(reflect.String.String())
	}

	min = Min<FTYPE>Ptr([]*<TYPE>{})
	if *min != 0 {
		t.Errorf("Min<FTYPE>Ptr failed. Expected=0, actual=%v", *min)
	}

	min = Min<FTYPE>Ptr(nil)
	if *min != 0 {
		t.Errorf("Min<FTYPE>Ptr failed. Expected=0, actual=%v", *min)
	}
}
`
}
