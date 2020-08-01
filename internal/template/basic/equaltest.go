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
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsNumberToStringTest is template to generate itself for different combination of data type.
func EqualMapsNumberToStringTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3

	var v10 <INPUT_TYPE2> = "10"
	var v20 <INPUT_TYPE2> = "20"
	var v30 <INPUT_TYPE2> = "30"

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsStringToNumberTest is template to generate itself for different combination of data type.
func EqualMapsStringToNumberTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsStringToBoolTest is template to generate itself for different combination of data type.
func EqualMapsStringToBoolTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"

	var v10 <INPUT_TYPE2> = true
	var v20 <INPUT_TYPE2> = false
	var v30 <INPUT_TYPE2> = true

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsBoolToStringTest is template to generate itself for different combination of data type.
func EqualMapsBoolToStringTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = true
	var v2 <INPUT_TYPE1> = false

	var v10 <INPUT_TYPE2> = "10"
	var v20 <INPUT_TYPE2> = "20"
	var v30 <INPUT_TYPE2> = "30"

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v2: v10}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsNumberToBoolTest is template to generate itself for different combination of data type.
func EqualMapsNumberToBoolTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3

	var v10 <INPUT_TYPE2> = true
	var v20 <INPUT_TYPE2> = false
	var v30 <INPUT_TYPE2> = true

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsBoolToNumberTest is template to generate itself for different combination of data type.
func EqualMapsBoolToNumberTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = true
	var v2 <INPUT_TYPE1> = false

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v2: v10}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsStringToStringTest is template to generate itself for different combination of data type.
func EqualMapsStringToStringTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v3: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20, v3: v30}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v3: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}

// EqualMapsBoolTest is template to generate itself for different combination of data type.
func EqualMapsBoolTest() string {
	return `
func TestEqualMaps<FINPUT_TYPE1><FINPUT_TYPE2>s(t *testing.T) {
	var v1 <INPUT_TYPE1> = true
	var v2 <INPUT_TYPE1> = false

	var v10 <INPUT_TYPE2> = true
	var v20 <INPUT_TYPE2> = false

	r := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if !r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v2: v10}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20, v2: v20})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(nil, map[<INPUT_TYPE1>]<INPUT_TYPE2>{})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10, v2: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	r = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v20}, map[<INPUT_TYPE1>]<INPUT_TYPE2>{v1: v10})
	if r {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P failed. Expected=false, actual=true")
	}

	// for pointer version of function
	rPtr := EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if !rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(nil, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}

	rPtr = EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v20}, map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v2: &v10})
	if rPtr {
		t.Errorf("EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr failed. Expected=false, actual=true")
	}
}
`
}
