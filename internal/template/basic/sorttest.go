package basic

// SortTest is template
func SortTest() string {
	return `
func TestSort<FTYPE>(t *testing.T) {
	expectedList := []<TYPE>{1, 2, 3, 4, 5}
	sortedList := Sort<FTYPE>s([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = Sort<FTYPE>s([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("Sotred<TYPE>s failed")
	}
}

func TestSort<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	sortedList := Sort<FTYPE>sPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotred<TYPE>s failed")
		}
	}

	sortedList = Sort<FTYPE>sPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("Sotred<TYPE>sPtr failed")
	}
}

func TestSort<FTYPE>Desc(t *testing.T) {
	expectedList := []<TYPE>{5, 4, 3, 2, 1}
	sortedList := Sort<FTYPE>sDesc([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = Sort<FTYPE>sDesc([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("Sotred<TYPE>sDesc failed")
	}
}

func TestSort<FTYPE>DescPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v5, &v4, &v3, &v2, &v1}
	sortedList := Sort<FTYPE>sDescPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotred<TYPE>s failed")
		}
	}

	sortedList = Sort<FTYPE>sDescPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("Sotred<TYPE>sDescPtr failed")
	}
}
`
}

// SortIntsTest is template
func SortIntsTest() string {
	return `
func TestSort<FTYPE>(t *testing.T) {
	expectedList := []<TYPE>{1, 2, 3, 4, 5}
	sortedList := SortInts([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = SortInts([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortInts failed")
	}
}

func TestSort<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortIntsPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})
	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}
	sortedList = SortIntsPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsPtr failed")
	}
}

func TestSort<FTYPE>Desc(t *testing.T) {
	expectedList := []<TYPE>{5, 4, 3, 2, 1}
	sortedList := SortIntsDesc([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}
	sortedList = SortIntsDesc([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsDesc failed")
	}
}

func TestSort<FTYPE>DescPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortIntsDescPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
		}
	}

	sortedList = SortIntsDescPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortIntsDescPtr failed")
	}
}
`
}

// SortFloats64Test is template
func SortFloats64Test() string {
	return `
func TestSort<FTYPE>(t *testing.T) {
	expectedList := []<TYPE>{1, 2, 3, 4, 5}
	sortedList := SortFloats64([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = SortFloats64([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64 failed")
	}
}

func TestSort<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortFloats64Ptr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}

	sortedList = SortFloats64Ptr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64Ptr failed")
	}
}

func TestSort<FTYPE>Desc(t *testing.T) {
	expectedList := []<TYPE>{5, 4, 3, 2, 1}
	sortedList := SortFloats64Desc([]<TYPE>{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = SortFloats64Desc([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64Desc failed")
	}
}

func TestSort<FTYPE>DescPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortFloats64DescPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
		}
	}

	sortedList = SortFloats64DescPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortFloats64DescPtr failed")
	}
}
`
}

// SortStrsTest is template
func SortStrsTest() string {
	return `
func TestSort<FTYPE>(t *testing.T) {
	expectedList := []<TYPE>{"1", "2", "3", "4", "5"}
	sortedList := SortStrs([]<TYPE>{"5", "1", "4", "2", "3"})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = SortStrs([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrs failed")
	}
}

func TestSort<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortStrsPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}

	sortedList = SortStrsPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsPtr failed")
	}
}

func TestSort<FTYPE>Desc(t *testing.T) {
	expectedList := []<TYPE>{"5", "4", "3", "2", "1"}
	sortedList := SortStrsDesc([]<TYPE>{"5", "1", "4", "2", "3"})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotred<TYPE>s failed")
	}

	sortedList = SortStrsDesc([]<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsDesc failed")
	}
}

func TestSort<FTYPE>DescPtr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedList := []*<TYPE>{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortStrsDescPtr([]*<TYPE>{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			t.Errorf("Sotred<TYPE>s failed")
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
		}
	}
	sortedList = SortStrsDescPtr([]*<TYPE>{})
	if len(sortedList) > 0 {
		t.Errorf("SortStrsDescPtr failed")
	}
}
`
}
