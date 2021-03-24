package basic

// RestPtrTest is template
func RestPtrTest() string {
	return `
func TestRest<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	

	list := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*<TYPE>{&v2, &v3, &v4, &v5}
	newList := Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*<TYPE>{&v1}
	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*<TYPE>{}
	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

`
}

// RestPtrTestBool is template
func RestPtrTestBool() string {
	return `
func TestRest<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false
	

	list := []*<TYPE>{&vt, &vf, &vf, &vf, &vt}
	expectedList := []*<TYPE>{&vf, &vf, &vf, &vt}
	newList := Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*<TYPE>{&vt}
	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*<TYPE>{}
	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*<TYPE>{}
	newList = Rest<FTYPE>Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("Rest<FTYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

`
}
