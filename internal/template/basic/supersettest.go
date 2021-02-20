package basic

// SupersetTest is template.
func SupersetTest() string {
	return `
func TestSuperset<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v1 <TYPE> = 1

	actual := Superset<FTYPE>([]<TYPE>{1}, []<TYPE>{})
	if actual == true {
		t.Errorf("Superset<FTYPE> failed")
		t.Errorf(reflect.String.String())
	}

	actual = Superset<FTYPE>([]<TYPE>{v1, v2}, []<TYPE>{v1, v8, v8})
	if actual == true {
		t.Errorf("Superset<FTYPE>Ram failed")
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual = Superset<FTYPE>(list1, list2)
	if !actual {
		t.Errorf("Superset<FTYPE> failed.")
	}

	actual2 := Superset<FTYPE>Ptr([]*<TYPE>{&v1}, []*<TYPE>{})
	if actual2 == true {
		t.Errorf("Superset<FTYPE>Ptr failed")
	}

	actual2 = Superset<FTYPE>Ptr([]*<TYPE>{&v1, &v2}, []*<TYPE>{&v8, &v1})
	if actual2 == true {
		t.Errorf("Superset<FTYPE>Ptr failed")
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = Superset<FTYPE>Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("Superset<FTYPE>Ptr failed.")
	}
}
`
}

// SupersetBoolTest is template.
func SupersetBoolTest() string {
	return `
func TestSuperset<FTYPE>(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = false
	var v1 <TYPE> = true

	actual := Superset<FTYPE>([]<TYPE>{true}, []<TYPE>{})
	if actual == true {
		t.Errorf("Superset<FTYPE> failed")
	}

	list1 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	list2 := []<TYPE>{v8, v2, v1, v1, v2, v8}
	actual = Superset<FTYPE>(list1, list2)
	if !actual {
		t.Errorf("Superset<FTYPE> failed.")
	}

	actual2 := Superset<FTYPE>Ptr([]*<TYPE>{&v1}, []*<TYPE>{})
	if actual2 == true {
		t.Errorf("Superset<FTYPE>Ptr failed")
	}

	list1Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	list2Ptr := []*<TYPE>{&v8, &v2, &v1, &v1, &v2, &v8}
	actual2 = Superset<FTYPE>Ptr(list1Ptr, list2Ptr)
	if !actual2 {
		t.Errorf("Superset<FTYPE>Ptr failed.")
	}
}
`
}
