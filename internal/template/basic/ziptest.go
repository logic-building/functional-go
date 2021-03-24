package basic

// ZIPNumberToNumber is template to generate itself for different combination of data type.
func ZIPNumberToNumber() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 := []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 10, 2: 20, 3: 30, 4: 40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{10, 20, 30}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 10, 2: 20, 3: 30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3}
	list2 = []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: 10, 2: 20, 3: 30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPNumberToStr is template to generate itself for different combination of data type.
func ZIPNumberToStr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 := []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "10", 2: "20", 3: "30", 4: "40"}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{"10", "20", "30"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "10", 2: "20", 3: "30"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3}
	list2 = []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: "10", 2: "20", 3: "30"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToNumber is template to generate itself for different combination of data type.
func ZIPStrToNumber() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 := []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 10, "2": 20, "3": 30, "4": 40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{10, 20, 30}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 10, "2": 20, "3": 30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3"}
	list2 = []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": 10, "2": 20, "3": 30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{10, 20, 30, 40}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToStr is template to generate itself for different combination of data type.
func ZIPStrToStr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 := []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "10", "2": "20", "3": "30", "4": "40"}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{"10", "20", "30"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "10", "2": "20", "3": "30"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3"}
	list2 = []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": "10", "2": "20", "3": "30"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{"10", "20", "30", "40"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToBool is template to generate itself for different combination of data type.
func ZIPStrToBool() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 := []<INPUT_TYPE2>{true, true, false, true}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true, "2": true, "3": false, "4": true}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{true, true, false}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true, "2": true, "3": false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3"}
	list2 = []<INPUT_TYPE2>{true, true, false, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{"1": true, "2": true, "3": false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{true, true, true, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{"1", "2", "3", "4"}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToStr is template to generate itself for different combination of data type.
func ZIPBoolToStr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{true, true, false, true}
	list2 := []<INPUT_TYPE2>{"1", "2", "3", "4"}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "4", false: "3"}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{"1", "2", "3"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "2", false: "3"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false}
	list2 = []<INPUT_TYPE2>{"1", "2", "3", "4"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: "2", false: "3"}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{"1", "2", "3", "4"}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPNumberToBool is template to generate itself for different combination of data type.
func ZIPNumberToBool() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 := []<INPUT_TYPE2>{true, true, false, true}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true, 2: true, 3: false, 4: true}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{true, true, false}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true, 2: true, 3: false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3}
	list2 = []<INPUT_TYPE2>{true, true, false, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{1: true, 2: true, 3: false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{true, true, true, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{1, 2, 3, 4}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToNumber is template to generate itself for different combination of data type.
func ZIPBoolToNumber() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{true, true, false, true}
	list2 := []<INPUT_TYPE2>{1, 2, 3, 4}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 4, false: 3}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{1, 2, 3}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 2, false: 3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false}
	list2 = []<INPUT_TYPE2>{1, 2, 3, 4}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: 2, false: 3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{1, 2, 3, 4}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToBool is template to generate itself for different combination of data type.
func ZIPBoolToBool() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>(t *testing.T) {

	list1 := []<INPUT_TYPE1>{true, true, false, true}
	list2 := []<INPUT_TYPE2>{true, true, false, true}

	expectedMap := map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{true, true, false}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false}
	list2 = []<INPUT_TYPE2>{true, true, false, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{true: true, false: false}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{true, true, false, true}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{true, true, false, true}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = []<INPUT_TYPE2>{}

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}
