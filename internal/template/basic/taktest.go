package basic

// TakeTest template.
func TakeTest() string {
	return `
func TestTake<FTYPE>(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v0 <TYPE> 

	expected := []<TYPE>{v8, v2, v8}
	list := []<TYPE>{v8, v2, v8, v0, v2, v0}
	actual := Take<FTYPE>(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE>s failed")
	}

	expected = []<TYPE>{}
	actual = Take<FTYPE>(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take <FTYPE> failed")
	}
}

func TestTake<FTYPE>Ptr(t *testing.T) {
	var v8 <TYPE> = 8
	var v2 <TYPE> = 2
	var v0 <TYPE>

	expected := []*<TYPE>{&v8, &v2, &v8}
	list := []*<TYPE>{&v8, &v2, &v8, &v0, &v2, &v0}
	actual := Take<FTYPE>Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE>Ptr failed")
	}

	expected = []*<TYPE>{}
	actual = Take<FTYPE>Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE> failed")
	}
}`
}

// TakeBoolTest removes duplicates.
func TakeBoolTest() string {
	return `
func TestTake<FTYPE>(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expected := []<TYPE>{vt, vf, vt}
	list := []<TYPE>{vt, vf, vt, vt, vt, vt}
	actual := Take<FTYPE>(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE>s failed")
	}

	expected = []<TYPE>{}
	actual = Take<FTYPE>(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take <FTYPE> failed")
	}
}

func TestTake<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expected := []*<TYPE>{&vt, &vf, &vt}
	list := []*<TYPE>{&vt, &vf, &vt, &vt, &vt, &vt}
	actual := Take<FTYPE>Ptr(3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE>Ptr failed")
	}

	expected = []*<TYPE>{}
	actual = Take<FTYPE>Ptr(-3, list)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test Take<FTYPE> failed")
	}
}`
}
