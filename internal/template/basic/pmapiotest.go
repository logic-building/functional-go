package basic

// PMapIONumber is template to generate itself for different combination of data type.
func PMapIONumber() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []<OUTPUT_TYPE>{2, 3, 4}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func plusOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	return <OUTPUT_TYPE>(num + 1)
}
`
}

// PMapIOStrNumber is template to generate itself for different combination of data type.
func PMapIOStrNumber() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num == "ten" {
		return <OUTPUT_TYPE>(10)
	}
	return 0
}
`
}

// PMapIONumberStr is template to generate itself for different combination of data type.
func PMapIONumberStr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num == 10 {
		return <OUTPUT_TYPE>("10")
	}
	return "0"
}
`
}

// PMapIONumberBool is template to generate itself for different combination of data type.
func PMapIONumberBool() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num > 0 {
		return true
	}
	return false
}
`
}

// PMapIOStrBool is template to generate itself for different combination of data type.
func PMapIOStrBool() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num != "0" {
		return true
	}
	return false
}
`
}

// PMapIOBoolNumber is template to generate itself for different combination of data type.
func PMapIOBoolNumber() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 0}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num == true {
		return 10
	}
	return 0
}
`
}

// PMapIOBoolStr is template to generate itself for different combination of data type.
func PMapIOBoolStr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "0"}
	newList := PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(PMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num == true {
		return "10"
	}
	return "0"
}
`
}
