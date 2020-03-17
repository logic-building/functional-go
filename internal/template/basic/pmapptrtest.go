package basic

// PMap<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrTest() string {
	return `
func TestPmap<FTYPE>Ptr(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v9 <TYPE> = 9

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList := PMap<FTYPE>Ptr(square<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>Ptr := func(num *<TYPE>) *<TYPE> {
		r := 5 + *num
		return &r
	}
	sumList := PMap<FTYPE>Ptr(partialAdd<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMap<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMap<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func square<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	r := *num * *num
	return &r
}

`
}

// PMap<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrBoolTest() string {
	return `
func TestPMap<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf}
	
	newList := PMap<FTYPE>Ptr(inverse<FTYPE>Ptr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Map<FTYPE>Ptr failed")
	}

	if len(PMap<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

`
}
