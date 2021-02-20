package basic

import "strings"

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

//**********************************Err***********************************

// FilterMapIONumberErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : some logic
	var v1 <INPUT_TYPE> = 1
	var v2 <INPUT_TYPE> = 2
	var v3 <INPUT_TYPE> = 3
	var v4 <INPUT_TYPE> = 4
	var v5 <INPUT_TYPE> = 5

	var vo5 <OUTPUT_TYPE> = 5
	var vo6 <OUTPUT_TYPE> = 6

	expectedList := []<OUTPUT_TYPE>{vo5, vo6}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{v1, v4, v5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})

	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{v2, v3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{v3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return num != 1, nil
}
func plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (<OUTPUT_TYPE>, error) {
	if num == 3 {
		return 0, errors.New("3 is not valid number for this test")
	}
	c := <OUTPUT_TYPE>(num + 1)
	return c, nil
}
`
}

// FilterMapIOStrNumberErrTest is template to generate itself for different combination of data type.
func FilterMapIOStrNumberErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vOne <INPUT_TYPE> = "one"
	var vTwo <INPUT_TYPE> = "two"
	var vThree <INPUT_TYPE> = "three"
	var vTen <INPUT_TYPE> = "ten"

	expectedList := []<OUTPUT_TYPE>{vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, []<INPUT_TYPE>{vOne, vTen})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, []<INPUT_TYPE>{vTwo, vThree})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrErr, []<INPUT_TYPE>{vThree})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrErr(num <INPUT_TYPE>) (bool, error) {
	if num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return num != "one", nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrErr(num string) (<OUTPUT_TYPE>, error) {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if num == "three" {
		return 0, errors.New("three is not valid value for this test")
	}
	if num == "ten" {
		r = <OUTPUT_TYPE>(10)
		return r, nil
	}
	return r, nil
}
`
}

// FilterMapIONumberStrErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberStrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	var ov10 <OUTPUT_TYPE> = "10"
	var iv1 <INPUT_TYPE> = 1
	var iv2 <INPUT_TYPE> = 2
	var iv3 <INPUT_TYPE> = 3
	var iv10 <INPUT_TYPE> = 10
	expectedList := []<OUTPUT_TYPE>{ov10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, []<INPUT_TYPE>{iv1, iv10})

	if newList[0] != expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, []<INPUT_TYPE>{iv1, iv2, iv10})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumErr, []<INPUT_TYPE>{iv1, iv3, iv10})

	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumErr(num <INPUT_TYPE>) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumErr(num <INPUT_TYPE>) (<OUTPUT_TYPE>, error) {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if num == 3 {
		return "0", errors.New("3 is not valid number for this test")
	}
	if num == 10 {
		r = "10"
		return r, nil
	}
	return r, nil
}
`
}

// FilterMapIONumberBoolErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberBoolErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE>

	expectedList := []<OUTPUT_TYPE>{vto, vfo}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi2, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi3, vi10, vi0})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return num != 1, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == 3 {
		return false, errors.New("3 is not valid number for this test")
	}
	r := num > 0
	return r, nil
}
`
}

// FilterMapIOStrBoolErrTest is template to generate itself for different combination of data type.
func FilterMapIOStrBoolErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic

	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = "1"
	var vi2 <INPUT_TYPE> = "2"
	var vi3 <INPUT_TYPE> = "3"
	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []<OUTPUT_TYPE>{vto, vfo}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi1, vi10, vi0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi1, vi10, vi2})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vi1, vi10, vi3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == "2" {
		return false, errors.New("2 is not valid value for this test")
	}
	return num != "1", nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == "3" {
		return false, errors.New("3 is not valid value for this test")
	}
	var t bool = true
	var f bool = false

	if num == "10" {
		return t, nil
	}
	return f, nil
}
`
}

// FilterMapIOBoolNumberErrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolNumberErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic

	var vo10 <OUTPUT_TYPE> = 10

	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []<OUTPUT_TYPE>{vo10, vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err2, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err2(num <INPUT_TYPE>) (bool, error) {
	if num == false {
		return true, nil
	}
	return true, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num bool) (<OUTPUT_TYPE>, error) {
	if num == false {
		return 0, errors.New("false is error for this test")
	}
	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE>

	if num == true {
		return v10, nil
	}
	return v0, nil
}
`
}

// FilterMapIOBoolStrErrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolStrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"

	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []<OUTPUT_TYPE>{vo10, vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err2, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{vit, vit, vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return num == true, nil
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Err2(num <INPUT_TYPE>) (bool, error) {
	if num == false {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num bool) (<OUTPUT_TYPE>, error) {
	if num == false {
		return "", errors.New("false is error in this test")
	}
	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	if num == true {
		return v10, nil
	}
	return v0, nil
}
`
}

// ReplaceActivityFilterMapIOErr replaces ...
func ReplaceActivityFilterMapIOErr(code string) string {
	s1 := `_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntInt64Err(t *testing.T) {`
	s2 := `"errors"
	"testing"
)

func TestFilterMapIntInt64Err(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
