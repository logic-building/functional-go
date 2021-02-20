package basic

// DropLastTest is template to generate itself for different combination of data type.
func DropLastTest() string {
	return `
func TestDropLast<FTYPE>(t *testing.T) {
	list := []<TYPE>{1, 2, 3, 4, 5}
	expectedList := []<TYPE>{1, 2, 3, 4}
	actualList := DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{1, 2}
	expectedList = []<TYPE>{1}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{1}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// DropLastBoolTest is template to generate itself for different combination of data type.
func DropLastBoolTest() string {
	return `
func TestDropLast<FTYPE>(t *testing.T) {
	list := []<TYPE>{true, true, true, true, false}
	expectedList := []<TYPE>{true, true, true, true}
	actualList := DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{true, true}
	expectedList = []<TYPE>{true}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{true}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// DropLastPtrTest is template to generate itself for different combination of data type.
func DropLastPtrTest() string {
	return `
func TestDropLast<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	list := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*<TYPE>{&v1, &v2, &v3, &v4}
	actualList := DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{&v1, &v2}
	expectedList = []*<TYPE>{&v1}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{&v1}
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{}
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// DropLastPtrBoolTest is template to generate itself for different combination of data type.
func DropLastPtrBoolTest() string {
	return `
func TestDropLast<FTYPE>Ptr(t *testing.T) {
	var true bool = true
	var false bool = false
	list := []*<TYPE>{&true, &true, &true, &true, &false}
	expectedList := []*<TYPE>{&true, &true, &true, &true}
	actualList := DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{&true, &true}
	expectedList = []*<TYPE>{&true}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{&true}
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []*<TYPE>{}
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []*<TYPE>{}
	actualList = DropLast<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE>Ptr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}
