package basic

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
}`
}
