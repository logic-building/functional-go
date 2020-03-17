package basic

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func MinMaxPtrTest() string {
	return `
func TestMinMax<FTYPE>Ptr(t *testing.T) {
	// Test : get min number in the list
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list := []*<TYPE>{&v8, &v2, &v10, &v4}
	min, max := MinMax<FTYPE>Ptr(list)
	if *min != 2 || *max != 10 {
		t.Errorf("MinMax<FTYPE>Ptr failed. Expected=2,10, actual=%v, %v", *min, *max)
		t.Errorf(reflect.String.String())
	}

	min, max = MinMax<FTYPE>Ptr([]*<TYPE>{})
	if *min != 0 || *max != 0{
		t.Errorf("MinMax<FTYPE>Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}

	min, max = MinMax<FTYPE>Ptr(nil)
	if *min != 0 || *max != 0{
		t.Errorf("MinMax<FTYPE>Ptr failed. Expected=0, actual=%v, %v", *min, *max)
	}
}
`
}
