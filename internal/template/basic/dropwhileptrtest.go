package basic

// Drop returns a new list after dropping the given item
func DropWhilePtrTest() string {
	return `
func TestDropWhile<FTYPE>Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedNewList := []*<TYPE>{&v3, &v4, &v5}
	NewList := DropWhile<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	if len(DropWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

`
}

// DistinctBoolPtr removes duplicates.
func DropWhilePtrBoolTest() string {
	return `
func TestDropWhile<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vf, &vt}
	NewList := DropWhile<FTYPE>Ptr(isTrue<FTYPE>Ptr, []*<TYPE>{&vt, &vf, &vt})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

`
}
