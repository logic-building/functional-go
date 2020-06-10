package basic

// KeysNumberToNumberTest is template
func KeysNumberToNumberTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 1}
	expectedList := []<INPUT_TYPE1>{1}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysNumberToStrTest is template
func KeysNumberToStrTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "1"}
	expectedList := []<INPUT_TYPE1>{1}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysStrToNumberTest is template
func KeysStrToNumberTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 1}
	expectedList := []<INPUT_TYPE1>{"1"}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysStrToBoolTest is template
func KeysStrToBoolTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true}
	expectedList := []<INPUT_TYPE1>{"1"}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysBoolToStrTest is template
func KeysBoolToStrTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1"}
	expectedList := []<INPUT_TYPE1>{true}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysNumberToBoolTest is template
func KeysNumberToBoolTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true}
	expectedList := []<INPUT_TYPE1>{1}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysBoolToNumberTest is template
func KeysBoolToNumberTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1}
	expectedList := []<INPUT_TYPE1>{true}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysBoolToBoolTest is template
func KeysBoolToBoolTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true}
	expectedList := []<INPUT_TYPE1>{true}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// KeysStrToStrTest is template
func KeysStrToStrTest() string {
	return `
func TestKeys<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	m := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"ram": "ram"}
	expectedList := []<INPUT_TYPE1>{"ram"}
	actualList := Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("Test Keys<FINPUT_TYPE1><FINPUT_TYPE2> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}
