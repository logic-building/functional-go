package basic

// FilterMapIONumber is template to generate itself for different combination of data type.
func FilterMapIONumber() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : some logic
	expectedList := []<OUTPUT_TYPE>{3, 4}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num != 1
}
`
}

// FilterMapIOStrNumber is template to generate itself for different combination of data type.
func FilterMapIOStrNumber() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"one", "ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num != "one"
}
`
}

// FilterMapIONumberStr is template to generate itself for different combination of data type.
func FilterMapIONumberStr() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num != 1
}
`
}

// FilterMapIONumberBool is template to generate itself for different combination of data type.
func FilterMapIONumberBool() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num != 1
}
`
}

// FilterMapIOStrBool is template to generate itself for different combination of data type.
func FilterMapIOStrBool() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"1", "10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num != "1"
}
`
}

// FilterMapIOBoolNumber is template to generate itself for different combination of data type.
func FilterMapIOBoolNumber() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num == true
}
`
}

// FilterMapIOBoolStr is template to generate itself for different combination of data type.
func FilterMapIOBoolStr() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "10"}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(notOne<FINPUT_TYPE><FOUTPUT_TYPE>, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) bool {
	return num == true
}
`
}
