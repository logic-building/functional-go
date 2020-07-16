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
