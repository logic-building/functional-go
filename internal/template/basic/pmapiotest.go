package basic

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
	} else {
		return 0
	}
}
`
}

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
	} else {
		return "0"
	}
}
`
}

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
	} else {
		return false
	}
}
`
}

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
	} else {
		return false
	}
}
`
}

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
	} else {
		return 0
	}
}
`
}

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
	} else {
		return "0"
	}
}
`
}
