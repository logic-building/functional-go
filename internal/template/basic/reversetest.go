package basic

// ReverseTest template.
func ReverseTest() string {
	return `
func TestReverse<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expected := []<TYPE>{v3, v2, v1}
	reversed := Reverse<FTYPE>s([]<TYPE>{v1, v2, v3})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*<TYPE>{&v3, &v2, &v1}
	reversedPtrList := Reverse<FTYPE>sPtr([]*<TYPE>{&v1, &v2, &v3})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}`
}

// ReverseBoolTest template.
func ReverseBoolTest() string {
	return `
func TestReverse<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expected := []<TYPE>{vt, vt, vf}
	reversed := Reverse<FTYPE>s([]<TYPE>{vf, vt, vt})
	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse<Type>s failed")
	}

	expectedPtrList := []*<TYPE>{&vt, &vt, &vf}
	reversedPtrList := Reverse<FTYPE>sPtr([]*<TYPE>{&vf, &vt, &vt})
	if !reflect.DeepEqual(reversedPtrList, expectedPtrList) {
		t.Errorf("Reverse<Type>s failed")
	}
}`
}
