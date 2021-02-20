package basic

// ZIPNumberToNumberPtr is template to generate itself for different combination of data type.
func ZIPNumberToNumberPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30
	var v40 <INPUT_TYPE2> = 40

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPNumberToStrPtr is template to generate itself for different combination of data type.
func ZIPNumberToStrPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4

	var v10 <INPUT_TYPE2> = "10"
	var v20 <INPUT_TYPE2> = "20"
	var v30 <INPUT_TYPE2> = "30"
	var v40 <INPUT_TYPE2> = "40"

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToNumberPtr is template to generate itself for different combination of data type.
func ZIPStrToNumberPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"
	var v4 <INPUT_TYPE1> = "4"

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30
	var v40 <INPUT_TYPE2> = 40

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToStrPtr is template to generate itself for different combination of data type.
func ZIPStrToStrPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"
	var v4 <INPUT_TYPE1> = "4"

	var v10 <INPUT_TYPE2> = "10"
	var v20 <INPUT_TYPE2> = "20"
	var v30 <INPUT_TYPE2> = "30"
	var v40 <INPUT_TYPE2> = "40"

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30, &v4: &v40}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v10, &v20, &v30, &v40}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPStrToBoolPtr is template to generate itself for different combination of data type.
func ZIPStrToBoolPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"
	var v4 <INPUT_TYPE1> = "4"

	var vt <INPUT_TYPE2> = true
	var vf <INPUT_TYPE2> = false

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vt, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToStrPtr is template to generate itself for different combination of data type.
func ZIPBoolToStrPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE2> = "1"
	var v2 <INPUT_TYPE2> = "2"
	var v3 <INPUT_TYPE2> = "3"
	var v4 <INPUT_TYPE2> = "4"

	var vt <INPUT_TYPE1> = true
	var vf <INPUT_TYPE1> = false

	list1 := []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 := []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v4, &vf: &v3}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2, &vf: &v3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2, &vf: &v3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPNumberToBoolPtr is template to generate itself for different combination of data type.
func ZIPNumberToBoolPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4

	var vt <INPUT_TYPE2> = true
	var vf <INPUT_TYPE2> = false

	list1 := []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 := []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf, &v4: &vt}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v2: &vt, &v3: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vt, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&v1, &v2, &v3, &v4}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToNumberPtr is template to generate itself for different combination of data type.
func ZIPBoolToNumberPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE2> = 1
	var v2 <INPUT_TYPE2> = 2
	var v3 <INPUT_TYPE2> = 3
	var v4 <INPUT_TYPE2> = 4

	var vt <INPUT_TYPE1> = true
	var vf <INPUT_TYPE1> = false

	list1 := []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 := []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v4, &vf: &v3}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2, &vf: &v3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2, &vf: &v3}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&v1, &v2, &v3, &v4}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}

// ZIPBoolToBoolPtr is template to generate itself for different combination of data type.
func ZIPBoolToBoolPtr() string {
	return `
func TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	list1 := []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 := []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actualMap := Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{&vt, &vt, &vf, &vt}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{&vt, &vt, &vf, &vt}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = []*<INPUT_TYPE2>{}

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = nil
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}

	list1 = []*<INPUT_TYPE1>{}
	list2 = nil

	expectedMap = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	actualMap = Zip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(list1, list2)

	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("TestZip<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expectedMap, actualMap)
	}
}
`
}
