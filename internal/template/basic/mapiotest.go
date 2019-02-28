package basic

// MapIONumber is template to generate itself for different combination of data type.
func MapIONumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []<OUTPUT_TYPE>{2, 3, 4}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrNumber is template to generate itself for different combination of data type.
func MapIOStrNumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberStr is template to generate itself for different combination of data type.
func MapIONumberStr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberBool is template to generate itself for different combination of data type.
func MapIONumberBool() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrBool is template to generate itself for different combination of data type.
func MapIOStrBool() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"10", "0"})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolNumber is template to generate itself for different combination of data type.
func MapIOBoolNumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 0}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolStr is template to generate itself for different combination of data type.
func MapIOBoolStr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "0"}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}
