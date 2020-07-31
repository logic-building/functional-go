package basic

// EqualTest is template to generate itself for different combination of data type.
func EqualTest() string {
	return `
func TestEqual<FTYPE>s(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	r := Equal<FTYPE>sP([]<TYPE>{v1, v2, v3}, []<TYPE>{v1, v2, v3})
	if !r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v3, v2}, []<TYPE>{v1, v2, v3})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP(nil, []<TYPE>{v1, v2, v3})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v2, v3}, []<TYPE>{})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP(nil, []<TYPE>{})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v2, v3}, []<TYPE>{v1, v2})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v2, &v3}, []*<TYPE>{&v1, &v2, &v3})
	if !rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v3, &v2}, []*<TYPE>{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr(nil, []*<TYPE>{&v1, &v2, &v3})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v2, &v3}, []*<TYPE>{})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr(nil, []*<TYPE>{})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v2, &v3}, []*<TYPE>{&v1, &v2})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualBoolTest is template to generate itself for different combination of data type.
func EqualBoolTest() string {
	return `
func TestEqual<FTYPE>s(t *testing.T) {
	var v1 <TYPE> = true
	var v2 <TYPE> = false

	r := Equal<FTYPE>sP([]<TYPE>{v1, v2}, []<TYPE>{v1, v2})
	if !r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v1}, []<TYPE>{v1, v2})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP(nil, []<TYPE>{v1, v2})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v2}, []<TYPE>{})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP(nil, []<TYPE>{})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	r = Equal<FTYPE>sP([]<TYPE>{v1, v2}, []<TYPE>{v1})
	if r {
		t.Errorf("Equal<FTYPE>sP failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v2}, []*<TYPE>{&v1, &v2})
	if !rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v1}, []*<TYPE>{&v1, &v2})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr(nil, []*<TYPE>{&v1, &v2})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1, &v2}, []*<TYPE>{})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr(nil, []*<TYPE>{})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}

	rPtr = Equal<FTYPE>sPPtr([]*<TYPE>{&v1}, []*<TYPE>{&v1, &v2})
	if rPtr {
		t.Errorf("Equal<FTYPE>sPPtr failed. Expected=false, actual=true")
	}
}
`
}

// For equalMaps

// EqualMapsTest is template to generate itself for different combination of data type.
func EqualMapsTest() string {
	return `
func TestEqualMaps<FTYPE>s(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3

	var v10 <TYPE> = 10
	var v20 <TYPE> = 20
	var v30 <TYPE> = 30

	r := EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30}, map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v10, v3: v3, v2: v20}, map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(nil, map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30}, map[<TYPE>]<TYPE>{})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(nil, map[<TYPE>]<TYPE>{})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v10, v2: v20, v3: v30}, map[<TYPE>]<TYPE>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v10, &v3: &v3, &v2: &v20}, map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(nil, map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<TYPE>]*<TYPE>{})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(nil, map[*<TYPE>]*<TYPE>{})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<TYPE>]*<TYPE>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsBoolTest is template to generate itself for different combination of data type.
func EqualMapsBoolTest() string {
	return `
func TestEqualMaps<FTYPE>s(t *testing.T) {
	var v1 <TYPE> = true
	var v2 <TYPE> = false

	r := EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v1, v2: v2}, map[<TYPE>]<TYPE>{v1: v1, v2: v2})
	if !r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v1, v1: v1}, map[<TYPE>]<TYPE>{v1: v1, v2: v2})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(nil, map[<TYPE>]<TYPE>{v1: v1, v2: v2})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v1, v2: v2}, map[<TYPE>]<TYPE>{})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(nil, map[<TYPE>]<TYPE>{})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v1, v2: v2}, map[<TYPE>]<TYPE>{v1: v1})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FTYPE>P(map[<TYPE>]<TYPE>{v1: v2}, map[<TYPE>]<TYPE>{v1: v1})
	if r {
		t.Errorf("EqualMap<FTYPE>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2}, map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2})
	if !rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v1, &v1: &v1}, map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(nil, map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2}, map[*<TYPE>]*<TYPE>{})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(nil, map[*<TYPE>]*<TYPE>{})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v1}, map[*<TYPE>]*<TYPE>{&v1: &v1, &v2: &v2})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FTYPE>PPtr(map[*<TYPE>]*<TYPE>{&v1: &v2}, map[*<TYPE>]*<TYPE>{&v1: &v1})
	if rPtr {
		t.Errorf("EqualMap<FTYPE>PPtr failed. Expected=false, actual=true")
	}
}
`
}
