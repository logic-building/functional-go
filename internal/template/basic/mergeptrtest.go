package basic

// MergeTestPtr is template to generate itself for different combination of data type.
func MergeTestPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4
	var v5 <INPUT_TYPE1> = 5

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30
	var v40 <INPUT_TYPE2> = 40
	var v50 <INPUT_TYPE2> = 50

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestNumbersToStringPtr is template to generate itself for different combination of data type.
func MergeTestNumbersToStringPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4
	var v5 <INPUT_TYPE1> = 5

	var v10 <INPUT_TYPE2> = "10"
	var v20 <INPUT_TYPE2> = "20"
	var v30 <INPUT_TYPE2> = "30"
	var v40 <INPUT_TYPE2> = "40"
	var v50 <INPUT_TYPE2> = "50"
	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStringToNumbersPtr is template to generate itself for different combination of data type.
func MergeTestStringToNumbersPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = "1"
	var v2 <INPUT_TYPE1> = "2"
	var v3 <INPUT_TYPE1> = "3"
	var v4 <INPUT_TYPE1> = "4"
	var v5 <INPUT_TYPE1> = "5"

	var v10 <INPUT_TYPE2> = 10
	var v20 <INPUT_TYPE2> = 20
	var v30 <INPUT_TYPE2> = 30
	var v40 <INPUT_TYPE2> = 40
	var v50 <INPUT_TYPE2> = 50

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v3: &v30}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &v10, &v2: &v20, &v4: &v40, &v5: &v50, &v3: &v30}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1>Ptr<FINPUT_TYPE2> failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &v40, &v5: &v50, &v3: &v30}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStringToBoolPtr is template to generate itself for different combination of data type.
func MergeTestStringToBoolPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v0 <INPUT_TYPE1> = "0"
	var v1 <INPUT_TYPE1> = "1"
	var v3 <INPUT_TYPE1> = "3"
	var v4 <INPUT_TYPE1> = "4"
	var v5 <INPUT_TYPE1> = "5"

	var vt <INPUT_TYPE2> = true
	var vf <INPUT_TYPE2> = false

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vt, &v3: &vt}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if *expected[&v0] != *actual[&v0] {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", *expected[&v0], *actual[&v0])
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToStringPtr is template to generate itself for different combination of data type.
func MergeTestBoolToStringPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {

	var v0 <INPUT_TYPE2> = "0"
	var v1 <INPUT_TYPE2> = "1"
	var v2 <INPUT_TYPE2> = "2"

	var vt <INPUT_TYPE1> = true
	var vf <INPUT_TYPE1> = false

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v2, &vf: &v0}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestNumberToBoolPtr is template to generate itself for different combination of data type.
func MergeTestNumberToBoolPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v0 <INPUT_TYPE1> = 0
	var v1 <INPUT_TYPE1> = 1
	var v3 <INPUT_TYPE1> = 3
	var v4 <INPUT_TYPE1> = 4
	var v5 <INPUT_TYPE1> = 5

	var vt <INPUT_TYPE2> = true
	var vf <INPUT_TYPE2> = false
	

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v0: &vf, &v3: &vt}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vf, &v3: &vt}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vt, &v0: &vf, &v4: &vt, &v5: &vf, &v3: &vt}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v4: &vt, &v5: &vt, &v3: &vt}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToNumberPtr is template to generate itself for different combination of data type.
func MergeTestBoolToNumberPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var vt <INPUT_TYPE1> = true
	var vf <INPUT_TYPE1> = false

	var v1 <INPUT_TYPE2> = 1
	var v0 <INPUT_TYPE2> = 2

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &v1, &vf: &v0}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestBoolToBoolPtr is template to generate itself for different combination of data type.
func MergeTestBoolToBoolPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {

	var vt bool = true
	var vf bool = false

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&vt: &vt, &vf: &vf}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}

// MergeTestStrToStrPtr is template to generate itself for different combination of data type.
func MergeTestStrToStrPtr() string {
	return `
func TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(t *testing.T) {
	var v1 <INPUT_TYPE1> = 1
	var v2 <INPUT_TYPE1> = 2

	var vOne <INPUT_TYPE2> = "One"
	var vTwo <INPUT_TYPE2> = "Two"

	map1 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	map2 := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v2: &vTwo}

	expected := map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	actual := Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = nil
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	map2 = nil

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	map1 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	map2 = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{}

	expected = map[*<INPUT_TYPE1>]*<INPUT_TYPE2>{&v1: &vOne, &v2: &vTwo}
	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(map1, map2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=%v, actual=%v", expected, actual)
	}

	actual = Merge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr(nil, nil)

	if len(actual) != 0 {
		t.Errorf("TestMerge<FINPUT_TYPE1><FINPUT_TYPE2>Ptr failed. Expected=empty mape, actual=%v", actual)
	}
}
`
}
