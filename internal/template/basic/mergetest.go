package basic

// MergeTest is template to generate itself for different combination of data type.
func MergeTest() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 10, 2: 20, 3: 30}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 10, 2: 20, 4: 40, 5: 50, 3: 30}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: 40, 5: 50, 3: 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestNumbersToString is template to generate itself for different combination of data type.
func MergeTestNumbersToString() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "10", 2: "20", 3: "30"}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "10", 2: "20", 4: "40", 5: "50", 3: "30"}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: "40", 5: "50", 3: "30"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStringToNumbers is template to generate itself for different combination of data type.
func MergeTestStringToNumbers() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 10, "2": 20, "3": 30}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 10, "2": 20, "4": 40, "5": 50, "3": 30}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": 40, "5": 50, "3": 30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStringToBool is template to generate itself for different combination of data type.
func MergeTestStringToBool() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true, "0": false, "3": true}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true, "0": false, "4": true, "5": true, "3": true}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"4": true, "5": true, "3": true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToString is template to generate itself for different combination of data type.
func MergeTestBoolToString() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "2"}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "2", false: "0"}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "1", false: "0"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestNumberToBool is template to generate itself for different combination of data type.
func MergeTestNumberToBool() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true, 0: false, 3: true}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true, 0: false, 4: true, 5: true, 3: true}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{4: true, 5: true, 3: true}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToNumber is template to generate itself for different combination of data type.
func MergeTestBoolToNumber() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 2}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 2, false: 0}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 1, false: 0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToBool is template to generate itself for different combination of data type.
func MergeTestBoolToBool() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStrToStr is template to generate itself for different combination of data type.
func MergeTestStrToStr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {
	map1 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "two"}
	map2 := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"2": "Two"}

	expected := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	map2 = nil

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	map2 = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}

	expected = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "One", "2": "Two"}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}
