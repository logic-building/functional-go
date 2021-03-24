package basic

// DropPtrTest returns a new list after dropping the given item
func DropPtrTest() string {
	return `
func TestDrop<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	expectedList := []*<TYPE>{&v2, &v3}
	newList := Drop<FTYPE>Ptr(&v1, []*<TYPE>{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("Drop<FTYPE>Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = Drop<FTYPE>Ptr(&v1, []*<TYPE>{})

	if len(newList) != 0 {
		t.Errorf("Drop<FTYPE>Ptr failed. Expected list=%v, actual list=%v", []*<TYPE>{}, newList)
	}

	newList = Drop<FTYPE>Ptr(&v1, nil)
	if len(newList) != 0 {
		t.Errorf("Drop<FTYPE>Ptr failed. Expected list=%v, actual list=%v", []*<TYPE>{}, newList)
		t.Errorf(reflect.String.String())
	}
}

func TestDrop<FTYPE>sPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4

	// Test : Drop number from the list
	expectedList := []*<TYPE>{&v2, &v3}
	newList := Drop<FTYPE>sPtr([]*<TYPE>{&v1, &v4}, []*<TYPE>{&v1, &v2, &v3, &v1, &v4})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("Drop<FTYPE>sPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}

	newList = Drop<FTYPE>sPtr([]*<TYPE>{&v1, &v4}, []*<TYPE>{})

	if len(newList) != 0 {
		t.Errorf("Drop<FTYPE>sPtr failed. Expected list=%v, actual list=%v", []*<TYPE>{}, newList)
	}

	newList = Drop<FTYPE>sPtr([]*<TYPE>{&v1, &v4}, nil)
	if len(newList) != 0 {
		t.Errorf("Drop<FTYPE>sPtr failed. Expected list=%v, actual list=%v", []*<TYPE>{}, newList)
	}

	newList = Drop<FTYPE>sPtr(nil, nil)
	if len(newList) != 0 {
		t.Errorf("Drop<FTYPE>sPtr failed. Expected list=%v, actual list=%v", []*<TYPE>{}, newList)
	}

	newList = Drop<FTYPE>sPtr(nil, []*<TYPE>{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("Drop<FTYPE>s failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}
`
}

// DropPtrBoolTest is template
func DropPtrBoolTest() string {
	return `
func TestDrop<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = false
	var v2 <TYPE> = true
	var v3 <TYPE> = true

	expectedList := []*<TYPE>{&v2, &v3}
	newList := Drop<FTYPE>Ptr(&v1, []*<TYPE>{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("Drop<FTYPE>Ptr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

func TestDrop<FTYPE>sPtr(t *testing.T) {
	var v1 <TYPE> = false
	var v2 <TYPE> = true
	var v3 <TYPE> = true

	expectedList := []*<TYPE>{&v2, &v3}
	newList := Drop<FTYPE>sPtr([]*<TYPE>{&v1}, []*<TYPE>{&v1, &v2, &v3, &v1})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("Drop<FTYPE>sPtr failed. Expected list=%v, actual list=%v", expectedList, newList)
	}
}

`
}
