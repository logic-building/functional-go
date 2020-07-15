package basic

// PMapIONumberPtr is template to generate itself for different combination of data type.
func PMapIONumberPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 <OUTPUT_TYPE> = 2
	var vo3 <OUTPUT_TYPE> = 3
	var vo4 <OUTPUT_TYPE> = 4

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3
	expectedList := []*<OUTPUT_TYPE>{&vo2, &vo3, &vo4}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOStrNumberPtr is template to generate itself for different combination of data type.
func PMapIOStrNumberPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vi10 <INPUT_TYPE> = "10"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIONumberStrPtr is template to generate itself for different combination of data type.
func PMapIONumberStrPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	var vi10 <INPUT_TYPE> = 10

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIONumberBoolPtr is template to generate itself for different combination of data type.
func PMapIONumberBoolPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = 10
	var v0 <INPUT_TYPE> = 0

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOStrBoolPtr is template to generate itself for different combination of data type.
func PMapIOStrBoolPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = "10"
	var v0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOBoolNumberPtr is template to generate itself for different combination of data type.
func PMapIOBoolNumberPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE> = 0

	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOBoolStrPtr is template to generate itself for different combination of data type.
func PMapIOBoolStrPtr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIONumberPtrErr is template to generate itself for different combination of data type.
func PMapIONumberPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 <OUTPUT_TYPE> = 2
	var vo3 <OUTPUT_TYPE> = 3

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3

	expectedList := []*<OUTPUT_TYPE>{&vo2, &vo3}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOStrNumberPtrErr is template to generate itself for different combination of data type.
func PMapIOStrNumberPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIONumberStrPtrErr is template to generate itself for different combination of data type.
func PMapIONumberStrPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE>

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIONumberBoolPtrErr is template to generate itself for different combination of data type.
func PMapIONumberBoolPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = 10
	var v0 <INPUT_TYPE>
	var v3 <INPUT_TYPE> = 3

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v3})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOStrBoolPtrErr is template to generate itself for different combination of data type.
func PMapIOStrBoolPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = "10"
	var v0 <INPUT_TYPE> = "0"
	var v3 <INPUT_TYPE> = "3"

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0, &v3})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// PMapIOBoolNumberPtrErr is template to generate itself for different combination of data type.
func PMapIOBoolNumberPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE> = 0

	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vf})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect
}
`
}

// PMapIOBoolStrPtrErr is template to generate itself for different combination of data type.
func PMapIOBoolStrPtrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vf})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}
