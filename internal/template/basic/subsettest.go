package basic

// SubsetTest is template.
func SubsetTest() string {
	return `
func TestSubset<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	actual := Subset<FTYPE>([]<TYPE>{1}, []<TYPE>{})
	if actual == true {
		t.Errorf("Subset<FTYPE> failed")
		t.Errorf(reflect.String.String())
	}

	actual = Subset<FTYPE>([]<TYPE>{v1, v2}, []<TYPE>{v1, v8, v8})
	if actual == true {
		t.Errorf("Subset<FTYPE>Ram failed")
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual = Subset<FTYPE>(list1, list2)
	if !actual {
		t.Errorf("Subset<FTYPE> failed.")
	}

	actual2 := Subset<FTYPE>Ptr([]*<TYPE>{&v1}, []*<TYPE>{})
	if actual2 == true {
		t.Errorf("Subset<FTYPE>Ptr failed")
	}

	actual2 = Subset<FTYPE>Ptr([]*<TYPE>{&v1, &v2}, []*<TYPE>{&v8, &v1})
	if actual2 == true {
		t.Errorf("Subset<FTYPE>Ptr failed")
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = Subset<FTYPE>Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("Subset<FTYPE>Ptr failed.")
	}
}
`
}

// SubsetBoolTest is template.
func SubsetBoolTest() string {
	return `
func TestSubset<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = false
	var v1 <TYPE> = true

	actual := Subset<FTYPE>([]<TYPE>{true}, []<TYPE>{})
	if actual == true {
		t.Errorf("Subset<FTYPE> failed")
		t.Errorf(reflect.String.String())
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual = Subset<FTYPE>(list1, list2)
	if !actual {
		t.Errorf("Subset<FTYPE> failed.")
	}

	actual2 := Subset<FTYPE>Ptr([]*<TYPE>{&v1}, []*<TYPE>{})
	if actual2 == true {
		t.Errorf("Subset<FTYPE>Ptr failed")
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = Subset<FTYPE>Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("Subset<FTYPE>Ptr failed.")
	}
}
`
}
