package basic

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func FilterPtrTest() string {
	return `
func TestFilter<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v10 <TYPE> = 10
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40


	// Test : even number in the list
	expectedFilteredList := []*<TYPE>{&v2, &v4}
	filteredList := Filter<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*<TYPE>{&v20, &v40}
	partialIsEven := func(num *<TYPE>) bool { return isEvenDivisibleBy<FTYPE>Ptr(num, &v10) }
	filteredList = Filter<FTYPE>Ptr(partialIsEven, []*<TYPE>{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(Filter<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEven<FTYPE>Ptr(num *<TYPE>) bool {
	return *num%2 == 0
}

func isEvenDivisibleBy<FTYPE>Ptr(num, divisibleBy *<TYPE>) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}

`
}

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func FilterPtrBoolTest() string {
	return `
func TestFilter<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true

	expectedSumList := []*<TYPE>{&vt}
	
	newList := Filter<FTYPE>Ptr(true<FTYPE>Ptr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Filter<FTYPE>Ptr failed")
	}

	if len(Map<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

func true<FTYPE>Ptr(num1 *<TYPE>) bool {
	return true
}

`
}
