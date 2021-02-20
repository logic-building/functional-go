package basic

// DistinctPtrTest is template
func DistinctPtrTest() string {
	return `
func TestDistinct<FTYPE>Ptr(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v0 <TYPE>

	// Test : Get distinct values
	expected := []*<TYPE>{&v8, &v2, &v0}
	list := []*<TYPE>{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := Distinct<FTYPE>Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("Distinct<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*<TYPE>{&v8, &v2, &v0}
	list = []*<TYPE>{&v8, &v2, &v0}
	distinct = Distinct<FTYPE>Ptr(list)
	if len(distinct) != 3 || *distinct[0] != v8 || *distinct[1] != v2 || *distinct[2] != v0 {
		t.Errorf("Distinct<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	expected = []*<TYPE>{}
	list = []*<TYPE>{}
	distinct = Distinct<FTYPE>Ptr(list)
	if len(distinct) != 0 {
		t.Errorf("Distinct<FTYPE>Ptr failed. Expected=%v, actual=%v", expected, distinct)
	}

	distinct = Distinct<FTYPE>Ptr(nil)
	if len(distinct) != 0 {
		t.Errorf("Distinct<TYPE> failed. Expected=%v, actual=%v", expected, distinct)
		t.Errorf(reflect.String.String())
	}
}
`
}

// DistinctPtrBoolTest is template
func DistinctPtrBoolTest() string {
	return `
func TestDistinct<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true

	newList := Distinct<FTYPE>Ptr([]*<TYPE>{&vt, &vt})
	if *newList[0] != vt {
		t.Errorf("Distinct<FTYPE>Ptr failed")
	}

	if len(Distinct<FTYPE>Ptr(nil)) > 0 {
		t.Errorf("Distinct<FTYPE>Ptr failed.")
	}
}
`
}

// DistinctP

// DistinctPTest is template
func DistinctPTest() string {
	return `
func TestDistinct<FTYPE>P(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v0 <TYPE>

	list0 := []<TYPE>{v8, v2, v8, v0, v2, v0}
	result := Distinct<FTYPE>P(list0)
	if result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=false, actual=true")
	}

	list0 = []<TYPE>{v8, v2, v0}
	result = Distinct<FTYPE>P(list0)
	if !result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=true, actual=false")
	}

	list0 = []<TYPE>{}
	result = Distinct<FTYPE>P(list0)
	if result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=false, actual=true")
	}

	result = Distinct<FTYPE>P(nil)
	if result {
		t.Errorf("Distinct<TYPE>P failed. Expected=false, actual=true")
	}

	list := []*<TYPE>{&v8, &v2, &v8, &v0, &v2, &v0}
	distinct := Distinct<FTYPE>PPtr(list)
	if distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	list = []*<TYPE>{&v8, &v2, &v0}
	distinct = Distinct<FTYPE>PPtr(list)
	if !distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=true, actual=false")
	}

	list = []*<TYPE>{}
	distinct = Distinct<FTYPE>PPtr(list)
	if distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	distinct = Distinct<FTYPE>PPtr(nil)
	if distinct {
		t.Errorf("Distinct<TYPE>PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}

// DistinctPBoolTest is template
func DistinctPBoolTest() string {
	return `
func TestDistinct<FTYPE>P(t *testing.T) {
	var v8 <TYPE> = true
	var v2 <TYPE> = false

	list0 := []<TYPE>{v8, v2, v8, v2}
	result := Distinct<FTYPE>P(list0)
	if result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=false, actual=true")
	}

	list0 = []<TYPE>{v8, v2}
	result = Distinct<FTYPE>P(list0)
	if !result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=true, actual=false")
	}

	list0 = []<TYPE>{}
	result = Distinct<FTYPE>P(list0)
	if result {
		t.Errorf("Distinct<FTYPE>P failed. Expected=false, actual=true")
	}

	result = Distinct<FTYPE>P(nil)
	if result {
		t.Errorf("Distinct<TYPE>P failed. Expected=false, actual=true")
	}

	list := []*<TYPE>{&v8, &v2, &v8, &v2}
	distinct := Distinct<FTYPE>PPtr(list)
	if distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	list = []*<TYPE>{&v8, &v2}
	distinct = Distinct<FTYPE>PPtr(list)
	if !distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=true, actual=false")
	}

	list = []*<TYPE>{}
	distinct = Distinct<FTYPE>PPtr(list)
	if distinct {
		t.Errorf("Distinct<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	distinct = Distinct<FTYPE>PPtr(nil)
	if distinct {
		t.Errorf("Distinct<TYPE>PPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}
