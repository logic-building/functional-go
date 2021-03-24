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

	expectedList = []<OUTPUT_TYPE>{2, 3, 4}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
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

	expectedList = []<OUTPUT_TYPE>{10}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
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

	expectedList = []<OUTPUT_TYPE>{"10"}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
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

	expectedList = []<OUTPUT_TYPE>{true, false}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
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

	expectedList = []<OUTPUT_TYPE>{true, false}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"10", "0"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
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

	expectedList = []<OUTPUT_TYPE>{10, 0}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
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

	expectedList = []<OUTPUT_TYPE>{"10", "0"}
	newList = PMap<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>(num <INPUT_TYPE>) <OUTPUT_TYPE> {
	if num == true {
		return "10"
	}
	return "0"
}
`
}

// PMapIONumberErr is template to generate itself for different combination of data type.
func PMapIONumberErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []<OUTPUT_TYPE>{2, 3}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 3, 2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 3, 2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 1, 3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{2, 3}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 3, 1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIOStrNumberErr is template to generate itself for different combination of data type.
func PMapIOStrNumberErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "0"})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0", "0", "ten"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0", "0", "ten"}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0", "ten", "0"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "ten", "0"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{10, 10}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "ten"}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "ten", "0"}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "ten", "0"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0", "0", "ten"}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0", "ten", "0"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten", "ten", "0"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIONumberStrErr is template to generate itself for different combination of data type.
func PMapIONumberStrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 0})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{0, 0, 10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{0, 10, 0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{"10", "10"}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{0, 0, 10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{0, 10, 0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIONumberBoolErr is template to generate itself for different combination of data type.
func PMapIONumberBoolErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 3, 3})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 3, 10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 10, 3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{true, true}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 3, 10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{3, 10, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 10, 3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIOStrBoolErr is template to generate itself for different combination of data type.
func PMapIOStrBoolErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "0", "3"})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10", "3"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"3", "3", "10"}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"3", "10", "3"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10", "3"}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{true, true}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10"}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10", "3"}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10", "3"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"3", "3", "10"}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"3", "10", "3"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "10", "3"}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIOBoolNumberErr is template to generate itself for different combination of data type.
func PMapIOBoolNumberErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 10}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, false})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, false, true}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{10, 10}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, false, true}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}

// PMapIOBoolStrErr is template to generate itself for different combination of data type.
func PMapIOBoolStrErr() string {
	return `
func TestPmap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "10"}
	newList, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, false})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, false, true}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	expectedList = []<OUTPUT_TYPE>{"10", "10"}
	newList, _ = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true}, Optional{RandomOrder: true})
	
	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if expectedList[i] == newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, false, true}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{false, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true, false}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
`
}
