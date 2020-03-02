package basic

// FilterMapIONumber is template to generate itself for different combination of data type.
func FilterMapIONumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : some logic
	var v1 <INPUT_TYPE> = 1
	var v2 <INPUT_TYPE> = 2
	var v3 <INPUT_TYPE> = 3

	var vo3 <OUTPUT_TYPE> = 3
	var vo4 <OUTPUT_TYPE> = 4

	expectedList := []*<OUTPUT_TYPE>{&vo3, &vo4}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != 1
}
func plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *<OUTPUT_TYPE> {
	c := <OUTPUT_TYPE>(*num + 1)
	return &c
}
`
}

// FilterMapIOStrNumberPtr is template to generate itself for different combination of data type.
func FilterMapIOStrNumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vOne <INPUT_TYPE> = "one"
	var vTen <INPUT_TYPE> = "ten"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr, []*<INPUT_TYPE>{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr(num *<INPUT_TYPE>) bool {
	return *num != "one"
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr(num *string) *<OUTPUT_TYPE> {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == "ten" {
		r = <OUTPUT_TYPE>(10)
		return &r
	}
	return &r
}
`
}

// FilterMapIONumberStrPtr is template to generate itself for different combination of data type.
func FilterMapIONumberStrPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var ov10 <OUTPUT_TYPE> = "10"
	var iv1 <INPUT_TYPE> = 1
	var iv10 <INPUT_TYPE> = 10
	expectedList := []*<OUTPUT_TYPE>{&ov10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr, []*<INPUT_TYPE>{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr(num *<INPUT_TYPE>) bool {
	return *num != 1
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr(num *<INPUT_TYPE>) *<OUTPUT_TYPE> {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}
`
}

// FilterMapIONumberBool is template to generate itself for different combination of data type.
func FilterMapIONumberBoolPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = 1
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE> = 0

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != 1
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *bool {
	r := *num > 0
	return &r
}
`
}

// FilterMapIOStrBoolPtr is template to generate itself for different combination of data type.
func FilterMapIOStrBoolPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic

	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = "1"
	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != "1"
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *bool {
	var t bool = true
	var f bool = false

	if *num == "10" {
		return &t
	} 
	return &f
}
`
}

// FilterMapIOBoolNumberPtr is template to generate itself for different combination of data type.
func FilterMapIOBoolNumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 <OUTPUT_TYPE> = 10
	
	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num == true
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *bool) *<OUTPUT_TYPE> {
	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE> = 0

	if *num == true {
		return &v10
	} else {
		return &v0
	}
}
`
}

// FilterMapIOBoolStrPtr is template to generate itself for different combination of data type.
func FilterMapIOBoolStrPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	
	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num == true
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *bool) *<OUTPUT_TYPE> {
	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	if *num == true {
		return &v10
	} else {
		return &v0
	}
}
`
}
